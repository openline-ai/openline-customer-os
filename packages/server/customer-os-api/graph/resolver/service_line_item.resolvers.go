package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// ServiceLineItemCreate is the resolver for the serviceLineItem_Create field.
func (r *mutationResolver) ServiceLineItemCreate(ctx context.Context, input model.ServiceLineItemInput) (*model.ServiceLineItem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ServiceLineItemCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	serviceLineItemId, err := r.Services.ServiceLineItemService.Create(ctx, &service.ServiceLineItemCreateData{
		ServiceLineItemEntity: mapper.MapServiceLineItemInputToEntity(input),
		ContractId:            input.ContractID,
		ExternalReference:     mapper.MapExternalSystemReferenceInputToRelationship(input.ExternalReference),
		Source:                entity.DataSourceOpenline,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create service line item")
		return &model.ServiceLineItem{ID: serviceLineItemId}, err
	}
	createdServiceLineItemEntity, err := r.Services.ServiceLineItemService.GetById(ctx, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Service line item details not yet available. Service line item id: %s", serviceLineItemId)
		return &model.ServiceLineItem{ID: serviceLineItemId}, nil
	}
	span.LogFields(log.String("response.serviceLineItemID", serviceLineItemId))
	return mapper.MapEntityToServiceLineItem(createdServiceLineItemEntity), nil
}

// ServiceLineItemUpdate is the resolver for the serviceLineItemUpdate field.
func (r *mutationResolver) ServiceLineItemUpdate(ctx context.Context, input model.ServiceLineItemUpdateInput) (*model.ServiceLineItem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ServiceLineItemUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.serviceLineItemId", input.ServiceLineItemID))

	err := r.Services.ServiceLineItemService.Update(ctx, mapper.MapServiceLineItemUpdateInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update service line item %s", input.ServiceLineItemID)
		return &model.ServiceLineItem{ID: input.ServiceLineItemID}, nil
	}
	serviceLineItemEntity, err := r.Services.ServiceLineItemService.GetById(ctx, input.ServiceLineItemID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed fetching service line item details. Service line item id: %s", input.ServiceLineItemID)
		return &model.ServiceLineItem{ID: input.ServiceLineItemID}, nil
	}

	return mapper.MapEntityToServiceLineItem(serviceLineItemEntity), nil
}

// ServiceLineItem is the resolver for the serviceLineItem field.
func (r *queryResolver) ServiceLineItem(ctx context.Context, id string) (*model.ServiceLineItem, error) {
	panic(fmt.Errorf("not implemented: ServiceLineItem - serviceLineItem"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *serviceLineItemResolver) CreatedBy(ctx context.Context, obj *model.ServiceLineItem) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *serviceLineItemResolver) ExternalLinks(ctx context.Context, obj *model.ServiceLineItem) ([]*model.ExternalSystem, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetExternalSystemsForServiceLineItem(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get external system for service line item %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get external system for service line item %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToExternalSystems(entities), nil
}

// ServiceLineItem returns generated.ServiceLineItemResolver implementation.
func (r *Resolver) ServiceLineItem() generated.ServiceLineItemResolver {
	return &serviceLineItemResolver{r}
}

type serviceLineItemResolver struct{ *Resolver }
