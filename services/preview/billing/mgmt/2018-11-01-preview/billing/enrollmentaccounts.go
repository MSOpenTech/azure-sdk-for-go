package billing

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

// EnrollmentAccountsClient is the billing client provides access to billing resources for Azure subscriptions.
type EnrollmentAccountsClient struct {
	BaseClient
}

// NewEnrollmentAccountsClient creates an instance of the EnrollmentAccountsClient client.
func NewEnrollmentAccountsClient(subscriptionID string) EnrollmentAccountsClient {
	return NewEnrollmentAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEnrollmentAccountsClientWithBaseURI creates an instance of the EnrollmentAccountsClient client.
func NewEnrollmentAccountsClientWithBaseURI(baseURI string, subscriptionID string) EnrollmentAccountsClient {
	return EnrollmentAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByEnrollmentAccountID get the enrollment account by id.
// Parameters:
// billingAccountName - billing Account Id.
// enrollmentAccountName - enrollment Account Id.
// expand - may be used to expand the Department.
// filter - the filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne',
// 'or', or 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:).
func (client EnrollmentAccountsClient) GetByEnrollmentAccountID(ctx context.Context, billingAccountName string, enrollmentAccountName string, expand string, filter string) (result EnrollmentAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentAccountsClient.GetByEnrollmentAccountID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByEnrollmentAccountIDPreparer(ctx, billingAccountName, enrollmentAccountName, expand, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.EnrollmentAccountsClient", "GetByEnrollmentAccountID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByEnrollmentAccountIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.EnrollmentAccountsClient", "GetByEnrollmentAccountID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByEnrollmentAccountIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.EnrollmentAccountsClient", "GetByEnrollmentAccountID", resp, "Failure responding to request")
	}

	return
}

// GetByEnrollmentAccountIDPreparer prepares the GetByEnrollmentAccountID request.
func (client EnrollmentAccountsClient) GetByEnrollmentAccountIDPreparer(ctx context.Context, billingAccountName string, enrollmentAccountName string, expand string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":    autorest.Encode("path", billingAccountName),
		"enrollmentAccountName": autorest.Encode("path", enrollmentAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByEnrollmentAccountIDSender sends the GetByEnrollmentAccountID request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentAccountsClient) GetByEnrollmentAccountIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByEnrollmentAccountIDResponder handles the response to the GetByEnrollmentAccountID request. The method always
// closes the http.Response Body.
func (client EnrollmentAccountsClient) GetByEnrollmentAccountIDResponder(resp *http.Response) (result EnrollmentAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccountName lists all Enrollment Accounts for a user which he has access to.
// Parameters:
// billingAccountName - billing Account Id.
// expand - may be used to expand the department.
// filter - the filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne',
// 'or', or 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:).
func (client EnrollmentAccountsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string, filter string) (result EnrollmentAccountListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentAccountsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName, expand, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.EnrollmentAccountsClient", "ListByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.EnrollmentAccountsClient", "ListByBillingAccountName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.EnrollmentAccountsClient", "ListByBillingAccountName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
func (client EnrollmentAccountsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, expand string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentAccountsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client EnrollmentAccountsClient) ListByBillingAccountNameResponder(resp *http.Response) (result EnrollmentAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
