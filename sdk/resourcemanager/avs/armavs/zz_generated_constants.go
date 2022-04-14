//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

const (
	moduleName    = "armavs"
	moduleVersion = "v0.3.0"
)

// AddonProvisioningState - The state of the addon provisioning
type AddonProvisioningState string

const (
	AddonProvisioningStateBuilding  AddonProvisioningState = "Building"
	AddonProvisioningStateCancelled AddonProvisioningState = "Cancelled"
	AddonProvisioningStateDeleting  AddonProvisioningState = "Deleting"
	AddonProvisioningStateFailed    AddonProvisioningState = "Failed"
	AddonProvisioningStateSucceeded AddonProvisioningState = "Succeeded"
	AddonProvisioningStateUpdating  AddonProvisioningState = "Updating"
)

// PossibleAddonProvisioningStateValues returns the possible values for the AddonProvisioningState const type.
func PossibleAddonProvisioningStateValues() []AddonProvisioningState {
	return []AddonProvisioningState{
		AddonProvisioningStateBuilding,
		AddonProvisioningStateCancelled,
		AddonProvisioningStateDeleting,
		AddonProvisioningStateFailed,
		AddonProvisioningStateSucceeded,
		AddonProvisioningStateUpdating,
	}
}

// AddonType - The type of private cloud addon
type AddonType string

const (
	AddonTypeHCX AddonType = "HCX"
	AddonTypeSRM AddonType = "SRM"
	AddonTypeVR  AddonType = "VR"
)

// PossibleAddonTypeValues returns the possible values for the AddonType const type.
func PossibleAddonTypeValues() []AddonType {
	return []AddonType{
		AddonTypeHCX,
		AddonTypeSRM,
		AddonTypeVR,
	}
}

// AffinityType - Placement policy affinity type
type AffinityType string

const (
	AffinityTypeAffinity     AffinityType = "Affinity"
	AffinityTypeAntiAffinity AffinityType = "AntiAffinity"
)

// PossibleAffinityTypeValues returns the possible values for the AffinityType const type.
func PossibleAffinityTypeValues() []AffinityType {
	return []AffinityType{
		AffinityTypeAffinity,
		AffinityTypeAntiAffinity,
	}
}

// AvailabilityStrategy - The availability strategy for the private cloud
type AvailabilityStrategy string

const (
	AvailabilityStrategyDualZone   AvailabilityStrategy = "DualZone"
	AvailabilityStrategySingleZone AvailabilityStrategy = "SingleZone"
)

// PossibleAvailabilityStrategyValues returns the possible values for the AvailabilityStrategy const type.
func PossibleAvailabilityStrategyValues() []AvailabilityStrategy {
	return []AvailabilityStrategy{
		AvailabilityStrategyDualZone,
		AvailabilityStrategySingleZone,
	}
}

// CloudLinkStatus - The state of the cloud link.
type CloudLinkStatus string

const (
	CloudLinkStatusActive       CloudLinkStatus = "Active"
	CloudLinkStatusBuilding     CloudLinkStatus = "Building"
	CloudLinkStatusDeleting     CloudLinkStatus = "Deleting"
	CloudLinkStatusDisconnected CloudLinkStatus = "Disconnected"
	CloudLinkStatusFailed       CloudLinkStatus = "Failed"
)

// PossibleCloudLinkStatusValues returns the possible values for the CloudLinkStatus const type.
func PossibleCloudLinkStatusValues() []CloudLinkStatus {
	return []CloudLinkStatus{
		CloudLinkStatusActive,
		CloudLinkStatusBuilding,
		CloudLinkStatusDeleting,
		CloudLinkStatusDisconnected,
		CloudLinkStatusFailed,
	}
}

// ClusterProvisioningState - The state of the cluster provisioning
type ClusterProvisioningState string

const (
	ClusterProvisioningStateCancelled ClusterProvisioningState = "Cancelled"
	ClusterProvisioningStateDeleting  ClusterProvisioningState = "Deleting"
	ClusterProvisioningStateFailed    ClusterProvisioningState = "Failed"
	ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"
	ClusterProvisioningStateUpdating  ClusterProvisioningState = "Updating"
)

// PossibleClusterProvisioningStateValues returns the possible values for the ClusterProvisioningState const type.
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return []ClusterProvisioningState{
		ClusterProvisioningStateCancelled,
		ClusterProvisioningStateDeleting,
		ClusterProvisioningStateFailed,
		ClusterProvisioningStateSucceeded,
		ClusterProvisioningStateUpdating,
	}
}

// DNSServiceLogLevelEnum - DNS Service log level.
type DNSServiceLogLevelEnum string

const (
	DNSServiceLogLevelEnumDEBUG   DNSServiceLogLevelEnum = "DEBUG"
	DNSServiceLogLevelEnumERROR   DNSServiceLogLevelEnum = "ERROR"
	DNSServiceLogLevelEnumFATAL   DNSServiceLogLevelEnum = "FATAL"
	DNSServiceLogLevelEnumINFO    DNSServiceLogLevelEnum = "INFO"
	DNSServiceLogLevelEnumWARNING DNSServiceLogLevelEnum = "WARNING"
)

