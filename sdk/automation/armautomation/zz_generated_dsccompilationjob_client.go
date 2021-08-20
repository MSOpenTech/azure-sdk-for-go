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
	"time"
)

// DscCompilationJobClient contains the methods for the DscCompilationJob group.
// Don't use this type directly, use NewDscCompilationJobClient() instead.
type DscCompilationJobClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDscCompilationJobClient creates a new instance of DscCompilationJobClient with the specified values.
func NewDscCompilationJobClient(con *armcore.Connection, subscriptionID string) *DscCompilationJobClient {
	return &DscCompilationJobClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreate - Creates the Dsc compilation job of the configuration.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscCompilationJobClient) BeginCreate(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobBeginCreateOptions) (DscCompilationJobCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, automationAccountName, compilationJobName, parameters, options)
	if err != nil {
		return DscCompilationJobCreatePollerResponse{}, err
	}
	result := DscCompilationJobCreatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("DscCompilationJobClient.Create", "", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return DscCompilationJobCreatePollerResponse{}, err
	}
	poller := &dscCompilationJobCreatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DscCompilationJobCreateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new DscCompilationJobCreatePoller from the specified resume token.
// token - The value must come from a previous call to DscCompilationJobCreatePoller.ResumeToken().
func (client *DscCompilationJobClient) ResumeCreate(ctx context.Context, token string) (DscCompilationJobCreatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("DscCompilationJobClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return DscCompilationJobCreatePollerResponse{}, err
	}
	poller := &dscCompilationJobCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DscCompilationJobCreatePollerResponse{}, err
	}
	result := DscCompilationJobCreatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DscCompilationJobCreateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Creates the Dsc compilation job of the configuration.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscCompilationJobClient) create(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, compilationJobName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *DscCompilationJobClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if compilationJobName == "" {
		return nil, errors.New("parameter compilationJobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{compilationJobName}", url.PathEscape(compilationJobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createHandleError handles the Create error response.
func (client *DscCompilationJobClient) createHandleError(resp *azcore.Response) error {
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

// Get - Retrieve the Dsc configuration compilation job identified by job id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscCompilationJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *DscCompilationJobGetOptions) (DscCompilationJobGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, compilationJobName, options)
	if err != nil {
		return DscCompilationJobGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DscCompilationJobGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DscCompilationJobGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscCompilationJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *DscCompilationJobGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if compilationJobName == "" {
		return nil, errors.New("parameter compilationJobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{compilationJobName}", url.PathEscape(compilationJobName))
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

// getHandleResponse handles the Get response.
func (client *DscCompilationJobClient) getHandleResponse(resp *azcore.Response) (DscCompilationJobGetResponse, error) {
	result := DscCompilationJobGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DscCompilationJob); err != nil {
		return DscCompilationJobGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DscCompilationJobClient) getHandleError(resp *azcore.Response) error {
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

// GetStream - Retrieve the job stream identified by job stream id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscCompilationJobClient) GetStream(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *DscCompilationJobGetStreamOptions) (DscCompilationJobGetStreamResponse, error) {
	req, err := client.getStreamCreateRequest(ctx, resourceGroupName, automationAccountName, jobID, jobStreamID, options)
	if err != nil {
		return DscCompilationJobGetStreamResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DscCompilationJobGetStreamResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DscCompilationJobGetStreamResponse{}, client.getStreamHandleError(resp)
	}
	return client.getStreamHandleResponse(resp)
}

// getStreamCreateRequest creates the GetStream request.
func (client *DscCompilationJobClient) getStreamCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *DscCompilationJobGetStreamOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams/{jobStreamId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if jobStreamID == "" {
		return nil, errors.New("parameter jobStreamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobStreamId}", url.PathEscape(jobStreamID))
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

// getStreamHandleResponse handles the GetStream response.
func (client *DscCompilationJobClient) getStreamHandleResponse(resp *azcore.Response) (DscCompilationJobGetStreamResponse, error) {
	result := DscCompilationJobGetStreamResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.JobStream); err != nil {
		return DscCompilationJobGetStreamResponse{}, err
	}
	return result, nil
}

// getStreamHandleError handles the GetStream error response.
func (client *DscCompilationJobClient) getStreamHandleError(resp *azcore.Response) error {
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

// ListByAutomationAccount - Retrieve a list of dsc compilation jobs.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscCompilationJobClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *DscCompilationJobListByAutomationAccountOptions) DscCompilationJobListByAutomationAccountPager {
	return &dscCompilationJobListByAutomationAccountPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
		},
		advancer: func(ctx context.Context, resp DscCompilationJobListByAutomationAccountResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DscCompilationJobListResult.NextLink)
		},
	}
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscCompilationJobClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscCompilationJobListByAutomationAccountOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscCompilationJobClient) listByAutomationAccountHandleResponse(resp *azcore.Response) (DscCompilationJobListByAutomationAccountResponse, error) {
	result := DscCompilationJobListByAutomationAccountResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DscCompilationJobListResult); err != nil {
		return DscCompilationJobListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *DscCompilationJobClient) listByAutomationAccountHandleError(resp *azcore.Response) error {
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
