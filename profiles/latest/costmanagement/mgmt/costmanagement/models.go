// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package costmanagement

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2020-06-01/costmanagement"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccumulatedType = original.AccumulatedType

const (
	False AccumulatedType = original.False
	True  AccumulatedType = original.True
)

type AlertCategory = original.AlertCategory

const (
	Billing AlertCategory = original.Billing
	Cost    AlertCategory = original.Cost
	System  AlertCategory = original.System
	Usage   AlertCategory = original.Usage
)

type AlertCriteria = original.AlertCriteria

const (
	CostThresholdExceeded          AlertCriteria = original.CostThresholdExceeded
	CreditThresholdApproaching     AlertCriteria = original.CreditThresholdApproaching
	CreditThresholdReached         AlertCriteria = original.CreditThresholdReached
	CrossCloudCollectionError      AlertCriteria = original.CrossCloudCollectionError
	CrossCloudNewDataAvailable     AlertCriteria = original.CrossCloudNewDataAvailable
	ForecastCostThresholdExceeded  AlertCriteria = original.ForecastCostThresholdExceeded
	ForecastUsageThresholdExceeded AlertCriteria = original.ForecastUsageThresholdExceeded
	GeneralThresholdError          AlertCriteria = original.GeneralThresholdError
	InvoiceDueDateApproaching      AlertCriteria = original.InvoiceDueDateApproaching
	InvoiceDueDateReached          AlertCriteria = original.InvoiceDueDateReached
	MultiCurrency                  AlertCriteria = original.MultiCurrency
	QuotaThresholdApproaching      AlertCriteria = original.QuotaThresholdApproaching
	QuotaThresholdReached          AlertCriteria = original.QuotaThresholdReached
	UsageThresholdExceeded         AlertCriteria = original.UsageThresholdExceeded
)

type AlertOperator = original.AlertOperator

const (
	EqualTo              AlertOperator = original.EqualTo
	GreaterThan          AlertOperator = original.GreaterThan
	GreaterThanOrEqualTo AlertOperator = original.GreaterThanOrEqualTo
	LessThan             AlertOperator = original.LessThan
	LessThanOrEqualTo    AlertOperator = original.LessThanOrEqualTo
	None                 AlertOperator = original.None
)

type AlertSource = original.AlertSource

const (
	Preset AlertSource = original.Preset
	User   AlertSource = original.User
)

type AlertStatus = original.AlertStatus

const (
	AlertStatusActive     AlertStatus = original.AlertStatusActive
	AlertStatusDismissed  AlertStatus = original.AlertStatusDismissed
	AlertStatusNone       AlertStatus = original.AlertStatusNone
	AlertStatusOverridden AlertStatus = original.AlertStatusOverridden
	AlertStatusResolved   AlertStatus = original.AlertStatusResolved
)

type AlertTimeGrainType = original.AlertTimeGrainType

const (
	AlertTimeGrainTypeAnnually       AlertTimeGrainType = original.AlertTimeGrainTypeAnnually
	AlertTimeGrainTypeBillingAnnual  AlertTimeGrainType = original.AlertTimeGrainTypeBillingAnnual
	AlertTimeGrainTypeBillingMonth   AlertTimeGrainType = original.AlertTimeGrainTypeBillingMonth
	AlertTimeGrainTypeBillingQuarter AlertTimeGrainType = original.AlertTimeGrainTypeBillingQuarter
	AlertTimeGrainTypeMonthly        AlertTimeGrainType = original.AlertTimeGrainTypeMonthly
	AlertTimeGrainTypeNone           AlertTimeGrainType = original.AlertTimeGrainTypeNone
	AlertTimeGrainTypeQuarterly      AlertTimeGrainType = original.AlertTimeGrainTypeQuarterly
)

type AlertType = original.AlertType

const (
	Budget         AlertType = original.Budget
	BudgetForecast AlertType = original.BudgetForecast
	Credit         AlertType = original.Credit
	General        AlertType = original.General
	Invoice        AlertType = original.Invoice
	Quota          AlertType = original.Quota
	XCloud         AlertType = original.XCloud
)

type ChartType = original.ChartType

const (
	Area          ChartType = original.Area
	GroupedColumn ChartType = original.GroupedColumn
	Line          ChartType = original.Line
	StackedColumn ChartType = original.StackedColumn
	Table         ChartType = original.Table
)

type Direction = original.Direction

const (
	Ascending  Direction = original.Ascending
	Descending Direction = original.Descending
)

type ExecutionStatus = original.ExecutionStatus

