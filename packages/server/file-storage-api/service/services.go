package service

import (
	"github.com/openline-ai/openline-customer-os/packages/server/file-storage-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/file-storage-api/repository"
	"gorm.io/gorm"
)

type Services struct {
	FileService FileService
}

func InitServices(cfg *config.Config, db *gorm.DB) *Services {
	repositories := repository.InitRepositories(db)

	return &Services{
		FileService: NewFileService(cfg, db, repositories),
	}
}
