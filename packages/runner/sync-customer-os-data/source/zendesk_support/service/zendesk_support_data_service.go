package service

import (
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/common"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	localEntity "github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/zendesk_support/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/zendesk_support/repository"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type zendeskSupportDataService struct {
	airbyteStoreDb *config.AirbyteStoreDB
	tenant         string
	instance       string
	users          map[string]localEntity.User
	organizations  map[string]localEntity.Organization
	contacts       map[string]localEntity.Contact
	tickets        map[string]localEntity.Ticket
}

func NewZendeskSupportDataService(airbyteStoreDb *config.AirbyteStoreDB, tenant string) common.SourceDataService {
	return &zendeskSupportDataService{
		airbyteStoreDb: airbyteStoreDb,
		tenant:         tenant,
		users:          map[string]localEntity.User{},
		contacts:       map[string]localEntity.Contact{},
		organizations:  map[string]localEntity.Organization{},
		tickets:        map[string]localEntity.Ticket{},
	}
}

func (s *zendeskSupportDataService) Refresh() {
	err := s.getDb().AutoMigrate(&localEntity.SyncStatusUser{})
	if err != nil {
		logrus.Error(err)
	}
	err = s.getDb().AutoMigrate(&localEntity.SyncStatusOrganization{})
	if err != nil {
		logrus.Error(err)
	}
	err = s.getDb().AutoMigrate(&localEntity.SyncStatusContact{})
	if err != nil {
		logrus.Error(err)
	}
	err = s.getDb().AutoMigrate(&localEntity.SyncStatusTicket{})
	if err != nil {
		logrus.Error(err)
	}
}

func (s *zendeskSupportDataService) getDb() *gorm.DB {
	schemaName := s.SourceId()

	if len(s.instance) > 0 {
		schemaName = schemaName + "_" + s.instance
	}
	schemaName = schemaName + "_" + s.tenant
	return s.airbyteStoreDb.GetDBHandler(&config.Context{
		Schema: schemaName,
	})
}

func (s *zendeskSupportDataService) Close() {
	s.users = make(map[string]localEntity.User)
	s.contacts = make(map[string]localEntity.Contact)
	s.organizations = make(map[string]localEntity.Organization)
	s.tickets = make(map[string]localEntity.Ticket)
}

func (s *zendeskSupportDataService) SourceId() string {
	return string(entity.AirbyteSourceZendeskSupport)
}

