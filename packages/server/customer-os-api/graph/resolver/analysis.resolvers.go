package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// Describes is the resolver for the describes field.
func (r *analysisResolver) Describes(ctx context.Context, obj *model.Analysis) ([]model.DescriptionNode, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	participantEntities, err := dataloader.For(ctx).GetDescribesForAnalysis(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToDescriptionNodes(participantEntities), nil
}

// AnalysisCreate is the resolver for the analysis_Create field.
func (r *mutationResolver) AnalysisCreate(ctx context.Context, analysis model.AnalysisInput) (*model.Analysis, error) {
	analysisCreated, err := r.Services.AnalysisService.Create(ctx, &service.AnalysisCreateData{
		AnalysisEntity: mapper.MapAnalysisInputToEntity(&analysis),
		Describes:      service.MapAnalysisDescriptionInputToDescriptionData(analysis.Describes),

		Source:        entity.DataSourceOpenline,
		SourceOfTruth: entity.DataSourceOpenline,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create Analysis")
		return nil, err
	}
	newAnalysis := mapper.MapEntityToAnalysis(analysisCreated)
	return newAnalysis, nil
}

// Analysis is the resolver for the analysis field.
func (r *queryResolver) Analysis(ctx context.Context, id string) (*model.Analysis, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	analysis, err := r.Services.AnalysisService.GetAnalysisById(ctx, id)
	if err != nil || analysis == nil {
		graphql.AddErrorf(ctx, "Analysis with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToAnalysis(analysis), nil
}

// Analysis returns generated.AnalysisResolver implementation.
func (r *Resolver) Analysis() generated.AnalysisResolver { return &analysisResolver{r} }

type analysisResolver struct{ *Resolver }
