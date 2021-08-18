//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztable

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"time"
)

type serviceClient struct {
	con *connection
}

// GetProperties - Gets the properties of an account's Table service, including properties for Analytics and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns the *TableServiceError error type.
func (client *serviceClient) GetProperties(ctx context.Context, options *ServiceGetPropertiesOptions) (TableServicePropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableServicePropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableServicePropertiesResponse{}, client.getPropertiesHandleError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *serviceClient) getPropertiesCreateRequest(ctx context.Context, options *ServiceGetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *serviceClient) getPropertiesHandleResponse(resp *azcore.Response) (TableServicePropertiesResponse, error) {
	var val *TableServiceProperties
	if err := resp.UnmarshalAsXML(&val); err != nil {
		return TableServicePropertiesResponse{}, err
	}
	result := TableServicePropertiesResponse{RawResponse: resp.Response, StorageServiceProperties: val}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// getPropertiesHandleError handles the GetProperties error response.
func (client *serviceClient) getPropertiesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetStatistics - Retrieves statistics related to replication for the Table service. It is only available on the secondary location endpoint when read-access
// geo-redundant replication is enabled for the account.
// If the operation fails it returns the *TableServiceError error type.
func (client *serviceClient) GetStatistics(ctx context.Context, options *ServiceGetStatisticsOptions) (TableServiceStatsResponse, error) {
	req, err := client.getStatisticsCreateRequest(ctx, options)
	if err != nil {
		return TableServiceStatsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableServiceStatsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableServiceStatsResponse{}, client.getStatisticsHandleError(resp)
	}
	return client.getStatisticsHandleResponse(resp)
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client *serviceClient) getStatisticsCreateRequest(ctx context.Context, options *ServiceGetStatisticsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "stats")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client *serviceClient) getStatisticsHandleResponse(resp *azcore.Response) (TableServiceStatsResponse, error) {
	var val *TableServiceStats
	if err := resp.UnmarshalAsXML(&val); err != nil {
		return TableServiceStatsResponse{}, err
	}
	result := TableServiceStatsResponse{RawResponse: resp.Response, StorageServiceStats: val}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return TableServiceStatsResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// getStatisticsHandleError handles the GetStatistics error response.
func (client *serviceClient) getStatisticsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// SetProperties - Sets properties for an account's Table service endpoint, including properties for Analytics and CORS (Cross-Origin Resource Sharing)
// rules.
// If the operation fails it returns the *TableServiceError error type.
func (client *serviceClient) SetProperties(ctx context.Context, tableServiceProperties TableServiceProperties, options *ServiceSetPropertiesOptions) (ServiceSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, tableServiceProperties, options)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return ServiceSetPropertiesResponse{}, client.setPropertiesHandleError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *serviceClient) setPropertiesCreateRequest(ctx context.Context, tableServiceProperties TableServiceProperties, options *ServiceSetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(tableServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *serviceClient) setPropertiesHandleResponse(resp *azcore.Response) (ServiceSetPropertiesResponse, error) {
	result := ServiceSetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

// setPropertiesHandleError handles the SetProperties error response.
func (client *serviceClient) setPropertiesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := TableServiceError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
