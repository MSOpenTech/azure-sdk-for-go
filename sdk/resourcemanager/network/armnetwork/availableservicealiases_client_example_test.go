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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/AvailableServiceAliasesList.json
func ExampleAvailableServiceAliasesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableServiceAliasesClient().NewListPager("westcentralus", nil)
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
		// page.AvailableServiceAliasesResult = armnetwork.AvailableServiceAliasesResult{
		// 	Value: []*armnetwork.AvailableServiceAlias{
		// 		{
		// 			Name: to.Ptr("servicesAzure"),
		// 			Type: to.Ptr("Microsoft.Network/AvailableServiceAliases"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/AvailableServiceAliases/servicesAzure"),
		// 			ResourceName: to.Ptr("/services/Azure"),
		// 		},
		// 		{
		// 			Name: to.Ptr("servicesAzureManagedInstance"),
		// 			Type: to.Ptr("Microsoft.Network/AvailableServiceAliases"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/AvailableServiceAliases/servicesAzureManagedInstance"),
		// 			ResourceName: to.Ptr("/services/Azure/ManagedInstance"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/AvailableServiceAliasesListByResourceGroup.json
func ExampleAvailableServiceAliasesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableServiceAliasesClient().NewListByResourceGroupPager("rg1", "westcentralus", nil)
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
		// page.AvailableServiceAliasesResult = armnetwork.AvailableServiceAliasesResult{
		// 	Value: []*armnetwork.AvailableServiceAlias{
		// 		{
		// 			Name: to.Ptr("servicesAzure"),
		// 			Type: to.Ptr("Microsoft.Network/AvailableServiceAliases"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/AvailableServiceAliases/servicesAzure"),
		// 			ResourceName: to.Ptr("/services/Azure"),
		// 		},
		// 		{
		// 			Name: to.Ptr("servicesAzureManagedInstance"),
		// 			Type: to.Ptr("Microsoft.Network/AvailableServiceAliases"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/AvailableServiceAliases/servicesAzureManagedInstance"),
		// 			ResourceName: to.Ptr("/services/Azure/ManagedInstance"),
		// 	}},
		// }
	}
}
