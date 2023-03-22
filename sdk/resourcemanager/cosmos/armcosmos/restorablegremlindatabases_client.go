//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

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

// RestorableGremlinDatabasesClient contains the methods for the RestorableGremlinDatabases group.
// Don't use this type directly, use NewRestorableGremlinDatabasesClient() instead.
type RestorableGremlinDatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableGremlinDatabasesClient creates a new instance of RestorableGremlinDatabasesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableGremlinDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableGremlinDatabasesClient, error) {
	cl, err := arm.NewClient(moduleName+".RestorableGremlinDatabasesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableGremlinDatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Show the event feed of all mutations done on all the Azure Cosmos DB Gremlin databases under the restorable
// account. This helps in scenario where database was accidentally deleted to get the deletion
// time. This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
//
// Generated from API version 2022-11-15-preview
//   - location - Cosmos DB region, with spaces between words and each word capitalized.
//   - instanceID - The instanceId GUID of a restorable database account.
//   - options - RestorableGremlinDatabasesClientListOptions contains the optional parameters for the RestorableGremlinDatabasesClient.NewListPager
//     method.
func (client *RestorableGremlinDatabasesClient) NewListPager(location string, instanceID string, options *RestorableGremlinDatabasesClientListOptions) *runtime.Pager[RestorableGremlinDatabasesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableGremlinDatabasesClientListResponse]{
		More: func(page RestorableGremlinDatabasesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableGremlinDatabasesClientListResponse) (RestorableGremlinDatabasesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableGremlinDatabasesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorableGremlinDatabasesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableGremlinDatabasesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RestorableGremlinDatabasesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableGremlinDatabasesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableGremlinDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableGremlinDatabasesClient) listHandleResponse(resp *http.Response) (RestorableGremlinDatabasesClientListResponse, error) {
	result := RestorableGremlinDatabasesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableGremlinDatabasesListResult); err != nil {
		return RestorableGremlinDatabasesClientListResponse{}, err
	}
	return result, nil
}
