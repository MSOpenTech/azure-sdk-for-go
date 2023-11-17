//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azeventgrid

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/messaging"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// Client contains the methods for the Client group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	endpoint string
}

// AcknowledgeCloudEvents - Acknowledge batch of Cloud Events. The server responds with an HTTP 200 status code if the request
// is successfully accepted. The response body will include the set of successfully acknowledged
// lockTokens, along with other failed lockTokens with their corresponding error information. Successfully acknowledged events
// will no longer be available to any consumer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - eventSubscriptionName - Event Subscription Name.
//   - acknowledgeOptions - acknowledgeOptions.
//   - options - AcknowledgeCloudEventsOptions contains the optional parameters for the Client.AcknowledgeCloudEvents method.
func (client *Client) internalAcknowledgeCloudEvents(ctx context.Context, topicName string, eventSubscriptionName string, acknowledgeOptions acknowledgeOptions, options *AcknowledgeCloudEventsOptions) (AcknowledgeCloudEventsResponse, error) {
	var err error
	req, err := client.acknowledgeCloudEventsCreateRequest(ctx, topicName, eventSubscriptionName, acknowledgeOptions, options)
	if err != nil {
		return AcknowledgeCloudEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AcknowledgeCloudEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AcknowledgeCloudEventsResponse{}, err
	}
	resp, err := client.acknowledgeCloudEventsHandleResponse(httpResp)
	return resp, err
}

// acknowledgeCloudEventsCreateRequest creates the AcknowledgeCloudEvents request.
func (client *Client) acknowledgeCloudEventsCreateRequest(ctx context.Context, topicName string, eventSubscriptionName string, acknowledgeOptions acknowledgeOptions, options *AcknowledgeCloudEventsOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}/eventsubscriptions/{eventSubscriptionName}:acknowledge"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, acknowledgeOptions); err != nil {
		return nil, err
	}
	return req, nil
}

// acknowledgeCloudEventsHandleResponse handles the AcknowledgeCloudEvents response.
func (client *Client) acknowledgeCloudEventsHandleResponse(resp *http.Response) (AcknowledgeCloudEventsResponse, error) {
	result := AcknowledgeCloudEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AcknowledgeResult); err != nil {
		return AcknowledgeCloudEventsResponse{}, err
	}
	return result, nil
}

// PublishCloudEvent - Publish Single Cloud Event to namespace topic. In case of success, the server responds with an HTTP
// 200 status code with an empty JSON object in response. Otherwise, the server can return various
// error codes. For example, 401: which indicates authorization failure, 403: which indicates quota exceeded or message is
// too large, 410: which indicates that specific topic is not found, 400: for bad
// request, and 500: for internal server error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - event - Single Cloud Event being published.
//   - options - PublishCloudEventOptions contains the optional parameters for the Client.PublishCloudEvent method.
func (client *Client) PublishCloudEvent(ctx context.Context, topicName string, event messaging.CloudEvent, options *PublishCloudEventOptions) (PublishCloudEventResponse, error) {
	var err error
	req, err := client.publishCloudEventCreateRequest(ctx, topicName, event, options)
	if err != nil {
		return PublishCloudEventResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PublishCloudEventResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PublishCloudEventResponse{}, err
	}
	return PublishCloudEventResponse{}, nil
}

// publishCloudEventCreateRequest creates the PublishCloudEvent request.
func (client *Client) publishCloudEventCreateRequest(ctx context.Context, topicName string, event messaging.CloudEvent, options *PublishCloudEventOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}:publish"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, event); err != nil {
		return nil, err
	}

	req.Raw().Header.Set("Content-type", "application/cloudevents+json; charset=utf-8")
	return req, nil
}

// PublishCloudEvents - Publish Batch Cloud Event to namespace topic. In case of success, the server responds with an HTTP
// 200 status code with an empty JSON object in response. Otherwise, the server can return various error
// codes. For example, 401: which indicates authorization failure, 403: which indicates quota exceeded or message is too large,
// 410: which indicates that specific topic is not found, 400: for bad
// request, and 500: for internal server error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - events - Array of Cloud Events being published.
//   - options - PublishCloudEventsOptions contains the optional parameters for the Client.PublishCloudEvents method.
func (client *Client) PublishCloudEvents(ctx context.Context, topicName string, events []messaging.CloudEvent, options *PublishCloudEventsOptions) (PublishCloudEventsResponse, error) {
	var err error
	req, err := client.publishCloudEventsCreateRequest(ctx, topicName, events, options)
	if err != nil {
		return PublishCloudEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PublishCloudEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PublishCloudEventsResponse{}, err
	}
	return PublishCloudEventsResponse{}, nil
}

