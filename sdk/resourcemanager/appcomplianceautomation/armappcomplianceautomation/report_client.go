//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ReportClient contains the methods for the Report group.
// Don't use this type directly, use NewReportClient() instead.
type ReportClient struct {
	internal *arm.Client
}

// NewReportClient creates a new instance of ReportClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReportClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReportClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReportClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new AppComplianceAutomation report or update an exiting AppComplianceAutomation report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - properties - Parameters for the create or update operation
//   - options - ReportClientBeginCreateOrUpdateOptions contains the optional parameters for the ReportClient.BeginCreateOrUpdate
//     method.
func (client *ReportClient) BeginCreateOrUpdate(ctx context.Context, reportName string, properties ReportResource, options *ReportClientBeginCreateOrUpdateOptions) (*runtime.Poller[ReportClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, reportName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReportClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReportClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new AppComplianceAutomation report or update an exiting AppComplianceAutomation report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ReportClient) createOrUpdate(ctx context.Context, reportName string, properties ReportResource, options *ReportClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ReportClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, reportName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ReportClient) createOrUpdateCreateRequest(ctx context.Context, reportName string, properties ReportResource, options *ReportClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an AppComplianceAutomation report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - ReportClientBeginDeleteOptions contains the optional parameters for the ReportClient.BeginDelete method.
func (client *ReportClient) BeginDelete(ctx context.Context, reportName string, options *ReportClientBeginDeleteOptions) (*runtime.Poller[ReportClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, reportName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReportClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReportClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete an AppComplianceAutomation report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ReportClient) deleteOperation(ctx context.Context, reportName string, options *ReportClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ReportClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, reportName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReportClient) deleteCreateRequest(ctx context.Context, reportName string, options *ReportClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginFix - Fix the AppComplianceAutomation report error. e.g: App Compliance Automation Tool service unregistered, automation
// removed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - ReportClientBeginFixOptions contains the optional parameters for the ReportClient.BeginFix method.
func (client *ReportClient) BeginFix(ctx context.Context, reportName string, options *ReportClientBeginFixOptions) (*runtime.Poller[ReportClientFixResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.fix(ctx, reportName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReportClientFixResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReportClientFixResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Fix - Fix the AppComplianceAutomation report error. e.g: App Compliance Automation Tool service unregistered, automation
// removed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ReportClient) fix(ctx context.Context, reportName string, options *ReportClientBeginFixOptions) (*http.Response, error) {
	var err error
	const operationName = "ReportClient.BeginFix"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fixCreateRequest(ctx, reportName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// fixCreateRequest creates the Fix request.
func (client *ReportClient) fixCreateRequest(ctx context.Context, reportName string, options *ReportClientBeginFixOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/fix"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the AppComplianceAutomation report and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - ReportClientGetOptions contains the optional parameters for the ReportClient.Get method.
func (client *ReportClient) Get(ctx context.Context, reportName string, options *ReportClientGetOptions) (ReportClientGetResponse, error) {
	var err error
	const operationName = "ReportClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, reportName, options)
	if err != nil {
		return ReportClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReportClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReportClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReportClient) getCreateRequest(ctx context.Context, reportName string, options *ReportClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReportClient) getHandleResponse(resp *http.Response) (ReportClientGetResponse, error) {
	result := ReportClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportResource); err != nil {
		return ReportClientGetResponse{}, err
	}
	return result, nil
}

// GetScopingQuestions - Fix the AppComplianceAutomation report error. e.g: App Compliance Automation Tool service unregistered,
// automation removed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - ReportClientGetScopingQuestionsOptions contains the optional parameters for the ReportClient.GetScopingQuestions
//     method.
func (client *ReportClient) GetScopingQuestions(ctx context.Context, reportName string, options *ReportClientGetScopingQuestionsOptions) (ReportClientGetScopingQuestionsResponse, error) {
	var err error
	const operationName = "ReportClient.GetScopingQuestions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getScopingQuestionsCreateRequest(ctx, reportName, options)
	if err != nil {
		return ReportClientGetScopingQuestionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReportClientGetScopingQuestionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReportClientGetScopingQuestionsResponse{}, err
	}
	resp, err := client.getScopingQuestionsHandleResponse(httpResp)
	return resp, err
}

// getScopingQuestionsCreateRequest creates the GetScopingQuestions request.
func (client *ReportClient) getScopingQuestionsCreateRequest(ctx context.Context, reportName string, options *ReportClientGetScopingQuestionsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/getScopingQuestions"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getScopingQuestionsHandleResponse handles the GetScopingQuestions response.
func (client *ReportClient) getScopingQuestionsHandleResponse(resp *http.Response) (ReportClientGetScopingQuestionsResponse, error) {
	result := ReportClientGetScopingQuestionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopingQuestions); err != nil {
		return ReportClientGetScopingQuestionsResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the AppComplianceAutomation report list for the tenant.
//
// Generated from API version 2024-06-27
//   - options - ReportClientListOptions contains the optional parameters for the ReportClient.NewListPager method.
func (client *ReportClient) NewListPager(options *ReportClientListOptions) *runtime.Pager[ReportClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReportClientListResponse]{
		More: func(page ReportClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportClientListResponse) (ReportClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReportClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ReportClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReportClient) listCreateRequest(ctx context.Context, options *ReportClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-27")
	if options != nil && options.OfferGUID != nil {
		reqQP.Set("offerGuid", *options.OfferGUID)
	}
	if options != nil && options.ReportCreatorTenantID != nil {
		reqQP.Set("reportCreatorTenantId", *options.ReportCreatorTenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReportClient) listHandleResponse(resp *http.Response) (ReportClientListResponse, error) {
	result := ReportClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportResourceListResult); err != nil {
		return ReportClientListResponse{}, err
	}
	return result, nil
}

// NestedResourceCheckNameAvailability - Checks the report's nested resource name availability, e.g: Webhooks, Evidences,
// Snapshots.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - body - NameAvailabilityRequest object.
//   - options - ReportClientNestedResourceCheckNameAvailabilityOptions contains the optional parameters for the ReportClient.NestedResourceCheckNameAvailability
//     method.
func (client *ReportClient) NestedResourceCheckNameAvailability(ctx context.Context, reportName string, body CheckNameAvailabilityRequest, options *ReportClientNestedResourceCheckNameAvailabilityOptions) (ReportClientNestedResourceCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "ReportClient.NestedResourceCheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.nestedResourceCheckNameAvailabilityCreateRequest(ctx, reportName, body, options)
	if err != nil {
		return ReportClientNestedResourceCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReportClientNestedResourceCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReportClientNestedResourceCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.nestedResourceCheckNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// nestedResourceCheckNameAvailabilityCreateRequest creates the NestedResourceCheckNameAvailability request.
func (client *ReportClient) nestedResourceCheckNameAvailabilityCreateRequest(ctx context.Context, reportName string, body CheckNameAvailabilityRequest, options *ReportClientNestedResourceCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/checkNameAvailability"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// nestedResourceCheckNameAvailabilityHandleResponse handles the NestedResourceCheckNameAvailability response.
func (client *ReportClient) nestedResourceCheckNameAvailabilityHandleResponse(resp *http.Response) (ReportClientNestedResourceCheckNameAvailabilityResponse, error) {
	result := ReportClientNestedResourceCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ReportClientNestedResourceCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginSyncCertRecord - Synchronize attestation record from app compliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - body - Parameters for synchronize certification record operation
//   - options - ReportClientBeginSyncCertRecordOptions contains the optional parameters for the ReportClient.BeginSyncCertRecord
//     method.
func (client *ReportClient) BeginSyncCertRecord(ctx context.Context, reportName string, body SyncCertRecordRequest, options *ReportClientBeginSyncCertRecordOptions) (*runtime.Poller[ReportClientSyncCertRecordResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.syncCertRecord(ctx, reportName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReportClientSyncCertRecordResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReportClientSyncCertRecordResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// SyncCertRecord - Synchronize attestation record from app compliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ReportClient) syncCertRecord(ctx context.Context, reportName string, body SyncCertRecordRequest, options *ReportClientBeginSyncCertRecordOptions) (*http.Response, error) {
	var err error
	const operationName = "ReportClient.BeginSyncCertRecord"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.syncCertRecordCreateRequest(ctx, reportName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// syncCertRecordCreateRequest creates the SyncCertRecord request.
func (client *ReportClient) syncCertRecordCreateRequest(ctx context.Context, reportName string, body SyncCertRecordRequest, options *ReportClientBeginSyncCertRecordOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/syncCertRecord"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Update an exiting AppComplianceAutomation report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - properties - Parameters for the create or update operation
//   - options - ReportClientBeginUpdateOptions contains the optional parameters for the ReportClient.BeginUpdate method.
func (client *ReportClient) BeginUpdate(ctx context.Context, reportName string, properties ReportResourcePatch, options *ReportClientBeginUpdateOptions) (*runtime.Poller[ReportClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, reportName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReportClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReportClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an exiting AppComplianceAutomation report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ReportClient) update(ctx context.Context, reportName string, properties ReportResourcePatch, options *ReportClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ReportClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, reportName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReportClient) updateCreateRequest(ctx context.Context, reportName string, properties ReportResourcePatch, options *ReportClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginVerify - Verify the AppComplianceAutomation report health status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - ReportClientBeginVerifyOptions contains the optional parameters for the ReportClient.BeginVerify method.
func (client *ReportClient) BeginVerify(ctx context.Context, reportName string, options *ReportClientBeginVerifyOptions) (*runtime.Poller[ReportClientVerifyResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.verify(ctx, reportName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReportClientVerifyResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReportClientVerifyResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Verify - Verify the AppComplianceAutomation report health status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
func (client *ReportClient) verify(ctx context.Context, reportName string, options *ReportClientBeginVerifyOptions) (*http.Response, error) {
	var err error
	const operationName = "ReportClient.BeginVerify"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.verifyCreateRequest(ctx, reportName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// verifyCreateRequest creates the Verify request.
func (client *ReportClient) verifyCreateRequest(ctx context.Context, reportName string, options *ReportClientBeginVerifyOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/verify"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
