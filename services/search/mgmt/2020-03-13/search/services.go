package search

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// ServicesClient is the client that can be used to manage Azure Cognitive Search services and API keys.
type ServicesClient struct {
	BaseClient
}

// NewServicesClient creates an instance of the ServicesClient client.
func NewServicesClient(subscriptionID string) ServicesClient {
	return NewServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServicesClientWithBaseURI creates an instance of the ServicesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return ServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks whether or not the given Search service name is available for use. Search service names
// must be globally unique since they are part of the service URI (https://<name>.search.windows.net).
// Parameters:
// checkNameAvailabilityInput - the resource name and type to check.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, clientRequestID *uuid.UUID) (result CheckNameAvailabilityOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "checkNameAvailabilityInput.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("search.ServicesClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, checkNameAvailabilityInput, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ServicesClient) CheckNameAvailabilityPreparer(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Search/checkNameAvailability", pathParameters),
		autorest.WithJSON(checkNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ServicesClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a Search service in the given resource group. If the Search service already
// exists, all properties will be updated with the given values.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service to create or update. Search service names
// must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one
// characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search
// service names must be globally unique since they are part of the service URI
// (https://<name>.search.windows.net). You cannot change the service name after the service is created.
// service - the definition of the Search service to create or update.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, clientRequestID *uuid.UUID) (result ServicesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: service,
			Constraints: []validation.Constraint{{Target: "service.ServiceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "service.ServiceProperties.ReplicaCount", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "service.ServiceProperties.ReplicaCount", Name: validation.InclusiveMaximum, Rule: int64(12), Chain: nil},
						{Target: "service.ServiceProperties.ReplicaCount", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
					}},
					{Target: "service.ServiceProperties.PartitionCount", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "service.ServiceProperties.PartitionCount", Name: validation.InclusiveMaximum, Rule: int64(12), Chain: nil},
							{Target: "service.ServiceProperties.PartitionCount", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("search.ServicesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, searchServiceName, service, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", pathParameters),
		autorest.WithJSON(service),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) CreateOrUpdateSender(req *http.Request) (future ServicesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServicesClient) (s Service, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("search.ServicesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		s.Response.Response, err = future.GetResult(sender)
		if s.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && s.Response.Response.StatusCode != http.StatusNoContent {
			s, err = client.CreateOrUpdateResponder(s.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "search.ServicesCreateOrUpdateFuture", "Result", s.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServicesClient) CreateOrUpdateResponder(resp *http.Response) (result Service, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Search service in the given resource group, along with its associated resources.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) Delete(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, searchServiceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the Search service with the given name in the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result Service, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, searchServiceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServicesClient) GetResponder(resp *http.Response) (result Service, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets a list of all Search services in the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, clientRequestID *uuid.UUID) (result ServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.slr.hasNextLink() && result.slr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ServicesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServicesClient) ListByResourceGroupResponder(resp *http.Response) (result ServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ServicesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ServiceListResult) (result ServiceListResult, err error) {
	req, err := lastResults.serviceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServicesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, clientRequestID *uuid.UUID) (result ServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, clientRequestID)
	return
}

// ListBySubscription gets a list of all Search services in the given subscription.
// Parameters:
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) ListBySubscription(ctx context.Context, clientRequestID *uuid.UUID) (result ServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.slr.hasNextLink() && result.slr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ServicesClient) ListBySubscriptionPreparer(ctx context.Context, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Search/searchServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ServicesClient) ListBySubscriptionResponder(resp *http.Response) (result ServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client ServicesClient) listBySubscriptionNextResults(ctx context.Context, lastResults ServiceListResult) (result ServiceListResult, err error) {
	req, err := lastResults.serviceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.ServicesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServicesClient) ListBySubscriptionComplete(ctx context.Context, clientRequestID *uuid.UUID) (result ServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, clientRequestID)
	return
}

// Update updates an existing Search service in the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service to update.
// service - the definition of the Search service to update.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client ServicesClient) Update(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, clientRequestID *uuid.UUID) (result Service, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, searchServiceName, service, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.ServicesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServicesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, service Service, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", pathParameters),
		autorest.WithJSON(service),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServicesClient) UpdateResponder(resp *http.Response) (result Service, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
