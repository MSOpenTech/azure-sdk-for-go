package securityandcompliance

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

// PrivateLinkServicesForM365ComplianceCenterClient is the security And Compliance Client
type PrivateLinkServicesForM365ComplianceCenterClient struct {
	BaseClient
}

// NewPrivateLinkServicesForM365ComplianceCenterClient creates an instance of the
// PrivateLinkServicesForM365ComplianceCenterClient client.
func NewPrivateLinkServicesForM365ComplianceCenterClient(subscriptionID string) PrivateLinkServicesForM365ComplianceCenterClient {
	return NewPrivateLinkServicesForM365ComplianceCenterClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkServicesForM365ComplianceCenterClientWithBaseURI creates an instance of the
// PrivateLinkServicesForM365ComplianceCenterClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrivateLinkServicesForM365ComplianceCenterClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkServicesForM365ComplianceCenterClient {
	return PrivateLinkServicesForM365ComplianceCenterClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update the metadata of a privateLinkServicesForM365ComplianceCenter instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// privateLinkServicesForM365ComplianceCenterDescription - the service instance metadata.
func (client PrivateLinkServicesForM365ComplianceCenterClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForM365ComplianceCenterDescription PrivateLinkServicesForM365ComplianceCenterDescription) (result PrivateLinkServicesForM365ComplianceCenterCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: privateLinkServicesForM365ComplianceCenterDescription,
			Constraints: []validation.Constraint{{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CosmosDbConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CosmosDbConfiguration.OfferThroughput", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CosmosDbConfiguration.OfferThroughput", Name: validation.InclusiveMaximum, Rule: int64(10000), Chain: nil},
							{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CosmosDbConfiguration.OfferThroughput", Name: validation.InclusiveMinimum, Rule: int64(400), Chain: nil},
						}},
					}},
					{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CorsConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CorsConfiguration.MaxAge", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CorsConfiguration.MaxAge", Name: validation.InclusiveMaximum, Rule: int64(99999), Chain: nil},
								{Target: "privateLinkServicesForM365ComplianceCenterDescription.Properties.CorsConfiguration.MaxAge", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, privateLinkServicesForM365ComplianceCenterDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateLinkServicesForM365ComplianceCenterClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForM365ComplianceCenterDescription PrivateLinkServicesForM365ComplianceCenterDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{resourceName}", pathParameters),
		autorest.WithJSON(privateLinkServicesForM365ComplianceCenterDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForM365ComplianceCenterClient) CreateOrUpdateSender(req *http.Request) (future PrivateLinkServicesForM365ComplianceCenterCreateOrUpdateFuture, err error) {
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
func (client PrivateLinkServicesForM365ComplianceCenterClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateLinkServicesForM365ComplianceCenterDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a service instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
func (client PrivateLinkServicesForM365ComplianceCenterClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateLinkServicesForM365ComplianceCenterDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateLinkServicesForM365ComplianceCenterClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForM365ComplianceCenterClient) DeleteSender(req *http.Request) (future PrivateLinkServicesForM365ComplianceCenterDeleteFuture, err error) {
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
func (client PrivateLinkServicesForM365ComplianceCenterClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the metadata of a privateLinkServicesForM365ComplianceCenter resource.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
func (client PrivateLinkServicesForM365ComplianceCenterClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result PrivateLinkServicesForM365ComplianceCenterDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkServicesForM365ComplianceCenterClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForM365ComplianceCenterClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForM365ComplianceCenterClient) GetResponder(resp *http.Response) (result PrivateLinkServicesForM365ComplianceCenterDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all the privateLinkServicesForM365ComplianceCenter instances in a subscription.
func (client PrivateLinkServicesForM365ComplianceCenterClient) List(ctx context.Context) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.List")
		defer func() {
			sc := -1
			if result.plsfmccdlr.Response.Response != nil {
				sc = result.plsfmccdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.plsfmccdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "List", resp, "Failure sending request")
		return
	}

	result.plsfmccdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "List", resp, "Failure responding to request")
		return
	}
	if result.plsfmccdlr.hasNextLink() && result.plsfmccdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListResponder(resp *http.Response) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateLinkServicesForM365ComplianceCenterClient) listNextResults(ctx context.Context, lastResults PrivateLinkServicesForM365ComplianceCenterDescriptionListResult) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResult, err error) {
	req, err := lastResults.privateLinkServicesForM365ComplianceCenterDescriptionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListComplete(ctx context.Context) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.List")
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

// ListByResourceGroup get all the service instances in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.plsfmccdlr.Response.Response != nil {
				sc = result.plsfmccdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.plsfmccdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.plsfmccdlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.plsfmccdlr.hasNextLink() && result.plsfmccdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListByResourceGroupResponder(resp *http.Response) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client PrivateLinkServicesForM365ComplianceCenterClient) listByResourceGroupNextResults(ctx context.Context, lastResults PrivateLinkServicesForM365ComplianceCenterDescriptionListResult) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResult, err error) {
	req, err := lastResults.privateLinkServicesForM365ComplianceCenterDescriptionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkServicesForM365ComplianceCenterClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result PrivateLinkServicesForM365ComplianceCenterDescriptionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.ListByResourceGroup")
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

// Update update the metadata of a privateLinkServicesForM365ComplianceCenter instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the service instance.
// resourceName - the name of the service instance.
// servicePatchDescription - the service instance metadata and security metadata.
func (client PrivateLinkServicesForM365ComplianceCenterClient) Update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription) (result PrivateLinkServicesForM365ComplianceCenterUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkServicesForM365ComplianceCenterClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, servicePatchDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityandcompliance.PrivateLinkServicesForM365ComplianceCenterClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PrivateLinkServicesForM365ComplianceCenterClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-11"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{resourceName}", pathParameters),
		autorest.WithJSON(servicePatchDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkServicesForM365ComplianceCenterClient) UpdateSender(req *http.Request) (future PrivateLinkServicesForM365ComplianceCenterUpdateFuture, err error) {
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
func (client PrivateLinkServicesForM365ComplianceCenterClient) UpdateResponder(resp *http.Response) (result PrivateLinkServicesForM365ComplianceCenterDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
