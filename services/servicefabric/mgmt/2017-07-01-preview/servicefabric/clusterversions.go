package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ClusterVersionsClient is the service Fabric Management Client
type ClusterVersionsClient struct {
	BaseClient
}

// NewClusterVersionsClient creates an instance of the ClusterVersionsClient client.
func NewClusterVersionsClient() ClusterVersionsClient {
	return NewClusterVersionsClientWithBaseURI(DefaultBaseURI)
}

// NewClusterVersionsClientWithBaseURI creates an instance of the ClusterVersionsClient client.
func NewClusterVersionsClientWithBaseURI(baseURI string) ClusterVersionsClient {
	return ClusterVersionsClient{NewWithBaseURI(baseURI)}
}

// Get get cluster code versions by location
//
// location is the location for the cluster code versions, this is different from cluster location subscriptionID
// is the customer subscription identifier clusterVersion is the cluster code version
func (client ClusterVersionsClient) Get(ctx context.Context, location string, subscriptionID string, clusterVersion string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.GetPreparer(ctx, location, subscriptionID, clusterVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClusterVersionsClient) GetPreparer(ctx context.Context, location string, subscriptionID string, clusterVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterVersion": autorest.Encode("path", clusterVersion),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) GetResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByEnvironment get cluster code versions by environment
//
// location is the location for the cluster code versions, this is different from cluster location environment is
// cluster operating system, the default means all subscriptionID is the customer subscription identifier
// clusterVersion is the cluster code version
func (client ClusterVersionsClient) GetByEnvironment(ctx context.Context, location string, environment string, subscriptionID string, clusterVersion string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.GetByEnvironmentPreparer(ctx, location, environment, subscriptionID, clusterVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "GetByEnvironment", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "GetByEnvironment", resp, "Failure sending request")
		return
	}

	result, err = client.GetByEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "GetByEnvironment", resp, "Failure responding to request")
	}

	return
}

// GetByEnvironmentPreparer prepares the GetByEnvironment request.
func (client ClusterVersionsClient) GetByEnvironmentPreparer(ctx context.Context, location string, environment string, subscriptionID string, clusterVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterVersion": autorest.Encode("path", clusterVersion),
		"environment":    autorest.Encode("path", environment),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByEnvironmentSender sends the GetByEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) GetByEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByEnvironmentResponder handles the response to the GetByEnvironment request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) GetByEnvironmentResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list cluster code versions by location
//
// location is the location for the cluster code versions, this is different from cluster location subscriptionID
// is the customer subscription identifier
func (client ClusterVersionsClient) List(ctx context.Context, location string, subscriptionID string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.ListPreparer(ctx, location, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ClusterVersionsClient) ListPreparer(ctx context.Context, location string, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEnvironment list cluster code versions by environment
//
// location is the location for the cluster code versions, this is different from cluster location environment is
// cluster operating system, the default means all subscriptionID is the customer subscription identifier
func (client ClusterVersionsClient) ListByEnvironment(ctx context.Context, location string, environment string, subscriptionID string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.ListByEnvironmentPreparer(ctx, location, environment, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure responding to request")
	}

	return
}

// ListByEnvironmentPreparer prepares the ListByEnvironment request.
func (client ClusterVersionsClient) ListByEnvironmentPreparer(ctx context.Context, location string, environment string, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environment":    autorest.Encode("path", environment),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEnvironmentSender sends the ListByEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) ListByEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByEnvironmentResponder handles the response to the ListByEnvironment request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListByEnvironmentResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
