//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azquery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LogsClient contains the methods for the Logs group.
// Don't use this type directly, use a constructor function instead.
type LogsClient struct {
	host     string
	internal *azcore.Client
}

// QueryBatch - Executes a batch of Analytics queries for data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-27_Preview
//   - body - The batch request body
//   - options - LogsClientQueryBatchOptions contains the optional parameters for the LogsClient.QueryBatch method.
func (client *LogsClient) QueryBatch(ctx context.Context, body BatchRequest, options *LogsClientQueryBatchOptions) (LogsClientQueryBatchResponse, error) {
	req, err := client.queryBatchCreateRequest(ctx, body, options)
	if err != nil {
		return LogsClientQueryBatchResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogsClientQueryBatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogsClientQueryBatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryBatchHandleResponse(resp)
}

// queryBatchCreateRequest creates the QueryBatch request.
func (client *LogsClient) queryBatchCreateRequest(ctx context.Context, body BatchRequest, options *LogsClientQueryBatchOptions) (*policy.Request, error) {
	urlPath := "/$batch"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// queryBatchHandleResponse handles the QueryBatch response.
func (client *LogsClient) queryBatchHandleResponse(resp *http.Response) (LogsClientQueryBatchResponse, error) {
	result := LogsClientQueryBatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchResponse); err != nil {
		return LogsClientQueryBatchResponse{}, err
	}
	return result, nil
}

// QueryResource - Executes an Analytics query for data in the context of a resource. Here [https://docs.microsoft.com/azure/azure-monitor/logs/api/azure-resource-queries]
// is an example for using POST with an Analytics
// query.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-27_Preview
//   - resourceID - The identifier of the resource.
//   - body - The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
//   - options - LogsClientQueryResourceOptions contains the optional parameters for the LogsClient.QueryResource method.
func (client *LogsClient) QueryResource(ctx context.Context, resourceID string, body Body, options *LogsClientQueryResourceOptions) (LogsClientQueryResourceResponse, error) {
	req, err := client.queryResourceCreateRequest(ctx, resourceID, body, options)
	if err != nil {
		return LogsClientQueryResourceResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogsClientQueryResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogsClientQueryResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryResourceHandleResponse(resp)
}

// queryResourceCreateRequest creates the QueryResource request.
func (client *LogsClient) queryResourceCreateRequest(ctx context.Context, resourceID string, body Body, options *LogsClientQueryResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/query"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.Options != nil {
		req.Raw().Header["Prefer"] = []string{options.Options.preferHeader()}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// queryResourceHandleResponse handles the QueryResource response.
func (client *LogsClient) queryResourceHandleResponse(resp *http.Response) (LogsClientQueryResourceResponse, error) {
	result := LogsClientQueryResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Results); err != nil {
		return LogsClientQueryResourceResponse{}, err
	}
	return result, nil
}

// QueryWorkspace - Executes an Analytics query for data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-27_Preview
//   - workspaceID - Primary Workspace ID of the query. This is Workspace ID from the Properties blade in the Azure portal
//   - body - The Analytics query. Learn more about the Analytics query syntax [https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/]
//   - options - LogsClientQueryWorkspaceOptions contains the optional parameters for the LogsClient.QueryWorkspace method.
func (client *LogsClient) QueryWorkspace(ctx context.Context, workspaceID string, body Body, options *LogsClientQueryWorkspaceOptions) (LogsClientQueryWorkspaceResponse, error) {
	req, err := client.queryWorkspaceCreateRequest(ctx, workspaceID, body, options)
	if err != nil {
		return LogsClientQueryWorkspaceResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogsClientQueryWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogsClientQueryWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryWorkspaceHandleResponse(resp)
}

// queryWorkspaceCreateRequest creates the QueryWorkspace request.
func (client *LogsClient) queryWorkspaceCreateRequest(ctx context.Context, workspaceID string, body Body, options *LogsClientQueryWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/workspaces/{workspaceId}/query"
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.Options != nil {
		req.Raw().Header["Prefer"] = []string{options.Options.preferHeader()}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// queryWorkspaceHandleResponse handles the QueryWorkspace response.
func (client *LogsClient) queryWorkspaceHandleResponse(resp *http.Response) (LogsClientQueryWorkspaceResponse, error) {
	result := LogsClientQueryWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Results); err != nil {
		return LogsClientQueryWorkspaceResponse{}, err
	}
	return result, nil
}
