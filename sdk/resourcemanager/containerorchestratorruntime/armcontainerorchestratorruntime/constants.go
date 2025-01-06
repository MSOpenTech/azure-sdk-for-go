// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcontainerorchestratorruntime

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerorchestratorruntime/armcontainerorchestratorruntime"
	moduleVersion = "v0.1.0"
)

// AccessMode - Storage Class Access Mode
type AccessMode string

const (
	// AccessModeReadWriteMany - Read Write Many (RWX) access mode
	AccessModeReadWriteMany AccessMode = "ReadWriteMany"
	// AccessModeReadWriteOnce - Read Write Once (RWO) access mode
	AccessModeReadWriteOnce AccessMode = "ReadWriteOnce"
)

// PossibleAccessModeValues returns the possible values for the AccessMode const type.
func PossibleAccessModeValues() []AccessMode {
	return []AccessMode{
		AccessModeReadWriteMany,
		AccessModeReadWriteOnce,
	}
}

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AdvertiseMode - Enum of advertise mode
type AdvertiseMode string

const (
	// AdvertiseModeARP - ARP advertise mode
	AdvertiseModeARP AdvertiseMode = "ARP"
	// AdvertiseModeBGP - BGP advertise mode
	AdvertiseModeBGP AdvertiseMode = "BGP"
	// AdvertiseModeBoth - both ARP and BGP advertise mode
	AdvertiseModeBoth AdvertiseMode = "Both"
)

