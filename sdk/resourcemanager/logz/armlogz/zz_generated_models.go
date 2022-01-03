//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows the OData
// error response format.).
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// The error object.
	InnerError *ErrorDetail `json:"error,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag. Exclusion takes priority over inclusion.
	Action *TagAction `json:"action,omitempty"`

	// The name (also known as the key) of the tag.
	Name *string `json:"name,omitempty"`

	// The value of the tag.
	Value *string `json:"value,omitempty"`
}

type IdentityProperties struct {
	Type *ManagedIdentityTypes `json:"type,omitempty"`

	// READ-ONLY; The identity ID.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all resources will be captured.
	// If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated
	// tags.
	FilteringTags []*FilteringTag `json:"filteringTags,omitempty"`

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *bool `json:"sendAadLogs,omitempty"`

	// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty"`

	// Flag specifying if subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type LogRules.
func (l LogRules) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "filteringTags", l.FilteringTags)
	populate(objectMap, "sendAadLogs", l.SendAADLogs)
	populate(objectMap, "sendActivityLogs", l.SendActivityLogs)
	populate(objectMap, "sendSubscriptionLogs", l.SendSubscriptionLogs)
	return json.Marshal(objectMap)
}

type LogzMonitorResource struct {
	// REQUIRED
	Location *string             `json:"location,omitempty"`
	Identity *IdentityProperties `json:"identity,omitempty"`

	// Properties specific to the monitor resource.
	Properties *MonitorProperties `json:"properties,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; ARM id of the monitor resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the monitor resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the monitor resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type LogzMonitorResource.
func (l LogzMonitorResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", l.ID)
	populate(objectMap, "identity", l.Identity)
	populate(objectMap, "location", l.Location)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "systemData", l.SystemData)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// LogzMonitorResourceListResponse - Response of a list operation.
type LogzMonitorResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*LogzMonitorResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type LogzMonitorResourceListResponse.
func (l LogzMonitorResourceListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// LogzMonitorResourceUpdateParameters - The parameters for a PATCH request to a monitor resource.
type LogzMonitorResourceUpdateParameters struct {
	// The set of properties that can be update in a PATCH request to a monitor resource.
	Properties *MonitorUpdateProperties `json:"properties,omitempty"`

	// The new tags of the monitor resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type LogzMonitorResourceUpdateParameters.
func (l LogzMonitorResourceUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "tags", l.Tags)
	return json.Marshal(objectMap)
}

type LogzOrganizationProperties struct {
	// Name of the Logz organization.
	CompanyName *string `json:"companyName,omitempty"`

	// The Id of the Enterprise App used for Single sign on.
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty"`

	// The login URL specific to this Logz Organization.
	SingleSignOnURL *string `json:"singleSignOnUrl,omitempty"`

	// READ-ONLY; Id of the Logz organization.
	ID *string `json:"id,omitempty" azure:"ro"`
}

