package keyvault

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

// MHSMPrivateLinkResourcesClient is the the Azure management API provides a RESTful set of web services that interact
// with Azure Key Vault.
type MHSMPrivateLinkResourcesClient struct {
	BaseClient
}

// NewMHSMPrivateLinkResourcesClient creates an instance of the MHSMPrivateLinkResourcesClient client.
func NewMHSMPrivateLinkResourcesClient(subscriptionID string) MHSMPrivateLinkResourcesClient {
	return NewMHSMPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMHSMPrivateLinkResourcesClientWithBaseURI creates an instance of the MHSMPrivateLinkResourcesClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewMHSMPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) MHSMPrivateLinkResourcesClient {
	return MHSMPrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByMHSMResource gets the private link resources supported for the managed hsm pool.
// Parameters:
// resourceGroupName - name of the resource group that contains the managed HSM pool.
// name - name of the managed HSM Pool
func (client MHSMPrivateLinkResourcesClient) ListByMHSMResource(ctx context.Context, resourceGroupName string, name string) (result MHSMPrivateLinkResourceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MHSMPrivateLinkResourcesClient.ListByMHSMResource")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByMHSMResourcePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.MHSMPrivateLinkResourcesClient", "ListByMHSMResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByMHSMResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.MHSMPrivateLinkResourcesClient", "ListByMHSMResource", resp, "Failure sending request")
		return
	}

	result, err = client.ListByMHSMResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.MHSMPrivateLinkResourcesClient", "ListByMHSMResource", resp, "Failure responding to request")
		return
	}

	return
}

// ListByMHSMResourcePreparer prepares the ListByMHSMResource request.
func (client MHSMPrivateLinkResourcesClient) ListByMHSMResourcePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByMHSMResourceSender sends the ListByMHSMResource request. The method will close the
// http.Response Body if it receives an error.
func (client MHSMPrivateLinkResourcesClient) ListByMHSMResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByMHSMResourceResponder handles the response to the ListByMHSMResource request. The method always
// closes the http.Response Body.
func (client MHSMPrivateLinkResourcesClient) ListByMHSMResourceResponder(resp *http.Response) (result MHSMPrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
