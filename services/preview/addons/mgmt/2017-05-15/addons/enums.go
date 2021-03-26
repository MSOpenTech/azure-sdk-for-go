package addons

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// PlanTypeName enumerates the values for plan type name.
type PlanTypeName string

const (
	// Advanced ...
	Advanced PlanTypeName = "Advanced"
	// Essential ...
	Essential PlanTypeName = "Essential"
	// Standard ...
	Standard PlanTypeName = "Standard"
)

// PossiblePlanTypeNameValues returns an array of possible values for the PlanTypeName const type.
func PossiblePlanTypeNameValues() []PlanTypeName {
	return []PlanTypeName{Advanced, Essential, Standard}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Cancelled ...
	Cancelled ProvisioningState = "Cancelled"
	// Cancelling ...
	Cancelling ProvisioningState = "Cancelling"
	// Downgrading ...
	Downgrading ProvisioningState = "Downgrading"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Purchasing ...
	Purchasing ProvisioningState = "Purchasing"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Upgrading ...
	Upgrading ProvisioningState = "Upgrading"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Cancelled, Cancelling, Downgrading, Failed, Purchasing, Succeeded, Upgrading}
}

// SupportPlanType enumerates the values for support plan type.
type SupportPlanType string

const (
	// SupportPlanTypeAdvanced ...
	SupportPlanTypeAdvanced SupportPlanType = "advanced"
	// SupportPlanTypeEssential ...
	SupportPlanTypeEssential SupportPlanType = "essential"
	// SupportPlanTypeStandard ...
	SupportPlanTypeStandard SupportPlanType = "standard"
)

// PossibleSupportPlanTypeValues returns an array of possible values for the SupportPlanType const type.
func PossibleSupportPlanTypeValues() []SupportPlanType {
	return []SupportPlanType{SupportPlanTypeAdvanced, SupportPlanTypeEssential, SupportPlanTypeStandard}
}
