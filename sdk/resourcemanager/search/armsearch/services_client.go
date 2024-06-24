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

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks whether or not the given search service name is available for use. Search service names
// must be globally unique since they are part of the service URI (https://.search.windows.net).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - checkNameAvailabilityInput - The resource name and type to check.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
//     method.
func (client *ServicesClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientCheckNameAvailabilityOptions) (ServicesClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "ServicesClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityInput, searchManagementRequestOptions, options)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServicesClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Search/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
	if err := runtime.MarshalAsJSON(req, checkNameAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServicesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServicesClientCheckNameAvailabilityResponse, error) {
	result := ServicesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates a search service in the given resource group. If the search service already exists,
// all properties will be updated with the given values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service to create or update. Search service names must only contain
//     lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot
//     contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique
//     since they are part of the service URI (https://.search.windows.net). You
//     cannot change the service name after the service is created.
//   - service - The definition of the search service to create or update.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate
//     method.
func (client *ServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, searchServiceName, service, searchManagementRequestOptions, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a search service in the given resource group. If the search service already exists,
// all properties will be updated with the given values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *ServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServicesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, searchServiceName, service, searchManagementRequestOptions, options)
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
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
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
	if err := runtime.MarshalAsJSON(req, service); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes a search service in the given resource group, along with its associated resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientDeleteOptions contains the optional parameters for the ServicesClient.Delete method.
func (client *ServicesClient) Delete(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientDeleteOptions) (ServicesClientDeleteResponse, error) {
	var err error
	const operationName = "ServicesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientDeleteResponse{}, err
	}
	return ServicesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
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

// Get - Gets the search service with the given name in the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	var err error
	const operationName = "ServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
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
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of all Search services in the given resource group.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
//     method.
func (client *ServicesClient) NewListByResourceGroupPager(resourceGroupName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientListByResourceGroupOptions) *runtime.Pager[ServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListByResourceGroupResponse]{
		More: func(page ServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListByResourceGroupResponse) (ServicesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServicesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, searchManagementRequestOptions, options)
			}, nil)
			if err != nil {
				return ServicesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServicesClientListByResourceGroupResponse, error) {
	result := ServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of all Search services in the given subscription.
//
// Generated from API version 2024-06-01-preview
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientListBySubscriptionOptions contains the optional parameters for the ServicesClient.NewListBySubscriptionPager
//     method.
func (client *ServicesClient) NewListBySubscriptionPager(searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientListBySubscriptionOptions) *runtime.Pager[ServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListBySubscriptionResponse]{
		More: func(page ServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListBySubscriptionResponse) (ServicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, searchManagementRequestOptions, options)
			}, nil)
			if err != nil {
				return ServicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ServicesClient) listBySubscriptionCreateRequest(ctx context.Context, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Search/searchServices"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (ServicesClientListBySubscriptionResponse, error) {
	result := ServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing search service in the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service to update.
//   - service - The definition of the search service to update.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ServicesClientUpdateOptions contains the optional parameters for the ServicesClient.Update method.
func (client *ServicesClient) Update(ctx context.Context, resourceGroupName string, searchServiceName string, service ServiceUpdate, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientUpdateOptions) (ServicesClientUpdateResponse, error) {
	var err error
	const operationName = "ServicesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, searchServiceName, service, searchManagementRequestOptions, options)
	if err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, service ServiceUpdate, searchManagementRequestOptions *SearchManagementRequestOptions, options *ServicesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
	if err := runtime.MarshalAsJSON(req, service); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ServicesClient) updateHandleResponse(resp *http.Response) (ServicesClientUpdateResponse, error) {
	result := ServicesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientUpdateResponse{}, err
	}
	return result, nil
}
