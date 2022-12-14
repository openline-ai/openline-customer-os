package service

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/repository"
	"gorm.io/gorm"
)

type Services struct {
	SyncService SyncService
	InitService InitService
}

func InitServices(driver *neo4j.Driver, controlDb *gorm.DB, airbyteStoreDb *config.AirbyteStoreDB) *Services {
	repositories := repository.InitRepos(driver, controlDb, airbyteStoreDb)

	services := new(Services)

	services.SyncService = NewSyncService(repositories, services)
	services.InitService = NewInitService(repositories, services)

	return services
}
