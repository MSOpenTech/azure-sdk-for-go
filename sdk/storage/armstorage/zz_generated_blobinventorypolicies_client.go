// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// BlobInventoryPoliciesClient contains the methods for the BlobInventoryPolicies group.
// Don't use this type directly, use NewBlobInventoryPoliciesClient() instead.
type BlobInventoryPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewBlobInventoryPoliciesClient creates a new instance of BlobInventoryPoliciesClient with the specified values.
func NewBlobInventoryPoliciesClient(con *armcore.Connection, subscriptionID string) *BlobInventoryPoliciesClient {
	return &BlobInventoryPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Sets the blob inventory policy to the specified storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BlobInventoryPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, properties BlobInventoryPolicy, options *BlobInventoryPoliciesCreateOrUpdateOptions) (BlobInventoryPolicyResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, blobInventoryPolicyName, properties, options)
	if err != nil {
		return BlobInventoryPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlobInventoryPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BlobInventoryPolicyResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BlobInventoryPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, properties BlobInventoryPolicy, options *BlobInventoryPoliciesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies/{blobInventoryPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if blobInventoryPolicyName == "" {
		return nil, errors.New("parameter blobInventoryPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobInventoryPolicyName}", url.PathEscape(string(blobInventoryPolicyName)))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(properties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BlobInventoryPoliciesClient) createOrUpdateHandleResponse(resp *azcore.Response) (BlobInventoryPolicyResponse, error) {
	var val *BlobInventoryPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BlobInventoryPolicyResponse{}, err
	}
	return BlobInventoryPolicyResponse{RawResponse: resp.Response, BlobInventoryPolicy: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *BlobInventoryPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes the blob inventory policy associated with the specified storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BlobInventoryPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, blobInventoryPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BlobInventoryPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies/{blobInventoryPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if blobInventoryPolicyName == "" {
		return nil, errors.New("parameter blobInventoryPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobInventoryPolicyName}", url.PathEscape(string(blobInventoryPolicyName)))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *BlobInventoryPoliciesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the blob inventory policy associated with the specified storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BlobInventoryPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesGetOptions) (BlobInventoryPolicyResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, blobInventoryPolicyName, options)
	if err != nil {
		return BlobInventoryPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BlobInventoryPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BlobInventoryPolicyResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BlobInventoryPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies/{blobInventoryPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if blobInventoryPolicyName == "" {
		return nil, errors.New("parameter blobInventoryPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobInventoryPolicyName}", url.PathEscape(string(blobInventoryPolicyName)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BlobInventoryPoliciesClient) getHandleResponse(resp *azcore.Response) (BlobInventoryPolicyResponse, error) {
	var val *BlobInventoryPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BlobInventoryPolicyResponse{}, err
	}
	return BlobInventoryPolicyResponse{RawResponse: resp.Response, BlobInventoryPolicy: val}, nil
}

// getHandleError handles the Get error response.
func (client *BlobInventoryPoliciesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets the blob inventory policy associated with the specified storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BlobInventoryPoliciesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *BlobInventoryPoliciesListOptions) (ListBlobInventoryPolicyResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return ListBlobInventoryPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ListBlobInventoryPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ListBlobInventoryPolicyResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BlobInventoryPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BlobInventoryPoliciesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BlobInventoryPoliciesClient) listHandleResponse(resp *azcore.Response) (ListBlobInventoryPolicyResponse, error) {
	var val *ListBlobInventoryPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListBlobInventoryPolicyResponse{}, err
	}
	return ListBlobInventoryPolicyResponse{RawResponse: resp.Response, ListBlobInventoryPolicy: val}, nil
}

// listHandleError handles the List error response.
func (client *BlobInventoryPoliciesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
