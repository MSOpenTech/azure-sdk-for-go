//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HierarchySettingsClient contains the methods for the HierarchySettings group.
// Don't use this type directly, use NewHierarchySettingsClient() instead.
type HierarchySettingsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewHierarchySettingsClient creates a new instance of HierarchySettingsClient with the specified values.
func NewHierarchySettingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *HierarchySettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &HierarchySettingsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates the hierarchy settings defined at the Management Group level.
// If the operation fails it returns the *ErrorResponse error type.
func (client *HierarchySettingsClient) CreateOrUpdate(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsCreateOrUpdateOptions) (HierarchySettingsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, groupID, createTenantSettingsRequest, options)
	if err != nil {
		return HierarchySettingsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HierarchySettingsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HierarchySettingsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HierarchySettingsClient) createOrUpdateCreateRequest(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, createTenantSettingsRequest)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HierarchySettingsClient) createOrUpdateHandleResponse(resp *http.Response) (HierarchySettingsCreateOrUpdateResponse, error) {
	result := HierarchySettingsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettings); err != nil {
		return HierarchySettingsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *HierarchySettingsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the hierarchy settings defined at the Management Group level.
// If the operation fails it returns the *ErrorResponse error type.
func (client *HierarchySettingsClient) Delete(ctx context.Context, groupID string, options *HierarchySettingsDeleteOptions) (HierarchySettingsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, groupID, options)
	if err != nil {
		return HierarchySettingsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HierarchySettingsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HierarchySettingsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return HierarchySettingsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HierarchySettingsClient) deleteCreateRequest(ctx context.Context, groupID string, options *HierarchySettingsDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *HierarchySettingsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management Group of the hierarchy.
// If the operation fails it returns the *ErrorResponse error type.
func (client *HierarchySettingsClient) Get(ctx context.Context, groupID string, options *HierarchySettingsGetOptions) (HierarchySettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groupID, options)
	if err != nil {
		return HierarchySettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HierarchySettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HierarchySettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HierarchySettingsClient) getCreateRequest(ctx context.Context, groupID string, options *HierarchySettingsGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
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

// getHandleResponse handles the Get response.
func (client *HierarchySettingsClient) getHandleResponse(resp *http.Response) (HierarchySettingsGetResponse, error) {
	result := HierarchySettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettings); err != nil {
		return HierarchySettingsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *HierarchySettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management Group of the hierarchy.
// If the operation fails it returns the *ErrorResponse error type.
func (client *HierarchySettingsClient) List(ctx context.Context, groupID string, options *HierarchySettingsListOptions) (HierarchySettingsListResponse, error) {
	req, err := client.listCreateRequest(ctx, groupID, options)
	if err != nil {
		return HierarchySettingsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HierarchySettingsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HierarchySettingsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *HierarchySettingsClient) listCreateRequest(ctx context.Context, groupID string, options *HierarchySettingsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
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
func (client *HierarchySettingsClient) listHandleResponse(resp *http.Response) (HierarchySettingsListResponse, error) {
	result := HierarchySettingsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettingsList); err != nil {
		return HierarchySettingsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *HierarchySettingsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates the hierarchy settings defined at the Management Group level.
// If the operation fails it returns the *ErrorResponse error type.
func (client *HierarchySettingsClient) Update(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsUpdateOptions) (HierarchySettingsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, groupID, createTenantSettingsRequest, options)
	if err != nil {
		return HierarchySettingsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HierarchySettingsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HierarchySettingsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *HierarchySettingsClient) updateCreateRequest(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, createTenantSettingsRequest)
}

// updateHandleResponse handles the Update response.
func (client *HierarchySettingsClient) updateHandleResponse(resp *http.Response) (HierarchySettingsUpdateResponse, error) {
	result := HierarchySettingsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettings); err != nil {
		return HierarchySettingsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *HierarchySettingsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
