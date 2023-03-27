//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ef6f2f06858cdbec7684968e1a54c7610da97dbb/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/LocationGetQuotas.json
func ExampleLocationClient_GetQuotas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationClient().GetQuotas(ctx, "japaneast", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LocationQuota = armbatch.LocationQuota{
	// 	AccountQuota: to.Ptr[int32](1),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ef6f2f06858cdbec7684968e1a54c7610da97dbb/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/LocationListVirtualMachineSkus.json
func ExampleLocationClient_NewListSupportedVirtualMachineSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListSupportedVirtualMachineSKUsPager("japaneast", &armbatch.LocationClientListSupportedVirtualMachineSKUsOptions{Maxresults: nil,
		Filter: nil,
	})
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
		// page.SupportedSKUsResult = armbatch.SupportedSKUsResult{
		// 	Value: []*armbatch.SupportedSKU{
		// 		{
		// 			Name: to.Ptr("Standard_D1_v2"),
		// 			Capabilities: []*armbatch.SKUCapability{
		// 				{
		// 					Name: to.Ptr("MaxResourceVolumeMB"),
		// 					Value: to.Ptr("20480"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUs"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("HyperVGenerations"),
		// 					Value: to.Ptr("V1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryGB"),
		// 					Value: to.Ptr("0.75"),
		// 				},
		// 				{
		// 					Name: to.Ptr("LowPriorityCapable"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUsAvailable"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("EphemeralOSDiskSupported"),
		// 					Value: to.Ptr("False"),
		// 			}},
		// 			FamilyName: to.Ptr("standardDFamily"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ef6f2f06858cdbec7684968e1a54c7610da97dbb/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/LocationListCloudServiceSkus.json
func ExampleLocationClient_NewListSupportedCloudServiceSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListSupportedCloudServiceSKUsPager("japaneast", &armbatch.LocationClientListSupportedCloudServiceSKUsOptions{Maxresults: nil,
		Filter: nil,
	})
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
		// page.SupportedSKUsResult = armbatch.SupportedSKUsResult{
		// 	Value: []*armbatch.SupportedSKU{
		// 		{
		// 			Name: to.Ptr("Small"),
		// 			Capabilities: []*armbatch.SKUCapability{
		// 				{
		// 					Name: to.Ptr("MaxResourceVolumeMB"),
		// 					Value: to.Ptr("20480"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUs"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("HyperVGenerations"),
		// 					Value: to.Ptr("V1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryGB"),
		// 					Value: to.Ptr("0.75"),
		// 				},
		// 				{
		// 					Name: to.Ptr("LowPriorityCapable"),
		// 					Value: to.Ptr("False"),
		// 				},
		// 				{
		// 					Name: to.Ptr("vCPUsAvailable"),
		// 					Value: to.Ptr("1"),
		// 				},
		// 				{
		// 					Name: to.Ptr("EphemeralOSDiskSupported"),
		// 					Value: to.Ptr("False"),
		// 			}},
		// 			FamilyName: to.Ptr("standardA0_A7Family"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ef6f2f06858cdbec7684968e1a54c7610da97dbb/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/LocationCheckNameAvailability_AlreadyExists.json
func ExampleLocationClient_CheckNameAvailability_locationCheckNameAvailabilityAlreadyExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationClient().CheckNameAvailability(ctx, "japaneast", armbatch.CheckNameAvailabilityParameters{
		Name: to.Ptr("existingaccountname"),
		Type: to.Ptr("Microsoft.Batch/batchAccounts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armbatch.CheckNameAvailabilityResult{
	// 	Message: to.Ptr("An account named 'existingaccountname' is already in use."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armbatch.NameAvailabilityReasonAlreadyExists),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ef6f2f06858cdbec7684968e1a54c7610da97dbb/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/LocationCheckNameAvailability_Available.json
func ExampleLocationClient_CheckNameAvailability_locationCheckNameAvailabilityAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationClient().CheckNameAvailability(ctx, "japaneast", armbatch.CheckNameAvailabilityParameters{
		Name: to.Ptr("newaccountname"),
		Type: to.Ptr("Microsoft.Batch/batchAccounts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armbatch.CheckNameAvailabilityResult{
	// 	NameAvailable: to.Ptr(true),
	// }
}
