//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsqlvirtualmachine

import "time"

// AADAuthenticationSettings - Enable AAD authentication for SQL VM.
type AADAuthenticationSettings struct {
	// The client Id of the Managed Identity to query Microsoft Graph API. An empty string must be used for the system assigned
	// Managed Identity
	ClientID *string
}

// AdditionalFeaturesServerConfigurations - Additional SQL Server feature settings.
type AdditionalFeaturesServerConfigurations struct {
	// Enable or disable R services (SQL 2016 onwards).
	IsRServicesEnabled *bool
}

// AgConfiguration - Availability group configuration.
type AgConfiguration struct {
	// Replica configurations.
	Replicas []*AgReplica
}

// AgReplica - Availability group replica configuration.
type AgReplica struct {
	// Replica commit mode in availability group.
	Commit *Commit

	// Replica failover mode in availability group.
	Failover *Failover

	// Replica readable secondary mode in availability group.
	ReadableSecondary *ReadableSecondary

	// Replica Role in availability group.
	Role *Role

	// Sql VirtualMachine Instance Id.
	SQLVirtualMachineInstanceID *string
}

// AssessmentSettings - Configure SQL best practices Assessment for databases in your SQL virtual machine.
type AssessmentSettings struct {
	// Enable or disable SQL best practices Assessment feature on SQL virtual machine.
	Enable *bool

	// Run SQL best practices Assessment immediately on SQL virtual machine.
	RunImmediately *bool

	// Schedule for SQL best practices Assessment.
	Schedule *Schedule
}

// AutoBackupSettings - Configure backups for databases in your SQL virtual machine.
type AutoBackupSettings struct {
	// Backup schedule type.
	BackupScheduleType *BackupScheduleType

	// Include or exclude system databases from auto backup.
	BackupSystemDbs *bool

	// Days of the week for the backups when FullBackupFrequency is set to Weekly.
	DaysOfWeek []*AutoBackupDaysOfWeek

	// Enable or disable autobackup on SQL virtual machine.
	Enable *bool

	// Enable or disable encryption for backup on SQL virtual machine.
	EnableEncryption *bool

	// Frequency of full backups. In both cases, full backups begin during the next scheduled time window.
	FullBackupFrequency *FullBackupFrequencyType

	// Start time of a given day during which full backups can take place. 0-23 hours.
	FullBackupStartTime *int32

	// Duration of the time window of a given day during which full backups can take place. 1-23 hours.
	FullBackupWindowHours *int32

	// Frequency of log backups. 5-60 minutes.
	LogBackupFrequency *int32

	// Password for encryption on backup.
	Password *string

	// Retention period of backup: 1-90 days.
	RetentionPeriod *int32

	// Storage account key where backup will be taken to.
	StorageAccessKey *string

	// Storage account url where backup will be taken to.
	StorageAccountURL *string

	// Storage container name where backup will be taken to.
	StorageContainerName *string
}

// AutoPatchingSettings - Set a patching window during which Windows and SQL patches will be applied.
type AutoPatchingSettings struct {
	// Day of week to apply the patch on.
	DayOfWeek *DayOfWeek

	// Enable or disable autopatching on SQL virtual machine.
	Enable *bool

	// Duration of patching.
	MaintenanceWindowDuration *int32

	// Hour of the day when patching is initiated. Local VM time.
	MaintenanceWindowStartingHour *int32
}

// AvailabilityGroupListener - A SQL Server availability group listener.
type AvailabilityGroupListener struct {
	// Resource properties.
	Properties *AvailabilityGroupListenerProperties

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// AvailabilityGroupListenerListResult - A list of availability group listeners.
type AvailabilityGroupListenerListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*AvailabilityGroupListener
}

// AvailabilityGroupListenerProperties - The properties of an availability group listener.
type AvailabilityGroupListenerProperties struct {
	// Availability Group configuration.
	AvailabilityGroupConfiguration *AgConfiguration

	// Name of the availability group.
	AvailabilityGroupName *string

	// Create a default availability group if it does not exist.
	CreateDefaultAvailabilityGroupIfNotExist *bool

	// List of load balancer configurations for an availability group listener.
	LoadBalancerConfigurations []*LoadBalancerConfiguration

	// List of multi subnet IP configurations for an AG listener.
	MultiSubnetIPConfigurations []*MultiSubnetIPConfiguration

	// Listener port.
	Port *int32

	// READ-ONLY; Provisioning state to track the async operation status.
	ProvisioningState *string
}

