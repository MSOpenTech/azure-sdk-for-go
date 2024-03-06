//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armconnectedvmware

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

// InventoryItemsClient contains the methods for the InventoryItems group.
// Don't use this type directly, use NewInventoryItemsClient() instead.
type InventoryItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInventoryItemsClient creates a new instance of InventoryItemsClient with the specified values.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInventoryItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InventoryItemsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InventoryItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create Or Update InventoryItem.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - vcenterName - Name of the vCenter.
//   - inventoryItemName - Name of the inventoryItem.
//   - body - Request payload.
//   - options - InventoryItemsClientCreateOptions contains the optional parameters for the InventoryItemsClient.Create method.
func (client *InventoryItemsClient) Create(ctx context.Context, resourceGroupName string, vcenterName string, inventoryItemName string, body InventoryItem, options *InventoryItemsClientCreateOptions) (InventoryItemsClientCreateResponse, error) {
	var err error
	const operationName = "InventoryItemsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, vcenterName, inventoryItemName, body, options)
	if err != nil {
		return InventoryItemsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InventoryItemsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InventoryItemsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *InventoryItemsClient) createCreateRequest(ctx context.Context, resourceGroupName string, vcenterName string, inventoryItemName string, body InventoryItem, options *InventoryItemsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/vcenters/{vcenterName}/inventoryItems/{inventoryItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	if inventoryItemName == "" {
		return nil, errors.New("parameter inventoryItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inventoryItemName}", url.PathEscape(inventoryItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *InventoryItemsClient) createHandleResponse(resp *http.Response) (InventoryItemsClientCreateResponse, error) {
	result := InventoryItemsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InventoryItem); err != nil {
		return InventoryItemsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Implements inventoryItem DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - vcenterName - Name of the vCenter.
//   - inventoryItemName - Name of the inventoryItem.
//   - options - InventoryItemsClientDeleteOptions contains the optional parameters for the InventoryItemsClient.Delete method.
func (client *InventoryItemsClient) Delete(ctx context.Context, resourceGroupName string, vcenterName string, inventoryItemName string, options *InventoryItemsClientDeleteOptions) (InventoryItemsClientDeleteResponse, error) {
	var err error
	const operationName = "InventoryItemsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vcenterName, inventoryItemName, options)
	if err != nil {
		return InventoryItemsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InventoryItemsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return InventoryItemsClientDeleteResponse{}, err
	}
	return InventoryItemsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InventoryItemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vcenterName string, inventoryItemName string, options *InventoryItemsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/vcenters/{vcenterName}/inventoryItems/{inventoryItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	if inventoryItemName == "" {
		return nil, errors.New("parameter inventoryItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inventoryItemName}", url.PathEscape(inventoryItemName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements InventoryItem GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - vcenterName - Name of the vCenter.
//   - inventoryItemName - Name of the inventoryItem.
//   - options - InventoryItemsClientGetOptions contains the optional parameters for the InventoryItemsClient.Get method.
func (client *InventoryItemsClient) Get(ctx context.Context, resourceGroupName string, vcenterName string, inventoryItemName string, options *InventoryItemsClientGetOptions) (InventoryItemsClientGetResponse, error) {
	var err error
	const operationName = "InventoryItemsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vcenterName, inventoryItemName, options)
	if err != nil {
		return InventoryItemsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InventoryItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InventoryItemsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InventoryItemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vcenterName string, inventoryItemName string, options *InventoryItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/vcenters/{vcenterName}/inventoryItems/{inventoryItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	if inventoryItemName == "" {
		return nil, errors.New("parameter inventoryItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inventoryItemName}", url.PathEscape(inventoryItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InventoryItemsClient) getHandleResponse(resp *http.Response) (InventoryItemsClientGetResponse, error) {
	result := InventoryItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InventoryItem); err != nil {
		return InventoryItemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVCenterPager - Returns the list of inventoryItems of the given vCenter.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - vcenterName - Name of the vCenter.
//   - options - InventoryItemsClientListByVCenterOptions contains the optional parameters for the InventoryItemsClient.NewListByVCenterPager
//     method.
func (client *InventoryItemsClient) NewListByVCenterPager(resourceGroupName string, vcenterName string, options *InventoryItemsClientListByVCenterOptions) *runtime.Pager[InventoryItemsClientListByVCenterResponse] {
	return runtime.NewPager(runtime.PagingHandler[InventoryItemsClientListByVCenterResponse]{
		More: func(page InventoryItemsClientListByVCenterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InventoryItemsClientListByVCenterResponse) (InventoryItemsClientListByVCenterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InventoryItemsClient.NewListByVCenterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVCenterCreateRequest(ctx, resourceGroupName, vcenterName, options)
			}, nil)
			if err != nil {
				return InventoryItemsClientListByVCenterResponse{}, err
			}
			return client.listByVCenterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVCenterCreateRequest creates the ListByVCenter request.
func (client *InventoryItemsClient) listByVCenterCreateRequest(ctx context.Context, resourceGroupName string, vcenterName string, options *InventoryItemsClientListByVCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/vcenters/{vcenterName}/inventoryItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVCenterHandleResponse handles the ListByVCenter response.
func (client *InventoryItemsClient) listByVCenterHandleResponse(resp *http.Response) (InventoryItemsClientListByVCenterResponse, error) {
	result := InventoryItemsClientListByVCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InventoryItemsList); err != nil {
		return InventoryItemsClientListByVCenterResponse{}, err
	}
	return result, nil
}
