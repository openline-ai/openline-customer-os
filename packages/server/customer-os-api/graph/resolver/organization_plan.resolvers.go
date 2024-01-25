package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"errors"
	"time"

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

// OrganizationPlanCreate is the resolver for the organizationPlan_Create field.
func (r *mutationResolver) OrganizationPlanCreate(ctx context.Context, input model.OrganizationPlanInput) (*model.OrganizationPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)
	// Create empty organization plan
	orgPlanId, err := r.Services.OrganizationPlanService.CreateOrganizationPlan(ctx, utils.IfNotNilString(input.Name), *input.MasterPlanID, input.OrganizationID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create organization plan")
		return &model.OrganizationPlan{ID: orgPlanId}, err
	}
	// get master plan milestones
	masterPlanMilestonesEntities, err := r.Services.MasterPlanService.GetMasterPlanMilestonesForMasterPlans(ctx, []string{*input.MasterPlanID})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get master plan milestones")
		return &model.OrganizationPlan{ID: orgPlanId}, err
	}
	// create organization plan milestones based on each master plan milestone
	for _, masterPlanMilestoneEntity := range *masterPlanMilestonesEntities {
		msDueDate := time.Now().UTC().Add(time.Hour * time.Duration(masterPlanMilestoneEntity.DurationHours))
		_, err := r.Services.OrganizationPlanService.CreateOrganizationPlanMilestone(ctx, orgPlanId, masterPlanMilestoneEntity.Name, &masterPlanMilestoneEntity.Order, &msDueDate, masterPlanMilestoneEntity.Optional, masterPlanMilestoneEntity.Items)
		if err != nil {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed to create organization plan milestone")
			return &model.OrganizationPlan{ID: orgPlanId}, err
		}
	}
	// get created organization plan
	createdOrgPlanEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanById(ctx, orgPlanId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Organization plan details not yet available. Organization plan id: %s", orgPlanId)
		return &model.OrganizationPlan{ID: orgPlanId}, nil
	}
	// return created organization plan
	span.LogFields(log.String("response.organizationPlanId", orgPlanId))
	return mapper.MapEntityToOrganizationPlan(createdOrgPlanEntity), nil
}

// OrganizationPlanUpdate is the resolver for the organizationPlan_Update field.
func (r *mutationResolver) OrganizationPlanUpdate(ctx context.Context, input model.OrganizationPlanUpdateInput) (*model.OrganizationPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.OrganizationPlanService.UpdateOrganizationPlan(ctx, input.ID, input.Name, input.Retired, input.StatusDetails)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update organization plan")
		return nil, err
	}

	updatedOrgPlanEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanById(ctx, input.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Organization plan with id %s", input.ID)
		return nil, nil
	}
	return mapper.MapEntityToOrganizationPlan(updatedOrgPlanEntity), nil
}

// OrganizationPlanDuplicate is the resolver for the organizationPlan_Duplicate field.
func (r *mutationResolver) OrganizationPlanDuplicate(ctx context.Context, id string) (*model.OrganizationPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanDuplicate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "id", id)

	opId, err := r.Services.OrganizationPlanService.DuplicateOrganizationPlan(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to duplicate org plan")
		return nil, err
	}

	createdOrgPlanEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanById(ctx, opId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Organization plan details not yet available. Organization plan id: %s", opId)
		return &model.OrganizationPlan{ID: opId}, nil
	}
	span.LogFields(log.String("response.organizationPlanId", opId))
	return mapper.MapEntityToOrganizationPlan(createdOrgPlanEntity), nil
}

// OrganizationPlanMilestoneCreate is the resolver for the organizationPlanMilestone_Create field.
func (r *mutationResolver) OrganizationPlanMilestoneCreate(ctx context.Context, input model.OrganizationPlanMilestoneInput) (*model.OrganizationPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanMilestoneCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	opMilestoneId, err := r.Services.OrganizationPlanService.CreateOrganizationPlanMilestone(ctx, input.OrganizationPlanID, utils.IfNotNilString(input.Name),
		&input.Order, &input.DueDate, input.Optional, input.Items)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create organization plan milestone")
		return &model.OrganizationPlanMilestone{ID: opMilestoneId}, err
	}

	createdMilestoneEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanMilestoneById(ctx, opMilestoneId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Organization plan milestone details not yet available. Organization plan milestone id: %s", opMilestoneId)
		return &model.OrganizationPlanMilestone{ID: opMilestoneId}, nil
	}
	span.LogFields(log.String("response.organizationPlanMilestoneId", opMilestoneId))
	return mapper.MapEntityToOrganizationPlanMilestone(createdMilestoneEntity), nil
}

