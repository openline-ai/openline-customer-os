package service

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	entityDashboard "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity/dashboard"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"math"
	"time"
)

type DashboardViewOrganizationsRequest struct {
	Where *model.Filter
	Sort  *model.SortBy
	Page  int
	Limit int
}

type DashboardService interface {
	GetDashboardViewOrganizationsData(ctx context.Context, requestDetails DashboardViewOrganizationsRequest) (*utils.Pagination, error)

	GetDashboardCustomerMapData(ctx context.Context) ([]*entityDashboard.DashboardCustomerMapData, error)
	GetDashboardMRRPerCustomerData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardDashboardMRRPerCustomerData, error)
	GetDashboardGrossRevenueRetentionData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardGrossRevenueRetentionData, error)
	GetDashboardARRBreakdownData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardARRBreakdownData, error)
	GetDashboardRevenueAtRiskData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardRevenueAtRiskData, error)
	GetDashboardRetentionRateData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardRetentionRateData, error)
	GetDashboardNewCustomersData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardNewCustomersData, error)
	GetDashboardAverageTimeToOnboardPerMonth(ctx context.Context, start, end time.Time) (*model.DashboardTimeToOnboard, error)
}

type dashboardService struct {
	log          logger.Logger
	repositories *repository.Repositories
	services     *Services
}

func NewDashboardService(log logger.Logger, repositories *repository.Repositories, services *Services) DashboardService {
	return &dashboardService{
		log:          log,
		repositories: repositories,
		services:     services,
	}
}

func (s *dashboardService) getNeo4jDriver() neo4j.DriverWithContext {
	return *s.repositories.Drivers.Neo4jDriver
}

func (s *dashboardService) GetDashboardViewOrganizationsData(ctx context.Context, requestDetails DashboardViewOrganizationsRequest) (*utils.Pagination, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardViewOrganizationsData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Int("page", requestDetails.Page), log.Int("limit", requestDetails.Limit))
	if requestDetails.Where != nil {
		span.LogFields(log.Object("filter", *requestDetails.Where))
	}
	if requestDetails.Sort != nil {
		span.LogFields(log.Object("sort", *requestDetails.Sort))
	}

	var paginatedResult = utils.Pagination{
		Limit: requestDetails.Limit,
		Page:  requestDetails.Page,
	}

	dbNodes, err := s.repositories.DashboardRepository.GetDashboardViewOrganizationData(ctx, common.GetContext(ctx).Tenant, paginatedResult.GetSkip(), paginatedResult.GetLimit(), requestDetails.Where, requestDetails.Sort)
	if err != nil {
		return nil, err
	}
	paginatedResult.SetTotalRows(dbNodes.Count)

	organizationEntities := entity.OrganizationEntities{}

	for _, v := range dbNodes.Nodes {
		organizationEntities = append(organizationEntities, *s.services.OrganizationService.mapDbNodeToOrganizationEntity(*v))
	}

	paginatedResult.SetRows(&organizationEntities)
	return &paginatedResult, nil
}

func (s *dashboardService) GetDashboardCustomerMapData(ctx context.Context) ([]*entityDashboard.DashboardCustomerMapData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardCustomerMapData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)

	response := make([]*entityDashboard.DashboardCustomerMapData, 0)

	data, err := s.repositories.DashboardRepository.GetDashboardCustomerMapData(ctx, common.GetContext(ctx).Tenant)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		organizationId, _ := record["organizationId"].(string)
		oldestServiceStartedAt, _ := record["oldestServiceStartedAt"].(time.Time)
		arr, _ := record["arr"].(float64)
		state, _ := record["state"].(string)

		response = append(response, &entityDashboard.DashboardCustomerMapData{
			OrganizationId:     organizationId,
			ContractSignedDate: oldestServiceStartedAt,
			State:              mapDashboardCustomerMapStateFromString(state),
			Arr:                arr,
		})
	}

	return response, nil
}

func mapDashboardCustomerMapStateFromString(state string) entityDashboard.DashboardCustomerMapState {
	switch state {
	case "OK":
		return entityDashboard.DashboardCustomerMapStateOk
	case "AT_RISK":
		return entityDashboard.DashboardCustomerMapStateAtRisk
	case "CHURNED":
		return entityDashboard.DashboardCustomerMapStateChurned
	default:
		return ""
	}
}

