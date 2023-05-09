//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// NamespacesClient contains the methods for the Namespaces group.
// Don't use this type directly, use NewNamespacesClient() instead.
type NamespacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNamespacesClient creates a new instance of NamespacesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNamespacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NamespacesClient, error) {
	cl, err := arm.NewClient(moduleName+".NamespacesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NamespacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks if resource name is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - checkNameAvailabilityRequest - The check name availability request.
//   - options - NamespacesClientCheckNameAvailabilityOptions contains the optional parameters for the NamespacesClient.CheckNameAvailability
//     method.
func (client *NamespacesClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, environmentName string, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *NamespacesClientCheckNameAvailabilityOptions) (NamespacesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, environmentName, checkNameAvailabilityRequest, options)
	if err != nil {
		return NamespacesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NamespacesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NamespacesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *NamespacesClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *NamespacesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *NamespacesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (NamespacesClientCheckNameAvailabilityResponse, error) {
	result := NamespacesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return NamespacesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}
