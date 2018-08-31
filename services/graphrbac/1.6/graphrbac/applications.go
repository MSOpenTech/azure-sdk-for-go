package graphrbac

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ApplicationsClient is the the Graph RBAC Management Client
type ApplicationsClient struct {
	BaseClient
}

// NewApplicationsClient creates an instance of the ApplicationsClient client.
func NewApplicationsClient(tenantID string) ApplicationsClient {
	return NewApplicationsClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewApplicationsClientWithBaseURI creates an instance of the ApplicationsClient client.
func NewApplicationsClientWithBaseURI(baseURI string, tenantID string) ApplicationsClient {
	return ApplicationsClient{NewWithBaseURI(baseURI, tenantID)}
}

// AddOwner add an owner to an application.
// Parameters:
// applicationObjectID - the object ID of the application to which to add the owner.
// parameters - the URL of the owner object, such as
// https://graph.windows.net/0b1f9851-1bf0-433f-aec3-cb9272f093dc/directoryObjects/f260bbc4-c254-447b-94cf-293b5ec434dd.
func (client ApplicationsClient) AddOwner(ctx context.Context, applicationObjectID string, parameters ApplicationAddOwnerParameters) (result autorest.Response, err error) {
	req, err := client.AddOwnerPreparer(ctx, applicationObjectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "AddOwner", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddOwnerSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "AddOwner", resp, "Failure sending request")
		return
	}

	result, err = client.AddOwnerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "AddOwner", resp, "Failure responding to request")
	}

	return
}

// AddOwnerPreparer prepares the AddOwner request.
func (client ApplicationsClient) AddOwnerPreparer(ctx context.Context, applicationObjectID string, parameters ApplicationAddOwnerParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/$links/owners", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddOwnerSender sends the AddOwner request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) AddOwnerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddOwnerResponder handles the response to the AddOwner request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) AddOwnerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create create a new application.
// Parameters:
// parameters - the parameters for creating an application.
func (client ApplicationsClient) Create(ctx context.Context, parameters ApplicationCreateParameters) (result Application, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AvailableToOtherTenants", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.IdentifierUris", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("graphrbac.ApplicationsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationsClient) CreatePreparer(ctx context.Context, parameters ApplicationCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) CreateResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an application.
// Parameters:
// applicationObjectID - application object ID.
func (client ApplicationsClient) Delete(ctx context.Context, applicationObjectID string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationsClient) DeletePreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an application by object ID.
// Parameters:
// applicationObjectID - application object ID.
func (client ApplicationsClient) Get(ctx context.Context, applicationObjectID string) (result Application, err error) {
	req, err := client.GetPreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationsClient) GetPreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) GetResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists applications by filter parameters.
// Parameters:
// filter - the filters to apply to the operation.
func (client ApplicationsClient) List(ctx context.Context, filter string) (result ApplicationListResultPage, err error) {
	result.fn = func(lastResult ApplicationListResult) (ApplicationListResult, error) {
		if lastResult.OdataNextLink == nil || len(to.String(lastResult.OdataNextLink)) < 1 {
			return ApplicationListResult{}, nil
		}
		return client.ListNext(ctx, *lastResult.OdataNextLink)
	}
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "List", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListResponder(resp *http.Response) (result ApplicationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationsClient) ListComplete(ctx context.Context, filter string) (result ApplicationListResultIterator, err error) {
	result.page, err = client.List(ctx, filter)
	return
}

// ListKeyCredentials get the keyCredentials associated with an application.
// Parameters:
// applicationObjectID - application object ID.
func (client ApplicationsClient) ListKeyCredentials(ctx context.Context, applicationObjectID string) (result KeyCredentialListResult, err error) {
	req, err := client.ListKeyCredentialsPreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListKeyCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeyCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListKeyCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeyCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListKeyCredentials", resp, "Failure responding to request")
	}

	return
}

// ListKeyCredentialsPreparer prepares the ListKeyCredentials request.
func (client ApplicationsClient) ListKeyCredentialsPreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/keyCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeyCredentialsSender sends the ListKeyCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListKeyCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListKeyCredentialsResponder handles the response to the ListKeyCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListKeyCredentialsResponder(resp *http.Response) (result KeyCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNext gets a list of applications from the current tenant.
// Parameters:
// nextLink - next link for the list operation.
func (client ApplicationsClient) ListNext(ctx context.Context, nextLink string) (result ApplicationListResult, err error) {
	req, err := client.ListNextPreparer(ctx, nextLink)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListNext", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListNext", resp, "Failure sending request")
		return
	}

	result, err = client.ListNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListNext", resp, "Failure responding to request")
	}

	return
}

