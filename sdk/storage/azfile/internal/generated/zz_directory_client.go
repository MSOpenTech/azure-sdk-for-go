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

// DirectoryClient contains the methods for the Directory group.
// Don't use this type directly, use NewDirectoryClient() instead.
type DirectoryClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// NewDirectoryClient creates a new instance of DirectoryClient with the specified values.
// endpoint - The URL of the service account, share, directory or file that is the target of the desired operation.
// pl - the pipeline used for sending requests and handling responses.
func NewDirectoryClient(endpoint string, pl runtime.Pipeline) *DirectoryClient {
	client := &DirectoryClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Create - Creates a new directory under the specified share or parent directory.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// fileAttributes - If specified, the provided file attributes shall be set. Default value: ‘Archive’ for file and ‘Directory’
// for directory. ‘None’ can also be specified as default.
// fileCreationTime - Creation time for the file/directory. Default value: Now.
// fileLastWriteTime - Last write time for the file/directory. Default value: Now.
// options - DirectoryClientCreateOptions contains the optional parameters for the DirectoryClient.Create method.
func (client *DirectoryClient) Create(ctx context.Context, fileAttributes string, fileCreationTime string, fileLastWriteTime string, options *DirectoryClientCreateOptions) (DirectoryClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, fileAttributes, fileCreationTime, fileLastWriteTime, options)
	if err != nil {
		return DirectoryClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return DirectoryClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DirectoryClient) createCreateRequest(ctx context.Context, fileAttributes string, fileCreationTime string, fileLastWriteTime string, options *DirectoryClientCreateOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "directory")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.Metadata != nil {
		for k, v := range options.Metadata {
			req.Raw().Header["x-ms-meta-"+k] = []string{v}
		}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.FilePermission != nil {
		req.Raw().Header["x-ms-file-permission"] = []string{*options.FilePermission}
	}
	if options != nil && options.FilePermissionKey != nil {
		req.Raw().Header["x-ms-file-permission-key"] = []string{*options.FilePermissionKey}
	}
	req.Raw().Header["x-ms-file-attributes"] = []string{fileAttributes}
	req.Raw().Header["x-ms-file-creation-time"] = []string{fileCreationTime}
	req.Raw().Header["x-ms-file-last-write-time"] = []string{fileLastWriteTime}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DirectoryClient) createHandleResponse(resp *http.Response) (DirectoryClientCreateResponse, error) {
	result := DirectoryClientCreateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DirectoryClientCreateResponse{}, err
		}
		result.LastModified = &lastModified
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
			return DirectoryClientCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return DirectoryClientCreateResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-file-permission-key"); val != "" {
		result.FilePermissionKey = &val
	}
	if val := resp.Header.Get("x-ms-file-attributes"); val != "" {
		result.FileAttributes = &val
	}
	if val := resp.Header.Get("x-ms-file-creation-time"); val != "" {
		result.FileCreationTime = &val
	}
	if val := resp.Header.Get("x-ms-file-last-write-time"); val != "" {
		result.FileLastWriteTime = &val
	}
	if val := resp.Header.Get("x-ms-file-change-time"); val != "" {
		result.FileChangeTime = &val
	}
	if val := resp.Header.Get("x-ms-file-id"); val != "" {
		result.FileID = &val
	}
	if val := resp.Header.Get("x-ms-file-parent-id"); val != "" {
		result.FileParentID = &val
	}
	return result, nil
}

// Delete - Removes the specified empty directory. Note that the directory must be empty before it can be deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// options - DirectoryClientDeleteOptions contains the optional parameters for the DirectoryClient.Delete method.
func (client *DirectoryClient) Delete(ctx context.Context, options *DirectoryClientDeleteOptions) (DirectoryClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return DirectoryClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return DirectoryClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *DirectoryClient) deleteCreateRequest(ctx context.Context, options *DirectoryClientDeleteOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodDelete, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "directory")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *DirectoryClient) deleteHandleResponse(resp *http.Response) (DirectoryClientDeleteResponse, error) {
	result := DirectoryClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DirectoryClientDeleteResponse{}, err
		}
		result.Date = &date
	}
	return result, nil
}

