//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// ReportsClient contains the methods for the Reports group.
// Don't use this type directly, use NewReportsClient() instead.
type ReportsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewReportsClient creates a new instance of ReportsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReportsClient, error) {
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
	client := &ReportsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByAPIPager - Lists report records by API.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - The filter to apply on the operation.
// options - ReportsClientListByAPIOptions contains the optional parameters for the ReportsClient.ListByAPI method.
func (client *ReportsClient) NewListByAPIPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListByAPIOptions) *runtime.Pager[ReportsClientListByAPIResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByAPIResponse]{
		More: func(page ReportsClientListByAPIResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByAPIResponse) (ReportsClientListByAPIResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAPICreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListByAPIResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByAPIResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByAPIResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAPIHandleResponse(resp)
		},
	})
}

// listByAPICreateRequest creates the ListByAPI request.
func (client *ReportsClient) listByAPICreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListByAPIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byApi"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAPIHandleResponse handles the ListByAPI response.
func (client *ReportsClient) listByAPIHandleResponse(resp *http.Response) (ReportsClientListByAPIResponse, error) {
	result := ReportsClientListByAPIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListByAPIResponse{}, err
	}
	return result, nil
}

// NewListByGeoPager - Lists report records by geography.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter | ge, le | |
// | country | select | | |
// | region | select | | |
// | zip | select | | |
// | apiRegion | filter | eq | |
// | userId | filter | eq | |
// | productId | filter | eq | |
// | subscriptionId | filter | eq | |
// | apiId | filter | eq | |
// | operationId | filter | eq | |
// | callCountSuccess | select | | |
// | callCountBlocked | select | | |
// | callCountFailed | select | | |
// | callCountOther | select | | |
// | bandwidth | select, orderBy | | |
// | cacheHitsCount | select | | |
// | cacheMissCount | select | | |
// | apiTimeAvg | select | | |
// | apiTimeMin | select | | |
// | apiTimeMax | select | | |
// | serviceTimeAvg | select | | |
// | serviceTimeMin | select | | |
// | serviceTimeMax | select | | |
// options - ReportsClientListByGeoOptions contains the optional parameters for the ReportsClient.ListByGeo method.
func (client *ReportsClient) NewListByGeoPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListByGeoOptions) *runtime.Pager[ReportsClientListByGeoResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByGeoResponse]{
		More: func(page ReportsClientListByGeoResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByGeoResponse) (ReportsClientListByGeoResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGeoCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListByGeoResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByGeoResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByGeoResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGeoHandleResponse(resp)
		},
	})
}

// listByGeoCreateRequest creates the ListByGeo request.
func (client *ReportsClient) listByGeoCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListByGeoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byGeo"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByGeoHandleResponse handles the ListByGeo response.
func (client *ReportsClient) listByGeoHandleResponse(resp *http.Response) (ReportsClientListByGeoResponse, error) {
	result := ReportsClientListByGeoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListByGeoResponse{}, err
	}
	return result, nil
}

// NewListByOperationPager - Lists report records by API Operations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter | ge, le | |
// | displayName | select, orderBy | | |
// | apiRegion | filter | eq | |
// | userId | filter | eq | |
// | productId | filter | eq | |
// | subscriptionId | filter | eq | |
// | apiId | filter | eq | |
// | operationId | select, filter | eq | |
// | callCountSuccess | select, orderBy | | |
// | callCountBlocked | select, orderBy | | |
// | callCountFailed | select, orderBy | | |
// | callCountOther | select, orderBy | | |
// | callCountTotal | select, orderBy | | |
// | bandwidth | select, orderBy | | |
// | cacheHitsCount | select | | |
// | cacheMissCount | select | | |
// | apiTimeAvg | select, orderBy | | |
// | apiTimeMin | select | | |
// | apiTimeMax | select | | |
// | serviceTimeAvg | select | | |
// | serviceTimeMin | select | | |
// | serviceTimeMax | select | | |
// options - ReportsClientListByOperationOptions contains the optional parameters for the ReportsClient.ListByOperation method.
func (client *ReportsClient) NewListByOperationPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListByOperationOptions) *runtime.Pager[ReportsClientListByOperationResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByOperationResponse]{
		More: func(page ReportsClientListByOperationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByOperationResponse) (ReportsClientListByOperationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByOperationCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListByOperationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByOperationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByOperationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByOperationHandleResponse(resp)
		},
	})
}

