package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jenum "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// ContractLineItemCreate is the resolver for the contractLineItem_Create field.
func (r *mutationResolver) ContractLineItemCreate(ctx context.Context, input model.ServiceLineItemInput) (*model.ServiceLineItem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ContractLineItemCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "request.input", input)

	billedType := neo4jenum.BilledTypeNone
	if input.Billed != nil {
		billedType = mapper.MapBilledTypeFromModel(*input.Billed)
	}

	data := service.ServiceLineItemCreateData{
		ContractId:        input.ContractID,
		ExternalReference: mapper.MapExternalSystemReferenceInputToRelationship(input.ExternalReference),
		Source:            neo4jentity.DataSourceOpenline,
		StartedAt:         input.StartedAt,
		EndedAt:           input.EndedAt,
		SliBilledType:     billedType,
		SliName:           utils.IfNotNilString(input.Name),
		SliPrice:          utils.IfNotNilFloat64(input.Price),
		SliQuantity:       utils.IfNotNilInt64(input.Quantity),
		SliVatRate:        utils.IfNotNilFloat64(input.VatRate),
	}
	if input.ServiceStarted != nil {
		data.StartedAt = input.ServiceStarted
	}
	if input.ServiceEnded != nil {
		data.EndedAt = input.ServiceEnded
	}
	if input.BillingCycle != nil {
		data.SliBilledType = mapper.MapBilledTypeFromModel(*input.BillingCycle)
	}
	if input.Description != nil {
		data.SliName = *input.Description
	}
	if input.Tax != nil {
		data.SliVatRate = (*input.Tax).TaxRate
	}

	serviceLineItemId, err := r.Services.ServiceLineItemService.Create(ctx, data)
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

// ContractLineItemUpdate is the resolver for the contractLineItem_Update field.
func (r *mutationResolver) ContractLineItemUpdate(ctx context.Context, input model.ServiceLineItemUpdateInput) (*model.ServiceLineItem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ContractLineItemUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "request.input", input)

	billedType := neo4jenum.BilledTypeNone
	if input.Billed != nil {
		billedType = mapper.MapBilledTypeFromModel(*input.Billed)
	}

	data := service.ServiceLineItemUpdateData{
		Id:                      input.ServiceLineItemID,
		IsRetroactiveCorrection: utils.IfNotNilBool(input.IsRetroactiveCorrection),
		SliName:                 utils.IfNotNilString(input.Name),
		SliPrice:                utils.IfNotNilFloat64(input.Price),
		SliQuantity:             utils.IfNotNilInt64(input.Quantity),
		SliBilledType:           billedType,
		SliComments:             utils.IfNotNilString(input.Comments),
		SliVatRate:              utils.IfNotNilFloat64(input.VatRate),
		Source:                  neo4jentity.DataSourceOpenline,
		AppSource:               utils.IfNotNilString(input.AppSource),
		StartedAt:               input.ServiceStarted,
	}
	if input.ID != nil {
		data.Id = *input.ID
	}
	if input.Tax != nil {
		data.SliVatRate = (*input.Tax).TaxRate
	}
	if input.BillingCycle != nil {
		data.SliBilledType = mapper.MapBilledTypeFromModel(*input.BillingCycle)
	}
	if input.Description != nil {
		data.SliName = *input.Description
	}

	err := r.Services.ServiceLineItemService.Update(ctx, data)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to update contract line item {%s}", input.ServiceLineItemID)
		return &model.ServiceLineItem{ID: input.ServiceLineItemID}, err
	}

	serviceLineItemEntity, err := r.Services.ServiceLineItemService.GetById(ctx, input.ServiceLineItemID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed fetching contract line item details. Contract line item id: {%s}", input.ServiceLineItemID)
		return &model.ServiceLineItem{ID: input.ServiceLineItemID}, nil
	}

	return mapper.MapEntityToServiceLineItem(serviceLineItemEntity), nil
}

// ContractLineItemClose is the resolver for the contractLineItem_Close field.
func (r *mutationResolver) ContractLineItemClose(ctx context.Context, input model.ServiceLineItemCloseInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ContractLineItemClose", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "request.input", input)

	endedAt := input.EndedAt
	if input.ServiceEnded != nil {
		endedAt = input.ServiceEnded
	}
	err := r.Services.ServiceLineItemService.Close(ctx, input.ID, endedAt)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to close service line item %s", input.ID)
		return input.ID, nil
	}
	return input.ID, nil
}

// ServiceLineItemDelete is the resolver for the serviceLineItem_Delete field.
func (r *mutationResolver) ServiceLineItemDelete(ctx context.Context, id string) (*model.DeleteResponse, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ServiceLineItemDelete", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.id", id))

	deletionCompleted, err := r.Services.ServiceLineItemService.Delete(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to delete service line item %s", id)
		return &model.DeleteResponse{Accepted: false, Completed: false}, nil
	}
	return &model.DeleteResponse{Accepted: true, Completed: deletionCompleted}, nil
}

