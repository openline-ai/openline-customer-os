package service

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jenum "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	commonpb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/common"
	servicelineitempb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-proto/gen/proto/go/api/grpc/v1/service_line_item"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

type ServiceLineItemService interface {
	Create(ctx context.Context, serviceLineItem *ServiceLineItemCreateData) (string, error)
	Update(ctx context.Context, serviceLineItem *entity.ServiceLineItemEntity, isRetroactiveCorrection bool) error
	Delete(ctx context.Context, serviceLineItemId string) (bool, error)
	GetById(ctx context.Context, id string) (*entity.ServiceLineItemEntity, error)
	GetServiceLineItemsForContracts(ctx context.Context, contractIds []string) (*entity.ServiceLineItemEntities, error)
	Close(ctx context.Context, serviceLineItemId string, endedAt *time.Time) error
	CreateOrUpdateInBulk(ctx context.Context, sliBulkData []*SLIBulkData) ([]string, error)
}
type serviceLineItemService struct {
	log          logger.Logger
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
	services     *Services
}

func NewServiceLineItemService(log logger.Logger, repositories *repository.Repositories, grpcClients *grpc_client.Clients, services *Services) ServiceLineItemService {
	return &serviceLineItemService{
		log:          log,
		repositories: repositories,
		grpcClients:  grpcClients,
		services:     services,
	}
}

type ServiceLineItemCreateData struct {
	ServiceLineItemEntity *entity.ServiceLineItemEntity
	ContractId            string
	ExternalReference     *entity.ExternalSystemEntity
	Source                neo4jentity.DataSource
	AppSource             string
	StartedAt             *time.Time
	EndedAt               *time.Time
}

func (s *serviceLineItemService) Create(ctx context.Context, serviceLineItemDetails *ServiceLineItemCreateData) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItem.Create")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("serviceLineItemDetails", serviceLineItemDetails))

	if serviceLineItemDetails.ServiceLineItemEntity == nil {
		err := fmt.Errorf("service line item entity is nil")
		tracing.TraceErr(span, err)
		return "", err
	}

	serviceLineItemId, err := s.createServiceLineItemWithEvents(ctx, serviceLineItemDetails)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return "", err
	}

	span.LogFields(log.String("output - createdServiceLineItemId", serviceLineItemId))
	return serviceLineItemId, nil
}

func (s *serviceLineItemService) createServiceLineItemWithEvents(ctx context.Context, serviceLineItemDetails *ServiceLineItemCreateData) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.createServiceLineItemWithEvents")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	createServiceLineItemRequest := servicelineitempb.CreateServiceLineItemGrpcRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		ContractId:     serviceLineItemDetails.ContractId,
		Name:           serviceLineItemDetails.ServiceLineItemEntity.Name,
		Quantity:       serviceLineItemDetails.ServiceLineItemEntity.Quantity,
		Price:          serviceLineItemDetails.ServiceLineItemEntity.Price,
		StartedAt:      utils.ConvertTimeToTimestampPtr(serviceLineItemDetails.StartedAt),
		EndedAt:        utils.ConvertTimeToTimestampPtr(serviceLineItemDetails.EndedAt),
		LoggedInUserId: common.GetUserIdFromContext(ctx),
		SourceFields: &commonpb.SourceFields{
			Source:    string(serviceLineItemDetails.Source),
			AppSource: utils.StringFirstNonEmpty(serviceLineItemDetails.AppSource, constants.AppSourceCustomerOsApi),
		},
	}

	switch serviceLineItemDetails.ServiceLineItemEntity.Billed {
	case neo4jenum.BilledTypeMonthly:
		createServiceLineItemRequest.Billed = commonpb.BilledType_MONTHLY_BILLED
	case neo4jenum.BilledTypeQuarterly:
		createServiceLineItemRequest.Billed = commonpb.BilledType_QUARTERLY_BILLED
	case neo4jenum.BilledTypeAnnually:
		createServiceLineItemRequest.Billed = commonpb.BilledType_ANNUALLY_BILLED
	case neo4jenum.BilledTypeOnce:
		createServiceLineItemRequest.Billed = commonpb.BilledType_ONCE_BILLED
	case neo4jenum.BilledTypeUsage:
		createServiceLineItemRequest.Billed = commonpb.BilledType_USAGE_BILLED
	default:
		createServiceLineItemRequest.Billed = commonpb.BilledType_NONE_BILLED
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	response, err := s.grpcClients.ServiceLineItemClient.CreateServiceLineItem(ctx, &createServiceLineItemRequest)
	if err != nil {
		tracing.TraceErr(span, err)
		return "", err
	}

	WaitForObjectCreationAndLogSpan(ctx, s.repositories, response.Id, neo4jutil.NodeLabelServiceLineItem, span)
	return response.Id, err
}

