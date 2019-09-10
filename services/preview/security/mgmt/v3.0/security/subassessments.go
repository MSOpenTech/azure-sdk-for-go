package security

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SubAssessmentsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type SubAssessmentsClient struct {
	BaseClient
}

// NewSubAssessmentsClient creates an instance of the SubAssessmentsClient client.
func NewSubAssessmentsClient(subscriptionID string, ascLocation string) SubAssessmentsClient {
	return NewSubAssessmentsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewSubAssessmentsClientWithBaseURI creates an instance of the SubAssessmentsClient client.
func NewSubAssessmentsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) SubAssessmentsClient {
	return SubAssessmentsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get get a security sub-assessment on your scanned resource
// Parameters:
// scope - scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or
// management group (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - the Assessment Key - Unique key for the assessment type
// subAssessmentName - the Sub-Assessment Key - Unique key for the sub-assessment type
func (client SubAssessmentsClient) Get(ctx context.Context, scope string, assessmentName string, subAssessmentName string) (result SubAssessment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubAssessmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, assessmentName, subAssessmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubAssessmentsClient) GetPreparer(ctx context.Context, scope string, assessmentName string, subAssessmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName":    autorest.Encode("path", assessmentName),
		"scope":             autorest.Encode("path", scope),
		"subAssessmentName": autorest.Encode("path", subAssessmentName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments/{subAssessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubAssessmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubAssessmentsClient) GetResponder(resp *http.Response) (result SubAssessment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get security sub-assessments on all your scanned resources inside a scope
// Parameters:
// scope - scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or
// management group (/providers/Microsoft.Management/managementGroups/mgName).
func (client SubAssessmentsClient) List(ctx context.Context, scope string) (result SubAssessmentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubAssessmentsClient.List")
		defer func() {
			sc := -1
			if result.sal.Response.Response != nil {
				sc = result.sal.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.sal, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SubAssessmentsClient) ListPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": autorest.Encode("path", scope),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Security/subAssessments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SubAssessmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SubAssessmentsClient) ListResponder(resp *http.Response) (result SubAssessmentList, err error) {
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
func (client SubAssessmentsClient) listNextResults(ctx context.Context, lastResults SubAssessmentList) (result SubAssessmentList, err error) {
	req, err := lastResults.subAssessmentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.SubAssessmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubAssessmentsClient) ListComplete(ctx context.Context, scope string) (result SubAssessmentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubAssessmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope)
	return
}
