package storage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessTier enumerates the values for access tier.
type AccessTier string

const (
	// Cool ...
	Cool AccessTier = "Cool"
	// Hot ...
	Hot AccessTier = "Hot"
)

// PossibleAccessTierValues returns an array of possible values for the AccessTier const type.
func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{Cool, Hot}
}

// AccountExpand enumerates the values for account expand.
type AccountExpand string

const (
	// AccountExpandGeoReplicationStats ...
	AccountExpandGeoReplicationStats AccountExpand = "geoReplicationStats"
)

// PossibleAccountExpandValues returns an array of possible values for the AccountExpand const type.
func PossibleAccountExpandValues() []AccountExpand {
	return []AccountExpand{AccountExpandGeoReplicationStats}
}

// AccountStatus enumerates the values for account status.
type AccountStatus string

const (
	// Available ...
	Available AccountStatus = "available"
	// Unavailable ...
	Unavailable AccountStatus = "unavailable"
)

// PossibleAccountStatusValues returns an array of possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{Available, Unavailable}
}

// Action enumerates the values for action.
type Action string

const (
	// Allow ...
	Allow Action = "Allow"
)

// PossibleActionValues returns an array of possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{Allow}
}

// Action1 enumerates the values for action 1.
type Action1 string

const (
	// Acquire ...
	Acquire Action1 = "Acquire"
	// Break ...
	Break Action1 = "Break"
	// Change ...
	Change Action1 = "Change"
	// Release ...
	Release Action1 = "Release"
	// Renew ...
	Renew Action1 = "Renew"
)

// PossibleAction1Values returns an array of possible values for the Action1 const type.
func PossibleAction1Values() []Action1 {
	return []Action1{Acquire, Break, Change, Release, Renew}
}

// Bypass enumerates the values for bypass.
type Bypass string

const (
	// AzureServices ...
	AzureServices Bypass = "AzureServices"
	// Logging ...
	Logging Bypass = "Logging"
	// Metrics ...
	Metrics Bypass = "Metrics"
	// None ...
	None Bypass = "None"
)

// PossibleBypassValues returns an array of possible values for the Bypass const type.
func PossibleBypassValues() []Bypass {
	return []Bypass{AzureServices, Logging, Metrics, None}
}

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// DefaultActionAllow ...
	DefaultActionAllow DefaultAction = "Allow"
	// DefaultActionDeny ...
	DefaultActionDeny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{DefaultActionAllow, DefaultActionDeny}
}

// GeoReplicationStatus enumerates the values for geo replication status.
type GeoReplicationStatus string

const (
	// GeoReplicationStatusBootstrap ...
	GeoReplicationStatusBootstrap GeoReplicationStatus = "Bootstrap"
	// GeoReplicationStatusLive ...
	GeoReplicationStatusLive GeoReplicationStatus = "Live"
	// GeoReplicationStatusUnavailable ...
	GeoReplicationStatusUnavailable GeoReplicationStatus = "Unavailable"
)

// PossibleGeoReplicationStatusValues returns an array of possible values for the GeoReplicationStatus const type.
func PossibleGeoReplicationStatusValues() []GeoReplicationStatus {
	return []GeoReplicationStatus{GeoReplicationStatusBootstrap, GeoReplicationStatusLive, GeoReplicationStatusUnavailable}
}

// HTTPProtocol enumerates the values for http protocol.
type HTTPProtocol string

const (
	// HTTPS ...
	HTTPS HTTPProtocol = "https"
	// Httpshttp ...
	Httpshttp HTTPProtocol = "https,http"
)

// PossibleHTTPProtocolValues returns an array of possible values for the HTTPProtocol const type.
func PossibleHTTPProtocolValues() []HTTPProtocol {
	return []HTTPProtocol{HTTPS, Httpshttp}
}

// ImmutabilityPolicyState enumerates the values for immutability policy state.
type ImmutabilityPolicyState string

const (
	// Locked ...
	Locked ImmutabilityPolicyState = "Locked"
	// Unlocked ...
	Unlocked ImmutabilityPolicyState = "Unlocked"
)

