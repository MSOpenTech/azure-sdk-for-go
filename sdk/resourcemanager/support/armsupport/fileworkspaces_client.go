//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

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

// FileWorkspacesClient contains the methods for the FileWorkspaces group.
// Don't use this type directly, use NewFileWorkspacesClient() instead.
type FileWorkspacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFileWorkspacesClient creates a new instance of FileWorkspacesClient with the specified values.
//   - subscriptionID - Azure subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFileWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FileWorkspacesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FileWorkspacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a new file workspace for the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - fileWorkspaceName - File workspace name.
//   - options - FileWorkspacesClientCreateOptions contains the optional parameters for the FileWorkspacesClient.Create method.
func (client *FileWorkspacesClient) Create(ctx context.Context, fileWorkspaceName string, options *FileWorkspacesClientCreateOptions) (FileWorkspacesClientCreateResponse, error) {
	var err error
	const operationName = "FileWorkspacesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, fileWorkspaceName, options)
	if err != nil {
		return FileWorkspacesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileWorkspacesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return FileWorkspacesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *FileWorkspacesClient) createCreateRequest(ctx context.Context, fileWorkspaceName string, options *FileWorkspacesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/fileWorkspaces/{fileWorkspaceName}"
	if fileWorkspaceName == "" {
		return nil, errors.New("parameter fileWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileWorkspaceName}", url.PathEscape(fileWorkspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *FileWorkspacesClient) createHandleResponse(resp *http.Response) (FileWorkspacesClientCreateResponse, error) {
	result := FileWorkspacesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileWorkspaceDetails); err != nil {
		return FileWorkspacesClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Gets details for a specific file workspace in an Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - fileWorkspaceName - File Workspace Name
//   - options - FileWorkspacesClientGetOptions contains the optional parameters for the FileWorkspacesClient.Get method.
func (client *FileWorkspacesClient) Get(ctx context.Context, fileWorkspaceName string, options *FileWorkspacesClientGetOptions) (FileWorkspacesClientGetResponse, error) {
	var err error
	const operationName = "FileWorkspacesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, fileWorkspaceName, options)
	if err != nil {
		return FileWorkspacesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FileWorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FileWorkspacesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FileWorkspacesClient) getCreateRequest(ctx context.Context, fileWorkspaceName string, options *FileWorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/fileWorkspaces/{fileWorkspaceName}"
	if fileWorkspaceName == "" {
		return nil, errors.New("parameter fileWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileWorkspaceName}", url.PathEscape(fileWorkspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FileWorkspacesClient) getHandleResponse(resp *http.Response) (FileWorkspacesClientGetResponse, error) {
	result := FileWorkspacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileWorkspaceDetails); err != nil {
		return FileWorkspacesClientGetResponse{}, err
	}
	return result, nil
}
