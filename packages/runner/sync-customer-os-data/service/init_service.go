package service

import (
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InitService interface {
	Init()
}

type initService struct {
	repositories *repository.Repositories
	services     *Services
}

func NewInitService(repositories *repository.Repositories, services *Services) InitService {
	return &initService{
		repositories: repositories,
		services:     services,
	}
}

func (s *initService) Init() {
	db := s.repositories.Dbs.ControlDb

	createAirbyteSourceEnum(db)

	err := db.AutoMigrate(&entity.TenantSyncSettings{})
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.AutoMigrate(&entity.SyncRun{})
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.AutoMigrate(&entity.ConversationEvent{})
	if err != nil {
		logrus.Fatal(err)
	}
}

func createAirbyteSourceEnum(db *gorm.DB) *gorm.DB {
	return db.Exec(`DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'airbyte_source') THEN
        CREATE TYPE airbyte_source AS ENUM
        (
            'hubspot','zendesk'
        );
    END IF;
END$$;`)
}
