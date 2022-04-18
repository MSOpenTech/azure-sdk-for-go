//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsans

const (
	moduleName    = "armelasticsans"
	moduleVersion = "v0.1.0"
)

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

// EncryptionType - The type of key used to encrypt the data of the disk.
type EncryptionType string

const (
	// EncryptionTypeEncryptionAtRestWithCustomerKey - Volume is encrypted at rest with Customer managed key that can be changed
	// and revoked by a customer.
	EncryptionTypeEncryptionAtRestWithCustomerKey EncryptionType = "EncryptionAtRestWithCustomerKey"
	// EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys - Volume is encrypted at rest with 2 layers of encryption. One
	// of the keys is Customer managed and the other key is Platform managed.
	EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys EncryptionType = "EncryptionAtRestWithPlatformAndCustomerKeys"
	// EncryptionTypeEncryptionAtRestWithPlatformKey - Volume is encrypted at rest with Platform managed key. It is the default
	// encryption type. This is not a valid encryption type for disk encryption sets.
	EncryptionTypeEncryptionAtRestWithPlatformKey EncryptionType = "EncryptionAtRestWithPlatformKey"
)

// PossibleEncryptionTypeValues returns the possible values for the EncryptionType const type.
func PossibleEncryptionTypeValues() []EncryptionType {
	return []EncryptionType{
		EncryptionTypeEncryptionAtRestWithCustomerKey,
		EncryptionTypeEncryptionAtRestWithPlatformAndCustomerKeys,
		EncryptionTypeEncryptionAtRestWithPlatformKey,
	}
}

// Name - The sku name.
type Name string

const (
	// NamePremiumLRS - Premium locally redundant storage
	NamePremiumLRS Name = "Premium_LRS"
	// NameStandardLRS - Standard locally redundant storage
	NameStandardLRS Name = "Standard_LRS"
	// NameStandardZRS - Standard zone redundant storage
	NameStandardZRS Name = "Standard_ZRS"
)

// PossibleNameValues returns the possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{
		NamePremiumLRS,
		NameStandardLRS,
		NameStandardZRS,
	}
}

// OperationalStatus - Operational status of the resource.
type OperationalStatus string

const (
	OperationalStatusHealthy            OperationalStatus = "Healthy"
	OperationalStatusInvalid            OperationalStatus = "Invalid"
	OperationalStatusRunning            OperationalStatus = "Running"
	OperationalStatusStopped            OperationalStatus = "Stopped"
	OperationalStatusStoppedDeallocated OperationalStatus = "Stopped (deallocated)"
	OperationalStatusUnhealthy          OperationalStatus = "Unhealthy"
	OperationalStatusUnknown            OperationalStatus = "Unknown"
	OperationalStatusUpdating           OperationalStatus = "Updating"
)

// PossibleOperationalStatusValues returns the possible values for the OperationalStatus const type.
func PossibleOperationalStatusValues() []OperationalStatus {
	return []OperationalStatus{
		OperationalStatusHealthy,
		OperationalStatusInvalid,
		OperationalStatusRunning,
		OperationalStatusStopped,
		OperationalStatusStoppedDeallocated,
		OperationalStatusUnhealthy,
		OperationalStatusUnknown,
		OperationalStatusUpdating,
	}
}

// ProvisioningStates - Provisioning state of the iSCSI Target.
type ProvisioningStates string

const (
	ProvisioningStatesCanceled  ProvisioningStates = "Canceled"
	ProvisioningStatesCreating  ProvisioningStates = "Creating"
	ProvisioningStatesDeleting  ProvisioningStates = "Deleting"
	ProvisioningStatesFailed    ProvisioningStates = "Failed"
	ProvisioningStatesInvalid   ProvisioningStates = "Invalid"
	ProvisioningStatesPending   ProvisioningStates = "Pending"
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	ProvisioningStatesUpdating  ProvisioningStates = "Updating"
)

// PossibleProvisioningStatesValues returns the possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{
		ProvisioningStatesCanceled,
		ProvisioningStatesCreating,
		ProvisioningStatesDeleting,
		ProvisioningStatesFailed,
		ProvisioningStatesInvalid,
		ProvisioningStatesPending,
		ProvisioningStatesSucceeded,
		ProvisioningStatesUpdating,
	}
}

// State - Gets the state of virtual network rule.
type State string

const (
	StateProvisioning         State = "provisioning"
	StateDeprovisioning       State = "deprovisioning"
	StateSucceeded            State = "succeeded"
	StateFailed               State = "failed"
	StateNetworkSourceDeleted State = "networkSourceDeleted"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateProvisioning,
		StateDeprovisioning,
		StateSucceeded,
		StateFailed,
		StateNetworkSourceDeleted,
	}
}

// StorageTargetType - Storage Target type.
type StorageTargetType string

const (
	StorageTargetTypeIscsi StorageTargetType = "Iscsi"
	StorageTargetTypeNone  StorageTargetType = "None"
)

// PossibleStorageTargetTypeValues returns the possible values for the StorageTargetType const type.
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return []StorageTargetType{
		StorageTargetTypeIscsi,
		StorageTargetTypeNone,
	}
}

// Tier - The sku tier.
type Tier string

const (
	// TierHero - Hero
	TierHero Tier = "Hero"
	// TierHub - Hub
	TierHub Tier = "Hub"
	// TierSatellite - Satellite
	TierSatellite Tier = "Satellite"
)

// PossibleTierValues returns the possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{
		TierHero,
		TierHub,
		TierSatellite,
	}
}

// VolumeCreateOption - This enumerates the possible sources of a volume creation.
type VolumeCreateOption string

const (
	VolumeCreateOptionNone             VolumeCreateOption = "None"
	VolumeCreateOptionFromVolume       VolumeCreateOption = "FromVolume"
	VolumeCreateOptionFromDiskSnapshot VolumeCreateOption = "FromDiskSnapshot"
	VolumeCreateOptionExport           VolumeCreateOption = "Export"
)

// PossibleVolumeCreateOptionValues returns the possible values for the VolumeCreateOption const type.
func PossibleVolumeCreateOptionValues() []VolumeCreateOption {
	return []VolumeCreateOption{
		VolumeCreateOptionNone,
		VolumeCreateOptionFromVolume,
		VolumeCreateOptionFromDiskSnapshot,
		VolumeCreateOptionExport,
	}
}
