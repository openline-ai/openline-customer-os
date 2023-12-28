package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// TimelineEvents is the resolver for the timelineEvents field.
func (r *queryResolver) TimelineEvents(ctx context.Context, ids []string) ([]model.TimelineEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.TimelineEvents", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("request.ids", ids))

	timelineEvents, err := r.Services.TimelineEventService.GetTimelineEventsWithIds(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch timeline events")
		return nil, nil
	}
	return mapper.MapEntitiesToTimelineEvents(timelineEvents), err
}
