//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

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

// TransactionNodesClient contains the methods for the TransactionNodes group.
// Don't use this type directly, use NewTransactionNodesClient() instead.
type TransactionNodesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTransactionNodesClient creates a new instance of TransactionNodesClient with the specified values.
// subscriptionID - Gets the subscription Id which uniquely identifies the Microsoft Azure subscription. The subscription
// ID is part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTransactionNodesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TransactionNodesClient, error) {
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
	client := &TransactionNodesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create or update the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// transactionNodeName - Transaction node name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientBeginCreateOptions contains the optional parameters for the TransactionNodesClient.BeginCreate
// method.
func (client *TransactionNodesClient) BeginCreate(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientBeginCreateOptions) (*runtime.Poller[TransactionNodesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[TransactionNodesClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[TransactionNodesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create or update the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
func (client *TransactionNodesClient) create(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
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
func (client *TransactionNodesClient) createCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.TransactionNode != nil {
		return req, runtime.MarshalAsJSON(req, *options.TransactionNode)
	}
	return req, nil
}

// BeginDelete - Delete the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// transactionNodeName - Transaction node name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientBeginDeleteOptions contains the optional parameters for the TransactionNodesClient.BeginDelete
// method.
func (client *TransactionNodesClient) BeginDelete(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientBeginDeleteOptions) (*runtime.Poller[TransactionNodesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[TransactionNodesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[TransactionNodesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
func (client *TransactionNodesClient) deleteOperation(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TransactionNodesClient) deleteCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get the details of the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// transactionNodeName - Transaction node name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientGetOptions contains the optional parameters for the TransactionNodesClient.Get method.
func (client *TransactionNodesClient) Get(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientGetOptions) (TransactionNodesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TransactionNodesClient) getCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TransactionNodesClient) getHandleResponse(resp *http.Response) (TransactionNodesClientGetResponse, error) {
	result := TransactionNodesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionNode); err != nil {
		return TransactionNodesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the transaction nodes for a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientListOptions contains the optional parameters for the TransactionNodesClient.List method.
func (client *TransactionNodesClient) NewListPager(blockchainMemberName string, resourceGroupName string, options *TransactionNodesClientListOptions) *runtime.Pager[TransactionNodesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransactionNodesClientListResponse]{
		More: func(page TransactionNodesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransactionNodesClientListResponse) (TransactionNodesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TransactionNodesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TransactionNodesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TransactionNodesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TransactionNodesClient) listCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *TransactionNodesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TransactionNodesClient) listHandleResponse(resp *http.Response) (TransactionNodesClientListResponse, error) {
	result := TransactionNodesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionNodeCollection); err != nil {
		return TransactionNodesClientListResponse{}, err
	}
	return result, nil
}

// ListAPIKeys - List the API keys for the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// transactionNodeName - Transaction node name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientListAPIKeysOptions contains the optional parameters for the TransactionNodesClient.ListAPIKeys
// method.
func (client *TransactionNodesClient) ListAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientListAPIKeysOptions) (TransactionNodesClientListAPIKeysResponse, error) {
	req, err := client.listAPIKeysCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesClientListAPIKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesClientListAPIKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesClientListAPIKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAPIKeysHandleResponse(resp)
}

// listAPIKeysCreateRequest creates the ListAPIKeys request.
func (client *TransactionNodesClient) listAPIKeysCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientListAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}/listApiKeys"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAPIKeysHandleResponse handles the ListAPIKeys response.
func (client *TransactionNodesClient) listAPIKeysHandleResponse(resp *http.Response) (TransactionNodesClientListAPIKeysResponse, error) {
	result := TransactionNodesClientListAPIKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyCollection); err != nil {
		return TransactionNodesClientListAPIKeysResponse{}, err
	}
	return result, nil
}

// ListRegenerateAPIKeys - Regenerate the API keys for the blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// transactionNodeName - Transaction node name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientListRegenerateAPIKeysOptions contains the optional parameters for the TransactionNodesClient.ListRegenerateAPIKeys
// method.
func (client *TransactionNodesClient) ListRegenerateAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientListRegenerateAPIKeysOptions) (TransactionNodesClientListRegenerateAPIKeysResponse, error) {
	req, err := client.listRegenerateAPIKeysCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesClientListRegenerateAPIKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesClientListRegenerateAPIKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesClientListRegenerateAPIKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listRegenerateAPIKeysHandleResponse(resp)
}

// listRegenerateAPIKeysCreateRequest creates the ListRegenerateAPIKeys request.
func (client *TransactionNodesClient) listRegenerateAPIKeysCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientListRegenerateAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}/regenerateApiKeys"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.APIKey != nil {
		return req, runtime.MarshalAsJSON(req, *options.APIKey)
	}
	return req, nil
}

// listRegenerateAPIKeysHandleResponse handles the ListRegenerateAPIKeys response.
func (client *TransactionNodesClient) listRegenerateAPIKeysHandleResponse(resp *http.Response) (TransactionNodesClientListRegenerateAPIKeysResponse, error) {
	result := TransactionNodesClientListRegenerateAPIKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyCollection); err != nil {
		return TransactionNodesClientListRegenerateAPIKeysResponse{}, err
	}
	return result, nil
}

// Update - Update the transaction node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// transactionNodeName - Transaction node name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - TransactionNodesClientUpdateOptions contains the optional parameters for the TransactionNodesClient.Update method.
func (client *TransactionNodesClient) Update(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientUpdateOptions) (TransactionNodesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, options)
	if err != nil {
		return TransactionNodesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TransactionNodesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TransactionNodesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TransactionNodesClient) updateCreateRequest(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, options *TransactionNodesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if transactionNodeName == "" {
		return nil, errors.New("parameter transactionNodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transactionNodeName}", url.PathEscape(transactionNodeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.TransactionNode != nil {
		return req, runtime.MarshalAsJSON(req, *options.TransactionNode)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TransactionNodesClient) updateHandleResponse(resp *http.Response) (TransactionNodesClientUpdateResponse, error) {
	result := TransactionNodesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionNode); err != nil {
		return TransactionNodesClientUpdateResponse{}, err
	}
	return result, nil
}
