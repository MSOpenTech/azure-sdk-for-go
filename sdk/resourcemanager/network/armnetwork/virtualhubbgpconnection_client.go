//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualHubBgpConnectionClient contains the methods for the VirtualHubBgpConnection group.
// Don't use this type directly, use NewVirtualHubBgpConnectionClient() instead.
type VirtualHubBgpConnectionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualHubBgpConnectionClient creates a new instance of VirtualHubBgpConnectionClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualHubBgpConnectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualHubBgpConnectionClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualHubBgpConnectionClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualHubBgpConnectionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a VirtualHubBgpConnection resource if it doesn't exist else updates the existing VirtualHubBgpConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - connectionName - The name of the connection.
//   - parameters - Parameters of Bgp connection.
//   - options - VirtualHubBgpConnectionClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualHubBgpConnectionClient.BeginCreateOrUpdate
//     method.
func (client *VirtualHubBgpConnectionClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters BgpConnection, options *VirtualHubBgpConnectionClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualHubBgpConnectionClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, connectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualHubBgpConnectionClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualHubBgpConnectionClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a VirtualHubBgpConnection resource if it doesn't exist else updates the existing VirtualHubBgpConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *VirtualHubBgpConnectionClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters BgpConnection, options *VirtualHubBgpConnectionClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualHubBgpConnectionClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualHubBgpConnectionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters BgpConnection, options *VirtualHubBgpConnectionClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a VirtualHubBgpConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The resource group name of the VirtualHubBgpConnection.
//   - virtualHubName - The name of the VirtualHub.
//   - connectionName - The name of the connection.
//   - options - VirtualHubBgpConnectionClientBeginDeleteOptions contains the optional parameters for the VirtualHubBgpConnectionClient.BeginDelete
//     method.
func (client *VirtualHubBgpConnectionClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionClientBeginDeleteOptions) (*runtime.Poller[VirtualHubBgpConnectionClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, connectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualHubBgpConnectionClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualHubBgpConnectionClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a VirtualHubBgpConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *VirtualHubBgpConnectionClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualHubBgpConnectionClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualHubBgpConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a Virtual Hub Bgp Connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - connectionName - The name of the connection.
//   - options - VirtualHubBgpConnectionClientGetOptions contains the optional parameters for the VirtualHubBgpConnectionClient.Get
//     method.
func (client *VirtualHubBgpConnectionClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionClientGetOptions) (VirtualHubBgpConnectionClientGetResponse, error) {
	var err error
	const operationName = "VirtualHubBgpConnectionClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return VirtualHubBgpConnectionClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualHubBgpConnectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualHubBgpConnectionClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualHubBgpConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualHubBgpConnectionClient) getHandleResponse(resp *http.Response) (VirtualHubBgpConnectionClientGetResponse, error) {
	result := VirtualHubBgpConnectionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BgpConnection); err != nil {
		return VirtualHubBgpConnectionClientGetResponse{}, err
	}
	return result, nil
}
