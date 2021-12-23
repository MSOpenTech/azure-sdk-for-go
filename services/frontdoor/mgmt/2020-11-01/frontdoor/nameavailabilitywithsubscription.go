package frontdoor

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

// NameAvailabilityWithSubscriptionClient is the frontDoor Client
type NameAvailabilityWithSubscriptionClient struct {
	BaseClient
}

// NewNameAvailabilityWithSubscriptionClient creates an instance of the NameAvailabilityWithSubscriptionClient client.
func NewNameAvailabilityWithSubscriptionClient(subscriptionID string) NameAvailabilityWithSubscriptionClient {
	return NewNameAvailabilityWithSubscriptionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNameAvailabilityWithSubscriptionClientWithBaseURI creates an instance of the
// NameAvailabilityWithSubscriptionClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNameAvailabilityWithSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) NameAvailabilityWithSubscriptionClient {
	return NameAvailabilityWithSubscriptionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Check check the availability of a Front Door subdomain.
// Parameters:
// checkFrontDoorNameAvailabilityInput - input to check.
func (client NameAvailabilityWithSubscriptionClient) Check(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NameAvailabilityWithSubscriptionClient.Check")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkFrontDoorNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkFrontDoorNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.NameAvailabilityWithSubscriptionClient", "Check", err.Error())
	}

	req, err := client.CheckPreparer(ctx, checkFrontDoorNameAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.NameAvailabilityWithSubscriptionClient", "Check", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.NameAvailabilityWithSubscriptionClient", "Check", resp, "Failure sending request")
		return
	}

	result, err = client.CheckResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.NameAvailabilityWithSubscriptionClient", "Check", resp, "Failure responding to request")
		return
	}

	return
}

// CheckPreparer prepares the Check request.
func (client NameAvailabilityWithSubscriptionClient) CheckPreparer(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/checkFrontDoorNameAvailability", pathParameters),
		autorest.WithJSON(checkFrontDoorNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckSender sends the Check request. The method will close the
// http.Response Body if it receives an error.
func (client NameAvailabilityWithSubscriptionClient) CheckSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckResponder handles the response to the Check request. The method always
// closes the http.Response Body.
func (client NameAvailabilityWithSubscriptionClient) CheckResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
