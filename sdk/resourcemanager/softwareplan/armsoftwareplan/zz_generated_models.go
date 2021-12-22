//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsoftwareplan

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// Error object returned by the RP
// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw string
	// Error code
	Code *ErrorCode `json:"code,omitempty"`

	// A user readable error message. Localized based on x-ms-effective-locale header in the request
	Message *string `json:"message,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// HybridUseBenefitCreateOptions contains the optional parameters for the HybridUseBenefit.Create method.
type HybridUseBenefitCreateOptions struct {
	// placeholder for future optional parameters
}

// HybridUseBenefitDeleteOptions contains the optional parameters for the HybridUseBenefit.Delete method.
type HybridUseBenefitDeleteOptions struct {
	// placeholder for future optional parameters
}

// HybridUseBenefitGetOptions contains the optional parameters for the HybridUseBenefit.Get method.
type HybridUseBenefitGetOptions struct {
	// placeholder for future optional parameters
}

// HybridUseBenefitListOptions contains the optional parameters for the HybridUseBenefit.List method.
type HybridUseBenefitListOptions struct {
	// Supports applying filter on the type of SKU
	Filter *string
}

// HybridUseBenefitListResult - List of hybrid use benefits
type HybridUseBenefitListResult struct {
	// Url to get the next page of items.
	NextLink *string `json:"nextLink,omitempty"`

	// List of hybrid use benefits
	Value []*HybridUseBenefitModel `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type HybridUseBenefitListResult.
func (h HybridUseBenefitListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", h.NextLink)
	populate(objectMap, "value", h.Value)
	return json.Marshal(objectMap)
}

// HybridUseBenefitModel - Response on GET of a hybrid use benefit
type HybridUseBenefitModel struct {
	Resource
	// REQUIRED; Hybrid use benefit SKU
	SKU *SKU `json:"sku,omitempty"`

	// Property bag for a hybrid use benefit response
	Properties *HybridUseBenefitProperties `json:"properties,omitempty"`

	// READ-ONLY; Indicates the revision of the hybrid use benefit
	Etag *int32 `json:"etag,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type HybridUseBenefitModel.
func (h HybridUseBenefitModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	h.Resource.marshalInternal(objectMap)
	populate(objectMap, "etag", h.Etag)
	populate(objectMap, "properties", h.Properties)
	populate(objectMap, "sku", h.SKU)
	return json.Marshal(objectMap)
}

// HybridUseBenefitProperties - Hybrid use benefit properties
type HybridUseBenefitProperties struct {
	// READ-ONLY; Created date
	CreatedDate *time.Time `json:"createdDate,omitempty" azure:"ro"`

	// READ-ONLY; Last updated date
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type HybridUseBenefitProperties.
func (h HybridUseBenefitProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdDate", h.CreatedDate)
	populateTimeRFC3339(objectMap, "lastUpdatedDate", h.LastUpdatedDate)
	populate(objectMap, "provisioningState", h.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type HybridUseBenefitProperties.
func (h *HybridUseBenefitProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdDate":
			err = unpopulateTimeRFC3339(val, &h.CreatedDate)
			delete(rawMsg, key)
		case "lastUpdatedDate":
			err = unpopulateTimeRFC3339(val, &h.LastUpdatedDate)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &h.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// HybridUseBenefitRevisionListOptions contains the optional parameters for the HybridUseBenefitRevision.List method.
type HybridUseBenefitRevisionListOptions struct {
	// placeholder for future optional parameters
}

// HybridUseBenefitUpdateOptions contains the optional parameters for the HybridUseBenefit.Update method.
type HybridUseBenefitUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationDisplay - Display fields for an operation.
type OperationDisplay struct {
	// Description of the operation
	Description *string `json:"description,omitempty"`

	// Operation to be performed
	Operation *string `json:"operation,omitempty"`

	// Resource Provider name
	Provider *string `json:"provider,omitempty"`

	// Resource that is acted upon
	Resource *string `json:"resource,omitempty"`
}

// OperationList - List all the operations.
type OperationList struct {
	// Url to get the next page of items.
	NextLink *string `json:"nextLink,omitempty"`

	// List of all operations
	Value []*OperationResponse `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationResponse - Operation response.
type OperationResponse struct {
	// Display properties for the operation
	Display *OperationDisplay `json:"display,omitempty"`

	// Name of the operation
	Name *string `json:"name,omitempty"`

	// Origin of the response
	Origin *string `json:"origin,omitempty"`
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// SKU - The SKU to be applied for this resource
type SKU struct {
	// Name of the SKU to be applied
	Name *string `json:"name,omitempty"`
}

// SoftwarePlanRegisterOptions contains the optional parameters for the SoftwarePlan.Register method.
type SoftwarePlanRegisterOptions struct {
	// placeholder for future optional parameters
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
