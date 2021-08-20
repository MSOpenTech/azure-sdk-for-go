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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
func NewUsagesClient(con *armcore.Connection, subscriptionID string) *UsagesClient {
	return &UsagesClient{con: con, subscriptionID: subscriptionID}
}

// ListByAutomationAccount - Retrieve the usage for the account id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsagesClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, options *UsagesListByAutomationAccountOptions) (UsagesListByAutomationAccountResponse, error) {
	req, err := client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return UsagesListByAutomationAccountResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return UsagesListByAutomationAccountResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return UsagesListByAutomationAccountResponse{}, client.listByAutomationAccountHandleError(resp)
	}
	return client.listByAutomationAccountHandleResponse(resp)
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *UsagesClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *UsagesListByAutomationAccountOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
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
	reqQP.Set("api-version", "2020-01-13-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *UsagesClient) listByAutomationAccountHandleResponse(resp *azcore.Response) (UsagesListByAutomationAccountResponse, error) {
	result := UsagesListByAutomationAccountResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.UsageListResult); err != nil {
		return UsagesListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *UsagesClient) listByAutomationAccountHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