func (s *serviceLineItemService) Update(ctx context.Context, serviceLineItem *entity.ServiceLineItemEntity, isRetroactiveCorrection bool) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.Update")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("serviceLineItem", serviceLineItem))

	if serviceLineItem == nil {
		err := fmt.Errorf("(ServiceLineItemService.Update) service line item entity is nil")
		s.log.Error(err.Error())
		tracing.TraceErr(span, err)
		return err
	} else if serviceLineItem.ID == "" {
		err := fmt.Errorf("(ServiceLineItemService.Update) service line item id is missing")
		s.log.Error(err.Error())
		tracing.TraceErr(span, err)
		return err
	}

	serviceLineItemExists, _ := s.repositories.Neo4jRepositories.CommonReadRepository.ExistsById(ctx, common.GetTenantFromContext(ctx), serviceLineItem.ID, neo4jutil.NodeLabelServiceLineItem)
	if !serviceLineItemExists {
		err := fmt.Errorf("(ServiceLineItemService.Update) service line item with id {%s} not found", serviceLineItem.ID)
		s.log.Error(err.Error())
		tracing.TraceErr(span, err)
		return err
	}

	serviceLineItemUpdateRequest := servicelineitempb.UpdateServiceLineItemGrpcRequest{
		Tenant:                  common.GetTenantFromContext(ctx),
		Id:                      serviceLineItem.ID,
		LoggedInUserId:          common.GetUserIdFromContext(ctx),
		Name:                    serviceLineItem.Name,
		Quantity:                serviceLineItem.Quantity,
		Price:                   serviceLineItem.Price,
		Comments:                serviceLineItem.Comments,
		IsRetroactiveCorrection: isRetroactiveCorrection,
		SourceFields: &commonpb.SourceFields{
			Source:    string(serviceLineItem.Source),
			AppSource: utils.StringFirstNonEmpty(serviceLineItem.AppSource, constants.AppSourceCustomerOsApi),
		},
	}
	switch serviceLineItem.Billed {
	case neo4jenum.BilledTypeMonthly:
		serviceLineItemUpdateRequest.Billed = commonpb.BilledType_MONTHLY_BILLED
	case neo4jenum.BilledTypeQuarterly:
		serviceLineItemUpdateRequest.Billed = commonpb.BilledType_QUARTERLY_BILLED
	case neo4jenum.BilledTypeAnnually:
		serviceLineItemUpdateRequest.Billed = commonpb.BilledType_ANNUALLY_BILLED
	case neo4jenum.BilledTypeOnce:
		serviceLineItemUpdateRequest.Billed = commonpb.BilledType_ONCE_BILLED
	case neo4jenum.BilledTypeUsage:
		serviceLineItemUpdateRequest.Billed = commonpb.BilledType_USAGE_BILLED
	default:
		serviceLineItemUpdateRequest.Billed = commonpb.BilledType_NONE_BILLED
	}
	// set contract id if it's not a retroactive correction
	if !isRetroactiveCorrection {
		contractDbNode, err := s.repositories.Neo4jRepositories.ContractReadRepository.GetContractByServiceLineItemId(ctx, common.GetTenantFromContext(ctx), serviceLineItem.ID)
		if err != nil {
			tracing.TraceErr(span, err)
			s.log.Errorf("Error on getting contract by service line item id {%s}: %s", serviceLineItem.ID, err.Error())
			return err
		}
		if contractDbNode == nil {
			err := fmt.Errorf("contract not found for service line item id {%s}", serviceLineItem.ID)
			tracing.TraceErr(span, err)
			s.log.Errorf(err.Error())
			return err
		}
		serviceLineItemUpdateRequest.ContractId = utils.GetStringPropOrEmpty(utils.GetPropsFromNode(*contractDbNode), "id")
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err := s.grpcClients.ServiceLineItemClient.UpdateServiceLineItem(ctx, &serviceLineItemUpdateRequest)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return err
	}

	return nil
}

