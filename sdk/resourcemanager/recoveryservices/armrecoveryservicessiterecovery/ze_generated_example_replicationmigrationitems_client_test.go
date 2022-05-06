//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_ListByReplicationProtectionContainers.json
func ExampleReplicationMigrationItemsClient_NewListByReplicationProtectionContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByReplicationProtectionContainersPager("<fabric-name>",
		"<protection-container-name>",
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientListByReplicationProtectionContainersOptions{SkipToken: nil,
			TakeToken: nil,
			Filter:    nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Get.json
func ExampleReplicationMigrationItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Create.json
func ExampleReplicationMigrationItemsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.EnableMigrationInput{
			Properties: &armrecoveryservicessiterecovery.EnableMigrationInputProperties{
				PolicyID: to.Ptr("<policy-id>"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtEnableMigrationInput{
					InstanceType:            to.Ptr("<instance-type>"),
					DataMoverRunAsAccountID: to.Ptr("<data-mover-run-as-account-id>"),
					DisksToInclude: []*armrecoveryservicessiterecovery.VMwareCbtDiskInput{
						{
							DiskID:                         to.Ptr("<disk-id>"),
							IsOSDisk:                       to.Ptr("<is-osdisk>"),
							LogStorageAccountID:            to.Ptr("<log-storage-account-id>"),
							LogStorageAccountSasSecretName: to.Ptr("<log-storage-account-sas-secret-name>"),
						}},
					SnapshotRunAsAccountID: to.Ptr("<snapshot-run-as-account-id>"),
					TargetNetworkID:        to.Ptr("<target-network-id>"),
					TargetResourceGroupID:  to.Ptr("<target-resource-group-id>"),
					VmwareMachineID:        to.Ptr("<vmware-machine-id>"),
				},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Delete.json
func ExampleReplicationMigrationItemsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginDeleteOptions{DeleteOption: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Update.json
func ExampleReplicationMigrationItemsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.UpdateMigrationItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateMigrationItemInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtUpdateMigrationItemInput{
					InstanceType: to.Ptr("<instance-type>"),
				},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Migrate.json
func ExampleReplicationMigrationItemsClient_BeginMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginMigrate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.MigrateInput{
			Properties: &armrecoveryservicessiterecovery.MigrateInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtMigrateInput{
					InstanceType:    to.Ptr("<instance-type>"),
					PerformShutdown: to.Ptr("<perform-shutdown>"),
				},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginMigrateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Resync.json
func ExampleReplicationMigrationItemsClient_BeginResync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResync(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.ResyncInput{
			Properties: &armrecoveryservicessiterecovery.ResyncInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtResyncInput{
					InstanceType: to.Ptr("<instance-type>"),
					SkipCbtReset: to.Ptr("<skip-cbt-reset>"),
				},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginResyncOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_TestMigrate.json
func ExampleReplicationMigrationItemsClient_BeginTestMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTestMigrate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.TestMigrateInput{
			Properties: &armrecoveryservicessiterecovery.TestMigrateInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtTestMigrateInput{
					InstanceType:    to.Ptr("<instance-type>"),
					NetworkID:       to.Ptr("<network-id>"),
					RecoveryPointID: to.Ptr("<recovery-point-id>"),
				},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginTestMigrateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_TestMigrateCleanup.json
func ExampleReplicationMigrationItemsClient_BeginTestMigrateCleanup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTestMigrateCleanup(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.TestMigrateCleanupInput{
			Properties: &armrecoveryservicessiterecovery.TestMigrateCleanupInputProperties{
				Comments: to.Ptr("<comments>"),
			},
		},
		&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientBeginTestMigrateCleanupOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_List.json
func ExampleReplicationMigrationItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armrecoveryservicessiterecovery.ReplicationMigrationItemsClientListOptions{SkipToken: nil,
		TakeToken: nil,
		Filter:    nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
