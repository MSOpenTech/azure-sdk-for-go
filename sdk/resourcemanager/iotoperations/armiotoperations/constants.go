// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
	moduleVersion = "v1.0.0"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// BrokerAuthenticationMethod - Broker Authentication Mode
type BrokerAuthenticationMethod string

const (
	// BrokerAuthenticationMethodCustom - Custom authentication configuration.
	BrokerAuthenticationMethodCustom BrokerAuthenticationMethod = "Custom"
	// BrokerAuthenticationMethodServiceAccountToken - ServiceAccountToken authentication configuration.
	BrokerAuthenticationMethodServiceAccountToken BrokerAuthenticationMethod = "ServiceAccountToken"
	// BrokerAuthenticationMethodX509 - X.509 authentication configuration.
	BrokerAuthenticationMethodX509 BrokerAuthenticationMethod = "X509"
)

// PossibleBrokerAuthenticationMethodValues returns the possible values for the BrokerAuthenticationMethod const type.
func PossibleBrokerAuthenticationMethodValues() []BrokerAuthenticationMethod {
	return []BrokerAuthenticationMethod{
		BrokerAuthenticationMethodCustom,
		BrokerAuthenticationMethodServiceAccountToken,
		BrokerAuthenticationMethodX509,
	}
}

// BrokerMemoryProfile - The memory profile settings of the Broker
type BrokerMemoryProfile string

const (
	// BrokerMemoryProfileHigh - High memory profile.
	BrokerMemoryProfileHigh BrokerMemoryProfile = "High"
	// BrokerMemoryProfileLow - Low memory profile.
	BrokerMemoryProfileLow BrokerMemoryProfile = "Low"
	// BrokerMemoryProfileMedium - Medium memory profile.
	BrokerMemoryProfileMedium BrokerMemoryProfile = "Medium"
	// BrokerMemoryProfileTiny - Tiny memory profile.
	BrokerMemoryProfileTiny BrokerMemoryProfile = "Tiny"
)

// PossibleBrokerMemoryProfileValues returns the possible values for the BrokerMemoryProfile const type.
func PossibleBrokerMemoryProfileValues() []BrokerMemoryProfile {
	return []BrokerMemoryProfile{
		BrokerMemoryProfileHigh,
		BrokerMemoryProfileLow,
		BrokerMemoryProfileMedium,
		BrokerMemoryProfileTiny,
	}
}

// BrokerProtocolType - Broker Protocol types
type BrokerProtocolType string

const (
	// BrokerProtocolTypeMqtt - protocol broker
	BrokerProtocolTypeMqtt BrokerProtocolType = "Mqtt"
	// BrokerProtocolTypeWebSockets - protocol websocket
	BrokerProtocolTypeWebSockets BrokerProtocolType = "WebSockets"
)

// PossibleBrokerProtocolTypeValues returns the possible values for the BrokerProtocolType const type.
func PossibleBrokerProtocolTypeValues() []BrokerProtocolType {
	return []BrokerProtocolType{
		BrokerProtocolTypeMqtt,
		BrokerProtocolTypeWebSockets,
	}
}

// BrokerResourceDefinitionMethods - BrokerResourceDefinitionMethods methods allowed
type BrokerResourceDefinitionMethods string

const (
	// BrokerResourceDefinitionMethodsConnect - Allowed Connecting to Broker
	BrokerResourceDefinitionMethodsConnect BrokerResourceDefinitionMethods = "Connect"
	// BrokerResourceDefinitionMethodsPublish - Allowed Publishing to Broker
	BrokerResourceDefinitionMethodsPublish BrokerResourceDefinitionMethods = "Publish"
	// BrokerResourceDefinitionMethodsSubscribe - Allowed Subscribing to Broker
	BrokerResourceDefinitionMethodsSubscribe BrokerResourceDefinitionMethods = "Subscribe"
)

// PossibleBrokerResourceDefinitionMethodsValues returns the possible values for the BrokerResourceDefinitionMethods const type.
func PossibleBrokerResourceDefinitionMethodsValues() []BrokerResourceDefinitionMethods {
	return []BrokerResourceDefinitionMethods{
		BrokerResourceDefinitionMethodsConnect,
		BrokerResourceDefinitionMethodsPublish,
		BrokerResourceDefinitionMethodsSubscribe,
	}
}

