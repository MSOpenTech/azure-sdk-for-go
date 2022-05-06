//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DscCompilationJobClient contains the methods for the DscCompilationJob group.
// Don't use this type directly, use NewDscCompilationJobClient() instead.
type DscCompilationJobClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDscCompilationJobClient creates a new instance of DscCompilationJobClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDscCompilationJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscCompilationJobClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DscCompilationJobClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates the Dsc compilation job of the configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// compilationJobName - The DSC configuration Id.
// parameters - The parameters supplied to the create compilation job operation.
// options - DscCompilationJobClientBeginCreateOptions contains the optional parameters for the DscCompilationJobClient.BeginCreate
// method.
func (client *DscCompilationJobClient) BeginCreate(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobClientBeginCreateOptions) (*armruntime.Poller[DscCompilationJobClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, automationAccountName, compilationJobName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DscCompilationJobClientCreateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DscCompilationJobClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates the Dsc compilation job of the configuration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DscCompilationJobClient) create(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, compilationJobName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *DscCompilationJobClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobClientBeginCreateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Retrieve the Dsc configuration compilation job identified by job id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// compilationJobName - The DSC configuration Id.
// options - DscCompilationJobClientGetOptions contains the optional parameters for the DscCompilationJobClient.Get method.
func (client *DscCompilationJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *DscCompilationJobClientGetOptions) (DscCompilationJobClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, compilationJobName, options)
	if err != nil {
		return DscCompilationJobClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscCompilationJobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscCompilationJobClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscCompilationJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *DscCompilationJobClientGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscCompilationJobClient) getHandleResponse(resp *http.Response) (DscCompilationJobClientGetResponse, error) {
	result := DscCompilationJobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscCompilationJob); err != nil {
		return DscCompilationJobClientGetResponse{}, err
	}
	return result, nil
}

// GetStream - Retrieve the job stream identified by job stream id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobID - The job id.
// jobStreamID - The job stream id.
// options - DscCompilationJobClientGetStreamOptions contains the optional parameters for the DscCompilationJobClient.GetStream
// method.
func (client *DscCompilationJobClient) GetStream(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *DscCompilationJobClientGetStreamOptions) (DscCompilationJobClientGetStreamResponse, error) {
	req, err := client.getStreamCreateRequest(ctx, resourceGroupName, automationAccountName, jobID, jobStreamID, options)
	if err != nil {
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscCompilationJobClientGetStreamResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStreamHandleResponse(resp)
}

// getStreamCreateRequest creates the GetStream request.
func (client *DscCompilationJobClient) getStreamCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *DscCompilationJobClientGetStreamOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStreamHandleResponse handles the GetStream response.
func (client *DscCompilationJobClient) getStreamHandleResponse(resp *http.Response) (DscCompilationJobClientGetStreamResponse, error) {
	result := DscCompilationJobClientGetStreamResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStream); err != nil {
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of dsc compilation jobs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - DscCompilationJobClientListByAutomationAccountOptions contains the optional parameters for the DscCompilationJobClient.ListByAutomationAccount
// method.
func (client *DscCompilationJobClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *DscCompilationJobClientListByAutomationAccountOptions) *runtime.Pager[DscCompilationJobClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PageProcessor[DscCompilationJobClientListByAutomationAccountResponse]{
		More: func(page DscCompilationJobClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscCompilationJobClientListByAutomationAccountResponse) (DscCompilationJobClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscCompilationJobClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DscCompilationJobClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscCompilationJobClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscCompilationJobClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscCompilationJobClientListByAutomationAccountOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscCompilationJobClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscCompilationJobClientListByAutomationAccountResponse, error) {
	result := DscCompilationJobClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscCompilationJobListResult); err != nil {
		return DscCompilationJobClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}
