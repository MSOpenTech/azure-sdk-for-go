//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

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

// SQLServerInstancesClient contains the methods for the SQLServerInstances group.
// Don't use this type directly, use NewSQLServerInstancesClient() instead.
type SQLServerInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLServerInstancesClient creates a new instance of SQLServerInstancesClient with the specified values.
// subscriptionID - The ID of the Azure subscription
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLServerInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLServerInstancesClient, error) {
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
	client := &SQLServerInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates or replaces a SQL Server Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group
// sqlServerInstanceName - Name of SQL Server Instance
// sqlServerInstance - The SQL Server Instance to be created or updated.
// options - SQLServerInstancesClientBeginCreateOptions contains the optional parameters for the SQLServerInstancesClient.BeginCreate
// method.
func (client *SQLServerInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, sqlServerInstance SQLServerInstance, options *SQLServerInstancesClientBeginCreateOptions) (*armruntime.Poller[SQLServerInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sqlServerInstanceName, sqlServerInstance, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SQLServerInstancesClientCreateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SQLServerInstancesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates or replaces a SQL Server Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLServerInstancesClient) create(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, sqlServerInstance SQLServerInstance, options *SQLServerInstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, sqlServerInstance, options)
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

// createCreateRequest creates the Create request.
func (client *SQLServerInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, sqlServerInstance SQLServerInstance, options *SQLServerInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sqlServerInstance)
}

// BeginDelete - Deletes a SQL Server Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group
// sqlServerInstanceName - Name of SQL Server Instance
// options - SQLServerInstancesClientBeginDeleteOptions contains the optional parameters for the SQLServerInstancesClient.BeginDelete
// method.
func (client *SQLServerInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, options *SQLServerInstancesClientBeginDeleteOptions) (*armruntime.Poller[SQLServerInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlServerInstanceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SQLServerInstancesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SQLServerInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a SQL Server Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLServerInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, options *SQLServerInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLServerInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, options *SQLServerInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves a SQL Server Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group
// sqlServerInstanceName - Name of SQL Server Instance
// options - SQLServerInstancesClientGetOptions contains the optional parameters for the SQLServerInstancesClient.Get method.
func (client *SQLServerInstancesClient) Get(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, options *SQLServerInstancesClientGetOptions) (SQLServerInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, options)
	if err != nil {
		return SQLServerInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLServerInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLServerInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLServerInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, options *SQLServerInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLServerInstancesClient) getHandleResponse(resp *http.Response) (SQLServerInstancesClientGetResponse, error) {
	result := SQLServerInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerInstance); err != nil {
		return SQLServerInstancesClientGetResponse{}, err
	}
	return result, nil
}

// List - List sqlServerInstance resources in the subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - SQLServerInstancesClientListOptions contains the optional parameters for the SQLServerInstancesClient.List method.
func (client *SQLServerInstancesClient) List(options *SQLServerInstancesClientListOptions) *runtime.Pager[SQLServerInstancesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[SQLServerInstancesClientListResponse]{
		More: func(page SQLServerInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLServerInstancesClientListResponse) (SQLServerInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLServerInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLServerInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLServerInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLServerInstancesClient) listCreateRequest(ctx context.Context, options *SQLServerInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/sqlServerInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLServerInstancesClient) listHandleResponse(resp *http.Response) (SQLServerInstancesClientListResponse, error) {
	result := SQLServerInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerInstanceListResult); err != nil {
		return SQLServerInstancesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Gets all sqlServerInstances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group
// options - SQLServerInstancesClientListByResourceGroupOptions contains the optional parameters for the SQLServerInstancesClient.ListByResourceGroup
// method.
func (client *SQLServerInstancesClient) ListByResourceGroup(resourceGroupName string, options *SQLServerInstancesClientListByResourceGroupOptions) *runtime.Pager[SQLServerInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[SQLServerInstancesClientListByResourceGroupResponse]{
		More: func(page SQLServerInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLServerInstancesClientListByResourceGroupResponse) (SQLServerInstancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLServerInstancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLServerInstancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLServerInstancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLServerInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLServerInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances"
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
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLServerInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLServerInstancesClientListByResourceGroupResponse, error) {
	result := SQLServerInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerInstanceListResult); err != nil {
		return SQLServerInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a SQL Server Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group
// sqlServerInstanceName - Name of SQL Server Instance
// parameters - The SQL Server Instance.
// options - SQLServerInstancesClientUpdateOptions contains the optional parameters for the SQLServerInstancesClient.Update
// method.
func (client *SQLServerInstancesClient) Update(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, parameters SQLServerInstanceUpdate, options *SQLServerInstancesClientUpdateOptions) (SQLServerInstancesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, parameters, options)
	if err != nil {
		return SQLServerInstancesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLServerInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLServerInstancesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SQLServerInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, parameters SQLServerInstanceUpdate, options *SQLServerInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SQLServerInstancesClient) updateHandleResponse(resp *http.Response) (SQLServerInstancesClientUpdateResponse, error) {
	result := SQLServerInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerInstance); err != nil {
		return SQLServerInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
