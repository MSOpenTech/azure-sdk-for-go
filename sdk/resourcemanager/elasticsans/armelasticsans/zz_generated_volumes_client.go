//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsans

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

// VolumesClient contains the methods for the Volumes group.
// Don't use this type directly, use NewVolumesClient() instead.
type VolumesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVolumesClient creates a new instance of VolumesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVolumesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VolumesClient, error) {
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
	client := &VolumesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create a Volume.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// volumeGroupName - The name of the VolumeGroup.
// volumeName - The name of the Volume.
// parameters - Volume object.
// options - VolumesClientBeginCreateOptions contains the optional parameters for the VolumesClient.BeginCreate method.
func (client *VolumesClient) BeginCreate(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, parameters Volume, options *VolumesClientBeginCreateOptions) (*armruntime.Poller[VolumesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VolumesClientCreateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VolumesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a Volume.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VolumesClient) create(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, parameters Volume, options *VolumesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *VolumesClient) createCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, parameters Volume, options *VolumesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete an Volume.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// volumeGroupName - The name of the VolumeGroup.
// volumeName - The name of the Volume.
// options - VolumesClientBeginDeleteOptions contains the optional parameters for the VolumesClient.BeginDelete method.
func (client *VolumesClient) BeginDelete(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*armruntime.Poller[VolumesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VolumesClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VolumesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an Volume.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VolumesClient) deleteOperation(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, options)
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
func (client *VolumesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get an Volume.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// volumeGroupName - The name of the VolumeGroup.
// volumeName - The name of the Volume.
// options - VolumesClientGetOptions contains the optional parameters for the VolumesClient.Get method.
func (client *VolumesClient) Get(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, options *VolumesClientGetOptions) (VolumesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, options)
	if err != nil {
		return VolumesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VolumesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VolumesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VolumesClient) getCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, options *VolumesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumesClient) getHandleResponse(resp *http.Response) (VolumesClientGetResponse, error) {
	result := VolumesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Volume); err != nil {
		return VolumesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVolumeGroupPager - List Volumes in a VolumeGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// volumeGroupName - The name of the VolumeGroup.
// options - VolumesClientListByVolumeGroupOptions contains the optional parameters for the VolumesClient.ListByVolumeGroup
// method.
func (client *VolumesClient) NewListByVolumeGroupPager(resourceGroupName string, elasticSanName string, volumeGroupName string, options *VolumesClientListByVolumeGroupOptions) *runtime.Pager[VolumesClientListByVolumeGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[VolumesClientListByVolumeGroupResponse]{
		More: func(page VolumesClientListByVolumeGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VolumesClientListByVolumeGroupResponse) (VolumesClientListByVolumeGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVolumeGroupCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VolumesClientListByVolumeGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VolumesClientListByVolumeGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumesClientListByVolumeGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVolumeGroupHandleResponse(resp)
		},
	})
}

// listByVolumeGroupCreateRequest creates the ListByVolumeGroup request.
func (client *VolumesClient) listByVolumeGroupCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, options *VolumesClientListByVolumeGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVolumeGroupHandleResponse handles the ListByVolumeGroup response.
func (client *VolumesClient) listByVolumeGroupHandleResponse(resp *http.Response) (VolumesClientListByVolumeGroupResponse, error) {
	result := VolumesClientListByVolumeGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeList); err != nil {
		return VolumesClientListByVolumeGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an Volume.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// volumeGroupName - The name of the VolumeGroup.
// volumeName - The name of the Volume.
// parameters - Volume object.
// options - VolumesClientBeginUpdateOptions contains the optional parameters for the VolumesClient.BeginUpdate method.
func (client *VolumesClient) BeginUpdate(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, parameters VolumeUpdate, options *VolumesClientBeginUpdateOptions) (*armruntime.Poller[VolumesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VolumesClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VolumesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update an Volume.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VolumesClient) update(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, parameters VolumeUpdate, options *VolumesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, elasticSanName, volumeGroupName, volumeName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VolumesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, volumeName string, parameters VolumeUpdate, options *VolumesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
