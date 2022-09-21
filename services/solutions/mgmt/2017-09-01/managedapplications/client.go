// Package managedapplications implements the Azure ARM Managedapplications service API version 2017-09-01.
//
// ARM applications
package managedapplications

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

const (
	// DefaultBaseURI is the default URI used for the service Managedapplications
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Managedapplications.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// ListOperations lists all of the available Microsoft.Solutions REST API operations.
func (client BaseClient) ListOperations(ctx context.Context) (result OperationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListOperations")
		defer func() {
			sc := -1
			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listOperationsNextResults
	req, err := client.ListOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.BaseClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.olr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.BaseClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result.olr, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.BaseClient", "ListOperations", resp, "Failure responding to request")
		return
	}
	if result.olr.hasNextLink() && result.olr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client BaseClient) ListOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Solutions/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client BaseClient) ListOperationsResponder(resp *http.Response) (result OperationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listOperationsNextResults retrieves the next set of results, if any.
func (client BaseClient) listOperationsNextResults(ctx context.Context, lastResults OperationListResult) (result OperationListResult, err error) {
	req, err := lastResults.operationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedapplications.BaseClient", "listOperationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedapplications.BaseClient", "listOperationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.BaseClient", "listOperationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListOperationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) ListOperationsComplete(ctx context.Context) (result OperationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListOperations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListOperations(ctx)
	return
}
