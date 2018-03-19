package automation

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// TestJobStreamsGroupClient is the automation Client
type TestJobStreamsGroupClient struct {
	BaseClient
}

// NewTestJobStreamsGroupClient creates an instance of the TestJobStreamsGroupClient client.
func NewTestJobStreamsGroupClient(subscriptionID string, automationAccountName string, automationAccountName1 string, clientRequestID string) TestJobStreamsGroupClient {
	return NewTestJobStreamsGroupClientWithBaseURI(DefaultBaseURI, subscriptionID, automationAccountName, automationAccountName1, clientRequestID)
}

// NewTestJobStreamsGroupClientWithBaseURI creates an instance of the TestJobStreamsGroupClient client.
func NewTestJobStreamsGroupClientWithBaseURI(baseURI string, subscriptionID string, automationAccountName string, automationAccountName1 string, clientRequestID string) TestJobStreamsGroupClient {
	return TestJobStreamsGroupClient{NewWithBaseURI(baseURI, subscriptionID, automationAccountName, automationAccountName1, clientRequestID)}
}

// Get retrieve a test job stream of the test job identified by runbook name and stream id.
//
// resourceGroupName is name of an Azure Resource group. automationAccountName is the automation account name.
// runbookName is the runbook name. jobStreamID is the job stream id.
func (client TestJobStreamsGroupClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (result JobStream, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.TestJobStreamsGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, runbookName, jobStreamID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TestJobStreamsGroupClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobStreamId":           autorest.Encode("path", jobStreamID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TestJobStreamsGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TestJobStreamsGroupClient) GetResponder(resp *http.Response) (result JobStream, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTestJob retrieve a list of test job streams identified by runbook name.
//
// resourceGroupName is name of an Azure Resource group. automationAccountName is the automation account name.
// runbookName is the runbook name. filter is the filter to apply on the operation.
func (client TestJobStreamsGroupClient) ListByTestJob(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (result JobStreamListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.TestJobStreamsGroupClient", "ListByTestJob", err.Error())
	}

	result.fn = client.listByTestJobNextResults
	req, err := client.ListByTestJobPreparer(ctx, resourceGroupName, automationAccountName, runbookName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "ListByTestJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTestJobSender(req)
	if err != nil {
		result.jslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "ListByTestJob", resp, "Failure sending request")
		return
	}

	result.jslr, err = client.ListByTestJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "ListByTestJob", resp, "Failure responding to request")
	}

	return
}

// ListByTestJobPreparer prepares the ListByTestJob request.
func (client TestJobStreamsGroupClient) ListByTestJobPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTestJobSender sends the ListByTestJob request. The method will close the
// http.Response Body if it receives an error.
func (client TestJobStreamsGroupClient) ListByTestJobSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByTestJobResponder handles the response to the ListByTestJob request. The method always
// closes the http.Response Body.
func (client TestJobStreamsGroupClient) ListByTestJobResponder(resp *http.Response) (result JobStreamListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTestJobNextResults retrieves the next set of results, if any.
func (client TestJobStreamsGroupClient) listByTestJobNextResults(lastResults JobStreamListResult) (result JobStreamListResult, err error) {
	req, err := lastResults.jobStreamListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "listByTestJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTestJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "listByTestJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTestJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsGroupClient", "listByTestJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTestJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client TestJobStreamsGroupClient) ListByTestJobComplete(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (result JobStreamListResultIterator, err error) {
	result.page, err = client.ListByTestJob(ctx, resourceGroupName, automationAccountName, runbookName, filter)
	return
}
