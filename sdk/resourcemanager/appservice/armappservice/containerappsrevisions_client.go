//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// ContainerAppsRevisionsClient contains the methods for the ContainerAppsRevisions group.
// Don't use this type directly, use NewContainerAppsRevisionsClient() instead.
type ContainerAppsRevisionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsRevisionsClient creates a new instance of ContainerAppsRevisionsClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsRevisionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsRevisionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsRevisionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// ActivateRevision - Activates a revision for a Container App
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - containerAppName - Name of the Container App.
//   - name - Name of the Container App Revision to activate
//   - options - ContainerAppsRevisionsClientActivateRevisionOptions contains the optional parameters for the ContainerAppsRevisionsClient.ActivateRevision
//     method.
func (client *ContainerAppsRevisionsClient) ActivateRevision(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientActivateRevisionOptions) (ContainerAppsRevisionsClientActivateRevisionResponse, error) {
	var err error
	const operationName = "ContainerAppsRevisionsClient.ActivateRevision"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.activateRevisionCreateRequest(ctx, resourceGroupName, containerAppName, name, options)
	if err != nil {
		return ContainerAppsRevisionsClientActivateRevisionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsRevisionsClientActivateRevisionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsRevisionsClientActivateRevisionResponse{}, err
	}
	return ContainerAppsRevisionsClientActivateRevisionResponse{}, nil
}

// activateRevisionCreateRequest creates the ActivateRevision request.
func (client *ContainerAppsRevisionsClient) activateRevisionCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientActivateRevisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{containerAppName}/revisions/{name}/activate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeactivateRevision - Deactivates a revision for a Container App
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - containerAppName - Name of the Container App.
//   - name - Name of the Container App Revision to deactivate
//   - options - ContainerAppsRevisionsClientDeactivateRevisionOptions contains the optional parameters for the ContainerAppsRevisionsClient.DeactivateRevision
//     method.
func (client *ContainerAppsRevisionsClient) DeactivateRevision(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientDeactivateRevisionOptions) (ContainerAppsRevisionsClientDeactivateRevisionResponse, error) {
	var err error
	const operationName = "ContainerAppsRevisionsClient.DeactivateRevision"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deactivateRevisionCreateRequest(ctx, resourceGroupName, containerAppName, name, options)
	if err != nil {
		return ContainerAppsRevisionsClientDeactivateRevisionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsRevisionsClientDeactivateRevisionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsRevisionsClientDeactivateRevisionResponse{}, err
	}
	return ContainerAppsRevisionsClientDeactivateRevisionResponse{}, nil
}

// deactivateRevisionCreateRequest creates the DeactivateRevision request.
func (client *ContainerAppsRevisionsClient) deactivateRevisionCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientDeactivateRevisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{containerAppName}/revisions/{name}/deactivate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetRevision - Get a revision of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - containerAppName - Name of the Container App.
//   - name - Name of the Container App Revision.
//   - options - ContainerAppsRevisionsClientGetRevisionOptions contains the optional parameters for the ContainerAppsRevisionsClient.GetRevision
//     method.
func (client *ContainerAppsRevisionsClient) GetRevision(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientGetRevisionOptions) (ContainerAppsRevisionsClientGetRevisionResponse, error) {
	var err error
	const operationName = "ContainerAppsRevisionsClient.GetRevision"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getRevisionCreateRequest(ctx, resourceGroupName, containerAppName, name, options)
	if err != nil {
		return ContainerAppsRevisionsClientGetRevisionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsRevisionsClientGetRevisionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsRevisionsClientGetRevisionResponse{}, err
	}
	resp, err := client.getRevisionHandleResponse(httpResp)
	return resp, err
}

// getRevisionCreateRequest creates the GetRevision request.
func (client *ContainerAppsRevisionsClient) getRevisionCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientGetRevisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{containerAppName}/revisions/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRevisionHandleResponse handles the GetRevision response.
func (client *ContainerAppsRevisionsClient) getRevisionHandleResponse(resp *http.Response) (ContainerAppsRevisionsClientGetRevisionResponse, error) {
	result := ContainerAppsRevisionsClientGetRevisionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Revision); err != nil {
		return ContainerAppsRevisionsClientGetRevisionResponse{}, err
	}
	return result, nil
}

// NewListRevisionsPager - Get the Revisions for a given Container App.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - containerAppName - Name of the Container App for which Revisions are needed.
//   - options - ContainerAppsRevisionsClientListRevisionsOptions contains the optional parameters for the ContainerAppsRevisionsClient.NewListRevisionsPager
//     method.
func (client *ContainerAppsRevisionsClient) NewListRevisionsPager(resourceGroupName string, containerAppName string, options *ContainerAppsRevisionsClientListRevisionsOptions) *runtime.Pager[ContainerAppsRevisionsClientListRevisionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsRevisionsClientListRevisionsResponse]{
		More: func(page ContainerAppsRevisionsClientListRevisionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsRevisionsClientListRevisionsResponse) (ContainerAppsRevisionsClientListRevisionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsRevisionsClient.NewListRevisionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listRevisionsCreateRequest(ctx, resourceGroupName, containerAppName, options)
			}, nil)
			if err != nil {
				return ContainerAppsRevisionsClientListRevisionsResponse{}, err
			}
			return client.listRevisionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listRevisionsCreateRequest creates the ListRevisions request.
func (client *ContainerAppsRevisionsClient) listRevisionsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsRevisionsClientListRevisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{containerAppName}/revisions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRevisionsHandleResponse handles the ListRevisions response.
func (client *ContainerAppsRevisionsClient) listRevisionsHandleResponse(resp *http.Response) (ContainerAppsRevisionsClientListRevisionsResponse, error) {
	result := ContainerAppsRevisionsClientListRevisionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RevisionCollection); err != nil {
		return ContainerAppsRevisionsClientListRevisionsResponse{}, err
	}
	return result, nil
}

// RestartRevision - Restarts a revision for a Container App
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - containerAppName - Name of the Container App.
//   - name - Name of the Container App Revision to restart
//   - options - ContainerAppsRevisionsClientRestartRevisionOptions contains the optional parameters for the ContainerAppsRevisionsClient.RestartRevision
//     method.
func (client *ContainerAppsRevisionsClient) RestartRevision(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientRestartRevisionOptions) (ContainerAppsRevisionsClientRestartRevisionResponse, error) {
	var err error
	const operationName = "ContainerAppsRevisionsClient.RestartRevision"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartRevisionCreateRequest(ctx, resourceGroupName, containerAppName, name, options)
	if err != nil {
		return ContainerAppsRevisionsClientRestartRevisionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsRevisionsClientRestartRevisionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsRevisionsClientRestartRevisionResponse{}, err
	}
	return ContainerAppsRevisionsClientRestartRevisionResponse{}, nil
}

// restartRevisionCreateRequest creates the RestartRevision request.
func (client *ContainerAppsRevisionsClient) restartRevisionCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *ContainerAppsRevisionsClientRestartRevisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{containerAppName}/revisions/{name}/restart"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
