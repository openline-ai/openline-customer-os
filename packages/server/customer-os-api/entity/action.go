package entity

import (
	"fmt"
	neo4jentity "github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-neo4j-repository/neo4jutil"
	"time"
)

type ActionEntity struct {
	Id            string
	CreatedAt     time.Time
	Content       string
	Metadata      string
	Type          ActionType
	Source        neo4jentity.DataSource
	AppSource     string
	DataloaderKey string
}

type ActionType string

const (
	ActionNA                                        ActionType = ""
	ActionCreated                                   ActionType = "CREATED"
	ActionRenewalLikelihoodUpdated                  ActionType = "RENEWAL_LIKELIHOOD_UPDATED"
	ActionRenewalForecastUpdated                    ActionType = "RENEWAL_FORECAST_UPDATED"
	ActionContractStatusUpdated                     ActionType = "CONTRACT_STATUS_UPDATED"
	ActionServiceLineItemPriceUpdated               ActionType = "SERVICE_LINE_ITEM_PRICE_UPDATED"
	ActionServiceLineItemQuantityUpdated            ActionType = "SERVICE_LINE_ITEM_QUANTITY_UPDATED"
	ActionServiceLineItemBilledTypeUpdated          ActionType = "SERVICE_LINE_ITEM_BILLED_TYPE_UPDATED"
	ActionServiceLineItemBilledTypeRecurringCreated ActionType = "SERVICE_LINE_ITEM_BILLED_TYPE_RECURRING_CREATED"
	ActionServiceLineItemBilledTypeOnceCreated      ActionType = "SERVICE_LINE_ITEM_BILLED_TYPE_ONCE_CREATED"
	ActionServiceLineItemBilledTypeUsageCreated     ActionType = "SERVICE_LINE_ITEM_BILLED_TYPE_USAGE_CREATED"
	ActionContractRenewed                           ActionType = "CONTRACT_RENEWED"
	ActionServiceLineItemRemoved                    ActionType = "SERVICE_LINE_ITEM_REMOVED"
	ActionOnboardingStatusChanged                   ActionType = "ONBOARDING_STATUS_CHANGED"
)

var AllActionType = []ActionType{
	ActionCreated,
	ActionRenewalLikelihoodUpdated,
	ActionRenewalForecastUpdated,
	ActionContractStatusUpdated,
	ActionServiceLineItemPriceUpdated,
	ActionServiceLineItemQuantityUpdated,
	ActionServiceLineItemBilledTypeUpdated,
	ActionServiceLineItemBilledTypeRecurringCreated,
	ActionServiceLineItemBilledTypeOnceCreated,
	ActionServiceLineItemBilledTypeUsageCreated,
	ActionContractRenewed,
	ActionServiceLineItemRemoved,
	ActionOnboardingStatusChanged,
}

func GetActionType(s string) ActionType {
	if IsValidActionType(s) {
		return ActionType(s)
	}
	return ActionNA
}

func IsValidActionType(s string) bool {
	for _, ds := range AllActionType {
		if ds == ActionType(s) {
			return true
		}
	}
	return false
}

func (action ActionEntity) ToString() string {
	return fmt.Sprintf("id: %s\nname: %s", action.Id, action.Type)
}

func (action ActionEntity) GetDataloaderKey() string {
	return action.DataloaderKey
}

func (action *ActionEntity) SetDataloaderKey(key string) {
	action.DataloaderKey = key
}

func (ActionEntity) IsTimelineEvent() {
}

func (ActionEntity) TimelineEventLabel() string {
	return neo4jutil.NodeLabelAction
}

type ActionEntities []ActionEntity

func (action ActionEntity) Labels(tenant string) []string {
	return []string{
		neo4jutil.NodeLabelAction,
		neo4jutil.NodeLabelAction + "_" + tenant,
		neo4jutil.NodeLabelTimelineEvent,
		neo4jutil.NodeLabelTimelineEvent + "_" + tenant,
	}
}
