//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

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

// LivePipelinesClient contains the methods for the LivePipelines group.
// Don't use this type directly, use NewLivePipelinesClient() instead.
type LivePipelinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLivePipelinesClient creates a new instance of LivePipelinesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLivePipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LivePipelinesClient, error) {
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
	client := &LivePipelinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginActivate - Activates a live pipeline with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// options - LivePipelinesClientBeginActivateOptions contains the optional parameters for the LivePipelinesClient.BeginActivate
// method.
func (client *LivePipelinesClient) BeginActivate(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientBeginActivateOptions) (*armruntime.Poller[LivePipelinesClientActivateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.activate(ctx, resourceGroupName, accountName, livePipelineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[LivePipelinesClientActivateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[LivePipelinesClientActivateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Activate - Activates a live pipeline with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LivePipelinesClient) activate(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientBeginActivateOptions) (*http.Response, error) {
	req, err := client.activateCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// activateCreateRequest creates the Activate request.
func (client *LivePipelinesClient) activateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientBeginActivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}/activate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// CreateOrUpdate - Creates a new live pipeline or updates an existing one, with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// parameters - The request parameters
// options - LivePipelinesClientCreateOrUpdateOptions contains the optional parameters for the LivePipelinesClient.CreateOrUpdate
// method.
func (client *LivePipelinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, parameters LivePipeline, options *LivePipelinesClientCreateOrUpdateOptions) (LivePipelinesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, parameters, options)
	if err != nil {
		return LivePipelinesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LivePipelinesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return LivePipelinesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LivePipelinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, parameters LivePipeline, options *LivePipelinesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LivePipelinesClient) createOrUpdateHandleResponse(resp *http.Response) (LivePipelinesClientCreateOrUpdateResponse, error) {
	result := LivePipelinesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivePipeline); err != nil {
		return LivePipelinesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDeactivate - Deactivates a live pipeline with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// options - LivePipelinesClientBeginDeactivateOptions contains the optional parameters for the LivePipelinesClient.BeginDeactivate
// method.
func (client *LivePipelinesClient) BeginDeactivate(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientBeginDeactivateOptions) (*armruntime.Poller[LivePipelinesClientDeactivateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deactivate(ctx, resourceGroupName, accountName, livePipelineName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[LivePipelinesClientDeactivateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[LivePipelinesClientDeactivateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Deactivate - Deactivates a live pipeline with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LivePipelinesClient) deactivate(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientBeginDeactivateOptions) (*http.Response, error) {
	req, err := client.deactivateCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deactivateCreateRequest creates the Deactivate request.
func (client *LivePipelinesClient) deactivateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientBeginDeactivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}/deactivate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Delete - Deletes a live pipeline with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// options - LivePipelinesClientDeleteOptions contains the optional parameters for the LivePipelinesClient.Delete method.
func (client *LivePipelinesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientDeleteOptions) (LivePipelinesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, options)
	if err != nil {
		return LivePipelinesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LivePipelinesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return LivePipelinesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return LivePipelinesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LivePipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves a specific live pipeline by name. If a live pipeline with that name has been previously created, the call
// will return the JSON representation of that instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// options - LivePipelinesClientGetOptions contains the optional parameters for the LivePipelinesClient.Get method.
func (client *LivePipelinesClient) Get(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientGetOptions) (LivePipelinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, options)
	if err != nil {
		return LivePipelinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LivePipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LivePipelinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LivePipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, options *LivePipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LivePipelinesClient) getHandleResponse(resp *http.Response) (LivePipelinesClientGetResponse, error) {
	result := LivePipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivePipeline); err != nil {
		return LivePipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves a list of live pipelines that have been created, along with their JSON representations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// options - LivePipelinesClientListOptions contains the optional parameters for the LivePipelinesClient.List method.
func (client *LivePipelinesClient) NewListPager(resourceGroupName string, accountName string, options *LivePipelinesClientListOptions) *runtime.Pager[LivePipelinesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[LivePipelinesClientListResponse]{
		More: func(page LivePipelinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LivePipelinesClientListResponse) (LivePipelinesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LivePipelinesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LivePipelinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LivePipelinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LivePipelinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *LivePipelinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LivePipelinesClient) listHandleResponse(resp *http.Response) (LivePipelinesClientListResponse, error) {
	result := LivePipelinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivePipelineCollection); err != nil {
		return LivePipelinesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing live pipeline with the given name. Properties that can be updated include: description, bitrateKbps,
// and parameter definitions. Only the description can be updated while the live
// pipeline is active.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// livePipelineName - Live pipeline unique identifier.
// parameters - The request parameters
// options - LivePipelinesClientUpdateOptions contains the optional parameters for the LivePipelinesClient.Update method.
func (client *LivePipelinesClient) Update(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, parameters LivePipelineUpdate, options *LivePipelinesClientUpdateOptions) (LivePipelinesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, livePipelineName, parameters, options)
	if err != nil {
		return LivePipelinesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LivePipelinesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LivePipelinesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *LivePipelinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, parameters LivePipelineUpdate, options *LivePipelinesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if livePipelineName == "" {
		return nil, errors.New("parameter livePipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{livePipelineName}", url.PathEscape(livePipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *LivePipelinesClient) updateHandleResponse(resp *http.Response) (LivePipelinesClientUpdateResponse, error) {
	result := LivePipelinesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LivePipeline); err != nil {
		return LivePipelinesClientUpdateResponse{}, err
	}
	return result, nil
}
