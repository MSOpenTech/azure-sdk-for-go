//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

const (
	moduleName    = "armmariadb"
	moduleVersion = "v0.3.0"
)

// CreateMode - The mode to create a new server.
type CreateMode string

const (
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeGeoRestore,
		CreateModePointInTimeRestore,
		CreateModeReplica,
	}
}

// GeoRedundantBackup - Enable Geo-redundant or not for server backup.
type GeoRedundantBackup string

const (
	GeoRedundantBackupDisabled GeoRedundantBackup = "Disabled"
	GeoRedundantBackupEnabled  GeoRedundantBackup = "Enabled"
)

// PossibleGeoRedundantBackupValues returns the possible values for the GeoRedundantBackup const type.
func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return []GeoRedundantBackup{
		GeoRedundantBackupDisabled,
		GeoRedundantBackupEnabled,
	}
}

// MinimalTLSVersionEnum - Enforce a minimal Tls version for the server.
type MinimalTLSVersionEnum string

const (
	MinimalTLSVersionEnumTLS10                  MinimalTLSVersionEnum = "TLS1_0"
	MinimalTLSVersionEnumTLS11                  MinimalTLSVersionEnum = "TLS1_1"
	MinimalTLSVersionEnumTLS12                  MinimalTLSVersionEnum = "TLS1_2"
	MinimalTLSVersionEnumTLSEnforcementDisabled MinimalTLSVersionEnum = "TLSEnforcementDisabled"
)

// PossibleMinimalTLSVersionEnumValues returns the possible values for the MinimalTLSVersionEnum const type.
func PossibleMinimalTLSVersionEnumValues() []MinimalTLSVersionEnum {
	return []MinimalTLSVersionEnum{
		MinimalTLSVersionEnumTLS10,
		MinimalTLSVersionEnumTLS11,
		MinimalTLSVersionEnumTLS12,
		MinimalTLSVersionEnumTLSEnforcementDisabled,
	}
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginNotSpecified OperationOrigin = "NotSpecified"
	OperationOriginSystem       OperationOrigin = "system"
	OperationOriginUser         OperationOrigin = "user"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginNotSpecified,
		OperationOriginSystem,
		OperationOriginUser,
	}
}

// PrivateEndpointProvisioningState - State of the private endpoint connection.
type PrivateEndpointProvisioningState string

const (
	PrivateEndpointProvisioningStateApproving PrivateEndpointProvisioningState = "Approving"
	PrivateEndpointProvisioningStateDropping  PrivateEndpointProvisioningState = "Dropping"
	PrivateEndpointProvisioningStateFailed    PrivateEndpointProvisioningState = "Failed"
	PrivateEndpointProvisioningStateReady     PrivateEndpointProvisioningState = "Ready"
	PrivateEndpointProvisioningStateRejecting PrivateEndpointProvisioningState = "Rejecting"
)

// PossiblePrivateEndpointProvisioningStateValues returns the possible values for the PrivateEndpointProvisioningState const type.
func PossiblePrivateEndpointProvisioningStateValues() []PrivateEndpointProvisioningState {
	return []PrivateEndpointProvisioningState{
		PrivateEndpointProvisioningStateApproving,
		PrivateEndpointProvisioningStateDropping,
		PrivateEndpointProvisioningStateFailed,
		PrivateEndpointProvisioningStateReady,
		PrivateEndpointProvisioningStateRejecting,
	}
}

// PrivateLinkServiceConnectionStateActionsRequire - The actions required for private link service connection.
type PrivateLinkServiceConnectionStateActionsRequire string

const (
	PrivateLinkServiceConnectionStateActionsRequireNone PrivateLinkServiceConnectionStateActionsRequire = "None"
)

