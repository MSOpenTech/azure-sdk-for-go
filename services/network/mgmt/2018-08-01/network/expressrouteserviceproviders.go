package network

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

// ExpressRouteServiceProvidersClient is the network Client
type ExpressRouteServiceProvidersClient struct {
	BaseClient
}

// NewExpressRouteServiceProvidersClient creates an instance of the ExpressRouteServiceProvidersClient client.
func NewExpressRouteServiceProvidersClient(subscriptionID string) ExpressRouteServiceProvidersClient {
	return NewExpressRouteServiceProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteServiceProvidersClientWithBaseURI creates an instance of the ExpressRouteServiceProvidersClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewExpressRouteServiceProvidersClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteServiceProvidersClient {
	return ExpressRouteServiceProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets all the available express route service providers.
func (client ExpressRouteServiceProvidersClient) List(ctx context.Context) (result ExpressRouteServiceProviderListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteServiceProvidersClient.List")
		defer func() {
			sc := -1
			if result.ersplr.Response.Response != nil {
				sc = result.ersplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ersplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersClient", "List", resp, "Failure sending request")
		return
	}

	result.ersplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ersplr.hasNextLink() && result.ersplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteServiceProvidersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteServiceProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteServiceProvidersClient) ListResponder(resp *http.Response) (result ExpressRouteServiceProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRouteServiceProvidersClient) listNextResults(ctx context.Context, lastResults ExpressRouteServiceProviderListResult) (result ExpressRouteServiceProviderListResult, err error) {
	req, err := lastResults.expressRouteServiceProviderListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteServiceProvidersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRouteServiceProvidersClient) ListComplete(ctx context.Context) (result ExpressRouteServiceProviderListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteServiceProvidersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
