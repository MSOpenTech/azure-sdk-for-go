//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventhub

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DisasterRecoveryConfigsClient contains the methods for the DisasterRecoveryConfigs group.
// Don't use this type directly, use NewDisasterRecoveryConfigsClient() instead.
type DisasterRecoveryConfigsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDisasterRecoveryConfigsClient creates a new instance of DisasterRecoveryConfigsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDisasterRecoveryConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DisasterRecoveryConfigsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armeventhub.DisasterRecoveryConfigsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DisasterRecoveryConfigsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetAuthorizationRule - Gets an AuthorizationRule for a Namespace by rule name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - authorizationRuleName - The authorization rule name.
//   - options - DisasterRecoveryConfigsClientGetAuthorizationRuleOptions contains the optional parameters for the DisasterRecoveryConfigsClient.GetAuthorizationRule
//     method.
func (client *DisasterRecoveryConfigsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientGetAuthorizationRuleOptions) (DisasterRecoveryConfigsClientGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientGetAuthorizationRuleResponse, error) {
	result := DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRule); err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// NewListAuthorizationRulesPager - Gets a list of authorization rules for a Namespace.
//
// Generated from API version 2017-04-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - options - DisasterRecoveryConfigsClientListAuthorizationRulesOptions contains the optional parameters for the DisasterRecoveryConfigsClient.NewListAuthorizationRulesPager
//     method.
func (client *DisasterRecoveryConfigsClient) NewListAuthorizationRulesPager(resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientListAuthorizationRulesOptions) *runtime.Pager[DisasterRecoveryConfigsClientListAuthorizationRulesResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisasterRecoveryConfigsClientListAuthorizationRulesResponse]{
		More: func(page DisasterRecoveryConfigsClientListAuthorizationRulesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisasterRecoveryConfigsClientListAuthorizationRulesResponse) (DisasterRecoveryConfigsClientListAuthorizationRulesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
	})
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientListAuthorizationRulesResponse, error) {
	result := DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRuleListResult); err != nil {
		return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// ListKeys - Gets the primary and secondary connection strings for the Namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - authorizationRuleName - The authorization rule name.
//   - options - DisasterRecoveryConfigsClientListKeysOptions contains the optional parameters for the DisasterRecoveryConfigsClient.ListKeys
//     method.
func (client *DisasterRecoveryConfigsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientListKeysOptions) (DisasterRecoveryConfigsClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *DisasterRecoveryConfigsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/authorizationRules/{authorizationRuleName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if alias == "" {
		return nil, errors.New("parameter alias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alias}", url.PathEscape(alias))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *DisasterRecoveryConfigsClient) listKeysHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientListKeysResponse, error) {
	result := DisasterRecoveryConfigsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	return result, nil
}
