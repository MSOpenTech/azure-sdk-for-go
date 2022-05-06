package hardwaresecuritymodules

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

// DedicatedHsmClient is the the Azure management API provides a RESTful set of web services that interact with Azure
// Dedicated HSM RP.
type DedicatedHsmClient struct {
	BaseClient
}

// NewDedicatedHsmClient creates an instance of the DedicatedHsmClient client.
func NewDedicatedHsmClient(subscriptionID string) DedicatedHsmClient {
	return NewDedicatedHsmClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDedicatedHsmClientWithBaseURI creates an instance of the DedicatedHsmClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDedicatedHsmClientWithBaseURI(baseURI string, subscriptionID string) DedicatedHsmClient {
	return DedicatedHsmClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or Update a dedicated HSM in the specified subscription.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the resource belongs.
// name - name of the dedicated Hsm
// parameters - parameters to create or update the dedicated hsm
func (client DedicatedHsmClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm) (result DedicatedHsmCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: name,
			Constraints: []validation.Constraint{{Target: "name", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DedicatedHsmProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hardwaresecuritymodules.DedicatedHsmClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DedicatedHsmClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsm) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) CreateOrUpdateSender(req *http.Request) (future DedicatedHsmCreateOrUpdateFuture, err error) {
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
func (client DedicatedHsmClient) CreateOrUpdateResponder(resp *http.Response) (result DedicatedHsm, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Azure Dedicated HSM.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the dedicated HSM belongs.
// name - the name of the dedicated HSM to delete
func (client DedicatedHsmClient) Delete(ctx context.Context, resourceGroupName string, name string) (result DedicatedHsmDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DedicatedHsmClient) DeletePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) DeleteSender(req *http.Request) (future DedicatedHsmDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DedicatedHsmClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Azure dedicated HSM.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the dedicated hsm belongs.
// name - the name of the dedicated HSM.
func (client DedicatedHsmClient) Get(ctx context.Context, resourceGroupName string, name string) (result DedicatedHsm, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DedicatedHsmClient) GetPreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DedicatedHsmClient) GetResponder(resp *http.Response) (result DedicatedHsm, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup the List operation gets information about the dedicated hsms associated with the subscription
// and within the specified resource group.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the dedicated HSM belongs.
// top - maximum number of results to return.
func (client DedicatedHsmClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result DedicatedHsmListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.dhlr.Response.Response != nil {
				sc = result.dhlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.dhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.dhlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.dhlr.hasNextLink() && result.dhlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DedicatedHsmClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DedicatedHsmClient) ListByResourceGroupResponder(resp *http.Response) (result DedicatedHsmListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DedicatedHsmClient) listByResourceGroupNextResults(ctx context.Context, lastResults DedicatedHsmListResult) (result DedicatedHsmListResult, err error) {
	req, err := lastResults.dedicatedHsmListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedHsmClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result DedicatedHsmListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, top)
	return
}

// ListBySubscription the List operation gets information about the dedicated HSMs associated with the subscription.
// Parameters:
// top - maximum number of results to return.
func (client DedicatedHsmClient) ListBySubscription(ctx context.Context, top *int32) (result DedicatedHsmListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.dhlr.Response.Response != nil {
				sc = result.dhlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.dhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.dhlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.dhlr.hasNextLink() && result.dhlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client DedicatedHsmClient) ListBySubscriptionPreparer(ctx context.Context, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client DedicatedHsmClient) ListBySubscriptionResponder(resp *http.Response) (result DedicatedHsmListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client DedicatedHsmClient) listBySubscriptionNextResults(ctx context.Context, lastResults DedicatedHsmListResult) (result DedicatedHsmListResult, err error) {
	req, err := lastResults.dedicatedHsmListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedHsmClient) ListBySubscriptionComplete(ctx context.Context, top *int32) (result DedicatedHsmListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, top)
	return
}

// ListOutboundNetworkDependenciesEndpoints gets a list of egress endpoints (network endpoints of all outbound
// dependencies) in the specified dedicated hsm resource. The operation returns properties of each egress endpoint.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the dedicated hsm belongs.
// name - the name of the dedicated HSM.
func (client DedicatedHsmClient) ListOutboundNetworkDependenciesEndpoints(ctx context.Context, resourceGroupName string, name string) (result OutboundEnvironmentEndpointCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.ListOutboundNetworkDependenciesEndpoints")
		defer func() {
			sc := -1
			if result.oeec.Response.Response != nil {
				sc = result.oeec.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listOutboundNetworkDependenciesEndpointsNextResults
	req, err := client.ListOutboundNetworkDependenciesEndpointsPreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOutboundNetworkDependenciesEndpointsSender(req)
	if err != nil {
		result.oeec.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListOutboundNetworkDependenciesEndpoints", resp, "Failure sending request")
		return
	}

	result.oeec, err = client.ListOutboundNetworkDependenciesEndpointsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "ListOutboundNetworkDependenciesEndpoints", resp, "Failure responding to request")
		return
	}
	if result.oeec.hasNextLink() && result.oeec.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListOutboundNetworkDependenciesEndpointsPreparer prepares the ListOutboundNetworkDependenciesEndpoints request.
func (client DedicatedHsmClient) ListOutboundNetworkDependenciesEndpointsPreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}/outboundNetworkDependenciesEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOutboundNetworkDependenciesEndpointsSender sends the ListOutboundNetworkDependenciesEndpoints request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) ListOutboundNetworkDependenciesEndpointsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListOutboundNetworkDependenciesEndpointsResponder handles the response to the ListOutboundNetworkDependenciesEndpoints request. The method always
// closes the http.Response Body.
func (client DedicatedHsmClient) ListOutboundNetworkDependenciesEndpointsResponder(resp *http.Response) (result OutboundEnvironmentEndpointCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listOutboundNetworkDependenciesEndpointsNextResults retrieves the next set of results, if any.
func (client DedicatedHsmClient) listOutboundNetworkDependenciesEndpointsNextResults(ctx context.Context, lastResults OutboundEnvironmentEndpointCollection) (result OutboundEnvironmentEndpointCollection, err error) {
	req, err := lastResults.outboundEnvironmentEndpointCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listOutboundNetworkDependenciesEndpointsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListOutboundNetworkDependenciesEndpointsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listOutboundNetworkDependenciesEndpointsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListOutboundNetworkDependenciesEndpointsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "listOutboundNetworkDependenciesEndpointsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListOutboundNetworkDependenciesEndpointsComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedHsmClient) ListOutboundNetworkDependenciesEndpointsComplete(ctx context.Context, resourceGroupName string, name string) (result OutboundEnvironmentEndpointCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.ListOutboundNetworkDependenciesEndpoints")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListOutboundNetworkDependenciesEndpoints(ctx, resourceGroupName, name)
	return
}

// Update update a dedicated HSM in the specified subscription.
// Parameters:
// resourceGroupName - the name of the Resource Group to which the server belongs.
// name - name of the dedicated HSM
// parameters - parameters to patch the dedicated HSM
func (client DedicatedHsmClient) Update(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters) (result DedicatedHsmUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: name,
			Constraints: []validation.Constraint{{Target: "name", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hardwaresecuritymodules.DedicatedHsmClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DedicatedHsmClient) UpdatePreparer(ctx context.Context, resourceGroupName string, name string, parameters DedicatedHsmPatchParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHsmClient) UpdateSender(req *http.Request) (future DedicatedHsmUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DedicatedHsmClient) UpdateResponder(resp *http.Response) (result DedicatedHsm, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
