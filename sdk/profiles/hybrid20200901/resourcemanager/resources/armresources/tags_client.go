//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresources

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/profiles/hybrid20200901/internal"
	"net/http"
	"net/url"
	"strings"
)

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use NewTagsClient() instead.
type TagsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTagsClient creates a new instance of TagsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTagsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TagsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(internal.ModuleName, internal.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &TagsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - This operation allows adding a name to the list of predefined tag names for the given subscription. A
// tag name can have a maximum of 512 characters and is case-insensitive. Tag names cannot have the
// following prefixes which are reserved for Azure use: 'microsoft', 'azure', 'windows'.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// tagName - The name of the tag to create.
// options - TagsClientCreateOrUpdateOptions contains the optional parameters for the TagsClient.CreateOrUpdate method.
func (client *TagsClient) CreateOrUpdate(ctx context.Context, tagName string, options *TagsClientCreateOrUpdateOptions) (TagsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, tagName, options)
	if err != nil {
		return TagsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TagsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagsClient) createOrUpdateCreateRequest(ctx context.Context, tagName string, options *TagsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagsClient) createOrUpdateHandleResponse(resp *http.Response) (TagsClientCreateOrUpdateResponse, error) {
	result := TagsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagDetails); err != nil {
		return TagsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAtScope - This operation allows adding or replacing the entire set of tags on the specified resource or subscription.
// The specified entity can have a maximum of 50 tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// scope - The resource scope.
// options - TagsClientCreateOrUpdateAtScopeOptions contains the optional parameters for the TagsClient.CreateOrUpdateAtScope
// method.
func (client *TagsClient) CreateOrUpdateAtScope(ctx context.Context, scope string, parameters TagsResource, options *TagsClientCreateOrUpdateAtScopeOptions) (TagsClientCreateOrUpdateAtScopeResponse, error) {
	req, err := client.createOrUpdateAtScopeCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return TagsClientCreateOrUpdateAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientCreateOrUpdateAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsClientCreateOrUpdateAtScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateAtScopeHandleResponse(resp)
}

// createOrUpdateAtScopeCreateRequest creates the CreateOrUpdateAtScope request.
func (client *TagsClient) createOrUpdateAtScopeCreateRequest(ctx context.Context, scope string, parameters TagsResource, options *TagsClientCreateOrUpdateAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAtScopeHandleResponse handles the CreateOrUpdateAtScope response.
func (client *TagsClient) createOrUpdateAtScopeHandleResponse(resp *http.Response) (TagsClientCreateOrUpdateAtScopeResponse, error) {
	result := TagsClientCreateOrUpdateAtScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResource); err != nil {
		return TagsClientCreateOrUpdateAtScopeResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateValue - This operation allows adding a value to the list of predefined values for an existing predefined
// tag name. A tag value can have a maximum of 256 characters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// tagName - The name of the tag.
// tagValue - The value of the tag to create.
// options - TagsClientCreateOrUpdateValueOptions contains the optional parameters for the TagsClient.CreateOrUpdateValue
// method.
func (client *TagsClient) CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string, options *TagsClientCreateOrUpdateValueOptions) (TagsClientCreateOrUpdateValueResponse, error) {
	req, err := client.createOrUpdateValueCreateRequest(ctx, tagName, tagValue, options)
	if err != nil {
		return TagsClientCreateOrUpdateValueResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientCreateOrUpdateValueResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TagsClientCreateOrUpdateValueResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateValueHandleResponse(resp)
}

// createOrUpdateValueCreateRequest creates the CreateOrUpdateValue request.
func (client *TagsClient) createOrUpdateValueCreateRequest(ctx context.Context, tagName string, tagValue string, options *TagsClientCreateOrUpdateValueOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if tagValue == "" {
		return nil, errors.New("parameter tagValue cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateValueHandleResponse handles the CreateOrUpdateValue response.
func (client *TagsClient) createOrUpdateValueHandleResponse(resp *http.Response) (TagsClientCreateOrUpdateValueResponse, error) {
	result := TagsClientCreateOrUpdateValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagValue); err != nil {
		return TagsClientCreateOrUpdateValueResponse{}, err
	}
	return result, nil
}

// Delete - This operation allows deleting a name from the list of predefined tag names for the given subscription. The name
// being deleted must not be in use as a tag name for any resource. All predefined values
// for the given name must have already been deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// tagName - The name of the tag.
// options - TagsClientDeleteOptions contains the optional parameters for the TagsClient.Delete method.
func (client *TagsClient) Delete(ctx context.Context, tagName string, options *TagsClientDeleteOptions) (TagsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, tagName, options)
	if err != nil {
		return TagsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TagsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return TagsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TagsClient) deleteCreateRequest(ctx context.Context, tagName string, options *TagsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAtScope - Deletes the entire set of tags on a resource or subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// scope - The resource scope.
// options - TagsClientDeleteAtScopeOptions contains the optional parameters for the TagsClient.DeleteAtScope method.
func (client *TagsClient) DeleteAtScope(ctx context.Context, scope string, options *TagsClientDeleteAtScopeOptions) (TagsClientDeleteAtScopeResponse, error) {
	req, err := client.deleteAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsClientDeleteAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientDeleteAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsClientDeleteAtScopeResponse{}, runtime.NewResponseError(resp)
	}
	return TagsClientDeleteAtScopeResponse{}, nil
}

// deleteAtScopeCreateRequest creates the DeleteAtScope request.
func (client *TagsClient) deleteAtScopeCreateRequest(ctx context.Context, scope string, options *TagsClientDeleteAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteValue - This operation allows deleting a value from the list of predefined values for an existing predefined tag
// name. The value being deleted must not be in use as a tag value for the given tag name for any
// resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// tagName - The name of the tag.
// tagValue - The value of the tag to delete.
// options - TagsClientDeleteValueOptions contains the optional parameters for the TagsClient.DeleteValue method.
func (client *TagsClient) DeleteValue(ctx context.Context, tagName string, tagValue string, options *TagsClientDeleteValueOptions) (TagsClientDeleteValueResponse, error) {
	req, err := client.deleteValueCreateRequest(ctx, tagName, tagValue, options)
	if err != nil {
		return TagsClientDeleteValueResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientDeleteValueResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TagsClientDeleteValueResponse{}, runtime.NewResponseError(resp)
	}
	return TagsClientDeleteValueResponse{}, nil
}

// deleteValueCreateRequest creates the DeleteValue request.
func (client *TagsClient) deleteValueCreateRequest(ctx context.Context, tagName string, tagValue string, options *TagsClientDeleteValueOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	if tagName == "" {
		return nil, errors.New("parameter tagName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	if tagValue == "" {
		return nil, errors.New("parameter tagValue cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetAtScope - Gets the entire set of tags on a resource or subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// scope - The resource scope.
// options - TagsClientGetAtScopeOptions contains the optional parameters for the TagsClient.GetAtScope method.
func (client *TagsClient) GetAtScope(ctx context.Context, scope string, options *TagsClientGetAtScopeOptions) (TagsClientGetAtScopeResponse, error) {
	req, err := client.getAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsClientGetAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientGetAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsClientGetAtScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtScopeHandleResponse(resp)
}

// getAtScopeCreateRequest creates the GetAtScope request.
func (client *TagsClient) getAtScopeCreateRequest(ctx context.Context, scope string, options *TagsClientGetAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtScopeHandleResponse handles the GetAtScope response.
func (client *TagsClient) getAtScopeHandleResponse(resp *http.Response) (TagsClientGetAtScopeResponse, error) {
	result := TagsClientGetAtScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResource); err != nil {
		return TagsClientGetAtScopeResponse{}, err
	}
	return result, nil
}

// NewListPager - This operation performs a union of predefined tags, resource tags, resource group tags and subscription
// tags, and returns a summary of usage for each tag name and value under the given subscription.
// In case of a large number of tags, this operation may return a previously cached result.
// Generated from API version 2019-10-01
// options - TagsClientListOptions contains the optional parameters for the TagsClient.List method.
func (client *TagsClient) NewListPager(options *TagsClientListOptions) *runtime.Pager[TagsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagsClientListResponse]{
		More: func(page TagsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TagsClientListResponse) (TagsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TagsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TagsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TagsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TagsClient) listCreateRequest(ctx context.Context, options *TagsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TagsClient) listHandleResponse(resp *http.Response) (TagsClientListResponse, error) {
	result := TagsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsListResult); err != nil {
		return TagsClientListResponse{}, err
	}
	return result, nil
}

// UpdateAtScope - This operation allows replacing, merging or selectively deleting tags on the specified resource or subscription.
// The specified entity can have a maximum of 50 tags at the end of the operation. The
// 'replace' option replaces the entire set of existing tags with a new set. The 'merge' option allows adding tags with new
// names and updating the values of tags with existing names. The 'delete' option
// allows selectively deleting tags based on given names or name/value pairs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// scope - The resource scope.
// options - TagsClientUpdateAtScopeOptions contains the optional parameters for the TagsClient.UpdateAtScope method.
func (client *TagsClient) UpdateAtScope(ctx context.Context, scope string, parameters TagsPatchResource, options *TagsClientUpdateAtScopeOptions) (TagsClientUpdateAtScopeResponse, error) {
	req, err := client.updateAtScopeCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return TagsClientUpdateAtScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientUpdateAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagsClientUpdateAtScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateAtScopeHandleResponse(resp)
}

// updateAtScopeCreateRequest creates the UpdateAtScope request.
func (client *TagsClient) updateAtScopeCreateRequest(ctx context.Context, scope string, parameters TagsPatchResource, options *TagsClientUpdateAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateAtScopeHandleResponse handles the UpdateAtScope response.
func (client *TagsClient) updateAtScopeHandleResponse(resp *http.Response) (TagsClientUpdateAtScopeResponse, error) {
	result := TagsClientUpdateAtScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResource); err != nil {
		return TagsClientUpdateAtScopeResponse{}, err
	}
	return result, nil
}
