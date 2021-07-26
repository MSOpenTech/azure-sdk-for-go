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
	"net/http"
	"net/url"
	"strings"
)

// DiskRestorePointClient contains the methods for the DiskRestorePoint group.
// Don't use this type directly, use NewDiskRestorePointClient() instead.
type DiskRestorePointClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDiskRestorePointClient creates a new instance of DiskRestorePointClient with the specified values.
func NewDiskRestorePointClient(con *armcore.Connection, subscriptionID string) *DiskRestorePointClient {
	return &DiskRestorePointClient{con: con, subscriptionID: subscriptionID}
}

// Get - Get disk restorePoint resource
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointGetOptions) (DiskRestorePointGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, diskRestorePointName, options)
	if err != nil {
		return DiskRestorePointGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiskRestorePointGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DiskRestorePointGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskRestorePointClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *DiskRestorePointGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	if diskRestorePointName == "" {
		return nil, errors.New("parameter diskRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskRestorePointName}", url.PathEscape(diskRestorePointName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskRestorePointClient) getHandleResponse(resp *azcore.Response) (DiskRestorePointGetResponse, error) {
	result := DiskRestorePointGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DiskRestorePoint); err != nil {
		return DiskRestorePointGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DiskRestorePointClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByRestorePoint - Lists diskRestorePoints under a vmRestorePoint.
// If the operation fails it returns the *CloudError error type.
func (client *DiskRestorePointClient) ListByRestorePoint(resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *DiskRestorePointListByRestorePointOptions) DiskRestorePointListByRestorePointPager {
	return &diskRestorePointListByRestorePointPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByRestorePointCreateRequest(ctx, resourceGroupName, restorePointCollectionName, vmRestorePointName, options)
		},
		advancer: func(ctx context.Context, resp DiskRestorePointListByRestorePointResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DiskRestorePointList.NextLink)
		},
	}
}

// listByRestorePointCreateRequest creates the ListByRestorePoint request.
func (client *DiskRestorePointClient) listByRestorePointCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *DiskRestorePointListByRestorePointOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if vmRestorePointName == "" {
		return nil, errors.New("parameter vmRestorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmRestorePointName}", url.PathEscape(vmRestorePointName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByRestorePointHandleResponse handles the ListByRestorePoint response.
func (client *DiskRestorePointClient) listByRestorePointHandleResponse(resp *azcore.Response) (DiskRestorePointListByRestorePointResponse, error) {
	result := DiskRestorePointListByRestorePointResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DiskRestorePointList); err != nil {
		return DiskRestorePointListByRestorePointResponse{}, err
	}
	return result, nil
}

// listByRestorePointHandleError handles the ListByRestorePoint error response.
func (client *DiskRestorePointClient) listByRestorePointHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