// OrganizationPlanMilestoneUpdate is the resolver for the organizationPlanMilestone_Update field.
func (r *mutationResolver) OrganizationPlanMilestoneUpdate(ctx context.Context, input model.OrganizationPlanMilestoneUpdateInput) (*model.OrganizationPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanMilestoneUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.OrganizationPlanService.UpdateOrganizationPlanMilestone(ctx, input.OrganizationPlanID, input.ID, input.Name,
		input.Order, &input.DueDate, input.Items, input.Optional, input.Retired, input.StatusDetails)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update organization plan milestone")
		return nil, err
	}

	updatedMilestoneEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanMilestoneById(ctx, input.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get organization plan milestone with id %s", input.ID)
		return nil, nil
	}
	return mapper.MapEntityToOrganizationPlanMilestone(updatedMilestoneEntity), nil
}

// OrganizationPlanMilestoneBulkUpdate is the resolver for the organizationPlanMilestone_BulkUpdate field.
func (r *mutationResolver) OrganizationPlanMilestoneBulkUpdate(ctx context.Context, input []*model.OrganizationPlanMilestoneUpdateInput) ([]*model.OrganizationPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanMilestoneBulkUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)
	var updatedOrgPlanMilestoneEntities []*model.OrganizationPlanMilestone
	var err error
	for _, opms := range input {
		err = r.Services.OrganizationPlanService.UpdateOrganizationPlanMilestone(ctx, opms.OrganizationPlanID, opms.ID, opms.Name,
			opms.Order, &opms.DueDate, opms.Items, opms.Optional, opms.Retired, opms.StatusDetails)
		if err != nil {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Failed to update org plan milestone")
			err = errors.Join(err)
		}

		updatedOPMilestoneEntity, errr := r.Services.OrganizationPlanService.GetOrganizationPlanMilestoneById(ctx, opms.ID)
		if errr != nil {
			tracing.TraceErr(span, errr)
			graphql.AddErrorf(ctx, "Failed to get org plan milestone with id %s", opms.ID)
			err = errors.Join(errr)
		}
		updatedOrgPlanMilestoneEntities = append(updatedOrgPlanMilestoneEntities, mapper.MapEntityToOrganizationPlanMilestone(updatedOPMilestoneEntity))
	}
	if err != nil {
		return nil, err
	}
	return updatedOrgPlanMilestoneEntities, nil
}

// OrganizationPlanMilestoneReorder is the resolver for the organizationPlanMilestone_Reorder field.
func (r *mutationResolver) OrganizationPlanMilestoneReorder(ctx context.Context, input model.OrganizationPlanMilestoneReorderInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanMilestoneReorder", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "input", input)

	err := r.Services.OrganizationPlanService.ReorderOrganizationPlanMilestones(ctx, input.OrganizationPlanID, input.OrderedIds)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to reorder org plan milestones")
		return "", err
	}
	return input.OrganizationPlanID, nil
}

// OrganizationPlanMilestoneDuplicate is the resolver for the organizationPlanMilestone_Duplicate field.
func (r *mutationResolver) OrganizationPlanMilestoneDuplicate(ctx context.Context, organizationPlanID string, id string) (*model.OrganizationPlanMilestone, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.OrganizationPlanMilestoneDuplicate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("organizationPlanID", organizationPlanID), log.String("id", id))

	opMilestoneId, err := r.Services.OrganizationPlanService.DuplicateOrganizationPlanMilestone(ctx, organizationPlanID, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to duplicate org plan milestone")
		return &model.OrganizationPlanMilestone{ID: opMilestoneId}, err
	}

	createdMilestoneEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanMilestoneById(ctx, opMilestoneId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Org plan milestone details not yet available. Org plan milestone id: %s", opMilestoneId)
		return &model.OrganizationPlanMilestone{ID: opMilestoneId}, nil
	}
	span.LogFields(log.String("response.OrganizationPlanMilestoneId", opMilestoneId))
	return mapper.MapEntityToOrganizationPlanMilestone(createdMilestoneEntity), nil
}

