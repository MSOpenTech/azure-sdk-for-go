//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9c51b17f1c544eea0f6a67c01a6b763995521f52/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_List.json
func ExampleIscsiPathsClient_NewListByPrivateCloudPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiPathsClient().NewListByPrivateCloudPager("group1", "cloud1", nil)
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
		// page.IscsiPathListResult = armavs.IscsiPathListResult{
		// 	Value: []*armavs.IscsiPath{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/iscsiPaths"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/iscsiPaths/default"),
		// 			Properties: &armavs.IscsiPathProperties{
		// 				NetworkBlock: to.Ptr("192.168.0.0/24"),
		// 				ProvisioningState: to.Ptr(armavs.IscsiPathProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9c51b17f1c544eea0f6a67c01a6b763995521f52/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_Get.json
func ExampleIscsiPathsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIscsiPathsClient().Get(ctx, "group1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IscsiPath = armavs.IscsiPath{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/iscsiPaths"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/iscsiPaths/default"),
	// 	Properties: &armavs.IscsiPathProperties{
	// 		NetworkBlock: to.Ptr("192.168.0.0/24"),
	// 		ProvisioningState: to.Ptr(armavs.IscsiPathProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9c51b17f1c544eea0f6a67c01a6b763995521f52/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_CreateOrUpdate.json
func ExampleIscsiPathsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiPathsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", armavs.IscsiPath{
		Properties: &armavs.IscsiPathProperties{
			NetworkBlock: to.Ptr("192.168.0.0/24"),
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
	// res.IscsiPath = armavs.IscsiPath{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/iscsiPaths"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/iscsiPaths/default"),
	// 	Properties: &armavs.IscsiPathProperties{
	// 		NetworkBlock: to.Ptr("192.168.0.0/24"),
	// 		ProvisioningState: to.Ptr(armavs.IscsiPathProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9c51b17f1c544eea0f6a67c01a6b763995521f52/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_Delete.json
func ExampleIscsiPathsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiPathsClient().BeginDelete(ctx, "group1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
