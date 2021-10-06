//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedDatabaseQueriesClient contains the methods for the ManagedDatabaseQueries group.
// Don't use this type directly, use NewManagedDatabaseQueriesClient() instead.
type ManagedDatabaseQueriesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedDatabaseQueriesClient creates a new instance of ManagedDatabaseQueriesClient with the specified values.
func NewManagedDatabaseQueriesClient(con *arm.Connection, subscriptionID string) *ManagedDatabaseQueriesClient {
	return &ManagedDatabaseQueriesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get query by query id.
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseQueriesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesGetOptions) (ManagedDatabaseQueriesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, queryID, options)
	if err != nil {
		return ManagedDatabaseQueriesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseQueriesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseQueriesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseQueriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/queries/{queryId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if queryID == "" {
		return nil, errors.New("parameter queryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryId}", url.PathEscape(queryID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseQueriesClient) getHandleResponse(resp *http.Response) (ManagedDatabaseQueriesGetResponse, error) {
	result := ManagedDatabaseQueriesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceQuery); err != nil {
		return ManagedDatabaseQueriesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedDatabaseQueriesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByQuery - Get query execution statistics by query id.
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseQueriesClient) ListByQuery(resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesListByQueryOptions) *ManagedDatabaseQueriesListByQueryPager {
	return &ManagedDatabaseQueriesListByQueryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByQueryCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, queryID, options)
		},
		advancer: func(ctx context.Context, resp ManagedDatabaseQueriesListByQueryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceQueryStatistics.NextLink)
		},
	}
}

// listByQueryCreateRequest creates the ListByQuery request.
func (client *ManagedDatabaseQueriesClient) listByQueryCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesListByQueryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/queries/{queryId}/statistics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if queryID == "" {
		return nil, errors.New("parameter queryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryId}", url.PathEscape(queryID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", string(*options.Interval))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByQueryHandleResponse handles the ListByQuery response.
func (client *ManagedDatabaseQueriesClient) listByQueryHandleResponse(resp *http.Response) (ManagedDatabaseQueriesListByQueryResponse, error) {
	result := ManagedDatabaseQueriesListByQueryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceQueryStatistics); err != nil {
		return ManagedDatabaseQueriesListByQueryResponse{}, err
	}
	return result, nil
}

// listByQueryHandleError handles the ListByQuery error response.
func (client *ManagedDatabaseQueriesClient) listByQueryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