// PossiblePrivateLinkServiceConnectionStateActionsRequireValues returns the possible values for the PrivateLinkServiceConnectionStateActionsRequire const type.
func PossiblePrivateLinkServiceConnectionStateActionsRequireValues() []PrivateLinkServiceConnectionStateActionsRequire {
	return []PrivateLinkServiceConnectionStateActionsRequire{
		PrivateLinkServiceConnectionStateActionsRequireNone,
	}
}

// PrivateLinkServiceConnectionStateStatus - The private link service connection status.
type PrivateLinkServiceConnectionStateStatus string

const (
	PrivateLinkServiceConnectionStateStatusApproved     PrivateLinkServiceConnectionStateStatus = "Approved"
	PrivateLinkServiceConnectionStateStatusDisconnected PrivateLinkServiceConnectionStateStatus = "Disconnected"
	PrivateLinkServiceConnectionStateStatusPending      PrivateLinkServiceConnectionStateStatus = "Pending"
	PrivateLinkServiceConnectionStateStatusRejected     PrivateLinkServiceConnectionStateStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStateStatusValues returns the possible values for the PrivateLinkServiceConnectionStateStatus const type.
func PossiblePrivateLinkServiceConnectionStateStatusValues() []PrivateLinkServiceConnectionStateStatus {
	return []PrivateLinkServiceConnectionStateStatus{
		PrivateLinkServiceConnectionStateStatusApproved,
		PrivateLinkServiceConnectionStateStatusDisconnected,
		PrivateLinkServiceConnectionStateStatusPending,
		PrivateLinkServiceConnectionStateStatusRejected,
	}
}

// PublicNetworkAccessEnum - Whether or not public network access is allowed for this server. Value is optional but if passed
// in, must be 'Enabled' or 'Disabled'
type PublicNetworkAccessEnum string

const (
	PublicNetworkAccessEnumDisabled PublicNetworkAccessEnum = "Disabled"
	PublicNetworkAccessEnumEnabled  PublicNetworkAccessEnum = "Enabled"
)

// PossiblePublicNetworkAccessEnumValues returns the possible values for the PublicNetworkAccessEnum const type.
func PossiblePublicNetworkAccessEnumValues() []PublicNetworkAccessEnum {
	return []PublicNetworkAccessEnum{
		PublicNetworkAccessEnumDisabled,
		PublicNetworkAccessEnumEnabled,
	}
}

// QueryPerformanceInsightResetDataResultState - Indicates result of the operation.
type QueryPerformanceInsightResetDataResultState string

const (
	QueryPerformanceInsightResetDataResultStateFailed    QueryPerformanceInsightResetDataResultState = "Failed"
	QueryPerformanceInsightResetDataResultStateSucceeded QueryPerformanceInsightResetDataResultState = "Succeeded"
)

// PossibleQueryPerformanceInsightResetDataResultStateValues returns the possible values for the QueryPerformanceInsightResetDataResultState const type.
func PossibleQueryPerformanceInsightResetDataResultStateValues() []QueryPerformanceInsightResetDataResultState {
	return []QueryPerformanceInsightResetDataResultState{
		QueryPerformanceInsightResetDataResultStateFailed,
		QueryPerformanceInsightResetDataResultStateSucceeded,
	}
}

// SKUTier - The tier of the particular SKU, e.g. Basic.
type SKUTier string

const (
	SKUTierBasic           SKUTier = "Basic"
	SKUTierGeneralPurpose  SKUTier = "GeneralPurpose"
	SKUTierMemoryOptimized SKUTier = "MemoryOptimized"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierGeneralPurpose,
		SKUTierMemoryOptimized,
	}
}

// SSLEnforcementEnum - Enable ssl enforcement or not when connect to server.
type SSLEnforcementEnum string

const (
	SSLEnforcementEnumEnabled  SSLEnforcementEnum = "Enabled"
	SSLEnforcementEnumDisabled SSLEnforcementEnum = "Disabled"
)

