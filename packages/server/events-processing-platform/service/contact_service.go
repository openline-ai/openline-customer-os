package service

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	commonmodel "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/command_handler"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	grpcerr "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_errors"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	contactpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/contact"
	"strings"
)

type contactService struct {
	contactpb.UnimplementedContactGrpcServiceServer
	log                    logger.Logger
	contactCommandHandlers *command_handler.CommandHandlers
	contactRequestHandler  contact.ContactRequestHandler
}

func NewContactService(log logger.Logger, contactCommandHandlers *command_handler.CommandHandlers, aggregateStore eventstore.AggregateStore, cfg *config.Config) *contactService {
	return &contactService{
		log:                    log,
		contactCommandHandlers: contactCommandHandlers,
		contactRequestHandler:  contact.NewContactRequestHandler(log, aggregateStore, cfg.Utils),
	}
}

func (s *contactService) UpsertContact(ctx context.Context, request *contactpb.UpsertContactGrpcRequest) (*contactpb.ContactIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ContactService.UpsertContact")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	request.Timezone = normalizeTimezone(request.Timezone)

	contactId := utils.NewUUIDIfEmpty(request.Id)

	dataFields := models.ContactDataFields{
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Name:            request.Name,
		Description:     request.Description,
		Prefix:          request.Prefix,
		Timezone:        request.Timezone,
		ProfilePhotoUrl: request.ProfilePhotoUrl,
	}
	sourceFields := commonmodel.Source{}
	sourceFields.FromGrpc(request.SourceFields)
	sourceFields.Source = utils.StringFirstNonEmpty(sourceFields.Source, request.Source)
	sourceFields.SourceOfTruth = utils.StringFirstNonEmpty(sourceFields.SourceOfTruth, request.SourceOfTruth)
	sourceFields.AppSource = utils.StringFirstNonEmpty(sourceFields.AppSource, request.AppSource)

	externalSystem := commonmodel.ExternalSystem{}
	externalSystem.FromGrpc(request.ExternalSystemFields)

	fieldsMask := extractContactFieldsMask(request.FieldsMask)

	cmd := command.NewUpsertContactCommand(contactId, request.Tenant, request.LoggedInUserId, sourceFields, externalSystem,
		dataFields, utils.TimestampProtoToTimePtr(request.CreatedAt), utils.TimestampProtoToTimePtr(request.UpdatedAt), request.Id == "", fieldsMask)
	if err := s.contactCommandHandlers.Upsert.Handle(ctx, cmd); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(UpsertContact.Handle) tenant:%s, contactID: %s, err: {%v}", request.Tenant, contactId, err)
		return nil, s.errResponse(err)
	}

	return &contactpb.ContactIdGrpcResponse{Id: contactId}, nil
}

func extractContactFieldsMask(fields []contactpb.ContactFieldMask) []string {
	fieldsMask := make([]string, 0)
	if len(fields) == 0 {
		return fieldsMask
	}
	for _, field := range fields {
		switch field {
		case contactpb.ContactFieldMask_CONTACT_FIELD_FIRST_NAME:
			fieldsMask = append(fieldsMask, event.FieldMaskFirstName)
		case contactpb.ContactFieldMask_CONTACT_FIELD_LAST_NAME:
			fieldsMask = append(fieldsMask, event.FieldMaskLastName)
		case contactpb.ContactFieldMask_CONTACT_FIELD_NAME:
			fieldsMask = append(fieldsMask, event.FieldMaskName)
		case contactpb.ContactFieldMask_CONTACT_FIELD_PREFIX:
			fieldsMask = append(fieldsMask, event.FieldMaskPrefix)
		case contactpb.ContactFieldMask_CONTACT_FIELD_DESCRIPTION:
			fieldsMask = append(fieldsMask, event.FieldMaskDescription)
		case contactpb.ContactFieldMask_CONTACT_FIELD_TIMEZONE:
			fieldsMask = append(fieldsMask, event.FieldMaskTimezone)
		case contactpb.ContactFieldMask_CONTACT_FIELD_PROFILE_PHOTO_URL:
			fieldsMask = append(fieldsMask, event.FieldMaskProfilePhotoUrl)
		}
	}
	return utils.RemoveDuplicates(fieldsMask)
}

