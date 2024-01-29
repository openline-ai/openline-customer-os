package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// InvoiceLines is the resolver for the invoiceLines field.
func (r *invoiceResolver) InvoiceLines(ctx context.Context, obj *model.Invoice) ([]*model.InvoiceLine, error) {
	panic(fmt.Errorf("not implemented: InvoiceLines - invoiceLines"))
}

// InvoiceNextDryRunForContract is the resolver for the invoice_NextDryRunForContract field.
func (r *mutationResolver) InvoiceNextDryRunForContract(ctx context.Context, contractID string) (string, error) {
	panic(fmt.Errorf("not implemented: InvoiceNextDryRunForContract - invoice_NextDryRunForContract"))
}

// InvoiceSimulate is the resolver for the invoice_Simulate field.
func (r *mutationResolver) InvoiceSimulate(ctx context.Context, input model.InvoiceSimulateInput) (string, error) {
	panic(fmt.Errorf("not implemented: InvoiceSimulate - invoice_Simulate"))
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
