package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// Milestones is the resolver for the milestones field.
func (r *masterPlanResolver) Milestones(ctx context.Context, obj *model.MasterPlan) ([]*model.MasterPlanMilestone, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	masterPlanMilestonesEntities, err := dataloader.For(ctx).GetMasterPlanMilestonesForMasterPlan(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get master plan milestones for master plan %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get master plan milestones for master plan %s", obj.ID)
		return nil, nil
	}
	allMasterPlanMilestones := mapper.MapEntitiesToMasterPlanMilestones(masterPlanMilestonesEntities)
	// filter out retired milestones
	var milestones []*model.MasterPlanMilestone
	for _, masterPlanMilestone := range allMasterPlanMilestones {
		if !masterPlanMilestone.Retired {
			milestones = append(milestones, masterPlanMilestone)
		}
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(log.Int("response.length", len(milestones)))
	}
	return milestones, nil
}

// RetiredMilestones is the resolver for the retiredMilestones field.
func (r *masterPlanResolver) RetiredMilestones(ctx context.Context, obj *model.MasterPlan) ([]*model.MasterPlanMilestone, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	masterPlanMilestonesEntities, err := dataloader.For(ctx).GetMasterPlanMilestonesForMasterPlan(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get retired master plan milestones for master plan %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get retired master plan milestones for master plan %s", obj.ID)
		return nil, nil
	}
	allMasterPlanMilestones := mapper.MapEntitiesToMasterPlanMilestones(masterPlanMilestonesEntities)
	// filter out non-retired milestones
	var milestones []*model.MasterPlanMilestone
	for _, masterPlanMilestone := range allMasterPlanMilestones {
		if masterPlanMilestone.Retired {
			milestones = append(milestones, masterPlanMilestone)
		}
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(log.Int("response.length", len(milestones)))
	}
	return milestones, nil
}

// MasterPlanCreate is the resolver for the masterPlan_Create field.
func (r *mutationResolver) MasterPlanCreate(ctx context.Context, input model.MasterPlanInput) (*model.MasterPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	masterPlanId, err := r.Services.MasterPlanService.CreateMasterPlan(ctx, utils.IfNotNilString(input.Name))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create master plan")
		return &model.MasterPlan{ID: masterPlanId}, err
	}

	createdMasterPlanEntity, err := r.Services.MasterPlanService.GetMasterPlanById(ctx, masterPlanId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Master plan details not yet available. Master plan id: %s", masterPlanId)
		return &model.MasterPlan{ID: masterPlanId}, nil
	}
	span.LogFields(log.String("response.masterPlanId", masterPlanId))
	return mapper.MapEntityToMasterPlan(createdMasterPlanEntity), nil
}

// MasterPlanUpdate is the resolver for the masterPlan_Update field.
func (r *mutationResolver) MasterPlanUpdate(ctx context.Context, input model.MasterPlanUpdateInput) (*model.MasterPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.MasterPlanService.UpdateMasterPlan(ctx, input.ID, input.Name, input.Retired)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update master plan")
		return nil, err
	}

	updatedMasterPlanEntity, err := r.Services.MasterPlanService.GetMasterPlanById(ctx, input.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Master plan with id %s", input.ID)
		return nil, nil
	}
	return mapper.MapEntityToMasterPlan(updatedMasterPlanEntity), nil
}

// MasterPlanDuplicate is the resolver for the masterPlan_Duplicate field.
func (r *mutationResolver) MasterPlanDuplicate(ctx context.Context, id string) (*model.MasterPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanDuplicate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "id", id)

	masterPlanId, err := r.Services.MasterPlanService.DuplicateMasterPlan(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to duplicate master plan")
		return nil, err
	}

	createdMasterPlanEntity, err := r.Services.MasterPlanService.GetMasterPlanById(ctx, masterPlanId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Master plan details not yet available. Master plan id: %s", masterPlanId)
		return &model.MasterPlan{ID: masterPlanId}, nil
	}
	span.LogFields(log.String("response.masterPlanId", masterPlanId))
	return mapper.MapEntityToMasterPlan(createdMasterPlanEntity), nil
}

