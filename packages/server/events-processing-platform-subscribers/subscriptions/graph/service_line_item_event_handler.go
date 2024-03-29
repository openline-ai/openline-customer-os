package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	neo4jenum "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/enum"
	neo4jmapper "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/mapper"
	neo4jmodel "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/model"
	neo4jrepository "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/helper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/repository"
	contracthandler "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/subscriptions/contract"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform-subscribers/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/event"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/service_line_item/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"strconv"
)

type ServiceLineItemEventHandler struct {
	log          logger.Logger
	repositories *repository.Repositories
	grpcClients  *grpc_client.Clients
}

func NewServiceLineItemEventHandler(log logger.Logger, repositories *repository.Repositories, grpcClients *grpc_client.Clients) *ServiceLineItemEventHandler {
	return &ServiceLineItemEventHandler{
		log:          log,
		repositories: repositories,
		grpcClients:  grpcClients,
	}
}

type userMetadata struct {
	UserId string `json:"user-id"`
}

type ActionPriceMetadata struct {
	UserName        string  `json:"user-name"`
	ServiceName     string  `json:"service-name"`
	Quantity        int64   `json:"quantity"`
	Price           float64 `json:"price"`
	PreviousPrice   float64 `json:"previousPrice"`
	BilledType      string  `json:"billedType"`
	Comment         string  `json:"comment"`
	ReasonForChange string  `json:"reasonForChange"`
}
type ActionQuantityMetadata struct {
	UserName         string  `json:"user-name"`
	ServiceName      string  `json:"service-name"`
	Quantity         int64   `json:"quantity"`
	PreviousQuantity int64   `json:"previousQuantity"`
	Price            float64 `json:"price"`
	BilledType       string  `json:"billedType"`
	Comment          string  `json:"comment"`
	ReasonForChange  string  `json:"reasonForChange"`
}
type ActionBilledTypeMetadata struct {
	UserName           string  `json:"user-name"`
	ServiceName        string  `json:"service-name"`
	Price              float64 `json:"price"`
	Quantity           int64   `json:"quantity"`
	BilledType         string  `json:"billedType"`
	PreviousBilledType string  `json:"previousBilledType"`
	Comment            string  `json:"comment"`
	ReasonForChange    string  `json:"reasonForChange"`
}
type ActionServiceLineItemRemovedMetadata struct {
	UserName    string `json:"user-name"`
	ServiceName string `json:"service-name"`
	Comment     string `json:"comment"`
}

