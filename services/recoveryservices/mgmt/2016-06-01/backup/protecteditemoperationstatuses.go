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

// ProtectedItemOperationStatusesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectedItemOperationStatusesClient struct {
	BaseClient
}

// NewProtectedItemOperationStatusesClient creates an instance of the ProtectedItemOperationStatusesClient client.
func NewProtectedItemOperationStatusesClient(subscriptionID string) ProtectedItemOperationStatusesClient {
	return NewProtectedItemOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectedItemOperationStatusesClientWithBaseURI creates an instance of the ProtectedItemOperationStatusesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewProtectedItemOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) ProtectedItemOperationStatusesClient {
	return ProtectedItemOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the status of an operation such as triggering a backup or restore. The status can be: In progress,
// Completed, or Failed. You can refer to the OperationStatus enum for all the possible states of the operation. Some
// operations create jobs. This method returns the list of jobs associated with the operation.
// Parameters:
// vaultName - the name of the Recovery Services vault.
// resourceGroupName - the name of the resource group associated with the Recovery Services vault.
// fabricName - the fabric name associated with the backup item.
// containerName - the container name associated with the backup item.
// protectedItemName - the name of backup item used in this GET operation.
// operationID - the OperationID used in this GET operation.
func (client ProtectedItemOperationStatusesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectedItemOperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemOperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectedItemOperationStatusesClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"operationId":       autorest.Encode("path", operationID),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/operationsStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectedItemOperationStatusesClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
