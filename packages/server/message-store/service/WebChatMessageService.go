package service

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store/proto/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/repository/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type webChatMessageStoreService struct {
	msProto.UnimplementedWebChatMessageStoreServiceServer
	driver               *neo4j.Driver
	postgresRepositories *repository.PostgresRepositories
	customerOSService    *customerOSService
}

// sender == contact -> find contact by email
// sender == user -> find user by email

// sender == contact -> find conversation by initiator = contact and channel = webchat
// sender == user -> find conversation by id
func (s *webChatMessageStoreService) SaveMessage(ctx context.Context, input *msProto.WebChatInputMessage) (*msProto.Message, error) {
	//var err error
	//var conversation *gen.Conversation
	//
	var conversationId string
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

	if input.SenderType == msProto.SenderType_CONTACT {

		contactId, err := s.getContactIdWithEmailOrCreate(tenant, *input.Email)
		if err != nil {
			return nil, err
		}
		participantId = contactId
	} else if input.SenderType == msProto.SenderType_USER {
		user, err := s.customerOSService.GetUserByEmail(*input.Email)
		if err != nil {
			return nil, err
		}
		participantId = user.Id
	}

	if conversationId == "" {
		convId, err := s.getActiveConversationOrCreate(tenant, participantId, input.SenderType)
		if err != nil {
			return nil, err
		}
		conversationId = convId
	}

	_, err := s.customerOSService.UpdateConversation(tenant, conversationId, participantId, convertSenderTypeToConversationSenderType(input.SenderType))
	if err != nil {
		return nil, err
	}

	conversationEvent := entity.ConversationEvent{
		TenantName:     tenant,
		ConversationId: conversationId,
		Type:           entity.WEB_CHAT,
		Subtype:        encodeConversationEventSubtype(input.Type),
		Content:        *input.Message,
		Source:         entity.OPENLINE,
		Direction:      encodeConversationEventDirection(input.Direction),
		CreateDate:     time.Time{},

		OriginalJson: "TODO",
	}

	if input.GetDirection() == msProto.MessageDirection_INBOUND {
		conversationEvent.SenderId = participantId
		conversationEvent.SenderType = entity.CONTACT
	} else {
		conversationEvent.SenderId = participantId
		conversationEvent.SenderType = entity.USER
	}

	s.postgresRepositories.ConversationEventRepository.Save(&conversationEvent)

	mi := &msProto.Message{
		Id:             conversationEvent.ID,
		ConversationId: conversationId,
		Type:           input.Type,
		Message:        *input.Message,
		Direction:      input.Direction,
		Channel:        msProto.MessageChannel_WEB_CHAT,
		SenderType:     input.SenderType,
		SenderId:       participantId,
		Time:           timestamppb.New(time.Now()),
	}
	return mi, nil
}

func (s *webChatMessageStoreService) getContactIdWithEmailOrCreate(tenant string, email string) (string, error) {
	contact, err := s.customerOSService.GetContactByEmail(email)
	if err != nil {
		contactId, err := s.customerOSService.CreateContactWithEmail(tenant, email)
		if err != nil {
			return "", err
		}
		return contactId, nil
	} else {
		return contact.Id, nil
	}
}

func (s *webChatMessageStoreService) getActiveConversationOrCreate(tenant string, participantId string, senderType msProto.SenderType) (string, error) {
	var conversationId string

	if senderType == msProto.SenderType_CONTACT {
		convId, err := s.customerOSService.GetWebChatConversationIdWithContactInitiator(tenant, participantId)
		if err != nil {
			return "", err
		}
		if convId != "" {
			conversationId = convId
		}
	} else if senderType == msProto.SenderType_USER {
		convId, err := s.customerOSService.GetWebChatConversationIdWithUserInitiator(tenant, participantId)
		if err != nil {
			return "", err
		}
		if convId != "" {
			conversationId = convId
		}
	}

	if conversationId == "" {
		convId, err := s.customerOSService.CreateConversation(tenant, participantId, convertSenderTypeToConversationSenderType(senderType), entity.WEB_CHAT)
		if err != nil {
			return "", err
		}
		conversationId = convId
	}

	return conversationId, nil
}

func convertSenderTypeToConversationSenderType(senderType msProto.SenderType) entity.SenderType {
	switch senderType {
	case msProto.SenderType_CONTACT:
		return entity.CONTACT
	case msProto.SenderType_USER:
		return entity.USER
	default:
		return entity.CONTACT
	}
}

func encodeConversationEventDirection(direction msProto.MessageDirection) entity.Direction {
	switch direction {
	case msProto.MessageDirection_INBOUND:
		return entity.INBOUND
	case msProto.MessageDirection_OUTBOUND:
		return entity.OUTBOUND
	default:
		return entity.OUTBOUND
	}
}

func encodeConversationEventSubtype(messageType msProto.MessageType) entity.EventSubtype {
	switch messageType {
	case msProto.MessageType_FILE:
		return entity.FILE
	case msProto.MessageType_MESSAGE:
		return entity.TEXT
	default:
		return entity.TEXT
	}
}

func NewWebChatMessageStoreService(driver *neo4j.Driver, postgresRepositories *repository.PostgresRepositories, customerOSService *customerOSService) *webChatMessageStoreService {
	ms := new(webChatMessageStoreService)
	ms.driver = driver
	ms.postgresRepositories = postgresRepositories
	ms.customerOSService = customerOSService
	return ms
}
