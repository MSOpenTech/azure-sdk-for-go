//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// ResourcePoolsClient contains the methods for the ResourcePools group.
// Don't use this type directly, use NewResourcePoolsClient() instead.
type ResourcePoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourcePoolsClient creates a new instance of ResourcePoolsClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourcePoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourcePoolsClient, error) {
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
	client := &ResourcePoolsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns resource pool templates by its name
// If the operation fails it returns an *azcore.ResponseError type.
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// resourcePoolName - resource pool id (vsphereId)
// options - ResourcePoolsClientGetOptions contains the optional parameters for the ResourcePoolsClient.Get method.
func (client *ResourcePoolsClient) Get(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *ResourcePoolsClientGetOptions) (ResourcePoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regionID, pcName, resourcePoolName, options)
	if err != nil {
		return ResourcePoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourcePoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourcePoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourcePoolsClient) getCreateRequest(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *ResourcePoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if resourcePoolName == "" {
		return nil, errors.New("parameter resourcePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourcePoolName}", url.PathEscape(resourcePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourcePoolsClient) getHandleResponse(resp *http.Response) (ResourcePoolsClientGetResponse, error) {
	result := ResourcePoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcePool); err != nil {
		return ResourcePoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns list of resource pools in region for private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// options - ResourcePoolsClientListOptions contains the optional parameters for the ResourcePoolsClient.List method.
func (client *ResourcePoolsClient) List(regionID string, pcName string, options *ResourcePoolsClientListOptions) *runtime.Pager[ResourcePoolsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourcePoolsClientListResponse]{
		More: func(page ResourcePoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourcePoolsClientListResponse) (ResourcePoolsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, regionID, pcName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourcePoolsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourcePoolsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourcePoolsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ResourcePoolsClient) listCreateRequest(ctx context.Context, regionID string, pcName string, options *ResourcePoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourcePoolsClient) listHandleResponse(resp *http.Response) (ResourcePoolsClientListResponse, error) {
	result := ResourcePoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcePoolsListResponse); err != nil {
		return ResourcePoolsClientListResponse{}, err
	}
	return result, nil
}
