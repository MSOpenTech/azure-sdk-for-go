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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Delete.json
func ExampleCloudServiceRoleInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServiceRoleInstancesClient().BeginDelete(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Get.json
func ExampleCloudServiceRoleInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudServiceRoleInstancesClient().Get(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", &armcompute.CloudServiceRoleInstancesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleInstance = armcompute.RoleInstance{
	// 	Name: to.Ptr("{roleInstance-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/{roleInstance-name}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armcompute.RoleInstanceProperties{
	// 		NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
	// 			NetworkInterfaces: []*armcompute.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/{roleInstance-name}/networkInterfaces/nic1"),
	// 			}},
	// 		},
	// 	},
	// 	SKU: &armcompute.InstanceSKU{
	// 		Name: to.Ptr("Standard_D1_v2"),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Get_InstanceView.json
func ExampleCloudServiceRoleInstancesClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudServiceRoleInstancesClient().GetInstanceView(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleInstanceView = armcompute.RoleInstanceView{
	// 	PlatformFaultDomain: to.Ptr[int32](0),
	// 	PlatformUpdateDomain: to.Ptr[int32](0),
	// 	PrivateID: to.Ptr("3491bc0c-1f6c-444f-b1d0-ec0751a74e3e"),
	// 	Statuses: []*armcompute.ResourceInstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("RoleState/RoleStateStarted"),
	// 			DisplayStatus: to.Ptr("RoleStateStarted"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Message: to.Ptr(""),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRolesInstance_List.json
func ExampleCloudServiceRoleInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudServiceRoleInstancesClient().NewListPager("ConstosoRG", "{cs-name}", &armcompute.CloudServiceRoleInstancesClientListOptions{Expand: nil})
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
		// page.RoleInstanceListResult = armcompute.RoleInstanceListResult{
		// 	Value: []*armcompute.RoleInstance{
		// 		{
		// 			Name: to.Ptr("ContosoFrontend_IN_0"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_0"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_0/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ContosoFrontend_IN_1"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_1"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoFrontend_IN_1/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ContosoBackend_IN_0"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_0"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_0/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ContosoBackend_IN_1"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices/roleInstances"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_1"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.RoleInstanceProperties{
		// 				NetworkProfile: &armcompute.RoleInstanceNetworkProfile{
		// 					NetworkInterfaces: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}/roleInstances/ContosoBackend_IN_1/networkInterfaces/nic1"),
		// 					}},
		// 				},
		// 			},
		// 			SKU: &armcompute.InstanceSKU{
		// 				Name: to.Ptr("Standard_D1_v2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Restart.json
func ExampleCloudServiceRoleInstancesClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServiceRoleInstancesClient().BeginRestart(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Reimage.json
func ExampleCloudServiceRoleInstancesClient_BeginReimage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServiceRoleInstancesClient().BeginReimage(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Rebuild.json
func ExampleCloudServiceRoleInstancesClient_BeginRebuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServiceRoleInstancesClient().BeginRebuild(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Get_RemoteDesktopFile.json
func ExampleCloudServiceRoleInstancesClient_GetRemoteDesktopFile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCloudServiceRoleInstancesClient().GetRemoteDesktopFile(ctx, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "rgcloudService", "aaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
