//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBManagedCassandraDataCenterList.json
func ExampleCassandraDataCentersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraDataCentersClient().NewListPager("cassandra-prod-rg", "cassandra-prod", nil)
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
		// page.ListDataCenters = armcosmos.ListDataCenters{
		// 	Value: []*armcosmos.DataCenterResource{
		// 		{
		// 			Name: to.Ptr("dc1"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/dataCenters"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/dataCenters"),
		// 			Properties: &armcosmos.DataCenterResourceProperties{
		// 				Base64EncodedCassandraYamlFragment: to.Ptr("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
		// 				DataCenterLocation: to.Ptr("West US 2"),
		// 				DelegatedSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1"),
		// 				NodeCount: to.Ptr[int32](9),
		// 				ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
		// 				SeedNodes: []*armcosmos.SeedNode{
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.2"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.3"),
		// 					},
		// 					{
		// 						IPAddress: to.Ptr("192.168.12.4"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBManagedCassandraDataCenterGet.json
func ExampleCassandraDataCentersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCassandraDataCentersClient().Get(ctx, "cassandra-prod-rg", "cassandra-prod", "dc1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataCenterResource = armcosmos.DataCenterResource{
	// 	Name: to.Ptr("dc1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/dataCenters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/dataCenters/dc1"),
	// 	Properties: &armcosmos.DataCenterResourceProperties{
	// 		Base64EncodedCassandraYamlFragment: to.Ptr("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
	// 		DataCenterLocation: to.Ptr("West US 2"),
	// 		DelegatedSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1"),
	// 		NodeCount: to.Ptr[int32](9),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBManagedCassandraDataCenterDelete.json
func ExampleCassandraDataCentersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraDataCentersClient().BeginDelete(ctx, "cassandra-prod-rg", "cassandra-prod", "dc1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBManagedCassandraDataCenterCreate.json
func ExampleCassandraDataCentersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraDataCentersClient().BeginCreateUpdate(ctx, "cassandra-prod-rg", "cassandra-prod", "dc1", armcosmos.DataCenterResource{
		Properties: &armcosmos.DataCenterResourceProperties{
			Base64EncodedCassandraYamlFragment: to.Ptr("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
			DataCenterLocation:                 to.Ptr("West US 2"),
			DelegatedSubnetID:                  to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1-subnet"),
			NodeCount:                          to.Ptr[int32](9),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataCenterResource = armcosmos.DataCenterResource{
	// 	Name: to.Ptr("dc1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/dataCenters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/dataCenters/dc1"),
	// 	Properties: &armcosmos.DataCenterResourceProperties{
	// 		DataCenterLocation: to.Ptr("West US 2"),
	// 		DelegatedSubnetID: to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1"),
	// 		NodeCount: to.Ptr[int32](9),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBManagedCassandraDataCenterPatch.json
func ExampleCassandraDataCentersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraDataCentersClient().BeginUpdate(ctx, "cassandra-prod-rg", "cassandra-prod", "dc1", armcosmos.DataCenterResource{
		Properties: &armcosmos.DataCenterResourceProperties{
			Base64EncodedCassandraYamlFragment: to.Ptr("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
			DataCenterLocation:                 to.Ptr("West US 2"),
			DelegatedSubnetID:                  to.Ptr("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1-subnet"),
			NodeCount:                          to.Ptr[int32](9),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataCenterResource = armcosmos.DataCenterResource{
	// 	Name: to.Ptr("dc1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/cassandraClusters/dataCenters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cassandra-prod-rg/providers/Microsoft.DocumentDB/cassandraClusters/cassandra-prod/dataCenters/dc1"),
	// 	Properties: &armcosmos.DataCenterResourceProperties{
	// 		Base64EncodedCassandraYamlFragment: to.Ptr("Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA="),
	// 		DataCenterLocation: to.Ptr("West US 2"),
	// 		DelegatedSubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1"),
	// 		NodeCount: to.Ptr[int32](9),
	// 		ProvisioningState: to.Ptr(armcosmos.ManagedCassandraProvisioningStateSucceeded),
	// 		SeedNodes: []*armcosmos.SeedNode{
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.2"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.3"),
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("192.168.12.4"),
	// 		}},
	// 	},
	// }
}
