//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

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
	"strconv"
	"strings"
)

// DataLakeStoreAccountsClient contains the methods for the DataLakeStoreAccounts group.
// Don't use this type directly, use NewDataLakeStoreAccountsClient() instead.
type DataLakeStoreAccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataLakeStoreAccountsClient creates a new instance of DataLakeStoreAccountsClient with the specified values.
// subscriptionID - Get subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataLakeStoreAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataLakeStoreAccountsClient, error) {
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
	client := &DataLakeStoreAccountsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Add - Updates the specified Data Lake Analytics account to include the additional Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// dataLakeStoreAccountName - The name of the Data Lake Store account to add.
// options - DataLakeStoreAccountsClientAddOptions contains the optional parameters for the DataLakeStoreAccountsClient.Add
// method.
func (client *DataLakeStoreAccountsClient) Add(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, options *DataLakeStoreAccountsClientAddOptions) (DataLakeStoreAccountsClientAddResponse, error) {
	req, err := client.addCreateRequest(ctx, resourceGroupName, accountName, dataLakeStoreAccountName, options)
	if err != nil {
		return DataLakeStoreAccountsClientAddResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataLakeStoreAccountsClientAddResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataLakeStoreAccountsClientAddResponse{}, runtime.NewResponseError(resp)
	}
	return DataLakeStoreAccountsClientAddResponse{}, nil
}

// addCreateRequest creates the Add request.
func (client *DataLakeStoreAccountsClient) addCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, options *DataLakeStoreAccountsClientAddOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if dataLakeStoreAccountName == "" {
		return nil, errors.New("parameter dataLakeStoreAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataLakeStoreAccountName}", url.PathEscape(dataLakeStoreAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// Delete - Updates the Data Lake Analytics account specified to remove the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// dataLakeStoreAccountName - The name of the Data Lake Store account to remove
// options - DataLakeStoreAccountsClientDeleteOptions contains the optional parameters for the DataLakeStoreAccountsClient.Delete
// method.
func (client *DataLakeStoreAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, options *DataLakeStoreAccountsClientDeleteOptions) (DataLakeStoreAccountsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, dataLakeStoreAccountName, options)
	if err != nil {
		return DataLakeStoreAccountsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataLakeStoreAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataLakeStoreAccountsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataLakeStoreAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataLakeStoreAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, options *DataLakeStoreAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if dataLakeStoreAccountName == "" {
		return nil, errors.New("parameter dataLakeStoreAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataLakeStoreAccountName}", url.PathEscape(dataLakeStoreAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified Data Lake Store account details in the specified Data Lake Analytics account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// dataLakeStoreAccountName - The name of the Data Lake Store account to retrieve
// options - DataLakeStoreAccountsClientGetOptions contains the optional parameters for the DataLakeStoreAccountsClient.Get
// method.
func (client *DataLakeStoreAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, options *DataLakeStoreAccountsClientGetOptions) (DataLakeStoreAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, dataLakeStoreAccountName, options)
	if err != nil {
		return DataLakeStoreAccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataLakeStoreAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataLakeStoreAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataLakeStoreAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, options *DataLakeStoreAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if dataLakeStoreAccountName == "" {
		return nil, errors.New("parameter dataLakeStoreAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataLakeStoreAccountName}", url.PathEscape(dataLakeStoreAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataLakeStoreAccountsClient) getHandleResponse(resp *http.Response) (DataLakeStoreAccountsClientGetResponse, error) {
	result := DataLakeStoreAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataLakeStoreAccountInformation); err != nil {
		return DataLakeStoreAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Gets the first page of Data Lake Store accounts linked to the specified Data Lake Analytics account.
// The response includes a link to the next page, if any.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// options - DataLakeStoreAccountsClientListByAccountOptions contains the optional parameters for the DataLakeStoreAccountsClient.ListByAccount
// method.
func (client *DataLakeStoreAccountsClient) NewListByAccountPager(resourceGroupName string, accountName string, options *DataLakeStoreAccountsClientListByAccountOptions) *runtime.Pager[DataLakeStoreAccountsClientListByAccountResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataLakeStoreAccountsClientListByAccountResponse]{
		More: func(page DataLakeStoreAccountsClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataLakeStoreAccountsClientListByAccountResponse) (DataLakeStoreAccountsClientListByAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataLakeStoreAccountsClientListByAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataLakeStoreAccountsClientListByAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataLakeStoreAccountsClientListByAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAccountHandleResponse(resp)
		},
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *DataLakeStoreAccountsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *DataLakeStoreAccountsClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *DataLakeStoreAccountsClient) listByAccountHandleResponse(resp *http.Response) (DataLakeStoreAccountsClientListByAccountResponse, error) {
	result := DataLakeStoreAccountsClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataLakeStoreAccountInformationListResult); err != nil {
		return DataLakeStoreAccountsClientListByAccountResponse{}, err
	}
	return result, nil
}
