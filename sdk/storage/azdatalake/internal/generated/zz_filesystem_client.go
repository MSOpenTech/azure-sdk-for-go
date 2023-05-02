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
	"net/http"
	"strconv"
	"strings"
	"time"
)

// FileSystemClient contains the methods for the FileSystem group.
// Don't use this type directly, use NewFileSystemClient() instead.
type FileSystemClient struct {
	endpoint string
	pl runtime.Pipeline
}

// NewFileSystemClient creates a new instance of FileSystemClient with the specified values.
//   - endpoint - The URL of the service account, container, or blob that is the target of the desired operation.
//   - pl - the pipeline used for sending requests and handling responses.
func NewFileSystemClient(endpoint string, pl runtime.Pipeline) *FileSystemClient {
	client := &FileSystemClient{
		endpoint: endpoint,
		pl: pl,
	}
	return client
}

// Create - Create a FileSystem rooted at the specified location. If the FileSystem already exists, the operation fails. This
// operation does not support conditional HTTP requests.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - FileSystemClientCreateOptions contains the optional parameters for the FileSystemClient.Create method.
func (client *FileSystemClient) Create(ctx context.Context, options *FileSystemClientCreateOptions) (FileSystemClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, options)
	if err != nil {
		return FileSystemClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileSystemClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return FileSystemClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *FileSystemClient) createCreateRequest(ctx context.Context, options *FileSystemClientCreateOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("resource", "filesystem")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.Properties != nil {
		req.Raw().Header["x-ms-properties"] = []string{*options.Properties}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *FileSystemClient) createHandleResponse(resp *http.Response) (FileSystemClientCreateResponse, error) {
	result := FileSystemClientCreateResponse{}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientCreateResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-namespace-enabled"); val != "" {
		result.NamespaceEnabled = &val
	}
	return result, nil
}

// Delete - Marks the FileSystem for deletion. When a FileSystem is deleted, a FileSystem with the same identifier cannot
// be created for at least 30 seconds. While the filesystem is being deleted, attempts to
// create a filesystem with the same identifier will fail with status code 409 (Conflict), with the service returning additional
// error information indicating that the filesystem is being deleted. All
// other operations, including operations on any files or directories within the filesystem, will fail with status code 404
// (Not Found) while the filesystem is being deleted. This operation supports
// conditional HTTP requests. For more information, see Specifying Conditional Headers for Blob Service Operations
// [https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations].
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - FileSystemClientDeleteOptions contains the optional parameters for the FileSystemClient.Delete method.
//   - ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the FileSystemClient.SetProperties
//     method.
func (client *FileSystemClient) Delete(ctx context.Context, options *FileSystemClientDeleteOptions, modifiedAccessConditions *ModifiedAccessConditions) (FileSystemClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, options, modifiedAccessConditions)
	if err != nil {
		return FileSystemClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileSystemClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return FileSystemClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *FileSystemClient) deleteCreateRequest(ctx context.Context, options *FileSystemClientDeleteOptions, modifiedAccessConditions *ModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodDelete, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("resource", "filesystem")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123)}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *FileSystemClient) deleteHandleResponse(resp *http.Response) (FileSystemClientDeleteResponse, error) {
	result := FileSystemClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientDeleteResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// GetProperties - All system and user-defined filesystem properties are specified in the response headers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - FileSystemClientGetPropertiesOptions contains the optional parameters for the FileSystemClient.GetProperties
//     method.
func (client *FileSystemClient) GetProperties(ctx context.Context, options *FileSystemClientGetPropertiesOptions) (FileSystemClientGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return FileSystemClientGetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileSystemClientGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileSystemClientGetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *FileSystemClient) getPropertiesCreateRequest(ctx context.Context, options *FileSystemClientGetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodHead, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("resource", "filesystem")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *FileSystemClient) getPropertiesHandleResponse(resp *http.Response) (FileSystemClientGetPropertiesResponse, error) {
	result := FileSystemClientGetPropertiesResponse{}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientGetPropertiesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientGetPropertiesResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-properties"); val != "" {
		result.Properties = &val
	}
	if val := resp.Header.Get("x-ms-namespace-enabled"); val != "" {
		result.NamespaceEnabled = &val
	}
	return result, nil
}