func (s *zendeskSupportDataService) GetContactsForSync(batchSize int, runId string) []entity.ContactData {
	zendeskContacts, err := repository.GetContacts(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	customerOsContacts := make([]entity.ContactData, 0, len(zendeskContacts))
	for _, v := range zendeskContacts {
		contactData := entity.ContactData{
			ExternalId:     strconv.FormatInt(v.Id, 10),
			ExternalSyncId: strconv.FormatInt(v.Id, 10),
			ExternalUrl:    v.Url,
			ExternalSystem: s.SourceId(),
			CreatedAt:      v.CreateDate.UTC(),
			UpdatedAt:      v.UpdatedDate.UTC(),
			PhoneNumber:    v.Phone,
		}
		if len(v.Email) > 0 && !strings.HasSuffix(v.Email, "@without-email.com") {
			contactData.AdditionalEmails = append(contactData.AdditionalEmails, v.Email)
		}
		if len(v.Notes) > 0 {
			contactData.Notes = append(contactData.Notes, entity.ContactNote{
				Note:        v.Notes,
				FieldSource: "notes",
			})
		}
		if len(v.Details) > 0 {
			contactData.Notes = append(contactData.Notes, entity.ContactNote{
				Note:        v.Details,
				FieldSource: "details",
			})
		}
		if len(v.Name) > 0 {
			contactData.TextCustomFields = append(contactData.TextCustomFields, entity.TextCustomField{
				Name:           "name",
				Value:          v.Name,
				ExternalSystem: s.SourceId(),
				CreatedAt:      v.CreateDate.UTC(),
			})
		}
		if v.OrganizationId != 0 {
			contactData.OrganizationsExternalIds = append(contactData.OrganizationsExternalIds, strconv.FormatInt(v.OrganizationId, 10))
		}
		var jsonObject map[string]string
		err = v.CustomFieldsAsJson.AssignTo(&jsonObject)
		if err == nil {
			for key, value := range jsonObject {
				if len(value) > 0 {
					contactData.TextCustomFields = append(contactData.TextCustomFields, entity.TextCustomField{
						Name:           key,
						Value:          value,
						ExternalSystem: s.SourceId(),
						CreatedAt:      v.CreateDate.UTC(),
					})
				}
			}
		}

		customerOsContacts = append(customerOsContacts, contactData)
		s.contacts[contactData.ExternalSyncId] = v
	}
	return customerOsContacts
}

func (s *zendeskSupportDataService) GetOrganizationsForSync(batchSize int, runId string) []entity.OrganizationData {
	zendeskOrganizations, err := repository.GetOrganizations(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	customerOsOrganizations := make([]entity.OrganizationData, 0, len(zendeskOrganizations))
	for _, v := range zendeskOrganizations {
		organizationData := entity.OrganizationData{
			ExternalId:     strconv.FormatInt(v.Id, 10),
			ExternalSyncId: strconv.FormatInt(v.Id, 10),
			ExternalSystem: s.SourceId(),
			CreatedAt:      v.CreateDate.UTC(),
			UpdatedAt:      v.UpdatedDate.UTC(),
			Name:           v.Name,
			NoteContent:    v.Details,
		}
		organizationData.Domains = utils.GetUniqueElements(utils.ConvertJsonbToStringSlice(v.DomainNames))

		customerOsOrganizations = append(customerOsOrganizations, organizationData)
		s.organizations[organizationData.ExternalSyncId] = v
	}
	return customerOsOrganizations
}

func (s *zendeskSupportDataService) GetUsersForSync(batchSize int, runId string) []entity.UserData {
	zendeskUsers, err := repository.GetUsers(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	customerOsUsers := make([]entity.UserData, 0, len(zendeskUsers))
	for _, v := range zendeskUsers {
		userData := entity.UserData{
			ExternalId:     strconv.FormatInt(v.Id, 10),
			ExternalSystem: s.SourceId(),
			Name:           v.Name,
			Email:          v.Email,
			PhoneNumber:    v.Phone,
			CreatedAt:      v.CreateDate.UTC(),
			UpdatedAt:      v.UpdatedDate.UTC(),
			ExternalSyncId: strconv.FormatInt(v.Id, 10),
		}
		customerOsUsers = append(customerOsUsers, userData)

		s.users[userData.ExternalSyncId] = v
	}
	return customerOsUsers
}

func (z zendeskSupportDataService) GetNotesForSync(batchSize int, runId string) []entity.NoteData {
	//TODO implement me
	return nil
}

func (z zendeskSupportDataService) GetEmailMessagesForSync(batchSize int, runId string) []entity.EmailMessageData {
	//TODO implement me
	return nil
}

func (s *zendeskSupportDataService) MarkContactProcessed(externalSyncId, runId string, synced bool) error {
	contact, ok := s.contacts[externalSyncId]
	if ok {
		err := repository.MarkContactProcessed(s.getDb(), contact, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking contact with external reference %s as synced for zendesk support", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *zendeskSupportDataService) MarkOrganizationProcessed(externalSyncId, runId string, synced bool) error {
	organization, ok := s.organizations[externalSyncId]
	if ok {
		err := repository.MarkOrganizationProcessed(s.getDb(), organization, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking organization with external reference %s as synced for zendesk support", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *zendeskSupportDataService) MarkUserProcessed(externalSyncId, runId string, synced bool) error {
	user, ok := s.users[externalSyncId]
	if ok {
		err := repository.MarkUserProcessed(s.getDb(), user, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking owner with external reference %s as synced for zendesk support", externalSyncId)
		}
		return err
	}
	return nil
}

func (z zendeskSupportDataService) MarkNoteProcessed(externalSyncId, runId string, synced bool) error {
	//TODO implement me
	return nil
}

func (z zendeskSupportDataService) MarkEmailMessageProcessed(externalSyncId, runId string, synced bool) error {
	//TODO implement me
	return nil
}

func (s *zendeskSupportDataService) GetTicketsForSync(batchSize int, runId string) []entity.TicketData {
	zendeskTickets, err := repository.GetTickets(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	ticketFields, err := repository.GetTicketFieldsNotInTypes(s.getDb(), []string{"tagger"})
	ticketFieldMap := make(map[int64]string)
	for _, ticketField := range ticketFields {
		ticketFieldMap[ticketField.Id] = ticketField.Title
	}

	ticketsToReturn := make([]entity.TicketData, 0, len(zendeskTickets))

	for _, v := range zendeskTickets {
		ticketData := entity.TicketData{
			ExternalId:     strconv.FormatInt(v.Id, 10),
			ExternalSyncId: strconv.FormatInt(v.Id, 10),
			ExternalSystem: s.SourceId(),
			ExternalUrl:    v.Url,
			CreatedAt:      v.CreateDate.UTC(),
			UpdatedAt:      v.UpdatedDate.UTC(),
			Subject:        v.Subject,
			Status:         v.Status,
			Priority:       v.Priority,
			Description:    v.Description,
		}
		ticketData.CollaboratorUserExternalIds = utils.GetUniqueElements(utils.ConvertJsonbToStringSlice(v.CollaboratorIds))
		ticketData.FollowerUserExternalIds = utils.GetUniqueElements(utils.ConvertJsonbToStringSlice(v.FollowerIds))
		if v.SubmitterId > 0 {
			ticketData.SubmitterExternalId = strconv.FormatInt(v.SubmitterId, 10)
		}
		if v.RequesterId > 0 {
			ticketData.RequesterExternalId = strconv.FormatInt(v.RequesterId, 10)
		}
		if v.AssigneeId > 0 {
			ticketData.AssigneeUserExternalId = strconv.FormatInt(v.AssigneeId, 10)
		}
		if len(v.Type) > 0 {
			ticketData.Tags = append(ticketData.Tags, "type:"+v.Type)
		}
		ticketData.Tags = append(ticketData.Tags, utils.GetUniqueElements(utils.ConvertJsonbToStringSlice(v.Tags))...)

		type TicketCustomField struct {
			Id    int64  `json:"id"`
			Value string `json:"value"`
		}
		var ticketCustomFields []TicketCustomField
		err = v.CustomFieldsAsJson.AssignTo(&ticketCustomFields)
		if err == nil {
			for _, customField := range ticketCustomFields {
				if _, ok := ticketFieldMap[customField.Id]; ok {
					if len(customField.Value) > 0 {
						ticketData.TextCustomFields = append(ticketData.TextCustomFields, entity.TextCustomField{
							Name:           ticketFieldMap[customField.Id],
							Value:          customField.Value,
							ExternalSystem: s.SourceId(),
							CreatedAt:      v.CreateDate.UTC(),
						})
					}
				}
			}
		}

		ticketsToReturn = append(ticketsToReturn, ticketData)
		s.tickets[ticketData.ExternalSyncId] = v
	}
	return ticketsToReturn
}

func (s *zendeskSupportDataService) MarkTicketProcessed(externalSyncId, runId string, synced bool) error {
	ticket, ok := s.tickets[externalSyncId]
	if ok {
		err := repository.MarkTicketProcessed(s.getDb(), ticket, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking ticket with external reference %s as synced for zendesk support", externalSyncId)
		}
		return err
	}
	return nil
}
