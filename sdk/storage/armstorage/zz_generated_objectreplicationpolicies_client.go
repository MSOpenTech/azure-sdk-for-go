//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ObjectReplicationPoliciesClient contains the methods for the ObjectReplicationPolicies group.
// Don't use this type directly, use NewObjectReplicationPoliciesClient() instead.
type ObjectReplicationPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewObjectReplicationPoliciesClient creates a new instance of ObjectReplicationPoliciesClient with the specified values.
func NewObjectReplicationPoliciesClient(con *arm.Connection, subscriptionID string) *ObjectReplicationPoliciesClient {
	return &ObjectReplicationPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update the object replication policy of the storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ObjectReplicationPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties ObjectReplicationPolicy, options *ObjectReplicationPoliciesCreateOrUpdateOptions) (ObjectReplicationPoliciesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, properties, options)
	if err != nil {
		return ObjectReplicationPoliciesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectReplicationPoliciesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ObjectReplicationPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties ObjectReplicationPolicy, options *ObjectReplicationPoliciesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
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
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ObjectReplicationPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ObjectReplicationPoliciesCreateOrUpdateResponse, error) {
	result := ObjectReplicationPoliciesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicy); err != nil {
		return ObjectReplicationPoliciesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ObjectReplicationPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the object replication policy associated with the specified storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ObjectReplicationPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesDeleteOptions) (ObjectReplicationPoliciesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, options)
	if err != nil {
		return ObjectReplicationPoliciesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ObjectReplicationPoliciesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ObjectReplicationPoliciesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ObjectReplicationPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
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
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ObjectReplicationPoliciesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the object replication policy of the storage account by policy ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ObjectReplicationPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesGetOptions) (ObjectReplicationPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, options)
	if err != nil {
		return ObjectReplicationPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectReplicationPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ObjectReplicationPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
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
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObjectReplicationPoliciesClient) getHandleResponse(resp *http.Response) (ObjectReplicationPoliciesGetResponse, error) {
	result := ObjectReplicationPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicy); err != nil {
		return ObjectReplicationPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ObjectReplicationPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List the object replication policies associated with the storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ObjectReplicationPoliciesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *ObjectReplicationPoliciesListOptions) (ObjectReplicationPoliciesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return ObjectReplicationPoliciesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectReplicationPoliciesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ObjectReplicationPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ObjectReplicationPoliciesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ObjectReplicationPoliciesClient) listHandleResponse(resp *http.Response) (ObjectReplicationPoliciesListResponse, error) {
	result := ObjectReplicationPoliciesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicies); err != nil {
		return ObjectReplicationPoliciesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ObjectReplicationPoliciesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
