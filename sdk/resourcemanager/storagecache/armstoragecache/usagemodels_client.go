//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

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

// UsageModelsClient contains the methods for the UsageModels group.
// Don't use this type directly, use NewUsageModelsClient() instead.
type UsageModelsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUsageModelsClient creates a new instance of UsageModelsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUsageModelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsageModelsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UsageModelsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get the list of cache usage models available to this subscription.
//
// Generated from API version 2024-03-01
//   - options - UsageModelsClientListOptions contains the optional parameters for the UsageModelsClient.NewListPager method.
func (client *UsageModelsClient) NewListPager(options *UsageModelsClientListOptions) *runtime.Pager[UsageModelsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsageModelsClientListResponse]{
		More: func(page UsageModelsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UsageModelsClientListResponse) (UsageModelsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UsageModelsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return UsageModelsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *UsageModelsClient) listCreateRequest(ctx context.Context, options *UsageModelsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/usageModels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageModelsClient) listHandleResponse(resp *http.Response) (UsageModelsClientListResponse, error) {
	result := UsageModelsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageModelsResult); err != nil {
		return UsageModelsClientListResponse{}, err
	}
	return result, nil
}
