//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchangeanalysis

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
	"time"
)

// ResourceChangesClient contains the methods for the ResourceChanges group.
// Don't use this type directly, use NewResourceChangesClient() instead.
type ResourceChangesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewResourceChangesClient creates a new instance of ResourceChangesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceChangesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceChangesClient, error) {
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
	client := &ResourceChangesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - List the changes of a resource within the specified time range. Customer data will be masked if the user
// doesn't have access.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The identifier of the resource.
// startTime - Specifies the start time of the changes request.
// endTime - Specifies the end time of the changes request.
// options - ResourceChangesClientListOptions contains the optional parameters for the ResourceChangesClient.List method.
func (client *ResourceChangesClient) NewListPager(resourceID string, startTime time.Time, endTime time.Time, options *ResourceChangesClientListOptions) *runtime.Pager[ResourceChangesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourceChangesClientListResponse]{
		More: func(page ResourceChangesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceChangesClientListResponse) (ResourceChangesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceID, startTime, endTime, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceChangesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceChangesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceChangesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ResourceChangesClient) listCreateRequest(ctx context.Context, resourceID string, startTime time.Time, endTime time.Time, options *ResourceChangesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.ChangeAnalysis/resourceChanges"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", url.PathEscape(resourceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	reqQP.Set("$startTime", startTime.Format(time.RFC3339Nano))
	reqQP.Set("$endTime", endTime.Format(time.RFC3339Nano))
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceChangesClient) listHandleResponse(resp *http.Response) (ResourceChangesClientListResponse, error) {
	result := ResourceChangesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChangeList); err != nil {
		return ResourceChangesClientListResponse{}, err
	}
	return result, nil
}
