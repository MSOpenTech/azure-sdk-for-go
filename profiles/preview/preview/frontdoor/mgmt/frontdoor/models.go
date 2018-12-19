// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package frontdoor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/frontdoor/mgmt/2018-08-01-preview/frontdoor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Action = original.Action

const (
	Allow Action = original.Allow
	Block Action = original.Block
	Log   Action = original.Log
)

type Availability = original.Availability

const (
	Available   Availability = original.Available
	Unavailable Availability = original.Unavailable
)

type CertificateSource = original.CertificateSource

const (
	CertificateSourceAzureKeyVault CertificateSource = original.CertificateSourceAzureKeyVault
	CertificateSourceFrontDoor     CertificateSource = original.CertificateSourceFrontDoor
)

type CertificateType = original.CertificateType

const (
	Dedicated CertificateType = original.Dedicated
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	Disabled  CustomHTTPSProvisioningState = original.Disabled
	Disabling CustomHTTPSProvisioningState = original.Disabling
	Enabled   CustomHTTPSProvisioningState = original.Enabled
	Enabling  CustomHTTPSProvisioningState = original.Enabling
	Failed    CustomHTTPSProvisioningState = original.Failed
)

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

type DestinationProtocol = original.DestinationProtocol

const (
	HTTP         DestinationProtocol = original.HTTP
	HTTPS        DestinationProtocol = original.HTTPS
	MatchRequest DestinationProtocol = original.MatchRequest
)

type DynamicCompressionEnabled = original.DynamicCompressionEnabled

const (
	DynamicCompressionEnabledDisabled DynamicCompressionEnabled = original.DynamicCompressionEnabledDisabled
	DynamicCompressionEnabledEnabled  DynamicCompressionEnabled = original.DynamicCompressionEnabledEnabled
)

type EnabledState = original.EnabledState

const (
	EnabledStateDisabled EnabledState = original.EnabledStateDisabled
	EnabledStateEnabled  EnabledState = original.EnabledStateEnabled
)

type EnabledStateEnum = original.EnabledStateEnum

const (
	EnabledStateEnumDisabled EnabledStateEnum = original.EnabledStateEnumDisabled
	EnabledStateEnumEnabled  EnabledStateEnum = original.EnabledStateEnumEnabled
)

type ForwardingProtocol = original.ForwardingProtocol

const (
	ForwardingProtocolHTTPOnly     ForwardingProtocol = original.ForwardingProtocolHTTPOnly
	ForwardingProtocolHTTPSOnly    ForwardingProtocol = original.ForwardingProtocolHTTPSOnly
	ForwardingProtocolMatchRequest ForwardingProtocol = original.ForwardingProtocolMatchRequest
)

type MatchCondition = original.MatchCondition

const (
	PostArgs      MatchCondition = original.PostArgs
	QueryString   MatchCondition = original.QueryString
	RemoteAddr    MatchCondition = original.RemoteAddr
	RequestBody   MatchCondition = original.RequestBody
	RequestHeader MatchCondition = original.RequestHeader
	RequestMethod MatchCondition = original.RequestMethod
	RequestURI    MatchCondition = original.RequestURI
)

type Mode = original.Mode

const (
	Detection  Mode = original.Detection
	Prevention Mode = original.Prevention
)

type NetworkOperationStatus = original.NetworkOperationStatus

const (
	NetworkOperationStatusFailed     NetworkOperationStatus = original.NetworkOperationStatusFailed
	NetworkOperationStatusInProgress NetworkOperationStatus = original.NetworkOperationStatusInProgress
	NetworkOperationStatusSucceeded  NetworkOperationStatus = original.NetworkOperationStatusSucceeded
)

type Operator = original.Operator

const (
	Any                Operator = original.Any
	BeginsWith         Operator = original.BeginsWith
	Contains           Operator = original.Contains
	EndsWith           Operator = original.EndsWith
	Equal              Operator = original.Equal
	GeoMatch           Operator = original.GeoMatch
	GreaterThan        Operator = original.GreaterThan
	GreaterThanOrEqual Operator = original.GreaterThanOrEqual
	IPMatch            Operator = original.IPMatch
	LessThan           Operator = original.LessThan
	LessThanOrEqual    Operator = original.LessThanOrEqual
)

type Protocol = original.Protocol

const (
	ProtocolHTTP  Protocol = original.ProtocolHTTP
	ProtocolHTTPS Protocol = original.ProtocolHTTPS
)

type Query = original.Query

const (
	StripAll  Query = original.StripAll
	StripNone Query = original.StripNone
)

type RedirectProtocol = original.RedirectProtocol

const (
	Found302             RedirectProtocol = original.Found302
	Moved301             RedirectProtocol = original.Moved301
	PermanentRedirect308 RedirectProtocol = original.PermanentRedirect308
	TemporaryRedirect307 RedirectProtocol = original.TemporaryRedirect307
)

type ResourceState = original.ResourceState

const (
	ResourceStateCreating  ResourceState = original.ResourceStateCreating
	ResourceStateDeleting  ResourceState = original.ResourceStateDeleting
	ResourceStateDisabled  ResourceState = original.ResourceStateDisabled
	ResourceStateDisabling ResourceState = original.ResourceStateDisabling
	ResourceStateEnabled   ResourceState = original.ResourceStateEnabled
	ResourceStateEnabling  ResourceState = original.ResourceStateEnabling
)

type ResourceType = original.ResourceType

const (
	MicrosoftNetworkfrontDoors                  ResourceType = original.MicrosoftNetworkfrontDoors
	MicrosoftNetworkfrontDoorsfrontendEndpoints ResourceType = original.MicrosoftNetworkfrontDoorsfrontendEndpoints
)

type RouteType = original.RouteType

const (
	Forward  RouteType = original.Forward
	Redirect RouteType = original.Redirect
)

type RuleGroupOverride = original.RuleGroupOverride

const (
	SQLInjection RuleGroupOverride = original.SQLInjection
	XSS          RuleGroupOverride = original.XSS
)

type RuleSetType = original.RuleSetType

const (
	RuleSetTypeAzureManagedRuleSet RuleSetType = original.RuleSetTypeAzureManagedRuleSet
	RuleSetTypeUnknown             RuleSetType = original.RuleSetTypeUnknown
)

type RuleType = original.RuleType

const (
	MatchRule     RuleType = original.MatchRule
	RateLimitRule RuleType = original.RateLimitRule
)

type SessionAffinityEnabledState = original.SessionAffinityEnabledState

const (
	SessionAffinityEnabledStateDisabled SessionAffinityEnabledState = original.SessionAffinityEnabledStateDisabled
	SessionAffinityEnabledStateEnabled  SessionAffinityEnabledState = original.SessionAffinityEnabledStateEnabled
)

type TLSProtocolType = original.TLSProtocolType

const (
	ServerNameIndication TLSProtocolType = original.ServerNameIndication
)

type Transform = original.Transform

const (
	HTMLEntityDecode Transform = original.HTMLEntityDecode
	Lowercase        Transform = original.Lowercase
	RemoveNulls      Transform = original.RemoveNulls
	Trim             Transform = original.Trim
	Uppercase        Transform = original.Uppercase
	URLDecode        Transform = original.URLDecode
	URLEncode        Transform = original.URLEncode
)

type WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicy

const (
	WebApplicationFirewallPolicyCreating  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyCreating
	WebApplicationFirewallPolicyDeleting  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyDeleting
	WebApplicationFirewallPolicyDisabled  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyDisabled
	WebApplicationFirewallPolicyDisabling WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyDisabling
	WebApplicationFirewallPolicyEnabled   WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyEnabled
	WebApplicationFirewallPolicyEnabling  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyEnabling
)

type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type AzureManagedOverrideRuleGroup = original.AzureManagedOverrideRuleGroup
type AzureManagedRuleSet = original.AzureManagedRuleSet
type Backend = original.Backend
type BackendPool = original.BackendPool
type BackendPoolListResult = original.BackendPoolListResult
type BackendPoolListResultIterator = original.BackendPoolListResultIterator
type BackendPoolListResultPage = original.BackendPoolListResultPage
type BackendPoolProperties = original.BackendPoolProperties
type BackendPoolUpdateParameters = original.BackendPoolUpdateParameters
type BackendPoolsClient = original.BackendPoolsClient
type BackendPoolsCreateOrUpdateFuture = original.BackendPoolsCreateOrUpdateFuture
type BackendPoolsDeleteFuture = original.BackendPoolsDeleteFuture
type BaseClient = original.BaseClient
type BasicManagedRuleSet = original.BasicManagedRuleSet
type CacheConfiguration = original.CacheConfiguration
type CertificateSourceParameters = original.CertificateSourceParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CustomHTTPSConfiguration = original.CustomHTTPSConfiguration
type CustomRule = original.CustomRule
type CustomRules = original.CustomRules
type EndpointsClient = original.EndpointsClient
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type FrontDoor = original.FrontDoor
type FrontDoorsClient = original.FrontDoorsClient
type FrontDoorsCreateOrUpdateFutureType = original.FrontDoorsCreateOrUpdateFutureType
type FrontDoorsDeleteFutureType = original.FrontDoorsDeleteFutureType
type FrontendEndpoint = original.FrontendEndpoint
type FrontendEndpointProperties = original.FrontendEndpointProperties
type FrontendEndpointUpdateParameters = original.FrontendEndpointUpdateParameters
type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink = original.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink
type FrontendEndpointsClient = original.FrontendEndpointsClient
type FrontendEndpointsCreateOrUpdateFuture = original.FrontendEndpointsCreateOrUpdateFuture
type FrontendEndpointsDeleteFuture = original.FrontendEndpointsDeleteFuture
type FrontendEndpointsDisableHTTPSFuture = original.FrontendEndpointsDisableHTTPSFuture
type FrontendEndpointsEnableHTTPSFuture = original.FrontendEndpointsEnableHTTPSFuture
type FrontendEndpointsListResult = original.FrontendEndpointsListResult
type FrontendEndpointsListResultIterator = original.FrontendEndpointsListResultIterator
type FrontendEndpointsListResultPage = original.FrontendEndpointsListResultPage
type HealthProbeSettingsClient = original.HealthProbeSettingsClient
type HealthProbeSettingsCreateOrUpdateFuture = original.HealthProbeSettingsCreateOrUpdateFuture
type HealthProbeSettingsDeleteFuture = original.HealthProbeSettingsDeleteFuture
type HealthProbeSettingsListResult = original.HealthProbeSettingsListResult
type HealthProbeSettingsListResultIterator = original.HealthProbeSettingsListResultIterator
type HealthProbeSettingsListResultPage = original.HealthProbeSettingsListResultPage
type HealthProbeSettingsModel = original.HealthProbeSettingsModel
type HealthProbeSettingsProperties = original.HealthProbeSettingsProperties
type HealthProbeSettingsUpdateParameters = original.HealthProbeSettingsUpdateParameters
type KeyVaultCertificateSourceParameters = original.KeyVaultCertificateSourceParameters
type KeyVaultCertificateSourceParametersVault = original.KeyVaultCertificateSourceParametersVault
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type LoadBalancingSettingsClient = original.LoadBalancingSettingsClient
type LoadBalancingSettingsCreateOrUpdateFuture = original.LoadBalancingSettingsCreateOrUpdateFuture
type LoadBalancingSettingsDeleteFuture = original.LoadBalancingSettingsDeleteFuture
type LoadBalancingSettingsListResult = original.LoadBalancingSettingsListResult
type LoadBalancingSettingsListResultIterator = original.LoadBalancingSettingsListResultIterator
type LoadBalancingSettingsListResultPage = original.LoadBalancingSettingsListResultPage
type LoadBalancingSettingsModel = original.LoadBalancingSettingsModel
type LoadBalancingSettingsProperties = original.LoadBalancingSettingsProperties
type LoadBalancingSettingsUpdateParameters = original.LoadBalancingSettingsUpdateParameters
type ManagedRuleSet = original.ManagedRuleSet
type ManagedRuleSets = original.ManagedRuleSets
type MatchCondition1 = original.MatchCondition1
type PoliciesClient = original.PoliciesClient
type PoliciesDeleteFuture = original.PoliciesDeleteFuture
type PolicySettings = original.PolicySettings
type Properties = original.Properties
type PurgeParameters = original.PurgeParameters
type RedirectConfiguration = original.RedirectConfiguration
type Resource = original.Resource
type RoutingRule = original.RoutingRule
type RoutingRuleListResult = original.RoutingRuleListResult
type RoutingRuleListResultIterator = original.RoutingRuleListResultIterator
type RoutingRuleListResultPage = original.RoutingRuleListResultPage
type RoutingRuleProperties = original.RoutingRuleProperties
type RoutingRuleUpdateParameters = original.RoutingRuleUpdateParameters
type RoutingRulesClient = original.RoutingRulesClient
type RoutingRulesCreateOrUpdateFuture = original.RoutingRulesCreateOrUpdateFuture
type RoutingRulesDeleteFuture = original.RoutingRulesDeleteFuture
type SubResource = original.SubResource
type TagsObject = original.TagsObject
type UpdateParameters = original.UpdateParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type WebApplicationFirewallPolicy1 = original.WebApplicationFirewallPolicy1
type WebApplicationFirewallPolicyListResult = original.WebApplicationFirewallPolicyListResult
type WebApplicationFirewallPolicyListResultIterator = original.WebApplicationFirewallPolicyListResultIterator
type WebApplicationFirewallPolicyListResultPage = original.WebApplicationFirewallPolicyListResultPage
type WebApplicationFirewallPolicyPropertiesFormat = original.WebApplicationFirewallPolicyPropertiesFormat

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewBackendPoolListResultIterator(page BackendPoolListResultPage) BackendPoolListResultIterator {
	return original.NewBackendPoolListResultIterator(page)
}
func NewBackendPoolListResultPage(getNextPage func(context.Context, BackendPoolListResult) (BackendPoolListResult, error)) BackendPoolListResultPage {
	return original.NewBackendPoolListResultPage(getNextPage)
}
func NewBackendPoolsClient(subscriptionID string) BackendPoolsClient {
	return original.NewBackendPoolsClient(subscriptionID)
}
func NewBackendPoolsClientWithBaseURI(baseURI string, subscriptionID string) BackendPoolsClient {
	return original.NewBackendPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontDoorsClient(subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClient(subscriptionID)
}
func NewFrontDoorsClientWithBaseURI(baseURI string, subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsClient(subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClient(subscriptionID)
}
func NewFrontendEndpointsClientWithBaseURI(baseURI string, subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsListResultIterator(page FrontendEndpointsListResultPage) FrontendEndpointsListResultIterator {
	return original.NewFrontendEndpointsListResultIterator(page)
}
func NewFrontendEndpointsListResultPage(getNextPage func(context.Context, FrontendEndpointsListResult) (FrontendEndpointsListResult, error)) FrontendEndpointsListResultPage {
	return original.NewFrontendEndpointsListResultPage(getNextPage)
}
func NewHealthProbeSettingsClient(subscriptionID string) HealthProbeSettingsClient {
	return original.NewHealthProbeSettingsClient(subscriptionID)
}
func NewHealthProbeSettingsClientWithBaseURI(baseURI string, subscriptionID string) HealthProbeSettingsClient {
	return original.NewHealthProbeSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewHealthProbeSettingsListResultIterator(page HealthProbeSettingsListResultPage) HealthProbeSettingsListResultIterator {
	return original.NewHealthProbeSettingsListResultIterator(page)
}
func NewHealthProbeSettingsListResultPage(getNextPage func(context.Context, HealthProbeSettingsListResult) (HealthProbeSettingsListResult, error)) HealthProbeSettingsListResultPage {
	return original.NewHealthProbeSettingsListResultPage(getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewLoadBalancingSettingsClient(subscriptionID string) LoadBalancingSettingsClient {
	return original.NewLoadBalancingSettingsClient(subscriptionID)
}
func NewLoadBalancingSettingsClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancingSettingsClient {
	return original.NewLoadBalancingSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLoadBalancingSettingsListResultIterator(page LoadBalancingSettingsListResultPage) LoadBalancingSettingsListResultIterator {
	return original.NewLoadBalancingSettingsListResultIterator(page)
}
func NewLoadBalancingSettingsListResultPage(getNextPage func(context.Context, LoadBalancingSettingsListResult) (LoadBalancingSettingsListResult, error)) LoadBalancingSettingsListResultPage {
	return original.NewLoadBalancingSettingsListResultPage(getNextPage)
}
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoutingRuleListResultIterator(page RoutingRuleListResultPage) RoutingRuleListResultIterator {
	return original.NewRoutingRuleListResultIterator(page)
}
func NewRoutingRuleListResultPage(getNextPage func(context.Context, RoutingRuleListResult) (RoutingRuleListResult, error)) RoutingRuleListResultPage {
	return original.NewRoutingRuleListResultPage(getNextPage)
}
func NewRoutingRulesClient(subscriptionID string) RoutingRulesClient {
	return original.NewRoutingRulesClient(subscriptionID)
}
func NewRoutingRulesClientWithBaseURI(baseURI string, subscriptionID string) RoutingRulesClient {
	return original.NewRoutingRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebApplicationFirewallPolicyListResultIterator(page WebApplicationFirewallPolicyListResultPage) WebApplicationFirewallPolicyListResultIterator {
	return original.NewWebApplicationFirewallPolicyListResultIterator(page)
}
func NewWebApplicationFirewallPolicyListResultPage(getNextPage func(context.Context, WebApplicationFirewallPolicyListResult) (WebApplicationFirewallPolicyListResult, error)) WebApplicationFirewallPolicyListResultPage {
	return original.NewWebApplicationFirewallPolicyListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleAvailabilityValues() []Availability {
	return original.PossibleAvailabilityValues()
}
func PossibleCertificateSourceValues() []CertificateSource {
	return original.PossibleCertificateSourceValues()
}
func PossibleCertificateTypeValues() []CertificateType {
	return original.PossibleCertificateTypeValues()
}
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}
func PossibleDestinationProtocolValues() []DestinationProtocol {
	return original.PossibleDestinationProtocolValues()
}
func PossibleDynamicCompressionEnabledValues() []DynamicCompressionEnabled {
	return original.PossibleDynamicCompressionEnabledValues()
}
func PossibleEnabledStateEnumValues() []EnabledStateEnum {
	return original.PossibleEnabledStateEnumValues()
}
func PossibleEnabledStateValues() []EnabledState {
	return original.PossibleEnabledStateValues()
}
func PossibleForwardingProtocolValues() []ForwardingProtocol {
	return original.PossibleForwardingProtocolValues()
}
func PossibleMatchConditionValues() []MatchCondition {
	return original.PossibleMatchConditionValues()
}
func PossibleModeValues() []Mode {
	return original.PossibleModeValues()
}
func PossibleNetworkOperationStatusValues() []NetworkOperationStatus {
	return original.PossibleNetworkOperationStatusValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleQueryValues() []Query {
	return original.PossibleQueryValues()
}
func PossibleRedirectProtocolValues() []RedirectProtocol {
	return original.PossibleRedirectProtocolValues()
}
func PossibleResourceStateValues() []ResourceState {
	return original.PossibleResourceStateValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleRouteTypeValues() []RouteType {
	return original.PossibleRouteTypeValues()
}
func PossibleRuleGroupOverrideValues() []RuleGroupOverride {
	return original.PossibleRuleGroupOverrideValues()
}
func PossibleRuleSetTypeValues() []RuleSetType {
	return original.PossibleRuleSetTypeValues()
}
func PossibleRuleTypeValues() []RuleType {
	return original.PossibleRuleTypeValues()
}
func PossibleSessionAffinityEnabledStateValues() []SessionAffinityEnabledState {
	return original.PossibleSessionAffinityEnabledStateValues()
}
func PossibleTLSProtocolTypeValues() []TLSProtocolType {
	return original.PossibleTLSProtocolTypeValues()
}
func PossibleTransformValues() []Transform {
	return original.PossibleTransformValues()
}
func PossibleWebApplicationFirewallPolicyValues() []WebApplicationFirewallPolicy {
	return original.PossibleWebApplicationFirewallPolicyValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
