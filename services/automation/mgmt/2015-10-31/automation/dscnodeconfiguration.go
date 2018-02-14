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

// DscNodeConfigurationClient is the automation Client
type DscNodeConfigurationClient struct {
	BaseClient
}

// NewDscNodeConfigurationClient creates an instance of the DscNodeConfigurationClient client.
func NewDscNodeConfigurationClient(subscriptionID string, resourceGroupName string) DscNodeConfigurationClient {
	return NewDscNodeConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName)
}

// NewDscNodeConfigurationClientWithBaseURI creates an instance of the DscNodeConfigurationClient client.
func NewDscNodeConfigurationClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) DscNodeConfigurationClient {
	return DscNodeConfigurationClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName)}
}

// CreateOrUpdate create the node configuration identified by node configuration name.
//
// automationAccountName is the automation account name. nodeConfigurationName is the create or update parameters for
// configuration. parameters is the create or update parameters for configuration.
func (client DscNodeConfigurationClient) CreateOrUpdate(ctx context.Context, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters) (result DscNodeConfiguration, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Source", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Source.Hash", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Source.Hash.Algorithm", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Source.Hash.Value", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}},
				{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Configuration", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.DscNodeConfigurationClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(ctx, automationAccountName, nodeConfigurationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DscNodeConfigurationClient) CreateOrUpdatePreparer(ctx context.Context, automationAccountName string, nodeConfigurationName string, parameters DscNodeConfigurationCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeConfigurationName": autorest.Encode("path", nodeConfigurationName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) CreateOrUpdateResponder(resp *http.Response) (result DscNodeConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the Dsc node configurations by node configuration.
//
// automationAccountName is the automation account name. nodeConfigurationName is the Dsc node configuration name.
func (client DscNodeConfigurationClient) Delete(ctx context.Context, automationAccountName string, nodeConfigurationName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.DscNodeConfigurationClient", "Delete")
	}

	req, err := client.DeletePreparer(ctx, automationAccountName, nodeConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DscNodeConfigurationClient) DeletePreparer(ctx context.Context, automationAccountName string, nodeConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeConfigurationName": autorest.Encode("path", nodeConfigurationName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the Dsc node configurations by node configuration.
//
// automationAccountName is the automation account name. nodeConfigurationName is the Dsc node configuration name.
func (client DscNodeConfigurationClient) Get(ctx context.Context, automationAccountName string, nodeConfigurationName string) (result DscNodeConfiguration, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.DscNodeConfigurationClient", "Get")
	}

	req, err := client.GetPreparer(ctx, automationAccountName, nodeConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DscNodeConfigurationClient) GetPreparer(ctx context.Context, automationAccountName string, nodeConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeConfigurationName": autorest.Encode("path", nodeConfigurationName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) GetResponder(resp *http.Response) (result DscNodeConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of dsc node configurations.
//
// automationAccountName is the automation account name. filter is the filter to apply on the operation.
func (client DscNodeConfigurationClient) ListByAutomationAccount(ctx context.Context, automationAccountName string, filter string) (result DscNodeConfigurationListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount")
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, automationAccountName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.dnclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.dnclr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client DscNodeConfigurationClient) ListByAutomationAccountPreparer(ctx context.Context, automationAccountName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeConfigurationClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client DscNodeConfigurationClient) ListByAutomationAccountResponder(resp *http.Response) (result DscNodeConfigurationListResult, err error) {
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
func (client DscNodeConfigurationClient) listByAutomationAccountNextResults(lastResults DscNodeConfigurationListResult) (result DscNodeConfigurationListResult, err error) {
	req, err := lastResults.dscNodeConfigurationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeConfigurationClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client DscNodeConfigurationClient) ListByAutomationAccountComplete(ctx context.Context, automationAccountName string, filter string) (result DscNodeConfigurationListResultIterator, err error) {
	result.page, err = client.ListByAutomationAccount(ctx, automationAccountName, filter)
	return
}