func (s *dashboardService) GetDashboardMRRPerCustomerData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardDashboardMRRPerCustomerData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardMRRPerCustomerData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start))
	span.LogFields(log.Object("end", end))

	response := entityDashboard.DashboardDashboardMRRPerCustomerData{}

	countCustomers, err := s.repositories.OrganizationRepository.CountCustomers(ctx, common.GetContext(ctx).Tenant)
	if err != nil {
		return nil, err
	}

	data, err := s.repositories.DashboardRepository.GetDashboardMRRPerCustomerData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		amountPerMonth, _ := record["amountPerMonth"].(float64)

		if amountPerMonth > 0 && countCustomers > 0 {
			amountPerMonth = amountPerMonth / float64(countCustomers)
		}

		newData := &entityDashboard.DashboardDashboardMRRPerCustomerPerMonthData{
			Year:  int(year),
			Month: int(month),
			Value: amountPerMonth,
		}

		response.Months = append(response.Months, newData)
	}

	currentMonth := 0.0
	previousMonth := 0.0

	if len(response.Months) == 1 {
		currentMonth = response.Months[len(response.Months)-1].Value
	} else if len(response.Months) > 1 {
		currentMonth = response.Months[len(response.Months)-1].Value
		previousMonth = response.Months[len(response.Months)-2].Value
	}

	response.MrrPerCustomer = currentMonth
	response.IncreasePercentage = ComputeNumbersDisplay(float64(previousMonth), float64(currentMonth))

	return &response, nil
}

func (s *dashboardService) GetDashboardGrossRevenueRetentionData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardGrossRevenueRetentionData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardGrossRevenueRetentionData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start))
	span.LogFields(log.Object("end", end))

	response := entityDashboard.DashboardGrossRevenueRetentionData{}

	//response.GrossRevenueRetention = 85
	//response.IncreasePercentage = 5.4

	current := start
	for current.Before(end) || current.Equal(end) {
		fmt.Println(current.Month(), current.Year())

		newData := &entityDashboard.DashboardGrossRevenueRetentionPerMonthData{
			Month:      int(current.Month()),
			Percentage: 0,
		}

		response.Months = append(response.Months, newData)

		current = current.AddDate(0, 1, 0)
	}

	contractsRenewalsData, err := s.repositories.DashboardRepository.GetDashboardRetentionRateContractsRenewalsData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range contractsRenewalsData {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		value, _ := record["value"].(float64)

		for _, monthData := range response.Months {
			if monthData.Year == int(year) && monthData.Month == int(month) {
				monthData.Percentage = value
			}
		}
	}

	return &response, nil
}

func (s *dashboardService) GetDashboardARRBreakdownData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardARRBreakdownData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardARRBreakdownData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start))
	span.LogFields(log.Object("end", end))

	response := entityDashboard.DashboardARRBreakdownData{}

	response.ArrBreakdown = 0
	response.IncreasePercentage = 0

	data, err := s.repositories.DashboardRepository.GetDashboardARRBreakdownData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		newlyContracted, _ := record["newlyContracted"].(float64)
		renewals, _ := record["renewals"].(float64)
		upsells, _ := record["upsells"].(float64)
		downgrades, _ := record["downgrades"].(float64)
		cancellations, _ := record["cancellations"].(float64)
		churned, _ := record["churned"].(float64)

		newData := &entityDashboard.DashboardARRBreakdownPerMonthData{
			Year:            int(year),
			Month:           int(month),
			NewlyContracted: newlyContracted,
			Renewals:        renewals,
			Upsells:         upsells,
			Downgrades:      downgrades,
			Cancellations:   cancellations,
			Churned:         churned,
		}

		response.Months = append(response.Months, newData)
	}

	upsells, err := s.repositories.DashboardRepository.GetDashboardARRBreakdownUpsellsAndDowngradesData(ctx, common.GetContext(ctx).Tenant, "UPSELLS", start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range upsells {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		value, _ := record["value"].(float64)

		for _, monthData := range response.Months {
			if monthData.Year == int(year) && monthData.Month == int(month) {
				monthData.Upsells = value
			}
		}
	}

	downgrades, err := s.repositories.DashboardRepository.GetDashboardARRBreakdownUpsellsAndDowngradesData(ctx, common.GetContext(ctx).Tenant, "DOWNGRADES", start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range downgrades {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		value, _ := record["value"].(float64)

		for _, monthData := range response.Months {
			if monthData.Year == int(year) && monthData.Month == int(month) {
				monthData.Downgrades = value
			}
		}
	}

	renewals, err := s.repositories.DashboardRepository.GetDashboardARRBreakdownRenewalsData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range renewals {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		value, _ := record["value"].(float64)

		for _, monthData := range response.Months {
			if monthData.Year == int(year) && monthData.Month == int(month) {
				monthData.Renewals = value
			}
		}
	}

	return &response, nil
}

func (s *dashboardService) GetDashboardRevenueAtRiskData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardRevenueAtRiskData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardRevenueAtRiskData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start))
	span.LogFields(log.Object("end", end))

	response := entityDashboard.DashboardRevenueAtRiskData{}

	data, err := s.repositories.DashboardRepository.GetDashboardRevenueAtRiskData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	high, _ := data[0]["high"].(float64)
	atRisk, _ := data[0]["atRisk"].(float64)

	response.HighConfidence = high
	response.AtRisk = atRisk

	return &response, nil
}