// PossibleDNSServiceLogLevelEnumValues returns the possible values for the DNSServiceLogLevelEnum const type.
func PossibleDNSServiceLogLevelEnumValues() []DNSServiceLogLevelEnum {
	return []DNSServiceLogLevelEnum{
		DNSServiceLogLevelEnumDEBUG,
		DNSServiceLogLevelEnumERROR,
		DNSServiceLogLevelEnumFATAL,
		DNSServiceLogLevelEnumINFO,
		DNSServiceLogLevelEnumWARNING,
	}
}

// DNSServiceStatusEnum - DNS Service status.
type DNSServiceStatusEnum string

const (
	DNSServiceStatusEnumFAILURE DNSServiceStatusEnum = "FAILURE"
	DNSServiceStatusEnumSUCCESS DNSServiceStatusEnum = "SUCCESS"
)

// PossibleDNSServiceStatusEnumValues returns the possible values for the DNSServiceStatusEnum const type.
func PossibleDNSServiceStatusEnumValues() []DNSServiceStatusEnum {
	return []DNSServiceStatusEnum{
		DNSServiceStatusEnumFAILURE,
		DNSServiceStatusEnumSUCCESS,
	}
}

// DatastoreProvisioningState - The state of the datastore provisioning
type DatastoreProvisioningState string

const (
	DatastoreProvisioningStateCancelled DatastoreProvisioningState = "Cancelled"
	DatastoreProvisioningStateCreating  DatastoreProvisioningState = "Creating"
	DatastoreProvisioningStateDeleting  DatastoreProvisioningState = "Deleting"
	DatastoreProvisioningStateFailed    DatastoreProvisioningState = "Failed"
	DatastoreProvisioningStatePending   DatastoreProvisioningState = "Pending"
	DatastoreProvisioningStateSucceeded DatastoreProvisioningState = "Succeeded"
	DatastoreProvisioningStateUpdating  DatastoreProvisioningState = "Updating"
)

// PossibleDatastoreProvisioningStateValues returns the possible values for the DatastoreProvisioningState const type.
func PossibleDatastoreProvisioningStateValues() []DatastoreProvisioningState {
	return []DatastoreProvisioningState{
		DatastoreProvisioningStateCancelled,
		DatastoreProvisioningStateCreating,
		DatastoreProvisioningStateDeleting,
		DatastoreProvisioningStateFailed,
		DatastoreProvisioningStatePending,
		DatastoreProvisioningStateSucceeded,
		DatastoreProvisioningStateUpdating,
	}
}

// DatastoreStatus - The operational status of the datastore
type DatastoreStatus string

const (
	DatastoreStatusAccessible        DatastoreStatus = "Accessible"
	DatastoreStatusAttached          DatastoreStatus = "Attached"
	DatastoreStatusDeadOrError       DatastoreStatus = "DeadOrError"
	DatastoreStatusDetached          DatastoreStatus = "Detached"
	DatastoreStatusInaccessible      DatastoreStatus = "Inaccessible"
	DatastoreStatusLostCommunication DatastoreStatus = "LostCommunication"
	DatastoreStatusUnknown           DatastoreStatus = "Unknown"
)

// PossibleDatastoreStatusValues returns the possible values for the DatastoreStatus const type.
func PossibleDatastoreStatusValues() []DatastoreStatus {
	return []DatastoreStatus{
		DatastoreStatusAccessible,
		DatastoreStatusAttached,
		DatastoreStatusDeadOrError,
		DatastoreStatusDetached,
		DatastoreStatusInaccessible,
		DatastoreStatusLostCommunication,
		DatastoreStatusUnknown,
	}
}

// DhcpTypeEnum - Type of DHCP: SERVER or RELAY.
type DhcpTypeEnum string

const (
	DhcpTypeEnumRELAY  DhcpTypeEnum = "RELAY"
	DhcpTypeEnumSERVER DhcpTypeEnum = "SERVER"
)

// PossibleDhcpTypeEnumValues returns the possible values for the DhcpTypeEnum const type.
func PossibleDhcpTypeEnumValues() []DhcpTypeEnum {
	return []DhcpTypeEnum{
		DhcpTypeEnumRELAY,
		DhcpTypeEnumSERVER,
	}
}

// EncryptionKeyStatus - The state of key provided
type EncryptionKeyStatus string

const (
	EncryptionKeyStatusAccessDenied EncryptionKeyStatus = "AccessDenied"
	EncryptionKeyStatusConnected    EncryptionKeyStatus = "Connected"
)

// PossibleEncryptionKeyStatusValues returns the possible values for the EncryptionKeyStatus const type.
func PossibleEncryptionKeyStatusValues() []EncryptionKeyStatus {
	return []EncryptionKeyStatus{
		EncryptionKeyStatusAccessDenied,
		EncryptionKeyStatusConnected,
	}
}

// EncryptionState - Status of customer managed encryption key
type EncryptionState string

const (
	EncryptionStateDisabled EncryptionState = "Disabled"
	EncryptionStateEnabled  EncryptionState = "Enabled"
)

// PossibleEncryptionStateValues returns the possible values for the EncryptionState const type.
func PossibleEncryptionStateValues() []EncryptionState {
	return []EncryptionState{
		EncryptionStateDisabled,
		EncryptionStateEnabled,
	}
}

// EncryptionVersionType - Property of the key if user provided or auto detected
type EncryptionVersionType string