func (h *ServiceLineItemEventHandler) OnCreate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemEventHandler.OnCreate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)
	var user *dbtype.Node
	var userEntity neo4jentity.UserEntity
	var message string
	var name string
	var priceChanged bool
	var quantityChanged bool
	var billedTypeChanged bool
	var eventData event.ServiceLineItemCreateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	serviceLineItemId := aggregate.GetServiceLineItemObjectID(evt.GetAggregateID(), eventData.Tenant)

	isNewVersionForExistingSLI := serviceLineItemId != eventData.ParentId && eventData.PreviousVersionId != ""
	previousPrice := float64(0)
	previousQuantity := int64(0)
	previousVatRate := float64(0)
	previousBilled := ""
	reasonForChange := eventData.Comments
	if isNewVersionForExistingSLI {
		//get the previous service line item to get the previous price and quantity
		previousSliDbNode, err := h.repositories.Neo4jRepositories.ServiceLineItemReadRepository.GetServiceLineItemById(ctx, eventData.Tenant, eventData.PreviousVersionId)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("error while getting latest service line item with parent id %s: %s", eventData.ParentId, err.Error())
		}
		if previousSliDbNode != nil {
			previousServiceLineItem := neo4jmapper.MapDbNodeToServiceLineItemEntity(previousSliDbNode)
			previousPrice = previousServiceLineItem.Price
			previousQuantity = previousServiceLineItem.Quantity
			previousBilled = previousServiceLineItem.Billed.String()
			previousVatRate = previousServiceLineItem.VatRate
			//use the booleans below to create the appropriate action message
			priceChanged = previousServiceLineItem.Price != eventData.Price
			quantityChanged = previousServiceLineItem.Quantity != eventData.Quantity
			billedTypeChanged = previousServiceLineItem.Billed.String() != eventData.Billed
		}
	}
	data := neo4jrepository.ServiceLineItemCreateFields{
		IsNewVersionForExistingSLI: isNewVersionForExistingSLI,
		PreviousQuantity:           previousQuantity,
		PreviousPrice:              previousPrice,
		PreviousBilled:             previousBilled,
		PreviousVatRate:            previousVatRate,
		SourceFields: neo4jmodel.Source{
			Source:        helper.GetSource(eventData.Source.Source),
			SourceOfTruth: helper.GetSourceOfTruth(eventData.Source.SourceOfTruth),
			AppSource:     helper.GetAppSource(eventData.Source.AppSource),
		},
		ContractId: eventData.ContractId,
		ParentId:   eventData.ParentId,
		CreatedAt:  eventData.CreatedAt,
		UpdatedAt:  eventData.UpdatedAt,
		StartedAt:  eventData.StartedAt,
		EndedAt:    eventData.EndedAt,
		Price:      eventData.Price,
		Quantity:   eventData.Quantity,
		Name:       eventData.Name,
		Billed:     eventData.Billed,
		Comments:   eventData.Comments,
		VatRate:    eventData.VatRate,
	}
	err := h.repositories.Neo4jRepositories.ServiceLineItemWriteRepository.CreateForContract(ctx, eventData.Tenant, serviceLineItemId, data)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while saving service line item %s: %s", serviceLineItemId, err.Error())
		return err
	}
	serviceLineItemDbNode, err := h.repositories.Neo4jRepositories.ServiceLineItemReadRepository.GetServiceLineItemById(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while getting service line item by id %s: %s", serviceLineItemId, err.Error())
		return err
	}
	serviceLineItemEntity := neo4jmapper.MapDbNodeToServiceLineItemEntity(serviceLineItemDbNode)

	contractHandler := contracthandler.NewContractHandler(h.log, h.repositories, h.grpcClients)
	err = contractHandler.UpdateActiveRenewalOpportunityArr(ctx, eventData.Tenant, eventData.ContractId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("error while updating renewal opportunity for contract %s: %s", eventData.ContractId, err.Error())
		return nil
	}
	contractDbNode, err := h.repositories.Neo4jRepositories.ContractReadRepository.GetContractById(ctx, eventData.Tenant, eventData.ContractId)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}
	contractEntity := neo4jmapper.MapDbNodeToContractEntity(contractDbNode)

	// get user
	usrMetadata := userMetadata{}
	if err = json.Unmarshal(evt.Metadata, &usrMetadata); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Unmarshal")
	} else {
		if usrMetadata.UserId != "" {
			user, err = h.repositories.Neo4jRepositories.UserReadRepository.GetUserById(ctx, eventData.Tenant, usrMetadata.UserId)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed to get user for service line item %s with userid %s", serviceLineItemId, usrMetadata.UserId)
			}
		}
		userEntity = *neo4jmapper.MapDbNodeToUserEntity(user)
	}
	if eventData.Name == "" {
		name = serviceLineItemEntity.Name
	}
	if serviceLineItemEntity.Name != "" {
		name = serviceLineItemEntity.Name
	}
	if name == "" {
		name = "Unnamed service"
	}

	userName := userEntity.GetFullName()
	metadataPrice, err := utils.ToJson(ActionPriceMetadata{
		UserName:        userName,
		ServiceName:     name,
		Quantity:        eventData.Quantity,
		BilledType:      eventData.Billed,
		PreviousPrice:   previousPrice,
		Price:           eventData.Price,
		Comment:         "price is " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + " for service " + name,
		ReasonForChange: reasonForChange,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed to serialize price metadata: %s", err.Error())
		return errors.Wrap(err, "Failed to serialize price metadata")
	}
	metadataQuantity, err := utils.ToJson(ActionQuantityMetadata{
		UserName:         userName,
		ServiceName:      name,
		Price:            eventData.Price,
		PreviousQuantity: previousQuantity,
		Quantity:         eventData.Quantity,
		BilledType:       eventData.Billed,
		Comment:          "quantity is " + strconv.FormatInt(serviceLineItemEntity.Quantity, 10) + " for service " + name,
		ReasonForChange:  reasonForChange,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed to serialize quantity metadata: %s", err.Error())
		return errors.Wrap(err, "Failed to serialize quantity metadata")
	}
	metadataBilledType, err := utils.ToJson(ActionBilledTypeMetadata{
		UserName:           userName,
		ServiceName:        name,
		BilledType:         eventData.Billed,
		PreviousBilledType: previousBilled,
		Quantity:           eventData.Quantity,
		Price:              eventData.Price,
		Comment:            "billed type is " + serviceLineItemEntity.Billed.String() + " for service " + name,
		ReasonForChange:    reasonForChange,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed to serialize billed type metadata: %s", err.Error())
		return errors.Wrap(err, "Failed to serialize billed type metadata")
	}
	extraActionProperties := map[string]interface{}{
		"comments": reasonForChange,
	}
	cycle := getBillingCycleNamingConvention(eventData.Billed)
	previousCycle := getBillingCycleNamingConvention(previousBilled)
	if previousCycle == "" {
		previousCycle = getBillingCycleNamingConvention(serviceLineItemEntity.Billed.String())
	}

	if !isNewVersionForExistingSLI {
		if serviceLineItemEntity.Billed.String() == model.AnnuallyBilled.String() || serviceLineItemEntity.Billed.String() == model.QuarterlyBilled.String() || serviceLineItemEntity.Billed.String() == model.MonthlyBilled.String() {
			message = userName + " added a recurring service to " + contractEntity.Name + ": " + name + " at " + strconv.FormatInt(serviceLineItemEntity.Quantity, 10) + " x " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + "/" + cycle
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, eventData.ContractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemBilledTypeRecurringCreated, message, metadataBilledType, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating recurring billed type service line item created action for contract %s: %s", eventData.ContractId, err.Error())
			}
		}
		if serviceLineItemEntity.Billed.String() == model.OnceBilled.String() {
			message = userName + " added an one time service to " + contractEntity.Name + ": " + name + " at " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price)
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, eventData.ContractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemBilledTypeOnceCreated, message, metadataBilledType, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating once billed type service line item created action for contract %s: %s", eventData.ContractId, err.Error())
			}
		}
		if serviceLineItemEntity.Billed.String() == model.UsageBilled.String() {
			message = userName + " added a per use service to " + contractEntity.Name + ": " + name + " at " + fmt.Sprintf("%.4f", serviceLineItemEntity.Price)
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, eventData.ContractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemBilledTypeUsageCreated, message, metadataBilledType, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating per use billed type service line item created action for contract %s: %s", eventData.ContractId, err.Error())
			}
		}
	}
	if isNewVersionForExistingSLI {
		if priceChanged && (eventData.Billed == model.AnnuallyBilled.String() || eventData.Billed == model.QuarterlyBilled.String() || eventData.Billed == model.MonthlyBilled.String()) {
			if eventData.Price > previousPrice {
				message = userName + " increased the price for " + name + " from " + fmt.Sprintf("%.2f", previousPrice) + "/" + previousCycle + " to " + fmt.Sprintf("%.2f", eventData.Price) + "/" + cycle
			}
			if eventData.Price < previousPrice {
				message = userName + " decreased the price for " + name + " from " + fmt.Sprintf("%.2f", previousPrice) + "/" + previousCycle + " to " + fmt.Sprintf("%.2f", eventData.Price) + "/" + cycle
			}
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractEntity.Id, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemPriceUpdated, message, metadataPrice, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating price update action for contract service line item %s: %s", contractEntity.Id, err.Error())
			}
		}

		if priceChanged && eventData.Billed == model.OnceBilled.String() {
			if eventData.Price > previousPrice {
				message = userName + " increased the price for " + name + " from " + fmt.Sprintf("%.2f", previousPrice) + " to " + fmt.Sprintf("%.2f", eventData.Price)
			}
			if eventData.Price < serviceLineItemEntity.Price {
				message = userName + " decreased the price for " + name + " from " + fmt.Sprintf("%.2f", previousPrice) + " to " + fmt.Sprintf("%.2f", eventData.Price)
			}
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractEntity.Id, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemPriceUpdated, message, metadataPrice, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating price update action for contract service line item %s: %s", contractEntity.Id, err.Error())
			}
		}
		if priceChanged && eventData.Billed == model.UsageBilled.String() {
			if eventData.Price > previousPrice {
				message = userName + " increased the price for " + name + " from " + fmt.Sprintf("%.4f", previousPrice) + " to " + fmt.Sprintf("%.4f", eventData.Price)
			}
			if eventData.Price < serviceLineItemEntity.Price {
				message = userName + " decreased the price for " + name + " from " + fmt.Sprintf("%.4f", previousPrice) + " to " + fmt.Sprintf("%.4f", eventData.Price)
			}
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractEntity.Id, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemPriceUpdated, message, metadataPrice, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating price update action for contract service line item %s: %s", contractEntity.Id, err.Error())
			}
		}
		if quantityChanged {
			if eventData.Quantity > previousQuantity {
				message = userName + " increased the quantity of " + name + " from " + strconv.FormatInt(previousQuantity, 10) + " to " + strconv.FormatInt(eventData.Quantity, 10)
			}
			if eventData.Quantity < previousQuantity {
				message = userName + " decreased the quantity of " + name + " from " + strconv.FormatInt(previousQuantity, 10) + " to " + strconv.FormatInt(eventData.Quantity, 10)
			}
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractEntity.Id, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemQuantityUpdated, message, metadataQuantity, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating quantity update action for contract service line item %s: %s", contractEntity.Id, err.Error())
			}
		}
		if billedTypeChanged && previousBilled != "" {
			message = userName + " changed the billing cycle for " + name + " from " + fmt.Sprintf("%.2f", previousPrice) + "/" + previousCycle + " to " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + "/" + cycle
			_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractEntity.Id, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemBilledTypeUpdated, message, metadataBilledType, utils.Now(), extraActionProperties)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed creating billed type update action for contract service line item %s: %s", contractEntity.Id, err.Error())
			}
		}
	}

	return nil
}

