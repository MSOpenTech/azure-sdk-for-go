// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// RoleEligibilitySchedulesClient contains the methods for the RoleEligibilitySchedules group.
// Don't use this type directly, use NewRoleEligibilitySchedulesClient() instead.
type RoleEligibilitySchedulesClient struct {
	con *armcore.Connection
}

// NewRoleEligibilitySchedulesClient creates a new instance of RoleEligibilitySchedulesClient with the specified values.
func NewRoleEligibilitySchedulesClient(con *armcore.Connection) *RoleEligibilitySchedulesClient {
	return &RoleEligibilitySchedulesClient{con: con}
}

// Get - Get the specified role eligibility schedule for a resource scope
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilitySchedulesClient) Get(ctx context.Context, scope string, roleEligibilityScheduleName string, options *RoleEligibilitySchedulesGetOptions) (RoleEligibilityScheduleResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleName, options)
	if err != nil {
		return RoleEligibilityScheduleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleEligibilityScheduleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleEligibilityScheduleResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilitySchedulesClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleName string, options *RoleEligibilitySchedulesGetOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilitySchedules/{roleEligibilityScheduleName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleName}", url.PathEscape(roleEligibilityScheduleName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilitySchedulesClient) getHandleResponse(resp *azcore.Response) (RoleEligibilityScheduleResponse, error) {
	var val *RoleEligibilitySchedule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleEligibilityScheduleResponse{}, err
	}
	return RoleEligibilityScheduleResponse{RawResponse: resp.Response, RoleEligibilitySchedule: val}, nil
}

// getHandleError handles the Get error response.
func (client *RoleEligibilitySchedulesClient) getHandleError(resp *azcore.Response) error {
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

// ListForScope - Gets role eligibility schedules for a resource scope.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilitySchedulesClient) ListForScope(scope string, options *RoleEligibilitySchedulesListForScopeOptions) RoleEligibilityScheduleListResultPager {
	return &roleEligibilityScheduleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		responder: client.listForScopeHandleResponse,
		errorer:   client.listForScopeHandleError,
		advancer: func(ctx context.Context, resp RoleEligibilityScheduleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleEligibilityScheduleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilitySchedulesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilitySchedulesListForScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilitySchedules"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilitySchedulesClient) listForScopeHandleResponse(resp *azcore.Response) (RoleEligibilityScheduleListResultResponse, error) {
	var val *RoleEligibilityScheduleListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleEligibilityScheduleListResultResponse{}, err
	}
	return RoleEligibilityScheduleListResultResponse{RawResponse: resp.Response, RoleEligibilityScheduleListResult: val}, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleEligibilitySchedulesClient) listForScopeHandleError(resp *azcore.Response) error {
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
