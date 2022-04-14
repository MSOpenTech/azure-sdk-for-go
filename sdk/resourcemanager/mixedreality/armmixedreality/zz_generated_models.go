//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

import "time"

// AccountKeyRegenerateRequest - Request for account key regeneration
type AccountKeyRegenerateRequest struct {
	// Serial of key to be regenerated
	Serial *Serial `json:"serial,omitempty"`
}

// AccountKeys - Developer Keys of account
type AccountKeys struct {
	// READ-ONLY; value of primary key.
	PrimaryKey *string `json:"primaryKey,omitempty" azure:"ro"`

	// READ-ONLY; value of secondary key.
	SecondaryKey *string `json:"secondaryKey,omitempty" azure:"ro"`
}

// AccountProperties - Common Properties shared by Mixed Reality Accounts
type AccountProperties struct {
	// The name of the storage account associated with this accountId
	StorageAccountName *string `json:"storageAccountName,omitempty"`

	// READ-ONLY; Correspond domain name of certain Spatial Anchors Account
	AccountDomain *string `json:"accountDomain,omitempty" azure:"ro"`

	// READ-ONLY; unique id of certain account.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`
}

// CheckNameAvailabilityRequest - Check Name Availability Request
type CheckNameAvailabilityRequest struct {
	// REQUIRED; Resource Name To Verify
	Name *string `json:"name,omitempty"`

	// REQUIRED; Fully qualified resource type which includes provider namespace
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - Check Name Availability Response
type CheckNameAvailabilityResponse struct {
	// REQUIRED; if name Available
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// detail message
	Message *string `json:"message,omitempty"`

	// Resource Name To Verify
	Reason *NameUnavailableReason `json:"reason,omitempty"`
}

// ClientCheckNameAvailabilityLocalOptions contains the optional parameters for the Client.CheckNameAvailabilityLocal method.
type ClientCheckNameAvailabilityLocalOptions struct {
	// placeholder for future optional parameters
}

// CloudError - An Error response.
type CloudError struct {
	// An Error response.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - An error response from Azure.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for displaying in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// LogSpecification - Specifications of the Log for Azure Monitoring
type LogSpecification struct {
	// Blob duration of the log
	BlobDuration *string `json:"blobDuration,omitempty"`

	// Localized friendly display name of the log
	DisplayName *string `json:"displayName,omitempty"`

	// Name of the log
	Name *string `json:"name,omitempty"`
}

// MetricDimension - Specifications of the Dimension of metrics
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string `json:"displayName,omitempty"`

	// Internal name of the dimension.
	InternalName *string `json:"internalName,omitempty"`

	// Name of the dimension
	Name *string `json:"name,omitempty"`

	// Flag to indicate export for Shoebox
	ToBeExportedForShoebox *bool `json:"toBeExportedForShoebox,omitempty"`
}

// MetricSpecification - Specifications of the Metrics for Azure Monitoring
type MetricSpecification struct {
	// Only provide one value for this field. Valid values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string `json:"aggregationType,omitempty"`

	// Metric category
	Category *string `json:"category,omitempty"`

	// Dimensions of the metric
	Dimensions []*MetricDimension `json:"dimensions,omitempty"`

	// Localized friendly description of the metric
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// Localized friendly display name of the metric
	DisplayName *string `json:"displayName,omitempty"`

	// Flag to indicate use of regional Mdm accounts
	EnableRegionalMdmAccount *bool `json:"enableRegionalMdmAccount,omitempty"`

	// Flag to determine is Zero is returned for time duration where no metric is emitted
	FillGapWithZero *bool `json:"fillGapWithZero,omitempty"`

	// Internal metric name.
	InternalMetricName *string `json:"internalMetricName,omitempty"`

	// Locked aggregation type of the metric
	LockedAggregationType *string `json:"lockedAggregationType,omitempty"`

	// Metric filter regex pattern
	MetricFilterPattern *string `json:"metricFilterPattern,omitempty"`

	// Name of the metric
	Name *string `json:"name,omitempty"`

	// Source mdm account
	SourceMdmAccount *string `json:"sourceMdmAccount,omitempty"`

	// Source mdm namespace
	SourceMdmNamespace *string `json:"sourceMdmNamespace,omitempty"`

	// Supported aggregation types. Valid values: Average, Minimum, Maximum, Total, Count.
	SupportedAggregationTypes []*string `json:"supportedAggregationTypes,omitempty"`

	// Supported time grains. Valid values: PT1M, PT5M, PT15M, PT30M, PT1H, PT6H, PT12H, P1D
	SupportedTimeGrainTypes []*string `json:"supportedTimeGrainTypes,omitempty"`

	// Unit that makes sense for the metric
	Unit *string `json:"unit,omitempty"`
}

// ObjectAnchorsAccount Response.
type ObjectAnchorsAccount struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string                       `json:"location,omitempty"`
	Identity *ObjectAnchorsAccountIdentity `json:"identity,omitempty"`

	// The kind of account, if supported
	Kind *SKU `json:"kind,omitempty"`

	// The plan associated with this account
	Plan *Identity `json:"plan,omitempty"`

	// Property bag.
	Properties *AccountProperties `json:"properties,omitempty"`

	// The sku associated with this account
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system metadata related to an object anchors account.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

type ObjectAnchorsAccountIdentity struct {
	// The identity type.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// ObjectAnchorsAccountPage - Result of the request to get resource collection. It contains a list of resources and a URL
// link to get the next set of results.
type ObjectAnchorsAccountPage struct {
	// URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources supported by the Resource Provider.
	Value []*ObjectAnchorsAccount `json:"value,omitempty"`
}

// ObjectAnchorsAccountsClientCreateOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Create method.
type ObjectAnchorsAccountsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientDeleteOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Delete method.
type ObjectAnchorsAccountsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientGetOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Get method.
type ObjectAnchorsAccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientListByResourceGroupOptions contains the optional parameters for the ObjectAnchorsAccountsClient.ListByResourceGroup
// method.
type ObjectAnchorsAccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientListBySubscriptionOptions contains the optional parameters for the ObjectAnchorsAccountsClient.ListBySubscription
// method.
type ObjectAnchorsAccountsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientListKeysOptions contains the optional parameters for the ObjectAnchorsAccountsClient.ListKeys
// method.
type ObjectAnchorsAccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientRegenerateKeysOptions contains the optional parameters for the ObjectAnchorsAccountsClient.RegenerateKeys
// method.
type ObjectAnchorsAccountsClientRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsClientUpdateOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Update method.
type ObjectAnchorsAccountsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Operation - REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Whether or not this is a data plane operation
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// The origin
	Origin *string `json:"origin,omitempty"`

	// Properties of the operation
	Properties *OperationProperties `json:"properties,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// REQUIRED; Description of operation
	Description *string `json:"description,omitempty"`

	// REQUIRED; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// REQUIRED; Service provider: Microsoft.ResourceProvider
	Provider *string `json:"provider,omitempty"`

	// REQUIRED; Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationPage - Result of the request to list Resource Provider operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationPage struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by the Resource Provider.
	Value []*Operation `json:"value,omitempty"`
}

// OperationProperties - Operation properties.
type OperationProperties struct {
	// Service specification.
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccount Response.
type RemoteRenderingAccount struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The identity associated with this account
	Identity *Identity `json:"identity,omitempty"`

	// The kind of account, if supported
	Kind *SKU `json:"kind,omitempty"`

	// The plan associated with this account
	Plan *Identity `json:"plan,omitempty"`

	// Property bag.
	Properties *AccountProperties `json:"properties,omitempty"`

	// The sku associated with this account
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; System metadata for this account
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RemoteRenderingAccountPage - Result of the request to get resource collection. It contains a list of resources and a URL
// link to get the next set of results.
type RemoteRenderingAccountPage struct {
	// URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources supported by the Resource Provider.
	Value []*RemoteRenderingAccount `json:"value,omitempty"`
}

// RemoteRenderingAccountsClientCreateOptions contains the optional parameters for the RemoteRenderingAccountsClient.Create
// method.
type RemoteRenderingAccountsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientDeleteOptions contains the optional parameters for the RemoteRenderingAccountsClient.Delete
// method.
type RemoteRenderingAccountsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientGetOptions contains the optional parameters for the RemoteRenderingAccountsClient.Get method.
type RemoteRenderingAccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientListByResourceGroupOptions contains the optional parameters for the RemoteRenderingAccountsClient.ListByResourceGroup
// method.
type RemoteRenderingAccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientListBySubscriptionOptions contains the optional parameters for the RemoteRenderingAccountsClient.ListBySubscription
// method.
type RemoteRenderingAccountsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientListKeysOptions contains the optional parameters for the RemoteRenderingAccountsClient.ListKeys
// method.
type RemoteRenderingAccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientRegenerateKeysOptions contains the optional parameters for the RemoteRenderingAccountsClient.RegenerateKeys
// method.
type RemoteRenderingAccountsClientRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsClientUpdateOptions contains the optional parameters for the RemoteRenderingAccountsClient.Update
// method.
type RemoteRenderingAccountsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SKU - The resource model definition representing SKU
type SKU struct {
	// REQUIRED; The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`

	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the
	// resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`

	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`

	// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required
	// on a PUT.
	Tier *SKUTier `json:"tier,omitempty"`
}

// ServiceSpecification - Service specification payload
type ServiceSpecification struct {
	// Specifications of the Log for Azure Monitoring
	LogSpecifications []*LogSpecification `json:"logSpecifications,omitempty"`

	// Specifications of the Metrics for Azure Monitoring
	MetricSpecifications []*MetricSpecification `json:"metricSpecifications,omitempty"`
}

// SpatialAnchorsAccount Response.
type SpatialAnchorsAccount struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The identity associated with this account
	Identity *Identity `json:"identity,omitempty"`

	// The kind of account, if supported
	Kind *SKU `json:"kind,omitempty"`

	// The plan associated with this account
	Plan *Identity `json:"plan,omitempty"`

	// Property bag.
	Properties *AccountProperties `json:"properties,omitempty"`

	// The sku associated with this account
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; System metadata for this account
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SpatialAnchorsAccountPage - Result of the request to get resource collection. It contains a list of resources and a URL
// link to get the next set of results.
type SpatialAnchorsAccountPage struct {
	// URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources supported by the Resource Provider.
	Value []*SpatialAnchorsAccount `json:"value,omitempty"`
}

// SpatialAnchorsAccountsClientCreateOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Create
// method.
type SpatialAnchorsAccountsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientDeleteOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Delete
// method.
type SpatialAnchorsAccountsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientGetOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Get method.
type SpatialAnchorsAccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientListByResourceGroupOptions contains the optional parameters for the SpatialAnchorsAccountsClient.ListByResourceGroup
// method.
type SpatialAnchorsAccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientListBySubscriptionOptions contains the optional parameters for the SpatialAnchorsAccountsClient.ListBySubscription
// method.
type SpatialAnchorsAccountsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientListKeysOptions contains the optional parameters for the SpatialAnchorsAccountsClient.ListKeys
// method.
type SpatialAnchorsAccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientRegenerateKeysOptions contains the optional parameters for the SpatialAnchorsAccountsClient.RegenerateKeys
// method.
type SpatialAnchorsAccountsClientRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsClientUpdateOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Update
// method.
type SpatialAnchorsAccountsClientUpdateOptions struct {
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

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}
