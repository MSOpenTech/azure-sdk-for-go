//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// GenerateDetailedCostReportClient contains the methods for the GenerateDetailedCostReport group.
// Don't use this type directly, use NewGenerateDetailedCostReportClient() instead.
type GenerateDetailedCostReportClient struct {
	internal *arm.Client
}

// NewGenerateDetailedCostReportClient creates a new instance of GenerateDetailedCostReportClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGenerateDetailedCostReportClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateDetailedCostReportClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GenerateDetailedCostReportClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOperation - Generates the detailed cost report for provided date range, billing period(only enterprise customers)
// or Invoice ID asynchronously at a certain scope. Call returns a 202 with header
// Azure-Consumption-AsyncOperation providing a link to the operation created. A call on the operation will provide the status
// and if the operation is completed the blob file where generated detailed
// cost report is being stored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - scope - The ARM Resource ID for subscription, resource group, billing account, or other billing scopes. For details, see
//     https://aka.ms/costmgmt/scopes.
//   - parameters - Parameters supplied to the Create detailed cost report operation.
//   - options - GenerateDetailedCostReportClientBeginCreateOperationOptions contains the optional parameters for the GenerateDetailedCostReportClient.BeginCreateOperation
//     method.
func (client *GenerateDetailedCostReportClient) BeginCreateOperation(ctx context.Context, scope string, parameters GenerateDetailedCostReportDefinition, options *GenerateDetailedCostReportClientBeginCreateOperationOptions) (*runtime.Poller[GenerateDetailedCostReportClientCreateOperationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOperation(ctx, scope, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GenerateDetailedCostReportClientCreateOperationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GenerateDetailedCostReportClientCreateOperationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOperation - Generates the detailed cost report for provided date range, billing period(only enterprise customers)
// or Invoice ID asynchronously at a certain scope. Call returns a 202 with header
// Azure-Consumption-AsyncOperation providing a link to the operation created. A call on the operation will provide the status
// and if the operation is completed the blob file where generated detailed
// cost report is being stored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
func (client *GenerateDetailedCostReportClient) createOperation(ctx context.Context, scope string, parameters GenerateDetailedCostReportDefinition, options *GenerateDetailedCostReportClientBeginCreateOperationOptions) (*http.Response, error) {
	var err error
	const operationName = "GenerateDetailedCostReportClient.BeginCreateOperation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOperationCreateRequest(ctx, scope, parameters, options)
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

// createOperationCreateRequest creates the CreateOperation request.
func (client *GenerateDetailedCostReportClient) createOperationCreateRequest(ctx context.Context, scope string, parameters GenerateDetailedCostReportDefinition, options *GenerateDetailedCostReportClientBeginCreateOperationOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/generateDetailedCostReport"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
