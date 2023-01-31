//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use NewServiceClient() instead.
type ServiceClient struct {
	endpoint string
	pl runtime.Pipeline
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
//   - endpoint - The URL of the service account, container, or blob that is the target of the desired operation.
//   - pl - the pipeline used for sending requests and handling responses.
func NewServiceClient(endpoint string, pl runtime.Pipeline) *ServiceClient {
	client := &ServiceClient{
		endpoint: endpoint,
		pl: pl,
	}
	return client
}

// FilterBlobs - The Filter Blobs operation enables callers to list blobs across all containers whose tags match a given search
// expression. Filter blobs searches across all containers within a storage account but can
// be scoped within the expression to a single container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - where - Filters the results to return only to return only blobs whose tags match the specified expression.
//   - options - ServiceClientFilterBlobsOptions contains the optional parameters for the ServiceClient.FilterBlobs method.
func (client *ServiceClient) FilterBlobs(ctx context.Context, where string, options *ServiceClientFilterBlobsOptions) (ServiceClientFilterBlobsResponse, error) {
	req, err := client.filterBlobsCreateRequest(ctx, where, options)
	if err != nil {
		return ServiceClientFilterBlobsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientFilterBlobsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientFilterBlobsResponse{}, runtime.NewResponseError(resp)
	}
	return client.filterBlobsHandleResponse(resp)
}

// filterBlobsCreateRequest creates the FilterBlobs request.
func (client *ServiceClient) filterBlobsCreateRequest(ctx context.Context, where string, options *ServiceClientFilterBlobsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "blobs")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	reqQP.Set("where", where)
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	req.Raw().URL.RawQuery = strings.Replace(reqQP.Encode(), "+", "%20", -1)
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// filterBlobsHandleResponse handles the FilterBlobs response.
func (client *ServiceClient) filterBlobsHandleResponse(resp *http.Response) (ServiceClientFilterBlobsResponse, error) {
	result := ServiceClientFilterBlobsResponse{}
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
			return ServiceClientFilterBlobsResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.FilterBlobSegment); err != nil {
		return ServiceClientFilterBlobsResponse{}, err
	}
	return result, nil
}

// GetAccountInfo - Returns the sku name and account kind
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - ServiceClientGetAccountInfoOptions contains the optional parameters for the ServiceClient.GetAccountInfo method.
func (client *ServiceClient) GetAccountInfo(ctx context.Context, options *ServiceClientGetAccountInfoOptions) (ServiceClientGetAccountInfoResponse, error) {
	req, err := client.getAccountInfoCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientGetAccountInfoResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientGetAccountInfoResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetAccountInfoResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAccountInfoHandleResponse(resp)
}

// getAccountInfoCreateRequest creates the GetAccountInfo request.
func (client *ServiceClient) getAccountInfoCreateRequest(ctx context.Context, options *ServiceClientGetAccountInfoOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "account")
	reqQP.Set("comp", "properties")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getAccountInfoHandleResponse handles the GetAccountInfo response.
func (client *ServiceClient) getAccountInfoHandleResponse(resp *http.Response) (ServiceClientGetAccountInfoResponse, error) {
	result := ServiceClientGetAccountInfoResponse{}
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
			return ServiceClientGetAccountInfoResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SKUName = (*SKUName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	if val := resp.Header.Get("x-ms-is-hns-enabled"); val != "" {
		isHierarchicalNamespaceEnabled, err := strconv.ParseBool(val)
		if err != nil {
			return ServiceClientGetAccountInfoResponse{}, err
		}
		result.IsHierarchicalNamespaceEnabled = &isHierarchicalNamespaceEnabled
	}
	return result, nil
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and
// CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - ServiceClientGetPropertiesOptions contains the optional parameters for the ServiceClient.GetProperties method.
func (client *ServiceClient) GetProperties(ctx context.Context, options *ServiceClientGetPropertiesOptions) (ServiceClientGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientGetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *ServiceClient) getPropertiesCreateRequest(ctx context.Context, options *ServiceClientGetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
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
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *ServiceClient) getPropertiesHandleResponse(resp *http.Response) (ServiceClientGetPropertiesResponse, error) {
	result := ServiceClientGetPropertiesResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.StorageServiceProperties); err != nil {
		return ServiceClientGetPropertiesResponse{}, err
	}
	return result, nil
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary
// location endpoint when read-access geo-redundant replication is enabled for the storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - ServiceClientGetStatisticsOptions contains the optional parameters for the ServiceClient.GetStatistics method.
func (client *ServiceClient) GetStatistics(ctx context.Context, options *ServiceClientGetStatisticsOptions) (ServiceClientGetStatisticsResponse, error) {
	req, err := client.getStatisticsCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientGetStatisticsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientGetStatisticsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetStatisticsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStatisticsHandleResponse(resp)
}

// getStatisticsCreateRequest creates the GetStatistics request.
func (client *ServiceClient) getStatisticsCreateRequest(ctx context.Context, options *ServiceClientGetStatisticsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
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
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getStatisticsHandleResponse handles the GetStatistics response.
func (client *ServiceClient) getStatisticsHandleResponse(resp *http.Response) (ServiceClientGetStatisticsResponse, error) {
	result := ServiceClientGetStatisticsResponse{}
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
			return ServiceClientGetStatisticsResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.StorageServiceStats); err != nil {
		return ServiceClientGetStatisticsResponse{}, err
	}
	return result, nil
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using
// bearer token authentication.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - keyInfo - Key information
//   - options - ServiceClientGetUserDelegationKeyOptions contains the optional parameters for the ServiceClient.GetUserDelegationKey
//     method.
func (client *ServiceClient) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, options *ServiceClientGetUserDelegationKeyOptions) (ServiceClientGetUserDelegationKeyResponse, error) {
	req, err := client.getUserDelegationKeyCreateRequest(ctx, keyInfo, options)
	if err != nil {
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetUserDelegationKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUserDelegationKeyHandleResponse(resp)
}

// getUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client *ServiceClient) getUserDelegationKeyCreateRequest(ctx context.Context, keyInfo KeyInfo, options *ServiceClientGetUserDelegationKeyOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "service")
	reqQP.Set("comp", "userdelegationkey")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, runtime.MarshalAsXML(req, keyInfo)
}

// getUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client *ServiceClient) getUserDelegationKeyHandleResponse(resp *http.Response) (ServiceClientGetUserDelegationKeyResponse, error) {
	result := ServiceClientGetUserDelegationKeyResponse{}
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
			return ServiceClientGetUserDelegationKeyResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.UserDelegationKey); err != nil {
		return ServiceClientGetUserDelegationKeyResponse{}, err
	}
	return result, nil
}

