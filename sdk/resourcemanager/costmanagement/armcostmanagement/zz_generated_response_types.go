//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

// AlertsClientDismissResponse contains the response from method AlertsClient.Dismiss.
type AlertsClientDismissResponse struct {
	Alert
}

// AlertsClientGetResponse contains the response from method AlertsClient.Get.
type AlertsClientGetResponse struct {
	Alert
}

// AlertsClientListExternalResponse contains the response from method AlertsClient.ListExternal.
type AlertsClientListExternalResponse struct {
	AlertsResult
}

// AlertsClientListResponse contains the response from method AlertsClient.List.
type AlertsClientListResponse struct {
	AlertsResult
}

// DimensionsClientByExternalCloudProviderTypeResponse contains the response from method DimensionsClient.ByExternalCloudProviderType.
type DimensionsClientByExternalCloudProviderTypeResponse struct {
	DimensionsListResult
}

// DimensionsClientListResponse contains the response from method DimensionsClient.List.
type DimensionsClientListResponse struct {
	DimensionsListResult
}

// ExportsClientCreateOrUpdateResponse contains the response from method ExportsClient.CreateOrUpdate.
type ExportsClientCreateOrUpdateResponse struct {
	Export
}

// ExportsClientDeleteResponse contains the response from method ExportsClient.Delete.
type ExportsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExportsClientExecuteResponse contains the response from method ExportsClient.Execute.
type ExportsClientExecuteResponse struct {
	// placeholder for future response values
}

// ExportsClientGetExecutionHistoryResponse contains the response from method ExportsClient.GetExecutionHistory.
type ExportsClientGetExecutionHistoryResponse struct {
	ExportExecutionListResult
}

// ExportsClientGetResponse contains the response from method ExportsClient.Get.
type ExportsClientGetResponse struct {
	Export
}

// ExportsClientListResponse contains the response from method ExportsClient.List.
type ExportsClientListResponse struct {
	ExportListResult
}

// ForecastClientExternalCloudProviderUsageResponse contains the response from method ForecastClient.ExternalCloudProviderUsage.
type ForecastClientExternalCloudProviderUsageResponse struct {
	QueryResult
}

// ForecastClientUsageResponse contains the response from method ForecastClient.Usage.
type ForecastClientUsageResponse struct {
	QueryResult
}

// GenerateDetailedCostReportClientCreateOperationResponse contains the response from method GenerateDetailedCostReportClient.CreateOperation.
type GenerateDetailedCostReportClientCreateOperationResponse struct {
	GenerateDetailedCostReportOperationResult
}

// GenerateDetailedCostReportOperationResultsClientGetResponse contains the response from method GenerateDetailedCostReportOperationResultsClient.Get.
type GenerateDetailedCostReportOperationResultsClientGetResponse struct {
	GenerateDetailedCostReportOperationResult
}

// GenerateDetailedCostReportOperationStatusClientGetResponse contains the response from method GenerateDetailedCostReportOperationStatusClient.Get.
type GenerateDetailedCostReportOperationStatusClientGetResponse struct {
	GenerateDetailedCostReportOperationStatuses
}

// GenerateReservationDetailsReportClientByBillingAccountIDResponse contains the response from method GenerateReservationDetailsReportClient.ByBillingAccountID.
type GenerateReservationDetailsReportClientByBillingAccountIDResponse struct {
	OperationStatus
}

// GenerateReservationDetailsReportClientByBillingProfileIDResponse contains the response from method GenerateReservationDetailsReportClient.ByBillingProfileID.
type GenerateReservationDetailsReportClientByBillingProfileIDResponse struct {
	OperationStatus
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// QueryClientUsageByExternalCloudProviderTypeResponse contains the response from method QueryClient.UsageByExternalCloudProviderType.
type QueryClientUsageByExternalCloudProviderTypeResponse struct {
	QueryResult
}

// QueryClientUsageResponse contains the response from method QueryClient.Usage.
type QueryClientUsageResponse struct {
	QueryResult
}

// ViewsClientCreateOrUpdateByScopeResponse contains the response from method ViewsClient.CreateOrUpdateByScope.
type ViewsClientCreateOrUpdateByScopeResponse struct {
	View
}

// ViewsClientCreateOrUpdateResponse contains the response from method ViewsClient.CreateOrUpdate.
type ViewsClientCreateOrUpdateResponse struct {
	View
}

// ViewsClientDeleteByScopeResponse contains the response from method ViewsClient.DeleteByScope.
type ViewsClientDeleteByScopeResponse struct {
	// placeholder for future response values
}

// ViewsClientDeleteResponse contains the response from method ViewsClient.Delete.
type ViewsClientDeleteResponse struct {
	// placeholder for future response values
}

// ViewsClientGetByScopeResponse contains the response from method ViewsClient.GetByScope.
type ViewsClientGetByScopeResponse struct {
	View
}

// ViewsClientGetResponse contains the response from method ViewsClient.Get.
type ViewsClientGetResponse struct {
	View
}

// ViewsClientListByScopeResponse contains the response from method ViewsClient.ListByScope.
type ViewsClientListByScopeResponse struct {
	ViewListResult
}

// ViewsClientListResponse contains the response from method ViewsClient.List.
type ViewsClientListResponse struct {
	ViewListResult
}
