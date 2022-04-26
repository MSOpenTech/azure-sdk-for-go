//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

const (
	moduleName    = "armrecoveryservices"
	moduleVersion = "v0.4.1"
)

// AuthType - Specifies the authentication type.
type AuthType string

const (
	AuthTypeAAD                  AuthType = "AAD"
	AuthTypeACS                  AuthType = "ACS"
	AuthTypeAccessControlService AuthType = "AccessControlService"
	AuthTypeAzureActiveDirectory AuthType = "AzureActiveDirectory"
	AuthTypeInvalid              AuthType = "Invalid"
)

// PossibleAuthTypeValues returns the possible values for the AuthType const type.
func PossibleAuthTypeValues() []AuthType {
	return []AuthType{
		AuthTypeAAD,
		AuthTypeACS,
		AuthTypeAccessControlService,
		AuthTypeAzureActiveDirectory,
		AuthTypeInvalid,
	}
}

// BackupStorageVersion - Backup storage version
type BackupStorageVersion string

const (
	BackupStorageVersionUnassigned BackupStorageVersion = "Unassigned"
	BackupStorageVersionV1         BackupStorageVersion = "V1"
	BackupStorageVersionV2         BackupStorageVersion = "V2"
)

// PossibleBackupStorageVersionValues returns the possible values for the BackupStorageVersion const type.
func PossibleBackupStorageVersionValues() []BackupStorageVersion {
	return []BackupStorageVersion{
		BackupStorageVersionUnassigned,
		BackupStorageVersionV1,
		BackupStorageVersionV2,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// InfrastructureEncryptionState - Enabling/Disabling the Double Encryption state
type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = "Disabled"
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = "Enabled"
)

// PossibleInfrastructureEncryptionStateValues returns the possible values for the InfrastructureEncryptionState const type.
func PossibleInfrastructureEncryptionStateValues() []InfrastructureEncryptionState {
	return []InfrastructureEncryptionState{
		InfrastructureEncryptionStateDisabled,
		InfrastructureEncryptionStateEnabled,
	}
}

// PrivateEndpointConnectionStatus - Gets or sets the status.
type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusApproved     PrivateEndpointConnectionStatus = "Approved"
	PrivateEndpointConnectionStatusDisconnected PrivateEndpointConnectionStatus = "Disconnected"
	PrivateEndpointConnectionStatusPending      PrivateEndpointConnectionStatus = "Pending"
	PrivateEndpointConnectionStatusRejected     PrivateEndpointConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointConnectionStatusValues returns the possible values for the PrivateEndpointConnectionStatus const type.
func PossiblePrivateEndpointConnectionStatusValues() []PrivateEndpointConnectionStatus {
	return []PrivateEndpointConnectionStatus{
		PrivateEndpointConnectionStatusApproved,
		PrivateEndpointConnectionStatusDisconnected,
		PrivateEndpointConnectionStatusPending,
		PrivateEndpointConnectionStatusRejected,
	}
}

// ProvisioningState - Gets or sets provisioning state of the private endpoint connection.
type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStatePending   ProvisioningState = "Pending"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStatePending,
		ProvisioningStateSucceeded,
	}
}

// ResourceIdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly
// created identity and a set of user-assigned identities. The type 'None' will remove any
// identities.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// ResourceMoveState - The State of the Resource after the move operation
type ResourceMoveState string

const (
	ResourceMoveStateCommitFailed    ResourceMoveState = "CommitFailed"
	ResourceMoveStateCommitTimedout  ResourceMoveState = "CommitTimedout"
	ResourceMoveStateCriticalFailure ResourceMoveState = "CriticalFailure"
	ResourceMoveStateFailure         ResourceMoveState = "Failure"
	ResourceMoveStateInProgress      ResourceMoveState = "InProgress"
	ResourceMoveStateMoveSucceeded   ResourceMoveState = "MoveSucceeded"
	ResourceMoveStatePartialSuccess  ResourceMoveState = "PartialSuccess"
	ResourceMoveStatePrepareFailed   ResourceMoveState = "PrepareFailed"
	ResourceMoveStatePrepareTimedout ResourceMoveState = "PrepareTimedout"
	ResourceMoveStateUnknown         ResourceMoveState = "Unknown"
)

