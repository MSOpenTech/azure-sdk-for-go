package network

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// VirtualNetworkGatewayConnectionsClient is the network Client
type VirtualNetworkGatewayConnectionsClient struct {
    BaseClient
}
// NewVirtualNetworkGatewayConnectionsClient creates an instance of the VirtualNetworkGatewayConnectionsClient client.
func NewVirtualNetworkGatewayConnectionsClient(subscriptionID string) VirtualNetworkGatewayConnectionsClient {
    return NewVirtualNetworkGatewayConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkGatewayConnectionsClientWithBaseURI creates an instance of the
// VirtualNetworkGatewayConnectionsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewVirtualNetworkGatewayConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewayConnectionsClient {
        return VirtualNetworkGatewayConnectionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate the Put VirtualNetworkGatewayConnection operation creates/updates a virtual network gateway
// connection in the specified resource group through Network resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // virtualNetworkGatewayConnectionName - the name of the virtual network gateway connection.
        // parameters - parameters supplied to the Begin Create or update Virtual Network Gateway connection operation
        // through Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (result VirtualNetworkGatewayConnectionsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualNetworkGatewayConnectionName": autorest.Encode("path",virtualNetworkGatewayConnectionName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdateSender(req *http.Request) (future VirtualNetworkGatewayConnectionsCreateOrUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkGatewayConnection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete the Delete VirtualNetworkGatewayConnection operation deletes the specified virtual network Gateway connection
// through Network resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // virtualNetworkGatewayConnectionName - the name of the virtual network gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result VirtualNetworkGatewayConnectionsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client VirtualNetworkGatewayConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualNetworkGatewayConnectionName": autorest.Encode("path",virtualNetworkGatewayConnectionName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) DeleteSender(req *http.Request) (future VirtualNetworkGatewayConnectionsDeleteFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get the Get VirtualNetworkGatewayConnection operation retrieves information about the specified virtual network
// gateway connection through Network resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // virtualNetworkGatewayConnectionName - the name of the virtual network gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result VirtualNetworkGatewayConnection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client VirtualNetworkGatewayConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualNetworkGatewayConnectionName": autorest.Encode("path",virtualNetworkGatewayConnectionName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) GetResponder(resp *http.Response) (result VirtualNetworkGatewayConnection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetSharedKey the Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified
// virtual network gateway connection shared key through Network resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // virtualNetworkGatewayConnectionName - the virtual network gateway connection shared key name.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result ConnectionSharedKeyResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.GetSharedKey")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetSharedKeyPreparer(ctx, resourceGroupName, virtualNetworkGatewayConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "GetSharedKey", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSharedKeySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "GetSharedKey", resp, "Failure sending request")
            return
            }

            result, err = client.GetSharedKeyResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "GetSharedKey", resp, "Failure responding to request")
            }

    return
    }

    // GetSharedKeyPreparer prepares the GetSharedKey request.
    func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeyPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualNetworkGatewayConnectionName": autorest.Encode("path",virtualNetworkGatewayConnectionName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSharedKeySender sends the GetSharedKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeySender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetSharedKeyResponder handles the response to the GetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeyResponder(resp *http.Response) (result ConnectionSharedKeyResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List the List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections
// created.
    // Parameters:
        // resourceGroupName - the name of the resource group.
func (client VirtualNetworkGatewayConnectionsClient) List(ctx context.Context, resourceGroupName string) (result VirtualNetworkGatewayConnectionListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.List")
        defer func() {
            sc := -1
            if result.vngclr.Response.Response != nil {
                sc = result.vngclr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.vngclr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure sending request")
            return
            }

            result.vngclr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client VirtualNetworkGatewayConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) ListResponder(resp *http.Response) (result VirtualNetworkGatewayConnectionListResult, err error) {
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
            func (client VirtualNetworkGatewayConnectionsClient) listNextResults(ctx context.Context, lastResults VirtualNetworkGatewayConnectionListResult) (result VirtualNetworkGatewayConnectionListResult, err error) {
            req, err := lastResults.virtualNetworkGatewayConnectionListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client VirtualNetworkGatewayConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string) (result VirtualNetworkGatewayConnectionListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName)
                return
        }

// ResetSharedKey the VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway
// connection shared key for passed virtual network gateway connection in the specified resource group through Network
// resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // virtualNetworkGatewayConnectionName - the virtual network gateway connection reset shared key Name.
        // parameters - parameters supplied to the Begin Reset Virtual Network Gateway connection shared key operation
        // through Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (result VirtualNetworkGatewayConnectionsResetSharedKeyFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.ResetSharedKey")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ResetSharedKeyPreparer(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", nil , "Failure preparing request")
    return
    }

            result, err = client.ResetSharedKeySender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // ResetSharedKeyPreparer prepares the ResetSharedKey request.
    func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeyPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualNetworkGatewayConnectionName": autorest.Encode("path",virtualNetworkGatewayConnectionName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ResetSharedKeySender sends the ResetSharedKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeySender(req *http.Request) (future VirtualNetworkGatewayConnectionsResetSharedKeyFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// ResetSharedKeyResponder handles the response to the ResetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeyResponder(resp *http.Response) (result ConnectionResetSharedKey, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// SetSharedKey the Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection
// shared key for passed virtual network gateway connection in the specified resource group through Network resource
// provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // virtualNetworkGatewayConnectionName - the virtual network gateway connection name.
        // parameters - parameters supplied to the Begin Set Virtual Network Gateway connection Shared key operation
        // throughNetwork resource provider.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (result VirtualNetworkGatewayConnectionsSetSharedKeyFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualNetworkGatewayConnectionsClient.SetSharedKey")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.SetSharedKeyPreparer(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "SetSharedKey", nil , "Failure preparing request")
    return
    }

            result, err = client.SetSharedKeySender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "SetSharedKey", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // SetSharedKeyPreparer prepares the SetSharedKey request.
    func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeyPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualNetworkGatewayConnectionName": autorest.Encode("path",virtualNetworkGatewayConnectionName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // SetSharedKeySender sends the SetSharedKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeySender(req *http.Request) (future VirtualNetworkGatewayConnectionsSetSharedKeyFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// SetSharedKeyResponder handles the response to the SetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeyResponder(resp *http.Response) (result ConnectionSharedKey, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

