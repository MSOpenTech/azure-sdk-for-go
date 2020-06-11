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

// FirewallPolicyRuleGroupsOperations contains the methods for the FirewallPolicyRuleGroups group.
type FirewallPolicyRuleGroupsOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup) (*FirewallPolicyRuleGroupResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FirewallPolicyRuleGroupPoller, error)
	// BeginDelete - Deletes the specified FirewallPolicyRuleGroup.
	BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified FirewallPolicyRuleGroup.
	Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*FirewallPolicyRuleGroupResponse, error)
	// List - Lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
	List(resourceGroupName string, firewallPolicyName string) (FirewallPolicyRuleGroupListResultPager, error)
}

// firewallPolicyRuleGroupsOperations implements the FirewallPolicyRuleGroupsOperations interface.
type firewallPolicyRuleGroupsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleGroup.
func (client *firewallPolicyRuleGroupsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup) (*FirewallPolicyRuleGroupResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, firewallPolicyName, ruleGroupName, parameters)
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
	pt, err := createPollingTracker("firewallPolicyRuleGroupsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &firewallPolicyRuleGroupPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FirewallPolicyRuleGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *firewallPolicyRuleGroupsOperations) ResumeCreateOrUpdate(token string) (FirewallPolicyRuleGroupPoller, error) {
	pt, err := resumePollingTracker("firewallPolicyRuleGroupsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyRuleGroupPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *firewallPolicyRuleGroupsOperations) createOrUpdateCreateRequest(resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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
func (client *firewallPolicyRuleGroupsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := FirewallPolicyRuleGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroup)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *firewallPolicyRuleGroupsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified FirewallPolicyRuleGroup.
func (client *firewallPolicyRuleGroupsOperations) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, firewallPolicyName, ruleGroupName)
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
	pt, err := createPollingTracker("firewallPolicyRuleGroupsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *firewallPolicyRuleGroupsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("firewallPolicyRuleGroupsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *firewallPolicyRuleGroupsOperations) deleteCreateRequest(resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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
func (client *firewallPolicyRuleGroupsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *firewallPolicyRuleGroupsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified FirewallPolicyRuleGroup.
func (client *firewallPolicyRuleGroupsOperations) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*FirewallPolicyRuleGroupResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, firewallPolicyName, ruleGroupName)
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
func (client *firewallPolicyRuleGroupsOperations) getCreateRequest(resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
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
func (client *firewallPolicyRuleGroupsOperations) getHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := FirewallPolicyRuleGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroup)
}

// getHandleError handles the Get error response.
func (client *firewallPolicyRuleGroupsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
func (client *firewallPolicyRuleGroupsOperations) List(resourceGroupName string, firewallPolicyName string) (FirewallPolicyRuleGroupListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, firewallPolicyName)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyRuleGroupListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *FirewallPolicyRuleGroupListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.FirewallPolicyRuleGroupListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.FirewallPolicyRuleGroupListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *firewallPolicyRuleGroupsOperations) listCreateRequest(resourceGroupName string, firewallPolicyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups"
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
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *firewallPolicyRuleGroupsOperations) listHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleGroupListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := FirewallPolicyRuleGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleGroupListResult)
}

// listHandleError handles the List error response.
func (client *firewallPolicyRuleGroupsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
