//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"time"
)

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use NewServiceClient() instead.
type ServiceClient struct {
	Con *connection
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
func NewServiceClient(con *connection) *ServiceClient {
	return &ServiceClient{Con: con}
}

// GetProperties - Gets the properties of an account's Table service, including properties for Analytics and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns the *TableServiceError error type.
func (client *ServiceClient) GetProperties(ctx context.Context, options *ServiceGetPropertiesOptions) (ServiceGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return ServiceGetPropertiesResponse{}, err
	}
	resp, err := 	client.Con.Pipeline().Do(req)
	if err != nil {
		return ServiceGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceGetPropertiesResponse{}, client.getPropertiesHandleError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *ServiceClient) getPropertiesCreateRequest(ctx context.Context, options *ServiceGetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.Con.Endpoint())
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *ServiceClient) getPropertiesHandleResponse(resp *http.Response) (ServiceGetPropertiesResponse, error) {
	result := ServiceGetPropertiesResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.TableServiceProperties); err != nil {
		return ServiceGetPropertiesResponse{}, err
	}
	return result, nil
}

// getPropertiesHandleError handles the GetProperties error response.
func (client *ServiceClient) getPropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := TableServiceError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetStatistics - Retrieves statistics related to replication for the Table service. It is only available on the secondary location endpoint when read-access
// geo-redundant replication is enabled for the account.
// If the operation fails it returns the *TableServiceError error type.
func (client *ServiceClient) GetStatistics(ctx context.Context, options *ServiceGetStatisticsOptions) (ServiceGetStatisticsResponse, error) {
	req, err := client.getStatisticsCreateRequest(ctx, options)
	if err != nil {
		return ServiceGetStatisticsResponse{}, err
	}
	resp, err := 	client.Con.Pipeline().Do(req)
	if err != nil {
		return ServiceGetStatisticsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceGetStatisticsResponse{}, client.getStatisticsHandleError(resp)
	}
	return client.getStatisticsHandleResponse(resp)
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client *ServiceClient) getStatisticsCreateRequest(ctx context.Context, options *ServiceGetStatisticsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.Con.Endpoint())
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "stats")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client *ServiceClient) getStatisticsHandleResponse(resp *http.Response) (ServiceGetStatisticsResponse, error) {
	result := ServiceGetStatisticsResponse{RawResponse: resp}
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
			return ServiceGetStatisticsResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.TableServiceStats); err != nil {
		return ServiceGetStatisticsResponse{}, err
	}
	return result, nil
}

// getStatisticsHandleError handles the GetStatistics error response.
func (client *ServiceClient) getStatisticsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := TableServiceError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// SetProperties - Sets properties for an account's Table service endpoint, including properties for Analytics and CORS (Cross-Origin Resource Sharing)
// rules.
// If the operation fails it returns the *TableServiceError error type.
func (client *ServiceClient) SetProperties(ctx context.Context, tableServiceProperties TableServiceProperties, options *ServiceSetPropertiesOptions) (ServiceSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, tableServiceProperties, options)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	resp, err := 	client.Con.Pipeline().Do(req)
	if err != nil {
		return ServiceSetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ServiceSetPropertiesResponse{}, client.setPropertiesHandleError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *ServiceClient) setPropertiesCreateRequest(ctx context.Context, tableServiceProperties TableServiceProperties, options *ServiceSetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.Con.Endpoint())
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", "2019-02-02")
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, runtime.MarshalAsXML(req, tableServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *ServiceClient) setPropertiesHandleResponse(resp *http.Response) (ServiceSetPropertiesResponse, error) {
	result := ServiceSetPropertiesResponse{RawResponse: resp}
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
func (client *ServiceClient) setPropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := TableServiceError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