// AvailabilityGroupListenersClientBeginCreateOrUpdateOptions contains the optional parameters for the AvailabilityGroupListenersClient.BeginCreateOrUpdate
// method.
type AvailabilityGroupListenersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AvailabilityGroupListenersClientBeginDeleteOptions contains the optional parameters for the AvailabilityGroupListenersClient.BeginDelete
// method.
type AvailabilityGroupListenersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AvailabilityGroupListenersClientGetOptions contains the optional parameters for the AvailabilityGroupListenersClient.Get
// method.
type AvailabilityGroupListenersClientGetOptions struct {
	// The child resources to include in the response.
	Expand *string
}

// AvailabilityGroupListenersClientListByGroupOptions contains the optional parameters for the AvailabilityGroupListenersClient.NewListByGroupPager
// method.
type AvailabilityGroupListenersClientListByGroupOptions struct {
	// placeholder for future optional parameters
}

// Group - A SQL virtual machine group.
type Group struct {
	// REQUIRED; Resource location.
	Location *string

	// Resource properties.
	Properties *GroupProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// GroupListResult - A list of SQL virtual machine groups.
type GroupListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*Group
}

// GroupProperties - The properties of a SQL virtual machine group.
type GroupProperties struct {
	// SQL image offer. Examples may include SQL2016-WS2016, SQL2017-WS2016.
	SQLImageOffer *string

	// SQL image sku.
	SQLImageSKU *SQLVMGroupImageSKU

	// Cluster Active Directory domain profile.
	WsfcDomainProfile *WsfcDomainProfile

	// READ-ONLY; Cluster type.
	ClusterConfiguration *ClusterConfiguration

	// READ-ONLY; Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the group and
	// the OS type.
	ClusterManagerType *ClusterManagerType

	// READ-ONLY; Provisioning state to track the async operation status.
	ProvisioningState *string

	// READ-ONLY; Scale type.
	ScaleType *ScaleType
}

// GroupUpdate - An update to a SQL virtual machine group.
type GroupUpdate struct {
	// Resource tags.
	Tags map[string]*string
}

// GroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupsClient.BeginCreateOrUpdate method.
type GroupsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GroupsClientBeginDeleteOptions contains the optional parameters for the GroupsClient.BeginDelete method.
type GroupsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GroupsClientBeginUpdateOptions contains the optional parameters for the GroupsClient.BeginUpdate method.
type GroupsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GroupsClientGetOptions contains the optional parameters for the GroupsClient.Get method.
type GroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// GroupsClientListByResourceGroupOptions contains the optional parameters for the GroupsClient.NewListByResourceGroupPager
// method.
type GroupsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// GroupsClientListOptions contains the optional parameters for the GroupsClient.NewListPager method.
type GroupsClientListOptions struct {
	// placeholder for future optional parameters
}

// KeyVaultCredentialSettings - Configure your SQL virtual machine to be able to connect to the Azure Key Vault service.
type KeyVaultCredentialSettings struct {
	// Azure Key Vault url.
	AzureKeyVaultURL *string

	// Credential name.
	CredentialName *string

	// Enable or disable key vault credential setting.
	Enable *bool

	// Service principal name to access key vault.
	ServicePrincipalName *string

	// Service principal name secret to access key vault.
	ServicePrincipalSecret *string
}

// ListResult - A list of SQL virtual machines.
type ListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*SQLVirtualMachine
}

// LoadBalancerConfiguration - A load balancer configuration for an availability group listener.
type LoadBalancerConfiguration struct {
	// Resource id of the load balancer.
	LoadBalancerResourceID *string

	// Private IP address.
	PrivateIPAddress *PrivateIPAddress

	// Probe port.
	ProbePort *int32

	// Resource id of the public IP.
	PublicIPAddressResourceID *string

	// List of the SQL virtual machine instance resource id's that are enrolled into the availability group listener.
	SQLVirtualMachineInstances []*string
}

// MultiSubnetIPConfiguration - Multi subnet ip configuration for an availability group listener.
type MultiSubnetIPConfiguration struct {
	// REQUIRED; Private IP address.
	PrivateIPAddress *PrivateIPAddress

	// REQUIRED; SQL virtual machine instance resource id that are enrolled into the availability group listener.
	SQLVirtualMachineInstance *string
}

