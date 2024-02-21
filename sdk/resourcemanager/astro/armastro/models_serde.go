//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armastro

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseDataOrganizationProperties.
func (l LiftrBaseDataOrganizationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "marketplace", l.Marketplace)
	populate(objectMap, "partnerOrganizationProperties", l.PartnerOrganizationProperties)
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "user", l.User)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseDataOrganizationProperties.
func (l *LiftrBaseDataOrganizationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "marketplace":
			err = unpopulate(val, "Marketplace", &l.Marketplace)
			delete(rawMsg, key)
		case "partnerOrganizationProperties":
			err = unpopulate(val, "PartnerOrganizationProperties", &l.PartnerOrganizationProperties)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &l.ProvisioningState)
			delete(rawMsg, key)
		case "user":
			err = unpopulate(val, "User", &l.User)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseDataPartnerOrganizationProperties.
func (l LiftrBaseDataPartnerOrganizationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "organizationId", l.OrganizationID)
	populate(objectMap, "organizationName", l.OrganizationName)
	populate(objectMap, "singleSignOnProperties", l.SingleSignOnProperties)
	populate(objectMap, "workspaceId", l.WorkspaceID)
	populate(objectMap, "workspaceName", l.WorkspaceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseDataPartnerOrganizationProperties.
func (l *LiftrBaseDataPartnerOrganizationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "organizationId":
			err = unpopulate(val, "OrganizationID", &l.OrganizationID)
			delete(rawMsg, key)
		case "organizationName":
			err = unpopulate(val, "OrganizationName", &l.OrganizationName)
			delete(rawMsg, key)
		case "singleSignOnProperties":
			err = unpopulate(val, "SingleSignOnProperties", &l.SingleSignOnProperties)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &l.WorkspaceID)
			delete(rawMsg, key)
		case "workspaceName":
			err = unpopulate(val, "WorkspaceName", &l.WorkspaceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseDataPartnerOrganizationPropertiesUpdate.
func (l LiftrBaseDataPartnerOrganizationPropertiesUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "organizationId", l.OrganizationID)
	populate(objectMap, "organizationName", l.OrganizationName)
	populate(objectMap, "singleSignOnProperties", l.SingleSignOnProperties)
	populate(objectMap, "workspaceId", l.WorkspaceID)
	populate(objectMap, "workspaceName", l.WorkspaceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseDataPartnerOrganizationPropertiesUpdate.
func (l *LiftrBaseDataPartnerOrganizationPropertiesUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "organizationId":
			err = unpopulate(val, "OrganizationID", &l.OrganizationID)
			delete(rawMsg, key)
		case "organizationName":
			err = unpopulate(val, "OrganizationName", &l.OrganizationName)
			delete(rawMsg, key)
		case "singleSignOnProperties":
			err = unpopulate(val, "SingleSignOnProperties", &l.SingleSignOnProperties)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &l.WorkspaceID)
			delete(rawMsg, key)
		case "workspaceName":
			err = unpopulate(val, "WorkspaceName", &l.WorkspaceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseMarketplaceDetails.
func (l LiftrBaseMarketplaceDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "offerDetails", l.OfferDetails)
	populate(objectMap, "subscriptionId", l.SubscriptionID)
	populate(objectMap, "subscriptionStatus", l.SubscriptionStatus)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseMarketplaceDetails.
func (l *LiftrBaseMarketplaceDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "offerDetails":
			err = unpopulate(val, "OfferDetails", &l.OfferDetails)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &l.SubscriptionID)
			delete(rawMsg, key)
		case "subscriptionStatus":
			err = unpopulate(val, "SubscriptionStatus", &l.SubscriptionStatus)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseOfferDetails.
func (l LiftrBaseOfferDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "offerId", l.OfferID)
	populate(objectMap, "planId", l.PlanID)
	populate(objectMap, "planName", l.PlanName)
	populate(objectMap, "publisherId", l.PublisherID)
	populate(objectMap, "termId", l.TermID)
	populate(objectMap, "termUnit", l.TermUnit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseOfferDetails.
func (l *LiftrBaseOfferDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "offerId":
			err = unpopulate(val, "OfferID", &l.OfferID)
			delete(rawMsg, key)
		case "planId":
			err = unpopulate(val, "PlanID", &l.PlanID)
			delete(rawMsg, key)
		case "planName":
			err = unpopulate(val, "PlanName", &l.PlanName)
			delete(rawMsg, key)
		case "publisherId":
			err = unpopulate(val, "PublisherID", &l.PublisherID)
			delete(rawMsg, key)
		case "termId":
			err = unpopulate(val, "TermID", &l.TermID)
			delete(rawMsg, key)
		case "termUnit":
			err = unpopulate(val, "TermUnit", &l.TermUnit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseSingleSignOnProperties.
func (l LiftrBaseSingleSignOnProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aadDomains", l.AADDomains)
	populate(objectMap, "enterpriseAppId", l.EnterpriseAppID)
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "singleSignOnState", l.SingleSignOnState)
	populate(objectMap, "singleSignOnUrl", l.SingleSignOnURL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseSingleSignOnProperties.
func (l *LiftrBaseSingleSignOnProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aadDomains":
			err = unpopulate(val, "AADDomains", &l.AADDomains)
			delete(rawMsg, key)
		case "enterpriseAppId":
			err = unpopulate(val, "EnterpriseAppID", &l.EnterpriseAppID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &l.ProvisioningState)
			delete(rawMsg, key)
		case "singleSignOnState":
			err = unpopulate(val, "SingleSignOnState", &l.SingleSignOnState)
			delete(rawMsg, key)
		case "singleSignOnUrl":
			err = unpopulate(val, "SingleSignOnURL", &l.SingleSignOnURL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseUserDetails.
func (l LiftrBaseUserDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "emailAddress", l.EmailAddress)
	populate(objectMap, "firstName", l.FirstName)
	populate(objectMap, "lastName", l.LastName)
	populate(objectMap, "phoneNumber", l.PhoneNumber)
	populate(objectMap, "upn", l.Upn)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseUserDetails.
func (l *LiftrBaseUserDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "emailAddress":
			err = unpopulate(val, "EmailAddress", &l.EmailAddress)
			delete(rawMsg, key)
		case "firstName":
			err = unpopulate(val, "FirstName", &l.FirstName)
			delete(rawMsg, key)
		case "lastName":
			err = unpopulate(val, "LastName", &l.LastName)
			delete(rawMsg, key)
		case "phoneNumber":
			err = unpopulate(val, "PhoneNumber", &l.PhoneNumber)
			delete(rawMsg, key)
		case "upn":
			err = unpopulate(val, "Upn", &l.Upn)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LiftrBaseUserDetailsUpdate.
func (l LiftrBaseUserDetailsUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "emailAddress", l.EmailAddress)
	populate(objectMap, "firstName", l.FirstName)
	populate(objectMap, "lastName", l.LastName)
	populate(objectMap, "phoneNumber", l.PhoneNumber)
	populate(objectMap, "upn", l.Upn)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LiftrBaseUserDetailsUpdate.
func (l *LiftrBaseUserDetailsUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "emailAddress":
			err = unpopulate(val, "EmailAddress", &l.EmailAddress)
			delete(rawMsg, key)
		case "firstName":
			err = unpopulate(val, "FirstName", &l.FirstName)
			delete(rawMsg, key)
		case "lastName":
			err = unpopulate(val, "LastName", &l.LastName)
			delete(rawMsg, key)
		case "phoneNumber":
			err = unpopulate(val, "PhoneNumber", &l.PhoneNumber)
			delete(rawMsg, key)
		case "upn":
			err = unpopulate(val, "Upn", &l.Upn)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagedServiceIdentity.
func (m ManagedServiceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", m.PrincipalID)
	populate(objectMap, "tenantId", m.TenantID)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagedServiceIdentity.
func (m *ManagedServiceIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &m.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &m.TenantID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		case "userAssignedIdentities":
			err = unpopulate(val, "UserAssignedIdentities", &m.UserAssignedIdentities)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actionType", o.ActionType)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
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
		case "actionType":
			err = unpopulate(val, "ActionType", &o.ActionType)
			delete(rawMsg, key)
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &o.Origin)
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

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type OrganizationResource.
func (o OrganizationResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "identity", o.Identity)
	populate(objectMap, "location", o.Location)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "systemData", o.SystemData)
	populate(objectMap, "tags", o.Tags)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationResource.
func (o *OrganizationResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &o.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &o.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &o.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &o.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &o.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &o.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OrganizationResourceListResult.
func (o OrganizationResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationResourceListResult.
func (o *OrganizationResourceListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type OrganizationResourceUpdate.
func (o OrganizationResourceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "identity", o.Identity)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "tags", o.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationResourceUpdate.
func (o *OrganizationResourceUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "identity":
			err = unpopulate(val, "Identity", &o.Identity)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &o.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &o.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OrganizationResourceUpdateProperties.
func (o OrganizationResourceUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "partnerOrganizationProperties", o.PartnerOrganizationProperties)
	populate(objectMap, "user", o.User)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationResourceUpdateProperties.
func (o *OrganizationResourceUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "partnerOrganizationProperties":
			err = unpopulate(val, "PartnerOrganizationProperties", &o.PartnerOrganizationProperties)
			delete(rawMsg, key)
		case "user":
			err = unpopulate(val, "User", &o.User)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
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

// MarshalJSON implements the json.Marshaller interface for type UserAssignedIdentity.
func (u UserAssignedIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "clientId", u.ClientID)
	populate(objectMap, "principalId", u.PrincipalID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserAssignedIdentity.
func (u *UserAssignedIdentity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clientId":
			err = unpopulate(val, "ClientID", &u.ClientID)
			delete(rawMsg, key)
		case "principalId":
			err = unpopulate(val, "PrincipalID", &u.PrincipalID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
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

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
