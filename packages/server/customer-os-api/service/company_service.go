package service

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

type CompanyService interface {
	Create(ctx context.Context, input *entity.CompanyEntity) (*entity.CompanyEntity, error)

	GetCompanyForRole(ctx context.Context, roleId string) (*entity.CompanyEntity, error)
	GetCompanyById(ctx context.Context, companyId string) (*entity.CompanyEntity, error)
	FindCompaniesByNameLike(ctx context.Context, page, limit int, companyName string) (*utils.Pagination, error)
}

type companyService struct {
	repositories *repository.Repositories
}

func NewCompanyService(repositories *repository.Repositories) CompanyService {
	return &companyService{
		repositories: repositories,
	}
}

func (s *companyService) Create(ctx context.Context, input *entity.CompanyEntity) (*entity.CompanyEntity, error) {
	session := utils.NewNeo4jWriteSession(*s.repositories.Drivers.Neo4jDriver)
	defer session.Close()

	dbNodePtr, err := s.repositories.CompanyRepository.Create(session, common.GetContext(ctx).Tenant, input)
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToCompanyEntity(*dbNodePtr), nil
}

func (s *companyService) FindCompaniesByNameLike(ctx context.Context, page, limit int, companyName string) (*utils.Pagination, error) {
	var paginatedResult = utils.Pagination{
		Limit: limit,
		Page:  page,
	}

	dbNodesWithTotalCount, err := s.repositories.CompanyRepository.GetPaginatedCompaniesWithNameLike(common.GetContext(ctx).Tenant, companyName, paginatedResult.GetSkip(), paginatedResult.GetLimit())
	if err != nil {
		return nil, err
	}
	paginatedResult.SetTotalRows(dbNodesWithTotalCount.Count)

	companyEntities := entity.CompanyEntities{}

	for _, v := range dbNodesWithTotalCount.Nodes {
		companyEntities = append(companyEntities, *s.mapDbNodeToCompanyEntity(*v))
	}
	paginatedResult.SetRows(&companyEntities)
	return &paginatedResult, nil
}

func (s *companyService) GetCompanyForRole(ctx context.Context, roleId string) (*entity.CompanyEntity, error) {
	session := utils.NewNeo4jReadSession(*s.repositories.Drivers.Neo4jDriver)
	defer session.Close()

	dbNode, err := s.repositories.CompanyRepository.GetCompanyForRole(session, common.GetContext(ctx).Tenant, roleId)
	if dbNode == nil || err != nil {
		return nil, err
	}
	return s.mapDbNodeToCompanyEntity(*dbNode), nil
}

func (s *companyService) GetCompanyById(ctx context.Context, companyId string) (*entity.CompanyEntity, error) {
	session := utils.NewNeo4jReadSession(*s.repositories.Drivers.Neo4jDriver)
	defer session.Close()

	dbNode, err := s.repositories.CompanyRepository.GetCompanyById(session, common.GetContext(ctx).Tenant, companyId)
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToCompanyEntity(*dbNode), nil
}

func (s *companyService) mapDbNodeToCompanyEntity(node dbtype.Node) *entity.CompanyEntity {
	props := utils.GetPropsFromNode(node)
	companyEntity := new(entity.CompanyEntity)
	companyEntity.Id = utils.GetStringPropOrEmpty(props, "id")
	companyEntity.Name = utils.GetStringPropOrEmpty(props, "name")
	companyEntity.Description = utils.GetStringPropOrEmpty(props, "description")
	companyEntity.Domain = utils.GetStringPropOrEmpty(props, "domain")
	companyEntity.Website = utils.GetStringPropOrEmpty(props, "website")
	companyEntity.Industry = utils.GetStringPropOrEmpty(props, "industry")
	companyEntity.IsPublic = utils.GetBoolPropOrFalse(props, "isPublic")
	companyEntity.CreatedAt = utils.GetTimePropOrNow(props, "createdAt")
	companyEntity.Readonly = utils.GetBoolPropOrFalse(props, "readonly")
	return companyEntity
}
