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
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualMachineImagesClient contains the methods for the VirtualMachineImages group.
// Don't use this type directly, use NewVirtualMachineImagesClient() instead.
type VirtualMachineImagesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineImagesClient creates a new instance of VirtualMachineImagesClient with the specified values.
func NewVirtualMachineImagesClient(con *armcore.Connection, subscriptionID string) VirtualMachineImagesClient {
	return VirtualMachineImagesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualMachineImagesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets a virtual machine image.
func (client VirtualMachineImagesClient) Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesGetOptions) (VirtualMachineImageResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publisherName, offer, skus, version, options)
	if err != nil {
		return VirtualMachineImageResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VirtualMachineImageResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VirtualMachineImagesClient) getCreateRequest(ctx context.Context, location string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client VirtualMachineImagesClient) getHandleResponse(resp *azcore.Response) (VirtualMachineImageResponse, error) {
	result := VirtualMachineImageResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineImage)
	return result, err
}

// getHandleError handles the Get error response.
func (client VirtualMachineImagesClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
func (client VirtualMachineImagesClient) List(ctx context.Context, location string, publisherName string, offer string, skus string, options *VirtualMachineImagesListOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listCreateRequest(ctx, location, publisherName, offer, skus, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listHandleError(resp)
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client VirtualMachineImagesClient) listCreateRequest(ctx context.Context, location string, publisherName string, offer string, skus string, options *VirtualMachineImagesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	if options != nil && options.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		query.Set("$orderby", *options.Orderby)
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client VirtualMachineImagesClient) listHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
	return result, err
}

// listHandleError handles the List error response.
func (client VirtualMachineImagesClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListOffers - Gets a list of virtual machine image offers for the specified location and publisher.
func (client VirtualMachineImagesClient) ListOffers(ctx context.Context, location string, publisherName string, options *VirtualMachineImagesListOffersOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listOffersCreateRequest(ctx, location, publisherName, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listOffersHandleError(resp)
	}
	result, err := client.listOffersHandleResponse(resp)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return result, nil
}

// listOffersCreateRequest creates the ListOffers request.
func (client VirtualMachineImagesClient) listOffersCreateRequest(ctx context.Context, location string, publisherName string, options *VirtualMachineImagesListOffersOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listOffersHandleResponse handles the ListOffers response.
func (client VirtualMachineImagesClient) listOffersHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
	return result, err
}

// listOffersHandleError handles the ListOffers error response.
func (client VirtualMachineImagesClient) listOffersHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location.
func (client VirtualMachineImagesClient) ListPublishers(ctx context.Context, location string, options *VirtualMachineImagesListPublishersOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listPublishersCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listPublishersHandleError(resp)
	}
	result, err := client.listPublishersHandleResponse(resp)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return result, nil
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client VirtualMachineImagesClient) listPublishersCreateRequest(ctx context.Context, location string, options *VirtualMachineImagesListPublishersOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listPublishersHandleResponse handles the ListPublishers response.
func (client VirtualMachineImagesClient) listPublishersHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
	return result, err
}

// listPublishersHandleError handles the ListPublishers error response.
func (client VirtualMachineImagesClient) listPublishersHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListSKUs - Gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
func (client VirtualMachineImagesClient) ListSKUs(ctx context.Context, location string, publisherName string, offer string, options *VirtualMachineImagesListSKUsOptions) (VirtualMachineImageResourceArrayResponse, error) {
	req, err := client.listSkUsCreateRequest(ctx, location, publisherName, offer, options)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineImageResourceArrayResponse{}, client.listSkUsHandleError(resp)
	}
	result, err := client.listSkUsHandleResponse(resp)
	if err != nil {
		return VirtualMachineImageResourceArrayResponse{}, err
	}
	return result, nil
}

// listSkUsCreateRequest creates the ListSKUs request.
func (client VirtualMachineImagesClient) listSkUsCreateRequest(ctx context.Context, location string, publisherName string, offer string, options *VirtualMachineImagesListSKUsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listSkUsHandleResponse handles the ListSKUs response.
func (client VirtualMachineImagesClient) listSkUsHandleResponse(resp *azcore.Response) (VirtualMachineImageResourceArrayResponse, error) {
	result := VirtualMachineImageResourceArrayResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineImageResourceArray)
	return result, err
}

// listSkUsHandleError handles the ListSKUs error response.
func (client VirtualMachineImagesClient) listSkUsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
