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

// SlackChannels is the resolver for the slack_Channels field.
func (r *queryResolver) SlackChannels(ctx context.Context, pagination *model.Pagination) (*model.SlackChannelPage, error) {
	panic(fmt.Errorf("not implemented: SlackChannels - slack_Channels"))
}

// Organization is the resolver for the organization field.
func (r *slackChannelResolver) Organization(ctx context.Context, obj *model.SlackChannel) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// SlackChannel returns generated.SlackChannelResolver implementation.
func (r *Resolver) SlackChannel() generated.SlackChannelResolver { return &slackChannelResolver{r} }

type slackChannelResolver struct{ *Resolver }
