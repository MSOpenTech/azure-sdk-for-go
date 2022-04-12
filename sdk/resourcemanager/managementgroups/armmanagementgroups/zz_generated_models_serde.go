//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CreateManagementGroupChildInfo.
func (c CreateManagementGroupChildInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "children", c.Children)
	populate(objectMap, "displayName", c.DisplayName)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CreateManagementGroupDetails.
func (c CreateManagementGroupDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "parent", c.Parent)
	populate(objectMap, "updatedBy", c.UpdatedBy)
	populateTimeRFC3339(objectMap, "updatedTime", c.UpdatedTime)
	populate(objectMap, "version", c.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateManagementGroupDetails.
func (c *CreateManagementGroupDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "parent":
			err = unpopulate(val, &c.Parent)
			delete(rawMsg, key)
		case "updatedBy":
			err = unpopulate(val, &c.UpdatedBy)
			delete(rawMsg, key)
		case "updatedTime":
			err = unpopulateTimeRFC3339(val, &c.UpdatedTime)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &c.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CreateManagementGroupProperties.
func (c CreateManagementGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "children", c.Children)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "displayName", c.DisplayName)
	populate(objectMap, "tenantId", c.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CreateOrUpdateSettingsRequest.
func (c CreateOrUpdateSettingsRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", c.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DescendantListResult.
func (d DescendantListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EntityHierarchyItemProperties.
func (e EntityHierarchyItemProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "children", e.Children)
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "permissions", e.Permissions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EntityInfoProperties.
func (e EntityInfoProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "inheritedPermissions", e.InheritedPermissions)
	populate(objectMap, "numberOfChildGroups", e.NumberOfChildGroups)
	populate(objectMap, "numberOfChildren", e.NumberOfChildren)
	populate(objectMap, "numberOfDescendants", e.NumberOfDescendants)
	populate(objectMap, "parent", e.Parent)
	populate(objectMap, "parentDisplayNameChain", e.ParentDisplayNameChain)
	populate(objectMap, "parentNameChain", e.ParentNameChain)
	populate(objectMap, "permissions", e.Permissions)
	populate(objectMap, "tenantId", e.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EntityListResult.
func (e EntityListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", e.Count)
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HierarchySettingsList.
func (h HierarchySettingsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", h.NextLink)
	populate(objectMap, "value", h.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ListSubscriptionUnderManagementGroup.
func (l ListSubscriptionUnderManagementGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagementGroupChildInfo.
func (m ManagementGroupChildInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "children", m.Children)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagementGroupDetails.
func (m ManagementGroupDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managementGroupAncestors", m.ManagementGroupAncestors)
	populate(objectMap, "managementGroupAncestorsChain", m.ManagementGroupAncestorsChain)
	populate(objectMap, "parent", m.Parent)
	populate(objectMap, "path", m.Path)
	populate(objectMap, "updatedBy", m.UpdatedBy)
	populateTimeRFC3339(objectMap, "updatedTime", m.UpdatedTime)
	populate(objectMap, "version", m.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManagementGroupDetails.
func (m *ManagementGroupDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "managementGroupAncestors":
			err = unpopulate(val, &m.ManagementGroupAncestors)
			delete(rawMsg, key)
		case "managementGroupAncestorsChain":
			err = unpopulate(val, &m.ManagementGroupAncestorsChain)
			delete(rawMsg, key)
		case "parent":
			err = unpopulate(val, &m.Parent)
			delete(rawMsg, key)
		case "path":
			err = unpopulate(val, &m.Path)
			delete(rawMsg, key)
		case "updatedBy":
			err = unpopulate(val, &m.UpdatedBy)
			delete(rawMsg, key)
		case "updatedTime":
			err = unpopulateTimeRFC3339(val, &m.UpdatedTime)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &m.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManagementGroupListResult.
func (m ManagementGroupListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagementGroupProperties.
func (m ManagementGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "children", m.Children)
	populate(objectMap, "details", m.Details)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "tenantId", m.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PatchManagementGroupRequest.
func (p PatchManagementGroupRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "parentGroupId", p.ParentGroupID)
	return json.Marshal(objectMap)
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