// PossibleAdvertiseModeValues returns the possible values for the AdvertiseMode const type.
func PossibleAdvertiseModeValues() []AdvertiseMode {
	return []AdvertiseMode{
		AdvertiseModeARP,
		AdvertiseModeBGP,
		AdvertiseModeBoth,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
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

// DataResilienceTier - Data resilience tier of a storage class
type DataResilienceTier string

const (
	// DataResilienceTierDataResilient - Data resilient
	DataResilienceTierDataResilient DataResilienceTier = "DataResilient"
	// DataResilienceTierNotDataResilient - Not data resilient
	DataResilienceTierNotDataResilient DataResilienceTier = "NotDataResilient"
)

// PossibleDataResilienceTierValues returns the possible values for the DataResilienceTier const type.
func PossibleDataResilienceTierValues() []DataResilienceTier {
	return []DataResilienceTier{
		DataResilienceTierDataResilient,
		DataResilienceTierNotDataResilient,
	}
}

// FailoverTier - Failover tier of a storage class
type FailoverTier string

const (
	// FailoverTierFast - Fast Failover Tier
	FailoverTierFast FailoverTier = "Fast"
	// FailoverTierNotAvailable - Not available Failover Tier
	FailoverTierNotAvailable FailoverTier = "NotAvailable"
	// FailoverTierSlow - Slow Failover Tier
	FailoverTierSlow FailoverTier = "Slow"
	// FailoverTierSuper - Super Failover Tier
	FailoverTierSuper FailoverTier = "Super"
)

// PossibleFailoverTierValues returns the possible values for the FailoverTier const type.
func PossibleFailoverTierValues() []FailoverTier {
	return []FailoverTier{
		FailoverTierFast,
		FailoverTierNotAvailable,
		FailoverTierSlow,
		FailoverTierSuper,
	}
}

// NfsDirectoryActionOnVolumeDeletion - The action to take when a NFS volume is deleted
type NfsDirectoryActionOnVolumeDeletion string

const (
	// NfsDirectoryActionOnVolumeDeletionDelete - When the volume is deleted, delete the directory
	NfsDirectoryActionOnVolumeDeletionDelete NfsDirectoryActionOnVolumeDeletion = "Delete"
	// NfsDirectoryActionOnVolumeDeletionRetain - When the volume is deleted, retain the directory
	NfsDirectoryActionOnVolumeDeletionRetain NfsDirectoryActionOnVolumeDeletion = "Retain"
)

// PossibleNfsDirectoryActionOnVolumeDeletionValues returns the possible values for the NfsDirectoryActionOnVolumeDeletion const type.
func PossibleNfsDirectoryActionOnVolumeDeletionValues() []NfsDirectoryActionOnVolumeDeletion {
	return []NfsDirectoryActionOnVolumeDeletion{
		NfsDirectoryActionOnVolumeDeletionDelete,
		NfsDirectoryActionOnVolumeDeletionRetain,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PerformanceTier - Performance tier of a storage class
type PerformanceTier string

const (
	// PerformanceTierBasic - Basic Performance Tier
	PerformanceTierBasic PerformanceTier = "Basic"
	// PerformanceTierPremium - Premium Performance Tier
	PerformanceTierPremium PerformanceTier = "Premium"
	// PerformanceTierStandard - Standard Performance Tier
	PerformanceTierStandard PerformanceTier = "Standard"
	// PerformanceTierUltra - Ultra Performance Tier
	PerformanceTierUltra PerformanceTier = "Ultra"
	// PerformanceTierUndefined - Undefined Performance Tier
	PerformanceTierUndefined PerformanceTier = "Undefined"
)

// PossiblePerformanceTierValues returns the possible values for the PerformanceTier const type.
func PossiblePerformanceTierValues() []PerformanceTier {
	return []PerformanceTier{
		PerformanceTierBasic,
		PerformanceTierPremium,
		PerformanceTierStandard,
		PerformanceTierUltra,
		PerformanceTierUndefined,
	}
}

// ProvisioningState - The status of the current operation.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Change accepted for processing
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Deletion in progress
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - Initial provisioning in progress
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Update in progress
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SCType - Type of a storage class
type SCType string

const (
	// SCTypeBlob - Blob storage class
	SCTypeBlob SCType = "Blob"
	// SCTypeNFS - NFS storage class
	SCTypeNFS SCType = "NFS"
	// SCTypeNative - Native storage class
	SCTypeNative SCType = "Native"
	// SCTypeRWX - RWX storage class
	SCTypeRWX SCType = "RWX"
	// SCTypeSMB - SMB storage class
	SCTypeSMB SCType = "SMB"
)

// PossibleSCTypeValues returns the possible values for the SCType const type.
func PossibleSCTypeValues() []SCType {
	return []SCType{
		SCTypeBlob,
		SCTypeNFS,
		SCTypeNative,
		SCTypeRWX,
		SCTypeSMB,
	}
}

// VolumeBindingMode - Storage class volume binding mode
type VolumeBindingMode string

const (
	// VolumeBindingModeImmediate - Immediate binding mode
	VolumeBindingModeImmediate VolumeBindingMode = "Immediate"
	// VolumeBindingModeWaitForFirstConsumer - Wait for first consumer binding mode
	VolumeBindingModeWaitForFirstConsumer VolumeBindingMode = "WaitForFirstConsumer"
)

// PossibleVolumeBindingModeValues returns the possible values for the VolumeBindingMode const type.
func PossibleVolumeBindingModeValues() []VolumeBindingMode {
	return []VolumeBindingMode{
		VolumeBindingModeImmediate,
		VolumeBindingModeWaitForFirstConsumer,
	}
}

// VolumeExpansion - Ability to expand volumes of a storage class
type VolumeExpansion string

const (
	// VolumeExpansionAllow - Allow volume expansion
	VolumeExpansionAllow VolumeExpansion = "Allow"
	// VolumeExpansionDisallow - Disallow volume expansion
	VolumeExpansionDisallow VolumeExpansion = "Disallow"
)

// PossibleVolumeExpansionValues returns the possible values for the VolumeExpansion const type.
func PossibleVolumeExpansionValues() []VolumeExpansion {
	return []VolumeExpansion{
		VolumeExpansionAllow,
		VolumeExpansionDisallow,
	}
}
