//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewBMSPrepareDataMoveOperationResultClient creates a new instance of BMSPrepareDataMoveOperationResultClient.
func (c *ClientFactory) NewBMSPrepareDataMoveOperationResultClient() *BMSPrepareDataMoveOperationResultClient {
	return &BMSPrepareDataMoveOperationResultClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupEnginesClient creates a new instance of BackupEnginesClient.
func (c *ClientFactory) NewBackupEnginesClient() *BackupEnginesClient {
	return &BackupEnginesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupJobsClient creates a new instance of BackupJobsClient.
func (c *ClientFactory) NewBackupJobsClient() *BackupJobsClient {
	return &BackupJobsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupOperationResultsClient creates a new instance of BackupOperationResultsClient.
func (c *ClientFactory) NewBackupOperationResultsClient() *BackupOperationResultsClient {
	return &BackupOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupOperationStatusesClient creates a new instance of BackupOperationStatusesClient.
func (c *ClientFactory) NewBackupOperationStatusesClient() *BackupOperationStatusesClient {
	return &BackupOperationStatusesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient.
func (c *ClientFactory) NewBackupPoliciesClient() *BackupPoliciesClient {
	return &BackupPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupProtectableItemsClient creates a new instance of BackupProtectableItemsClient.
func (c *ClientFactory) NewBackupProtectableItemsClient() *BackupProtectableItemsClient {
	return &BackupProtectableItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupProtectedItemsClient creates a new instance of BackupProtectedItemsClient.
func (c *ClientFactory) NewBackupProtectedItemsClient() *BackupProtectedItemsClient {
	return &BackupProtectedItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupProtectionContainersClient creates a new instance of BackupProtectionContainersClient.
func (c *ClientFactory) NewBackupProtectionContainersClient() *BackupProtectionContainersClient {
	return &BackupProtectionContainersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupProtectionIntentClient creates a new instance of BackupProtectionIntentClient.
func (c *ClientFactory) NewBackupProtectionIntentClient() *BackupProtectionIntentClient {
	return &BackupProtectionIntentClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupResourceEncryptionConfigsClient creates a new instance of BackupResourceEncryptionConfigsClient.
func (c *ClientFactory) NewBackupResourceEncryptionConfigsClient() *BackupResourceEncryptionConfigsClient {
	return &BackupResourceEncryptionConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupResourceStorageConfigsNonCRRClient creates a new instance of BackupResourceStorageConfigsNonCRRClient.
func (c *ClientFactory) NewBackupResourceStorageConfigsNonCRRClient() *BackupResourceStorageConfigsNonCRRClient {
	return &BackupResourceStorageConfigsNonCRRClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupResourceVaultConfigsClient creates a new instance of BackupResourceVaultConfigsClient.
func (c *ClientFactory) NewBackupResourceVaultConfigsClient() *BackupResourceVaultConfigsClient {
	return &BackupResourceVaultConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupStatusClient creates a new instance of BackupStatusClient.
func (c *ClientFactory) NewBackupStatusClient() *BackupStatusClient {
	return &BackupStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupUsageSummariesClient creates a new instance of BackupUsageSummariesClient.
func (c *ClientFactory) NewBackupUsageSummariesClient() *BackupUsageSummariesClient {
	return &BackupUsageSummariesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupWorkloadItemsClient creates a new instance of BackupWorkloadItemsClient.
func (c *ClientFactory) NewBackupWorkloadItemsClient() *BackupWorkloadItemsClient {
	return &BackupWorkloadItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupsClient creates a new instance of BackupsClient.
func (c *ClientFactory) NewBackupsClient() *BackupsClient {
	return &BackupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeletedProtectionContainersClient creates a new instance of DeletedProtectionContainersClient.
func (c *ClientFactory) NewDeletedProtectionContainersClient() *DeletedProtectionContainersClient {
	return &DeletedProtectionContainersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExportJobsOperationResultsClient creates a new instance of ExportJobsOperationResultsClient.
func (c *ClientFactory) NewExportJobsOperationResultsClient() *ExportJobsOperationResultsClient {
	return &ExportJobsOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFeatureSupportClient creates a new instance of FeatureSupportClient.
func (c *ClientFactory) NewFeatureSupportClient() *FeatureSupportClient {
	return &FeatureSupportClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFetchTieringCostClient creates a new instance of FetchTieringCostClient.
func (c *ClientFactory) NewFetchTieringCostClient() *FetchTieringCostClient {
	return &FetchTieringCostClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGetTieringCostOperationResultClient creates a new instance of GetTieringCostOperationResultClient.
func (c *ClientFactory) NewGetTieringCostOperationResultClient() *GetTieringCostOperationResultClient {
	return &GetTieringCostOperationResultClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewItemLevelRecoveryConnectionsClient creates a new instance of ItemLevelRecoveryConnectionsClient.
func (c *ClientFactory) NewItemLevelRecoveryConnectionsClient() *ItemLevelRecoveryConnectionsClient {
	return &ItemLevelRecoveryConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobCancellationsClient creates a new instance of JobCancellationsClient.
func (c *ClientFactory) NewJobCancellationsClient() *JobCancellationsClient {
	return &JobCancellationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobDetailsClient creates a new instance of JobDetailsClient.
func (c *ClientFactory) NewJobDetailsClient() *JobDetailsClient {
	return &JobDetailsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobOperationResultsClient creates a new instance of JobOperationResultsClient.
func (c *ClientFactory) NewJobOperationResultsClient() *JobOperationResultsClient {
	return &JobOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobsClient creates a new instance of JobsClient.
func (c *ClientFactory) NewJobsClient() *JobsClient {
	return &JobsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationClient creates a new instance of OperationClient.
func (c *ClientFactory) NewOperationClient() *OperationClient {
	return &OperationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointClient creates a new instance of PrivateEndpointClient.
func (c *ClientFactory) NewPrivateEndpointClient() *PrivateEndpointClient {
	return &PrivateEndpointClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionClient creates a new instance of PrivateEndpointConnectionClient.
func (c *ClientFactory) NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient {
	return &PrivateEndpointConnectionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectableContainersClient creates a new instance of ProtectableContainersClient.
func (c *ClientFactory) NewProtectableContainersClient() *ProtectableContainersClient {
	return &ProtectableContainersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectedItemOperationResultsClient creates a new instance of ProtectedItemOperationResultsClient.
func (c *ClientFactory) NewProtectedItemOperationResultsClient() *ProtectedItemOperationResultsClient {
	return &ProtectedItemOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectedItemOperationStatusesClient creates a new instance of ProtectedItemOperationStatusesClient.
func (c *ClientFactory) NewProtectedItemOperationStatusesClient() *ProtectedItemOperationStatusesClient {
	return &ProtectedItemOperationStatusesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectedItemsClient creates a new instance of ProtectedItemsClient.
func (c *ClientFactory) NewProtectedItemsClient() *ProtectedItemsClient {
	return &ProtectedItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionContainerOperationResultsClient creates a new instance of ProtectionContainerOperationResultsClient.
func (c *ClientFactory) NewProtectionContainerOperationResultsClient() *ProtectionContainerOperationResultsClient {
	return &ProtectionContainerOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionContainerRefreshOperationResultsClient creates a new instance of ProtectionContainerRefreshOperationResultsClient.
func (c *ClientFactory) NewProtectionContainerRefreshOperationResultsClient() *ProtectionContainerRefreshOperationResultsClient {
	return &ProtectionContainerRefreshOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionContainersClient creates a new instance of ProtectionContainersClient.
func (c *ClientFactory) NewProtectionContainersClient() *ProtectionContainersClient {
	return &ProtectionContainersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionIntentClient creates a new instance of ProtectionIntentClient.
func (c *ClientFactory) NewProtectionIntentClient() *ProtectionIntentClient {
	return &ProtectionIntentClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionPoliciesClient creates a new instance of ProtectionPoliciesClient.
func (c *ClientFactory) NewProtectionPoliciesClient() *ProtectionPoliciesClient {
	return &ProtectionPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionPolicyOperationResultsClient creates a new instance of ProtectionPolicyOperationResultsClient.
func (c *ClientFactory) NewProtectionPolicyOperationResultsClient() *ProtectionPolicyOperationResultsClient {
	return &ProtectionPolicyOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProtectionPolicyOperationStatusesClient creates a new instance of ProtectionPolicyOperationStatusesClient.
func (c *ClientFactory) NewProtectionPolicyOperationStatusesClient() *ProtectionPolicyOperationStatusesClient {
	return &ProtectionPolicyOperationStatusesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRecoveryPointsClient creates a new instance of RecoveryPointsClient.
func (c *ClientFactory) NewRecoveryPointsClient() *RecoveryPointsClient {
	return &RecoveryPointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRecoveryPointsRecommendedForMoveClient creates a new instance of RecoveryPointsRecommendedForMoveClient.
func (c *ClientFactory) NewRecoveryPointsRecommendedForMoveClient() *RecoveryPointsRecommendedForMoveClient {
	return &RecoveryPointsRecommendedForMoveClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewResourceGuardProxiesClient creates a new instance of ResourceGuardProxiesClient.
func (c *ClientFactory) NewResourceGuardProxiesClient() *ResourceGuardProxiesClient {
	return &ResourceGuardProxiesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewResourceGuardProxyClient creates a new instance of ResourceGuardProxyClient.
func (c *ClientFactory) NewResourceGuardProxyClient() *ResourceGuardProxyClient {
	return &ResourceGuardProxyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRestoresClient creates a new instance of RestoresClient.
func (c *ClientFactory) NewRestoresClient() *RestoresClient {
	return &RestoresClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSecurityPINsClient creates a new instance of SecurityPINsClient.
func (c *ClientFactory) NewSecurityPINsClient() *SecurityPINsClient {
	return &SecurityPINsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTieringCostOperationStatusClient creates a new instance of TieringCostOperationStatusClient.
func (c *ClientFactory) NewTieringCostOperationStatusClient() *TieringCostOperationStatusClient {
	return &TieringCostOperationStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewValidateOperationClient creates a new instance of ValidateOperationClient.
func (c *ClientFactory) NewValidateOperationClient() *ValidateOperationClient {
	return &ValidateOperationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewValidateOperationResultsClient creates a new instance of ValidateOperationResultsClient.
func (c *ClientFactory) NewValidateOperationResultsClient() *ValidateOperationResultsClient {
	return &ValidateOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewValidateOperationStatusesClient creates a new instance of ValidateOperationStatusesClient.
func (c *ClientFactory) NewValidateOperationStatusesClient() *ValidateOperationStatusesClient {
	return &ValidateOperationStatusesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
