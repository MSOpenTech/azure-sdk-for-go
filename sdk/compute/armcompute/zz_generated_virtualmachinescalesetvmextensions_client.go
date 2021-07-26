// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineScaleSetVMExtensionsClient contains the methods for the VirtualMachineScaleSetVMExtensions group.
// Don't use this type directly, use NewVirtualMachineScaleSetVMExtensionsClient() instead.
type VirtualMachineScaleSetVMExtensionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetVMExtensionsClient creates a new instance of VirtualMachineScaleSetVMExtensionsClient with the specified values.
func NewVirtualMachineScaleSetVMExtensionsClient(con *armcore.Connection, subscriptionID string) *VirtualMachineScaleSetVMExtensionsClient {
	return &VirtualMachineScaleSetVMExtensionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - The operation to create or update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualMachineScaleSetVMExtensionsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineScaleSetVMExtensionsCreateOrUpdatePoller.ResumeToken().
func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineScaleSetVMExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - The operation to delete the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (VirtualMachineScaleSetVMExtensionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsDeletePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualMachineScaleSetVMExtensionsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsDeletePollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new VirtualMachineScaleSetVMExtensionsDeletePoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineScaleSetVMExtensionsDeletePoller.ResumeToken().
func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeDelete(ctx context.Context, token string) (VirtualMachineScaleSetVMExtensionsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsDeletePollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsDeletePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - The operation to delete the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineScaleSetVMExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - The operation to get the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsGetOptions) (VirtualMachineScaleSetVMExtensionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineScaleSetVMExtensionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetVMExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMExtensionsClient) getHandleResponse(resp *azcore.Response) (VirtualMachineScaleSetVMExtensionsGetResponse, error) {
	result := VirtualMachineScaleSetVMExtensionsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetVMExtension); err != nil {
		return VirtualMachineScaleSetVMExtensionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - The operation to get all extensions of an instance in Virtual Machine Scaleset.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) List(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMExtensionsListOptions) (VirtualMachineScaleSetVMExtensionsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineScaleSetVMExtensionsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetVMExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMExtensionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMExtensionsClient) listHandleResponse(resp *azcore.Response) (VirtualMachineScaleSetVMExtensionsListResponse, error) {
	result := VirtualMachineScaleSetVMExtensionsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetVMExtensionsListResult); err != nil {
		return VirtualMachineScaleSetVMExtensionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - The operation to update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (VirtualMachineScaleSetVMExtensionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionsUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VirtualMachineScaleSetVMExtensionsClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsUpdatePollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionsUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionsUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new VirtualMachineScaleSetVMExtensionsUpdatePoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineScaleSetVMExtensionsUpdatePoller.ResumeToken().
func (client *VirtualMachineScaleSetVMExtensionsClient) ResumeUpdate(ctx context.Context, token string) (VirtualMachineScaleSetVMExtensionsUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsUpdatePollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionsUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionsUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - The operation to update the VMSS VM extension.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMExtensionsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// updateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetVMExtensionsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
