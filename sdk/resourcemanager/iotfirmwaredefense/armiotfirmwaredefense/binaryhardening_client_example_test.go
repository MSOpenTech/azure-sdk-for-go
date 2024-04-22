//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/BinaryHardening_ListByFirmware_MaximumSet_Gen.json
func ExampleBinaryHardeningClient_NewListByFirmwarePager_binaryHardeningListByFirmwareMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBinaryHardeningClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
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
		// page.BinaryHardeningListResult = armiotfirmwaredefense.BinaryHardeningListResult{
		// 	Value: []*armiotfirmwaredefense.BinaryHardeningResource{
		// 		{
		// 			Name: to.Ptr("cd31ca40-6772-44f8-9180-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/binaryHardeningResults"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/binaryHardeningResults/cd31ca40-6772-44f8-9180-000000000000"),
		// 			Properties: &armiotfirmwaredefense.BinaryHardeningResult{
		// 				Architecture: to.Ptr("ARM"),
		// 				BinaryHardeningID: to.Ptr("cd31ca40-6772-44f8-9180-000000000000"),
		// 				Class: to.Ptr("Bit32"),
		// 				Features: &armiotfirmwaredefense.BinaryHardeningFeatures{
		// 					Canary: to.Ptr(false),
		// 					Nx: to.Ptr(true),
		// 					Pie: to.Ptr(false),
		// 					Relro: to.Ptr(true),
		// 					Stripped: to.Ptr(true),
		// 				},
		// 				FilePath: to.Ptr("/path/to/squashfs-root/bin/bcm_boot_launcher"),
		// 				Rpath: to.Ptr("no"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5924e71d-5f38-444e-8b08-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/binaryHardeningResults"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/binaryHardeningResults/5924e71d-5f38-444e-8b08-000000000000"),
		// 			Properties: &armiotfirmwaredefense.BinaryHardeningResult{
		// 				Architecture: to.Ptr("ARM"),
		// 				BinaryHardeningID: to.Ptr("5924e71d-5f38-444e-8b08-000000000000"),
		// 				Class: to.Ptr("Bit32"),
		// 				Features: &armiotfirmwaredefense.BinaryHardeningFeatures{
		// 					Canary: to.Ptr(false),
		// 					Nx: to.Ptr(true),
		// 					Pie: to.Ptr(false),
		// 					Relro: to.Ptr(true),
		// 					Stripped: to.Ptr(true),
		// 				},
		// 				FilePath: to.Ptr("/path/to/squashfs-root/bin/archerctl"),
		// 				Rpath: to.Ptr("no"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/BinaryHardening_ListByFirmware_MinimumSet_Gen.json
func ExampleBinaryHardeningClient_NewListByFirmwarePager_binaryHardeningListByFirmwareMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBinaryHardeningClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
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
		// page.BinaryHardeningListResult = armiotfirmwaredefense.BinaryHardeningListResult{
		// }
	}
}
