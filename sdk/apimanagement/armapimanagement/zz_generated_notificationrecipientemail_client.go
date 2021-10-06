//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// NotificationRecipientEmailClient contains the methods for the NotificationRecipientEmail group.
// Don't use this type directly, use NewNotificationRecipientEmailClient() instead.
type NotificationRecipientEmailClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNotificationRecipientEmailClient creates a new instance of NotificationRecipientEmailClient with the specified values.
func NewNotificationRecipientEmailClient(con *arm.Connection, subscriptionID string) *NotificationRecipientEmailClient {
	return &NotificationRecipientEmailClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CheckEntityExists - Determine if Notification Recipient Email subscribed to the notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientEmailClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailCheckEntityExistsOptions) (NotificationRecipientEmailCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, notificationName, email, options)
	if err != nil {
		return NotificationRecipientEmailCheckEntityExistsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientEmailCheckEntityExistsResponse{}, err
	}
	result := NotificationRecipientEmailCheckEntityExistsResponse{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *NotificationRecipientEmailClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails/{email}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// CreateOrUpdate - Adds the Email address to the list of Recipients for the Notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientEmailClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailCreateOrUpdateOptions) (NotificationRecipientEmailCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, notificationName, email, options)
	if err != nil {
		return NotificationRecipientEmailCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientEmailCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return NotificationRecipientEmailCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationRecipientEmailClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails/{email}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationRecipientEmailClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationRecipientEmailCreateOrUpdateResponse, error) {
	result := NotificationRecipientEmailCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientEmailContract); err != nil {
		return NotificationRecipientEmailCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NotificationRecipientEmailClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Removes the email from the list of Notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientEmailClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailDeleteOptions) (NotificationRecipientEmailDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, notificationName, email, options)
	if err != nil {
		return NotificationRecipientEmailDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientEmailDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return NotificationRecipientEmailDeleteResponse{}, client.deleteHandleError(resp)
	}
	return NotificationRecipientEmailDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotificationRecipientEmailClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails/{email}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *NotificationRecipientEmailClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByNotification - Gets the list of the Notification Recipient Emails subscribed to a notification.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotificationRecipientEmailClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientEmailListByNotificationOptions) (NotificationRecipientEmailListByNotificationResponse, error) {
	req, err := client.listByNotificationCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationRecipientEmailListByNotificationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationRecipientEmailListByNotificationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationRecipientEmailListByNotificationResponse{}, client.listByNotificationHandleError(resp)
	}
	return client.listByNotificationHandleResponse(resp)
}

// listByNotificationCreateRequest creates the ListByNotification request.
func (client *NotificationRecipientEmailClient) listByNotificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientEmailListByNotificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNotificationHandleResponse handles the ListByNotification response.
func (client *NotificationRecipientEmailClient) listByNotificationHandleResponse(resp *http.Response) (NotificationRecipientEmailListByNotificationResponse, error) {
	result := NotificationRecipientEmailListByNotificationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientEmailCollection); err != nil {
		return NotificationRecipientEmailListByNotificationResponse{}, err
	}
	return result, nil
}

// listByNotificationHandleError handles the ListByNotification error response.
func (client *NotificationRecipientEmailClient) listByNotificationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