// publishCloudEventsCreateRequest creates the PublishCloudEvents request.
func (client *Client) publishCloudEventsCreateRequest(ctx context.Context, topicName string, events []messaging.CloudEvent, options *PublishCloudEventsOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}:publish"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, events); err != nil {
		return nil, err
	}

	req.Raw().Header.Set("Content-type", "application/cloudevents-batch+json; charset=utf-8")
	return req, nil
}

// ReceiveCloudEvents - Receive Batch of Cloud Events from the Event Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - eventSubscriptionName - Event Subscription Name.
//   - options - ReceiveCloudEventsOptions contains the optional parameters for the Client.ReceiveCloudEvents method.
func (client *Client) ReceiveCloudEvents(ctx context.Context, topicName string, eventSubscriptionName string, options *ReceiveCloudEventsOptions) (ReceiveCloudEventsResponse, error) {
	var err error
	req, err := client.receiveCloudEventsCreateRequest(ctx, topicName, eventSubscriptionName, options)
	if err != nil {
		return ReceiveCloudEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReceiveCloudEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReceiveCloudEventsResponse{}, err
	}
	resp, err := client.receiveCloudEventsHandleResponse(httpResp)
	return resp, err
}

// receiveCloudEventsCreateRequest creates the ReceiveCloudEvents request.
func (client *Client) receiveCloudEventsCreateRequest(ctx context.Context, topicName string, eventSubscriptionName string, options *ReceiveCloudEventsOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}/eventsubscriptions/{eventSubscriptionName}:receive"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.MaxEvents != nil {
		reqQP.Set("maxEvents", strconv.FormatInt(int64(*options.MaxEvents), 10))
	}
	if options != nil && options.MaxWaitTime != nil {
		reqQP.Set("maxWaitTime", strconv.FormatInt(int64(*options.MaxWaitTime), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// receiveCloudEventsHandleResponse handles the ReceiveCloudEvents response.
func (client *Client) receiveCloudEventsHandleResponse(resp *http.Response) (ReceiveCloudEventsResponse, error) {
	result := ReceiveCloudEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReceiveResult); err != nil {
		return ReceiveCloudEventsResponse{}, err
	}
	return result, nil
}

// RejectCloudEvents - Reject batch of Cloud Events. The server responds with an HTTP 200 status code if the request is successfully
// accepted. The response body will include the set of successfully rejected lockTokens,
// along with other failed lockTokens with their corresponding error information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - eventSubscriptionName - Event Subscription Name.
//   - rejectOptions - rejectOptions
//   - options - RejectCloudEventsOptions contains the optional parameters for the Client.RejectCloudEvents method.
func (client *Client) internalRejectCloudEvents(ctx context.Context, topicName string, eventSubscriptionName string, rejectOptions rejectOptions, options *RejectCloudEventsOptions) (RejectCloudEventsResponse, error) {
	var err error
	req, err := client.rejectCloudEventsCreateRequest(ctx, topicName, eventSubscriptionName, rejectOptions, options)
	if err != nil {
		return RejectCloudEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RejectCloudEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RejectCloudEventsResponse{}, err
	}
	resp, err := client.rejectCloudEventsHandleResponse(httpResp)
	return resp, err
}

// rejectCloudEventsCreateRequest creates the RejectCloudEvents request.
func (client *Client) rejectCloudEventsCreateRequest(ctx context.Context, topicName string, eventSubscriptionName string, rejectOptions rejectOptions, options *RejectCloudEventsOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}/eventsubscriptions/{eventSubscriptionName}:reject"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, rejectOptions); err != nil {
		return nil, err
	}
	return req, nil
}

