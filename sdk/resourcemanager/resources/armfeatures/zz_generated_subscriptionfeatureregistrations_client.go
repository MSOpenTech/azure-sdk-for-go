//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

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

// SubscriptionFeatureRegistrationsClient contains the methods for the SubscriptionFeatureRegistrations group.
// Don't use this type directly, use NewSubscriptionFeatureRegistrationsClient() instead.
type SubscriptionFeatureRegistrationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSubscriptionFeatureRegistrationsClient creates a new instance of SubscriptionFeatureRegistrationsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubscriptionFeatureRegistrationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionFeatureRegistrationsClient, error) {
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
	client := &SubscriptionFeatureRegistrationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a feature registration.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The provider namespace.
// featureName - The feature name.
// options - SubscriptionFeatureRegistrationsClientCreateOrUpdateOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.CreateOrUpdate
// method.
func (client *SubscriptionFeatureRegistrationsClient) CreateOrUpdate(ctx context.Context, providerNamespace string, featureName string, options *SubscriptionFeatureRegistrationsClientCreateOrUpdateOptions) (SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, featureName, options)
	if err != nil {
		return SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubscriptionFeatureRegistrationsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, featureName string, options *SubscriptionFeatureRegistrationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.SubscriptionFeatureRegistrationType != nil {
		return req, runtime.MarshalAsJSON(req, *options.SubscriptionFeatureRegistrationType)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SubscriptionFeatureRegistrationsClient) createOrUpdateHandleResponse(resp *http.Response) (SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse, error) {
	result := SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionFeatureRegistration); err != nil {
		return SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a feature registration
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The provider namespace.
// featureName - The feature name.
// options - SubscriptionFeatureRegistrationsClientDeleteOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.Delete
// method.
func (client *SubscriptionFeatureRegistrationsClient) Delete(ctx context.Context, providerNamespace string, featureName string, options *SubscriptionFeatureRegistrationsClientDeleteOptions) (SubscriptionFeatureRegistrationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, providerNamespace, featureName, options)
	if err != nil {
		return SubscriptionFeatureRegistrationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionFeatureRegistrationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SubscriptionFeatureRegistrationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionFeatureRegistrationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SubscriptionFeatureRegistrationsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, featureName string, options *SubscriptionFeatureRegistrationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns a feature registration
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The provider namespace.
// featureName - The feature name.
// options - SubscriptionFeatureRegistrationsClientGetOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.Get
// method.
func (client *SubscriptionFeatureRegistrationsClient) Get(ctx context.Context, providerNamespace string, featureName string, options *SubscriptionFeatureRegistrationsClientGetOptions) (SubscriptionFeatureRegistrationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, providerNamespace, featureName, options)
	if err != nil {
		return SubscriptionFeatureRegistrationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionFeatureRegistrationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionFeatureRegistrationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionFeatureRegistrationsClient) getCreateRequest(ctx context.Context, providerNamespace string, featureName string, options *SubscriptionFeatureRegistrationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations/{featureName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionFeatureRegistrationsClient) getHandleResponse(resp *http.Response) (SubscriptionFeatureRegistrationsClientGetResponse, error) {
	result := SubscriptionFeatureRegistrationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionFeatureRegistration); err != nil {
		return SubscriptionFeatureRegistrationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListAllBySubscriptionPager - Returns subscription feature registrations for given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionFeatureRegistrationsClientListAllBySubscriptionOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.ListAllBySubscription
// method.
func (client *SubscriptionFeatureRegistrationsClient) NewListAllBySubscriptionPager(options *SubscriptionFeatureRegistrationsClientListAllBySubscriptionOptions) *runtime.Pager[SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse]{
		More: func(page SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse) (SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllBySubscriptionHandleResponse(resp)
		},
	})
}

// listAllBySubscriptionCreateRequest creates the ListAllBySubscription request.
func (client *SubscriptionFeatureRegistrationsClient) listAllBySubscriptionCreateRequest(ctx context.Context, options *SubscriptionFeatureRegistrationsClientListAllBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/subscriptionFeatureRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllBySubscriptionHandleResponse handles the ListAllBySubscription response.
func (client *SubscriptionFeatureRegistrationsClient) listAllBySubscriptionHandleResponse(resp *http.Response) (SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse, error) {
	result := SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionFeatureRegistrationList); err != nil {
		return SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns subscription feature registrations for given subscription and provider namespace.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The provider namespace.
// options - SubscriptionFeatureRegistrationsClientListBySubscriptionOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.ListBySubscription
// method.
func (client *SubscriptionFeatureRegistrationsClient) NewListBySubscriptionPager(providerNamespace string, options *SubscriptionFeatureRegistrationsClientListBySubscriptionOptions) *runtime.Pager[SubscriptionFeatureRegistrationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[SubscriptionFeatureRegistrationsClientListBySubscriptionResponse]{
		More: func(page SubscriptionFeatureRegistrationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionFeatureRegistrationsClientListBySubscriptionResponse) (SubscriptionFeatureRegistrationsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, providerNamespace, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionFeatureRegistrationsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubscriptionFeatureRegistrationsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionFeatureRegistrationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SubscriptionFeatureRegistrationsClient) listBySubscriptionCreateRequest(ctx context.Context, providerNamespace string, options *SubscriptionFeatureRegistrationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SubscriptionFeatureRegistrationsClient) listBySubscriptionHandleResponse(resp *http.Response) (SubscriptionFeatureRegistrationsClientListBySubscriptionResponse, error) {
	result := SubscriptionFeatureRegistrationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionFeatureRegistrationList); err != nil {
		return SubscriptionFeatureRegistrationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