const (
	EncryptionVersionTypeAutoDetected EncryptionVersionType = "AutoDetected"
	EncryptionVersionTypeFixed        EncryptionVersionType = "Fixed"
)

// PossibleEncryptionVersionTypeValues returns the possible values for the EncryptionVersionType const type.
func PossibleEncryptionVersionTypeValues() []EncryptionVersionType {
	return []EncryptionVersionType{
		EncryptionVersionTypeAutoDetected,
		EncryptionVersionTypeFixed,
	}
}

// ExpressRouteAuthorizationProvisioningState - The state of the ExpressRoute Circuit Authorization provisioning
type ExpressRouteAuthorizationProvisioningState string

const (
	ExpressRouteAuthorizationProvisioningStateFailed    ExpressRouteAuthorizationProvisioningState = "Failed"
	ExpressRouteAuthorizationProvisioningStateSucceeded ExpressRouteAuthorizationProvisioningState = "Succeeded"
	ExpressRouteAuthorizationProvisioningStateUpdating  ExpressRouteAuthorizationProvisioningState = "Updating"
)

// PossibleExpressRouteAuthorizationProvisioningStateValues returns the possible values for the ExpressRouteAuthorizationProvisioningState const type.
func PossibleExpressRouteAuthorizationProvisioningStateValues() []ExpressRouteAuthorizationProvisioningState {
	return []ExpressRouteAuthorizationProvisioningState{
		ExpressRouteAuthorizationProvisioningStateFailed,
		ExpressRouteAuthorizationProvisioningStateSucceeded,
		ExpressRouteAuthorizationProvisioningStateUpdating,
	}
}

// GlobalReachConnectionProvisioningState - The state of the ExpressRoute Circuit Authorization provisioning
type GlobalReachConnectionProvisioningState string

const (
	GlobalReachConnectionProvisioningStateFailed    GlobalReachConnectionProvisioningState = "Failed"
	GlobalReachConnectionProvisioningStateSucceeded GlobalReachConnectionProvisioningState = "Succeeded"
	GlobalReachConnectionProvisioningStateUpdating  GlobalReachConnectionProvisioningState = "Updating"
)

// PossibleGlobalReachConnectionProvisioningStateValues returns the possible values for the GlobalReachConnectionProvisioningState const type.
func PossibleGlobalReachConnectionProvisioningStateValues() []GlobalReachConnectionProvisioningState {
	return []GlobalReachConnectionProvisioningState{
		GlobalReachConnectionProvisioningStateFailed,
		GlobalReachConnectionProvisioningStateSucceeded,
		GlobalReachConnectionProvisioningStateUpdating,
	}
}

// GlobalReachConnectionStatus - The connection status of the global reach connection
type GlobalReachConnectionStatus string

const (
	GlobalReachConnectionStatusConnected    GlobalReachConnectionStatus = "Connected"
	GlobalReachConnectionStatusConnecting   GlobalReachConnectionStatus = "Connecting"
	GlobalReachConnectionStatusDisconnected GlobalReachConnectionStatus = "Disconnected"
)

// PossibleGlobalReachConnectionStatusValues returns the possible values for the GlobalReachConnectionStatus const type.
func PossibleGlobalReachConnectionStatusValues() []GlobalReachConnectionStatus {
	return []GlobalReachConnectionStatus{
		GlobalReachConnectionStatusConnected,
		GlobalReachConnectionStatusConnecting,
		GlobalReachConnectionStatusDisconnected,
	}
}

// HcxEnterpriseSiteStatus - The status of the HCX Enterprise Site
type HcxEnterpriseSiteStatus string

const (
	HcxEnterpriseSiteStatusAvailable   HcxEnterpriseSiteStatus = "Available"
	HcxEnterpriseSiteStatusConsumed    HcxEnterpriseSiteStatus = "Consumed"
	HcxEnterpriseSiteStatusDeactivated HcxEnterpriseSiteStatus = "Deactivated"
	HcxEnterpriseSiteStatusDeleted     HcxEnterpriseSiteStatus = "Deleted"
)

// PossibleHcxEnterpriseSiteStatusValues returns the possible values for the HcxEnterpriseSiteStatus const type.
func PossibleHcxEnterpriseSiteStatusValues() []HcxEnterpriseSiteStatus {
	return []HcxEnterpriseSiteStatus{
		HcxEnterpriseSiteStatusAvailable,
		HcxEnterpriseSiteStatusConsumed,
		HcxEnterpriseSiteStatusDeactivated,
		HcxEnterpriseSiteStatusDeleted,
	}
}

// InternetEnum - Connectivity to internet is enabled or disabled
type InternetEnum string

const (
	InternetEnumDisabled InternetEnum = "Disabled"
	InternetEnumEnabled  InternetEnum = "Enabled"
)

// PossibleInternetEnumValues returns the possible values for the InternetEnum const type.
func PossibleInternetEnumValues() []InternetEnum {
	return []InternetEnum{
		InternetEnumDisabled,
		InternetEnumEnabled,
	}
}

// MountOptionEnum - Mode that describes whether the LUN has to be mounted as a datastore or attached as a LUN
type MountOptionEnum string

const (
	MountOptionEnumATTACH MountOptionEnum = "ATTACH"
	MountOptionEnumMOUNT  MountOptionEnum = "MOUNT"
)

