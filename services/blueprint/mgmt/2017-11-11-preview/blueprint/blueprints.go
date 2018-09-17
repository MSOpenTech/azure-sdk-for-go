package blueprint

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

// BlueprintsClient is the blueprint Client
type BlueprintsClient struct {
	BaseClient
}

// NewBlueprintsClient creates an instance of the BlueprintsClient client.
func NewBlueprintsClient() BlueprintsClient {
	return NewBlueprintsClientWithBaseURI(DefaultBaseURI)
}

// NewBlueprintsClientWithBaseURI creates an instance of the BlueprintsClient client.
func NewBlueprintsClientWithBaseURI(baseURI string) BlueprintsClient {
	return BlueprintsClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate create or update Blueprint definition.
// Parameters:
// managementGroupName - managementGroup where blueprint stores.
// blueprintName - name of the blueprint.
// blueprint - blueprint definition.
func (client BlueprintsClient) CreateOrUpdate(ctx context.Context, managementGroupName string, blueprintName string, blueprint Model) (result Model, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: blueprint,
			Constraints: []validation.Constraint{{Target: "blueprint.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("blueprint.BlueprintsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, managementGroupName, blueprintName, blueprint)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BlueprintsClient) CreateOrUpdatePreparer(ctx context.Context, managementGroupName string, blueprintName string, blueprint Model) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName":             autorest.Encode("path", blueprintName),
		"managementGroupName":       autorest.Encode("path", managementGroupName),
		"managementGroupsNamespace": autorest.Encode("path", "Microsoft.Management"),
	}

	const APIVersion = "2017-11-11-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupName}/providers/Microsoft.Blueprint/Blueprints/{blueprintName}", pathParameters),
		autorest.WithJSON(blueprint),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BlueprintsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BlueprintsClient) CreateOrUpdateResponder(resp *http.Response) (result Model, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a blueprint definition.
// Parameters:
// managementGroupName - managementGroup where blueprint stores.
// blueprintName - name of the blueprint.
func (client BlueprintsClient) Delete(ctx context.Context, managementGroupName string, blueprintName string) (result Model, err error) {
	req, err := client.DeletePreparer(ctx, managementGroupName, blueprintName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BlueprintsClient) DeletePreparer(ctx context.Context, managementGroupName string, blueprintName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName":             autorest.Encode("path", blueprintName),
		"managementGroupName":       autorest.Encode("path", managementGroupName),
		"managementGroupsNamespace": autorest.Encode("path", "Microsoft.Management"),
	}

	const APIVersion = "2017-11-11-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupName}/providers/Microsoft.Blueprint/Blueprints/{blueprintName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BlueprintsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BlueprintsClient) DeleteResponder(resp *http.Response) (result Model, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a blueprint definition.
// Parameters:
// managementGroupName - managementGroup where blueprint stores.
// blueprintName - name of the blueprint.
func (client BlueprintsClient) Get(ctx context.Context, managementGroupName string, blueprintName string) (result Model, err error) {
	req, err := client.GetPreparer(ctx, managementGroupName, blueprintName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BlueprintsClient) GetPreparer(ctx context.Context, managementGroupName string, blueprintName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName":             autorest.Encode("path", blueprintName),
		"managementGroupName":       autorest.Encode("path", managementGroupName),
		"managementGroupsNamespace": autorest.Encode("path", "Microsoft.Management"),
	}

	const APIVersion = "2017-11-11-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupName}/providers/Microsoft.Blueprint/Blueprints/{blueprintName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BlueprintsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BlueprintsClient) GetResponder(resp *http.Response) (result Model, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List create or update blueprint definition.
// Parameters:
// managementGroupName - managementGroup where blueprint stores.
func (client BlueprintsClient) List(ctx context.Context, managementGroupName string) (result ListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, managementGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.l.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "List", resp, "Failure sending request")
		return
	}

	result.l, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BlueprintsClient) ListPreparer(ctx context.Context, managementGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupName":       autorest.Encode("path", managementGroupName),
		"managementGroupsNamespace": autorest.Encode("path", "Microsoft.Management"),
	}

	const APIVersion = "2017-11-11-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupName}/providers/Microsoft.Blueprint/Blueprints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BlueprintsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BlueprintsClient) ListResponder(resp *http.Response) (result List, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BlueprintsClient) listNextResults(lastResults List) (result List, err error) {
	req, err := lastResults.listPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.BlueprintsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BlueprintsClient) ListComplete(ctx context.Context, managementGroupName string) (result ListIterator, err error) {
	result.page, err = client.List(ctx, managementGroupName)
	return
}
