//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkloadClassifiersClient contains the methods for the WorkloadClassifiers group.
// Don't use this type directly, use NewWorkloadClassifiersClient() instead.
type WorkloadClassifiersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkloadClassifiersClient creates a new instance of WorkloadClassifiersClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkloadClassifiersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkloadClassifiersClient, error) {
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
	client := &WorkloadClassifiersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a workload classifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// workloadGroupName - The name of the workload group from which to receive the classifier from.
// workloadClassifierName - The name of the workload classifier to create/update.
// parameters - The properties of the workload classifier.
// options - WorkloadClassifiersClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkloadClassifiersClient.BeginCreateOrUpdate
// method.
func (client *WorkloadClassifiersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *WorkloadClassifiersClientBeginCreateOrUpdateOptions) (*armruntime.Poller[WorkloadClassifiersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[WorkloadClassifiersClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[WorkloadClassifiersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a workload classifier.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkloadClassifiersClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *WorkloadClassifiersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkloadClassifiersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier, options *WorkloadClassifiersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a workload classifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// workloadGroupName - The name of the workload group from which to receive the classifier from.
// workloadClassifierName - The name of the workload classifier to delete.
// options - WorkloadClassifiersClientBeginDeleteOptions contains the optional parameters for the WorkloadClassifiersClient.BeginDelete
// method.
func (client *WorkloadClassifiersClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersClientBeginDeleteOptions) (*armruntime.Poller[WorkloadClassifiersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[WorkloadClassifiersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[WorkloadClassifiersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a workload classifier.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkloadClassifiersClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, options)
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
func (client *WorkloadClassifiersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a workload classifier
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// workloadGroupName - The name of the workload group from which to receive the classifier from.
// workloadClassifierName - The name of the workload classifier.
// options - WorkloadClassifiersClientGetOptions contains the optional parameters for the WorkloadClassifiersClient.Get method.
func (client *WorkloadClassifiersClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersClientGetOptions) (WorkloadClassifiersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, options)
	if err != nil {
		return WorkloadClassifiersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkloadClassifiersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkloadClassifiersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkloadClassifiersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, options *WorkloadClassifiersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if workloadClassifierName == "" {
		return nil, errors.New("parameter workloadClassifierName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadClassifierName}", url.PathEscape(workloadClassifierName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkloadClassifiersClient) getHandleResponse(resp *http.Response) (WorkloadClassifiersClientGetResponse, error) {
	result := WorkloadClassifiersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadClassifier); err != nil {
		return WorkloadClassifiersClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkloadGroup - Gets the list of workload classifiers for a workload group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// workloadGroupName - The name of the workload group from which to receive the classifiers from.
// options - WorkloadClassifiersClientListByWorkloadGroupOptions contains the optional parameters for the WorkloadClassifiersClient.ListByWorkloadGroup
// method.
func (client *WorkloadClassifiersClient) ListByWorkloadGroup(resourceGroupName string, serverName string, databaseName string, workloadGroupName string, options *WorkloadClassifiersClientListByWorkloadGroupOptions) *runtime.Pager[WorkloadClassifiersClientListByWorkloadGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[WorkloadClassifiersClientListByWorkloadGroupResponse]{
		More: func(page WorkloadClassifiersClientListByWorkloadGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkloadClassifiersClientListByWorkloadGroupResponse) (WorkloadClassifiersClientListByWorkloadGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkloadGroupCreateRequest(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkloadClassifiersClientListByWorkloadGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkloadClassifiersClientListByWorkloadGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkloadClassifiersClientListByWorkloadGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkloadGroupHandleResponse(resp)
		},
	})
}

// listByWorkloadGroupCreateRequest creates the ListByWorkloadGroup request.
func (client *WorkloadClassifiersClient) listByWorkloadGroupCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, options *WorkloadClassifiersClientListByWorkloadGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers"
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
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkloadGroupHandleResponse handles the ListByWorkloadGroup response.
func (client *WorkloadClassifiersClient) listByWorkloadGroupHandleResponse(resp *http.Response) (WorkloadClassifiersClientListByWorkloadGroupResponse, error) {
	result := WorkloadClassifiersClientListByWorkloadGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadClassifierListResult); err != nil {
		return WorkloadClassifiersClientListByWorkloadGroupResponse{}, err
	}
	return result, nil
}
