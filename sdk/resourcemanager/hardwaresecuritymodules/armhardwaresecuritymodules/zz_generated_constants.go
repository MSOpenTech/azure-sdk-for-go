//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

const (
	moduleName    = "armhardwaresecuritymodules"
	moduleVersion = "v0.3.0"
)

// IdentityType - The type of identity.
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "Application"
	IdentityTypeKey             IdentityType = "Key"
	IdentityTypeManagedIdentity IdentityType = "ManagedIdentity"
	IdentityTypeUser            IdentityType = "User"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// JSONWebKeyType - Provisioning state.
type JSONWebKeyType string

const (
	// JSONWebKeyTypeAllocating - A device is currently being allocated for the dedicated HSM resource.
	JSONWebKeyTypeAllocating JSONWebKeyType = "Allocating"
	// JSONWebKeyTypeCheckingQuota - Validating the subscription has sufficient quota to allocate a dedicated HSM device.
	JSONWebKeyTypeCheckingQuota JSONWebKeyType = "CheckingQuota"
	// JSONWebKeyTypeConnecting - The dedicated HSM is being connected to the virtual network.
	JSONWebKeyTypeConnecting JSONWebKeyType = "Connecting"
	// JSONWebKeyTypeDeleting - The dedicated HSM is currently being deleted.
	JSONWebKeyTypeDeleting JSONWebKeyType = "Deleting"
	// JSONWebKeyTypeFailed - Provisioning of the dedicated HSM has failed.
	JSONWebKeyTypeFailed JSONWebKeyType = "Failed"
	// JSONWebKeyTypeProvisioning - The dedicated HSM is currently being provisioned.
	JSONWebKeyTypeProvisioning JSONWebKeyType = "Provisioning"
	// JSONWebKeyTypeSucceeded - The dedicated HSM has been full provisioned.
	JSONWebKeyTypeSucceeded JSONWebKeyType = "Succeeded"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeAllocating,
		JSONWebKeyTypeCheckingQuota,
		JSONWebKeyTypeConnecting,
		JSONWebKeyTypeDeleting,
		JSONWebKeyTypeFailed,
		JSONWebKeyTypeProvisioning,
		JSONWebKeyTypeSucceeded,
	}
}

// SKUName - SKU of the dedicated HSM
type SKUName string

const (
	// SKUNamePayShield10KLMK1CPS250 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 1 local master key which supports up to 250 calls per second.
	SKUNamePayShield10KLMK1CPS250 SKUName = "payShield10K_LMK1_CPS250"
	// SKUNamePayShield10KLMK1CPS2500 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 1 local master key which supports up to 2500 calls per second.
	SKUNamePayShield10KLMK1CPS2500 SKUName = "payShield10K_LMK1_CPS2500"
	// SKUNamePayShield10KLMK1CPS60 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 1 local master key which supports up to 60 calls per second.
	SKUNamePayShield10KLMK1CPS60 SKUName = "payShield10K_LMK1_CPS60"
	// SKUNamePayShield10KLMK2CPS250 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 2 local master keys which supports up to 250 calls per second.
	SKUNamePayShield10KLMK2CPS250 SKUName = "payShield10K_LMK2_CPS250"
	// SKUNamePayShield10KLMK2CPS2500 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 2 local master keys which supports up to 2500 calls per second.
	SKUNamePayShield10KLMK2CPS2500 SKUName = "payShield10K_LMK2_CPS2500"
	// SKUNamePayShield10KLMK2CPS60 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 2 local master keys which supports up to 60 calls per second.
	SKUNamePayShield10KLMK2CPS60 SKUName = "payShield10K_LMK2_CPS60"
	// SKUNameSafeNetLunaNetworkHSMA790 - The dedicated HSM is a Safenet Luna Network HSM A790 device.
	SKUNameSafeNetLunaNetworkHSMA790 SKUName = "SafeNet Luna Network HSM A790"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePayShield10KLMK1CPS250,
		SKUNamePayShield10KLMK1CPS2500,
		SKUNamePayShield10KLMK1CPS60,
		SKUNamePayShield10KLMK2CPS250,
		SKUNamePayShield10KLMK2CPS2500,
		SKUNamePayShield10KLMK2CPS60,
		SKUNameSafeNetLunaNetworkHSMA790,
	}
}