const (
	Completed           ExecutionStatus = original.Completed
	DataNotAvailable    ExecutionStatus = original.DataNotAvailable
	Failed              ExecutionStatus = original.Failed
	InProgress          ExecutionStatus = original.InProgress
	NewDataNotAvailable ExecutionStatus = original.NewDataNotAvailable
	Queued              ExecutionStatus = original.Queued
	Timeout             ExecutionStatus = original.Timeout
)

type ExecutionType = original.ExecutionType

const (
	OnDemand  ExecutionType = original.OnDemand
	Scheduled ExecutionType = original.Scheduled
)

type ExportType = original.ExportType

const (
	ExportTypeActualCost    ExportType = original.ExportTypeActualCost
	ExportTypeAmortizedCost ExportType = original.ExportTypeAmortizedCost
	ExportTypeUsage         ExportType = original.ExportTypeUsage
)

type ExternalCloudProviderType = original.ExternalCloudProviderType

const (
	ExternalBillingAccounts ExternalCloudProviderType = original.ExternalBillingAccounts
	ExternalSubscriptions   ExternalCloudProviderType = original.ExternalSubscriptions
)

type ForecastTimeframeType = original.ForecastTimeframeType

const (
	BillingMonthToDate  ForecastTimeframeType = original.BillingMonthToDate
	Custom              ForecastTimeframeType = original.Custom
	MonthToDate         ForecastTimeframeType = original.MonthToDate
	TheLastBillingMonth ForecastTimeframeType = original.TheLastBillingMonth
	TheLastMonth        ForecastTimeframeType = original.TheLastMonth
	WeekToDate          ForecastTimeframeType = original.WeekToDate
)

type ForecastType = original.ForecastType

const (
	ForecastTypeActualCost    ForecastType = original.ForecastTypeActualCost
	ForecastTypeAmortizedCost ForecastType = original.ForecastTypeAmortizedCost
	ForecastTypeUsage         ForecastType = original.ForecastTypeUsage
)

type FormatType = original.FormatType

const (
	Csv FormatType = original.Csv
)

type GranularityType = original.GranularityType

const (
	Daily GranularityType = original.Daily
)

type KpiTypeType = original.KpiTypeType

const (
	KpiTypeTypeBudget   KpiTypeType = original.KpiTypeTypeBudget
	KpiTypeTypeForecast KpiTypeType = original.KpiTypeTypeForecast
)

type MetricType = original.MetricType

const (
	ActualCost    MetricType = original.ActualCost
	AHUB          MetricType = original.AHUB
	AmortizedCost MetricType = original.AmortizedCost
)

type OperatorType = original.OperatorType

const (
	Contains OperatorType = original.Contains
	In       OperatorType = original.In
)

type PivotTypeType = original.PivotTypeType

const (
	PivotTypeTypeDimension PivotTypeType = original.PivotTypeTypeDimension
	PivotTypeTypeTagKey    PivotTypeType = original.PivotTypeTypeTagKey
)

type QueryColumnType = original.QueryColumnType

const (
	QueryColumnTypeDimension QueryColumnType = original.QueryColumnTypeDimension
	QueryColumnTypeTag       QueryColumnType = original.QueryColumnTypeTag
)

type RecurrenceType = original.RecurrenceType

const (
	RecurrenceTypeAnnually RecurrenceType = original.RecurrenceTypeAnnually
	RecurrenceTypeDaily    RecurrenceType = original.RecurrenceTypeDaily
	RecurrenceTypeMonthly  RecurrenceType = original.RecurrenceTypeMonthly
	RecurrenceTypeWeekly   RecurrenceType = original.RecurrenceTypeWeekly
)

type ReportConfigColumnType = original.ReportConfigColumnType

const (
	ReportConfigColumnTypeDimension ReportConfigColumnType = original.ReportConfigColumnTypeDimension
	ReportConfigColumnTypeTag       ReportConfigColumnType = original.ReportConfigColumnTypeTag
)

type ReportGranularityType = original.ReportGranularityType

const (
	ReportGranularityTypeDaily   ReportGranularityType = original.ReportGranularityTypeDaily
	ReportGranularityTypeMonthly ReportGranularityType = original.ReportGranularityTypeMonthly
)

type ReportTimeframeType = original.ReportTimeframeType

