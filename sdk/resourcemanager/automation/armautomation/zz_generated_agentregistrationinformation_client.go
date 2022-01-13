//go:build go1.16
// +build go1.16

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AgentRegistrationInformationClient contains the methods for the AgentRegistrationInformation group.
// Don't use this type directly, use NewAgentRegistrationInformationClient() instead.
type AgentRegistrationInformationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAgentRegistrationInformationClient creates a new instance of AgentRegistrationInformationClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAgentRegistrationInformationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AgentRegistrationInformationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AgentRegistrationInformationClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Retrieve the automation agent registration information.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - AgentRegistrationInformationClientGetOptions contains the optional parameters for the AgentRegistrationInformationClient.Get
// method.
func (client *AgentRegistrationInformationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, options *AgentRegistrationInformationClientGetOptions) (AgentRegistrationInformationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return AgentRegistrationInformationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentRegistrationInformationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentRegistrationInformationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentRegistrationInformationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *AgentRegistrationInformationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/agentRegistrationInformation"
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
func (client *AgentRegistrationInformationClient) getHandleResponse(resp *http.Response) (AgentRegistrationInformationClientGetResponse, error) {
	result := AgentRegistrationInformationClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentRegistration); err != nil {
		return AgentRegistrationInformationClientGetResponse{}, err
	}
	return result, nil
}

// RegenerateKey - Regenerate a primary or secondary agent registration key
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// parameters - The name of the agent registration key to be regenerated
// options - AgentRegistrationInformationClientRegenerateKeyOptions contains the optional parameters for the AgentRegistrationInformationClient.RegenerateKey
// method.
func (client *AgentRegistrationInformationClient) RegenerateKey(ctx context.Context, resourceGroupName string, automationAccountName string, parameters AgentRegistrationRegenerateKeyParameter, options *AgentRegistrationInformationClientRegenerateKeyOptions) (AgentRegistrationInformationClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, automationAccountName, parameters, options)
	if err != nil {
		return AgentRegistrationInformationClientRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentRegistrationInformationClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentRegistrationInformationClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *AgentRegistrationInformationClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, parameters AgentRegistrationRegenerateKeyParameter, options *AgentRegistrationInformationClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/agentRegistrationInformation/regenerateKey"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *AgentRegistrationInformationClient) regenerateKeyHandleResponse(resp *http.Response) (AgentRegistrationInformationClientRegenerateKeyResponse, error) {
	result := AgentRegistrationInformationClientRegenerateKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentRegistration); err != nil {
		return AgentRegistrationInformationClientRegenerateKeyResponse{}, err
	}
	return result, nil
}
