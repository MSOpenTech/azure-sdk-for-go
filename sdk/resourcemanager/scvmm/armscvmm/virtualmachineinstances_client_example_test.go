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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GetVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineInstancesClient().Get(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineInstance = armscvmm.VirtualMachineInstance{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ScVmm/VirtualMachineInstances"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.ScVmm/virtualMachineInstances/default"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineInstanceProperties{
	// 		HardwareProfile: &armscvmm.HardwareProfile{
	// 			CPUCount: to.Ptr[int32](4),
	// 			MemoryMB: to.Ptr[int32](4196),
	// 		},
	// 		InfrastructureProfile: &armscvmm.InfrastructureProfile{
	// 			BiosGUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			CloudID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
	// 			TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
	// 			UUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 		},
	// 		OSProfile: &armscvmm.OsProfileForVMInstance{
	// 			ComputerName: to.Ptr("DemoVM"),
	// 			OSSKU: to.Ptr("Windows Server 2022"),
	// 			OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 			OSVersion: to.Ptr("10.0.10101"),
	// 		},
	// 		PowerState: to.Ptr("Running"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/CreateVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginCreateOrUpdate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginCreateOrUpdateOptions{Body: &armscvmm.VirtualMachineInstance{
		ExtendedLocation: &armscvmm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
			Type: to.Ptr("customLocation"),
		},
		Properties: &armscvmm.VirtualMachineInstanceProperties{
			HardwareProfile: &armscvmm.HardwareProfile{
				CPUCount: to.Ptr[int32](4),
				MemoryMB: to.Ptr[int32](4196),
			},
			InfrastructureProfile: &armscvmm.InfrastructureProfile{
				CloudID:     to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
				TemplateID:  to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
				VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
			},
		},
	},
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
	// res.VirtualMachineInstance = armscvmm.VirtualMachineInstance{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ScVmm/VirtualMachineInstances"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.ScVmm/virtualMachineInstances/default"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineInstanceProperties{
	// 		HardwareProfile: &armscvmm.HardwareProfile{
	// 			CPUCount: to.Ptr[int32](4),
	// 			MemoryMB: to.Ptr[int32](4196),
	// 		},
	// 		InfrastructureProfile: &armscvmm.InfrastructureProfile{
	// 			BiosGUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			CloudID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
	// 			TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
	// 			UUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 		},
	// 		OSProfile: &armscvmm.OsProfileForVMInstance{
	// 			ComputerName: to.Ptr("DemoVM"),
	// 			OSSKU: to.Ptr("Windows Server 2022"),
	// 			OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 			OSVersion: to.Ptr("10.0.10101"),
	// 		},
	// 		PowerState: to.Ptr("Running"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/UpdateVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginUpdate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginUpdateOptions{Body: &armscvmm.VirtualMachineInstanceUpdate{
		Properties: &armscvmm.VirtualMachineInstanceUpdateProperties{
			HardwareProfile: &armscvmm.HardwareProfileUpdate{
				CPUCount: to.Ptr[int32](4),
				MemoryMB: to.Ptr[int32](4196),
			},
		},
	},
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
	// res.VirtualMachineInstance = armscvmm.VirtualMachineInstance{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ScVmm/VirtualMachineInstances"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.ScVmm/virtualMachineInstances/default"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineInstanceProperties{
	// 		HardwareProfile: &armscvmm.HardwareProfile{
	// 			CPUCount: to.Ptr[int32](4),
	// 			MemoryMB: to.Ptr[int32](4196),
	// 		},
	// 		InfrastructureProfile: &armscvmm.InfrastructureProfile{
	// 			BiosGUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			CloudID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
	// 			TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
	// 			UUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 		},
	// 		OSProfile: &armscvmm.OsProfileForVMInstance{
	// 			ComputerName: to.Ptr("DemoVM"),
	// 			OSSKU: to.Ptr("Windows Server 2022"),
	// 			OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 			OSVersion: to.Ptr("10.0.10101"),
	// 		},
	// 		PowerState: to.Ptr("Running"),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginDelete(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginDeleteOptions{Force: nil,
		DeleteFromHost: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/ListVirtualMachineInstances.json
func ExampleVirtualMachineInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineInstancesClient().NewListPager("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", nil)
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
		// page.VirtualMachineInstanceListResult = armscvmm.VirtualMachineInstanceListResult{
		// 	Value: []*armscvmm.VirtualMachineInstance{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.ScVmm/VirtualMachineInstances"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.ScVmm/virtualMachineInstances/default"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.VirtualMachineInstanceProperties{
		// 				HardwareProfile: &armscvmm.HardwareProfile{
		// 					CPUCount: to.Ptr[int32](4),
		// 					MemoryMB: to.Ptr[int32](4196),
		// 				},
		// 				InfrastructureProfile: &armscvmm.InfrastructureProfile{
		// 					BiosGUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 					CloudID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud"),
		// 					TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate"),
		// 					UUID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 					VmmServerID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
		// 				},
		// 				OSProfile: &armscvmm.OsProfileForVMInstance{
		// 					ComputerName: to.Ptr("DemoVM"),
		// 					OSSKU: to.Ptr("Windows Server 2022"),
		// 					OSType: to.Ptr(armscvmm.OsTypeWindows),
		// 					OSVersion: to.Ptr("10.0.10101"),
		// 				},
		// 				PowerState: to.Ptr("Running"),
		// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/StopVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginStop(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginStopOptions{Body: &armscvmm.StopVirtualMachineOptions{
		SkipShutdown: to.Ptr(armscvmm.SkipShutdownTrue),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/StartVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginStart(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/RestartVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginRestart(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/CreateCheckpointVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginCreateCheckpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginCreateCheckpoint(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginCreateCheckpointOptions{Body: &armscvmm.VirtualMachineCreateCheckpoint{
		Name:        to.Ptr("Demo Checkpoint name"),
		Description: to.Ptr("Demo Checkpoint description"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteCheckpointVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginDeleteCheckpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginDeleteCheckpoint(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginDeleteCheckpointOptions{Body: &armscvmm.VirtualMachineDeleteCheckpoint{
		ID: to.Ptr("Demo CheckpointID"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c74e0b0b6aa869fdd6f6d76984fc2b2610bc64a8/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/RestoreCheckpointVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_BeginRestoreCheckpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginRestoreCheckpoint(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", &armscvmm.VirtualMachineInstancesClientBeginRestoreCheckpointOptions{Body: &armscvmm.VirtualMachineRestoreCheckpoint{
		ID: to.Ptr("Demo CheckpointID"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
