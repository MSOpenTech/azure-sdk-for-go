//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers

import "time"

// ActiveDirectoryAdministrator - Represents an Active Directory administrator.
type ActiveDirectoryAdministrator struct {
	// REQUIRED; Properties of the active directory administrator.
	Properties *AdministratorProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ActiveDirectoryAdministratorAdd - Represents an Active Directory administrator.
type ActiveDirectoryAdministratorAdd struct {
	// Properties of the active directory administrator.
	Properties *AdministratorPropertiesForAdd
}

// AdministratorListResult - A list of active directory administrators.
type AdministratorListResult struct {
	// The link used to get the next page of active directory.
	NextLink *string

	// The list of active directory administrators
	Value []*ActiveDirectoryAdministrator
}

// AdministratorProperties - The properties of an Active Directory administrator.
type AdministratorProperties struct {
	// The objectId of the Active Directory administrator.
	ObjectID *string

	// Active Directory administrator principal name.
	PrincipalName *string

	// The principal type used to represent the type of Active Directory Administrator.
	PrincipalType *PrincipalType

	// The tenantId of the Active Directory administrator.
	TenantID *string
}

// AdministratorPropertiesForAdd - The properties of an Active Directory administrator.
type AdministratorPropertiesForAdd struct {
	// Active Directory administrator principal name.
	PrincipalName *string

	// The principal type used to represent the type of Active Directory Administrator.
	PrincipalType *PrincipalType

	// The tenantId of the Active Directory administrator.
	TenantID *string
}

// AdministratorsClientBeginCreateOptions contains the optional parameters for the AdministratorsClient.BeginCreate method.
type AdministratorsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AdministratorsClientBeginDeleteOptions contains the optional parameters for the AdministratorsClient.BeginDelete method.
type AdministratorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AdministratorsClientGetOptions contains the optional parameters for the AdministratorsClient.Get method.
type AdministratorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AdministratorsClientListByServerOptions contains the optional parameters for the AdministratorsClient.NewListByServerPager
// method.
type AdministratorsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// AuthConfig - Authentication configuration properties of a server
type AuthConfig struct {
	// If Enabled, Azure Active Directory authentication is enabled.
	ActiveDirectoryAuth *ActiveDirectoryAuthEnum

	// If Enabled, Password authentication is enabled.
	PasswordAuth *PasswordAuthEnum

	// Tenant id of the server.
	TenantID *string
}

// Backup properties of a server
type Backup struct {
	// Backup retention days for the server.
	BackupRetentionDays *int32

	// A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *GeoRedundantBackupEnum

	// READ-ONLY; The earliest restore point time (ISO8601 format) for server.
	EarliestRestoreDate *time.Time
}

// BackupsClientGetOptions contains the optional parameters for the BackupsClient.Get method.
type BackupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BackupsClientListByServerOptions contains the optional parameters for the BackupsClient.NewListByServerPager method.
type BackupsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesListResult - location capability
type CapabilitiesListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; A list of supported capabilities.
	Value []*CapabilityProperties
}

// CapabilityProperties - Location capabilities.
type CapabilityProperties struct {
	// READ-ONLY; A value indicating whether fast provisioning is supported in this region.
	FastProvisioningSupported *bool

	// READ-ONLY; A value indicating whether a new server in this region can have geo-backups to paired region.
	GeoBackupSupported *bool

	// READ-ONLY; The status
	Status *string

	// READ-ONLY
	SupportedFastProvisioningEditions []*FastProvisioningEditionCapability

	// READ-ONLY
	SupportedFlexibleServerEditions []*FlexibleServerEditionCapability

	// READ-ONLY; Supported high availability mode
	SupportedHAMode []*string

	// READ-ONLY
	SupportedHyperscaleNodeEditions []*HyperscaleNodeEditionCapability

	// READ-ONLY; zone name
	Zone *string

	// READ-ONLY; A value indicating whether a new server in this region can have geo-backups to paired region.
	ZoneRedundantHaAndGeoBackupSupported *bool

	// READ-ONLY; A value indicating whether a new server in this region can support multi zone HA.
	ZoneRedundantHaSupported *bool
}

// CheckNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityClient.Execute
// method.
type CheckNameAvailabilityClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityWithLocationClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityWithLocationClient.Execute
// method.
type CheckNameAvailabilityWithLocationClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// Configuration - Represents a Configuration.
type Configuration struct {
	// The properties of a configuration.
	Properties *ConfigurationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ConfigurationForUpdate - Represents a Configuration.
type ConfigurationForUpdate struct {
	// The properties of a configuration.
	Properties *ConfigurationProperties
}

// ConfigurationListResult - A list of server configurations.
type ConfigurationListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of server configurations.
	Value []*Configuration
}

// ConfigurationProperties - The properties of a configuration.
type ConfigurationProperties struct {
	// Source of the configuration.
	Source *string

	// Value of the configuration.
	Value *string

	// READ-ONLY; Allowed values of the configuration.
	AllowedValues *string

	// READ-ONLY; Data type of the configuration.
	DataType *ConfigurationDataType

	// READ-ONLY; Default value of the configuration.
	DefaultValue *string

	// READ-ONLY; Description of the configuration.
	Description *string

	// READ-ONLY; Configuration documentation link.
	DocumentationLink *string

	// READ-ONLY; Configuration is pending restart or not.
	IsConfigPendingRestart *bool

	// READ-ONLY; Configuration dynamic or static.
	IsDynamicConfig *bool

	// READ-ONLY; Configuration read-only or not.
	IsReadOnly *bool

	// READ-ONLY; Configuration unit.
	Unit *string
}

// ConfigurationsClientBeginPutOptions contains the optional parameters for the ConfigurationsClient.BeginPut method.
type ConfigurationsClientBeginPutOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientBeginUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginUpdate method.
type ConfigurationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
type ConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientListByServerOptions contains the optional parameters for the ConfigurationsClient.NewListByServerPager
// method.
type ConfigurationsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// DataEncryption - Data encryption properties of a server
type DataEncryption struct {
	// URI for the key for data encryption for primary server.
	PrimaryKeyURI *string

	// Resource Id for the User assigned identity to be used for data encryption for primary server.
	PrimaryUserAssignedIdentityID *string

	// Data encryption type to depict if it is System Managed vs Azure Key vault.
	Type *ArmServerKeyType
}

