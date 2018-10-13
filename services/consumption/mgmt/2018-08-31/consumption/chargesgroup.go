package consumption

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
	"net/http"
)

// ChargesGroupClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type ChargesGroupClient struct {
	BaseClient
}

// NewChargesGroupClient creates an instance of the ChargesGroupClient client.
func NewChargesGroupClient(subscriptionID string) ChargesGroupClient {
	return NewChargesGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewChargesGroupClientWithBaseURI creates an instance of the ChargesGroupClient client.
func NewChargesGroupClientWithBaseURI(baseURI string, subscriptionID string) ChargesGroupClient {
	return ChargesGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByDepartment lists the charges by departmentId.
// Parameters:
// billingAccountID - billingAccount ID
// departmentID - department ID
// filter - may be used to filter charges by properties/usageEnd (Utc time), properties/usageStart (Utc time).
// The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or
// 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:).
func (client ChargesGroupClient) ListByDepartment(ctx context.Context, billingAccountID string, departmentID string, filter string) (result ChargesListResult, err error) {
	req, err := client.ListByDepartmentPreparer(ctx, billingAccountID, departmentID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListByDepartment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDepartmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListByDepartment", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDepartmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListByDepartment", resp, "Failure responding to request")
	}

	return
}

// ListByDepartmentPreparer prepares the ListByDepartment request.
func (client ChargesGroupClient) ListByDepartmentPreparer(ctx context.Context, billingAccountID string, departmentID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
		"departmentId":     autorest.Encode("path", departmentID),
	}

	const APIVersion = "2018-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}/providers/Microsoft.Consumption/charges", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDepartmentSender sends the ListByDepartment request. The method will close the
// http.Response Body if it receives an error.
func (client ChargesGroupClient) ListByDepartmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByDepartmentResponder handles the response to the ListByDepartment request. The method always
// closes the http.Response Body.
func (client ChargesGroupClient) ListByDepartmentResponder(resp *http.Response) (result ChargesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEnrollmentAccount lists the charges by enrollmentAccountId.
// Parameters:
// billingAccountID - billingAccount ID
// enrollmentAccountID - enrollmentAccount ID
// filter - may be used to filter charges by properties/usageEnd (Utc time), properties/usageStart (Utc time).
// The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or
// 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:).
func (client ChargesGroupClient) ListByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string) (result ChargesListResult, err error) {
	req, err := client.ListByEnrollmentAccountPreparer(ctx, billingAccountID, enrollmentAccountID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListByEnrollmentAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEnrollmentAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListByEnrollmentAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEnrollmentAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListByEnrollmentAccount", resp, "Failure responding to request")
	}

	return
}

// ListByEnrollmentAccountPreparer prepares the ListByEnrollmentAccount request.
func (client ChargesGroupClient) ListByEnrollmentAccountPreparer(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId":    autorest.Encode("path", billingAccountID),
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2018-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.Consumption/charges", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEnrollmentAccountSender sends the ListByEnrollmentAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ChargesGroupClient) ListByEnrollmentAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByEnrollmentAccountResponder handles the response to the ListByEnrollmentAccount request. The method always
// closes the http.Response Body.
func (client ChargesGroupClient) ListByEnrollmentAccountResponder(resp *http.Response) (result ChargesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForBillingPeriodByDepartment lists the charges based on departmentId by billing period.
// Parameters:
// billingAccountID - billingAccount ID
// departmentID - department ID
// billingPeriodName - billing Period Name.
// filter - may be used to filter charges by properties/usageEnd (Utc time), properties/usageStart (Utc time).
// The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or
// 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:).
func (client ChargesGroupClient) ListForBillingPeriodByDepartment(ctx context.Context, billingAccountID string, departmentID string, billingPeriodName string, filter string) (result ChargeSummary, err error) {
	req, err := client.ListForBillingPeriodByDepartmentPreparer(ctx, billingAccountID, departmentID, billingPeriodName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListForBillingPeriodByDepartment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForBillingPeriodByDepartmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListForBillingPeriodByDepartment", resp, "Failure sending request")
		return
	}

	result, err = client.ListForBillingPeriodByDepartmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListForBillingPeriodByDepartment", resp, "Failure responding to request")
	}

	return
}

// ListForBillingPeriodByDepartmentPreparer prepares the ListForBillingPeriodByDepartment request.
func (client ChargesGroupClient) ListForBillingPeriodByDepartmentPreparer(ctx context.Context, billingAccountID string, departmentID string, billingPeriodName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId":  autorest.Encode("path", billingAccountID),
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"departmentId":      autorest.Encode("path", departmentID),
	}

	const APIVersion = "2018-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/charges", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForBillingPeriodByDepartmentSender sends the ListForBillingPeriodByDepartment request. The method will close the
// http.Response Body if it receives an error.
func (client ChargesGroupClient) ListForBillingPeriodByDepartmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForBillingPeriodByDepartmentResponder handles the response to the ListForBillingPeriodByDepartment request. The method always
// closes the http.Response Body.
func (client ChargesGroupClient) ListForBillingPeriodByDepartmentResponder(resp *http.Response) (result ChargeSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListForBillingPeriodByEnrollmentAccount lists the charges based on enrollmentAccountId by billing period.
// Parameters:
// billingAccountID - billingAccount ID
// enrollmentAccountID - enrollmentAccount ID
// billingPeriodName - billing Period Name.
// filter - may be used to filter charges by properties/usageEnd (Utc time), properties/usageStart (Utc time).
// The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or
// 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:).
func (client ChargesGroupClient) ListForBillingPeriodByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, billingPeriodName string, filter string) (result ChargeSummary, err error) {
	req, err := client.ListForBillingPeriodByEnrollmentAccountPreparer(ctx, billingAccountID, enrollmentAccountID, billingPeriodName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListForBillingPeriodByEnrollmentAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForBillingPeriodByEnrollmentAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListForBillingPeriodByEnrollmentAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListForBillingPeriodByEnrollmentAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ChargesGroupClient", "ListForBillingPeriodByEnrollmentAccount", resp, "Failure responding to request")
	}

	return
}

// ListForBillingPeriodByEnrollmentAccountPreparer prepares the ListForBillingPeriodByEnrollmentAccount request.
func (client ChargesGroupClient) ListForBillingPeriodByEnrollmentAccountPreparer(ctx context.Context, billingAccountID string, enrollmentAccountID string, billingPeriodName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId":    autorest.Encode("path", billingAccountID),
		"billingPeriodName":   autorest.Encode("path", billingPeriodName),
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2018-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/charges", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForBillingPeriodByEnrollmentAccountSender sends the ListForBillingPeriodByEnrollmentAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ChargesGroupClient) ListForBillingPeriodByEnrollmentAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForBillingPeriodByEnrollmentAccountResponder handles the response to the ListForBillingPeriodByEnrollmentAccount request. The method always
// closes the http.Response Body.
func (client ChargesGroupClient) ListForBillingPeriodByEnrollmentAccountResponder(resp *http.Response) (result ChargeSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
