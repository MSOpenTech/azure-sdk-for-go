//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_List.json
func ExampleWebPubSubPrivateEndpointConnectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwebpubsub.NewWebPubSubPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<resource-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("PrivateEndpointConnection.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Get.json
func ExampleWebPubSubPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwebpubsub.NewWebPubSubPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnection.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Update.json
func ExampleWebPubSubPrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwebpubsub.NewWebPubSubPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		armwebpubsub.PrivateEndpointConnection{
			Properties: &armwebpubsub.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armwebpubsub.PrivateEndpoint{
					ID: to.StringPtr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armwebpubsub.PrivateLinkServiceConnectionState{
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          armwebpubsub.PrivateLinkServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnection.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Delete.json
func ExampleWebPubSubPrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwebpubsub.NewWebPubSubPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<private-endpoint-connection-name>",
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