// Database - Represents a Database.
type Database struct {
	// The properties of a database.
	Properties *DatabaseProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DatabaseListResult - A List of databases.
type DatabaseListResult struct {
	// The link used to get the next page of databases.
	NextLink *string

	// The list of databases housed in a server
	Value []*Database
}

// DatabaseProperties - The properties of a database.
type DatabaseProperties struct {
	// The charset of the database.
	Charset *string

	// The collation of the database.
	Collation *string
}

// DatabasesClientBeginCreateOptions contains the optional parameters for the DatabasesClient.BeginCreate method.
type DatabasesClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabasesClientBeginDeleteOptions contains the optional parameters for the DatabasesClient.BeginDelete method.
type DatabasesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabasesClientGetOptions contains the optional parameters for the DatabasesClient.Get method.
type DatabasesClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatabasesClientListByServerOptions contains the optional parameters for the DatabasesClient.NewListByServerPager method.
type DatabasesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// DelegatedSubnetUsage - Delegated subnet usage data.
type DelegatedSubnetUsage struct {
	// READ-ONLY; Name of the delegated subnet for which IP addresses are in use
	SubnetName *string

	// READ-ONLY; Number of IP addresses used by the delegated subnet
	Usage *int64
}

type FastProvisioningEditionCapability struct {
	// READ-ONLY; Fast provisioning supported sku name
	SupportedSKU *string

	// READ-ONLY; Fast provisioning supported version
	SupportedServerVersions *string

	// READ-ONLY; Fast provisioning supported storage in Gb
	SupportedStorageGb *int64
}

// FirewallRule - Represents a server firewall rule.
type FirewallRule struct {
	// REQUIRED; The properties of a firewall rule.
	Properties *FirewallRuleProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FirewallRuleListResult - A list of firewall rules.
type FirewallRuleListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of firewall rules in a server.
	Value []*FirewallRule
}

// FirewallRuleProperties - The properties of a server firewall rule.
type FirewallRuleProperties struct {
	// REQUIRED; The end IP address of the server firewall rule. Must be IPv4 format.
	EndIPAddress *string

	// REQUIRED; The start IP address of the server firewall rule. Must be IPv4 format.
	StartIPAddress *string
}

// FirewallRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.BeginCreateOrUpdate
// method.
type FirewallRulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FirewallRulesClientBeginDeleteOptions contains the optional parameters for the FirewallRulesClient.BeginDelete method.
type FirewallRulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
type FirewallRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientListByServerOptions contains the optional parameters for the FirewallRulesClient.NewListByServerPager
// method.
type FirewallRulesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// FlexibleServerEditionCapability - Flexible server edition capabilities.
type FlexibleServerEditionCapability struct {
	// READ-ONLY; Server edition name
	Name *string

	// READ-ONLY; The status
	Status *string

	// READ-ONLY; The list of server versions supported by this server edition.
	SupportedServerVersions []*ServerVersionCapability

	// READ-ONLY; The list of editions supported by this server edition.
	SupportedStorageEditions []*StorageEditionCapability
}

// GetPrivateDNSZoneSuffixClientExecuteOptions contains the optional parameters for the GetPrivateDNSZoneSuffixClient.Execute
// method.
type GetPrivateDNSZoneSuffixClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// HighAvailability - High availability properties of a server
type HighAvailability struct {
	// The HA mode for the server.
	Mode *HighAvailabilityMode

	// availability zone information of the standby.
	StandbyAvailabilityZone *string

	// READ-ONLY; A state of a HA server that is visible to user.
	State *ServerHAState
}

// HyperscaleNodeEditionCapability - Hyperscale node edition capabilities.
type HyperscaleNodeEditionCapability struct {
	// READ-ONLY; Server edition name
	Name *string

	// READ-ONLY; The status
	Status *string

	// READ-ONLY; The list of Node Types supported by this server edition.
	SupportedNodeTypes []*NodeTypeCapability

	// READ-ONLY; The list of server versions supported by this server edition.
	SupportedServerVersions []*ServerVersionCapability

	// READ-ONLY; The list of editions supported by this server edition.
	SupportedStorageEditions []*StorageEditionCapability
}

// LocationBasedCapabilitiesClientExecuteOptions contains the optional parameters for the LocationBasedCapabilitiesClient.NewExecutePager
// method.
type LocationBasedCapabilitiesClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// MaintenanceWindow - Maintenance window properties of a server.
type MaintenanceWindow struct {
	// indicates whether custom window is enabled or disabled
	CustomWindow *string

	// day of week for maintenance window
	DayOfWeek *int32

	// start hour for maintenance window
	StartHour *int32

	// start minute for maintenance window
	StartMinute *int32
}

// NameAvailability - Represents a resource name availability.
type NameAvailability struct {
	// Detailed reason why the given name is available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason

	// READ-ONLY; name of the PostgreSQL server.
	Name *string

	// READ-ONLY; type of the server
	Type *string
}

// Network properties of a server
type Network struct {
	// delegated subnet arm resource id.
	DelegatedSubnetResourceID *string

	// private dns zone arm resource id.
	PrivateDNSZoneArmResourceID *string

	// READ-ONLY; public network access is enabled or not
	PublicNetworkAccess *ServerPublicNetworkAccessState
}

// NodeTypeCapability - node type capability
type NodeTypeCapability struct {
	// READ-ONLY; note type name
	Name *string

	// READ-ONLY; note type
	NodeType *string

	// READ-ONLY; The status
	Status *string
}

// Operation - REST API operation definition.
type Operation struct {
	// Indicates whether the operation is a data action
	IsDataAction *bool

	// READ-ONLY; The localized display information for this particular operation or action.
	Display *OperationDisplay

	// READ-ONLY; The name of the operation being performed on this particular object.
	Name *string

	// READ-ONLY; The intended executor of the operation.
	Origin *OperationOrigin

	// READ-ONLY; Additional descriptions for the operation.
	Properties map[string]any
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// READ-ONLY; Operation description.
	Description *string

	// READ-ONLY; Localized friendly name for the operation.
	Operation *string

	// READ-ONLY; Operation resource provider name.
	Provider *string

	// READ-ONLY; Resource on which the operation is performed.
	Resource *string
}

// OperationListResult - A list of resource provider operations.
type OperationListResult struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// Collection of available operation details
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ReplicasClientListByServerOptions contains the optional parameters for the ReplicasClient.NewListByServerPager method.
type ReplicasClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// RestartParameter - Represents server restart parameters.
type RestartParameter struct {
	// Failover mode.
	FailoverMode *FailoverMode

	// Indicates whether to restart the server with failover.
	RestartWithFailover *bool
}

// SKU - Sku information related properties of a server.
type SKU struct {
	// REQUIRED; The name of the sku, typically, tier + family + cores, e.g. StandardD4sv3.
	Name *string

	// REQUIRED; The tier of the particular SKU, e.g. Burstable.
	Tier *SKUTier
}

// Server - Represents a server.
type Server struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Describes the identity of the application.
	Identity *UserAssignedIdentity

	// Properties of the server.
	Properties *ServerProperties

	// The SKU (pricing tier) of the server.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServerBackup - Server backup properties
type ServerBackup struct {
	// The properties of a server backup.
	Properties *ServerBackupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServerBackupListResult - A list of server backups.
type ServerBackupListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of backups of a server.
	Value []*ServerBackup
}

// ServerBackupProperties - The properties of a server backup.
type ServerBackupProperties struct {
	// Backup type.
	BackupType *Origin

	// Backup completed time (ISO8601 format).
	CompletedTime *time.Time

	// Backup source
	Source *string
}

// ServerForUpdate - Represents a server to be updated.
type ServerForUpdate struct {
	// Describes the identity of the application.
	Identity *UserAssignedIdentity

	// Properties of the server.
	Properties *ServerPropertiesForUpdate

	// The SKU (pricing tier) of the server.
	SKU *SKU

	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// ServerListResult - A list of servers.
type ServerListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of flexible servers
	Value []*Server
}

// ServerProperties - The properties of a server.
type ServerProperties struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for
	// creation).
	AdministratorLogin *string

	// The administrator login password (required for server creation).
	AdministratorLoginPassword *string

	// AuthConfig properties of a server.
	AuthConfig *AuthConfig

	// availability zone information of the server.
	AvailabilityZone *string

	// Backup properties of a server.
	Backup *Backup

	// The mode to create a new PostgreSQL server.
	CreateMode *CreateMode

	// Data encryption properties of a server.
	DataEncryption *DataEncryption

	// High availability properties of a server.
	HighAvailability *HighAvailability

	// Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow

	// Network properties of a server.
	Network *Network

	// Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when 'createMode' is 'PointInTimeRestore'
	// or 'GeoRestore'.
	PointInTimeUTC *time.Time

	// Replicas allowed for a server.
	ReplicaCapacity *int32

	// Replication role of the server
	ReplicationRole *ReplicationRole

	// The source server resource ID to restore from. It's required when 'createMode' is 'PointInTimeRestore' or 'GeoRestore'
	// or 'Replica'.
	SourceServerResourceID *string

	// Storage properties of a server.
	Storage *Storage

	// PostgreSQL Server version.
	Version *ServerVersion

	// READ-ONLY; The fully qualified domain name of a server.
	FullyQualifiedDomainName *string

	// READ-ONLY; The minor version of the server.
	MinorVersion *string

	// READ-ONLY; A state of a server that is visible to user.
	State *ServerState
}