// rejectCloudEventsHandleResponse handles the RejectCloudEvents response.
func (client *Client) rejectCloudEventsHandleResponse(resp *http.Response) (RejectCloudEventsResponse, error) {
	result := RejectCloudEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RejectResult); err != nil {
		return RejectCloudEventsResponse{}, err
	}
	return result, nil
}

// ReleaseCloudEvents - Release batch of Cloud Events. The server responds with an HTTP 200 status code if the request is
// successfully accepted. The response body will include the set of successfully released lockTokens,
// along with other failed lockTokens with their corresponding error information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - eventSubscriptionName - Event Subscription Name.
//   - releaseOptions - releaseOptions
//   - options - ReleaseCloudEventsOptions contains the optional parameters for the Client.ReleaseCloudEvents method.
func (client *Client) internalReleaseCloudEvents(ctx context.Context, topicName string, eventSubscriptionName string, releaseOptions releaseOptions, options *ReleaseCloudEventsOptions) (ReleaseCloudEventsResponse, error) {
	var err error
	req, err := client.releaseCloudEventsCreateRequest(ctx, topicName, eventSubscriptionName, releaseOptions, options)
	if err != nil {
		return ReleaseCloudEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReleaseCloudEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReleaseCloudEventsResponse{}, err
	}
	resp, err := client.releaseCloudEventsHandleResponse(httpResp)
	return resp, err
}

// releaseCloudEventsCreateRequest creates the ReleaseCloudEvents request.
func (client *Client) releaseCloudEventsCreateRequest(ctx context.Context, topicName string, eventSubscriptionName string, releaseOptions releaseOptions, options *ReleaseCloudEventsOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}/eventsubscriptions/{eventSubscriptionName}:release"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.ReleaseDelayInSeconds != nil {
		reqQP.Set("releaseDelayInSeconds", fmt.Sprintf("%d", *options.ReleaseDelayInSeconds))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, releaseOptions); err != nil {
		return nil, err
	}
	return req, nil
}

// releaseCloudEventsHandleResponse handles the ReleaseCloudEvents response.
func (client *Client) releaseCloudEventsHandleResponse(resp *http.Response) (ReleaseCloudEventsResponse, error) {
	result := ReleaseCloudEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReleaseResult); err != nil {
		return ReleaseCloudEventsResponse{}, err
	}
	return result, nil
}

// RenewCloudEventLocks - Renew lock for batch of Cloud Events. The server responds with an HTTP 200 status code if the request
// is successfully accepted. The response body will include the set of successfully renewed
// lockTokens, along with other failed lockTokens with their corresponding error information.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - topicName - Topic Name.
//   - eventSubscriptionName - Event Subscription Name.
//   - renewLockOptions - RenewLockOptions
//   - options - RenewCloudEventLocksOptions contains the optional parameters for the Client.RenewCloudEventLocks method.
func (client *Client) RenewCloudEventLocks(ctx context.Context, topicName string, eventSubscriptionName string, renewLockOptions RenewLockOptions, options *RenewCloudEventLocksOptions) (RenewCloudEventLocksResponse, error) {
	var err error
	req, err := client.renewCloudEventLocksCreateRequest(ctx, topicName, eventSubscriptionName, renewLockOptions, options)
	if err != nil {
		return RenewCloudEventLocksResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RenewCloudEventLocksResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RenewCloudEventLocksResponse{}, err
	}
	resp, err := client.renewCloudEventLocksHandleResponse(httpResp)
	return resp, err
}

// renewCloudEventLocksCreateRequest creates the RenewCloudEventLocks request.
func (client *Client) renewCloudEventLocksCreateRequest(ctx context.Context, topicName string, eventSubscriptionName string, renewLockOptions RenewLockOptions, options *RenewCloudEventLocksOptions) (*policy.Request, error) {
	urlPath := "/topics/{topicName}/eventsubscriptions/{eventSubscriptionName}:renewLock"
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, renewLockOptions); err != nil {
		return nil, err
	}
	return req, nil
}

// renewCloudEventLocksHandleResponse handles the RenewCloudEventLocks response.
func (client *Client) renewCloudEventLocksHandleResponse(resp *http.Response) (RenewCloudEventLocksResponse, error) {
	result := RenewCloudEventLocksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RenewCloudEventLocksResult); err != nil {
		return RenewCloudEventLocksResponse{}, err
	}
	return result, nil
}
