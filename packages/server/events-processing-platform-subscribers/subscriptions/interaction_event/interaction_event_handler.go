package interactionEvent

import (
	"context"
	"encoding/json"
	"fmt"
	aiConfig "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-ai/config"
	ai "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-ai/service"
	commonEntity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/postgres/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/subscriptions"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/interaction_event/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/interaction_event/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	interactioneventpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/interaction_event"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

func NewInteractionEventHandler(repositories *repository.Repositories, log logger.Logger, cfg *config.Config, grpcClients *grpc_client.Clients) *interactionEventHandler {
	aiCfg := aiConfig.Config{
		OpenAi: aiConfig.AiModelConfigOpenAi{},
		Anthropic: aiConfig.AiModelConfigAnthropic{
			ApiPath: cfg.Services.Anthropic.ApiPath,
			ApiKey:  cfg.Services.Anthropic.ApiKey,
		},
	}
	return &interactionEventHandler{
		repositories: repositories,
		log:          log,
		cfg:          cfg,
		aiModel:      ai.NewAiModel(ai.AnthropicModelType, aiCfg),
		grpcClients:  grpcClients,
	}
}

type interactionEventHandler struct {
	repositories *repository.Repositories
	log          logger.Logger
	cfg          *config.Config
	aiModel      ai.AiModel
	grpcClients  *grpc_client.Clients
}

func (h *interactionEventHandler) GenerateSummaryForEmail(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventHandler.GenerateSummaryForEmail")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData event.InteractionEventRequestSummaryEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	interactionEventId := aggregate.GetInteractionEventObjectID(evt.AggregateID, eventData.Tenant)
	span.LogFields(log.String("interactionEventId", interactionEventId))

	interactionEvent, err := h.repositories.Neo4jRepositories.InteractionEventReadRepository.GetInteractionEvent(ctx, eventData.Tenant, interactionEventId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error getting interaction event with id %s: %v", interactionEvent, err)
		return nil
	}

	interactionEventChannel := utils.GetStringPropOrEmpty(interactionEvent.Props, "channel")
	interactionEventContent := utils.GetStringPropOrEmpty(interactionEvent.Props, "content")

	if interactionEventChannel != "EMAIL" {
		tracing.TraceErr(span, errors.New("interaction event is not an email"))
		h.log.Warnf("Interaction event with id %s is not an email, skipping", interactionEventId)
		return nil
	}
	if interactionEventContent == "" {
		tracing.TraceErr(span, errors.New("interaction event content is empty"))
		h.log.Warnf("Interaction event with id %s has no content, skipping", interactionEventId)
		return nil
	}

	summaryPrompt := fmt.Sprintf(h.cfg.Services.Anthropic.EmailSummaryPrompt, interactionEventContent)

	promptLog := commonEntity.AiPromptLog{
		CreatedAt:      utils.Now(),
		AppSource:      constants.AppSourceEventProcessingPlatform,
		Provider:       constants.Anthropic,
		Model:          "claude-2",
		PromptType:     constants.PromptType_EmailSummary,
		Tenant:         &eventData.Tenant,
		NodeId:         &interactionEventId,
		NodeLabel:      utils.StringPtr(neo4jutil.NodeLabelInteractionEvent),
		PromptTemplate: &h.cfg.Services.Anthropic.EmailSummaryPrompt,
		Prompt:         summaryPrompt,
	}
	promptStoreLogId, err := h.repositories.CommonRepositories.AiPromptLogRepository.Store(promptLog)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error storing prompt log: %v", err)
	} else {
		span.LogFields(log.String("promptStoreLogId", promptStoreLogId))
	}

	aiResponse, err := h.aiModel.Inference(ctx, summaryPrompt) // ai.InvokeAnthropic(ctx, h.cfg, h.log, summaryPrompt)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error invoking AI: %v", err.Error())
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId, err.Error())
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with error: %v", storeErr)
		}
		return nil
	} else {
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateResponse(promptStoreLogId, aiResponse)
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with ai response: %v", storeErr)
		}
	}
	summary := utils.ExtractAfterColon(aiResponse)

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err = subscriptions.CallEventsPlatformGRPCWithRetry[*interactioneventpb.InteractionEventIdGrpcResponse](func() (*interactioneventpb.InteractionEventIdGrpcResponse, error) {
		return h.grpcClients.InteractionEventClient.ReplaceSummary(ctx, &interactioneventpb.ReplaceSummaryGrpcRequest{
			Tenant:             eventData.Tenant,
			InteractionEventId: interactionEventId,
			AppSource:          constants.AppSourceEventProcessingPlatform,
			Summary:            summary,
			ContentType:        "text/plain",
		})
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error replacing summary: %v", err)
		return err
	}

	return nil
}

