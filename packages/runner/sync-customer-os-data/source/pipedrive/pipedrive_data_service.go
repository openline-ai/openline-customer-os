package pipedrive

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/common"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/logger"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source"
	source_entity "github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source/repository"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	UsersTableSuffix         = "users"
	OrganizationsTableSuffix = "organizations"
	PersonsTableSuffix       = "persons"
	NotesTableSuffix         = "notes"
)

var sourceTableSuffixByDataType = map[string][]string{
	string(common.USERS):         {UsersTableSuffix},
	string(common.ORGANIZATIONS): {OrganizationsTableSuffix},
	string(common.CONTACTS):      {PersonsTableSuffix},
	string(common.NOTES):         {NotesTableSuffix},
}

type pipedriveDataService struct {
	airbyteStoreDb *config.RawDataStoreDB
	tenant         string
	instance       string
	processingIds  map[string]source.ProcessingEntity
	dataFuncs      map[common.SyncedEntityType]func(context.Context, int, string) []any
	log            logger.Logger
}

func NewPipedriveDataService(airbyteStoreDb *config.RawDataStoreDB, tenant string, log logger.Logger) source.SourceDataService {
	dataService := pipedriveDataService{
		airbyteStoreDb: airbyteStoreDb,
		tenant:         tenant,
		processingIds:  map[string]source.ProcessingEntity{},
		log:            log,
	}
	dataService.dataFuncs = map[common.SyncedEntityType]func(context.Context, int, string) []any{}
	dataService.dataFuncs[common.USERS] = dataService.GetUsersForSync
	dataService.dataFuncs[common.ORGANIZATIONS] = dataService.GetOrganizationsForSync
	dataService.dataFuncs[common.CONTACTS] = dataService.GetContactsForSync
	dataService.dataFuncs[common.NOTES] = dataService.GetNotesForSync
	return &dataService
}

func (s *pipedriveDataService) GetDataForSync(ctx context.Context, dataType common.SyncedEntityType, batchSize int, runId string) []interface{} {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PipedriveDataService.GetDataForSync")
	defer span.Finish()
	tracing.SetDefaultSyncServiceSpanTags(ctx, span)
	span.LogFields(log.String("dataType", string(dataType)), log.Int("batchSize", batchSize))

	if ok := s.dataFuncs[dataType]; ok != nil {
		return s.dataFuncs[dataType](ctx, batchSize, runId)
	} else {
		s.log.Warnf("No %s data function for %s", s.SourceId(), dataType)
		return nil
	}
}

func (s *pipedriveDataService) Init() {
	err := s.getDb().AutoMigrate(&source_entity.SyncStatusForAirbyte{})
	if err != nil {
		s.log.Error(err)
	}
}

func (s *pipedriveDataService) getDb() *gorm.DB {
	schemaName := s.SourceId()

	if len(s.instance) > 0 {
		schemaName = schemaName + "_" + s.instance
	}
	schemaName = schemaName + "_" + s.tenant
	return s.airbyteStoreDb.GetDBHandler(&config.Context{
		Schema: schemaName,
	})
}

func (s *pipedriveDataService) SourceId() string {
	return string(entity.AirbyteSourcePipedrive)
}

func (s *pipedriveDataService) Close() {
	s.processingIds = make(map[string]source.ProcessingEntity)
}

func (s *pipedriveDataService) GetUsersForSync(ctx context.Context, batchSize int, runId string) []any {
	s.processingIds = make(map[string]source.ProcessingEntity)
	currentEntity := string(common.USERS)
	var users []any
	for _, sourceTableSuffix := range sourceTableSuffixByDataType[currentEntity] {
		airbyteRecords, err := repository.GetAirbyteUnprocessedRawRecords(ctx, s.getDb(), batchSize, runId, currentEntity, sourceTableSuffix)
		if err != nil {
			s.log.Fatal(err) // alexb handle errors
			return nil
		}
		for _, v := range airbyteRecords {
			if len(users) >= batchSize {
				break
			}
			outputJSON, err := MapUser(v.AirbyteData)
			user, err := source.MapJsonToUser(outputJSON, v.AirbyteAbId, s.SourceId())
			if err != nil {
				s.log.Fatal(err) // alexb handle errors
				continue
			}
			s.processingIds[v.AirbyteAbId] = source.ProcessingEntity{
				ExternalId:  user.ExternalId,
				Entity:      currentEntity,
				TableSuffix: sourceTableSuffix,
			}
			users = append(users, user)
		}
	}
	return users
}