// ListNextPreparer prepares the ListNext request.
func (client ApplicationsClient) ListNextPreparer(ctx context.Context, nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink": nextLink,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/{nextLink}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListNextSender sends the ListNext request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListNextSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListNextResponder handles the response to the ListNext request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListNextResponder(resp *http.Response) (result ApplicationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOwners the owners are a set of non-admin users who are allowed to modify this object.
// Parameters:
// applicationObjectID - the object ID of the application for which to get owners.
func (client ApplicationsClient) ListOwners(ctx context.Context, applicationObjectID string) (result DirectoryObjectListResultPage, err error) {
	result.fn = client.listOwnersNextResults
	req, err := client.ListOwnersPreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListOwners", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOwnersSender(req)
	if err != nil {
		result.dolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListOwners", resp, "Failure sending request")
		return
	}

	result.dolr, err = client.ListOwnersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListOwners", resp, "Failure responding to request")
	}

	return
}

// ListOwnersPreparer prepares the ListOwners request.
func (client ApplicationsClient) ListOwnersPreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/owners", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOwnersSender sends the ListOwners request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListOwnersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOwnersResponder handles the response to the ListOwners request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListOwnersResponder(resp *http.Response) (result DirectoryObjectListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listOwnersNextResults retrieves the next set of results, if any.
func (client ApplicationsClient) listOwnersNextResults(lastResults DirectoryObjectListResult) (result DirectoryObjectListResult, err error) {
	req, err := lastResults.directoryObjectListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "listOwnersNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListOwnersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "listOwnersNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListOwnersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "listOwnersNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListOwnersComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationsClient) ListOwnersComplete(ctx context.Context, applicationObjectID string) (result DirectoryObjectListResultIterator, err error) {
	result.page, err = client.ListOwners(ctx, applicationObjectID)
	return
}

// ListPasswordCredentials get the passwordCredentials associated with an application.
// Parameters:
// applicationObjectID - application object ID.
func (client ApplicationsClient) ListPasswordCredentials(ctx context.Context, applicationObjectID string) (result PasswordCredentialListResult, err error) {
	req, err := client.ListPasswordCredentialsPreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPasswordCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.ListPasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// ListPasswordCredentialsPreparer prepares the ListPasswordCredentials request.
func (client ApplicationsClient) ListPasswordCredentialsPreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/passwordCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPasswordCredentialsSender sends the ListPasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListPasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListPasswordCredentialsResponder handles the response to the ListPasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListPasswordCredentialsResponder(resp *http.Response) (result PasswordCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch update an existing application.
// Parameters:
// applicationObjectID - application object ID.
// parameters - parameters to update an existing application.
func (client ApplicationsClient) Patch(ctx context.Context, applicationObjectID string, parameters ApplicationUpdateParameters) (result autorest.Response, err error) {
	req, err := client.PatchPreparer(ctx, applicationObjectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client ApplicationsClient) PatchPreparer(ctx context.Context, applicationObjectID string, parameters ApplicationUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) PatchResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdateKeyCredentials update the keyCredentials associated with an application.
// Parameters:
// applicationObjectID - application object ID.
// parameters - parameters to update the keyCredentials of an existing application.
func (client ApplicationsClient) UpdateKeyCredentials(ctx context.Context, applicationObjectID string, parameters KeyCredentialsUpdateParameters) (result autorest.Response, err error) {
	req, err := client.UpdateKeyCredentialsPreparer(ctx, applicationObjectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdateKeyCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateKeyCredentialsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdateKeyCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateKeyCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdateKeyCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdateKeyCredentialsPreparer prepares the UpdateKeyCredentials request.
func (client ApplicationsClient) UpdateKeyCredentialsPreparer(ctx context.Context, applicationObjectID string, parameters KeyCredentialsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/keyCredentials", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateKeyCredentialsSender sends the UpdateKeyCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) UpdateKeyCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateKeyCredentialsResponder handles the response to the UpdateKeyCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) UpdateKeyCredentialsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdatePasswordCredentials update passwordCredentials associated with an application.
// Parameters:
// applicationObjectID - application object ID.
// parameters - parameters to update passwordCredentials of an existing application.
func (client ApplicationsClient) UpdatePasswordCredentials(ctx context.Context, applicationObjectID string, parameters PasswordCredentialsUpdateParameters) (result autorest.Response, err error) {
	req, err := client.UpdatePasswordCredentialsPreparer(ctx, applicationObjectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdatePasswordCredentialsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure sending request")
		return
	}

	result, err = client.UpdatePasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdatePasswordCredentialsPreparer prepares the UpdatePasswordCredentials request.
func (client ApplicationsClient) UpdatePasswordCredentialsPreparer(ctx context.Context, applicationObjectID string, parameters PasswordCredentialsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/passwordCredentials", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePasswordCredentialsSender sends the UpdatePasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) UpdatePasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdatePasswordCredentialsResponder handles the response to the UpdatePasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) UpdatePasswordCredentialsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
