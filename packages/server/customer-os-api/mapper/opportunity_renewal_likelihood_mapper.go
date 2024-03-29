package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

var opportunityRenewalLikelihoodByModel = map[model.OpportunityRenewalLikelihood]entity.OpportunityRenewalLikelihood{
	model.OpportunityRenewalLikelihoodHighRenewal:   entity.OpportunityRenewalLikelihoodHigh,
	model.OpportunityRenewalLikelihoodMediumRenewal: entity.OpportunityRenewalLikelihoodMedium,
	model.OpportunityRenewalLikelihoodLowRenewal:    entity.OpportunityRenewalLikelihoodLow,
	model.OpportunityRenewalLikelihoodZeroRenewal:   entity.OpportunityRenewalLikelihoodZero,
}

var opportunityRenewalLikelihoodByValue = utils.ReverseMap(opportunityRenewalLikelihoodByModel)

func MapOpportunityRenewalLikelihoodFromModel(input *model.OpportunityRenewalLikelihood) entity.OpportunityRenewalLikelihood {
	if input == nil {
		return ""
	}
	return opportunityRenewalLikelihoodByModel[*input]
}

func MapOpportunityRenewalLikelihoodToModel(input entity.OpportunityRenewalLikelihood) model.OpportunityRenewalLikelihood {
	return opportunityRenewalLikelihoodByValue[input]
}

func MapOpportunityRenewalLikelihoodFromString(input *string) string {
	if input == nil {
		return ""
	}
	if v, exists := opportunityRenewalLikelihoodByModel[model.OpportunityRenewalLikelihood(*input)]; exists {
		return string(v)
	} else {
		return ""
	}
}

func MapOpportunityRenewalLikelihoodToModelPtr(input string) *model.OpportunityRenewalLikelihood {
	switch input {
	case string(entity.OpportunityRenewalLikelihoodHigh):
		return utils.Ptr(model.OpportunityRenewalLikelihoodHighRenewal)
	case string(entity.OpportunityRenewalLikelihoodMedium):
		return utils.Ptr(model.OpportunityRenewalLikelihoodMediumRenewal)
	case string(entity.OpportunityRenewalLikelihoodLow):
		return utils.Ptr(model.OpportunityRenewalLikelihoodLowRenewal)
	case string(entity.OpportunityRenewalLikelihoodZero):
		return utils.Ptr(model.OpportunityRenewalLikelihoodZeroRenewal)
	default:
		return nil
	}
}
