// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// RegionsOperations contains the methods for the Regions group.
type RegionsOperations interface {
	// ListBySku - Gets the available Regions for a given sku
	ListBySku(sku string) (MessagingRegionsListResultPager, error)
}

// regionsOperations implements the RegionsOperations interface.
type regionsOperations struct {
	*Client
	subscriptionID string
}

// ListBySku - Gets the available Regions for a given sku
func (client *regionsOperations) ListBySku(sku string) (MessagingRegionsListResultPager, error) {
	req, err := client.listBySkuCreateRequest(sku)
	if err != nil {
		return nil, err
	}
	return &messagingRegionsListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listBySkuHandleResponse,
		advancer: func(resp *MessagingRegionsListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.MessagingRegionsListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.MessagingRegionsListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listBySkuCreateRequest creates the ListBySku request.
func (client *regionsOperations) listBySkuCreateRequest(sku string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/sku/{sku}/regions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{sku}", url.PathEscape(sku))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2017-04-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listBySkuHandleResponse handles the ListBySku response.
func (client *regionsOperations) listBySkuHandleResponse(resp *azcore.Response) (*MessagingRegionsListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listBySkuHandleError(resp)
	}
	result := MessagingRegionsListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.MessagingRegionsListResult)
}

// listBySkuHandleError handles the ListBySku error response.
func (client *regionsOperations) listBySkuHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
