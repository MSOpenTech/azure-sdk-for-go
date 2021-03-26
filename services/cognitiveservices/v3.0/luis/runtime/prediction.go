package runtime

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
	"github.com/gofrs/uuid"
	"net/http"
)

// PredictionClient is the client for the Prediction methods of the Runtime service.
type PredictionClient struct {
	BaseClient
}

// NewPredictionClient creates an instance of the PredictionClient client.
func NewPredictionClient(endpoint string) PredictionClient {
	return PredictionClient{New(endpoint)}
}

// GetSlotPrediction gets the predictions for an application slot.
// Parameters:
// appID - the application ID.
// slotName - the application slot name.
// predictionRequest - the prediction request parameters.
// verbose - indicates whether to get extra metadata for the entities predictions or not.
// showAllIntents - indicates whether to return all the intents in the response or just the top intent.
// logParameter - indicates whether to log the endpoint query or not.
func (client PredictionClient) GetSlotPrediction(ctx context.Context, appID uuid.UUID, slotName string, predictionRequest PredictionRequest, verbose *bool, showAllIntents *bool, logParameter *bool) (result PredictionResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionClient.GetSlotPrediction")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: predictionRequest,
			Constraints: []validation.Constraint{{Target: "predictionRequest.Query", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("runtime.PredictionClient", "GetSlotPrediction", err.Error())
	}

	req, err := client.GetSlotPredictionPreparer(ctx, appID, slotName, predictionRequest, verbose, showAllIntents, logParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "GetSlotPrediction", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSlotPredictionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "GetSlotPrediction", resp, "Failure sending request")
		return
	}

	result, err = client.GetSlotPredictionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "GetSlotPrediction", resp, "Failure responding to request")
		return
	}

	return
}

// GetSlotPredictionPreparer prepares the GetSlotPrediction request.
func (client PredictionClient) GetSlotPredictionPreparer(ctx context.Context, appID uuid.UUID, slotName string, predictionRequest PredictionRequest, verbose *bool, showAllIntents *bool, logParameter *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":    autorest.Encode("path", appID),
		"slotName": autorest.Encode("path", slotName),
	}

	queryParameters := map[string]interface{}{}
	if verbose != nil {
		queryParameters["verbose"] = autorest.Encode("query", *verbose)
	}
	if showAllIntents != nil {
		queryParameters["show-all-intents"] = autorest.Encode("query", *showAllIntents)
	}
	if logParameter != nil {
		queryParameters["log"] = autorest.Encode("query", *logParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/prediction/v3.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/slots/{slotName}/predict", pathParameters),
		autorest.WithJSON(predictionRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSlotPredictionSender sends the GetSlotPrediction request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionClient) GetSlotPredictionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSlotPredictionResponder handles the response to the GetSlotPrediction request. The method always
// closes the http.Response Body.
func (client PredictionClient) GetSlotPredictionResponder(resp *http.Response) (result PredictionResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetVersionPrediction gets the predictions for an application version.
// Parameters:
// appID - the application ID.
// versionID - the application version ID.
// predictionRequest - the prediction request parameters.
// verbose - indicates whether to get extra metadata for the entities predictions or not.
// showAllIntents - indicates whether to return all the intents in the response or just the top intent.
// logParameter - indicates whether to log the endpoint query or not.
func (client PredictionClient) GetVersionPrediction(ctx context.Context, appID uuid.UUID, versionID string, predictionRequest PredictionRequest, verbose *bool, showAllIntents *bool, logParameter *bool) (result PredictionResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionClient.GetVersionPrediction")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: predictionRequest,
			Constraints: []validation.Constraint{{Target: "predictionRequest.Query", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("runtime.PredictionClient", "GetVersionPrediction", err.Error())
	}

	req, err := client.GetVersionPredictionPreparer(ctx, appID, versionID, predictionRequest, verbose, showAllIntents, logParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "GetVersionPrediction", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetVersionPredictionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "GetVersionPrediction", resp, "Failure sending request")
		return
	}

	result, err = client.GetVersionPredictionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runtime.PredictionClient", "GetVersionPrediction", resp, "Failure responding to request")
		return
	}

	return
}

// GetVersionPredictionPreparer prepares the GetVersionPrediction request.
func (client PredictionClient) GetVersionPredictionPreparer(ctx context.Context, appID uuid.UUID, versionID string, predictionRequest PredictionRequest, verbose *bool, showAllIntents *bool, logParameter *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if verbose != nil {
		queryParameters["verbose"] = autorest.Encode("query", *verbose)
	}
	if showAllIntents != nil {
		queryParameters["show-all-intents"] = autorest.Encode("query", *showAllIntents)
	}
	if logParameter != nil {
		queryParameters["log"] = autorest.Encode("query", *logParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/prediction/v3.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/predict", pathParameters),
		autorest.WithJSON(predictionRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetVersionPredictionSender sends the GetVersionPrediction request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionClient) GetVersionPredictionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetVersionPredictionResponder handles the response to the GetVersionPrediction request. The method always
// closes the http.Response Body.
func (client PredictionClient) GetVersionPredictionResponder(resp *http.Response) (result PredictionResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
