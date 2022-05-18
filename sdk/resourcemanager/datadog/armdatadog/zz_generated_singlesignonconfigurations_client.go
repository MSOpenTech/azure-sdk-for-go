//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

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

// SingleSignOnConfigurationsClient contains the methods for the SingleSignOnConfigurations group.
// Don't use this type directly, use NewSingleSignOnConfigurationsClient() instead.
type SingleSignOnConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSingleSignOnConfigurationsClient creates a new instance of SingleSignOnConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSingleSignOnConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SingleSignOnConfigurationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SingleSignOnConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Configures single-sign-on for this resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// configurationName - Configuration name
// options - SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the SingleSignOnConfigurationsClient.BeginCreateOrUpdate
// method.
func (client *SingleSignOnConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SingleSignOnConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, monitorName, configurationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SingleSignOnConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SingleSignOnConfigurationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Configures single-sign-on for this resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-01
func (client *SingleSignOnConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
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
func (client *SingleSignOnConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// Get - Gets the datadog single sign-on resource for the given Monitor.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// configurationName - Configuration name
// options - SingleSignOnConfigurationsClientGetOptions contains the optional parameters for the SingleSignOnConfigurationsClient.Get
// method.
func (client *SingleSignOnConfigurationsClient) Get(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnConfigurationsClientGetOptions) (SingleSignOnConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return SingleSignOnConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SingleSignOnConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SingleSignOnConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SingleSignOnConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SingleSignOnConfigurationsClient) getHandleResponse(resp *http.Response) (SingleSignOnConfigurationsClientGetResponse, error) {
	result := SingleSignOnConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SingleSignOnResource); err != nil {
		return SingleSignOnConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the single sign-on configurations for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - SingleSignOnConfigurationsClientListOptions contains the optional parameters for the SingleSignOnConfigurationsClient.List
// method.
func (client *SingleSignOnConfigurationsClient) NewListPager(resourceGroupName string, monitorName string, options *SingleSignOnConfigurationsClientListOptions) *runtime.Pager[SingleSignOnConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SingleSignOnConfigurationsClientListResponse]{
		More: func(page SingleSignOnConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SingleSignOnConfigurationsClientListResponse) (SingleSignOnConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SingleSignOnConfigurationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SingleSignOnConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SingleSignOnConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SingleSignOnConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *SingleSignOnConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/singleSignOnConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SingleSignOnConfigurationsClient) listHandleResponse(resp *http.Response) (SingleSignOnConfigurationsClientListResponse, error) {
	result := SingleSignOnConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SingleSignOnResourceListResponse); err != nil {
		return SingleSignOnConfigurationsClientListResponse{}, err
	}
	return result, nil
}
