//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"errors"
	"fmt"
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
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDisasterRecoveryConfigsClient creates a new instance of DisasterRecoveryConfigsClient with the specified values.
func NewDisasterRecoveryConfigsClient(con *arm.Connection, subscriptionID string) *DisasterRecoveryConfigsClient {
	return &DisasterRecoveryConfigsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BreakPairing - This operation disables the Disaster Recovery and stops replicating changes from primary to secondary namespaces
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) BreakPairing(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsBreakPairingOptions) (DisasterRecoveryConfigsBreakPairingResponse, error) {
	req, err := client.breakPairingCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsBreakPairingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsBreakPairingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsBreakPairingResponse{}, client.breakPairingHandleError(resp)
	}
	return DisasterRecoveryConfigsBreakPairingResponse{RawResponse: resp}, nil
}

// breakPairingCreateRequest creates the BreakPairing request.
func (client *DisasterRecoveryConfigsClient) breakPairingCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsBreakPairingOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// breakPairingHandleError handles the BreakPairing error response.
func (client *DisasterRecoveryConfigsClient) breakPairingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CheckNameAvailability - Check the give Namespace name availability.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailabilityParameter, options *DisasterRecoveryConfigsCheckNameAvailabilityOptions) (DisasterRecoveryConfigsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, namespaceName, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, parameters CheckNameAvailabilityParameter, options *DisasterRecoveryConfigsCheckNameAvailabilityOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DisasterRecoveryConfigsCheckNameAvailabilityResponse, error) {
	result := DisasterRecoveryConfigsCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return DisasterRecoveryConfigsCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *DisasterRecoveryConfigsClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CreateOrUpdate - Creates or updates a new Alias(Disaster Recovery configuration)
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsCreateOrUpdateOptions) (DisasterRecoveryConfigsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, alias, parameters, options)
	if err != nil {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisasterRecoveryConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters ArmDisasterRecovery, options *DisasterRecoveryConfigsCreateOrUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DisasterRecoveryConfigsClient) createOrUpdateHandleResponse(resp *http.Response) (DisasterRecoveryConfigsCreateOrUpdateResponse, error) {
	result := DisasterRecoveryConfigsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DisasterRecoveryConfigsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an Alias(Disaster Recovery configuration)
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsDeleteOptions) (DisasterRecoveryConfigsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DisasterRecoveryConfigsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DisasterRecoveryConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsDeleteOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DisasterRecoveryConfigsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// FailOver - Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) FailOver(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsFailOverOptions) (DisasterRecoveryConfigsFailOverResponse, error) {
	req, err := client.failOverCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsFailOverResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsFailOverResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsFailOverResponse{}, client.failOverHandleError(resp)
	}
	return DisasterRecoveryConfigsFailOverResponse{RawResponse: resp}, nil
}

// failOverCreateRequest creates the FailOver request.
func (client *DisasterRecoveryConfigsClient) failOverCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsFailOverOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// failOverHandleError handles the FailOver error response.
func (client *DisasterRecoveryConfigsClient) failOverHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsGetOptions) (DisasterRecoveryConfigsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
	if err != nil {
		return DisasterRecoveryConfigsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DisasterRecoveryConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisasterRecoveryConfigsClient) getHandleResponse(resp *http.Response) (DisasterRecoveryConfigsGetResponse, error) {
	result := DisasterRecoveryConfigsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecovery); err != nil {
		return DisasterRecoveryConfigsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DisasterRecoveryConfigsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetAuthorizationRule - Gets an AuthorizationRule for a Namespace by rule name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsGetAuthorizationRuleOptions) (DisasterRecoveryConfigsGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, client.getAuthorizationRuleHandleError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsGetAuthorizationRuleOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleHandleResponse(resp *http.Response) (DisasterRecoveryConfigsGetAuthorizationRuleResponse, error) {
	result := DisasterRecoveryConfigsGetAuthorizationRuleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRule); err != nil {
		return DisasterRecoveryConfigsGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// getAuthorizationRuleHandleError handles the GetAuthorizationRule error response.
func (client *DisasterRecoveryConfigsClient) getAuthorizationRuleHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all Alias(Disaster Recovery configurations)
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) List(resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsListOptions) *DisasterRecoveryConfigsListPager {
	return &DisasterRecoveryConfigsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp DisasterRecoveryConfigsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ArmDisasterRecoveryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DisasterRecoveryConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *DisasterRecoveryConfigsListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisasterRecoveryConfigsClient) listHandleResponse(resp *http.Response) (DisasterRecoveryConfigsListResponse, error) {
	result := DisasterRecoveryConfigsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmDisasterRecoveryListResult); err != nil {
		return DisasterRecoveryConfigsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DisasterRecoveryConfigsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAuthorizationRules - Gets a list of authorization rules for a Namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsListAuthorizationRulesOptions) *DisasterRecoveryConfigsListAuthorizationRulesPager {
	return &DisasterRecoveryConfigsListAuthorizationRulesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, alias, options)
		},
		advancer: func(ctx context.Context, resp DisasterRecoveryConfigsListAuthorizationRulesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AuthorizationRuleListResult.NextLink)
		},
	}
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, options *DisasterRecoveryConfigsListAuthorizationRulesOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesHandleResponse(resp *http.Response) (DisasterRecoveryConfigsListAuthorizationRulesResponse, error) {
	result := DisasterRecoveryConfigsListAuthorizationRulesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRuleListResult); err != nil {
		return DisasterRecoveryConfigsListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// listAuthorizationRulesHandleError handles the ListAuthorizationRules error response.
func (client *DisasterRecoveryConfigsClient) listAuthorizationRulesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListKeys - Gets the primary and secondary connection strings for the Namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DisasterRecoveryConfigsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsListKeysOptions) (DisasterRecoveryConfigsListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, alias, authorizationRuleName, options)
	if err != nil {
		return DisasterRecoveryConfigsListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisasterRecoveryConfigsListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisasterRecoveryConfigsListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *DisasterRecoveryConfigsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string, options *DisasterRecoveryConfigsListKeysOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *DisasterRecoveryConfigsClient) listKeysHandleResponse(resp *http.Response) (DisasterRecoveryConfigsListKeysResponse, error) {
	result := DisasterRecoveryConfigsListKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return DisasterRecoveryConfigsListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *DisasterRecoveryConfigsClient) listKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
