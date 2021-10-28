//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProviderClient contains the methods for the Provider group.
// Don't use this type directly, use NewProviderClient() instead.
type ProviderClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProviderClient creates a new instance of ProviderClient with the specified values.
func NewProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProviderClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ProviderClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetAvailableStacks - Description for Get available application frameworks and their versions
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) GetAvailableStacks(options *ProviderGetAvailableStacksOptions) *ProviderGetAvailableStacksPager {
	return &ProviderGetAvailableStacksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getAvailableStacksCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderGetAvailableStacksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationStackCollection.NextLink)
		},
	}
}

// getAvailableStacksCreateRequest creates the GetAvailableStacks request.
func (client *ProviderClient) getAvailableStacksCreateRequest(ctx context.Context, options *ProviderGetAvailableStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/availableStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OSTypeSelected != nil {
		reqQP.Set("osTypeSelected", string(*options.OSTypeSelected))
	}
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAvailableStacksHandleResponse handles the GetAvailableStacks response.
func (client *ProviderClient) getAvailableStacksHandleResponse(resp *http.Response) (ProviderGetAvailableStacksResponse, error) {
	result := ProviderGetAvailableStacksResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationStackCollection); err != nil {
		return ProviderGetAvailableStacksResponse{}, err
	}
	return result, nil
}

// getAvailableStacksHandleError handles the GetAvailableStacks error response.
func (client *ProviderClient) getAvailableStacksHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetAvailableStacksOnPrem - Description for Get available application frameworks and their versions
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) GetAvailableStacksOnPrem(options *ProviderGetAvailableStacksOnPremOptions) *ProviderGetAvailableStacksOnPremPager {
	return &ProviderGetAvailableStacksOnPremPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getAvailableStacksOnPremCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderGetAvailableStacksOnPremResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationStackCollection.NextLink)
		},
	}
}

// getAvailableStacksOnPremCreateRequest creates the GetAvailableStacksOnPrem request.
func (client *ProviderClient) getAvailableStacksOnPremCreateRequest(ctx context.Context, options *ProviderGetAvailableStacksOnPremOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OSTypeSelected != nil {
		reqQP.Set("osTypeSelected", string(*options.OSTypeSelected))
	}
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAvailableStacksOnPremHandleResponse handles the GetAvailableStacksOnPrem response.
func (client *ProviderClient) getAvailableStacksOnPremHandleResponse(resp *http.Response) (ProviderGetAvailableStacksOnPremResponse, error) {
	result := ProviderGetAvailableStacksOnPremResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationStackCollection); err != nil {
		return ProviderGetAvailableStacksOnPremResponse{}, err
	}
	return result, nil
}

// getAvailableStacksOnPremHandleError handles the GetAvailableStacksOnPrem error response.
func (client *ProviderClient) getAvailableStacksOnPremHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetFunctionAppStacks - Description for Get available Function app frameworks and their versions
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) GetFunctionAppStacks(options *ProviderGetFunctionAppStacksOptions) *ProviderGetFunctionAppStacksPager {
	return &ProviderGetFunctionAppStacksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getFunctionAppStacksCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderGetFunctionAppStacksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FunctionAppStackCollection.NextLink)
		},
	}
}

// getFunctionAppStacksCreateRequest creates the GetFunctionAppStacks request.
func (client *ProviderClient) getFunctionAppStacksCreateRequest(ctx context.Context, options *ProviderGetFunctionAppStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/functionAppStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFunctionAppStacksHandleResponse handles the GetFunctionAppStacks response.
func (client *ProviderClient) getFunctionAppStacksHandleResponse(resp *http.Response) (ProviderGetFunctionAppStacksResponse, error) {
	result := ProviderGetFunctionAppStacksResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FunctionAppStackCollection); err != nil {
		return ProviderGetFunctionAppStacksResponse{}, err
	}
	return result, nil
}

// getFunctionAppStacksHandleError handles the GetFunctionAppStacks error response.
func (client *ProviderClient) getFunctionAppStacksHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetFunctionAppStacksForLocation - Description for Get available Function app frameworks and their versions for location
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) GetFunctionAppStacksForLocation(location string, options *ProviderGetFunctionAppStacksForLocationOptions) *ProviderGetFunctionAppStacksForLocationPager {
	return &ProviderGetFunctionAppStacksForLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getFunctionAppStacksForLocationCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp ProviderGetFunctionAppStacksForLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.FunctionAppStackCollection.NextLink)
		},
	}
}