// PossibleMountOptionEnumValues returns the possible values for the MountOptionEnum const type.
func PossibleMountOptionEnumValues() []MountOptionEnum {
	return []MountOptionEnum{
		MountOptionEnumATTACH,
		MountOptionEnumMOUNT,
	}
}

// OptionalParamEnum - Is this parameter required or optional
type OptionalParamEnum string

const (
	OptionalParamEnumOptional OptionalParamEnum = "Optional"
	OptionalParamEnumRequired OptionalParamEnum = "Required"
)

// PossibleOptionalParamEnumValues returns the possible values for the OptionalParamEnum const type.
func PossibleOptionalParamEnumValues() []OptionalParamEnum {
	return []OptionalParamEnum{
		OptionalParamEnumOptional,
		OptionalParamEnumRequired,
	}
}

// PlacementPolicyProvisioningState - The provisioning state
type PlacementPolicyProvisioningState string

const (
	PlacementPolicyProvisioningStateBuilding  PlacementPolicyProvisioningState = "Building"
	PlacementPolicyProvisioningStateDeleting  PlacementPolicyProvisioningState = "Deleting"
	PlacementPolicyProvisioningStateFailed    PlacementPolicyProvisioningState = "Failed"
	PlacementPolicyProvisioningStateSucceeded PlacementPolicyProvisioningState = "Succeeded"
	PlacementPolicyProvisioningStateUpdating  PlacementPolicyProvisioningState = "Updating"
)

// PossiblePlacementPolicyProvisioningStateValues returns the possible values for the PlacementPolicyProvisioningState const type.
func PossiblePlacementPolicyProvisioningStateValues() []PlacementPolicyProvisioningState {
	return []PlacementPolicyProvisioningState{
		PlacementPolicyProvisioningStateBuilding,
		PlacementPolicyProvisioningStateDeleting,
		PlacementPolicyProvisioningStateFailed,
		PlacementPolicyProvisioningStateSucceeded,
		PlacementPolicyProvisioningStateUpdating,
	}
}

// PlacementPolicyState - Whether the placement policy is enabled or disabled
type PlacementPolicyState string

const (
	PlacementPolicyStateDisabled PlacementPolicyState = "Disabled"
	PlacementPolicyStateEnabled  PlacementPolicyState = "Enabled"
)

// PossiblePlacementPolicyStateValues returns the possible values for the PlacementPolicyState const type.
func PossiblePlacementPolicyStateValues() []PlacementPolicyState {
	return []PlacementPolicyState{
		PlacementPolicyStateDisabled,
		PlacementPolicyStateEnabled,
	}
}

// PlacementPolicyType - placement policy type
type PlacementPolicyType string

const (
	PlacementPolicyTypeVMHost PlacementPolicyType = "VmHost"
	PlacementPolicyTypeVMVM   PlacementPolicyType = "VmVm"
)

// PossiblePlacementPolicyTypeValues returns the possible values for the PlacementPolicyType const type.
func PossiblePlacementPolicyTypeValues() []PlacementPolicyType {
	return []PlacementPolicyType{
		PlacementPolicyTypeVMHost,
		PlacementPolicyTypeVMVM,
	}
}

// PortMirroringDirectionEnum - Direction of port mirroring profile.
type PortMirroringDirectionEnum string

const (
	PortMirroringDirectionEnumBIDIRECTIONAL PortMirroringDirectionEnum = "BIDIRECTIONAL"
	PortMirroringDirectionEnumEGRESS        PortMirroringDirectionEnum = "EGRESS"
	PortMirroringDirectionEnumINGRESS       PortMirroringDirectionEnum = "INGRESS"
)

// PossiblePortMirroringDirectionEnumValues returns the possible values for the PortMirroringDirectionEnum const type.
func PossiblePortMirroringDirectionEnumValues() []PortMirroringDirectionEnum {
	return []PortMirroringDirectionEnum{
		PortMirroringDirectionEnumBIDIRECTIONAL,
		PortMirroringDirectionEnumEGRESS,
		PortMirroringDirectionEnumINGRESS,
	}
}

// PortMirroringStatusEnum - Port Mirroring Status.
type PortMirroringStatusEnum string

const (
	PortMirroringStatusEnumFAILURE PortMirroringStatusEnum = "FAILURE"
	PortMirroringStatusEnumSUCCESS PortMirroringStatusEnum = "SUCCESS"
)

// PossiblePortMirroringStatusEnumValues returns the possible values for the PortMirroringStatusEnum const type.
func PossiblePortMirroringStatusEnumValues() []PortMirroringStatusEnum {
	return []PortMirroringStatusEnum{
		PortMirroringStatusEnumFAILURE,
		PortMirroringStatusEnumSUCCESS,
	}
}

// PrivateCloudProvisioningState - The provisioning state
type PrivateCloudProvisioningState string

const (
	PrivateCloudProvisioningStateBuilding  PrivateCloudProvisioningState = "Building"
	PrivateCloudProvisioningStateCancelled PrivateCloudProvisioningState = "Cancelled"
	PrivateCloudProvisioningStateDeleting  PrivateCloudProvisioningState = "Deleting"
	PrivateCloudProvisioningStateFailed    PrivateCloudProvisioningState = "Failed"
	PrivateCloudProvisioningStatePending   PrivateCloudProvisioningState = "Pending"
	PrivateCloudProvisioningStateSucceeded PrivateCloudProvisioningState = "Succeeded"
	PrivateCloudProvisioningStateUpdating  PrivateCloudProvisioningState = "Updating"
)

