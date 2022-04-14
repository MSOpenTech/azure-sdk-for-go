//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

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

// HybridConnectionsClient contains the methods for the HybridConnections group.
// Don't use this type directly, use NewHybridConnectionsClient() instead.
type HybridConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHybridConnectionsClient creates a new instance of HybridConnectionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHybridConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridConnectionsClient, error) {
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
	client := &HybridConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a service hybrid connection. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// parameters - Parameters supplied to create a hybrid connection.
// options - HybridConnectionsClientCreateOrUpdateOptions contains the optional parameters for the HybridConnectionsClient.CreateOrUpdate
// method.
func (client *HybridConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, parameters HybridConnection, options *HybridConnectionsClientCreateOrUpdateOptions) (HybridConnectionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, parameters, options)
	if err != nil {
		return HybridConnectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridConnectionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HybridConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, parameters HybridConnection, options *HybridConnectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HybridConnectionsClient) createOrUpdateHandleResponse(resp *http.Response) (HybridConnectionsClientCreateOrUpdateResponse, error) {
	result := HybridConnectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridConnection); err != nil {
		return HybridConnectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAuthorizationRule - Creates or updates an authorization rule for a hybrid connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// authorizationRuleName - The authorization rule name.
// parameters - The authorization rule parameters.
// options - HybridConnectionsClientCreateOrUpdateAuthorizationRuleOptions contains the optional parameters for the HybridConnectionsClient.CreateOrUpdateAuthorizationRule
// method.
func (client *HybridConnectionsClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters AuthorizationRule, options *HybridConnectionsClientCreateOrUpdateAuthorizationRuleOptions) (HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse, error) {
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, authorizationRuleName, parameters, options)
	if err != nil {
		return HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateAuthorizationRuleHandleResponse(resp)
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *HybridConnectionsClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters AuthorizationRule, options *HybridConnectionsClientCreateOrUpdateAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *HybridConnectionsClient) createOrUpdateAuthorizationRuleHandleResponse(resp *http.Response) (HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse, error) {
	result := HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRule); err != nil {
		return HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a hybrid connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// options - HybridConnectionsClientDeleteOptions contains the optional parameters for the HybridConnectionsClient.Delete
// method.
func (client *HybridConnectionsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *HybridConnectionsClientDeleteOptions) (HybridConnectionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, options)
	if err != nil {
		return HybridConnectionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HybridConnectionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HybridConnectionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *HybridConnectionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteAuthorizationRule - Deletes a hybrid connection authorization rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// authorizationRuleName - The authorization rule name.
// options - HybridConnectionsClientDeleteAuthorizationRuleOptions contains the optional parameters for the HybridConnectionsClient.DeleteAuthorizationRule
// method.
func (client *HybridConnectionsClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *HybridConnectionsClientDeleteAuthorizationRuleOptions) (HybridConnectionsClientDeleteAuthorizationRuleResponse, error) {
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, authorizationRuleName, options)
	if err != nil {
		return HybridConnectionsClientDeleteAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientDeleteAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HybridConnectionsClientDeleteAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return HybridConnectionsClientDeleteAuthorizationRuleResponse{}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *HybridConnectionsClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *HybridConnectionsClientDeleteAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the description for the specified hybrid connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// options - HybridConnectionsClientGetOptions contains the optional parameters for the HybridConnectionsClient.Get method.
func (client *HybridConnectionsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *HybridConnectionsClientGetOptions) (HybridConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, options)
	if err != nil {
		return HybridConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HybridConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *HybridConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridConnectionsClient) getHandleResponse(resp *http.Response) (HybridConnectionsClientGetResponse, error) {
	result := HybridConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridConnection); err != nil {
		return HybridConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthorizationRule - Hybrid connection authorization rule for a hybrid connection by name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// authorizationRuleName - The authorization rule name.
// options - HybridConnectionsClientGetAuthorizationRuleOptions contains the optional parameters for the HybridConnectionsClient.GetAuthorizationRule
// method.
func (client *HybridConnectionsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *HybridConnectionsClientGetAuthorizationRuleOptions) (HybridConnectionsClientGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, authorizationRuleName, options)
	if err != nil {
		return HybridConnectionsClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridConnectionsClientGetAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *HybridConnectionsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *HybridConnectionsClientGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *HybridConnectionsClient) getAuthorizationRuleHandleResponse(resp *http.Response) (HybridConnectionsClientGetAuthorizationRuleResponse, error) {
	result := HybridConnectionsClientGetAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRule); err != nil {
		return HybridConnectionsClientGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// ListAuthorizationRules - Authorization rules for a hybrid connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// options - HybridConnectionsClientListAuthorizationRulesOptions contains the optional parameters for the HybridConnectionsClient.ListAuthorizationRules
// method.
func (client *HybridConnectionsClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, hybridConnectionName string, options *HybridConnectionsClientListAuthorizationRulesOptions) *runtime.Pager[HybridConnectionsClientListAuthorizationRulesResponse] {
	return runtime.NewPager(runtime.PageProcessor[HybridConnectionsClientListAuthorizationRulesResponse]{
		More: func(page HybridConnectionsClientListAuthorizationRulesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridConnectionsClientListAuthorizationRulesResponse) (HybridConnectionsClientListAuthorizationRulesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridConnectionsClientListAuthorizationRulesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HybridConnectionsClientListAuthorizationRulesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridConnectionsClientListAuthorizationRulesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
	})
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *HybridConnectionsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *HybridConnectionsClientListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *HybridConnectionsClient) listAuthorizationRulesHandleResponse(resp *http.Response) (HybridConnectionsClientListAuthorizationRulesResponse, error) {
	result := HybridConnectionsClientListAuthorizationRulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRuleListResult); err != nil {
		return HybridConnectionsClientListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// ListByNamespace - Lists the hybrid connection within the namespace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// options - HybridConnectionsClientListByNamespaceOptions contains the optional parameters for the HybridConnectionsClient.ListByNamespace
// method.
func (client *HybridConnectionsClient) ListByNamespace(resourceGroupName string, namespaceName string, options *HybridConnectionsClientListByNamespaceOptions) *runtime.Pager[HybridConnectionsClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PageProcessor[HybridConnectionsClientListByNamespaceResponse]{
		More: func(page HybridConnectionsClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridConnectionsClientListByNamespaceResponse) (HybridConnectionsClientListByNamespaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridConnectionsClientListByNamespaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HybridConnectionsClientListByNamespaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridConnectionsClientListByNamespaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNamespaceHandleResponse(resp)
		},
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *HybridConnectionsClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *HybridConnectionsClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *HybridConnectionsClient) listByNamespaceHandleResponse(resp *http.Response) (HybridConnectionsClientListByNamespaceResponse, error) {
	result := HybridConnectionsClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridConnectionListResult); err != nil {
		return HybridConnectionsClientListByNamespaceResponse{}, err
	}
	return result, nil
}

// ListKeys - Primary and secondary connection strings to the hybrid connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// authorizationRuleName - The authorization rule name.
// options - HybridConnectionsClientListKeysOptions contains the optional parameters for the HybridConnectionsClient.ListKeys
// method.
func (client *HybridConnectionsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *HybridConnectionsClientListKeysOptions) (HybridConnectionsClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, authorizationRuleName, options)
	if err != nil {
		return HybridConnectionsClientListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridConnectionsClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *HybridConnectionsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *HybridConnectionsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *HybridConnectionsClient) listKeysHandleResponse(resp *http.Response) (HybridConnectionsClientListKeysResponse, error) {
	result := HybridConnectionsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return HybridConnectionsClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - Regenerates the primary or secondary connection strings to the hybrid connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// hybridConnectionName - The hybrid connection name.
// authorizationRuleName - The authorization rule name.
// parameters - Parameters supplied to regenerate authorization rule.
// options - HybridConnectionsClientRegenerateKeysOptions contains the optional parameters for the HybridConnectionsClient.RegenerateKeys
// method.
func (client *HybridConnectionsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *HybridConnectionsClientRegenerateKeysOptions) (HybridConnectionsClientRegenerateKeysResponse, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, hybridConnectionName, authorizationRuleName, parameters, options)
	if err != nil {
		return HybridConnectionsClientRegenerateKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridConnectionsClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridConnectionsClientRegenerateKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeysHandleResponse(resp)
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *HybridConnectionsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *HybridConnectionsClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if hybridConnectionName == "" {
		return nil, errors.New("parameter hybridConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridConnectionName}", url.PathEscape(hybridConnectionName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *HybridConnectionsClient) regenerateKeysHandleResponse(resp *http.Response) (HybridConnectionsClientRegenerateKeysResponse, error) {
	result := HybridConnectionsClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return HybridConnectionsClientRegenerateKeysResponse{}, err
	}
	return result, nil
}
