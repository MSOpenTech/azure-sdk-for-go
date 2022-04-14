//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationServiceEnvironmentManagedAPIOperationsClient contains the methods for the IntegrationServiceEnvironmentManagedAPIOperations group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentManagedAPIOperationsClient() instead.
type IntegrationServiceEnvironmentManagedAPIOperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationServiceEnvironmentManagedAPIOperationsClient creates a new instance of IntegrationServiceEnvironmentManagedAPIOperationsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationServiceEnvironmentManagedAPIOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationServiceEnvironmentManagedAPIOperationsClient, error) {
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
	client := &IntegrationServiceEnvironmentManagedAPIOperationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Gets the managed Api operations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - The resource group.
// integrationServiceEnvironmentName - The integration service environment name.
// apiName - The api name.
// options - IntegrationServiceEnvironmentManagedAPIOperationsClientListOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedAPIOperationsClient.List
// method.
func (client *IntegrationServiceEnvironmentManagedAPIOperationsClient) List(resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedAPIOperationsClientListOptions) *runtime.Pager[IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse]{
		More: func(page IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse) (IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentManagedAPIOperationsClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedAPIOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}/apiOperations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentManagedAPIOperationsClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse, error) {
	result := IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIOperationListResult); err != nil {
		return IntegrationServiceEnvironmentManagedAPIOperationsClientListResponse{}, err
	}
	return result, nil
}