func (h *ServiceLineItemEventHandler) OnUpdate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemEventHandler.OnUpdate")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)
	var contractId string
	var user *dbtype.Node
	var userEntity neo4jentity.UserEntity
	var name string
	var message string
	var eventData event.ServiceLineItemUpdateEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	serviceLineItemId := aggregate.GetServiceLineItemObjectID(evt.GetAggregateID(), eventData.Tenant)
	serviceLineItemDbNode, err := h.repositories.Neo4jRepositories.ServiceLineItemReadRepository.GetServiceLineItemById(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}
	serviceLineItemEntity := neo4jmapper.MapDbNodeToServiceLineItemEntity(serviceLineItemDbNode)
	//we will use the following booleans below to check if the price, quantity, billed type has changed
	priceChanged := serviceLineItemEntity.Price != eventData.Price
	quantityChanged := serviceLineItemEntity.Quantity != eventData.Quantity
	billedTypeChanged := serviceLineItemEntity.Billed.String() != eventData.Billed

	data := neo4jrepository.ServiceLineItemUpdateFields{
		Price:     eventData.Price,
		Quantity:  eventData.Quantity,
		Billed:    eventData.Billed,
		Comments:  eventData.Comments,
		Name:      eventData.Name,
		Source:    helper.GetSource(eventData.Source.Source),
		UpdatedAt: eventData.UpdatedAt,
		VatRate:   eventData.VatRate,
		StartedAt: eventData.StartedAt,
	}
	err = h.repositories.Neo4jRepositories.ServiceLineItemWriteRepository.Update(ctx, eventData.Tenant, serviceLineItemId, data)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while updating service line item %s: %s", serviceLineItemId, err.Error())
		return err
	}

	contractDbNode, err := h.repositories.Neo4jRepositories.ContractReadRepository.GetContractByServiceLineItemId(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("error while getting contract for service line item %s: %s", serviceLineItemId, err.Error())
		return nil
	}
	if contractDbNode != nil {
		contract := neo4jmapper.MapDbNodeToContractEntity(contractDbNode)
		contractId = contract.Id
		contractHandler := contracthandler.NewContractHandler(h.log, h.repositories, h.grpcClients)
		err = contractHandler.UpdateActiveRenewalOpportunityArr(ctx, eventData.Tenant, contract.Id)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("error while updating renewal opportunity for contract %s: %s", contract.Id, err.Error())
			return nil
		}
	}
	//check to make sure the name displays correctly in the action message
	if eventData.Name == "" {
		name = serviceLineItemEntity.Name
	} else {
		name = eventData.Name
	}
	if name == "" {
		name = "Unnamed service"
	}
	// get user
	usrMetadata := userMetadata{}
	if err = json.Unmarshal(evt.Metadata, &usrMetadata); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Unmarshal")
	} else {
		if usrMetadata.UserId != "" {
			user, err = h.repositories.Neo4jRepositories.UserReadRepository.GetUserById(ctx, eventData.Tenant, usrMetadata.UserId)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed to get user for service line item %s with userid %s", serviceLineItemId, usrMetadata.UserId)
			}
		}
		userEntity = *neo4jmapper.MapDbNodeToUserEntity(user)
	}

	metadataPrice, err := utils.ToJson(ActionPriceMetadata{
		UserName:        userEntity.GetFullName(),
		ServiceName:     serviceLineItemEntity.Name,
		Price:           eventData.Price,
		PreviousPrice:   serviceLineItemEntity.Price,
		BilledType:      serviceLineItemEntity.Billed.String(),
		Quantity:        serviceLineItemEntity.Quantity,
		Comment:         "price changed is " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + " for service " + name,
		ReasonForChange: eventData.Comments,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed to serialize price metadata: %s", err.Error())
		return errors.Wrap(err, "Failed to serialize price metadata")
	}
	metadataQuantity, err := utils.ToJson(ActionQuantityMetadata{
		UserName:         userEntity.GetFullName(),
		ServiceName:      serviceLineItemEntity.Name,
		PreviousQuantity: serviceLineItemEntity.Quantity,
		Quantity:         eventData.Quantity,
		Price:            serviceLineItemEntity.Price,
		BilledType:       serviceLineItemEntity.Billed.String(),
		Comment:          "quantity changed is " + strconv.FormatInt(serviceLineItemEntity.Quantity, 10) + " for service " + name,
		ReasonForChange:  eventData.Comments,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed to serialize quantity metadata: %s", err.Error())
		return errors.Wrap(err, "Failed to serialize quantity metadata")
	}
	metadataBilledType, err := utils.ToJson(ActionBilledTypeMetadata{
		UserName:           userEntity.GetFullName(),
		ServiceName:        serviceLineItemEntity.Name,
		BilledType:         eventData.Billed,
		PreviousBilledType: serviceLineItemEntity.Billed.String(),
		Quantity:           serviceLineItemEntity.Quantity,
		Price:              serviceLineItemEntity.Price,
		Comment:            "billed type changed is " + serviceLineItemEntity.Billed.String() + " for service " + name,
		ReasonForChange:    eventData.Comments,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed to serialize billed type metadata: %s", err.Error())
		return errors.Wrap(err, "Failed to serialize billed type metadata")
	}
	oldCycle := getBillingCycleNamingConvention(serviceLineItemEntity.Billed.String())
	cycle := getBillingCycleNamingConvention(eventData.Billed)
	extraActionProperties := map[string]interface{}{
		"comments": eventData.Comments,
	}

	if priceChanged && (eventData.Billed == model.AnnuallyBilled.String() || eventData.Billed == model.QuarterlyBilled.String() || eventData.Billed == model.MonthlyBilled.String()) {
		if eventData.Price > serviceLineItemEntity.Price {
			message = userEntity.GetFullName() + " retroactively increased the price for " + name + " from " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + "/" + oldCycle + " to " + fmt.Sprintf("%.2f", eventData.Price) + "/" + cycle
		}
		if eventData.Price < serviceLineItemEntity.Price {
			message = userEntity.GetFullName() + " retroactively decreased the price for " + name + " from " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + "/" + oldCycle + " to " + fmt.Sprintf("%.2f", eventData.Price) + "/" + cycle
		}
		_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemPriceUpdated, message, metadataPrice, utils.Now(), extraActionProperties)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Failed creating price update action for contract service line item %s: %s", contractId, err.Error())
		}
	}

	if priceChanged && eventData.Billed == model.OnceBilled.String() {
		if eventData.Price > serviceLineItemEntity.Price {
			message = userEntity.GetFullName() + " retroactively increased the price for " + name + " from " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + " to " + fmt.Sprintf("%.2f", eventData.Price)
		}
		if eventData.Price < serviceLineItemEntity.Price {
			message = userEntity.GetFullName() + " retroactively decreased the price for " + name + " from " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + " to " + fmt.Sprintf("%.2f", eventData.Price)
		}
		_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemPriceUpdated, message, metadataPrice, utils.Now(), extraActionProperties)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Failed creating price update action for contract service line item %s: %s", contractId, err.Error())
		}
	}
	if priceChanged && eventData.Billed == model.UsageBilled.String() {
		if eventData.Price > serviceLineItemEntity.Price {
			message = userEntity.GetFullName() + " retroactively increased the price for " + name + " from " + fmt.Sprintf("%.4f", serviceLineItemEntity.Price) + " to " + fmt.Sprintf("%.4f", eventData.Price)
		}
		if eventData.Price < serviceLineItemEntity.Price {
			message = userEntity.GetFullName() + " retroactively decreased the price for " + name + " from " + fmt.Sprintf("%.4f", serviceLineItemEntity.Price) + " to " + fmt.Sprintf("%.4f", eventData.Price)
		}
		_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemPriceUpdated, message, metadataPrice, utils.Now(), extraActionProperties)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Failed creating price update action for contract service line item %s: %s", contractId, err.Error())
		}
	}

	if quantityChanged {
		if eventData.Quantity > serviceLineItemEntity.Quantity {
			message = userEntity.GetFullName() + " retroactively increased the quantity of " + name + " from " + strconv.FormatInt(serviceLineItemEntity.Quantity, 10) + " to " + strconv.FormatInt(eventData.Quantity, 10)
		}
		if eventData.Quantity < serviceLineItemEntity.Quantity {
			message = userEntity.GetFullName() + " retroactively decreased the quantity of " + name + " from " + strconv.FormatInt(serviceLineItemEntity.Quantity, 10) + " to " + strconv.FormatInt(eventData.Quantity, 10)
		}
		_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemQuantityUpdated, message, metadataQuantity, utils.Now(), extraActionProperties)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Failed creating quantity update action for contract service line item %s: %s", contractId, err.Error())
		}
	}
	if billedTypeChanged && serviceLineItemEntity.Billed != "" {
		message = userEntity.GetFullName() + " changed the billing cycle for " + name + " from " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + "/" + oldCycle + " to " + fmt.Sprintf("%.2f", serviceLineItemEntity.Price) + "/" + cycle
		_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.CreateWithProperties(ctx, eventData.Tenant, contractId, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemBilledTypeUpdated, message, metadataBilledType, utils.Now(), extraActionProperties)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("Failed creating billed type update action for contract service line item %s: %s", contractId, err.Error())
		}
	}
	return nil
}

