//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

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

// NotebookWorkspacesClient contains the methods for the NotebookWorkspaces group.
// Don't use this type directly, use NewNotebookWorkspacesClient() instead.
type NotebookWorkspacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNotebookWorkspacesClient creates a new instance of NotebookWorkspacesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNotebookWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NotebookWorkspacesClient, error) {
	cl, err := arm.NewClient(moduleName+".NotebookWorkspacesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NotebookWorkspacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - notebookCreateUpdateParameters - The notebook workspace to create for the current database account.
//   - options - NotebookWorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the NotebookWorkspacesClient.BeginCreateOrUpdate
//     method.
func (client *NotebookWorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[NotebookWorkspacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, notebookWorkspaceName, notebookCreateUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookWorkspacesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookWorkspacesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
func (client *NotebookWorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, notebookCreateUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotebookWorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, notebookCreateUpdateParameters)
}

// BeginDelete - Deletes the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientBeginDeleteOptions contains the optional parameters for the NotebookWorkspacesClient.BeginDelete
//     method.
func (client *NotebookWorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginDeleteOptions) (*runtime.Poller[NotebookWorkspacesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookWorkspacesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookWorkspacesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
func (client *NotebookWorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotebookWorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientGetOptions contains the optional parameters for the NotebookWorkspacesClient.Get method.
func (client *NotebookWorkspacesClient) Get(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientGetOptions) (NotebookWorkspacesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotebookWorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotebookWorkspacesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NotebookWorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NotebookWorkspacesClient) getHandleResponse(resp *http.Response) (NotebookWorkspacesClientGetResponse, error) {
	result := NotebookWorkspacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspace); err != nil {
		return NotebookWorkspacesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabaseAccountPager - Gets the notebook workspace resources of an existing Cosmos DB account.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - options - NotebookWorkspacesClientListByDatabaseAccountOptions contains the optional parameters for the NotebookWorkspacesClient.NewListByDatabaseAccountPager
//     method.
func (client *NotebookWorkspacesClient) NewListByDatabaseAccountPager(resourceGroupName string, accountName string, options *NotebookWorkspacesClientListByDatabaseAccountOptions) *runtime.Pager[NotebookWorkspacesClientListByDatabaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[NotebookWorkspacesClientListByDatabaseAccountResponse]{
		More: func(page NotebookWorkspacesClientListByDatabaseAccountResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *NotebookWorkspacesClientListByDatabaseAccountResponse) (NotebookWorkspacesClientListByDatabaseAccountResponse, error) {
			req, err := client.listByDatabaseAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return NotebookWorkspacesClientListByDatabaseAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NotebookWorkspacesClientListByDatabaseAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NotebookWorkspacesClientListByDatabaseAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseAccountHandleResponse(resp)
		},
	})
}

// listByDatabaseAccountCreateRequest creates the ListByDatabaseAccount request.
func (client *NotebookWorkspacesClient) listByDatabaseAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *NotebookWorkspacesClientListByDatabaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseAccountHandleResponse handles the ListByDatabaseAccount response.
func (client *NotebookWorkspacesClient) listByDatabaseAccountHandleResponse(resp *http.Response) (NotebookWorkspacesClientListByDatabaseAccountResponse, error) {
	result := NotebookWorkspacesClientListByDatabaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspaceListResult); err != nil {
		return NotebookWorkspacesClientListByDatabaseAccountResponse{}, err
	}
	return result, nil
}

// ListConnectionInfo - Retrieves the connection info for the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientListConnectionInfoOptions contains the optional parameters for the NotebookWorkspacesClient.ListConnectionInfo
//     method.
func (client *NotebookWorkspacesClient) ListConnectionInfo(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientListConnectionInfoOptions) (NotebookWorkspacesClientListConnectionInfoResponse, error) {
	req, err := client.listConnectionInfoCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, runtime.NewResponseError(resp)
	}
	return client.listConnectionInfoHandleResponse(resp)
}

// listConnectionInfoCreateRequest creates the ListConnectionInfo request.
func (client *NotebookWorkspacesClient) listConnectionInfoCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientListConnectionInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/listConnectionInfo"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConnectionInfoHandleResponse handles the ListConnectionInfo response.
func (client *NotebookWorkspacesClient) listConnectionInfoHandleResponse(resp *http.Response) (NotebookWorkspacesClientListConnectionInfoResponse, error) {
	result := NotebookWorkspacesClientListConnectionInfoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspaceConnectionInfoResult); err != nil {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	return result, nil
}

// BeginRegenerateAuthToken - Regenerates the auth token for the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientBeginRegenerateAuthTokenOptions contains the optional parameters for the NotebookWorkspacesClient.BeginRegenerateAuthToken
//     method.
func (client *NotebookWorkspacesClient) BeginRegenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (*runtime.Poller[NotebookWorkspacesClientRegenerateAuthTokenResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.regenerateAuthToken(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookWorkspacesClientRegenerateAuthTokenResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookWorkspacesClientRegenerateAuthTokenResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RegenerateAuthToken - Regenerates the auth token for the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
func (client *NotebookWorkspacesClient) regenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (*http.Response, error) {
	req, err := client.regenerateAuthTokenCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// regenerateAuthTokenCreateRequest creates the RegenerateAuthToken request.
func (client *NotebookWorkspacesClient) regenerateAuthTokenCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/regenerateAuthToken"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - Starts the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientBeginStartOptions contains the optional parameters for the NotebookWorkspacesClient.BeginStart
//     method.
func (client *NotebookWorkspacesClient) BeginStart(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginStartOptions) (*runtime.Poller[NotebookWorkspacesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[NotebookWorkspacesClientStartResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[NotebookWorkspacesClientStartResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Start - Starts the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-15-preview
func (client *NotebookWorkspacesClient) start(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *NotebookWorkspacesClient) startCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/start"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
