//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
func NewPrivateLinkResourcesClient(con *armcore.Connection, subscriptionID string) *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{con: con, subscriptionID: subscriptionID}
}

// Automation - Gets the private link resources that need to be created for Automation account.
// If the operation fails it returns a generic error.
func (client *PrivateLinkResourcesClient) Automation(ctx context.Context, resourceGroupName string, automationAccountName string, options *PrivateLinkResourcesAutomationOptions) (PrivateLinkResourcesAutomationResponse, error) {
	req, err := client.automationCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return PrivateLinkResourcesAutomationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesAutomationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateLinkResourcesAutomationResponse{}, client.automationHandleError(resp)
	}
	return client.automationHandleResponse(resp)
}

// automationCreateRequest creates the Automation request.
func (client *PrivateLinkResourcesClient) automationCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *PrivateLinkResourcesAutomationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// automationHandleResponse handles the Automation response.
func (client *PrivateLinkResourcesClient) automationHandleResponse(resp *azcore.Response) (PrivateLinkResourcesAutomationResponse, error) {
	result := PrivateLinkResourcesAutomationResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesAutomationResponse{}, err
	}
	return result, nil
}

// automationHandleError handles the Automation error response.
func (client *PrivateLinkResourcesClient) automationHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
