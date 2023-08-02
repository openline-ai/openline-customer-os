package commands

import (
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
)

type InteractionEventCommands struct {
	RequestSummary RequestSummaryCommandHandler
	ReplaceSummary ReplaceSummaryCommandHandler
}

func NewInteractionEventCommands(log logger.Logger, cfg *config.Config, es eventstore.AggregateStore) *InteractionEventCommands {
	return &InteractionEventCommands{
		RequestSummary: NewRequestSummaryCommandHandler(log, cfg, es),
		ReplaceSummary: NewReplaceSummaryCommandHandler(log, cfg, es),
	}
}
