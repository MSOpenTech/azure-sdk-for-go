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

// FirewallPoliciesOperations contains the methods for the FirewallPolicies group.
type FirewallPoliciesOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Firewall Policy.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy) (*FirewallPolicyResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FirewallPolicyPoller, error)
	// BeginDelete - Deletes the specified Firewall Policy.
	BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Firewall Policy.
	Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, firewallPoliciesGetOptions *FirewallPoliciesGetOptions) (*FirewallPolicyResponse, error)
	// List - Lists all Firewall Policies in a resource group.
	List(resourceGroupName string) (FirewallPolicyListResultPager, error)
	// ListAll - Gets all the Firewall Policies in a subscription.
	ListAll() (FirewallPolicyListResultPager, error)
}

// firewallPoliciesOperations implements the FirewallPoliciesOperations interface.
type firewallPoliciesOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates the specified Firewall Policy.
func (client *firewallPoliciesOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy) (*FirewallPolicyResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, firewallPolicyName, parameters)
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
	pt, err := createPollingTracker("firewallPoliciesOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &firewallPolicyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FirewallPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *firewallPoliciesOperations) ResumeCreateOrUpdate(token string) (FirewallPolicyPoller, error) {
	pt, err := resumePollingTracker("firewallPoliciesOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *firewallPoliciesOperations) createOrUpdateCreateRequest(resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *firewallPoliciesOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*FirewallPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := FirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicy)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *firewallPoliciesOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified Firewall Policy.
func (client *firewallPoliciesOperations) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, firewallPolicyName)
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
	pt, err := createPollingTracker("firewallPoliciesOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *firewallPoliciesOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("firewallPoliciesOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *firewallPoliciesOperations) deleteCreateRequest(resourceGroupName string, firewallPolicyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
func (client *firewallPoliciesOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *firewallPoliciesOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified Firewall Policy.
func (client *firewallPoliciesOperations) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, firewallPoliciesGetOptions *FirewallPoliciesGetOptions) (*FirewallPolicyResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, firewallPolicyName, firewallPoliciesGetOptions)
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
func (client *firewallPoliciesOperations) getCreateRequest(resourceGroupName string, firewallPolicyName string, firewallPoliciesGetOptions *FirewallPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if firewallPoliciesGetOptions != nil && firewallPoliciesGetOptions.Expand != nil {
		query.Set("$expand", *firewallPoliciesGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *firewallPoliciesOperations) getHandleResponse(resp *azcore.Response) (*FirewallPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := FirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicy)
}

// getHandleError handles the Get error response.
func (client *firewallPoliciesOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all Firewall Policies in a resource group.
func (client *firewallPoliciesOperations) List(resourceGroupName string) (FirewallPolicyListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *FirewallPolicyListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.FirewallPolicyListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.FirewallPolicyListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *firewallPoliciesOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *firewallPoliciesOperations) listHandleResponse(resp *azcore.Response) (*FirewallPolicyListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := FirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyListResult)
}

// listHandleError handles the List error response.
func (client *firewallPoliciesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the Firewall Policies in a subscription.
func (client *firewallPoliciesOperations) ListAll() (FirewallPolicyListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &firewallPolicyListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *FirewallPolicyListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.FirewallPolicyListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.FirewallPolicyListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *firewallPoliciesOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/firewallPolicies"
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

// listAllHandleResponse handles the ListAll response.
func (client *firewallPoliciesOperations) listAllHandleResponse(resp *azcore.Response) (*FirewallPolicyListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := FirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *firewallPoliciesOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
