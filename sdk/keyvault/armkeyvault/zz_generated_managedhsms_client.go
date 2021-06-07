// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ManagedHsmsClient contains the methods for the ManagedHsms group.
// Don't use this type directly, use NewManagedHsmsClient() instead.
type ManagedHsmsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewManagedHsmsClient creates a new instance of ManagedHsmsClient with the specified values.
func NewManagedHsmsClient(con *armcore.Connection, subscriptionID string) *ManagedHsmsClient {
	return &ManagedHsmsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a managed HSM Pool in the specified subscription.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsBeginCreateOrUpdateOptions) (ManagedHsmPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	result := ManagedHsmPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ManagedHsmsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	poller := &managedHsmPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedHsmResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ManagedHsmPoller from the specified resume token.
// token - The value must come from a previous call to ManagedHsmPoller.ResumeToken().
func (client *ManagedHsmsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ManagedHsmPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ManagedHsmsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	poller := &managedHsmPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	result := ManagedHsmPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedHsmResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Create or update a managed HSM Pool in the specified subscription.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedHsmsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ManagedHsmsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified managed HSM Pool.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ManagedHsmsClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ManagedHsmsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ManagedHsmsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified managed HSM Pool.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *ManagedHsmsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ManagedHsmsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified managed HSM Pool.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) Get(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsGetOptions) (ManagedHsmResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return ManagedHsmResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagedHsmResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ManagedHsmResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedHsmsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedHsmsClient) getHandleResponse(resp *azcore.Response) (ManagedHsmResponse, error) {
	var val *ManagedHsm
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ManagedHsmResponse{}, err
	}
	return ManagedHsmResponse{RawResponse: resp.Response, ManagedHsm: val}, nil
}

// getHandleError handles the Get error response.
func (client *ManagedHsmsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetDeleted - Gets the specified deleted managed HSM.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) GetDeleted(ctx context.Context, name string, location string, options *ManagedHsmsGetDeletedOptions) (DeletedManagedHsmResponse, error) {
	req, err := client.getDeletedCreateRequest(ctx, name, location, options)
	if err != nil {
		return DeletedManagedHsmResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DeletedManagedHsmResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DeletedManagedHsmResponse{}, client.getDeletedHandleError(resp)
	}
	return client.getDeletedHandleResponse(resp)
}

// getDeletedCreateRequest creates the GetDeleted request.
func (client *ManagedHsmsClient) getDeletedCreateRequest(ctx context.Context, name string, location string, options *ManagedHsmsGetDeletedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedManagedHSMs/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDeletedHandleResponse handles the GetDeleted response.
func (client *ManagedHsmsClient) getDeletedHandleResponse(resp *azcore.Response) (DeletedManagedHsmResponse, error) {
	var val *DeletedManagedHsm
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DeletedManagedHsmResponse{}, err
	}
	return DeletedManagedHsmResponse{RawResponse: resp.Response, DeletedManagedHsm: val}, nil
}

// getDeletedHandleError handles the GetDeleted error response.
func (client *ManagedHsmsClient) getDeletedHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - The List operation gets information about the managed HSM Pools associated with the subscription and within the specified resource
// group.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) ListByResourceGroup(resourceGroupName string, options *ManagedHsmsListByResourceGroupOptions) ManagedHsmListResultPager {
	return &managedHsmListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ManagedHsmListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedHsmListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedHsmsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedHsmsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedHsmsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ManagedHsmListResultResponse, error) {
	var val *ManagedHsmListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ManagedHsmListResultResponse{}, err
	}
	return ManagedHsmListResultResponse{RawResponse: resp.Response, ManagedHsmListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ManagedHsmsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListBySubscription - The List operation gets information about the managed HSM Pools associated with the subscription.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) ListBySubscription(options *ManagedHsmsListBySubscriptionOptions) ManagedHsmListResultPager {
	return &managedHsmListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp ManagedHsmListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ManagedHsmListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagedHsmsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagedHsmsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/managedHSMs"
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagedHsmsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (ManagedHsmListResultResponse, error) {
	var val *ManagedHsmListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ManagedHsmListResultResponse{}, err
	}
	return ManagedHsmListResultResponse{RawResponse: resp.Response, ManagedHsmListResult: val}, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ManagedHsmsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListDeleted - The List operation gets information about the deleted managed HSMs associated with the subscription.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) ListDeleted(options *ManagedHsmsListDeletedOptions) DeletedManagedHsmListResultPager {
	return &deletedManagedHsmListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listDeletedCreateRequest(ctx, options)
		},
		responder: client.listDeletedHandleResponse,
		errorer:   client.listDeletedHandleError,
		advancer: func(ctx context.Context, resp DeletedManagedHsmListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DeletedManagedHsmListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listDeletedCreateRequest creates the ListDeleted request.
func (client *ManagedHsmsClient) listDeletedCreateRequest(ctx context.Context, options *ManagedHsmsListDeletedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedManagedHSMs"
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
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listDeletedHandleResponse handles the ListDeleted response.
func (client *ManagedHsmsClient) listDeletedHandleResponse(resp *azcore.Response) (DeletedManagedHsmListResultResponse, error) {
	var val *DeletedManagedHsmListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DeletedManagedHsmListResultResponse{}, err
	}
	return DeletedManagedHsmListResultResponse{RawResponse: resp.Response, DeletedManagedHsmListResult: val}, nil
}

// listDeletedHandleError handles the ListDeleted error response.
func (client *ManagedHsmsClient) listDeletedHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginPurgeDeleted - Permanently deletes the specified managed HSM.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) BeginPurgeDeleted(ctx context.Context, name string, location string, options *ManagedHsmsBeginPurgeDeletedOptions) (HTTPPollerResponse, error) {
	resp, err := client.purgeDeleted(ctx, name, location, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ManagedHsmsClient.PurgeDeleted", "", resp, client.purgeDeletedHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePurgeDeleted creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ManagedHsmsClient) ResumePurgeDeleted(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ManagedHsmsClient.PurgeDeleted", token, client.purgeDeletedHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// PurgeDeleted - Permanently deletes the specified managed HSM.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) purgeDeleted(ctx context.Context, name string, location string, options *ManagedHsmsBeginPurgeDeletedOptions) (*azcore.Response, error) {
	req, err := client.purgeDeletedCreateRequest(ctx, name, location, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.purgeDeletedHandleError(resp)
	}
	return resp, nil
}

// purgeDeletedCreateRequest creates the PurgeDeleted request.
func (client *ManagedHsmsClient) purgeDeletedCreateRequest(ctx context.Context, name string, location string, options *ManagedHsmsBeginPurgeDeletedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedManagedHSMs/{name}/purge"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// purgeDeletedHandleError handles the PurgeDeleted error response.
func (client *ManagedHsmsClient) purgeDeletedHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - Update a managed HSM Pool in the specified subscription.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) BeginUpdate(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsBeginUpdateOptions) (ManagedHsmPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	result := ManagedHsmPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ManagedHsmsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	poller := &managedHsmPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedHsmResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ManagedHsmPoller from the specified resume token.
// token - The value must come from a previous call to ManagedHsmPoller.ResumeToken().
func (client *ManagedHsmsClient) ResumeUpdate(ctx context.Context, token string) (ManagedHsmPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ManagedHsmsClient.Update", token, client.updateHandleError)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	poller := &managedHsmPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ManagedHsmPollerResponse{}, err
	}
	result := ManagedHsmPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ManagedHsmResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Update a managed HSM Pool in the specified subscription.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *ManagedHsmsClient) update(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedHsmsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleError handles the Update error response.
func (client *ManagedHsmsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
