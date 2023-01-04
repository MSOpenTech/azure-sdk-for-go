//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/profiles/hybrid20200901/internal"
	"net/http"
	"net/url"
	"strings"
)

// ContainerServicesClient contains the methods for the ContainerServices group.
// Don't use this type directly, use NewContainerServicesClient() instead.
type ContainerServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainerServicesClient creates a new instance of ContainerServicesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContainerServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerServicesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(internal.ModuleName, internal.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a container service with the specified configuration of orchestrator, masters,
// and agents.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-07-01
// resourceGroupName - The name of the resource group.
// containerServiceName - The name of the container service in the specified subscription and resource group.
// parameters - Parameters supplied to the Create or Update a Container Service operation.
// options - ContainerServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerServicesClient.BeginCreateOrUpdate
// method.
func (client *ContainerServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainerServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, containerServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerServicesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerServicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a container service with the specified configuration of orchestrator, masters, and
// agents.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-07-01
func (client *ContainerServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContainerServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerServiceName == "" {
		return nil, errors.New("parameter containerServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified container service in the specified subscription and resource group. The operation does
// not delete other resources created as part of creating a container service, including
// storage accounts, VMs, and availability sets. All the other resources created with the container service are part of the
// same resource group and can be deleted individually.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-07-01
// resourceGroupName - The name of the resource group.
// containerServiceName - The name of the container service in the specified subscription and resource group.
// options - ContainerServicesClientBeginDeleteOptions contains the optional parameters for the ContainerServicesClient.BeginDelete
// method.
func (client *ContainerServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesClientBeginDeleteOptions) (*runtime.Poller[ContainerServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerServicesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified container service in the specified subscription and resource group. The operation does not
// delete other resources created as part of creating a container service, including
// storage accounts, VMs, and availability sets. All the other resources created with the container service are part of the
// same resource group and can be deleted individually.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-07-01
func (client *ContainerServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainerServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerServiceName == "" {
		return nil, errors.New("parameter containerServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the properties of the specified container service in the specified subscription and resource group. The operation
// returns the properties including state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-07-01
// resourceGroupName - The name of the resource group.
// containerServiceName - The name of the container service in the specified subscription and resource group.
// options - ContainerServicesClientGetOptions contains the optional parameters for the ContainerServicesClient.Get method.
func (client *ContainerServicesClient) Get(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesClientGetOptions) (ContainerServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return ContainerServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContainerServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerServiceName == "" {
		return nil, errors.New("parameter containerServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerServicesClient) getHandleResponse(resp *http.Response) (ContainerServicesClientGetResponse, error) {
	result := ContainerServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerService); err != nil {
		return ContainerServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of container services in the specified subscription. The operation returns properties of each
// container service including state, orchestrator, number of masters and agents, and FQDNs of
// masters and agents.
// Generated from API version 2017-07-01
// options - ContainerServicesClientListOptions contains the optional parameters for the ContainerServicesClient.List method.
func (client *ContainerServicesClient) NewListPager(options *ContainerServicesClientListOptions) *runtime.Pager[ContainerServicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerServicesClientListResponse]{
		More: func(page ContainerServicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerServicesClientListResponse) (ContainerServicesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerServicesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainerServicesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerServicesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ContainerServicesClient) listCreateRequest(ctx context.Context, options *ContainerServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContainerServicesClient) listHandleResponse(resp *http.Response) (ContainerServicesClientListResponse, error) {
	result := ContainerServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ContainerServicesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of container services in the specified subscription and resource group. The operation
// returns properties of each container service including state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
// Generated from API version 2017-07-01
// resourceGroupName - The name of the resource group.
// options - ContainerServicesClientListByResourceGroupOptions contains the optional parameters for the ContainerServicesClient.ListByResourceGroup
// method.
func (client *ContainerServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *ContainerServicesClientListByResourceGroupOptions) *runtime.Pager[ContainerServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerServicesClientListByResourceGroupResponse]{
		More: func(page ContainerServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerServicesClientListByResourceGroupResponse) (ContainerServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainerServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices"
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
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ContainerServicesClientListByResourceGroupResponse, error) {
	result := ContainerServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ContainerServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListOrchestrators - Gets a list of supported orchestrators in the specified subscription. The operation returns properties
// of each orchestrator including version, available upgrades and whether that version or upgrades
// are in preview.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// location - The name of a supported Azure region.
// options - ContainerServicesClientListOrchestratorsOptions contains the optional parameters for the ContainerServicesClient.ListOrchestrators
// method.
func (client *ContainerServicesClient) ListOrchestrators(ctx context.Context, location string, options *ContainerServicesClientListOrchestratorsOptions) (ContainerServicesClientListOrchestratorsResponse, error) {
	req, err := client.listOrchestratorsCreateRequest(ctx, location, options)
	if err != nil {
		return ContainerServicesClientListOrchestratorsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerServicesClientListOrchestratorsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerServicesClientListOrchestratorsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listOrchestratorsHandleResponse(resp)
}

// listOrchestratorsCreateRequest creates the ListOrchestrators request.
func (client *ContainerServicesClient) listOrchestratorsCreateRequest(ctx context.Context, location string, options *ContainerServicesClientListOrchestratorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	if options != nil && options.ResourceType != nil {
		reqQP.Set("resource-type", *options.ResourceType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOrchestratorsHandleResponse handles the ListOrchestrators response.
func (client *ContainerServicesClient) listOrchestratorsHandleResponse(resp *http.Response) (ContainerServicesClientListOrchestratorsResponse, error) {
	result := ContainerServicesClientListOrchestratorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrchestratorVersionProfileListResult); err != nil {
		return ContainerServicesClientListOrchestratorsResponse{}, err
	}
	return result, nil
}
