package azurestack

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// CustomerSubscriptionsClient is the azure Stack
type CustomerSubscriptionsClient struct {
	BaseClient
}

// NewCustomerSubscriptionsClient creates an instance of the CustomerSubscriptionsClient client.
func NewCustomerSubscriptionsClient(subscriptionID string) CustomerSubscriptionsClient {
	return NewCustomerSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomerSubscriptionsClientWithBaseURI creates an instance of the CustomerSubscriptionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewCustomerSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) CustomerSubscriptionsClient {
	return CustomerSubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new customer subscription under a registration.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// customerSubscriptionName - name of the product.
// customerCreationParameters - parameters use to create a customer subscription.
func (client CustomerSubscriptionsClient) Create(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, customerCreationParameters CustomerSubscription) (result CustomerSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomerSubscriptionsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroup, registrationName, customerSubscriptionName, customerCreationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client CustomerSubscriptionsClient) CreatePreparer(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, customerCreationParameters CustomerSubscription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customerSubscriptionName": autorest.Encode("path", customerSubscriptionName),
		"registrationName":         autorest.Encode("path", registrationName),
		"resourceGroup":            autorest.Encode("path", resourceGroup),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", pathParameters),
		autorest.WithJSON(customerCreationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CustomerSubscriptionsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CustomerSubscriptionsClient) CreateResponder(resp *http.Response) (result CustomerSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a customer subscription under a registration.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// customerSubscriptionName - name of the product.
func (client CustomerSubscriptionsClient) Delete(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomerSubscriptionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroup, registrationName, customerSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CustomerSubscriptionsClient) DeletePreparer(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customerSubscriptionName": autorest.Encode("path", customerSubscriptionName),
		"registrationName":         autorest.Encode("path", registrationName),
		"resourceGroup":            autorest.Encode("path", resourceGroup),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CustomerSubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomerSubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the specified product.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// customerSubscriptionName - name of the product.
func (client CustomerSubscriptionsClient) Get(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string) (result CustomerSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomerSubscriptionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroup, registrationName, customerSubscriptionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomerSubscriptionsClient) GetPreparer(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customerSubscriptionName": autorest.Encode("path", customerSubscriptionName),
		"registrationName":         autorest.Encode("path", registrationName),
		"resourceGroup":            autorest.Encode("path", resourceGroup),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomerSubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomerSubscriptionsClient) GetResponder(resp *http.Response) (result CustomerSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns a list of products.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
func (client CustomerSubscriptionsClient) List(ctx context.Context, resourceGroup string, registrationName string) (result CustomerSubscriptionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomerSubscriptionsClient.List")
		defer func() {
			sc := -1
			if result.csl.Response.Response != nil {
				sc = result.csl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroup, registrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.csl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "List", resp, "Failure sending request")
		return
	}

	result.csl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.csl.hasNextLink() && result.csl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CustomerSubscriptionsClient) ListPreparer(ctx context.Context, resourceGroup string, registrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CustomerSubscriptionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CustomerSubscriptionsClient) ListResponder(resp *http.Response) (result CustomerSubscriptionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CustomerSubscriptionsClient) listNextResults(ctx context.Context, lastResults CustomerSubscriptionList) (result CustomerSubscriptionList, err error) {
	req, err := lastResults.customerSubscriptionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.CustomerSubscriptionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomerSubscriptionsClient) ListComplete(ctx context.Context, resourceGroup string, registrationName string) (result CustomerSubscriptionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomerSubscriptionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroup, registrationName)
	return
}
