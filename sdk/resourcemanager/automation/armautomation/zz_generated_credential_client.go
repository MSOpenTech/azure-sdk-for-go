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

// CredentialClient contains the methods for the Credential group.
// Don't use this type directly, use NewCredentialClient() instead.
type CredentialClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCredentialClient creates a new instance of CredentialClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCredentialClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CredentialClient, error) {
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
	client := &CredentialClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create a credential.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// credentialName - The parameters supplied to the create or update credential operation.
// parameters - The parameters supplied to the create or update credential operation.
// options - CredentialClientCreateOrUpdateOptions contains the optional parameters for the CredentialClient.CreateOrUpdate
// method.
func (client *CredentialClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, parameters CredentialCreateOrUpdateParameters, options *CredentialClientCreateOrUpdateOptions) (CredentialClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, credentialName, parameters, options)
	if err != nil {
		return CredentialClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CredentialClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CredentialClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, parameters CredentialCreateOrUpdateParameters, options *CredentialClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CredentialClient) createOrUpdateHandleResponse(resp *http.Response) (CredentialClientCreateOrUpdateResponse, error) {
	result := CredentialClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Credential); err != nil {
		return CredentialClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the credential.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// credentialName - The name of credential.
// options - CredentialClientDeleteOptions contains the optional parameters for the CredentialClient.Delete method.
func (client *CredentialClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, options *CredentialClientDeleteOptions) (CredentialClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, credentialName, options)
	if err != nil {
		return CredentialClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CredentialClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CredentialClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CredentialClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, options *CredentialClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieve the credential identified by credential name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// credentialName - The name of credential.
// options - CredentialClientGetOptions contains the optional parameters for the CredentialClient.Get method.
func (client *CredentialClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, options *CredentialClientGetOptions) (CredentialClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, credentialName, options)
	if err != nil {
		return CredentialClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CredentialClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CredentialClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, options *CredentialClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
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
func (client *CredentialClient) getHandleResponse(resp *http.Response) (CredentialClientGetResponse, error) {
	result := CredentialClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Credential); err != nil {
		return CredentialClientGetResponse{}, err
	}
	return result, nil
}

// ListByAutomationAccount - Retrieve a list of credentials.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - CredentialClientListByAutomationAccountOptions contains the optional parameters for the CredentialClient.ListByAutomationAccount
// method.
func (client *CredentialClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *CredentialClientListByAutomationAccountOptions) *runtime.Pager[CredentialClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PageProcessor[CredentialClientListByAutomationAccountResponse]{
		More: func(page CredentialClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CredentialClientListByAutomationAccountResponse) (CredentialClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CredentialClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CredentialClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CredentialClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *CredentialClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *CredentialClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials"
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
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *CredentialClient) listByAutomationAccountHandleResponse(resp *http.Response) (CredentialClientListByAutomationAccountResponse, error) {
	result := CredentialClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialListResult); err != nil {
		return CredentialClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update a credential.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// credentialName - The parameters supplied to the Update credential operation.
// parameters - The parameters supplied to the Update credential operation.
// options - CredentialClientUpdateOptions contains the optional parameters for the CredentialClient.Update method.
func (client *CredentialClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, parameters CredentialUpdateParameters, options *CredentialClientUpdateOptions) (CredentialClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, credentialName, parameters, options)
	if err != nil {
		return CredentialClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CredentialClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CredentialClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, parameters CredentialUpdateParameters, options *CredentialClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CredentialClient) updateHandleResponse(resp *http.Response) (CredentialClientUpdateResponse, error) {
	result := CredentialClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Credential); err != nil {
		return CredentialClientUpdateResponse{}, err
	}
	return result, nil
}
