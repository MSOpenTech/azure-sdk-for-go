//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpolicyinsights

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
	"strconv"
	"strings"
)

// AttestationsClient contains the methods for the Attestations group.
// Don't use this type directly, use NewAttestationsClient() instead.
type AttestationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAttestationsClient creates a new instance of AttestationsClient with the specified values.
// subscriptionID - Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAttestationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AttestationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AttestationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdateAtResource - Creates or updates an attestation at resource scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceID - Resource ID.
// attestationName - The name of the attestation.
// parameters - The attestation parameters.
// options - AttestationsClientBeginCreateOrUpdateAtResourceOptions contains the optional parameters for the AttestationsClient.BeginCreateOrUpdateAtResource
// method.
func (client *AttestationsClient) BeginCreateOrUpdateAtResource(ctx context.Context, resourceID string, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtResourceOptions) (*runtime.Poller[AttestationsClientCreateOrUpdateAtResourceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateAtResource(ctx, resourceID, attestationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AttestationsClientCreateOrUpdateAtResourceResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AttestationsClientCreateOrUpdateAtResourceResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateAtResource - Creates or updates an attestation at resource scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
func (client *AttestationsClient) createOrUpdateAtResource(ctx context.Context, resourceID string, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtResourceOptions) (*http.Response, error) {
	req, err := client.createOrUpdateAtResourceCreateRequest(ctx, resourceID, attestationName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateAtResourceCreateRequest creates the CreateOrUpdateAtResource request.
func (client *AttestationsClient) createOrUpdateAtResourceCreateRequest(ctx context.Context, resourceID string, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdateAtResourceGroup - Creates or updates an attestation at resource group scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// attestationName - The name of the attestation.
// parameters - The attestation parameters.
// options - AttestationsClientBeginCreateOrUpdateAtResourceGroupOptions contains the optional parameters for the AttestationsClient.BeginCreateOrUpdateAtResourceGroup
// method.
func (client *AttestationsClient) BeginCreateOrUpdateAtResourceGroup(ctx context.Context, resourceGroupName string, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtResourceGroupOptions) (*runtime.Poller[AttestationsClientCreateOrUpdateAtResourceGroupResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateAtResourceGroup(ctx, resourceGroupName, attestationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AttestationsClientCreateOrUpdateAtResourceGroupResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AttestationsClientCreateOrUpdateAtResourceGroupResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateAtResourceGroup - Creates or updates an attestation at resource group scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
func (client *AttestationsClient) createOrUpdateAtResourceGroup(ctx context.Context, resourceGroupName string, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtResourceGroupOptions) (*http.Response, error) {
	req, err := client.createOrUpdateAtResourceGroupCreateRequest(ctx, resourceGroupName, attestationName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateAtResourceGroupCreateRequest creates the CreateOrUpdateAtResourceGroup request.
func (client *AttestationsClient) createOrUpdateAtResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdateAtSubscription - Creates or updates an attestation at subscription scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// attestationName - The name of the attestation.
// parameters - The attestation parameters.
// options - AttestationsClientBeginCreateOrUpdateAtSubscriptionOptions contains the optional parameters for the AttestationsClient.BeginCreateOrUpdateAtSubscription
// method.
func (client *AttestationsClient) BeginCreateOrUpdateAtSubscription(ctx context.Context, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtSubscriptionOptions) (*runtime.Poller[AttestationsClientCreateOrUpdateAtSubscriptionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateAtSubscription(ctx, attestationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AttestationsClientCreateOrUpdateAtSubscriptionResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AttestationsClientCreateOrUpdateAtSubscriptionResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateAtSubscription - Creates or updates an attestation at subscription scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
func (client *AttestationsClient) createOrUpdateAtSubscription(ctx context.Context, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtSubscriptionOptions) (*http.Response, error) {
	req, err := client.createOrUpdateAtSubscriptionCreateRequest(ctx, attestationName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateAtSubscriptionCreateRequest creates the CreateOrUpdateAtSubscription request.
func (client *AttestationsClient) createOrUpdateAtSubscriptionCreateRequest(ctx context.Context, attestationName string, parameters Attestation, options *AttestationsClientBeginCreateOrUpdateAtSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// DeleteAtResource - Deletes an existing attestation at individual resource scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceID - Resource ID.
// attestationName - The name of the attestation.
// options - AttestationsClientDeleteAtResourceOptions contains the optional parameters for the AttestationsClient.DeleteAtResource
// method.
func (client *AttestationsClient) DeleteAtResource(ctx context.Context, resourceID string, attestationName string, options *AttestationsClientDeleteAtResourceOptions) (AttestationsClientDeleteAtResourceResponse, error) {
	req, err := client.deleteAtResourceCreateRequest(ctx, resourceID, attestationName, options)
	if err != nil {
		return AttestationsClientDeleteAtResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationsClientDeleteAtResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AttestationsClientDeleteAtResourceResponse{}, runtime.NewResponseError(resp)
	}
	return AttestationsClientDeleteAtResourceResponse{}, nil
}

// deleteAtResourceCreateRequest creates the DeleteAtResource request.
func (client *AttestationsClient) deleteAtResourceCreateRequest(ctx context.Context, resourceID string, attestationName string, options *AttestationsClientDeleteAtResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAtResourceGroup - Deletes an existing attestation at resource group scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// attestationName - The name of the attestation.
// options - AttestationsClientDeleteAtResourceGroupOptions contains the optional parameters for the AttestationsClient.DeleteAtResourceGroup
// method.
func (client *AttestationsClient) DeleteAtResourceGroup(ctx context.Context, resourceGroupName string, attestationName string, options *AttestationsClientDeleteAtResourceGroupOptions) (AttestationsClientDeleteAtResourceGroupResponse, error) {
	req, err := client.deleteAtResourceGroupCreateRequest(ctx, resourceGroupName, attestationName, options)
	if err != nil {
		return AttestationsClientDeleteAtResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationsClientDeleteAtResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AttestationsClientDeleteAtResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	return AttestationsClientDeleteAtResourceGroupResponse{}, nil
}

// deleteAtResourceGroupCreateRequest creates the DeleteAtResourceGroup request.
func (client *AttestationsClient) deleteAtResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, attestationName string, options *AttestationsClientDeleteAtResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAtSubscription - Deletes an existing attestation at subscription scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// attestationName - The name of the attestation.
// options - AttestationsClientDeleteAtSubscriptionOptions contains the optional parameters for the AttestationsClient.DeleteAtSubscription
// method.
func (client *AttestationsClient) DeleteAtSubscription(ctx context.Context, attestationName string, options *AttestationsClientDeleteAtSubscriptionOptions) (AttestationsClientDeleteAtSubscriptionResponse, error) {
	req, err := client.deleteAtSubscriptionCreateRequest(ctx, attestationName, options)
	if err != nil {
		return AttestationsClientDeleteAtSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationsClientDeleteAtSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AttestationsClientDeleteAtSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return AttestationsClientDeleteAtSubscriptionResponse{}, nil
}

// deleteAtSubscriptionCreateRequest creates the DeleteAtSubscription request.
func (client *AttestationsClient) deleteAtSubscriptionCreateRequest(ctx context.Context, attestationName string, options *AttestationsClientDeleteAtSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetAtResource - Gets an existing attestation at resource scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceID - Resource ID.
// attestationName - The name of the attestation.
// options - AttestationsClientGetAtResourceOptions contains the optional parameters for the AttestationsClient.GetAtResource
// method.
func (client *AttestationsClient) GetAtResource(ctx context.Context, resourceID string, attestationName string, options *AttestationsClientGetAtResourceOptions) (AttestationsClientGetAtResourceResponse, error) {
	req, err := client.getAtResourceCreateRequest(ctx, resourceID, attestationName, options)
	if err != nil {
		return AttestationsClientGetAtResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationsClientGetAtResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationsClientGetAtResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtResourceHandleResponse(resp)
}

// getAtResourceCreateRequest creates the GetAtResource request.
func (client *AttestationsClient) getAtResourceCreateRequest(ctx context.Context, resourceID string, attestationName string, options *AttestationsClientGetAtResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtResourceHandleResponse handles the GetAtResource response.
func (client *AttestationsClient) getAtResourceHandleResponse(resp *http.Response) (AttestationsClientGetAtResourceResponse, error) {
	result := AttestationsClientGetAtResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Attestation); err != nil {
		return AttestationsClientGetAtResourceResponse{}, err
	}
	return result, nil
}

// GetAtResourceGroup - Gets an existing attestation at resource group scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// attestationName - The name of the attestation.
// options - AttestationsClientGetAtResourceGroupOptions contains the optional parameters for the AttestationsClient.GetAtResourceGroup
// method.
func (client *AttestationsClient) GetAtResourceGroup(ctx context.Context, resourceGroupName string, attestationName string, options *AttestationsClientGetAtResourceGroupOptions) (AttestationsClientGetAtResourceGroupResponse, error) {
	req, err := client.getAtResourceGroupCreateRequest(ctx, resourceGroupName, attestationName, options)
	if err != nil {
		return AttestationsClientGetAtResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationsClientGetAtResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationsClientGetAtResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtResourceGroupHandleResponse(resp)
}

// getAtResourceGroupCreateRequest creates the GetAtResourceGroup request.
func (client *AttestationsClient) getAtResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, attestationName string, options *AttestationsClientGetAtResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtResourceGroupHandleResponse handles the GetAtResourceGroup response.
func (client *AttestationsClient) getAtResourceGroupHandleResponse(resp *http.Response) (AttestationsClientGetAtResourceGroupResponse, error) {
	result := AttestationsClientGetAtResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Attestation); err != nil {
		return AttestationsClientGetAtResourceGroupResponse{}, err
	}
	return result, nil
}

// GetAtSubscription - Gets an existing attestation at subscription scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// attestationName - The name of the attestation.
// options - AttestationsClientGetAtSubscriptionOptions contains the optional parameters for the AttestationsClient.GetAtSubscription
// method.
func (client *AttestationsClient) GetAtSubscription(ctx context.Context, attestationName string, options *AttestationsClientGetAtSubscriptionOptions) (AttestationsClientGetAtSubscriptionResponse, error) {
	req, err := client.getAtSubscriptionCreateRequest(ctx, attestationName, options)
	if err != nil {
		return AttestationsClientGetAtSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttestationsClientGetAtSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttestationsClientGetAtSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtSubscriptionHandleResponse(resp)
}

// getAtSubscriptionCreateRequest creates the GetAtSubscription request.
func (client *AttestationsClient) getAtSubscriptionCreateRequest(ctx context.Context, attestationName string, options *AttestationsClientGetAtSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/attestations/{attestationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if attestationName == "" {
		return nil, errors.New("parameter attestationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attestationName}", url.PathEscape(attestationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtSubscriptionHandleResponse handles the GetAtSubscription response.
func (client *AttestationsClient) getAtSubscriptionHandleResponse(resp *http.Response) (AttestationsClientGetAtSubscriptionResponse, error) {
	result := AttestationsClientGetAtSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Attestation); err != nil {
		return AttestationsClientGetAtSubscriptionResponse{}, err
	}
	return result, nil
}

// NewListForResourcePager - Gets all attestations for a resource.
// Generated from API version 2022-09-01
// resourceID - Resource ID.
// QueryOptions - QueryOptions contains a group of parameters for the PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup
// method.
// options - AttestationsClientListForResourceOptions contains the optional parameters for the AttestationsClient.ListForResource
// method.
func (client *AttestationsClient) NewListForResourcePager(resourceID string, queryOptions *QueryOptions, options *AttestationsClientListForResourceOptions) *runtime.Pager[AttestationsClientListForResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AttestationsClientListForResourceResponse]{
		More: func(page AttestationsClientListForResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttestationsClientListForResourceResponse) (AttestationsClientListForResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceCreateRequest(ctx, resourceID, queryOptions, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AttestationsClientListForResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttestationsClientListForResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttestationsClientListForResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceHandleResponse(resp)
		},
	})
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *AttestationsClient) listForResourceCreateRequest(ctx context.Context, resourceID string, queryOptions *QueryOptions, options *AttestationsClientListForResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.PolicyInsights/attestations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if queryOptions != nil && queryOptions.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*queryOptions.Top), 10))
	}
	if queryOptions != nil && queryOptions.Filter != nil {
		reqQP.Set("$filter", *queryOptions.Filter)
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *AttestationsClient) listForResourceHandleResponse(resp *http.Response) (AttestationsClientListForResourceResponse, error) {
	result := AttestationsClientListForResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationListResult); err != nil {
		return AttestationsClientListForResourceResponse{}, err
	}
	return result, nil
}

// NewListForResourceGroupPager - Gets all attestations for the resource group.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// QueryOptions - QueryOptions contains a group of parameters for the PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup
// method.
// options - AttestationsClientListForResourceGroupOptions contains the optional parameters for the AttestationsClient.ListForResourceGroup
// method.
func (client *AttestationsClient) NewListForResourceGroupPager(resourceGroupName string, queryOptions *QueryOptions, options *AttestationsClientListForResourceGroupOptions) *runtime.Pager[AttestationsClientListForResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AttestationsClientListForResourceGroupResponse]{
		More: func(page AttestationsClientListForResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttestationsClientListForResourceGroupResponse) (AttestationsClientListForResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForResourceGroupCreateRequest(ctx, resourceGroupName, queryOptions, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AttestationsClientListForResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttestationsClientListForResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttestationsClientListForResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForResourceGroupHandleResponse(resp)
		},
	})
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *AttestationsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, queryOptions *QueryOptions, options *AttestationsClientListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/attestations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if queryOptions != nil && queryOptions.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*queryOptions.Top), 10))
	}
	if queryOptions != nil && queryOptions.Filter != nil {
		reqQP.Set("$filter", *queryOptions.Filter)
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *AttestationsClient) listForResourceGroupHandleResponse(resp *http.Response) (AttestationsClientListForResourceGroupResponse, error) {
	result := AttestationsClientListForResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationListResult); err != nil {
		return AttestationsClientListForResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListForSubscriptionPager - Gets all attestations for the subscription.
// Generated from API version 2022-09-01
// QueryOptions - QueryOptions contains a group of parameters for the PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup
// method.
// options - AttestationsClientListForSubscriptionOptions contains the optional parameters for the AttestationsClient.ListForSubscription
// method.
func (client *AttestationsClient) NewListForSubscriptionPager(queryOptions *QueryOptions, options *AttestationsClientListForSubscriptionOptions) *runtime.Pager[AttestationsClientListForSubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AttestationsClientListForSubscriptionResponse]{
		More: func(page AttestationsClientListForSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttestationsClientListForSubscriptionResponse) (AttestationsClientListForSubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForSubscriptionCreateRequest(ctx, queryOptions, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AttestationsClientListForSubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttestationsClientListForSubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttestationsClientListForSubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForSubscriptionHandleResponse(resp)
		},
	})
}

// listForSubscriptionCreateRequest creates the ListForSubscription request.
func (client *AttestationsClient) listForSubscriptionCreateRequest(ctx context.Context, queryOptions *QueryOptions, options *AttestationsClientListForSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/attestations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if queryOptions != nil && queryOptions.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*queryOptions.Top), 10))
	}
	if queryOptions != nil && queryOptions.Filter != nil {
		reqQP.Set("$filter", *queryOptions.Filter)
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForSubscriptionHandleResponse handles the ListForSubscription response.
func (client *AttestationsClient) listForSubscriptionHandleResponse(resp *http.Response) (AttestationsClientListForSubscriptionResponse, error) {
	result := AttestationsClientListForSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttestationListResult); err != nil {
		return AttestationsClientListForSubscriptionResponse{}, err
	}
	return result, nil
}
