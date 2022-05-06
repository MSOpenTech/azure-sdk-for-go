//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

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

// TracksClient contains the methods for the Tracks group.
// Don't use this type directly, use NewTracksClient() instead.
type TracksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTracksClient creates a new instance of TracksClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTracksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TracksClient, error) {
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
	client := &TracksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Track in the asset
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// trackName - The Asset Track name.
// parameters - The request parameters
// options - TracksClientBeginCreateOrUpdateOptions contains the optional parameters for the TracksClient.BeginCreateOrUpdate
// method.
func (client *TracksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, parameters AssetTrack, options *TracksClientBeginCreateOrUpdateOptions) (*armruntime.Poller[TracksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, assetName, trackName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[TracksClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[TracksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a Track in the asset
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TracksClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, parameters AssetTrack, options *TracksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, assetName, trackName, parameters, options)
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
func (client *TracksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, parameters AssetTrack, options *TracksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if trackName == "" {
		return nil, errors.New("parameter trackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trackName}", url.PathEscape(trackName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a Track in the asset
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// trackName - The Asset Track name.
// options - TracksClientBeginDeleteOptions contains the optional parameters for the TracksClient.BeginDelete method.
func (client *TracksClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientBeginDeleteOptions) (*armruntime.Poller[TracksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, assetName, trackName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[TracksClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[TracksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Track in the asset
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TracksClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, assetName, trackName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TracksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if trackName == "" {
		return nil, errors.New("parameter trackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trackName}", url.PathEscape(trackName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the details of a Track in the Asset
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// trackName - The Asset Track name.
// options - TracksClientGetOptions contains the optional parameters for the TracksClient.Get method.
func (client *TracksClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientGetOptions) (TracksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, assetName, trackName, options)
	if err != nil {
		return TracksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TracksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TracksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TracksClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if trackName == "" {
		return nil, errors.New("parameter trackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trackName}", url.PathEscape(trackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TracksClient) getHandleResponse(resp *http.Response) (TracksClientGetResponse, error) {
	result := TracksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetTrack); err != nil {
		return TracksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the Tracks in the asset
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// options - TracksClientListOptions contains the optional parameters for the TracksClient.List method.
func (client *TracksClient) NewListPager(resourceGroupName string, accountName string, assetName string, options *TracksClientListOptions) *runtime.Pager[TracksClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[TracksClientListResponse]{
		More: func(page TracksClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TracksClientListResponse) (TracksClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
			if err != nil {
				return TracksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TracksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TracksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TracksClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *TracksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TracksClient) listHandleResponse(resp *http.Response) (TracksClientListResponse, error) {
	result := TracksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetTrackCollection); err != nil {
		return TracksClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing Track in the asset
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// trackName - The Asset Track name.
// parameters - The request parameters
// options - TracksClientBeginUpdateOptions contains the optional parameters for the TracksClient.BeginUpdate method.
func (client *TracksClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, parameters AssetTrack, options *TracksClientBeginUpdateOptions) (*armruntime.Poller[TracksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, assetName, trackName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[TracksClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[TracksClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates an existing Track in the asset
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TracksClient) update(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, parameters AssetTrack, options *TracksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, assetName, trackName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *TracksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, parameters AssetTrack, options *TracksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if trackName == "" {
		return nil, errors.New("parameter trackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trackName}", url.PathEscape(trackName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdateTrackData - Update the track data. Call this API after any changes are made to the track data stored in the
// asset container. For example, you have modified the WebVTT captions file in the Azure blob storage
// container for the asset, viewers will not see the new version of the captions unless this API is called. Note, the changes
// may not be reflected immediately. CDN cache may also need to be purged if
// applicable.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// trackName - The Asset Track name.
// options - TracksClientBeginUpdateTrackDataOptions contains the optional parameters for the TracksClient.BeginUpdateTrackData
// method.
func (client *TracksClient) BeginUpdateTrackData(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientBeginUpdateTrackDataOptions) (*armruntime.Poller[TracksClientUpdateTrackDataResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTrackData(ctx, resourceGroupName, accountName, assetName, trackName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[TracksClientUpdateTrackDataResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[TracksClientUpdateTrackDataResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateTrackData - Update the track data. Call this API after any changes are made to the track data stored in the asset
// container. For example, you have modified the WebVTT captions file in the Azure blob storage
// container for the asset, viewers will not see the new version of the captions unless this API is called. Note, the changes
// may not be reflected immediately. CDN cache may also need to be purged if
// applicable.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TracksClient) updateTrackData(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientBeginUpdateTrackDataOptions) (*http.Response, error) {
	req, err := client.updateTrackDataCreateRequest(ctx, resourceGroupName, accountName, assetName, trackName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateTrackDataCreateRequest creates the UpdateTrackData request.
func (client *TracksClient) updateTrackDataCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, options *TracksClientBeginUpdateTrackDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}/updateTrackData"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if trackName == "" {
		return nil, errors.New("parameter trackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trackName}", url.PathEscape(trackName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
