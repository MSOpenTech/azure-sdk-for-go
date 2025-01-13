//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Disconnects the private endpoint connection and deletes it from the search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - privateEndpointConnectionName - The name of the private endpoint connection to the Azure AI Search service with the specified
//     resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - PrivateEndpointConnectionsClientDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Delete
//     method.
func (client *PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientDeleteOptions) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, searchServiceName, privateEndpointConnectionName, searchManagementRequestOptions, options)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *PrivateEndpointConnectionsClient) deleteHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	result := PrivateEndpointConnectionsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the private endpoint connection to the search service in the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - privateEndpointConnectionName - The name of the private endpoint connection to the Azure AI Search service with the specified
//     resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
//     method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, searchServiceName, privateEndpointConnectionName, searchManagementRequestOptions, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Gets a list of all private endpoint connections in the given service.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - PrivateEndpointConnectionsClientListByServiceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByServicePager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByServicePager(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientListByServiceOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByServiceResponse]{
		More: func(page PrivateEndpointConnectionsClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByServiceResponse) (PrivateEndpointConnectionsClientListByServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateEndpointConnectionsClient.NewListByServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServiceCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
			}, nil)
			if err != nil {
				return PrivateEndpointConnectionsClientListByServiceResponse{}, err
			}
			return client.listByServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PrivateEndpointConnectionsClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PrivateEndpointConnectionsClient) listByServiceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByServiceResponse, error) {
	result := PrivateEndpointConnectionsClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates a private endpoint connection to the search service in the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - privateEndpointConnectionName - The name of the private endpoint connection to the Azure AI Search service with the specified
//     resource group.
//   - privateEndpointConnection - The definition of the private endpoint connection to update.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - PrivateEndpointConnectionsClientUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Update
//     method.
func (client *PrivateEndpointConnectionsClient) Update(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientUpdateOptions) (PrivateEndpointConnectionsClientUpdateResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, searchServiceName, privateEndpointConnectionName, privateEndpointConnection, searchManagementRequestOptions, options)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *PrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateEndpointConnectionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	if err := runtime.MarshalAsJSON(req, privateEndpointConnection); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PrivateEndpointConnectionsClient) updateHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientUpdateResponse, error) {
	result := PrivateEndpointConnectionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	return result, nil
}