// Operation - SQL REST API operation definition.
type Operation struct {
	// READ-ONLY; The localized display information for this particular operation / action.
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
	// READ-ONLY; The localized friendly description for the operation.
	Description *string

	// READ-ONLY; The localized friendly name for the operation.
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name.
	Provider *string

	// READ-ONLY; The localized friendly form of the resource type related to this action/operation.
	Resource *string
}

// OperationListResult - Result of the request to list SQL operations.
type OperationListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateIPAddress - A private IP address bound to the availability group listener.
type PrivateIPAddress struct {
	// Private IP address bound to the availability group listener.
	IPAddress *string

	// Subnet used to include private IP.
	SubnetResourceID *string
}

// Properties - The SQL virtual machine properties.
type Properties struct {
	// SQL best practices Assessment Settings.
	AssessmentSettings *AssessmentSettings

	// Auto backup settings for SQL Server.
	AutoBackupSettings *AutoBackupSettings

	// Auto patching settings for applying critical security updates to SQL virtual machine.
	AutoPatchingSettings *AutoPatchingSettings

	// Enable automatic upgrade of Sql IaaS extension Agent.
	EnableAutomaticUpgrade *bool

	// Key vault credential settings.
	KeyVaultCredentialSettings *KeyVaultCredentialSettings

	// SQL IaaS Agent least privilege mode.
	LeastPrivilegeMode *LeastPrivilegeMode

	// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
	SQLImageOffer *string

	// SQL Server edition type.
	SQLImageSKU *SQLImageSKU

	// SQL Server Management type.
	SQLManagement *SQLManagementMode

	// SQL Server license type.
	SQLServerLicenseType *SQLServerLicenseType

	// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
	SQLVirtualMachineGroupResourceID *string

	// SQL Server configuration management settings.
	ServerConfigurationsManagementSettings *ServerConfigurationsManagementSettings

	// Storage Configuration Settings.
	StorageConfigurationSettings *StorageConfigurationSettings

	// ARM Resource id of underlying virtual machine created from SQL marketplace image.
	VirtualMachineResourceID *string

	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcDomainCredentials *WsfcDomainCredentials

	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcStaticIP *string

	// READ-ONLY; Provisioning state to track the async operation status.
	ProvisioningState *string

	// READ-ONLY; Troubleshooting status
	TroubleshootingStatus *TroubleshootingStatus
}

// ResourceIdentity - Azure Active Directory identity configuration for a resource.
type ResourceIdentity struct {
	// The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal
	// for the resource.
	Type *IdentityType

	// READ-ONLY; The Azure Active Directory principal id.
	PrincipalID *string

	// READ-ONLY; The Azure Active Directory tenant id.
	TenantID *string
}

// SQLConnectivityUpdateSettings - Set the access level and network port settings for SQL Server.
type SQLConnectivityUpdateSettings struct {
	// SQL Server connectivity option.
	ConnectivityType *ConnectivityType

	// SQL Server port.
	Port *int32

	// SQL Server sysadmin login password.
	SQLAuthUpdatePassword *string

	// SQL Server sysadmin login to create.
	SQLAuthUpdateUserName *string
}

// SQLInstanceSettings - Set the server/instance-level settings for SQL Server.
type SQLInstanceSettings struct {
	// SQL Server Collation.
	Collation *string

	// SQL Server IFI.
	IsIfiEnabled *bool

	// SQL Server LPIM.
	IsLpimEnabled *bool

	// SQL Server Optimize for Adhoc workloads.
	IsOptimizeForAdHocWorkloadsEnabled *bool

	// SQL Server MAXDOP.
	MaxDop *int32

	// SQL Server maximum memory.
	MaxServerMemoryMB *int32

	// SQL Server minimum memory.
	MinServerMemoryMB *int32
}

// SQLStorageSettings - Set disk storage settings for SQL Server.
type SQLStorageSettings struct {
	// SQL Server default file path
	DefaultFilePath *string

	// Logical Unit Numbers for the disks.
	Luns []*int32
}

