//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armelastic

import "time"

// AllTrafficFiltersClientListOptions contains the optional parameters for the AllTrafficFiltersClient.List method.
type AllTrafficFiltersClientListOptions struct {
	// placeholder for future optional parameters
}

// AssociateTrafficFilterClientBeginAssociateOptions contains the optional parameters for the AssociateTrafficFilterClient.BeginAssociate
// method.
type AssociateTrafficFilterClientBeginAssociateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
	// Ruleset Id of the filter
	RulesetID *string
}

// CloudDeployment - Details of the user's elastic deployment associated with the monitor resource.
type CloudDeployment struct {
	// READ-ONLY; Associated Azure subscription Id for the elastic deployment.
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty" azure:"ro"`

	// READ-ONLY; Elastic deployment Id
	DeploymentID *string `json:"deploymentId,omitempty" azure:"ro"`

	// READ-ONLY; Region where Deployment at Elastic side took place.
	ElasticsearchRegion *string `json:"elasticsearchRegion,omitempty" azure:"ro"`

	// READ-ONLY; Elasticsearch ingestion endpoint of the Elastic deployment.
	ElasticsearchServiceURL *string `json:"elasticsearchServiceUrl,omitempty" azure:"ro"`

	// READ-ONLY; Kibana endpoint of the Elastic deployment.
	KibanaServiceURL *string `json:"kibanaServiceUrl,omitempty" azure:"ro"`

	// READ-ONLY; Kibana dashboard sso URL of the Elastic deployment.
	KibanaSsoURL *string `json:"kibanaSsoUrl,omitempty" azure:"ro"`

	// READ-ONLY; Elastic deployment name
	Name *string `json:"name,omitempty" azure:"ro"`
}

