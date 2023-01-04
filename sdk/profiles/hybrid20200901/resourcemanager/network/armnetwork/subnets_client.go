//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// SubnetsClient contains the methods for the Subnets group.
// Don't use this type directly, use NewSubnetsClient() instead.
type SubnetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSubnetsClient creates a new instance of SubnetsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubnetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubnetsClient, error) {
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
	client := &SubnetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a subnet in the specified virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// subnetName - The name of the subnet.
// subnetParameters - Parameters supplied to the create or update subnet operation.
// options - SubnetsClientBeginCreateOrUpdateOptions contains the optional parameters for the SubnetsClient.BeginCreateOrUpdate
// method.
func (client *SubnetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet, options *SubnetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SubnetsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SubnetsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SubnetsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a subnet in the specified virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *SubnetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet, options *SubnetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetParameters, options)
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
func (client *SubnetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet, options *SubnetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, subnetParameters)
}

// BeginDelete - Deletes the specified subnet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// subnetName - The name of the subnet.
// options - SubnetsClientBeginDeleteOptions contains the optional parameters for the SubnetsClient.BeginDelete method.
func (client *SubnetsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsClientBeginDeleteOptions) (*runtime.Poller[SubnetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SubnetsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SubnetsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified subnet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *SubnetsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
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
func (client *SubnetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified subnet by virtual network and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// subnetName - The name of the subnet.
// options - SubnetsClientGetOptions contains the optional parameters for the SubnetsClient.Get method.
func (client *SubnetsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsClientGetOptions) (SubnetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return SubnetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubnetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubnetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubnetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *SubnetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubnetsClient) getHandleResponse(resp *http.Response) (SubnetsClientGetResponse, error) {
	result := SubnetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Subnet); err != nil {
		return SubnetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all subnets in a virtual network.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// options - SubnetsClientListOptions contains the optional parameters for the SubnetsClient.List method.
func (client *SubnetsClient) NewListPager(resourceGroupName string, virtualNetworkName string, options *SubnetsClientListOptions) *runtime.Pager[SubnetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubnetsClientListResponse]{
		More: func(page SubnetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubnetsClientListResponse) (SubnetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubnetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubnetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubnetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SubnetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *SubnetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubnetsClient) listHandleResponse(resp *http.Response) (SubnetsClientListResponse, error) {
	result := SubnetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubnetListResult); err != nil {
		return SubnetsClientListResponse{}, err
	}
	return result, nil
}
