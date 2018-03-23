package consumption

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
	"net/http"
)

// MarketplacesClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type MarketplacesClient struct {
	BaseClient
}

// NewMarketplacesClient creates an instance of the MarketplacesClient client.
func NewMarketplacesClient(subscriptionID string, billingAccountID string) MarketplacesClient {
	return NewMarketplacesClientWithBaseURI(DefaultBaseURI, subscriptionID, billingAccountID)
}

// NewMarketplacesClientWithBaseURI creates an instance of the MarketplacesClient client.
func NewMarketplacesClientWithBaseURI(baseURI string, subscriptionID string, billingAccountID string) MarketplacesClient {
	return MarketplacesClient{NewWithBaseURI(baseURI, subscriptionID, billingAccountID)}
}

// List lists the marketplaces for a scope by subscriptionId. Marketplaces are available via this API only for May 1,
// 2014 or later.
//
// filter is may be used to filter marketplaces by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'. top is may be used to
// limit the number of results to the most recent N marketplaces. skiptoken is skiptoken is only used if a previous
// operation returned a partial result. If a previous response contains a nextLink element, the value of the
// nextLink element will include a skiptoken parameter that specifies a starting point to use for subsequent calls.
func (client MarketplacesClient) List(ctx context.Context, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "List", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MarketplacesClient) ListPreparer(ctx context.Context, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MarketplacesClient) ListResponder(resp *http.Response) (result MarketplacesListResult, err error) {
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
func (client MarketplacesClient) listNextResults(lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesClient) ListComplete(ctx context.Context, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	result.page, err = client.List(ctx, filter, top, skiptoken)
	return
}

// ListByBillingPeriod lists the marketplaces for a scope by billing period and subscripotionId. Marketplaces are
// available via this API only for May 1, 2014 or later.
//
// billingPeriodName is billing Period Name. filter is may be used to filter marketplaces by properties/usageEnd
// (Utc time), properties/usageStart (Utc time), properties/resourceGroup, properties/instanceName or
// properties/instanceId. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently
// support 'ne', 'or', or 'not'. top is may be used to limit the number of results to the most recent N
// marketplaces. skiptoken is skiptoken is only used if a previous operation returned a partial result. If a
// previous response contains a nextLink element, the value of the nextLink element will include a skiptoken
// parameter that specifies a starting point to use for subsequent calls.
func (client MarketplacesClient) ListByBillingPeriod(ctx context.Context, billingPeriodName string, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesClient", "ListByBillingPeriod", err.Error())
	}

	result.fn = client.listByBillingPeriodNextResults
	req, err := client.ListByBillingPeriodPreparer(ctx, billingPeriodName, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "ListByBillingPeriod", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "ListByBillingPeriod", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "ListByBillingPeriod", resp, "Failure responding to request")
	}

	return
}

// ListByBillingPeriodPreparer prepares the ListByBillingPeriod request.
func (client MarketplacesClient) ListByBillingPeriodPreparer(ctx context.Context, billingPeriodName string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingPeriodSender sends the ListByBillingPeriod request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesClient) ListByBillingPeriodSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByBillingPeriodResponder handles the response to the ListByBillingPeriod request. The method always
// closes the http.Response Body.
func (client MarketplacesClient) ListByBillingPeriodResponder(resp *http.Response) (result MarketplacesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingPeriodNextResults retrieves the next set of results, if any.
func (client MarketplacesClient) listByBillingPeriodNextResults(lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listByBillingPeriodNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listByBillingPeriodNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesClient", "listByBillingPeriodNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingPeriodComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesClient) ListByBillingPeriodComplete(ctx context.Context, billingPeriodName string, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	result.page, err = client.ListByBillingPeriod(ctx, billingPeriodName, filter, top, skiptoken)
	return
}
