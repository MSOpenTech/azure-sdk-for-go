// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualMachineImagesOperations contains the methods for the VirtualMachineImages group.
type VirtualMachineImagesOperations interface {
	// Get - Gets a virtual machine image.
	Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (*VirtualMachineImageResponse, error)
	// List - Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
	List(ctx context.Context, location string, publisherName string, offer string, skus string, virtualMachineImagesListOptions *VirtualMachineImagesListOptions) (*VirtualMachineImageResourceArrayResponse, error)
	// ListOffers - Gets a list of virtual machine image offers for the specified location and publisher.
	ListOffers(ctx context.Context, location string, publisherName string) (*VirtualMachineImageResourceArrayResponse, error)
	// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location.
	ListPublishers(ctx context.Context, location string) (*VirtualMachineImageResourceArrayResponse, error)
	// ListSkus - Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
	ListSkus(ctx context.Context, location string, publisherName string, offer string) (*VirtualMachineImageResourceArrayResponse, error)
}

// virtualMachineImagesOperations implements the VirtualMachineImagesOperations interface.
type virtualMachineImagesOperations struct {
	*Client
	subscriptionID string
}

// Get - Gets a virtual machine image.
func (client *virtualMachineImagesOperations) Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (*VirtualMachineImageResponse, error) {
	req, err := client.getCreateRequest(location, publisherName, offer, skus, version)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *virtualMachineImagesOperations) getCreateRequest(location string, publisherName string, offer string, skus string, version string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *virtualMachineImagesOperations) getHandleResponse(resp *azcore.Response) (*VirtualMachineImageResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VirtualMachineImageResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineImage)
}

// getHandleError handles the Get error response.
func (client *virtualMachineImagesOperations) getHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// List - Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
func (client *virtualMachineImagesOperations) List(ctx context.Context, location string, publisherName string, offer string, skus string, virtualMachineImagesListOptions *VirtualMachineImagesListOptions) (*VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listCreateRequest(location, publisherName, offer, skus, virtualMachineImagesListOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client *virtualMachineImagesOperations) listCreateRequest(location string, publisherName string, offer string, skus string, virtualMachineImagesListOptions *VirtualMachineImagesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if virtualMachineImagesListOptions != nil && virtualMachineImagesListOptions.Expand != nil {
		query.Set("$expand", *virtualMachineImagesListOptions.Expand)
	}
	if virtualMachineImagesListOptions != nil && virtualMachineImagesListOptions.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*virtualMachineImagesListOptions.Top), 10))
	}
	if virtualMachineImagesListOptions != nil && virtualMachineImagesListOptions.Orderby != nil {
		query.Set("$orderby", *virtualMachineImagesListOptions.Orderby)
	}
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *virtualMachineImagesOperations) listHandleResponse(resp *azcore.Response) (*VirtualMachineImageResourceArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
}

// listHandleError handles the List error response.
func (client *virtualMachineImagesOperations) listHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// ListOffers - Gets a list of virtual machine image offers for the specified location and publisher.
func (client *virtualMachineImagesOperations) ListOffers(ctx context.Context, location string, publisherName string) (*VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listOffersCreateRequest(location, publisherName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listOffersHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listOffersCreateRequest creates the ListOffers request.
func (client *virtualMachineImagesOperations) listOffersCreateRequest(location string, publisherName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listOffersHandleResponse handles the ListOffers response.
func (client *virtualMachineImagesOperations) listOffersHandleResponse(resp *azcore.Response) (*VirtualMachineImageResourceArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listOffersHandleError(resp)
	}
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
}

// listOffersHandleError handles the ListOffers error response.
func (client *virtualMachineImagesOperations) listOffersHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location.
func (client *virtualMachineImagesOperations) ListPublishers(ctx context.Context, location string) (*VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listPublishersCreateRequest(location)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listPublishersHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client *virtualMachineImagesOperations) listPublishersCreateRequest(location string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listPublishersHandleResponse handles the ListPublishers response.
func (client *virtualMachineImagesOperations) listPublishersHandleResponse(resp *azcore.Response) (*VirtualMachineImageResourceArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listPublishersHandleError(resp)
	}
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
}

// listPublishersHandleError handles the ListPublishers error response.
func (client *virtualMachineImagesOperations) listPublishersHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}

// ListSkus - Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
func (client *virtualMachineImagesOperations) ListSkus(ctx context.Context, location string, publisherName string, offer string) (*VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listSkusCreateRequest(location, publisherName, offer)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listSkusHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listSkusCreateRequest creates the ListSkus request.
func (client *virtualMachineImagesOperations) listSkusCreateRequest(location string, publisherName string, offer string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listSkusHandleResponse handles the ListSkus response.
func (client *virtualMachineImagesOperations) listSkusHandleResponse(resp *azcore.Response) (*VirtualMachineImageResourceArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listSkusHandleError(resp)
	}
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
}

// listSkusHandleError handles the ListSkus error response.
func (client *virtualMachineImagesOperations) listSkusHandleError(resp *azcore.Response) error {
	return errors.New(resp.Status)
}