type LogzSingleSignOnProperties struct {
	// The Id of the Enterprise App used for Single sign-on.
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty"`

	// Various states of the SSO resource
	SingleSignOnState *SingleSignOnStates `json:"singleSignOnState,omitempty"`

	// The login URL specific to this Logz Organization.
	SingleSignOnURL *string `json:"singleSignOnUrl,omitempty"`

	// READ-ONLY; Flag specifying if the resource provisioning state as tracked by ARM.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

type LogzSingleSignOnResource struct {
	Properties *LogzSingleSignOnProperties `json:"properties,omitempty"`

	// READ-ONLY; ARM id of the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the configuration.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// LogzSingleSignOnResourceListResponse - Response of a list operation.
type LogzSingleSignOnResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*LogzSingleSignOnResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type LogzSingleSignOnResourceListResponse.
func (l LogzSingleSignOnResourceListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MonitorListVMHostUpdateOptions contains the optional parameters for the Monitor.ListVMHostUpdate method.
type MonitorListVMHostUpdateOptions struct {
	// Request body to update the collection for agent installed in the given monitor.
	Body *VMHostUpdateRequest
}

// MonitorListVMHostsOptions contains the optional parameters for the Monitor.ListVMHosts method.
type MonitorListVMHostsOptions struct {
	// placeholder for future optional parameters
}

// MonitorProperties - Properties specific to the monitor resource.
type MonitorProperties struct {
	LogzOrganizationProperties *LogzOrganizationProperties `json:"logzOrganizationProperties,omitempty"`

	// Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in time, the resource will go in Suspended state.
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus `json:"marketplaceSubscriptionStatus,omitempty"`

	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus `json:"monitoringStatus,omitempty"`
	PlanData         *PlanData         `json:"planData,omitempty"`
	UserInfo         *UserInfo         `json:"userInfo,omitempty"`

	// READ-ONLY
	LiftrResourceCategory *LiftrResourceCategories `json:"liftrResourceCategory,omitempty" azure:"ro"`

	// READ-ONLY; The priority of the resource.
	LiftrResourcePreference *int32 `json:"liftrResourcePreference,omitempty" azure:"ro"`

	// READ-ONLY; Flag specifying if the resource provisioning state as tracked by ARM.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MonitorUpdateProperties - The set of properties that can be update in a PATCH request to a monitor resource.
type MonitorUpdateProperties struct {
	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus `json:"monitoringStatus,omitempty"`
}

// MonitorVMHostPayloadOptions contains the optional parameters for the Monitor.VMHostPayload method.
type MonitorVMHostPayloadOptions struct {
	// placeholder for future optional parameters
}

// MonitoredResource - The properties of a resource currently being monitored by the Logz monitor resource.
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string `json:"id,omitempty"`

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string `json:"reasonForLogsStatus,omitempty"`

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string `json:"reasonForMetricsStatus,omitempty"`

	// Flag indicating if resource is sending logs to Logz.
	SendingLogs *bool `json:"sendingLogs,omitempty"`

	// Flag indicating if resource is sending metrics to Logz.
	SendingMetrics *bool `json:"sendingMetrics,omitempty"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MonitoredResourceListResponse - Response of a list operation.
type MonitoredResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*MonitoredResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MonitoredResourceListResponse.
func (m MonitoredResourceListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MonitoringTagRules - Capture logs and metrics of Azure resources based on ARM tags.
type MonitoringTagRules struct {
	// Definition of the properties for a TagRules resource.
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

// MarshalJSON implements the json.Marshaller interface for type MonitoringTagRulesListResponse.
func (m MonitoringTagRulesListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MonitoringTagRulesProperties - Definition of the properties for a TagRules resource.
type MonitoringTagRulesProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules `json:"logRules,omitempty"`

	// READ-ONLY; Flag specifying if the resource provisioning state as tracked by ARM.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MonitorsBeginCreateOptions contains the optional parameters for the Monitors.BeginCreate method.
type MonitorsBeginCreateOptions struct {
	Body *LogzMonitorResource
}

// MonitorsBeginDeleteOptions contains the optional parameters for the Monitors.BeginDelete method.
type MonitorsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// MonitorsGetOptions contains the optional parameters for the Monitors.Get method.
type MonitorsGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsListByResourceGroupOptions contains the optional parameters for the Monitors.ListByResourceGroup method.
type MonitorsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsListBySubscriptionOptions contains the optional parameters for the Monitors.ListBySubscription method.
type MonitorsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// MonitorsListMonitoredResourcesOptions contains the optional parameters for the Monitors.ListMonitoredResources method.
type MonitorsListMonitoredResourcesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsListUserRolesOptions contains the optional parameters for the Monitors.ListUserRoles method.
type MonitorsListUserRolesOptions struct {
	Body *UserRoleRequest
}

// MonitorsUpdateOptions contains the optional parameters for the Monitors.Update method.
type MonitorsUpdateOptions struct {
	Body *LogzMonitorResourceUpdateParameters
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write monitors'.
	Description *string `json:"description,omitempty"`

	// Operation type, e.g., read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider, i.e., Microsoft.Logz.
	Provider *string `json:"provider,omitempty"`

	// Type on which the operation is performed, e.g., 'monitors'.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of GET request to list the Microsoft.Logz operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by the Microsoft.Logz provider.
	Value []*OperationResult `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationResult - A Microsoft.Logz REST API operation.
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

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

type PlanData struct {
	// different billing cycles like MONTHLY/WEEKLY. this could be enum
	BillingCycle *string `json:"billingCycle,omitempty"`

	// date when plan was applied
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// plan id as published by Logz
	PlanDetails *string `json:"planDetails,omitempty"`

	// different usage type like PAYG/COMMITTED. this could be enum
	UsageType *string `json:"usageType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PlanData.
func (p PlanData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "billingCycle", p.BillingCycle)
	populateTimeRFC3339(objectMap, "effectiveDate", p.EffectiveDate)
	populate(objectMap, "planDetails", p.PlanDetails)
	populate(objectMap, "usageType", p.UsageType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanData.
func (p *PlanData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "billingCycle":
			err = unpopulate(val, &p.BillingCycle)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, &p.EffectiveDate)
			delete(rawMsg, key)
		case "planDetails":
			err = unpopulate(val, &p.PlanDetails)
			delete(rawMsg, key)
		case "usageType":
			err = unpopulate(val, &p.UsageType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// SingleSignOnBeginCreateOrUpdateOptions contains the optional parameters for the SingleSignOn.BeginCreateOrUpdate method.
type SingleSignOnBeginCreateOrUpdateOptions struct {
	Body *LogzSingleSignOnResource
}

// SingleSignOnGetOptions contains the optional parameters for the SingleSignOn.Get method.
type SingleSignOnGetOptions struct {
	// placeholder for future optional parameters
}

// SingleSignOnListOptions contains the optional parameters for the SingleSignOn.List method.
type SingleSignOnListOptions struct {
	// placeholder for future optional parameters
}

// SubAccountBeginCreateOptions contains the optional parameters for the SubAccount.BeginCreate method.
type SubAccountBeginCreateOptions struct {
	Body *LogzMonitorResource
}

// SubAccountBeginDeleteOptions contains the optional parameters for the SubAccount.BeginDelete method.
type SubAccountBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// SubAccountGetOptions contains the optional parameters for the SubAccount.Get method.
type SubAccountGetOptions struct {
	// placeholder for future optional parameters
}

// SubAccountListMonitoredResourcesOptions contains the optional parameters for the SubAccount.ListMonitoredResources method.
type SubAccountListMonitoredResourcesOptions struct {
	// placeholder for future optional parameters
}

// SubAccountListOptions contains the optional parameters for the SubAccount.List method.
type SubAccountListOptions struct {
	// placeholder for future optional parameters
}

// SubAccountListVMHostUpdateOptions contains the optional parameters for the SubAccount.ListVMHostUpdate method.
type SubAccountListVMHostUpdateOptions struct {
	// Request body to update the collection for agent installed in the given monitor.
	Body *VMHostUpdateRequest
}

// SubAccountListVMHostsOptions contains the optional parameters for the SubAccount.ListVMHosts method.
type SubAccountListVMHostsOptions struct {
	// placeholder for future optional parameters
}

// SubAccountTagRulesCreateOrUpdateOptions contains the optional parameters for the SubAccountTagRules.CreateOrUpdate method.
type SubAccountTagRulesCreateOrUpdateOptions struct {
	Body *MonitoringTagRules
}

// SubAccountTagRulesDeleteOptions contains the optional parameters for the SubAccountTagRules.Delete method.
type SubAccountTagRulesDeleteOptions struct {
	// placeholder for future optional parameters
}

// SubAccountTagRulesGetOptions contains the optional parameters for the SubAccountTagRules.Get method.
type SubAccountTagRulesGetOptions struct {
	// placeholder for future optional parameters
}

// SubAccountTagRulesListOptions contains the optional parameters for the SubAccountTagRules.List method.
type SubAccountTagRulesListOptions struct {
	// placeholder for future optional parameters
}

// SubAccountUpdateOptions contains the optional parameters for the SubAccount.Update method.
type SubAccountUpdateOptions struct {
	Body *LogzMonitorResourceUpdateParameters
}

// SubAccountVMHostPayloadOptions contains the optional parameters for the SubAccount.VMHostPayload method.
type SubAccountVMHostPayloadOptions struct {
	// placeholder for future optional parameters
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

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TagRulesCreateOrUpdateOptions contains the optional parameters for the TagRules.CreateOrUpdate method.
type TagRulesCreateOrUpdateOptions struct {
	Body *MonitoringTagRules
}

// TagRulesDeleteOptions contains the optional parameters for the TagRules.Delete method.
type TagRulesDeleteOptions struct {
	// placeholder for future optional parameters
}

// TagRulesGetOptions contains the optional parameters for the TagRules.Get method.
type TagRulesGetOptions struct {
	// placeholder for future optional parameters
}

// TagRulesListOptions contains the optional parameters for the TagRules.List method.
type TagRulesListOptions struct {
	// placeholder for future optional parameters
}

type UserInfo struct {
	// Email of the user used by Logz for contacting them if needed
	EmailAddress *string `json:"emailAddress,omitempty"`

	// First Name of the user
	FirstName *string `json:"firstName,omitempty"`

	// Last Name of the user
	LastName *string `json:"lastName,omitempty"`

	// Phone number of the user used by Logz for contacting them if needed
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// UserRoleListResponse - Response for list of user's role for Logz.io account.
type UserRoleListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of user roles for Logz.io account.
	Value []*UserRoleResponse `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type UserRoleListResponse.
func (u UserRoleListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// UserRoleRequest - Request for checking user's role for Logz.io account.
type UserRoleRequest struct {
	// Email of the user used by Logz for contacting them if needed
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// UserRoleResponse - Response for checking user's role for Logz.io account.
type UserRoleResponse struct {
	// User roles on configured in Logz.io account.
	Role *UserRole `json:"role,omitempty"`
}

// VMExtensionPayload - Response of payload to be passed while installing VM agent.
type VMExtensionPayload struct {
	// API Key corresponding to the resource.
	APIKey *string `json:"apiKey,omitempty"`

	// Logz.io region where the resource has been created.
	Region *string `json:"region,omitempty"`
}

// VMHostUpdateRequest - Request of a list VM Host Update Operation.
type VMHostUpdateRequest struct {
	// Specifies the state of the operation - install/ delete.
	State *VMHostUpdateStates `json:"state,omitempty"`

	// Request of a list vm host update operation.
	VMResourceIDs []*VMResources `json:"vmResourceIds,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type VMHostUpdateRequest.
func (v VMHostUpdateRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "state", v.State)
	populate(objectMap, "vmResourceIds", v.VMResourceIDs)
	return json.Marshal(objectMap)
}

// VMResources - VM Resource Ids
type VMResources struct {
	// Version of the Logz agent installed on the VM.
	AgentVersion *string `json:"agentVersion,omitempty"`

	// Request of a list vm host update operation.
	ID *string `json:"id,omitempty"`
}

// VMResourcesListResponse - Response of a list VM Host Update Operation.
type VMResourcesListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Response of a list vm host update operation.
	Value []*VMResources `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type VMResourcesListResponse.
func (v VMResourcesListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
