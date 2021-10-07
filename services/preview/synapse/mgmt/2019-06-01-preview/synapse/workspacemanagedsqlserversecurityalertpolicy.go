package synapse

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

// WorkspaceManagedSQLServerSecurityAlertPolicyClient is the azure Synapse Analytics Management Client
type WorkspaceManagedSQLServerSecurityAlertPolicyClient struct {
	BaseClient
}

// NewWorkspaceManagedSQLServerSecurityAlertPolicyClient creates an instance of the
// WorkspaceManagedSQLServerSecurityAlertPolicyClient client.
func NewWorkspaceManagedSQLServerSecurityAlertPolicyClient(subscriptionID string) WorkspaceManagedSQLServerSecurityAlertPolicyClient {
	return NewWorkspaceManagedSQLServerSecurityAlertPolicyClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedSQLServerSecurityAlertPolicyClientWithBaseURI creates an instance of the
// WorkspaceManagedSQLServerSecurityAlertPolicyClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspaceManagedSQLServerSecurityAlertPolicyClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedSQLServerSecurityAlertPolicyClient {
	return WorkspaceManagedSQLServerSecurityAlertPolicyClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or Update a workspace managed sql server's threat detection policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// parameters - the workspace managed sql server security alert policy.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters ServerSecurityAlertPolicy) (result WorkspaceManagedSQLServerSecurityAlertPolicyCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerSecurityAlertPolicyClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters ServerSecurityAlertPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"securityAlertPolicyName": autorest.Encode("path", "Default"),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceName":           autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies/{securityAlertPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) CreateOrUpdateSender(req *http.Request) (future WorkspaceManagedSQLServerSecurityAlertPolicyCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
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
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) CreateOrUpdateResponder(resp *http.Response) (result ServerSecurityAlertPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a workspace managed sql server's security alert policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (result ServerSecurityAlertPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerSecurityAlertPolicyClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"securityAlertPolicyName": autorest.Encode("path", "Default"),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceName":           autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies/{securityAlertPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) GetResponder(resp *http.Response) (result ServerSecurityAlertPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get workspace managed sql server's threat detection policies.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result ServerSecurityAlertPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerSecurityAlertPolicyClient.List")
		defer func() {
			sc := -1
			if result.ssaplr.Response.Response != nil {
				sc = result.ssaplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ssaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "List", resp, "Failure sending request")
		return
	}

	result.ssaplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ssaplr.hasNextLink() && result.ssaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) ListResponder(resp *http.Response) (result ServerSecurityAlertPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) listNextResults(ctx context.Context, lastResults ServerSecurityAlertPolicyListResult) (result ServerSecurityAlertPolicyListResult, err error) {
	req, err := lastResults.serverSecurityAlertPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspaceManagedSQLServerSecurityAlertPolicyClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result ServerSecurityAlertPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerSecurityAlertPolicyClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName)
	return
}
