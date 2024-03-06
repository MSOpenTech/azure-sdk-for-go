//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"errors"
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
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DisasterRecoveryConfigsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BreakPairing - This operation disables the Disaster Recovery and stops replicating changes from primary to secondary namespaces
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - options - DisasterRecoveryConfigsClientBreakPairingOptions contains the optional parameters for the DisasterRecoveryConfigsClient.BreakPairing
//     method.
func (client *DisasterRecoveryConfigsClient) BreakPairing(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientBreakPairingOptions) (DisasterRecoveryConfigsClientBreakPairingResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.BreakPairing"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.breakPairingCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientBreakPairingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientBreakPairingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientBreakPairingResponse{}, err
	}
	return DisasterRecoveryConfigsClientBreakPairingResponse{}, nil
}

// breakPairingCreateRequest creates the BreakPairing request.
func (client *DisasterRecoveryConfigsClient) breakPairingCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientBreakPairingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/breakPairing"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CheckNameAvailability - Check the give Namespace name availability.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - parameters - Parameters to check availability of the given Alias name
//   - options - DisasterRecoveryConfigsClientCheckNameAvailabilityOptions contains the optional parameters for the DisasterRecoveryConfigsClient.CheckNameAvailability
//     method.
func (client *DisasterRecoveryConfigsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailabilityParameter, options *DisasterRecoveryConfigsClientCheckNameAvailabilityOptions) (DisasterRecoveryConfigsClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, namespaceName, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailabilityParameter, options *DisasterRecoveryConfigsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientCheckNameAvailabilityResponse, error) {
	result := DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return DisasterRecoveryConfigsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a new Alias(Disaster Recovery configuration)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - parameters - Parameters required to create an Alias(Disaster Recovery configuration)
//   - options - DisasterRecoveryConfigsClientCreateOrUpdateOptions contains the optional parameters for the DisasterRecoveryConfigsClient.CreateOrUpdate
//     method.
func (client *DisasterRecoveryConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsClientCreateOrUpdateOptions) (DisasterRecoveryConfigsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, alias, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisasterRecoveryConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DisasterRecoveryConfigsClient) createOrUpdateHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientCreateOrUpdateResponse, error) {
	result := DisasterRecoveryConfigsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Alias(Disaster Recovery configuration)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - options - DisasterRecoveryConfigsClientDeleteOptions contains the optional parameters for the DisasterRecoveryConfigsClient.Delete
//     method.
func (client *DisasterRecoveryConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientDeleteOptions) (DisasterRecoveryConfigsClientDeleteResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientDeleteResponse{}, err
	}
	return DisasterRecoveryConfigsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DisasterRecoveryConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// FailOver - Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - options - DisasterRecoveryConfigsClientFailOverOptions contains the optional parameters for the DisasterRecoveryConfigsClient.FailOver
//     method.
func (client *DisasterRecoveryConfigsClient) FailOver(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientFailOverOptions) (DisasterRecoveryConfigsClientFailOverResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.FailOver"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.failOverCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientFailOverResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientFailOverResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientFailOverResponse{}, err
	}
	return DisasterRecoveryConfigsClientFailOverResponse{}, nil
}

// failOverCreateRequest creates the FailOver request.
func (client *DisasterRecoveryConfigsClient) failOverCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientFailOverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/failover"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - options - DisasterRecoveryConfigsClientGetOptions contains the optional parameters for the DisasterRecoveryConfigsClient.Get
//     method.
func (client *DisasterRecoveryConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientGetOptions) (DisasterRecoveryConfigsClientGetResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DisasterRecoveryConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}"
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
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisasterRecoveryConfigsClient) getHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientGetResponse, error) {
	result := DisasterRecoveryConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthorizationRule - Gets an AuthorizationRule for a Namespace by rule name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - authorizationRuleName - The authorization rule name.
//   - options - DisasterRecoveryConfigsClientGetAuthorizationRuleOptions contains the optional parameters for the DisasterRecoveryConfigsClient.GetAuthorizationRule
//     method.
func (client *DisasterRecoveryConfigsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientGetAuthorizationRuleOptions) (DisasterRecoveryConfigsClientGetAuthorizationRuleResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.GetAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.getAuthorizationRuleHandleResponse(httpResp)
	return resp, err
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
	reqQP.Set("api-version", "2022-10-01-preview")
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

// NewListPager - Gets all Alias(Disaster Recovery configurations)
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - options - DisasterRecoveryConfigsClientListOptions contains the optional parameters for the DisasterRecoveryConfigsClient.NewListPager
//     method.
func (client *DisasterRecoveryConfigsClient) NewListPager(resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsClientListOptions) *runtime.Pager[DisasterRecoveryConfigsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisasterRecoveryConfigsClientListResponse]{
		More: func(page DisasterRecoveryConfigsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisasterRecoveryConfigsClientListResponse) (DisasterRecoveryConfigsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DisasterRecoveryConfigsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
			}, nil)
			if err != nil {
				return DisasterRecoveryConfigsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DisasterRecoveryConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisasterRecoveryConfigsClient) listHandleResponse(resp *http.Response) (DisasterRecoveryConfigsClientListResponse, error) {
	result := DisasterRecoveryConfigsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecoveryListResult); err != nil {
		return DisasterRecoveryConfigsClientListResponse{}, err
	}
	return result, nil
}

// NewListAuthorizationRulesPager - Gets a list of authorization rules for a Namespace.
//
// Generated from API version 2022-10-01-preview
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
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DisasterRecoveryConfigsClient.NewListAuthorizationRulesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
			}, nil)
			if err != nil {
				return DisasterRecoveryConfigsClientListAuthorizationRulesResponse{}, err
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
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
	reqQP.Set("api-version", "2022-10-01-preview")
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
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - alias - The Disaster Recovery configuration name
//   - authorizationRuleName - The authorization rule name.
//   - options - DisasterRecoveryConfigsClientListKeysOptions contains the optional parameters for the DisasterRecoveryConfigsClient.ListKeys
//     method.
func (client *DisasterRecoveryConfigsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsClientListKeysOptions) (DisasterRecoveryConfigsClientListKeysResponse, error) {
	var err error
	const operationName = "DisasterRecoveryConfigsClient.ListKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisasterRecoveryConfigsClientListKeysResponse{}, err
	}
	resp, err := client.listKeysHandleResponse(httpResp)
	return resp, err
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
	reqQP.Set("api-version", "2022-10-01-preview")
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