// CertManagerIssuerKind - CertManagerIssuerKind properties
type CertManagerIssuerKind string

const (
	// CertManagerIssuerKindClusterIssuer - ClusterIssuer kind.
	CertManagerIssuerKindClusterIssuer CertManagerIssuerKind = "ClusterIssuer"
	// CertManagerIssuerKindIssuer - Issuer kind.
	CertManagerIssuerKindIssuer CertManagerIssuerKind = "Issuer"
)

// PossibleCertManagerIssuerKindValues returns the possible values for the CertManagerIssuerKind const type.
func PossibleCertManagerIssuerKindValues() []CertManagerIssuerKind {
	return []CertManagerIssuerKind{
		CertManagerIssuerKindClusterIssuer,
		CertManagerIssuerKindIssuer,
	}
}

// CloudEventAttributeType - How to map events to the cloud.
type CloudEventAttributeType string

const (
	// CloudEventAttributeTypeCreateOrRemap - CreateOrRemap type
	CloudEventAttributeTypeCreateOrRemap CloudEventAttributeType = "CreateOrRemap"
	// CloudEventAttributeTypePropagate - Propagate type
	CloudEventAttributeTypePropagate CloudEventAttributeType = "Propagate"
)

// PossibleCloudEventAttributeTypeValues returns the possible values for the CloudEventAttributeType const type.
func PossibleCloudEventAttributeTypeValues() []CloudEventAttributeType {
	return []CloudEventAttributeType{
		CloudEventAttributeTypeCreateOrRemap,
		CloudEventAttributeTypePropagate,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DataExplorerAuthMethod - DataflowEndpoint Data Explorer Authentication Method properties
type DataExplorerAuthMethod string

const (
	// DataExplorerAuthMethodSystemAssignedManagedIdentity - SystemAssignedManagedIdentity type
	DataExplorerAuthMethodSystemAssignedManagedIdentity DataExplorerAuthMethod = "SystemAssignedManagedIdentity"
	// DataExplorerAuthMethodUserAssignedManagedIdentity - UserAssignedManagedIdentity type
	DataExplorerAuthMethodUserAssignedManagedIdentity DataExplorerAuthMethod = "UserAssignedManagedIdentity"
)

// PossibleDataExplorerAuthMethodValues returns the possible values for the DataExplorerAuthMethod const type.
func PossibleDataExplorerAuthMethodValues() []DataExplorerAuthMethod {
	return []DataExplorerAuthMethod{
		DataExplorerAuthMethodSystemAssignedManagedIdentity,
		DataExplorerAuthMethodUserAssignedManagedIdentity,
	}
}

// DataLakeStorageAuthMethod - DataflowEndpoint Data Lake Storage Authentication Method properties
type DataLakeStorageAuthMethod string

const (
	// DataLakeStorageAuthMethodAccessToken - AccessToken Option
	DataLakeStorageAuthMethodAccessToken DataLakeStorageAuthMethod = "AccessToken"
	// DataLakeStorageAuthMethodSystemAssignedManagedIdentity - SystemAssignedManagedIdentity type
	DataLakeStorageAuthMethodSystemAssignedManagedIdentity DataLakeStorageAuthMethod = "SystemAssignedManagedIdentity"
	// DataLakeStorageAuthMethodUserAssignedManagedIdentity - UserAssignedManagedIdentity type
	DataLakeStorageAuthMethodUserAssignedManagedIdentity DataLakeStorageAuthMethod = "UserAssignedManagedIdentity"
)

// PossibleDataLakeStorageAuthMethodValues returns the possible values for the DataLakeStorageAuthMethod const type.
func PossibleDataLakeStorageAuthMethodValues() []DataLakeStorageAuthMethod {
	return []DataLakeStorageAuthMethod{
		DataLakeStorageAuthMethodAccessToken,
		DataLakeStorageAuthMethodSystemAssignedManagedIdentity,
		DataLakeStorageAuthMethodUserAssignedManagedIdentity,
	}
}

// DataflowEndpointAuthenticationSaslType - DataflowEndpoint Authentication Sasl Type properties
type DataflowEndpointAuthenticationSaslType string

const (
	// DataflowEndpointAuthenticationSaslTypePlain - PLAIN Type
	DataflowEndpointAuthenticationSaslTypePlain DataflowEndpointAuthenticationSaslType = "Plain"
	// DataflowEndpointAuthenticationSaslTypeScramSHA256 - SCRAM_SHA_256 Type
	DataflowEndpointAuthenticationSaslTypeScramSHA256 DataflowEndpointAuthenticationSaslType = "ScramSha256"
	// DataflowEndpointAuthenticationSaslTypeScramSHA512 - SCRAM_SHA_512 Type
	DataflowEndpointAuthenticationSaslTypeScramSHA512 DataflowEndpointAuthenticationSaslType = "ScramSha512"
)

// PossibleDataflowEndpointAuthenticationSaslTypeValues returns the possible values for the DataflowEndpointAuthenticationSaslType const type.
func PossibleDataflowEndpointAuthenticationSaslTypeValues() []DataflowEndpointAuthenticationSaslType {
	return []DataflowEndpointAuthenticationSaslType{
		DataflowEndpointAuthenticationSaslTypePlain,
		DataflowEndpointAuthenticationSaslTypeScramSHA256,
		DataflowEndpointAuthenticationSaslTypeScramSHA512,
	}
}

// DataflowEndpointFabricPathType - DataflowEndpoint Fabric Path Type properties
type DataflowEndpointFabricPathType string

const (
	// DataflowEndpointFabricPathTypeFiles - FILES Type
	DataflowEndpointFabricPathTypeFiles DataflowEndpointFabricPathType = "Files"
	// DataflowEndpointFabricPathTypeTables - TABLES Type
	DataflowEndpointFabricPathTypeTables DataflowEndpointFabricPathType = "Tables"
)

// PossibleDataflowEndpointFabricPathTypeValues returns the possible values for the DataflowEndpointFabricPathType const type.
func PossibleDataflowEndpointFabricPathTypeValues() []DataflowEndpointFabricPathType {
	return []DataflowEndpointFabricPathType{
		DataflowEndpointFabricPathTypeFiles,
		DataflowEndpointFabricPathTypeTables,
	}
}

// DataflowEndpointKafkaAcks - DataflowEndpoint Kafka Acks properties
type DataflowEndpointKafkaAcks string

const (
	// DataflowEndpointKafkaAcksAll - ALL Option
	DataflowEndpointKafkaAcksAll DataflowEndpointKafkaAcks = "All"
	// DataflowEndpointKafkaAcksOne - ONE Option
	DataflowEndpointKafkaAcksOne DataflowEndpointKafkaAcks = "One"
	// DataflowEndpointKafkaAcksZero - ZERO Option
	DataflowEndpointKafkaAcksZero DataflowEndpointKafkaAcks = "Zero"
)

// PossibleDataflowEndpointKafkaAcksValues returns the possible values for the DataflowEndpointKafkaAcks const type.
func PossibleDataflowEndpointKafkaAcksValues() []DataflowEndpointKafkaAcks {
	return []DataflowEndpointKafkaAcks{
		DataflowEndpointKafkaAcksAll,
		DataflowEndpointKafkaAcksOne,
		DataflowEndpointKafkaAcksZero,
	}
}

// DataflowEndpointKafkaCompression - Kafka endpoint Compression properties
type DataflowEndpointKafkaCompression string

const (
	// DataflowEndpointKafkaCompressionGzip - Gzip Option
	DataflowEndpointKafkaCompressionGzip DataflowEndpointKafkaCompression = "Gzip"
	// DataflowEndpointKafkaCompressionLz4 - LZ4 Option
	DataflowEndpointKafkaCompressionLz4 DataflowEndpointKafkaCompression = "Lz4"
	// DataflowEndpointKafkaCompressionNone - NONE Option
	DataflowEndpointKafkaCompressionNone DataflowEndpointKafkaCompression = "None"
	// DataflowEndpointKafkaCompressionSnappy - SNAPPY Option
	DataflowEndpointKafkaCompressionSnappy DataflowEndpointKafkaCompression = "Snappy"
)

// PossibleDataflowEndpointKafkaCompressionValues returns the possible values for the DataflowEndpointKafkaCompression const type.
func PossibleDataflowEndpointKafkaCompressionValues() []DataflowEndpointKafkaCompression {
	return []DataflowEndpointKafkaCompression{
		DataflowEndpointKafkaCompressionGzip,
		DataflowEndpointKafkaCompressionLz4,
		DataflowEndpointKafkaCompressionNone,
		DataflowEndpointKafkaCompressionSnappy,
	}
}

// DataflowEndpointKafkaPartitionStrategy - DataflowEndpoint Kafka Partition Strategy properties
type DataflowEndpointKafkaPartitionStrategy string

const (
	// DataflowEndpointKafkaPartitionStrategyDefault - Default: Assigns messages to random partitions, using a round-robin algorithm.
	DataflowEndpointKafkaPartitionStrategyDefault DataflowEndpointKafkaPartitionStrategy = "Default"
	// DataflowEndpointKafkaPartitionStrategyProperty - PROPERTY Option
	DataflowEndpointKafkaPartitionStrategyProperty DataflowEndpointKafkaPartitionStrategy = "Property"
	// DataflowEndpointKafkaPartitionStrategyStatic - Static: Assigns messages to a fixed partition number that's derived from
	// the instance ID of the dataflow.
	DataflowEndpointKafkaPartitionStrategyStatic DataflowEndpointKafkaPartitionStrategy = "Static"
	// DataflowEndpointKafkaPartitionStrategyTopic - TOPIC Option
	DataflowEndpointKafkaPartitionStrategyTopic DataflowEndpointKafkaPartitionStrategy = "Topic"
)

// PossibleDataflowEndpointKafkaPartitionStrategyValues returns the possible values for the DataflowEndpointKafkaPartitionStrategy const type.
func PossibleDataflowEndpointKafkaPartitionStrategyValues() []DataflowEndpointKafkaPartitionStrategy {
	return []DataflowEndpointKafkaPartitionStrategy{
		DataflowEndpointKafkaPartitionStrategyDefault,
		DataflowEndpointKafkaPartitionStrategyProperty,
		DataflowEndpointKafkaPartitionStrategyStatic,
		DataflowEndpointKafkaPartitionStrategyTopic,
	}
}

// DataflowMappingType - Dataflow type mapping properties
type DataflowMappingType string

const (
	// DataflowMappingTypeBuiltInFunction - Built in function type
	DataflowMappingTypeBuiltInFunction DataflowMappingType = "BuiltInFunction"
	// DataflowMappingTypeCompute - Compute type
	DataflowMappingTypeCompute DataflowMappingType = "Compute"
	// DataflowMappingTypeNewProperties - New Properties type
	DataflowMappingTypeNewProperties DataflowMappingType = "NewProperties"
	// DataflowMappingTypePassThrough - Pass-through type
	DataflowMappingTypePassThrough DataflowMappingType = "PassThrough"
	// DataflowMappingTypeRename - Rename type
	DataflowMappingTypeRename DataflowMappingType = "Rename"
)

// PossibleDataflowMappingTypeValues returns the possible values for the DataflowMappingType const type.
func PossibleDataflowMappingTypeValues() []DataflowMappingType {
	return []DataflowMappingType{
		DataflowMappingTypeBuiltInFunction,
		DataflowMappingTypeCompute,
		DataflowMappingTypeNewProperties,
		DataflowMappingTypePassThrough,
		DataflowMappingTypeRename,
	}
}

// EndpointType - DataflowEndpoint Type properties
type EndpointType string

const (
	// EndpointTypeDataExplorer - Azure Data Explorer Type
	EndpointTypeDataExplorer EndpointType = "DataExplorer"
	// EndpointTypeDataLakeStorage - Azure Data Lake Type
	EndpointTypeDataLakeStorage EndpointType = "DataLakeStorage"
	// EndpointTypeFabricOneLake - Microsoft Fabric Type
	EndpointTypeFabricOneLake EndpointType = "FabricOneLake"
	// EndpointTypeKafka - Kafka Type
	EndpointTypeKafka EndpointType = "Kafka"
	// EndpointTypeLocalStorage - Local Storage Type
	EndpointTypeLocalStorage EndpointType = "LocalStorage"
	// EndpointTypeMqtt - Broker Type
	EndpointTypeMqtt EndpointType = "Mqtt"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeDataExplorer,
		EndpointTypeDataLakeStorage,
		EndpointTypeFabricOneLake,
		EndpointTypeKafka,
		EndpointTypeLocalStorage,
		EndpointTypeMqtt,
	}
}

// ExtendedLocationType - The enum defining type of ExtendedLocation accepted.
type ExtendedLocationType string

const (
	// ExtendedLocationTypeCustomLocation - CustomLocation type
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
)

// PossibleExtendedLocationTypeValues returns the possible values for the ExtendedLocationType const type.
func PossibleExtendedLocationTypeValues() []ExtendedLocationType {
	return []ExtendedLocationType{
		ExtendedLocationTypeCustomLocation,
	}
}

// FabricOneLakeAuthMethod - DataflowEndpoint Fabric One Lake Authentication Method properties
type FabricOneLakeAuthMethod string

const (
	// FabricOneLakeAuthMethodSystemAssignedManagedIdentity - SystemAssignedManagedIdentity type
	FabricOneLakeAuthMethodSystemAssignedManagedIdentity FabricOneLakeAuthMethod = "SystemAssignedManagedIdentity"
	// FabricOneLakeAuthMethodUserAssignedManagedIdentity - UserAssignedManagedIdentity type
	FabricOneLakeAuthMethodUserAssignedManagedIdentity FabricOneLakeAuthMethod = "UserAssignedManagedIdentity"
)

// PossibleFabricOneLakeAuthMethodValues returns the possible values for the FabricOneLakeAuthMethod const type.
func PossibleFabricOneLakeAuthMethodValues() []FabricOneLakeAuthMethod {
	return []FabricOneLakeAuthMethod{
		FabricOneLakeAuthMethodSystemAssignedManagedIdentity,
		FabricOneLakeAuthMethodUserAssignedManagedIdentity,
	}
}

// FilterType - Filter Type properties
type FilterType string

const (
	// FilterTypeFilter - Filter type
	FilterTypeFilter FilterType = "Filter"
)

// PossibleFilterTypeValues returns the possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{
		FilterTypeFilter,
	}
}

