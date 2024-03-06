//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azmetrics

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client contains the methods for the Metrics group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	endpoint string
}

// QueryResources - Lists the metric values for multiple resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - subscriptionID - The subscription identifier for the resources in this batch.
//   - metricNamespace - Metric namespace that contains the requested metric names.
//   - metricNames - The names of the metrics (comma separated) to retrieve.
//   - resourceIDs - Metrics batch body including the list of resource ids
//   - options - QueryResourcesOptions contains the optional parameters for the Client.QueryResources method.
func (client *Client) QueryResources(ctx context.Context, subscriptionID string, metricNamespace string, metricNames []string, resourceIDs ResourceIDList, options *QueryResourcesOptions) (QueryResourcesResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "Client.QueryResources", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.queryResourcesCreateRequest(ctx, subscriptionID, metricNamespace, metricNames, resourceIDs, options)
	if err != nil {
		return QueryResourcesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueryResourcesResponse{}, err
	}
	resp, err := client.queryResourcesHandleResponse(httpResp)
	return resp, err
}

// queryResourcesCreateRequest creates the QueryResources request.
func (client *Client) queryResourcesCreateRequest(ctx context.Context, subscriptionID string, metricNamespace string, metricNames []string, resourceIDs ResourceIDList, options *QueryResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/metrics:getBatch"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.StartTime != nil {
		reqQP.Set("starttime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endtime", *options.EndTime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	reqQP.Set("metricnamespace", metricNamespace)
	reqQP.Set("metricnames", strings.Join(metricNames, ","))
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderby", *options.OrderBy)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.RollUpBy != nil {
		reqQP.Set("rollupby", *options.RollUpBy)
	}
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resourceIDs); err != nil {
		return nil, err
	}
	return req, nil
}

// queryResourcesHandleResponse handles the QueryResources response.
func (client *Client) queryResourcesHandleResponse(resp *http.Response) (QueryResourcesResponse, error) {
	result := QueryResourcesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricResults); err != nil {
		return QueryResourcesResponse{}, err
	}
	return result, nil
}
