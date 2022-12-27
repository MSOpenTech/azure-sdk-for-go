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
	"encoding/xml"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// QueueClient contains the methods for the Queue group.
// Don't use this type directly, use NewQueueClient() instead.
type QueueClient struct {
	endpoint string
	pl runtime.Pipeline
}

// NewQueueClient creates a new instance of QueueClient with the specified values.
// endpoint - The URL of the service account, queue or message that is the target of the desired operation.
// pl - the pipeline used for sending requests and handling responses.
func NewQueueClient(endpoint string, pl runtime.Pipeline) *QueueClient {
	client := &QueueClient{
		endpoint: endpoint,
		pl: pl,
	}
	return client
}

// Create - creates a new queue under the given account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-28
// queueName - The queue name.
// options - QueueClientCreateOptions contains the optional parameters for the QueueClient.Create method.
func (client *QueueClient) Create(ctx context.Context, queueName string, options *QueueClientCreateOptions) (QueueClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, queueName, options)
	if err != nil {
		return QueueClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated, http.StatusNoContent) {
		return QueueClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *QueueClient) createCreateRequest(ctx context.Context, queueName string, options *QueueClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{queueName}"
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.Metadata != nil {
		for k, v := range options.Metadata {
			req.Raw().Header["x-ms-meta-"+k] = []string{v}
		}
	}
	req.Raw().Header["x-ms-version"] = []string{"2018-03-28"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *QueueClient) createHandleResponse(resp *http.Response) (QueueClientCreateResponse, error) {
	result := QueueClientCreateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return QueueClientCreateResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// Delete - operation permanently deletes the specified queue
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-28
// queueName - The queue name.
// options - QueueClientDeleteOptions contains the optional parameters for the QueueClient.Delete method.
func (client *QueueClient) Delete(ctx context.Context, queueName string, options *QueueClientDeleteOptions) (QueueClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, queueName, options)
	if err != nil {
		return QueueClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return QueueClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *QueueClient) deleteCreateRequest(ctx context.Context, queueName string, options *QueueClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{queueName}"
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2018-03-28"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *QueueClient) deleteHandleResponse(resp *http.Response) (QueueClientDeleteResponse, error) {
	result := QueueClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return QueueClientDeleteResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// GetAccessPolicy - returns details about any stored access policies specified on the queue that may be used with Shared
// Access Signatures.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-28
// queueName - The queue name.
// options - QueueClientGetAccessPolicyOptions contains the optional parameters for the QueueClient.GetAccessPolicy method.
func (client *QueueClient) GetAccessPolicy(ctx context.Context, queueName string, options *QueueClientGetAccessPolicyOptions) (QueueClientGetAccessPolicyResponse, error) {
	req, err := client.getAccessPolicyCreateRequest(ctx, queueName, options)
	if err != nil {
		return QueueClientGetAccessPolicyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueClientGetAccessPolicyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueueClientGetAccessPolicyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAccessPolicyHandleResponse(resp)
}

// getAccessPolicyCreateRequest creates the GetAccessPolicy request.
func (client *QueueClient) getAccessPolicyCreateRequest(ctx context.Context, queueName string, options *QueueClientGetAccessPolicyOptions) (*policy.Request, error) {
	urlPath := "/{queueName}"
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "acl")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2018-03-28"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getAccessPolicyHandleResponse handles the GetAccessPolicy response.
func (client *QueueClient) getAccessPolicyHandleResponse(resp *http.Response) (QueueClientGetAccessPolicyResponse, error) {
	result := QueueClientGetAccessPolicyResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return QueueClientGetAccessPolicyResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result); err != nil {
		return QueueClientGetAccessPolicyResponse{}, err
	}
	return result, nil
}

// GetProperties - Retrieves user-defined metadata and queue properties on the specified queue. Metadata is associated with
// the queue as name-values pairs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-28
// queueName - The queue name.
// options - QueueClientGetPropertiesOptions contains the optional parameters for the QueueClient.GetProperties method.
func (client *QueueClient) GetProperties(ctx context.Context, queueName string, options *QueueClientGetPropertiesOptions) (QueueClientGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, queueName, options)
	if err != nil {
		return QueueClientGetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueClientGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueueClientGetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *QueueClient) getPropertiesCreateRequest(ctx context.Context, queueName string, options *QueueClientGetPropertiesOptions) (*policy.Request, error) {
	urlPath := "/{queueName}"
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "metadata")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2018-03-28"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *QueueClient) getPropertiesHandleResponse(resp *http.Response) (QueueClientGetPropertiesResponse, error) {
	result := QueueClientGetPropertiesResponse{}
	for hh := range resp.Header {
		if len(hh) > len("x-ms-meta-") && strings.EqualFold(hh[:len("x-ms-meta-")], "x-ms-meta-") {
			if result.Metadata == nil {
				result.Metadata = map[string]string{}
			}
			result.Metadata[hh[len("x-ms-meta-"):]] = resp.Header.Get(hh)
		}
	}
	if val := resp.Header.Get("x-ms-approximate-messages-count"); val != "" {
		approximateMessagesCount32, err := strconv.ParseInt(val, 10, 32)
		approximateMessagesCount := int32(approximateMessagesCount32)
		if err != nil {
			return QueueClientGetPropertiesResponse{}, err
		}
		result.ApproximateMessagesCount = &approximateMessagesCount
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
			return QueueClientGetPropertiesResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// SetAccessPolicy - sets stored access policies for the queue that may be used with Shared Access Signatures
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-28
// queueName - The queue name.
// queueACL - the acls for the queue
// options - QueueClientSetAccessPolicyOptions contains the optional parameters for the QueueClient.SetAccessPolicy method.
func (client *QueueClient) SetAccessPolicy(ctx context.Context, queueName string, queueACL []*SignedIdentifier, options *QueueClientSetAccessPolicyOptions) (QueueClientSetAccessPolicyResponse, error) {
	req, err := client.setAccessPolicyCreateRequest(ctx, queueName, queueACL, options)
	if err != nil {
		return QueueClientSetAccessPolicyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueClientSetAccessPolicyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return QueueClientSetAccessPolicyResponse{}, runtime.NewResponseError(resp)
	}
	return client.setAccessPolicyHandleResponse(resp)
}

// setAccessPolicyCreateRequest creates the SetAccessPolicy request.
func (client *QueueClient) setAccessPolicyCreateRequest(ctx context.Context, queueName string, queueACL []*SignedIdentifier, options *QueueClientSetAccessPolicyOptions) (*policy.Request, error) {
	urlPath := "/{queueName}"
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "acl")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2018-03-28"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	type wrapper struct {
		XMLName xml.Name `xml:"SignedIdentifiers"`
		QueueACL *[]*SignedIdentifier `xml:"SignedIdentifier"`
	}
	return req, runtime.MarshalAsXML(req, wrapper{QueueACL: &queueACL})
}

// setAccessPolicyHandleResponse handles the SetAccessPolicy response.
func (client *QueueClient) setAccessPolicyHandleResponse(resp *http.Response) (QueueClientSetAccessPolicyResponse, error) {
	result := QueueClientSetAccessPolicyResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return QueueClientSetAccessPolicyResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// SetMetadata - sets user-defined metadata on the specified queue. Metadata is associated with the queue as name-value pairs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-28
// queueName - The queue name.
// options - QueueClientSetMetadataOptions contains the optional parameters for the QueueClient.SetMetadata method.
func (client *QueueClient) SetMetadata(ctx context.Context, queueName string, options *QueueClientSetMetadataOptions) (QueueClientSetMetadataResponse, error) {
	req, err := client.setMetadataCreateRequest(ctx, queueName, options)
	if err != nil {
		return QueueClientSetMetadataResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueueClientSetMetadataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return QueueClientSetMetadataResponse{}, runtime.NewResponseError(resp)
	}
	return client.setMetadataHandleResponse(resp)
}

// setMetadataCreateRequest creates the SetMetadata request.
func (client *QueueClient) setMetadataCreateRequest(ctx context.Context, queueName string, options *QueueClientSetMetadataOptions) (*policy.Request, error) {
	urlPath := "/{queueName}"
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "metadata")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.Metadata != nil {
		for k, v := range options.Metadata {
			req.Raw().Header["x-ms-meta-"+k] = []string{v}
		}
	}
	req.Raw().Header["x-ms-version"] = []string{"2018-03-28"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// setMetadataHandleResponse handles the SetMetadata response.
func (client *QueueClient) setMetadataHandleResponse(resp *http.Response) (QueueClientSetMetadataResponse, error) {
	result := QueueClientSetMetadataResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return QueueClientSetMetadataResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