// CloudUser - Details of the user's elastic account.
type CloudUser struct {
	// READ-ONLY; Elastic cloud default dashboard sso URL of the Elastic user account.
	ElasticCloudSsoDefaultURL *string `json:"elasticCloudSsoDefaultUrl,omitempty" azure:"ro"`

	// READ-ONLY; Email of the Elastic User Account.
	EmailAddress *string `json:"emailAddress,omitempty" azure:"ro"`

	// READ-ONLY; User Id of the elastic account of the User.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// CompanyInfo - Company information of the user to be passed to partners.
type CompanyInfo struct {
	// Business of the company
	Business *string `json:"business,omitempty"`

	// Country of the company location.
	Country *string `json:"country,omitempty"`

	// Domain of the company
	Domain *string `json:"domain,omitempty"`

	// Number of employees in the company
	EmployeesNumber *string `json:"employeesNumber,omitempty"`

	// State of the company location.
	State *string `json:"state,omitempty"`
}

// CreateAndAssociateIPFilterClientBeginCreateOptions contains the optional parameters for the CreateAndAssociateIPFilterClient.BeginCreate
// method.
type CreateAndAssociateIPFilterClientBeginCreateOptions struct {
	// List of ips
	IPs *string
	// Name of the traffic filter
	Name *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CreateAndAssociatePLFilterClientBeginCreateOptions contains the optional parameters for the CreateAndAssociatePLFilterClient.BeginCreate
// method.
type CreateAndAssociatePLFilterClientBeginCreateOptions struct {
	// Name of the traffic filter
	Name *string
	// Guid of the private endpoint
	PrivateEndpointGUID *string
	// Name of the private endpoint
	PrivateEndpointName *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentInfoClientListOptions contains the optional parameters for the DeploymentInfoClient.List method.
type DeploymentInfoClientListOptions struct {
	// placeholder for future optional parameters
}

// DeploymentInfoResponse - The properties of deployment in Elastic cloud corresponding to the Elastic monitor resource.
type DeploymentInfoResponse struct {
	// READ-ONLY; Deployment URL of the elasticsearch in Elastic cloud deployment.
	DeploymentURL *string `json:"deploymentUrl,omitempty" azure:"ro"`

	// READ-ONLY; Disk capacity of the elasticsearch in Elastic cloud deployment.
	DiskCapacity *string `json:"diskCapacity,omitempty" azure:"ro"`

	// READ-ONLY; Marketplace SaaS Info of the resource.
	MarketplaceSaasInfo *MarketplaceSaaSInfo `json:"marketplaceSaasInfo,omitempty" azure:"ro"`

	// READ-ONLY; RAM capacity of the elasticsearch in Elastic cloud deployment.
	MemoryCapacity *string `json:"memoryCapacity,omitempty" azure:"ro"`

	// READ-ONLY; The Elastic deployment status.
	Status *ElasticDeploymentStatus `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; Version of the elasticsearch in Elastic cloud deployment.
	Version *string `json:"version,omitempty" azure:"ro"`
}

// DetachAndDeleteTrafficFilterClientDeleteOptions contains the optional parameters for the DetachAndDeleteTrafficFilterClient.Delete
// method.
type DetachAndDeleteTrafficFilterClientDeleteOptions struct {
	// Ruleset Id of the filter
	RulesetID *string
}

// DetachTrafficFilterClientBeginUpdateOptions contains the optional parameters for the DetachTrafficFilterClient.BeginUpdate
// method.
type DetachTrafficFilterClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
	// Ruleset Id of the filter
	RulesetID *string
}

// ExternalUserClientCreateOrUpdateOptions contains the optional parameters for the ExternalUserClient.CreateOrUpdate method.
type ExternalUserClientCreateOrUpdateOptions struct {
	// Elastic External User Creation Parameters
	Body *ExternalUserInfo
}

// ExternalUserCreationResponse - The properties of the response we got from elastic while creating external user
type ExternalUserCreationResponse struct {
	// READ-ONLY; Shows if user is created or updated
	Created *bool `json:"created,omitempty" azure:"ro"`
}

// ExternalUserInfo - The properties of the request required for creating user on elastic side
type ExternalUserInfo struct {
	// Email id of the user to be created or updated
	EmailID *string `json:"emailId,omitempty"`

	// Full name of the user to be created or updated
	FullName *string `json:"fullName,omitempty"`

	// Password of the user to be created or updated
	Password *string `json:"password,omitempty"`

	// Roles to be assigned for created or updated user
	Roles []*string `json:"roles,omitempty"`

	// Username of the user to be created or updated
	UserName *string `json:"userName,omitempty"`
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them
// from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag.
	Action *TagAction `json:"action,omitempty"`

	// The name (also known as the key) of the tag.
	Name *string `json:"name,omitempty"`

	// The value of the tag.
	Value *string `json:"value,omitempty"`
}

// IdentityProperties - Identity properties.
type IdentityProperties struct {
	// Managed identity type.
	Type *ManagedIdentityTypes `json:"type,omitempty"`

	// READ-ONLY; The identity ID.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// ListAssociatedTrafficFiltersClientListOptions contains the optional parameters for the ListAssociatedTrafficFiltersClient.List
// method.
type ListAssociatedTrafficFiltersClientListOptions struct {
	// placeholder for future optional parameters
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty,
	// all resources will be captured. If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
	// resources with the associated tags.
	FilteringTags []*FilteringTag `json:"filteringTags,omitempty"`

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *bool `json:"sendAadLogs,omitempty"`

	// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty"`

	// Flag specifying if subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty"`
}

// MarketplaceSaaSInfo - Marketplace SAAS Info of the resource.
type MarketplaceSaaSInfo struct {
	// Subscription Details: Marketplace SAAS Name
	MarketplaceName *string `json:"marketplaceName,omitempty"`

	// Subscription Details: Marketplace Resource URI
	MarketplaceResourceID *string `json:"marketplaceResourceId,omitempty"`

	// Marketplace Subscription Id
	MarketplaceSubscription *MarketplaceSaaSInfoMarketplaceSubscription `json:"marketplaceSubscription,omitempty"`
}

// MarketplaceSaaSInfoMarketplaceSubscription - Marketplace Subscription Id
type MarketplaceSaaSInfoMarketplaceSubscription struct {
	// Marketplace Subscription Id. This is a GUID-formatted string.
	ID *string `json:"id,omitempty"`
}

// MonitorClientBeginUpgradeOptions contains the optional parameters for the MonitorClient.BeginUpgrade method.
type MonitorClientBeginUpgradeOptions struct {
	// Elastic Monitor Upgrade Parameters
	Body *MonitorUpgrade
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorProperties - Properties specific to the monitor resource.
type MonitorProperties struct {
	// Elastic cloud properties.
	ElasticProperties *Properties `json:"elasticProperties,omitempty"`

	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus `json:"monitoringStatus,omitempty"`

	// Provisioning state of the monitor resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// User information.
	UserInfo *UserInfo `json:"userInfo,omitempty"`

	// Version of elastic of the monitor resource
	Version *string `json:"version,omitempty"`

	// READ-ONLY
	LiftrResourceCategory *LiftrResourceCategories `json:"liftrResourceCategory,omitempty" azure:"ro"`

	// READ-ONLY; The priority of the resource.
	LiftrResourcePreference *int32 `json:"liftrResourcePreference,omitempty" azure:"ro"`
}

// MonitorResource - Monitor resource.
type MonitorResource struct {
	// REQUIRED; The location of the monitor resource
	Location *string `json:"location,omitempty"`

	// Identity properties of the monitor resource.
	Identity *IdentityProperties `json:"identity,omitempty"`

	// Properties of the monitor resource.
	Properties *MonitorProperties `json:"properties,omitempty"`

	// SKU of the monitor resource.
	SKU *ResourceSKU `json:"sku,omitempty"`

	// The tags of the monitor resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Flag to determine if User API Key has to be generated and shared.
	GenerateAPIKey *bool `json:"generateApiKey,omitempty" azure:"ro"`

	// READ-ONLY; ARM id of the monitor resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the monitor resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the monitor resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MonitorResourceListResponse - Response of a list operation.
type MonitorResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*MonitorResource `json:"value,omitempty"`
}

// MonitorResourceUpdateParameters - Monitor resource update parameters.
type MonitorResourceUpdateParameters struct {
	// elastic monitor resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MonitorUpgrade - Upgrade elastic monitor version
type MonitorUpgrade struct {
	// Version to which the elastic monitor should be upgraded to
	Version *string `json:"version,omitempty"`
}

// MonitoredResource - The properties of a resource currently being monitored by the Elastic monitor resource.
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string `json:"id,omitempty"`

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string `json:"reasonForLogsStatus,omitempty"`

	// Flag indicating the status of the resource for sending logs operation to Elastic.
	SendingLogs *SendingLogs `json:"sendingLogs,omitempty"`
}

// MonitoredResourceListResponse - Response of a list operation.
type MonitoredResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*MonitoredResource `json:"value,omitempty"`
}

// MonitoredResourcesClientListOptions contains the optional parameters for the MonitoredResourcesClient.NewListPager method.
type MonitoredResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitoringTagRules - Capture logs and metrics of Azure resources based on ARM tags.
type MonitoringTagRules struct {
	// Properties of the monitoring tag rules.
	Properties *MonitoringTagRulesProperties `json:"properties,omitempty"`

	// READ-ONLY; The id of the rule set.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the rule set.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the rule set.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MonitoringTagRulesListResponse - Response of a list operation.
type MonitoringTagRulesListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*MonitoringTagRules `json:"value,omitempty"`
}

// MonitoringTagRulesProperties - Definition of the properties for a TagRules resource.
type MonitoringTagRulesProperties struct {
	// Rules for sending logs.
	LogRules *LogRules `json:"logRules,omitempty"`

	// Provisioning state of the monitoring tag rules.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

// MonitorsClientBeginCreateOptions contains the optional parameters for the MonitorsClient.BeginCreate method.
type MonitorsClientBeginCreateOptions struct {
	// Elastic monitor resource model
	Body *MonitorResource
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
type MonitorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
type MonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.NewListByResourceGroupPager
// method.
type MonitorsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListOptions contains the optional parameters for the MonitorsClient.NewListPager method.
type MonitorsClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientUpdateOptions contains the optional parameters for the MonitorsClient.Update method.
type MonitorsClientUpdateOptions struct {
	// Elastic resource model update parameters.
	Body *MonitorResourceUpdateParameters
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write monitors'.
	Description *string `json:"description,omitempty"`

	// Operation type, e.g., read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider, i.e., Microsoft.Elastic.
	Provider *string `json:"provider,omitempty"`

	// Type on which the operation is performed, e.g., 'monitors'.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of GET request to list the Microsoft.Elastic operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by the Microsoft.Elastic provider.
	Value []*OperationResult `json:"value,omitempty"`
}

// OperationResult - A Microsoft.Elastic REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name, i.e., {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`

	// Origin of the operation
	Origin *string `json:"origin,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OrganizationsClientGetAPIKeyOptions contains the optional parameters for the OrganizationsClient.GetAPIKey method.
type OrganizationsClientGetAPIKeyOptions struct {
	// Email Id parameter of the User Organization, of which the API Key must be returned
	Body *UserEmailID
}

// Properties - Elastic Resource Properties.
type Properties struct {
	// Details of the elastic cloud deployment.
	ElasticCloudDeployment *CloudDeployment `json:"elasticCloudDeployment,omitempty"`

	// Details of the user's elastic account.
	ElasticCloudUser *CloudUser `json:"elasticCloudUser,omitempty"`
}

// ResourceSKU - Microsoft.Elastic SKU.
type ResourceSKU struct {
	// REQUIRED; Name of the SKU.
	Name *string `json:"name,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TagRulesClientBeginDeleteOptions contains the optional parameters for the TagRulesClient.BeginDelete method.
type TagRulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TagRulesClientCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.CreateOrUpdate method.
type TagRulesClientCreateOrUpdateOptions struct {
	// request body of MonitoringTagRules
	Body *MonitoringTagRules
}

// TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
type TagRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientListOptions contains the optional parameters for the TagRulesClient.NewListPager method.
type TagRulesClientListOptions struct {
	// placeholder for future optional parameters
}

// TrafficFilter - Elastic traffic filter object
type TrafficFilter struct {
	// Description of the elastic filter
	Description *string `json:"description,omitempty"`

	// Id of the elastic filter
	ID *string `json:"id,omitempty"`

	// IncludeByDefault for the elastic filter
	IncludeByDefault *bool `json:"includeByDefault,omitempty"`

	// Name of the elastic filter
	Name *string `json:"name,omitempty"`

	// Region of the elastic filter
	Region *string `json:"region,omitempty"`

	// Rules in the elastic filter
	Rules []*TrafficFilterRule `json:"rules,omitempty"`

	// Type of the elastic filter
	Type *Type `json:"type,omitempty"`
}

// TrafficFilterResponse - List of elastic traffic filters in the account
type TrafficFilterResponse struct {
	// List of elastic traffic filters in the account
	Rulesets []*TrafficFilter `json:"rulesets,omitempty"`
}

// TrafficFilterRule - Elastic traffic filter rule object
type TrafficFilterRule struct {
	// Guid of Private Endpoint in the elastic filter rule
	AzureEndpointGUID *string `json:"azureEndpointGuid,omitempty"`

	// Name of the Private Endpoint in the elastic filter rule
	AzureEndpointName *string `json:"azureEndpointName,omitempty"`

	// Description of the elastic filter rule
	Description *string `json:"description,omitempty"`

	// Id of the elastic filter rule
	ID *string `json:"id,omitempty"`

	// IP of the elastic filter rule
	Source *string `json:"source,omitempty"`
}

// TrafficFiltersClientDeleteOptions contains the optional parameters for the TrafficFiltersClient.Delete method.
type TrafficFiltersClientDeleteOptions struct {
	// Ruleset Id of the filter
	RulesetID *string
}

// UpgradableVersionsClientDetailsOptions contains the optional parameters for the UpgradableVersionsClient.Details method.
type UpgradableVersionsClientDetailsOptions struct {
	// placeholder for future optional parameters
}

// UpgradableVersionsList - Stack Versions that this version can upgrade to
type UpgradableVersionsList struct {
	// Current version of the elastic monitor
	CurrentVersion *string `json:"currentVersion,omitempty"`

	// Stack Versions that this version can upgrade to
	UpgradableVersions []*string `json:"upgradableVersions,omitempty"`
}

// UserAPIKeyResponse - The User Api Key created for the Organization associated with the User Email Id that was passed in
// the request
type UserAPIKeyResponse struct {
	// The User Api Key Generated based on ReturnApiKey flag. This is applicable for non-Portal clients only.
	APIKey *string `json:"apiKey,omitempty"`
}

// UserEmailID - Email Id of the User Organization, of which the API Key must be returned
type UserEmailID struct {
	// The User email Id
	EmailID *string `json:"emailId,omitempty"`
}

// UserInfo - User Information to be passed to partners.
type UserInfo struct {
	// Company information of the user to be passed to partners.
	CompanyInfo *CompanyInfo `json:"companyInfo,omitempty"`

	// Company name of the user
	CompanyName *string `json:"companyName,omitempty"`

	// Email of the user used by Elastic for contacting them if needed
	EmailAddress *string `json:"emailAddress,omitempty"`

	// First name of the user
	FirstName *string `json:"firstName,omitempty"`

	// Last name of the user
	LastName *string `json:"lastName,omitempty"`
}

// VMCollectionClientUpdateOptions contains the optional parameters for the VMCollectionClient.Update method.
type VMCollectionClientUpdateOptions struct {
	// VM resource Id
	Body *VMCollectionUpdate
}

// VMCollectionUpdate - Update VM resource collection.
type VMCollectionUpdate struct {
	// Operation to be performed for given VM.
	OperationName *OperationName `json:"operationName,omitempty"`

	// ARM id of the VM resource.
	VMResourceID *string `json:"vmResourceId,omitempty"`
}

// VMHostClientListOptions contains the optional parameters for the VMHostClient.NewListPager method.
type VMHostClientListOptions struct {
	// placeholder for future optional parameters
}

// VMHostListResponse - Response of a list operation.
type VMHostListResponse struct {
	// Link to the next Vm resource Id, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*VMResources `json:"value,omitempty"`
}

// VMIngestionClientDetailsOptions contains the optional parameters for the VMIngestionClient.Details method.
type VMIngestionClientDetailsOptions struct {
	// placeholder for future optional parameters
}

// VMIngestionDetailsResponse - The vm ingestion details to install an agent.
type VMIngestionDetailsResponse struct {
	// The cloudId of given Elastic monitor resource.
	CloudID *string `json:"cloudId,omitempty"`

	// Ingestion details to install agent on given VM.
	IngestionKey *string `json:"ingestionKey,omitempty"`
}

// VMResources - The vm resource properties that is currently being monitored by the Elastic monitor resource.
type VMResources struct {
	// The ARM id of the VM resource.
	VMResourceID *string `json:"vmResourceId,omitempty"`
}
