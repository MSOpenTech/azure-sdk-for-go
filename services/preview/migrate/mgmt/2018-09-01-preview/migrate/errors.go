package migrate

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

// ErrorsClient is the migrate your workloads to Azure.
type ErrorsClient struct {
	BaseClient
}

// NewErrorsClient creates an instance of the ErrorsClient client.
func NewErrorsClient(subscriptionID string, acceptLanguage string) ErrorsClient {
	return NewErrorsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewErrorsClientWithBaseURI creates an instance of the ErrorsClient client.
func NewErrorsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) ErrorsClient {
	return ErrorsClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// EnumerateEvents sends the enumerate events request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// continuationToken - the continuation token.
func (client ErrorsClient) EnumerateEvents(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string) (result EventCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ErrorsClient.EnumerateEvents")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnumerateEventsPreparer(ctx, resourceGroupName, migrateProjectName, continuationToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ErrorsClient", "EnumerateEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnumerateEventsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ErrorsClient", "EnumerateEvents", resp, "Failure sending request")
		return
	}

	result, err = client.EnumerateEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ErrorsClient", "EnumerateEvents", resp, "Failure responding to request")
	}

	return
}

// EnumerateEventsPreparer prepares the EnumerateEvents request.
func (client ErrorsClient) EnumerateEventsPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(continuationToken) > 0 {
		queryParameters["continuationToken"] = autorest.Encode("query", continuationToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/MigrateProjects/{migrateProjectName}/MigrateEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnumerateEventsSender sends the EnumerateEvents request. The method will close the
// http.Response Body if it receives an error.
func (client ErrorsClient) EnumerateEventsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// EnumerateEventsResponder handles the response to the EnumerateEvents request. The method always
// closes the http.Response Body.
func (client ErrorsClient) EnumerateEventsResponder(resp *http.Response) (result EventCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
