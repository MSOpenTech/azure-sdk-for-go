//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetWorkloadGroup.json
func ExampleWorkloadGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadGroupsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "smallrc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkloadGroup = armsql.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/smallrc"),
	// 	Properties: &armsql.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadGroupMax.json
func ExampleWorkloadGroupsClient_BeginCreateOrUpdate_createAWorkloadGroupWithAllPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadGroupsClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "smallrc", armsql.WorkloadGroup{
		Properties: &armsql.WorkloadGroupProperties{
			Importance:                   to.Ptr("normal"),
			MaxResourcePercent:           to.Ptr[int32](100),
			MaxResourcePercentPerRequest: to.Ptr[float64](3),
			MinResourcePercent:           to.Ptr[int32](0),
			MinResourcePercentPerRequest: to.Ptr[float64](3),
			QueryExecutionTimeout:        to.Ptr[int32](0),
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
	// res.WorkloadGroup = armsql.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/smallrc"),
	// 	Properties: &armsql.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadGroupMin.json
func ExampleWorkloadGroupsClient_BeginCreateOrUpdate_createAWorkloadGroupWithTheRequiredPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadGroupsClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "smallrc", armsql.WorkloadGroup{
		Properties: &armsql.WorkloadGroupProperties{
			MaxResourcePercent:           to.Ptr[int32](100),
			MinResourcePercent:           to.Ptr[int32](0),
			MinResourcePercentPerRequest: to.Ptr[float64](3),
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
	// res.WorkloadGroup = armsql.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/smallrc"),
	// 	Properties: &armsql.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeleteWorkloadGroup.json
func ExampleWorkloadGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkloadGroupsClient().BeginDelete(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetWorkloadGroupList.json
func ExampleWorkloadGroupsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkloadGroupsClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "testdb", nil)
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
		// page.WorkloadGroupListResult = armsql.WorkloadGroupListResult{
		// 	Value: []*armsql.WorkloadGroup{
		// 		{
		// 			Name: to.Ptr("smallrc"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/smallrc"),
		// 			Properties: &armsql.WorkloadGroupProperties{
		// 				Importance: to.Ptr("normal"),
		// 				MaxResourcePercent: to.Ptr[int32](100),
		// 				MaxResourcePercentPerRequest: to.Ptr[float64](5),
		// 				MinResourcePercent: to.Ptr[int32](0),
		// 				MinResourcePercentPerRequest: to.Ptr[float64](5),
		// 				QueryExecutionTimeout: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mediumrc"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/mediumrc"),
		// 			Properties: &armsql.WorkloadGroupProperties{
		// 				Importance: to.Ptr("normal"),
		// 				MaxResourcePercent: to.Ptr[int32](100),
		// 				MaxResourcePercentPerRequest: to.Ptr[float64](10),
		// 				MinResourcePercent: to.Ptr[int32](0),
		// 				MinResourcePercentPerRequest: to.Ptr[float64](10),
		// 				QueryExecutionTimeout: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("largerc"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/largerc"),
		// 			Properties: &armsql.WorkloadGroupProperties{
		// 				Importance: to.Ptr("high"),
		// 				MaxResourcePercent: to.Ptr[int32](100),
		// 				MaxResourcePercentPerRequest: to.Ptr[float64](20),
		// 				MinResourcePercent: to.Ptr[int32](0),
		// 				MinResourcePercentPerRequest: to.Ptr[float64](20),
		// 				QueryExecutionTimeout: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
