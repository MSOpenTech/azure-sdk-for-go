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

// AzureFirewallFqdnTagsClient is the network Client
type AzureFirewallFqdnTagsClient struct {
	BaseClient
}

// NewAzureFirewallFqdnTagsClient creates an instance of the AzureFirewallFqdnTagsClient client.
func NewAzureFirewallFqdnTagsClient(subscriptionID string) AzureFirewallFqdnTagsClient {
	return NewAzureFirewallFqdnTagsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAzureFirewallFqdnTagsClientWithBaseURI creates an instance of the AzureFirewallFqdnTagsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewAzureFirewallFqdnTagsClientWithBaseURI(baseURI string, subscriptionID string) AzureFirewallFqdnTagsClient {
	return AzureFirewallFqdnTagsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListAll gets all the Azure Firewall FQDN Tags in a subscription.
func (client AzureFirewallFqdnTagsClient) ListAll(ctx context.Context) (result AzureFirewallFqdnTagListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureFirewallFqdnTagsClient.ListAll")
		defer func() {
			sc := -1
			if result.afftlr.Response.Response != nil {
				sc = result.afftlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AzureFirewallFqdnTagsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.afftlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.AzureFirewallFqdnTagsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.afftlr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AzureFirewallFqdnTagsClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client AzureFirewallFqdnTagsClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client AzureFirewallFqdnTagsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client AzureFirewallFqdnTagsClient) ListAllResponder(resp *http.Response) (result AzureFirewallFqdnTagListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client AzureFirewallFqdnTagsClient) listAllNextResults(ctx context.Context, lastResults AzureFirewallFqdnTagListResult) (result AzureFirewallFqdnTagListResult, err error) {
	req, err := lastResults.azureFirewallFqdnTagListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.AzureFirewallFqdnTagsClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.AzureFirewallFqdnTagsClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AzureFirewallFqdnTagsClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client AzureFirewallFqdnTagsClient) ListAllComplete(ctx context.Context) (result AzureFirewallFqdnTagListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureFirewallFqdnTagsClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}