func (h *interactionEventHandler) GenerateActionItemsForEmail(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventHandler.GenerateActionItemsForEmail")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData event.InteractionEventRequestSummaryEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	interactionEventId := aggregate.GetInteractionEventObjectID(evt.AggregateID, eventData.Tenant)
	span.LogFields(log.String("interactionEventId", interactionEventId))

	interactionEvent, err := h.repositories.Neo4jRepositories.InteractionEventReadRepository.GetInteractionEvent(ctx, eventData.Tenant, interactionEventId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error getting interaction event with id %s: %v", interactionEvent, err)
		return nil
	}

	interactionEventChannel := utils.GetStringPropOrEmpty(interactionEvent.Props, "channel")
	interactionEventContent := utils.GetStringPropOrEmpty(interactionEvent.Props, "content")

	if interactionEventChannel != "EMAIL" {
		tracing.TraceErr(span, errors.New("interaction event is not an email"))
		h.log.Warnf("Interaction event with id %s is not an email, skipping", interactionEventId)
		return nil
	}
	if interactionEventContent == "" {
		tracing.TraceErr(span, errors.New("interaction event content is empty"))
		h.log.Warnf("Interaction event with id %s has no content, skipping", interactionEventId)
		return nil
	}

	actionItemsPrompt := fmt.Sprintf(h.cfg.Services.Anthropic.EmailActionsItemsPrompt, interactionEventContent)

	promptLog := commonEntity.AiPromptLog{
		CreatedAt:      utils.Now(),
		AppSource:      constants.AppSourceEventProcessingPlatform,
		Provider:       constants.Anthropic,
		Model:          "claude-2",
		PromptType:     constants.PromptType_EmailActionItems,
		Tenant:         &eventData.Tenant,
		NodeId:         &interactionEventId,
		NodeLabel:      utils.StringPtr(neo4jutil.NodeLabelInteractionEvent),
		PromptTemplate: &h.cfg.Services.Anthropic.EmailActionsItemsPrompt,
		Prompt:         actionItemsPrompt,
	}
	promptStoreLogId, err := h.repositories.CommonRepositories.AiPromptLogRepository.Store(promptLog)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error storing prompt log: %v", err)
	} else {
		span.LogFields(log.String("promptStoreLogId", promptStoreLogId))
	}

	aiResponse, err := h.aiModel.Inference(ctx, actionItemsPrompt) // ai.InvokeAnthropic(ctx, h.cfg, h.log, actionItemsPrompt)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error invoking AI: %v", err.Error())
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId, err.Error())
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with error: %v", storeErr)
		}
		return nil
	} else {
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateResponse(promptStoreLogId, aiResponse)
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with ai response: %v", storeErr)
		}
	}

	actionItems, err := extractActionItemsFromAiResponse(aiResponse)
	span.LogFields(log.Object("output - actionItems", actionItems))
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error extracting action items from ai response: %v", err)
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId, err.Error())
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with error: %v", storeErr)
		}
		return nil
	}
	if len(actionItems) == 0 {
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId, err.Error())
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with error: %v", storeErr)
		}
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err = subscriptions.CallEventsPlatformGRPCWithRetry[*interactioneventpb.InteractionEventIdGrpcResponse](func() (*interactioneventpb.InteractionEventIdGrpcResponse, error) {
		return h.grpcClients.InteractionEventClient.ReplaceActionItems(ctx, &interactioneventpb.ReplaceActionItemsGrpcRequest{
			Tenant:             eventData.Tenant,
			InteractionEventId: interactionEventId,
			AppSource:          constants.AppSourceEventProcessingPlatform,
			ActionItems:        actionItems,
		})
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error replacing action items: %v", err)
		return err
	}

	return nil
}

func extractActionItemsFromAiResponse(str string) ([]string, error) {
	jsonStr, err := utils.ExtractJsonFromString(str)
	if err != nil {
		return []string{}, err
	}

	var data map[string]interface{}

	json.Unmarshal([]byte(jsonStr), &data)

	items, ok := data["items"].([]interface{})
	if !ok {
		return []string{}, fmt.Errorf("invalid JSON format")
	}

	var actionItems []string
	for _, item := range items {
		actionItems = append(actionItems, item.(string))
	}

	return actionItems, nil
}