// PossiblePrivateCloudProvisioningStateValues returns the possible values for the PrivateCloudProvisioningState const type.
func PossiblePrivateCloudProvisioningStateValues() []PrivateCloudProvisioningState {
	return []PrivateCloudProvisioningState{
		PrivateCloudProvisioningStateBuilding,
		PrivateCloudProvisioningStateCancelled,
		PrivateCloudProvisioningStateDeleting,
		PrivateCloudProvisioningStateFailed,
		PrivateCloudProvisioningStatePending,
		PrivateCloudProvisioningStateSucceeded,
		PrivateCloudProvisioningStateUpdating,
	}
}

// QuotaEnabled - Host quota is active for current subscription
type QuotaEnabled string

const (
	QuotaEnabledDisabled QuotaEnabled = "Disabled"
	QuotaEnabledEnabled  QuotaEnabled = "Enabled"
)

// PossibleQuotaEnabledValues returns the possible values for the QuotaEnabled const type.
func PossibleQuotaEnabledValues() []QuotaEnabled {
	return []QuotaEnabled{
		QuotaEnabledDisabled,
		QuotaEnabledEnabled,
	}
}

// ResourceIdentityType - The type of identity used for the private cloud. The type 'SystemAssigned' refers to an implicitly
// created identity. The type 'None' will remove any identities from the Private Cloud.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
	}
}

// SSLEnum - Protect LDAP communication using SSL certificate (LDAPS)
type SSLEnum string

const (
	SSLEnumDisabled SSLEnum = "Disabled"
	SSLEnumEnabled  SSLEnum = "Enabled"
)

// PossibleSSLEnumValues returns the possible values for the SSLEnum const type.
func PossibleSSLEnumValues() []SSLEnum {
	return []SSLEnum{
		SSLEnumDisabled,
		SSLEnumEnabled,
	}
}

// ScriptExecutionParameterType - The type of execution parameter
type ScriptExecutionParameterType string

const (
	ScriptExecutionParameterTypeCredential  ScriptExecutionParameterType = "Credential"
	ScriptExecutionParameterTypeSecureValue ScriptExecutionParameterType = "SecureValue"
	ScriptExecutionParameterTypeValue       ScriptExecutionParameterType = "Value"
)

// PossibleScriptExecutionParameterTypeValues returns the possible values for the ScriptExecutionParameterType const type.
func PossibleScriptExecutionParameterTypeValues() []ScriptExecutionParameterType {
	return []ScriptExecutionParameterType{
		ScriptExecutionParameterTypeCredential,
		ScriptExecutionParameterTypeSecureValue,
		ScriptExecutionParameterTypeValue,
	}
}

// ScriptExecutionProvisioningState - The state of the script execution resource
type ScriptExecutionProvisioningState string

const (
	ScriptExecutionProvisioningStateCancelled  ScriptExecutionProvisioningState = "Cancelled"
	ScriptExecutionProvisioningStateCancelling ScriptExecutionProvisioningState = "Cancelling"
	ScriptExecutionProvisioningStateDeleting   ScriptExecutionProvisioningState = "Deleting"
	ScriptExecutionProvisioningStateFailed     ScriptExecutionProvisioningState = "Failed"
	ScriptExecutionProvisioningStatePending    ScriptExecutionProvisioningState = "Pending"
	ScriptExecutionProvisioningStateRunning    ScriptExecutionProvisioningState = "Running"
	ScriptExecutionProvisioningStateSucceeded  ScriptExecutionProvisioningState = "Succeeded"
)

// PossibleScriptExecutionProvisioningStateValues returns the possible values for the ScriptExecutionProvisioningState const type.
func PossibleScriptExecutionProvisioningStateValues() []ScriptExecutionProvisioningState {
	return []ScriptExecutionProvisioningState{
		ScriptExecutionProvisioningStateCancelled,
		ScriptExecutionProvisioningStateCancelling,
		ScriptExecutionProvisioningStateDeleting,
		ScriptExecutionProvisioningStateFailed,
		ScriptExecutionProvisioningStatePending,
		ScriptExecutionProvisioningStateRunning,
		ScriptExecutionProvisioningStateSucceeded,
	}
}

type ScriptOutputStreamType string

const (
	ScriptOutputStreamTypeError       ScriptOutputStreamType = "Error"
	ScriptOutputStreamTypeInformation ScriptOutputStreamType = "Information"
	ScriptOutputStreamTypeOutput      ScriptOutputStreamType = "Output"
	ScriptOutputStreamTypeWarning     ScriptOutputStreamType = "Warning"
)

// PossibleScriptOutputStreamTypeValues returns the possible values for the ScriptOutputStreamType const type.
func PossibleScriptOutputStreamTypeValues() []ScriptOutputStreamType {
	return []ScriptOutputStreamType{
		ScriptOutputStreamTypeError,
		ScriptOutputStreamTypeInformation,
		ScriptOutputStreamTypeOutput,
		ScriptOutputStreamTypeWarning,
	}
}

