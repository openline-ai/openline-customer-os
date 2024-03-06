package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/model"
)

// Milestones is the resolver for the milestones field.
func (r *masterPlanResolver) Milestones(ctx context.Context, obj *model.MasterPlan) ([]*model.MasterPlanMilestone, error) {
	panic(fmt.Errorf("not implemented: Milestones - milestones"))
}

// RetiredMilestones is the resolver for the retiredMilestones field.
func (r *masterPlanResolver) RetiredMilestones(ctx context.Context, obj *model.MasterPlan) ([]*model.MasterPlanMilestone, error) {
	panic(fmt.Errorf("not implemented: RetiredMilestones - retiredMilestones"))
}

// MasterPlanCreate is the resolver for the masterPlan_Create field.
func (r *mutationResolver) MasterPlanCreate(ctx context.Context, input model.MasterPlanInput) (*model.MasterPlan, error) {
	panic(fmt.Errorf("not implemented: MasterPlanCreate - masterPlan_Create"))
}

// MasterPlanCreateDefault is the resolver for the masterPlan_CreateDefault field.
func (r *mutationResolver) MasterPlanCreateDefault(ctx context.Context, input model.MasterPlanInput) (*model.MasterPlan, error) {
	panic(fmt.Errorf("not implemented: MasterPlanCreateDefault - masterPlan_CreateDefault"))
}

// MasterPlanUpdate is the resolver for the masterPlan_Update field.
func (r *mutationResolver) MasterPlanUpdate(ctx context.Context, input model.MasterPlanUpdateInput) (*model.MasterPlan, error) {
	panic(fmt.Errorf("not implemented: MasterPlanUpdate - masterPlan_Update"))
}

// MasterPlanDuplicate is the resolver for the masterPlan_Duplicate field.
func (r *mutationResolver) MasterPlanDuplicate(ctx context.Context, id string) (*model.MasterPlan, error) {
	panic(fmt.Errorf("not implemented: MasterPlanDuplicate - masterPlan_Duplicate"))
}

// MasterPlanMilestoneCreate is the resolver for the masterPlanMilestone_Create field.
func (r *mutationResolver) MasterPlanMilestoneCreate(ctx context.Context, input model.MasterPlanMilestoneInput) (*model.MasterPlanMilestone, error) {
	panic(fmt.Errorf("not implemented: MasterPlanMilestoneCreate - masterPlanMilestone_Create"))
}

// MasterPlanMilestoneUpdate is the resolver for the masterPlanMilestone_Update field.
func (r *mutationResolver) MasterPlanMilestoneUpdate(ctx context.Context, input model.MasterPlanMilestoneUpdateInput) (*model.MasterPlanMilestone, error) {
	panic(fmt.Errorf("not implemented: MasterPlanMilestoneUpdate - masterPlanMilestone_Update"))
}

// MasterPlanMilestoneBulkUpdate is the resolver for the masterPlanMilestone_BulkUpdate field.
func (r *mutationResolver) MasterPlanMilestoneBulkUpdate(ctx context.Context, input []*model.MasterPlanMilestoneUpdateInput) ([]*model.MasterPlanMilestone, error) {
	panic(fmt.Errorf("not implemented: MasterPlanMilestoneBulkUpdate - masterPlanMilestone_BulkUpdate"))
}

// MasterPlanMilestoneReorder is the resolver for the masterPlanMilestone_Reorder field.
func (r *mutationResolver) MasterPlanMilestoneReorder(ctx context.Context, input model.MasterPlanMilestoneReorderInput) (string, error) {
	panic(fmt.Errorf("not implemented: MasterPlanMilestoneReorder - masterPlanMilestone_Reorder"))
}

// MasterPlanMilestoneDuplicate is the resolver for the masterPlanMilestone_Duplicate field.
func (r *mutationResolver) MasterPlanMilestoneDuplicate(ctx context.Context, masterPlanID string, id string) (*model.MasterPlanMilestone, error) {
	panic(fmt.Errorf("not implemented: MasterPlanMilestoneDuplicate - masterPlanMilestone_Duplicate"))
}

// MasterPlan is the resolver for the masterPlan field.
func (r *queryResolver) MasterPlan(ctx context.Context, id string) (*model.MasterPlan, error) {
	panic(fmt.Errorf("not implemented: MasterPlan - masterPlan"))
}

// MasterPlans is the resolver for the masterPlans field.
func (r *queryResolver) MasterPlans(ctx context.Context, retired *bool) ([]*model.MasterPlan, error) {
	panic(fmt.Errorf("not implemented: MasterPlans - masterPlans"))
}

// MasterPlan returns generated.MasterPlanResolver implementation.
func (r *Resolver) MasterPlan() generated.MasterPlanResolver { return &masterPlanResolver{r} }

type masterPlanResolver struct{ *Resolver }