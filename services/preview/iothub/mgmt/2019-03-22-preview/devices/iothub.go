package devices

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IotHubClient is the use this API to manage the IoT hubs in your Azure subscription.
type IotHubClient struct {
	BaseClient
}

// NewIotHubClient creates an instance of the IotHubClient client.
func NewIotHubClient(subscriptionID string) IotHubClient {
	return NewIotHubClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIotHubClientWithBaseURI creates an instance of the IotHubClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIotHubClientWithBaseURI(baseURI string, subscriptionID string) IotHubClient {
	return IotHubClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ManualFailover perform manual fail over of given hub
// Parameters:
// iotHubName - iotHub to fail over
// failoverInput - region to failover to. Must be a azure DR pair
// resourceGroupName - resource group which Iot Hub belongs to
func (client IotHubClient) ManualFailover(ctx context.Context, iotHubName string, failoverInput FailoverInput, resourceGroupName string) (result IotHubManualFailoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotHubClient.ManualFailover")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: failoverInput,
			Constraints: []validation.Constraint{{Target: "failoverInput.FailoverRegion", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("devices.IotHubClient", "ManualFailover", err.Error())
	}

	req, err := client.ManualFailoverPreparer(ctx, iotHubName, failoverInput, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.IotHubClient", "ManualFailover", nil, "Failure preparing request")
		return
	}

	result, err = client.ManualFailoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.IotHubClient", "ManualFailover", nil, "Failure sending request")
		return
	}

	return
}

// ManualFailoverPreparer prepares the ManualFailover request.
func (client IotHubClient) ManualFailoverPreparer(ctx context.Context, iotHubName string, failoverInput FailoverInput, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"iotHubName":        autorest.Encode("path", iotHubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-22-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/failover", pathParameters),
		autorest.WithJSON(failoverInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ManualFailoverSender sends the ManualFailover request. The method will close the
// http.Response Body if it receives an error.
func (client IotHubClient) ManualFailoverSender(req *http.Request) (future IotHubManualFailoverFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client IotHubClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devices.IotHubManualFailoverFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("devices.IotHubManualFailoverFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// ManualFailoverResponder handles the response to the ManualFailover request. The method always
// closes the http.Response Body.
func (client IotHubClient) ManualFailoverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
