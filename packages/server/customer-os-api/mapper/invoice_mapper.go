package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
)

func MapEntityToInvoice(entity *neo4jentity.InvoiceEntity) *model.Invoice {
	if entity == nil {
		return nil
	}
	return &model.Invoice{
		ID:               entity.Id,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
		Source:           MapDataSourceToModel(entity.Source),
		SourceOfTruth:    MapDataSourceToModel(entity.SourceOfTruth),
		AppSource:        entity.AppSource,
		DryRun:           entity.DryRun,
		Number:           entity.Number,
		Date:             entity.Date,
		DueDate:          entity.DueDate,
		Amount:           entity.Amount,
		Vat:              entity.Vat,
		Total:            entity.Total,
		Currency:         entity.Currency.String(),
		RepositoryFileID: entity.RepositoryFileId,
	}
}

func MapEntityToInvoiceLine(entity *neo4jentity.InvoiceLineEntity) *model.InvoiceLine {
	if entity == nil {
		return nil
	}
	return &model.InvoiceLine{
		ID:        entity.Id,
		CreatedAt: entity.CreatedAt,
		Name:      entity.Name,
		Price:     entity.Price,
		Quantity:  int(entity.Quantity),
		Amount:    entity.Amount,
		Vat:       entity.Vat,
		Total:     entity.Total,
	}
}

func MapEntitiesToInvoiceLines(entities *neo4jentity.InvoiceLineEntities) []*model.InvoiceLine {
	var output []*model.InvoiceLine
	for _, v := range *entities {
		output = append(output, MapEntityToInvoiceLine(&v))
	}
	return output
}