// MasterPlanMilestoneCreate is the resolver for the masterPlanMilestone_Create field.
func (r *mutationResolver) MasterPlanMilestoneCreate(ctx context.Context, input model.MasterPlanMilestoneInput) (*model.MasterPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanMilestoneCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	masterPlanMilestoneId, err := r.Services.MasterPlanService.CreateMasterPlanMilestone(ctx, input.MasterPlanID, utils.IfNotNilString(input.Name),
		input.Order, input.DurationHours, input.Optional, input.Items)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create master plan milestone")
		return &model.MasterPlanMilestone{ID: masterPlanMilestoneId}, err
	}

	createdMasterPlanMilestoneEntity, err := r.Services.MasterPlanService.GetMasterPlanMilestoneById(ctx, masterPlanMilestoneId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Master plan milestone details not yet available. Master plan milestone id: %s", masterPlanMilestoneId)
		return &model.MasterPlanMilestone{ID: masterPlanMilestoneId}, nil
	}
	span.LogFields(log.String("response.masterPlanMilestoneId", masterPlanMilestoneId))
	return mapper.MapEntityToMasterPlanMilestone(createdMasterPlanMilestoneEntity), nil
}

// MasterPlanMilestoneUpdate is the resolver for the masterPlanMilestone_Update field.
func (r *mutationResolver) MasterPlanMilestoneUpdate(ctx context.Context, input model.MasterPlanMilestoneUpdateInput) (*model.MasterPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanMilestoneUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.MasterPlanService.UpdateMasterPlanMilestone(ctx, input.MasterPlanID, input.ID, input.Name,
		input.Order, input.DurationHours, input.Items, input.Optional, input.Retired)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update master plan milestone")
		return nil, err
	}

	updatedMasterPlanMilestoneEntity, err := r.Services.MasterPlanService.GetMasterPlanMilestoneById(ctx, input.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get master plan milestone with id %s", input.ID)
		return nil, nil
	}
	return mapper.MapEntityToMasterPlanMilestone(updatedMasterPlanMilestoneEntity), nil
}

// MasterPlanMilestoneReorder is the resolver for the masterPlanMilestone_Reorder field.
func (r *mutationResolver) MasterPlanMilestoneReorder(ctx context.Context, input model.MasterPlanMilestoneReorderInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanMilestoneReorder", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.MasterPlanService.ReorderMasterPlanMilestones(ctx, input.MasterPlanID, input.OrderedIds)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to reorder master plan milestones")
		return "", err
	}
	return input.MasterPlanID, nil
}

// MasterPlanMilestoneDuplicate is the resolver for the masterPlanMilestone_Duplicate field.
func (r *mutationResolver) MasterPlanMilestoneDuplicate(ctx context.Context, masterPlanID string, id string) (*model.MasterPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.MasterPlanMilestoneDuplicate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("masterPlanID", masterPlanID), log.String("id", id))

	masterPlanMilestoneId, err := r.Services.MasterPlanService.DuplicateMasterPlanMilestone(ctx, masterPlanID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to duplicate master plan milestone")
		return &model.MasterPlanMilestone{ID: masterPlanMilestoneId}, err
	}

	createdMasterPlanMilestoneEntity, err := r.Services.MasterPlanService.GetMasterPlanMilestoneById(ctx, masterPlanMilestoneId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Master plan milestone details not yet available. Master plan milestone id: %s", masterPlanMilestoneId)
		return &model.MasterPlanMilestone{ID: masterPlanMilestoneId}, nil
	}
	span.LogFields(log.String("response.masterPlanMilestoneId", masterPlanMilestoneId))
	return mapper.MapEntityToMasterPlanMilestone(createdMasterPlanMilestoneEntity), nil
}

// MasterPlan is the resolver for the masterPlan field.
func (r *queryResolver) MasterPlan(ctx context.Context, id string) (*model.MasterPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.MasterPlan", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.SetTag(tracing.SpanTagEntityId, id)

	masterPlanEntity, err := r.Services.MasterPlanService.GetMasterPlanById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Master plan with id %s", id)
		return nil, nil
	}
	if masterPlanEntity == nil {
		graphql.AddErrorf(ctx, "Master plan with id %s not found", id)
		return nil, nil
	}
	return mapper.MapEntityToMasterPlan(masterPlanEntity), nil
}

// MasterPlans is the resolver for the masterPlans field.
func (r *queryResolver) MasterPlans(ctx context.Context, retired *bool) ([]*model.MasterPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.MasterPlans", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	masterPlanEntities, err := r.Services.MasterPlanService.GetMasterPlans(ctx, retired)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Master plans")
		return nil, nil
	}
	return mapper.MapEntitiesToMasterPlans(masterPlanEntities), nil
}

// MasterPlan returns generated.MasterPlanResolver implementation.
func (r *Resolver) MasterPlan() generated.MasterPlanResolver { return &masterPlanResolver{r} }

type masterPlanResolver struct{ *Resolver }
