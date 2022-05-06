//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AttachedDataNetworksClient contains the methods for the AttachedDataNetworks group.
// Don't use this type directly, use NewAttachedDataNetworksClient() instead.
type AttachedDataNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAttachedDataNetworksClient creates a new instance of AttachedDataNetworksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAttachedDataNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AttachedDataNetworksClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AttachedDataNetworksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// packetCoreDataPlaneName - The name of the packet core data plane.
// attachedDataNetworkName - The name of the attached data network.
// parameters - Parameters supplied to the create or update attached data network operation.
// options - AttachedDataNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the AttachedDataNetworksClient.BeginCreateOrUpdate
// method.
func (client *AttachedDataNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters AttachedDataNetwork, options *AttachedDataNetworksClientBeginCreateOrUpdateOptions) (*armruntime.Poller[AttachedDataNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[AttachedDataNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[AttachedDataNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AttachedDataNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters AttachedDataNetwork, options *AttachedDataNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AttachedDataNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters AttachedDataNetwork, options *AttachedDataNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// packetCoreDataPlaneName - The name of the packet core data plane.
// attachedDataNetworkName - The name of the attached data network.
// options - AttachedDataNetworksClientBeginDeleteOptions contains the optional parameters for the AttachedDataNetworksClient.BeginDelete
// method.
func (client *AttachedDataNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientBeginDeleteOptions) (*armruntime.Poller[AttachedDataNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[AttachedDataNetworksClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[AttachedDataNetworksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AttachedDataNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AttachedDataNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// packetCoreDataPlaneName - The name of the packet core data plane.
// attachedDataNetworkName - The name of the attached data network.
// options - AttachedDataNetworksClientGetOptions contains the optional parameters for the AttachedDataNetworksClient.Get
// method.
func (client *AttachedDataNetworksClient) Get(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientGetOptions) (AttachedDataNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, options)
	if err != nil {
		return AttachedDataNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttachedDataNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttachedDataNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AttachedDataNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AttachedDataNetworksClient) getHandleResponse(resp *http.Response) (AttachedDataNetworksClientGetResponse, error) {
	result := AttachedDataNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDataNetwork); err != nil {
		return AttachedDataNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPacketCoreDataPlanePager - Gets all the data networks associated with a packet core data plane.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// packetCoreDataPlaneName - The name of the packet core data plane.
// options - AttachedDataNetworksClientListByPacketCoreDataPlaneOptions contains the optional parameters for the AttachedDataNetworksClient.ListByPacketCoreDataPlane
// method.
func (client *AttachedDataNetworksClient) NewListByPacketCoreDataPlanePager(resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, options *AttachedDataNetworksClientListByPacketCoreDataPlaneOptions) *runtime.Pager[AttachedDataNetworksClientListByPacketCoreDataPlaneResponse] {
	return runtime.NewPager(runtime.PageProcessor[AttachedDataNetworksClientListByPacketCoreDataPlaneResponse]{
		More: func(page AttachedDataNetworksClientListByPacketCoreDataPlaneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttachedDataNetworksClientListByPacketCoreDataPlaneResponse) (AttachedDataNetworksClientListByPacketCoreDataPlaneResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPacketCoreDataPlaneCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPacketCoreDataPlaneHandleResponse(resp)
		},
	})
}

// listByPacketCoreDataPlaneCreateRequest creates the ListByPacketCoreDataPlane request.
func (client *AttachedDataNetworksClient) listByPacketCoreDataPlaneCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, options *AttachedDataNetworksClientListByPacketCoreDataPlaneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPacketCoreDataPlaneHandleResponse handles the ListByPacketCoreDataPlane response.
func (client *AttachedDataNetworksClient) listByPacketCoreDataPlaneHandleResponse(resp *http.Response) (AttachedDataNetworksClientListByPacketCoreDataPlaneResponse, error) {
	result := AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDataNetworkListResult); err != nil {
		return AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an attached data network update tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// packetCoreControlPlaneName - The name of the packet core control plane.
// packetCoreDataPlaneName - The name of the packet core data plane.
// attachedDataNetworkName - The name of the attached data network.
// parameters - Parameters supplied to update attached data network tags.
// options - AttachedDataNetworksClientUpdateTagsOptions contains the optional parameters for the AttachedDataNetworksClient.UpdateTags
// method.
func (client *AttachedDataNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters TagsObject, options *AttachedDataNetworksClientUpdateTagsOptions) (AttachedDataNetworksClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, parameters, options)
	if err != nil {
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AttachedDataNetworksClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *AttachedDataNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters TagsObject, options *AttachedDataNetworksClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *AttachedDataNetworksClient) updateTagsHandleResponse(resp *http.Response) (AttachedDataNetworksClientUpdateTagsResponse, error) {
	result := AttachedDataNetworksClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDataNetwork); err != nil {
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	return result, nil
}