const (
	ReportTimeframeTypeCustom      ReportTimeframeType = original.ReportTimeframeTypeCustom
	ReportTimeframeTypeMonthToDate ReportTimeframeType = original.ReportTimeframeTypeMonthToDate
	ReportTimeframeTypeWeekToDate  ReportTimeframeType = original.ReportTimeframeTypeWeekToDate
	ReportTimeframeTypeYearToDate  ReportTimeframeType = original.ReportTimeframeTypeYearToDate
)

type StatusType = original.StatusType

const (
	Active   StatusType = original.Active
	Inactive StatusType = original.Inactive
)

type TimeframeType = original.TimeframeType

const (
	TimeframeTypeBillingMonthToDate  TimeframeType = original.TimeframeTypeBillingMonthToDate
	TimeframeTypeCustom              TimeframeType = original.TimeframeTypeCustom
	TimeframeTypeMonthToDate         TimeframeType = original.TimeframeTypeMonthToDate
	TimeframeTypeTheLastBillingMonth TimeframeType = original.TimeframeTypeTheLastBillingMonth
	TimeframeTypeTheLastMonth        TimeframeType = original.TimeframeTypeTheLastMonth
	TimeframeTypeWeekToDate          TimeframeType = original.TimeframeTypeWeekToDate
)

