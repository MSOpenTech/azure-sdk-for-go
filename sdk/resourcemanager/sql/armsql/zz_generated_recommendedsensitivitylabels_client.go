//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RecommendedSensitivityLabelsClient contains the methods for the RecommendedSensitivityLabels group.
// Don't use this type directly, use NewRecommendedSensitivityLabelsClient() instead.
type RecommendedSensitivityLabelsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRecommendedSensitivityLabelsClient creates a new instance of RecommendedSensitivityLabelsClient with the specified values.
func NewRecommendedSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RecommendedSensitivityLabelsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RecommendedSensitivityLabelsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Update - Update recommended sensitivity labels states of a given database using an operations batch.
// If the operation fails it returns a generic error.
func (client *RecommendedSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *RecommendedSensitivityLabelsUpdateOptions) (RecommendedSensitivityLabelsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, parameters, options)
	if err != nil {
		return RecommendedSensitivityLabelsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendedSensitivityLabelsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendedSensitivityLabelsUpdateResponse{}, client.updateHandleError(resp)
	}
	return RecommendedSensitivityLabelsUpdateResponse{RawResponse: resp}, nil
}

// updateCreateRequest creates the Update request.
func (client *RecommendedSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *RecommendedSensitivityLabelsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/recommendedSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *RecommendedSensitivityLabelsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
