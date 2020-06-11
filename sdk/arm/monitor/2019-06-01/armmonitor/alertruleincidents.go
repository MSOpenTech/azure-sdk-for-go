// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AlertRuleIncidentsOperations contains the methods for the AlertRuleIncidents group.
type AlertRuleIncidentsOperations interface {
	// Get - Gets an incident associated to an alert rule
	Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string) (*IncidentResponse, error)
	// ListByAlertRule - Gets a list of incidents associated to an alert rule
	ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string) (*IncidentListResultResponse, error)
}

// alertRuleIncidentsOperations implements the AlertRuleIncidentsOperations interface.
type alertRuleIncidentsOperations struct {
	*Client
	subscriptionID string
}

// Get - Gets an incident associated to an alert rule
func (client *alertRuleIncidentsOperations) Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string) (*IncidentResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, ruleName, incidentName)
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
func (client *alertRuleIncidentsOperations) getCreateRequest(resourceGroupName string, ruleName string, incidentName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents/{incidentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{incidentName}", url.PathEscape(incidentName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2016-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *alertRuleIncidentsOperations) getHandleResponse(resp *azcore.Response) (*IncidentResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := IncidentResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Incident)
}

// getHandleError handles the Get error response.
func (client *alertRuleIncidentsOperations) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByAlertRule - Gets a list of incidents associated to an alert rule
func (client *alertRuleIncidentsOperations) ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string) (*IncidentListResultResponse, error) {
	req, err := client.listByAlertRuleCreateRequest(resourceGroupName, ruleName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listByAlertRuleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listByAlertRuleCreateRequest creates the ListByAlertRule request.
func (client *alertRuleIncidentsOperations) listByAlertRuleCreateRequest(resourceGroupName string, ruleName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2016-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByAlertRuleHandleResponse handles the ListByAlertRule response.
func (client *alertRuleIncidentsOperations) listByAlertRuleHandleResponse(resp *azcore.Response) (*IncidentListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByAlertRuleHandleError(resp)
	}
	result := IncidentListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.IncidentListResult)
}

// listByAlertRuleHandleError handles the ListByAlertRule error response.
func (client *alertRuleIncidentsOperations) listByAlertRuleHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}
