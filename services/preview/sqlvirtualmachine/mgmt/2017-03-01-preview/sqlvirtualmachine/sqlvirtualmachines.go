package sqlvirtualmachine

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

// SQLVirtualMachinesClient is the the SQL virtual machine management API provides a RESTful set of web APIs that
// interact with Azure Compute, Network & Storage services to manage your SQL Server virtual machine. The API enables
// users to create, delete and retrieve a SQL virtual machine, SQL virtual machine group or availability group
// listener.
type SQLVirtualMachinesClient struct {
	BaseClient
}

// NewSQLVirtualMachinesClient creates an instance of the SQLVirtualMachinesClient client.
func NewSQLVirtualMachinesClient(subscriptionID string) SQLVirtualMachinesClient {
	return NewSQLVirtualMachinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLVirtualMachinesClientWithBaseURI creates an instance of the SQLVirtualMachinesClient client.
func NewSQLVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) SQLVirtualMachinesClient {
	return SQLVirtualMachinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a SQL virtual machine.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLVirtualMachineName - name of the SQL virtual machine.
// parameters - the SQL virtual machine.
func (client SQLVirtualMachinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, parameters SQLVirtualMachine) (result SQLVirtualMachinesCreateOrUpdateFutureType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, SQLVirtualMachineName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLVirtualMachinesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, parameters SQLVirtualMachine) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlVirtualMachineName": autorest.Encode("path", SQLVirtualMachineName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) CreateOrUpdateSender(req *http.Request) (future SQLVirtualMachinesCreateOrUpdateFutureType, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) CreateOrUpdateResponder(resp *http.Response) (result SQLVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a SQL virtual machine.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLVirtualMachineName - name of the SQL virtual machine.
func (client SQLVirtualMachinesClient) Delete(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string) (result SQLVirtualMachinesDeleteFutureType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, SQLVirtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLVirtualMachinesClient) DeletePreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlVirtualMachineName": autorest.Encode("path", SQLVirtualMachineName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) DeleteSender(req *http.Request) (future SQLVirtualMachinesDeleteFutureType, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a SQL virtual machine.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLVirtualMachineName - name of the SQL virtual machine.
// expand - the child resources to include in the response.
func (client SQLVirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, expand string) (result SQLVirtualMachine, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, SQLVirtualMachineName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLVirtualMachinesClient) GetPreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlVirtualMachineName": autorest.Encode("path", SQLVirtualMachineName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) GetResponder(resp *http.Response) (result SQLVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all SQL virtual machines in a subscription.
func (client SQLVirtualMachinesClient) List(ctx context.Context) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.List")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "List", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SQLVirtualMachinesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) ListResponder(resp *http.Response) (result ListResult, err error) {
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
func (client SQLVirtualMachinesClient) listNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLVirtualMachinesClient) ListComplete(ctx context.Context) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets all SQL virtual machines in a resource group.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
func (client SQLVirtualMachinesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SQLVirtualMachinesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) ListByResourceGroupResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SQLVirtualMachinesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLVirtualMachinesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySQLVMGroup gets the list of sql virtual machines in a SQL virtual machine group.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLVirtualMachineGroupName - name of the SQL virtual machine group.
func (client SQLVirtualMachinesClient) ListBySQLVMGroup(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.ListBySQLVMGroup")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySQLVMGroupNextResults
	req, err := client.ListBySQLVMGroupPreparer(ctx, resourceGroupName, SQLVirtualMachineGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "ListBySQLVMGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySQLVMGroupSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "ListBySQLVMGroup", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListBySQLVMGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "ListBySQLVMGroup", resp, "Failure responding to request")
	}

	return
}

// ListBySQLVMGroupPreparer prepares the ListBySQLVMGroup request.
func (client SQLVirtualMachinesClient) ListBySQLVMGroupPreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"sqlVirtualMachineGroupName": autorest.Encode("path", SQLVirtualMachineGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/sqlVirtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySQLVMGroupSender sends the ListBySQLVMGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) ListBySQLVMGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListBySQLVMGroupResponder handles the response to the ListBySQLVMGroup request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) ListBySQLVMGroupResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySQLVMGroupNextResults retrieves the next set of results, if any.
func (client SQLVirtualMachinesClient) listBySQLVMGroupNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listBySQLVMGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySQLVMGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listBySQLVMGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySQLVMGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "listBySQLVMGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySQLVMGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLVirtualMachinesClient) ListBySQLVMGroupComplete(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.ListBySQLVMGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySQLVMGroup(ctx, resourceGroupName, SQLVirtualMachineGroupName)
	return
}

// Update updates a SQL virtual machine.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLVirtualMachineName - name of the SQL virtual machine.
// parameters - the SQL virtual machine.
func (client SQLVirtualMachinesClient) Update(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, parameters Update) (result SQLVirtualMachinesUpdateFutureType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLVirtualMachinesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, SQLVirtualMachineName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlvirtualmachine.SQLVirtualMachinesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SQLVirtualMachinesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, parameters Update) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlVirtualMachineName": autorest.Encode("path", SQLVirtualMachineName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SQLVirtualMachinesClient) UpdateSender(req *http.Request) (future SQLVirtualMachinesUpdateFutureType, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SQLVirtualMachinesClient) UpdateResponder(resp *http.Response) (result SQLVirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
