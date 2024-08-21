//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_CreateOrUpdate.json
func ExampleCapacityReservationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myCapacityReservationGroup", "myCapacityReservation", armcompute.CapacityReservation{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_DS1_v2"),
			Capacity: to.Ptr[int64](4),
		},
		Zones: []*string{
			to.Ptr("1")},
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
	// res.CapacityReservation = armcompute.CapacityReservation{
	// 	Name: to.Ptr("myCapacityReservation"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.CapacityReservationProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		ReservationID: to.Ptr("{GUID}"),
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_DS1_v2"),
	// 		Capacity: to.Ptr[int64](4),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_Update_MaximumSet_Gen.json
func ExampleCapacityReservationsClient_BeginUpdate_capacityReservationUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaaa", "aaaaaaaaaaaaaaaaaaa", armcompute.CapacityReservationUpdate{
		Tags: map[string]*string{
			"key4974": to.Ptr("aaaaaaaaaaaaaaaa"),
		},
		Properties: &armcompute.CapacityReservationProperties{
			InstanceView: &armcompute.CapacityReservationInstanceView{
				Statuses: []*armcompute.InstanceViewStatus{
					{
						Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
						DisplayStatus: to.Ptr("aaaaaa"),
						Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
						Message:       to.Ptr("a"),
						Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
					}},
				UtilizationInfo: &armcompute.CapacityReservationUtilization{},
			},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_DS1_v2"),
			Capacity: to.Ptr[int64](7),
			Tier:     to.Ptr("aaa"),
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
	// res.CapacityReservation = armcompute.CapacityReservation{
	// 	Name: to.Ptr("myCapacityReservation"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcompute.CapacityReservationProperties{
	// 		InstanceView: &armcompute.CapacityReservationInstanceView{
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 					DisplayStatus: to.Ptr("aaaaaa"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Message: to.Ptr("a"),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 			}},
	// 			UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 				VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 					{
	// 						ID: to.Ptr("aaaa"),
	// 				}},
	// 			},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		ReservationID: to.Ptr("{GUID}"),
	// 		VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("aaaa"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_DS1_v2"),
	// 		Capacity: to.Ptr[int64](4),
	// 		Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_Update_MinimumSet_Gen.json
func ExampleCapacityReservationsClient_BeginUpdate_capacityReservationUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aaa", armcompute.CapacityReservationUpdate{}, nil)
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
	// res.CapacityReservation = armcompute.CapacityReservation{
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armcompute.SKU{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_Delete_MaximumSet_Gen.json
func ExampleCapacityReservationsClient_BeginDelete_capacityReservationDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginDelete(ctx, "rgcompute", "aaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_Delete_MinimumSet_Gen.json
func ExampleCapacityReservationsClient_BeginDelete_capacityReservationDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginDelete(ctx, "rgcompute", "aaa", "aaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_Get.json
func ExampleCapacityReservationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationsClient().Get(ctx, "myResourceGroup", "myCapacityReservationGroup", "myCapacityReservation", &armcompute.CapacityReservationsClientGetOptions{Expand: to.Ptr(armcompute.CapacityReservationInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapacityReservation = armcompute.CapacityReservation{
	// 	Name: to.Ptr("myCapacityReservation"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.CapacityReservationProperties{
	// 		InstanceView: &armcompute.CapacityReservationInstanceView{
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			}},
	// 			UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 				CurrentCapacity: to.Ptr[int32](5),
	// 				VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 				}},
	// 			},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		ReservationID: to.Ptr("{GUID}"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
	// 		VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_DS1_v2"),
	// 		Capacity: to.Ptr[int64](4),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_ListByReservationGroup.json
func ExampleCapacityReservationsClient_NewListByCapacityReservationGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacityReservationsClient().NewListByCapacityReservationGroupPager("myResourceGroup", "myCapacityReservationGroup", nil)
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
		// page.CapacityReservationListResult = armcompute.CapacityReservationListResult{
		// 	Value: []*armcompute.CapacityReservation{
		// 		{
		// 			Name: to.Ptr("{capacityReservationName}"),
		// 			Type: to.Ptr("Microsoft.Compute/CapacityReservations"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/CapacityReservation/{capacityReservationName}"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("HR"),
		// 			},
		// 			Properties: &armcompute.CapacityReservationProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
		// 				ReservationID: to.Ptr("{GUID}"),
		// 				VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM3"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Standard_DS1_v2"),
		// 				Capacity: to.Ptr[int64](4),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 			},
		// 			{
		// 				Name: to.Ptr("{capacityReservationName}"),
		// 				Type: to.Ptr("Microsoft.Compute/CapacityReservations"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/CapacityReservation/{capacityReservationName}"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("HR"),
		// 				},
		// 				Properties: &armcompute.CapacityReservationProperties{
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.313Z"); return t}()),
		// 					ReservationID: to.Ptr("{GUID}"),
		// 					VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM4"),
		// 					}},
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_A1_v2"),
		// 					Capacity: to.Ptr[int64](4),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1")},
		// 			}},
		// 		}
	}
}
