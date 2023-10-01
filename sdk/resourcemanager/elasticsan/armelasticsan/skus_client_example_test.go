//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Skus_List_MaximumSet_Gen.json
func ExampleSKUsClient_NewListPager_skusListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(&armelasticsan.SKUsClientListOptions{Filter: to.Ptr("obwwdrkq")})
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
		// page.SKUInformationList = armelasticsan.SKUInformationList{
		// 	Value: []*armelasticsan.SKUInformation{
		// 		{
		// 			Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
		// 			Capabilities: []*armelasticsan.SKUCapability{
		// 				{
		// 					Name: to.Ptr("qkht"),
		// 					Value: to.Ptr("eoayvlyzyjjziecxymlpk"),
		// 			}},
		// 			LocationInfo: []*armelasticsan.SKULocationInfo{
		// 				{
		// 					Location: to.Ptr("ngycrsoihxdfctigejlf"),
		// 					Zones: []*string{
		// 						to.Ptr("1")},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("ceszpxwqyifrxobqykplm")},
		// 					ResourceType: to.Ptr("tlqickysdtjahoanstgancifxfu"),
		// 					Tier: to.Ptr(armelasticsan.SKUTierPremium),
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Skus_List_MinimumSet_Gen.json
func ExampleSKUsClient_NewListPager_skusListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(&armelasticsan.SKUsClientListOptions{Filter: nil})
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
		// page.SKUInformationList = armelasticsan.SKUInformationList{
		// }
	}
}