// NewListBlobHierarchySegmentPager - The List Blobs operation returns a list of the blobs under the specified container
//
// Generated from API version 2020-10-02
//   - options - FileSystemClientListBlobHierarchySegmentOptions contains the optional parameters for the FileSystemClient.NewListBlobHierarchySegmentPager
//     method.
//
// ListBlobHierarchySegmentCreateRequest creates the ListBlobHierarchySegment request.
func (client *FileSystemClient) ListBlobHierarchySegmentCreateRequest(ctx context.Context, options *FileSystemClientListBlobHierarchySegmentOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "container")
	reqQP.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Delimiter != nil {
		reqQP.Set("delimiter", *options.Delimiter)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxResults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	if options != nil && options.Showonly != nil {
		reqQP.Set("showonly", "deleted")
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

// listBlobHierarchySegmentHandleResponse handles the ListBlobHierarchySegment response.
func (client *FileSystemClient) ListBlobHierarchySegmentHandleResponse(resp *http.Response) (FileSystemClientListBlobHierarchySegmentResponse, error) {
	result := FileSystemClientListBlobHierarchySegmentResponse{}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
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
			return FileSystemClientListBlobHierarchySegmentResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ListBlobsHierarchySegmentResponse); err != nil {
		return FileSystemClientListBlobHierarchySegmentResponse{}, err
	}
	return result, nil
}

// NewListPathsPager - List FileSystem paths and their properties.
//
// Generated from API version 2020-10-02
//   - recursive - Required
//   - options - FileSystemClientListPathsOptions contains the optional parameters for the FileSystemClient.NewListPathsPager
//     method.
//
// ListPathsCreateRequest creates the ListPaths request.
func (client *FileSystemClient) ListPathsCreateRequest(ctx context.Context, recursive bool, options *FileSystemClientListPathsOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("resource", "filesystem")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Continuation != nil {
		reqQP.Set("continuation", *options.Continuation)
	}
	if options != nil && options.Path != nil {
		reqQP.Set("directory", *options.Path)
	}
	reqQP.Set("recursive", strconv.FormatBool(recursive))
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxResults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	if options != nil && options.Upn != nil {
		reqQP.Set("upn", strconv.FormatBool(*options.Upn))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listPathsHandleResponse handles the ListPaths response.
func (client *FileSystemClient) ListPathsHandleResponse(resp *http.Response) (FileSystemClientListPathsResponse, error) {
	result := FileSystemClientListPathsResponse{}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientListPathsResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientListPathsResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-continuation"); val != "" {
		result.Continuation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PathList); err != nil {
		return FileSystemClientListPathsResponse{}, err
	}
	return result, nil
}

// SetProperties - Set properties for the FileSystem. This operation supports conditional HTTP requests. For more information,
// see Specifying Conditional Headers for Blob Service Operations
// [https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations].
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-02
//   - options - FileSystemClientSetPropertiesOptions contains the optional parameters for the FileSystemClient.SetProperties
//     method.
//   - ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the FileSystemClient.SetProperties
//     method.
func (client *FileSystemClient) SetProperties(ctx context.Context, options *FileSystemClientSetPropertiesOptions, modifiedAccessConditions *ModifiedAccessConditions) (FileSystemClientSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, options, modifiedAccessConditions)
	if err != nil {
		return FileSystemClientSetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileSystemClientSetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileSystemClientSetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *FileSystemClient) setPropertiesCreateRequest(ctx context.Context, options *FileSystemClientSetPropertiesOptions, modifiedAccessConditions *ModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPatch, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("resource", "filesystem")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.Properties != nil {
		req.Raw().Header["x-ms-properties"] = []string{*options.Properties}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123)}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *FileSystemClient) setPropertiesHandleResponse(resp *http.Response) (FileSystemClientSetPropertiesResponse, error) {
	result := FileSystemClientSetPropertiesResponse{}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientSetPropertiesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return FileSystemClientSetPropertiesResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return result, nil
}

