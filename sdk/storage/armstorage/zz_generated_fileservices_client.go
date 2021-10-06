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
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// FileServicesClient contains the methods for the FileServices group.
// Don't use this type directly, use NewFileServicesClient() instead.
type FileServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewFileServicesClient creates a new instance of FileServicesClient with the specified values.
func NewFileServicesClient(con *arm.Connection, subscriptionID string) *FileServicesClient {
	return &FileServicesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetServiceProperties - Gets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *FileServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesGetServicePropertiesOptions) (FileServicesGetServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return FileServicesGetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileServicesGetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileServicesGetServicePropertiesResponse{}, client.getServicePropertiesHandleError(resp)
	}
	return client.getServicePropertiesHandleResponse(resp)
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *FileServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesGetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{FileServicesName}", url.PathEscape("default"))
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

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *FileServicesClient) getServicePropertiesHandleResponse(resp *http.Response) (FileServicesGetServicePropertiesResponse, error) {
	result := FileServicesGetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileServiceProperties); err != nil {
		return FileServicesGetServicePropertiesResponse{}, err
	}
	return result, nil
}

// getServicePropertiesHandleError handles the GetServiceProperties error response.
func (client *FileServicesClient) getServicePropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List all file services in storage accounts
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *FileServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesListOptions) (FileServicesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return FileServicesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileServicesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileServicesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *FileServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *FileServicesClient) listHandleResponse(resp *http.Response) (FileServicesListResponse, error) {
	result := FileServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileServiceItems); err != nil {
		return FileServicesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *FileServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// SetServiceProperties - Sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns the *CloudErrorAutoGenerated error type.
func (client *FileServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters FileServiceProperties, options *FileServicesSetServicePropertiesOptions) (FileServicesSetServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return FileServicesSetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileServicesSetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileServicesSetServicePropertiesResponse{}, client.setServicePropertiesHandleError(resp)
	}
	return client.setServicePropertiesHandleResponse(resp)
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *FileServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters FileServiceProperties, options *FileServicesSetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{FileServicesName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *FileServicesClient) setServicePropertiesHandleResponse(resp *http.Response) (FileServicesSetServicePropertiesResponse, error) {
	result := FileServicesSetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileServiceProperties); err != nil {
		return FileServicesSetServicePropertiesResponse{}, err
	}
	return result, nil
}

// setServicePropertiesHandleError handles the SetServiceProperties error response.
func (client *FileServicesClient) setServicePropertiesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudErrorAutoGenerated{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
