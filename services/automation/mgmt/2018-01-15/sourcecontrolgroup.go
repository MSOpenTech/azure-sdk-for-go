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

// SourceControlGroupClient is the automation Client
type SourceControlGroupClient struct {
	BaseClient
}

// NewSourceControlGroupClient creates an instance of the SourceControlGroupClient client.
func NewSourceControlGroupClient(subscriptionID string, automationAccountName string, automationAccountName1 string, clientRequestID string) SourceControlGroupClient {
	return NewSourceControlGroupClientWithBaseURI(DefaultBaseURI, subscriptionID, automationAccountName, automationAccountName1, clientRequestID)
}

// NewSourceControlGroupClientWithBaseURI creates an instance of the SourceControlGroupClient client.
func NewSourceControlGroupClientWithBaseURI(baseURI string, subscriptionID string, automationAccountName string, automationAccountName1 string, clientRequestID string) SourceControlGroupClient {
	return SourceControlGroupClient{NewWithBaseURI(baseURI, subscriptionID, automationAccountName, automationAccountName1, clientRequestID)}
}

// CreateOrUpdate create a source control.
//
// resourceGroupName is name of an Azure Resource group. sourceControlName is the source control name. parameters
// is the parameters supplied to the create or update source control operation.
func (client SourceControlGroupClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, sourceControlName string, parameters SourceControlCreateOrUpdateParameters) (result SourceControl, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties.RepoURL", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties.RepoURL", Name: validation.MaxLength, Rule: 2000, Chain: nil}}},
					{Target: "parameters.SourceControlCreateOrUpdateProperties.Branch", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties.Branch", Name: validation.MaxLength, Rule: 255, Chain: nil}}},
					{Target: "parameters.SourceControlCreateOrUpdateProperties.FolderPath", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties.FolderPath", Name: validation.MaxLength, Rule: 255, Chain: nil}}},
					{Target: "parameters.SourceControlCreateOrUpdateProperties.SecurityToken", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties.SecurityToken", Name: validation.MaxLength, Rule: 1024, Chain: nil}}},
					{Target: "parameters.SourceControlCreateOrUpdateProperties.Description", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.SourceControlCreateOrUpdateProperties.Description", Name: validation.MaxLength, Rule: 512, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlGroupClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, sourceControlName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SourceControlGroupClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, sourceControlName string, parameters SourceControlCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", client.AutomationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sourceControlName":     autorest.Encode("path", sourceControlName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlGroupClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SourceControlGroupClient) CreateOrUpdateResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the source control.
//
// resourceGroupName is name of an Azure Resource group. sourceControlName is the name of source control.
func (client SourceControlGroupClient) Delete(ctx context.Context, resourceGroupName string, sourceControlName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlGroupClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, sourceControlName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SourceControlGroupClient) DeletePreparer(ctx context.Context, resourceGroupName string, sourceControlName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", client.AutomationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sourceControlName":     autorest.Encode("path", sourceControlName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SourceControlGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the source control identified by source control name.
//
// resourceGroupName is name of an Azure Resource group. sourceControlName is the name of source control.
func (client SourceControlGroupClient) Get(ctx context.Context, resourceGroupName string, sourceControlName string) (result SourceControl, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, sourceControlName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SourceControlGroupClient) GetPreparer(ctx context.Context, resourceGroupName string, sourceControlName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", client.AutomationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sourceControlName":     autorest.Encode("path", sourceControlName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SourceControlGroupClient) GetResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of source controls.
//
// resourceGroupName is name of an Azure Resource group. filter is the filter to apply on the operation.
func (client SourceControlGroupClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, filter string) (result SourceControlListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlGroupClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.sclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.sclr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client SourceControlGroupClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", client.AutomationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlGroupClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client SourceControlGroupClient) ListByAutomationAccountResponder(resp *http.Response) (result SourceControlListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client SourceControlGroupClient) listByAutomationAccountNextResults(lastResults SourceControlListResult) (result SourceControlListResult, err error) {
	req, err := lastResults.sourceControlListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client SourceControlGroupClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, filter string) (result SourceControlListResultIterator, err error) {
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, filter)
	return
}

// Update update a source control.
//
// resourceGroupName is name of an Azure Resource group. sourceControlName is the source control name. parameters
// is the parameters supplied to the update source control operation.
func (client SourceControlGroupClient) Update(ctx context.Context, resourceGroupName string, sourceControlName string, parameters SourceControlUpdateParameters) (result SourceControl, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.SourceControlGroupClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, sourceControlName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.SourceControlGroupClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SourceControlGroupClient) UpdatePreparer(ctx context.Context, resourceGroupName string, sourceControlName string, parameters SourceControlUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", client.AutomationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sourceControlName":     autorest.Encode("path", sourceControlName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-05-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlGroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SourceControlGroupClient) UpdateResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
