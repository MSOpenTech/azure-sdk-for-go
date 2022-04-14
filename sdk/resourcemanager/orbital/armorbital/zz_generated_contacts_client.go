//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ContactsClient contains the methods for the Contacts group.
// Don't use this type directly, use NewContactsClient() instead.
type ContactsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContactsClient creates a new instance of ContactsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContactsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContactsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ContactsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a contact.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// spacecraftName - Spacecraft ID
// contactName - Contact Name
// parameters - The parameters to provide for the created contact.
// options - ContactsClientBeginCreateOptions contains the optional parameters for the ContactsClient.BeginCreate method.
func (client *ContactsClient) BeginCreate(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, parameters Contact, options *ContactsClientBeginCreateOptions) (*armruntime.Poller[ContactsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, spacecraftName, contactName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ContactsClientCreateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ContactsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a contact.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ContactsClient) create(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, parameters Contact, options *ContactsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, spacecraftName, contactName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ContactsClient) createCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, parameters Contact, options *ContactsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts/{contactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	if contactName == "" {
		return nil, errors.New("parameter contactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactName}", url.PathEscape(contactName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a specified contact
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// spacecraftName - Spacecraft ID
// contactName - Contact Name
// options - ContactsClientBeginDeleteOptions contains the optional parameters for the ContactsClient.BeginDelete method.
func (client *ContactsClient) BeginDelete(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsClientBeginDeleteOptions) (*armruntime.Poller[ContactsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, spacecraftName, contactName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ContactsClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ContactsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a specified contact
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ContactsClient) deleteOperation(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, spacecraftName, contactName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContactsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts/{contactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	if contactName == "" {
		return nil, errors.New("parameter contactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactName}", url.PathEscape(contactName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified contact in a specified resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// spacecraftName - Spacecraft ID
// contactName - Contact Name
// options - ContactsClientGetOptions contains the optional parameters for the ContactsClient.Get method.
func (client *ContactsClient) Get(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsClientGetOptions) (ContactsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, spacecraftName, contactName, options)
	if err != nil {
		return ContactsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContactsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContactsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContactsClient) getCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts/{contactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	if contactName == "" {
		return nil, errors.New("parameter contactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactName}", url.PathEscape(contactName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContactsClient) getHandleResponse(resp *http.Response) (ContactsClientGetResponse, error) {
	result := ContactsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Contact); err != nil {
		return ContactsClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns list of contacts by spacecraftName
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// spacecraftName - Spacecraft ID
// options - ContactsClientListOptions contains the optional parameters for the ContactsClient.List method.
func (client *ContactsClient) List(resourceGroupName string, spacecraftName string, options *ContactsClientListOptions) *runtime.Pager[ContactsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ContactsClientListResponse]{
		More: func(page ContactsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ContactsClientListResponse) (ContactsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, spacecraftName, options)
			if err != nil {
				return ContactsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContactsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContactsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ContactsClient) listCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, options *ContactsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContactsClient) listHandleResponse(resp *http.Response) (ContactsClientListResponse, error) {
	result := ContactsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContactListResult); err != nil {
		return ContactsClientListResponse{}, err
	}
	return result, nil
}
