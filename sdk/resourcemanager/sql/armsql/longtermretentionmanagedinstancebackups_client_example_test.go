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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceLongTermRetentionBackupGet.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().Get(ctx, "japaneast", "testInstance", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceLongTermRetentionBackup = armsql.ManagedInstanceLongTermRetentionBackup{
	// 	Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
	// 	Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
	// 		BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-23T08:00:00.000Z"); return t}()),
	// 		DatabaseName: to.Ptr("testDatabase"),
	// 		ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-01T08:00:00.000Z"); return t}()),
	// 		ManagedInstanceName: to.Ptr("testInstance"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceLongTermRetentionBackupDelete.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().BeginDelete(ctx, "japaneast", "testInstance", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceLongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByDatabasePager("japaneast", "testInstance", "testDatabase", &armsql.LongTermRetentionManagedInstanceBackupsClientListByDatabaseOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-23T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-23T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-30T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-30T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-09-06T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-09-06T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-06T08:00:00.000Z"); return t}()),
		// 				DatabaseDeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-07T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceLongTermRetentionBackupListByInstance.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByInstancePager("japaneast", "testInstance", &armsql.LongTermRetentionManagedInstanceBackupsClientListByInstanceOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase1/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase1"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("12341234-1234-1234-1234-123123123123;131657960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase2/longTermRetentionManagedInstanceBackups/12341234-1234-1234-1234-123123123123;131657960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase2"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("43214321-4321-4321-4321-321321321321;131667960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase3/longTermRetentionManagedInstanceBackups/43214321-4321-4321-4321-321321321321;131677960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase3"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceLongTermRetentionBackupListByLocation.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByLocationPager("japaneast", &armsql.LongTermRetentionManagedInstanceBackupsClientListByLocationOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;2017-08-23T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testserver1/longTermRetentionDatabases/testDatabase1/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase1"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("12341234-1234-1234-1234-123123123123;2017-08-30T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testserver2/longTermRetentionDatabases/testDatabase2/longTermRetentionManagedInstanceBackups/12341234-1234-1234-1234-123123123123;131657960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase2"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("43214321-4321-4321-4321-321321321321;2017-09-06T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testserver3/longTermRetentionDatabases/testDatabase3/longTermRetentionManagedInstanceBackups/43214321-4321-4321-4321-321321321321;131677960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase3"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance3"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupGet.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_GetByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().GetByResourceGroup(ctx, "testResourceGroup", "japaneast", "testInstance", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceLongTermRetentionBackup = armsql.ManagedInstanceLongTermRetentionBackup{
	// 	Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
	// 	Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
	// 		BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
	// 		DatabaseName: to.Ptr("testDatabase"),
	// 		ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
	// 		ManagedInstanceName: to.Ptr("testInstance"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupDelete.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_BeginDeleteByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().BeginDeleteByResourceGroup(ctx, "testResourceGroup", "japaneast", "testInstance", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByResourceGroupDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByResourceGroupDatabasePager("testResourceGroup", "japaneast", "testInstance", "testDatabase", &armsql.LongTermRetentionManagedInstanceBackupsClientListByResourceGroupDatabaseOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-23T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-23T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-30T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-30T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-09-06T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-09-06T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-06T08:00:00.000Z"); return t}()),
		// 				DatabaseDeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-07T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByInstance.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByResourceGroupInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByResourceGroupInstancePager("testResourceGroup", "japaneast", "testInstance", &armsql.LongTermRetentionManagedInstanceBackupsClientListByResourceGroupInstanceOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase1/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase1"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("12341234-1234-1234-1234-123123123123;131657960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase2/longTermRetentionManagedInstanceBackups/12341234-1234-1234-1234-123123123123;131657960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase2"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("43214321-4321-4321-4321-321321321321;131667960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase3/longTermRetentionManagedInstanceBackups/43214321-4321-4321-4321-321321321321;131677960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase3"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByLocation.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByResourceGroupLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByResourceGroupLocationPager("testResourceGroup", "japaneast", &armsql.LongTermRetentionManagedInstanceBackupsClientListByResourceGroupLocationOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;2017-08-23T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance1/longTermRetentionDatabases/testDatabase1/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase1"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("12341234-1234-1234-1234-123123123123;2017-08-30T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance2/longTermRetentionDatabases/testDatabase2/longTermRetentionManagedInstanceBackups/12341234-1234-1234-1234-123123123123;131657960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase2"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("43214321-4321-4321-4321-321321321321;2017-09-06T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance3/longTermRetentionDatabases/testDatabase3/longTermRetentionManagedInstanceBackups/43214321-4321-4321-4321-321321321321;131677960820000000"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase3"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-10T08:00:00.000Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance3"),
		// 			},
		// 	}},
		// }
	}
}
