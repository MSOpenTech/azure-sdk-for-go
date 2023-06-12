//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PipelineRunsClient contains the methods for the PipelineRuns group.
// Don't use this type directly, use NewPipelineRunsClient() instead.
type PipelineRunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPipelineRunsClient creates a new instance of PipelineRunsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelineRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelineRunsClient, error) {
	cl, err := arm.NewClient(moduleName+".PipelineRunsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PipelineRunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - pipelineRunName - The name of the pipeline run.
//   - pipelineRunCreateParameters - The parameters for creating a pipeline run.
//   - options - PipelineRunsClientBeginCreateOptions contains the optional parameters for the PipelineRunsClient.BeginCreate
//     method.
func (client *PipelineRunsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsClientBeginCreateOptions) (*runtime.Poller[PipelineRunsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PipelineRunsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PipelineRunsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *PipelineRunsClient) create(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "PipelineRunsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *PipelineRunsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, pipelineRunCreateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a pipeline run from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - pipelineRunName - The name of the pipeline run.
//   - options - PipelineRunsClientBeginDeleteOptions contains the optional parameters for the PipelineRunsClient.BeginDelete
//     method.
func (client *PipelineRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientBeginDeleteOptions) (*runtime.Poller[PipelineRunsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, pipelineRunName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PipelineRunsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PipelineRunsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a pipeline run from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *PipelineRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PipelineRunsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelineRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the detailed information for a given pipeline run.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - pipelineRunName - The name of the pipeline run.
//   - options - PipelineRunsClientGetOptions contains the optional parameters for the PipelineRunsClient.Get method.
func (client *PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientGetOptions) (PipelineRunsClientGetResponse, error) {
	var err error
	const operationName = "PipelineRunsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PipelineRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelineRunsClient) getHandleResponse(resp *http.Response) (PipelineRunsClientGetResponse, error) {
	result := PipelineRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRun); err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the pipeline runs for the specified container registry.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - options - PipelineRunsClientListOptions contains the optional parameters for the PipelineRunsClient.NewListPager method.
func (client *PipelineRunsClient) NewListPager(resourceGroupName string, registryName string, options *PipelineRunsClientListOptions) *runtime.Pager[PipelineRunsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelineRunsClientListResponse]{
		More: func(page PipelineRunsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelineRunsClientListResponse) (PipelineRunsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PipelineRunsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelineRunsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PipelineRunsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelineRunsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PipelineRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *PipelineRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineRunsClient) listHandleResponse(resp *http.Response) (PipelineRunsClientListResponse, error) {
	result := PipelineRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRunListResult); err != nil {
		return PipelineRunsClientListResponse{}, err
	}
	return result, nil
}