func (h *ServiceLineItemEventHandler) OnDelete(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemEventHandler.OnDelete")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)
	var user *dbtype.Node
	var userEntity neo4jentity.UserEntity
	var serviceLineItemName string
	var contractName string

	var eventData event.ServiceLineItemDeleteEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	serviceLineItemId := aggregate.GetServiceLineItemObjectID(evt.GetAggregateID(), eventData.Tenant)
	serviceLineItemDbNode, err := h.repositories.Neo4jRepositories.ServiceLineItemReadRepository.GetServiceLineItemById(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}
	serviceLineItemEntity := neo4jmapper.MapDbNodeToServiceLineItemEntity(serviceLineItemDbNode)
	if serviceLineItemEntity.Name != "" {
		serviceLineItemName = serviceLineItemEntity.Name
	} else {
		serviceLineItemName = "Unnamed service"
	}

	// get user
	usrMetadata := userMetadata{}
	if err := json.Unmarshal(evt.Metadata, &usrMetadata); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "json.Unmarshal")
	} else {
		if usrMetadata.UserId != "" {
			user, err = h.repositories.Neo4jRepositories.UserReadRepository.GetUserById(ctx, eventData.Tenant, usrMetadata.UserId)
			if err != nil {
				tracing.TraceErr(span, err)
				h.log.Errorf("Failed to get user for service line item %s with userid %s", serviceLineItemId, usrMetadata.UserId)
			}
		}
		userEntity = *neo4jmapper.MapDbNodeToUserEntity(user)
	}

	contractDbNode, err := h.repositories.Neo4jRepositories.ContractReadRepository.GetContractByServiceLineItemId(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("error while getting contract for service line item %s: %s", serviceLineItemId, err.Error())
		return nil
	}
	contract := neo4jmapper.MapDbNodeToContractEntity(contractDbNode)
	if contract.Name != "" {
		contractName = contract.Name
	} else {
		contractName = "Unnamed contract"
	}

	err = h.repositories.Neo4jRepositories.ServiceLineItemWriteRepository.Delete(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while deleting service line item %s: %s", serviceLineItemId, err.Error())
		return err
	}

	if contractDbNode != nil {
		contractHandler := contracthandler.NewContractHandler(h.log, h.repositories, h.grpcClients)
		err = contractHandler.UpdateActiveRenewalOpportunityArr(ctx, eventData.Tenant, contract.Id)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("error while updating renewal opportunity for contract %s: %s", contract.Id, err.Error())
			return nil
		}
	}
	metadata, err := utils.ToJson(ActionServiceLineItemRemovedMetadata{
		UserName:    userEntity.GetFullName(),
		ServiceName: serviceLineItemName,
		Comment:     "service line item removed is " + serviceLineItemName + " from " + contractName + " by " + userEntity.GetFullName(),
	})
	message := userEntity.GetFullName() + " removed " + serviceLineItemName + " from " + contractName

	_, err = h.repositories.Neo4jRepositories.ActionWriteRepository.Create(ctx, eventData.Tenant, contract.Id, neo4jenum.CONTRACT, neo4jenum.ActionServiceLineItemRemoved, message, metadata, utils.Now())
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Failed remove service line item action for contract %s: %s", contract.Id, err.Error())
	}

	return nil
}

