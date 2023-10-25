//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazureintegrationspaces

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BusinessProcessesClient contains the methods for the BusinessProcesses group.
// Don't use this type directly, use NewBusinessProcessesClient() instead.
type BusinessProcessesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBusinessProcessesClient creates a new instance of BusinessProcessesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBusinessProcessesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BusinessProcessesClient, error) {
	cl, err := arm.NewClient(moduleName+".BusinessProcessesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BusinessProcessesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a BusinessProcess
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-14-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spaceName - The name of the space
//   - applicationName - The name of the Application
//   - businessProcessName - The name of the business process
//   - resource - Resource create parameters.
//   - options - BusinessProcessesClientCreateOrUpdateOptions contains the optional parameters for the BusinessProcessesClient.CreateOrUpdate
//     method.
func (client *BusinessProcessesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, resource BusinessProcess, options *BusinessProcessesClientCreateOrUpdateOptions) (BusinessProcessesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, spaceName, applicationName, businessProcessName, resource, options)
	if err != nil {
		return BusinessProcessesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessProcessesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return BusinessProcessesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BusinessProcessesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, resource BusinessProcess, options *BusinessProcessesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IntegrationSpaces/spaces/{spaceName}/applications/{applicationName}/businessProcesses/{businessProcessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spaceName == "" {
		return nil, errors.New("parameter spaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spaceName}", url.PathEscape(spaceName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if businessProcessName == "" {
		return nil, errors.New("parameter businessProcessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessProcessName}", url.PathEscape(businessProcessName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-14-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BusinessProcessesClient) createOrUpdateHandleResponse(resp *http.Response) (BusinessProcessesClientCreateOrUpdateResponse, error) {
	result := BusinessProcessesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessProcess); err != nil {
		return BusinessProcessesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a BusinessProcess
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-14-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spaceName - The name of the space
//   - applicationName - The name of the Application
//   - businessProcessName - The name of the business process
//   - options - BusinessProcessesClientDeleteOptions contains the optional parameters for the BusinessProcessesClient.Delete
//     method.
func (client *BusinessProcessesClient) Delete(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, options *BusinessProcessesClientDeleteOptions) (BusinessProcessesClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, spaceName, applicationName, businessProcessName, options)
	if err != nil {
		return BusinessProcessesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessProcessesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BusinessProcessesClientDeleteResponse{}, err
	}
	return BusinessProcessesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BusinessProcessesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, options *BusinessProcessesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IntegrationSpaces/spaces/{spaceName}/applications/{applicationName}/businessProcesses/{businessProcessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spaceName == "" {
		return nil, errors.New("parameter spaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spaceName}", url.PathEscape(spaceName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if businessProcessName == "" {
		return nil, errors.New("parameter businessProcessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessProcessName}", url.PathEscape(businessProcessName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-14-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a BusinessProcess
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-14-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spaceName - The name of the space
//   - applicationName - The name of the Application
//   - businessProcessName - The name of the business process
//   - options - BusinessProcessesClientGetOptions contains the optional parameters for the BusinessProcessesClient.Get method.
func (client *BusinessProcessesClient) Get(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, options *BusinessProcessesClientGetOptions) (BusinessProcessesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, spaceName, applicationName, businessProcessName, options)
	if err != nil {
		return BusinessProcessesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessProcessesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BusinessProcessesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BusinessProcessesClient) getCreateRequest(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, options *BusinessProcessesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IntegrationSpaces/spaces/{spaceName}/applications/{applicationName}/businessProcesses/{businessProcessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spaceName == "" {
		return nil, errors.New("parameter spaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spaceName}", url.PathEscape(spaceName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if businessProcessName == "" {
		return nil, errors.New("parameter businessProcessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessProcessName}", url.PathEscape(businessProcessName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-14-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BusinessProcessesClient) getHandleResponse(resp *http.Response) (BusinessProcessesClientGetResponse, error) {
	result := BusinessProcessesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessProcess); err != nil {
		return BusinessProcessesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByApplicationPager - List BusinessProcess resources by Application
//
// Generated from API version 2023-11-14-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spaceName - The name of the space
//   - applicationName - The name of the Application
//   - options - BusinessProcessesClientListByApplicationOptions contains the optional parameters for the BusinessProcessesClient.NewListByApplicationPager
//     method.
func (client *BusinessProcessesClient) NewListByApplicationPager(resourceGroupName string, spaceName string, applicationName string, options *BusinessProcessesClientListByApplicationOptions) *runtime.Pager[BusinessProcessesClientListByApplicationResponse] {
	return runtime.NewPager(runtime.PagingHandler[BusinessProcessesClientListByApplicationResponse]{
		More: func(page BusinessProcessesClientListByApplicationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BusinessProcessesClientListByApplicationResponse) (BusinessProcessesClientListByApplicationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByApplicationCreateRequest(ctx, resourceGroupName, spaceName, applicationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BusinessProcessesClientListByApplicationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BusinessProcessesClientListByApplicationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BusinessProcessesClientListByApplicationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByApplicationHandleResponse(resp)
		},
	})
}

// listByApplicationCreateRequest creates the ListByApplication request.
func (client *BusinessProcessesClient) listByApplicationCreateRequest(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, options *BusinessProcessesClientListByApplicationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IntegrationSpaces/spaces/{spaceName}/applications/{applicationName}/businessProcesses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spaceName == "" {
		return nil, errors.New("parameter spaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spaceName}", url.PathEscape(spaceName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-14-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Select != nil {
		for _, qv := range options.Select {
			reqQP.Add("select", qv)
		}
	}
	if options != nil && options.Expand != nil {
		for _, qv := range options.Expand {
			reqQP.Add("expand", qv)
		}
	}
	if options != nil && options.Orderby != nil {
		for _, qv := range options.Orderby {
			reqQP.Add("orderby", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByApplicationHandleResponse handles the ListByApplication response.
func (client *BusinessProcessesClient) listByApplicationHandleResponse(resp *http.Response) (BusinessProcessesClientListByApplicationResponse, error) {
	result := BusinessProcessesClientListByApplicationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessProcessListResult); err != nil {
		return BusinessProcessesClientListByApplicationResponse{}, err
	}
	return result, nil
}

// Patch - Update a BusinessProcess
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-14-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spaceName - The name of the space
//   - applicationName - The name of the Application
//   - businessProcessName - The name of the business process
//   - properties - The resource properties to be updated.
//   - options - BusinessProcessesClientPatchOptions contains the optional parameters for the BusinessProcessesClient.Patch method.
func (client *BusinessProcessesClient) Patch(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, properties BusinessProcessUpdate, options *BusinessProcessesClientPatchOptions) (BusinessProcessesClientPatchResponse, error) {
	var err error
	req, err := client.patchCreateRequest(ctx, resourceGroupName, spaceName, applicationName, businessProcessName, properties, options)
	if err != nil {
		return BusinessProcessesClientPatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessProcessesClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BusinessProcessesClientPatchResponse{}, err
	}
	resp, err := client.patchHandleResponse(httpResp)
	return resp, err
}

// patchCreateRequest creates the Patch request.
func (client *BusinessProcessesClient) patchCreateRequest(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, businessProcessName string, properties BusinessProcessUpdate, options *BusinessProcessesClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IntegrationSpaces/spaces/{spaceName}/applications/{applicationName}/businessProcesses/{businessProcessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spaceName == "" {
		return nil, errors.New("parameter spaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spaceName}", url.PathEscape(spaceName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if businessProcessName == "" {
		return nil, errors.New("parameter businessProcessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessProcessName}", url.PathEscape(businessProcessName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-14-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// patchHandleResponse handles the Patch response.
func (client *BusinessProcessesClient) patchHandleResponse(resp *http.Response) (BusinessProcessesClientPatchResponse, error) {
	result := BusinessProcessesClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessProcess); err != nil {
		return BusinessProcessesClientPatchResponse{}, err
	}
	return result, nil
}
