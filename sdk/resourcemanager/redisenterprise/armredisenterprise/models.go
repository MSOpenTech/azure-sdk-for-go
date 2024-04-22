//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredisenterprise

// AccessKeys - The secret access keys used for authenticating connections to redis
type AccessKeys struct {
	// READ-ONLY; The current primary key that clients can use to authenticate
	PrimaryKey *string

	// READ-ONLY; The current secondary key that clients can use to authenticate
	SecondaryKey *string
}

// Cluster - Describes the RedisEnterprise cluster
type Cluster struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The SKU to create, which affects price, performance, and features.
	SKU *SKU

	// The identity of the resource.
	Identity *ManagedServiceIdentity

	// Other properties of the cluster.
	Properties *ClusterProperties

	// Resource tags.
	Tags map[string]*string

	// The Availability Zones where this cluster will be deployed.
	Zones []*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ClusterList - The response of a list-all operation.
type ClusterList struct {
	// List of clusters.
	Value []*Cluster

	// READ-ONLY; The URI to fetch the next page of results.
	NextLink *string
}

// ClusterProperties - Properties of RedisEnterprise clusters, as opposed to general resource properties like location, tags
type ClusterProperties struct {
	// Encryption-at-rest configuration for the cluster.
	Encryption *ClusterPropertiesEncryption

	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTLSVersion *TLSVersion

	// READ-ONLY; DNS name of the cluster endpoint
	HostName *string

	// READ-ONLY; List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; Current provisioning status of the cluster
	ProvisioningState *ProvisioningState

	// READ-ONLY; Version of redis the cluster supports, e.g. '6'
	RedisVersion *string

	// READ-ONLY; Current resource status of the cluster
	ResourceState *ResourceState
}

// ClusterPropertiesEncryption - Encryption-at-rest configuration for the cluster.
type ClusterPropertiesEncryption struct {
	// All Customer-managed key encryption properties for the resource. Set this to an empty object to use Microsoft-managed key
	// encryption.
	CustomerManagedKeyEncryption *ClusterPropertiesEncryptionCustomerManagedKeyEncryption
}

// ClusterPropertiesEncryptionCustomerManagedKeyEncryption - All Customer-managed key encryption properties for the resource.
// Set this to an empty object to use Microsoft-managed key encryption.
type ClusterPropertiesEncryptionCustomerManagedKeyEncryption struct {
	// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
	KeyEncryptionKeyIdentity *ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity

	// Key encryption key Url, versioned only. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
	KeyEncryptionKeyURL *string
}

// ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity - All identity configuration for Customer-managed key
// settings defining which identity should be used to auth to Key Vault.
type ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity struct {
	// Only userAssignedIdentity is supported in this API version; other types may be supported in the future
	IdentityType *CmkIdentityType

	// User assigned identity to use for accessing key encryption key Url. Ex: /subscriptions//resourceGroups//providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
	UserAssignedIdentityResourceID *string
}

// ClusterUpdate - A partial update to the RedisEnterprise cluster
type ClusterUpdate struct {
	// The identity of the resource.
	Identity *ManagedServiceIdentity

	// Other properties of the cluster.
	Properties *ClusterProperties

	// The SKU to create, which affects price, performance, and features.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string
}

// Database - Describes a database on the RedisEnterprise cluster
type Database struct {
	// Other properties of the database.
	Properties *DatabaseProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DatabaseList - The response of a list-all operation.
type DatabaseList struct {
	// List of databases
	Value []*Database

	// READ-ONLY; The URI to fetch the next page of results.
	NextLink *string
}

// DatabaseProperties - Properties of RedisEnterprise databases, as opposed to general resource properties like location,
// tags
type DatabaseProperties struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted.
	ClientProtocol *Protocol

	// Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *ClusteringPolicy

	// Redis eviction policy - default is VolatileLRU
	EvictionPolicy *EvictionPolicy

	// Optional set of properties to configure geo replication for this database.
	GeoReplication *DatabasePropertiesGeoReplication

	// Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []*Module

	// Persistence settings
	Persistence *Persistence

	// TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int32

	// READ-ONLY; Current provisioning status of the database
	ProvisioningState *ProvisioningState

	// READ-ONLY; Current resource status of the database
	ResourceState *ResourceState
}

// DatabasePropertiesGeoReplication - Optional set of properties to configure geo replication for this database.
type DatabasePropertiesGeoReplication struct {
	// Name for the group of linked database resources
	GroupNickname *string

	// List of database resources to link with this database
	LinkedDatabases []*LinkedDatabase
}

// DatabaseUpdate - A partial update to the RedisEnterprise database
type DatabaseUpdate struct {
	// Properties of the database.
	Properties *DatabaseProperties
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// ExportClusterParameters - Parameters for a Redis Enterprise export operation.
type ExportClusterParameters struct {
	// REQUIRED; SAS URI for the target directory to export to
	SasURI *string
}

// FlushParameters - Parameters for a Redis Enterprise active geo-replication flush operation
type FlushParameters struct {
	// The identifiers of all the other database resources in the georeplication group to be flushed.
	IDs []*string
}

// ForceUnlinkParameters - Parameters for a Redis Enterprise Active Geo Replication Force Unlink operation.
type ForceUnlinkParameters struct {
	// REQUIRED; The resource IDs of the database resources to be unlinked.
	IDs []*string
}

// ImportClusterParameters - Parameters for a Redis Enterprise import operation.
type ImportClusterParameters struct {
	// REQUIRED; SAS URIs for the target blobs to import from
	SasUris []*string
}

// LinkedDatabase - Specifies details of a linked database resource.
type LinkedDatabase struct {
	// Resource ID of a database resource to link with this database.
	ID *string

	// READ-ONLY; State of the link between the database resources.
	State *LinkState
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// Module - Specifies configuration of a redis module
type Module struct {
	// REQUIRED; The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name *string

	// Configuration options for the module, e.g. 'ERRORRATE 0.01 INITIALSIZE 400'.
	Args *string

	// READ-ONLY; The version of the module, e.g. '1.0'.
	Version *string
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

// OperationStatus - The status of a long-running operation.
type OperationStatus struct {
	// The end time of the operation.
	EndTime *string

	// Error response describing why the operation failed.
	Error *ErrorResponse

	// The operation's unique id.
	ID *string

	// The operation's name.
	Name *string

	// The start time of the operation.
	StartTime *string

	// The current status of the operation.
	Status *string
}

// Persistence-related configuration for the RedisEnterprise database
type Persistence struct {
	// Sets whether AOF is enabled.
	AofEnabled *bool

	// Sets the frequency at which data is written to disk.
	AofFrequency *AofFrequency

	// Sets whether RDB is enabled.
	RdbEnabled *bool

	// Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *RdbFrequency
}

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RegenerateKeyParameters - Specifies which access keys to reset to a new random value.
type RegenerateKeyParameters struct {
	// REQUIRED; Which access key to regenerate.
	KeyType *AccessKeyType
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SKU parameters supplied to the create RedisEnterprise operation.
type SKU struct {
	// REQUIRED; The type of RedisEnterprise cluster to deploy. Possible values: (EnterpriseE10, EnterpriseFlashF300 etc.)
	Name *SKUName

	// The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, …) for Enterprise
	// SKUs and (3, 9, 15, …) for Flash SKUs.
	Capacity *int32
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