// NewListContainersSegmentPager - The List Containers Segment operation returns a list of the containers under the specified
// account
//
// Generated from API version 2020-10-02
//   - options - ServiceClientListContainersSegmentOptions contains the optional parameters for the ServiceClient.NewListContainersSegmentPager
//     method.
// listContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client *ServiceClient) ListContainersSegmentCreateRequest(ctx context.Context, options *ServiceClientListContainersSegmentOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// listContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client *ServiceClient) ListContainersSegmentHandleResponse(resp *http.Response) (ServiceClientListContainersSegmentResponse, error) {
	result := ServiceClientListContainersSegmentResponse{}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ListContainersSegmentResponse); err != nil {
		return ServiceClientListContainersSegmentResponse{}, err
	}
	return result, nil
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - storageServiceProperties - The StorageService properties.
//   - options - ServiceClientSetPropertiesOptions contains the optional parameters for the ServiceClient.SetProperties method.
func (client *ServiceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceClientSetPropertiesOptions) (ServiceClientSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, storageServiceProperties, options)
	if err != nil {
		return ServiceClientSetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientSetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ServiceClientSetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *ServiceClient) setPropertiesCreateRequest(ctx context.Context, storageServiceProperties StorageServiceProperties, options *ServiceClientSetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
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
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, runtime.MarshalAsXML(req, storageServiceProperties)
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *ServiceClient) setPropertiesHandleResponse(resp *http.Response) (ServiceClientSetPropertiesResponse, error) {
	result := ServiceClientSetPropertiesResponse{}
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

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - contentLength - The length of the request.
//   - multipartContentType - Required. The value of this header must be multipart/mixed with a batch boundary. Example header
//     value: multipart/mixed; boundary=batch_
//   - body - Initial data
//   - options - ServiceClientSubmitBatchOptions contains the optional parameters for the ServiceClient.SubmitBatch method.
func (client *ServiceClient) SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body io.ReadSeekCloser, options *ServiceClientSubmitBatchOptions) (ServiceClientSubmitBatchResponse, error) {
	req, err := client.submitBatchCreateRequest(ctx, contentLength, multipartContentType, body, options)
	if err != nil {
		return ServiceClientSubmitBatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientSubmitBatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientSubmitBatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.submitBatchHandleResponse(resp)
}

// submitBatchCreateRequest creates the SubmitBatch request.
func (client *ServiceClient) submitBatchCreateRequest(ctx context.Context, contentLength int64, multipartContentType string, body io.ReadSeekCloser, options *ServiceClientSubmitBatchOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "batch")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Content-Length"] = []string{strconv.FormatInt(contentLength, 10)}
	req.Raw().Header["Content-Type"] = []string{multipartContentType}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, req.SetBody(body, "application/xml")
}

// submitBatchHandleResponse handles the SubmitBatch response.
func (client *ServiceClient) submitBatchHandleResponse(resp *http.Response) (ServiceClientSubmitBatchResponse, error) {
	result := ServiceClientSubmitBatchResponse{Body: resp.Body}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

