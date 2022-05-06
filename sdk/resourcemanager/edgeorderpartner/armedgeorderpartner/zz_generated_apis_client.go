//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner

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

// APISClient contains the methods for the EdgeOrderPartnerAPIS group.
// Don't use this type directly, use NewAPISClient() instead.
type APISClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPISClient creates a new instance of APISClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPISClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APISClient, error) {
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
	client := &APISClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListOperationsPartnerPager - This method gets all the operations that are exposed for customer.
// If the operation fails it returns an *azcore.ResponseError type.
// options - APISClientListOperationsPartnerOptions contains the optional parameters for the APISClient.ListOperationsPartner
// method.
func (client *APISClient) NewListOperationsPartnerPager(options *APISClientListOperationsPartnerOptions) *runtime.Pager[APISClientListOperationsPartnerResponse] {
	return runtime.NewPager(runtime.PageProcessor[APISClientListOperationsPartnerResponse]{
		More: func(page APISClientListOperationsPartnerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APISClientListOperationsPartnerResponse) (APISClientListOperationsPartnerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOperationsPartnerCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APISClientListOperationsPartnerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APISClientListOperationsPartnerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APISClientListOperationsPartnerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOperationsPartnerHandleResponse(resp)
		},
	})
}

// listOperationsPartnerCreateRequest creates the ListOperationsPartner request.
func (client *APISClient) listOperationsPartnerCreateRequest(ctx context.Context, options *APISClientListOperationsPartnerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EdgeOrderPartner/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsPartnerHandleResponse handles the ListOperationsPartner response.
func (client *APISClient) listOperationsPartnerHandleResponse(resp *http.Response) (APISClientListOperationsPartnerResponse, error) {
	result := APISClientListOperationsPartnerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return APISClientListOperationsPartnerResponse{}, err
	}
	return result, nil
}

// BeginManageInventoryMetadata - API for updating inventory metadata and inventory configuration
// If the operation fails it returns an *azcore.ResponseError type.
// familyIdentifier - Unique identifier for the product family
// location - The location of the resource
// serialNumber - The serial number of the device
// manageInventoryMetadataRequest - Updates inventory metadata and inventory configuration
// options - APISClientBeginManageInventoryMetadataOptions contains the optional parameters for the APISClient.BeginManageInventoryMetadata
// method.
func (client *APISClient) BeginManageInventoryMetadata(ctx context.Context, familyIdentifier string, location string, serialNumber string, manageInventoryMetadataRequest ManageInventoryMetadataRequest, options *APISClientBeginManageInventoryMetadataOptions) (*armruntime.Poller[APISClientManageInventoryMetadataResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.manageInventoryMetadata(ctx, familyIdentifier, location, serialNumber, manageInventoryMetadataRequest, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[APISClientManageInventoryMetadataResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[APISClientManageInventoryMetadataResponse](options.ResumeToken, client.pl, nil)
	}
}

// ManageInventoryMetadata - API for updating inventory metadata and inventory configuration
// If the operation fails it returns an *azcore.ResponseError type.
func (client *APISClient) manageInventoryMetadata(ctx context.Context, familyIdentifier string, location string, serialNumber string, manageInventoryMetadataRequest ManageInventoryMetadataRequest, options *APISClientBeginManageInventoryMetadataOptions) (*http.Response, error) {
	req, err := client.manageInventoryMetadataCreateRequest(ctx, familyIdentifier, location, serialNumber, manageInventoryMetadataRequest, options)
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

// manageInventoryMetadataCreateRequest creates the ManageInventoryMetadata request.
func (client *APISClient) manageInventoryMetadataCreateRequest(ctx context.Context, familyIdentifier string, location string, serialNumber string, manageInventoryMetadataRequest ManageInventoryMetadataRequest, options *APISClientBeginManageInventoryMetadataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrderPartner/locations/{location}/productFamilies/{familyIdentifier}/inventories/{serialNumber}/manageInventoryMetadata"
	if familyIdentifier == "" {
		return nil, errors.New("parameter familyIdentifier cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{familyIdentifier}", url.PathEscape(familyIdentifier))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if serialNumber == "" {
		return nil, errors.New("parameter serialNumber cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serialNumber}", url.PathEscape(serialNumber))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, manageInventoryMetadataRequest)
}

// ManageLink - API for linking management resource with inventory
// If the operation fails it returns an *azcore.ResponseError type.
// familyIdentifier - Unique identifier for the product family
// location - The location of the resource
// serialNumber - The serial number of the device
// manageLinkRequest - Links the management resource to the inventory
// options - APISClientManageLinkOptions contains the optional parameters for the APISClient.ManageLink method.
func (client *APISClient) ManageLink(ctx context.Context, familyIdentifier string, location string, serialNumber string, manageLinkRequest ManageLinkRequest, options *APISClientManageLinkOptions) (APISClientManageLinkResponse, error) {
	req, err := client.manageLinkCreateRequest(ctx, familyIdentifier, location, serialNumber, manageLinkRequest, options)
	if err != nil {
		return APISClientManageLinkResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APISClientManageLinkResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return APISClientManageLinkResponse{}, runtime.NewResponseError(resp)
	}
	return APISClientManageLinkResponse{}, nil
}

// manageLinkCreateRequest creates the ManageLink request.
func (client *APISClient) manageLinkCreateRequest(ctx context.Context, familyIdentifier string, location string, serialNumber string, manageLinkRequest ManageLinkRequest, options *APISClientManageLinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrderPartner/locations/{location}/productFamilies/{familyIdentifier}/inventories/{serialNumber}/manageLink"
	if familyIdentifier == "" {
		return nil, errors.New("parameter familyIdentifier cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{familyIdentifier}", url.PathEscape(familyIdentifier))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if serialNumber == "" {
		return nil, errors.New("parameter serialNumber cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serialNumber}", url.PathEscape(serialNumber))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, manageLinkRequest)
}

// NewSearchInventoriesPager - API for Search inventories
// If the operation fails it returns an *azcore.ResponseError type.
// searchInventoriesRequest - Searches inventories with the given filters and returns in the form of a list
// options - APISClientSearchInventoriesOptions contains the optional parameters for the APISClient.SearchInventories method.
func (client *APISClient) NewSearchInventoriesPager(searchInventoriesRequest SearchInventoriesRequest, options *APISClientSearchInventoriesOptions) *runtime.Pager[APISClientSearchInventoriesResponse] {
	return runtime.NewPager(runtime.PageProcessor[APISClientSearchInventoriesResponse]{
		More: func(page APISClientSearchInventoriesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APISClientSearchInventoriesResponse) (APISClientSearchInventoriesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.searchInventoriesCreateRequest(ctx, searchInventoriesRequest, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APISClientSearchInventoriesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APISClientSearchInventoriesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APISClientSearchInventoriesResponse{}, runtime.NewResponseError(resp)
			}
			return client.searchInventoriesHandleResponse(resp)
		},
	})
}

// searchInventoriesCreateRequest creates the SearchInventories request.
func (client *APISClient) searchInventoriesCreateRequest(ctx context.Context, searchInventoriesRequest SearchInventoriesRequest, options *APISClientSearchInventoriesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrderPartner/searchInventories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, searchInventoriesRequest)
}

// searchInventoriesHandleResponse handles the SearchInventories response.
func (client *APISClient) searchInventoriesHandleResponse(resp *http.Response) (APISClientSearchInventoriesResponse, error) {
	result := APISClientSearchInventoriesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerInventoryList); err != nil {
		return APISClientSearchInventoriesResponse{}, err
	}
	return result, nil
}
