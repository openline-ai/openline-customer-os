package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api-sdk/graph/model"
)

// CreatedBy is the resolver for the createdBy field.
func (r *logEntryResolver) CreatedBy(ctx context.Context, obj *model.LogEntry) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// Tags is the resolver for the tags field.
func (r *logEntryResolver) Tags(ctx context.Context, obj *model.LogEntry) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *logEntryResolver) ExternalLinks(ctx context.Context, obj *model.LogEntry) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// LogEntryCreateForOrganization is the resolver for the logEntry_CreateForOrganization field.
func (r *mutationResolver) LogEntryCreateForOrganization(ctx context.Context, organizationID string, input model.LogEntryInput) (string, error) {
	panic(fmt.Errorf("not implemented: LogEntryCreateForOrganization - logEntry_CreateForOrganization"))
}

// LogEntryUpdate is the resolver for the logEntry_Update field.
func (r *mutationResolver) LogEntryUpdate(ctx context.Context, id string, input model.LogEntryUpdateInput) (string, error) {
	panic(fmt.Errorf("not implemented: LogEntryUpdate - logEntry_Update"))
}

// LogEntryAddTag is the resolver for the logEntry_AddTag field.
func (r *mutationResolver) LogEntryAddTag(ctx context.Context, id string, input model.TagIDOrNameInput) (string, error) {
	panic(fmt.Errorf("not implemented: LogEntryAddTag - logEntry_AddTag"))
}

// LogEntryRemoveTag is the resolver for the logEntry_RemoveTag field.
func (r *mutationResolver) LogEntryRemoveTag(ctx context.Context, id string, input model.TagIDOrNameInput) (string, error) {
	panic(fmt.Errorf("not implemented: LogEntryRemoveTag - logEntry_RemoveTag"))
}

// LogEntryResetTags is the resolver for the logEntry_ResetTags field.
func (r *mutationResolver) LogEntryResetTags(ctx context.Context, id string, input []*model.TagIDOrNameInput) (string, error) {
	panic(fmt.Errorf("not implemented: LogEntryResetTags - logEntry_ResetTags"))
}

// LogEntry is the resolver for the logEntry field.
func (r *queryResolver) LogEntry(ctx context.Context, id string) (*model.LogEntry, error) {
	panic(fmt.Errorf("not implemented: LogEntry - logEntry"))
}

// LogEntry returns generated.LogEntryResolver implementation.
func (r *Resolver) LogEntry() generated.LogEntryResolver { return &logEntryResolver{r} }

type logEntryResolver struct{ *Resolver }
