//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoSkus.json
func ExampleSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewSKUsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("westus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SKUDescriptionList = armkusto.SKUDescriptionList{
		// 	Value: []*armkusto.SKUDescription{
		// 		{
		// 			Name: to.Ptr("Standard_L8s_v3"),
		// 			LocationInfo: []*armkusto.SKULocationInfoItem{
		// 				{
		// 					Location: to.Ptr("West US"),
		// 					Zones: []*string{
		// 						to.Ptr("1"),
		// 						to.Ptr("2"),
		// 						to.Ptr("3")},
		// 					},
		// 					{
		// 						Location: to.Ptr("West US"),
		// 						Zones: []*string{
		// 						},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("West US")},
		// 					Tier: to.Ptr("Standard"),
		// 				},
		// 				{
		// 					Name: to.Ptr("Standard_L16s_v3"),
		// 					LocationInfo: []*armkusto.SKULocationInfoItem{
		// 						{
		// 							Location: to.Ptr("West US"),
		// 							Zones: []*string{
		// 								to.Ptr("1"),
		// 								to.Ptr("2"),
		// 								to.Ptr("3")},
		// 							},
		// 							{
		// 								Location: to.Ptr("West US"),
		// 								Zones: []*string{
		// 								},
		// 						}},
		// 						Locations: []*string{
		// 							to.Ptr("West US")},
		// 							Tier: to.Ptr("Standard"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Standard_L8as_v3"),
		// 							LocationInfo: []*armkusto.SKULocationInfoItem{
		// 								{
		// 									Location: to.Ptr("West US"),
		// 									Zones: []*string{
		// 										to.Ptr("1"),
		// 										to.Ptr("2"),
		// 										to.Ptr("3")},
		// 									},
		// 									{
		// 										Location: to.Ptr("West US"),
		// 										Zones: []*string{
		// 										},
		// 								}},
		// 								Locations: []*string{
		// 									to.Ptr("West US"),
		// 									to.Ptr("West Europe")},
		// 									Tier: to.Ptr("Standard"),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_L16as_v3"),
		// 									LocationInfo: []*armkusto.SKULocationInfoItem{
		// 										{
		// 											Location: to.Ptr("West US"),
		// 											Zones: []*string{
		// 												to.Ptr("1"),
		// 												to.Ptr("2"),
		// 												to.Ptr("3")},
		// 											},
		// 											{
		// 												Location: to.Ptr("West US"),
		// 												Zones: []*string{
		// 												},
		// 										}},
		// 										Locations: []*string{
		// 											to.Ptr("West US"),
		// 											to.Ptr("West Europe")},
		// 											Tier: to.Ptr("Standard"),
		// 									}},
		// 								}
	}
}
