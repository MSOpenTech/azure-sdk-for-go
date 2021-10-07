package managedservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RegistrationAssignmentsClient is the specification for ManagedServices.
type RegistrationAssignmentsClient struct {
	BaseClient
}

// NewRegistrationAssignmentsClient creates an instance of the RegistrationAssignmentsClient client.
func NewRegistrationAssignmentsClient() RegistrationAssignmentsClient {
	return NewRegistrationAssignmentsClientWithBaseURI(DefaultBaseURI)
}

// NewRegistrationAssignmentsClientWithBaseURI creates an instance of the RegistrationAssignmentsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRegistrationAssignmentsClientWithBaseURI(baseURI string) RegistrationAssignmentsClient {
	return RegistrationAssignmentsClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate creates or updates a registration assignment.
// Parameters:
// scope - scope of the resource.
// registrationAssignmentID - guid of the registration assignment.
// APIVersion - the API version to use for this operation.
// requestBody - the parameters required to create new registration assignment.
func (client RegistrationAssignmentsClient) CreateOrUpdate(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, requestBody RegistrationAssignment) (result RegistrationAssignmentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: requestBody,
			Constraints: []validation.Constraint{{Target: "requestBody.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "requestBody.Properties.RegistrationDefinitionID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "requestBody.Properties.RegistrationDefinition", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "requestBody.Properties.RegistrationDefinition.Plan", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "requestBody.Properties.RegistrationDefinition.Plan.Name", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "requestBody.Properties.RegistrationDefinition.Plan.Publisher", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "requestBody.Properties.RegistrationDefinition.Plan.Product", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "requestBody.Properties.RegistrationDefinition.Plan.Version", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("managedservices.RegistrationAssignmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, scope, registrationAssignmentID, APIVersion, requestBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegistrationAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, requestBody RegistrationAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationAssignmentId": autorest.Encode("path", registrationAssignmentID),
		"scope":                    scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	requestBody.ID = nil
	requestBody.Type = nil
	requestBody.Name = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", pathParameters),
		autorest.WithJSON(requestBody),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationAssignmentsClient) CreateOrUpdateSender(req *http.Request) (future RegistrationAssignmentsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RegistrationAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result RegistrationAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified registration assignment.
// Parameters:
// scope - scope of the resource.
// registrationAssignmentID - guid of the registration assignment.
// APIVersion - the API version to use for this operation.
func (client RegistrationAssignmentsClient) Delete(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string) (result RegistrationAssignmentsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, registrationAssignmentID, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegistrationAssignmentsClient) DeletePreparer(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationAssignmentId": autorest.Encode("path", registrationAssignmentID),
		"scope":                    scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationAssignmentsClient) DeleteSender(req *http.Request) (future RegistrationAssignmentsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegistrationAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of specified registration assignment.
// Parameters:
// scope - scope of the resource.
// registrationAssignmentID - guid of the registration assignment.
// APIVersion - the API version to use for this operation.
// expandRegistrationDefinition - tells whether to return registration definition details also along with
// registration assignment details.
func (client RegistrationAssignmentsClient) Get(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, expandRegistrationDefinition *bool) (result RegistrationAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, registrationAssignmentID, APIVersion, expandRegistrationDefinition)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegistrationAssignmentsClient) GetPreparer(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, expandRegistrationDefinition *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationAssignmentId": autorest.Encode("path", registrationAssignmentID),
		"scope":                    scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if expandRegistrationDefinition != nil {
		queryParameters["$expandRegistrationDefinition"] = autorest.Encode("query", *expandRegistrationDefinition)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegistrationAssignmentsClient) GetResponder(resp *http.Response) (result RegistrationAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of the registration assignments.
// Parameters:
// scope - scope of the resource.
// APIVersion - the API version to use for this operation.
// expandRegistrationDefinition - tells whether to return registration definition details also along with
// registration assignment details.
func (client RegistrationAssignmentsClient) List(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (result RegistrationAssignmentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.ral.Response.Response != nil {
				sc = result.ral.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, APIVersion, expandRegistrationDefinition)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ral.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.ral, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ral.hasNextLink() && result.ral.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RegistrationAssignmentsClient) ListPreparer(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if expandRegistrationDefinition != nil {
		queryParameters["$expandRegistrationDefinition"] = autorest.Encode("query", *expandRegistrationDefinition)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegistrationAssignmentsClient) ListResponder(resp *http.Response) (result RegistrationAssignmentList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RegistrationAssignmentsClient) listNextResults(ctx context.Context, lastResults RegistrationAssignmentList) (result RegistrationAssignmentList, err error) {
	req, err := lastResults.registrationAssignmentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegistrationAssignmentsClient) ListComplete(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (result RegistrationAssignmentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope, APIVersion, expandRegistrationDefinition)
	return
}
