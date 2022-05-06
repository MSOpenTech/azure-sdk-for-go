//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

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
	"strings"
)

// TimeSeriesDatabaseConnectionsClient contains the methods for the TimeSeriesDatabaseConnections group.
// Don't use this type directly, use NewTimeSeriesDatabaseConnectionsClient() instead.
type TimeSeriesDatabaseConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTimeSeriesDatabaseConnectionsClient creates a new instance of TimeSeriesDatabaseConnectionsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTimeSeriesDatabaseConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TimeSeriesDatabaseConnectionsClient, error) {
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
	client := &TimeSeriesDatabaseConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a time series database connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// timeSeriesDatabaseConnectionName - Name of time series database connection.
// timeSeriesDatabaseConnectionDescription - The time series database connection description.
// options - TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.BeginCreateOrUpdate
// method.
func (client *TimeSeriesDatabaseConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, timeSeriesDatabaseConnectionDescription TimeSeriesDatabaseConnection, options *TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[TimeSeriesDatabaseConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, timeSeriesDatabaseConnectionName, timeSeriesDatabaseConnectionDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[TimeSeriesDatabaseConnectionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[TimeSeriesDatabaseConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a time series database connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TimeSeriesDatabaseConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, timeSeriesDatabaseConnectionDescription TimeSeriesDatabaseConnection, options *TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, timeSeriesDatabaseConnectionName, timeSeriesDatabaseConnectionDescription, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TimeSeriesDatabaseConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, timeSeriesDatabaseConnectionDescription TimeSeriesDatabaseConnection, options *TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/timeSeriesDatabaseConnections/{timeSeriesDatabaseConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if timeSeriesDatabaseConnectionName == "" {
		return nil, errors.New("parameter timeSeriesDatabaseConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timeSeriesDatabaseConnectionName}", url.PathEscape(timeSeriesDatabaseConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, timeSeriesDatabaseConnectionDescription)
}

// BeginDelete - Delete a time series database connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// timeSeriesDatabaseConnectionName - Name of time series database connection.
// options - TimeSeriesDatabaseConnectionsClientBeginDeleteOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.BeginDelete
// method.
func (client *TimeSeriesDatabaseConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *TimeSeriesDatabaseConnectionsClientBeginDeleteOptions) (*armruntime.Poller[TimeSeriesDatabaseConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, timeSeriesDatabaseConnectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[TimeSeriesDatabaseConnectionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[TimeSeriesDatabaseConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a time series database connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TimeSeriesDatabaseConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *TimeSeriesDatabaseConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, timeSeriesDatabaseConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TimeSeriesDatabaseConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *TimeSeriesDatabaseConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/timeSeriesDatabaseConnections/{timeSeriesDatabaseConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if timeSeriesDatabaseConnectionName == "" {
		return nil, errors.New("parameter timeSeriesDatabaseConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timeSeriesDatabaseConnectionName}", url.PathEscape(timeSeriesDatabaseConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the description of an existing time series database connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// timeSeriesDatabaseConnectionName - Name of time series database connection.
// options - TimeSeriesDatabaseConnectionsClientGetOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.Get
// method.
func (client *TimeSeriesDatabaseConnectionsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *TimeSeriesDatabaseConnectionsClientGetOptions) (TimeSeriesDatabaseConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, timeSeriesDatabaseConnectionName, options)
	if err != nil {
		return TimeSeriesDatabaseConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TimeSeriesDatabaseConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TimeSeriesDatabaseConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TimeSeriesDatabaseConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *TimeSeriesDatabaseConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/timeSeriesDatabaseConnections/{timeSeriesDatabaseConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if timeSeriesDatabaseConnectionName == "" {
		return nil, errors.New("parameter timeSeriesDatabaseConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timeSeriesDatabaseConnectionName}", url.PathEscape(timeSeriesDatabaseConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TimeSeriesDatabaseConnectionsClient) getHandleResponse(resp *http.Response) (TimeSeriesDatabaseConnectionsClientGetResponse, error) {
	result := TimeSeriesDatabaseConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeSeriesDatabaseConnection); err != nil {
		return TimeSeriesDatabaseConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all existing time series database connections for this DigitalTwins instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
// resourceName - The name of the DigitalTwinsInstance.
// options - TimeSeriesDatabaseConnectionsClientListOptions contains the optional parameters for the TimeSeriesDatabaseConnectionsClient.List
// method.
func (client *TimeSeriesDatabaseConnectionsClient) NewListPager(resourceGroupName string, resourceName string, options *TimeSeriesDatabaseConnectionsClientListOptions) *runtime.Pager[TimeSeriesDatabaseConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TimeSeriesDatabaseConnectionsClientListResponse]{
		More: func(page TimeSeriesDatabaseConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TimeSeriesDatabaseConnectionsClientListResponse) (TimeSeriesDatabaseConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TimeSeriesDatabaseConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TimeSeriesDatabaseConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TimeSeriesDatabaseConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TimeSeriesDatabaseConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *TimeSeriesDatabaseConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/timeSeriesDatabaseConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TimeSeriesDatabaseConnectionsClient) listHandleResponse(resp *http.Response) (TimeSeriesDatabaseConnectionsClientListResponse, error) {
	result := TimeSeriesDatabaseConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeSeriesDatabaseConnectionListResult); err != nil {
		return TimeSeriesDatabaseConnectionsClientListResponse{}, err
	}
	return result, nil
}
