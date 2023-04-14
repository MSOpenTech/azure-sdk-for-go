//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopologySubscription_example.json
func ExampleTopologyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopologyClient().NewListPager(nil)
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
		// page.TopologyList = armsecurity.TopologyList{
		// 	Value: []*armsecurity.TopologyResource{
		// 		{
		// 			Location: to.Ptr("westus"),
		// 			Name: to.Ptr("vnets"),
		// 			Type: to.Ptr("Microsoft.Security/locations/topologies"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/vnets"),
		// 			Properties: &armsecurity.TopologyResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.5755270Z"); return t}()),
		// 				TopologyResources: []*armsecurity.TopologySingleResource{
		// 					{
		// 						Children: []*armsecurity.TopologySingleResourceChild{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		// 						}},
		// 						Location: to.Ptr("westus"),
		// 						NetworkZones: to.Ptr("Internal"),
		// 						RecommendationsExist: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
		// 						Severity: to.Ptr("Healthy"),
		// 						TopologyScore: to.Ptr[int32](0),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("westus"),
		// 			Name: to.Ptr("subnets"),
		// 			Type: to.Ptr("Microsoft.Security/locations/topologies"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/subnets"),
		// 			Properties: &armsecurity.TopologyResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.5755270Z"); return t}()),
		// 				TopologyResources: []*armsecurity.TopologySingleResource{
		// 					{
		// 						Location: to.Ptr("westus"),
		// 						NetworkZones: to.Ptr("Internal"),
		// 						Parents: []*armsecurity.TopologySingleResourceParent{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
		// 						}},
		// 						RecommendationsExist: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		// 						Severity: to.Ptr("Healthy"),
		// 						TopologyScore: to.Ptr[int32](5),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopologySubscriptionLocation_example.json
func ExampleTopologyClient_NewListByHomeRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopologyClient().NewListByHomeRegionPager("centralus", nil)
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
		// page.TopologyList = armsecurity.TopologyList{
		// 	Value: []*armsecurity.TopologyResource{
		// 		{
		// 			Location: to.Ptr("westus"),
		// 			Name: to.Ptr("vnets"),
		// 			Type: to.Ptr("Microsoft.Security/locations/topologies"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/vnets"),
		// 			Properties: &armsecurity.TopologyResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.5755270Z"); return t}()),
		// 				TopologyResources: []*armsecurity.TopologySingleResource{
		// 					{
		// 						Children: []*armsecurity.TopologySingleResourceChild{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		// 						}},
		// 						Location: to.Ptr("westus"),
		// 						NetworkZones: to.Ptr("Internal"),
		// 						RecommendationsExist: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
		// 						Severity: to.Ptr("Healthy"),
		// 						TopologyScore: to.Ptr[int32](0),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("westus"),
		// 			Name: to.Ptr("subnets"),
		// 			Type: to.Ptr("Microsoft.Security/locations/topologies"),
		// 			ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/subnets"),
		// 			Properties: &armsecurity.TopologyResourceProperties{
		// 				CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.5755270Z"); return t}()),
		// 				TopologyResources: []*armsecurity.TopologySingleResource{
		// 					{
		// 						Location: to.Ptr("westus"),
		// 						NetworkZones: to.Ptr("Internal"),
		// 						Parents: []*armsecurity.TopologySingleResourceParent{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
		// 						}},
		// 						RecommendationsExist: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		// 						Severity: to.Ptr("Healthy"),
		// 						TopologyScore: to.Ptr[int32](5),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopology_example.json
func ExampleTopologyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopologyClient().Get(ctx, "myservers", "centralus", "vnets", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TopologyResource = armsecurity.TopologyResource{
	// 	Location: to.Ptr("westus"),
	// 	Name: to.Ptr("vnets"),
	// 	Type: to.Ptr("Microsoft.Security/locations/topologies"),
	// 	ID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Security/locations/centralus/topologies/vnets"),
	// 	Properties: &armsecurity.TopologyResourceProperties{
	// 		CalculatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-10T13:56:10.5755270Z"); return t}()),
	// 		TopologyResources: []*armsecurity.TopologySingleResource{
	// 			{
	// 				Children: []*armsecurity.TopologySingleResourceChild{
	// 					{
	// 						ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
	// 				}},
	// 				Location: to.Ptr("westus"),
	// 				NetworkZones: to.Ptr("InternetFacing"),
	// 				RecommendationsExist: to.Ptr(false),
	// 				ResourceID: to.Ptr("/subscriptions/3eeab341-f466-499c-a8be-85427e154bad/resourceGroups/myservers/providers/Microsoft.Network/virtualNetworks/myvnet"),
	// 				Severity: to.Ptr("Healthy"),
	// 				TopologyScore: to.Ptr[int32](0),
	// 		}},
	// 	},
	// }
}