func (h *ServiceLineItemEventHandler) OnClose(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceLineItemEventHandler.OnClose")
	defer span.Finish()
	setEventSpanTagsAndLogFields(span, evt)

	var eventData event.ServiceLineItemCloseEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	serviceLineItemId := aggregate.GetServiceLineItemObjectID(evt.GetAggregateID(), eventData.Tenant)
	err := h.repositories.Neo4jRepositories.ServiceLineItemWriteRepository.Close(ctx, eventData.Tenant, serviceLineItemId, eventData.UpdatedAt, eventData.EndedAt, eventData.IsCanceled)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("Error while closing service line item %s: %s", serviceLineItemId, err.Error())
		return err
	}

	contractDbNode, err := h.repositories.Neo4jRepositories.ContractReadRepository.GetContractByServiceLineItemId(ctx, eventData.Tenant, serviceLineItemId)
	if err != nil {
		tracing.TraceErr(span, err)
		h.log.Errorf("error while getting contract for service line item %s: %s", serviceLineItemId, err.Error())
		return nil
	}
	if contractDbNode != nil {
		contract := neo4jmapper.MapDbNodeToContractEntity(contractDbNode)
		contractHandler := contracthandler.NewContractHandler(h.log, h.repositories, h.grpcClients)
		err = contractHandler.UpdateActiveRenewalOpportunityArr(ctx, eventData.Tenant, contract.Id)
		if err != nil {
			tracing.TraceErr(span, err)
			h.log.Errorf("error while updating renewal opportunity for contract %s: %s", contract.Id, err.Error())
			return nil
		}
	}

	return nil
}

func getBillingCycleNamingConvention(billedType string) string {
	switch billedType {
	case model.AnnuallyBilled.String():
		return "year"
	case model.QuarterlyBilled.String():
		return "quarter"
	case model.MonthlyBilled.String():
		return "month"
	default:
		return ""
	}
}
