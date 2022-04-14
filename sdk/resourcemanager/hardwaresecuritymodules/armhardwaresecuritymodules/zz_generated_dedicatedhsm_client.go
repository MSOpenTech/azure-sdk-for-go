//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

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

// DedicatedHsmClient contains the methods for the DedicatedHsm group.
// Don't use this type directly, use NewDedicatedHsmClient() instead.
type DedicatedHsmClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDedicatedHsmClient creates a new instance of DedicatedHsmClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDedicatedHsmClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DedicatedHsmClient, error) {
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
	client := &DedicatedHsmClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the resource belongs.
// name - Name of the dedicated Hsm
// parameters - Parameters to create or update the dedicated hsm
// options - DedicatedHsmClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedHsmClient.BeginCreateOrUpdate
// method.
func (client *DedicatedHsmClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm, options *DedicatedHsmClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DedicatedHsmClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, name, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DedicatedHsmClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DedicatedHsmClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedHsmClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm, options *DedicatedHsmClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, parameters, options)
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
func (client *DedicatedHsmClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm, options *DedicatedHsmClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Azure Dedicated HSM.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the dedicated HSM belongs.
// name - The name of the dedicated HSM to delete
// options - DedicatedHsmClientBeginDeleteOptions contains the optional parameters for the DedicatedHsmClient.BeginDelete
// method.
func (client *DedicatedHsmClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmClientBeginDeleteOptions) (*armruntime.Poller[DedicatedHsmClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DedicatedHsmClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DedicatedHsmClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified Azure Dedicated HSM.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedHsmClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *DedicatedHsmClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified Azure dedicated HSM.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the dedicated hsm belongs.
// name - The name of the dedicated HSM.
// options - DedicatedHsmClientGetOptions contains the optional parameters for the DedicatedHsmClient.Get method.
func (client *DedicatedHsmClient) Get(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmClientGetOptions) (DedicatedHsmClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return DedicatedHsmClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHsmClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHsmClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHsmClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHsmClient) getHandleResponse(resp *http.Response) (DedicatedHsmClientGetResponse, error) {
	result := DedicatedHsmClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHsm); err != nil {
		return DedicatedHsmClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - The List operation gets information about the dedicated hsms associated with the subscription and
// within the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the dedicated HSM belongs.
// options - DedicatedHsmClientListByResourceGroupOptions contains the optional parameters for the DedicatedHsmClient.ListByResourceGroup
// method.
func (client *DedicatedHsmClient) ListByResourceGroup(resourceGroupName string, options *DedicatedHsmClientListByResourceGroupOptions) *runtime.Pager[DedicatedHsmClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[DedicatedHsmClientListByResourceGroupResponse]{
		More: func(page DedicatedHsmClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHsmClientListByResourceGroupResponse) (DedicatedHsmClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedHsmClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedHsmClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedHsmClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedHsmClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DedicatedHsmClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedHsmClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedHsmClientListByResourceGroupResponse, error) {
	result := DedicatedHsmClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHsmListResult); err != nil {
		return DedicatedHsmClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - The List operation gets information about the dedicated HSMs associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DedicatedHsmClientListBySubscriptionOptions contains the optional parameters for the DedicatedHsmClient.ListBySubscription
// method.
func (client *DedicatedHsmClient) ListBySubscription(options *DedicatedHsmClientListBySubscriptionOptions) *runtime.Pager[DedicatedHsmClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[DedicatedHsmClientListBySubscriptionResponse]{
		More: func(page DedicatedHsmClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHsmClientListBySubscriptionResponse) (DedicatedHsmClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedHsmClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedHsmClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedHsmClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedHsmClient) listBySubscriptionCreateRequest(ctx context.Context, options *DedicatedHsmClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedHsmClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedHsmClientListBySubscriptionResponse, error) {
	result := DedicatedHsmClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHsmListResult); err != nil {
		return DedicatedHsmClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListOutboundNetworkDependenciesEndpoints - Gets a list of egress endpoints (network endpoints of all outbound dependencies)
// in the specified dedicated hsm resource. The operation returns properties of each egress endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the dedicated hsm belongs.
// name - The name of the dedicated HSM.
// options - DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DedicatedHsmClient.ListOutboundNetworkDependenciesEndpoints
// method.
func (client *DedicatedHsmClient) ListOutboundNetworkDependenciesEndpoints(resourceGroupName string, name string, options *DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions) *runtime.Pager[DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse] {
	return runtime.NewPager(runtime.PageProcessor[DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse]{
		More: func(page DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse) (DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOutboundNetworkDependenciesEndpointsCreateRequest(ctx, resourceGroupName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
		},
	})
}

// listOutboundNetworkDependenciesEndpointsCreateRequest creates the ListOutboundNetworkDependenciesEndpoints request.
func (client *DedicatedHsmClient) listOutboundNetworkDependenciesEndpointsCreateRequest(ctx context.Context, resourceGroupName string, name string, options *DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}/outboundNetworkDependenciesEndpoints"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOutboundNetworkDependenciesEndpointsHandleResponse handles the ListOutboundNetworkDependenciesEndpoints response.
func (client *DedicatedHsmClient) listOutboundNetworkDependenciesEndpointsHandleResponse(resp *http.Response) (DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse, error) {
	result := DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointCollection); err != nil {
		return DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Resource Group to which the server belongs.
// name - Name of the dedicated HSM
// parameters - Parameters to patch the dedicated HSM
// options - DedicatedHsmClientBeginUpdateOptions contains the optional parameters for the DedicatedHsmClient.BeginUpdate
// method.
func (client *DedicatedHsmClient) BeginUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters, options *DedicatedHsmClientBeginUpdateOptions) (*armruntime.Poller[DedicatedHsmClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, name, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DedicatedHsmClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DedicatedHsmClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a dedicated HSM in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedHsmClient) update(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters, options *DedicatedHsmClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHsmClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters, options *DedicatedHsmClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
