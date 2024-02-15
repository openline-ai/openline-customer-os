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
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// Organization is the resolver for the organization field.
func (r *invoiceResolver) Organization(ctx context.Context, obj *model.Invoice) (*model.Organization, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	organizationEntity, err := dataloader.For(ctx).GetOrganizationForInvoice(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("error fetching organization for invoice %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Error fetching organization for invoice %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToOrganization(organizationEntity), nil
}

// InvoiceLineItems is the resolver for the invoiceLineItems field.
func (r *invoiceResolver) InvoiceLineItems(ctx context.Context, obj *model.Invoice) ([]*model.InvoiceLine, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetInvoiceLinesForInvoice(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("failed to get invoice lines for invoice %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to fetch invoice lines for invoice %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToInvoiceLines(entities), nil
}

// InvoiceLines is the resolver for the invoiceLines field.
func (r *invoiceResolver) InvoiceLines(ctx context.Context, obj *model.Invoice) ([]*model.InvoiceLine, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetInvoiceLinesForInvoice(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(opentracing.SpanFromContext(ctx), err)
		r.log.Errorf("Failed to get invoice lines for invoice %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get invoice lines for invoice %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToInvoiceLines(entities), nil
}

// InvoiceNextDryRunForContract is the resolver for the invoice_NextDryRunForContract field.
func (r *mutationResolver) InvoiceNextDryRunForContract(ctx context.Context, contractID string) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InvoiceResolver.InvoiceNextInvoiceDryRun", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("contractID", contractID))

	invoiceId, err := r.Services.InvoiceService.NextInvoiceDryRun(ctx, contractID)

	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to dry run next invoice for contract")
		return "", err
	}
	return invoiceId, nil
}

// InvoiceSimulate is the resolver for the invoice_Simulate field.
func (r *mutationResolver) InvoiceSimulate(ctx context.Context, input model.InvoiceSimulateInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InvoiceResolver.InvoiceSimulate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.Object("request.input", input))

	simulateInvoiceData := service.SimulateInvoiceData{
		ContractId: input.ContractID,
		Date:       input.PeriodStartDate,
	}
	for _, invoiceLine := range input.InvoiceLines {
		simulateInvoiceData.InvoiceLines = append(simulateInvoiceData.InvoiceLines, service.SimulateInvoiceLineData{
			ServiceLineItemID: invoiceLine.ServiceLineItemID,
			Name:              invoiceLine.Name,
			Billed:            mapper.MapBilledTypeFromModel(invoiceLine.Billed),
			Price:             invoiceLine.Price,
			Quantity:          invoiceLine.Quantity,
		})
	}

	invoiceId, err := r.Services.InvoiceService.SimulateInvoice(ctx, &simulateInvoiceData)

	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to simulate invoice")
		return "", err
	}
	return invoiceId, nil
}

// Invoice is the resolver for the invoice field.
func (r *queryResolver) Invoice(ctx context.Context, id string) (*model.Invoice, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InvoiceResolver.Invoice", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.invoiceID", id))

	if id == "" {
		tracing.TraceErr(span, errors.New("Missing invoice input id"))
		graphql.AddErrorf(ctx, "Missing invoice input id")
		return nil, nil
	}

	invoiceEntityPtr, err := r.Services.InvoiceService.GetById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get contract by id %s", id)
		return nil, err
	}
	return mapper.MapEntityToInvoice(invoiceEntityPtr), nil
}

// Invoices is the resolver for the invoices field.
func (r *queryResolver) Invoices(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy, organizationID *string) (*model.InvoicesPage, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InvoiceResolver.Invoices", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	if where != nil {
		tracing.LogObjectAsJson(span, "request.where", where)
	}
	if sort != nil {
		tracing.LogObjectAsJson(span, "request.sort", sort)
	}

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	span.LogFields(log.Int("request.pagination.page", pagination.Page), log.Int("request.pagination.limit", pagination.Limit))

	paginatedResult, err := r.Services.InvoiceService.GetInvoices(ctx, utils.IfNotNilString(organizationID), pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get invoices")
		return nil, err
	}
	return &model.InvoicesPage{
		Content:       mapper.MapEntitiesToInvoices(paginatedResult.Rows.(*neo4jentity.InvoiceEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Invoice returns generated.InvoiceResolver implementation.
func (r *Resolver) Invoice() generated.InvoiceResolver { return &invoiceResolver{r} }

type invoiceResolver struct{ *Resolver }
