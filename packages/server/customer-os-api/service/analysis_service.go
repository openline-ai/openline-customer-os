package service

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"golang.org/x/exp/slices"
	"golang.org/x/net/context"
)

type AnalysisService interface {
	GetAnalysisById(ctx context.Context, id string) (*entity.AnalysisEntity, error)

	Create(ctx context.Context, newAnalysis *AnalysisCreateData) (*entity.AnalysisEntity, error)
	GetDescribesForAnalysis(ctx context.Context, ids []string) (*entity.AnalysisDescribes, error)
	convertDbNodesAnalysisDescribes(records []*utils.DbNodeAndId) entity.AnalysisDescribes
	mapDbNodeToAnalysisEntity(node dbtype.Node) *entity.AnalysisEntity
}

type AnalysisDescriptionData struct {
	InteractionEventId   *string
	InteractionSessionId *string
}

type AnalysisCreateData struct {
	AnalysisEntity *entity.AnalysisEntity
	Describes      []AnalysisDescriptionData
	Source         entity.DataSource
	SourceOfTruth  entity.DataSource
}

type analysisService struct {
	repositories *repository.Repositories
	services     *Services
}

func NewAnalysisService(repositories *repository.Repositories, services *Services) AnalysisService {
	return &analysisService{
		repositories: repositories,
		services:     services,
	}
}

func (s *analysisService) GetDescribesForAnalysis(ctx context.Context, ids []string) (*entity.AnalysisDescribes, error) {
	records, err := s.repositories.AnalysisRepository.GetDescribesForAnalysis(ctx, common.GetTenantFromContext(ctx), ids)
	if err != nil {
		return nil, err
	}

	analysisDescribes := s.convertDbNodesAnalysisDescribes(records)

	return &analysisDescribes, nil
}

func (s *analysisService) Create(ctx context.Context, newAnalysis *AnalysisCreateData) (*entity.AnalysisEntity, error) {
	session := utils.NewNeo4jWriteSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	interactionEventDbNode, err := session.ExecuteWrite(ctx, s.createAnalysisInDBTxWork(ctx, newAnalysis))
	if err != nil {
		return nil, err
	}
	return s.mapDbNodeToAnalysisEntity(*interactionEventDbNode.(*dbtype.Node)), nil
}

func (s *analysisService) createAnalysisInDBTxWork(ctx context.Context, newAnalysis *AnalysisCreateData) func(tx neo4j.ManagedTransaction) (any, error) {
	return func(tx neo4j.ManagedTransaction) (any, error) {
		tenant := common.GetContext(ctx).Tenant
		analysisDbNode, err := s.repositories.AnalysisRepository.Create(ctx, tx, tenant, *newAnalysis.AnalysisEntity, newAnalysis.Source, newAnalysis.SourceOfTruth)
		if err != nil {
			return nil, err
		}
		var analysisId = utils.GetPropsFromNode(*analysisDbNode)["id"].(string)

		for _, describes := range newAnalysis.Describes {
			if describes.InteractionSessionId != nil {
				err := s.repositories.AnalysisRepository.LinkWithDescribesXXInTx(ctx, tx, tenant, repository.INTERACTION_SESSION, analysisId, *describes.InteractionSessionId)
				if err != nil {
					return nil, err
				}
			}
			if describes.InteractionEventId != nil {
				err := s.repositories.AnalysisRepository.LinkWithDescribesXXInTx(ctx, tx, tenant, repository.INTERACTION_EVENT, analysisId, *describes.InteractionEventId)
				if err != nil {
					return nil, err
				}
			}
		}

		return analysisDbNode, nil
	}
}

func (s *analysisService) GetAnalysisById(ctx context.Context, id string) (*entity.AnalysisEntity, error) {
	session := utils.NewNeo4jReadSession(ctx, s.getNeo4jDriver())
	defer session.Close(ctx)

	queryResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (a:Analysis_%s {id:$id}) RETURN a`,
			common.GetTenantFromContext(ctx)),
			map[string]interface{}{
				"id": id,
			})
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}

	return s.mapDbNodeToAnalysisEntity(queryResult.(dbtype.Node)), nil
}

func (s *analysisService) mapDbNodeToAnalysisEntity(node dbtype.Node) *entity.AnalysisEntity {
	props := utils.GetPropsFromNode(node)
	createdAt := utils.GetTimePropOrEpochStart(props, "createdAt")
	analysisEntity := entity.AnalysisEntity{
		Id:            utils.GetStringPropOrEmpty(props, "id"),
		CreatedAt:     &createdAt,
		AnalysisType:  utils.GetStringPropOrEmpty(props, "analysisType"),
		Content:       utils.GetStringPropOrEmpty(props, "content"),
		ContentType:   utils.GetStringPropOrEmpty(props, "contentType"),
		AppSource:     utils.GetStringPropOrEmpty(props, "appSource"),
		Source:        entity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth: entity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
	}
	return &analysisEntity
}

func (s *analysisService) convertDbNodesToAnalysis(records []*utils.DbNodeAndId) entity.AnalysisEntitys {
	analysises := entity.AnalysisEntitys{}
	for _, v := range records {
		analysis := s.mapDbNodeToAnalysisEntity(*v.Node)
		analysis.DataloaderKey = v.LinkedNodeId
		analysises = append(analysises, *analysis)

	}
	return analysises
}

func (s *analysisService) convertDbNodesAnalysisDescribes(records []*utils.DbNodeAndId) entity.AnalysisDescribes {
	analysisDescribes := entity.AnalysisDescribes{}
	for _, v := range records {
		if slices.Contains(v.Node.Labels, entity.NodeLabel_InteractionSession) {
			sessionEntity := s.services.InteractionSessionService.mapDbNodeToInteractionSessionEntity(*v.Node)
			sessionEntity.DataloaderKey = v.LinkedNodeId
			analysisDescribes = append(analysisDescribes, sessionEntity)
		} else if slices.Contains(v.Node.Labels, entity.NodeLabel_InteractionEvent) {
			participant := s.services.InteractionEventService.mapDbNodeToInteractionEventEntity(*v.Node)
			participant.DataloaderKey = v.LinkedNodeId
			analysisDescribes = append(analysisDescribes, participant)
		}
	}
	return analysisDescribes
}

func MapAnalysisDescriptionInputToDescriptionData(input []*model.AnalysisDescriptionInput) []AnalysisDescriptionData {
	var inputData []AnalysisDescriptionData
	for _, participant := range input {
		inputData = append(inputData, AnalysisDescriptionData{
			InteractionEventId:   participant.InteractionEventID,
			InteractionSessionId: participant.InteractionSessionID,
		})
	}
	return inputData
}
func (s *analysisService) getNeo4jDriver() neo4j.DriverWithContext {
	return *s.repositories.Drivers.Neo4jDriver
}