// ScriptParameterTypes - The type of parameter the script is expecting. psCredential is a PSCredentialObject
type ScriptParameterTypes string

const (
	ScriptParameterTypesBool         ScriptParameterTypes = "Bool"
	ScriptParameterTypesCredential   ScriptParameterTypes = "Credential"
	ScriptParameterTypesFloat        ScriptParameterTypes = "Float"
	ScriptParameterTypesInt          ScriptParameterTypes = "Int"
	ScriptParameterTypesSecureString ScriptParameterTypes = "SecureString"
	ScriptParameterTypesString       ScriptParameterTypes = "String"
)

// PossibleScriptParameterTypesValues returns the possible values for the ScriptParameterTypes const type.
func PossibleScriptParameterTypesValues() []ScriptParameterTypes {
	return []ScriptParameterTypes{
		ScriptParameterTypesBool,
		ScriptParameterTypesCredential,
		ScriptParameterTypesFloat,
		ScriptParameterTypesInt,
		ScriptParameterTypesSecureString,
		ScriptParameterTypesString,
	}
}

// SegmentStatusEnum - Segment status.
type SegmentStatusEnum string

const (
	SegmentStatusEnumFAILURE SegmentStatusEnum = "FAILURE"
	SegmentStatusEnumSUCCESS SegmentStatusEnum = "SUCCESS"
)

// PossibleSegmentStatusEnumValues returns the possible values for the SegmentStatusEnum const type.
func PossibleSegmentStatusEnumValues() []SegmentStatusEnum {
	return []SegmentStatusEnum{
		SegmentStatusEnumFAILURE,
		SegmentStatusEnumSUCCESS,
	}
}

// TrialStatus - Trial status
type TrialStatus string

const (
	TrialStatusTrialAvailable TrialStatus = "TrialAvailable"
	TrialStatusTrialDisabled  TrialStatus = "TrialDisabled"
	TrialStatusTrialUsed      TrialStatus = "TrialUsed"
)

// PossibleTrialStatusValues returns the possible values for the TrialStatus const type.
func PossibleTrialStatusValues() []TrialStatus {
	return []TrialStatus{
		TrialStatusTrialAvailable,
		TrialStatusTrialDisabled,
		TrialStatusTrialUsed,
	}
}

// VMGroupStatusEnum - VM Group status.
type VMGroupStatusEnum string

const (
	VMGroupStatusEnumFAILURE VMGroupStatusEnum = "FAILURE"
	VMGroupStatusEnumSUCCESS VMGroupStatusEnum = "SUCCESS"
)

// PossibleVMGroupStatusEnumValues returns the possible values for the VMGroupStatusEnum const type.
func PossibleVMGroupStatusEnumValues() []VMGroupStatusEnum {
	return []VMGroupStatusEnum{
		VMGroupStatusEnumFAILURE,
		VMGroupStatusEnumSUCCESS,
	}
}

// VMTypeEnum - Virtual machine type.
type VMTypeEnum string

const (
	VMTypeEnumEDGE    VMTypeEnum = "EDGE"
	VMTypeEnumREGULAR VMTypeEnum = "REGULAR"
	VMTypeEnumSERVICE VMTypeEnum = "SERVICE"
)

// PossibleVMTypeEnumValues returns the possible values for the VMTypeEnum const type.
func PossibleVMTypeEnumValues() []VMTypeEnum {
	return []VMTypeEnum{
		VMTypeEnumEDGE,
		VMTypeEnumREGULAR,
		VMTypeEnumSERVICE,
	}
}

// VirtualMachineRestrictMovementState - Whether VM DRS-driven movement is restricted (enabled) or not (disabled)
type VirtualMachineRestrictMovementState string

const (
	VirtualMachineRestrictMovementStateDisabled VirtualMachineRestrictMovementState = "Disabled"
	VirtualMachineRestrictMovementStateEnabled  VirtualMachineRestrictMovementState = "Enabled"
)

// PossibleVirtualMachineRestrictMovementStateValues returns the possible values for the VirtualMachineRestrictMovementState const type.
func PossibleVirtualMachineRestrictMovementStateValues() []VirtualMachineRestrictMovementState {
	return []VirtualMachineRestrictMovementState{
		VirtualMachineRestrictMovementStateDisabled,
		VirtualMachineRestrictMovementStateEnabled,
	}
}

// VisibilityParameterEnum - Should this parameter be visible to arm and passed in the parameters argument when executing
type VisibilityParameterEnum string

const (
	VisibilityParameterEnumHidden  VisibilityParameterEnum = "Hidden"
	VisibilityParameterEnumVisible VisibilityParameterEnum = "Visible"
)

// PossibleVisibilityParameterEnumValues returns the possible values for the VisibilityParameterEnum const type.
func PossibleVisibilityParameterEnumValues() []VisibilityParameterEnum {
	return []VisibilityParameterEnum{
		VisibilityParameterEnumHidden,
		VisibilityParameterEnumVisible,
	}
}

// WorkloadNetworkDNSServiceProvisioningState - The provisioning state
type WorkloadNetworkDNSServiceProvisioningState string

