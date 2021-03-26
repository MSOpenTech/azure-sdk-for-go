package storagecache

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CacheIdentityType enumerates the values for cache identity type.
type CacheIdentityType string

const (
	// None ...
	None CacheIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned CacheIdentityType = "SystemAssigned"
)

// PossibleCacheIdentityTypeValues returns an array of possible values for the CacheIdentityType const type.
func PossibleCacheIdentityTypeValues() []CacheIdentityType {
	return []CacheIdentityType{None, SystemAssigned}
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

// FirmwareStatusType enumerates the values for firmware status type.
type FirmwareStatusType string

const (
	// Available ...
	Available FirmwareStatusType = "available"
	// Unavailable ...
	Unavailable FirmwareStatusType = "unavailable"
)

// PossibleFirmwareStatusTypeValues returns an array of possible values for the FirmwareStatusType const type.
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return []FirmwareStatusType{Available, Unavailable}
}

// HealthStateType enumerates the values for health state type.
type HealthStateType string

const (
	// Degraded ...
	Degraded HealthStateType = "Degraded"
	// Down ...
	Down HealthStateType = "Down"
	// Flushing ...
	Flushing HealthStateType = "Flushing"
	// Healthy ...
	Healthy HealthStateType = "Healthy"
	// Stopped ...
	Stopped HealthStateType = "Stopped"
	// Stopping ...
	Stopping HealthStateType = "Stopping"
	// Transitioning ...
	Transitioning HealthStateType = "Transitioning"
	// Unknown ...
	Unknown HealthStateType = "Unknown"
	// Upgrading ...
	Upgrading HealthStateType = "Upgrading"
)

// PossibleHealthStateTypeValues returns an array of possible values for the HealthStateType const type.
func PossibleHealthStateTypeValues() []HealthStateType {
	return []HealthStateType{Degraded, Down, Flushing, Healthy, Stopped, Stopping, Transitioning, Unknown, Upgrading}
}

// MetricAggregationType enumerates the values for metric aggregation type.
type MetricAggregationType string

const (
	// MetricAggregationTypeAverage ...
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	// MetricAggregationTypeCount ...
	MetricAggregationTypeCount MetricAggregationType = "Count"
	// MetricAggregationTypeMaximum ...
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
	// MetricAggregationTypeMinimum ...
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	// MetricAggregationTypeNone ...
	MetricAggregationTypeNone MetricAggregationType = "None"
	// MetricAggregationTypeNotSpecified ...
	MetricAggregationTypeNotSpecified MetricAggregationType = "NotSpecified"
	// MetricAggregationTypeTotal ...
	MetricAggregationTypeTotal MetricAggregationType = "Total"
)

// PossibleMetricAggregationTypeValues returns an array of possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{MetricAggregationTypeAverage, MetricAggregationTypeCount, MetricAggregationTypeMaximum, MetricAggregationTypeMinimum, MetricAggregationTypeNone, MetricAggregationTypeNotSpecified, MetricAggregationTypeTotal}
}

// ProvisioningStateType enumerates the values for provisioning state type.
type ProvisioningStateType string

const (
	// Cancelled ...
	Cancelled ProvisioningStateType = "Cancelled"
	// Creating ...
	Creating ProvisioningStateType = "Creating"
	// Deleting ...
	Deleting ProvisioningStateType = "Deleting"
	// Failed ...
	Failed ProvisioningStateType = "Failed"
	// Succeeded ...
	Succeeded ProvisioningStateType = "Succeeded"
	// Updating ...
	Updating ProvisioningStateType = "Updating"
)

// PossibleProvisioningStateTypeValues returns an array of possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{Cancelled, Creating, Deleting, Failed, Succeeded, Updating}
}

// ReasonCode enumerates the values for reason code.
type ReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns an array of possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{NotAvailableForSubscription, QuotaID}
}

// TargetType enumerates the values for target type.
type TargetType string

const (
	// TargetTypeClfs ...
	TargetTypeClfs TargetType = "clfs"
	// TargetTypeNfs3 ...
	TargetTypeNfs3 TargetType = "nfs3"
	// TargetTypeStorageTargetProperties ...
	TargetTypeStorageTargetProperties TargetType = "StorageTargetProperties"
	// TargetTypeUnknown ...
	TargetTypeUnknown TargetType = "unknown"
)

// PossibleTargetTypeValues returns an array of possible values for the TargetType const type.
func PossibleTargetTypeValues() []TargetType {
	return []TargetType{TargetTypeClfs, TargetTypeNfs3, TargetTypeStorageTargetProperties, TargetTypeUnknown}
}
