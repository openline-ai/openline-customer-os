package service

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

type LocationService interface {
	GetAllForContact(ctx context.Context, contactId string) (*entity.LocationEntities, error)
	GetAllForOrganization(ctx context.Context, organizationId string) (*entity.LocationEntities, error)
}

type locationService struct {
	repositories *repository.Repositories
}

func NewLocationService(repositories *repository.Repositories) LocationService {
	return &locationService{
		repositories: repositories,
	}
}

func (s *locationService) getNeo4jDriver() neo4j.Driver {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *locationService) GetAllForContact(ctx context.Context, contactId string) (*entity.LocationEntities, error) {
	dbNodes, err := s.repositories.LocationRepository.GetAllForContact(common.GetTenantFromContext(ctx), contactId)
	if err != nil {
		return nil, err
	}

	locationEntities := entity.LocationEntities{}
	for _, dbNode := range dbNodes {
		locationEntities = append(locationEntities, *s.mapDbNodeToLocationEntity(*dbNode))
	}
	return &locationEntities, nil
}

func (s *locationService) GetAllForOrganization(ctx context.Context, organizationId string) (*entity.LocationEntities, error) {
	dbNodes, err := s.repositories.LocationRepository.GetAllForOrganization(common.GetContext(ctx).Tenant, organizationId)
	if err != nil {
		return nil, err
	}

	locationEntities := entity.LocationEntities{}
	for _, dbNode := range dbNodes {
		locationEntities = append(locationEntities, *s.mapDbNodeToLocationEntity(*dbNode))
	}
	return &locationEntities, nil
}

func (s *locationService) mapDbNodeToLocationEntity(node dbtype.Node) *entity.LocationEntity {
	props := utils.GetPropsFromNode(node)
	result := entity.LocationEntity{
		Id:        utils.GetStringPropOrEmpty(props, "id"),
		Name:      utils.GetStringPropOrEmpty(props, "name"),
		CreatedAt: utils.GetTimePropOrNow(props, "createdAt"),
		UpdatedAt: utils.GetTimePropOrNow(props, "updatedAt"),
		Source:    entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		AppSource: utils.GetStringPropOrEmpty(props, "appSource"),
	}
	return &result
}