// PossibleResourceMoveStateValues returns the possible values for the ResourceMoveState const type.
func PossibleResourceMoveStateValues() []ResourceMoveState {
	return []ResourceMoveState{
		ResourceMoveStateCommitFailed,
		ResourceMoveStateCommitTimedout,
		ResourceMoveStateCriticalFailure,
		ResourceMoveStateFailure,
		ResourceMoveStateInProgress,
		ResourceMoveStateMoveSucceeded,
		ResourceMoveStatePartialSuccess,
		ResourceMoveStatePrepareFailed,
		ResourceMoveStatePrepareTimedout,
		ResourceMoveStateUnknown,
	}
}

// SKUName - The Sku name.
type SKUName string

const (
	SKUNameRS0      SKUName = "RS0"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameRS0,
		SKUNameStandard,
	}
}

// TriggerType - The way the vault upgrade was triggered.
type TriggerType string

const (
	TriggerTypeForcedUpgrade TriggerType = "ForcedUpgrade"
	TriggerTypeUserTriggered TriggerType = "UserTriggered"
)

// PossibleTriggerTypeValues returns the possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{
		TriggerTypeForcedUpgrade,
		TriggerTypeUserTriggered,
	}
}

// UsagesUnit - Unit of the usage.
type UsagesUnit string

const (
	UsagesUnitBytes          UsagesUnit = "Bytes"
	UsagesUnitBytesPerSecond UsagesUnit = "BytesPerSecond"
	UsagesUnitCount          UsagesUnit = "Count"
	UsagesUnitCountPerSecond UsagesUnit = "CountPerSecond"
	UsagesUnitPercent        UsagesUnit = "Percent"
	UsagesUnitSeconds        UsagesUnit = "Seconds"
)

// PossibleUsagesUnitValues returns the possible values for the UsagesUnit const type.
func PossibleUsagesUnitValues() []UsagesUnit {
	return []UsagesUnit{
		UsagesUnitBytes,
		UsagesUnitBytesPerSecond,
		UsagesUnitCount,
		UsagesUnitCountPerSecond,
		UsagesUnitPercent,
		UsagesUnitSeconds,
	}
}

// VaultPrivateEndpointState - Private endpoint state for backup.
type VaultPrivateEndpointState string

const (
	VaultPrivateEndpointStateEnabled VaultPrivateEndpointState = "Enabled"
	VaultPrivateEndpointStateNone    VaultPrivateEndpointState = "None"
)

// PossibleVaultPrivateEndpointStateValues returns the possible values for the VaultPrivateEndpointState const type.
func PossibleVaultPrivateEndpointStateValues() []VaultPrivateEndpointState {
	return []VaultPrivateEndpointState{
		VaultPrivateEndpointStateEnabled,
		VaultPrivateEndpointStateNone,
	}
}

// VaultUpgradeState - Status of the vault upgrade operation.
type VaultUpgradeState string

const (
	VaultUpgradeStateFailed     VaultUpgradeState = "Failed"
	VaultUpgradeStateInProgress VaultUpgradeState = "InProgress"
	VaultUpgradeStateUnknown    VaultUpgradeState = "Unknown"
	VaultUpgradeStateUpgraded   VaultUpgradeState = "Upgraded"
)

// PossibleVaultUpgradeStateValues returns the possible values for the VaultUpgradeState const type.
func PossibleVaultUpgradeStateValues() []VaultUpgradeState {
	return []VaultUpgradeState{
		VaultUpgradeStateFailed,
		VaultUpgradeStateInProgress,
		VaultUpgradeStateUnknown,
		VaultUpgradeStateUpgraded,
	}
}