type ServerPropertiesForUpdate struct {
	// The password of the administrator login.
	AdministratorLoginPassword *string

	// AuthConfig properties of a server.
	AuthConfig *AuthConfig

	// Backup properties of a server.
	Backup *Backup

	// The mode to update a new PostgreSQL server.
	CreateMode *CreateModeForUpdate

	// Data encryption properties of a server.
	DataEncryption *DataEncryption

	// High availability properties of a server.
	HighAvailability *HighAvailability

	// Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow

	// Replication role of the server
	ReplicationRole *ReplicationRole

	// Storage properties of a server.
	Storage *Storage

	// PostgreSQL Server version.
	Version *ServerVersion
}

// ServerVersionCapability - Server version capabilities.
type ServerVersionCapability struct {
	// READ-ONLY; server version
	Name *string

	// READ-ONLY; The status
	Status *string

	// READ-ONLY
	SupportedVcores []*VcoreCapability

	// READ-ONLY; Supported servers versions to upgrade
	SupportedVersionsToUpgrade []*string
}

// ServersClientBeginCreateOptions contains the optional parameters for the ServersClient.BeginCreate method.
type ServersClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginDeleteOptions contains the optional parameters for the ServersClient.BeginDelete method.
type ServersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginRestartOptions contains the optional parameters for the ServersClient.BeginRestart method.
type ServersClientBeginRestartOptions struct {
	// The parameters for restarting a server.
	Parameters *RestartParameter
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginStartOptions contains the optional parameters for the ServersClient.BeginStart method.
type ServersClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginStopOptions contains the optional parameters for the ServersClient.BeginStop method.
type ServersClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginUpdateOptions contains the optional parameters for the ServersClient.BeginUpdate method.
type ServersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientGetOptions contains the optional parameters for the ServersClient.Get method.
type ServersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServersClientListByResourceGroupOptions contains the optional parameters for the ServersClient.NewListByResourceGroupPager
// method.
type ServersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServersClientListOptions contains the optional parameters for the ServersClient.NewListPager method.
type ServersClientListOptions struct {
	// placeholder for future optional parameters
}

// Storage properties of a server
type Storage struct {
	// Max storage allowed for a server.
	StorageSizeGB *int32
}

// StorageEditionCapability - storage edition capability
type StorageEditionCapability struct {
	// READ-ONLY; storage edition name
	Name *string

	// READ-ONLY; The status
	Status *string

	// READ-ONLY
	SupportedStorageMB []*StorageMBCapability
}

// StorageMBCapability - storage size in MB capability
type StorageMBCapability struct {
	// READ-ONLY; storage MB name
	Name *string

	// READ-ONLY; The status
	Status *string

	// READ-ONLY; storage size in MB
	StorageSizeMB *int64

	// READ-ONLY; supported IOPS
	SupportedIops *int64

	// READ-ONLY
	SupportedUpgradableTierList []*StorageTierCapability
}

type StorageTierCapability struct {
	// READ-ONLY; Supported IOPS for this storage tier
	Iops *int64

	// READ-ONLY; Indicates if this is a baseline storage tier or not
	IsBaseline *bool

	// READ-ONLY; Name to represent Storage tier capability
	Name *string

	// READ-ONLY; Status os this storage tier
	Status *string

	// READ-ONLY; Storage tier name
	TierName *string
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

// UserAssignedIdentity - Information describing the identities associated with this application.
type UserAssignedIdentity struct {
	// REQUIRED; the types of identities associated with this resource; currently restricted to 'SystemAssigned and UserAssigned'
	Type *IdentityType

	// represents user assigned identities map.
	UserAssignedIdentities map[string]*UserIdentity
}

// UserIdentity - Describes a single user-assigned identity associated with the application.
type UserIdentity struct {
	// the client identifier of the Service Principal which this identity represents.
	ClientID *string

	// the object identifier of the Service Principal which this identity represents.
	PrincipalID *string
}

// VcoreCapability - Vcores capability
type VcoreCapability struct {
	// READ-ONLY; vCore name
	Name *string

	// READ-ONLY; The status
	Status *string

	// READ-ONLY; supported IOPS
	SupportedIops *int64

	// READ-ONLY; supported memory per vCore in MB
	SupportedMemoryPerVcoreMB *int64

	// READ-ONLY; supported vCores
	VCores *int64
}

// VirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the VirtualNetworkSubnetUsageClient.Execute
// method.
type VirtualNetworkSubnetUsageClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkSubnetUsageParameter - Virtual network subnet usage parameter
type VirtualNetworkSubnetUsageParameter struct {
	// Virtual network resource id.
	VirtualNetworkArmResourceID *string
}

// VirtualNetworkSubnetUsageResult - Virtual network subnet usage data.
type VirtualNetworkSubnetUsageResult struct {
	// READ-ONLY
	DelegatedSubnetsUsage []*DelegatedSubnetUsage

	// READ-ONLY; location of the delegated subnet usage
	Location *string

	// READ-ONLY; subscriptionId of the delegated subnet usage
	SubscriptionID *string
}
