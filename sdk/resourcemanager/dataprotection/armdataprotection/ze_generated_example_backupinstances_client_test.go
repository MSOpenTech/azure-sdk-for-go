//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/ListBackupInstances.json
func ExampleBackupInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("<resource-group-name>",
		"<vault-name>",
		nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/GetBackupInstance.json
func ExampleBackupInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/PutBackupInstance.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.BackupInstanceResource{
			Properties: &armdataprotection.BackupInstance{
				DataSourceInfo: &armdataprotection.Datasource{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DataSourceSetInfo: &armdataprotection.DatasourceSet{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
					ObjectType: to.Ptr("<object-type>"),
					SecretStoreResource: &armdataprotection.SecretStoreResource{
						SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
						URI:             to.Ptr("<uri>"),
					},
				},
				FriendlyName: to.Ptr("<friendly-name>"),
				ObjectType:   to.Ptr("<object-type>"),
				PolicyInfo: &armdataprotection.PolicyInfo{
					PolicyID: to.Ptr("<policy-id>"),
					PolicyParameters: &armdataprotection.PolicyParameters{
						DataStoreParametersList: []armdataprotection.DataStoreParametersClassification{
							&armdataprotection.AzureOperationalStoreParameters{
								DataStoreType:   to.Ptr(armdataprotection.DataStoreTypesOperationalStore),
								ObjectType:      to.Ptr("<object-type>"),
								ResourceGroupID: to.Ptr("<resource-group-id>"),
							}},
					},
				},
				ValidationType: to.Ptr(armdataprotection.ValidationTypeShallowValidation),
			},
		},
		&armdataprotection.BackupInstancesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/DeleteBackupInstance.json
func ExampleBackupInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		&armdataprotection.BackupInstancesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/TriggerBackup.json
func ExampleBackupInstancesClient_BeginAdhocBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAdhocBackup(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.TriggerBackupRequest{
			BackupRuleOptions: &armdataprotection.AdHocBackupRuleOptions{
				RuleName: to.Ptr("<rule-name>"),
				TriggerOption: &armdataprotection.AdhocBackupTriggerOption{
					RetentionTagOverride: to.Ptr("<retention-tag-override>"),
				},
			},
		},
		&armdataprotection.BackupInstancesClientBeginAdhocBackupOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/ValidateForBackup.json
func ExampleBackupInstancesClient_BeginValidateForBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginValidateForBackup(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armdataprotection.ValidateForBackupRequest{
			BackupInstance: &armdataprotection.BackupInstance{
				DataSourceInfo: &armdataprotection.Datasource{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DataSourceSetInfo: &armdataprotection.DatasourceSet{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
					ObjectType: to.Ptr("<object-type>"),
					SecretStoreResource: &armdataprotection.SecretStoreResource{
						SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
						URI:             to.Ptr("<uri>"),
					},
				},
				FriendlyName: to.Ptr("<friendly-name>"),
				ObjectType:   to.Ptr("<object-type>"),
				PolicyInfo: &armdataprotection.PolicyInfo{
					PolicyID: to.Ptr("<policy-id>"),
				},
			},
		},
		&armdataprotection.BackupInstancesClientBeginValidateForBackupOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/GetBackupInstanceOperationResult.json
func ExampleBackupInstancesClient_GetBackupInstanceOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetBackupInstanceOperationResult(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/TriggerRehydrate.json
func ExampleBackupInstancesClient_BeginTriggerRehydrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRehydrate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.AzureBackupRehydrationRequest{
			RecoveryPointID:              to.Ptr("<recovery-point-id>"),
			RehydrationPriority:          to.Ptr(armdataprotection.RehydrationPriorityHigh),
			RehydrationRetentionDuration: to.Ptr("<rehydration-retention-duration>"),
		},
		&armdataprotection.BackupInstancesClientBeginTriggerRehydrateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/TriggerRestore.json
func ExampleBackupInstancesClient_BeginTriggerRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginTriggerRestore(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		&armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
			ObjectType: to.Ptr("<object-type>"),
			RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
				ObjectType:      to.Ptr("<object-type>"),
				RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
				RestoreLocation: to.Ptr("<restore-location>"),
				DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
					ObjectType: to.Ptr("<object-type>"),
					SecretStoreResource: &armdataprotection.SecretStoreResource{
						SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
						URI:             to.Ptr("<uri>"),
					},
				},
				DatasourceInfo: &armdataprotection.Datasource{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
				DatasourceSetInfo: &armdataprotection.DatasourceSet{
					DatasourceType:   to.Ptr("<datasource-type>"),
					ObjectType:       to.Ptr("<object-type>"),
					ResourceID:       to.Ptr("<resource-id>"),
					ResourceLocation: to.Ptr("<resource-location>"),
					ResourceName:     to.Ptr("<resource-name>"),
					ResourceType:     to.Ptr("<resource-type>"),
					ResourceURI:      to.Ptr("<resource-uri>"),
				},
			},
			SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
			SourceResourceID:    to.Ptr("<source-resource-id>"),
			RecoveryPointID:     to.Ptr("<recovery-point-id>"),
		},
		&armdataprotection.BackupInstancesClientBeginTriggerRestoreOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/ResumeBackups.json
func ExampleBackupInstancesClient_BeginResumeBackups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResumeBackups(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		&armdataprotection.BackupInstancesClientBeginResumeBackupsOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/ResumeProtection.json
func ExampleBackupInstancesClient_BeginResumeProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResumeProtection(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		&armdataprotection.BackupInstancesClientBeginResumeProtectionOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/StopProtection.json
func ExampleBackupInstancesClient_BeginStopProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStopProtection(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		&armdataprotection.BackupInstancesClientBeginStopProtectionOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/SuspendBackups.json
func ExampleBackupInstancesClient_BeginSuspendBackups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSuspendBackups(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		&armdataprotection.BackupInstancesClientBeginSuspendBackupsOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/SyncBackupInstance.json
func ExampleBackupInstancesClient_BeginSyncBackupInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSyncBackupInstance(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.SyncBackupInstanceRequest{
			SyncType: to.Ptr(armdataprotection.SyncTypeDefault),
		},
		&armdataprotection.BackupInstancesClientBeginSyncBackupInstanceOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-02-01-preview/examples/BackupInstanceOperations/ValidateRestore.json
func ExampleBackupInstancesClient_BeginValidateForRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewBackupInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginValidateForRestore(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<backup-instance-name>",
		armdataprotection.ValidateRestoreRequestObject{
			RestoreRequestObject: &armdataprotection.AzureBackupRecoveryPointBasedRestoreRequest{
				ObjectType: to.Ptr("<object-type>"),
				RestoreTargetInfo: &armdataprotection.RestoreTargetInfo{
					ObjectType:      to.Ptr("<object-type>"),
					RecoveryOption:  to.Ptr(armdataprotection.RecoveryOptionFailIfExists),
					RestoreLocation: to.Ptr("<restore-location>"),
					DatasourceAuthCredentials: &armdataprotection.SecretStoreBasedAuthCredentials{
						ObjectType: to.Ptr("<object-type>"),
						SecretStoreResource: &armdataprotection.SecretStoreResource{
							SecretStoreType: to.Ptr(armdataprotection.SecretStoreTypeAzureKeyVault),
							URI:             to.Ptr("<uri>"),
						},
					},
					DatasourceInfo: &armdataprotection.Datasource{
						DatasourceType:   to.Ptr("<datasource-type>"),
						ObjectType:       to.Ptr("<object-type>"),
						ResourceID:       to.Ptr("<resource-id>"),
						ResourceLocation: to.Ptr("<resource-location>"),
						ResourceName:     to.Ptr("<resource-name>"),
						ResourceType:     to.Ptr("<resource-type>"),
						ResourceURI:      to.Ptr("<resource-uri>"),
					},
					DatasourceSetInfo: &armdataprotection.DatasourceSet{
						DatasourceType:   to.Ptr("<datasource-type>"),
						ObjectType:       to.Ptr("<object-type>"),
						ResourceID:       to.Ptr("<resource-id>"),
						ResourceLocation: to.Ptr("<resource-location>"),
						ResourceName:     to.Ptr("<resource-name>"),
						ResourceType:     to.Ptr("<resource-type>"),
						ResourceURI:      to.Ptr("<resource-uri>"),
					},
				},
				SourceDataStoreType: to.Ptr(armdataprotection.SourceDataStoreTypeVaultStore),
				SourceResourceID:    to.Ptr("<source-resource-id>"),
				RecoveryPointID:     to.Ptr("<recovery-point-id>"),
			},
		},
		&armdataprotection.BackupInstancesClientBeginValidateForRestoreOptions{ResumeToken: ""})
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
