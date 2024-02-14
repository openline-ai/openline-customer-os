package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
)

func MapEntityToInvoice(entity *neo4jentity.InvoiceEntity) *model.Invoice {
	if entity == nil {
		return nil
	}
	invoice := model.Invoice{
		ID:                            entity.Id,
		CreatedAt:                     entity.CreatedAt,
		UpdatedAt:                     entity.UpdatedAt,
		Source:                        MapDataSourceToModel(entity.Source),
		SourceOfTruth:                 MapDataSourceToModel(entity.SourceOfTruth),
		AppSource:                     entity.AppSource,
		DryRun:                        entity.DryRun,
		Number:                        entity.Number,
		PeriodStartDate:               entity.PeriodStartDate,
		PeriodEndDate:                 entity.PeriodEndDate,
		DueDate:                       entity.DueDate,
		Amount:                        entity.Amount,
		Vat:                           entity.Vat,
		Postpaid:                      entity.Postpaid,
		OffCycle:                      entity.OffCycle,
		TotalAmount:                   entity.TotalAmount,
		Currency:                      entity.Currency.String(),
		RepositoryFileID:              entity.RepositoryFileId,
		Status:                        utils.ToPtr(MapInvoiceStatusToModel(entity.Status)),
		Note:                          utils.StringPtrNillable(entity.Note),
		DomesticPaymentsBankInfo:      utils.StringPtrNillable(entity.DomesticPaymentsBankInfo),
		InternationalPaymentsBankInfo: utils.StringPtrNillable(entity.InternationalPaymentsBankInfo),
		Customer: &model.InvoiceCustomer{
			Name:            utils.StringPtrNillable(entity.Customer.Name),
			Email:           utils.StringPtrNillable(entity.Customer.Email),
			AddressLine1:    utils.StringPtrNillable(entity.Customer.AddressLine1),
			AddressLine2:    utils.StringPtrNillable(entity.Customer.AddressLine2),
			AddressZip:      utils.StringPtrNillable(entity.Customer.Zip),
			AddressLocality: utils.StringPtrNillable(entity.Customer.Locality),
			AddressCountry:  utils.StringPtrNillable(entity.Customer.Country),
		},
		Provider: &model.InvoiceProvider{
			LogoURL:         utils.StringPtrNillable(entity.Provider.LogoUrl),
			Name:            utils.StringPtrNillable(entity.Provider.Name),
			AddressLine1:    utils.StringPtrNillable(entity.Provider.AddressLine1),
			AddressLine2:    utils.StringPtrNillable(entity.Provider.AddressLine2),
			AddressZip:      utils.StringPtrNillable(entity.Provider.Zip),
			AddressLocality: utils.StringPtrNillable(entity.Provider.Locality),
			AddressCountry:  utils.StringPtrNillable(entity.Provider.Country),
		},
	}
	return &invoice
}

func MapEntityToInvoiceLine(entity *neo4jentity.InvoiceLineEntity) *model.InvoiceLine {
	if entity == nil {
		return nil
	}
	return &model.InvoiceLine{
		Metadata: &model.Metadata{
			ID:            entity.Id,
			Created:       entity.CreatedAt,
			LastUpdated:   entity.UpdatedAt,
			Source:        MapDataSourceToModel(entity.Source),
			SourceOfTruth: MapDataSourceToModel(entity.SourceOfTruth),
			AppSource:     entity.AppSource,
		},
		Description: entity.Name,
		Price:       entity.Price,
		Quantity:    int(entity.Quantity),
		Total:       entity.TotalAmount,
		Subtotal:    entity.Amount,
		TaxDue:      entity.Vat,

		//Deprecated all below
		ID:          entity.Id,
		CreatedAt:   entity.CreatedAt,
		Name:        entity.Name,
		Amount:      entity.Amount,
		Vat:         entity.Vat,
		TotalAmount: entity.TotalAmount,
	}
}

func MapEntitiesToInvoices(entities *neo4jentity.InvoiceEntities) []*model.Invoice {
	var output []*model.Invoice
	for _, v := range *entities {
		output = append(output, MapEntityToInvoice(&v))
	}
	return output
}

func MapEntitiesToInvoiceLines(entities *neo4jentity.InvoiceLineEntities) []*model.InvoiceLine {
	var output []*model.InvoiceLine
	for _, v := range *entities {
		output = append(output, MapEntityToInvoiceLine(&v))
	}
	return output
}