func (s *pipedriveDataService) GetOrganizationsForSync(ctx context.Context, batchSize int, runId string) []any {
	s.processingIds = make(map[string]source.ProcessingEntity)
	currentEntity := string(common.ORGANIZATIONS)

	var organizations []any
	for _, sourceTableSuffix := range sourceTableSuffixByDataType[currentEntity] {
		airbyteRecords, err := repository.GetAirbyteUnprocessedRawRecords(ctx, s.getDb(), batchSize, runId, currentEntity, sourceTableSuffix)
		if err != nil {
			s.log.Fatal(err) // alexb handle errors
			return nil
		}
		for _, v := range airbyteRecords {
			if len(organizations) >= batchSize {
				break
			}
			outputJSON, err := MapOrganization(v.AirbyteData)
			organization, err := source.MapJsonToOrganization(outputJSON, v.AirbyteAbId, s.SourceId())
			if err != nil {
				s.log.Fatal(err) // alexb handle errors
				continue
			}

			s.processingIds[v.AirbyteAbId] = source.ProcessingEntity{
				ExternalId:  organization.ExternalId,
				Entity:      currentEntity,
				TableSuffix: sourceTableSuffix,
			}
			organizations = append(organizations, organization)
		}
	}
	return organizations
}

func (s *pipedriveDataService) GetContactsForSync(ctx context.Context, batchSize int, runId string) []any {
	s.processingIds = make(map[string]source.ProcessingEntity)
	currentEntity := string(common.CONTACTS)

	var contacts []any
	for _, sourceTableSuffix := range sourceTableSuffixByDataType[currentEntity] {
		airbyteRecords, err := repository.GetAirbyteUnprocessedRawRecords(ctx, s.getDb(), batchSize, runId, currentEntity, sourceTableSuffix)
		if err != nil {
			s.log.Fatal(err) // alexb handle errors
			return nil
		}
		for _, v := range airbyteRecords {
			if len(contacts) >= batchSize {
				break
			}
			outputJSON, err := MapContact(v.AirbyteData)
			contact, err := source.MapJsonToContact(outputJSON, v.AirbyteAbId, s.SourceId())
			if err != nil {
				s.log.Fatal(err) // alexb handle errors
				continue
			}

			s.processingIds[v.AirbyteAbId] = source.ProcessingEntity{
				ExternalId:  contact.ExternalId,
				Entity:      currentEntity,
				TableSuffix: sourceTableSuffix,
			}
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

func (s *pipedriveDataService) GetNotesForSync(ctx context.Context, batchSize int, runId string) []any {
	s.processingIds = make(map[string]source.ProcessingEntity)
	currentEntity := string(common.NOTES)
	var notes []any
	for _, sourceTableSuffix := range sourceTableSuffixByDataType[currentEntity] {
		airbyteRecords, err := repository.GetAirbyteUnprocessedRawRecords(ctx, s.getDb(), batchSize, runId, currentEntity, sourceTableSuffix)
		if err != nil {
			s.log.Fatal(err) // alexb handle errors
			return nil
		}
		for _, v := range airbyteRecords {
			if len(notes) >= batchSize {
				break
			}
			outputJSON, err := MapNote(v.AirbyteData)
			note, err := source.MapJsonToNote(outputJSON, v.AirbyteAbId, s.SourceId())
			if err != nil {
				s.log.Fatal(err) // alexb handle errors
				continue
			}

			s.processingIds[v.AirbyteAbId] = source.ProcessingEntity{
				ExternalId:  note.ExternalId,
				Entity:      currentEntity,
				TableSuffix: sourceTableSuffix,
			}
			notes = append(notes, note)
		}
	}
	return notes
}

func (s *pipedriveDataService) MarkProcessed(ctx context.Context, syncId, runId string, synced, skipped bool, reason string) error {
	v, ok := s.processingIds[syncId]
	if ok {
		err := repository.MarkAirbyteRawRecordProcessed(ctx, s.getDb(), v.Entity, v.TableSuffix, syncId, synced, skipped, runId, v.ExternalId, reason)
		if err != nil {
			s.log.Errorf("error while marking %s with external reference %s as synced for %s", v.Entity, v.ExternalId, s.SourceId())
		}
		return err
	}
	return nil
}