type Alert = original.Alert
type AlertProperties = original.AlertProperties
type AlertPropertiesDefinition = original.AlertPropertiesDefinition
type AlertPropertiesDetails = original.AlertPropertiesDetails
type AlertsClient = original.AlertsClient
type AlertsResult = original.AlertsResult
type BaseClient = original.BaseClient
type CommonExportProperties = original.CommonExportProperties
type Dimension = original.Dimension
type DimensionProperties = original.DimensionProperties
type DimensionsClient = original.DimensionsClient
type DimensionsListResult = original.DimensionsListResult
type DismissAlertPayload = original.DismissAlertPayload
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Export = original.Export
type ExportDataset = original.ExportDataset
type ExportDatasetConfiguration = original.ExportDatasetConfiguration
type ExportDefinition = original.ExportDefinition
type ExportDeliveryDestination = original.ExportDeliveryDestination
type ExportDeliveryInfo = original.ExportDeliveryInfo
type ExportExecution = original.ExportExecution
type ExportExecutionListResult = original.ExportExecutionListResult
type ExportExecutionProperties = original.ExportExecutionProperties
type ExportListResult = original.ExportListResult
type ExportProperties = original.ExportProperties
type ExportRecurrencePeriod = original.ExportRecurrencePeriod
type ExportSchedule = original.ExportSchedule
type ExportTimePeriod = original.ExportTimePeriod
type ExportsClient = original.ExportsClient
type ForecastClient = original.ForecastClient
type ForecastDataset = original.ForecastDataset
type ForecastDefinition = original.ForecastDefinition
type KpiProperties = original.KpiProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PivotProperties = original.PivotProperties
type ProxyResource = original.ProxyResource
type QueryAggregation = original.QueryAggregation
type QueryClient = original.QueryClient
type QueryColumn = original.QueryColumn
type QueryComparisonExpression = original.QueryComparisonExpression
type QueryDataset = original.QueryDataset
type QueryDatasetConfiguration = original.QueryDatasetConfiguration
type QueryDefinition = original.QueryDefinition
type QueryFilter = original.QueryFilter
type QueryGrouping = original.QueryGrouping
type QueryProperties = original.QueryProperties
type QueryResult = original.QueryResult
type QueryTimePeriod = original.QueryTimePeriod
type ReportConfigAggregation = original.ReportConfigAggregation
type ReportConfigComparisonExpression = original.ReportConfigComparisonExpression
type ReportConfigDataset = original.ReportConfigDataset
type ReportConfigDatasetConfiguration = original.ReportConfigDatasetConfiguration
type ReportConfigDefinition = original.ReportConfigDefinition
type ReportConfigFilter = original.ReportConfigFilter
type ReportConfigGrouping = original.ReportConfigGrouping
type ReportConfigSorting = original.ReportConfigSorting
type ReportConfigTimePeriod = original.ReportConfigTimePeriod
type Resource = original.Resource
type View = original.View
type ViewListResult = original.ViewListResult
type ViewListResultIterator = original.ViewListResultIterator
type ViewListResultPage = original.ViewListResultPage
type ViewProperties = original.ViewProperties
type ViewsClient = original.ViewsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDimensionsClient(subscriptionID string) DimensionsClient {
	return original.NewDimensionsClient(subscriptionID)
}
func NewDimensionsClientWithBaseURI(baseURI string, subscriptionID string) DimensionsClient {
	return original.NewDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExportsClient(subscriptionID string) ExportsClient {
	return original.NewExportsClient(subscriptionID)
}
func NewExportsClientWithBaseURI(baseURI string, subscriptionID string) ExportsClient {
	return original.NewExportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewForecastClient(subscriptionID string) ForecastClient {
	return original.NewForecastClient(subscriptionID)
}
func NewForecastClientWithBaseURI(baseURI string, subscriptionID string) ForecastClient {
	return original.NewForecastClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueryClient(subscriptionID string) QueryClient {
	return original.NewQueryClient(subscriptionID)
}
func NewQueryClientWithBaseURI(baseURI string, subscriptionID string) QueryClient {
	return original.NewQueryClientWithBaseURI(baseURI, subscriptionID)
}
func NewViewListResultIterator(page ViewListResultPage) ViewListResultIterator {
	return original.NewViewListResultIterator(page)
}
func NewViewListResultPage(cur ViewListResult, getNextPage func(context.Context, ViewListResult) (ViewListResult, error)) ViewListResultPage {
	return original.NewViewListResultPage(cur, getNextPage)
}
func NewViewsClient(subscriptionID string) ViewsClient {
	return original.NewViewsClient(subscriptionID)
}
func NewViewsClientWithBaseURI(baseURI string, subscriptionID string) ViewsClient {
	return original.NewViewsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccumulatedTypeValues() []AccumulatedType {
	return original.PossibleAccumulatedTypeValues()
}
func PossibleAlertCategoryValues() []AlertCategory {
	return original.PossibleAlertCategoryValues()
}
func PossibleAlertCriteriaValues() []AlertCriteria {
	return original.PossibleAlertCriteriaValues()
}
func PossibleAlertOperatorValues() []AlertOperator {
	return original.PossibleAlertOperatorValues()
}
func PossibleAlertSourceValues() []AlertSource {
	return original.PossibleAlertSourceValues()
}
func PossibleAlertStatusValues() []AlertStatus {
	return original.PossibleAlertStatusValues()
}
func PossibleAlertTimeGrainTypeValues() []AlertTimeGrainType {
	return original.PossibleAlertTimeGrainTypeValues()
}
func PossibleAlertTypeValues() []AlertType {
	return original.PossibleAlertTypeValues()
}
func PossibleChartTypeValues() []ChartType {
	return original.PossibleChartTypeValues()
}
func PossibleDirectionValues() []Direction {
	return original.PossibleDirectionValues()
}
func PossibleExecutionStatusValues() []ExecutionStatus {
	return original.PossibleExecutionStatusValues()
}
func PossibleExecutionTypeValues() []ExecutionType {
	return original.PossibleExecutionTypeValues()
}
func PossibleExportTypeValues() []ExportType {
	return original.PossibleExportTypeValues()
}
func PossibleExternalCloudProviderTypeValues() []ExternalCloudProviderType {
	return original.PossibleExternalCloudProviderTypeValues()
}
func PossibleForecastTimeframeTypeValues() []ForecastTimeframeType {
	return original.PossibleForecastTimeframeTypeValues()
}
func PossibleForecastTypeValues() []ForecastType {
	return original.PossibleForecastTypeValues()
}
func PossibleFormatTypeValues() []FormatType {
	return original.PossibleFormatTypeValues()
}
func PossibleGranularityTypeValues() []GranularityType {
	return original.PossibleGranularityTypeValues()
}
func PossibleKpiTypeTypeValues() []KpiTypeType {
	return original.PossibleKpiTypeTypeValues()
}
func PossibleMetricTypeValues() []MetricType {
	return original.PossibleMetricTypeValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossiblePivotTypeTypeValues() []PivotTypeType {
	return original.PossiblePivotTypeTypeValues()
}
func PossibleQueryColumnTypeValues() []QueryColumnType {
	return original.PossibleQueryColumnTypeValues()
}
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return original.PossibleRecurrenceTypeValues()
}
func PossibleReportConfigColumnTypeValues() []ReportConfigColumnType {
	return original.PossibleReportConfigColumnTypeValues()
}
func PossibleReportGranularityTypeValues() []ReportGranularityType {
	return original.PossibleReportGranularityTypeValues()
}
func PossibleReportTimeframeTypeValues() []ReportTimeframeType {
	return original.PossibleReportTimeframeTypeValues()
}
func PossibleStatusTypeValues() []StatusType {
	return original.PossibleStatusTypeValues()
}
func PossibleTimeframeTypeValues() []TimeframeType {
	return original.PossibleTimeframeTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