func (s *serviceLineItemService) Delete(ctx context.Context, serviceLineItemId string) (completed bool, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.Delete")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId))

	sliExists, err := s.repositories.Neo4jRepositories.CommonReadRepository.ExistsById(ctx, common.GetTenantFromContext(ctx), serviceLineItemId, neo4jutil.NodeLabelServiceLineItem)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("error on checking if service line item exists: %s", err.Error())
		return false, err
	}
	if !sliExists {
		err := fmt.Errorf("service line item with id {%s} not found", serviceLineItemId)
		tracing.TraceErr(span, err)
		s.log.Errorf(err.Error())
		return false, err
	}

	deleteRequest := servicelineitempb.DeleteServiceLineItemGrpcRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		Id:             serviceLineItemId,
		LoggedInUserId: common.GetUserIdFromContext(ctx),
		AppSource:      constants.AppSourceCustomerOsApi,
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err = s.grpcClients.ServiceLineItemClient.DeleteServiceLineItem(ctx, &deleteRequest)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return false, err
	}

	// wait for service line item to be deleted from graph db
	for i := 1; i <= constants.MaxRetriesCheckDataInNeo4jAfterEventRequest; i++ {
		serviceLineItemFound, findErr := s.repositories.Neo4jRepositories.CommonReadRepository.ExistsById(ctx, common.GetTenantFromContext(ctx), serviceLineItemId, neo4jutil.NodeLabelServiceLineItem)
		if findErr != nil {
			tracing.TraceErr(span, findErr)
			s.log.Errorf("error on checking if service line item exists: %s", findErr.Error())
		} else if !serviceLineItemFound {
			span.LogFields(log.Bool("serviceLineItemDeletedFromGraphDb", true))
			return true, nil
		}
		time.Sleep(utils.BackOffIncrementalDelay(i))
	}

	return false, nil
}

func (s *serviceLineItemService) Close(ctx context.Context, serviceLineItemId string, endedAt *time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.Close")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId))

	sliExists, err := s.repositories.Neo4jRepositories.CommonReadRepository.ExistsById(ctx, common.GetTenantFromContext(ctx), serviceLineItemId, neo4jutil.NodeLabelServiceLineItem)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("error on checking if service line item exists: %s", err.Error())
		return err
	}
	if !sliExists {
		err := fmt.Errorf("service line item with id {%s} not found", serviceLineItemId)
		tracing.TraceErr(span, err)
		s.log.Errorf(err.Error())
		return err
	}

	closeRequest := servicelineitempb.CloseServiceLineItemGrpcRequest{
		Tenant:         common.GetTenantFromContext(ctx),
		Id:             serviceLineItemId,
		LoggedInUserId: common.GetUserIdFromContext(ctx),
		AppSource:      constants.AppSourceCustomerOsApi,
		EndedAt:        utils.ConvertTimeToTimestampPtr(endedAt),
	}

	ctx = tracing.InjectSpanContextIntoGrpcMetadata(ctx, span)
	_, err = s.grpcClients.ServiceLineItemClient.CloseServiceLineItem(ctx, &closeRequest)
	if err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Error from events processing: %s", err.Error())
		return err
	}

	return nil
}

func (s *serviceLineItemService) GetById(ctx context.Context, serviceLineItemId string) (*entity.ServiceLineItemEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.GetById")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.String("serviceLineItemId", serviceLineItemId))

	if serviceLineItemDbNode, err := s.repositories.ServiceLineItemRepository.GetById(ctx, common.GetContext(ctx).Tenant, serviceLineItemId); err != nil {
		tracing.TraceErr(span, err)
		wrappedErr := errors.Wrap(err, fmt.Sprintf("service line item with id {%s} not found", serviceLineItemId))
		return nil, wrappedErr
	} else {
		return s.mapDbNodeToServiceLineItemEntity(*serviceLineItemDbNode), nil
	}
}

func (s *serviceLineItemService) GetServiceLineItemsForContracts(ctx context.Context, contractIDs []string) (*entity.ServiceLineItemEntities, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.GetServiceLineItemsForContracts")
	defer span.Finish()
	span.LogFields(log.Object("contractIDs", contractIDs))

	serviceLineItems, err := s.repositories.ServiceLineItemRepository.GetForContracts(ctx, common.GetTenantFromContext(ctx), contractIDs)
	if err != nil {
		return nil, err
	}
	serviceLineItemEntities := make(entity.ServiceLineItemEntities, 0, len(serviceLineItems))
	for _, v := range serviceLineItems {
		serviceLineItemEntity := s.mapDbNodeToServiceLineItemEntity(*v.Node)
		serviceLineItemEntity.DataloaderKey = v.LinkedNodeId
		serviceLineItemEntities = append(serviceLineItemEntities, *serviceLineItemEntity)
	}
	return &serviceLineItemEntities, nil
}