const (
	WorkloadNetworkDNSServiceProvisioningStateBuilding  WorkloadNetworkDNSServiceProvisioningState = "Building"
	WorkloadNetworkDNSServiceProvisioningStateDeleting  WorkloadNetworkDNSServiceProvisioningState = "Deleting"
	WorkloadNetworkDNSServiceProvisioningStateFailed    WorkloadNetworkDNSServiceProvisioningState = "Failed"
	WorkloadNetworkDNSServiceProvisioningStateSucceeded WorkloadNetworkDNSServiceProvisioningState = "Succeeded"
	WorkloadNetworkDNSServiceProvisioningStateUpdating  WorkloadNetworkDNSServiceProvisioningState = "Updating"
)

// PossibleWorkloadNetworkDNSServiceProvisioningStateValues returns the possible values for the WorkloadNetworkDNSServiceProvisioningState const type.
func PossibleWorkloadNetworkDNSServiceProvisioningStateValues() []WorkloadNetworkDNSServiceProvisioningState {
	return []WorkloadNetworkDNSServiceProvisioningState{
		WorkloadNetworkDNSServiceProvisioningStateBuilding,
		WorkloadNetworkDNSServiceProvisioningStateDeleting,
		WorkloadNetworkDNSServiceProvisioningStateFailed,
		WorkloadNetworkDNSServiceProvisioningStateSucceeded,
		WorkloadNetworkDNSServiceProvisioningStateUpdating,
	}
}

// WorkloadNetworkDNSZoneProvisioningState - The provisioning state
type WorkloadNetworkDNSZoneProvisioningState string

const (
	WorkloadNetworkDNSZoneProvisioningStateBuilding  WorkloadNetworkDNSZoneProvisioningState = "Building"
	WorkloadNetworkDNSZoneProvisioningStateDeleting  WorkloadNetworkDNSZoneProvisioningState = "Deleting"
	WorkloadNetworkDNSZoneProvisioningStateFailed    WorkloadNetworkDNSZoneProvisioningState = "Failed"
	WorkloadNetworkDNSZoneProvisioningStateSucceeded WorkloadNetworkDNSZoneProvisioningState = "Succeeded"
	WorkloadNetworkDNSZoneProvisioningStateUpdating  WorkloadNetworkDNSZoneProvisioningState = "Updating"
)

// PossibleWorkloadNetworkDNSZoneProvisioningStateValues returns the possible values for the WorkloadNetworkDNSZoneProvisioningState const type.
func PossibleWorkloadNetworkDNSZoneProvisioningStateValues() []WorkloadNetworkDNSZoneProvisioningState {
	return []WorkloadNetworkDNSZoneProvisioningState{
		WorkloadNetworkDNSZoneProvisioningStateBuilding,
		WorkloadNetworkDNSZoneProvisioningStateDeleting,
		WorkloadNetworkDNSZoneProvisioningStateFailed,
		WorkloadNetworkDNSZoneProvisioningStateSucceeded,
		WorkloadNetworkDNSZoneProvisioningStateUpdating,
	}
}

// WorkloadNetworkDhcpProvisioningState - The provisioning state
type WorkloadNetworkDhcpProvisioningState string

const (
	WorkloadNetworkDhcpProvisioningStateBuilding  WorkloadNetworkDhcpProvisioningState = "Building"
	WorkloadNetworkDhcpProvisioningStateDeleting  WorkloadNetworkDhcpProvisioningState = "Deleting"
	WorkloadNetworkDhcpProvisioningStateFailed    WorkloadNetworkDhcpProvisioningState = "Failed"
	WorkloadNetworkDhcpProvisioningStateSucceeded WorkloadNetworkDhcpProvisioningState = "Succeeded"
	WorkloadNetworkDhcpProvisioningStateUpdating  WorkloadNetworkDhcpProvisioningState = "Updating"
)

// PossibleWorkloadNetworkDhcpProvisioningStateValues returns the possible values for the WorkloadNetworkDhcpProvisioningState const type.
func PossibleWorkloadNetworkDhcpProvisioningStateValues() []WorkloadNetworkDhcpProvisioningState {
	return []WorkloadNetworkDhcpProvisioningState{
		WorkloadNetworkDhcpProvisioningStateBuilding,
		WorkloadNetworkDhcpProvisioningStateDeleting,
		WorkloadNetworkDhcpProvisioningStateFailed,
		WorkloadNetworkDhcpProvisioningStateSucceeded,
		WorkloadNetworkDhcpProvisioningStateUpdating,
	}
}

// WorkloadNetworkPortMirroringProvisioningState - The provisioning state
type WorkloadNetworkPortMirroringProvisioningState string

const (
	WorkloadNetworkPortMirroringProvisioningStateBuilding  WorkloadNetworkPortMirroringProvisioningState = "Building"
	WorkloadNetworkPortMirroringProvisioningStateDeleting  WorkloadNetworkPortMirroringProvisioningState = "Deleting"
	WorkloadNetworkPortMirroringProvisioningStateFailed    WorkloadNetworkPortMirroringProvisioningState = "Failed"
	WorkloadNetworkPortMirroringProvisioningStateSucceeded WorkloadNetworkPortMirroringProvisioningState = "Succeeded"
	WorkloadNetworkPortMirroringProvisioningStateUpdating  WorkloadNetworkPortMirroringProvisioningState = "Updating"
)

