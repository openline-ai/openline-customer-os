package organization

import (
	"context"
	"fmt"
	"strings"
	"time"

	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jmapper "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/subscriptions"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	organizationpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/organization"

	ai "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-ai/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/data"
	commonEntity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/postgres/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/caches"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

const (
	Unknown = "Unknown"
)

type WebscrapeRequest struct {
	Domain string `json:"scrape"`
}

type Socials struct {
	Github    string `json:"github,omitempty"`
	Linkedin  string `json:"linkedin,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Facebook  string `json:"facebook,omitempty"`
}

type WebscrapeResponseV1 struct {
	CompanyName          string `json:"companyName,omitempty"`
	Website              string `json:"website,omitempty"`
	Market               string `json:"market,omitempty"`
	Industry             string `json:"industry,omitempty"`
	IndustryGroup        string `json:"industryGroup,omitempty"`
	SubIndustry          string `json:"subIndustry,omitempty"`
	TargetAudience       string `json:"targetAudience,omitempty"`
	ValueProposition     string `json:"valueProposition,omitempty"`
	Github               string `json:"github,omitempty"`
	Linkedin             string `json:"linkedin,omitempty"`
	Twitter              string `json:"twitter,omitempty"`
	Youtube              string `json:"youtube,omitempty"`
	Instagram            string `json:"instagram,omitempty"`
	Facebook             string `json:"facebook,omitempty"`
	CompanySize          int64  `json:"companySize,omitempty"`
	EmployeeGrowthRate   string `json:"employeeGrowthRate,omitempty"`
	HeadquartersLocation string `json:"headquartersLocation,omitempty"`
	YearFounded          int64  `json:"yearFounded,omitempty"`
	LogoUrl              string `json:"logoUrl,omitempty"`
}

type organizationEventHandler struct {
	repositories  *repository.Repositories
	log           logger.Logger
	cfg           *config.Config
	caches        caches.Cache
	domainScraper WebScraper
	aiModel       ai.AiModel
	grpcClients   *grpc_client.Clients
}

func NewOrganizationEventHandler(repositories *repository.Repositories, log logger.Logger, cfg *config.Config, caches caches.Cache, domainScraper WebScraper, aiModel ai.AiModel, grpcClients *grpc_client.Clients) *organizationEventHandler {
	return &organizationEventHandler{
		repositories:  repositories,
		log:           log,
		cfg:           cfg,
		caches:        caches,
		domainScraper: domainScraper,
		aiModel:       aiModel,
		grpcClients:   grpcClients,
	}
}

func (h *organizationEventHandler) WebScrapeOrganizationByDomain(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.WebScrapeOrganizationByDomain")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData events.OrganizationLinkDomainEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	organizationId := aggregate.GetOrganizationObjectID(evt.AggregateID, eventData.Tenant)
	span.LogFields(log.String("organizationId", organizationId))
	if eventData.Domain == "" {
		tracing.TraceErr(span, errors.New("domain is empty"))
		h.log.Errorf("Missing domain in event data: %v", eventData)
		return nil
	}

	return h.webScrapeOrganization(ctx, eventData.Tenant, organizationId, eventData.Domain)
}

func (h *organizationEventHandler) WebScrapeOrganizationByWebsite(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.WebScrapeOrganizationByWebsite")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData events.OrganizationRequestScrapeByWebsite
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	organizationId := aggregate.GetOrganizationObjectID(evt.AggregateID, eventData.Tenant)
	span.SetTag(tracing.SpanTagEntityId, organizationId)

	organizationNodeAvailable := subscriptions.WaitCheckNodeExistsInNeo4j(ctx, h.repositories.Neo4jRepositories, eventData.Tenant, organizationId, neo4jutil.NodeLabelOrganization)
	if !organizationNodeAvailable {
		err := errors.Errorf("%s node %s not available in neo4j", neo4jutil.NodeLabelOrganization, organizationId)
		tracing.TraceErr(span, err)
		return nil
	}

	span.LogFields(log.String("website", eventData.Website))
	if eventData.Website == "" {
		tracing.TraceErr(span, errors.New("website is empty"))
		h.log.Errorf("Missing website in event data: %v", eventData)
		return nil
	}

	return h.webScrapeOrganization(ctx, eventData.Tenant, organizationId, eventData.Website)
}

func (h *organizationEventHandler) webScrapeOrganization(ctx context.Context, tenant, organizationId, url string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.webScrapeOrganization")
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, tenant)
	span.LogFields(log.String("organizationId", organizationId), log.String("url", url))

	organizationDbNode, err := h.repositories.Neo4jRepositories.OrganizationReadRepository.GetOrganization(ctx, tenant, organizationId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error getting organization with id %s: %v", organizationId, err)
		return nil
	}
	if organizationDbNode == nil {
		tracing.TraceErr(span, errors.New("organization not found"))
		h.log.Errorf("Organization with id %s not found", organizationId)
		return nil
	}
	organization := neo4jmapper.MapDbNodeToOrganizationEntity(organizationDbNode)

	// if already web scraped for this url, skip
	if organization.WebScrapeDetails.WebScrapedUrl == url {
		h.log.Infof("Organization {%s} already web scraped for url {%s}", organizationId, url)
		return nil
	}

	// register web scraping request and attempts
	attempt := int64(1)
	if organization.WebScrapeDetails.WebScrapeLastRequestedUrl == url {
		attempt = organization.WebScrapeDetails.WebScrapeAttempts + 1
	}
	err = h.repositories.Neo4jRepositories.OrganizationWriteRepository.WebScrapeRequested(ctx, tenant, organizationId, url, attempt, utils.Now())
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error registering web scrape request: %v", err)
	}

	result, err := h.domainScraper.Scrape(url, tenant, organizationId, false)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error scraping url %s: %v", url, err)
		// if organization name is empty set it to the domain name
		h.updateOrganizationNameIfEmpty(ctx, tenant, url, organization, span)
		return nil
	}

	fieldsMask := []string{model.FieldMaskMarket, model.FieldMaskIndustry, model.FieldMaskIndustryGroup, model.FieldMaskSubIndustry, model.FieldMaskTargetAudience, model.FieldMaskValueProposition}
	// name organization name if missing
	if organization.Name == "" {
		fieldsMask = append(fieldsMask, model.FieldMaskName)
	}
	// set website if missing
	if organization.Website == "" {
		fieldsMask = append(fieldsMask, model.FieldMaskWebsite)
	}

	// add fields to fields mask
	fieldsMask = *h.addFieldMasks(&model.OrganizationDataFields{
		Employees:          result.CompanySize,
		YearFounded:        &result.YearFounded,
		Headquarters:       result.HeadquartersLocation,
		EmployeeGrowthRate: result.EmployeeGrowthRate,
		LogoUrl:            result.LogoUrl,
	}, fieldsMask)

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	updateGrpcRequest := organizationpb.UpdateOrganizationGrpcRequest{
		Tenant:         tenant,
		OrganizationId: organizationId,
		SourceFields: &commonpb.SourceFields{
			AppSource: constants.AppSourceEventProcessingPlatform,
			Source:    constants.SourceWebscrape,
		},
		FieldsMask:         prepareFieldsMaskForGrpcRequests(fieldsMask),
		WebScrapedUrl:      url,
		Market:             result.Market,
		Industry:           result.Industry,
		IndustryGroup:      result.IndustryGroup,
		SubIndustry:        result.SubIndustry,
		TargetAudience:     result.TargetAudience,
		ValueProposition:   result.ValueProposition,
		LogoUrl:            result.LogoUrl,
		Headquarters:       result.HeadquartersLocation,
		EmployeeGrowthRate: result.EmployeeGrowthRate,
		Employees:          result.CompanySize,
		YearFounded:        &result.YearFounded,
		Name:               result.CompanyName,
		Website:            result.Website,
	}
	_, err = subscriptions.CallEventsPlatformGRPCWithRetry[*organizationpb.OrganizationIdGrpcResponse](func() (*organizationpb.OrganizationIdGrpcResponse, error) {
		return h.grpcClients.OrganizationClient.UpdateOrganization(ctx, &updateGrpcRequest)
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error updating organization: %s", err.Error())
	}

	if result.Youtube != "" {
		h.addSocial(ctx, organizationId, tenant, "youtube", result.Youtube)
	}
	if result.Twitter != "" {
		h.addSocial(ctx, organizationId, tenant, "twitter", result.Twitter)
	}
	if result.Linkedin != "" {
		h.addSocial(ctx, organizationId, tenant, "linkedin", result.Linkedin)
	}
	if result.Github != "" {
		h.addSocial(ctx, organizationId, tenant, "github", result.Github)
	}
	if result.Instagram != "" {
		h.addSocial(ctx, organizationId, tenant, "instagram", result.Instagram)
	}
	if result.Facebook != "" {
		h.addSocial(ctx, organizationId, tenant, "facebook", result.Facebook)
	}

	return nil
}

func (h *organizationEventHandler) addFieldMasks(orgFields *model.OrganizationDataFields, fieldMasks []string) *[]string {
	if orgFields.Employees != 0 {
		fieldMasks = append(fieldMasks, model.FieldMaskEmployees)
	}
	if *orgFields.YearFounded != 0 {
		fieldMasks = append(fieldMasks, model.FieldMaskYearFounded)
	}
	if orgFields.Headquarters != "" {
		fieldMasks = append(fieldMasks, model.FieldMaskHeadquarters)
	}
	if orgFields.EmployeeGrowthRate != "" {
		fieldMasks = append(fieldMasks, model.FieldMaskEmployeeGrowthRate)
	}
	if orgFields.LogoUrl != "" {
		fieldMasks = append(fieldMasks, model.FieldMaskLogoUrl)
	}
	return &fieldMasks
}

func (h *organizationEventHandler) updateOrganizationNameIfEmpty(ctx context.Context, tenant, url string, organization *neo4jentity.OrganizationEntity, span opentracing.Span) {
	if organization.Name == "" && strings.Contains(url, ".") {
		ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
		_, err := subscriptions.CallEventsPlatformGRPCWithRetry[*organizationpb.OrganizationIdGrpcResponse](func() (*organizationpb.OrganizationIdGrpcResponse, error) {
			return h.grpcClients.OrganizationClient.UpdateOrganization(ctx, &organizationpb.UpdateOrganizationGrpcRequest{
				Tenant:         tenant,
				OrganizationId: organization.ID,
				SourceFields: &commonpb.SourceFields{
					AppSource: constants.AppSourceEventProcessingPlatform,
					Source:    constants.SourceWebscrape,
				},
				Name:       utils.ExtractFirstPart(utils.ExtractDomain(url), "."),
				FieldsMask: []organizationpb.OrganizationMaskField{organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_NAME},
			})
		})
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Error updating organization: %v", err)
		}
	}
}

func (h *organizationEventHandler) addSocial(ctx context.Context, organizationId, tenant, platform, url string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.addSocial")
	defer span.Finish()
	span.LogFields(log.String("organizationId", organizationId), log.String("tenant", tenant), log.String("platform", platform), log.String("url", url))
	span.SetTag(tracing.SpanTagEntityId, organizationId)

	_, err := subscriptions.CallEventsPlatformGRPCWithRetry[*organizationpb.OrganizationIdGrpcResponse](func() (*organizationpb.OrganizationIdGrpcResponse, error) {
		return h.grpcClients.OrganizationClient.AddSocial(ctx, &organizationpb.AddSocialGrpcRequest{
			Tenant:         tenant,
			OrganizationId: organizationId,
			SourceFields: &commonpb.SourceFields{
				AppSource:     constants.AppSourceEventProcessingPlatform,
				Source:        constants.SourceWebscrape,
				SourceOfTruth: constants.SourceWebscrape,
			},
			Platform: platform,
			Url:      url,
		})
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error adding %s social: %s", platform, err.Error())
	}
}

func (h *organizationEventHandler) AdjustNewOrganizationFields(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.AdjustNewOrganizationFields")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData events.OrganizationCreateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	organizationId := aggregate.GetOrganizationObjectID(evt.AggregateID, eventData.Tenant)
	span.SetTag(tracing.SpanTagEntityId, organizationId)
	span.SetTag(tracing.SpanTagTenant, eventData.Tenant)

	market := h.mapMarketValue(eventData.Market)
	industry := h.mapIndustryToGICS(ctx, eventData.Tenant, organizationId, eventData.Industry)

	// wait for organization to be created in neo4j before updating it
	for attempt := 1; attempt <= constants.MaxRetriesCheckDataInNeo4j; attempt++ {
		exists, err := h.repositories.Neo4jRepositories.CommonReadRepository.ExistsById(ctx, eventData.Tenant, organizationId, neo4jutil.NodeLabelOrganization)
		if err == nil && exists {
			break
		}
		time.Sleep(utils.BackOffExponentialDelay(attempt))
	}

	if eventData.Market != market || eventData.Industry != industry {
		err := h.callUpdateOrganizationCommand(ctx, eventData.Tenant, organizationId, eventData.SourceOfTruth, market, industry, span)
		if err != nil {
			tracing.TraceErr(span, err)
			return err
		}
	} else {
		h.log.Infof("No need to update organization %s", organizationId)
	}
	return nil
}

func (h *organizationEventHandler) AdjustUpdatedOrganizationFields(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.AdjustUpdatedOrganizationFields")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))
	tracing.LogObjectAsJson(span, "eventData", evt)

	var eventData events.OrganizationUpdateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	organizationId := aggregate.GetOrganizationObjectID(evt.AggregateID, eventData.Tenant)

	market := h.mapMarketValue(eventData.Market)
	industry := h.mapIndustryToGICS(ctx, eventData.Tenant, organizationId, eventData.Industry)

	if eventData.Market != market || eventData.Industry != industry {
		err := h.callUpdateOrganizationCommand(ctx, eventData.Tenant, organizationId, eventData.Source, market, industry, span)
		if err != nil {
			tracing.TraceErr(span, err)
			return err
		}
	} else {
		h.log.Infof("No need to update organization %s", organizationId)
	}
	return nil
}

func (h *organizationEventHandler) callUpdateOrganizationCommand(ctx context.Context, tenant, organizationId, source, market, industry string, span opentracing.Span) error {
	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err := subscriptions.CallEventsPlatformGRPCWithRetry[*organizationpb.OrganizationIdGrpcResponse](func() (*organizationpb.OrganizationIdGrpcResponse, error) {
		return h.grpcClients.OrganizationClient.UpdateOrganization(ctx, &organizationpb.UpdateOrganizationGrpcRequest{
			Tenant:         tenant,
			OrganizationId: organizationId,
			SourceFields: &commonpb.SourceFields{
				AppSource: constants.AppSourceEventProcessingPlatform,
				Source:    source,
			},
			Market:   market,
			Industry: industry,
			FieldsMask: []organizationpb.OrganizationMaskField{
				organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_MARKET,
				organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_INDUSTRY,
			},
		})
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error updating organization %s: %s", organizationId, err.Error())
		return err
	}
	return nil
}

func (h *organizationEventHandler) mapMarketValue(inputMarket string) string {
	return data.AdjustOrganizationMarket(inputMarket)
}

func (h *organizationEventHandler) mapIndustryToGICS(ctx context.Context, tenant, orgId, inputIndustry string) string {
	trimmedInputIndustry := strings.TrimSpace(inputIndustry)

	if inputIndustry == "" {
		return ""
	}

	var industry string
	if industryValue, ok := h.caches.GetIndustry(trimmedInputIndustry); ok {
		industry = industryValue
	} else {
		h.log.Infof("Industry %s not found in cache, asking AI", trimmedInputIndustry)
		industry = h.mapIndustryToGICSWithAI(ctx, tenant, orgId, trimmedInputIndustry)
		if industry != "" && len(industry) < 45 {
			h.caches.SetIndustry(trimmedInputIndustry, industry)
			h.log.Infof("Industry %s mapped to %s", trimmedInputIndustry, industry)
		} else {
			h.log.Warnf("Industry %s mapped wrongly to (%s) with AI, returning input value", industry, trimmedInputIndustry)
			return trimmedInputIndustry
		}
	}
	if industry == Unknown {
		h.log.Infof("Unknown industry %s, returning as is", trimmedInputIndustry)
		return trimmedInputIndustry
	}

	return strings.TrimSpace(industry)
}

func (h *organizationEventHandler) mapIndustryToGICSWithAI(ctx context.Context, tenant, orgId, inputIndustry string) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OrganizationEventHandler.mapIndustryToGICSWithAI")
	defer span.Finish()
	span.LogFields(log.String("inputIndustry", inputIndustry))

	firstPrompt := fmt.Sprintf(h.cfg.Services.Anthropic.IndustryLookupPrompt1, inputIndustry)

	promptLog1 := commonEntity.AiPromptLog{
		CreatedAt:      utils.Now(),
		AppSource:      constants.AppSourceEventProcessingPlatform,
		Provider:       constants.Anthropic,
		Model:          "claude-2",
		PromptType:     constants.PromptType_MapIndustry,
		Tenant:         &tenant,
		NodeId:         &orgId,
		NodeLabel:      utils.StringPtr(neo4jutil.NodeLabelOrganization),
		PromptTemplate: &h.cfg.Services.Anthropic.IndustryLookupPrompt1,
		Prompt:         firstPrompt,
	}
	promptStoreLogId1, err := h.repositories.CommonRepositories.AiPromptLogRepository.Store(promptLog1)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error storing prompt log: %v", err)
	} else {
		span.LogFields(log.String("promptStoreLogId1", promptStoreLogId1))
	}

	firstResult, err := h.aiModel.Inference(ctx, firstPrompt) // ai.InvokeAnthropic(ctx, h.cfg, h.log, firstPrompt)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error invoking AI: %v", err)
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId1, err.Error())
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with error: %v", storeErr)
		}
		return ""
	} else {
		storeErr := h.repositories.CommonRepositories.AiPromptLogRepository.UpdateResponse(promptStoreLogId1, firstResult)
		if storeErr != nil {
			tracing.TraceErr(span, storeErr)
			h.log.Errorf("Error updating prompt log with ai response: %v", storeErr)
		}
	}
	if firstResult == "" || firstResult == Unknown {
		return firstResult
	}
	secondPrompt := fmt.Sprintf(h.cfg.Services.Anthropic.IndustryLookupPrompt2, firstResult)

	promptLog2 := commonEntity.AiPromptLog{
		CreatedAt:      utils.Now(),
		AppSource:      constants.AppSourceEventProcessingPlatform,
		Provider:       constants.Anthropic,
		Model:          "claude-2",
		PromptType:     constants.PromptType_ExtractIndustryValue,
		Tenant:         &tenant,
		NodeId:         &orgId,
		NodeLabel:      utils.StringPtr(neo4jutil.NodeLabelOrganization),
		PromptTemplate: &h.cfg.Services.Anthropic.IndustryLookupPrompt2,
		Prompt:         secondPrompt,
	}
	promptStoreLogId2, err := h.repositories.CommonRepositories.AiPromptLogRepository.Store(promptLog2)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error storing prompt log with error: %v", err)
	}
	secondResult, err := h.aiModel.Inference(ctx, secondPrompt) // ai.InvokeAnthropic(ctx, h.cfg, h.log, secondPrompt)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error invoking AI: %v", err)
		err = h.repositories.CommonRepositories.AiPromptLogRepository.UpdateError(promptStoreLogId2, err.Error())
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Error updating prompt log with error: %v", err)
		}
		return ""
	} else {
		err = h.repositories.CommonRepositories.AiPromptLogRepository.UpdateResponse(promptStoreLogId2, secondResult)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Error updating prompt log with ai response: %v", err)
		}
	}
	return secondResult
}

func prepareFieldsMaskForGrpcRequests(fields []string) []organizationpb.OrganizationMaskField {
	organizationFieldsMask := make([]organizationpb.OrganizationMaskField, 0)
	for _, field := range fields {
		switch field {
		case model.FieldMaskDescription:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_DESCRIPTION)
		case model.FieldMaskName:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_NAME)
		case model.FieldMaskWebsite:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_WEBSITE)
		case model.FieldMaskMarket:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_MARKET)
		case model.FieldMaskIndustry:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_INDUSTRY)
		case model.FieldMaskIndustryGroup:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_INDUSTRY_GROUP)
		case model.FieldMaskSubIndustry:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_SUB_INDUSTRY)
		case model.FieldMaskTargetAudience:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_TARGET_AUDIENCE)
		case model.FieldMaskValueProposition:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_VALUE_PROPOSITION)
		case model.FieldMaskEmployees:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_EMPLOYEES)
		case model.FieldMaskYearFounded:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_YEAR_FOUNDED)
		case model.FieldMaskHeadquarters:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_HEADQUARTERS)
		case model.FieldMaskEmployeeGrowthRate:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_EMPLOYEE_GROWTH_RATE)
		case model.FieldMaskLogoUrl:
			organizationFieldsMask = append(organizationFieldsMask, organizationpb.OrganizationMaskField_ORGANIZATION_PROPERTY_LOGO_URL)
		}
	}
	return organizationFieldsMask
}
