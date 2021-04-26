// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PrivateCloudsClient contains the methods for the PrivateClouds group.
// Don't use this type directly, use NewPrivateCloudsClient() instead.
type PrivateCloudsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateCloudsClient creates a new instance of PrivateCloudsClient with the specified values.
func NewPrivateCloudsClient(con *armcore.Connection, subscriptionID string) *PrivateCloudsClient {
	return &PrivateCloudsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a private cloud
func (client *PrivateCloudsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsBeginCreateOrUpdateOptions) (PrivateCloudPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, privateCloud, options)
	if err != nil {
		return PrivateCloudPollerResponse{}, err
	}
	result := PrivateCloudPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateCloudsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return PrivateCloudPollerResponse{}, err
	}
	poller := &privateCloudPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateCloudResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new PrivateCloudPoller from the specified resume token.
// token - The value must come from a previous call to PrivateCloudPoller.ResumeToken().
func (client *PrivateCloudsClient) ResumeCreateOrUpdate(token string) (PrivateCloudPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateCloudsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &privateCloudPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a private cloud
func (client *PrivateCloudsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, privateCloud, options)
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
func (client *PrivateCloudsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(privateCloud)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateCloudsClient) createOrUpdateHandleResponse(resp *azcore.Response) (PrivateCloudResponse, error) {
	var val *PrivateCloud
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateCloudResponse{}, err
	}
	return PrivateCloudResponse{RawResponse: resp.Response, PrivateCloud: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateCloudsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Delete a private cloud
func (client *PrivateCloudsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateCloudsClient.Delete", "", resp, client.deleteHandleError)
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
func (client *PrivateCloudsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateCloudsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Delete a private cloud
func (client *PrivateCloudsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, options)
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
func (client *PrivateCloudsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateCloudsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Get a private cloud
func (client *PrivateCloudsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsGetOptions) (PrivateCloudResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateCloudResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateCloudResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateCloudsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateCloudsClient) getHandleResponse(resp *azcore.Response) (PrivateCloudResponse, error) {
	var val *PrivateCloud
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateCloudResponse{}, err
	}
	return PrivateCloudResponse{RawResponse: resp.Response, PrivateCloud: val}, nil
}

// getHandleError handles the Get error response.
func (client *PrivateCloudsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List private clouds in a resource group
func (client *PrivateCloudsClient) List(resourceGroupName string, options *PrivateCloudsListOptions) PrivateCloudListPager {
	return &privateCloudListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp PrivateCloudListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateCloudList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateCloudsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateCloudsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateCloudsClient) listHandleResponse(resp *azcore.Response) (PrivateCloudListResponse, error) {
	var val *PrivateCloudList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateCloudListResponse{}, err
	}
	return PrivateCloudListResponse{RawResponse: resp.Response, PrivateCloudList: val}, nil
}

// listHandleError handles the List error response.
func (client *PrivateCloudsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAdminCredentials - List the admin credentials for the private cloud
func (client *PrivateCloudsClient) ListAdminCredentials(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsListAdminCredentialsOptions) (AdminCredentialsResponse, error) {
	req, err := client.listAdminCredentialsCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return AdminCredentialsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AdminCredentialsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AdminCredentialsResponse{}, client.listAdminCredentialsHandleError(resp)
	}
	return client.listAdminCredentialsHandleResponse(resp)
}

// listAdminCredentialsCreateRequest creates the ListAdminCredentials request.
func (client *PrivateCloudsClient) listAdminCredentialsCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsListAdminCredentialsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/listAdminCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAdminCredentialsHandleResponse handles the ListAdminCredentials response.
func (client *PrivateCloudsClient) listAdminCredentialsHandleResponse(resp *azcore.Response) (AdminCredentialsResponse, error) {
	var val *AdminCredentials
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AdminCredentialsResponse{}, err
	}
	return AdminCredentialsResponse{RawResponse: resp.Response, AdminCredentials: val}, nil
}

// listAdminCredentialsHandleError handles the ListAdminCredentials error response.
func (client *PrivateCloudsClient) listAdminCredentialsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListInSubscription - List private clouds in a subscription
func (client *PrivateCloudsClient) ListInSubscription(options *PrivateCloudsListInSubscriptionOptions) PrivateCloudListPager {
	return &privateCloudListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listInSubscriptionCreateRequest(ctx, options)
		},
		responder: client.listInSubscriptionHandleResponse,
		errorer:   client.listInSubscriptionHandleError,
		advancer: func(ctx context.Context, resp PrivateCloudListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateCloudList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listInSubscriptionCreateRequest creates the ListInSubscription request.
func (client *PrivateCloudsClient) listInSubscriptionCreateRequest(ctx context.Context, options *PrivateCloudsListInSubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AVS/privateClouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listInSubscriptionHandleResponse handles the ListInSubscription response.
func (client *PrivateCloudsClient) listInSubscriptionHandleResponse(resp *azcore.Response) (PrivateCloudListResponse, error) {
	var val *PrivateCloudList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateCloudListResponse{}, err
	}
	return PrivateCloudListResponse{RawResponse: resp.Response, PrivateCloudList: val}, nil
}

// listInSubscriptionHandleError handles the ListInSubscription error response.
func (client *PrivateCloudsClient) listInSubscriptionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdate - Update a private cloud
func (client *PrivateCloudsClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsBeginUpdateOptions) (PrivateCloudPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, privateCloudName, privateCloudUpdate, options)
	if err != nil {
		return PrivateCloudPollerResponse{}, err
	}
	result := PrivateCloudPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateCloudsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return PrivateCloudPollerResponse{}, err
	}
	poller := &privateCloudPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateCloudResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new PrivateCloudPoller from the specified resume token.
// token - The value must come from a previous call to PrivateCloudPoller.ResumeToken().
func (client *PrivateCloudsClient) ResumeUpdate(token string) (PrivateCloudPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateCloudsClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &privateCloudPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - Update a private cloud
func (client *PrivateCloudsClient) update(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateCloudName, privateCloudUpdate, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *PrivateCloudsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-20")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(privateCloudUpdate)
}

// updateHandleResponse handles the Update response.
func (client *PrivateCloudsClient) updateHandleResponse(resp *azcore.Response) (PrivateCloudResponse, error) {
	var val *PrivateCloud
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateCloudResponse{}, err
	}
	return PrivateCloudResponse{RawResponse: resp.Response, PrivateCloud: val}, nil
}

// updateHandleError handles the Update error response.
func (client *PrivateCloudsClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
