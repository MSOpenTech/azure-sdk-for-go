//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// DataCollectionRulesClient contains the methods for the DataCollectionRules group.
// Don't use this type directly, use NewDataCollectionRulesClient() instead.
type DataCollectionRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataCollectionRulesClient creates a new instance of DataCollectionRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataCollectionRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataCollectionRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".DataCollectionRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataCollectionRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or updates a data collection rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataCollectionRuleName - The name of the data collection rule. The name is case insensitive.
//   - options - DataCollectionRulesClientCreateOptions contains the optional parameters for the DataCollectionRulesClient.Create
//     method.
func (client *DataCollectionRulesClient) Create(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientCreateOptions) (DataCollectionRulesClientCreateResponse, error) {
	var err error
	const operationName = "DataCollectionRulesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRulesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRulesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRulesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *DataCollectionRulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DataCollectionRulesClient) createHandleResponse(resp *http.Response) (DataCollectionRulesClientCreateResponse, error) {
	result := DataCollectionRulesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleResource); err != nil {
		return DataCollectionRulesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a data collection rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataCollectionRuleName - The name of the data collection rule. The name is case insensitive.
//   - options - DataCollectionRulesClientDeleteOptions contains the optional parameters for the DataCollectionRulesClient.Delete
//     method.
func (client *DataCollectionRulesClient) Delete(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientDeleteOptions) (DataCollectionRulesClientDeleteResponse, error) {
	var err error
	const operationName = "DataCollectionRulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRulesClientDeleteResponse{}, err
	}
	return DataCollectionRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataCollectionRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the specified data collection rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataCollectionRuleName - The name of the data collection rule. The name is case insensitive.
//   - options - DataCollectionRulesClientGetOptions contains the optional parameters for the DataCollectionRulesClient.Get method.
func (client *DataCollectionRulesClient) Get(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientGetOptions) (DataCollectionRulesClientGetResponse, error) {
	var err error
	const operationName = "DataCollectionRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataCollectionRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataCollectionRulesClient) getHandleResponse(resp *http.Response) (DataCollectionRulesClientGetResponse, error) {
	result := DataCollectionRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleResource); err != nil {
		return DataCollectionRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all data collection rules in the specified resource group.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DataCollectionRulesClientListByResourceGroupOptions contains the optional parameters for the DataCollectionRulesClient.NewListByResourceGroupPager
//     method.
func (client *DataCollectionRulesClient) NewListByResourceGroupPager(resourceGroupName string, options *DataCollectionRulesClientListByResourceGroupOptions) *runtime.Pager[DataCollectionRulesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionRulesClientListByResourceGroupResponse]{
		More: func(page DataCollectionRulesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRulesClientListByResourceGroupResponse) (DataCollectionRulesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataCollectionRulesClient.NewListByResourceGroupPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionRulesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataCollectionRulesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionRulesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataCollectionRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DataCollectionRulesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataCollectionRulesClient) listByResourceGroupHandleResponse(resp *http.Response) (DataCollectionRulesClientListByResourceGroupResponse, error) {
	result := DataCollectionRulesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleResourceListResult); err != nil {
		return DataCollectionRulesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all data collection rules in the specified subscription.
//
// Generated from API version 2022-06-01
//   - options - DataCollectionRulesClientListBySubscriptionOptions contains the optional parameters for the DataCollectionRulesClient.NewListBySubscriptionPager
//     method.
func (client *DataCollectionRulesClient) NewListBySubscriptionPager(options *DataCollectionRulesClientListBySubscriptionOptions) *runtime.Pager[DataCollectionRulesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionRulesClientListBySubscriptionResponse]{
		More: func(page DataCollectionRulesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRulesClientListBySubscriptionResponse) (DataCollectionRulesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataCollectionRulesClient.NewListBySubscriptionPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionRulesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataCollectionRulesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionRulesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DataCollectionRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DataCollectionRulesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/dataCollectionRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DataCollectionRulesClient) listBySubscriptionHandleResponse(resp *http.Response) (DataCollectionRulesClientListBySubscriptionResponse, error) {
	result := DataCollectionRulesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleResourceListResult); err != nil {
		return DataCollectionRulesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates part of a data collection rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataCollectionRuleName - The name of the data collection rule. The name is case insensitive.
//   - options - DataCollectionRulesClientUpdateOptions contains the optional parameters for the DataCollectionRulesClient.Update
//     method.
func (client *DataCollectionRulesClient) Update(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientUpdateOptions) (DataCollectionRulesClientUpdateResponse, error) {
	var err error
	const operationName = "DataCollectionRulesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRulesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRulesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DataCollectionRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DataCollectionRulesClient) updateHandleResponse(resp *http.Response) (DataCollectionRulesClientUpdateResponse, error) {
	result := DataCollectionRulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleResource); err != nil {
		return DataCollectionRulesClientUpdateResponse{}, err
	}
	return result, nil
}
