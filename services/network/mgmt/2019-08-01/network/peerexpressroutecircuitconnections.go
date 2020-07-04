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

// PeerExpressRouteCircuitConnectionsClient is the network Client
type PeerExpressRouteCircuitConnectionsClient struct {
	BaseClient
}

// NewPeerExpressRouteCircuitConnectionsClient creates an instance of the PeerExpressRouteCircuitConnectionsClient
// client.
func NewPeerExpressRouteCircuitConnectionsClient(subscriptionID string) PeerExpressRouteCircuitConnectionsClient {
	return NewPeerExpressRouteCircuitConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPeerExpressRouteCircuitConnectionsClientWithBaseURI creates an instance of the
// PeerExpressRouteCircuitConnectionsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPeerExpressRouteCircuitConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PeerExpressRouteCircuitConnectionsClient {
	return PeerExpressRouteCircuitConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the express route circuit.
// peeringName - the name of the peering.
// connectionName - the name of the peer express route circuit connection.
func (client PeerExpressRouteCircuitConnectionsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string) (result PeerExpressRouteCircuitConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeerExpressRouteCircuitConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, circuitName, peeringName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PeerExpressRouteCircuitConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"connectionName":    autorest.Encode("path", connectionName),
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PeerExpressRouteCircuitConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PeerExpressRouteCircuitConnectionsClient) GetResponder(resp *http.Response) (result PeerExpressRouteCircuitConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all global reach peer connections associated with a private peering in an express route circuit.
// Parameters:
// resourceGroupName - the name of the resource group.
// circuitName - the name of the circuit.
// peeringName - the name of the peering.
func (client PeerExpressRouteCircuitConnectionsClient) List(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result PeerExpressRouteCircuitConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeerExpressRouteCircuitConnectionsClient.List")
		defer func() {
			sc := -1
			if result.percclr.Response.Response != nil {
				sc = result.percclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, circuitName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.percclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.percclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PeerExpressRouteCircuitConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       autorest.Encode("path", circuitName),
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PeerExpressRouteCircuitConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PeerExpressRouteCircuitConnectionsClient) ListResponder(resp *http.Response) (result PeerExpressRouteCircuitConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PeerExpressRouteCircuitConnectionsClient) listNextResults(ctx context.Context, lastResults PeerExpressRouteCircuitConnectionListResult) (result PeerExpressRouteCircuitConnectionListResult, err error) {
	req, err := lastResults.peerExpressRouteCircuitConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PeerExpressRouteCircuitConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PeerExpressRouteCircuitConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result PeerExpressRouteCircuitConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeerExpressRouteCircuitConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, circuitName, peeringName)
	return
}
