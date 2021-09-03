//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
func NewUsagesClient(con *arm.Connection, subscriptionID string) *UsagesClient {
	return &UsagesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListByLocation - Gets the current usage count and the limit for the resources of the location under the subscription.
// If the operation fails it returns a generic error.
func (client *UsagesClient) ListByLocation(ctx context.Context, location string, options *UsagesListByLocationOptions) (UsagesListByLocationResponse, error) {
	req, err := client.listByLocationCreateRequest(ctx, location, options)
	if err != nil {
		return UsagesListByLocationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UsagesListByLocationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UsagesListByLocationResponse{}, client.listByLocationHandleError(resp)
	}
	return client.listByLocationHandleResponse(resp)
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *UsagesClient) listByLocationCreateRequest(ctx context.Context, location string, options *UsagesListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *UsagesClient) listByLocationHandleResponse(resp *http.Response) (UsagesListByLocationResponse, error) {
	result := UsagesListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return UsagesListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *UsagesClient) listByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
