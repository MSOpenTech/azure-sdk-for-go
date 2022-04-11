//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AlertsClient contains the methods for the Alerts group.
// Don't use this type directly, use NewAlertsClient() instead.
type AlertsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertsClient creates a new instance of AlertsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AlertsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ChangeState - Change the state of an alert.
// If the operation fails it returns an *azcore.ResponseError type.
// alertID - Unique ID of an alert instance.
// newState - New state of the alert.
// options - AlertsClientChangeStateOptions contains the optional parameters for the AlertsClient.ChangeState method.
func (client *AlertsClient) ChangeState(ctx context.Context, alertID string, newState AlertState, options *AlertsClientChangeStateOptions) (AlertsClientChangeStateResponse, error) {
	req, err := client.changeStateCreateRequest(ctx, alertID, newState, options)
	if err != nil {
		return AlertsClientChangeStateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientChangeStateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientChangeStateResponse{}, runtime.NewResponseError(resp)
	}
	return client.changeStateHandleResponse(resp)
}

// changeStateCreateRequest creates the ChangeState request.
func (client *AlertsClient) changeStateCreateRequest(ctx context.Context, alertID string, newState AlertState, options *AlertsClientChangeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}/changestate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertID == "" {
		return nil, errors.New("parameter alertID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", url.PathEscape(alertID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-05-preview")
	reqQP.Set("newState", string(newState))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// changeStateHandleResponse handles the ChangeState response.
func (client *AlertsClient) changeStateHandleResponse(resp *http.Response) (AlertsClientChangeStateResponse, error) {
	result := AlertsClientChangeStateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientChangeStateResponse{}, err
	}
	return result, nil
}

// GetAll - List all existing alerts, where the results can be filtered on the basis of multiple parameters (e.g. time range).
// The results can then be sorted on the basis specific fields, with the default being
// lastModifiedDateTime.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AlertsClientGetAllOptions contains the optional parameters for the AlertsClient.GetAll method.
func (client *AlertsClient) GetAll(options *AlertsClientGetAllOptions) *runtime.Pager[AlertsClientGetAllResponse] {
	return runtime.NewPager(runtime.PageProcessor[AlertsClientGetAllResponse]{
		More: func(page AlertsClientGetAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertsClientGetAllResponse) (AlertsClientGetAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertsClientGetAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertsClientGetAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertsClientGetAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.getAllHandleResponse(resp)
		},
	})
}

// getAllCreateRequest creates the GetAll request.
func (client *AlertsClient) getAllCreateRequest(ctx context.Context, options *AlertsClientGetAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.TargetResource != nil {
		reqQP.Set("targetResource", *options.TargetResource)
	}
	if options != nil && options.TargetResourceType != nil {
		reqQP.Set("targetResourceType", *options.TargetResourceType)
	}
	if options != nil && options.TargetResourceGroup != nil {
		reqQP.Set("targetResourceGroup", *options.TargetResourceGroup)
	}
	if options != nil && options.MonitorService != nil {
		reqQP.Set("monitorService", string(*options.MonitorService))
	}
	if options != nil && options.MonitorCondition != nil {
		reqQP.Set("monitorCondition", string(*options.MonitorCondition))
	}
	if options != nil && options.Severity != nil {
		reqQP.Set("severity", string(*options.Severity))
	}
	if options != nil && options.AlertState != nil {
		reqQP.Set("alertState", string(*options.AlertState))
	}
	if options != nil && options.AlertRule != nil {
		reqQP.Set("alertRule", *options.AlertRule)
	}
	if options != nil && options.SmartGroupID != nil {
		reqQP.Set("smartGroupId", *options.SmartGroupID)
	}
	if options != nil && options.IncludeContext != nil {
		reqQP.Set("includeContext", strconv.FormatBool(*options.IncludeContext))
	}
	if options != nil && options.IncludeEgressConfig != nil {
		reqQP.Set("includeEgressConfig", strconv.FormatBool(*options.IncludeEgressConfig))
	}
	if options != nil && options.PageCount != nil {
		reqQP.Set("pageCount", strconv.FormatInt(*options.PageCount, 10))
	}
	if options != nil && options.SortBy != nil {
		reqQP.Set("sortBy", string(*options.SortBy))
	}
	if options != nil && options.SortOrder != nil {
		reqQP.Set("sortOrder", string(*options.SortOrder))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("select", *options.Select)
	}
	if options != nil && options.TimeRange != nil {
		reqQP.Set("timeRange", string(*options.TimeRange))
	}
	if options != nil && options.CustomTimeRange != nil {
		reqQP.Set("customTimeRange", *options.CustomTimeRange)
	}
	reqQP.Set("api-version", "2019-05-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAllHandleResponse handles the GetAll response.
func (client *AlertsClient) getAllHandleResponse(resp *http.Response) (AlertsClientGetAllResponse, error) {
	result := AlertsClientGetAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsList); err != nil {
		return AlertsClientGetAllResponse{}, err
	}
	return result, nil
}

// GetByID - Get information related to a specific alert
// If the operation fails it returns an *azcore.ResponseError type.
// alertID - Unique ID of an alert instance.
// options - AlertsClientGetByIDOptions contains the optional parameters for the AlertsClient.GetByID method.
func (client *AlertsClient) GetByID(ctx context.Context, alertID string, options *AlertsClientGetByIDOptions) (AlertsClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, alertID, options)
	if err != nil {
		return AlertsClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *AlertsClient) getByIDCreateRequest(ctx context.Context, alertID string, options *AlertsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertID == "" {
		return nil, errors.New("parameter alertID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", url.PathEscape(alertID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AlertsClient) getByIDHandleResponse(resp *http.Response) (AlertsClientGetByIDResponse, error) {
	result := AlertsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientGetByIDResponse{}, err
	}
	return result, nil
}

// GetHistory - Get the history of an alert, which captures any monitor condition changes (Fired/Resolved) and alert state
// changes (New/Acknowledged/Closed).
// If the operation fails it returns an *azcore.ResponseError type.
// alertID - Unique ID of an alert instance.
// options - AlertsClientGetHistoryOptions contains the optional parameters for the AlertsClient.GetHistory method.
func (client *AlertsClient) GetHistory(ctx context.Context, alertID string, options *AlertsClientGetHistoryOptions) (AlertsClientGetHistoryResponse, error) {
	req, err := client.getHistoryCreateRequest(ctx, alertID, options)
	if err != nil {
		return AlertsClientGetHistoryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientGetHistoryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientGetHistoryResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHistoryHandleResponse(resp)
}

// getHistoryCreateRequest creates the GetHistory request.
func (client *AlertsClient) getHistoryCreateRequest(ctx context.Context, alertID string, options *AlertsClientGetHistoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}/history"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertID == "" {
		return nil, errors.New("parameter alertID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", url.PathEscape(alertID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHistoryHandleResponse handles the GetHistory response.
func (client *AlertsClient) getHistoryHandleResponse(resp *http.Response) (AlertsClientGetHistoryResponse, error) {
	result := AlertsClientGetHistoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertModification); err != nil {
		return AlertsClientGetHistoryResponse{}, err
	}
	return result, nil
}

// GetSummary - Get a summarized count of your alerts grouped by various parameters (e.g. grouping by 'Severity' returns the
// count of alerts for each severity).
// If the operation fails it returns an *azcore.ResponseError type.
// groupby - This parameter allows the result set to be grouped by input fields (Maximum 2 comma separated fields supported).
// For example, groupby=severity or groupby=severity,alertstate.
// options - AlertsClientGetSummaryOptions contains the optional parameters for the AlertsClient.GetSummary method.
func (client *AlertsClient) GetSummary(ctx context.Context, groupby AlertsSummaryGroupByFields, options *AlertsClientGetSummaryOptions) (AlertsClientGetSummaryResponse, error) {
	req, err := client.getSummaryCreateRequest(ctx, groupby, options)
	if err != nil {
		return AlertsClientGetSummaryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientGetSummaryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientGetSummaryResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSummaryHandleResponse(resp)
}

// getSummaryCreateRequest creates the GetSummary request.
func (client *AlertsClient) getSummaryCreateRequest(ctx context.Context, groupby AlertsSummaryGroupByFields, options *AlertsClientGetSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alertsSummary"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("groupby", string(groupby))
	if options != nil && options.IncludeSmartGroupsCount != nil {
		reqQP.Set("includeSmartGroupsCount", strconv.FormatBool(*options.IncludeSmartGroupsCount))
	}
	if options != nil && options.TargetResource != nil {
		reqQP.Set("targetResource", *options.TargetResource)
	}
	if options != nil && options.TargetResourceType != nil {
		reqQP.Set("targetResourceType", *options.TargetResourceType)
	}
	if options != nil && options.TargetResourceGroup != nil {
		reqQP.Set("targetResourceGroup", *options.TargetResourceGroup)
	}
	if options != nil && options.MonitorService != nil {
		reqQP.Set("monitorService", string(*options.MonitorService))
	}
	if options != nil && options.MonitorCondition != nil {
		reqQP.Set("monitorCondition", string(*options.MonitorCondition))
	}
	if options != nil && options.Severity != nil {
		reqQP.Set("severity", string(*options.Severity))
	}
	if options != nil && options.AlertState != nil {
		reqQP.Set("alertState", string(*options.AlertState))
	}
	if options != nil && options.AlertRule != nil {
		reqQP.Set("alertRule", *options.AlertRule)
	}
	if options != nil && options.TimeRange != nil {
		reqQP.Set("timeRange", string(*options.TimeRange))
	}
	if options != nil && options.CustomTimeRange != nil {
		reqQP.Set("customTimeRange", *options.CustomTimeRange)
	}
	reqQP.Set("api-version", "2019-05-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSummaryHandleResponse handles the GetSummary response.
func (client *AlertsClient) getSummaryHandleResponse(resp *http.Response) (AlertsClientGetSummaryResponse, error) {
	result := AlertsClientGetSummaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsSummary); err != nil {
		return AlertsClientGetSummaryResponse{}, err
	}
	return result, nil
}

// MetaData - List alerts meta data information based on value of identifier parameter.
// If the operation fails it returns an *azcore.ResponseError type.
// identifier - Identification of the information to be retrieved by API call.
// options - AlertsClientMetaDataOptions contains the optional parameters for the AlertsClient.MetaData method.
func (client *AlertsClient) MetaData(ctx context.Context, identifier Identifier, options *AlertsClientMetaDataOptions) (AlertsClientMetaDataResponse, error) {
	req, err := client.metaDataCreateRequest(ctx, identifier, options)
	if err != nil {
		return AlertsClientMetaDataResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertsClientMetaDataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertsClientMetaDataResponse{}, runtime.NewResponseError(resp)
	}
	return client.metaDataHandleResponse(resp)
}

// metaDataCreateRequest creates the MetaData request.
func (client *AlertsClient) metaDataCreateRequest(ctx context.Context, identifier Identifier, options *AlertsClientMetaDataOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AlertsManagement/alertsMetaData"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-05-preview")
	reqQP.Set("identifier", string(identifier))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// metaDataHandleResponse handles the MetaData response.
func (client *AlertsClient) metaDataHandleResponse(resp *http.Response) (AlertsClientMetaDataResponse, error) {
	result := AlertsClientMetaDataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsMetaData); err != nil {
		return AlertsClientMetaDataResponse{}, err
	}
	return result, nil
}
