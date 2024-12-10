//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceregistry

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
	moduleVersion = "v0.1.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
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

// DataPointsObservabilityMode - An indication of how the data point should be mapped to OpenTelemetry.
type DataPointsObservabilityMode string

const (
	// DataPointsObservabilityModeCounter - Map as counter to OpenTelemetry.
	DataPointsObservabilityModeCounter DataPointsObservabilityMode = "counter"
	// DataPointsObservabilityModeGauge - Map as gauge to OpenTelemetry.
	DataPointsObservabilityModeGauge DataPointsObservabilityMode = "gauge"
	// DataPointsObservabilityModeHistogram - Map as histogram to OpenTelemetry.
	DataPointsObservabilityModeHistogram DataPointsObservabilityMode = "histogram"
	// DataPointsObservabilityModeLog - Map as log to OpenTelemetry.
	DataPointsObservabilityModeLog DataPointsObservabilityMode = "log"
	// DataPointsObservabilityModeNone - No mapping to OpenTelemetry.
	DataPointsObservabilityModeNone DataPointsObservabilityMode = "none"
)

// PossibleDataPointsObservabilityModeValues returns the possible values for the DataPointsObservabilityMode const type.
func PossibleDataPointsObservabilityModeValues() []DataPointsObservabilityMode {
	return []DataPointsObservabilityMode{
		DataPointsObservabilityModeCounter,
		DataPointsObservabilityModeGauge,
		DataPointsObservabilityModeHistogram,
		DataPointsObservabilityModeLog,
		DataPointsObservabilityModeNone,
	}
}

// EventsObservabilityMode - An indication of how the event should be mapped to OpenTelemetry.
type EventsObservabilityMode string

const (
	// EventsObservabilityModeLog - Map as log to OpenTelemetry.
	EventsObservabilityModeLog EventsObservabilityMode = "log"
	// EventsObservabilityModeNone - No mapping to OpenTelemetry.
	EventsObservabilityModeNone EventsObservabilityMode = "none"
)

// PossibleEventsObservabilityModeValues returns the possible values for the EventsObservabilityMode const type.
func PossibleEventsObservabilityModeValues() []EventsObservabilityMode {
	return []EventsObservabilityMode{
		EventsObservabilityModeLog,
		EventsObservabilityModeNone,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
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

// ProvisioningState - The provisioning status of the resource.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Resource has been accepted by the server.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// UserAuthenticationMode - Defines the method to authenticate the user of the client at the server.
type UserAuthenticationMode string

const (
	// UserAuthenticationModeAnonymous - The user authentication method is anonymous.
	UserAuthenticationModeAnonymous UserAuthenticationMode = "Anonymous"
	// UserAuthenticationModeCertificate - The user authentication method is an x509 certificate.
	UserAuthenticationModeCertificate UserAuthenticationMode = "Certificate"
	// UserAuthenticationModeUsernamePassword - The user authentication method is a username and password.
	UserAuthenticationModeUsernamePassword UserAuthenticationMode = "UsernamePassword"
)

// PossibleUserAuthenticationModeValues returns the possible values for the UserAuthenticationMode const type.
func PossibleUserAuthenticationModeValues() []UserAuthenticationMode {
	return []UserAuthenticationMode{
		UserAuthenticationModeAnonymous,
		UserAuthenticationModeCertificate,
		UserAuthenticationModeUsernamePassword,
	}
}
