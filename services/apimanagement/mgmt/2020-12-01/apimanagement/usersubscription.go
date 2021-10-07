package apimanagement

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

// UserSubscriptionClient is the apiManagement Client
type UserSubscriptionClient struct {
	BaseClient
}

// NewUserSubscriptionClient creates an instance of the UserSubscriptionClient client.
func NewUserSubscriptionClient(subscriptionID string) UserSubscriptionClient {
	return NewUserSubscriptionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserSubscriptionClientWithBaseURI creates an instance of the UserSubscriptionClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewUserSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) UserSubscriptionClient {
	return UserSubscriptionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified Subscription entity associated with a particular user.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// userID - user identifier. Must be unique in the current API Management service instance.
// sid - subscription entity Identifier. The entity represents the association between a user and a product in
// API Management.
func (client UserSubscriptionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, userID string, sid string) (result SubscriptionContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: sid,
			Constraints: []validation.Constraint{{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserSubscriptionClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, userID, sid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client UserSubscriptionClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, userID string, sid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/subscriptions/{sid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UserSubscriptionClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UserSubscriptionClient) GetResponder(resp *http.Response) (result SubscriptionContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the collection of subscriptions of the specified user.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// userID - user identifier. Must be unique in the current API Management service instance.
// filter - | Field     |     Usage     |     Supported operators    | Supported functions
// |</br>|-------------|------------------------|-----------------------------------|</br>|name | filter | ge,
// le, eq, ne, gt, lt | substringof, contains, startswith, endswith |</br>|displayName | filter | ge, le, eq,
// ne, gt, lt | substringof, contains, startswith, endswith |</br>|stateComment | filter | ge, le, eq, ne, gt,
// lt | substringof, contains, startswith, endswith |</br>|ownerId | filter | ge, le, eq, ne, gt, lt |
// substringof, contains, startswith, endswith |</br>|scope | filter | ge, le, eq, ne, gt, lt | substringof,
// contains, startswith, endswith |</br>|userId | filter | ge, le, eq, ne, gt, lt | substringof, contains,
// startswith, endswith |</br>|productId | filter | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |</br>
// top - number of records to return.
// skip - number of records to skip.
func (client UserSubscriptionClient) List(ctx context.Context, resourceGroupName string, serviceName string, userID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.sc.Response.Response != nil {
				sc = result.sc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserSubscriptionClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, userID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "List", resp, "Failure sending request")
		return
	}

	result.sc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sc.hasNextLink() && result.sc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client UserSubscriptionClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, userID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UserSubscriptionClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UserSubscriptionClient) ListResponder(resp *http.Response) (result SubscriptionCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UserSubscriptionClient) listNextResults(ctx context.Context, lastResults SubscriptionCollection) (result SubscriptionCollection, err error) {
	req, err := lastResults.subscriptionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UserSubscriptionClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, userID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, serviceName, userID, filter, top, skip)
	return
}
