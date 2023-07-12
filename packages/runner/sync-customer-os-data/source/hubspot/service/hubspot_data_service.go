package service

import (
	"encoding/json"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/common"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/hubspot"
	localEntity "github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/hubspot/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/hubspot/repository"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type hubspotDataService struct {
	airbyteStoreDb *config.AirbyteStoreDB
	tenant         string
	instance       string
	contacts       map[string]localEntity.Contact
	companiesRaw   map[string]string
	ownersRaw      map[string]string
	notesRaw       map[string]string
	emails         map[string]localEntity.Email
	meetings       map[string]localEntity.Meeting
}

func NewHubspotDataService(airbyteStoreDb *config.AirbyteStoreDB, tenant string) common.SourceDataService {
	return &hubspotDataService{
		airbyteStoreDb: airbyteStoreDb,
		tenant:         tenant,
		contacts:       map[string]localEntity.Contact{},
		companiesRaw:   map[string]string{},
		ownersRaw:      map[string]string{},
		notesRaw:       map[string]string{},
		emails:         map[string]localEntity.Email{},
		meetings:       map[string]localEntity.Meeting{},
	}
}

func (s *hubspotDataService) Refresh() {
	err := s.getDb().AutoMigrate(&localEntity.SyncStatus{})
	if err != nil {
		logrus.Error(err)
	}
	err = s.getDb().AutoMigrate(&localEntity.SyncStatusContact{})
	if err != nil {
		logrus.Error(err)
	}
	err = s.getDb().AutoMigrate(&localEntity.SyncStatusEmail{})
	if err != nil {
		logrus.Error(err)
	}
	err = s.getDb().AutoMigrate(&localEntity.SyncStatusMeeting{})
	if err != nil {
		logrus.Error(err)
	}
}

func (s *hubspotDataService) getDb() *gorm.DB {
	schemaName := s.SourceId()

	if len(s.instance) > 0 {
		schemaName = schemaName + "_" + s.instance
	}
	schemaName = schemaName + "_" + s.tenant
	return s.airbyteStoreDb.GetDBHandler(&config.Context{
		Schema: schemaName,
	})
}

func (s *hubspotDataService) SourceId() string {
	return string(entity.AirbyteSourceHubspot)
}

func (s *hubspotDataService) Close() {
	s.ownersRaw = make(map[string]string)
	s.contacts = make(map[string]localEntity.Contact)
	s.companiesRaw = make(map[string]string)
	s.notesRaw = make(map[string]string)
	s.emails = make(map[string]localEntity.Email)
	s.meetings = make(map[string]localEntity.Meeting)
}

func (s *hubspotDataService) GetContactsForSync(batchSize int, runId string) []entity.ContactData {
	hubspotContacts, err := repository.GetContacts(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	customerOsContacts := make([]entity.ContactData, 0, len(hubspotContacts))
	for _, v := range hubspotContacts {
		hubspotContactProperties, err := repository.GetContactProperties(s.getDb(), v.AirbyteAbId, v.AirbyteContactsHashid)
		if err != nil {
			logrus.Error(err)
			continue
		}
		// set main contact fields
		contactForCustomerOs := entity.ContactData{
			ExternalId:          v.Id,
			ExternalSyncId:      v.Id,
			ExternalSystem:      s.SourceId(),
			FirstName:           hubspotContactProperties.FirstName,
			LastName:            hubspotContactProperties.LastName,
			JobTitle:            hubspotContactProperties.JobTitle,
			CreatedAt:           v.CreateDate.UTC(),
			UpdatedAt:           v.UpdatedDate.UTC(),
			PrimaryEmail:        hubspotContactProperties.Email,
			AdditionalEmails:    strings.Split(hubspotContactProperties.AdditionalEmails, ";"),
			PhoneNumber:         hubspotContactProperties.PhoneNumber,
			UserExternalOwnerId: hubspotContactProperties.OwnerId,
			Country:             hubspotContactProperties.Country,
			Region:              hubspotContactProperties.State,
			Locality:            hubspotContactProperties.City,
			Address:             hubspotContactProperties.Address,
			Zip:                 hubspotContactProperties.Zip,
		}
		// set reference to linked organizations
		contactForCustomerOs.OrganizationsExternalIds = utils.ConvertJsonbToStringSlice(v.CompaniesExternalIds)
		// set reference to primary organization
		if hubspotContactProperties.PrimaryCompanyExternalId.Valid {
			contactForCustomerOs.PrimaryOrganizationExternalId = strconv.FormatFloat(hubspotContactProperties.PrimaryCompanyExternalId.Float64, 'f', 0, 64)
		}
		// add primary organization to organizations list
		contactForCustomerOs.OrganizationsExternalIds = append(contactForCustomerOs.OrganizationsExternalIds, contactForCustomerOs.PrimaryOrganizationExternalId)
		// remove any duplicated organizations
		contactForCustomerOs.OrganizationsExternalIds = utils.GetUniqueElements(contactForCustomerOs.OrganizationsExternalIds)

		// set custom fields
		var textCustomFields []entity.TextCustomField
		if len(hubspotContactProperties.LifecycleStage) > 0 {
			textCustomFields = append(textCustomFields, entity.TextCustomField{
				Name:           "Hubspot Lifecycle Stage",
				Value:          hubspotContactProperties.LifecycleStage,
				ExternalSystem: s.SourceId(),
				CreatedAt:      v.CreateDate.UTC(),
			})
		}
		contactForCustomerOs.TextCustomFields = textCustomFields

		// set contact's tags
		if isCustomerTag(hubspotContactProperties.LifecycleStage) {
			contactForCustomerOs.Tags = append(contactForCustomerOs.Tags, "CUSTOMER")
		} else if isProspectTag(hubspotContactProperties.LifecycleStage) {
			contactForCustomerOs.Tags = append(contactForCustomerOs.Tags, "PROSPECT")
		}

		customerOsContacts = append(customerOsContacts, contactForCustomerOs)
		s.contacts[contactForCustomerOs.ExternalSyncId] = v
	}
	return customerOsContacts
}

func isCustomerTag(hubspotLifecycleStage string) bool {
	customerLifecycleStages := map[string]bool{
		"customer": true}
	return customerLifecycleStages[hubspotLifecycleStage]
}

func isProspectTag(hubspotLifecycleStage string) bool {
	prospectLifecycleStages := map[string]bool{
		"lead": true, "subscriber": true, "marketingqualifiedlead": true, "salesqualifiedlead": true,
		"opportunity": true}
	return prospectLifecycleStages[hubspotLifecycleStage]
}

func (s *hubspotDataService) GetOrganizationsForSync(batchSize int, runId string) []entity.OrganizationData {
	airbyteRecords, err := repository.GetCompanies(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	var organizations []entity.OrganizationData
	for _, v := range airbyteRecords {
		organization := entity.OrganizationData{}
		outputJSON, err := hubspot.MapOrganization(v.AirbyteData)
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal([]byte(outputJSON), &organization)
		if err != nil {
			logrus.Error(err)
			continue
		}
		organization.ExternalSyncId = organization.ExternalId
		organization.ExternalSystem = s.SourceId()
		organization.Id = ""

		s.companiesRaw[organization.ExternalSyncId] = v.AirbyteAbId
		organizations = append(organizations, organization)
	}
	return organizations
}

func (s *hubspotDataService) GetUsersForSync(batchSize int, runId string) []entity.UserData {
	airbyteRecords, err := repository.GetOwners(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	var users []entity.UserData
	for _, v := range airbyteRecords {
		user := entity.UserData{}
		outputJSON, err := hubspot.MapUser(v.AirbyteData)
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal([]byte(outputJSON), &user)
		if err != nil {
			logrus.Error(err)
			continue
		}
		user.ExternalSyncId = user.ExternalId
		user.ExternalSystem = s.SourceId()
		user.Id = ""

		s.ownersRaw[user.ExternalSyncId] = v.AirbyteAbId
		users = append(users, user)
	}
	return users
}

func (s *hubspotDataService) GetNotesForSync(batchSize int, runId string) []entity.NoteData {
	airbyteRecords, err := repository.GetNotes(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	var notes []entity.NoteData
	for _, v := range airbyteRecords {
		note := entity.NoteData{}
		outputJSON, err := hubspot.MapNote(v.AirbyteData)
		if err != nil {
			logrus.Error(err)
			continue
		}
		err = json.Unmarshal([]byte(outputJSON), &note)
		if err != nil {
			logrus.Error(err)
			continue
		}
		note.ExternalSyncId = note.ExternalId
		note.ExternalSystem = s.SourceId()
		note.Id = ""

		s.notesRaw[note.ExternalSyncId] = v.AirbyteAbId
		notes = append(notes, note)
	}
	return notes
}

func (s *hubspotDataService) GetEmailMessagesForSync(batchSize int, runId string) []entity.EmailMessageData {
	hubspotEmails, err := repository.GetEmails(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	customerOsEmails := []entity.EmailMessageData{}
	for _, v := range hubspotEmails {
		hubspotEmailProperties, err := repository.GetEmailProperties(s.getDb(), v.AirbyteAbId, v.AirbyteEmailsHashid)
		if err != nil {
			logrus.Error(err)
			continue
		}
		// set main fields
		emailForCustomerOS := entity.EmailMessageData{
			Html:           hubspotEmailProperties.EmailHtml,
			Text:           hubspotEmailProperties.EmailText,
			Subject:        hubspotEmailProperties.EmailSubject,
			CreatedAt:      v.CreateDate.UTC(),
			ExternalId:     v.Id,
			ExternalSystem: s.SourceId(),
			EmailThreadId:  hubspotEmailProperties.EmailThreadId,
			EmailMessageId: hubspotEmailProperties.EmailMessageId,
			FromEmail:      hubspotEmailProperties.EmailFromEmail,
			ToEmail:        emailsStringToArray(hubspotEmailProperties.EmailToEmail),
			CcEmail:        emailsStringToArray(hubspotEmailProperties.EmailCcEmail),
			BccEmail:       emailsStringToArray(hubspotEmailProperties.EmailBccEmail),
			FromFirstName:  hubspotEmailProperties.EmailFromFirstName,
			FromLastName:   hubspotEmailProperties.EmailFromLastName,
		}
		// set user id
		if hubspotEmailProperties.CreatedByUserId.Valid {
			emailForCustomerOS.UserExternalId = strconv.FormatFloat(hubspotEmailProperties.CreatedByUserId.Float64, 'f', 0, 64)
		}
		// set email message direction
		if hubspotEmailProperties.EmailDirection == "INCOMING_EMAIL" {
			emailForCustomerOS.Direction = entity.INBOUND
		} else {
			emailForCustomerOS.Direction = entity.OUTBOUND
		}
		// set reference to all linked contacts
		emailForCustomerOS.ContactsExternalIds = utils.ConvertJsonbToStringSlice(v.ContactsExternalIds)
		customerOsEmails = append(customerOsEmails, emailForCustomerOS)
		s.emails[v.Id] = v
	}
	return customerOsEmails
}

func emailsStringToArray(str string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ";")
}

func (s *hubspotDataService) GetIssuesForSync(batchSize int, runId string) []entity.IssueData {
	// no need to implement
	return nil
}

func (s *hubspotDataService) GetMeetingsForSync(batchSize int, runId string) []entity.MeetingData {
	hubspotMeetings, err := repository.GetMeetings(s.getDb(), batchSize, runId)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	customerOsMeetings := []entity.MeetingData{}
	for _, v := range hubspotMeetings {
		hubspotMeetingProperties, err := repository.GetMeetingProperties(s.getDb(), v.AirbyteAbId, v.AirbyteMeetingsHashid)
		if err != nil {
			logrus.Error(err)
			continue
		}
		// set main fields
		meetingForCustomerOS := entity.MeetingData{
			ExternalId:         v.Id,
			ExternalSyncId:     v.Id,
			ExternalSystem:     s.SourceId(),
			CreatedAt:          v.CreateDate.UTC(),
			UpdatedAt:          v.UpdatedDate.UTC(),
			Name:               hubspotMeetingProperties.Title,
			MeetingExternalUrl: hubspotMeetingProperties.MeetingExternalUrl,
			StartedAt:          hubspotMeetingProperties.StartedAt.UTC(),
			EndedAt:            hubspotMeetingProperties.EndedAt.UTC(),
		}
		if len(hubspotMeetingProperties.Location) > 0 {
			if strings.HasPrefix(hubspotMeetingProperties.Location, "https://") {
				meetingForCustomerOS.ConferenceUrl = hubspotMeetingProperties.Location
			} else {
				meetingForCustomerOS.Location = hubspotMeetingProperties.Location
			}
		}
		if len(hubspotMeetingProperties.MeetingHtml) > 0 {
			meetingForCustomerOS.Agenda = hubspotMeetingProperties.MeetingHtml
			meetingForCustomerOS.AgendaContentType = "text/html"
		} else if len(hubspotMeetingProperties.MeetingText) > 0 {
			meetingForCustomerOS.Agenda = hubspotMeetingProperties.MeetingText
			meetingForCustomerOS.AgendaContentType = "text/plain"
		}
		// set user id
		if hubspotMeetingProperties.CreatedByUserId.Valid {
			meetingForCustomerOS.UserCreatorExternalId = strconv.FormatFloat(hubspotMeetingProperties.CreatedByUserId.Float64, 'f', 0, 64)
		}
		// set reference to all linked contacts
		meetingForCustomerOS.ContactsExternalIds = utils.ConvertJsonbToStringSlice(v.ContactsExternalIds)
		customerOsMeetings = append(customerOsMeetings, meetingForCustomerOS)
		s.meetings[v.Id] = v
	}
	return customerOsMeetings
}

func (s *hubspotDataService) MarkContactProcessed(externalSyncId, runId string, synced bool) error {
	contact, ok := s.contacts[externalSyncId]
	if ok {
		err := repository.MarkContactProcessed(s.getDb(), contact, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking contact with external reference %s as synced for hubspot", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *hubspotDataService) MarkOrganizationProcessed(externalSyncId, runId string, synced bool) error {
	airbyteAbId, ok := s.companiesRaw[externalSyncId]
	if ok {
		err := repository.MarkProcessed(s.getDb(), repository.CompanyEntity, airbyteAbId, synced, runId, externalSyncId)
		if err != nil {
			logrus.Errorf("error while marking company with external reference %s as synced for hubspot", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *hubspotDataService) MarkUserProcessed(externalSyncId, runId string, synced bool) error {
	airbyteAbId, ok := s.ownersRaw[externalSyncId]
	if ok {
		err := repository.MarkProcessed(s.getDb(), repository.OwnerEntity, airbyteAbId, synced, runId, externalSyncId)
		if err != nil {
			logrus.Errorf("error while marking owner with external reference %s as synced for hubspot", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *hubspotDataService) MarkNoteProcessed(externalSyncId, runId string, synced bool) error {
	airbyteAbId, ok := s.notesRaw[externalSyncId]
	if ok {
		err := repository.MarkProcessed(s.getDb(), repository.NoteEntity, airbyteAbId, synced, runId, externalSyncId)
		if err != nil {
			logrus.Errorf("error while marking note with external reference %s as synced for hubspot", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *hubspotDataService) MarkIssueProcessed(externalSyncId, runId string, synced bool) error {
	// no need to implement
	return nil
}

func (s *hubspotDataService) MarkEmailMessageProcessed(externalSyncId, runId string, synced bool) error {
	email, ok := s.emails[externalSyncId]
	if ok {
		err := repository.MarkEmailProcessed(s.getDb(), email, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking email with external reference %s as synced for hubspot", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *hubspotDataService) MarkMeetingProcessed(externalSyncId, runId string, synced bool) error {
	meeting, ok := s.meetings[externalSyncId]
	if ok {
		err := repository.MarkMeetingProcessed(s.getDb(), meeting, synced, runId)
		if err != nil {
			logrus.Errorf("error while marking meeting with external reference %s as synced for hubspot", externalSyncId)
		}
		return err
	}
	return nil
}

func (s *hubspotDataService) GetInteractionEventsForSync(batchSize int, runId string) []entity.InteractionEventData {
	return nil
}

func (s *hubspotDataService) MarkInteractionEventProcessed(externalSyncId, runId string, synced bool) error {
	return nil
}