// SQLStorageUpdateSettings - Set disk storage settings for SQL Server.
type SQLStorageUpdateSettings struct {
	// Disk configuration to apply to SQL Server.
	DiskConfigurationType *DiskConfigurationType

	// Virtual machine disk count.
	DiskCount *int32

	// Device id of the first disk to be updated.
	StartingDeviceID *int32
}

// SQLTempDbSettings - Set tempDb storage settings for SQL Server.
type SQLTempDbSettings struct {
	// SQL Server tempdb data file count
	DataFileCount *int32

	// SQL Server tempdb data file size
	DataFileSize *int32

	// SQL Server tempdb data file autoGrowth size
	DataGrowth *int32

	// SQL Server default file path
	DefaultFilePath *string

	// SQL Server tempdb log file size
	LogFileSize *int32

	// SQL Server tempdb log file autoGrowth size
	LogGrowth *int32

	// Logical Unit Numbers for the disks.
	Luns []*int32

	// SQL Server tempdb persist folder choice
	PersistFolder *bool

	// SQL Server tempdb persist folder location
	PersistFolderPath *string
}

// SQLVMTroubleshooting - Details required for SQL VM troubleshooting
type SQLVMTroubleshooting struct {
	// End time in UTC timezone.
	EndTimeUTC *time.Time

	// Troubleshooting properties
	Properties *TroubleshootingAdditionalProperties

	// Start time in UTC timezone.
	StartTimeUTC *time.Time

	// SQL VM troubleshooting scenario.
	TroubleshootingScenario *TroubleshootingScenario

	// READ-ONLY; Virtual machine resource id for response.
	VirtualMachineResourceID *string
}

