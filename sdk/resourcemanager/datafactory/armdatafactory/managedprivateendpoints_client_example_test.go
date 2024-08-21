//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_ListByFactory.json
func ExampleManagedPrivateEndpointsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedPrivateEndpointsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", nil)
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
		// page.ManagedPrivateEndpointListResponse = armdatafactory.ManagedPrivateEndpointListResponse{
		// 	Value: []*armdatafactory.ManagedPrivateEndpointResource{
		// 		{
		// 			Name: to.Ptr("exampleManagedPrivateEndpointName"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks/managedPrivateEndpoints"),
		// 			Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName/managedPrivateEndpoints/exampleManagedPrivateEndpointName"),
		// 			Properties: &armdatafactory.ManagedPrivateEndpoint{
		// 				ConnectionState: &armdatafactory.ConnectionStateProperties{
		// 					Description: to.Ptr(""),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr("Pending"),
		// 				},
		// 				Fqdns: []*string{
		// 				},
		// 				GroupID: to.Ptr("blob"),
		// 				PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Create.json
func ExampleManagedPrivateEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedPrivateEndpointsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", "exampleManagedPrivateEndpointName", armdatafactory.ManagedPrivateEndpointResource{
		Properties: &armdatafactory.ManagedPrivateEndpoint{
			Fqdns:                 []*string{},
			GroupID:               to.Ptr("blob"),
			PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
		},
	}, &armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedPrivateEndpointResource = armdatafactory.ManagedPrivateEndpointResource{
	// 	Name: to.Ptr("exampleManagedPrivateEndpointName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks/managedPrivateEndpoints"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName/managedPrivateEndpoints/exampleManagedPrivateEndpointName"),
	// 	Properties: &armdatafactory.ManagedPrivateEndpoint{
	// 		ConnectionState: &armdatafactory.ConnectionStateProperties{
	// 			Description: to.Ptr(""),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Pending"),
	// 		},
	// 		Fqdns: []*string{
	// 		},
	// 		GroupID: to.Ptr("blob"),
	// 		PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Get.json
func ExampleManagedPrivateEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedPrivateEndpointsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", "exampleManagedPrivateEndpointName", &armdatafactory.ManagedPrivateEndpointsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedPrivateEndpointResource = armdatafactory.ManagedPrivateEndpointResource{
	// 	Name: to.Ptr("exampleManagedPrivateEndpointName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks/managedPrivateEndpoints"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName/managedPrivateEndpoints/exampleManagedPrivateEndpointName"),
	// 	Properties: &armdatafactory.ManagedPrivateEndpoint{
	// 		ConnectionState: &armdatafactory.ConnectionStateProperties{
	// 			Description: to.Ptr(""),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Pending"),
	// 		},
	// 		Fqdns: []*string{
	// 		},
	// 		GroupID: to.Ptr("blob"),
	// 		PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Delete.json
func ExampleManagedPrivateEndpointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagedPrivateEndpointsClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", "exampleManagedPrivateEndpointName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
