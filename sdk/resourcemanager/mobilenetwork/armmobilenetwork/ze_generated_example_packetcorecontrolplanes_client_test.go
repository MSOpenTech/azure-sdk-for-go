//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneDelete.json
func ExamplePacketCoreControlPlanesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		&armmobilenetwork.PacketCoreControlPlanesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneGet.json
func ExamplePacketCoreControlPlanesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneCreate.json
func ExamplePacketCoreControlPlanesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		armmobilenetwork.PacketCoreControlPlane{
			Location: to.Ptr("<location>"),
			Properties: &armmobilenetwork.PacketCoreControlPlanePropertiesFormat{
				ControlPlaneAccessInterface: &armmobilenetwork.InterfaceProperties{
					Name: to.Ptr("<name>"),
				},
				CoreNetworkTechnology: to.Ptr(armmobilenetwork.CoreNetworkTypeFiveGC),
				CustomLocation: &armmobilenetwork.CustomLocationResourceID{
					ID: to.Ptr("<id>"),
				},
				MobileNetwork: &armmobilenetwork.ResourceID{
					ID: to.Ptr("<id>"),
				},
				Version: to.Ptr("<version>"),
			},
		},
		&armmobilenetwork.PacketCoreControlPlanesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneUpdateTags.json
func ExamplePacketCoreControlPlanesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		armmobilenetwork.TagsObject{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneListBySubscription.json
func ExamplePacketCoreControlPlanesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListBySubscription(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneListByResourceGroup.json
func ExamplePacketCoreControlPlanesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
