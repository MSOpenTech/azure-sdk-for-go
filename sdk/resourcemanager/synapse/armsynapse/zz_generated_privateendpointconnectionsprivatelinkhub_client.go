//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateEndpointConnectionsPrivateLinkHubClient contains the methods for the PrivateEndpointConnectionsPrivateLinkHub group.
// Don't use this type directly, use NewPrivateEndpointConnectionsPrivateLinkHubClient() instead.
type PrivateEndpointConnectionsPrivateLinkHubClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateEndpointConnectionsPrivateLinkHubClient creates a new instance of PrivateEndpointConnectionsPrivateLinkHubClient with the specified values.
func NewPrivateEndpointConnectionsPrivateLinkHubClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionsPrivateLinkHubClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PrivateEndpointConnectionsPrivateLinkHubClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get all PrivateEndpointConnection in the PrivateLinkHub by name
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) Get(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsPrivateLinkHubGetOptions) (PrivateEndpointConnectionsPrivateLinkHubGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateLinkHubName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsPrivateLinkHubGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsPrivateLinkHubGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateLinkHubName == "" {
		return nil, errors.New("parameter privateLinkHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkHubName}", url.PathEscape(privateLinkHubName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsPrivateLinkHubGetResponse, error) {
	result := PrivateEndpointConnectionsPrivateLinkHubGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionForPrivateLinkHub); err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) getHandleError(resp *http.Response) error {
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

// List - Get all PrivateEndpointConnections in the PrivateLinkHub
// If the operation fails it returns the *ErrorResponse error type.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) List(resourceGroupName string, privateLinkHubName string, options *PrivateEndpointConnectionsPrivateLinkHubListOptions) *PrivateEndpointConnectionsPrivateLinkHubListPager {
	return &PrivateEndpointConnectionsPrivateLinkHubListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateLinkHubName, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndpointConnectionsPrivateLinkHubListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateEndpointConnectionsPrivateLinkHubListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateLinkHubName == "" {
		return nil, errors.New("parameter privateLinkHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkHubName}", url.PathEscape(privateLinkHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) listHandleResponse(resp *http.Response) (PrivateEndpointConnectionsPrivateLinkHubListResponse, error) {
	result := PrivateEndpointConnectionsPrivateLinkHubListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponse); err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) listHandleError(resp *http.Response) error {
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
