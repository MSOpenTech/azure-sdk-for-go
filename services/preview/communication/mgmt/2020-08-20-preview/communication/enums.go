package communication

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AggregationType enumerates the values for aggregation type.
type AggregationType string

const (
	// Average ...
	Average AggregationType = "Average"
	// Count ...
	Count AggregationType = "Count"
	// Maximum ...
	Maximum AggregationType = "Maximum"
	// Minimum ...
	Minimum AggregationType = "Minimum"
	// Total ...
	Total AggregationType = "Total"
)

// PossibleAggregationTypeValues returns an array of possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{Average, Count, Maximum, Minimum, Total}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "Primary"
	// Secondary ...
	Secondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Unknown ...
	Unknown ProvisioningState = "Unknown"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Deleting, Failed, Moving, Running, Succeeded, Unknown, Updating}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusCanceled ...
	StatusCanceled Status = "Canceled"
	// StatusCreating ...
	StatusCreating Status = "Creating"
	// StatusDeleting ...
	StatusDeleting Status = "Deleting"
	// StatusFailed ...
	StatusFailed Status = "Failed"
	// StatusMoving ...
	StatusMoving Status = "Moving"
	// StatusSucceeded ...
	StatusSucceeded Status = "Succeeded"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusCanceled, StatusCreating, StatusDeleting, StatusFailed, StatusMoving, StatusSucceeded}
}