// PossibleSSLEnforcementEnumValues returns the possible values for the SSLEnforcementEnum const type.
func PossibleSSLEnforcementEnumValues() []SSLEnforcementEnum {
	return []SSLEnforcementEnum{
		SSLEnforcementEnumEnabled,
		SSLEnforcementEnumDisabled,
	}
}

type SecurityAlertPolicyName string

const (
	SecurityAlertPolicyNameDefault SecurityAlertPolicyName = "Default"
)

// PossibleSecurityAlertPolicyNameValues returns the possible values for the SecurityAlertPolicyName const type.
func PossibleSecurityAlertPolicyNameValues() []SecurityAlertPolicyName {
	return []SecurityAlertPolicyName{
		SecurityAlertPolicyNameDefault,
	}
}

// ServerSecurityAlertPolicyState - Specifies the state of the policy, whether it is enabled or disabled.
type ServerSecurityAlertPolicyState string

const (
	ServerSecurityAlertPolicyStateEnabled  ServerSecurityAlertPolicyState = "Enabled"
	ServerSecurityAlertPolicyStateDisabled ServerSecurityAlertPolicyState = "Disabled"
)

// PossibleServerSecurityAlertPolicyStateValues returns the possible values for the ServerSecurityAlertPolicyState const type.
func PossibleServerSecurityAlertPolicyStateValues() []ServerSecurityAlertPolicyState {
	return []ServerSecurityAlertPolicyState{
		ServerSecurityAlertPolicyStateEnabled,
		ServerSecurityAlertPolicyStateDisabled,
	}
}

// ServerState - A state of a server that is visible to user.
type ServerState string

const (
	ServerStateDisabled ServerState = "Disabled"
	ServerStateDropping ServerState = "Dropping"
	ServerStateReady    ServerState = "Ready"
)

// PossibleServerStateValues returns the possible values for the ServerState const type.
func PossibleServerStateValues() []ServerState {
	return []ServerState{
		ServerStateDisabled,
		ServerStateDropping,
		ServerStateReady,
	}
}

// ServerVersion - The version of a server.
type ServerVersion string

const (
	ServerVersionTen2 ServerVersion = "10.2"
	ServerVersionTen3 ServerVersion = "10.3"
)

// PossibleServerVersionValues returns the possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{
		ServerVersionTen2,
		ServerVersionTen3,
	}
}

// StorageAutogrow - Enable Storage Auto Grow.
type StorageAutogrow string

const (
	StorageAutogrowDisabled StorageAutogrow = "Disabled"
	StorageAutogrowEnabled  StorageAutogrow = "Enabled"
)

// PossibleStorageAutogrowValues returns the possible values for the StorageAutogrow const type.
func PossibleStorageAutogrowValues() []StorageAutogrow {
	return []StorageAutogrow{
		StorageAutogrowDisabled,
		StorageAutogrowEnabled,
	}
}

// VirtualNetworkRuleState - Virtual Network Rule State
type VirtualNetworkRuleState string

const (
	VirtualNetworkRuleStateDeleting     VirtualNetworkRuleState = "Deleting"
	VirtualNetworkRuleStateInProgress   VirtualNetworkRuleState = "InProgress"
	VirtualNetworkRuleStateInitializing VirtualNetworkRuleState = "Initializing"
	VirtualNetworkRuleStateReady        VirtualNetworkRuleState = "Ready"
	VirtualNetworkRuleStateUnknown      VirtualNetworkRuleState = "Unknown"
)

// PossibleVirtualNetworkRuleStateValues returns the possible values for the VirtualNetworkRuleState const type.
func PossibleVirtualNetworkRuleStateValues() []VirtualNetworkRuleState {
	return []VirtualNetworkRuleState{
		VirtualNetworkRuleStateDeleting,
		VirtualNetworkRuleStateInProgress,
		VirtualNetworkRuleStateInitializing,
		VirtualNetworkRuleStateReady,
		VirtualNetworkRuleStateUnknown,
	}
}
