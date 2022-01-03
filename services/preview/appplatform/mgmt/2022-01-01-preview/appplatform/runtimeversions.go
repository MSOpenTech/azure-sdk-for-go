package appplatform

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

// RuntimeVersionsClient is the REST API for Azure Spring Cloud
type RuntimeVersionsClient struct {
	BaseClient
}

// NewRuntimeVersionsClient creates an instance of the RuntimeVersionsClient client.
func NewRuntimeVersionsClient(subscriptionID string) RuntimeVersionsClient {
	return NewRuntimeVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRuntimeVersionsClientWithBaseURI creates an instance of the RuntimeVersionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRuntimeVersionsClientWithBaseURI(baseURI string, subscriptionID string) RuntimeVersionsClient {
	return RuntimeVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListRuntimeVersions lists all of the available runtime versions supported by Microsoft.AppPlatform provider.
func (client RuntimeVersionsClient) ListRuntimeVersions(ctx context.Context) (result AvailableRuntimeVersions, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuntimeVersionsClient.ListRuntimeVersions")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRuntimeVersionsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.RuntimeVersionsClient", "ListRuntimeVersions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRuntimeVersionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.RuntimeVersionsClient", "ListRuntimeVersions", resp, "Failure sending request")
		return
	}

	result, err = client.ListRuntimeVersionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.RuntimeVersionsClient", "ListRuntimeVersions", resp, "Failure responding to request")
		return
	}

	return
}

// ListRuntimeVersionsPreparer prepares the ListRuntimeVersions request.
func (client RuntimeVersionsClient) ListRuntimeVersionsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.AppPlatform/runtimeVersions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRuntimeVersionsSender sends the ListRuntimeVersions request. The method will close the
// http.Response Body if it receives an error.
func (client RuntimeVersionsClient) ListRuntimeVersionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListRuntimeVersionsResponder handles the response to the ListRuntimeVersions request. The method always
// closes the http.Response Body.
func (client RuntimeVersionsClient) ListRuntimeVersionsResponder(resp *http.Response) (result AvailableRuntimeVersions, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