// PossibleImmutabilityPolicyStateValues returns an array of possible values for the ImmutabilityPolicyState const type.
func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return []ImmutabilityPolicyState{Locked, Unlocked}
}

// ImmutabilityPolicyUpdateType enumerates the values for immutability policy update type.
type ImmutabilityPolicyUpdateType string

const (
	// Extend ...
	Extend ImmutabilityPolicyUpdateType = "extend"
	// Lock ...
	Lock ImmutabilityPolicyUpdateType = "lock"
	// Put ...
	Put ImmutabilityPolicyUpdateType = "put"
)

// PossibleImmutabilityPolicyUpdateTypeValues returns an array of possible values for the ImmutabilityPolicyUpdateType const type.
func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return []ImmutabilityPolicyUpdateType{Extend, Lock, Put}
}

// KeyPermission enumerates the values for key permission.
type KeyPermission string

const (
	// Full ...
	Full KeyPermission = "Full"
	// Read ...
	Read KeyPermission = "Read"
)

// PossibleKeyPermissionValues returns an array of possible values for the KeyPermission const type.
func PossibleKeyPermissionValues() []KeyPermission {
	return []KeyPermission{Full, Read}
}

// KeySource enumerates the values for key source.
type KeySource string

const (
	// MicrosoftKeyvault ...
	MicrosoftKeyvault KeySource = "Microsoft.Keyvault"
	// MicrosoftStorage ...
	MicrosoftStorage KeySource = "Microsoft.Storage"
)

// PossibleKeySourceValues returns an array of possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{MicrosoftKeyvault, MicrosoftStorage}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// BlobStorage ...
	BlobStorage Kind = "BlobStorage"
	// BlockBlobStorage ...
	BlockBlobStorage Kind = "BlockBlobStorage"
	// FileStorage ...
	FileStorage Kind = "FileStorage"
	// Storage ...
	Storage Kind = "Storage"
	// StorageV2 ...
	StorageV2 Kind = "StorageV2"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{BlobStorage, BlockBlobStorage, FileStorage, Storage, StorageV2}
}

// LeaseDuration enumerates the values for lease duration.
type LeaseDuration string

const (
	// Fixed ...
	Fixed LeaseDuration = "Fixed"
	// Infinite ...
	Infinite LeaseDuration = "Infinite"
)

// PossibleLeaseDurationValues returns an array of possible values for the LeaseDuration const type.
func PossibleLeaseDurationValues() []LeaseDuration {
	return []LeaseDuration{Fixed, Infinite}
}

// LeaseState enumerates the values for lease state.
type LeaseState string

const (
	// LeaseStateAvailable ...
	LeaseStateAvailable LeaseState = "Available"
	// LeaseStateBreaking ...
	LeaseStateBreaking LeaseState = "Breaking"
	// LeaseStateBroken ...
	LeaseStateBroken LeaseState = "Broken"
	// LeaseStateExpired ...
	LeaseStateExpired LeaseState = "Expired"
	// LeaseStateLeased ...
	LeaseStateLeased LeaseState = "Leased"
)

// PossibleLeaseStateValues returns an array of possible values for the LeaseState const type.
func PossibleLeaseStateValues() []LeaseState {
	return []LeaseState{LeaseStateAvailable, LeaseStateBreaking, LeaseStateBroken, LeaseStateExpired, LeaseStateLeased}
}

// LeaseStatus enumerates the values for lease status.
type LeaseStatus string

const (
	// LeaseStatusLocked ...
	LeaseStatusLocked LeaseStatus = "Locked"
	// LeaseStatusUnlocked ...
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

// PossibleLeaseStatusValues returns an array of possible values for the LeaseStatus const type.
func PossibleLeaseStatusValues() []LeaseStatus {
	return []LeaseStatus{LeaseStatusLocked, LeaseStatusUnlocked}
}

// Permissions enumerates the values for permissions.
type Permissions string

const (
	// A ...
	A Permissions = "a"
	// C ...
	C Permissions = "c"
	// D ...
	D Permissions = "d"
	// L ...
	L Permissions = "l"
	// P ...
	P Permissions = "p"
	// R ...
	R Permissions = "r"
	// U ...
	U Permissions = "u"
	// W ...
	W Permissions = "w"
)

// PossiblePermissionsValues returns an array of possible values for the Permissions const type.
func PossiblePermissionsValues() []Permissions {
	return []Permissions{A, C, D, L, P, R, U, W}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, ResolvingDNS, Succeeded}
}

