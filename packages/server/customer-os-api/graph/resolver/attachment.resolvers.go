package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// AttachmentCreate is the resolver for the attachment_Create field.
func (r *mutationResolver) AttachmentCreate(ctx context.Context, input model.AttachmentInput) (*model.Attachment, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.AttachmentCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	attachmentCreated, err := r.Services.AttachmentService.Create(ctx, mapper.MapAttachmentInputToEntity(&input), entity.DataSourceOpenline, entity.DataSourceOpenline)

	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create Attachment")
		return nil, err
	}
	newAttachment := mapper.MapEntityToAttachment(attachmentCreated)
	return newAttachment, nil
}

// Attachment is the resolver for the attachment field.
func (r *queryResolver) Attachment(ctx context.Context, id string) (*model.Attachment, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Attachment", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.ID", id))

	analysis, err := r.Services.AttachmentService.GetAttachmentById(ctx, id)
	if err != nil || analysis == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Attachment with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToAttachment(analysis), nil
}
