// Package prediction implements the Azure ARM Prediction service API version 3.0.
//
//
package prediction

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
	"github.com/gofrs/uuid"
	"io"
	"net/http"
)

// BaseClient is the base client for Prediction.
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

// ClassifyImage sends the classify image request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageData - binary image data. Supported formats are JPEG, GIF, PNG, and BMP. Supports images up to 4MB.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) ClassifyImage(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ClassifyImage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ClassifyImagePreparer(ctx, projectID, publishedName, imageData, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClassifyImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImage", resp, "Failure sending request")
		return
	}

	result, err = client.ClassifyImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImage", resp, "Failure responding to request")
		return
	}

	return
}

// ClassifyImagePreparer prepares the ClassifyImage request.
func (client BaseClient) ClassifyImagePreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	formDataParameters := map[string]interface{}{
		"imageData": imageData,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/classify/iterations/{publishedName}/image", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClassifyImageSender sends the ClassifyImage request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ClassifyImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ClassifyImageResponder handles the response to the ClassifyImage request. The method always
// closes the http.Response Body.
func (client BaseClient) ClassifyImageResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ClassifyImageURL sends the classify image url request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageURL - an ImageUrl that contains the url of the image to be evaluated.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) ClassifyImageURL(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ClassifyImageURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("prediction.BaseClient", "ClassifyImageURL", err.Error())
	}

	req, err := client.ClassifyImageURLPreparer(ctx, projectID, publishedName, imageURL, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClassifyImageURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageURL", resp, "Failure sending request")
		return
	}

	result, err = client.ClassifyImageURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageURL", resp, "Failure responding to request")
		return
	}

	return
}

// ClassifyImageURLPreparer prepares the ClassifyImageURL request.
func (client BaseClient) ClassifyImageURLPreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/classify/iterations/{publishedName}/url", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClassifyImageURLSender sends the ClassifyImageURL request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ClassifyImageURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ClassifyImageURLResponder handles the response to the ClassifyImageURL request. The method always
// closes the http.Response Body.
func (client BaseClient) ClassifyImageURLResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ClassifyImageURLWithNoStore sends the classify image url with no store request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageURL - an {Iris.Web.Api.Models.ImageUrl} that contains the url of the image to be evaluated.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) ClassifyImageURLWithNoStore(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ClassifyImageURLWithNoStore")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("prediction.BaseClient", "ClassifyImageURLWithNoStore", err.Error())
	}

	req, err := client.ClassifyImageURLWithNoStorePreparer(ctx, projectID, publishedName, imageURL, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageURLWithNoStore", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClassifyImageURLWithNoStoreSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageURLWithNoStore", resp, "Failure sending request")
		return
	}

	result, err = client.ClassifyImageURLWithNoStoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageURLWithNoStore", resp, "Failure responding to request")
		return
	}

	return
}

// ClassifyImageURLWithNoStorePreparer prepares the ClassifyImageURLWithNoStore request.
func (client BaseClient) ClassifyImageURLWithNoStorePreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/classify/iterations/{publishedName}/url/nostore", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClassifyImageURLWithNoStoreSender sends the ClassifyImageURLWithNoStore request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ClassifyImageURLWithNoStoreSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ClassifyImageURLWithNoStoreResponder handles the response to the ClassifyImageURLWithNoStore request. The method always
// closes the http.Response Body.
func (client BaseClient) ClassifyImageURLWithNoStoreResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ClassifyImageWithNoStore sends the classify image with no store request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageData - binary image data. Supported formats are JPEG, GIF, PNG, and BMP. Supports images up to 0MB.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) ClassifyImageWithNoStore(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ClassifyImageWithNoStore")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ClassifyImageWithNoStorePreparer(ctx, projectID, publishedName, imageData, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageWithNoStore", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClassifyImageWithNoStoreSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageWithNoStore", resp, "Failure sending request")
		return
	}

	result, err = client.ClassifyImageWithNoStoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "ClassifyImageWithNoStore", resp, "Failure responding to request")
		return
	}

	return
}

// ClassifyImageWithNoStorePreparer prepares the ClassifyImageWithNoStore request.
func (client BaseClient) ClassifyImageWithNoStorePreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	formDataParameters := map[string]interface{}{
		"imageData": imageData,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/classify/iterations/{publishedName}/image/nostore", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClassifyImageWithNoStoreSender sends the ClassifyImageWithNoStore request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ClassifyImageWithNoStoreSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ClassifyImageWithNoStoreResponder handles the response to the ClassifyImageWithNoStore request. The method always
// closes the http.Response Body.
func (client BaseClient) ClassifyImageWithNoStoreResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectImage sends the detect image request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageData - binary image data. Supported formats are JPEG, GIF, PNG, and BMP. Supports images up to 4MB.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) DetectImage(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DetectImage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetectImagePreparer(ctx, projectID, publishedName, imageData, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImage", resp, "Failure sending request")
		return
	}

	result, err = client.DetectImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImage", resp, "Failure responding to request")
		return
	}

	return
}

