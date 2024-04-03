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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_CreateOrUpdate.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", armcompute.VirtualMachineRunCommand{
		Location: to.Ptr("West US"),
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			AsyncExecution:           to.Ptr(false),
			ErrorBlobManagedIdentity: &armcompute.RunCommandManagedIdentity{},
			ErrorBlobURI:             to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
			OutputBlobManagedIdentity: &armcompute.RunCommandManagedIdentity{
				ClientID: to.Ptr("22d35efb-0c99-4041-8c5b-6d24db33a69a"),
			},
			OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
			Parameters: []*armcompute.RunCommandInputParameter{
				{
					Name:  to.Ptr("param1"),
					Value: to.Ptr("value1"),
				},
				{
					Name:  to.Ptr("param2"),
					Value: to.Ptr("value2"),
				}},
			RunAsPassword: to.Ptr("<runAsPassword>"),
			RunAsUser:     to.Ptr("user1"),
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
				ScriptURIManagedIdentity: &armcompute.RunCommandManagedIdentity{
					ObjectID: to.Ptr("4231e4d2-33e4-4e23-96b2-17888afa6072"),
				},
			},
			TimeoutInSeconds:                to.Ptr[int32](3600),
			TreatFailureAsDeploymentFailure: to.Ptr(true),
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
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Update.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().BeginUpdate(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", armcompute.VirtualMachineRunCommandUpdate{
		Properties: &armcompute.VirtualMachineRunCommandProperties{
			Source: &armcompute.VirtualMachineRunCommandScriptSource{
				ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
				ScriptURIManagedIdentity: &armcompute.RunCommandManagedIdentity{
					ObjectID: to.Ptr("4231e4d2-33e4-4e23-96b2-17888afa6072"),
				},
			},
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
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			ScriptURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Delete.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().BeginDelete(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Get.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().Get(ctx, "myResourceGroup", "myvmScaleSet", "0", "myRunCommand", &armcompute.VirtualMachineScaleSetVMRunCommandsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineRunCommand = armcompute.VirtualMachineRunCommand{
	// 	Name: to.Ptr("myRunCommand"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineRunCommandProperties{
	// 		AsyncExecution: to.Ptr(false),
	// 		ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
	// 		OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
	// 		Parameters: []*armcompute.RunCommandInputParameter{
	// 			{
	// 				Name: to.Ptr("param1"),
	// 				Value: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("param2"),
	// 				Value: to.Ptr("value2"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RunAsUser: to.Ptr("user1"),
	// 		Source: &armcompute.VirtualMachineRunCommandScriptSource{
	// 			ScriptURI: to.Ptr("https://gist.githubusercontent.com/myusername/75fd3634w7511116063c60bcc50bee0/raw/04a4c68ac9e1d36asfasdc64bd1d889b104c7abdb8/HelloWorld.ps1"),
	// 		},
	// 		TimeoutInSeconds: to.Ptr[int32](3600),
	// 		TreatFailureAsDeploymentFailure: to.Ptr(false),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_List.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetVMRunCommandsClient().NewListPager("myResourceGroup", "myvmScaleSet", "0", &armcompute.VirtualMachineScaleSetVMRunCommandsClientListOptions{Expand: nil})
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
		// page.VirtualMachineRunCommandsListResult = armcompute.VirtualMachineRunCommandsListResult{
		// 	Value: []*armcompute.VirtualMachineRunCommand{
		// 		{
		// 			Name: to.Ptr("myRunCommand"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines/runCommands"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/myvmScaleSet/virtualMachines/0/runCommands/myRunCommand"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcompute.VirtualMachineRunCommandProperties{
		// 				AsyncExecution: to.Ptr(false),
		// 				ErrorBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
		// 				OutputBlobURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
		// 				Parameters: []*armcompute.RunCommandInputParameter{
		// 					{
		// 						Name: to.Ptr("param1"),
		// 						Value: to.Ptr("value1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("param2"),
		// 						Value: to.Ptr("value2"),
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RunAsUser: to.Ptr("user1"),
		// 				Source: &armcompute.VirtualMachineRunCommandScriptSource{
		// 					Script: to.Ptr("Write-Host Hello World! ; Remove-Item C:	est	estFile.txt"),
		// 				},
		// 				TimeoutInSeconds: to.Ptr[int32](0),
		// 				TreatFailureAsDeploymentFailure: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}