func (s *dashboardService) GetDashboardRetentionRateData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardRetentionRateData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardRetentionRateData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start))
	span.LogFields(log.Object("end", end))

	response := entityDashboard.DashboardRetentionRateData{}

	current := start
	for current.Before(end) || current.Equal(end) {
		fmt.Println(current.Month(), current.Year())

		newData := &entityDashboard.DashboardRetentionRatePerMonthData{
			Year:       current.Year(),
			Month:      int(current.Month()),
			RenewCount: 0,
			ChurnCount: 0,
		}

		response.Months = append(response.Months, newData)

		current = current.AddDate(0, 1, 0)
	}

	contractsRenewalsData, err := s.repositories.DashboardRepository.GetDashboardRetentionRateContractsRenewalsData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range contractsRenewalsData {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		renewCount, _ := record["value"].(float64)

		for _, monthData := range response.Months {
			if monthData.Year == int(year) && monthData.Month == int(month) {
				monthData.RenewCount = int(renewCount)
			}
		}
	}

	contractsChurnedData, err := s.repositories.DashboardRepository.GetDashboardRetentionRateContractsChurnedData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range contractsChurnedData {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		churnCount, _ := record["value"].(float64)

		for _, monthData := range response.Months {
			if monthData.Year == int(year) && monthData.Month == int(month) {
				monthData.ChurnCount = int(churnCount)
			}
		}
	}

	currentRenew := 0
	previousRenew := 0
	currentChurn := 0
	previousChurn := 0

	if len(response.Months) == 1 {
		currentRenew = response.Months[len(response.Months)-1].RenewCount
		currentChurn = response.Months[len(response.Months)-1].ChurnCount
	} else if len(response.Months) > 1 {
		currentRenew = response.Months[len(response.Months)-1].RenewCount
		previousRenew = response.Months[len(response.Months)-2].RenewCount
		currentChurn = response.Months[len(response.Months)-1].ChurnCount
		previousChurn = response.Months[len(response.Months)-2].ChurnCount
	}

	currentRetentionRate := float64(currentRenew) / float64(currentRenew+currentChurn) * 100
	previousRetentionRate := float64(previousRenew) / float64(previousRenew+previousChurn) * 100

	if math.IsNaN(currentRetentionRate) {
		currentRetentionRate = 0
	}
	if math.IsNaN(previousRetentionRate) {
		previousRetentionRate = 0
	}

	if currentRenew == 0 && currentChurn == 0 {
		if previousRenew == 0 && previousChurn == 0 {
			response.RetentionRate = 0
		} else {
			response.RetentionRate = -100
		}
	} else {
		response.RetentionRate = currentRetentionRate
	}

	response.IncreasePercentage = ComputePercentagesDisplay(previousRetentionRate, response.RetentionRate)

	return &response, nil
}

func (s *dashboardService) GetDashboardNewCustomersData(ctx context.Context, start, end time.Time) (*entityDashboard.DashboardNewCustomersData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardNewCustomersData")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start))
	span.LogFields(log.Object("end", end))

	response := entityDashboard.DashboardNewCustomersData{}

	data, err := s.repositories.DashboardRepository.GetDashboardNewCustomersData(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)
		count, _ := record["count"].(int64)

		newData := &entityDashboard.DashboardNewCustomerMonthData{
			Year:  int(year),
			Month: int(month),
			Count: int(count),
		}

		response.Months = append(response.Months, newData)
	}

	currentMonthCount := 0
	previousMonthCount := 0

	if len(response.Months) == 1 {
		currentMonthCount = response.Months[len(response.Months)-1].Count
	} else if len(response.Months) > 1 {
		currentMonthCount = response.Months[len(response.Months)-1].Count
		previousMonthCount = response.Months[len(response.Months)-2].Count
	}

	response.ThisMonthCount = currentMonthCount
	response.ThisMonthIncreasePercentage = ComputeNumbersDisplay(float64(previousMonthCount), float64(currentMonthCount))

	return &response, nil

}

