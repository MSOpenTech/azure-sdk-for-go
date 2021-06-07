// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package cognitiveservices

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2021-04-30/cognitiveservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	ActionTypeInternal ActionType = original.ActionTypeInternal
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type KeyName = original.KeyName

const (
	KeyNameKey1 KeyName = original.KeyNameKey1
	KeyNameKey2 KeyName = original.KeyNameKey2
)

type KeySource = original.KeySource

const (
	KeySourceMicrosoftCognitiveServices KeySource = original.KeySourceMicrosoftCognitiveServices
	KeySourceMicrosoftKeyVault          KeySource = original.KeySourceMicrosoftKeyVault
)

type NetworkRuleAction = original.NetworkRuleAction

const (
	NetworkRuleActionAllow NetworkRuleAction = original.NetworkRuleActionAllow
	NetworkRuleActionDeny  NetworkRuleAction = original.NetworkRuleActionDeny
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoving       ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateResolvingDNS ProvisioningState = original.ProvisioningStateResolvingDNS
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type QuotaUsageStatus = original.QuotaUsageStatus

const (
	QuotaUsageStatusBlocked   QuotaUsageStatus = original.QuotaUsageStatusBlocked
	QuotaUsageStatusIncluded  QuotaUsageStatus = original.QuotaUsageStatusIncluded
	QuotaUsageStatusInOverage QuotaUsageStatus = original.QuotaUsageStatusInOverage
	QuotaUsageStatusUnknown   QuotaUsageStatus = original.QuotaUsageStatusUnknown
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription
	ResourceSkuRestrictionsReasonCodeQuotaID                     ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCodeQuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	ResourceSkuRestrictionsTypeLocation ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsTypeLocation
	ResourceSkuRestrictionsTypeZone     ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsTypeZone
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic      SkuTier = original.SkuTierBasic
	SkuTierEnterprise SkuTier = original.SkuTierEnterprise
	SkuTierFree       SkuTier = original.SkuTierFree
	SkuTierPremium    SkuTier = original.SkuTierPremium
	SkuTierStandard   SkuTier = original.SkuTierStandard
)

type UnitType = original.UnitType

const (
	UnitTypeBytes          UnitType = original.UnitTypeBytes
	UnitTypeBytesPerSecond UnitType = original.UnitTypeBytesPerSecond
	UnitTypeCount          UnitType = original.UnitTypeCount
	UnitTypeCountPerSecond UnitType = original.UnitTypeCountPerSecond
	UnitTypeMilliseconds   UnitType = original.UnitTypeMilliseconds
	UnitTypePercent        UnitType = original.UnitTypePercent
	UnitTypeSeconds        UnitType = original.UnitTypeSeconds
)

type APIKeys = original.APIKeys
type APIProperties = original.APIProperties
type Account = original.Account
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountProperties = original.AccountProperties
type AccountSku = original.AccountSku
type AccountSkuListResult = original.AccountSkuListResult
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CallRateLimit = original.CallRateLimit
type CheckDomainAvailabilityParameter = original.CheckDomainAvailabilityParameter
type CheckSkuAvailabilityParameter = original.CheckSkuAvailabilityParameter
type DeletedAccountsClient = original.DeletedAccountsClient
type DeletedAccountsPurgeFuture = original.DeletedAccountsPurgeFuture
type DomainAvailability = original.DomainAvailability
type Encryption = original.Encryption
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type IPRule = original.IPRule
type Identity = original.Identity
type KeyVaultProperties = original.KeyVaultProperties
type MetricName = original.MetricName
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type QuotaLimit = original.QuotaLimit
type RegenerateKeyParameters = original.RegenerateKeyParameters
type RequestMatchPattern = original.RequestMatchPattern
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type ResourceSkuListResult = original.ResourceSkuListResult
type ResourceSkuListResultIterator = original.ResourceSkuListResultIterator
type ResourceSkuListResultPage = original.ResourceSkuListResultPage
type ResourceSkuRestrictionInfo = original.ResourceSkuRestrictionInfo
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkusClient = original.ResourceSkusClient
type Sku = original.Sku
type SkuAvailability = original.SkuAvailability
type SkuAvailabilityListResult = original.SkuAvailabilityListResult
type SkuCapability = original.SkuCapability
type SkuChangeInfo = original.SkuChangeInfo
type SystemData = original.SystemData
type ThrottlingRule = original.ThrottlingRule
type TrackedResource = original.TrackedResource
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UserAssignedIdentity = original.UserAssignedIdentity
type UserOwnedStorage = original.UserOwnedStorage
type VirtualNetworkRule = original.VirtualNetworkRule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListResultIterator(page AccountListResultPage) AccountListResultIterator {
	return original.NewAccountListResultIterator(page)
}
func NewAccountListResultPage(cur AccountListResult, getNextPage func(context.Context, AccountListResult) (AccountListResult, error)) AccountListResultPage {
	return original.NewAccountListResultPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeletedAccountsClient(subscriptionID string) DeletedAccountsClient {
	return original.NewDeletedAccountsClient(subscriptionID)
}
func NewDeletedAccountsClientWithBaseURI(baseURI string, subscriptionID string) DeletedAccountsClient {
	return original.NewDeletedAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkuListResultIterator(page ResourceSkuListResultPage) ResourceSkuListResultIterator {
	return original.NewResourceSkuListResultIterator(page)
}
func NewResourceSkuListResultPage(cur ResourceSkuListResult, getNextPage func(context.Context, ResourceSkuListResult) (ResourceSkuListResult, error)) ResourceSkuListResultPage {
	return original.NewResourceSkuListResultPage(cur, getNextPage)
}
func NewResourceSkusClient(subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClient(subscriptionID)
}
func NewResourceSkusClientWithBaseURI(baseURI string, subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleKeyNameValues() []KeyName {
	return original.PossibleKeyNameValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return original.PossibleNetworkRuleActionValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleQuotaUsageStatusValues() []QuotaUsageStatus {
	return original.PossibleQuotaUsageStatusValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnitTypeValues() []UnitType {
	return original.PossibleUnitTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
