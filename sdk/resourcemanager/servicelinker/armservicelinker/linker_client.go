//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

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

// LinkerClient contains the methods for the Linker group.
// Don't use this type directly, use NewLinkerClient() instead.
type LinkerClient struct {
	internal *arm.Client
}

// NewLinkerClient creates a new instance of LinkerClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLinkerClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LinkerClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update Linker resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - linkerName - The name Linker resource.
//   - parameters - Linker details.
//   - options - LinkerClientBeginCreateOrUpdateOptions contains the optional parameters for the LinkerClient.BeginCreateOrUpdate
//     method.
func (client *LinkerClient) BeginCreateOrUpdate(ctx context.Context, resourceURI string, linkerName string, parameters LinkerResource, options *LinkerClientBeginCreateOrUpdateOptions) (*runtime.Poller[LinkerClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceURI, linkerName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LinkerClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LinkerClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update Linker resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *LinkerClient) createOrUpdate(ctx context.Context, resourceURI string, linkerName string, parameters LinkerResource, options *LinkerClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "LinkerClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, linkerName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkerClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, linkerName string, parameters LinkerResource, options *LinkerClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - linkerName - The name Linker resource.
//   - options - LinkerClientBeginDeleteOptions contains the optional parameters for the LinkerClient.BeginDelete method.
func (client *LinkerClient) BeginDelete(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginDeleteOptions) (*runtime.Poller[LinkerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceURI, linkerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LinkerClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LinkerClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *LinkerClient) deleteOperation(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "LinkerClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkerClient) deleteCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns Linker resource for a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - linkerName - The name Linker resource.
//   - options - LinkerClientGetOptions contains the optional parameters for the LinkerClient.Get method.
func (client *LinkerClient) Get(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientGetOptions) (LinkerClientGetResponse, error) {
	var err error
	const operationName = "LinkerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return LinkerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LinkerClient) getCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkerClient) getHandleResponse(resp *http.Response) (LinkerClientGetResponse, error) {
	result := LinkerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkerResource); err != nil {
		return LinkerClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of Linkers which connects to the resource. which supports to config both application and target
// service during the resource provision.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - options - LinkerClientListOptions contains the optional parameters for the LinkerClient.NewListPager method.
func (client *LinkerClient) NewListPager(resourceURI string, options *LinkerClientListOptions) *runtime.Pager[LinkerClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkerClientListResponse]{
		More: func(page LinkerClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkerClientListResponse) (LinkerClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LinkerClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return LinkerClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LinkerClient) listCreateRequest(ctx context.Context, resourceURI string, options *LinkerClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LinkerClient) listHandleResponse(resp *http.Response) (LinkerClientListResponse, error) {
	result := LinkerClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceList); err != nil {
		return LinkerClientListResponse{}, err
	}
	return result, nil
}

// ListConfigurations - list source configurations for a Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - linkerName - The name Linker resource.
//   - options - LinkerClientListConfigurationsOptions contains the optional parameters for the LinkerClient.ListConfigurations
//     method.
func (client *LinkerClient) ListConfigurations(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientListConfigurationsOptions) (LinkerClientListConfigurationsResponse, error) {
	var err error
	const operationName = "LinkerClient.ListConfigurations"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listConfigurationsCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkerClientListConfigurationsResponse{}, err
	}
	resp, err := client.listConfigurationsHandleResponse(httpResp)
	return resp, err
}

// listConfigurationsCreateRequest creates the ListConfigurations request.
func (client *LinkerClient) listConfigurationsCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientListConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}/listConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConfigurationsHandleResponse handles the ListConfigurations response.
func (client *LinkerClient) listConfigurationsHandleResponse(resp *http.Response) (LinkerClientListConfigurationsResponse, error) {
	result := LinkerClientListConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationResult); err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Operation to update an existing Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - linkerName - The name Linker resource.
//   - parameters - Linker details.
//   - options - LinkerClientBeginUpdateOptions contains the optional parameters for the LinkerClient.BeginUpdate method.
func (client *LinkerClient) BeginUpdate(ctx context.Context, resourceURI string, linkerName string, parameters LinkerPatch, options *LinkerClientBeginUpdateOptions) (*runtime.Poller[LinkerClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceURI, linkerName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LinkerClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LinkerClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Operation to update an existing Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *LinkerClient) update(ctx context.Context, resourceURI string, linkerName string, parameters LinkerPatch, options *LinkerClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "LinkerClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceURI, linkerName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *LinkerClient) updateCreateRequest(ctx context.Context, resourceURI string, linkerName string, parameters LinkerPatch, options *LinkerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginValidate - Validate a Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - linkerName - The name Linker resource.
//   - options - LinkerClientBeginValidateOptions contains the optional parameters for the LinkerClient.BeginValidate method.
func (client *LinkerClient) BeginValidate(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*runtime.Poller[LinkerClientValidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validate(ctx, resourceURI, linkerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LinkerClientValidateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LinkerClientValidateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Validate - Validate a Linker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *LinkerClient) validate(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*http.Response, error) {
	var err error
	const operationName = "LinkerClient.BeginValidate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// validateCreateRequest creates the Validate request.
func (client *LinkerClient) validateCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}/validateLinker"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
