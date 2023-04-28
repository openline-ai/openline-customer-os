package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// Tags is the resolver for the tags field.
func (r *issueResolver) Tags(ctx context.Context, obj *model.Issue) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// MentionedByNotes is the resolver for the mentionedByNotes field.
func (r *issueResolver) MentionedByNotes(ctx context.Context, obj *model.Issue) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: MentionedByNotes - mentionedByNotes"))
}

// Issue is the resolver for the issue field.
func (r *queryResolver) Issue(ctx context.Context, id string) (*model.Issue, error) {
	panic(fmt.Errorf("not implemented: Issue - issue"))
}

// Issue returns generated.IssueResolver implementation.
func (r *Resolver) Issue() generated.IssueResolver { return &issueResolver{r} }

type issueResolver struct{ *Resolver }
