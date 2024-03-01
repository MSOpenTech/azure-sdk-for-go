//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type APIKey.
func (a APIKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "createdAt", a.CreatedAt)
	populate(objectMap, "key", a.Key)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type APIKey.
func (a *APIKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateDateTimeRFC3339(val, "CreatedAt", &a.CreatedAt)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, "Key", &a.Key)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type APIKeys.
func (a APIKeys) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "keys", a.Keys)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type APIKeys.
func (a *APIKeys) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "keys":
			err = unpopulate(val, "Keys", &a.Keys)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CheckNameAvailabilityParameters.
func (c CheckNameAvailabilityParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckNameAvailabilityParameters.
func (c *CheckNameAvailabilityParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CheckNameAvailabilityResult.
func (c CheckNameAvailabilityResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "nameAvailable", c.NameAvailable)
	populate(objectMap, "reason", c.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckNameAvailabilityResult.
func (c *CheckNameAvailabilityResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
			err = unpopulate(val, "Message", &c.Message)
			delete(rawMsg, key)
		case "nameAvailable":
			err = unpopulate(val, "NameAvailable", &c.NameAvailable)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &c.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "info", e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
			err = unpopulate(val, "Info", &e.Info)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
			err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
			delete(rawMsg, key)
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListKeysResult.
func (l ListKeysResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "apiKeyEnabled", l.APIKeyEnabled)
	populate(objectMap, "primaryConnectionString", l.PrimaryConnectionString)
	populate(objectMap, "primaryKey", l.PrimaryKey)
	populate(objectMap, "secondaryConnectionString", l.SecondaryConnectionString)
	populate(objectMap, "secondaryKey", l.SecondaryKey)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListKeysResult.
func (l *ListKeysResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "apiKeyEnabled":
			err = unpopulate(val, "APIKeyEnabled", &l.APIKeyEnabled)
			delete(rawMsg, key)
		case "primaryConnectionString":
			err = unpopulate(val, "PrimaryConnectionString", &l.PrimaryConnectionString)
			delete(rawMsg, key)
		case "primaryKey":
			err = unpopulate(val, "PrimaryKey", &l.PrimaryKey)
			delete(rawMsg, key)
		case "secondaryConnectionString":
			err = unpopulate(val, "SecondaryConnectionString", &l.SecondaryConnectionString)
			delete(rawMsg, key)
		case "secondaryKey":
			err = unpopulate(val, "SecondaryKey", &l.SecondaryKey)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OfferingsListResult.
func (o OfferingsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OfferingsListResult.
func (o *OfferingsListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationsList.
func (o OperationsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationsList.
func (o *OperationsList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PricingDetail.
func (p PricingDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PricingDetail.
func (p *PricingDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &p.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PricingDimension.
func (p PricingDimension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PricingDimension.
func (p *PricingDimension) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Provider.
func (p Provider) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "applicationName", p.ApplicationName)
	populate(objectMap, "instanceUri", p.InstanceURI)
	populate(objectMap, "providerId", p.ProviderID)
	populate(objectMap, "providerSku", p.ProviderSKU)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "resourceUsageId", p.ResourceUsageID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Provider.
func (p *Provider) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "applicationName":
			err = unpopulate(val, "ApplicationName", &p.ApplicationName)
			delete(rawMsg, key)
		case "instanceUri":
			err = unpopulate(val, "InstanceURI", &p.InstanceURI)
			delete(rawMsg, key)
		case "providerId":
			err = unpopulate(val, "ProviderID", &p.ProviderID)
			delete(rawMsg, key)
		case "providerSku":
			err = unpopulate(val, "ProviderSKU", &p.ProviderSKU)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		case "resourceUsageId":
			err = unpopulate(val, "ResourceUsageID", &p.ResourceUsageID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProviderDescription.
func (p ProviderDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProviderDescription.
func (p *ProviderDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProviderProperties.
func (p ProviderProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aad", p.AAD)
	populate(objectMap, "company", p.Company)
	populate(objectMap, "defaultEndpoint", p.DefaultEndpoint)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "managedApplication", p.ManagedApplication)
	populate(objectMap, "pricingDimensions", p.PricingDimensions)
	populate(objectMap, "providerType", p.ProviderType)
	populate(objectMap, "quotaDimensions", p.QuotaDimensions)
	populate(objectMap, "skus", p.SKUs)
	populate(objectMap, "targets", p.Targets)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProviderProperties.
func (p *ProviderProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aad":
			err = unpopulate(val, "AAD", &p.AAD)
			delete(rawMsg, key)
		case "company":
			err = unpopulate(val, "Company", &p.Company)
			delete(rawMsg, key)
		case "defaultEndpoint":
			err = unpopulate(val, "DefaultEndpoint", &p.DefaultEndpoint)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &p.Description)
			delete(rawMsg, key)
		case "managedApplication":
			err = unpopulate(val, "ManagedApplication", &p.ManagedApplication)
			delete(rawMsg, key)
		case "pricingDimensions":
			err = unpopulate(val, "PricingDimensions", &p.PricingDimensions)
			delete(rawMsg, key)
		case "providerType":
			err = unpopulate(val, "ProviderType", &p.ProviderType)
			delete(rawMsg, key)
		case "quotaDimensions":
			err = unpopulate(val, "QuotaDimensions", &p.QuotaDimensions)
			delete(rawMsg, key)
		case "skus":
			err = unpopulate(val, "SKUs", &p.SKUs)
			delete(rawMsg, key)
		case "targets":
			err = unpopulate(val, "Targets", &p.Targets)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProviderPropertiesAAD.
func (p ProviderPropertiesAAD) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "applicationId", p.ApplicationID)
	populate(objectMap, "tenantId", p.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProviderPropertiesAAD.
func (p *ProviderPropertiesAAD) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "applicationId":
			err = unpopulate(val, "ApplicationID", &p.ApplicationID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &p.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProviderPropertiesManagedApplication.
func (p ProviderPropertiesManagedApplication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "offerId", p.OfferID)
	populate(objectMap, "publisherId", p.PublisherID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProviderPropertiesManagedApplication.
func (p *ProviderPropertiesManagedApplication) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "offerId":
			err = unpopulate(val, "OfferID", &p.OfferID)
			delete(rawMsg, key)
		case "publisherId":
			err = unpopulate(val, "PublisherID", &p.PublisherID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QuotaDimension.
func (q QuotaDimension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", q.Description)
	populate(objectMap, "id", q.ID)
	populate(objectMap, "name", q.Name)
	populate(objectMap, "period", q.Period)
	populate(objectMap, "quota", q.Quota)
	populate(objectMap, "scope", q.Scope)
	populate(objectMap, "unit", q.Unit)
	populate(objectMap, "unitPlural", q.UnitPlural)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QuotaDimension.
func (q *QuotaDimension) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", q, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &q.Description)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &q.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &q.Name)
			delete(rawMsg, key)
		case "period":
			err = unpopulate(val, "Period", &q.Period)
			delete(rawMsg, key)
		case "quota":
			err = unpopulate(val, "Quota", &q.Quota)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, "Scope", &q.Scope)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &q.Unit)
			delete(rawMsg, key)
		case "unitPlural":
			err = unpopulate(val, "UnitPlural", &q.UnitPlural)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", q, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Resource.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &r.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUDescription.
func (s SKUDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "autoAdd", s.AutoAdd)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "pricingDetails", s.PricingDetails)
	populate(objectMap, "quotaDimensions", s.QuotaDimensions)
	populate(objectMap, "restrictedAccessUri", s.RestrictedAccessURI)
	populate(objectMap, "targets", s.Targets)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUDescription.
func (s *SKUDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "autoAdd":
			err = unpopulate(val, "AutoAdd", &s.AutoAdd)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &s.Description)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "pricingDetails":
			err = unpopulate(val, "PricingDetails", &s.PricingDetails)
			delete(rawMsg, key)
		case "quotaDimensions":
			err = unpopulate(val, "QuotaDimensions", &s.QuotaDimensions)
			delete(rawMsg, key)
		case "restrictedAccessUri":
			err = unpopulate(val, "RestrictedAccessURI", &s.RestrictedAccessURI)
			delete(rawMsg, key)
		case "targets":
			err = unpopulate(val, "Targets", &s.Targets)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateDateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateDateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TagsObject.
func (t *TagsObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TargetDescription.
func (t TargetDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "acceptedContentEncodings", t.AcceptedContentEncodings)
	populate(objectMap, "acceptedDataFormats", t.AcceptedDataFormats)
	populate(objectMap, "description", t.Description)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "name", t.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TargetDescription.
func (t *TargetDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "acceptedContentEncodings":
			err = unpopulate(val, "AcceptedContentEncodings", &t.AcceptedContentEncodings)
			delete(rawMsg, key)
		case "acceptedDataFormats":
			err = unpopulate(val, "AcceptedDataFormats", &t.AcceptedDataFormats)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &t.Description)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TrackedResource.
func (t *TrackedResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &t.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &t.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &t.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "identity", w.Identity)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Workspace.
func (w *Workspace) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &w.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &w.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &w.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &w.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &w.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &w.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &w.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &w.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceIdentity.
func (w WorkspaceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", w.PrincipalID)
	populate(objectMap, "tenantId", w.TenantID)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceIdentity.
func (w *WorkspaceIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &w.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &w.TenantID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &w.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceListResult.
func (w WorkspaceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", w.NextLink)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceListResult.
func (w *WorkspaceListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &w.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &w.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceResourceProperties.
func (w WorkspaceResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "apiKeyEnabled", w.APIKeyEnabled)
	populate(objectMap, "endpointUri", w.EndpointURI)
	populate(objectMap, "providers", w.Providers)
	populate(objectMap, "provisioningState", w.ProvisioningState)
	populate(objectMap, "storageAccount", w.StorageAccount)
	populate(objectMap, "usable", w.Usable)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceResourceProperties.
func (w *WorkspaceResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "apiKeyEnabled":
			err = unpopulate(val, "APIKeyEnabled", &w.APIKeyEnabled)
			delete(rawMsg, key)
		case "endpointUri":
			err = unpopulate(val, "EndpointURI", &w.EndpointURI)
			delete(rawMsg, key)
		case "providers":
			err = unpopulate(val, "Providers", &w.Providers)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &w.ProvisioningState)
			delete(rawMsg, key)
		case "storageAccount":
			err = unpopulate(val, "StorageAccount", &w.StorageAccount)
			delete(rawMsg, key)
		case "usable":
			err = unpopulate(val, "Usable", &w.Usable)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
