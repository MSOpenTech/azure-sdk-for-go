//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersListByResourceGroup.json
func ExampleClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListByResourceGroupPager("oiautorest6685", nil)
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
		// page.ClusterListResult = armoperationalinsights.ClusterListResult{
		// 	Value: []*armoperationalinsights.Cluster{
		// 		{
		// 			Name: to.Ptr("TestResourceLock"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
		// 			ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/clusters/testresourcelock"),
		// 			Location: to.Ptr("eastus"),
		// 			Identity: &armoperationalinsights.Identity{
		// 				Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Properties: &armoperationalinsights.ClusterProperties{
		// 				ClusterID: to.Ptr("5b02755b-5bf4-430c-9487-45502a2a7e62"),
		// 				KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
		// 					KeyName: to.Ptr("aztest2170cert"),
		// 					KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
		// 					KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
		// 				},
		// 				ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
		// 			},
		// 			SKU: &armoperationalinsights.ClusterSKU{
		// 				Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
		// 				Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersSubscriptionList.json
func ExampleClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListPager(nil)
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
		// page.ClusterListResult = armoperationalinsights.ClusterListResult{
		// 	Value: []*armoperationalinsights.Cluster{
		// 		{
		// 			Name: to.Ptr("TestResourceLock"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
		// 			ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/clusters/testresourcelock"),
		// 			Location: to.Ptr("eastus"),
		// 			Identity: &armoperationalinsights.Identity{
		// 				Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Properties: &armoperationalinsights.ClusterProperties{
		// 				ClusterID: to.Ptr("5b02755b-5bf4-430c-9487-45502a2a7e62"),
		// 				KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
		// 					KeyName: to.Ptr("aztest2170cert"),
		// 					KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
		// 					KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
		// 				},
		// 				ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
		// 			},
		// 			SKU: &armoperationalinsights.ClusterSKU{
		// 				Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
		// 				Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersCreate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.Cluster{
		Location: to.Ptr("australiasoutheast"),
		Tags: map[string]*string{
			"tag1": to.Ptr("val1"),
		},
		SKU: &armoperationalinsights.ClusterSKU{
			Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
			Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
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
	// res.Cluster = armoperationalinsights.Cluster{
	// 	Name: to.Ptr("oiautorest6685"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
	// 	ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/clusters/oiautorest6685"),
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Identity: &armoperationalinsights.Identity{
	// 		Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Properties: &armoperationalinsights.ClusterProperties{
	// 		AssociatedWorkspaces: []*armoperationalinsights.AssociatedWorkspace{
	// 		},
	// 		BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 		CapacityReservationProperties: &armoperationalinsights.CapacityReservationProperties{
	// 			LastSKUUpdate: to.Ptr("Thu, 01 Jan 1970 00:00:00 GMT"),
	// 			MinCapacity: to.Ptr[int64](500),
	// 		},
	// 		ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 		CreatedDate: to.Ptr("Mon, 13 Jan 2020 14:40:33 GMT"),
	// 		KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
	// 			KeyName: to.Ptr("aztest2170cert"),
	// 			KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
	// 			KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
	// 		},
	// 		LastModifiedDate: to.Ptr("Sun, 04 Jan 2020 17:10:56 GMT"),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
	// 	},
	// 	SKU: &armoperationalinsights.ClusterSKU{
	// 		Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
	// 		Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersDelete.json
func ExampleClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginDelete(ctx, "oiautorest6685", "oiautorest6685", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersGet.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().Get(ctx, "oiautorest6685", "oiautorest6685", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armoperationalinsights.Cluster{
	// 	Name: to.Ptr("TestResourceLock"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/clusters"),
	// 	ID: to.Ptr("/subscriptions/594038b5-1093-476e-a366-482775671c11/resourcegroups/calbot-rg/providers/microsoft.operationalinsights/clusters/testresourcelock"),
	// 	Location: to.Ptr("eastus"),
	// 	Identity: &armoperationalinsights.Identity{
	// 		Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("356d057d-cba5-44dd-8a30-b2e547bc416b"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Properties: &armoperationalinsights.ClusterProperties{
	// 		AssociatedWorkspaces: []*armoperationalinsights.AssociatedWorkspace{
	// 			{
	// 				AssociateDate: to.Ptr("Tue, 07 Jul 2020 07:35:51 GMT"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs1"),
	// 				WorkspaceID: to.Ptr("942bdefd-e6c9-411c-ac69-70ffad564363"),
	// 				WorkspaceName: to.Ptr("testWs1"),
	// 			},
	// 			{
	// 				AssociateDate: to.Ptr("Mon, 13 Jan 2020 16:03:39 GMT"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs2"),
	// 				WorkspaceID: to.Ptr("c7edb8f8-67f7-41f2-bddb-aecf22507e3f"),
	// 				WorkspaceName: to.Ptr("testWs2"),
	// 		}},
	// 		BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 		CapacityReservationProperties: &armoperationalinsights.CapacityReservationProperties{
	// 			LastSKUUpdate: to.Ptr("Thu, 01 Jan 1970 00:00:00 GMT"),
	// 			MinCapacity: to.Ptr[int64](500),
	// 		},
	// 		ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 		CreatedDate: to.Ptr("Mon, 13 Jan 2020 14:40:33 GMT"),
	// 		IsAvailabilityZonesEnabled: to.Ptr(false),
	// 		IsDoubleEncryptionEnabled: to.Ptr(false),
	// 		KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
	// 			KeyName: to.Ptr("aztest2170cert"),
	// 			KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
	// 			KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
	// 		},
	// 		LastModifiedDate: to.Ptr("Sun, 04 Jan 2020 17:10:56 GMT"),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
	// 	},
	// 	SKU: &armoperationalinsights.ClusterSKU{
	// 		Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
	// 		Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersUpdate.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.ClusterPatch{
		Identity: &armoperationalinsights.Identity{
			Type: to.Ptr(armoperationalinsights.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armoperationalinsights.UserIdentityProperties{
				"/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
			},
		},
		Properties: &armoperationalinsights.ClusterPatchProperties{
			KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
				KeyName:     to.Ptr("aztest2170cert"),
				KeyRsaSize:  to.Ptr[int32](1024),
				KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
				KeyVersion:  to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
			},
		},
		SKU: &armoperationalinsights.ClusterSKU{
			Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
			Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("val1"),
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
	// res.Cluster = armoperationalinsights.Cluster{
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("val1"),
	// 	},
	// 	Identity: &armoperationalinsights.Identity{
	// 		Type: to.Ptr(armoperationalinsights.IdentityTypeUserAssigned),
	// 		TenantID: to.Ptr("72f999bf-acf1-41af-91ab-2d7cd011db47"),
	// 		UserAssignedIdentities: map[string]*armoperationalinsights.UserIdentityProperties{
	// 			"/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": &armoperationalinsights.UserIdentityProperties{
	// 				ClientID: to.Ptr("eb3a943d-6b12-48a6-b585-ac2316e15ab2"),
	// 				PrincipalID: to.Ptr("b31776d4-ee80-4860-9433-ec0101be1891"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armoperationalinsights.ClusterProperties{
	// 		AssociatedWorkspaces: []*armoperationalinsights.AssociatedWorkspace{
	// 			{
	// 				AssociateDate: to.Ptr("Tue, 07 Jul 2020 07:35:51 GMT"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs1"),
	// 				WorkspaceID: to.Ptr("942bdefd-e6c9-411c-ac69-70ffad564363"),
	// 				WorkspaceName: to.Ptr("testWs1"),
	// 			},
	// 			{
	// 				AssociateDate: to.Ptr("Mon, 13 Jan 2020 16:03:39 GMT"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/testRG/providers/microsoft.operationalinsights/workspaces/testWs2"),
	// 				WorkspaceID: to.Ptr("c7edb8f8-67f7-41f2-bddb-aecf22507e3f"),
	// 				WorkspaceName: to.Ptr("testWs2"),
	// 		}},
	// 		BillingType: to.Ptr(armoperationalinsights.BillingTypeCluster),
	// 		CapacityReservationProperties: &armoperationalinsights.CapacityReservationProperties{
	// 			LastSKUUpdate: to.Ptr("Thu, 01 Jan 1970 00:00:00 GMT"),
	// 			MinCapacity: to.Ptr[int64](500),
	// 		},
	// 		ClusterID: to.Ptr("1c3f2afc-9ec1-4e74-bfe8-022b83bcd7b1"),
	// 		CreatedDate: to.Ptr("Mon, 13 Jan 2020 14:40:33 GMT"),
	// 		KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
	// 			KeyName: to.Ptr("aztest2170cert"),
	// 			KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
	// 			KeyVersion: to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
	// 		},
	// 		LastModifiedDate: to.Ptr("Sun, 04 Jan 2020 17:10:56 GMT"),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.ClusterEntityStatusSucceeded),
	// 	},
	// 	SKU: &armoperationalinsights.ClusterSKU{
	// 		Name: to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
	// 		Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
	// 	},
	// }
}
