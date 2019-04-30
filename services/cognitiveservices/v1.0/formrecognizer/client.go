// Package formrecognizer implements the Azure ARM Formrecognizer service API version v1.0-preview.
//
// Extracts information from forms and images into structured data based on a model created by a set of representative
// training forms.
package formrecognizer

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
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

// BaseClient is the base client for Formrecognizer.
type BaseClient struct {
	autorest.Client
	Endpoint string
}

// New creates an instance of the BaseClient client.
func New(endpoint string) BaseClient {
	return NewWithoutDefaults(endpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(endpoint string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		Endpoint: endpoint,
	}
}

// AnalyzeWithModel the document to analyze must be of a supported content type - 'application/pdf', 'image/jpeg' or
// 'image/png'. The response contains not just the extracted information of the analyzed form but also information
// about content that was not extracted along with a reason.
// Parameters:
// ID - model Identifier to analyze the document with.
// formStream - a pdf document or image (jpg,png) file to analyze.
// keys - an optional list of known keys to extract the values for.
func (client BaseClient) AnalyzeWithModel(ctx context.Context, ID uuid.UUID, formStream io.ReadCloser, keys []string) (result AnalyzeResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.AnalyzeWithModel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AnalyzeWithModelPreparer(ctx, ID, formStream, keys)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "AnalyzeWithModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.AnalyzeWithModelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "AnalyzeWithModel", resp, "Failure sending request")
		return
	}

	result, err = client.AnalyzeWithModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "AnalyzeWithModel", resp, "Failure responding to request")
	}

	return
}

// AnalyzeWithModelPreparer prepares the AnalyzeWithModel request.
func (client BaseClient) AnalyzeWithModelPreparer(ctx context.Context, ID uuid.UUID, formStream io.ReadCloser, keys []string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	queryParameters := map[string]interface{}{}
	if keys != nil && len(keys) > 0 {
		queryParameters["keys"] = autorest.Encode("query", keys, ",")
	}

	formDataParameters := map[string]interface{}{
		"form_stream": formStream,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}/analyze", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AnalyzeWithModelSender sends the AnalyzeWithModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) AnalyzeWithModelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AnalyzeWithModelResponder handles the response to the AnalyzeWithModel request. The method always
// closes the http.Response Body.
func (client BaseClient) AnalyzeWithModelResponder(resp *http.Response) (result AnalyzeResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteModel delete model artifacts.
// Parameters:
// ID - the identifier of the model to delete.
func (client BaseClient) DeleteModel(ctx context.Context, ID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DeleteModel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteModelPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "DeleteModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteModelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "DeleteModel", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "DeleteModel", resp, "Failure responding to request")
	}

	return
}

// DeleteModelPreparer prepares the DeleteModel request.
func (client BaseClient) DeleteModelPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteModelSender sends the DeleteModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DeleteModelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteModelResponder handles the response to the DeleteModel request. The method always
// closes the http.Response Body.
func (client BaseClient) DeleteModelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetExtractedKeys use the API to retrieve the keys that were
// extracted by the specified model.
// Parameters:
// ID - model identifier.
func (client BaseClient) GetExtractedKeys(ctx context.Context, ID uuid.UUID) (result KeysResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetExtractedKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetExtractedKeysPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetExtractedKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExtractedKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetExtractedKeys", resp, "Failure sending request")
		return
	}

	result, err = client.GetExtractedKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetExtractedKeys", resp, "Failure responding to request")
	}

	return
}

// GetExtractedKeysPreparer prepares the GetExtractedKeys request.
func (client BaseClient) GetExtractedKeysPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}/keys", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetExtractedKeysSender sends the GetExtractedKeys request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetExtractedKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetExtractedKeysResponder handles the response to the GetExtractedKeys request. The method always
// closes the http.Response Body.
func (client BaseClient) GetExtractedKeysResponder(resp *http.Response) (result KeysResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetModel get information about a model.
// Parameters:
// ID - model identifier.
func (client BaseClient) GetModel(ctx context.Context, ID uuid.UUID) (result ModelResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetModel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetModelPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetModelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetModel", resp, "Failure sending request")
		return
	}

	result, err = client.GetModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetModel", resp, "Failure responding to request")
	}

	return
}

// GetModelPreparer prepares the GetModel request.
func (client BaseClient) GetModelPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPathParameters("/custom/models/{id}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetModelSender sends the GetModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetModelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetModelResponder handles the response to the GetModel request. The method always
// closes the http.Response Body.
func (client BaseClient) GetModelResponder(resp *http.Response) (result ModelResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetModels get information about all trained models
func (client BaseClient) GetModels(ctx context.Context) (result ModelsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetModels")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetModelsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetModels", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetModelsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetModels", resp, "Failure sending request")
		return
	}

	result, err = client.GetModelsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "GetModels", resp, "Failure responding to request")
	}

	return
}

// GetModelsPreparer prepares the GetModels request.
func (client BaseClient) GetModelsPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPath("/custom/models"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetModelsSender sends the GetModels request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetModelsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetModelsResponder handles the response to the GetModels request. The method always
// closes the http.Response Body.
func (client BaseClient) GetModelsResponder(resp *http.Response) (result ModelsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TrainModel the train request must include a source parameter that is either an externally accessible Azure Storage
// blob container Uri (preferably a Shared Access Signature Uri) or valid path to a data folder in a locally mounted
// drive. When local paths are specified, they must follow the Linux/Unix path format and be an absolute path rooted to
// the input mount configuration
// setting value e.g., if '{Mounts:Input}' configuration setting value is '/input' then a valid source path would be
// '/input/contosodataset'. All data to be trained are expected to be under the source. Models are trained using
// documents that are of the following content type - 'application/pdf', 'image/jpeg' and 'image/png'."
// Other content is ignored when training a model.
// Parameters:
// trainRequest - request object for training.
func (client BaseClient) TrainModel(ctx context.Context, trainRequest TrainRequest) (result TrainResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.TrainModel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: trainRequest,
			Constraints: []validation.Constraint{{Target: "trainRequest.Source", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "trainRequest.Source", Name: validation.MaxLength, Rule: 2048, Chain: nil},
					{Target: "trainRequest.Source", Name: validation.MinLength, Rule: 0, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("formrecognizer.BaseClient", "TrainModel", err.Error())
	}

	req, err := client.TrainModelPreparer(ctx, trainRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "TrainModel", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainModelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "TrainModel", resp, "Failure sending request")
		return
	}

	result, err = client.TrainModelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "formrecognizer.BaseClient", "TrainModel", resp, "Failure responding to request")
	}

	return
}

// TrainModelPreparer prepares the TrainModel request.
func (client BaseClient) TrainModelPreparer(ctx context.Context, trainRequest TrainRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json-patch+json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/formrecognizer/v1.0-preview", urlParameters),
		autorest.WithPath("/custom/train"),
		autorest.WithJSON(trainRequest))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainModelSender sends the TrainModel request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) TrainModelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TrainModelResponder handles the response to the TrainModel request. The method always
// closes the http.Response Body.
func (client BaseClient) TrainModelResponder(resp *http.Response) (result TrainResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
