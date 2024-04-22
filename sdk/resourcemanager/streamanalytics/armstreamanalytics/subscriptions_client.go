//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionsClient contains the methods for the Subscriptions group.
// Don't use this type directly, use NewSubscriptionsClient() instead.
type SubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubscriptionsClient creates a new instance of SubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CompileQuery - Compile the Stream Analytics query.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
//   - location - The region to which the request is sent. You can find out which regions Azure Stream Analytics is supported
//     in here: https://azure.microsoft.com/en-us/regions/
//   - compileQuery - The query compilation object which defines the input, output, and transformation for the query compilation.
//   - options - SubscriptionsClientCompileQueryOptions contains the optional parameters for the SubscriptionsClient.CompileQuery
//     method.
func (client *SubscriptionsClient) CompileQuery(ctx context.Context, location string, compileQuery CompileQuery, options *SubscriptionsClientCompileQueryOptions) (SubscriptionsClientCompileQueryResponse, error) {
	var err error
	const operationName = "SubscriptionsClient.CompileQuery"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.compileQueryCreateRequest(ctx, location, compileQuery, options)
	if err != nil {
		return SubscriptionsClientCompileQueryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionsClientCompileQueryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionsClientCompileQueryResponse{}, err
	}
	resp, err := client.compileQueryHandleResponse(httpResp)
	return resp, err
}

// compileQueryCreateRequest creates the CompileQuery request.
func (client *SubscriptionsClient) compileQueryCreateRequest(ctx context.Context, location string, compileQuery CompileQuery, options *SubscriptionsClientCompileQueryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/compileQuery"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, compileQuery); err != nil {
		return nil, err
	}
	return req, nil
}

// compileQueryHandleResponse handles the CompileQuery response.
func (client *SubscriptionsClient) compileQueryHandleResponse(resp *http.Response) (SubscriptionsClientCompileQueryResponse, error) {
	result := SubscriptionsClientCompileQueryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryCompilationResult); err != nil {
		return SubscriptionsClientCompileQueryResponse{}, err
	}
	return result, nil
}

// ListQuotas - Retrieves the subscription's current quota information in a particular region.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
//   - location - The region to which the request is sent. You can find out which regions Azure Stream Analytics is supported
//     in here: https://azure.microsoft.com/en-us/regions/
//   - options - SubscriptionsClientListQuotasOptions contains the optional parameters for the SubscriptionsClient.ListQuotas
//     method.
func (client *SubscriptionsClient) ListQuotas(ctx context.Context, location string, options *SubscriptionsClientListQuotasOptions) (SubscriptionsClientListQuotasResponse, error) {
	var err error
	const operationName = "SubscriptionsClient.ListQuotas"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQuotasCreateRequest(ctx, location, options)
	if err != nil {
		return SubscriptionsClientListQuotasResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionsClientListQuotasResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionsClientListQuotasResponse{}, err
	}
	resp, err := client.listQuotasHandleResponse(httpResp)
	return resp, err
}

// listQuotasCreateRequest creates the ListQuotas request.
func (client *SubscriptionsClient) listQuotasCreateRequest(ctx context.Context, location string, options *SubscriptionsClientListQuotasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/quotas"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQuotasHandleResponse handles the ListQuotas response.
func (client *SubscriptionsClient) listQuotasHandleResponse(resp *http.Response) (SubscriptionsClientListQuotasResponse, error) {
	result := SubscriptionsClientListQuotasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionQuotasListResult); err != nil {
		return SubscriptionsClientListQuotasResponse{}, err
	}
	return result, nil
}