// listByOperationCreateRequest creates the ListByOperation request.
func (client *ReportsClient) listByOperationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListByOperationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byOperation"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByOperationHandleResponse handles the ListByOperation response.
func (client *ReportsClient) listByOperationHandleResponse(resp *http.Response) (ReportsClientListByOperationResponse, error) {
	result := ReportsClientListByOperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListByOperationResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - Lists report records by Product.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter | ge, le | |
// | displayName | select, orderBy | | |
// | apiRegion | filter | eq | |
// | userId | filter | eq | |
// | productId | select, filter | eq | |
// | subscriptionId | filter | eq | |
// | callCountSuccess | select, orderBy | | |
// | callCountBlocked | select, orderBy | | |
// | callCountFailed | select, orderBy | | |
// | callCountOther | select, orderBy | | |
// | callCountTotal | select, orderBy | | |
// | bandwidth | select, orderBy | | |
// | cacheHitsCount | select | | |
// | cacheMissCount | select | | |
// | apiTimeAvg | select, orderBy | | |
// | apiTimeMin | select | | |
// | apiTimeMax | select | | |
// | serviceTimeAvg | select | | |
// | serviceTimeMin | select | | |
// | serviceTimeMax | select | | |
// options - ReportsClientListByProductOptions contains the optional parameters for the ReportsClient.ListByProduct method.
func (client *ReportsClient) NewListByProductPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListByProductOptions) *runtime.Pager[ReportsClientListByProductResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByProductResponse]{
		More: func(page ReportsClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByProductResponse) (ReportsClientListByProductResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListByProductResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByProductResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByProductResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProductHandleResponse(resp)
		},
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ReportsClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byProduct"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *ReportsClient) listByProductHandleResponse(resp *http.Response) (ReportsClientListByProductResponse, error) {
	result := ReportsClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListByProductResponse{}, err
	}
	return result, nil
}

// NewListByRequestPager - Lists report records by Request.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter | ge, le | |
// | apiId | filter | eq | |
// | operationId | filter | eq | |
// | productId | filter | eq | |
// | userId | filter | eq | |
// | apiRegion | filter | eq | |
// | subscriptionId | filter | eq | |
// options - ReportsClientListByRequestOptions contains the optional parameters for the ReportsClient.ListByRequest method.
func (client *ReportsClient) NewListByRequestPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListByRequestOptions) *runtime.Pager[ReportsClientListByRequestResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByRequestResponse]{
		More: func(page ReportsClientListByRequestResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByRequestResponse) (ReportsClientListByRequestResponse, error) {
			req, err := client.listByRequestCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			if err != nil {
				return ReportsClientListByRequestResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByRequestResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByRequestResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRequestHandleResponse(resp)
		},
	})
}

