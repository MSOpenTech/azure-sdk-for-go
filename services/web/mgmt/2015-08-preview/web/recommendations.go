package web

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
	"net/http"
)

// RecommendationsClient is the webSite Management Client
type RecommendationsClient struct {
	BaseClient
}

// NewRecommendationsClient creates an instance of the RecommendationsClient client.
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return NewRecommendationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendationsClientWithBaseURI creates an instance of the RecommendationsClient client.
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return RecommendationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetRecommendationBySubscription sends the get recommendation by subscription request.
//
// featured is if set, this API returns only the most critical recommendation among the others. Otherwise this API
// returns all recommendations available filter is return only channels specified in the filter. Filter is specified by
// using OData syntax. Example: $filter=channels eq 'Api' or channel eq 'Notification'
func (client RecommendationsClient) GetRecommendationBySubscription(ctx context.Context, featured *bool, filter string) (result ListRecommendation, err error) {
	req, err := client.GetRecommendationBySubscriptionPreparer(ctx, featured, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendationBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRecommendationBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendationBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.GetRecommendationBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendationBySubscription", resp, "Failure responding to request")
	}

	return
}

// GetRecommendationBySubscriptionPreparer prepares the GetRecommendationBySubscription request.
func (client RecommendationsClient) GetRecommendationBySubscriptionPreparer(ctx context.Context, featured *bool, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if featured != nil {
		queryParameters["featured"] = autorest.Encode("query", *featured)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRecommendationBySubscriptionSender sends the GetRecommendationBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetRecommendationBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetRecommendationBySubscriptionResponder handles the response to the GetRecommendationBySubscription request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetRecommendationBySubscriptionResponder(resp *http.Response) (result ListRecommendation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRecommendationHistoryForSite sends the get recommendation history for site request.
//
// resourceGroupName is resource group name siteName is site name startTime is the start time of a time range to query,
// e.g. $filter=startTime eq '2015-01-01T00:00:00Z' and endTime eq '2015-01-02T00:00:00Z' endTime is the end time of a
// time range to query, e.g. $filter=startTime eq '2015-01-01T00:00:00Z' and endTime eq '2015-01-02T00:00:00Z'
func (client RecommendationsClient) GetRecommendationHistoryForSite(ctx context.Context, resourceGroupName string, siteName string, startTime string, endTime string) (result ListRecommendation, err error) {
	req, err := client.GetRecommendationHistoryForSitePreparer(ctx, resourceGroupName, siteName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendationHistoryForSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRecommendationHistoryForSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendationHistoryForSite", resp, "Failure sending request")
		return
	}

	result, err = client.GetRecommendationHistoryForSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendationHistoryForSite", resp, "Failure responding to request")
	}

	return
}

// GetRecommendationHistoryForSitePreparer prepares the GetRecommendationHistoryForSite request.
func (client RecommendationsClient) GetRecommendationHistoryForSitePreparer(ctx context.Context, resourceGroupName string, siteName string, startTime string, endTime string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(startTime) > 0 {
		queryParameters["startTime"] = autorest.Encode("query", startTime)
	}
	if len(endTime) > 0 {
		queryParameters["endTime"] = autorest.Encode("query", endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRecommendationHistoryForSiteSender sends the GetRecommendationHistoryForSite request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetRecommendationHistoryForSiteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetRecommendationHistoryForSiteResponder handles the response to the GetRecommendationHistoryForSite request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetRecommendationHistoryForSiteResponder(resp *http.Response) (result ListRecommendation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRecommendedRulesForSite sends the get recommended rules for site request.
//
// resourceGroupName is resource group name siteName is site name featured is if set, this API returns only the most
// critical recommendation among the others. Otherwise this API returns all recommendations available siteSku is the
// name of site SKU. numSlots is the number of site slots associated to the site
func (client RecommendationsClient) GetRecommendedRulesForSite(ctx context.Context, resourceGroupName string, siteName string, featured *bool, siteSku string, numSlots *int32) (result ListRecommendation, err error) {
	req, err := client.GetRecommendedRulesForSitePreparer(ctx, resourceGroupName, siteName, featured, siteSku, numSlots)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendedRulesForSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRecommendedRulesForSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendedRulesForSite", resp, "Failure sending request")
		return
	}

	result, err = client.GetRecommendedRulesForSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRecommendedRulesForSite", resp, "Failure responding to request")
	}

	return
}

// GetRecommendedRulesForSitePreparer prepares the GetRecommendedRulesForSite request.
func (client RecommendationsClient) GetRecommendedRulesForSitePreparer(ctx context.Context, resourceGroupName string, siteName string, featured *bool, siteSku string, numSlots *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if featured != nil {
		queryParameters["featured"] = autorest.Encode("query", *featured)
	}
	if len(siteSku) > 0 {
		queryParameters["siteSku"] = autorest.Encode("query", siteSku)
	}
	if numSlots != nil {
		queryParameters["numSlots"] = autorest.Encode("query", *numSlots)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRecommendedRulesForSiteSender sends the GetRecommendedRulesForSite request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetRecommendedRulesForSiteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetRecommendedRulesForSiteResponder handles the response to the GetRecommendedRulesForSite request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetRecommendedRulesForSiteResponder(resp *http.Response) (result ListRecommendation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRuleDetailsBySiteName sends the get rule details by site name request.
//
// resourceGroupName is resource group name siteName is site name name is recommendation rule name
func (client RecommendationsClient) GetRuleDetailsBySiteName(ctx context.Context, resourceGroupName string, siteName string, name string) (result RecommendationRule, err error) {
	req, err := client.GetRuleDetailsBySiteNamePreparer(ctx, resourceGroupName, siteName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRuleDetailsBySiteName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRuleDetailsBySiteNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRuleDetailsBySiteName", resp, "Failure sending request")
		return
	}

	result, err = client.GetRuleDetailsBySiteNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.RecommendationsClient", "GetRuleDetailsBySiteName", resp, "Failure responding to request")
	}

	return
}

// GetRuleDetailsBySiteNamePreparer prepares the GetRuleDetailsBySiteName request.
func (client RecommendationsClient) GetRuleDetailsBySiteNamePreparer(ctx context.Context, resourceGroupName string, siteName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRuleDetailsBySiteNameSender sends the GetRuleDetailsBySiteName request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetRuleDetailsBySiteNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetRuleDetailsBySiteNameResponder handles the response to the GetRuleDetailsBySiteName request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetRuleDetailsBySiteNameResponder(resp *http.Response) (result RecommendationRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
