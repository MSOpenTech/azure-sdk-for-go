//go:build go1.16
// +build go1.16

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DataNetworksClient contains the methods for the DataNetworks group.
// Don't use this type directly, use NewDataNetworksClient() instead.
type DataNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataNetworksClient creates a new instance of DataNetworksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataNetworksClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DataNetworksClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a mobile network dataNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// dataNetworkName - The name of the mobile network dataNetwork.
// parameters - Parameters supplied to the create or update mobile network dataNetwork operation.
// options - DataNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the DataNetworksClient.BeginCreateOrUpdate
// method.
func (client *DataNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters DataNetwork, options *DataNetworksClientBeginCreateOrUpdateOptions) (DataNetworksClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, parameters, options)
	if err != nil {
		return DataNetworksClientCreateOrUpdatePollerResponse{}, err
	}
	result := DataNetworksClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataNetworksClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return DataNetworksClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DataNetworksClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a mobile network dataNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters DataNetwork, options *DataNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, parameters, options)
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
func (client *DataNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters DataNetwork, options *DataNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified mobile network dataNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// dataNetworkName - The name of the mobile network dataNetwork.
// options - DataNetworksClientBeginDeleteOptions contains the optional parameters for the DataNetworksClient.BeginDelete
// method.
func (client *DataNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientBeginDeleteOptions) (DataNetworksClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, options)
	if err != nil {
		return DataNetworksClientDeletePollerResponse{}, err
	}
	result := DataNetworksClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataNetworksClient.Delete", "location", resp, client.pl)
	if err != nil {
		return DataNetworksClientDeletePollerResponse{}, err
	}
	result.Poller = &DataNetworksClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified mobile network dataNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DataNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, options)
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
func (client *DataNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified mobile network dataNetwork.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// dataNetworkName - The name of the mobile network dataNetwork.
// options - DataNetworksClientGetOptions contains the optional parameters for the DataNetworksClient.Get method.
func (client *DataNetworksClient) Get(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientGetOptions) (DataNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, options)
	if err != nil {
		return DataNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataNetworksClient) getHandleResponse(resp *http.Response) (DataNetworksClientGetResponse, error) {
	result := DataNetworksClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataNetwork); err != nil {
		return DataNetworksClientGetResponse{}, err
	}
	return result, nil
}

// ListByMobileNetwork - Lists all dataNetworks in the mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// options - DataNetworksClientListByMobileNetworkOptions contains the optional parameters for the DataNetworksClient.ListByMobileNetwork
// method.
func (client *DataNetworksClient) ListByMobileNetwork(resourceGroupName string, mobileNetworkName string, options *DataNetworksClientListByMobileNetworkOptions) *DataNetworksClientListByMobileNetworkPager {
	return &DataNetworksClientListByMobileNetworkPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByMobileNetworkCreateRequest(ctx, resourceGroupName, mobileNetworkName, options)
		},
		advancer: func(ctx context.Context, resp DataNetworksClientListByMobileNetworkResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DataNetworkListResult.NextLink)
		},
	}
}

// listByMobileNetworkCreateRequest creates the ListByMobileNetwork request.
func (client *DataNetworksClient) listByMobileNetworkCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, options *DataNetworksClientListByMobileNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByMobileNetworkHandleResponse handles the ListByMobileNetwork response.
func (client *DataNetworksClient) listByMobileNetworkHandleResponse(resp *http.Response) (DataNetworksClientListByMobileNetworkResponse, error) {
	result := DataNetworksClientListByMobileNetworkResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataNetworkListResult); err != nil {
		return DataNetworksClientListByMobileNetworkResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update data network tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// dataNetworkName - The name of the mobile network dataNetwork.
// parameters - Parameters supplied to update data network tags.
// options - DataNetworksClientUpdateTagsOptions contains the optional parameters for the DataNetworksClient.UpdateTags method.
func (client *DataNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters TagsObject, options *DataNetworksClientUpdateTagsOptions) (DataNetworksClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, parameters, options)
	if err != nil {
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataNetworksClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DataNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters TagsObject, options *DataNetworksClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DataNetworksClient) updateTagsHandleResponse(resp *http.Response) (DataNetworksClientUpdateTagsResponse, error) {
	result := DataNetworksClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataNetwork); err != nil {
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	return result, nil
}
