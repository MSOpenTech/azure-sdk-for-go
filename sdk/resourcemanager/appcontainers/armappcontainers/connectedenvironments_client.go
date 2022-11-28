//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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
	"strings"
)

// ConnectedEnvironmentsClient contains the methods for the ConnectedEnvironments group.
// Don't use this type directly, use NewConnectedEnvironmentsClient() instead.
type ConnectedEnvironmentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectedEnvironmentsClient creates a new instance of ConnectedEnvironmentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConnectedEnvironmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedEnvironmentsClient, error) {
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
	client := &ConnectedEnvironmentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Checks if resource connectedEnvironmentName is available.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// connectedEnvironmentName - Name of the Managed Environment.
// checkNameAvailabilityRequest - The check connectedEnvironmentName availability request.
// options - ConnectedEnvironmentsClientCheckNameAvailabilityOptions contains the optional parameters for the ConnectedEnvironmentsClient.CheckNameAvailability
// method.
func (client *ConnectedEnvironmentsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *ConnectedEnvironmentsClientCheckNameAvailabilityOptions) (ConnectedEnvironmentsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, checkNameAvailabilityRequest, options)
	if err != nil {
		return ConnectedEnvironmentsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedEnvironmentsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedEnvironmentsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ConnectedEnvironmentsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *ConnectedEnvironmentsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ConnectedEnvironmentsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ConnectedEnvironmentsClientCheckNameAvailabilityResponse, error) {
	result := ConnectedEnvironmentsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ConnectedEnvironmentsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates an connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// connectedEnvironmentName - Name of the connectedEnvironment.
// environmentEnvelope - Configuration details of the connectedEnvironment.
// options - ConnectedEnvironmentsClientBeginCreateOrUpdateOptions contains the optional parameters for the ConnectedEnvironmentsClient.BeginCreateOrUpdate
// method.
func (client *ConnectedEnvironmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, environmentEnvelope ConnectedEnvironment, options *ConnectedEnvironmentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ConnectedEnvironmentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, connectedEnvironmentName, environmentEnvelope, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ConnectedEnvironmentsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedEnvironmentsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *ConnectedEnvironmentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, environmentEnvelope ConnectedEnvironment, options *ConnectedEnvironmentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, environmentEnvelope, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectedEnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, environmentEnvelope ConnectedEnvironment, options *ConnectedEnvironmentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, environmentEnvelope)
}

// BeginDelete - Delete an connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// connectedEnvironmentName - Name of the connectedEnvironment.
// options - ConnectedEnvironmentsClientBeginDeleteOptions contains the optional parameters for the ConnectedEnvironmentsClient.BeginDelete
// method.
func (client *ConnectedEnvironmentsClient) BeginDelete(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientBeginDeleteOptions) (*runtime.Poller[ConnectedEnvironmentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, connectedEnvironmentName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ConnectedEnvironmentsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedEnvironmentsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *ConnectedEnvironmentsClient) deleteOperation(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
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
func (client *ConnectedEnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of an connectedEnvironment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// connectedEnvironmentName - Name of the connectedEnvironment.
// options - ConnectedEnvironmentsClientGetOptions contains the optional parameters for the ConnectedEnvironmentsClient.Get
// method.
func (client *ConnectedEnvironmentsClient) Get(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientGetOptions) (ConnectedEnvironmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
	if err != nil {
		return ConnectedEnvironmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedEnvironmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedEnvironmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectedEnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedEnvironmentsClient) getHandleResponse(resp *http.Response) (ConnectedEnvironmentsClientGetResponse, error) {
	result := ConnectedEnvironmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironment); err != nil {
		return ConnectedEnvironmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all connectedEnvironments in a resource group.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ConnectedEnvironmentsClientListByResourceGroupOptions contains the optional parameters for the ConnectedEnvironmentsClient.ListByResourceGroup
// method.
func (client *ConnectedEnvironmentsClient) NewListByResourceGroupPager(resourceGroupName string, options *ConnectedEnvironmentsClientListByResourceGroupOptions) *runtime.Pager[ConnectedEnvironmentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedEnvironmentsClientListByResourceGroupResponse]{
		More: func(page ConnectedEnvironmentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedEnvironmentsClientListByResourceGroupResponse) (ConnectedEnvironmentsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectedEnvironmentsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectedEnvironmentsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedEnvironmentsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConnectedEnvironmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConnectedEnvironmentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments"
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
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConnectedEnvironmentsClient) listByResourceGroupHandleResponse(resp *http.Response) (ConnectedEnvironmentsClientListByResourceGroupResponse, error) {
	result := ConnectedEnvironmentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironmentCollection); err != nil {
		return ConnectedEnvironmentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get all connectedEnvironments for a subscription.
// Generated from API version 2022-06-01-preview
// options - ConnectedEnvironmentsClientListBySubscriptionOptions contains the optional parameters for the ConnectedEnvironmentsClient.ListBySubscription
// method.
func (client *ConnectedEnvironmentsClient) NewListBySubscriptionPager(options *ConnectedEnvironmentsClientListBySubscriptionOptions) *runtime.Pager[ConnectedEnvironmentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedEnvironmentsClientListBySubscriptionResponse]{
		More: func(page ConnectedEnvironmentsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedEnvironmentsClientListBySubscriptionResponse) (ConnectedEnvironmentsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectedEnvironmentsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectedEnvironmentsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedEnvironmentsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConnectedEnvironmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConnectedEnvironmentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.App/connectedEnvironments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConnectedEnvironmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (ConnectedEnvironmentsClientListBySubscriptionResponse, error) {
	result := ConnectedEnvironmentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironmentCollection); err != nil {
		return ConnectedEnvironmentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches a Managed Environment. Only patching of tags is supported currently
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// connectedEnvironmentName - Name of the connectedEnvironment.
// options - ConnectedEnvironmentsClientUpdateOptions contains the optional parameters for the ConnectedEnvironmentsClient.Update
// method.
func (client *ConnectedEnvironmentsClient) Update(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientUpdateOptions) (ConnectedEnvironmentsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
	if err != nil {
		return ConnectedEnvironmentsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedEnvironmentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedEnvironmentsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ConnectedEnvironmentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ConnectedEnvironmentsClient) updateHandleResponse(resp *http.Response) (ConnectedEnvironmentsClientUpdateResponse, error) {
	result := ConnectedEnvironmentsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedEnvironment); err != nil {
		return ConnectedEnvironmentsClientUpdateResponse{}, err
	}
	return result, nil
}
