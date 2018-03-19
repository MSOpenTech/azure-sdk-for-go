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

// FieldsGroupClient is the automation Client
type FieldsGroupClient struct {
	BaseClient
}

// NewFieldsGroupClient creates an instance of the FieldsGroupClient client.
func NewFieldsGroupClient(subscriptionID string, automationAccountName string, automationAccountName1 string, clientRequestID string) FieldsGroupClient {
	return NewFieldsGroupClientWithBaseURI(DefaultBaseURI, subscriptionID, automationAccountName, automationAccountName1, clientRequestID)
}

// NewFieldsGroupClientWithBaseURI creates an instance of the FieldsGroupClient client.
func NewFieldsGroupClientWithBaseURI(baseURI string, subscriptionID string, automationAccountName string, automationAccountName1 string, clientRequestID string) FieldsGroupClient {
	return FieldsGroupClient{NewWithBaseURI(baseURI, subscriptionID, automationAccountName, automationAccountName1, clientRequestID)}
}

// ListByType retrieve a list of fields of a given type identified by module name.
//
// resourceGroupName is name of an Azure Resource group. automationAccountName is the automation account name.
// moduleName is the name of module. typeName is the name of type.
func (client FieldsGroupClient) ListByType(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string) (result TypeFieldListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.FieldsGroupClient", "ListByType", err.Error())
	}

	req, err := client.ListByTypePreparer(ctx, resourceGroupName, automationAccountName, moduleName, typeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.FieldsGroupClient", "ListByType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.FieldsGroupClient", "ListByType", resp, "Failure sending request")
		return
	}

	result, err = client.ListByTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.FieldsGroupClient", "ListByType", resp, "Failure responding to request")
	}

	return
}

// ListByTypePreparer prepares the ListByType request.
func (client FieldsGroupClient) ListByTypePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"moduleName":            autorest.Encode("path", moduleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"typeName":              autorest.Encode("path", typeName),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/types/{typeName}/fields", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTypeSender sends the ListByType request. The method will close the
// http.Response Body if it receives an error.
func (client FieldsGroupClient) ListByTypeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByTypeResponder handles the response to the ListByType request. The method always
// closes the http.Response Body.
func (client FieldsGroupClient) ListByTypeResponder(resp *http.Response) (result TypeFieldListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
