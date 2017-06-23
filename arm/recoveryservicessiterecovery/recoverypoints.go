package recoveryservicessiterecovery

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// RecoveryPointsClient is the client for the RecoveryPoints methods of the
// Recoveryservicessiterecovery service.
type RecoveryPointsClient struct {
	ManagementClient
}

// NewRecoveryPointsClient creates an instance of the RecoveryPointsClient
// client.
func NewRecoveryPointsClient(subscriptionID string, resourceGroupName string, resourceName string) RecoveryPointsClient {
	return NewRecoveryPointsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewRecoveryPointsClientWithBaseURI creates an instance of the
// RecoveryPointsClient client.
func NewRecoveryPointsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) RecoveryPointsClient {
	return RecoveryPointsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get get the details of specified recovery point.
//
// fabricName is the fabric name. protectionContainerName is the protection
// container name. replicatedProtectedItemName is the replication protected
// item's name. recoveryPointName is the recovery point name.
func (client RecoveryPointsClient) Get(fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string) (result RecoveryPoint, err error) {
	req, err := client.GetPreparer(fabricName, protectionContainerName, replicatedProtectedItemName, recoveryPointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecoveryPointsClient) GetPreparer(fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                  autorest.Encode("path", fabricName),
		"protectionContainerName":     autorest.Encode("path", protectionContainerName),
		"recoveryPointName":           autorest.Encode("path", recoveryPointName),
		"replicatedProtectedItemName": autorest.Encode("path", replicatedProtectedItemName),
		"resourceGroupName":           autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                autorest.Encode("path", client.ResourceName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints/{recoveryPointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) GetResponder(resp *http.Response) (result RecoveryPoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationProtectedItems lists the available recovery points for a
// replication protected item.
//
// fabricName is the fabric name. protectionContainerName is the protection
// container name. replicatedProtectedItemName is the replication protected
// item's name.
func (client RecoveryPointsClient) ListByReplicationProtectedItems(fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result RecoveryPointCollection, err error) {
	req, err := client.ListByReplicationProtectedItemsPreparer(fabricName, protectionContainerName, replicatedProtectedItemName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationProtectedItemsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", resp, "Failure sending request")
		return
	}

	result, err = client.ListByReplicationProtectedItemsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationProtectedItemsPreparer prepares the ListByReplicationProtectedItems request.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsPreparer(fabricName string, protectionContainerName string, replicatedProtectedItemName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                  autorest.Encode("path", fabricName),
		"protectionContainerName":     autorest.Encode("path", protectionContainerName),
		"replicatedProtectedItemName": autorest.Encode("path", replicatedProtectedItemName),
		"resourceGroupName":           autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                autorest.Encode("path", client.ResourceName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByReplicationProtectedItemsSender sends the ListByReplicationProtectedItems request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByReplicationProtectedItemsResponder handles the response to the ListByReplicationProtectedItems request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsResponder(resp *http.Response) (result RecoveryPointCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationProtectedItemsNextResults retrieves the next set of results, if any.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsNextResults(lastResults RecoveryPointCollection) (result RecoveryPointCollection, err error) {
	req, err := lastResults.RecoveryPointCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByReplicationProtectedItemsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", resp, "Failure sending next results request")
	}

	result, err = client.ListByReplicationProtectedItemsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.RecoveryPointsClient", "ListByReplicationProtectedItems", resp, "Failure responding to next results request")
	}

	return
}

// ListByReplicationProtectedItemsComplete gets all elements from the list without paging.
func (client RecoveryPointsClient) ListByReplicationProtectedItemsComplete(fabricName string, protectionContainerName string, replicatedProtectedItemName string, cancel <-chan struct{}) (<-chan RecoveryPoint, <-chan error) {
	resultChan := make(chan RecoveryPoint)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByReplicationProtectedItems(fabricName, protectionContainerName, replicatedProtectedItemName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByReplicationProtectedItemsNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
