//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BillingProfilesClient contains the methods for the BillingProfiles group.
// Don't use this type directly, use NewBillingProfilesClient() instead.
type BillingProfilesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewBillingProfilesClient creates a new instance of BillingProfilesClient with the specified values.
func NewBillingProfilesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *BillingProfilesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BillingProfilesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates a billing profile. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement
// or Microsoft Partner Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingProfilesClient) BeginCreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, parameters BillingProfile, options *BillingProfilesBeginCreateOrUpdateOptions) (BillingProfilesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, billingAccountName, billingProfileName, parameters, options)
	if err != nil {
		return BillingProfilesCreateOrUpdatePollerResponse{}, err
	}
	result := BillingProfilesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("BillingProfilesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return BillingProfilesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &BillingProfilesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a billing profile. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement
// or Microsoft Partner Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingProfilesClient) createOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, parameters BillingProfile, options *BillingProfilesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, billingAccountName, billingProfileName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BillingProfilesClient) createOrUpdateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, parameters BillingProfile, options *BillingProfilesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *BillingProfilesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a billing profile by its ID. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft
// Partner Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingProfilesClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, options *BillingProfilesGetOptions) (BillingProfilesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return BillingProfilesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BillingProfilesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingProfilesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BillingProfilesClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *BillingProfilesGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BillingProfilesClient) getHandleResponse(resp *http.Response) (BillingProfilesGetResponse, error) {
	result := BillingProfilesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingProfile); err != nil {
		return BillingProfilesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BillingProfilesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByBillingAccount - Lists the billing profiles that a user has access to. The operation is supported for billing accounts with agreement type Microsoft
// Customer Agreement or Microsoft Partner Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingProfilesClient) ListByBillingAccount(billingAccountName string, options *BillingProfilesListByBillingAccountOptions) *BillingProfilesListByBillingAccountPager {
	return &BillingProfilesListByBillingAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
		},
		advancer: func(ctx context.Context, resp BillingProfilesListByBillingAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BillingProfileListResult.NextLink)
		},
	}
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *BillingProfilesClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *BillingProfilesListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *BillingProfilesClient) listByBillingAccountHandleResponse(resp *http.Response) (BillingProfilesListByBillingAccountResponse, error) {
	result := BillingProfilesListByBillingAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingProfileListResult); err != nil {
		return BillingProfilesListByBillingAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByBillingAccountHandleError handles the ListByBillingAccount error response.
func (client *BillingProfilesClient) listByBillingAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
