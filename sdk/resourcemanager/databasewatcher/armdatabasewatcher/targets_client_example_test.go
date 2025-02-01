//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/databasewatcher/resource-manager/Microsoft.DatabaseWatcher/stable/2025-01-02/examples/Targets_ListByWatcher_MaximumSet_Gen.json
func ExampleTargetsClient_NewListByWatcherPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTargetsClient().NewListByWatcherPager("apiTest-ddat4p", "databasemo3ej9ih", nil)
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
		// page.TargetListResult = armdatabasewatcher.TargetListResult{
		// 	Value: []*armdatabasewatcher.Target{
		// 		{
		// 			Name: to.Ptr("monitoringzkndgm"),
		// 			Type: to.Ptr("microsoft.databasewatcher/watchers/targets"),
		// 			ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-px9ma7/providers/Microsoft.DatabaseWatcher/watchers/databasemo3d9sgt/targets/monitoringzkndgm"),
		// 			SystemData: &armdatabasewatcher.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
		// 				CreatedBy: to.Ptr("enbpvlpqbwd"),
		// 				CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("mxp"),
		// 				LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdatabasewatcher.SQLDbSingleDatabaseTargetProperties{
		// 				ConnectionServerName: to.Ptr("sqlServero1ihe2"),
		// 				ProvisioningState: to.Ptr(armdatabasewatcher.ResourceProvisioningStateSucceeded),
		// 				TargetAuthenticationType: to.Ptr(armdatabasewatcher.TargetAuthenticationTypeAAD),
		// 				TargetType: to.Ptr("SqlDb"),
		// 				SQLDbResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/databasewatcher/resource-manager/Microsoft.DatabaseWatcher/stable/2025-01-02/examples/Targets_Get_MaximumSet_Gen.json
func ExampleTargetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetsClient().Get(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Target = armdatabasewatcher.Target{
	// 	Name: to.Ptr("monitoringzkndgm"),
	// 	Type: to.Ptr("microsoft.databasewatcher/watchers/targets"),
	// 	ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-px9ma7/providers/Microsoft.DatabaseWatcher/watchers/databasemo3d9sgt/targets/monitoringzkndgm"),
	// 	SystemData: &armdatabasewatcher.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 		CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("mxp"),
	// 		LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdatabasewatcher.SQLDbSingleDatabaseTargetProperties{
	// 		ConnectionServerName: to.Ptr("sqlServero1ihe2"),
	// 		ProvisioningState: to.Ptr(armdatabasewatcher.ResourceProvisioningStateSucceeded),
	// 		TargetAuthenticationType: to.Ptr(armdatabasewatcher.TargetAuthenticationTypeAAD),
	// 		TargetType: to.Ptr("SqlDb"),
	// 		SQLDbResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/databasewatcher/resource-manager/Microsoft.DatabaseWatcher/stable/2025-01-02/examples/Targets_CreateOrUpdate_MaximumSet_Gen.json
func ExampleTargetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetsClient().CreateOrUpdate(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", armdatabasewatcher.Target{
		Properties: &armdatabasewatcher.SQLDbSingleDatabaseTargetProperties{
			ConnectionServerName:     to.Ptr("sqlServero1ihe2"),
			TargetAuthenticationType: to.Ptr(armdatabasewatcher.TargetAuthenticationTypeAAD),
			TargetType:               to.Ptr("SqlDb"),
			SQLDbResourceID:          to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Target = armdatabasewatcher.Target{
	// 	Name: to.Ptr("monitoringzkndgm"),
	// 	Type: to.Ptr("microsoft.databasewatcher/watchers/targets"),
	// 	ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-px9ma7/providers/Microsoft.DatabaseWatcher/watchers/databasemo3d9sgt/targets/monitoringzkndgm"),
	// 	SystemData: &armdatabasewatcher.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 		CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("mxp"),
	// 		LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdatabasewatcher.SQLDbSingleDatabaseTargetProperties{
	// 		ConnectionServerName: to.Ptr("sqlServero1ihe2"),
	// 		ProvisioningState: to.Ptr(armdatabasewatcher.ResourceProvisioningStateSucceeded),
	// 		TargetAuthenticationType: to.Ptr(armdatabasewatcher.TargetAuthenticationTypeAAD),
	// 		TargetType: to.Ptr("SqlDb"),
	// 		SQLDbResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/databasewatcher/resource-manager/Microsoft.DatabaseWatcher/stable/2025-01-02/examples/Targets_Delete_MaximumSet_Gen.json
func ExampleTargetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTargetsClient().Delete(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
