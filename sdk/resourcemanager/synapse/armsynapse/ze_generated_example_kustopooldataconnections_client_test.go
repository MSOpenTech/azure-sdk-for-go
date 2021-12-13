//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCheckNameAvailability.json
func ExampleKustoPoolDataConnectionsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	_, err = client.CheckNameAvailability(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		armsynapse.DataConnectionCheckNameRequest{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionValidation.json
func ExampleKustoPoolDataConnectionsClient_BeginDataConnectionValidation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDataConnectionValidation(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		armsynapse.DataConnectionValidation{
			DataConnectionName: to.StringPtr("<data-connection-name>"),
			Properties: &armsynapse.EventHubDataConnection{
				DataConnection: armsynapse.DataConnection{
					Kind: armsynapse.DataConnectionKindEventHub.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsListByDatabase.json
func ExampleKustoPoolDataConnectionsClient_ListByDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	_, err = client.ListByDatabase(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsGet.json
func ExampleKustoPoolDataConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		"<data-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataConnectionClassification.GetDataConnection().ID: %s\n", *res.GetDataConnection().ID)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCreateOrUpdate.json
func ExampleKustoPoolDataConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		"<data-connection-name>",
		&armsynapse.EventHubDataConnection{
			DataConnection: armsynapse.DataConnection{
				Kind:     armsynapse.DataConnectionKindEventHub.ToPtr(),
				Location: to.StringPtr("<location>"),
			},
			Properties: &armsynapse.EventHubConnectionProperties{
				ConsumerGroup:      to.StringPtr("<consumer-group>"),
				EventHubResourceID: to.StringPtr("<event-hub-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataConnectionClassification.GetDataConnection().ID: %s\n", *res.GetDataConnection().ID)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsUpdate.json
func ExampleKustoPoolDataConnectionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		"<data-connection-name>",
		&armsynapse.EventHubDataConnection{
			DataConnection: armsynapse.DataConnection{
				Kind:     armsynapse.DataConnectionKindEventHub.ToPtr(),
				Location: to.StringPtr("<location>"),
			},
			Properties: &armsynapse.EventHubConnectionProperties{
				ConsumerGroup:      to.StringPtr("<consumer-group>"),
				EventHubResourceID: to.StringPtr("<event-hub-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataConnectionClassification.GetDataConnection().ID: %s\n", *res.GetDataConnection().ID)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsDelete.json
func ExampleKustoPoolDataConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolDataConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<kusto-pool-name>",
		"<database-name>",
		"<data-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
