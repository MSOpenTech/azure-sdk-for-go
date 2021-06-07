package backup

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

// RecoveryPointsRecommendedForMoveClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type RecoveryPointsRecommendedForMoveClient struct {
	BaseClient
}

// NewRecoveryPointsRecommendedForMoveClient creates an instance of the RecoveryPointsRecommendedForMoveClient client.
func NewRecoveryPointsRecommendedForMoveClient(subscriptionID string) RecoveryPointsRecommendedForMoveClient {
	return NewRecoveryPointsRecommendedForMoveClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecoveryPointsRecommendedForMoveClientWithBaseURI creates an instance of the
// RecoveryPointsRecommendedForMoveClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRecoveryPointsRecommendedForMoveClientWithBaseURI(baseURI string, subscriptionID string) RecoveryPointsRecommendedForMoveClient {
	return RecoveryPointsRecommendedForMoveClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the recovery points recommended for move to another tier
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// parameters - list Recovery points Recommended for Move Request
func (client RecoveryPointsRecommendedForMoveClient) List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters ListRecoveryPointsRecommendedForMoveRequest) (result RecoveryPointResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsRecommendedForMoveClient.List")
		defer func() {
			sc := -1
			if result.rprl.Response.Response != nil {
				sc = result.rprl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsRecommendedForMoveClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rprl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsRecommendedForMoveClient", "List", resp, "Failure sending request")
		return
	}

	result.rprl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsRecommendedForMoveClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rprl.hasNextLink() && result.rprl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RecoveryPointsRecommendedForMoveClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters ListRecoveryPointsRecommendedForMoveRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPointsRecommendedForMove", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsRecommendedForMoveClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecoveryPointsRecommendedForMoveClient) ListResponder(resp *http.Response) (result RecoveryPointResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecoveryPointsRecommendedForMoveClient) listNextResults(ctx context.Context, lastResults RecoveryPointResourceList) (result RecoveryPointResourceList, err error) {
	req, err := lastResults.recoveryPointResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "backup.RecoveryPointsRecommendedForMoveClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "backup.RecoveryPointsRecommendedForMoveClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsRecommendedForMoveClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecoveryPointsRecommendedForMoveClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters ListRecoveryPointsRecommendedForMoveRequest) (result RecoveryPointResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsRecommendedForMoveClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, parameters)
	return
}
