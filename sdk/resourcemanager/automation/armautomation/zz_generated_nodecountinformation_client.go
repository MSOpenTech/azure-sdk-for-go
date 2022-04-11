//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// NodeCountInformationClient contains the methods for the NodeCountInformation group.
// Don't use this type directly, use NewNodeCountInformationClient() instead.
type NodeCountInformationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNodeCountInformationClient creates a new instance of NodeCountInformationClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNodeCountInformationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NodeCountInformationClient, error) {
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
	client := &NodeCountInformationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Retrieve counts for Dsc Nodes.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// countType - The type of counts to retrieve
// options - NodeCountInformationClientGetOptions contains the optional parameters for the NodeCountInformationClient.Get
// method.
func (client *NodeCountInformationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, countType CountType, options *NodeCountInformationClientGetOptions) (NodeCountInformationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, countType, options)
	if err != nil {
		return NodeCountInformationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NodeCountInformationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NodeCountInformationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NodeCountInformationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, countType CountType, options *NodeCountInformationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodecounts/{countType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if countType == "" {
		return nil, errors.New("parameter countType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{countType}", url.PathEscape(string(countType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NodeCountInformationClient) getHandleResponse(resp *http.Response) (NodeCountInformationClientGetResponse, error) {
	result := NodeCountInformationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NodeCounts); err != nil {
		return NodeCountInformationClientGetResponse{}, err
	}
	return result, nil
}
