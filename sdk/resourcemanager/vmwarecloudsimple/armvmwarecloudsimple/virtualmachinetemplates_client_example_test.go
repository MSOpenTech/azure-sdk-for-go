//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListVirtualMachineTemplates.json
func ExampleVirtualMachineTemplatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineTemplatesClient().NewListPager("myPrivateCloud", "westus2", "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26", nil)
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
		// page.VirtualMachineTemplateListResponse = armvmwarecloudsimple.VirtualMachineTemplateListResponse{
		// 	Value: []*armvmwarecloudsimple.VirtualMachineTemplate{
		// 		{
		// 			Name: to.Ptr("centos-template"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachineTemplates"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualMachineTemplateProperties{
		// 				Path: to.Ptr("Datacenter/Workload VMs"),
		// 				AmountOfRAM: to.Ptr[int32](4096),
		// 				Controllers: []*armvmwarecloudsimple.VirtualDiskController{
		// 					{
		// 						Name: to.Ptr("SCSI controller 0"),
		// 						Type: to.Ptr("SCSI"),
		// 						ID: to.Ptr("1000"),
		// 						SubType: to.Ptr("LSI_PARALEL"),
		// 				}},
		// 				Disks: []*armvmwarecloudsimple.VirtualDisk{
		// 					{
		// 						ControllerID: to.Ptr("1000"),
		// 						IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
		// 						TotalSize: to.Ptr[int32](10485760),
		// 						VirtualDiskID: to.Ptr("2000"),
		// 						VirtualDiskName: to.Ptr("Hard disk 1"),
		// 				}},
		// 				GuestOS: to.Ptr("Other (32-bit)"),
		// 				GuestOSType: to.Ptr("other"),
		// 				Nics: []*armvmwarecloudsimple.VirtualNic{
		// 					{
		// 						MacAddress: to.Ptr("00:50:56:a6:7e:93"),
		// 						Network: &armvmwarecloudsimple.VirtualNetwork{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
		// 						},
		// 						NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
		// 						PowerOnBoot: to.Ptr(true),
		// 						VirtualNicID: to.Ptr("4000"),
		// 						VirtualNicName: to.Ptr("Network adapter 1"),
		// 				}},
		// 				NumberOfCores: to.Ptr[int32](2),
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 				Vmwaretools: to.Ptr("10309"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachineTemplate.json
func ExampleVirtualMachineTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineTemplatesClient().Get(ctx, "westus2", "myPrivateCloud", "vm-34", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armvmwarecloudsimple.VirtualMachineTemplate{
	// 	Name: to.Ptr("centos-template"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualMachineTemplates"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armvmwarecloudsimple.VirtualMachineTemplateProperties{
	// 		Path: to.Ptr("Datacenter/Workload VMs"),
	// 		AmountOfRAM: to.Ptr[int32](4096),
	// 		Controllers: []*armvmwarecloudsimple.VirtualDiskController{
	// 			{
	// 				Name: to.Ptr("SCSI controller 0"),
	// 				Type: to.Ptr("SCSI"),
	// 				ID: to.Ptr("1000"),
	// 				SubType: to.Ptr("LSI_PARALEL"),
	// 		}},
	// 		Disks: []*armvmwarecloudsimple.VirtualDisk{
	// 			{
	// 				ControllerID: to.Ptr("1000"),
	// 				IndependenceMode: to.Ptr(armvmwarecloudsimple.DiskIndependenceModePersistent),
	// 				TotalSize: to.Ptr[int32](10485760),
	// 				VirtualDiskID: to.Ptr("2000"),
	// 				VirtualDiskName: to.Ptr("Hard disk 1"),
	// 		}},
	// 		GuestOS: to.Ptr("Other (32-bit)"),
	// 		GuestOSType: to.Ptr("other"),
	// 		Nics: []*armvmwarecloudsimple.VirtualNic{
	// 			{
	// 				MacAddress: to.Ptr("00:50:56:a6:7e:93"),
	// 				Network: &armvmwarecloudsimple.VirtualNetwork{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
	// 				},
	// 				NicType: to.Ptr(armvmwarecloudsimple.NICTypeE1000),
	// 				PowerOnBoot: to.Ptr(true),
	// 				VirtualNicID: to.Ptr("4000"),
	// 				VirtualNicName: to.Ptr("Network adapter 1"),
	// 		}},
	// 		NumberOfCores: to.Ptr[int32](2),
	// 		PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
	// 		Vmwaretools: to.Ptr("10309"),
	// 	},
	// }
}
