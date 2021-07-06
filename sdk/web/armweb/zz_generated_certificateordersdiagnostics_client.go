// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

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

// CertificateOrdersDiagnosticsClient contains the methods for the CertificateOrdersDiagnostics group.
// Don't use this type directly, use NewCertificateOrdersDiagnosticsClient() instead.
type CertificateOrdersDiagnosticsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewCertificateOrdersDiagnosticsClient creates a new instance of CertificateOrdersDiagnosticsClient with the specified values.
func NewCertificateOrdersDiagnosticsClient(con *armcore.Connection, subscriptionID string) *CertificateOrdersDiagnosticsClient {
	return &CertificateOrdersDiagnosticsClient{con: con, subscriptionID: subscriptionID}
}

// GetAppServiceCertificateOrderDetectorResponse - Description for Microsoft.CertificateRegistration call to get a detector response from App Lens.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponse(ctx context.Context, resourceGroupName string, certificateOrderName string, detectorName string, options *CertificateOrdersDiagnosticsGetAppServiceCertificateOrderDetectorResponseOptions) (DetectorResponseResponse, error) {
	req, err := client.getAppServiceCertificateOrderDetectorResponseCreateRequest(ctx, resourceGroupName, certificateOrderName, detectorName, options)
	if err != nil {
		return DetectorResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DetectorResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DetectorResponseResponse{}, client.getAppServiceCertificateOrderDetectorResponseHandleError(resp)
	}
	return client.getAppServiceCertificateOrderDetectorResponseHandleResponse(resp)
}

// getAppServiceCertificateOrderDetectorResponseCreateRequest creates the GetAppServiceCertificateOrderDetectorResponse request.
func (client *CertificateOrdersDiagnosticsClient) getAppServiceCertificateOrderDetectorResponseCreateRequest(ctx context.Context, resourceGroupName string, certificateOrderName string, detectorName string, options *CertificateOrdersDiagnosticsGetAppServiceCertificateOrderDetectorResponseOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/detectors/{detectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if certificateOrderName == "" {
		return nil, errors.New("parameter certificateOrderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateOrderName}", url.PathEscape(certificateOrderName))
	if detectorName == "" {
		return nil, errors.New("parameter detectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{detectorName}", url.PathEscape(detectorName))
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
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", options.StartTime.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", options.EndTime.Format(time.RFC3339Nano))
	}
	if options != nil && options.TimeGrain != nil {
		reqQP.Set("timeGrain", *options.TimeGrain)
	}
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getAppServiceCertificateOrderDetectorResponseHandleResponse handles the GetAppServiceCertificateOrderDetectorResponse response.
func (client *CertificateOrdersDiagnosticsClient) getAppServiceCertificateOrderDetectorResponseHandleResponse(resp *azcore.Response) (DetectorResponseResponse, error) {
	var val *DetectorResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DetectorResponseResponse{}, err
	}
	return DetectorResponseResponse{RawResponse: resp.Response, DetectorResponse: val}, nil
}

// getAppServiceCertificateOrderDetectorResponseHandleError handles the GetAppServiceCertificateOrderDetectorResponse error response.
func (client *CertificateOrdersDiagnosticsClient) getAppServiceCertificateOrderDetectorResponseHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListAppServiceCertificateOrderDetectorResponse - Description for Microsoft.CertificateRegistration to get the list of detectors for this RP.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponse(resourceGroupName string, certificateOrderName string, options *CertificateOrdersDiagnosticsListAppServiceCertificateOrderDetectorResponseOptions) DetectorResponseCollectionPager {
	return &detectorResponseCollectionPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAppServiceCertificateOrderDetectorResponseCreateRequest(ctx, resourceGroupName, certificateOrderName, options)
		},
		responder: client.listAppServiceCertificateOrderDetectorResponseHandleResponse,
		errorer:   client.listAppServiceCertificateOrderDetectorResponseHandleError,
		advancer: func(ctx context.Context, resp DetectorResponseCollectionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DetectorResponseCollection.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAppServiceCertificateOrderDetectorResponseCreateRequest creates the ListAppServiceCertificateOrderDetectorResponse request.
func (client *CertificateOrdersDiagnosticsClient) listAppServiceCertificateOrderDetectorResponseCreateRequest(ctx context.Context, resourceGroupName string, certificateOrderName string, options *CertificateOrdersDiagnosticsListAppServiceCertificateOrderDetectorResponseOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/detectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if certificateOrderName == "" {
		return nil, errors.New("parameter certificateOrderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateOrderName}", url.PathEscape(certificateOrderName))
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAppServiceCertificateOrderDetectorResponseHandleResponse handles the ListAppServiceCertificateOrderDetectorResponse response.
func (client *CertificateOrdersDiagnosticsClient) listAppServiceCertificateOrderDetectorResponseHandleResponse(resp *azcore.Response) (DetectorResponseCollectionResponse, error) {
	var val *DetectorResponseCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DetectorResponseCollectionResponse{}, err
	}
	return DetectorResponseCollectionResponse{RawResponse: resp.Response, DetectorResponseCollection: val}, nil
}

// listAppServiceCertificateOrderDetectorResponseHandleError handles the ListAppServiceCertificateOrderDetectorResponse error response.
func (client *CertificateOrdersDiagnosticsClient) listAppServiceCertificateOrderDetectorResponseHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
