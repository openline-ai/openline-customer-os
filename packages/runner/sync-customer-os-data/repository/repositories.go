package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"gorm.io/gorm"
)

type Dbs struct {
	ControlDb      *gorm.DB
	Neo4jDriver    *neo4j.DriverWithContext
	AirbyteStoreDB *config.AirbyteStoreDB
}

type Repositories struct {
	Dbs                          Dbs
	TenantSyncSettingsRepository TenantSyncSettingsRepository
	SyncRunRepository            SyncRunRepository
	ConversationEventRepository  ConversationEventRepository

	ContactRepository          ContactRepository
	ExternalSystemRepository   ExternalSystemRepository
	OrganizationRepository     OrganizationRepository
	RoleRepository             JobRoleRepository
	UserRepository             UserRepository
	NoteRepository             NoteRepository
	InteractionEventRepository InteractionEventRepository
	TicketRepository           TicketRepository
}

func InitRepos(driver *neo4j.DriverWithContext, controlDb *gorm.DB, airbyteStoreDb *config.AirbyteStoreDB) *Repositories {
	repositories := Repositories{
		Dbs: Dbs{
			Neo4jDriver:    driver,
			ControlDb:      controlDb,
			AirbyteStoreDB: airbyteStoreDb,
		},
		TenantSyncSettingsRepository: NewTenantSyncSettingsRepository(controlDb),
		SyncRunRepository:            NewSyncRunRepository(controlDb),
		ConversationEventRepository:  NewConversationEventRepository(controlDb),
		ContactRepository:            NewContactRepository(driver),
		ExternalSystemRepository:     NewExternalSystemRepository(driver),
		OrganizationRepository:       NewOrganizationRepository(driver),
		RoleRepository:               NewJobRoleRepository(driver),
		UserRepository:               NewUserRepository(driver),
		NoteRepository:               NewNoteRepository(driver),
		InteractionEventRepository:   NewInteractionEventRepository(driver),
		TicketRepository:             NewTicketRepository(driver),
	}
	return &repositories
}
