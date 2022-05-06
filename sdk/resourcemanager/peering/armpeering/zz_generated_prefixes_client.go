//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// PrefixesClient contains the methods for the Prefixes group.
// Don't use this type directly, use NewPrefixesClient() instead.
type PrefixesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrefixesClient creates a new instance of PrefixesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrefixesClient, error) {
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
	client := &PrefixesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new prefix with the specified name under the given subscription, resource group and peering
// service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// prefixName - The name of the prefix.
// peeringServicePrefix - The properties needed to create a prefix.
// options - PrefixesClientCreateOrUpdateOptions contains the optional parameters for the PrefixesClient.CreateOrUpdate method.
func (client *PrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix ServicePrefix, options *PrefixesClientCreateOrUpdateOptions) (PrefixesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringServiceName, prefixName, peeringServicePrefix, options)
	if err != nil {
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrefixesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix ServicePrefix, options *PrefixesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if prefixName == "" {
		return nil, errors.New("parameter prefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{prefixName}", url.PathEscape(prefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, peeringServicePrefix)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrefixesClient) createOrUpdateHandleResponse(resp *http.Response) (PrefixesClientCreateOrUpdateResponse, error) {
	result := PrefixesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefix); err != nil {
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing prefix with the specified name under the given subscription, resource group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// prefixName - The name of the prefix.
// options - PrefixesClientDeleteOptions contains the optional parameters for the PrefixesClient.Delete method.
func (client *PrefixesClient) Delete(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientDeleteOptions) (PrefixesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringServiceName, prefixName, options)
	if err != nil {
		return PrefixesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrefixesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PrefixesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PrefixesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if prefixName == "" {
		return nil, errors.New("parameter prefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{prefixName}", url.PathEscape(prefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets an existing prefix with the specified name under the given subscription, resource group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// prefixName - The name of the prefix.
// options - PrefixesClientGetOptions contains the optional parameters for the PrefixesClient.Get method.
func (client *PrefixesClient) Get(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientGetOptions) (PrefixesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringServiceName, prefixName, options)
	if err != nil {
		return PrefixesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrefixesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrefixesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if prefixName == "" {
		return nil, errors.New("parameter prefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{prefixName}", url.PathEscape(prefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrefixesClient) getHandleResponse(resp *http.Response) (PrefixesClientGetResponse, error) {
	result := PrefixesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefix); err != nil {
		return PrefixesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPeeringServicePager - Lists all prefixes under the given subscription, resource group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// peeringServiceName - The name of the peering service.
// options - PrefixesClientListByPeeringServiceOptions contains the optional parameters for the PrefixesClient.ListByPeeringService
// method.
func (client *PrefixesClient) NewListByPeeringServicePager(resourceGroupName string, peeringServiceName string, options *PrefixesClientListByPeeringServiceOptions) *runtime.Pager[PrefixesClientListByPeeringServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrefixesClientListByPeeringServiceResponse]{
		More: func(page PrefixesClientListByPeeringServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrefixesClientListByPeeringServiceResponse) (PrefixesClientListByPeeringServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPeeringServiceCreateRequest(ctx, resourceGroupName, peeringServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrefixesClientListByPeeringServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrefixesClientListByPeeringServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrefixesClientListByPeeringServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPeeringServiceHandleResponse(resp)
		},
	})
}

// listByPeeringServiceCreateRequest creates the ListByPeeringService request.
func (client *PrefixesClient) listByPeeringServiceCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, options *PrefixesClientListByPeeringServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPeeringServiceHandleResponse handles the ListByPeeringService response.
func (client *PrefixesClient) listByPeeringServiceHandleResponse(resp *http.Response) (PrefixesClientListByPeeringServiceResponse, error) {
	result := PrefixesClientListByPeeringServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefixListResult); err != nil {
		return PrefixesClientListByPeeringServiceResponse{}, err
	}
	return result, nil
}
