//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/CloudServiceSwapGet.json
func ExampleVipSwapClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVipSwapClient().Get(ctx, "rg1", "testCloudService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SwapResource = armnetwork.SwapResource{
	// 	Name: to.Ptr("swap"),
	// 	Type: to.Ptr("Microsoft.Network/cloudServiceSlots"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/cloudServices/testCloudService/providers/Microsoft.Network/cloudServiceSlots/swap"),
	// 	Properties: &armnetwork.SwapResourceProperties{
	// 		SlotType: to.Ptr(armnetwork.SlotTypeStaging),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/CloudServiceSwapPut.json
func ExampleVipSwapClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVipSwapClient().BeginCreate(ctx, "rg1", "testCloudService", armnetwork.SwapResource{
		Properties: &armnetwork.SwapResourceProperties{
			SlotType: to.Ptr(armnetwork.SlotTypeProduction),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/CloudServiceSwapList.json
func ExampleVipSwapClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVipSwapClient().List(ctx, "rg1", "testCloudService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SwapResourceListResult = armnetwork.SwapResourceListResult{
	// 	Value: []*armnetwork.SwapResource{
	// 		{
	// 			Name: to.Ptr("swap"),
	// 			Type: to.Ptr("Microsoft.Network/cloudServiceSlots"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/cloudServices/testCloudService/providers/Microsoft.Network/cloudServiceSlots/swap"),
	// 			Properties: &armnetwork.SwapResourceProperties{
	// 				SlotType: to.Ptr(armnetwork.SlotTypeStaging),
	// 			},
	// 	}},
	// }
}
