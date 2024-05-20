//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// NamespaceTopicEventSubscriptionsClient contains the methods for the NamespaceTopicEventSubscriptions group.
// Don't use this type directly, use NewNamespaceTopicEventSubscriptionsClient() instead.
type NamespaceTopicEventSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNamespaceTopicEventSubscriptionsClient creates a new instance of NamespaceTopicEventSubscriptionsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNamespaceTopicEventSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NamespaceTopicEventSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NamespaceTopicEventSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates or updates an event subscription of a namespace topic with the specified parameters.
// Existing event subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 50
//     characters in length and use alphanumeric letters only.
//   - eventSubscriptionInfo - Event subscription properties containing the delivery mode, filter information, and others.
//   - options - NamespaceTopicEventSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.BeginCreateOrUpdate
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, eventSubscriptionInfo Subscription, options *NamespaceTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NamespaceTopicEventSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, eventSubscriptionInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NamespaceTopicEventSubscriptionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NamespaceTopicEventSubscriptionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Asynchronously creates or updates an event subscription of a namespace topic with the specified parameters.
// Existing event subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *NamespaceTopicEventSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, eventSubscriptionInfo Subscription, options *NamespaceTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NamespaceTopicEventSubscriptionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, eventSubscriptionInfo, options)
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
func (client *NamespaceTopicEventSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, eventSubscriptionInfo Subscription, options *NamespaceTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSubscriptionInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an existing event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - eventSubscriptionName - Name of the event subscription to be deleted.
//   - options - NamespaceTopicEventSubscriptionsClientBeginDeleteOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.BeginDelete
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[NamespaceTopicEventSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NamespaceTopicEventSubscriptionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NamespaceTopicEventSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete an existing event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *NamespaceTopicEventSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NamespaceTopicEventSubscriptionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, options)
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
func (client *NamespaceTopicEventSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of an event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - eventSubscriptionName - Name of the event subscription to be found.
//   - options - NamespaceTopicEventSubscriptionsClientGetOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.Get
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientGetOptions) (NamespaceTopicEventSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "NamespaceTopicEventSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, options)
	if err != nil {
		return NamespaceTopicEventSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamespaceTopicEventSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NamespaceTopicEventSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NamespaceTopicEventSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NamespaceTopicEventSubscriptionsClient) getHandleResponse(resp *http.Response) (NamespaceTopicEventSubscriptionsClientGetResponse, error) {
	result := NamespaceTopicEventSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Subscription); err != nil {
		return NamespaceTopicEventSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeliveryAttributes - Get all delivery attributes for an event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - eventSubscriptionName - Name of the event subscription.
//   - options - NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.GetDeliveryAttributes
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	var err error
	const operationName = "NamespaceTopicEventSubscriptionsClient.GetDeliveryAttributes"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDeliveryAttributesCreateRequest(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, options)
	if err != nil {
		return NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	resp, err := client.getDeliveryAttributesHandleResponse(httpResp)
	return resp, err
}

// getDeliveryAttributesCreateRequest creates the GetDeliveryAttributes request.
func (client *NamespaceTopicEventSubscriptionsClient) getDeliveryAttributesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeliveryAttributesHandleResponse handles the GetDeliveryAttributes response.
func (client *NamespaceTopicEventSubscriptionsClient) getDeliveryAttributesHandleResponse(resp *http.Response) (NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	result := NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeliveryAttributeListResult); err != nil {
		return NamespaceTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	return result, nil
}

// GetFullURL - Get the full endpoint URL for an event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - eventSubscriptionName - Name of the event subscription.
//   - options - NamespaceTopicEventSubscriptionsClientGetFullURLOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.GetFullURL
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientGetFullURLOptions) (NamespaceTopicEventSubscriptionsClientGetFullURLResponse, error) {
	var err error
	const operationName = "NamespaceTopicEventSubscriptionsClient.GetFullURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, options)
	if err != nil {
		return NamespaceTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamespaceTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NamespaceTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	resp, err := client.getFullURLHandleResponse(httpResp)
	return resp, err
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *NamespaceTopicEventSubscriptionsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, options *NamespaceTopicEventSubscriptionsClientGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *NamespaceTopicEventSubscriptionsClient) getFullURLHandleResponse(resp *http.Response) (NamespaceTopicEventSubscriptionsClientGetFullURLResponse, error) {
	result := NamespaceTopicEventSubscriptionsClientGetFullURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionFullURL); err != nil {
		return NamespaceTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	return result, nil
}

// NewListByNamespaceTopicPager - List event subscriptions that belong to a specific namespace topic.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - options - NamespaceTopicEventSubscriptionsClientListByNamespaceTopicOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.NewListByNamespaceTopicPager
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) NewListByNamespaceTopicPager(resourceGroupName string, namespaceName string, topicName string, options *NamespaceTopicEventSubscriptionsClientListByNamespaceTopicOptions) *runtime.Pager[NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse] {
	return runtime.NewPager(runtime.PagingHandler[NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse]{
		More: func(page NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse) (NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NamespaceTopicEventSubscriptionsClient.NewListByNamespaceTopicPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNamespaceTopicCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
			}, nil)
			if err != nil {
				return NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse{}, err
			}
			return client.listByNamespaceTopicHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNamespaceTopicCreateRequest creates the ListByNamespaceTopic request.
func (client *NamespaceTopicEventSubscriptionsClient) listByNamespaceTopicCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *NamespaceTopicEventSubscriptionsClientListByNamespaceTopicOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNamespaceTopicHandleResponse handles the ListByNamespaceTopic response.
func (client *NamespaceTopicEventSubscriptionsClient) listByNamespaceTopicHandleResponse(resp *http.Response) (NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse, error) {
	result := NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionsListResult); err != nil {
		return NamespaceTopicEventSubscriptionsClientListByNamespaceTopicResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an existing event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicName - Name of the namespace topic.
//   - eventSubscriptionName - Name of the event subscription to be updated.
//   - eventSubscriptionUpdateParameters - Updated event subscription information.
//   - options - NamespaceTopicEventSubscriptionsClientBeginUpdateOptions contains the optional parameters for the NamespaceTopicEventSubscriptionsClient.BeginUpdate
//     method.
func (client *NamespaceTopicEventSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters SubscriptionUpdateParameters, options *NamespaceTopicEventSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[NamespaceTopicEventSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NamespaceTopicEventSubscriptionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NamespaceTopicEventSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an existing event subscription of a namespace topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *NamespaceTopicEventSubscriptionsClient) update(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters SubscriptionUpdateParameters, options *NamespaceTopicEventSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NamespaceTopicEventSubscriptionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, namespaceName, topicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *NamespaceTopicEventSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters SubscriptionUpdateParameters, options *NamespaceTopicEventSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topics/{topicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSubscriptionUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
