//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package kusto

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2021-08-27/kusto"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AzureScaleType = original.AzureScaleType

const (
	AzureScaleTypeAutomatic AzureScaleType = original.AzureScaleTypeAutomatic
	AzureScaleTypeManual    AzureScaleType = original.AzureScaleTypeManual
	AzureScaleTypeNone      AzureScaleType = original.AzureScaleTypeNone
)

type AzureSkuName = original.AzureSkuName

const (
	AzureSkuNameDevNoSLAStandardD11V2 AzureSkuName = original.AzureSkuNameDevNoSLAStandardD11V2
	AzureSkuNameDevNoSLAStandardE2aV4 AzureSkuName = original.AzureSkuNameDevNoSLAStandardE2aV4
	AzureSkuNameStandardD11V2         AzureSkuName = original.AzureSkuNameStandardD11V2
	AzureSkuNameStandardD12V2         AzureSkuName = original.AzureSkuNameStandardD12V2
	AzureSkuNameStandardD13V2         AzureSkuName = original.AzureSkuNameStandardD13V2
	AzureSkuNameStandardD14V2         AzureSkuName = original.AzureSkuNameStandardD14V2
	AzureSkuNameStandardDS13V21TBPS   AzureSkuName = original.AzureSkuNameStandardDS13V21TBPS
	AzureSkuNameStandardDS13V22TBPS   AzureSkuName = original.AzureSkuNameStandardDS13V22TBPS
	AzureSkuNameStandardDS14V23TBPS   AzureSkuName = original.AzureSkuNameStandardDS14V23TBPS
	AzureSkuNameStandardDS14V24TBPS   AzureSkuName = original.AzureSkuNameStandardDS14V24TBPS
	AzureSkuNameStandardE16asV43TBPS  AzureSkuName = original.AzureSkuNameStandardE16asV43TBPS
	AzureSkuNameStandardE16asV44TBPS  AzureSkuName = original.AzureSkuNameStandardE16asV44TBPS
	AzureSkuNameStandardE16aV4        AzureSkuName = original.AzureSkuNameStandardE16aV4
	AzureSkuNameStandardE2aV4         AzureSkuName = original.AzureSkuNameStandardE2aV4
	AzureSkuNameStandardE4aV4         AzureSkuName = original.AzureSkuNameStandardE4aV4
	AzureSkuNameStandardE64iV3        AzureSkuName = original.AzureSkuNameStandardE64iV3
	AzureSkuNameStandardE80idsV4      AzureSkuName = original.AzureSkuNameStandardE80idsV4
	AzureSkuNameStandardE8asV41TBPS   AzureSkuName = original.AzureSkuNameStandardE8asV41TBPS
	AzureSkuNameStandardE8asV42TBPS   AzureSkuName = original.AzureSkuNameStandardE8asV42TBPS
	AzureSkuNameStandardE8aV4         AzureSkuName = original.AzureSkuNameStandardE8aV4
	AzureSkuNameStandardL16s          AzureSkuName = original.AzureSkuNameStandardL16s
	AzureSkuNameStandardL16sV2        AzureSkuName = original.AzureSkuNameStandardL16sV2
	AzureSkuNameStandardL4s           AzureSkuName = original.AzureSkuNameStandardL4s
	AzureSkuNameStandardL8s           AzureSkuName = original.AzureSkuNameStandardL8s
	AzureSkuNameStandardL8sV2         AzureSkuName = original.AzureSkuNameStandardL8sV2
)

type AzureSkuTier = original.AzureSkuTier

const (
	AzureSkuTierBasic    AzureSkuTier = original.AzureSkuTierBasic
	AzureSkuTierStandard AzureSkuTier = original.AzureSkuTierStandard
)

type BlobStorageEventType = original.BlobStorageEventType

const (
	BlobStorageEventTypeMicrosoftStorageBlobCreated BlobStorageEventType = original.BlobStorageEventTypeMicrosoftStorageBlobCreated
	BlobStorageEventTypeMicrosoftStorageBlobRenamed BlobStorageEventType = original.BlobStorageEventTypeMicrosoftStorageBlobRenamed
)

type ClusterNetworkAccessFlag = original.ClusterNetworkAccessFlag

