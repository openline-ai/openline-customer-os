package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
)

func MapEntityToContract(entity *neo4jentity.ContractEntity) *model.Contract {
	if entity == nil {
		return nil
	}
	return &model.Contract{
		Metadata: &model.Metadata{
			ID:            entity.Id,
			Created:       entity.CreatedAt,
			LastUpdated:   entity.UpdatedAt,
			Source:        MapDataSourceToModel(entity.Source),
			SourceOfTruth: MapDataSourceToModel(entity.SourceOfTruth),
			AppSource:     entity.AppSource,
		},
		BillingDetails: &model.BillingDetails{
			BillingCycle:           utils.ToPtr(MapContractBillingCycleToModel(entity.BillingCycle)),
			InvoicingStarted:       entity.InvoicingStartDate,
			NextInvoicing:          entity.NextInvoiceDate,
			AddressLine1:           utils.ToPtr(entity.AddressLine1),
			AddressLine2:           utils.ToPtr(entity.AddressLine2),
			Locality:               utils.ToPtr(entity.Locality),
			Region:                 utils.ToPtr(entity.Region),
			Country:                utils.ToPtr(entity.Country),
			PostalCode:             utils.ToPtr(entity.Zip),
			OrganizationLegalName:  utils.ToPtr(entity.OrganizationLegalName),
			BillingEmail:           utils.ToPtr(entity.InvoiceEmail),
			InvoiceNote:            utils.ToPtr(entity.InvoiceNote),
			CanPayWithCard:         utils.ToPtr(entity.CanPayWithCard),
			CanPayWithDirectDebit:  utils.ToPtr(entity.CanPayWithDirectDebit),
			CanPayWithBankTransfer: utils.ToPtr(entity.CanPayWithBankTransfer),
			PayOnline:              utils.ToPtr(entity.PayOnline),
			PayAutomatically:       utils.ToPtr(entity.PayAutomatically),
			Check:                  utils.ToPtr(entity.Check),
			DueDays:                utils.ToPtr(entity.DueDays),
		},
		CommittedPeriods:     entity.RenewalPeriods,
		ContractEnded:        entity.EndedAt,
		ContractName:         entity.Name,
		ContractRenewalCycle: MapContractRenewalCycleToModel(entity.RenewalCycle),
		ContractSigned:       entity.SignedAt,
		ContractURL:          utils.StringPtrNillable(entity.ContractUrl),
		Currency:             utils.ToPtr(MapCurrencyToModel(entity.Currency)),
		BillingEnabled:       entity.InvoicingEnabled,
		ServiceStarted:       entity.ServiceStartedAt,
		ContractStatus:       MapContractStatusToModel(entity.ContractStatus),
		AutoRenew:            entity.AutoRenew,

		// All below are deprecated
		ID:                    entity.Id,
		Name:                  entity.Name,
		CreatedAt:             entity.CreatedAt,
		UpdatedAt:             entity.UpdatedAt,
		Source:                MapDataSourceToModel(entity.Source),
		SourceOfTruth:         MapDataSourceToModel(entity.SourceOfTruth),
		AppSource:             entity.AppSource,
		Status:                MapContractStatusToModel(entity.ContractStatus),
		RenewalCycle:          MapContractRenewalCycleToModel(entity.RenewalCycle),
		RenewalPeriods:        entity.RenewalPeriods,
		ServiceStartedAt:      entity.ServiceStartedAt,
		SignedAt:              entity.SignedAt,
		EndedAt:               entity.EndedAt,
		InvoicingStartDate:    entity.InvoicingStartDate,
		BillingCycle:          utils.ToPtr(MapContractBillingCycleToModel(entity.BillingCycle)),
		AddressLine1:          utils.ToPtr(entity.AddressLine1),
		AddressLine2:          utils.ToPtr(entity.AddressLine2),
		Zip:                   utils.ToPtr(entity.Zip),
		Country:               utils.ToPtr(entity.Country),
		Locality:              utils.ToPtr(entity.Locality),
		OrganizationLegalName: utils.ToPtr(entity.OrganizationLegalName),
		InvoiceEmail:          utils.ToPtr(entity.InvoiceEmail),
		InvoiceNote:           utils.ToPtr(entity.InvoiceNote),
	}
}

func MapEntitiesToContracts(entities *neo4jentity.ContractEntities) []*model.Contract {
	var contracts []*model.Contract
	for _, contractEntity := range *entities {
		contracts = append(contracts, MapEntityToContract(&contractEntity))
	}
	return contracts
}
