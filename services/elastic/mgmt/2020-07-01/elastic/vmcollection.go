package elastic

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

// VMCollectionClient is the client for the VMCollection methods of the Elastic service.
type VMCollectionClient struct {
	BaseClient
}

// NewVMCollectionClient creates an instance of the VMCollectionClient client.
func NewVMCollectionClient(subscriptionID string) VMCollectionClient {
	return NewVMCollectionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVMCollectionClientWithBaseURI creates an instance of the VMCollectionClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVMCollectionClientWithBaseURI(baseURI string, subscriptionID string) VMCollectionClient {
	return VMCollectionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Update sends the update request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
// body - VM resource Id
func (client VMCollectionClient) Update(ctx context.Context, resourceGroupName string, monitorName string, body *VMCollectionUpdate) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VMCollectionClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, monitorName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMCollectionClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "elastic.VMCollectionClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMCollectionClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VMCollectionClient) UpdatePreparer(ctx context.Context, resourceGroupName string, monitorName string, body *VMCollectionUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/vmCollectionUpdate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VMCollectionClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VMCollectionClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
