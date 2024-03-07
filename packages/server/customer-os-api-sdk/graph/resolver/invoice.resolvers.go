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

// Organization is the resolver for the organization field.
func (r *invoiceResolver) Organization(ctx context.Context, obj *model.Invoice) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Contract is the resolver for the contract field.
func (r *invoiceResolver) Contract(ctx context.Context, obj *model.Invoice) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: Contract - contract"))
}

// InvoiceLineItems is the resolver for the invoiceLineItems field.
func (r *invoiceResolver) InvoiceLineItems(ctx context.Context, obj *model.Invoice) ([]*model.InvoiceLine, error) {
	panic(fmt.Errorf("not implemented: InvoiceLineItems - invoiceLineItems"))
}

// InvoiceNextDryRunForContract is the resolver for the invoice_NextDryRunForContract field.
func (r *mutationResolver) InvoiceNextDryRunForContract(ctx context.Context, contractID string) (string, error) {
	panic(fmt.Errorf("not implemented: InvoiceNextDryRunForContract - invoice_NextDryRunForContract"))
}

// InvoiceSimulate is the resolver for the invoice_Simulate field.
func (r *mutationResolver) InvoiceSimulate(ctx context.Context, input model.InvoiceSimulateInput) (string, error) {
	panic(fmt.Errorf("not implemented: InvoiceSimulate - invoice_Simulate"))
}

// InvoiceUpdate is the resolver for the invoice_Update field.
func (r *mutationResolver) InvoiceUpdate(ctx context.Context, input model.InvoiceUpdateInput) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: InvoiceUpdate - invoice_Update"))
}

// InvoiceVoid is the resolver for the invoice_Void field.
func (r *mutationResolver) InvoiceVoid(ctx context.Context, id string) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: InvoiceVoid - invoice_Void"))
}

// Invoice is the resolver for the invoice field.
func (r *queryResolver) Invoice(ctx context.Context, id string) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: Invoice - invoice"))
}

// Invoices is the resolver for the invoices field.
func (r *queryResolver) Invoices(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy, organizationID *string) (*model.InvoicesPage, error) {
	panic(fmt.Errorf("not implemented: Invoices - invoices"))
}

// Invoice returns generated.InvoiceResolver implementation.
func (r *Resolver) Invoice() generated.InvoiceResolver { return &invoiceResolver{r} }

type invoiceResolver struct{ *Resolver }