// DetectImagePreparer prepares the DetectImage request.
func (client BaseClient) DetectImagePreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	formDataParameters := map[string]interface{}{
		"imageData": imageData,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/detect/iterations/{publishedName}/image", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectImageSender sends the DetectImage request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectImageResponder handles the response to the DetectImage request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectImageResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectImageURL sends the detect image url request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageURL - an ImageUrl that contains the url of the image to be evaluated.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) DetectImageURL(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DetectImageURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("prediction.BaseClient", "DetectImageURL", err.Error())
	}

	req, err := client.DetectImageURLPreparer(ctx, projectID, publishedName, imageURL, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectImageURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageURL", resp, "Failure sending request")
		return
	}

	result, err = client.DetectImageURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageURL", resp, "Failure responding to request")
		return
	}

	return
}

// DetectImageURLPreparer prepares the DetectImageURL request.
func (client BaseClient) DetectImageURLPreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/detect/iterations/{publishedName}/url", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectImageURLSender sends the DetectImageURL request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectImageURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectImageURLResponder handles the response to the DetectImageURL request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectImageURLResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectImageURLWithNoStore sends the detect image url with no store request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageURL - an {Iris.Web.Api.Models.ImageUrl} that contains the url of the image to be evaluated.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) DetectImageURLWithNoStore(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DetectImageURLWithNoStore")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("prediction.BaseClient", "DetectImageURLWithNoStore", err.Error())
	}

	req, err := client.DetectImageURLWithNoStorePreparer(ctx, projectID, publishedName, imageURL, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageURLWithNoStore", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectImageURLWithNoStoreSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageURLWithNoStore", resp, "Failure sending request")
		return
	}

	result, err = client.DetectImageURLWithNoStoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageURLWithNoStore", resp, "Failure responding to request")
		return
	}

	return
}

// DetectImageURLWithNoStorePreparer prepares the DetectImageURLWithNoStore request.
func (client BaseClient) DetectImageURLWithNoStorePreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageURL ImageURL, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/detect/iterations/{publishedName}/url/nostore", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectImageURLWithNoStoreSender sends the DetectImageURLWithNoStore request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectImageURLWithNoStoreSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectImageURLWithNoStoreResponder handles the response to the DetectImageURLWithNoStore request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectImageURLWithNoStoreResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectImageWithNoStore sends the detect image with no store request.
// Parameters:
// projectID - the project id.
// publishedName - specifies the name of the model to evaluate against.
// imageData - binary image data. Supported formats are JPEG, GIF, PNG, and BMP. Supports images up to 0MB.
// application - optional. Specifies the name of application using the endpoint.
func (client BaseClient) DetectImageWithNoStore(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (result ImagePrediction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DetectImageWithNoStore")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetectImageWithNoStorePreparer(ctx, projectID, publishedName, imageData, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageWithNoStore", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectImageWithNoStoreSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageWithNoStore", resp, "Failure sending request")
		return
	}

	result, err = client.DetectImageWithNoStoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "prediction.BaseClient", "DetectImageWithNoStore", resp, "Failure responding to request")
		return
	}

	return
}

// DetectImageWithNoStorePreparer prepares the DetectImageWithNoStore request.
func (client BaseClient) DetectImageWithNoStorePreparer(ctx context.Context, projectID uuid.UUID, publishedName string, imageData io.ReadCloser, application string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"projectId":     autorest.Encode("path", projectID),
		"publishedName": autorest.Encode("path", publishedName),
	}

	queryParameters := map[string]interface{}{}
	if len(application) > 0 {
		queryParameters["application"] = autorest.Encode("query", application)
	}

	formDataParameters := map[string]interface{}{
		"imageData": imageData,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/customvision/v3.0/prediction", urlParameters),
		autorest.WithPathParameters("/{projectId}/detect/iterations/{publishedName}/image/nostore", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectImageWithNoStoreSender sends the DetectImageWithNoStore request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectImageWithNoStoreSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectImageWithNoStoreResponder handles the response to the DetectImageWithNoStore request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectImageWithNoStoreResponder(resp *http.Response) (result ImagePrediction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