func (s *dashboardService) GetDashboardAverageTimeToOnboardPerMonth(ctx context.Context, start, end time.Time) (*model.DashboardTimeToOnboard, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DashboardService.GetDashboardAverageTimeToOnboardPerMonth")
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("start", start), log.Object("end", end))

	data, err := s.repositories.DashboardRepository.GetDashboardAverageTimeToOnboardPerMonth(ctx, common.GetContext(ctx).Tenant, start, end)
	if err != nil {
		return nil, err
	}

	response := model.DashboardTimeToOnboard{}

	for _, record := range data {
		year, _ := record["year"].(int64)
		month, _ := record["month"].(int64)

		newData := &model.DashboardTimeToOnboardPerMonth{
			Year:  int(year),
			Month: int(month),
		}
		_, ok := record["duration"]
		if ok {
			duration := record["duration"].(neo4j.Duration)
			totalSeconds := duration.Seconds + duration.Days*86400 + duration.Months*30*86400
			days := float64(float64(totalSeconds) / 86400.0) // 86400 seconds in a day
			roundedDays := float64(int64(days*10+0.5)) / 10  // Round to one decimal place
			if roundedDays == 0.0 && totalSeconds > 0 {
				roundedDays = 0.1
			}
			newData.Value = roundedDays
		} else {
			newData.Value = 0.0
		}

		response.PerMonth = append(response.PerMonth, newData)
	}

	currentMonth := 0.0
	previousMonth := 0.0

	if len(response.PerMonth) == 1 {
		currentMonth = response.PerMonth[len(response.PerMonth)-1].Value
	} else if len(response.PerMonth) > 1 {
		currentMonth = response.PerMonth[len(response.PerMonth)-1].Value
		previousMonth = response.PerMonth[len(response.PerMonth)-2].Value
	}
	if currentMonth == 0.0 {
		response.TimeToOnboard = nil
	} else {
		response.TimeToOnboard = &currentMonth
	}
	if currentMonth == 0.0 || previousMonth == 0.0 {
		response.IncreasePercentage = nil
	} else {
		percentageChange := calculatePercentageChange(previousMonth, currentMonth)
		response.IncreasePercentage = &percentageChange
	}

	return &response, nil
}

func ComputeNumbersDisplay(previousMonthCount, currentMonthCount float64) string {
	var increase, percentage float64

	if previousMonthCount == 0 {
		increase = float64(currentMonthCount)
		percentage = increase * 100
	} else {
		increase = float64(currentMonthCount - previousMonthCount)
		percentage = math.Round((increase / float64(previousMonthCount)) * 100)
	}

	if math.Abs(percentage) > 100 {
		if previousMonthCount == 0 {
			return fmt.Sprintf("+%.0f", increase)
		}
		a := math.Abs(percentage) / 100
		return printFloat(a, false) + "×"
	}

	return printFloat(percentage, true) + "%"
}

func ComputePercentagesDisplay(previous, current float64) string {
	if math.IsNaN(current) {
		return "0"
	}
	if math.IsNaN(previous) {
		return printFloat(current, true)
	}

	diff := current - previous

	if diff > 100 {
		diff = 100
	}
	if diff < -100 {
		diff = -100
	}

	return printFloat(diff, true)
}

func printFloat(number float64, withSign bool) string {
	if number == 0 {
		return fmt.Sprintf("%.0f", number)
	} else {
		sign := ""
		if withSign && number > 0 {
			sign = "+"
		}
		if hasSingleDecimal(number) {
			return fmt.Sprintf(sign+"%.1f", number)
		} else if hasDecimals(number) {
			return fmt.Sprintf(sign+"%.2f", number)
		} else {
			return fmt.Sprintf(sign+"%.0f", number)
		}
	}
}

func hasSingleDecimal(number float64) bool {
	if number != math.Floor(number) {
		decimalPart := number - math.Floor(number)
		return decimalPart >= 0.1 && decimalPart < 1.0
	}
	return false
}

func hasDecimals(number float64) bool {
	return number != float64(int(number))
}

func calculatePercentageChange(a, b float64) float64 {
	if a == 0 {
		if b == 0 {
			return 0.0
		}
		return math.Round((b-a)/b*1000) / 10 // Keep only one decimal place
	}
	return math.Round((b-a)/a*1000) / 10 // Keep only one decimal place
}
