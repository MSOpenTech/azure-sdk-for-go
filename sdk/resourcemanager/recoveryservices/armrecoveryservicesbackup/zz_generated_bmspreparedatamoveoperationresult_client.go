//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// BMSPrepareDataMoveOperationResultClient contains the methods for the BMSPrepareDataMoveOperationResult group.
// Don't use this type directly, use NewBMSPrepareDataMoveOperationResultClient() instead.
type BMSPrepareDataMoveOperationResultClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewBMSPrepareDataMoveOperationResultClient creates a new instance of BMSPrepareDataMoveOperationResultClient with the specified values.
func NewBMSPrepareDataMoveOperationResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BMSPrepareDataMoveOperationResultClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BMSPrepareDataMoveOperationResultClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Fetches Operation Result for Prepare Data Move
// If the operation fails it returns the *NewErrorResponse error type.
func (client *BMSPrepareDataMoveOperationResultClient) Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *BMSPrepareDataMoveOperationResultGetOptions) (BMSPrepareDataMoveOperationResultGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, operationID, options)
	if err != nil {
		return BMSPrepareDataMoveOperationResultGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BMSPrepareDataMoveOperationResultGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return BMSPrepareDataMoveOperationResultGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BMSPrepareDataMoveOperationResultClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *BMSPrepareDataMoveOperationResultGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig/operationResults/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BMSPrepareDataMoveOperationResultClient) getHandleResponse(resp *http.Response) (BMSPrepareDataMoveOperationResultGetResponse, error) {
	result := BMSPrepareDataMoveOperationResultGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return BMSPrepareDataMoveOperationResultGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BMSPrepareDataMoveOperationResultClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := NewErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
