// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ApplicationSecurityGroupsClient contains the methods for the ApplicationSecurityGroups group.
// Don't use this type directly, use NewApplicationSecurityGroupsClient() instead.
type ApplicationSecurityGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewApplicationSecurityGroupsClient creates a new instance of ApplicationSecurityGroupsClient with the specified values.
func NewApplicationSecurityGroupsClient(con *armcore.Connection, subscriptionID string) ApplicationSecurityGroupsClient {
	return ApplicationSecurityGroupsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client ApplicationSecurityGroupsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates an application security group.
func (client ApplicationSecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsBeginCreateOrUpdateOptions) (ApplicationSecurityGroupPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return ApplicationSecurityGroupPollerResponse{}, err
	}
	result := ApplicationSecurityGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ApplicationSecurityGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ApplicationSecurityGroupPollerResponse{}, err
	}
	poller := &applicationSecurityGroupPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ApplicationSecurityGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ApplicationSecurityGroupPoller from the specified resume token.
// token - The value must come from a previous call to ApplicationSecurityGroupPoller.ResumeToken().
func (client ApplicationSecurityGroupsClient) ResumeCreateOrUpdate(token string) (ApplicationSecurityGroupPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ApplicationSecurityGroupsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &applicationSecurityGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an application security group.
func (client ApplicationSecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client ApplicationSecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client ApplicationSecurityGroupsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupResponse, error) {
	result := ApplicationSecurityGroupResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ApplicationSecurityGroup)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client ApplicationSecurityGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified application security group.
func (client ApplicationSecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ApplicationSecurityGroupsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client ApplicationSecurityGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ApplicationSecurityGroupsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified application security group.
func (client ApplicationSecurityGroupsClient) delete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client ApplicationSecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client ApplicationSecurityGroupsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets information about the specified application security group.
func (client ApplicationSecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsGetOptions) (ApplicationSecurityGroupResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ApplicationSecurityGroupResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client ApplicationSecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client ApplicationSecurityGroupsClient) getHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupResponse, error) {
	result := ApplicationSecurityGroupResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ApplicationSecurityGroup)
	return result, err
}

// getHandleError handles the Get error response.
func (client ApplicationSecurityGroupsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the application security groups in a resource group.
func (client ApplicationSecurityGroupsClient) List(resourceGroupName string, options *ApplicationSecurityGroupsListOptions) ApplicationSecurityGroupListResultPager {
	return &applicationSecurityGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ApplicationSecurityGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ApplicationSecurityGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client ApplicationSecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationSecurityGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client ApplicationSecurityGroupsClient) listHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupListResultResponse, error) {
	result := ApplicationSecurityGroupListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ApplicationSecurityGroupListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client ApplicationSecurityGroupsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all application security groups in a subscription.
func (client ApplicationSecurityGroupsClient) ListAll(options *ApplicationSecurityGroupsListAllOptions) ApplicationSecurityGroupListResultPager {
	return &applicationSecurityGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp ApplicationSecurityGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ApplicationSecurityGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client ApplicationSecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *ApplicationSecurityGroupsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client ApplicationSecurityGroupsClient) listAllHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupListResultResponse, error) {
	result := ApplicationSecurityGroupListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ApplicationSecurityGroupListResult)
	return result, err
}

// listAllHandleError handles the ListAll error response.
func (client ApplicationSecurityGroupsClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates an application security group's tags.
func (client ApplicationSecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsUpdateTagsOptions) (ApplicationSecurityGroupResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ApplicationSecurityGroupResponse{}, client.updateTagsHandleError(resp)
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client ApplicationSecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client ApplicationSecurityGroupsClient) updateTagsHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupResponse, error) {
	result := ApplicationSecurityGroupResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ApplicationSecurityGroup)
	return result, err
}

// updateTagsHandleError handles the UpdateTags error response.
func (client ApplicationSecurityGroupsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