// getFunctionAppStacksForLocationCreateRequest creates the GetFunctionAppStacksForLocation request.
func (client *ProviderClient) getFunctionAppStacksForLocationCreateRequest(ctx context.Context, location string, options *ProviderGetFunctionAppStacksForLocationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/locations/{location}/functionAppStacks"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFunctionAppStacksForLocationHandleResponse handles the GetFunctionAppStacksForLocation response.
func (client *ProviderClient) getFunctionAppStacksForLocationHandleResponse(resp *http.Response) (ProviderGetFunctionAppStacksForLocationResponse, error) {
	result := ProviderGetFunctionAppStacksForLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FunctionAppStackCollection); err != nil {
		return ProviderGetFunctionAppStacksForLocationResponse{}, err
	}
	return result, nil
}

// getFunctionAppStacksForLocationHandleError handles the GetFunctionAppStacksForLocation error response.
func (client *ProviderClient) getFunctionAppStacksForLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetWebAppStacks - Description for Get available Web app frameworks and their versions
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) GetWebAppStacks(options *ProviderGetWebAppStacksOptions) *ProviderGetWebAppStacksPager {
	return &ProviderGetWebAppStacksPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getWebAppStacksCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderGetWebAppStacksResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebAppStackCollection.NextLink)
		},
	}
}

// getWebAppStacksCreateRequest creates the GetWebAppStacks request.
func (client *ProviderClient) getWebAppStacksCreateRequest(ctx context.Context, options *ProviderGetWebAppStacksOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/webAppStacks"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWebAppStacksHandleResponse handles the GetWebAppStacks response.
func (client *ProviderClient) getWebAppStacksHandleResponse(resp *http.Response) (ProviderGetWebAppStacksResponse, error) {
	result := ProviderGetWebAppStacksResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppStackCollection); err != nil {
		return ProviderGetWebAppStacksResponse{}, err
	}
	return result, nil
}

// getWebAppStacksHandleError handles the GetWebAppStacks error response.
func (client *ProviderClient) getWebAppStacksHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetWebAppStacksForLocation - Description for Get available Web app frameworks and their versions for location
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) GetWebAppStacksForLocation(location string, options *ProviderGetWebAppStacksForLocationOptions) *ProviderGetWebAppStacksForLocationPager {
	return &ProviderGetWebAppStacksForLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getWebAppStacksForLocationCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp ProviderGetWebAppStacksForLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebAppStackCollection.NextLink)
		},
	}
}

// getWebAppStacksForLocationCreateRequest creates the GetWebAppStacksForLocation request.
func (client *ProviderClient) getWebAppStacksForLocationCreateRequest(ctx context.Context, location string, options *ProviderGetWebAppStacksForLocationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/locations/{location}/webAppStacks"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StackOsType != nil {
		reqQP.Set("stackOsType", string(*options.StackOsType))
	}
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWebAppStacksForLocationHandleResponse handles the GetWebAppStacksForLocation response.
func (client *ProviderClient) getWebAppStacksForLocationHandleResponse(resp *http.Response) (ProviderGetWebAppStacksForLocationResponse, error) {
	result := ProviderGetWebAppStacksForLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppStackCollection); err != nil {
		return ProviderGetWebAppStacksForLocationResponse{}, err
	}
	return result, nil
}

// getWebAppStacksForLocationHandleError handles the GetWebAppStacksForLocation error response.
func (client *ProviderClient) getWebAppStacksForLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListOperations - Description for Gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric definitions
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *ProviderClient) ListOperations(options *ProviderListOperationsOptions) *ProviderListOperationsPager {
	return &ProviderListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProviderListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CsmOperationCollection.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *ProviderClient) listOperationsCreateRequest(ctx context.Context, options *ProviderListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Web/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *ProviderClient) listOperationsHandleResponse(resp *http.Response) (ProviderListOperationsResponse, error) {
	result := ProviderListOperationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmOperationCollection); err != nil {
		return ProviderListOperationsResponse{}, err
	}
	return result, nil
}

// listOperationsHandleError handles the ListOperations error response.
func (client *ProviderClient) listOperationsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
