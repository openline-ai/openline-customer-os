package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/file-store-api/test/graph/model"
)

// CreatedBy is the resolver for the createdBy field.
func (r *columnDefResolver) CreatedBy(ctx context.Context, obj *model.ColumnDef) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *columnTypeResolver) CreatedBy(ctx context.Context, obj *model.ColumnType) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// TableViewDefs is the resolver for the tableViewDefs field.
func (r *queryResolver) TableViewDefs(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort *model.SortBy) (*model.TableViewDefPage, error) {
	panic(fmt.Errorf("not implemented: TableViewDefs - tableViewDefs"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *tableViewDefResolver) CreatedBy(ctx context.Context, obj *model.TableViewDef) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *viewTypeResolver) CreatedBy(ctx context.Context, obj *model.ViewType) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// ColumnDef returns generated.ColumnDefResolver implementation.
func (r *Resolver) ColumnDef() generated.ColumnDefResolver { return &columnDefResolver{r} }

// ColumnType returns generated.ColumnTypeResolver implementation.
func (r *Resolver) ColumnType() generated.ColumnTypeResolver { return &columnTypeResolver{r} }

// TableViewDef returns generated.TableViewDefResolver implementation.
func (r *Resolver) TableViewDef() generated.TableViewDefResolver { return &tableViewDefResolver{r} }

// ViewType returns generated.ViewTypeResolver implementation.
func (r *Resolver) ViewType() generated.ViewTypeResolver { return &viewTypeResolver{r} }

type columnDefResolver struct{ *Resolver }
type columnTypeResolver struct{ *Resolver }
type tableViewDefResolver struct{ *Resolver }
type viewTypeResolver struct{ *Resolver }