// PossibleWorkloadNetworkPortMirroringProvisioningStateValues returns the possible values for the WorkloadNetworkPortMirroringProvisioningState const type.
func PossibleWorkloadNetworkPortMirroringProvisioningStateValues() []WorkloadNetworkPortMirroringProvisioningState {
	return []WorkloadNetworkPortMirroringProvisioningState{
		WorkloadNetworkPortMirroringProvisioningStateBuilding,
		WorkloadNetworkPortMirroringProvisioningStateDeleting,
		WorkloadNetworkPortMirroringProvisioningStateFailed,
		WorkloadNetworkPortMirroringProvisioningStateSucceeded,
		WorkloadNetworkPortMirroringProvisioningStateUpdating,
	}
}

// WorkloadNetworkPublicIPProvisioningState - The provisioning state
type WorkloadNetworkPublicIPProvisioningState string

const (
	WorkloadNetworkPublicIPProvisioningStateBuilding  WorkloadNetworkPublicIPProvisioningState = "Building"
	WorkloadNetworkPublicIPProvisioningStateDeleting  WorkloadNetworkPublicIPProvisioningState = "Deleting"
	WorkloadNetworkPublicIPProvisioningStateFailed    WorkloadNetworkPublicIPProvisioningState = "Failed"
	WorkloadNetworkPublicIPProvisioningStateSucceeded WorkloadNetworkPublicIPProvisioningState = "Succeeded"
	WorkloadNetworkPublicIPProvisioningStateUpdating  WorkloadNetworkPublicIPProvisioningState = "Updating"
)

// PossibleWorkloadNetworkPublicIPProvisioningStateValues returns the possible values for the WorkloadNetworkPublicIPProvisioningState const type.
func PossibleWorkloadNetworkPublicIPProvisioningStateValues() []WorkloadNetworkPublicIPProvisioningState {
	return []WorkloadNetworkPublicIPProvisioningState{
		WorkloadNetworkPublicIPProvisioningStateBuilding,
		WorkloadNetworkPublicIPProvisioningStateDeleting,
		WorkloadNetworkPublicIPProvisioningStateFailed,
		WorkloadNetworkPublicIPProvisioningStateSucceeded,
		WorkloadNetworkPublicIPProvisioningStateUpdating,
	}
}

// WorkloadNetworkSegmentProvisioningState - The provisioning state
type WorkloadNetworkSegmentProvisioningState string

const (
	WorkloadNetworkSegmentProvisioningStateBuilding  WorkloadNetworkSegmentProvisioningState = "Building"
	WorkloadNetworkSegmentProvisioningStateDeleting  WorkloadNetworkSegmentProvisioningState = "Deleting"
	WorkloadNetworkSegmentProvisioningStateFailed    WorkloadNetworkSegmentProvisioningState = "Failed"
	WorkloadNetworkSegmentProvisioningStateSucceeded WorkloadNetworkSegmentProvisioningState = "Succeeded"
	WorkloadNetworkSegmentProvisioningStateUpdating  WorkloadNetworkSegmentProvisioningState = "Updating"
)

// PossibleWorkloadNetworkSegmentProvisioningStateValues returns the possible values for the WorkloadNetworkSegmentProvisioningState const type.
func PossibleWorkloadNetworkSegmentProvisioningStateValues() []WorkloadNetworkSegmentProvisioningState {
	return []WorkloadNetworkSegmentProvisioningState{
		WorkloadNetworkSegmentProvisioningStateBuilding,
		WorkloadNetworkSegmentProvisioningStateDeleting,
		WorkloadNetworkSegmentProvisioningStateFailed,
		WorkloadNetworkSegmentProvisioningStateSucceeded,
		WorkloadNetworkSegmentProvisioningStateUpdating,
	}
}

// WorkloadNetworkVMGroupProvisioningState - The provisioning state
type WorkloadNetworkVMGroupProvisioningState string

const (
	WorkloadNetworkVMGroupProvisioningStateBuilding  WorkloadNetworkVMGroupProvisioningState = "Building"
	WorkloadNetworkVMGroupProvisioningStateDeleting  WorkloadNetworkVMGroupProvisioningState = "Deleting"
	WorkloadNetworkVMGroupProvisioningStateFailed    WorkloadNetworkVMGroupProvisioningState = "Failed"
	WorkloadNetworkVMGroupProvisioningStateSucceeded WorkloadNetworkVMGroupProvisioningState = "Succeeded"
	WorkloadNetworkVMGroupProvisioningStateUpdating  WorkloadNetworkVMGroupProvisioningState = "Updating"
)

// PossibleWorkloadNetworkVMGroupProvisioningStateValues returns the possible values for the WorkloadNetworkVMGroupProvisioningState const type.
func PossibleWorkloadNetworkVMGroupProvisioningStateValues() []WorkloadNetworkVMGroupProvisioningState {
	return []WorkloadNetworkVMGroupProvisioningState{
		WorkloadNetworkVMGroupProvisioningStateBuilding,
		WorkloadNetworkVMGroupProvisioningStateDeleting,
		WorkloadNetworkVMGroupProvisioningStateFailed,
		WorkloadNetworkVMGroupProvisioningStateSucceeded,
		WorkloadNetworkVMGroupProvisioningStateUpdating,
	}
}
