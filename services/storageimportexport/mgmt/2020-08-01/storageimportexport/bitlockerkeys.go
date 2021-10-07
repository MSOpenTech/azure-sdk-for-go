package storageimportexport

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

// BitLockerKeysClient is the the Storage Import/Export Resource Provider API.
type BitLockerKeysClient struct {
	BaseClient
}

// NewBitLockerKeysClient creates an instance of the BitLockerKeysClient client.
func NewBitLockerKeysClient(subscriptionID string, acceptLanguage string) BitLockerKeysClient {
	return NewBitLockerKeysClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewBitLockerKeysClientWithBaseURI creates an instance of the BitLockerKeysClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBitLockerKeysClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) BitLockerKeysClient {
	return BitLockerKeysClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// List returns the BitLocker Keys for all drives in the specified job.
// Parameters:
// jobName - the name of the import/export job.
// resourceGroupName - the resource group name uniquely identifies the resource group within the user
// subscription.
func (client BitLockerKeysClient) List(ctx context.Context, jobName string, resourceGroupName string) (result GetBitLockerKeysResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BitLockerKeysClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, jobName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.BitLockerKeysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storageimportexport.BitLockerKeysClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageimportexport.BitLockerKeysClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BitLockerKeysClient) ListPreparer(ctx context.Context, jobName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}/listBitLockerKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BitLockerKeysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BitLockerKeysClient) ListResponder(resp *http.Response) (result GetBitLockerKeysResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
