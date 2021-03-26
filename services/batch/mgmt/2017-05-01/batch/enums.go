package batch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccountKeyType enumerates the values for account key type.
type AccountKeyType string

const (
	// Primary ...
	Primary AccountKeyType = "Primary"
	// Secondary ...
	Secondary AccountKeyType = "Secondary"
)

// PossibleAccountKeyTypeValues returns an array of possible values for the AccountKeyType const type.
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return []AccountKeyType{Primary, Secondary}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameAvailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameAvailabilityReason = "Invalid"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{AlreadyExists, Invalid}
}

// PackageState enumerates the values for package state.
type PackageState string

const (
	// Active ...
	Active PackageState = "active"
	// Pending ...
	Pending PackageState = "pending"
	// Unmapped ...
	Unmapped PackageState = "unmapped"
)

// PossiblePackageStateValues returns an array of possible values for the PackageState const type.
func PossiblePackageStateValues() []PackageState {
	return []PackageState{Active, Pending, Unmapped}
}

// PoolAllocationMode enumerates the values for pool allocation mode.
type PoolAllocationMode string

const (
	// BatchService ...
	BatchService PoolAllocationMode = "BatchService"
	// UserSubscription ...
	UserSubscription PoolAllocationMode = "UserSubscription"
)

// PossiblePoolAllocationModeValues returns an array of possible values for the PoolAllocationMode const type.
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return []PoolAllocationMode{BatchService, UserSubscription}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCancelled ...
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateInvalid ...
	ProvisioningStateInvalid ProvisioningState = "Invalid"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCancelled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateInvalid, ProvisioningStateSucceeded}
}
