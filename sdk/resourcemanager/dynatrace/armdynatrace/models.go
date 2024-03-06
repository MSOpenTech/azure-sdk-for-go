//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdynatrace

import "time"

// AccountInfo - Dynatrace Account Information
type AccountInfo struct {
	// Account Id of the account this environment is linked to
	AccountID *string

	// Region in which the account is created
	RegionID *string
}

// AppServiceInfo - Details of App Services having Dynatrace OneAgent installed
type AppServiceInfo struct {
	// Update settings of OneAgent.
	AutoUpdateSetting *AutoUpdateSetting

	// The availability state of OneAgent.
	AvailabilityState *AvailabilityState

	// The name of the host group
	HostGroup *string

	// The name of the host
	HostName *string

	// Tells whether log modules are enabled or not
	LogModule *LogModule

	// The monitoring mode of OneAgent
	MonitoringType *MonitoringType

	// App service resource ID
	ResourceID *string

	// The current update status of OneAgent.
	UpdateStatus *UpdateStatus

	// Version of the Dynatrace agent installed on the App Service.
	Version *string
}

// AppServiceListResponse - Response of a list App Services Operation.
type AppServiceListResponse struct {
	// The link to the next page of items
	NextLink *string

	// The items on this page
	Value []*AppServiceInfo
}

// EnvironmentInfo - Dynatrace Environment Information
type EnvironmentInfo struct {
	// Id of the environment created
	EnvironmentID *string

	// Ingestion key of the environment
	IngestionKey *string

	// Landing URL for Dynatrace environment
	LandingURL *string

	// Ingestion endpoint used for sending logs
	LogsIngestionEndpoint *string
}

// EnvironmentProperties - Properties of the Dynatrace environment.
type EnvironmentProperties struct {
	// Dynatrace Account Information
	AccountInfo *AccountInfo

	// Dynatrace Environment Information
	EnvironmentInfo *EnvironmentInfo

	// The details of a Dynatrace single sign-on.
	SingleSignOnProperties *SingleSignOnProperties

	// User id
	UserID *string
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

// IdentityProperties - The properties of the managed service identities assigned to this resource.
type IdentityProperties struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedIdentityType

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The active directory identifier of this principal.
	PrincipalID *string

	// READ-ONLY; The Active Directory tenant id of the principal.
	TenantID *string
}

// LinkableEnvironmentListResponse - Response for getting all the linkable environments
type LinkableEnvironmentListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// List of environments for which user is an admin
	Value []*LinkableEnvironmentResponse
}

// LinkableEnvironmentRequest - Request for getting all the linkable environments for a user
type LinkableEnvironmentRequest struct {
	// REQUIRED; Azure region in which we want to link the environment
	Region *string

	// REQUIRED; Tenant Id of the user in which they want to link the environment
	TenantID *string

	// REQUIRED; user principal id of the user
	UserPrincipal *string
}

// LinkableEnvironmentResponse - Response for getting all the linkable environments
type LinkableEnvironmentResponse struct {
	// environment id for which user is an admin
	EnvironmentID *string

	// Name of the environment
	EnvironmentName *string

	// Billing plan information.
	PlanData *PlanData
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty,
	// all resources will be captured. If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
	// resources with the associated tags.
	FilteringTags []*FilteringTag

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *SendAADLogsStatus

	// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
	SendActivityLogs *SendActivityLogsStatus

	// Flag specifying if subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *SendSubscriptionLogsStatus
}

// MarketplaceSaaSResourceDetailsRequest - Request for getting Marketplace SaaS resource details for a tenant Id
type MarketplaceSaaSResourceDetailsRequest struct {
	// REQUIRED; Tenant Id
	TenantID *string
}

// MarketplaceSaaSResourceDetailsResponse - Marketplace SaaS resource details linked to the given tenant Id
type MarketplaceSaaSResourceDetailsResponse struct {
	// Id of the Marketplace SaaS Resource
	MarketplaceSaaSResourceID *string

	// Marketplace subscription status
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus

	// Id of the plan
	PlanID *string
}

// MetricRules - Set of rules for sending metrics for the Monitor resource.
type MetricRules struct {
	// List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action
	// is specified, the rules will apply to the list of all available resources. If
	// Include actions are specified, the rules will only include resources with the associated tags.
	FilteringTags []*FilteringTag

	// Flag specifying if metrics from Azure resources should be sent for the Monitor resource.
	SendingMetrics *SendingMetricsStatus
}

// MetricsStatusResponse - Response of get metrics status operation
type MetricsStatusResponse struct {
	// Azure resource IDs
	AzureResourceIDs []*string
}

