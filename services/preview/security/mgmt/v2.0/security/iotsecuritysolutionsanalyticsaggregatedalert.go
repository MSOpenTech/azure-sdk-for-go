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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IoTSecuritySolutionsAnalyticsAggregatedAlertClient is the API spec for Microsoft.Security (Azure Security Center)
// resource provider
type IoTSecuritySolutionsAnalyticsAggregatedAlertClient struct {
	BaseClient
}

// NewIoTSecuritySolutionsAnalyticsAggregatedAlertClient creates an instance of the
// IoTSecuritySolutionsAnalyticsAggregatedAlertClient client.
func NewIoTSecuritySolutionsAnalyticsAggregatedAlertClient(subscriptionID string, ascLocation string) IoTSecuritySolutionsAnalyticsAggregatedAlertClient {
	return NewIoTSecuritySolutionsAnalyticsAggregatedAlertClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewIoTSecuritySolutionsAnalyticsAggregatedAlertClientWithBaseURI creates an instance of the
// IoTSecuritySolutionsAnalyticsAggregatedAlertClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIoTSecuritySolutionsAnalyticsAggregatedAlertClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) IoTSecuritySolutionsAnalyticsAggregatedAlertClient {
	return IoTSecuritySolutionsAnalyticsAggregatedAlertClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Dismiss security Analytics of a security solution
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the solution manager name
// aggregatedAlertName - identifier of the aggregated alert
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) Dismiss(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsAnalyticsAggregatedAlertClient.Dismiss")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Dismiss", err.Error())
	}

	req, err := client.DismissPreparer(ctx, resourceGroupName, solutionName, aggregatedAlertName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Dismiss", nil, "Failure preparing request")
		return
	}

	resp, err := client.DismissSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Dismiss", resp, "Failure sending request")
		return
	}

	result, err = client.DismissResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Dismiss", resp, "Failure responding to request")
	}

	return
}

// DismissPreparer prepares the Dismiss request.
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) DismissPreparer(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aggregatedAlertName": autorest.Encode("path", aggregatedAlertName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"solutionName":        autorest.Encode("path", solutionName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts/{aggregatedAlertName}/dismiss", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DismissSender sends the Dismiss request. The method will close the
// http.Response Body if it receives an error.
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) DismissSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DismissResponder handles the response to the Dismiss request. The method always
// closes the http.Response Body.
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) DismissResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get security Analytics of a security solution
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the solution manager name
// aggregatedAlertName - identifier of the aggregated alert
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) Get(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string) (result IoTSecurityAggregatedAlert, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsAnalyticsAggregatedAlertClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, solutionName, aggregatedAlertName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsAggregatedAlertClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) GetPreparer(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aggregatedAlertName": autorest.Encode("path", aggregatedAlertName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"solutionName":        autorest.Encode("path", solutionName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts/{aggregatedAlertName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IoTSecuritySolutionsAnalyticsAggregatedAlertClient) GetResponder(resp *http.Response) (result IoTSecurityAggregatedAlert, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
