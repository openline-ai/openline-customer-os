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

type ContactRoleService interface {
	FindAllForContact(ctx context.Context, contactId string) (*entity.ContactRoleEntities, error)
	DeleteContactRole(ctx context.Context, contactId, roleId string) (bool, error)
	CreateContactRole(ctx context.Context, contactId string, companyId *string, input *entity.ContactRoleEntity) (*entity.ContactRoleEntity, error)
}

type contactRoleService struct {
	repositories *repository.Repositories
}

func NewContactRoleService(repositories *repository.Repositories) ContactRoleService {
	return &contactRoleService{
		repositories: repositories,
	}
}

func (s *contactRoleService) getDriver() neo4j.Driver {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *contactRoleService) FindAllForContact(ctx context.Context, contactId string) (*entity.ContactRoleEntities, error) {
	session := utils.NewNeo4jReadSession(s.getDriver())
	defer session.Close()

	dbNodes, err := s.repositories.ContactRoleRepository.GetRolesForContact(session, common.GetContext(ctx).Tenant, contactId)
	if err != nil {
		return nil, err
	}

	contactRoleEntities := entity.ContactRoleEntities{}
	for _, dbNode := range dbNodes {
		contactRoleEntities = append(contactRoleEntities, *s.mapDbNodeToContactRoleEntity(dbNode))
	}
	return &contactRoleEntities, nil
}

func (s *contactRoleService) CreateContactRole(ctx context.Context, contactId string, companyId *string, input *entity.ContactRoleEntity) (*entity.ContactRoleEntity, error) {
	session := utils.NewNeo4jWriteSession(*s.repositories.Drivers.Neo4jDriver)
	defer session.Close()

	dbNode, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		if input.Primary == true {
			s.repositories.ContactRoleRepository.SetOtherRolesNonPrimaryInTx(tx, common.GetContext(ctx).Tenant, contactId)
		}

		roleDbNode, err := s.repositories.ContactRoleRepository.CreateContactRole(tx, common.GetContext(ctx).Tenant, contactId, input)
		if err != nil {
			return nil, err
		}
		var roleId = utils.GetPropsFromNode(*roleDbNode)["id"].(string)

		if companyId != nil {
			if err = s.repositories.ContactRoleRepository.LinkWithCompany(tx, common.GetContext(ctx).Tenant, roleId, *companyId); err != nil {
				return nil, err
			}
		}
		return roleDbNode, nil
	})
	if err != nil {
		return nil, err
	}

	return s.mapDbNodeToContactRoleEntity(dbNode.(*dbtype.Node)), nil
}

func (s *contactRoleService) DeleteContactRole(ctx context.Context, contactId, roleId string) (bool, error) {
	session := utils.NewNeo4jWriteSession(*s.repositories.Drivers.Neo4jDriver)
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
		return nil, s.repositories.ContactRoleRepository.DeleteContactRoleInTx(tx, common.GetContext(ctx).Tenant, contactId, roleId)
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *contactRoleService) mapDbNodeToContactRoleEntity(node *dbtype.Node) *entity.ContactRoleEntity {
	props := utils.GetPropsFromNode(*node)
	result := entity.ContactRoleEntity{
		Id:       utils.GetStringPropOrEmpty(props, "id"),
		JobTitle: utils.GetStringPropOrEmpty(props, "jobTitle"),
		Primary:  utils.GetBoolPropOrFalse(props, "primary"),
	}
	return &result
}

//func (s *companyService) UpdateCompanyPosition(ctx context.Context, contactId, companyPositionId string, input *entity.CompanyPositionEntity) (*entity.CompanyPositionEntity, error) {
//var err error
//var companyDbNode *dbtype.Node
//var positionDbRelationship *dbtype.Relationship
//tenant := common.GetContext(ctx).Tenant
//
//session := utils.NewNeo4jWriteSession(*s.repositories.Drivers.Neo4jDriver)
//defer session.Close()
//
//currentPositionDtls, err := s.repositories.CompanyRepository.GetCompanyPositionForContact(session, tenant, contactId, companyPositionId)
//if err != nil {
//	return nil, err
//}
//currentPositionId := utils.GetStringPropOrEmpty(utils.GetPropsFromRelationship(*currentPositionDtls.Position), "id")
//
//updatedPosition, err := session.WriteTransaction(func(tx neo4j.Transaction) (any, error) {
//	if len(input.Company.Id) == 0 {
//		err := s.repositories.CompanyRepository.DeleteCompanyPositionInTx(tx, tenant, contactId, currentPositionId)
//		if err != nil {
//			return nil, err
//		}
//		companyDbNode, positionDbRelationship, err = s.repositories.CompanyRepository.LinkNewCompanyToContactInTx(tx, tenant, contactId, input.Company.Name, input.JobTitle)
//	} else if input.Company.Id == currentPositionId {
//		companyDbNode, positionDbRelationship, err = s.repositories.CompanyRepository.UpdateCompanyPositionInTx(tx, common.GetContext(ctx).Tenant, contactId, companyPositionId, input.JobTitle)
//	} else {
//		err := s.repositories.CompanyRepository.DeleteCompanyPositionInTx(tx, tenant, contactId, currentPositionId)
//		if err != nil {
//			return nil, err
//		}
//		companyDbNode, positionDbRelationship, err = s.repositories.CompanyRepository.LinkExistingCompanyToContactInTx(tx, tenant, contactId, input.Company.Id, input.JobTitle)
//	}
//	companyPositionEntity := s.mapCompanyPositionDbRelationshipToEntity(positionDbRelationship)
//	companyPositionEntity.Company = *s.mapCompanyDbNodeToEntity(companyDbNode)
//	return companyPositionEntity, nil
//})
//
//return updatedPosition.(*entity.CompanyPositionEntity), err
//return nil, nil
//}
