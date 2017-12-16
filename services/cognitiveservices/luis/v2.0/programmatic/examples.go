package programmatic

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
	"github.com/satori/go.uuid"
	"net/http"
)

// ExamplesClient is the client for the Examples methods of the Programmatic service.
type ExamplesClient struct {
	BaseClient
}

// NewExamplesClient creates an instance of the ExamplesClient client.
func NewExamplesClient(azureRegion AzureRegions) ExamplesClient {
	return ExamplesClient{New(azureRegion)}
}

// Add adds a labeled example to the application.
//
// appID is the application ID. versionID is the version ID. exampleLabelObject is an example label with the expected
// intent and entities.
func (client ExamplesClient) Add(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObject ExampleLabelObject) (result LabelExampleResponse, err error) {
	req, err := client.AddPreparer(ctx, appID, versionID, exampleLabelObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client ExamplesClient) AddPreparer(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObject ExampleLabelObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/example", pathParameters),
		autorest.WithJSON(exampleLabelObject))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) AddSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client ExamplesClient) AddResponder(resp *http.Response) (result LabelExampleResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Batch adds a batch of labeled examples to the application.
//
// appID is the application ID. versionID is the version ID. exampleLabelObjectArray is array of examples.
func (client ExamplesClient) Batch(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObjectArray []ExampleLabelObject) (result ListBatchLabelExample, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: exampleLabelObjectArray,
			Constraints: []validation.Constraint{{Target: "exampleLabelObjectArray", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "programmatic.ExamplesClient", "Batch")
	}

	req, err := client.BatchPreparer(ctx, appID, versionID, exampleLabelObjectArray)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Batch", nil, "Failure preparing request")
		return
	}

	resp, err := client.BatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Batch", resp, "Failure sending request")
		return
	}

	result, err = client.BatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Batch", resp, "Failure responding to request")
	}

	return
}

// BatchPreparer prepares the Batch request.
func (client ExamplesClient) BatchPreparer(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObjectArray []ExampleLabelObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/examples", pathParameters),
		autorest.WithJSON(exampleLabelObjectArray))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BatchSender sends the Batch request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) BatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// BatchResponder handles the response to the Batch request. The method always
// closes the http.Response Body.
func (client ExamplesClient) BatchResponder(resp *http.Response) (result ListBatchLabelExample, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the labeled example with the specified ID.
//
// appID is the application ID. versionID is the version ID. exampleID is the example ID.
func (client ExamplesClient) Delete(ctx context.Context, appID uuid.UUID, versionID string, exampleID int32) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, appID, versionID, exampleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExamplesClient) DeletePreparer(ctx context.Context, appID uuid.UUID, versionID string, exampleID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"exampleId": autorest.Encode("path", exampleID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/examples/{exampleId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExamplesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List returns examples to be reviewed.
//
// appID is the application ID. versionID is the version ID. skip is the number of entries to skip. Default value is 0.
// take is the number of entries to return. Maximum page size is 500. Default is 100.
func (client ExamplesClient) List(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result ListLabeledUtterance, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}},
		{TargetValue: take,
			Constraints: []validation.Constraint{{Target: "take", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "take", Name: validation.InclusiveMaximum, Rule: 500, Chain: nil},
					{Target: "take", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "programmatic.ExamplesClient", "List")
	}

	req, err := client.ListPreparer(ctx, appID, versionID, skip, take)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.ExamplesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExamplesClient) ListPreparer(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if skip != nil {
		queryParameters["skip"] = autorest.Encode("query", *skip)
	}
	if take != nil {
		queryParameters["take"] = autorest.Encode("query", *take)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/examples", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExamplesClient) ListResponder(resp *http.Response) (result ListLabeledUtterance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
