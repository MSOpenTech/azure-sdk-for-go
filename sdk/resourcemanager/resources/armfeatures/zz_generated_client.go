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

// Client contains the methods for the Features group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
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
	client := &Client{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the preview feature with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceProviderNamespace - The resource provider namespace for the feature.
// featureName - The name of the feature to get.
// options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, resourceProviderNamespace string, featureName string, options *ClientGetOptions) (ClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceProviderNamespace, featureName, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, resourceProviderNamespace string, featureName string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureResult); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the preview features in a provider namespace that are available through AFEC for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceProviderNamespace - The namespace of the resource provider for getting features.
// options - ClientListOptions contains the optional parameters for the Client.List method.
func (client *Client) NewListPager(resourceProviderNamespace string, options *ClientListOptions) *runtime.Pager[ClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ClientListResponse]{
		More: func(page ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListResponse) (ClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceProviderNamespace, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ClientListResponse, error) {
	result := ClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureOperationsListResult); err != nil {
		return ClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the preview features that are available through AFEC for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ClientListAllOptions contains the optional parameters for the Client.ListAll method.
func (client *Client) NewListAllPager(options *ClientListAllOptions) *runtime.Pager[ClientListAllResponse] {
	return runtime.NewPager(runtime.PageProcessor[ClientListAllResponse]{
		More: func(page ClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListAllResponse) (ClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *Client) listAllCreateRequest(ctx context.Context, options *ClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/features"
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *Client) listAllHandleResponse(resp *http.Response) (ClientListAllResponse, error) {
	result := ClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureOperationsListResult); err != nil {
		return ClientListAllResponse{}, err
	}
	return result, nil
}

// Register - Registers the preview feature for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceProviderNamespace - The namespace of the resource provider.
// featureName - The name of the feature to register.
// options - ClientRegisterOptions contains the optional parameters for the Client.Register method.
func (client *Client) Register(ctx context.Context, resourceProviderNamespace string, featureName string, options *ClientRegisterOptions) (ClientRegisterResponse, error) {
	req, err := client.registerCreateRequest(ctx, resourceProviderNamespace, featureName, options)
	if err != nil {
		return ClientRegisterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientRegisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientRegisterResponse{}, runtime.NewResponseError(resp)
	}
	return client.registerHandleResponse(resp)
}

// registerCreateRequest creates the Register request.
func (client *Client) registerCreateRequest(ctx context.Context, resourceProviderNamespace string, featureName string, options *ClientRegisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}/register"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// registerHandleResponse handles the Register response.
func (client *Client) registerHandleResponse(resp *http.Response) (ClientRegisterResponse, error) {
	result := ClientRegisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureResult); err != nil {
		return ClientRegisterResponse{}, err
	}
	return result, nil
}

// Unregister - Unregisters the preview feature for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceProviderNamespace - The namespace of the resource provider.
// featureName - The name of the feature to unregister.
// options - ClientUnregisterOptions contains the optional parameters for the Client.Unregister method.
func (client *Client) Unregister(ctx context.Context, resourceProviderNamespace string, featureName string, options *ClientUnregisterOptions) (ClientUnregisterResponse, error) {
	req, err := client.unregisterCreateRequest(ctx, resourceProviderNamespace, featureName, options)
	if err != nil {
		return ClientUnregisterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientUnregisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientUnregisterResponse{}, runtime.NewResponseError(resp)
	}
	return client.unregisterHandleResponse(resp)
}

// unregisterCreateRequest creates the Unregister request.
func (client *Client) unregisterCreateRequest(ctx context.Context, resourceProviderNamespace string, featureName string, options *ClientUnregisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}/unregister"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if featureName == "" {
		return nil, errors.New("parameter featureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{featureName}", url.PathEscape(featureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// unregisterHandleResponse handles the Unregister response.
func (client *Client) unregisterHandleResponse(resp *http.Response) (ClientUnregisterResponse, error) {
	result := ClientUnregisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeatureResult); err != nil {
		return ClientUnregisterResponse{}, err
	}
	return result, nil
}
