// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DedicatedHostsClient contains the methods for the DedicatedHosts group.
// Don't use this type directly, use NewDedicatedHostsClient() instead.
type DedicatedHostsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDedicatedHostsClient creates a new instance of DedicatedHostsClient with the specified values.
func NewDedicatedHostsClient(con *armcore.Connection, subscriptionID string) *DedicatedHostsClient {
	return &DedicatedHostsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a dedicated host .
func (client *DedicatedHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsBeginCreateOrUpdateOptions) (DedicatedHostPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return DedicatedHostPollerResponse{}, err
	}
	result := DedicatedHostPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DedicatedHostsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return DedicatedHostPollerResponse{}, err
	}
	poller := &dedicatedHostPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DedicatedHostResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new DedicatedHostPoller from the specified resume token.
// token - The value must come from a previous call to DedicatedHostPoller.ResumeToken().
func (client *DedicatedHostsClient) ResumeCreateOrUpdate(token string) (DedicatedHostPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DedicatedHostsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &dedicatedHostPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a dedicated host .
func (client *DedicatedHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DedicatedHostsClient) createOrUpdateHandleResponse(resp *azcore.Response) (DedicatedHostResponse, error) {
	var val *DedicatedHost
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DedicatedHostResponse{}, err
	}
	return DedicatedHostResponse{RawResponse: resp.Response, DedicatedHost: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DedicatedHostsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Delete a dedicated host.
func (client *DedicatedHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DedicatedHostsClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *DedicatedHostsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DedicatedHostsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Delete a dedicated host.
func (client *DedicatedHostsClient) delete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DedicatedHostsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Retrieves information about a dedicated host.
func (client *DedicatedHostsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsGetOptions) (DedicatedHostResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return DedicatedHostResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DedicatedHostResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DedicatedHostResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", "instanceView")
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostsClient) getHandleResponse(resp *azcore.Response) (DedicatedHostResponse, error) {
	var val *DedicatedHost
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DedicatedHostResponse{}, err
	}
	return DedicatedHostResponse{RawResponse: resp.Response, DedicatedHost: val}, nil
}

// getHandleError handles the Get error response.
func (client *DedicatedHostsClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByHostGroup - Lists all of the dedicated hosts in the specified dedicated host group. Use the nextLink property in the response to get the next page
// of dedicated hosts.
func (client *DedicatedHostsClient) ListByHostGroup(resourceGroupName string, hostGroupName string, options *DedicatedHostsListByHostGroupOptions) DedicatedHostListResultPager {
	return &dedicatedHostListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByHostGroupCreateRequest(ctx, resourceGroupName, hostGroupName, options)
		},
		responder: client.listByHostGroupHandleResponse,
		errorer:   client.listByHostGroupHandleError,
		advancer: func(ctx context.Context, resp DedicatedHostListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DedicatedHostListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByHostGroupCreateRequest creates the ListByHostGroup request.
func (client *DedicatedHostsClient) listByHostGroupCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostsListByHostGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByHostGroupHandleResponse handles the ListByHostGroup response.
func (client *DedicatedHostsClient) listByHostGroupHandleResponse(resp *azcore.Response) (DedicatedHostListResultResponse, error) {
	var val *DedicatedHostListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DedicatedHostListResultResponse{}, err
	}
	return DedicatedHostListResultResponse{RawResponse: resp.Response, DedicatedHostListResult: val}, nil
}

// listByHostGroupHandleError handles the ListByHostGroup error response.
func (client *DedicatedHostsClient) listByHostGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Update an dedicated host .
func (client *DedicatedHostsClient) BeginUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsBeginUpdateOptions) (DedicatedHostPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return DedicatedHostPollerResponse{}, err
	}
	result := DedicatedHostPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DedicatedHostsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return DedicatedHostPollerResponse{}, err
	}
	poller := &dedicatedHostPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DedicatedHostResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new DedicatedHostPoller from the specified resume token.
// token - The value must come from a previous call to DedicatedHostPoller.ResumeToken().
func (client *DedicatedHostsClient) ResumeUpdate(token string) (DedicatedHostPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DedicatedHostsClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &dedicatedHostPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update an dedicated host .
func (client *DedicatedHostsClient) update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *DedicatedHostsClient) updateHandleResponse(resp *azcore.Response) (DedicatedHostResponse, error) {
	var val *DedicatedHost
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DedicatedHostResponse{}, err
	}
	return DedicatedHostResponse{RawResponse: resp.Response, DedicatedHost: val}, nil
}

// updateHandleError handles the Update error response.
func (client *DedicatedHostsClient) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
