package service

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	commonAuthService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-auth/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	commonService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-webhooks/config"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-webhooks/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-webhooks/repository"
)

type Services struct {
	cfg *config.Config

	CommonServices     *commonService.Services
	CommonAuthServices *commonAuthService.Services

	TenantService      TenantService
	EmailService       EmailService
	PhoneNumberService PhoneNumberService
	UserService        UserService
	SyncStatusService  SyncStatusService
}

func InitServices(log logger.Logger, driver *neo4j.DriverWithContext, cfg *config.Config, commonServices *commonService.Services, commonAuthServices *commonAuthService.Services, grpcClients *grpc_client.Clients) *Services {
	repositories := repository.InitRepos(driver)

	services := Services{
		CommonServices:     commonServices,
		CommonAuthServices: commonAuthServices,
		TenantService:      NewTenantService(log, repositories),
		EmailService:       NewEmailService(log, repositories, grpcClients),
		PhoneNumberService: NewPhoneNumberService(log, repositories, grpcClients),
		SyncStatusService:  NewSyncStatusService(log, repositories),
	}

	services.UserService = NewUserService(log, repositories, grpcClients, &services)
	services.cfg = cfg
	return &services
}
