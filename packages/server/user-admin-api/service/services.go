package service

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	authServices "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-auth/service"
	commonService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/openline-ai/openline-customer-os/packages/server/user-admin-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/user-admin-api/grpc_client"
	"gorm.io/gorm"
)

type Services struct {
	GrpcClients *grpc_client.Clients

	CommonServices *commonService.Services
	AuthServices   *authServices.Services

	CustomerOsClient CustomerOsClient
}

func InitServices(cfg *config.Config, db *gorm.DB, driver *neo4j.DriverWithContext, grpcClients *grpc_client.Clients) *Services {
	services := Services{
		GrpcClients:      grpcClients,
		CustomerOsClient: NewCustomerOsClient(cfg, driver),
	}

	services.CommonServices = commonService.InitServices(db, driver)
	services.AuthServices = authServices.InitServices(nil, services.CommonServices, db)

	return &services
}
