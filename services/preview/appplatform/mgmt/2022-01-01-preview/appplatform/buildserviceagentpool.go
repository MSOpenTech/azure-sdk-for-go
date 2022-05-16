package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// BuildServiceAgentPoolClient is the REST API for Azure Spring Cloud
type BuildServiceAgentPoolClient struct {
	BaseClient
}

// NewBuildServiceAgentPoolClient creates an instance of the BuildServiceAgentPoolClient client.
func NewBuildServiceAgentPoolClient(subscriptionID string) BuildServiceAgentPoolClient {
	return NewBuildServiceAgentPoolClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBuildServiceAgentPoolClientWithBaseURI creates an instance of the BuildServiceAgentPoolClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewBuildServiceAgentPoolClientWithBaseURI(baseURI string, subscriptionID string) BuildServiceAgentPoolClient {
	return BuildServiceAgentPoolClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get build service agent pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// buildServiceName - the name of the build service resource.
// agentPoolName - the name of the build service agent pool resource.
func (client BuildServiceAgentPoolClient) Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string) (result BuildServiceAgentPoolResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BuildServiceAgentPoolClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, buildServiceName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client BuildServiceAgentPoolClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"buildServiceName":  autorest.Encode("path", buildServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BuildServiceAgentPoolClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BuildServiceAgentPoolClient) GetResponder(resp *http.Response) (result BuildServiceAgentPoolResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list build service agent pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// buildServiceName - the name of the build service resource.
func (client BuildServiceAgentPoolClient) List(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result BuildServiceAgentPoolResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BuildServiceAgentPoolClient.List")
		defer func() {
			sc := -1
			if result.bsaprc.Response.Response != nil {
				sc = result.bsaprc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, buildServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.bsaprc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "List", resp, "Failure sending request")
		return
	}

	result.bsaprc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "List", resp, "Failure responding to request")
		return
	}
	if result.bsaprc.hasNextLink() && result.bsaprc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BuildServiceAgentPoolClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildServiceName":  autorest.Encode("path", buildServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BuildServiceAgentPoolClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BuildServiceAgentPoolClient) ListResponder(resp *http.Response) (result BuildServiceAgentPoolResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BuildServiceAgentPoolClient) listNextResults(ctx context.Context, lastResults BuildServiceAgentPoolResourceCollection) (result BuildServiceAgentPoolResourceCollection, err error) {
	req, err := lastResults.buildServiceAgentPoolResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BuildServiceAgentPoolClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result BuildServiceAgentPoolResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BuildServiceAgentPoolClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, serviceName, buildServiceName)
	return
}

// UpdatePut create or update build service agent pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// buildServiceName - the name of the build service resource.
// agentPoolName - the name of the build service agent pool resource.
// agentPoolResource - parameters for the update operation
func (client BuildServiceAgentPoolClient) UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, agentPoolResource BuildServiceAgentPoolResource) (result BuildServiceAgentPoolUpdatePutFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BuildServiceAgentPoolClient.UpdatePut")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePutPreparer(ctx, resourceGroupName, serviceName, buildServiceName, agentPoolName, agentPoolResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "UpdatePut", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdatePutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.BuildServiceAgentPoolClient", "UpdatePut", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePutPreparer prepares the UpdatePut request.
func (client BuildServiceAgentPoolClient) UpdatePutPreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, agentPoolResource BuildServiceAgentPoolResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"buildServiceName":  autorest.Encode("path", buildServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithJSON(agentPoolResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePutSender sends the UpdatePut request. The method will close the
// http.Response Body if it receives an error.
func (client BuildServiceAgentPoolClient) UpdatePutSender(req *http.Request) (future BuildServiceAgentPoolUpdatePutFuture, err error) {
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

// UpdatePutResponder handles the response to the UpdatePut request. The method always
// closes the http.Response Body.
func (client BuildServiceAgentPoolClient) UpdatePutResponder(resp *http.Response) (result BuildServiceAgentPoolResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
