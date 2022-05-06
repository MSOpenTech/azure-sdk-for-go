//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// SystemTopicEventSubscriptionsClient contains the methods for the SystemTopicEventSubscriptions group.
// Don't use this type directly, use NewSystemTopicEventSubscriptionsClient() instead.
type SystemTopicEventSubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSystemTopicEventSubscriptionsClient creates a new instance of SystemTopicEventSubscriptionsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSystemTopicEventSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SystemTopicEventSubscriptionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SystemTopicEventSubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates or updates an event subscription with the specified parameters. Existing event
// subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// eventSubscriptionInfo - Event subscription properties containing the destination and filter information.
// options - SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.BeginCreateOrUpdate
// method.
func (client *SystemTopicEventSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[SystemTopicEventSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SystemTopicEventSubscriptionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SystemTopicEventSubscriptionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Asynchronously creates or updates an event subscription with the specified parameters. Existing event
// subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SystemTopicEventSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SystemTopicEventSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventSubscriptionInfo)
}

// BeginDelete - Delete an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// options - SystemTopicEventSubscriptionsClientBeginDeleteOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.BeginDelete
// method.
func (client *SystemTopicEventSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientBeginDeleteOptions) (*armruntime.Poller[SystemTopicEventSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SystemTopicEventSubscriptionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SystemTopicEventSubscriptionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SystemTopicEventSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
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
func (client *SystemTopicEventSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get an event subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// options - SystemTopicEventSubscriptionsClientGetOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.Get
// method.
func (client *SystemTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetOptions) (SystemTopicEventSubscriptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SystemTopicEventSubscriptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SystemTopicEventSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemTopicEventSubscriptionsClient) getHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientGetResponse, error) {
	result := SystemTopicEventSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscription); err != nil {
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeliveryAttributes - Get all delivery attributes for an event subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// options - SystemTopicEventSubscriptionsClientGetDeliveryAttributesOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.GetDeliveryAttributes
// method.
func (client *SystemTopicEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	req, err := client.getDeliveryAttributesCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeliveryAttributesHandleResponse(resp)
}

// getDeliveryAttributesCreateRequest creates the GetDeliveryAttributes request.
func (client *SystemTopicEventSubscriptionsClient) getDeliveryAttributesCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDeliveryAttributesHandleResponse handles the GetDeliveryAttributes response.
func (client *SystemTopicEventSubscriptionsClient) getDeliveryAttributesHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	result := SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeliveryAttributeListResult); err != nil {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	return result, nil
}

// GetFullURL - Get the full endpoint URL for an event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// options - SystemTopicEventSubscriptionsClientGetFullURLOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.GetFullURL
// method.
func (client *SystemTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetFullURLOptions) (SystemTopicEventSubscriptionsClientGetFullURLResponse, error) {
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFullURLHandleResponse(resp)
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *SystemTopicEventSubscriptionsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *SystemTopicEventSubscriptionsClient) getFullURLHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientGetFullURLResponse, error) {
	result := SystemTopicEventSubscriptionsClientGetFullURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionFullURL); err != nil {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	return result, nil
}

// NewListBySystemTopicPager - List event subscriptions that belong to a specific system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// options - SystemTopicEventSubscriptionsClientListBySystemTopicOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.ListBySystemTopic
// method.
func (client *SystemTopicEventSubscriptionsClient) NewListBySystemTopicPager(resourceGroupName string, systemTopicName string, options *SystemTopicEventSubscriptionsClientListBySystemTopicOptions) *runtime.Pager[SystemTopicEventSubscriptionsClientListBySystemTopicResponse] {
	return runtime.NewPager(runtime.PageProcessor[SystemTopicEventSubscriptionsClientListBySystemTopicResponse]{
		More: func(page SystemTopicEventSubscriptionsClientListBySystemTopicResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemTopicEventSubscriptionsClientListBySystemTopicResponse) (SystemTopicEventSubscriptionsClientListBySystemTopicResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySystemTopicCreateRequest(ctx, resourceGroupName, systemTopicName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySystemTopicHandleResponse(resp)
		},
	})
}

// listBySystemTopicCreateRequest creates the ListBySystemTopic request.
func (client *SystemTopicEventSubscriptionsClient) listBySystemTopicCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicEventSubscriptionsClientListBySystemTopicOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySystemTopicHandleResponse handles the ListBySystemTopic response.
func (client *SystemTopicEventSubscriptionsClient) listBySystemTopicHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientListBySystemTopicResponse, error) {
	result := SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionsListResult); err != nil {
		return SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// systemTopicName - Name of the system topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// eventSubscriptionUpdateParameters - Updated event subscription information.
// options - SystemTopicEventSubscriptionsClientBeginUpdateOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.BeginUpdate
// method.
func (client *SystemTopicEventSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *SystemTopicEventSubscriptionsClientBeginUpdateOptions) (*armruntime.Poller[SystemTopicEventSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SystemTopicEventSubscriptionsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SystemTopicEventSubscriptionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SystemTopicEventSubscriptionsClient) update(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *SystemTopicEventSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SystemTopicEventSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *SystemTopicEventSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventSubscriptionUpdateParameters)
}
