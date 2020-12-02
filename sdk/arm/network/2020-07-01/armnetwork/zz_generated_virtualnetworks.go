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

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
func NewVirtualNetworksClient(con *armcore.Connection, subscriptionID string) VirtualNetworksClient {
	return VirtualNetworksClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualNetworksClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CheckIPAddressAvailability - Checks whether a private IP address is available for use.
func (client VirtualNetworksClient) CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *VirtualNetworksCheckIPAddressAvailabilityOptions) (IPAddressAvailabilityResultResponse, error) {
	req, err := client.checkIPAddressAvailabilityCreateRequest(ctx, resourceGroupName, virtualNetworkName, ipAddress, options)
	if err != nil {
		return IPAddressAvailabilityResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return IPAddressAvailabilityResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IPAddressAvailabilityResultResponse{}, client.checkIPAddressAvailabilityHandleError(resp)
	}
	result, err := client.checkIPAddressAvailabilityHandleResponse(resp)
	if err != nil {
		return IPAddressAvailabilityResultResponse{}, err
	}
	return result, nil
}

// checkIPAddressAvailabilityCreateRequest creates the CheckIPAddressAvailability request.
func (client VirtualNetworksClient) checkIPAddressAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, ipAddress string, options *VirtualNetworksCheckIPAddressAvailabilityOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("ipAddress", ipAddress)
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// checkIPAddressAvailabilityHandleResponse handles the CheckIPAddressAvailability response.
func (client VirtualNetworksClient) checkIPAddressAvailabilityHandleResponse(resp *azcore.Response) (IPAddressAvailabilityResultResponse, error) {
	result := IPAddressAvailabilityResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.IPAddressAvailabilityResult)
	return result, err
}

// checkIPAddressAvailabilityHandleError handles the CheckIPAddressAvailability error response.
func (client VirtualNetworksClient) checkIPAddressAvailabilityHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginCreateOrUpdate - Creates or updates a virtual network in the specified resource group.
func (client VirtualNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksBeginCreateOrUpdateOptions) (VirtualNetworkPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return VirtualNetworkPollerResponse{}, err
	}
	result := VirtualNetworkPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworksClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkPollerResponse{}, err
	}
	poller := &virtualNetworkPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualNetworkPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkPoller.ResumeToken().
func (client VirtualNetworksClient) ResumeCreateOrUpdate(token string) (VirtualNetworkPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworksClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a virtual network in the specified resource group.
func (client VirtualNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, parameters, options)
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
func (client VirtualNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, options *VirtualNetworksBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client VirtualNetworksClient) createOrUpdateHandleResponse(resp *azcore.Response) (VirtualNetworkResponse, error) {
	result := VirtualNetworkResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualNetwork)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VirtualNetworksClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified virtual network.
func (client VirtualNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworksClient.Delete", "location", resp, client.deleteHandleError)
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
func (client VirtualNetworksClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworksClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified virtual network.
func (client VirtualNetworksClient) delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
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
func (client VirtualNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client VirtualNetworksClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified virtual network by resource group.
func (client VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksGetOptions) (VirtualNetworkResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworkResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualNetworkResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VirtualNetworkResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VirtualNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client VirtualNetworksClient) getHandleResponse(resp *azcore.Response) (VirtualNetworkResponse, error) {
	result := VirtualNetworkResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualNetwork)
	return result, err
}

// getHandleError handles the Get error response.
func (client VirtualNetworksClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all virtual networks in a resource group.
func (client VirtualNetworksClient) List(resourceGroupName string, options *VirtualNetworksListOptions) VirtualNetworkListResultPager {
	return &virtualNetworkListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client VirtualNetworksClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client VirtualNetworksClient) listHandleResponse(resp *azcore.Response) (VirtualNetworkListResultResponse, error) {
	result := VirtualNetworkListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualNetworkListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client VirtualNetworksClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all virtual networks in a subscription.
func (client VirtualNetworksClient) ListAll(options *VirtualNetworksListAllOptions) VirtualNetworkListResultPager {
	return &virtualNetworkListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client VirtualNetworksClient) listAllCreateRequest(ctx context.Context, options *VirtualNetworksListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks"
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

// listAllHandleResponse handles the ListAll response.
func (client VirtualNetworksClient) listAllHandleResponse(resp *azcore.Response) (VirtualNetworkListResultResponse, error) {
	result := VirtualNetworkListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualNetworkListResult)
	return result, err
}

// listAllHandleError handles the ListAll error response.
func (client VirtualNetworksClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListUsage - Lists usage stats.
func (client VirtualNetworksClient) ListUsage(resourceGroupName string, virtualNetworkName string, options *VirtualNetworksListUsageOptions) VirtualNetworkListUsageResultPager {
	return &virtualNetworkListUsageResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listUsageCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
		},
		responder: client.listUsageHandleResponse,
		errorer:   client.listUsageHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkListUsageResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkListUsageResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listUsageCreateRequest creates the ListUsage request.
func (client VirtualNetworksClient) listUsageCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksListUsageOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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

// listUsageHandleResponse handles the ListUsage response.
func (client VirtualNetworksClient) listUsageHandleResponse(resp *azcore.Response) (VirtualNetworkListUsageResultResponse, error) {
	result := VirtualNetworkListUsageResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualNetworkListUsageResult)
	return result, err
}

// listUsageHandleError handles the ListUsage error response.
func (client VirtualNetworksClient) listUsageHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a virtual network tags.
func (client VirtualNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksUpdateTagsOptions) (VirtualNetworkResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualNetworkName, parameters, options)
	if err != nil {
		return VirtualNetworkResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualNetworkResponse{}, client.updateTagsHandleError(resp)
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return VirtualNetworkResponse{}, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client VirtualNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters TagsObject, options *VirtualNetworksUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// updateTagsHandleResponse handles the UpdateTags response.
func (client VirtualNetworksClient) updateTagsHandleResponse(resp *azcore.Response) (VirtualNetworkResponse, error) {
	result := VirtualNetworkResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualNetwork)
	return result, err
}

// updateTagsHandleError handles the UpdateTags error response.
func (client VirtualNetworksClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
