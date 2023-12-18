//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GetAvailabilitySet.json
func ExampleAvailabilitySetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Get(ctx, "testrg", "HRAvailabilitySet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armscvmm.AvailabilitySet{
	// 	Name: to.Ptr("HRAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.ScVmm/AvailabilitySets"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/AvailabilitySets/HRAvailabilitySet"),
	// 	Location: to.Ptr("East US"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.AvailabilitySetProperties{
	// 		AvailabilitySetName: to.Ptr("hr-avset"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/CreateAvailabilitySet.json
func ExampleAvailabilitySetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvailabilitySetsClient().BeginCreateOrUpdate(ctx, "testrg", "HRAvailabilitySet", armscvmm.AvailabilitySet{
		Location: to.Ptr("East US"),
		ExtendedLocation: &armscvmm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
			Type: to.Ptr("customLocation"),
		},
		Properties: &armscvmm.AvailabilitySetProperties{
			AvailabilitySetName: to.Ptr("hr-avset"),
			VmmServerID:         to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
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
	// res.AvailabilitySet = armscvmm.AvailabilitySet{
	// 	Name: to.Ptr("HRAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.ScVmm/AvailabilitySets"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/AvailabilitySets/HRAvailabilitySet"),
	// 	Location: to.Ptr("East US"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.AvailabilitySetProperties{
	// 		AvailabilitySetName: to.Ptr("hr-avset"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteAvailabilitySet.json
func ExampleAvailabilitySetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvailabilitySetsClient().BeginDelete(ctx, "testrg", "HRAvailabilitySet", &armscvmm.AvailabilitySetsClientBeginDeleteOptions{Force: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/UpdateAvailabilitySet.json
func ExampleAvailabilitySetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvailabilitySetsClient().BeginUpdate(ctx, "testrg", "HRAvailabilitySet", armscvmm.ResourcePatch{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.AvailabilitySet = armscvmm.AvailabilitySet{
	// 	Name: to.Ptr("HRAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.ScVmm/AvailabilitySets"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/AvailabilitySets/HRAvailabilitySet"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.AvailabilitySetProperties{
	// 		AvailabilitySetName: to.Ptr("hr-avset"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/ListAvailabilitySetsByResourceGroup.json
func ExampleAvailabilitySetsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListByResourceGroupPager("testrg", nil)
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
		// page.AvailabilitySetListResult = armscvmm.AvailabilitySetListResult{
		// 	Value: []*armscvmm.AvailabilitySet{
		// 		{
		// 			Name: to.Ptr("HRAvailabilitySet"),
		// 			Type: to.Ptr("Microsoft.ScVmm/AvailabilitySets"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/AvailabilitySets/HRAvailabilitySet"),
		// 			Location: to.Ptr("East US"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.AvailabilitySetProperties{
		// 				AvailabilitySetName: to.Ptr("hr-avset"),
		// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 				VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/ListAvailabilitySetsBySubscription.json
func ExampleAvailabilitySetsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListBySubscriptionPager(nil)
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
		// page.AvailabilitySetListResult = armscvmm.AvailabilitySetListResult{
		// 	Value: []*armscvmm.AvailabilitySet{
		// 		{
		// 			Name: to.Ptr("HRAvailabilitySet"),
		// 			Type: to.Ptr("Microsoft.ScVmm/AvailabilitySets"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/AvailabilitySets/HRAvailabilitySet"),
		// 			Location: to.Ptr("East US"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.AvailabilitySetProperties{
		// 				AvailabilitySetName: to.Ptr("hr-avset"),
		// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 				VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
		// 			},
		// 	}},
		// }
	}
}
