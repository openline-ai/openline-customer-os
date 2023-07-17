package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-gmail/entity"
	"gorm.io/gorm"
)

type Repositories struct {
	Neo4jDriver *neo4j.DriverWithContext

	ApiKeyRepository                   ApiKeyRepository
	TenantRepository                   TenantRepository
	UserRepository                     UserRepository
	EmailRepository                    EmailRepository
	ExternalSystemRepository           ExternalSystemRepository
	InteractionEventRepository         InteractionEventRepository
	UserGmailImportPageTokenRepository UserGmailImportPageTokenRepository
	ContactRepository                  ContactRepository
	WorkspaceRepository                WorkspaceRepository
}

func InitRepos(driver *neo4j.DriverWithContext, gormDb *gorm.DB) *Repositories {
	repositories := Repositories{

		Neo4jDriver:                        driver,
		ApiKeyRepository:                   NewApiKeyRepository(gormDb),
		TenantRepository:                   NewTenantRepository(driver),
		UserRepository:                     NewUserRepository(driver),
		EmailRepository:                    NewEmailRepository(driver),
		ExternalSystemRepository:           NewExternalSystemRepository(driver),
		InteractionEventRepository:         NewInteractionEventRepository(driver),
		UserGmailImportPageTokenRepository: NewUserGmailImportPageTokenRepository(gormDb),
		ContactRepository:                  NewContactRepository(driver),
		WorkspaceRepository:                NewWorkspaceRepository(driver),
	}

	var err error

	err = gormDb.AutoMigrate(&entity.UserGmailImportPageToken{})
	if err != nil {
		panic(err)
	}

	return &repositories
}
