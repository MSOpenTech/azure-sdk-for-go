package policy

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AliasPathAttributes enumerates the values for alias path attributes.
type AliasPathAttributes string

const (
	// Modifiable The token that the alias path is referring to is modifiable by policies with 'modify' effect.
	Modifiable AliasPathAttributes = "Modifiable"
	// None The token that the alias path is referring to has no attributes.
	None AliasPathAttributes = "None"
)

// PossibleAliasPathAttributesValues returns an array of possible values for the AliasPathAttributes const type.
func PossibleAliasPathAttributesValues() []AliasPathAttributes {
	return []AliasPathAttributes{Modifiable, None}
}

// AliasPathTokenType enumerates the values for alias path token type.
type AliasPathTokenType string

const (
	// Any The token type can be anything.
	Any AliasPathTokenType = "Any"
	// Array The token type is array.
	Array AliasPathTokenType = "Array"
	// Boolean The token type is boolean.
	Boolean AliasPathTokenType = "Boolean"
	// Integer The token type is integer.
	Integer AliasPathTokenType = "Integer"
	// NotSpecified The token type is not specified.
	NotSpecified AliasPathTokenType = "NotSpecified"
	// Number The token type is number.
	Number AliasPathTokenType = "Number"
	// Object The token type is object.
	Object AliasPathTokenType = "Object"
	// String The token type is string.
	String AliasPathTokenType = "String"
)

// PossibleAliasPathTokenTypeValues returns an array of possible values for the AliasPathTokenType const type.
func PossibleAliasPathTokenTypeValues() []AliasPathTokenType {
	return []AliasPathTokenType{Any, Array, Boolean, Integer, NotSpecified, Number, Object, String}
}

// AliasPatternType enumerates the values for alias pattern type.
type AliasPatternType string

const (
	// AliasPatternTypeExtract Extract is the only allowed value.
	AliasPatternTypeExtract AliasPatternType = "Extract"
	// AliasPatternTypeNotSpecified NotSpecified is not allowed.
	AliasPatternTypeNotSpecified AliasPatternType = "NotSpecified"
)

// PossibleAliasPatternTypeValues returns an array of possible values for the AliasPatternType const type.
func PossibleAliasPatternTypeValues() []AliasPatternType {
	return []AliasPatternType{AliasPatternTypeExtract, AliasPatternTypeNotSpecified}
}

// AliasType enumerates the values for alias type.
type AliasType string

const (
	// AliasTypeMask Alias value is secret.
	AliasTypeMask AliasType = "Mask"
	// AliasTypeNotSpecified Alias type is unknown (same as not providing alias type).
	AliasTypeNotSpecified AliasType = "NotSpecified"
	// AliasTypePlainText Alias value is not secret.
	AliasTypePlainText AliasType = "PlainText"
)

// PossibleAliasTypeValues returns an array of possible values for the AliasType const type.
func PossibleAliasTypeValues() []AliasType {
	return []AliasType{AliasTypeMask, AliasTypeNotSpecified, AliasTypePlainText}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// EnforcementMode enumerates the values for enforcement mode.
type EnforcementMode string

const (
	// Default The policy effect is enforced during resource creation or update.
	Default EnforcementMode = "Default"
	// DoNotEnforce The policy effect is not enforced during resource creation or update.
	DoNotEnforce EnforcementMode = "DoNotEnforce"
)

// PossibleEnforcementModeValues returns an array of possible values for the EnforcementMode const type.
func PossibleEnforcementModeValues() []EnforcementMode {
	return []EnforcementMode{Default, DoNotEnforce}
}

// ExemptionCategory enumerates the values for exemption category.
type ExemptionCategory string

const (
	// Mitigated This category of exemptions usually means the mitigation actions have been applied to the
	// scope.
	Mitigated ExemptionCategory = "Mitigated"
	// Waiver This category of exemptions usually means the scope is not applicable for the policy.
	Waiver ExemptionCategory = "Waiver"
)

// PossibleExemptionCategoryValues returns an array of possible values for the ExemptionCategory const type.
func PossibleExemptionCategoryValues() []ExemptionCategory {
	return []ExemptionCategory{Mitigated, Waiver}
}

// ParameterType enumerates the values for parameter type.
type ParameterType string

const (
	// ParameterTypeArray ...
	ParameterTypeArray ParameterType = "Array"
	// ParameterTypeBoolean ...
	ParameterTypeBoolean ParameterType = "Boolean"
	// ParameterTypeDateTime ...
	ParameterTypeDateTime ParameterType = "DateTime"
	// ParameterTypeFloat ...
	ParameterTypeFloat ParameterType = "Float"
	// ParameterTypeInteger ...
	ParameterTypeInteger ParameterType = "Integer"
	// ParameterTypeObject ...
	ParameterTypeObject ParameterType = "Object"
	// ParameterTypeString ...
	ParameterTypeString ParameterType = "String"
)

// PossibleParameterTypeValues returns an array of possible values for the ParameterType const type.
func PossibleParameterTypeValues() []ParameterType {
	return []ParameterType{ParameterTypeArray, ParameterTypeBoolean, ParameterTypeDateTime, ParameterTypeFloat, ParameterTypeInteger, ParameterTypeObject, ParameterTypeString}
}

// PricingTier enumerates the values for pricing tier.
type PricingTier string

const (
	// Advanced The pricing tier gives the user ability to use policy exemption feature.
	Advanced PricingTier = "Advanced"
	// Defender The pricing tier gives the user ability to use policy exemption feature. This pricing tier is
	// managed by Azure Security Center.
	Defender PricingTier = "Defender"
)

// PossiblePricingTierValues returns an array of possible values for the PricingTier const type.
func PossiblePricingTierValues() []PricingTier {
	return []PricingTier{Advanced, Defender}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted When creating policy pricing at the management group level, the provisioning state will be
	// accepted initially.
	Accepted ProvisioningState = "Accepted"
	// Deleting When deleting policy pricing at the management group level, the provisioning state will be
	// deleting.
	Deleting ProvisioningState = "Deleting"
	// Succeeded The policy pricing is provisioned.
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Deleting, Succeeded}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone Indicates that no identity is associated with the resource or that the existing
	// identity should be removed.
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned Indicates that a system assigned identity is associated with the
	// resource.
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeUserAssigned Indicates that a system assigned identity is associated with the
	// resource.
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeUserAssigned}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeBuiltIn ...
	TypeBuiltIn Type = "BuiltIn"
	// TypeCustom ...
	TypeCustom Type = "Custom"
	// TypeNotSpecified ...
	TypeNotSpecified Type = "NotSpecified"
	// TypeStatic ...
	TypeStatic Type = "Static"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeBuiltIn, TypeCustom, TypeNotSpecified, TypeStatic}
}