// SQLVirtualMachine - A SQL virtual machine.
type SQLVirtualMachine struct {
	// REQUIRED; Resource location.
	Location *string

	// Azure Active Directory identity of the server.
	Identity *ResourceIdentity

	// Resource properties.
	Properties *Properties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// SQLVirtualMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginCreateOrUpdate
// method.
type SQLVirtualMachinesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLVirtualMachinesClientBeginDeleteOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginDelete
// method.
type SQLVirtualMachinesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLVirtualMachinesClientBeginRedeployOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginRedeploy
// method.
type SQLVirtualMachinesClientBeginRedeployOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLVirtualMachinesClientBeginStartAssessmentOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginStartAssessment
// method.
type SQLVirtualMachinesClientBeginStartAssessmentOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLVirtualMachinesClientBeginUpdateOptions contains the optional parameters for the SQLVirtualMachinesClient.BeginUpdate
// method.
type SQLVirtualMachinesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLVirtualMachinesClientGetOptions contains the optional parameters for the SQLVirtualMachinesClient.Get method.
type SQLVirtualMachinesClientGetOptions struct {
	// The child resources to include in the response.
	Expand *string
}

// SQLVirtualMachinesClientListByResourceGroupOptions contains the optional parameters for the SQLVirtualMachinesClient.NewListByResourceGroupPager
// method.
type SQLVirtualMachinesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesClientListBySQLVMGroupOptions contains the optional parameters for the SQLVirtualMachinesClient.NewListBySQLVMGroupPager
// method.
type SQLVirtualMachinesClientListBySQLVMGroupOptions struct {
	// placeholder for future optional parameters
}

// SQLVirtualMachinesClientListOptions contains the optional parameters for the SQLVirtualMachinesClient.NewListPager method.
type SQLVirtualMachinesClientListOptions struct {
	// placeholder for future optional parameters
}

// SQLWorkloadTypeUpdateSettings - Set workload type to optimize storage for SQL Server.
type SQLWorkloadTypeUpdateSettings struct {
	// SQL Server workload type.
	SQLWorkloadType *SQLWorkloadType
}

// Schedule - Set assessment schedule for SQL Server.
type Schedule struct {
	// Day of the week to run assessment.
	DayOfWeek *AssessmentDayOfWeek

	// Enable or disable assessment schedule on SQL virtual machine.
	Enable *bool

	// Occurrence of the DayOfWeek day within a month to schedule assessment. Takes values: 1,2,3,4 and -1. Use -1 for last DayOfWeek
	// day of the month
	MonthlyOccurrence *int32

	// Time of the day in HH:mm format. Eg. 17:30
	StartTime *string

	// Number of weeks to schedule between 2 assessment runs. Takes value from 1-6
	WeeklyInterval *int32
}

// ServerConfigurationsManagementSettings - Set the connectivity, storage and workload settings.
type ServerConfigurationsManagementSettings struct {
	// Additional SQL feature settings.
	AdditionalFeaturesServerConfigurations *AdditionalFeaturesServerConfigurations

	// Azure AD authentication Settings.
	AzureAdAuthenticationSettings *AADAuthenticationSettings

	// SQL connectivity type settings.
	SQLConnectivityUpdateSettings *SQLConnectivityUpdateSettings

	// SQL Instance settings.
	SQLInstanceSettings *SQLInstanceSettings

	// SQL storage update settings.
	SQLStorageUpdateSettings *SQLStorageUpdateSettings

	// SQL workload type settings.
	SQLWorkloadTypeUpdateSettings *SQLWorkloadTypeUpdateSettings
}

// StorageConfigurationSettings - Storage Configurations for SQL Data, Log and TempDb.
type StorageConfigurationSettings struct {
	// Disk configuration to apply to SQL Server.
	DiskConfigurationType *DiskConfigurationType

	// SQL Server Data Storage Settings.
	SQLDataSettings *SQLStorageSettings

	// SQL Server Log Storage Settings.
	SQLLogSettings *SQLStorageSettings

	// SQL Server SystemDb Storage on DataPool if true.
	SQLSystemDbOnDataDisk *bool

	// SQL Server TempDb Storage Settings.
	SQLTempDbSettings *SQLTempDbSettings

	// Storage workload type.
	StorageWorkloadType *StorageWorkloadType
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

// TroubleshootClientBeginTroubleshootOptions contains the optional parameters for the TroubleshootClient.BeginTroubleshoot
// method.
type TroubleshootClientBeginTroubleshootOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TroubleshootingAdditionalProperties - SQL VM Troubleshooting additional properties.
type TroubleshootingAdditionalProperties struct {
	// The unhealthy replica information
	UnhealthyReplicaInfo *UnhealthyReplicaInfo
}

// TroubleshootingStatus - Status of last troubleshooting operation on this SQL VM
type TroubleshootingStatus struct {
	// READ-ONLY; End time in UTC timezone.
	EndTimeUTC *time.Time

	// READ-ONLY; Last troubleshooting trigger time in UTC timezone
	LastTriggerTimeUTC *time.Time

	// READ-ONLY; Troubleshooting properties
	Properties *TroubleshootingAdditionalProperties

	// READ-ONLY; Root cause of the issue
	RootCause *string

	// READ-ONLY; Start time in UTC timezone.
	StartTimeUTC *time.Time

	// READ-ONLY; SQL VM troubleshooting scenario.
	TroubleshootingScenario *TroubleshootingScenario
}

// UnhealthyReplicaInfo - SQL VM Troubleshoot UnhealthyReplica scenario information.
type UnhealthyReplicaInfo struct {
	// The name of the availability group
	AvailabilityGroupName *string
}

// Update - An update to a SQL virtual machine.
type Update struct {
	// Resource tags.
	Tags map[string]*string
}

// WsfcDomainCredentials - Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
type WsfcDomainCredentials struct {
	// Cluster bootstrap account password.
	ClusterBootstrapAccountPassword *string

	// Cluster operator account password.
	ClusterOperatorAccountPassword *string

	// SQL service account password.
	SQLServiceAccountPassword *string
}

// WsfcDomainProfile - Active Directory account details to operate Windows Server Failover Cluster.
type WsfcDomainProfile struct {
	// Account name used for creating cluster (at minimum needs permissions to 'Create Computer Objects' in domain).
	ClusterBootstrapAccount *string

	// Account name used for operating cluster i.e. will be part of administrators group on all the participating virtual machines
	// in the cluster.
	ClusterOperatorAccount *string

	// Cluster subnet type.
	ClusterSubnetType *ClusterSubnetType

	// Fully qualified name of the domain.
	DomainFqdn *string

	// Optional path for fileshare witness.
	FileShareWitnessPath *string

	// Organizational Unit path in which the nodes and cluster will be present.
	OuPath *string

	// Account name under which SQL service will run on all participating SQL virtual machines in the cluster.
	SQLServiceAccount *string

	// Primary key of the witness storage account.
	StorageAccountPrimaryKey *string

	// Fully qualified ARM resource id of the witness storage account.
	StorageAccountURL *string
}