// MonitorProperties - Properties specific to the monitor resource.
type MonitorProperties struct {
	// Properties of the Dynatrace environment.
	DynatraceEnvironmentProperties *EnvironmentProperties

	// Marketplace subscription status.
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus

	// Status of the monitor.
	MonitoringStatus *MonitoringStatus

	// Billing plan information.
	PlanData *PlanData

	// User info.
	UserInfo *UserInfo

	// READ-ONLY; Liftr Resource category.
	LiftrResourceCategory *LiftrResourceCategories

	// READ-ONLY; The priority of the resource.
	LiftrResourcePreference *int32

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// MonitorResource - Dynatrace Monitor Resource
type MonitorResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitorProperties

	// The managed service identities assigned to this resource.
	Identity *IdentityProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; System metadata for this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// MonitorResourceListResult - The response of a MonitorResource list operation.
type MonitorResourceListResult struct {
	// REQUIRED; The items on this page
	Value []*MonitorResource

	// The link to the next page of items
	NextLink *string
}

// MonitorResourceUpdate - The updatable properties of the MonitorResource.
type MonitorResourceUpdate struct {
	// Resource tags.
	Tags map[string]*string
}

// MonitoredResource - Details of resource being monitored by Dynatrace monitor resource
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string

	// Flag indicating if resource is sending logs to Dynatrace.
	SendingLogs *SendingLogsStatus

	// Flag indicating if resource is sending metrics to Dynatrace.
	SendingMetrics *SendingMetricsStatus
}

// MonitoredResourceListResponse - List of all the resources being monitored by Dynatrace monitor resource
type MonitoredResourceListResponse struct {
	// The link to the next page of items
	NextLink *string

	// The items on this page
	Value []*MonitoredResource
}

// MonitoringTagRulesProperties - Properties for the Tag rules resource of a Monitor account.
type MonitoringTagRulesProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// PlanData - Billing plan information.
type PlanData struct {
	// different billing cycles like MONTHLY/WEEKLY. this could be enum
	BillingCycle *string

	// date when plan was applied
	EffectiveDate *time.Time

	// plan id as published by Dynatrace
	PlanDetails *string

	// different usage type like PAYG/COMMITTED. this could be enum
	UsageType *string
}

// SSODetailsRequest - Request for getting sso details for a user
type SSODetailsRequest struct {
	// REQUIRED; user principal id of the user
	UserPrincipal *string
}

// SSODetailsResponse - SSO details from the Dynatrace partner
type SSODetailsResponse struct {
	// array of Aad(azure active directory) domains
	AADDomains []*string

	// Array of admin user emails.
	AdminUsers []*string

	// Whether the SSO is enabled for this resource or not.
	IsSsoEnabled *SSOStatus

	// URL for Azure AD metadata
	MetadataURL *string

	// The login URL specific to this Dynatrace Environment
	SingleSignOnURL *string
}

// SingleSignOnProperties - The details of a Dynatrace single sign-on.
type SingleSignOnProperties struct {
	// array of Aad(azure active directory) domains
	AADDomains []*string

	// Version of the Dynatrace agent installed on the VM.
	EnterpriseAppID *string

	// State of Single Sign On
	SingleSignOnState *SingleSignOnStates

	// The login URL specific to this Dynatrace Environment
	SingleSignOnURL *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// SingleSignOnResource - Single sign-on configurations for a given monitor resource.
type SingleSignOnResource struct {
	// REQUIRED; The resource-specific properties for this resource.
	Properties *SingleSignOnProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; System metadata for this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SingleSignOnResourceListResult - The response of a DynatraceSingleSignOnResource list operation.
type SingleSignOnResourceListResult struct {
	// REQUIRED; The items on this page
	Value []*SingleSignOnResource

	// The link to the next page of items
	NextLink *string
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

// TagRule - Tag rules for a monitor resource
type TagRule struct {
	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitoringTagRulesProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; System metadata for this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TagRuleListResult - The response of a TagRule list operation.
type TagRuleListResult struct {
	// REQUIRED; The items on this page
	Value []*TagRule

	// The link to the next page of items
	NextLink *string
}

// UserAssignedIdentity - A managed identity assigned by the user.
type UserAssignedIdentity struct {
	// REQUIRED; The active directory client identifier for this principal.
	ClientID *string

	// REQUIRED; The active directory identifier for this principal.
	PrincipalID *string
}

// UserInfo - User info.
type UserInfo struct {
	// Country of the user
	Country *string

	// Email of the user used by Dynatrace for contacting them if needed
	EmailAddress *string

	// First Name of the user
	FirstName *string

	// Last Name of the user
	LastName *string

	// Phone number of the user used by Dynatrace for contacting them if needed
	PhoneNumber *string
}

// VMExtensionPayload - Response of payload to be passed while installing VM agent.
type VMExtensionPayload struct {
	// Id of the environment created
	EnvironmentID *string

	// Ingestion key of the environment
	IngestionKey *string
}

// VMHostsListResponse - Response of a list VM Host Operation.
type VMHostsListResponse struct {
	// The link to the next page of items
	NextLink *string

	// The items on this page
	Value []*VMInfo
}

// VMInfo - Details of VM Resource having Dynatrace OneAgent installed
type VMInfo struct {
	// Update settings of OneAgent.
	AutoUpdateSetting *AutoUpdateSetting

	// The availability state of OneAgent.
	AvailabilityState *AvailabilityState

	// The name of the host group
	HostGroup *string

	// The name of the host
	HostName *string

	// Tells whether log modules are enabled or not
	LogModule *LogModule

	// The monitoring mode of OneAgent
	MonitoringType *MonitoringType

	// Azure VM resource ID
	ResourceID *string

	// The current update status of OneAgent.
	UpdateStatus *UpdateStatus

	// Version of the Dynatrace agent installed on the VM.
	Version *string
}
