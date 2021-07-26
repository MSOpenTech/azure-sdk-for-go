// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineSizesClient contains the methods for the VirtualMachineSizes group.
// Don't use this type directly, use NewVirtualMachineSizesClient() instead.
type VirtualMachineSizesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient with the specified values.
func NewVirtualMachineSizesClient(con *armcore.Connection, subscriptionID string) *VirtualMachineSizesClient {
	return &VirtualMachineSizesClient{con: con, subscriptionID: subscriptionID}
}

// List - This API is deprecated. Use Resources Skus [https://docs.microsoft.com/rest/api/compute/resourceskus/list]
// If the operation fails it returns a generic error.
func (client *VirtualMachineSizesClient) List(ctx context.Context, location string, options *VirtualMachineSizesListOptions) (VirtualMachineSizesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineSizesListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineSizesListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineSizesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineSizesClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineSizesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineSizesClient) listHandleResponse(resp *azcore.Response) (VirtualMachineSizesListResponse, error) {
	result := VirtualMachineSizesListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.VirtualMachineSizeListResult); err != nil {
		return VirtualMachineSizesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineSizesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