// ServiceLineItemBulkUpdate is the resolver for the serviceLineItemBulkUpdate field.
func (r *mutationResolver) ServiceLineItemBulkUpdate(ctx context.Context, input model.ServiceLineItemBulkUpdateInput) ([]string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ServiceLineItemBulkUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "request.input", input)

	sliEntitiesForBulkSync := service.MapServiceLineItemBulkItemsToData(input.ServiceLineItems)

	updatedServiceLineItemIds, err := r.Services.ServiceLineItemService.CreateOrUpdateOrCloseInBulk(ctx, input.ContractID, sliEntitiesForBulkSync)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to bulk update service line items")
		return nil, nil
	}

	if input.InvoiceNote != nil && input.ContractID != "" {
		err = r.Services.ContractService.Update(ctx, model.ContractUpdateInput{
			Patch:       utils.ToPtr(true),
			ContractID:  input.ContractID,
			InvoiceNote: input.InvoiceNote,
		})
		if err != nil {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "failed to update contract invoice note")
			return nil, nil
		}
	}

	return updatedServiceLineItemIds, nil
}

// ServiceLineItemClose is the resolver for the serviceLineItem_Close field.
func (r *mutationResolver) ServiceLineItemClose(ctx context.Context, input model.ServiceLineItemCloseInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ServiceLineItemClose", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	tracing.LogObjectAsJson(span, "request.input", input)

	endedAt := input.EndedAt
	if input.ServiceEnded != nil {
		endedAt = input.ServiceEnded
	}
	err := r.Services.ServiceLineItemService.Close(ctx, input.ID, endedAt)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "failed to close service line item %s", input.ID)
		return input.ID, nil
	}
	return input.ID, nil
}

// ServiceLineItemCreate is the resolver for the serviceLineItem_Create field.
func (r *mutationResolver) ServiceLineItemCreate(ctx context.Context, input model.ServiceLineItemInput) (*model.ServiceLineItem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ServiceLineItemCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	billedType := neo4jenum.BilledTypeNone
	if input.Billed != nil {
		billedType = mapper.MapBilledTypeFromModel(*input.Billed)
	}

	data := service.ServiceLineItemCreateData{
		ContractId:        input.ContractID,
		ExternalReference: mapper.MapExternalSystemReferenceInputToRelationship(input.ExternalReference),
		Source:            neo4jentity.DataSourceOpenline,
		StartedAt:         input.StartedAt,
		EndedAt:           input.EndedAt,
		SliBilledType:     billedType,
		SliName:           utils.IfNotNilString(input.Name),
		SliPrice:          utils.IfNotNilFloat64(input.Price),
		SliQuantity:       utils.IfNotNilInt64(input.Quantity),
		SliVatRate:        utils.IfNotNilFloat64(input.VatRate),
	}
	if input.ServiceStarted != nil {
		data.StartedAt = input.ServiceStarted
	}
	if input.ServiceEnded != nil {
		data.EndedAt = input.ServiceEnded
	}
	if input.BillingCycle != nil {
		data.SliBilledType = mapper.MapBilledTypeFromModel(*input.BillingCycle)
	}
	if input.Description != nil {
		data.SliName = *input.Description
	}
	if input.Tax != nil {
		data.SliVatRate = (*input.Tax).TaxRate
	}

	serviceLineItemId, err := r.Services.ServiceLineItemService.Create(ctx, data)
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

	billedType := neo4jenum.BilledTypeNone
	if input.Billed != nil {
		billedType = mapper.MapBilledTypeFromModel(*input.Billed)
	}

	data := service.ServiceLineItemUpdateData{
		Id:                      input.ServiceLineItemID,
		IsRetroactiveCorrection: utils.IfNotNilBool(input.IsRetroactiveCorrection),
		SliName:                 utils.IfNotNilString(input.Name),
		SliPrice:                utils.IfNotNilFloat64(input.Price),
		SliQuantity:             utils.IfNotNilInt64(input.Quantity),
		SliBilledType:           billedType,
		SliComments:             utils.IfNotNilString(input.Comments),
		SliVatRate:              utils.IfNotNilFloat64(input.VatRate),
		Source:                  neo4jentity.DataSourceOpenline,
		AppSource:               utils.IfNotNilString(input.AppSource),
		StartedAt:               input.ServiceStarted,
	}
	if input.ID != nil {
		data.Id = *input.ID
	}
	if input.Tax != nil {
		data.SliVatRate = (*input.Tax).TaxRate
	}
	if input.BillingCycle != nil {
		data.SliBilledType = mapper.MapBilledTypeFromModel(*input.BillingCycle)
	}
	if input.Description != nil {
		data.SliName = *input.Description
	}

	err := r.Services.ServiceLineItemService.Update(ctx, data)
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
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.ServiceLineItem", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.serviceLineItemID", id))

	if id == "" {
		tracing.TraceErr(span, errors.New("missing service line item input id"))
		graphql.AddErrorf(ctx, "Missing service line item input id")
		return nil, nil
	}

	serviceLineItemEntityPtr, err := r.Services.ServiceLineItemService.GetById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get service line item by id %s", id)
		return nil, err
	}
	return mapper.MapEntityToServiceLineItem(serviceLineItemEntityPtr), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *serviceLineItemResolver) CreatedBy(ctx context.Context, obj *model.ServiceLineItem) (*model.User, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	userEntityNillable, err := dataloader.For(ctx).GetUserCreatorForServiceLineItem(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("error fetching user creator for service line item %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "error fetching user creator for service line item %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToUser(userEntityNillable), nil
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *serviceLineItemResolver) ExternalLinks(ctx context.Context, obj *model.ServiceLineItem) ([]*model.ExternalSystem, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetExternalSystemsForServiceLineItem(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
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
