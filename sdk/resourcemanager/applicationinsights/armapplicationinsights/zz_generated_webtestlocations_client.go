//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// WebTestLocationsClient contains the methods for the WebTestLocations group.
// Don't use this type directly, use NewWebTestLocationsClient() instead.
type WebTestLocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWebTestLocationsClient creates a new instance of WebTestLocationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWebTestLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebTestLocationsClient, error) {
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
	client := &WebTestLocationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Gets a list of web test locations available to this Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - WebTestLocationsClientListOptions contains the optional parameters for the WebTestLocationsClient.List method.
func (client *WebTestLocationsClient) List(resourceGroupName string, resourceName string, options *WebTestLocationsClientListOptions) *runtime.Pager[WebTestLocationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[WebTestLocationsClientListResponse]{
		More: func(page WebTestLocationsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WebTestLocationsClientListResponse) (WebTestLocationsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			if err != nil {
				return WebTestLocationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebTestLocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebTestLocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WebTestLocationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *WebTestLocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/syntheticmonitorlocations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebTestLocationsClient) listHandleResponse(resp *http.Response) (WebTestLocationsClientListResponse, error) {
	result := WebTestLocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestLocationsListResult); err != nil {
		return WebTestLocationsClientListResponse{}, err
	}
	return result, nil
}
