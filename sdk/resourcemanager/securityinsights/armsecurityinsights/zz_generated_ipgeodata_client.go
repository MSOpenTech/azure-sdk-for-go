//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// IPGeodataClient contains the methods for the IPGeodata group.
// Don't use this type directly, use NewIPGeodataClient() instead.
type IPGeodataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIPGeodataClient creates a new instance of IPGeodataClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIPGeodataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPGeodataClient, error) {
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
	client := &IPGeodataClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get geodata for a single IP address
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ipAddress - IP address (v4 or v6) to be enriched
// options - IPGeodataClientGetOptions contains the optional parameters for the IPGeodataClient.Get method.
func (client *IPGeodataClient) Get(ctx context.Context, resourceGroupName string, ipAddress string, options *IPGeodataClientGetOptions) (IPGeodataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipAddress, options)
	if err != nil {
		return IPGeodataClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IPGeodataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IPGeodataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IPGeodataClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipAddress string, options *IPGeodataClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityInsights/enrichment/ip/geodata/"
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
	reqQP.Set("api-version", "2022-04-01-preview")
	reqQP.Set("ipAddress", ipAddress)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPGeodataClient) getHandleResponse(resp *http.Response) (IPGeodataClientGetResponse, error) {
	result := IPGeodataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnrichmentIPGeodata); err != nil {
		return IPGeodataClientGetResponse{}, err
	}
	return result, nil
}
