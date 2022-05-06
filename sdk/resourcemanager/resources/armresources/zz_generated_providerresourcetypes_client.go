//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

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

// ProviderResourceTypesClient contains the methods for the ProviderResourceTypes group.
// Don't use this type directly, use NewProviderResourceTypesClient() instead.
type ProviderResourceTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProviderResourceTypesClient creates a new instance of ProviderResourceTypesClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProviderResourceTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderResourceTypesClient, error) {
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
	client := &ProviderResourceTypesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - List the resource types for a specified resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceProviderNamespace - The namespace of the resource provider.
// options - ProviderResourceTypesClientListOptions contains the optional parameters for the ProviderResourceTypesClient.List
// method.
func (client *ProviderResourceTypesClient) List(ctx context.Context, resourceProviderNamespace string, options *ProviderResourceTypesClientListOptions) (ProviderResourceTypesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProviderResourceTypesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderResourceTypesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderResourceTypesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ProviderResourceTypesClient) listCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProviderResourceTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/resourceTypes"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
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
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProviderResourceTypesClient) listHandleResponse(resp *http.Response) (ProviderResourceTypesClientListResponse, error) {
	result := ProviderResourceTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderResourceTypeListResult); err != nil {
		return ProviderResourceTypesClientListResponse{}, err
	}
	return result, nil
}
