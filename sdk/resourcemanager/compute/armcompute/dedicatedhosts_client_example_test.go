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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_CreateOrUpdate.json
func ExampleDedicatedHostsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", "myDedicatedHost", armcompute.DedicatedHost{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		Properties: &armcompute.DedicatedHostProperties{
			PlatformFaultDomain: to.Ptr[int32](1),
		},
		SKU: &armcompute.SKU{
			Name: to.Ptr("DSv3-Type1"),
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
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Name: to.Ptr("myDedicatedHost"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostProperties{
	// 		AutoReplaceOnFailure: to.Ptr(false),
	// 		HostID: to.Ptr("{GUID}"),
	// 		LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("DSv3-Type1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Update_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_BeginUpdate_dedicatedHostUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", armcompute.DedicatedHostUpdate{
		Tags: map[string]*string{
			"key8813": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		},
		Properties: &armcompute.DedicatedHostProperties{
			AutoReplaceOnFailure: to.Ptr(true),
			InstanceView: &armcompute.DedicatedHostInstanceView{
				AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
					AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
						{
							Count:  to.Ptr[float64](26),
							VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
						}},
				},
				Statuses: []*armcompute.InstanceViewStatus{
					{
						Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
						DisplayStatus: to.Ptr("aaaaaa"),
						Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
						Message:       to.Ptr("a"),
						Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
					}},
			},
			LicenseType:         to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
			PlatformFaultDomain: to.Ptr[int32](1),
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
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Name: to.Ptr("myDedicatedHost"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcompute.DedicatedHostProperties{
	// 		AutoReplaceOnFailure: to.Ptr(true),
	// 		HostID: to.Ptr("{GUID}"),
	// 		InstanceView: &armcompute.DedicatedHostInstanceView{
	// 			AssetID: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 			AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 				AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 					{
	// 						Count: to.Ptr[float64](26),
	// 						VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 				}},
	// 			},
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 					DisplayStatus: to.Ptr("aaaaaa"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Message: to.Ptr("a"),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 			}},
	// 		},
	// 		LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.526Z"); return t}()),
	// 		VirtualMachines: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("aaaa"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("DSv3-Type1"),
	// 		Capacity: to.Ptr[int64](7),
	// 		Tier: to.Ptr("aaa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Update_MinimumSet_Gen.json
func ExampleDedicatedHostsClient_BeginUpdate_dedicatedHostUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginUpdate(ctx, "rgcompute", "aa", "aaaaaaaaaaaaaaaaaaaaaaaaaa", armcompute.DedicatedHostUpdate{}, nil)
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
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armcompute.SKU{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Update_Resize.json
func ExampleDedicatedHostsClient_BeginUpdate_dedicatedHostUpdateResize() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", armcompute.DedicatedHostUpdate{
		SKU: &armcompute.SKU{
			Name: to.Ptr("DSv3-Type1"),
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
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Name: to.Ptr("myDedicatedHost"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcompute.DedicatedHostProperties{
	// 		AutoReplaceOnFailure: to.Ptr(true),
	// 		HostID: to.Ptr("{GUID}"),
	// 		LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.526Z"); return t}()),
	// 		VirtualMachines: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("aaaa"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("DSv3-Type1"),
	// 		Capacity: to.Ptr[int64](7),
	// 		Tier: to.Ptr("aaa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Delete_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_BeginDelete_dedicatedHostDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginDelete(ctx, "rgcompute", "aaaaaa", "aaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Delete_MinimumSet_Gen.json
func ExampleDedicatedHostsClient_BeginDelete_dedicatedHostDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginDelete(ctx, "rgcompute", "aaaaaaaaaaaaaaa", "aaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Get.json
func ExampleDedicatedHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", "myHost", &armcompute.DedicatedHostsClientGetOptions{Expand: to.Ptr(armcompute.InstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHost = armcompute.DedicatedHost{
	// 	Name: to.Ptr("myHost"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myHost"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("HR"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostProperties{
	// 		AutoReplaceOnFailure: to.Ptr(true),
	// 		HostID: to.Ptr("{GUID}"),
	// 		InstanceView: &armcompute.DedicatedHostInstanceView{
	// 			AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 			AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 				AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 					{
	// 						Count: to.Ptr[float64](10),
	// 						VMSize: to.Ptr("Standard_A1"),
	// 				}},
	// 			},
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				},
	// 				{
	// 					Code: to.Ptr("HealthState/available"),
	// 					DisplayStatus: to.Ptr("Host available"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			}},
	// 		},
	// 		PlatformFaultDomain: to.Ptr[int32](1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-27T01:02:38.3138469+00:00"); return t}()),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-27T01:02:38.3138469+00:00"); return t}()),
	// 		VirtualMachines: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("DSv3-Type1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_ListByHostGroup_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_NewListByHostGroupPager_dedicatedHostListByHostGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostsClient().NewListByHostGroupPager("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
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
		// page.DedicatedHostListResult = armcompute.DedicatedHostListResult{
		// 	Value: []*armcompute.DedicatedHost{
		// 		{
		// 			Name: to.Ptr("myDedicatedHost"),
		// 			Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.DedicatedHostProperties{
		// 				AutoReplaceOnFailure: to.Ptr(true),
		// 				HostID: to.Ptr("{GUID}"),
		// 				InstanceView: &armcompute.DedicatedHostInstanceView{
		// 					AssetID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
		// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
		// 							{
		// 								Count: to.Ptr[float64](26),
		// 								VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 						}},
		// 					},
		// 					Statuses: []*armcompute.InstanceViewStatus{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 							DisplayStatus: to.Ptr("aaaaaa"),
		// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 							Message: to.Ptr("a"),
		// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 					}},
		// 				},
		// 				LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
		// 				PlatformFaultDomain: to.Ptr[int32](1),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.526Z"); return t}()),
		// 				VirtualMachines: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("aaaa"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("DSv3-Type1"),
		// 				Capacity: to.Ptr[int64](7),
		// 				Tier: to.Ptr("aaa"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_ListByHostGroup_MinimumSet_Gen.json
func ExampleDedicatedHostsClient_NewListByHostGroupPager_dedicatedHostListByHostGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostsClient().NewListByHostGroupPager("rgcompute", "aaaa", nil)
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
		// page.DedicatedHostListResult = armcompute.DedicatedHostListResult{
		// 	Value: []*armcompute.DedicatedHost{
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myHost"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armcompute.SKU{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_Restart.json
func ExampleDedicatedHostsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginRestart(ctx, "myResourceGroup", "myDedicatedHostGroup", "myHost", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/dedicatedHostExamples/DedicatedHost_ListAvailableSizes.json
func ExampleDedicatedHostsClient_NewListAvailableSizesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostsClient().NewListAvailableSizesPager("myResourceGroup", "myDedicatedHostGroup", "myHost", nil)
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
		// page.DedicatedHostSizeListResult = armcompute.DedicatedHostSizeListResult{
		// 	Value: []*string{
		// 		to.Ptr("Dsv3-Type1"),
		// 		to.Ptr("Esv3-Type1")},
		// 	}
	}
}