const (
	ClusterNetworkAccessFlagDisabled ClusterNetworkAccessFlag = original.ClusterNetworkAccessFlagDisabled
	ClusterNetworkAccessFlagEnabled  ClusterNetworkAccessFlag = original.ClusterNetworkAccessFlagEnabled
)

type ClusterPrincipalRole = original.ClusterPrincipalRole

const (
	ClusterPrincipalRoleAllDatabasesAdmin  ClusterPrincipalRole = original.ClusterPrincipalRoleAllDatabasesAdmin
	ClusterPrincipalRoleAllDatabasesViewer ClusterPrincipalRole = original.ClusterPrincipalRoleAllDatabasesViewer
)

type Compression = original.Compression

const (
	CompressionGZip Compression = original.CompressionGZip
	CompressionNone Compression = original.CompressionNone
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DatabasePrincipalRole = original.DatabasePrincipalRole

const (
	DatabasePrincipalRoleAdmin              DatabasePrincipalRole = original.DatabasePrincipalRoleAdmin
	DatabasePrincipalRoleIngestor           DatabasePrincipalRole = original.DatabasePrincipalRoleIngestor
	DatabasePrincipalRoleMonitor            DatabasePrincipalRole = original.DatabasePrincipalRoleMonitor
	DatabasePrincipalRoleUnrestrictedViewer DatabasePrincipalRole = original.DatabasePrincipalRoleUnrestrictedViewer
	DatabasePrincipalRoleUser               DatabasePrincipalRole = original.DatabasePrincipalRoleUser
	DatabasePrincipalRoleViewer             DatabasePrincipalRole = original.DatabasePrincipalRoleViewer
)

type DatabasePrincipalType = original.DatabasePrincipalType

const (
	DatabasePrincipalTypeApp   DatabasePrincipalType = original.DatabasePrincipalTypeApp
	DatabasePrincipalTypeGroup DatabasePrincipalType = original.DatabasePrincipalTypeGroup
	DatabasePrincipalTypeUser  DatabasePrincipalType = original.DatabasePrincipalTypeUser
)

type DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKind

const (
	DefaultPrincipalsModificationKindNone    DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKindNone
	DefaultPrincipalsModificationKindReplace DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKindReplace
	DefaultPrincipalsModificationKindUnion   DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKindUnion
)

type EngineType = original.EngineType

const (
	EngineTypeV2 EngineType = original.EngineTypeV2
	EngineTypeV3 EngineType = original.EngineTypeV3
)

type EventGridDataFormat = original.EventGridDataFormat

const (
	EventGridDataFormatAPACHEAVRO EventGridDataFormat = original.EventGridDataFormatAPACHEAVRO
	EventGridDataFormatAVRO       EventGridDataFormat = original.EventGridDataFormatAVRO
	EventGridDataFormatCSV        EventGridDataFormat = original.EventGridDataFormatCSV
	EventGridDataFormatJSON       EventGridDataFormat = original.EventGridDataFormatJSON
	EventGridDataFormatMULTIJSON  EventGridDataFormat = original.EventGridDataFormatMULTIJSON
	EventGridDataFormatORC        EventGridDataFormat = original.EventGridDataFormatORC
	EventGridDataFormatPARQUET    EventGridDataFormat = original.EventGridDataFormatPARQUET
	EventGridDataFormatPSV        EventGridDataFormat = original.EventGridDataFormatPSV
	EventGridDataFormatRAW        EventGridDataFormat = original.EventGridDataFormatRAW
	EventGridDataFormatSCSV       EventGridDataFormat = original.EventGridDataFormatSCSV
	EventGridDataFormatSINGLEJSON EventGridDataFormat = original.EventGridDataFormatSINGLEJSON
	EventGridDataFormatSOHSV      EventGridDataFormat = original.EventGridDataFormatSOHSV
	EventGridDataFormatTSV        EventGridDataFormat = original.EventGridDataFormatTSV
	EventGridDataFormatTSVE       EventGridDataFormat = original.EventGridDataFormatTSVE
	EventGridDataFormatTXT        EventGridDataFormat = original.EventGridDataFormatTXT
	EventGridDataFormatW3CLOGFILE EventGridDataFormat = original.EventGridDataFormatW3CLOGFILE
)

type EventHubDataFormat = original.EventHubDataFormat

const (
	EventHubDataFormatAPACHEAVRO EventHubDataFormat = original.EventHubDataFormatAPACHEAVRO
	EventHubDataFormatAVRO       EventHubDataFormat = original.EventHubDataFormatAVRO
	EventHubDataFormatCSV        EventHubDataFormat = original.EventHubDataFormatCSV
	EventHubDataFormatJSON       EventHubDataFormat = original.EventHubDataFormatJSON
	EventHubDataFormatMULTIJSON  EventHubDataFormat = original.EventHubDataFormatMULTIJSON
	EventHubDataFormatORC        EventHubDataFormat = original.EventHubDataFormatORC
	EventHubDataFormatPARQUET    EventHubDataFormat = original.EventHubDataFormatPARQUET
	EventHubDataFormatPSV        EventHubDataFormat = original.EventHubDataFormatPSV
	EventHubDataFormatRAW        EventHubDataFormat = original.EventHubDataFormatRAW
	EventHubDataFormatSCSV       EventHubDataFormat = original.EventHubDataFormatSCSV
	EventHubDataFormatSINGLEJSON EventHubDataFormat = original.EventHubDataFormatSINGLEJSON
	EventHubDataFormatSOHSV      EventHubDataFormat = original.EventHubDataFormatSOHSV
	EventHubDataFormatTSV        EventHubDataFormat = original.EventHubDataFormatTSV
	EventHubDataFormatTSVE       EventHubDataFormat = original.EventHubDataFormatTSVE
	EventHubDataFormatTXT        EventHubDataFormat = original.EventHubDataFormatTXT
	EventHubDataFormatW3CLOGFILE EventHubDataFormat = original.EventHubDataFormatW3CLOGFILE
)

type IdentityType = original.IdentityType

const (
	IdentityTypeNone                       IdentityType = original.IdentityTypeNone
	IdentityTypeSystemAssigned             IdentityType = original.IdentityTypeSystemAssigned
	IdentityTypeSystemAssignedUserAssigned IdentityType = original.IdentityTypeSystemAssignedUserAssigned
	IdentityTypeUserAssigned               IdentityType = original.IdentityTypeUserAssigned
)

type IotHubDataFormat = original.IotHubDataFormat

const (
	IotHubDataFormatAPACHEAVRO IotHubDataFormat = original.IotHubDataFormatAPACHEAVRO
	IotHubDataFormatAVRO       IotHubDataFormat = original.IotHubDataFormatAVRO
	IotHubDataFormatCSV        IotHubDataFormat = original.IotHubDataFormatCSV
	IotHubDataFormatJSON       IotHubDataFormat = original.IotHubDataFormatJSON
	IotHubDataFormatMULTIJSON  IotHubDataFormat = original.IotHubDataFormatMULTIJSON
	IotHubDataFormatORC        IotHubDataFormat = original.IotHubDataFormatORC
	IotHubDataFormatPARQUET    IotHubDataFormat = original.IotHubDataFormatPARQUET
	IotHubDataFormatPSV        IotHubDataFormat = original.IotHubDataFormatPSV
	IotHubDataFormatRAW        IotHubDataFormat = original.IotHubDataFormatRAW
	IotHubDataFormatSCSV       IotHubDataFormat = original.IotHubDataFormatSCSV
	IotHubDataFormatSINGLEJSON IotHubDataFormat = original.IotHubDataFormatSINGLEJSON
	IotHubDataFormatSOHSV      IotHubDataFormat = original.IotHubDataFormatSOHSV
	IotHubDataFormatTSV        IotHubDataFormat = original.IotHubDataFormatTSV
	IotHubDataFormatTSVE       IotHubDataFormat = original.IotHubDataFormatTSVE
	IotHubDataFormatTXT        IotHubDataFormat = original.IotHubDataFormatTXT
	IotHubDataFormatW3CLOGFILE IotHubDataFormat = original.IotHubDataFormatW3CLOGFILE
)

type Kind = original.Kind

const (
	KindDatabase          Kind = original.KindDatabase
	KindReadOnlyFollowing Kind = original.KindReadOnlyFollowing
	KindReadWrite         Kind = original.KindReadWrite
)

type KindBasicDataConnection = original.KindBasicDataConnection

const (
	KindBasicDataConnectionKindDataConnection KindBasicDataConnection = original.KindBasicDataConnectionKindDataConnection
	KindBasicDataConnectionKindEventGrid      KindBasicDataConnection = original.KindBasicDataConnectionKindEventGrid
	KindBasicDataConnectionKindEventHub       KindBasicDataConnection = original.KindBasicDataConnectionKindEventHub
	KindBasicDataConnectionKindIotHub         KindBasicDataConnection = original.KindBasicDataConnectionKindIotHub
)

type LanguageExtensionName = original.LanguageExtensionName

const (
	LanguageExtensionNamePYTHON LanguageExtensionName = original.LanguageExtensionNamePYTHON
	LanguageExtensionNameR      LanguageExtensionName = original.LanguageExtensionNameR
)

type PrincipalType = original.PrincipalType

const (
	PrincipalTypeApp   PrincipalType = original.PrincipalTypeApp
	PrincipalTypeGroup PrincipalType = original.PrincipalTypeGroup
	PrincipalTypeUser  PrincipalType = original.PrincipalTypeUser
)

type PrincipalsModificationKind = original.PrincipalsModificationKind

const (
	PrincipalsModificationKindNone    PrincipalsModificationKind = original.PrincipalsModificationKindNone
	PrincipalsModificationKindReplace PrincipalsModificationKind = original.PrincipalsModificationKindReplace
	PrincipalsModificationKindUnion   PrincipalsModificationKind = original.PrincipalsModificationKindUnion
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoving    ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateRunning   ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type Reason = original.Reason

const (
	ReasonAlreadyExists Reason = original.ReasonAlreadyExists
	ReasonInvalid       Reason = original.ReasonInvalid
)

type State = original.State

const (
	StateCreating    State = original.StateCreating
	StateDeleted     State = original.StateDeleted
	StateDeleting    State = original.StateDeleting
	StateRunning     State = original.StateRunning
	StateStarting    State = original.StateStarting
	StateStopped     State = original.StateStopped
	StateStopping    State = original.StateStopping
	StateUnavailable State = original.StateUnavailable
	StateUpdating    State = original.StateUpdating
)

type Status = original.Status

const (
	StatusCanceled  Status = original.StatusCanceled
	StatusFailed    Status = original.StatusFailed
	StatusRunning   Status = original.StatusRunning
	StatusSucceeded Status = original.StatusSucceeded
)

type Type = original.Type

const (
	TypeMicrosoftKustoclustersattachedDatabaseConfigurations Type = original.TypeMicrosoftKustoclustersattachedDatabaseConfigurations
	TypeMicrosoftKustoclustersdatabases                      Type = original.TypeMicrosoftKustoclustersdatabases
)

type AcceptedAudiences = original.AcceptedAudiences
type AttachedDatabaseConfiguration = original.AttachedDatabaseConfiguration
type AttachedDatabaseConfigurationListResult = original.AttachedDatabaseConfigurationListResult
type AttachedDatabaseConfigurationProperties = original.AttachedDatabaseConfigurationProperties
type AttachedDatabaseConfigurationsCheckNameRequest = original.AttachedDatabaseConfigurationsCheckNameRequest
type AttachedDatabaseConfigurationsClient = original.AttachedDatabaseConfigurationsClient
type AttachedDatabaseConfigurationsCreateOrUpdateFuture = original.AttachedDatabaseConfigurationsCreateOrUpdateFuture
type AttachedDatabaseConfigurationsDeleteFuture = original.AttachedDatabaseConfigurationsDeleteFuture
type AzureCapacity = original.AzureCapacity
type AzureEntityResource = original.AzureEntityResource
type AzureResourceSku = original.AzureResourceSku
type AzureSku = original.AzureSku
type BaseClient = original.BaseClient
type BasicDataConnection = original.BasicDataConnection
type BasicDatabase = original.BasicDatabase
type CheckNameRequest = original.CheckNameRequest
type CheckNameResult = original.CheckNameResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Cluster = original.Cluster
type ClusterCheckNameRequest = original.ClusterCheckNameRequest
type ClusterListResult = original.ClusterListResult
type ClusterPrincipalAssignment = original.ClusterPrincipalAssignment
type ClusterPrincipalAssignmentCheckNameRequest = original.ClusterPrincipalAssignmentCheckNameRequest
type ClusterPrincipalAssignmentListResult = original.ClusterPrincipalAssignmentListResult
type ClusterPrincipalAssignmentsClient = original.ClusterPrincipalAssignmentsClient
type ClusterPrincipalAssignmentsCreateOrUpdateFuture = original.ClusterPrincipalAssignmentsCreateOrUpdateFuture
type ClusterPrincipalAssignmentsDeleteFuture = original.ClusterPrincipalAssignmentsDeleteFuture
type ClusterPrincipalProperties = original.ClusterPrincipalProperties
type ClusterProperties = original.ClusterProperties
type ClusterUpdate = original.ClusterUpdate
type ClustersAddLanguageExtensionsFuture = original.ClustersAddLanguageExtensionsFuture
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersDetachFollowerDatabasesFuture = original.ClustersDetachFollowerDatabasesFuture
type ClustersDiagnoseVirtualNetworkFuture = original.ClustersDiagnoseVirtualNetworkFuture
type ClustersRemoveLanguageExtensionsFuture = original.ClustersRemoveLanguageExtensionsFuture
type ClustersStartFuture = original.ClustersStartFuture
type ClustersStopFuture = original.ClustersStopFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type DataConnection = original.DataConnection
type DataConnectionCheckNameRequest = original.DataConnectionCheckNameRequest
type DataConnectionListResult = original.DataConnectionListResult
type DataConnectionModel = original.DataConnectionModel
type DataConnectionValidation = original.DataConnectionValidation
type DataConnectionValidationListResult = original.DataConnectionValidationListResult
type DataConnectionValidationResult = original.DataConnectionValidationResult
type DataConnectionsClient = original.DataConnectionsClient
type DataConnectionsCreateOrUpdateFuture = original.DataConnectionsCreateOrUpdateFuture
type DataConnectionsDataConnectionValidationMethodFuture = original.DataConnectionsDataConnectionValidationMethodFuture
type DataConnectionsDeleteFuture = original.DataConnectionsDeleteFuture
type DataConnectionsUpdateFuture = original.DataConnectionsUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseModel = original.DatabaseModel
type DatabasePrincipal = original.DatabasePrincipal
type DatabasePrincipalAssignment = original.DatabasePrincipalAssignment
type DatabasePrincipalAssignmentCheckNameRequest = original.DatabasePrincipalAssignmentCheckNameRequest
type DatabasePrincipalAssignmentListResult = original.DatabasePrincipalAssignmentListResult
type DatabasePrincipalAssignmentsClient = original.DatabasePrincipalAssignmentsClient
type DatabasePrincipalAssignmentsCreateOrUpdateFuture = original.DatabasePrincipalAssignmentsCreateOrUpdateFuture
type DatabasePrincipalAssignmentsDeleteFuture = original.DatabasePrincipalAssignmentsDeleteFuture
type DatabasePrincipalListRequest = original.DatabasePrincipalListRequest
type DatabasePrincipalListResult = original.DatabasePrincipalListResult
type DatabasePrincipalProperties = original.DatabasePrincipalProperties
type DatabaseStatistics = original.DatabaseStatistics
type DatabasesClient = original.DatabasesClient
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type DatabasesUpdateFuture = original.DatabasesUpdateFuture
type DiagnoseVirtualNetworkResult = original.DiagnoseVirtualNetworkResult
type EndpointDependency = original.EndpointDependency
type EndpointDetail = original.EndpointDetail
type EventGridConnectionProperties = original.EventGridConnectionProperties
type EventGridDataConnection = original.EventGridDataConnection
type EventHubConnectionProperties = original.EventHubConnectionProperties
type EventHubDataConnection = original.EventHubDataConnection
type FollowerDatabaseDefinition = original.FollowerDatabaseDefinition
type FollowerDatabaseListResult = original.FollowerDatabaseListResult
type Identity = original.Identity
type IdentityUserAssignedIdentitiesValue = original.IdentityUserAssignedIdentitiesValue
type IotHubConnectionProperties = original.IotHubConnectionProperties
type IotHubDataConnection = original.IotHubDataConnection
type KeyVaultProperties = original.KeyVaultProperties
type LanguageExtension = original.LanguageExtension
type LanguageExtensionsList = original.LanguageExtensionsList
type ListResourceSkusResult = original.ListResourceSkusResult
type ManagedPrivateEndpoint = original.ManagedPrivateEndpoint
type ManagedPrivateEndpointListResult = original.ManagedPrivateEndpointListResult
type ManagedPrivateEndpointProperties = original.ManagedPrivateEndpointProperties
type ManagedPrivateEndpointsCheckNameRequest = original.ManagedPrivateEndpointsCheckNameRequest
type ManagedPrivateEndpointsClient = original.ManagedPrivateEndpointsClient
type ManagedPrivateEndpointsCreateOrUpdateFuture = original.ManagedPrivateEndpointsCreateOrUpdateFuture
type ManagedPrivateEndpointsDeleteFuture = original.ManagedPrivateEndpointsDeleteFuture
type ManagedPrivateEndpointsUpdateFuture = original.ManagedPrivateEndpointsUpdateFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationResultErrorProperties = original.OperationResultErrorProperties
type OperationResultProperties = original.OperationResultProperties
type OperationsClient = original.OperationsClient
type OperationsResultsClient = original.OperationsResultsClient
type OptimizedAutoscale = original.OptimizedAutoscale
type OutboundNetworkDependenciesEndpoint = original.OutboundNetworkDependenciesEndpoint
type OutboundNetworkDependenciesEndpointListResult = original.OutboundNetworkDependenciesEndpointListResult
type OutboundNetworkDependenciesEndpointListResultIterator = original.OutboundNetworkDependenciesEndpointListResultIterator
type OutboundNetworkDependenciesEndpointListResultPage = original.OutboundNetworkDependenciesEndpointListResultPage
type OutboundNetworkDependenciesEndpointProperties = original.OutboundNetworkDependenciesEndpointProperties
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointProperty = original.PrivateEndpointProperty
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionStateProperty = original.PrivateLinkServiceConnectionStateProperty
type ProxyResource = original.ProxyResource
type ReadOnlyFollowingDatabase = original.ReadOnlyFollowingDatabase
type ReadOnlyFollowingDatabaseProperties = original.ReadOnlyFollowingDatabaseProperties
type ReadWriteDatabase = original.ReadWriteDatabase
type ReadWriteDatabaseProperties = original.ReadWriteDatabaseProperties
type Resource = original.Resource
type Script = original.Script
type ScriptCheckNameRequest = original.ScriptCheckNameRequest
type ScriptListResult = original.ScriptListResult
type ScriptProperties = original.ScriptProperties
type ScriptsClient = original.ScriptsClient
type ScriptsCreateOrUpdateFuture = original.ScriptsCreateOrUpdateFuture
type ScriptsDeleteFuture = original.ScriptsDeleteFuture
type ScriptsUpdateFuture = original.ScriptsUpdateFuture
type SkuDescription = original.SkuDescription
type SkuDescriptionList = original.SkuDescriptionList
type SkuLocationInfoItem = original.SkuLocationInfoItem
type SystemData = original.SystemData
type TableLevelSharingProperties = original.TableLevelSharingProperties
type TrackedResource = original.TrackedResource
type TrustedExternalTenant = original.TrustedExternalTenant
type VirtualNetworkConfiguration = original.VirtualNetworkConfiguration

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAttachedDatabaseConfigurationsClient(subscriptionID string) AttachedDatabaseConfigurationsClient {
	return original.NewAttachedDatabaseConfigurationsClient(subscriptionID)
}
func NewAttachedDatabaseConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) AttachedDatabaseConfigurationsClient {
	return original.NewAttachedDatabaseConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterPrincipalAssignmentsClient(subscriptionID string) ClusterPrincipalAssignmentsClient {
	return original.NewClusterPrincipalAssignmentsClient(subscriptionID)
}
func NewClusterPrincipalAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ClusterPrincipalAssignmentsClient {
	return original.NewClusterPrincipalAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataConnectionsClient(subscriptionID string) DataConnectionsClient {
	return original.NewDataConnectionsClient(subscriptionID)
}
func NewDataConnectionsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectionsClient {
	return original.NewDataConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasePrincipalAssignmentsClient(subscriptionID string) DatabasePrincipalAssignmentsClient {
	return original.NewDatabasePrincipalAssignmentsClient(subscriptionID)
}
func NewDatabasePrincipalAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) DatabasePrincipalAssignmentsClient {
	return original.NewDatabasePrincipalAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedPrivateEndpointsClient(subscriptionID string) ManagedPrivateEndpointsClient {
	return original.NewManagedPrivateEndpointsClient(subscriptionID)
}
func NewManagedPrivateEndpointsClientWithBaseURI(baseURI string, subscriptionID string) ManagedPrivateEndpointsClient {
	return original.NewManagedPrivateEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsResultsClient(subscriptionID string) OperationsResultsClient {
	return original.NewOperationsResultsClient(subscriptionID)
}
func NewOperationsResultsClientWithBaseURI(baseURI string, subscriptionID string) OperationsResultsClient {
	return original.NewOperationsResultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutboundNetworkDependenciesEndpointListResultIterator(page OutboundNetworkDependenciesEndpointListResultPage) OutboundNetworkDependenciesEndpointListResultIterator {
	return original.NewOutboundNetworkDependenciesEndpointListResultIterator(page)
}
func NewOutboundNetworkDependenciesEndpointListResultPage(cur OutboundNetworkDependenciesEndpointListResult, getNextPage func(context.Context, OutboundNetworkDependenciesEndpointListResult) (OutboundNetworkDependenciesEndpointListResult, error)) OutboundNetworkDependenciesEndpointListResultPage {
	return original.NewOutboundNetworkDependenciesEndpointListResultPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewScriptsClient(subscriptionID string) ScriptsClient {
	return original.NewScriptsClient(subscriptionID)
}
func NewScriptsClientWithBaseURI(baseURI string, subscriptionID string) ScriptsClient {
	return original.NewScriptsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAzureScaleTypeValues() []AzureScaleType {
	return original.PossibleAzureScaleTypeValues()
}
func PossibleAzureSkuNameValues() []AzureSkuName {
	return original.PossibleAzureSkuNameValues()
}
func PossibleAzureSkuTierValues() []AzureSkuTier {
	return original.PossibleAzureSkuTierValues()
}
func PossibleBlobStorageEventTypeValues() []BlobStorageEventType {
	return original.PossibleBlobStorageEventTypeValues()
}
func PossibleClusterNetworkAccessFlagValues() []ClusterNetworkAccessFlag {
	return original.PossibleClusterNetworkAccessFlagValues()
}
func PossibleClusterPrincipalRoleValues() []ClusterPrincipalRole {
	return original.PossibleClusterPrincipalRoleValues()
}
func PossibleCompressionValues() []Compression {
	return original.PossibleCompressionValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDatabasePrincipalRoleValues() []DatabasePrincipalRole {
	return original.PossibleDatabasePrincipalRoleValues()
}
func PossibleDatabasePrincipalTypeValues() []DatabasePrincipalType {
	return original.PossibleDatabasePrincipalTypeValues()
}
func PossibleDefaultPrincipalsModificationKindValues() []DefaultPrincipalsModificationKind {
	return original.PossibleDefaultPrincipalsModificationKindValues()
}
func PossibleEngineTypeValues() []EngineType {
	return original.PossibleEngineTypeValues()
}
func PossibleEventGridDataFormatValues() []EventGridDataFormat {
	return original.PossibleEventGridDataFormatValues()
}
func PossibleEventHubDataFormatValues() []EventHubDataFormat {
	return original.PossibleEventHubDataFormatValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleIotHubDataFormatValues() []IotHubDataFormat {
	return original.PossibleIotHubDataFormatValues()
}
func PossibleKindBasicDataConnectionValues() []KindBasicDataConnection {
	return original.PossibleKindBasicDataConnectionValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLanguageExtensionNameValues() []LanguageExtensionName {
	return original.PossibleLanguageExtensionNameValues()
}
func PossiblePrincipalTypeValues() []PrincipalType {
	return original.PossiblePrincipalTypeValues()
}
func PossiblePrincipalsModificationKindValues() []PrincipalsModificationKind {
	return original.PossiblePrincipalsModificationKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
