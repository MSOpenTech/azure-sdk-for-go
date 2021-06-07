package elastic

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

// MonitorsClient is the client for the Monitors methods of the Elastic service.
type MonitorsClient struct {
	BaseClient
}

// NewMonitorsClient creates an instance of the MonitorsClient client.
func NewMonitorsClient(subscriptionID string) MonitorsClient {
	return NewMonitorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitorsClientWithBaseURI creates an instance of the MonitorsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMonitorsClientWithBaseURI(baseURI string, subscriptionID string) MonitorsClient {
	return MonitorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
// body - elastic monitor resource model
func (client MonitorsClient) Create(ctx context.Context, resourceGroupName string, monitorName string, body *MonitorResource) (result MonitorsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Sku", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Sku.Name", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "body.Properties", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Properties.UserInfo", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.FirstName", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.FirstName", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
								{Target: "body.Properties.UserInfo.LastName", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.LastName", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
								{Target: "body.Properties.UserInfo.CompanyName", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyName", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
								{Target: "body.Properties.UserInfo.EmailAddress", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.EmailAddress", Name: validation.Pattern, Rule: `^([^<>()\[\]\.,;:\s@"]+(\.[^<>()\[\]\.,;:\s@"]+)*)@(([a-zA-Z-_0-9]+\.)+[a-zA-Z]{2,})$`, Chain: nil}}},
								{Target: "body.Properties.UserInfo.CompanyInfo", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyInfo.Domain", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyInfo.Domain", Name: validation.MaxLength, Rule: 250, Chain: nil}}},
										{Target: "body.Properties.UserInfo.CompanyInfo.Business", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyInfo.Business", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
										{Target: "body.Properties.UserInfo.CompanyInfo.EmployeesNumber", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyInfo.EmployeesNumber", Name: validation.MaxLength, Rule: 20, Chain: nil}}},
										{Target: "body.Properties.UserInfo.CompanyInfo.State", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyInfo.State", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
										{Target: "body.Properties.UserInfo.CompanyInfo.Country", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "body.Properties.UserInfo.CompanyInfo.Country", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
									}},
							}},
						}},
					{Target: "body.Location", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("elastic.MonitorsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, monitorName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MonitorsClient) CreatePreparer(ctx context.Context, resourceGroupName string, monitorName string, body *MonitorResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	body.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) CreateSender(req *http.Request) (future MonitorsCreateFuture, err error) {
	var resp *http.Response
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client MonitorsClient) CreateResponder(resp *http.Response) (result MonitorResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
func (client MonitorsClient) Delete(ctx context.Context, resourceGroupName string, monitorName string) (result MonitorsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MonitorsClient) DeletePreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) DeleteSender(req *http.Request) (future MonitorsDeleteFuture, err error) {
	var resp *http.Response
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MonitorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
func (client MonitorsClient) Get(ctx context.Context, resourceGroupName string, monitorName string) (result MonitorResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MonitorsClient) GetPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MonitorsClient) GetResponder(resp *http.Response) (result MonitorResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client MonitorsClient) List(ctx context.Context) (result MonitorResourceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.List")
		defer func() {
			sc := -1
			if result.mrlr.Response.Response != nil {
				sc = result.mrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "List", resp, "Failure sending request")
		return
	}

	result.mrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.mrlr.hasNextLink() && result.mrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client MonitorsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Elastic/monitors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MonitorsClient) ListResponder(resp *http.Response) (result MonitorResourceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MonitorsClient) listNextResults(ctx context.Context, lastResults MonitorResourceListResponse) (result MonitorResourceListResponse, err error) {
	req, err := lastResults.monitorResourceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "elastic.MonitorsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "elastic.MonitorsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MonitorsClient) ListComplete(ctx context.Context) (result MonitorResourceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.List")
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

// ListByResourceGroup sends the list by resource group request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
func (client MonitorsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result MonitorResourceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.mrlr.Response.Response != nil {
				sc = result.mrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.mrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.mrlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.mrlr.hasNextLink() && result.mrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client MonitorsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client MonitorsClient) ListByResourceGroupResponder(resp *http.Response) (result MonitorResourceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client MonitorsClient) listByResourceGroupNextResults(ctx context.Context, lastResults MonitorResourceListResponse) (result MonitorResourceListResponse, err error) {
	req, err := lastResults.monitorResourceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "elastic.MonitorsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "elastic.MonitorsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client MonitorsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result MonitorResourceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.ListByResourceGroup")
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

// Update sends the update request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
// body - elastic resource model update parameters.
func (client MonitorsClient) Update(ctx context.Context, resourceGroupName string, monitorName string, body *MonitorResourceUpdateParameters) (result MonitorResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, monitorName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.MonitorsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MonitorsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, monitorName string, body *MonitorResourceUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MonitorsClient) UpdateResponder(resp *http.Response) (result MonitorResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
