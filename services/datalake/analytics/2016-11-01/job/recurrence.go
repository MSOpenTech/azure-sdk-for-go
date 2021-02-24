package job

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// RecurrenceClient is the creates an Azure Data Lake Analytics job client.
type RecurrenceClient struct {
	BaseClient
}

// NewRecurrenceClient creates an instance of the RecurrenceClient client.
func NewRecurrenceClient() RecurrenceClient {
	return RecurrenceClient{New()}
}

// Get gets the recurrence information for the specified recurrence ID.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// recurrenceIdentity - recurrence ID.
// startDateTime - the start date for when to get the recurrence and aggregate its data. The startDateTime and
// endDateTime can be no more than 30 days apart.
// endDateTime - the end date for when to get recurrence and aggregate its data. The startDateTime and
// endDateTime can be no more than 30 days apart.
func (client RecurrenceClient) Get(ctx context.Context, accountName string, recurrenceIdentity uuid.UUID, startDateTime *date.Time, endDateTime *date.Time) (result RecurrenceInformation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecurrenceClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, accountName, recurrenceIdentity, startDateTime, endDateTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecurrenceClient) GetPreparer(ctx context.Context, accountName string, recurrenceIdentity uuid.UUID, startDateTime *date.Time, endDateTime *date.Time) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"recurrenceIdentity": autorest.Encode("path", recurrenceIdentity),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startDateTime != nil {
		queryParameters["startDateTime"] = autorest.Encode("query", *startDateTime)
	}
	if endDateTime != nil {
		queryParameters["endDateTime"] = autorest.Encode("query", *endDateTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/recurrences/{recurrenceIdentity}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecurrenceClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecurrenceClient) GetResponder(resp *http.Response) (result RecurrenceInformation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all recurrences.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// startDateTime - the start date for when to get the list of recurrences. The startDateTime and endDateTime
// can be no more than 30 days apart.
// endDateTime - the end date for when to get the list of recurrences. The startDateTime and endDateTime can be
// no more than 30 days apart.
func (client RecurrenceClient) List(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result RecurrenceInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecurrenceClient.List")
		defer func() {
			sc := -1
			if result.rilr.Response.Response != nil {
				sc = result.rilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, accountName, startDateTime, endDateTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "List", resp, "Failure sending request")
		return
	}

	result.rilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rilr.hasNextLink() && result.rilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RecurrenceClient) ListPreparer(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startDateTime != nil {
		queryParameters["startDateTime"] = autorest.Encode("query", *startDateTime)
	}
	if endDateTime != nil {
		queryParameters["endDateTime"] = autorest.Encode("query", *endDateTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/recurrences"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecurrenceClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecurrenceClient) ListResponder(resp *http.Response) (result RecurrenceInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecurrenceClient) listNextResults(ctx context.Context, lastResults RecurrenceInformationListResult) (result RecurrenceInformationListResult, err error) {
	req, err := lastResults.recurrenceInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "job.RecurrenceClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "job.RecurrenceClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.RecurrenceClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecurrenceClient) ListComplete(ctx context.Context, accountName string, startDateTime *date.Time, endDateTime *date.Time) (result RecurrenceInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecurrenceClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, accountName, startDateTime, endDateTime)
	return
}