func (s *contactService) LinkPhoneNumberToContact(ctx context.Context, request *contactpb.LinkPhoneNumberToContactGrpcRequest) (*contactpb.ContactIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ContactService.LinkPhoneNumberToContact")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	cmd := command.NewLinkPhoneNumberCommand(request.ContactId, request.Tenant, request.LoggedInUserId, request.PhoneNumberId, request.Label, request.AppSource, request.Primary)
	if err := s.contactCommandHandlers.LinkPhoneNumber.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkPhoneNumberCommand.Handle) tenant:{%s}, contact ID: {%s}, err: {%v}", request.Tenant, request.ContactId, err.Error())
		return nil, s.errResponse(err)
	}

	return &contactpb.ContactIdGrpcResponse{Id: request.ContactId}, nil
}

func (s *contactService) LinkEmailToContact(ctx context.Context, request *contactpb.LinkEmailToContactGrpcRequest) (*contactpb.ContactIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ContactService.LinkEmailToContact")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	cmd := command.NewLinkEmailCommand(request.ContactId, request.Tenant, request.LoggedInUserId, request.EmailId, request.Label, request.AppSource, request.Primary)
	if err := s.contactCommandHandlers.LinkEmail.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkEmailCommand.Handle) tenant:{%s}, contact ID: {%s}, err: {%v}", request.Tenant, request.ContactId, err.Error())
		return nil, s.errResponse(err)
	}

	return &contactpb.ContactIdGrpcResponse{Id: request.ContactId}, nil
}

func (s *contactService) LinkLocationToContact(ctx context.Context, request *contactpb.LinkLocationToContactGrpcRequest) (*contactpb.ContactIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ContactService.LinkLocationToContact")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	cmd := command.NewLinkLocationCommand(request.ContactId, request.Tenant, request.LoggedInUserId, request.LocationId, request.AppSource)
	if err := s.contactCommandHandlers.LinkLocation.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkLocationCommand.Handle) tenant:{%s}, contact ID: {%s}, err: {%v}", request.Tenant, request.ContactId, err.Error())
		return nil, s.errResponse(err)
	}

	return &contactpb.ContactIdGrpcResponse{Id: request.ContactId}, nil
}

func (s *contactService) LinkWithOrganization(ctx context.Context, request *contactpb.LinkWithOrganizationGrpcRequest) (*contactpb.ContactIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ContactService.LinkWithOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)

	sourceFields := commonmodel.Source{}
	sourceFields.FromGrpc(request.SourceFields)

	jobRoleFields := models.JobRole{
		JobTitle:    request.JobTitle,
		Description: request.Description,
		Primary:     request.Primary,
		StartedAt:   utils.TimestampProtoToTimePtr(request.StartedAt),
		EndedAt:     utils.TimestampProtoToTimePtr(request.EndedAt),
	}

	cmd := command.NewLinkOrganizationCommand(request.ContactId, request.Tenant, request.LoggedInUserId, request.OrganizationId, sourceFields, jobRoleFields,
		utils.TimestampProtoToTimePtr(request.CreatedAt), utils.TimestampProtoToTimePtr(request.UpdatedAt))
	if err := s.contactCommandHandlers.LinkOrganization.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkOrganizationCommand.Handle) tenant:{%s}, contact ID: {%s}, err: {%v}", request.Tenant, request.ContactId, err.Error())
		return nil, s.errResponse(err)
	}

	return &contactpb.ContactIdGrpcResponse{Id: request.ContactId}, nil
}

func (s *contactService) AddSocial(ctx context.Context, request *contactpb.ContactAddSocialGrpcRequest) (*contactpb.SocialIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "ContactService.AddSocial")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	tracing.LogObjectAsJson(span, "request", request)
	span.SetTag(tracing.SpanTagEntityId, request.ContactId)

	socialIdAny, err := s.contactRequestHandler.HandleWithRetry(ctx, request.Tenant, request.ContactId, true, request)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(AddSocial.Handle) tenant:{%v}, err: %v", request.Tenant, err.Error())
		return nil, grpcerr.ErrResponse(err)
	}
	socialId := ""
	if socialIdAny != nil {
		socialId = socialIdAny.(string)
	}

	return &contactpb.SocialIdGrpcResponse{Id: socialId}, nil
}

func (s *contactService) errResponse(err error) error {
	return grpcerr.ErrResponse(err)
}

func normalizeTimezone(timezone string) string {
	if timezone == "" {
		return ""
	}
	output := strings.Replace(timezone, "_slash_", "/", -1)
	output = utils.CapitalizeAllParts(output, []string{"/", "_"})
	return output
}
