package adhybridhealthservice

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

// AddsServicesReplicationStatusClient is the REST APIs for Azure Active Directory Connect Health
type AddsServicesReplicationStatusClient struct {
	BaseClient
}

// NewAddsServicesReplicationStatusClient creates an instance of the AddsServicesReplicationStatusClient client.
func NewAddsServicesReplicationStatusClient() AddsServicesReplicationStatusClient {
	return NewAddsServicesReplicationStatusClientWithBaseURI(DefaultBaseURI)
}

// NewAddsServicesReplicationStatusClientWithBaseURI creates an instance of the AddsServicesReplicationStatusClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewAddsServicesReplicationStatusClientWithBaseURI(baseURI string) AddsServicesReplicationStatusClient {
	return AddsServicesReplicationStatusClient{NewWithBaseURI(baseURI)}
}

// Get gets Replication status for a given Active Directory Domain Service, that is onboarded to Azure Active Directory
// Connect Health.
// Parameters:
// serviceName - the name of the service.
func (client AddsServicesReplicationStatusClient) Get(ctx context.Context, serviceName string) (result ReplicationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddsServicesReplicationStatusClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesReplicationStatusClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesReplicationStatusClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AddsServicesReplicationStatusClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddsServicesReplicationStatusClient) GetPreparer(ctx context.Context, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationstatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddsServicesReplicationStatusClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddsServicesReplicationStatusClient) GetResponder(resp *http.Response) (result ReplicationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