// KafkaAuthMethod - DataflowEndpoint Kafka Authentication Method properties
type KafkaAuthMethod string

const (
	// KafkaAuthMethodAnonymous - Anonymous Option
	KafkaAuthMethodAnonymous KafkaAuthMethod = "Anonymous"
	// KafkaAuthMethodSasl - Sasl Option
	KafkaAuthMethodSasl KafkaAuthMethod = "Sasl"
	// KafkaAuthMethodSystemAssignedManagedIdentity - SystemAssignedManagedIdentity type
	KafkaAuthMethodSystemAssignedManagedIdentity KafkaAuthMethod = "SystemAssignedManagedIdentity"
	// KafkaAuthMethodUserAssignedManagedIdentity - UserAssignedManagedIdentity type
	KafkaAuthMethodUserAssignedManagedIdentity KafkaAuthMethod = "UserAssignedManagedIdentity"
	// KafkaAuthMethodX509Certificate - x509Certificate Option
	KafkaAuthMethodX509Certificate KafkaAuthMethod = "X509Certificate"
)

// PossibleKafkaAuthMethodValues returns the possible values for the KafkaAuthMethod const type.
func PossibleKafkaAuthMethodValues() []KafkaAuthMethod {
	return []KafkaAuthMethod{
		KafkaAuthMethodAnonymous,
		KafkaAuthMethodSasl,
		KafkaAuthMethodSystemAssignedManagedIdentity,
		KafkaAuthMethodUserAssignedManagedIdentity,
		KafkaAuthMethodX509Certificate,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	// ManagedServiceIdentityTypeNone - No managed identity.
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	// ManagedServiceIdentityTypeSystemAndUserAssigned - System and user assigned managed identity.
	ManagedServiceIdentityTypeSystemAndUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	// ManagedServiceIdentityTypeSystemAssigned - System assigned managed identity.
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	// ManagedServiceIdentityTypeUserAssigned - User assigned managed identity.
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAndUserAssigned,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// MqttAuthMethod - DataflowEndpoint Mqtt Authentication Method properties
type MqttAuthMethod string

const (
	// MqttAuthMethodAnonymous - Anonymous Option
	MqttAuthMethodAnonymous MqttAuthMethod = "Anonymous"
	// MqttAuthMethodServiceAccountToken - ServiceAccountToken Option
	MqttAuthMethodServiceAccountToken MqttAuthMethod = "ServiceAccountToken"
	// MqttAuthMethodSystemAssignedManagedIdentity - SystemAssignedManagedIdentity type
	MqttAuthMethodSystemAssignedManagedIdentity MqttAuthMethod = "SystemAssignedManagedIdentity"
	// MqttAuthMethodUserAssignedManagedIdentity - UserAssignedManagedIdentity type
	MqttAuthMethodUserAssignedManagedIdentity MqttAuthMethod = "UserAssignedManagedIdentity"
	// MqttAuthMethodX509Certificate - x509Certificate Option
	MqttAuthMethodX509Certificate MqttAuthMethod = "X509Certificate"
)

// PossibleMqttAuthMethodValues returns the possible values for the MqttAuthMethod const type.
func PossibleMqttAuthMethodValues() []MqttAuthMethod {
	return []MqttAuthMethod{
		MqttAuthMethodAnonymous,
		MqttAuthMethodServiceAccountToken,
		MqttAuthMethodSystemAssignedManagedIdentity,
		MqttAuthMethodUserAssignedManagedIdentity,
		MqttAuthMethodX509Certificate,
	}
}

// MqttRetainType - Broker Retain types
type MqttRetainType string

const (
	// MqttRetainTypeKeep - Retain the messages.
	MqttRetainTypeKeep MqttRetainType = "Keep"
	// MqttRetainTypeNever - Never retain messages.
	MqttRetainTypeNever MqttRetainType = "Never"
)

// PossibleMqttRetainTypeValues returns the possible values for the MqttRetainType const type.
func PossibleMqttRetainTypeValues() []MqttRetainType {
	return []MqttRetainType{
		MqttRetainTypeKeep,
		MqttRetainTypeNever,
	}
}

// OperationType - Dataflow Operation Type properties
type OperationType string

const (
	// OperationTypeBuiltInTransformation - Dataflow BuiltIn Transformation Operation
	OperationTypeBuiltInTransformation OperationType = "BuiltInTransformation"
	// OperationTypeDestination - Dataflow Destination Operation
	OperationTypeDestination OperationType = "Destination"
	// OperationTypeSource - Dataflow Source Operation
	OperationTypeSource OperationType = "Source"
)

// PossibleOperationTypeValues returns the possible values for the OperationType const type.
func PossibleOperationTypeValues() []OperationType {
	return []OperationType{
		OperationTypeBuiltInTransformation,
		OperationTypeDestination,
		OperationTypeSource,
	}
}

// OperationalMode - Mode properties
type OperationalMode string

const (
	// OperationalModeDisabled - Disabled is equivalent to False.
	OperationalModeDisabled OperationalMode = "Disabled"
	// OperationalModeEnabled - Enabled is equivalent to True
	OperationalModeEnabled OperationalMode = "Enabled"
)

// PossibleOperationalModeValues returns the possible values for the OperationalMode const type.
func PossibleOperationalModeValues() []OperationalMode {
	return []OperationalMode{
		OperationalModeDisabled,
		OperationalModeEnabled,
	}
}

// OperatorValues - Valid operators are In, NotIn, Exists and DoesNotExist.
type OperatorValues string

const (
	// OperatorValuesDoesNotExist - DoesNotExist operator.
	OperatorValuesDoesNotExist OperatorValues = "DoesNotExist"
	// OperatorValuesExists - Exists operator.
	OperatorValuesExists OperatorValues = "Exists"
	// OperatorValuesIn - In operator.
	OperatorValuesIn OperatorValues = "In"
	// OperatorValuesNotIn - NotIn operator.
	OperatorValuesNotIn OperatorValues = "NotIn"
)

// PossibleOperatorValuesValues returns the possible values for the OperatorValues const type.
func PossibleOperatorValuesValues() []OperatorValues {
	return []OperatorValues{
		OperatorValuesDoesNotExist,
		OperatorValuesExists,
		OperatorValuesIn,
		OperatorValuesNotIn,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PrivateKeyAlgorithm - Private key algorithm types.
type PrivateKeyAlgorithm string

const (
	// PrivateKeyAlgorithmEc256 - Algorithm - ec256.
	PrivateKeyAlgorithmEc256 PrivateKeyAlgorithm = "Ec256"
	// PrivateKeyAlgorithmEc384 - Algorithm - ec384.
	PrivateKeyAlgorithmEc384 PrivateKeyAlgorithm = "Ec384"
	// PrivateKeyAlgorithmEc521 - Algorithm - ec521.
	PrivateKeyAlgorithmEc521 PrivateKeyAlgorithm = "Ec521"
	// PrivateKeyAlgorithmEd25519 - Algorithm - ed25519.
	PrivateKeyAlgorithmEd25519 PrivateKeyAlgorithm = "Ed25519"
	// PrivateKeyAlgorithmRsa2048 - Algorithm - rsa2048.
	PrivateKeyAlgorithmRsa2048 PrivateKeyAlgorithm = "Rsa2048"
	// PrivateKeyAlgorithmRsa4096 - Algorithm - rsa4096.
	PrivateKeyAlgorithmRsa4096 PrivateKeyAlgorithm = "Rsa4096"
	// PrivateKeyAlgorithmRsa8192 - Algorithm - rsa8192.
	PrivateKeyAlgorithmRsa8192 PrivateKeyAlgorithm = "Rsa8192"
)

// PossiblePrivateKeyAlgorithmValues returns the possible values for the PrivateKeyAlgorithm const type.
func PossiblePrivateKeyAlgorithmValues() []PrivateKeyAlgorithm {
	return []PrivateKeyAlgorithm{
		PrivateKeyAlgorithmEc256,
		PrivateKeyAlgorithmEc384,
		PrivateKeyAlgorithmEc521,
		PrivateKeyAlgorithmEd25519,
		PrivateKeyAlgorithmRsa2048,
		PrivateKeyAlgorithmRsa4096,
		PrivateKeyAlgorithmRsa8192,
	}
}

// PrivateKeyRotationPolicy - Private key rotation policy.
type PrivateKeyRotationPolicy string

const (
	// PrivateKeyRotationPolicyAlways - Rotation Policy - Always.
	PrivateKeyRotationPolicyAlways PrivateKeyRotationPolicy = "Always"
	// PrivateKeyRotationPolicyNever - Rotation Policy - Never.
	PrivateKeyRotationPolicyNever PrivateKeyRotationPolicy = "Never"
)

// PossiblePrivateKeyRotationPolicyValues returns the possible values for the PrivateKeyRotationPolicy const type.
func PossiblePrivateKeyRotationPolicyValues() []PrivateKeyRotationPolicy {
	return []PrivateKeyRotationPolicy{
		PrivateKeyRotationPolicyAlways,
		PrivateKeyRotationPolicyNever,
	}
}

// ProvisioningState - The enum defining status of resource.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Resource has been Accepted.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Resource is Deleting.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - Resource is getting provisioned.
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Resource is Updating.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ServiceType - Kubernetes Service Types supported by Listener
type ServiceType string

const (
	// ServiceTypeClusterIP - Cluster IP Service.
	ServiceTypeClusterIP ServiceType = "ClusterIp"
	// ServiceTypeLoadBalancer - Load Balancer Service.
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	// ServiceTypeNodePort - Node Port Service.
	ServiceTypeNodePort ServiceType = "NodePort"
)

// PossibleServiceTypeValues returns the possible values for the ServiceType const type.
func PossibleServiceTypeValues() []ServiceType {
	return []ServiceType{
		ServiceTypeClusterIP,
		ServiceTypeLoadBalancer,
		ServiceTypeNodePort,
	}
}

// SourceSerializationFormat - Serialization Format properties
type SourceSerializationFormat string

const (
	// SourceSerializationFormatJSON - JSON Format
	SourceSerializationFormatJSON SourceSerializationFormat = "Json"
)

// PossibleSourceSerializationFormatValues returns the possible values for the SourceSerializationFormat const type.
func PossibleSourceSerializationFormatValues() []SourceSerializationFormat {
	return []SourceSerializationFormat{
		SourceSerializationFormatJSON,
	}
}

// StateStoreResourceDefinitionMethods - StateStoreResourceDefinitionMethods methods allowed
type StateStoreResourceDefinitionMethods string

const (
	// StateStoreResourceDefinitionMethodsRead - Get/KeyNotify from Store
	StateStoreResourceDefinitionMethodsRead StateStoreResourceDefinitionMethods = "Read"
	// StateStoreResourceDefinitionMethodsReadWrite - Allowed all operations on Store - Get/KeyNotify/Set/Delete
	StateStoreResourceDefinitionMethodsReadWrite StateStoreResourceDefinitionMethods = "ReadWrite"
	// StateStoreResourceDefinitionMethodsWrite - Set/Delete in Store
	StateStoreResourceDefinitionMethodsWrite StateStoreResourceDefinitionMethods = "Write"
)

// PossibleStateStoreResourceDefinitionMethodsValues returns the possible values for the StateStoreResourceDefinitionMethods const type.
func PossibleStateStoreResourceDefinitionMethodsValues() []StateStoreResourceDefinitionMethods {
	return []StateStoreResourceDefinitionMethods{
		StateStoreResourceDefinitionMethodsRead,
		StateStoreResourceDefinitionMethodsReadWrite,
		StateStoreResourceDefinitionMethodsWrite,
	}
}

// StateStoreResourceKeyTypes - StateStoreResourceKeyTypes properties
type StateStoreResourceKeyTypes string

const (
	// StateStoreResourceKeyTypesBinary - Key type - binary
	StateStoreResourceKeyTypesBinary StateStoreResourceKeyTypes = "Binary"
	// StateStoreResourceKeyTypesPattern - Key type - pattern
	StateStoreResourceKeyTypesPattern StateStoreResourceKeyTypes = "Pattern"
	// StateStoreResourceKeyTypesString - Key type - string
	StateStoreResourceKeyTypesString StateStoreResourceKeyTypes = "String"
)

// PossibleStateStoreResourceKeyTypesValues returns the possible values for the StateStoreResourceKeyTypes const type.
func PossibleStateStoreResourceKeyTypesValues() []StateStoreResourceKeyTypes {
	return []StateStoreResourceKeyTypes{
		StateStoreResourceKeyTypesBinary,
		StateStoreResourceKeyTypesPattern,
		StateStoreResourceKeyTypesString,
	}
}

// SubscriberMessageDropStrategy - The enum defining strategies for dropping messages from the subscriber queue.
type SubscriberMessageDropStrategy string

const (
	// SubscriberMessageDropStrategyDropOldest - The oldest message is dropped.
	SubscriberMessageDropStrategyDropOldest SubscriberMessageDropStrategy = "DropOldest"
	// SubscriberMessageDropStrategyNone - Messages are never dropped.
	SubscriberMessageDropStrategyNone SubscriberMessageDropStrategy = "None"
)

// PossibleSubscriberMessageDropStrategyValues returns the possible values for the SubscriberMessageDropStrategy const type.
func PossibleSubscriberMessageDropStrategyValues() []SubscriberMessageDropStrategy {
	return []SubscriberMessageDropStrategy{
		SubscriberMessageDropStrategyDropOldest,
		SubscriberMessageDropStrategyNone,
	}
}

// TLSCertMethodMode - Broker Authentication Mode
type TLSCertMethodMode string

const (
	// TLSCertMethodModeAutomatic - Automatic TLS server certificate configuration.
	TLSCertMethodModeAutomatic TLSCertMethodMode = "Automatic"
	// TLSCertMethodModeManual - Manual TLS server certificate configuration.
	TLSCertMethodModeManual TLSCertMethodMode = "Manual"
)

// PossibleTLSCertMethodModeValues returns the possible values for the TLSCertMethodMode const type.
func PossibleTLSCertMethodModeValues() []TLSCertMethodMode {
	return []TLSCertMethodMode{
		TLSCertMethodModeAutomatic,
		TLSCertMethodModeManual,
	}
}

// TransformationSerializationFormat - Transformation Format properties
type TransformationSerializationFormat string

const (
	// TransformationSerializationFormatDelta - Delta Format
	TransformationSerializationFormatDelta TransformationSerializationFormat = "Delta"
	// TransformationSerializationFormatJSON - JSON Format
	TransformationSerializationFormatJSON TransformationSerializationFormat = "Json"
	// TransformationSerializationFormatParquet - Parquet Format
	TransformationSerializationFormatParquet TransformationSerializationFormat = "Parquet"
)

// PossibleTransformationSerializationFormatValues returns the possible values for the TransformationSerializationFormat const type.
func PossibleTransformationSerializationFormatValues() []TransformationSerializationFormat {
	return []TransformationSerializationFormat{
		TransformationSerializationFormatDelta,
		TransformationSerializationFormatJSON,
		TransformationSerializationFormatParquet,
	}
}