// BeginSampleInput - Sample the Stream Analytics input data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
//   - location - The region to which the request is sent. You can find out which regions Azure Stream Analytics is supported
//     in here: https://azure.microsoft.com/en-us/regions/
//   - sampleInput - Defines the necessary parameters for sampling the Stream Analytics input data.
//   - options - SubscriptionsClientBeginSampleInputOptions contains the optional parameters for the SubscriptionsClient.BeginSampleInput
//     method.
func (client *SubscriptionsClient) BeginSampleInput(ctx context.Context, location string, sampleInput SampleInput, options *SubscriptionsClientBeginSampleInputOptions) (*runtime.Poller[SubscriptionsClientSampleInputResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.sampleInput(ctx, location, sampleInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubscriptionsClientSampleInputResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubscriptionsClientSampleInputResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// SampleInput - Sample the Stream Analytics input data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
func (client *SubscriptionsClient) sampleInput(ctx context.Context, location string, sampleInput SampleInput, options *SubscriptionsClientBeginSampleInputOptions) (*http.Response, error) {
	var err error
	const operationName = "SubscriptionsClient.BeginSampleInput"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sampleInputCreateRequest(ctx, location, sampleInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// sampleInputCreateRequest creates the SampleInput request.
func (client *SubscriptionsClient) sampleInputCreateRequest(ctx context.Context, location string, sampleInput SampleInput, options *SubscriptionsClientBeginSampleInputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/sampleInput"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sampleInput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginTestInput - Test the Stream Analytics input.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
//   - location - The region to which the request is sent. You can find out which regions Azure Stream Analytics is supported
//     in here: https://azure.microsoft.com/en-us/regions/
//   - testInput - Defines the necessary parameters for testing the Stream Analytics input.
//   - options - SubscriptionsClientBeginTestInputOptions contains the optional parameters for the SubscriptionsClient.BeginTestInput
//     method.
func (client *SubscriptionsClient) BeginTestInput(ctx context.Context, location string, testInput TestInput, options *SubscriptionsClientBeginTestInputOptions) (*runtime.Poller[SubscriptionsClientTestInputResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.testInput(ctx, location, testInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubscriptionsClientTestInputResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubscriptionsClientTestInputResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TestInput - Test the Stream Analytics input.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
func (client *SubscriptionsClient) testInput(ctx context.Context, location string, testInput TestInput, options *SubscriptionsClientBeginTestInputOptions) (*http.Response, error) {
	var err error
	const operationName = "SubscriptionsClient.BeginTestInput"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.testInputCreateRequest(ctx, location, testInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// testInputCreateRequest creates the TestInput request.
func (client *SubscriptionsClient) testInputCreateRequest(ctx context.Context, location string, testInput TestInput, options *SubscriptionsClientBeginTestInputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/testInput"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, testInput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginTestOutput - Test the Stream Analytics output.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
//   - location - The region to which the request is sent. You can find out which regions Azure Stream Analytics is supported
//     in here: https://azure.microsoft.com/en-us/regions/
//   - testOutput - Defines the necessary parameters for testing the Stream Analytics output.
//   - options - SubscriptionsClientBeginTestOutputOptions contains the optional parameters for the SubscriptionsClient.BeginTestOutput
//     method.
func (client *SubscriptionsClient) BeginTestOutput(ctx context.Context, location string, testOutput TestOutput, options *SubscriptionsClientBeginTestOutputOptions) (*runtime.Poller[SubscriptionsClientTestOutputResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.testOutput(ctx, location, testOutput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubscriptionsClientTestOutputResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubscriptionsClientTestOutputResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TestOutput - Test the Stream Analytics output.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
func (client *SubscriptionsClient) testOutput(ctx context.Context, location string, testOutput TestOutput, options *SubscriptionsClientBeginTestOutputOptions) (*http.Response, error) {
	var err error
	const operationName = "SubscriptionsClient.BeginTestOutput"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.testOutputCreateRequest(ctx, location, testOutput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// testOutputCreateRequest creates the TestOutput request.
func (client *SubscriptionsClient) testOutputCreateRequest(ctx context.Context, location string, testOutput TestOutput, options *SubscriptionsClientBeginTestOutputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/testOutput"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, testOutput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginTestQuery - Test the Stream Analytics query on a sample input.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
//   - location - The region to which the request is sent. You can find out which regions Azure Stream Analytics is supported
//     in here: https://azure.microsoft.com/en-us/regions/
//   - testQuery - The query testing object that defines the input, output, and transformation for the query testing.
//   - options - SubscriptionsClientBeginTestQueryOptions contains the optional parameters for the SubscriptionsClient.BeginTestQuery
//     method.
func (client *SubscriptionsClient) BeginTestQuery(ctx context.Context, location string, testQuery TestQuery, options *SubscriptionsClientBeginTestQueryOptions) (*runtime.Poller[SubscriptionsClientTestQueryResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.testQuery(ctx, location, testQuery, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubscriptionsClientTestQueryResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubscriptionsClientTestQueryResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TestQuery - Test the Stream Analytics query on a sample input.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01-preview
func (client *SubscriptionsClient) testQuery(ctx context.Context, location string, testQuery TestQuery, options *SubscriptionsClientBeginTestQueryOptions) (*http.Response, error) {
	var err error
	const operationName = "SubscriptionsClient.BeginTestQuery"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.testQueryCreateRequest(ctx, location, testQuery, options)
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

// testQueryCreateRequest creates the TestQuery request.
func (client *SubscriptionsClient) testQueryCreateRequest(ctx context.Context, location string, testQuery TestQuery, options *SubscriptionsClientBeginTestQueryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/testQuery"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, testQuery); err != nil {
		return nil, err
	}
	return req, nil
}
