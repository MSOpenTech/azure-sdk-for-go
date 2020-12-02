// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SSHPublicKeysClient contains the methods for the SSHPublicKeys group.
// Don't use this type directly, use NewSSHPublicKeysClient() instead.
type SSHPublicKeysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSSHPublicKeysClient creates a new instance of SSHPublicKeysClient with the specified values.
func NewSSHPublicKeysClient(con *armcore.Connection, subscriptionID string) SSHPublicKeysClient {
	return SSHPublicKeysClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client SSHPublicKeysClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Create - Creates a new SSH public key resource.
func (client SSHPublicKeysClient) Create(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource, options *SSHPublicKeysCreateOptions) (SSHPublicKeyResourceResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters, options)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return SSHPublicKeyResourceResponse{}, client.createHandleError(resp)
	}
	result, err := client.createHandleResponse(resp)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	return result, nil
}

// createCreateRequest creates the Create request.
func (client SSHPublicKeysClient) createCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource, options *SSHPublicKeysCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createHandleResponse handles the Create response.
func (client SSHPublicKeysClient) createHandleResponse(resp *azcore.Response) (SSHPublicKeyResourceResponse, error) {
	result := SSHPublicKeyResourceResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.SSHPublicKeyResource)
	return result, err
}

// createHandleError handles the Create error response.
func (client SSHPublicKeysClient) createHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Delete an SSH public key.
func (client SSHPublicKeysClient) Delete(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client SSHPublicKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client SSHPublicKeysClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GenerateKeyPair - Generates and returns a public/private key pair and populates the SSH public key resource with the public key. The length of the key
// will be 3072 bits. This operation can only be performed once per
// SSH public key resource.
func (client SSHPublicKeysClient) GenerateKeyPair(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGenerateKeyPairOptions) (SSHPublicKeyGenerateKeyPairResultResponse, error) {
	req, err := client.generateKeyPairCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeyGenerateKeyPairResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeyGenerateKeyPairResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SSHPublicKeyGenerateKeyPairResultResponse{}, client.generateKeyPairHandleError(resp)
	}
	result, err := client.generateKeyPairHandleResponse(resp)
	if err != nil {
		return SSHPublicKeyGenerateKeyPairResultResponse{}, err
	}
	return result, nil
}

// generateKeyPairCreateRequest creates the GenerateKeyPair request.
func (client SSHPublicKeysClient) generateKeyPairCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGenerateKeyPairOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}/generateKeyPair"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// generateKeyPairHandleResponse handles the GenerateKeyPair response.
func (client SSHPublicKeysClient) generateKeyPairHandleResponse(resp *azcore.Response) (SSHPublicKeyGenerateKeyPairResultResponse, error) {
	result := SSHPublicKeyGenerateKeyPairResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.SSHPublicKeyGenerateKeyPairResult)
	return result, err
}

// generateKeyPairHandleError handles the GenerateKeyPair error response.
func (client SSHPublicKeysClient) generateKeyPairHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Retrieves information about an SSH public key.
func (client SSHPublicKeysClient) Get(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGetOptions) (SSHPublicKeyResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sshPublicKeyName, options)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SSHPublicKeyResourceResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client SSHPublicKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, options *SSHPublicKeysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client SSHPublicKeysClient) getHandleResponse(resp *azcore.Response) (SSHPublicKeyResourceResponse, error) {
	result := SSHPublicKeyResourceResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.SSHPublicKeyResource)
	return result, err
}

// getHandleError handles the Get error response.
func (client SSHPublicKeysClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Lists all of the SSH public keys in the specified resource group. Use the nextLink property in the response to get the next page
// of SSH public keys.
func (client SSHPublicKeysClient) ListByResourceGroup(resourceGroupName string, options *SSHPublicKeysListByResourceGroupOptions) SSHPublicKeysGroupListResultPager {
	return &sshPublicKeysGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp SSHPublicKeysGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SSHPublicKeysGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client SSHPublicKeysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SSHPublicKeysListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client SSHPublicKeysClient) listByResourceGroupHandleResponse(resp *azcore.Response) (SSHPublicKeysGroupListResultResponse, error) {
	result := SSHPublicKeysGroupListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.SSHPublicKeysGroupListResult)
	return result, err
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client SSHPublicKeysClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySubscription - Lists all of the SSH public keys in the subscription. Use the nextLink property in the response to get the next page of SSH public
// keys.
func (client SSHPublicKeysClient) ListBySubscription(options *SSHPublicKeysListBySubscriptionOptions) SSHPublicKeysGroupListResultPager {
	return &sshPublicKeysGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp SSHPublicKeysGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SSHPublicKeysGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client SSHPublicKeysClient) listBySubscriptionCreateRequest(ctx context.Context, options *SSHPublicKeysListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/sshPublicKeys"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client SSHPublicKeysClient) listBySubscriptionHandleResponse(resp *azcore.Response) (SSHPublicKeysGroupListResultResponse, error) {
	result := SSHPublicKeysGroupListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.SSHPublicKeysGroupListResult)
	return result, err
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client SSHPublicKeysClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Update - Updates a new SSH public key resource.
func (client SSHPublicKeysClient) Update(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource, options *SSHPublicKeysUpdateOptions) (SSHPublicKeyResourceResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters, options)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SSHPublicKeyResourceResponse{}, client.updateHandleError(resp)
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return SSHPublicKeyResourceResponse{}, err
	}
	return result, nil
}

// updateCreateRequest creates the Update request.
func (client SSHPublicKeysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource, options *SSHPublicKeysUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client SSHPublicKeysClient) updateHandleResponse(resp *azcore.Response) (SSHPublicKeyResourceResponse, error) {
	result := SSHPublicKeyResourceResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.SSHPublicKeyResource)
	return result, err
}

// updateHandleError handles the Update error response.
func (client SSHPublicKeysClient) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
