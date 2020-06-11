// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DisksOperations contains the methods for the Disks group.
type DisksOperations interface {
	// BeginCreateOrUpdate - Creates or updates a disk.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk) (*DiskResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (DiskPoller, error)
	// BeginDelete - Deletes a disk.
	BeginDelete(ctx context.Context, resourceGroupName string, diskName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about a disk.
	Get(ctx context.Context, resourceGroupName string, diskName string) (*DiskResponse, error)
	// BeginGrantAccess - Grants access to a disk.
	BeginGrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData) (*AccessURIResponse, error)
	// ResumeGrantAccess - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeGrantAccess(token string) (AccessURIPoller, error)
	// List - Lists all the disks under a subscription.
	List() (DiskListPager, error)
	// ListByResourceGroup - Lists all the disks under a resource group.
	ListByResourceGroup(resourceGroupName string) (DiskListPager, error)
	// BeginRevokeAccess - Revokes access to a disk.
	BeginRevokeAccess(ctx context.Context, resourceGroupName string, diskName string) (*HTTPResponse, error)
	// ResumeRevokeAccess - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeRevokeAccess(token string) (HTTPPoller, error)
	// BeginUpdate - Updates (patches) a disk.
	BeginUpdate(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate) (*DiskResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (DiskPoller, error)
}

// disksOperations implements the DisksOperations interface.
type disksOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a disk.
func (client *disksOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk) (*DiskResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, diskName, disk)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("disksOperations.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &diskPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*DiskResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *disksOperations) ResumeCreateOrUpdate(token string) (DiskPoller, error) {
	pt, err := resumePollingTracker("disksOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &diskPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *disksOperations) createOrUpdateCreateRequest(resourceGroupName string, diskName string, disk Disk) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(disk)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *disksOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*DiskResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := DiskResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Disk)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *disksOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// Delete - Deletes a disk.
func (client *disksOperations) BeginDelete(ctx context.Context, resourceGroupName string, diskName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, diskName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("disksOperations.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *disksOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("disksOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *disksOperations) deleteCreateRequest(resourceGroupName string, diskName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *disksOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *disksOperations) deleteHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// Get - Gets information about a disk.
func (client *disksOperations) Get(ctx context.Context, resourceGroupName string, diskName string) (*DiskResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, diskName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *disksOperations) getCreateRequest(resourceGroupName string, diskName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *disksOperations) getHandleResponse(resp *azcore.Response) (*DiskResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := DiskResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Disk)
}

// getHandleError handles the Get error response.
func (client *disksOperations) getHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// GrantAccess - Grants access to a disk.
func (client *disksOperations) BeginGrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData) (*AccessURIResponse, error) {
	req, err := client.grantAccessCreateRequest(resourceGroupName, diskName, grantAccessData)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.grantAccessHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("disksOperations.GrantAccess", "location", resp, client.grantAccessHandleError)
	if err != nil {
		return nil, err
	}
	poller := &accessUriPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*AccessURIResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *disksOperations) ResumeGrantAccess(token string) (AccessURIPoller, error) {
	pt, err := resumePollingTracker("disksOperations.GrantAccess", token, client.grantAccessHandleError)
	if err != nil {
		return nil, err
	}
	return &accessUriPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *disksOperations) grantAccessCreateRequest(resourceGroupName string, diskName string, grantAccessData GrantAccessData) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/beginGetAccess"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(grantAccessData)
}

// grantAccessHandleResponse handles the GrantAccess response.
func (client *disksOperations) grantAccessHandleResponse(resp *azcore.Response) (*AccessURIResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.grantAccessHandleError(resp)
	}
	result := AccessURIResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AccessURI)
}

// grantAccessHandleError handles the GrantAccess error response.
func (client *disksOperations) grantAccessHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// List - Lists all the disks under a subscription.
func (client *disksOperations) List() (DiskListPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &diskListPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *DiskListResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.DiskList.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.DiskList.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *disksOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *disksOperations) listHandleResponse(resp *azcore.Response) (*DiskListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := DiskListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskList)
}

// listHandleError handles the List error response.
func (client *disksOperations) listHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// ListByResourceGroup - Lists all the disks under a resource group.
func (client *disksOperations) ListByResourceGroup(resourceGroupName string) (DiskListPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &diskListPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *DiskListResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.DiskList.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.DiskList.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *disksOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *disksOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*DiskListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := DiskListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiskList)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *disksOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// RevokeAccess - Revokes access to a disk.
func (client *disksOperations) BeginRevokeAccess(ctx context.Context, resourceGroupName string, diskName string) (*HTTPResponse, error) {
	req, err := client.revokeAccessCreateRequest(resourceGroupName, diskName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.revokeAccessHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("disksOperations.RevokeAccess", "location", resp, client.revokeAccessHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *disksOperations) ResumeRevokeAccess(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("disksOperations.RevokeAccess", token, client.revokeAccessHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *disksOperations) revokeAccessCreateRequest(resourceGroupName string, diskName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/endGetAccess"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// revokeAccessHandleResponse handles the RevokeAccess response.
func (client *disksOperations) revokeAccessHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.revokeAccessHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// revokeAccessHandleError handles the RevokeAccess error response.
func (client *disksOperations) revokeAccessHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// Update - Updates (patches) a disk.
func (client *disksOperations) BeginUpdate(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate) (*DiskResponse, error) {
	req, err := client.updateCreateRequest(resourceGroupName, diskName, disk)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("disksOperations.Update", "", resp, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &diskPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*DiskResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *disksOperations) ResumeUpdate(token string) (DiskPoller, error) {
	pt, err := resumePollingTracker("disksOperations.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &diskPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// updateCreateRequest creates the Update request.
func (client *disksOperations) updateCreateRequest(resourceGroupName string, diskName string, disk DiskUpdate) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-11-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(disk)
}

// updateHandleResponse handles the Update response.
func (client *disksOperations) updateHandleResponse(resp *azcore.Response) (*DiskResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.updateHandleError(resp)
	}
	result := DiskResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Disk)
}

// updateHandleError handles the Update error response.
func (client *disksOperations) updateHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}
