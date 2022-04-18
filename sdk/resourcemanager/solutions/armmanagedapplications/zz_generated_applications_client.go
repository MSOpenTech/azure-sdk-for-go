//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

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

// ApplicationsClient contains the methods for the Applications group.
// Don't use this type directly, use NewApplicationsClient() instead.
type ApplicationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationsClient creates a new instance of ApplicationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationsClient, error) {
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
	client := &ApplicationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new managed application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationName - The name of the managed application.
// parameters - Parameters supplied to the create or update a managed application.
// options - ApplicationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationsClient.BeginCreateOrUpdate
// method.
func (client *ApplicationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters Application, options *ApplicationsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ApplicationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ApplicationsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ApplicationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a new managed application.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ApplicationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters Application, options *ApplicationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationName, parameters, options)
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
func (client *ApplicationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, parameters Application, options *ApplicationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the managed application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationName - The name of the managed application.
// options - ApplicationsClientBeginDeleteOptions contains the optional parameters for the ApplicationsClient.BeginDelete
// method.
func (client *ApplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientBeginDeleteOptions) (*armruntime.Poller[ApplicationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, applicationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ApplicationsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ApplicationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the managed application.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ApplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the managed application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationName - The name of the managed application.
// options - ApplicationsClientGetOptions contains the optional parameters for the ApplicationsClient.Get method.
func (client *ApplicationsClient) Get(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientGetOptions) (ApplicationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return ApplicationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationsClient) getHandleResponse(resp *http.Response) (ApplicationsClientGetResponse, error) {
	result := ApplicationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ApplicationsClientGetResponse{}, err
	}
	return result, nil
}

// ListAllowedUpgradePlans - List allowed upgrade plans for application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationName - The name of the managed application.
// options - ApplicationsClientListAllowedUpgradePlansOptions contains the optional parameters for the ApplicationsClient.ListAllowedUpgradePlans
// method.
func (client *ApplicationsClient) ListAllowedUpgradePlans(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientListAllowedUpgradePlansOptions) (ApplicationsClientListAllowedUpgradePlansResponse, error) {
	req, err := client.listAllowedUpgradePlansCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return ApplicationsClientListAllowedUpgradePlansResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationsClientListAllowedUpgradePlansResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationsClientListAllowedUpgradePlansResponse{}, runtime.NewResponseError(resp)
	}
	return ApplicationsClientListAllowedUpgradePlansResponse{}, nil
}

// listAllowedUpgradePlansCreateRequest creates the ListAllowedUpgradePlans request.
func (client *ApplicationsClient) listAllowedUpgradePlansCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientListAllowedUpgradePlansOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}/listAllowedUpgradePlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// NewListByResourceGroupPager - Gets all the applications within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ApplicationsClientListByResourceGroupOptions contains the optional parameters for the ApplicationsClient.ListByResourceGroup
// method.
func (client *ApplicationsClient) NewListByResourceGroupPager(resourceGroupName string, options *ApplicationsClientListByResourceGroupOptions) *runtime.Pager[ApplicationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ApplicationsClientListByResourceGroupResponse]{
		More: func(page ApplicationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationsClientListByResourceGroupResponse) (ApplicationsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications"
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
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationsClientListByResourceGroupResponse, error) {
	result := ApplicationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationListResult); err != nil {
		return ApplicationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all the applications within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ApplicationsClientListBySubscriptionOptions contains the optional parameters for the ApplicationsClient.ListBySubscription
// method.
func (client *ApplicationsClient) NewListBySubscriptionPager(options *ApplicationsClientListBySubscriptionOptions) *runtime.Pager[ApplicationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[ApplicationsClientListBySubscriptionResponse]{
		More: func(page ApplicationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationsClientListBySubscriptionResponse) (ApplicationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/applications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationsClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationsClientListBySubscriptionResponse, error) {
	result := ApplicationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationListResult); err != nil {
		return ApplicationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRefreshPermissions - Refresh Permissions for application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationName - The name of the managed application.
// options - ApplicationsClientBeginRefreshPermissionsOptions contains the optional parameters for the ApplicationsClient.BeginRefreshPermissions
// method.
func (client *ApplicationsClient) BeginRefreshPermissions(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientBeginRefreshPermissionsOptions) (*armruntime.Poller[ApplicationsClientRefreshPermissionsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refreshPermissions(ctx, resourceGroupName, applicationName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ApplicationsClientRefreshPermissionsResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ApplicationsClientRefreshPermissionsResponse](options.ResumeToken, client.pl, nil)
	}
}

// RefreshPermissions - Refresh Permissions for application.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ApplicationsClient) refreshPermissions(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientBeginRefreshPermissionsOptions) (*http.Response, error) {
	req, err := client.refreshPermissionsCreateRequest(ctx, resourceGroupName, applicationName, options)
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

// refreshPermissionsCreateRequest creates the RefreshPermissions request.
func (client *ApplicationsClient) refreshPermissionsCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientBeginRefreshPermissionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}/refreshPermissions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Update - Updates an existing managed application. The only value that can be updated via PATCH currently is the tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationName - The name of the managed application.
// options - ApplicationsClientUpdateOptions contains the optional parameters for the ApplicationsClient.Update method.
func (client *ApplicationsClient) Update(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientUpdateOptions) (ApplicationsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return ApplicationsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return ApplicationsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *ApplicationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ApplicationsClient) updateHandleResponse(resp *http.Response) (ApplicationsClientUpdateResponse, error) {
	result := ApplicationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ApplicationsClientUpdateResponse{}, err
	}
	return result, nil
}
