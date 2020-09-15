package postgresqlflexibleservers

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
	"net/http"
)

// VirtualNetworkSubnetUsageClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure PostgreSQL resources including servers, databases, firewall rules, VNET rules, security
// alert policies, log files and configurations with new business model.
type VirtualNetworkSubnetUsageClient struct {
	BaseClient
}

// NewVirtualNetworkSubnetUsageClient creates an instance of the VirtualNetworkSubnetUsageClient client.
func NewVirtualNetworkSubnetUsageClient(subscriptionID string) VirtualNetworkSubnetUsageClient {
	return NewVirtualNetworkSubnetUsageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkSubnetUsageClientWithBaseURI creates an instance of the VirtualNetworkSubnetUsageClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewVirtualNetworkSubnetUsageClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkSubnetUsageClient {
	return VirtualNetworkSubnetUsageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Execute get virtual network subnet usage for a given vNet resource id.
// Parameters:
// locationName - the name of the location.
// parameters - the required parameters for creating or updating a server.
func (client VirtualNetworkSubnetUsageClient) Execute(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter) (result VirtualNetworkSubnetUsageResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkSubnetUsageClient.Execute")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlflexibleservers.VirtualNetworkSubnetUsageClient", "Execute", err.Error())
	}

	req, err := client.ExecutePreparer(ctx, locationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.VirtualNetworkSubnetUsageClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.VirtualNetworkSubnetUsageClient", "Execute", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlflexibleservers.VirtualNetworkSubnetUsageClient", "Execute", resp, "Failure responding to request")
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client VirtualNetworkSubnetUsageClient) ExecutePreparer(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSql/locations/{locationName}/checkVirtualNetworkSubnetUsage", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkSubnetUsageClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client VirtualNetworkSubnetUsageClient) ExecuteResponder(resp *http.Response) (result VirtualNetworkSubnetUsageResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
