// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongocluster_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
	"log"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster/TempTypeSpecFiles/DocumentDB.MongoCluster.Management/examples/2024-06-01-preview/MongoClusters_PrivateLinkResourceList.json
func ExamplePrivateLinksClient_NewListByMongoClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinksClient().NewListByMongoClusterPager("TestGroup", "myMongoCluster", nil)
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
		// page = armmongocluster.PrivateLinksClientListByMongoClusterResponse{
		// 	PrivateLinkResourceListResult: armmongocluster.PrivateLinkResourceListResult{
		// 		Value: []*armmongocluster.PrivateLinkResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster/privateLinkResources/MongoCluster"),
		// 				Name: to.Ptr("MongoCluster"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/mongoClusters/privateLinkResources"),
		// 				Properties: &armmongocluster.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("MongoCluster"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("c.p7ex3v2euquypn"),
		// 					},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.mongocluster.cosmos.azure.com"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
