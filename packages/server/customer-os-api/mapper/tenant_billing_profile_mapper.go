package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
)

func MapEntityToTenantBillingProfile(entity *neo4jentity.TenantBillingProfileEntity) *model.TenantBillingProfile {
	if entity == nil {
		return nil
	}
	return &model.TenantBillingProfile{
		ID:                            entity.Id,
		CreatedAt:                     entity.CreatedAt,
		UpdatedAt:                     entity.UpdatedAt,
		Source:                        MapDataSourceToModel(entity.Source),
		SourceOfTruth:                 MapDataSourceToModel(entity.SourceOfTruth),
		AppSource:                     entity.AppSource,
		LegalName:                     entity.LegalName,
		Email:                         entity.Email,
		Phone:                         entity.Phone,
		AddressLine1:                  entity.AddressLine1,
		AddressLine2:                  entity.AddressLine2,
		AddressLine3:                  entity.AddressLine3,
		Locality:                      entity.Locality,
		Country:                       entity.Country,
		Zip:                           entity.Zip,
		DomesticPaymentsBankInfo:      entity.DomesticPaymentsBankInfo,
		InternationalPaymentsBankInfo: entity.InternationalPaymentsBankInfo,
	}
}

func MapEntitiesToTenantBillingProfiles(entities *neo4jentity.TenantBillingProfileEntities) []*model.TenantBillingProfile {
	var models []*model.TenantBillingProfile
	for _, entity := range *entities {
		models = append(models, MapEntityToTenantBillingProfile(&entity))
	}
	return models
}