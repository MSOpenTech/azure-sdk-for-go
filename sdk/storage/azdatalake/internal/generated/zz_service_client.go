// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use a constructor function instead.
type ServiceClient struct {
	internal *azcore.Client
	endpoint string
}

// NewListFileSystemsPager - List filesystems and their properties in given account.
//
// Generated from API version 2024-05-04
//   - options - ServiceClientListFileSystemsOptions contains the optional parameters for the ServiceClient.NewListFileSystemsPager
//     method.
//
// ListFileSystemsCreateRequest creates the ListFileSystems request.
func (client *ServiceClient) ListFileSystemsCreateRequest(ctx context.Context, options *ServiceClientListFileSystemsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Continuation != nil {
		reqQP.Set("continuation", *options.Continuation)
	}
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxResults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	reqQP.Set("resource", "account")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = strings.Replace(reqQP.Encode(), "+", "%20", -1)
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["x-ms-version"] = []string{ServiceVersion}
	return req, nil
}

// listFileSystemsHandleResponse handles the ListFileSystems response.
func (client *ServiceClient) ListFileSystemsHandleResponse(resp *http.Response) (ServiceClientListFileSystemsResponse, error) {
	result := ServiceClientListFileSystemsResponse{}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-continuation"); val != "" {
		result.Continuation = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return ServiceClientListFileSystemsResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileSystemList); err != nil {
		return ServiceClientListFileSystemsResponse{}, err
	}
	return result, nil
}
