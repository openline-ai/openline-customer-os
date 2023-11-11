package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

func MapEntityToContract(entity *entity.ContractEntity) *model.Contract {
	if entity == nil {
		return nil
	}
	return &model.Contract{
		ID:                              entity.ID,
		ReferenceID:                     utils.StringPtrNillable(entity.ReferenceId),
		CustomerOsID:                    entity.CustomerOsId,
		Name:                            entity.Name,
		CreatedAt:                       entity.CreatedAt,
		UpdatedAt:                       entity.UpdatedAt,
		Source:                          MapDataSourceToModel(entity.Source),
		Status:                          MapContractStatusToModel(entity.ContractStatus),
		RenewalCycle:                    MapContractRenewalCycleToModel(entity.ContractRenewalCycle),
		AppSource:                       entity.AppSource,
		ServiceStartedAt:                *entity.ServiceStartedAt,
		ServiceStartedAtTimelineEventID: entity.ServiceStartedId,
	}
}
