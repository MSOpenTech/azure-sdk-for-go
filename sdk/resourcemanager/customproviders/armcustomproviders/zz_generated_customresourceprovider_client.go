//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomproviders

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CustomResourceProviderClient contains the methods for the CustomResourceProvider group.
// Don't use this type directly, use NewCustomResourceProviderClient() instead.
type CustomResourceProviderClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCustomResourceProviderClient creates a new instance of CustomResourceProviderClient with the specified values.
func NewCustomResourceProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CustomResourceProviderClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CustomResourceProviderClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates the custom resource provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest, options *CustomResourceProviderBeginCreateOrUpdateOptions) (CustomResourceProviderCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceProviderName, resourceProvider, options)
	if err != nil {
		return CustomResourceProviderCreateOrUpdatePollerResponse{}, err
	}
	result := CustomResourceProviderCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomResourceProviderClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return CustomResourceProviderCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CustomResourceProviderCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the custom resource provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest, options *CustomResourceProviderBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceProviderName, resourceProvider, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomResourceProviderClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest, options *CustomResourceProviderBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resourceProvider)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CustomResourceProviderClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the custom resource provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderBeginDeleteOptions) (CustomResourceProviderDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, resourceProviderName, options)
	if err != nil {
		return CustomResourceProviderDeletePollerResponse{}, err
	}
	result := CustomResourceProviderDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomResourceProviderClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CustomResourceProviderDeletePollerResponse{}, err
	}
	result.Poller = &CustomResourceProviderDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the custom resource provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceProviderName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomResourceProviderClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomResourceProviderClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the custom resource provider manifest.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) Get(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderGetOptions) (CustomResourceProviderGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceProviderName, options)
	if err != nil {
		return CustomResourceProviderGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomResourceProviderGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomResourceProviderGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomResourceProviderClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, options *CustomResourceProviderGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomResourceProviderClient) getHandleResponse(resp *http.Response) (CustomResourceProviderGetResponse, error) {
	result := CustomResourceProviderGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRPManifest); err != nil {
		return CustomResourceProviderGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CustomResourceProviderClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets all the custom resource providers within a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) ListByResourceGroup(resourceGroupName string, options *CustomResourceProviderListByResourceGroupOptions) *CustomResourceProviderListByResourceGroupPager {
	return &CustomResourceProviderListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomResourceProviderListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListByCustomRPManifest.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomResourceProviderClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomResourceProviderListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomResourceProviderClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomResourceProviderListByResourceGroupResponse, error) {
	result := CustomResourceProviderListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListByCustomRPManifest); err != nil {
		return CustomResourceProviderListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CustomResourceProviderClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Gets all the custom resource providers within a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) ListBySubscription(options *CustomResourceProviderListBySubscriptionOptions) *CustomResourceProviderListBySubscriptionPager {
	return &CustomResourceProviderListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomResourceProviderListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListByCustomRPManifest.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomResourceProviderClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomResourceProviderListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CustomProviders/resourceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomResourceProviderClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomResourceProviderListBySubscriptionResponse, error) {
	result := CustomResourceProviderListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListByCustomRPManifest); err != nil {
		return CustomResourceProviderListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *CustomResourceProviderClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates an existing custom resource provider. The only value that can be updated via PATCH currently is the tags.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CustomResourceProviderClient) Update(ctx context.Context, resourceGroupName string, resourceProviderName string, patchableResource ResourceProvidersUpdate, options *CustomResourceProviderUpdateOptions) (CustomResourceProviderUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceProviderName, patchableResource, options)
	if err != nil {
		return CustomResourceProviderUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomResourceProviderUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomResourceProviderUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CustomResourceProviderClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderName string, patchableResource ResourceProvidersUpdate, options *CustomResourceProviderUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, patchableResource)
}

// updateHandleResponse handles the Update response.
func (client *CustomResourceProviderClient) updateHandleResponse(resp *http.Response) (CustomResourceProviderUpdateResponse, error) {
	result := CustomResourceProviderUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRPManifest); err != nil {
		return CustomResourceProviderUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *CustomResourceProviderClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