// listByRequestCreateRequest creates the ListByRequest request.
func (client *ReportsClient) listByRequestCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListByRequestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byRequest"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRequestHandleResponse handles the ListByRequest response.
func (client *ReportsClient) listByRequestHandleResponse(resp *http.Response) (ReportsClientListByRequestResponse, error) {
	result := ReportsClientListByRequestResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestReportCollection); err != nil {
		return ReportsClientListByRequestResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists report records by subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter | ge, le | |
// | displayName | select, orderBy | | |
// | apiRegion | filter | eq | |
// | userId | select, filter | eq | |
// | productId | select, filter | eq | |
// | subscriptionId | select, filter | eq | |
// | callCountSuccess | select, orderBy | | |
// | callCountBlocked | select, orderBy | | |
// | callCountFailed | select, orderBy | | |
// | callCountOther | select, orderBy | | |
// | callCountTotal | select, orderBy | | |
// | bandwidth | select, orderBy | | |
// | cacheHitsCount | select | | |
// | cacheMissCount | select | | |
// | apiTimeAvg | select, orderBy | | |
// | apiTimeMin | select | | |
// | apiTimeMax | select | | |
// | serviceTimeAvg | select | | |
// | serviceTimeMin | select | | |
// | serviceTimeMax | select | | |
// options - ReportsClientListBySubscriptionOptions contains the optional parameters for the ReportsClient.ListBySubscription
// method.
func (client *ReportsClient) NewListBySubscriptionPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListBySubscriptionOptions) *runtime.Pager[ReportsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListBySubscriptionResponse]{
		More: func(page ReportsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListBySubscriptionResponse) (ReportsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ReportsClient) listBySubscriptionCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/bySubscription"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ReportsClient) listBySubscriptionHandleResponse(resp *http.Response) (ReportsClientListBySubscriptionResponse, error) {
	result := ReportsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListByTimePager - Lists report records by Time.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter, select | ge, le | |
// | interval | select | | |
// | apiRegion | filter | eq | |
// | userId | filter | eq | |
// | productId | filter | eq | |
// | subscriptionId | filter | eq | |
// | apiId | filter | eq | |
// | operationId | filter | eq | |
// | callCountSuccess | select | | |
// | callCountBlocked | select | | |
// | callCountFailed | select | | |
// | callCountOther | select | | |
// | bandwidth | select, orderBy | | |
// | cacheHitsCount | select | | |
// | cacheMissCount | select | | |
// | apiTimeAvg | select | | |
// | apiTimeMin | select | | |
// | apiTimeMax | select | | |
// | serviceTimeAvg | select | | |
// | serviceTimeMin | select | | |
// | serviceTimeMax | select | | |
// interval - By time interval. Interval must be multiple of 15 minutes and may not be zero. The value should be in ISO 8601
// format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert
// TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds)).
// options - ReportsClientListByTimeOptions contains the optional parameters for the ReportsClient.ListByTime method.
func (client *ReportsClient) NewListByTimePager(resourceGroupName string, serviceName string, filter string, interval string, options *ReportsClientListByTimeOptions) *runtime.Pager[ReportsClientListByTimeResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByTimeResponse]{
		More: func(page ReportsClientListByTimeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByTimeResponse) (ReportsClientListByTimeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTimeCreateRequest(ctx, resourceGroupName, serviceName, filter, interval, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListByTimeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByTimeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByTimeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTimeHandleResponse(resp)
		},
	})
}

// listByTimeCreateRequest creates the ListByTime request.
func (client *ReportsClient) listByTimeCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, interval string, options *ReportsClientListByTimeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byTime"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("interval", interval)
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTimeHandleResponse handles the ListByTime response.
func (client *ReportsClient) listByTimeHandleResponse(resp *http.Response) (ReportsClientListByTimeResponse, error) {
	result := ReportsClientListByTimeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListByTimeResponse{}, err
	}
	return result, nil
}

// NewListByUserPager - Lists report records by User.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// filter - | Field | Usage | Supported operators | Supported functions |
// |-------------|-------------|-------------|-------------|
// | timestamp | filter | ge, le | |
// | displayName | select, orderBy | | |
// | userId | select, filter | eq | |
// | apiRegion | filter | eq | |
// | productId | filter | eq | |
// | subscriptionId | filter | eq | |
// | apiId | filter | eq | |
// | operationId | filter | eq | |
// | callCountSuccess | select, orderBy | | |
// | callCountBlocked | select, orderBy | | |
// | callCountFailed | select, orderBy | | |
// | callCountOther | select, orderBy | | |
// | callCountTotal | select, orderBy | | |
// | bandwidth | select, orderBy | | |
// | cacheHitsCount | select | | |
// | cacheMissCount | select | | |
// | apiTimeAvg | select, orderBy | | |
// | apiTimeMin | select | | |
// | apiTimeMax | select | | |
// | serviceTimeAvg | select | | |
// | serviceTimeMin | select | | |
// | serviceTimeMax | select | | |
// options - ReportsClientListByUserOptions contains the optional parameters for the ReportsClient.ListByUser method.
func (client *ReportsClient) NewListByUserPager(resourceGroupName string, serviceName string, filter string, options *ReportsClientListByUserOptions) *runtime.Pager[ReportsClientListByUserResponse] {
	return runtime.NewPager(runtime.PageProcessor[ReportsClientListByUserResponse]{
		More: func(page ReportsClientListByUserResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListByUserResponse) (ReportsClientListByUserResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByUserCreateRequest(ctx, resourceGroupName, serviceName, filter, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListByUserResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListByUserResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListByUserResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByUserHandleResponse(resp)
		},
	})
}

// listByUserCreateRequest creates the ListByUser request.
func (client *ReportsClient) listByUserCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, filter string, options *ReportsClientListByUserOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byUser"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByUserHandleResponse handles the ListByUser response.
func (client *ReportsClient) listByUserHandleResponse(resp *http.Response) (ReportsClientListByUserResponse, error) {
	result := ReportsClientListByUserResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportCollection); err != nil {
		return ReportsClientListByUserResponse{}, err
	}
	return result, nil
}
