//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredisenterprise

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// AccessKeys - The secret access keys used for authenticating connections to redis
type AccessKeys struct {
	// READ-ONLY; The current primary key that clients can use to authenticate
	PrimaryKey *string `json:"primaryKey,omitempty" azure:"ro"`

	// READ-ONLY; The current secondary key that clients can use to authenticate
	SecondaryKey *string `json:"secondaryKey,omitempty" azure:"ro"`
}

// Cluster - Describes the RedisEnterprise cluster
type Cluster struct {
	TrackedResource
	// REQUIRED; The SKU to create, which affects price, performance, and features.
	SKU *SKU `json:"sku,omitempty"`

	// Other properties of the cluster.
	Properties *ClusterProperties `json:"properties,omitempty"`

	// The Availability Zones where this cluster will be deployed.
	Zones []*string `json:"zones,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "zones", c.Zones)
	return json.Marshal(objectMap)
}

// ClusterList - The response of a list-all operation.
type ClusterList struct {
	// List of clusters.
	Value []*Cluster `json:"value,omitempty"`

	// READ-ONLY; The URI to fetch the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ClusterList.
func (c ClusterList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ClusterProperties - Properties of RedisEnterprise clusters, as opposed to general resource properties like location, tags
type ClusterProperties struct {
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTLSVersion *TLSVersion `json:"minimumTlsVersion,omitempty"`

	// READ-ONLY; DNS name of the cluster endpoint
	HostName *string `json:"hostName,omitempty" azure:"ro"`

	// READ-ONLY; List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections []*PrivateEndpointConnection `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// READ-ONLY; Current provisioning status of the cluster
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Version of redis the cluster supports, e.g. '6'
	RedisVersion *string `json:"redisVersion,omitempty" azure:"ro"`

	// READ-ONLY; Current resource status of the cluster
	ResourceState *ResourceState `json:"resourceState,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ClusterProperties.
func (c ClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hostName", c.HostName)
	populate(objectMap, "minimumTlsVersion", c.MinimumTLSVersion)
	populate(objectMap, "privateEndpointConnections", c.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "redisVersion", c.RedisVersion)
	populate(objectMap, "resourceState", c.ResourceState)
	return json.Marshal(objectMap)
}

// ClusterUpdate - A partial update to the RedisEnterprise cluster
type ClusterUpdate struct {
	// Other properties of the cluster.
	Properties *ClusterProperties `json:"properties,omitempty"`

	// The SKU to create, which affects price, performance, and features.
	SKU *SKU `json:"sku,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ClusterUpdate.
func (c ClusterUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// Database - Describes a database on the RedisEnterprise cluster
type Database struct {
	ProxyResource
	// Other properties of the database.
	Properties *DatabaseProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Database.
func (d Database) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	d.ProxyResource.marshalInternal(objectMap)
	populate(objectMap, "properties", d.Properties)
	return json.Marshal(objectMap)
}

// DatabaseList - The response of a list-all operation.
type DatabaseList struct {
	// List of databases
	Value []*Database `json:"value,omitempty"`

	// READ-ONLY; The URI to fetch the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type DatabaseList.
func (d DatabaseList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// DatabaseProperties - Properties of RedisEnterprise databases, as opposed to general resource properties like location, tags
type DatabaseProperties struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted.
	ClientProtocol *Protocol `json:"clientProtocol,omitempty"`

	// Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *ClusteringPolicy `json:"clusteringPolicy,omitempty"`

	// Redis eviction policy - default is VolatileLRU
	EvictionPolicy *EvictionPolicy `json:"evictionPolicy,omitempty"`

	// Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []*Module `json:"modules,omitempty"`

	// Persistence settings
	Persistence *Persistence `json:"persistence,omitempty"`

	// TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int32 `json:"port,omitempty"`

	// READ-ONLY; Current provisioning status of the database
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Current resource status of the database
	ResourceState *ResourceState `json:"resourceState,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type DatabaseProperties.
func (d DatabaseProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clientProtocol", d.ClientProtocol)
	populate(objectMap, "clusteringPolicy", d.ClusteringPolicy)
	populate(objectMap, "evictionPolicy", d.EvictionPolicy)
	populate(objectMap, "modules", d.Modules)
	populate(objectMap, "persistence", d.Persistence)
	populate(objectMap, "port", d.Port)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "resourceState", d.ResourceState)
	return json.Marshal(objectMap)
}

// DatabaseUpdate - A partial update to the RedisEnterprise database
type DatabaseUpdate struct {
	// Properties of the database.
	Properties *DatabaseProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DatabaseUpdate.
func (d DatabaseUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", d.Properties)
	return json.Marshal(objectMap)
}

// DatabasesBeginCreateOptions contains the optional parameters for the Databases.BeginCreate method.
type DatabasesBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// DatabasesBeginDeleteOptions contains the optional parameters for the Databases.BeginDelete method.
type DatabasesBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// DatabasesBeginExportOptions contains the optional parameters for the Databases.BeginExport method.
type DatabasesBeginExportOptions struct {
	// placeholder for future optional parameters
}

// DatabasesBeginImportOptions contains the optional parameters for the Databases.BeginImport method.
type DatabasesBeginImportOptions struct {
	// placeholder for future optional parameters
}

// DatabasesBeginRegenerateKeyOptions contains the optional parameters for the Databases.BeginRegenerateKey method.
type DatabasesBeginRegenerateKeyOptions struct {
	// placeholder for future optional parameters
}

// DatabasesBeginUpdateOptions contains the optional parameters for the Databases.BeginUpdate method.
type DatabasesBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// DatabasesGetOptions contains the optional parameters for the Databases.Get method.
type DatabasesGetOptions struct {
	// placeholder for future optional parameters
}

// DatabasesListByClusterOptions contains the optional parameters for the Databases.ListByCluster method.
type DatabasesListByClusterOptions struct {
	// placeholder for future optional parameters
}

// DatabasesListKeysOptions contains the optional parameters for the Databases.ListKeys method.
type DatabasesListKeysOptions struct {
	// placeholder for future optional parameters
}

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

// ExportClusterParameters - Parameters for a Redis Enterprise export operation.
type ExportClusterParameters struct {
	// REQUIRED; SAS URI for the target directory to export to
	SasURI *string `json:"sasUri,omitempty"`
}

// ImportClusterParameters - Parameters for a Redis Enterprise import operation.
type ImportClusterParameters struct {
	// REQUIRED; SAS URIs for the target blobs to import from
	SasUris []*string `json:"sasUris,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ImportClusterParameters.
func (i ImportClusterParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "sasUris", i.SasUris)
	return json.Marshal(objectMap)
}

// Module - Specifies configuration of a redis module
type Module struct {
	// REQUIRED; The name of the module, e.g. 'RedisBloom', 'RediSearch', 'RedisTimeSeries'
	Name *string `json:"name,omitempty"`

	// Configuration options for the module, e.g. 'ERRORRATE 0.00 INITIALSIZE 400'.
	Args *string `json:"args,omitempty"`

	// READ-ONLY; The version of the module, e.g. '1.0'.
	Version *string `json:"version,omitempty" azure:"ro"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write", "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual Machine", "Restart Virtual
	// Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationStatus - The status of a long-running operation.
type OperationStatus struct {
	// The end time of the operation.
	EndTime *string `json:"endTime,omitempty"`

	// Error response describing why the operation failed.
	Error *ErrorResponse `json:"error,omitempty"`

	// The operation's unique id.
	ID *string `json:"id,omitempty"`

	// The operation's name.
	Name *string `json:"name,omitempty"`

	// The start time of the operation.
	StartTime *string `json:"startTime,omitempty"`

	// The current status of the operation.
	Status *string `json:"status,omitempty"`
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// OperationsStatusGetOptions contains the optional parameters for the OperationsStatus.Get method.
type OperationsStatusGetOptions struct {
	// placeholder for future optional parameters
}

// Persistence-related configuration for the RedisEnterprise database
type Persistence struct {
	// Sets whether AOF is enabled.
	AofEnabled *bool `json:"aofEnabled,omitempty"`

	// Sets the frequency at which data is written to disk.
	AofFrequency *AofFrequency `json:"aofFrequency,omitempty"`

	// Sets whether RDB is enabled.
	RdbEnabled *bool `json:"rdbEnabled,omitempty"`

	// Sets the frequency at which a snapshot of the database is created.
	RdbFrequency *RdbFrequency `json:"rdbFrequency,omitempty"`
}

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string `json:"id,omitempty" azure:"ro"`
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	Resource
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnection.
func (p PrivateEndpointConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	p.Resource.marshalInternal(objectMap)
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsBeginPutOptions contains the optional parameters for the PrivateEndpointConnections.BeginPut method.
type PrivateEndpointConnectionsBeginPutOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsDeleteOptions contains the optional parameters for the PrivateEndpointConnections.Delete method.
type PrivateEndpointConnectionsDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsGetOptions contains the optional parameters for the PrivateEndpointConnections.Get method.
type PrivateEndpointConnectionsGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsListOptions contains the optional parameters for the PrivateEndpointConnections.List method.
type PrivateEndpointConnectionsListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	Resource
	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResource.
func (p PrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	p.Resource.marshalInternal(objectMap)
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty"`

	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// PrivateLinkResourcesListByClusterOptions contains the optional parameters for the PrivateLinkResources.ListByCluster method.
type PrivateLinkResourcesListByClusterOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a location
type ProxyResource struct {
	Resource
}

func (p ProxyResource) marshalInternal(objectMap map[string]interface{}) {
	p.Resource.marshalInternal(objectMap)
}

// RedisEnterpriseBeginCreateOptions contains the optional parameters for the RedisEnterprise.BeginCreate method.
type RedisEnterpriseBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// RedisEnterpriseBeginDeleteOptions contains the optional parameters for the RedisEnterprise.BeginDelete method.
type RedisEnterpriseBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// RedisEnterpriseBeginUpdateOptions contains the optional parameters for the RedisEnterprise.BeginUpdate method.
type RedisEnterpriseBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// RedisEnterpriseGetOptions contains the optional parameters for the RedisEnterprise.Get method.
type RedisEnterpriseGetOptions struct {
	// placeholder for future optional parameters
}

// RedisEnterpriseListByResourceGroupOptions contains the optional parameters for the RedisEnterprise.ListByResourceGroup method.
type RedisEnterpriseListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// RedisEnterpriseListOptions contains the optional parameters for the RedisEnterprise.List method.
type RedisEnterpriseListOptions struct {
	// placeholder for future optional parameters
}

// RegenerateKeyParameters - Specifies which access keys to reset to a new random value.
type RegenerateKeyParameters struct {
	// REQUIRED; Which access key to regenerate.
	KeyType *AccessKeyType `json:"keyType,omitempty"`
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

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// SKU parameters supplied to the create RedisEnterprise operation.
type SKU struct {
	// REQUIRED; The type of RedisEnterprise cluster to deploy. Possible values: (EnterpriseE10, EnterpriseFlashF300 etc.)
	Name *SKUName `json:"name,omitempty"`

	// The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, …) for Enterprise SKUs and (3, 9, 15, …) for
	// Flash SKUs.
	Capacity *int32 `json:"capacity,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
type TrackedResource struct {
	Resource
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
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
