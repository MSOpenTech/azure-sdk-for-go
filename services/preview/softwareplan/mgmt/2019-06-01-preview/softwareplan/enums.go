package softwareplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ErrorCode enumerates the values for error code.
type ErrorCode string

const (
	// InvalidRequestParameter ...
	InvalidRequestParameter ErrorCode = "InvalidRequestParameter"
	// MissingRequestParameter ...
	MissingRequestParameter ErrorCode = "MissingRequestParameter"
)

// PossibleErrorCodeValues returns an array of possible values for the ErrorCode const type.
func PossibleErrorCodeValues() []ErrorCode {
	return []ErrorCode{InvalidRequestParameter, MissingRequestParameter}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Cancelled ...
	Cancelled ProvisioningState = "Cancelled"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Cancelled, Failed, Succeeded}
}
