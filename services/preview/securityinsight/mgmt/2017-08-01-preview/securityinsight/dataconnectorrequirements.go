package securityinsight

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

// DataConnectorRequirementsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource
// provider
type DataConnectorRequirementsClient struct {
	BaseClient
}

// NewDataConnectorRequirementsClient creates an instance of the DataConnectorRequirementsClient client.
func NewDataConnectorRequirementsClient(subscriptionID string) DataConnectorRequirementsClient {
	return NewDataConnectorRequirementsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataConnectorRequirementsClientWithBaseURI creates an instance of the DataConnectorRequirementsClient client.
func NewDataConnectorRequirementsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectorRequirementsClient {
	return DataConnectorRequirementsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get requirements state for a data connector type.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// workspaceName - the name of the workspace.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// dataConnectorsCheckRequirements - the paramertes for requirements check message
func (client DataConnectorRequirementsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, operationalInsightsResourceProvider string, dataConnectorsCheckRequirements DataConnectorsCheckRequirements) (result DataConnectorRequirementsState, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectorRequirementsClient.List")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.DataConnectorRequirementsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, operationalInsightsResourceProvider, dataConnectorsCheckRequirements)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.DataConnectorRequirementsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.DataConnectorRequirementsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.DataConnectorRequirementsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DataConnectorRequirementsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, operationalInsightsResourceProvider string, dataConnectorsCheckRequirements DataConnectorsCheckRequirements) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/checkDataConnectorRequirements", pathParameters),
		autorest.WithJSON(dataConnectorsCheckRequirements),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectorRequirementsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DataConnectorRequirementsClient) ListResponder(resp *http.Response) (result DataConnectorRequirementsState, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
