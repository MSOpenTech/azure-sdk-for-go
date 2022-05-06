package healthcareapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionType enumerates the values for action type.
type ActionType string

const (
	// ActionTypeInternal ...
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{ActionTypeInternal}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// FhirServiceKind enumerates the values for fhir service kind.
type FhirServiceKind string

const (
	// FhirServiceKindFhirR4 ...
	FhirServiceKindFhirR4 FhirServiceKind = "fhir-R4"
	// FhirServiceKindFhirStu3 ...
	FhirServiceKindFhirStu3 FhirServiceKind = "fhir-Stu3"
)

// PossibleFhirServiceKindValues returns an array of possible values for the FhirServiceKind const type.
func PossibleFhirServiceKindValues() []FhirServiceKind {
	return []FhirServiceKind{FhirServiceKindFhirR4, FhirServiceKindFhirStu3}
}

// IotIdentityResolutionType enumerates the values for iot identity resolution type.
type IotIdentityResolutionType string

const (
	// IotIdentityResolutionTypeCreate ...
	IotIdentityResolutionTypeCreate IotIdentityResolutionType = "Create"
	// IotIdentityResolutionTypeLookup ...
	IotIdentityResolutionTypeLookup IotIdentityResolutionType = "Lookup"
)

// PossibleIotIdentityResolutionTypeValues returns an array of possible values for the IotIdentityResolutionType const type.
func PossibleIotIdentityResolutionTypeValues() []IotIdentityResolutionType {
	return []IotIdentityResolutionType{IotIdentityResolutionTypeCreate, IotIdentityResolutionTypeLookup}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindFhir ...
	KindFhir Kind = "fhir"
	// KindFhirR4 ...
	KindFhirR4 Kind = "fhir-R4"
	// KindFhirStu3 ...
	KindFhirStu3 Kind = "fhir-Stu3"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindFhir, KindFhirR4, KindFhirStu3}
}

// ManagedServiceIdentityType enumerates the values for managed service identity type.
type ManagedServiceIdentityType string

const (
	// ManagedServiceIdentityTypeNone ...
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	// ManagedServiceIdentityTypeSystemAssigned ...
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns an array of possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{ManagedServiceIdentityTypeNone, ManagedServiceIdentityTypeSystemAssigned}
}

// OperationResultStatus enumerates the values for operation result status.
type OperationResultStatus string

const (
	// OperationResultStatusCanceled ...
	OperationResultStatusCanceled OperationResultStatus = "Canceled"
	// OperationResultStatusFailed ...
	OperationResultStatusFailed OperationResultStatus = "Failed"
	// OperationResultStatusRequested ...
	OperationResultStatusRequested OperationResultStatus = "Requested"
	// OperationResultStatusRunning ...
	OperationResultStatusRunning OperationResultStatus = "Running"
	// OperationResultStatusSucceeded ...
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
)

// PossibleOperationResultStatusValues returns an array of possible values for the OperationResultStatus const type.
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{OperationResultStatusCanceled, OperationResultStatusFailed, OperationResultStatusRequested, OperationResultStatusRunning, OperationResultStatusSucceeded}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating ...
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting ...
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed ...
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded ...
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateCreating, PrivateEndpointConnectionProvisioningStateDeleting, PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// PrivateEndpointServiceConnectionStatusApproved ...
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	// PrivateEndpointServiceConnectionStatusPending ...
	PrivateEndpointServiceConnectionStatusPending PrivateEndpointServiceConnectionStatus = "Pending"
	// PrivateEndpointServiceConnectionStatusRejected ...
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{PrivateEndpointServiceConnectionStatusApproved, PrivateEndpointServiceConnectionStatusPending, PrivateEndpointServiceConnectionStatusRejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateAccepted ...
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateDeprovisioned ...
	ProvisioningStateDeprovisioned ProvisioningState = "Deprovisioned"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMoving ...
	ProvisioningStateMoving ProvisioningState = "Moving"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateSuspended ...
	ProvisioningStateSuspended ProvisioningState = "Suspended"
	// ProvisioningStateSystemMaintenance ...
	ProvisioningStateSystemMaintenance ProvisioningState = "SystemMaintenance"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
	// ProvisioningStateVerifying ...
	ProvisioningStateVerifying ProvisioningState = "Verifying"
	// ProvisioningStateWarned ...
	ProvisioningStateWarned ProvisioningState = "Warned"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateAccepted, ProvisioningStateCanceled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateDeprovisioned, ProvisioningStateFailed, ProvisioningStateMoving, ProvisioningStateSucceeded, ProvisioningStateSuspended, ProvisioningStateSystemMaintenance, ProvisioningStateUpdating, ProvisioningStateVerifying, ProvisioningStateWarned}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled ...
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	// PublicNetworkAccessEnabled ...
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{PublicNetworkAccessDisabled, PublicNetworkAccessEnabled}
}

// ServiceNameUnavailabilityReason enumerates the values for service name unavailability reason.
type ServiceNameUnavailabilityReason string

const (
	// ServiceNameUnavailabilityReasonAlreadyExists ...
	ServiceNameUnavailabilityReasonAlreadyExists ServiceNameUnavailabilityReason = "AlreadyExists"
	// ServiceNameUnavailabilityReasonInvalid ...
	ServiceNameUnavailabilityReasonInvalid ServiceNameUnavailabilityReason = "Invalid"
)

// PossibleServiceNameUnavailabilityReasonValues returns an array of possible values for the ServiceNameUnavailabilityReason const type.
func PossibleServiceNameUnavailabilityReasonValues() []ServiceNameUnavailabilityReason {
	return []ServiceNameUnavailabilityReason{ServiceNameUnavailabilityReasonAlreadyExists, ServiceNameUnavailabilityReasonInvalid}
}