// ForceCloseHandles - Closes all handles open for given directory.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// handleID - Specifies handle ID opened on the file or directory to be closed. Asterisk (‘*’) is a wildcard that specifies
// all handles.
// options - DirectoryClientForceCloseHandlesOptions contains the optional parameters for the DirectoryClient.ForceCloseHandles
// method.
func (client *DirectoryClient) ForceCloseHandles(ctx context.Context, handleID string, options *DirectoryClientForceCloseHandlesOptions) (DirectoryClientForceCloseHandlesResponse, error) {
	req, err := client.forceCloseHandlesCreateRequest(ctx, handleID, options)
	if err != nil {
		return DirectoryClientForceCloseHandlesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientForceCloseHandlesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DirectoryClientForceCloseHandlesResponse{}, runtime.NewResponseError(resp)
	}
	return client.forceCloseHandlesHandleResponse(resp)
}

// forceCloseHandlesCreateRequest creates the ForceCloseHandles request.
func (client *DirectoryClient) forceCloseHandlesCreateRequest(ctx context.Context, handleID string, options *DirectoryClientForceCloseHandlesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "forceclosehandles")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Sharesnapshot != nil {
		reqQP.Set("sharesnapshot", *options.Sharesnapshot)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-handle-id"] = []string{handleID}
	if options != nil && options.Recursive != nil {
		req.Raw().Header["x-ms-recursive"] = []string{strconv.FormatBool(*options.Recursive)}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// forceCloseHandlesHandleResponse handles the ForceCloseHandles response.
func (client *DirectoryClient) forceCloseHandlesHandleResponse(resp *http.Response) (DirectoryClientForceCloseHandlesResponse, error) {
	result := DirectoryClientForceCloseHandlesResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DirectoryClientForceCloseHandlesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-marker"); val != "" {
		result.Marker = &val
	}
	if val := resp.Header.Get("x-ms-number-of-handles-closed"); val != "" {
		numberOfHandlesClosed32, err := strconv.ParseInt(val, 10, 32)
		numberOfHandlesClosed := int32(numberOfHandlesClosed32)
		if err != nil {
			return DirectoryClientForceCloseHandlesResponse{}, err
		}
		result.NumberOfHandlesClosed = &numberOfHandlesClosed
	}
	if val := resp.Header.Get("x-ms-number-of-handles-failed"); val != "" {
		numberOfHandlesFailedToClose32, err := strconv.ParseInt(val, 10, 32)
		numberOfHandlesFailedToClose := int32(numberOfHandlesFailedToClose32)
		if err != nil {
			return DirectoryClientForceCloseHandlesResponse{}, err
		}
		result.NumberOfHandlesFailedToClose = &numberOfHandlesFailedToClose
	}
	return result, nil
}

// GetProperties - Returns all system properties for the specified directory, and can also be used to check the existence
// of a directory. The data returned does not include the files in the directory or any
// subdirectories.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// options - DirectoryClientGetPropertiesOptions contains the optional parameters for the DirectoryClient.GetProperties method.
func (client *DirectoryClient) GetProperties(ctx context.Context, options *DirectoryClientGetPropertiesOptions) (DirectoryClientGetPropertiesResponse, error) {
	req, err := client.getPropertiesCreateRequest(ctx, options)
	if err != nil {
		return DirectoryClientGetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientGetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DirectoryClientGetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPropertiesHandleResponse(resp)
}

// getPropertiesCreateRequest creates the GetProperties request.
func (client *DirectoryClient) getPropertiesCreateRequest(ctx context.Context, options *DirectoryClientGetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "directory")
	if options != nil && options.Sharesnapshot != nil {
		reqQP.Set("sharesnapshot", *options.Sharesnapshot)
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// getPropertiesHandleResponse handles the GetProperties response.
func (client *DirectoryClient) getPropertiesHandleResponse(resp *http.Response) (DirectoryClientGetPropertiesResponse, error) {
	result := DirectoryClientGetPropertiesResponse{}
	for hh := range resp.Header {
		if len(hh) > len("x-ms-meta-") && strings.EqualFold(hh[:len("x-ms-meta-")], "x-ms-meta-") {
			if result.Metadata == nil {
				result.Metadata = map[string]string{}
			}
			result.Metadata[hh[len("x-ms-meta-"):]] = resp.Header.Get(hh)
		}
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DirectoryClientGetPropertiesResponse{}, err
		}
		result.LastModified = &lastModified
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
			return DirectoryClientGetPropertiesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return DirectoryClientGetPropertiesResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-file-attributes"); val != "" {
		result.FileAttributes = &val
	}
	if val := resp.Header.Get("x-ms-file-creation-time"); val != "" {
		result.FileCreationTime = &val
	}
	if val := resp.Header.Get("x-ms-file-last-write-time"); val != "" {
		result.FileLastWriteTime = &val
	}
	if val := resp.Header.Get("x-ms-file-change-time"); val != "" {
		result.FileChangeTime = &val
	}
	if val := resp.Header.Get("x-ms-file-permission-key"); val != "" {
		result.FilePermissionKey = &val
	}
	if val := resp.Header.Get("x-ms-file-id"); val != "" {
		result.FileID = &val
	}
	if val := resp.Header.Get("x-ms-file-parent-id"); val != "" {
		result.FileParentID = &val
	}
	return result, nil
}

// NewListFilesAndDirectoriesSegmentPager - Returns a list of files or directories under the specified share or directory.
// It lists the contents only for a single level of the directory hierarchy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// options - DirectoryClientListFilesAndDirectoriesSegmentOptions contains the optional parameters for the DirectoryClient.ListFilesAndDirectoriesSegment
// method.
func (client *DirectoryClient) NewListFilesAndDirectoriesSegmentPager(options *DirectoryClientListFilesAndDirectoriesSegmentOptions) *runtime.Pager[DirectoryClientListFilesAndDirectoriesSegmentResponse] {
	return runtime.NewPager(runtime.PagingHandler[DirectoryClientListFilesAndDirectoriesSegmentResponse]{
		More: func(page DirectoryClientListFilesAndDirectoriesSegmentResponse) bool {
			return page.NextMarker != nil && len(*page.NextMarker) > 0
		},
		Fetcher: func(ctx context.Context, page *DirectoryClientListFilesAndDirectoriesSegmentResponse) (DirectoryClientListFilesAndDirectoriesSegmentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listFilesAndDirectoriesSegmentCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextMarker)
			}
			if err != nil {
				return DirectoryClientListFilesAndDirectoriesSegmentResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DirectoryClientListFilesAndDirectoriesSegmentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DirectoryClientListFilesAndDirectoriesSegmentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listFilesAndDirectoriesSegmentHandleResponse(resp)
		},
	})
}

// listFilesAndDirectoriesSegmentCreateRequest creates the ListFilesAndDirectoriesSegment request.
func (client *DirectoryClient) listFilesAndDirectoriesSegmentCreateRequest(ctx context.Context, options *DirectoryClientListFilesAndDirectoriesSegmentOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "directory")
	reqQP.Set("comp", "list")
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.Sharesnapshot != nil {
		reqQP.Set("sharesnapshot", *options.Sharesnapshot)
	}
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Include != nil {
		reqQP.Set("include", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.Include), "[]")), ","))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.IncludeExtendedInfo != nil {
		req.Raw().Header["x-ms-file-extended-info"] = []string{strconv.FormatBool(*options.IncludeExtendedInfo)}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// listFilesAndDirectoriesSegmentHandleResponse handles the ListFilesAndDirectoriesSegment response.
func (client *DirectoryClient) listFilesAndDirectoriesSegmentHandleResponse(resp *http.Response) (DirectoryClientListFilesAndDirectoriesSegmentResponse, error) {
	result := DirectoryClientListFilesAndDirectoriesSegmentResponse{}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
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
			return DirectoryClientListFilesAndDirectoriesSegmentResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ListFilesAndDirectoriesSegmentResponse); err != nil {
		return DirectoryClientListFilesAndDirectoriesSegmentResponse{}, err
	}
	return result, nil
}

