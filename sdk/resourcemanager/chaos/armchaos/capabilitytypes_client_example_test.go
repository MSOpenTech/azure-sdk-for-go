//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/ListCapabilityTypes.json
func ExampleCapabilityTypesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapabilityTypesClient().NewListPager("westus2", "Microsoft-VirtualMachine", &armchaos.CapabilityTypesClientListOptions{ContinuationToken: nil})
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
		// page.CapabilityTypeListResult = armchaos.CapabilityTypeListResult{
		// 	Value: []*armchaos.CapabilityType{
		// 		{
		// 			Name: to.Ptr("Shutdown-1.0"),
		// 			Type: to.Ptr("Microsoft.Chaos/locations/targetTypes/capabilityTypes"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/providers/Microsoft.Chaos/locations/westus2/targetTypes/Microsoft-VirtualMachine/capabilityTypes/Shutdown-1.0"),
		// 			Properties: &armchaos.CapabilityTypeProperties{
		// 				Description: to.Ptr("Shutdown an Azure Virtual Machine for a defined period of time."),
		// 				DisplayName: to.Ptr("Shutdown VM"),
		// 				Kind: to.Ptr("fault"),
		// 				ParametersSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0.json"),
		// 				Publisher: to.Ptr("Microsoft"),
		// 				RuntimeProperties: &armchaos.CapabilityTypePropertiesRuntimeProperties{
		// 					Kind: to.Ptr("continuous"),
		// 				},
		// 				TargetType: to.Ptr("VirtualMachine"),
		// 				Urn: to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/GetCapabilityType.json
func ExampleCapabilityTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapabilityTypesClient().Get(ctx, "westus2", "Microsoft-VirtualMachine", "Shutdown-1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapabilityType = armchaos.CapabilityType{
	// 	Name: to.Ptr("Shutdown-1.0"),
	// 	Type: to.Ptr("Microsoft.Chaos/locations/targetTypes/capabilityTypes"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/Microsoft.Chaos/locations/westus2/targetTypes/Microsoft-VirtualMachine/capabilityTypes/Shutdown-1.0"),
	// 	Properties: &armchaos.CapabilityTypeProperties{
	// 		Description: to.Ptr("Shutdown an Azure Virtual Machine for a defined period of time."),
	// 		DisplayName: to.Ptr("Shutdown VM"),
	// 		Kind: to.Ptr("fault"),
	// 		ParametersSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0.json"),
	// 		Publisher: to.Ptr("Microsoft"),
	// 		RuntimeProperties: &armchaos.CapabilityTypePropertiesRuntimeProperties{
	// 			Kind: to.Ptr("continuous"),
	// 		},
	// 		TargetType: to.Ptr("VirtualMachine"),
	// 		Urn: to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
	// 	},
	// }
}
