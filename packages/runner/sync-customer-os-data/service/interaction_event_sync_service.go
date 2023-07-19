package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/common"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/logger"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/repository"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type interactionEventSyncService struct {
	repositories *repository.Repositories
	log          logger.Logger
}

func NewDefaultInteractionEventSyncService(repositories *repository.Repositories, log logger.Logger) SyncService {
	return &interactionEventSyncService{
		repositories: repositories,
		log:          log,
	}
}

func (s *interactionEventSyncService) Sync(ctx context.Context, dataService source.SourceDataService, syncDate time.Time, tenant, runId string, batchSize int) (int, int, int) {
	span, ctx := tracing.StartTracerSpan(ctx, "InteractionEventSyncService.Sync")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	completed, failed, skipped := 0, 0, 0
	for {
		interactionEvents := dataService.GetDataForSync(ctx, common.INTERACTION_EVENTS, batchSize, runId)
		if len(interactionEvents) == 0 {
			s.log.Debugf("no interaction found for sync from %s for tenant %s", dataService.SourceId(), tenant)
			break
		}
		s.log.Infof("syncing %d interaction events from %s for tenant %s", len(interactionEvents), dataService.SourceId(), tenant)

		for _, v := range interactionEvents {
			s.syncInteractionEvent(ctx, v.(entity.InteractionEventData), dataService, syncDate, tenant, runId, &completed, &failed, &skipped)
		}

		if len(interactionEvents) < batchSize {
			break
		}
	}
	return completed, failed, skipped
}

func (s *interactionEventSyncService) syncInteractionEvent(ctx context.Context, interactionEventInput entity.InteractionEventData, dataService source.SourceDataService, syncDate time.Time, tenant, runId string, completed, failed, skipped *int) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventSyncService.syncInteractionEvent")
	defer span.Finish()
	tracing.SetDefaultSyncServiceSpanTags(ctx, span)

	var failedSync = false
	var reason string
	interactionEventInput.Normalize()

	if interactionEventInput.Skip {
		if err := dataService.MarkProcessed(ctx, interactionEventInput.SyncId, runId, true, true, interactionEventInput.SkipReason); err != nil {
			*failed++
			span.LogFields(log.Bool("failedSync", true))
			return
		}
		*skipped++
		span.LogFields(log.Bool("skippedSync", true))
		return
	}

	interactionEventId, err := s.repositories.InteractionEventRepository.GetMatchedInteractionEvent(ctx, tenant, interactionEventInput)
	if err != nil {
		failedSync = true
		tracing.TraceErr(span, err)
		reason = fmt.Sprintf("failed finding existing matched interaction event with external reference id %v for tenant %v :%v", interactionEventInput.ExternalId, tenant, err)
		s.log.Error(reason)
	}

	// Create new note id if not found
	if interactionEventId == "" {
		ieUuid, _ := uuid.NewRandom()
		interactionEventId = ieUuid.String()
	}
	interactionEventInput.Id = interactionEventId

	if !failedSync {
		err = s.repositories.InteractionEventRepository.MergeInteractionEvent(ctx, tenant, syncDate, interactionEventInput)
		if err != nil {
			failedSync = true
			tracing.TraceErr(span, err)
			reason = fmt.Sprintf("failed merge interaction event with external reference %v for tenant %v :%v", interactionEventInput.ExternalId, tenant, err)
			s.log.Error(reason)
		}
	}

	if !failedSync && interactionEventInput.IsPartOf() {
		err = s.repositories.InteractionEventRepository.LinkInteractionEventAsPartOfByExternalId(ctx, tenant, interactionEventInput)
		if err != nil {
			failedSync = true
			tracing.TraceErr(span, err)
			reason = fmt.Sprintf("failed link interaction event as part of by external reference %v for tenant %v :%v", interactionEventInput.ExternalId, tenant, err)
			s.log.Error(reason)
		}
	}

	if !failedSync && interactionEventInput.HasSender() {
		err = s.repositories.InteractionEventRepository.LinkInteractionEventWithSenderByExternalId(ctx, tenant, interactionEventId, interactionEventInput.ExternalSystem, interactionEventInput.SentBy)
		if err != nil {
			failedSync = true
			tracing.TraceErr(span, err)
			reason = fmt.Sprintf("failed link interaction event with sender by external reference %v for tenant %v :%v", interactionEventInput.ExternalId, tenant, err)
			s.log.Error(reason)
		}
	}

	if !failedSync && interactionEventInput.HasRecipients() {
		for _, recipient := range interactionEventInput.SentTo {
			err = s.repositories.InteractionEventRepository.LinkInteractionEventWithRecipientByExternalId(ctx, tenant, interactionEventId, interactionEventInput.ExternalSystem, recipient)
			if err != nil {
				failedSync = true
				tracing.TraceErr(span, err)
				reason = fmt.Sprintf("failed link interaction event with recipient by external reference %v for tenant %v :%v", interactionEventInput.ExternalId, tenant, err)
				s.log.Error(reason)
			}
		}
	}

	if failedSync == false {
		s.log.Debugf("successfully merged interaction event with id %v for tenant %v from %v", interactionEventId, tenant, dataService.SourceId())
	}
	if err := dataService.MarkProcessed(ctx, interactionEventInput.SyncId, runId, failedSync == false, false, reason); err != nil {
		*failed++
		span.LogFields(log.Bool("failedSync", true))
		return
	}
	if failedSync == true {
		*failed++
	} else {
		*completed++
	}
	span.LogFields(log.Bool("failedSync", failedSync))
}
