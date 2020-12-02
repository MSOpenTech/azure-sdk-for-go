// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualRouterPeeringsClient contains the methods for the VirtualRouterPeerings group.
// Don't use this type directly, use NewVirtualRouterPeeringsClient() instead.
type VirtualRouterPeeringsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualRouterPeeringsClient creates a new instance of VirtualRouterPeeringsClient with the specified values.
func NewVirtualRouterPeeringsClient(con *armcore.Connection, subscriptionID string) VirtualRouterPeeringsClient {
	return VirtualRouterPeeringsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualRouterPeeringsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates the specified Virtual Router Peering.
func (client VirtualRouterPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsBeginCreateOrUpdateOptions) (VirtualRouterPeeringPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualRouterName, peeringName, parameters, options)
	if err != nil {
		return VirtualRouterPeeringPollerResponse{}, err
	}
	result := VirtualRouterPeeringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualRouterPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualRouterPeeringPollerResponse{}, err
	}
	poller := &virtualRouterPeeringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualRouterPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualRouterPeeringPoller from the specified resume token.
// token - The value must come from a previous call to VirtualRouterPeeringPoller.ResumeToken().
func (client VirtualRouterPeeringsClient) ResumeCreateOrUpdate(token string) (VirtualRouterPeeringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualRouterPeeringsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualRouterPeeringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Virtual Router Peering.
func (client VirtualRouterPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client VirtualRouterPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, parameters VirtualRouterPeering, options *VirtualRouterPeeringsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client VirtualRouterPeeringsClient) createOrUpdateHandleResponse(resp *azcore.Response) (VirtualRouterPeeringResponse, error) {
	result := VirtualRouterPeeringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualRouterPeering)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VirtualRouterPeeringsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified peering from a Virtual Router.
func (client VirtualRouterPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualRouterPeeringsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client VirtualRouterPeeringsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualRouterPeeringsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified peering from a Virtual Router.
func (client VirtualRouterPeeringsClient) delete(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client VirtualRouterPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client VirtualRouterPeeringsClient) deleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Virtual Router Peering.
func (client VirtualRouterPeeringsClient) Get(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsGetOptions) (VirtualRouterPeeringResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualRouterName, peeringName, options)
	if err != nil {
		return VirtualRouterPeeringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualRouterPeeringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualRouterPeeringResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VirtualRouterPeeringResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VirtualRouterPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, peeringName string, options *VirtualRouterPeeringsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client VirtualRouterPeeringsClient) getHandleResponse(resp *azcore.Response) (VirtualRouterPeeringResponse, error) {
	result := VirtualRouterPeeringResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualRouterPeering)
	return result, err
}

// getHandleError handles the Get error response.
func (client VirtualRouterPeeringsClient) getHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all Virtual Router Peerings in a Virtual Router resource.
func (client VirtualRouterPeeringsClient) List(resourceGroupName string, virtualRouterName string, options *VirtualRouterPeeringsListOptions) VirtualRouterPeeringListResultPager {
	return &virtualRouterPeeringListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualRouterName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp VirtualRouterPeeringListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualRouterPeeringListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client VirtualRouterPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualRouterName string, options *VirtualRouterPeeringsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualRouterName}", url.PathEscape(virtualRouterName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client VirtualRouterPeeringsClient) listHandleResponse(resp *azcore.Response) (VirtualRouterPeeringListResultResponse, error) {
	result := VirtualRouterPeeringListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualRouterPeeringListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client VirtualRouterPeeringsClient) listHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
