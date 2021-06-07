package maps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/maps/mgmt/2017-01-01-preview/maps"

// Account an Azure resource which represents access to a suite of Maps REST APIs.
type Account struct {
	autorest.Response `json:"-"`
	// Location - READ-ONLY; The location of the resource.
	Location *string `json:"location,omitempty"`
	// Tags - READ-ONLY; Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Sku - READ-ONLY; The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
	// Properties - READ-ONLY; The map account properties.
	Properties *AccountProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The fully qualified Maps Account resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the Maps Account, which is unique within a Resource Group.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Azure resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountCreateParameters parameters used to create a new Maps Account.
type AccountCreateParameters struct {
	// Location - The location of the resource.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	if acp.Sku != nil {
		objectMap["sku"] = acp.Sku
	}
	return json.Marshal(objectMap)
}

// AccountKeys the set of keys which can be used to access the Maps REST APIs. Two keys are provided for
// key rotation without interruption.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The full Azure resource identifier of the Maps Account.
	ID *string `json:"id,omitempty"`
	// PrimaryKey - READ-ONLY; The primary key for accessing the Maps REST APIs.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - READ-ONLY; The secondary key for accessing the Maps REST APIs.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountKeys.
func (ak AccountKeys) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountProperties additional Map account properties
type AccountProperties struct {
	// XMsClientID - A unique identifier for the maps account
	XMsClientID *string `json:"x-ms-client-id,omitempty"`
}

// Accounts a list of Maps Accounts.
type Accounts struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; a Maps Account.
	Value *[]Account `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for Accounts.
func (a Accounts) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountsMoveRequest the description of what resources to move between resource groups.
type AccountsMoveRequest struct {
	// TargetResourceGroup - The name of the destination resource group.
	TargetResourceGroup *string `json:"targetResourceGroup,omitempty"`
	// ResourceIds - A list of resource names to move from the source resource group.
	ResourceIds *[]string `json:"resourceIds,omitempty"`
}

// AccountUpdateParameters parameters used to update an existing Maps Account.
type AccountUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	if aup.Sku != nil {
		objectMap["sku"] = aup.Sku
	}
	return json.Marshal(objectMap)
}

// Error this object is returned when an error occurs in the Maps API
type Error struct {
	// Code - READ-ONLY; Error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; If available, a human readable description of the error.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; If available, the component generating the error.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; If available, a list of additional details about the error.
	Details *[]ErrorDetailsItem `json:"details,omitempty"`
}

// MarshalJSON is the custom marshaler for Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorDetailsItem ...
type ErrorDetailsItem struct {
	// Code - READ-ONLY; Error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; If available, a human readable description of the error.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; If available, the component generating the error.
	Target *string `json:"target,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorDetailsItem.
func (eI ErrorDetailsItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// KeySpecification whether the operation refers to the primary or secondary key.
type KeySpecification struct {
	// KeyType - Whether the operation refers to the primary or secondary key. Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// Operations the set of operations available for Maps.
type Operations struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; An operation available for Maps.
	Value *[]OperationsValueItem `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for Operations.
func (o Operations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationsValueItem ...
type OperationsValueItem struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The human-readable description of the operation.
	Display *OperationsValueItemDisplay `json:"display,omitempty"`
	// Origin - READ-ONLY; The origin of the operation.
	Origin *string `json:"origin,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationsValueItem.
func (oI OperationsValueItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if oI.Display != nil {
		objectMap["display"] = oI.Display
	}
	return json.Marshal(objectMap)
}

// OperationsValueItemDisplay the human-readable description of the operation.
type OperationsValueItemDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft Maps.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; The description of the operation.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationsValueItemDisplay.
func (oI OperationsValueItemDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Resource an Azure resource
type Resource struct {
	// ID - READ-ONLY; The fully qualified Maps Account resource identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the Maps Account, which is unique within a Resource Group.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Azure resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Sku the SKU of the Maps Account.
type Sku struct {
	// Name - The name of the SKU, in standard format (such as S0).
	Name *string `json:"name,omitempty"`
	// Tier - READ-ONLY; Gets the sku tier. This is based on the SKU name.
	Tier *string `json:"tier,omitempty"`
}

// MarshalJSON is the custom marshaler for Sku.
func (s Sku) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.Name != nil {
		objectMap["name"] = s.Name
	}
	return json.Marshal(objectMap)
}
