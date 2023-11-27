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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/GetShortTermRetentionPolicy.json
func ExampleBackupShortTermRetentionPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupShortTermRetentionPoliciesClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", armsql.ShortTermRetentionPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupShortTermRetentionPolicy = armsql.BackupShortTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/backupShortTermRetentionPolicies/default"),
	// 	Properties: &armsql.BackupShortTermRetentionPolicyProperties{
	// 		DiffBackupIntervalInHours: to.Ptr(armsql.DiffBackupIntervalInHours(24)),
	// 		RetentionDays: to.Ptr[int32](7),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/UpdateShortTermRetentionPolicy.json
func ExampleBackupShortTermRetentionPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupShortTermRetentionPoliciesClient().BeginCreateOrUpdate(ctx, "resourceGroup", "testsvr", "testdb", armsql.ShortTermRetentionPolicyNameDefault, armsql.BackupShortTermRetentionPolicy{
		Properties: &armsql.BackupShortTermRetentionPolicyProperties{
			DiffBackupIntervalInHours: to.Ptr(armsql.DiffBackupIntervalInHours(24)),
			RetentionDays:             to.Ptr[int32](7),
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
	// res.BackupShortTermRetentionPolicy = armsql.BackupShortTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testsvr/databases/testdb/backupShortTermRetentionPolicies/default"),
	// 	Properties: &armsql.BackupShortTermRetentionPolicyProperties{
	// 		DiffBackupIntervalInHours: to.Ptr(armsql.DiffBackupIntervalInHours(24)),
	// 		RetentionDays: to.Ptr[int32](7),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListShortTermRetentionPoliciesByDatabase.json
func ExampleBackupShortTermRetentionPoliciesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupShortTermRetentionPoliciesClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "testdb", nil)
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
		// page.BackupShortTermRetentionPolicyListResult = armsql.BackupShortTermRetentionPolicyListResult{
		// 	Value: []*armsql.BackupShortTermRetentionPolicy{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/backupShortTermRetentionPolicies/default"),
		// 			Properties: &armsql.BackupShortTermRetentionPolicyProperties{
		// 				DiffBackupIntervalInHours: to.Ptr(armsql.DiffBackupIntervalInHours(24)),
		// 				RetentionDays: to.Ptr[int32](7),
		// 			},
		// 	}},
		// }
	}
}