// Milestones is the resolver for the milestones field.
func (r *organizationPlanResolver) Milestones(ctx context.Context, obj *model.OrganizationPlan) ([]*model.OrganizationPlanMilestone, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	orgPlanMilestonesEntities, err := dataloader.For(ctx).GetOrganizationPlanMilestonesForOrganizationPlan(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get org plan milestones for org plan %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get org plan milestones for org plan %s", obj.ID)
		return nil, nil
	}
	allOrgPlanMilestones := mapper.MapEntitiesToOrganizationPlanMilestones(orgPlanMilestonesEntities)
	// filter out retired milestones
	var milestones []*model.OrganizationPlanMilestone
	for _, orgPlanMilestone := range allOrgPlanMilestones {
		if !orgPlanMilestone.Retired {
			milestones = append(milestones, orgPlanMilestone)
		}
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(log.Int("response.length", len(milestones)))
	}
	return milestones, nil
}

// RetiredMilestones is the resolver for the retiredMilestones field.
func (r *organizationPlanResolver) RetiredMilestones(ctx context.Context, obj *model.OrganizationPlan) ([]*model.OrganizationPlanMilestone, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	orgPlanMilestonesEntities, err := dataloader.For(ctx).GetOrganizationPlanMilestonesForOrganizationPlan(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get retired org plan milestones for org plan %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get retired org plan milestones for org plan %s", obj.ID)
		return nil, nil
	}
	allOrgPlanMilestones := mapper.MapEntitiesToOrganizationPlanMilestones(orgPlanMilestonesEntities)
	// filter out non-retired milestones
	var milestones []*model.OrganizationPlanMilestone
	for _, orgPlanMilestone := range allOrgPlanMilestones {
		if orgPlanMilestone.Retired {
			milestones = append(milestones, orgPlanMilestone)
		}
	}
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(log.Int("response.length", len(milestones)))
	}
	return milestones, nil
}

// OrganizationPlan is the resolver for the organizationPlan field.
func (r *queryResolver) OrganizationPlan(ctx context.Context, id string) (*model.OrganizationPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.OrganizationPlan", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.SetTag(tracing.SpanTagEntityId, id)

	orgPlanEntity, err := r.Services.OrganizationPlanService.GetOrganizationPlanById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Org plan with id %s", id)
		return nil, nil
	}
	if orgPlanEntity == nil {
		graphql.AddErrorf(ctx, "Org plan with id %s not found", id)
		return nil, nil
	}
	return mapper.MapEntityToOrganizationPlan(orgPlanEntity), nil
}

// OrganizationPlansForOrganization is the resolver for the organizationPlansForOrganization field.
func (r *queryResolver) OrganizationPlansForOrganization(ctx context.Context, organizationID string) ([]*model.OrganizationPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.OrganizationPlansForOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.SetTag(tracing.SpanTagEntityId, organizationID)

	orgPlanEntities, err := r.Services.OrganizationPlanService.GetOrganizationPlansForOrganization(ctx, organizationID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Org plans for organization %s", organizationID)
		return nil, nil
	}
	return mapper.MapEntitiesToOrganizationPlans(orgPlanEntities), nil
}

// OrganizationPlans is the resolver for the organizationPlans field.
func (r *queryResolver) OrganizationPlans(ctx context.Context, retired *bool) ([]*model.OrganizationPlan, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.OrganizationPlans", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	orgPlanEntities, err := r.Services.OrganizationPlanService.GetOrganizationPlans(ctx, retired)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get Organization plans")
		return nil, nil
	}
	return mapper.MapEntitiesToOrganizationPlans(orgPlanEntities), nil
}

// OrganizationPlan returns generated.OrganizationPlanResolver implementation.
func (r *Resolver) OrganizationPlan() generated.OrganizationPlanResolver {
	return &organizationPlanResolver{r}
}

type organizationPlanResolver struct{ *Resolver }
