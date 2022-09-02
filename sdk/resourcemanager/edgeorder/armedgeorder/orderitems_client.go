//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armedgeorder

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// OrderItemsClient contains the methods for the OrderItems group.
// Don't use this type directly, use NewOrderItemsClient() instead.
type OrderItemsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOrderItemsClient creates a new instance of OrderItemsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOrderItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrderItemsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &OrderItemsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Cancel - Cancel order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// orderItemName - The name of the order item.
// cancellationReason - Reason for cancellation.
// options - OrderItemsClientCancelOptions contains the optional parameters for the OrderItemsClient.Cancel method.
func (client *OrderItemsClient) Cancel(ctx context.Context, resourceGroupName string, orderItemName string, cancellationReason CancellationReason, options *OrderItemsClientCancelOptions) (OrderItemsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, orderItemName, cancellationReason, options)
	if err != nil {
		return OrderItemsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrderItemsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return OrderItemsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return OrderItemsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *OrderItemsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, cancellationReason CancellationReason, options *OrderItemsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, cancellationReason)
}

// BeginCreate - Create an order item. Existing order item cannot be updated with this api and should instead be updated with
// the Update order item API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// orderItemName - The name of the order item.
// orderItemResource - Order item details from request body.
// options - OrderItemsClientBeginCreateOptions contains the optional parameters for the OrderItemsClient.BeginCreate method.
func (client *OrderItemsClient) BeginCreate(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource OrderItemResource, options *OrderItemsClientBeginCreateOptions) (*runtime.Poller[OrderItemsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, orderItemName, orderItemResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OrderItemsClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OrderItemsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create an order item. Existing order item cannot be updated with this api and should instead be updated with the
// Update order item API.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *OrderItemsClient) create(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource OrderItemResource, options *OrderItemsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, orderItemName, orderItemResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OrderItemsClient) createCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource OrderItemResource, options *OrderItemsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, orderItemResource)
}

// BeginDelete - Delete an order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// orderItemName - The name of the order item.
// options - OrderItemsClientBeginDeleteOptions contains the optional parameters for the OrderItemsClient.BeginDelete method.
func (client *OrderItemsClient) BeginDelete(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientBeginDeleteOptions) (*runtime.Poller[OrderItemsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, orderItemName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OrderItemsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OrderItemsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *OrderItemsClient) deleteOperation(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, orderItemName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OrderItemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// orderItemName - The name of the order item.
// options - OrderItemsClientGetOptions contains the optional parameters for the OrderItemsClient.Get method.
func (client *OrderItemsClient) Get(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientGetOptions) (OrderItemsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, orderItemName, options)
	if err != nil {
		return OrderItemsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrderItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrderItemsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OrderItemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrderItemsClient) getHandleResponse(resp *http.Response) (OrderItemsClientGetResponse, error) {
	result := OrderItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderItemResource); err != nil {
		return OrderItemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List order items at resource group level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - OrderItemsClientListByResourceGroupOptions contains the optional parameters for the OrderItemsClient.ListByResourceGroup
// method.
func (client *OrderItemsClient) NewListByResourceGroupPager(resourceGroupName string, options *OrderItemsClientListByResourceGroupOptions) *runtime.Pager[OrderItemsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrderItemsClientListByResourceGroupResponse]{
		More: func(page OrderItemsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrderItemsClientListByResourceGroupResponse) (OrderItemsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrderItemsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrderItemsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrderItemsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OrderItemsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OrderItemsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OrderItemsClient) listByResourceGroupHandleResponse(resp *http.Response) (OrderItemsClientListByResourceGroupResponse, error) {
	result := OrderItemsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderItemResourceList); err != nil {
		return OrderItemsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List order items at subscription level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// options - OrderItemsClientListBySubscriptionOptions contains the optional parameters for the OrderItemsClient.ListBySubscription
// method.
func (client *OrderItemsClient) NewListBySubscriptionPager(options *OrderItemsClientListBySubscriptionOptions) *runtime.Pager[OrderItemsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrderItemsClientListBySubscriptionResponse]{
		More: func(page OrderItemsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrderItemsClientListBySubscriptionResponse) (OrderItemsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrderItemsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrderItemsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrderItemsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OrderItemsClient) listBySubscriptionCreateRequest(ctx context.Context, options *OrderItemsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/orderItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OrderItemsClient) listBySubscriptionHandleResponse(resp *http.Response) (OrderItemsClientListBySubscriptionResponse, error) {
	result := OrderItemsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderItemResourceList); err != nil {
		return OrderItemsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginReturn - Return order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// orderItemName - The name of the order item.
// returnOrderItemDetails - Return order item details.
// options - OrderItemsClientBeginReturnOptions contains the optional parameters for the OrderItemsClient.BeginReturn method.
func (client *OrderItemsClient) BeginReturn(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails ReturnOrderItemDetails, options *OrderItemsClientBeginReturnOptions) (*runtime.Poller[OrderItemsClientReturnResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.returnOperation(ctx, resourceGroupName, orderItemName, returnOrderItemDetails, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OrderItemsClientReturnResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OrderItemsClientReturnResponse](options.ResumeToken, client.pl, nil)
	}
}

// Return - Return order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *OrderItemsClient) returnOperation(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails ReturnOrderItemDetails, options *OrderItemsClientBeginReturnOptions) (*http.Response, error) {
	req, err := client.returnCreateRequest(ctx, resourceGroupName, orderItemName, returnOrderItemDetails, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// returnCreateRequest creates the Return request.
func (client *OrderItemsClient) returnCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails ReturnOrderItemDetails, options *OrderItemsClientBeginReturnOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}/return"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, returnOrderItemDetails)
}

// BeginUpdate - Update the properties of an existing order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// orderItemName - The name of the order item.
// orderItemUpdateParameter - Order item update parameters from request body.
// options - OrderItemsClientBeginUpdateOptions contains the optional parameters for the OrderItemsClient.BeginUpdate method.
func (client *OrderItemsClient) BeginUpdate(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter OrderItemUpdateParameter, options *OrderItemsClientBeginUpdateOptions) (*runtime.Poller[OrderItemsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, orderItemName, orderItemUpdateParameter, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OrderItemsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[OrderItemsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the properties of an existing order item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *OrderItemsClient) update(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter OrderItemUpdateParameter, options *OrderItemsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, orderItemName, orderItemUpdateParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OrderItemsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter OrderItemUpdateParameter, options *OrderItemsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, orderItemUpdateParameter)
}
