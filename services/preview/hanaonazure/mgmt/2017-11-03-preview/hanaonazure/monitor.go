package hanaonazure

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

// MonitorClient is the HANA on Azure Client
type MonitorClient struct {
	BaseClient
}

// NewMonitorClient creates an instance of the MonitorClient client.
func NewMonitorClient(subscriptionID string) MonitorClient {
	return NewMonitorClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitorClientWithBaseURI creates an instance of the MonitorClient client.
func NewMonitorClientWithBaseURI(baseURI string, subscriptionID string) MonitorClient {
	return MonitorClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// HanaInstancesMethod the operation to monitor a SAP HANA instance.
// Parameters:
// resourceGroupName - name of the resource group.
// hanaInstanceName - name of the SAP HANA on Azure instance.
// monitoringParameter - request body that only contains monitoring attributes
func (client MonitorClient) HanaInstancesMethod(ctx context.Context, resourceGroupName string, hanaInstanceName string, monitoringParameter MonitoringDetails) (result MonitorHanaInstancesMethodFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorClient.HanaInstancesMethod")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.HanaInstancesMethodPreparer(ctx, resourceGroupName, hanaInstanceName, monitoringParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.MonitorClient", "HanaInstancesMethod", nil, "Failure preparing request")
		return
	}

	result, err = client.HanaInstancesMethodSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.MonitorClient", "HanaInstancesMethod", result.Response(), "Failure sending request")
		return
	}

	return
}

// HanaInstancesMethodPreparer prepares the HanaInstancesMethod request.
func (client MonitorClient) HanaInstancesMethodPreparer(ctx context.Context, resourceGroupName string, hanaInstanceName string, monitoringParameter MonitoringDetails) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hanaInstanceName":  autorest.Encode("path", hanaInstanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}/monitoring", pathParameters),
		autorest.WithJSON(monitoringParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// HanaInstancesMethodSender sends the HanaInstancesMethod request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorClient) HanaInstancesMethodSender(req *http.Request) (future MonitorHanaInstancesMethodFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// HanaInstancesMethodResponder handles the response to the HanaInstancesMethod request. The method always
// closes the http.Response Body.
func (client MonitorClient) HanaInstancesMethodResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
