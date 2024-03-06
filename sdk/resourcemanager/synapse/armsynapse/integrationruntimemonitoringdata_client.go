//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// IntegrationRuntimeMonitoringDataClient contains the methods for the IntegrationRuntimeMonitoringData group.
// Don't use this type directly, use NewIntegrationRuntimeMonitoringDataClient() instead.
type IntegrationRuntimeMonitoringDataClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationRuntimeMonitoringDataClient creates a new instance of IntegrationRuntimeMonitoringDataClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationRuntimeMonitoringDataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeMonitoringDataClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationRuntimeMonitoringDataClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Get monitoring data for an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - integrationRuntimeName - Integration runtime name
//   - options - IntegrationRuntimeMonitoringDataClientListOptions contains the optional parameters for the IntegrationRuntimeMonitoringDataClient.List
//     method.
func (client *IntegrationRuntimeMonitoringDataClient) List(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeMonitoringDataClientListOptions) (IntegrationRuntimeMonitoringDataClientListResponse, error) {
	var err error
	const operationName = "IntegrationRuntimeMonitoringDataClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeMonitoringDataClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeMonitoringDataClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationRuntimeMonitoringDataClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *IntegrationRuntimeMonitoringDataClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeMonitoringDataClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/monitoringData"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationRuntimeMonitoringDataClient) listHandleResponse(resp *http.Response) (IntegrationRuntimeMonitoringDataClientListResponse, error) {
	result := IntegrationRuntimeMonitoringDataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeMonitoringData); err != nil {
		return IntegrationRuntimeMonitoringDataClientListResponse{}, err
	}
	return result, nil
}
