package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// SearchBasic is the resolver for the search_Basic field.
func (r *queryResolver) SearchBasic(ctx context.Context, keyword string) ([]*model.SearchBasicResultItem, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	searchResultEntities, err := r.Services.SearchService.SearchBasic(ctx, keyword)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed basic search for keyword %s", keyword)
		return nil, err
	}
	result := make([]*model.SearchBasicResultItem, 0)
	for _, v := range *searchResultEntities {
		resultItem := model.SearchBasicResultItem{
			Score: v.Score,
		}
		switch v.EntityType {
		case entity.SearchResultEntityTypeContact:
			resultItem.Result = mapper.MapEntityToContact(v.Node.(*entity.ContactEntity))
		case entity.SearchResultEntityTypeOrganization:
			resultItem.Result = mapper.MapEntityToOrganization(v.Node.(*entity.OrganizationEntity))
		case entity.SearchResultEntityTypeEmail:
			resultItem.Result = mapper.MapEntityToEmail(v.Node.(*entity.EmailEntity))
		}
		result = append(result, &resultItem)
	}
	return result, nil
}