// ListHandles - Lists handles for directory.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// options - DirectoryClientListHandlesOptions contains the optional parameters for the DirectoryClient.ListHandles method.
func (client *DirectoryClient) ListHandles(ctx context.Context, options *DirectoryClientListHandlesOptions) (DirectoryClientListHandlesResponse, error) {
	req, err := client.listHandlesCreateRequest(ctx, options)
	if err != nil {
		return DirectoryClientListHandlesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientListHandlesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DirectoryClientListHandlesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandlesHandleResponse(resp)
}

// listHandlesCreateRequest creates the ListHandles request.
func (client *DirectoryClient) listHandlesCreateRequest(ctx context.Context, options *DirectoryClientListHandlesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodGet, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "listhandles")
	if options != nil && options.Marker != nil {
		reqQP.Set("marker", *options.Marker)
	}
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	if options != nil && options.Sharesnapshot != nil {
		reqQP.Set("sharesnapshot", *options.Sharesnapshot)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.Recursive != nil {
		req.Raw().Header["x-ms-recursive"] = []string{strconv.FormatBool(*options.Recursive)}
	}
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// listHandlesHandleResponse handles the ListHandles response.
func (client *DirectoryClient) listHandlesHandleResponse(resp *http.Response) (DirectoryClientListHandlesResponse, error) {
	result := DirectoryClientListHandlesResponse{}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
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
			return DirectoryClientListHandlesResponse{}, err
		}
		result.Date = &date
	}
	if err := runtime.UnmarshalAsXML(resp, &result.ListHandlesResponse); err != nil {
		return DirectoryClientListHandlesResponse{}, err
	}
	return result, nil
}

