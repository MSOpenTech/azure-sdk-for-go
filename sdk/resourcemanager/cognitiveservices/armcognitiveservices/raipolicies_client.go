//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// RaiPoliciesClient contains the methods for the RaiPolicies group.
// Don't use this type directly, use NewRaiPoliciesClient() instead.
type RaiPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRaiPoliciesClient creates a new instance of RaiPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRaiPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RaiPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RaiPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Update the state of specified Content Filters associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiPolicyName - The name of the RaiPolicy associated with the Cognitive Services Account
//   - raiPolicy - Properties describing the Content Filters.
//   - options - RaiPoliciesClientCreateOrUpdateOptions contains the optional parameters for the RaiPoliciesClient.CreateOrUpdate
//     method.
func (client *RaiPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, raiPolicy RaiPolicy, options *RaiPoliciesClientCreateOrUpdateOptions) (RaiPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "RaiPoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, raiPolicyName, raiPolicy, options)
	if err != nil {
		return RaiPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RaiPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RaiPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RaiPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, raiPolicy RaiPolicy, options *RaiPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiPolicies/{raiPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiPolicyName == "" {
		return nil, errors.New("parameter raiPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiPolicyName}", url.PathEscape(raiPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, raiPolicy); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RaiPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (RaiPoliciesClientCreateOrUpdateResponse, error) {
	result := RaiPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiPolicy); err != nil {
		return RaiPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified Content Filters associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiPolicyName - The name of the RaiPolicy associated with the Cognitive Services Account
//   - options - RaiPoliciesClientBeginDeleteOptions contains the optional parameters for the RaiPoliciesClient.BeginDelete method.
func (client *RaiPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, options *RaiPoliciesClientBeginDeleteOptions) (*runtime.Poller[RaiPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, raiPolicyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RaiPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RaiPoliciesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified Content Filters associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *RaiPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, options *RaiPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RaiPoliciesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, raiPolicyName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RaiPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, options *RaiPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiPolicies/{raiPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiPolicyName == "" {
		return nil, errors.New("parameter raiPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiPolicyName}", url.PathEscape(raiPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Content Filters associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiPolicyName - The name of the RaiPolicy associated with the Cognitive Services Account
//   - options - RaiPoliciesClientGetOptions contains the optional parameters for the RaiPoliciesClient.Get method.
func (client *RaiPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, options *RaiPoliciesClientGetOptions) (RaiPoliciesClientGetResponse, error) {
	var err error
	const operationName = "RaiPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, raiPolicyName, options)
	if err != nil {
		return RaiPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RaiPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RaiPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RaiPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiPolicyName string, options *RaiPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiPolicies/{raiPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiPolicyName == "" {
		return nil, errors.New("parameter raiPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiPolicyName}", url.PathEscape(raiPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RaiPoliciesClient) getHandleResponse(resp *http.Response) (RaiPoliciesClientGetResponse, error) {
	result := RaiPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiPolicy); err != nil {
		return RaiPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the content filters associated with the Azure OpenAI account.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - options - RaiPoliciesClientListOptions contains the optional parameters for the RaiPoliciesClient.NewListPager method.
func (client *RaiPoliciesClient) NewListPager(resourceGroupName string, accountName string, options *RaiPoliciesClientListOptions) *runtime.Pager[RaiPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RaiPoliciesClientListResponse]{
		More: func(page RaiPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RaiPoliciesClientListResponse) (RaiPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RaiPoliciesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return RaiPoliciesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RaiPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *RaiPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RaiPoliciesClient) listHandleResponse(resp *http.Response) (RaiPoliciesClientListResponse, error) {
	result := RaiPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiPolicyListResult); err != nil {
		return RaiPoliciesClientListResponse{}, err
	}
	return result, nil
}