func (s *serviceLineItemService) mapDbNodeToServiceLineItemEntity(dbNode dbtype.Node) *entity.ServiceLineItemEntity {
	props := utils.GetPropsFromNode(dbNode)
	serviceLineItem := entity.ServiceLineItemEntity{
		ID:            utils.GetStringPropOrEmpty(props, "id"),
		Name:          utils.GetStringPropOrEmpty(props, "name"),
		CreatedAt:     utils.GetTimePropOrEpochStart(props, "createdAt"),
		UpdatedAt:     utils.GetTimePropOrEpochStart(props, "updatedAt"),
		StartedAt:     utils.GetTimePropOrEpochStart(props, "startedAt"),
		EndedAt:       utils.GetTimePropOrNil(props, "endedAt"),
		Billed:        neo4jenum.DecodeBilledType(utils.GetStringPropOrEmpty(props, "billed")),
		Price:         utils.GetFloatPropOrZero(props, "price"),
		Quantity:      utils.GetInt64PropOrZero(props, "quantity"),
		Comments:      utils.GetStringPropOrEmpty(props, "comments"),
		Source:        neo4jentity.GetDataSource(utils.GetStringPropOrEmpty(props, "source")),
		SourceOfTruth: neo4jentity.GetDataSource(utils.GetStringPropOrEmpty(props, "sourceOfTruth")),
		AppSource:     utils.GetStringPropOrEmpty(props, "appSource"),
		ParentID:      utils.GetStringPropOrEmpty(props, "parentId"),
	}
	return &serviceLineItem
}

type SLIBulkData struct {
	Id                      string
	Name                    string
	Price                   float64
	Quantity                int64
	Billed                  neo4jenum.BilledType
	Comments                string
	IsRetroactiveCorrection bool
}

func (s *serviceLineItemService) CreateOrUpdateInBulk(ctx context.Context, sliBulkData []*SLIBulkData) ([]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemService.CreateOrUpdateInBulk")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	if len(sliBulkData) == 0 {
		return []string{}, nil
	}

	var responseIds []string

	for _, serviceLineItem := range sliBulkData {
		if serviceLineItem.Id == "" {
			itemId, err := s.Create(ctx, &ServiceLineItemCreateData{
				ServiceLineItemEntity: &entity.ServiceLineItemEntity{
					Name:          serviceLineItem.Name,
					Price:         serviceLineItem.Price,
					Quantity:      serviceLineItem.Quantity,
					Billed:        serviceLineItem.Billed,
					Comments:      serviceLineItem.Comments,
					Source:        neo4jentity.DataSourceOpenline,
					SourceOfTruth: neo4jentity.DataSourceOpenline,
					AppSource:     constants.AppSourceCustomerOsApi,
				},
			})
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error from events processing: %s", err.Error())
				return []string{}, err
			}
			responseIds = append(responseIds, itemId)

		} else {
			responseIds = append(responseIds, serviceLineItem.Id)
			err := s.Update(ctx, &entity.ServiceLineItemEntity{
				ID:        serviceLineItem.Id,
				Name:      serviceLineItem.Name,
				Price:     serviceLineItem.Price,
				Quantity:  serviceLineItem.Quantity,
				Comments:  serviceLineItem.Comments,
				Billed:    serviceLineItem.Billed,
				Source:    neo4jentity.DataSourceOpenline,
				AppSource: constants.AppSourceCustomerOsApi,
			}, serviceLineItem.IsRetroactiveCorrection)
			if err != nil {
				tracing.TraceErr(span, err)
				s.log.Errorf("Error from events processing: %s", err.Error())
				return []string{}, err
			}
		}
	}

	return responseIds, nil
}

func MapServiceLineItemBulkItemsToData(input []*model.ServiceLineItemBulkUpdateItem) []*SLIBulkData {
	var arr []*SLIBulkData
	for _, item := range input {
		sli := MapServiceLineItemBulkItemToData(item)
		if sli != nil {
			arr = append(arr, sli)
		}
	}
	return arr
}

func MapServiceLineItemBulkItemToData(input *model.ServiceLineItemBulkUpdateItem) *SLIBulkData {
	if input == nil {
		return nil
	}
	billed := neo4jenum.BilledTypeNone
	if input.Billed != nil {
		billed = mapper.MapBilledTypeFromModel(*input.Billed)
	}
	return &SLIBulkData{
		Id:                      input.ServiceLineItemID,
		Name:                    utils.IfNotNilString(input.Name),
		Price:                   utils.IfNotNilFloat64(input.Price),
		Quantity:                utils.IfNotNilInt64(input.Quantity),
		Billed:                  billed,
		Comments:                utils.IfNotNilString(input.Comments),
		IsRetroactiveCorrection: utils.IfNotNilBool(input.IsRetroactiveCorrection),
	}
}