// PublicAccess enumerates the values for public access.
type PublicAccess string

const (
	// PublicAccessBlob ...
	PublicAccessBlob PublicAccess = "Blob"
	// PublicAccessContainer ...
	PublicAccessContainer PublicAccess = "Container"
	// PublicAccessNone ...
	PublicAccessNone PublicAccess = "None"
)

// PossiblePublicAccessValues returns an array of possible values for the PublicAccess const type.
func PossiblePublicAccessValues() []PublicAccess {
	return []PublicAccess{PublicAccessBlob, PublicAccessContainer, PublicAccessNone}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AccountNameInvalid, AlreadyExists}
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

// Services enumerates the values for services.
type Services string

const (
	// B ...
	B Services = "b"
	// F ...
	F Services = "f"
	// Q ...
	Q Services = "q"
	// T ...
	T Services = "t"
)

// PossibleServicesValues returns an array of possible values for the Services const type.
func PossibleServicesValues() []Services {
	return []Services{B, F, Q, T}
}

// SignedResource enumerates the values for signed resource.
type SignedResource string

const (
	// SignedResourceB ...
	SignedResourceB SignedResource = "b"
	// SignedResourceC ...
	SignedResourceC SignedResource = "c"
	// SignedResourceF ...
	SignedResourceF SignedResource = "f"
	// SignedResourceS ...
	SignedResourceS SignedResource = "s"
)

// PossibleSignedResourceValues returns an array of possible values for the SignedResource const type.
func PossibleSignedResourceValues() []SignedResource {
	return []SignedResource{SignedResourceB, SignedResourceC, SignedResourceF, SignedResourceS}
}

// SignedResourceTypes enumerates the values for signed resource types.
type SignedResourceTypes string

const (
	// SignedResourceTypesC ...
	SignedResourceTypesC SignedResourceTypes = "c"
	// SignedResourceTypesO ...
	SignedResourceTypesO SignedResourceTypes = "o"
	// SignedResourceTypesS ...
	SignedResourceTypesS SignedResourceTypes = "s"
)

// PossibleSignedResourceTypesValues returns an array of possible values for the SignedResourceTypes const type.
func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return []SignedResourceTypes{SignedResourceTypesC, SignedResourceTypesO, SignedResourceTypesS}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// PremiumLRS ...
	PremiumLRS SkuName = "Premium_LRS"
	// PremiumZRS ...
	PremiumZRS SkuName = "Premium_ZRS"
	// StandardGRS ...
	StandardGRS SkuName = "Standard_GRS"
	// StandardLRS ...
	StandardLRS SkuName = "Standard_LRS"
	// StandardRAGRS ...
	StandardRAGRS SkuName = "Standard_RAGRS"
	// StandardZRS ...
	StandardZRS SkuName = "Standard_ZRS"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{PremiumLRS, PremiumZRS, StandardGRS, StandardLRS, StandardRAGRS, StandardZRS}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Premium ...
	Premium SkuTier = "Premium"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Premium, Standard}
}

// State enumerates the values for state.
type State string

const (
	// StateDeprovisioning ...
	StateDeprovisioning State = "deprovisioning"
	// StateFailed ...
	StateFailed State = "failed"
	// StateNetworkSourceDeleted ...
	StateNetworkSourceDeleted State = "networkSourceDeleted"
	// StateProvisioning ...
	StateProvisioning State = "provisioning"
	// StateSucceeded ...
	StateSucceeded State = "succeeded"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateDeprovisioning, StateFailed, StateNetworkSourceDeleted, StateProvisioning, StateSucceeded}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes ...
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count ...
	Count UsageUnit = "Count"
	// CountsPerSecond ...
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent ...
	Percent UsageUnit = "Percent"
	// Seconds ...
	Seconds UsageUnit = "Seconds"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{Bytes, BytesPerSecond, Count, CountsPerSecond, Percent, Seconds}
}
