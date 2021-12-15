package digitaltwins

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

// ModelsClient is the a service for managing and querying digital twins and digital twin models.
type ModelsClient struct {
	BaseClient
}

// NewModelsClient creates an instance of the ModelsClient client.
func NewModelsClient() ModelsClient {
	return NewModelsClientWithBaseURI(DefaultBaseURI)
}

// NewModelsClientWithBaseURI creates an instance of the ModelsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewModelsClientWithBaseURI(baseURI string) ModelsClient {
	return ModelsClient{NewWithBaseURI(baseURI)}
}

// Add uploads one or more models. When any error occurs, no models are uploaded.
// Status codes:
// * 201 Created
// * 400 Bad Request
// * DTDLParserError - The models provided are not valid DTDL.
// * InvalidArgument - The model id is invalid.
// * LimitExceeded - The maximum number of model ids allowed in 'dependenciesFor' has been reached.
// * ModelVersionNotSupported - The version of DTDL used is not supported.
// * 409 Conflict
// * ModelAlreadyExists - The model provided already exists.
// Parameters:
// models - an array of models to add.
// traceparent - identifies the request in a distributed tracing system.
// tracestate - provides vendor-specific trace identification information and is a companion to traceparent.
func (client ModelsClient) Add(ctx context.Context, models []interface{}, traceparent string, tracestate string) (result ListModelData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ModelsClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: models,
			Constraints: []validation.Constraint{{Target: "models", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "models", Name: validation.MinItems, Rule: 1, Chain: nil},
					{Target: "models", Name: validation.UniqueItems, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("digitaltwins.ModelsClient", "Add", err.Error())
	}

	req, err := client.AddPreparer(ctx, models, traceparent, tracestate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Add", resp, "Failure responding to request")
		return
	}

	return
}

// AddPreparer prepares the Add request.
func (client ModelsClient) AddPreparer(ctx context.Context, models []interface{}, traceparent string, tracestate string) (*http.Request, error) {
	const APIVersion = "2020-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/models"),
		autorest.WithQueryParameters(queryParameters))
	if models != nil && len(models) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(models))
	}
	if len(traceparent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("traceparent", autorest.String(traceparent)))
	}
	if len(tracestate) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("tracestate", autorest.String(tracestate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client ModelsClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client ModelsClient) AddResponder(resp *http.Response) (result ListModelData, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a model. A model can only be deleted if no other models reference it.
// Status codes:
// * 204 No Content
// * 400 Bad Request
// * InvalidArgument - The model id is invalid.
// * MissingArgument - The model id was not provided.
// * 404 Not Found
// * ModelNotFound - The model was not found.
// * 409 Conflict
// * ModelReferencesNotDeleted - The model refers to models that are not deleted.
// Parameters:
// ID - the id for the model. The id is globally unique and case sensitive.
// traceparent - identifies the request in a distributed tracing system.
// tracestate - provides vendor-specific trace identification information and is a companion to traceparent.
func (client ModelsClient) Delete(ctx context.Context, ID string, traceparent string, tracestate string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ModelsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, ID, traceparent, tracestate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ModelsClient) DeletePreparer(ctx context.Context, ID string, traceparent string, tracestate string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	const APIVersion = "2020-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/models/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(traceparent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("traceparent", autorest.String(traceparent)))
	}
	if len(tracestate) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("tracestate", autorest.String(tracestate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ModelsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ModelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetByID retrieves model metadata and optionally the model definition.
// Status codes:
// * 200 OK
// * 400 Bad Request
// * InvalidArgument - The model id is invalid.
// * MissingArgument - The model id was not provided.
// * 404 Not Found
// * ModelNotFound - The model was not found.
// Parameters:
// ID - the id for the model. The id is globally unique and case sensitive.
// includeModelDefinition - when true the model definition will be returned as part of the result.
// traceparent - identifies the request in a distributed tracing system.
// tracestate - provides vendor-specific trace identification information and is a companion to traceparent.
func (client ModelsClient) GetByID(ctx context.Context, ID string, includeModelDefinition *bool, traceparent string, tracestate string) (result ModelData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ModelsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, ID, includeModelDefinition, traceparent, tracestate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "GetByID", resp, "Failure responding to request")
		return
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client ModelsClient) GetByIDPreparer(ctx context.Context, ID string, includeModelDefinition *bool, traceparent string, tracestate string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	const APIVersion = "2020-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if includeModelDefinition != nil {
		queryParameters["includeModelDefinition"] = autorest.Encode("query", *includeModelDefinition)
	} else {
		queryParameters["includeModelDefinition"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/models/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(traceparent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("traceparent", autorest.String(traceparent)))
	}
	if len(tracestate) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("tracestate", autorest.String(tracestate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client ModelsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client ModelsClient) GetByIDResponder(resp *http.Response) (result ModelData, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves model metadata and, optionally, model definitions.
// Status codes:
// * 200 OK
// * 400 Bad Request
// * InvalidArgument - The model id is invalid.
// * LimitExceeded - The maximum number of model ids allowed in 'dependenciesFor' has been reached.
// * 404 Not Found
// * ModelNotFound - The model was not found.
// Parameters:
// dependenciesFor - the set of the models which will have their dependencies retrieved. If omitted, all models
// are retrieved.
// includeModelDefinition - when true the model definition will be returned as part of the result.
// maxItemsPerPage - the maximum number of items to retrieve per request. The server may choose to return less
// than the requested number.
// traceparent - identifies the request in a distributed tracing system.
// tracestate - provides vendor-specific trace identification information and is a companion to traceparent.
func (client ModelsClient) List(ctx context.Context, dependenciesFor []string, includeModelDefinition *bool, maxItemsPerPage *int32, traceparent string, tracestate string) (result PagedDigitalTwinsModelDataCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ModelsClient.List")
		defer func() {
			sc := -1
			if result.pdtmdc.Response.Response != nil {
				sc = result.pdtmdc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, dependenciesFor, includeModelDefinition, maxItemsPerPage, traceparent, tracestate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pdtmdc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "List", resp, "Failure sending request")
		return
	}

	result.pdtmdc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pdtmdc.hasNextLink() && result.pdtmdc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ModelsClient) ListPreparer(ctx context.Context, dependenciesFor []string, includeModelDefinition *bool, maxItemsPerPage *int32, traceparent string, tracestate string) (*http.Request, error) {
	const APIVersion = "2020-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if dependenciesFor != nil && len(dependenciesFor) > 0 {
		queryParameters["dependenciesFor"] = dependenciesFor
	}
	if includeModelDefinition != nil {
		queryParameters["includeModelDefinition"] = autorest.Encode("query", *includeModelDefinition)
	} else {
		queryParameters["includeModelDefinition"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/models"),
		autorest.WithQueryParameters(queryParameters))
	if maxItemsPerPage != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("max-items-per-page", autorest.String(*maxItemsPerPage)))
	}
	if len(traceparent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("traceparent", autorest.String(traceparent)))
	}
	if len(tracestate) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("tracestate", autorest.String(tracestate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ModelsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ModelsClient) ListResponder(resp *http.Response) (result PagedDigitalTwinsModelDataCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ModelsClient) listNextResults(ctx context.Context, lastResults PagedDigitalTwinsModelDataCollection) (result PagedDigitalTwinsModelDataCollection, err error) {
	req, err := lastResults.pagedDigitalTwinsModelDataCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ModelsClient) ListComplete(ctx context.Context, dependenciesFor []string, includeModelDefinition *bool, maxItemsPerPage *int32, traceparent string, tracestate string) (result PagedDigitalTwinsModelDataCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ModelsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, dependenciesFor, includeModelDefinition, maxItemsPerPage, traceparent, tracestate)
	return
}

// Update updates the metadata for a model.
// Status codes:
// * 204 No Content
// * 400 Bad Request
// * InvalidArgument - The model id is invalid.
// * JsonPatchInvalid - The JSON Patch provided is invalid.
// * MissingArgument - The model id was not provided.
// * 404 Not Found
// * ModelNotFound - The model was not found.
// * 409 Conflict
// * ModelReferencesNotDecommissioned - The model refers to models that are not decommissioned.
// Parameters:
// ID - the id for the model. The id is globally unique and case sensitive.
// updateModel - an update specification described by JSON Patch. Only the decommissioned property can be
// replaced.
// traceparent - identifies the request in a distributed tracing system.
// tracestate - provides vendor-specific trace identification information and is a companion to traceparent.
func (client ModelsClient) Update(ctx context.Context, ID string, updateModel []interface{}, traceparent string, tracestate string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ModelsClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: updateModel,
			Constraints: []validation.Constraint{{Target: "updateModel", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("digitaltwins.ModelsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, ID, updateModel, traceparent, tracestate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.ModelsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ModelsClient) UpdatePreparer(ctx context.Context, ID string, updateModel []interface{}, traceparent string, tracestate string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	const APIVersion = "2020-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json-patch+json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/models/{id}", pathParameters),
		autorest.WithJSON(updateModel),
		autorest.WithQueryParameters(queryParameters))
	if len(traceparent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("traceparent", autorest.String(traceparent)))
	}
	if len(tracestate) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("tracestate", autorest.String(tracestate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ModelsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ModelsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
