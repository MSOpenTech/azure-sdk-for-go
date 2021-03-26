package search

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2015-02-28/search"

// AdminKeyResult response containing the primary and secondary API keys for a given Azure Search service.
type AdminKeyResult struct {
	autorest.Response `json:"-"`
	// PrimaryKey - READ-ONLY; The primary API key of the Search service.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - READ-ONLY; The secondary API key of the Search service.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// ListQueryKeysResult response containing the query API keys for a given Azure Search service.
type ListQueryKeysResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The query keys for the Azure Search service.
	Value *[]QueryKey `json:"value,omitempty"`
}

// QueryKey describes an API key for a given Azure Search service that has permissions for query operations
// only.
type QueryKey struct {
	// Name - READ-ONLY; The name of the query API key; may be empty.
	Name *string `json:"name,omitempty"`
	// Key - READ-ONLY; The value of the query API key.
	Key *string `json:"key,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ServiceCreateOrUpdateParameters properties that describe an Azure Search service.
type ServiceCreateOrUpdateParameters struct {
	// Location - The geographic location of the Search service.
	Location *string `json:"location,omitempty"`
	// Tags - Tags to help categorize the Search service in the Azure Portal.
	Tags map[string]*string `json:"tags"`
	// Properties - Properties of the Search service.
	Properties *ServiceProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceCreateOrUpdateParameters.
func (scoup ServiceCreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if scoup.Location != nil {
		objectMap["location"] = scoup.Location
	}
	if scoup.Tags != nil {
		objectMap["tags"] = scoup.Tags
	}
	if scoup.Properties != nil {
		objectMap["properties"] = scoup.Properties
	}
	return json.Marshal(objectMap)
}

// ServiceListResult response containing a list of Azure Search services for a given resource group.
type ServiceListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The Search services in the resource group.
	Value *[]ServiceResource `json:"value,omitempty"`
}

// ServiceProperties defines properties of an Azure Search service that can be modified.
type ServiceProperties struct {
	// Sku - The SKU of the Search Service, which determines price tier and capacity limits.
	Sku *Sku `json:"sku,omitempty"`
	// ReplicaCount - The number of replicas in the Search service. If specified, it must be a value between 1 and 6 inclusive.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`
	// PartitionCount - The number of partitions in the Search service; if specified, it can be 1, 2, 3, 4, 6, or 12.
	PartitionCount *int32 `json:"partitionCount,omitempty"`
}

// ServiceReadableProperties defines all the properties of an Azure Search service.
type ServiceReadableProperties struct {
	// Status - READ-ONLY; The status of the Search service. Possible values include: 'ServiceStatusRunning', 'ServiceStatusProvisioning', 'ServiceStatusDeleting', 'ServiceStatusDegraded', 'ServiceStatusDisabled', 'ServiceStatusError'
	Status ServiceStatus `json:"status,omitempty"`
	// StatusDetails - READ-ONLY; The details of the Search service status.
	StatusDetails *string `json:"statusDetails,omitempty"`
	// ProvisioningState - READ-ONLY; The state of the last provisioning operation performed on the Search service. Possible values include: 'Succeeded', 'Provisioning', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Sku - The SKU of the Search Service, which determines price tier and capacity limits.
	Sku *Sku `json:"sku,omitempty"`
	// ReplicaCount - The number of replicas in the Search service. If specified, it must be a value between 1 and 6 inclusive.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`
	// PartitionCount - The number of partitions in the Search service; if specified, it can be 1, 2, 3, 4, 6, or 12.
	PartitionCount *int32 `json:"partitionCount,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceReadableProperties.
func (srp ServiceReadableProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if srp.Sku != nil {
		objectMap["sku"] = srp.Sku
	}
	if srp.ReplicaCount != nil {
		objectMap["replicaCount"] = srp.ReplicaCount
	}
	if srp.PartitionCount != nil {
		objectMap["partitionCount"] = srp.PartitionCount
	}
	return json.Marshal(objectMap)
}

// ServiceResource describes an Azure Search service and its current state.
type ServiceResource struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The resource Id of the Azure Search service.
	ID *string `json:"id,omitempty"`
	// Name - The name of the Search service.
	Name *string `json:"name,omitempty"`
	// Location - The geographic location of the Search service.
	Location *string `json:"location,omitempty"`
	// Tags - Tags to help categorize the Search service in the Azure Portal.
	Tags map[string]*string `json:"tags"`
	// Properties - READ-ONLY; Properties of the Search service.
	Properties *ServiceReadableProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceResource.
func (sr ServiceResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sr.Name != nil {
		objectMap["name"] = sr.Name
	}
	if sr.Location != nil {
		objectMap["location"] = sr.Location
	}
	if sr.Tags != nil {
		objectMap["tags"] = sr.Tags
	}
	return json.Marshal(objectMap)
}

// Sku defines the SKU of an Azure Search Service, which determines price tier and capacity limits.
type Sku struct {
	// Name - The SKU of the Search service. Possible values include: 'Free', 'Standard', 'Standard2'
	Name SkuType `json:"name,omitempty"`
}

// SubResource ...
type SubResource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}
