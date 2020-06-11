// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ExpressRouteCrossConnectionPeeringsOperations contains the methods for the ExpressRouteCrossConnectionPeerings group.
type ExpressRouteCrossConnectionPeeringsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering) (*ExpressRouteCrossConnectionPeeringResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPeeringPoller, error)
	// BeginDelete - Deletes the specified peering from the ExpressRouteCrossConnection.
	BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified peering for the ExpressRouteCrossConnection.
	Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (*ExpressRouteCrossConnectionPeeringResponse, error)
	// List - Gets all peerings in a specified ExpressRouteCrossConnection.
	List(resourceGroupName string, crossConnectionName string) (ExpressRouteCrossConnectionPeeringListPager, error)
}

// expressRouteCrossConnectionPeeringsOperations implements the ExpressRouteCrossConnectionPeeringsOperations interface.
type expressRouteCrossConnectionPeeringsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a peering in the specified ExpressRouteCrossConnection.
func (client *expressRouteCrossConnectionPeeringsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, crossConnectionName, peeringName, peeringParameters)
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
	pt, err := createPollingTracker("expressRouteCrossConnectionPeeringsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionPeeringPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *expressRouteCrossConnectionPeeringsOperations) ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPeeringPoller, error) {
	pt, err := resumePollingTracker("expressRouteCrossConnectionPeeringsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPeeringPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *expressRouteCrossConnectionPeeringsOperations) createOrUpdateCreateRequest(resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(peeringParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *expressRouteCrossConnectionPeeringsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := ExpressRouteCrossConnectionPeeringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeering)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *expressRouteCrossConnectionPeeringsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified peering from the ExpressRouteCrossConnection.
func (client *expressRouteCrossConnectionPeeringsOperations) BeginDelete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, crossConnectionName, peeringName)
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
	pt, err := createPollingTracker("expressRouteCrossConnectionPeeringsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *expressRouteCrossConnectionPeeringsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("expressRouteCrossConnectionPeeringsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *expressRouteCrossConnectionPeeringsOperations) deleteCreateRequest(resourceGroupName string, crossConnectionName string, peeringName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *expressRouteCrossConnectionPeeringsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *expressRouteCrossConnectionPeeringsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified peering for the ExpressRouteCrossConnection.
func (client *expressRouteCrossConnectionPeeringsOperations) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, crossConnectionName, peeringName)
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
func (client *expressRouteCrossConnectionPeeringsOperations) getCreateRequest(resourceGroupName string, crossConnectionName string, peeringName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *expressRouteCrossConnectionPeeringsOperations) getHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPeeringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ExpressRouteCrossConnectionPeeringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeering)
}

// getHandleError handles the Get error response.
func (client *expressRouteCrossConnectionPeeringsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all peerings in a specified ExpressRouteCrossConnection.
func (client *expressRouteCrossConnectionPeeringsOperations) List(resourceGroupName string, crossConnectionName string) (ExpressRouteCrossConnectionPeeringListPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, crossConnectionName)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPeeringListPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ExpressRouteCrossConnectionPeeringListResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRouteCrossConnectionPeeringList.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRouteCrossConnectionPeeringList.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *expressRouteCrossConnectionPeeringsOperations) listCreateRequest(resourceGroupName string, crossConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *expressRouteCrossConnectionPeeringsOperations) listHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionPeeringListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ExpressRouteCrossConnectionPeeringListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionPeeringList)
}

// listHandleError handles the List error response.
func (client *expressRouteCrossConnectionPeeringsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
