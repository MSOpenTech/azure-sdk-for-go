//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple8000series

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// CloudAppliancesClient contains the methods for the CloudAppliances group.
// Don't use this type directly, use NewCloudAppliancesClient() instead.
type CloudAppliancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCloudAppliancesClient creates a new instance of CloudAppliancesClient with the specified values.
// subscriptionID - The subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCloudAppliancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudAppliancesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CloudAppliancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListSupportedConfigurationsPager - Lists supported cloud appliance models and supported configurations.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-06-01
// resourceGroupName - The resource group name
// managerName - The manager name
// options - CloudAppliancesClientListSupportedConfigurationsOptions contains the optional parameters for the CloudAppliancesClient.ListSupportedConfigurations
// method.
func (client *CloudAppliancesClient) NewListSupportedConfigurationsPager(resourceGroupName string, managerName string, options *CloudAppliancesClientListSupportedConfigurationsOptions) *runtime.Pager[CloudAppliancesClientListSupportedConfigurationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudAppliancesClientListSupportedConfigurationsResponse]{
		More: func(page CloudAppliancesClientListSupportedConfigurationsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CloudAppliancesClientListSupportedConfigurationsResponse) (CloudAppliancesClientListSupportedConfigurationsResponse, error) {
			req, err := client.listSupportedConfigurationsCreateRequest(ctx, resourceGroupName, managerName, options)
			if err != nil {
				return CloudAppliancesClientListSupportedConfigurationsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CloudAppliancesClientListSupportedConfigurationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CloudAppliancesClientListSupportedConfigurationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSupportedConfigurationsHandleResponse(resp)
		},
	})
}

// listSupportedConfigurationsCreateRequest creates the ListSupportedConfigurations request.
func (client *CloudAppliancesClient) listSupportedConfigurationsCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *CloudAppliancesClientListSupportedConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/cloudApplianceConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSupportedConfigurationsHandleResponse handles the ListSupportedConfigurations response.
func (client *CloudAppliancesClient) listSupportedConfigurationsHandleResponse(resp *http.Response) (CloudAppliancesClientListSupportedConfigurationsResponse, error) {
	result := CloudAppliancesClientListSupportedConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudApplianceConfigurationList); err != nil {
		return CloudAppliancesClientListSupportedConfigurationsResponse{}, err
	}
	return result, nil
}

// BeginProvision - Provisions cloud appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-06-01
// resourceGroupName - The resource group name
// managerName - The manager name
// parameters - The cloud appliance
// options - CloudAppliancesClientBeginProvisionOptions contains the optional parameters for the CloudAppliancesClient.BeginProvision
// method.
func (client *CloudAppliancesClient) BeginProvision(ctx context.Context, resourceGroupName string, managerName string, parameters CloudAppliance, options *CloudAppliancesClientBeginProvisionOptions) (*runtime.Poller[CloudAppliancesClientProvisionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.provision(ctx, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CloudAppliancesClientProvisionResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CloudAppliancesClientProvisionResponse](options.ResumeToken, client.pl, nil)
	}
}

// Provision - Provisions cloud appliance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-06-01
func (client *CloudAppliancesClient) provision(ctx context.Context, resourceGroupName string, managerName string, parameters CloudAppliance, options *CloudAppliancesClientBeginProvisionOptions) (*http.Response, error) {
	req, err := client.provisionCreateRequest(ctx, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// provisionCreateRequest creates the Provision request.
func (client *CloudAppliancesClient) provisionCreateRequest(ctx context.Context, resourceGroupName string, managerName string, parameters CloudAppliance, options *CloudAppliancesClientBeginProvisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/provisionCloudAppliance"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}
