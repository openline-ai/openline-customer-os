package service

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/generated/proto"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/repository/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type webChatMessageStoreService struct {
	proto.UnimplementedWebChatMessageStoreServiceServer
	driver               *neo4j.Driver
	postgresRepositories *repository.PostgresRepositories
	customerOSService    *customerOSService
}

// sender == contact -> find contact by email
// sender == user -> find user by email

// sender == contact -> find conversation by initiator = contact and channel = webchat
// sender == user -> find conversation by id
func (s *webChatMessageStoreService) SaveMessage(ctx context.Context, input *proto.WebChatInputMessage) (*proto.Message, error) {
	//var err error
	//var conversation *gen.Conversation
	//
	var conversationId string
	var contactId string
	var userId string
	var participantId string

	if input.ConversationId == nil && input.Email == nil {
		return nil, errors.New("conversationId or email must be provided")
	}
	if input.Message == nil && input.Bytes == nil {
		return nil, errors.New("message or bytes must be provided")
	}

	tenant := "openline" //TODO get tenant from context

	if input.ConversationId != nil {
		exists, err := s.customerOSService.ConversationByIdExists(tenant, *input.ConversationId)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("conversation not found")
		}
		conversationId = *input.ConversationId
	}

	if input.SenderType == proto.SenderType_CONTACT {
		contact, err := s.customerOSService.GetContactByEmail(*input.Email)
		if err != nil {
			contactId, err = s.customerOSService.CreateContactWithEmail(tenant, *input.Email)
			if err != nil {
				return nil, err
			}
		} else {
			contactId = contact.Id
		}
		participantId = contactId

		if conversationId == "" {
			convId, err := s.customerOSService.GetWebChatConversationIdWithContactInitiator(tenant, contactId)
			if err != nil {
				return nil, err
			}
			if convId != "" {
				conversationId = convId
			} else {
				convId, err = s.customerOSService.CreateConversation(tenant, participantId, convertSenderTypeToConversationSenderType(input.SenderType), entity.WEB_CHAT)
				if err != nil {
					return nil, err
				}
				conversationId = convId
			}
		}
	} else if input.SenderType == proto.SenderType_USER {
		user, err := s.customerOSService.GetUserByEmail(*input.Email)
		if err != nil {
			return nil, err
		} else {
			userId = user.Id
		}
		participantId = userId

		if conversationId == "" {
			convId, err := s.customerOSService.GetWebChatConversationIdWithContactInitiator(tenant, userId)
			if err != nil {
				return nil, err
			}
			if convId != "" {
				conversationId = convId
			} else {
				convId, err = s.customerOSService.CreateConversation(tenant, userId, convertSenderTypeToConversationSenderType(input.SenderType), entity.WEB_CHAT)
				if err != nil {
					return nil, err
				}
				conversationId = convId
			}
		}
	}

	_, err := s.customerOSService.UpdateConversation(tenant, conversationId, participantId, convertSenderTypeToConversationSenderType(input.SenderType))
	if err != nil {
		return nil, err
	}

	conversationEvent := entity.ConversationEvent{
		TenantId:       tenant,
		ConversationId: conversationId,
		Type:           entity.WEB_CHAT,
		Content:        *input.Message,
		Source:         entity.OPENLINE,
		Direction:      encodeConversationEventDirection(input.Direction),
		CreateDate:     time.Time{},
	}

	if input.GetDirection() == proto.MessageDirection_INBOUND {
		conversationEvent.SenderId = contactId
		conversationEvent.SenderType = entity.CONTACT
	} else {
		conversationEvent.SenderId = userId
		conversationEvent.SenderType = entity.USER
	}

	s.postgresRepositories.ConversationEventRepository.Save(&conversationEvent)

	mi := &proto.Message{
		Id:             conversationEvent.ID,
		ConversationId: conversationId,
		Type:           input.Type,
		Message:        *input.Message,
		Direction:      input.Direction,
		Channel:        proto.MessageChannel_WEB_CHAT,
		SenderType:     input.SenderType,
		SenderId:       participantId,
		Time:           timestamppb.New(time.Now()),
	}
	return mi, nil
}

func convertSenderTypeToConversationSenderType(senderType proto.SenderType) entity.SenderType {
	switch senderType {
	case proto.SenderType_CONTACT:
		return entity.CONTACT
	case proto.SenderType_USER:
		return entity.USER
	default:
		return entity.CONTACT
	}
}

func encodeConversationEventDirection(direction proto.MessageDirection) entity.Direction {
	switch direction {
	case proto.MessageDirection_INBOUND:
		return entity.INBOUND
	case proto.MessageDirection_OUTBOUND:
		return entity.OUTBOUND
	default:
		return entity.OUTBOUND
	}
}

func NewWebChatMessageStoreService(driver *neo4j.Driver, postgresRepositories *repository.PostgresRepositories, customerOSService *customerOSService) *webChatMessageStoreService {
	ms := new(webChatMessageStoreService)
	ms.driver = driver
	ms.postgresRepositories = postgresRepositories
	ms.customerOSService = customerOSService
	return ms
}
