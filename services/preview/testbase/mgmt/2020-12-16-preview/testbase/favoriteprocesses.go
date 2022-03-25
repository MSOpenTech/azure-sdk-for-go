package testbase

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

// FavoriteProcessesClient is the test Base
type FavoriteProcessesClient struct {
	BaseClient
}

// NewFavoriteProcessesClient creates an instance of the FavoriteProcessesClient client.
func NewFavoriteProcessesClient(subscriptionID string) FavoriteProcessesClient {
	return NewFavoriteProcessesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFavoriteProcessesClientWithBaseURI creates an instance of the FavoriteProcessesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewFavoriteProcessesClientWithBaseURI(baseURI string, subscriptionID string) FavoriteProcessesClient {
	return FavoriteProcessesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create or replace a favorite process for a Test Base Package.
// Parameters:
// parameters - parameters supplied to create a favorite process in a package.
// resourceGroupName - the name of the resource group that contains the resource.
// testBaseAccountName - the resource name of the Test Base Account.
// packageName - the resource name of the Test Base Package.
// favoriteProcessResourceName - the resource name of a favorite process in a package. If the process name
// contains characters that are not allowed in Azure Resource Name, we use 'actualProcessName' in request body
// to submit the name.
func (client FavoriteProcessesClient) Create(ctx context.Context, parameters FavoriteProcessResource, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (result FavoriteProcessResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FavoriteProcessesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FavoriteProcessProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.FavoriteProcessProperties.ActualProcessName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("testbase.FavoriteProcessesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, parameters, resourceGroupName, testBaseAccountName, packageName, favoriteProcessResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client FavoriteProcessesClient) CreatePreparer(ctx context.Context, parameters FavoriteProcessResource, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteProcessResourceName": autorest.Encode("path", favoriteProcessResourceName),
		"packageName":                 autorest.Encode("path", packageName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
		"testBaseAccountName":         autorest.Encode("path", testBaseAccountName),
	}

	const APIVersion = "2020-12-16-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/favoriteProcesses/{favoriteProcessResourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteProcessesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client FavoriteProcessesClient) CreateResponder(resp *http.Response) (result FavoriteProcessResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a favorite process for a specific package.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource.
// testBaseAccountName - the resource name of the Test Base Account.
// packageName - the resource name of the Test Base Package.
// favoriteProcessResourceName - the resource name of a favorite process in a package. If the process name
// contains characters that are not allowed in Azure Resource Name, we use 'actualProcessName' in request body
// to submit the name.
func (client FavoriteProcessesClient) Delete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FavoriteProcessesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, testBaseAccountName, packageName, favoriteProcessResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FavoriteProcessesClient) DeletePreparer(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteProcessResourceName": autorest.Encode("path", favoriteProcessResourceName),
		"packageName":                 autorest.Encode("path", packageName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
		"testBaseAccountName":         autorest.Encode("path", testBaseAccountName),
	}

	const APIVersion = "2020-12-16-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/favoriteProcesses/{favoriteProcessResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteProcessesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FavoriteProcessesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a favorite process for a Test Base Package.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource.
// testBaseAccountName - the resource name of the Test Base Account.
// packageName - the resource name of the Test Base Package.
// favoriteProcessResourceName - the resource name of a favorite process in a package. If the process name
// contains characters that are not allowed in Azure Resource Name, we use 'actualProcessName' in request body
// to submit the name.
func (client FavoriteProcessesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (result FavoriteProcessResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FavoriteProcessesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, testBaseAccountName, packageName, favoriteProcessResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FavoriteProcessesClient) GetPreparer(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, favoriteProcessResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteProcessResourceName": autorest.Encode("path", favoriteProcessResourceName),
		"packageName":                 autorest.Encode("path", packageName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
		"testBaseAccountName":         autorest.Encode("path", testBaseAccountName),
	}

	const APIVersion = "2020-12-16-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/favoriteProcesses/{favoriteProcessResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteProcessesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FavoriteProcessesClient) GetResponder(resp *http.Response) (result FavoriteProcessResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the favorite processes for a specific package.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource.
// testBaseAccountName - the resource name of the Test Base Account.
// packageName - the resource name of the Test Base Package.
func (client FavoriteProcessesClient) List(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result FavoriteProcessListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FavoriteProcessesClient.List")
		defer func() {
			sc := -1
			if result.fplr.Response.Response != nil {
				sc = result.fplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, testBaseAccountName, packageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.fplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "List", resp, "Failure sending request")
		return
	}

	result.fplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.fplr.hasNextLink() && result.fplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client FavoriteProcessesClient) ListPreparer(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"packageName":         autorest.Encode("path", packageName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"testBaseAccountName": autorest.Encode("path", testBaseAccountName),
	}

	const APIVersion = "2020-12-16-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/favoriteProcesses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteProcessesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FavoriteProcessesClient) ListResponder(resp *http.Response) (result FavoriteProcessListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client FavoriteProcessesClient) listNextResults(ctx context.Context, lastResults FavoriteProcessListResult) (result FavoriteProcessListResult, err error) {
	req, err := lastResults.favoriteProcessListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.FavoriteProcessesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FavoriteProcessesClient) ListComplete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string) (result FavoriteProcessListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FavoriteProcessesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, testBaseAccountName, packageName)
	return
}
