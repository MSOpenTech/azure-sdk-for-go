//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_ListBySub.json
func ExampleFleetsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListBySubscriptionPager(nil)
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
		// page.FleetListResult = armcontainerservicefleet.FleetListResult{
		// 	Value: []*armcontainerservicefleet.Fleet{
		// 		{
		// 			Name: to.Ptr("fleet-1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/fleets"),
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet1"),
		// 			SystemData: &armcontainerservicefleet.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				CreatedBy: to.Ptr("someUser"),
		// 				CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("someOtherUser"),
		// 				LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"archv2": to.Ptr(""),
		// 				"tier": to.Ptr("production"),
		// 			},
		// 			ETag: to.Ptr("23ujdflewrj3="),
		// 			Properties: &armcontainerservicefleet.FleetProperties{
		// 				HubProfile: &armcontainerservicefleet.FleetHubProfile{
		// 					DNSPrefix: to.Ptr("dnsprefix1"),
		// 					Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
		// 					KubernetesVersion: to.Ptr("1.22.4"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_ListByResourceGroup.json
func ExampleFleetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.FleetListResult = armcontainerservicefleet.FleetListResult{
		// 	Value: []*armcontainerservicefleet.Fleet{
		// 		{
		// 			Name: to.Ptr("fleet1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/fleets"),
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
		// 			SystemData: &armcontainerservicefleet.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				CreatedBy: to.Ptr("someUser"),
		// 				CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("someOtherUser"),
		// 				LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"archv2": to.Ptr(""),
		// 				"tier": to.Ptr("production"),
		// 			},
		// 			ETag: to.Ptr("23ujdflewrj3="),
		// 			Properties: &armcontainerservicefleet.FleetProperties{
		// 				HubProfile: &armcontainerservicefleet.FleetHubProfile{
		// 					DNSPrefix: to.Ptr("dnsprefix1"),
		// 					Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
		// 					KubernetesVersion: to.Ptr("1.22.4"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_Get.json
func ExampleFleetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().Get(ctx, "rg1", "fleet1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fleet = armcontainerservicefleet.Fleet{
	// 	Name: to.Ptr("fleet-1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/fleets"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
	// 	SystemData: &armcontainerservicefleet.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		CreatedBy: to.Ptr("someUser"),
	// 		CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("someOtherUser"),
	// 		LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"archv2": to.Ptr(""),
	// 		"tier": to.Ptr("production"),
	// 	},
	// 	ETag: to.Ptr("23ujdflewrj3="),
	// 	Properties: &armcontainerservicefleet.FleetProperties{
	// 		HubProfile: &armcontainerservicefleet.FleetHubProfile{
	// 			DNSPrefix: to.Ptr("dnsprefix1"),
	// 			Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 			KubernetesVersion: to.Ptr("1.22.4"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_CreateOrUpdate.json
func ExampleFleetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginCreateOrUpdate(ctx, "rg1", "fleet1", armcontainerservicefleet.Fleet{
		Location: to.Ptr("East US"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Properties: &armcontainerservicefleet.FleetProperties{
			HubProfile: &armcontainerservicefleet.FleetHubProfile{
				DNSPrefix: to.Ptr("dnsprefix1"),
			},
		},
	}, &armcontainerservicefleet.FleetsClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.Fleet = armcontainerservicefleet.Fleet{
	// 	Name: to.Ptr("fleet-1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/fleets"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
	// 	SystemData: &armcontainerservicefleet.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		CreatedBy: to.Ptr("someUser"),
	// 		CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("someOtherUser"),
	// 		LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"archv2": to.Ptr(""),
	// 		"tier": to.Ptr("production"),
	// 	},
	// 	ETag: to.Ptr("23ujdflewrj3="),
	// 	Properties: &armcontainerservicefleet.FleetProperties{
	// 		HubProfile: &armcontainerservicefleet.FleetHubProfile{
	// 			DNSPrefix: to.Ptr("dnsprefix1"),
	// 			Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 			KubernetesVersion: to.Ptr("1.22.4"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_PatchTags.json
func ExampleFleetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().Update(ctx, "rg1", "fleet1", armcontainerservicefleet.FleetPatch{
		Tags: map[string]*string{
			"env":  to.Ptr("prod"),
			"tier": to.Ptr("secure"),
		},
	}, &armcontainerservicefleet.FleetsClientUpdateOptions{IfMatch: to.Ptr("dfjkwelr7384")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fleet = armcontainerservicefleet.Fleet{
	// 	Name: to.Ptr("fleet-1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/fleets"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
	// 	SystemData: &armcontainerservicefleet.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		CreatedBy: to.Ptr("someUser"),
	// 		CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("someOtherUser"),
	// 		LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"env": to.Ptr("prod"),
	// 		"tier": to.Ptr("secure"),
	// 	},
	// 	ETag: to.Ptr("23ujdflewrj3="),
	// 	Properties: &armcontainerservicefleet.FleetProperties{
	// 		HubProfile: &armcontainerservicefleet.FleetHubProfile{
	// 			DNSPrefix: to.Ptr("dnsprefix1"),
	// 			Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 			KubernetesVersion: to.Ptr("1.22.4"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_Delete.json
func ExampleFleetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginDelete(ctx, "rg1", "fleet1", &armcontainerservicefleet.FleetsClientBeginDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de14cb8751b978b1877597b13292818e80f8c661/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/examples/Fleets_ListCredentialsResult.json
func ExampleFleetsClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().ListCredentials(ctx, "rg1", "fleet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FleetCredentialResults = armcontainerservicefleet.FleetCredentialResults{
	// 	Kubeconfigs: []*armcontainerservicefleet.FleetCredentialResult{
	// 		{
	// 			Name: to.Ptr("credentialName1"),
	// 			Value: []byte("Y3JlZGVudGlhbFZhbHVlMQ=="),
	// 	}},
	// }
}
