//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

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

// ConfigurationClient contains the methods for the Configuration group.
// Don't use this type directly, use NewConfigurationClient() instead.
type ConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationClient creates a new instance of ConfigurationClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationClient, error) {
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
	client := &ConfigurationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get all Event Hubs Cluster settings - a collection of key/value pairs which represent the quotas and settings imposed
// on the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// options - ConfigurationClientGetOptions contains the optional parameters for the ConfigurationClient.Get method.
func (client *ConfigurationClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ConfigurationClientGetOptions) (ConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ConfigurationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/quotaConfiguration/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationClient) getHandleResponse(resp *http.Response) (ConfigurationClientGetResponse, error) {
	result := ConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterQuotaConfigurationProperties); err != nil {
		return ConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// Patch - Replace all specified Event Hubs Cluster settings with those contained in the request body. Leaves the settings
// not specified in the request body unmodified.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group within the azure subscription.
// clusterName - The name of the Event Hubs Cluster.
// parameters - Parameters for creating an Event Hubs Cluster resource.
// options - ConfigurationClientPatchOptions contains the optional parameters for the ConfigurationClient.Patch method.
func (client *ConfigurationClient) Patch(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterQuotaConfigurationProperties, options *ConfigurationClientPatchOptions) (ConfigurationClientPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return ConfigurationClientPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return ConfigurationClientPatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *ConfigurationClient) patchCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterQuotaConfigurationProperties, options *ConfigurationClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/quotaConfiguration/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// patchHandleResponse handles the Patch response.
func (client *ConfigurationClient) patchHandleResponse(resp *http.Response) (ConfigurationClientPatchResponse, error) {
	result := ConfigurationClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterQuotaConfigurationProperties); err != nil {
		return ConfigurationClientPatchResponse{}, err
	}
	return result, nil
}
