package network

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateLinkServicesClient is the network Client
type PrivateLinkServicesClient struct {
	BaseClient
}

// NewPrivateLinkServicesClient creates an instance of the PrivateLinkServicesClient client.
func NewPrivateLinkServicesClient(subscriptionID string) PrivateLinkServicesClient {
	return NewPrivateLinkServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkServicesClientWithBaseURI creates an instance of the PrivateLinkServicesClient client.
func NewPrivateLinkServicesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkServicesClient {
	return PrivateLinkServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an private link service in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the private link service.
// parameters - parameters supplied to the create or update private link service operation
func (client PrivateLinkServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters PrivateLinkService) (result PrivateLinkServicesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateLinkServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters PrivateLinkService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) CreateOrUpdateSender(req *http.Request) (future PrivateLinkServicesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateLinkService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified private link service.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the private link service.
func (client PrivateLinkServicesClient) Delete(ctx context.Context, resourceGroupName string, serviceName string) (result PrivateLinkServicesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateLinkServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) DeleteSender(req *http.Request) (future PrivateLinkServicesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeletePrivateEndpointConnection delete private end point connection for a private link service in a subscription.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the private link service.
// peConnectionName - the name of the private end point connection.
func (client PrivateLinkServicesClient) DeletePrivateEndpointConnection(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string) (result PrivateLinkServicesDeletePrivateEndpointConnectionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.DeletePrivateEndpointConnection")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePrivateEndpointConnectionPreparer(ctx, resourceGroupName, serviceName, peConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "DeletePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	result, err = client.DeletePrivateEndpointConnectionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "DeletePrivateEndpointConnection", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePrivateEndpointConnectionPreparer prepares the DeletePrivateEndpointConnection request.
func (client PrivateLinkServicesClient) DeletePrivateEndpointConnectionPreparer(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peConnectionName":  autorest.Encode("path", peConnectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}/privateEndpointConnections/{peConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeletePrivateEndpointConnectionSender sends the DeletePrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) DeletePrivateEndpointConnectionSender(req *http.Request) (future PrivateLinkServicesDeletePrivateEndpointConnectionFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeletePrivateEndpointConnectionResponder handles the response to the DeletePrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) DeletePrivateEndpointConnectionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified private link service by resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the private link service.
// expand - expands referenced resources.
func (client PrivateLinkServicesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, expand string) (result PrivateLinkService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) GetResponder(resp *http.Response) (result PrivateLinkService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all private link services in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client PrivateLinkServicesClient) List(ctx context.Context, resourceGroupName string) (result PrivateLinkServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.List")
		defer func() {
			sc := -1
			if result.plslr.Response.Response != nil {
				sc = result.plslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.plslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "List", resp, "Failure sending request")
		return
	}

	result.plslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateLinkServicesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) ListResponder(resp *http.Response) (result PrivateLinkServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateLinkServicesClient) listNextResults(ctx context.Context, lastResults PrivateLinkServiceListResult) (result PrivateLinkServiceListResult, err error) {
	req, err := lastResults.privateLinkServiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkServicesClient) ListComplete(ctx context.Context, resourceGroupName string) (result PrivateLinkServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListBySubscription gets all private link service in a subscription.
func (client PrivateLinkServicesClient) ListBySubscription(ctx context.Context) (result PrivateLinkServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.plslr.Response.Response != nil {
				sc = result.plslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.plslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.plslr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client PrivateLinkServicesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateLinkServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) ListBySubscriptionResponder(resp *http.Response) (result PrivateLinkServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client PrivateLinkServicesClient) listBySubscriptionNextResults(ctx context.Context, lastResults PrivateLinkServiceListResult) (result PrivateLinkServiceListResult, err error) {
	req, err := lastResults.privateLinkServiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkServicesClient) ListBySubscriptionComplete(ctx context.Context) (result PrivateLinkServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// UpdatePrivateEndpointConnection approve or reject private end point connection for a private link service in a
// subscription.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the private link service.
// peConnectionName - the name of the private end point connection.
// parameters - parameters supplied to approve or reject the private end point connection.
func (client PrivateLinkServicesClient) UpdatePrivateEndpointConnection(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string, parameters PrivateEndpointConnection) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesClient.UpdatePrivateEndpointConnection")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePrivateEndpointConnectionPreparer(ctx, resourceGroupName, serviceName, peConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "UpdatePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdatePrivateEndpointConnectionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "UpdatePrivateEndpointConnection", resp, "Failure sending request")
		return
	}

	result, err = client.UpdatePrivateEndpointConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateLinkServicesClient", "UpdatePrivateEndpointConnection", resp, "Failure responding to request")
	}

	return
}

// UpdatePrivateEndpointConnectionPreparer prepares the UpdatePrivateEndpointConnection request.
func (client PrivateLinkServicesClient) UpdatePrivateEndpointConnectionPreparer(ctx context.Context, resourceGroupName string, serviceName string, peConnectionName string, parameters PrivateEndpointConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peConnectionName":  autorest.Encode("path", peConnectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}/privateEndpointConnections/{peConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePrivateEndpointConnectionSender sends the UpdatePrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesClient) UpdatePrivateEndpointConnectionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdatePrivateEndpointConnectionResponder handles the response to the UpdatePrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesClient) UpdatePrivateEndpointConnectionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
