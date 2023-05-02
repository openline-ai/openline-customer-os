package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	commonEntity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/neo4j/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// GcliSearch is the resolver for the gcli_search field.
func (r *queryResolver) GcliSearch(ctx context.Context, keyword string, limit *int) ([]*model.GCliSearchResultItem, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	searchResultEntities, err := r.Services.SearchService.GCliSearch(ctx, keyword, limit)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed basic search for keyword %s", keyword)
		return nil, err
	}
	result := make([]*model.GCliSearchResultItem, 0)
	for _, v := range *searchResultEntities {
		resultItem := model.GCliSearchResultItem{
			Score:  v.Score,
			Result: new(model.GCliSearchResult),
		}

		switch v.EntityType {
		case entity.SearchResultEntityTypeContact:
			resultItem.Result.ID = v.Node.(*entity.ContactEntity).Id
			resultItem.Result.Type = model.GCliSearchResultTypeContact
			if v.Node.(*entity.ContactEntity).FirstName != "" {
				resultItem.Result.Display = v.Node.(*entity.ContactEntity).FirstName + " " + v.Node.(*entity.ContactEntity).LastName
			} else if v.Node.(*entity.ContactEntity).Name != "" {
				resultItem.Result.Display = v.Node.(*entity.ContactEntity).Name
			} else {
				continue // skip this result
			}
		case entity.SearchResultEntityTypeOrganization:
			resultItem.Result.ID = v.Node.(*entity.OrganizationEntity).ID
			resultItem.Result.Type = model.GCliSearchResultTypeOrganization
			resultItem.Result.Display = v.Node.(*entity.OrganizationEntity).Name
		case entity.SearchResultEntityTypeEmail:
			resultItem.Result.ID = v.Node.(*entity.EmailEntity).Id
			resultItem.Result.Type = model.GCliSearchResultTypeEmail
			resultItem.Result.Display = utils.StringFirstNonEmpty(v.Node.(*entity.EmailEntity).Email, v.Node.(*entity.EmailEntity).RawEmail)
		case entity.SearchResultEntityTypeState:
			resultItem.Result.ID = v.Node.(*commonEntity.StateEntity).Id
			resultItem.Result.Type = model.GCliSearchResultTypeState
			resultItem.Result.Display = v.Node.(*commonEntity.StateEntity).Name
			data := []*model.GCliAttributeKeyValuePair{}
			data = append(data, &model.GCliAttributeKeyValuePair{
				Key:   "code",
				Value: v.Node.(*commonEntity.StateEntity).Code,
			})
			resultItem.Result.Data = data
		}
		result = append(result, &resultItem)
	}
	return result, nil
}
