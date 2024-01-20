package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

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
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// ServiceLineItems is the resolver for the serviceLineItems field.
func (r *contractResolver) ServiceLineItems(ctx context.Context, obj *model.Contract) ([]*model.ServiceLineItem, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	serviceLineItemEntities, err := dataloader.For(ctx).GetServiceLineItemsForContract(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get service line items for contract %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get service line items for contract %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToServiceLineItems(serviceLineItemEntities), nil
}

// Opportunities is the resolver for the opportunities field.
func (r *contractResolver) Opportunities(ctx context.Context, obj *model.Contract) ([]*model.Opportunity, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	opportunityEntities, err := dataloader.For(ctx).GetOpportunitiesForContract(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get opportunities for contract %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get opportunities for contract %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToOpportunities(opportunityEntities), nil
}

// Owner is the resolver for the owner field.
func (r *contractResolver) Owner(ctx context.Context, obj *model.Contract) (*model.User, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "ContractResolver.Owner", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contractID", obj.ID))

	owner, err := r.Services.UserService.GetContactOwner(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get owner for contact %s", obj.ID)
		return nil, err
	}
	if owner == nil {
		return nil, nil
	}
	return mapper.MapEntityToUser(owner), err
}

// CreatedBy is the resolver for the createdBy field.
func (r *contractResolver) CreatedBy(ctx context.Context, obj *model.Contract) (*model.User, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	userEntityNillable, err := dataloader.For(ctx).GetUserCreatorForContract(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("error fetching user creator for service line item %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "error fetching user creator for service line item %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToUser(userEntityNillable), nil
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *contractResolver) ExternalLinks(ctx context.Context, obj *model.Contract) ([]*model.ExternalSystem, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetExternalSystemsForContract(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get external system for contract %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get external system for contract %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToExternalSystems(entities), nil
}

// ContractCreate is the resolver for the contract_Create field.
func (r *mutationResolver) ContractCreate(ctx context.Context, input model.ContractInput) (*model.Contract, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ContractCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	contractId, err := r.Services.ContractService.Create(ctx, &service.ContractCreateData{
		ContractEntity:    mapper.MapContractInputToEntity(input),
		OrganizationId:    input.OrganizationID,
		ExternalReference: mapper.MapExternalSystemReferenceInputToRelationship(input.ExternalReference),
		Source:            neo4jentity.DataSourceOpenline,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create contract")
		return &model.Contract{ID: contractId}, err
	}
	createdContractEntity, err := r.Services.ContractService.GetById(ctx, contractId)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Contract details not yet available. Contract id: %s", contractId)
		return &model.Contract{ID: contractId}, nil
	}
	span.LogFields(log.String("response.contractID", contractId))
	return mapper.MapEntityToContract(createdContractEntity), nil
}

// ContractUpdate is the resolver for the contract_Update field.
func (r *mutationResolver) ContractUpdate(ctx context.Context, input model.ContractUpdateInput) (*model.Contract, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.ContractUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contractId", input.ContractID))

	err := r.Services.ContractService.Update(ctx, mapper.MapContractUpdateInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update contract %s", input.ContractID)
		return &model.Contract{ID: input.ContractID}, nil
	}
	contractEntity, err := r.Services.ContractService.GetById(ctx, input.ContractID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed fetching contract details. Contract id: %s", input.ContractID)
		return &model.Contract{ID: input.ContractID}, nil
	}

	return mapper.MapEntityToContract(contractEntity), nil
}

// Contract is the resolver for the contract field.
func (r *queryResolver) Contract(ctx context.Context, id string) (*model.Contract, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Contract", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.contractID", id))

	if id == "" {
		tracing.TraceErr(span, errors.New("Missing contract input id"))
		graphql.AddErrorf(ctx, "Missing contract input id")
		return nil, nil
	}

	contractEntityPtr, err := r.Services.ContractService.GetById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get contract by id %s", id)
		return nil, err
	}
	return mapper.MapEntityToContract(contractEntityPtr), nil
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }
