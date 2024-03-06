//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

import "time"

type APIKey struct {
	// REQUIRED; The value of the API key.
	Key *string

	// The time of creation of the API key.
	Created *string

	// The user that created the API key.
	CreatedBy *string

	// The name of the API key.
	Name *string
}

// APIKeyListResponse - Response of a list operation.
type APIKeyListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*APIKey
}

// AgreementProperties - Terms properties.
type AgreementProperties struct {
	// If any version of the terms have been accepted, otherwise false.
	Accepted *bool

	// Link to HTML with Microsoft and Publisher terms.
	LicenseTextLink *string

	// Plan identifier string.
	Plan *string

	// Link to the privacy policy of the publisher.
	PrivacyPolicyLink *string

	// Product identifier string.
	Product *string

	// Publisher identifier string.
	Publisher *string

	// Date and time in UTC of when the terms were accepted. This is empty if Accepted is false.
	RetrieveDatetime *time.Time

	// Terms signature.
	Signature *string
}

type AgreementResource struct {
	// Represents the properties of the resource.
	Properties *AgreementProperties

	// READ-ONLY; ARM id of the resource.
	ID *string

	// READ-ONLY; Name of the agreement.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource.
	Type *string
}

// AgreementResourceListResponse - Response of a list operation.
type AgreementResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*AgreementResource
}

// CreateResourceSupportedProperties - Datadog resource can be created or not properties.
type CreateResourceSupportedProperties struct {
	// READ-ONLY; Indicates if selected subscription supports Datadog resource creation, if not it is already being monitored
	// for the selected organization via multi subscription feature.
	CreationSupported *bool

	// READ-ONLY; The ARM id of the subscription.
	Name *string
}

// CreateResourceSupportedResponse - Datadog resource can be created or not.
type CreateResourceSupportedResponse struct {
	// Represents the properties of the resource.
	Properties *CreateResourceSupportedProperties
}

type CreateResourceSupportedResponseList struct {
	Value []*CreateResourceSupportedResponse
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them
// from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag. Exclusion takes priority over inclusion.
	Action *TagAction

	// The name (also known as the key) of the tag.
	Name *string

	// The value of the tag.
	Value *string
}

type Host struct {
	// The aliases for the host installed via the Datadog agent.
	Aliases []*string

	// The Datadog integrations reporting metrics for the host.
	Apps []*string
	Meta *HostMetadata

	// The name of the host.
	Name *string
}

// HostListResponse - Response of a list operation.
type HostListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*Host
}

type HostMetadata struct {
	// The agent version.
	AgentVersion  *string
	InstallMethod *InstallMethod
	LogsAgent     *LogsAgent
}

type IdentityProperties struct {
	// Specifies the identity type of the Datadog Monitor. At this time the only allowed value is 'SystemAssigned'.
	Type *ManagedIdentityTypes

	// READ-ONLY; The identity ID.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

type InstallMethod struct {
	// The installer version.
	InstallerVersion *string

	// The tool.
	Tool *string

	// The tool version.
	ToolVersion *string
}

// LinkedResource - The definition of a linked resource.
type LinkedResource struct {
	// The ARM id of the linked resource.
	ID *string
}

// LinkedResourceListResponse - Response of a list operation.
type LinkedResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*LinkedResource
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendResourceLogs flag is enabled. If empty,
	// all resources will be captured. If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
	// resources with the associated tags.
	FilteringTags []*FilteringTag

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *bool

	// Flag specifying if Azure resource logs should be sent for the Monitor resource.
	SendResourceLogs *bool

	// Flag specifying if Azure subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *bool
}

type LogsAgent struct {
	// The transport.
	Transport *string
}

// MetricRules - Set of rules for sending metrics for the Monitor resource.
type MetricRules struct {
	// List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action
	// is specified, the rules will apply to the list of all available resources. If
	// Include actions are specified, the rules will only include resources with the associated tags.
	FilteringTags []*FilteringTag
}

// MonitorProperties - Properties specific to the monitor resource.
type MonitorProperties struct {
	// Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey
	// is required as well.
	DatadogOrganizationProperties *OrganizationProperties

	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus

	// Includes name, email and optionally, phone number. User Information can't be null.
	UserInfo *UserInfo

	// READ-ONLY
	LiftrResourceCategory *LiftrResourceCategories

	// READ-ONLY; The priority of the resource.
	LiftrResourcePreference *int32

	// READ-ONLY; Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in time, the resource
	// will go in Suspended state.
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

type MonitorResource struct {
	// REQUIRED
	Location *string
	Identity *IdentityProperties

	// Properties specific to the monitor resource.
	Properties *MonitorProperties
	SKU        *ResourceSKU

	// Dictionary of
	Tags map[string]*string

	// READ-ONLY; ARM id of the monitor resource.
	ID *string

	// READ-ONLY; Name of the monitor resource.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the monitor resource.
	Type *string
}

// MonitorResourceListResponse - Response of a list operation.
type MonitorResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*MonitorResource
}

// MonitorResourceUpdateParameters - The parameters for a PATCH request to a monitor resource.
type MonitorResourceUpdateParameters struct {
	// The set of properties that can be update in a PATCH request to a monitor resource.
	Properties *MonitorUpdateProperties
	SKU        *ResourceSKU

	// The new tags of the monitor resource.
	Tags map[string]*string
}

// MonitorUpdateProperties - The set of properties that can be update in a PATCH request to a monitor resource.
type MonitorUpdateProperties struct {
	// The new cloud security posture management value of the monitor resource. This collects configuration information for all
	// resources in a subscription and track conformance to industry benchmarks.
	Cspm *bool

	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus
}

// MonitoredResource - The properties of a resource currently being monitored by the Datadog monitor resource.
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string

	// Flag indicating if resource is sending logs to Datadog.
	SendingLogs *bool

	// Flag indicating if resource is sending metrics to Datadog.
	SendingMetrics *bool
}

// MonitoredResourceListResponse - Response of a list operation.
type MonitoredResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*MonitoredResource
}

// MonitoredSubscription - The list of subscriptions and it's monitoring status by current Datadog monitor.
type MonitoredSubscription struct {
	// The reason of not monitoring the subscription.
	Error *string

	// The state of monitoring.
	Status *Status

	// The subscriptionId to be monitored.
	SubscriptionID *string

	// Definition of the properties for a TagRules resource.
	TagRules *MonitoringTagRulesProperties
}

// MonitoredSubscriptionProperties - The request to update subscriptions needed to be monitored by the Datadog monitor resource.
type MonitoredSubscriptionProperties struct {
	// The request to update subscriptions needed to be monitored by the Datadog monitor resource.
	Properties *SubscriptionList

	// READ-ONLY; The id of the monitored subscription resource.
	ID *string

	// READ-ONLY; Name of the monitored subscription resource.
	Name *string

	// READ-ONLY; The type of the monitored subscription resource.
	Type *string
}

type MonitoredSubscriptionPropertiesList struct {
	Value []*MonitoredSubscriptionProperties
}

// MonitoringTagRules - Capture logs and metrics of Azure resources based on ARM tags.
type MonitoringTagRules struct {
	// Definition of the properties for a TagRules resource.
	Properties *MonitoringTagRulesProperties

	// READ-ONLY; The id of the rule set.
	ID *string

	// READ-ONLY; Name of the rule set.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the rule set.
	Type *string
}

// MonitoringTagRulesListResponse - Response of a list operation.
type MonitoringTagRulesListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*MonitoringTagRules
}

// MonitoringTagRulesProperties - Definition of the properties for a TagRules resource.
type MonitoringTagRulesProperties struct {
	// Configuration to enable/disable auto-muting flag
	Automuting *bool

	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write monitors'.
	Description *string

	// Operation type, e.g., read, write, delete, etc.
	Operation *string

	// Service provider, i.e., Microsoft.Datadog.
	Provider *string

	// Type on which the operation is performed, e.g., 'monitors'.
	Resource *string
}

// OperationListResult - Result of GET request to list the Microsoft.Datadog operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of operations supported by the Microsoft.Datadog provider.
	Value []*OperationResult
}

// OperationResult - A Microsoft.Datadog REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name, i.e., {provider}/{resource}/{operation}.
	Name *string
}

// OrganizationProperties - Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey,
// and Applicationkey is required as well.
type OrganizationProperties struct {
	// Api key associated to the Datadog organization.
	APIKey *string

	// Application key associated to the Datadog organization.
	ApplicationKey *string

	// The configuration which describes the state of cloud security posture management. This collects configuration information
	// for all resources in a subscription and track conformance to industry
	// benchmarks.
	Cspm *bool

	// The Id of the Enterprise App used for Single sign on.
	EnterpriseAppID *string

	// Id of the Datadog organization.
	ID *string

	// The auth code used to linking to an existing Datadog organization.
	LinkingAuthCode *string

	// The client_id from an existing in exchange for an auth token to link organization.
	LinkingClientID *string

	// Name of the Datadog organization.
	Name *string

	// The redirect URI for linking.
	RedirectURI *string
}

type ResourceSKU struct {
	// REQUIRED; Name of the SKU in {PlanId} format. For Terraform, the only allowed value is 'linking'.
	Name *string
}

type SetPasswordLink struct {
	SetPasswordLink *string
}

type SingleSignOnProperties struct {
	// The Id of the Enterprise App used for Single sign-on.
	EnterpriseAppID *string

	// Various states of the SSO resource
	SingleSignOnState *SingleSignOnStates

	// READ-ONLY
	ProvisioningState *ProvisioningState

	// READ-ONLY; The login URL specific to this Datadog Organization.
	SingleSignOnURL *string
}

type SingleSignOnResource struct {
	Properties *SingleSignOnProperties

	// READ-ONLY; ARM id of the resource.
	ID *string

	// READ-ONLY; Name of the configuration.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource.
	Type *string
}

// SingleSignOnResourceListResponse - Response of a list operation.
type SingleSignOnResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*SingleSignOnResource
}

// SubscriptionList - The request to update subscriptions needed to be monitored by the Datadog monitor resource.
type SubscriptionList struct {
	// List of subscriptions and the state of the monitoring.
	MonitoredSubscriptionList []*MonitoredSubscription

	// The operation for the patch on the resource.
	Operation *Operation
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserInfo - Includes name, email and optionally, phone number. User Information can't be null.
type UserInfo struct {
	// Email of the user used by Datadog for contacting them if needed
	EmailAddress *string

	// Name of the user
	Name *string

	// Phone number of the user used by Datadog for contacting them if needed
	PhoneNumber *string
}