// SetMetadata - Updates user defined metadata for the specified directory.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// options - DirectoryClientSetMetadataOptions contains the optional parameters for the DirectoryClient.SetMetadata method.
func (client *DirectoryClient) SetMetadata(ctx context.Context, options *DirectoryClientSetMetadataOptions) (DirectoryClientSetMetadataResponse, error) {
	req, err := client.setMetadataCreateRequest(ctx, options)
	if err != nil {
		return DirectoryClientSetMetadataResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientSetMetadataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DirectoryClientSetMetadataResponse{}, runtime.NewResponseError(resp)
	}
	return client.setMetadataHandleResponse(resp)
}

// setMetadataCreateRequest creates the SetMetadata request.
func (client *DirectoryClient) setMetadataCreateRequest(ctx context.Context, options *DirectoryClientSetMetadataOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "directory")
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
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// setMetadataHandleResponse handles the SetMetadata response.
func (client *DirectoryClient) setMetadataHandleResponse(resp *http.Response) (DirectoryClientSetMetadataResponse, error) {
	result := DirectoryClientSetMetadataResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
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
			return DirectoryClientSetMetadataResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return DirectoryClientSetMetadataResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	return result, nil
}

// SetProperties - Sets properties on the directory.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-02
// fileAttributes - If specified, the provided file attributes shall be set. Default value: ‘Archive’ for file and ‘Directory’
// for directory. ‘None’ can also be specified as default.
// fileCreationTime - Creation time for the file/directory. Default value: Now.
// fileLastWriteTime - Last write time for the file/directory. Default value: Now.
// options - DirectoryClientSetPropertiesOptions contains the optional parameters for the DirectoryClient.SetProperties method.
func (client *DirectoryClient) SetProperties(ctx context.Context, fileAttributes string, fileCreationTime string, fileLastWriteTime string, options *DirectoryClientSetPropertiesOptions) (DirectoryClientSetPropertiesResponse, error) {
	req, err := client.setPropertiesCreateRequest(ctx, fileAttributes, fileCreationTime, fileLastWriteTime, options)
	if err != nil {
		return DirectoryClientSetPropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DirectoryClientSetPropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DirectoryClientSetPropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.setPropertiesHandleResponse(resp)
}

// setPropertiesCreateRequest creates the SetProperties request.
func (client *DirectoryClient) setPropertiesCreateRequest(ctx context.Context, fileAttributes string, fileCreationTime string, fileLastWriteTime string, options *DirectoryClientSetPropertiesOptions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("restype", "directory")
	reqQP.Set("comp", "properties")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2020-10-02"}
	if options != nil && options.FilePermission != nil {
		req.Raw().Header["x-ms-file-permission"] = []string{*options.FilePermission}
	}
	if options != nil && options.FilePermissionKey != nil {
		req.Raw().Header["x-ms-file-permission-key"] = []string{*options.FilePermissionKey}
	}
	req.Raw().Header["x-ms-file-attributes"] = []string{fileAttributes}
	req.Raw().Header["x-ms-file-creation-time"] = []string{fileCreationTime}
	req.Raw().Header["x-ms-file-last-write-time"] = []string{fileLastWriteTime}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// setPropertiesHandleResponse handles the SetProperties response.
func (client *DirectoryClient) setPropertiesHandleResponse(resp *http.Response) (DirectoryClientSetPropertiesResponse, error) {
	result := DirectoryClientSetPropertiesResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DirectoryClientSetPropertiesResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return DirectoryClientSetPropertiesResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return DirectoryClientSetPropertiesResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-file-permission-key"); val != "" {
		result.FilePermissionKey = &val
	}
	if val := resp.Header.Get("x-ms-file-attributes"); val != "" {
		result.FileAttributes = &val
	}
	if val := resp.Header.Get("x-ms-file-creation-time"); val != "" {
		result.FileCreationTime = &val
	}
	if val := resp.Header.Get("x-ms-file-last-write-time"); val != "" {
		result.FileLastWriteTime = &val
	}
	if val := resp.Header.Get("x-ms-file-change-time"); val != "" {
		result.FileChangeTime = &val
	}
	if val := resp.Header.Get("x-ms-file-id"); val != "" {
		result.FileID = &val
	}
	if val := resp.Header.Get("x-ms-file-parent-id"); val != "" {
		result.FileParentID = &val
	}
	return result, nil
}
