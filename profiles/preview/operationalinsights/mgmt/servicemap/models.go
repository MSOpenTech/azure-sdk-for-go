// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package servicemap

import original "github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2015-11-01-preview/servicemap"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type MachineGroupsClient = original.MachineGroupsClient
type MapsClient = original.MapsClient
type PortsClient = original.PortsClient
type ProcessesClient = original.ProcessesClient
type SummariesClient = original.SummariesClient
type ClientGroupsClient = original.ClientGroupsClient
type MachinesClient = original.MachinesClient
type Accuracy = original.Accuracy

const (
	Actual    Accuracy = original.Actual
	Estimated Accuracy = original.Estimated
)

type Bitness = original.Bitness

const (
	SixFourbit  Bitness = original.SixFourbit
	ThreeTwobit Bitness = original.ThreeTwobit
)

type ConnectionFailureState = original.ConnectionFailureState

const (
	Failed ConnectionFailureState = original.Failed
	Mixed  ConnectionFailureState = original.Mixed
	Ok     ConnectionFailureState = original.Ok
)

type HypervisorType = original.HypervisorType

const (
	Hyperv  HypervisorType = original.Hyperv
	Unknown HypervisorType = original.Unknown
)

type Kind = original.Kind

const (
	KindRefmachine          Kind = original.KindRefmachine
	KindRefmachinewithhints Kind = original.KindRefmachinewithhints
	KindRefport             Kind = original.KindRefport
	KindRefprocess          Kind = original.KindRefprocess
	KindResourceReference   Kind = original.KindResourceReference
)

type KindBasicCoreResource = original.KindBasicCoreResource

const (
	KindClientGroup  KindBasicCoreResource = original.KindClientGroup
	KindCoreResource KindBasicCoreResource = original.KindCoreResource
	KindMachine      KindBasicCoreResource = original.KindMachine
	KindMachineGroup KindBasicCoreResource = original.KindMachineGroup
	KindPort         KindBasicCoreResource = original.KindPort
	KindProcess      KindBasicCoreResource = original.KindProcess
)

type KindBasicMapRequest = original.KindBasicMapRequest

const (
	KindMapmachineGroupDependency  KindBasicMapRequest = original.KindMapmachineGroupDependency
	KindMapRequest                 KindBasicMapRequest = original.KindMapRequest
	KindMapsingleMachineDependency KindBasicMapRequest = original.KindMapsingleMachineDependency
)

type KindBasicRelationship = original.KindBasicRelationship

const (
	KindRelacceptor   KindBasicRelationship = original.KindRelacceptor
	KindRelationship  KindBasicRelationship = original.KindRelationship
	KindRelconnection KindBasicRelationship = original.KindRelconnection
)

type MachineRebootStatus = original.MachineRebootStatus

const (
	MachineRebootStatusNotRebooted MachineRebootStatus = original.MachineRebootStatusNotRebooted
	MachineRebootStatusRebooted    MachineRebootStatus = original.MachineRebootStatusRebooted
	MachineRebootStatusUnknown     MachineRebootStatus = original.MachineRebootStatusUnknown
)

type MonitoringState = original.MonitoringState

const (
	Discovered MonitoringState = original.Discovered
	Monitored  MonitoringState = original.Monitored
)

type OperatingSystemFamily = original.OperatingSystemFamily

const (
	OperatingSystemFamilyAix     OperatingSystemFamily = original.OperatingSystemFamilyAix
	OperatingSystemFamilyLinux   OperatingSystemFamily = original.OperatingSystemFamilyLinux
	OperatingSystemFamilySolaris OperatingSystemFamily = original.OperatingSystemFamilySolaris
	OperatingSystemFamilyUnknown OperatingSystemFamily = original.OperatingSystemFamilyUnknown
	OperatingSystemFamilyWindows OperatingSystemFamily = original.OperatingSystemFamilyWindows
)

type ProcessRole = original.ProcessRole

const (
	AppServer      ProcessRole = original.AppServer
	DatabaseServer ProcessRole = original.DatabaseServer
	LdapServer     ProcessRole = original.LdapServer
	SmbServer      ProcessRole = original.SmbServer
	WebServer      ProcessRole = original.WebServer
)

type VirtualizationState = original.VirtualizationState

const (
	VirtualizationStateHypervisor VirtualizationState = original.VirtualizationStateHypervisor
	VirtualizationStatePhysical   VirtualizationState = original.VirtualizationStatePhysical
	VirtualizationStateUnknown    VirtualizationState = original.VirtualizationStateUnknown
	VirtualizationStateVirtual    VirtualizationState = original.VirtualizationStateVirtual
)

type VirtualMachineType = original.VirtualMachineType

const (
	VirtualMachineTypeHyperv    VirtualMachineType = original.VirtualMachineTypeHyperv
	VirtualMachineTypeLdom      VirtualMachineType = original.VirtualMachineTypeLdom
	VirtualMachineTypeLpar      VirtualMachineType = original.VirtualMachineTypeLpar
	VirtualMachineTypeUnknown   VirtualMachineType = original.VirtualMachineTypeUnknown
	VirtualMachineTypeVirtualPc VirtualMachineType = original.VirtualMachineTypeVirtualPc
	VirtualMachineTypeVmware    VirtualMachineType = original.VirtualMachineTypeVmware
	VirtualMachineTypeXen       VirtualMachineType = original.VirtualMachineTypeXen
)

type Acceptor = original.Acceptor
type AcceptorProperties = original.AcceptorProperties
type AgentConfiguration = original.AgentConfiguration
type ClientGroup = original.ClientGroup
type ClientGroupMember = original.ClientGroupMember
type ClientGroupMemberProperties = original.ClientGroupMemberProperties
type ClientGroupMembersCollection = original.ClientGroupMembersCollection
type ClientGroupMembersCollectionIterator = original.ClientGroupMembersCollectionIterator
type ClientGroupMembersCollectionPage = original.ClientGroupMembersCollectionPage
type ClientGroupMembersCount = original.ClientGroupMembersCount
type ClientGroupProperties = original.ClientGroupProperties
type Connection = original.Connection
type ConnectionCollection = original.ConnectionCollection
type ConnectionCollectionIterator = original.ConnectionCollectionIterator
type ConnectionCollectionPage = original.ConnectionCollectionPage
type ConnectionProperties = original.ConnectionProperties
type BasicCoreResource = original.BasicCoreResource
type CoreResource = original.CoreResource
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type HypervisorConfiguration = original.HypervisorConfiguration
type Ipv4NetworkInterface = original.Ipv4NetworkInterface
type Ipv6NetworkInterface = original.Ipv6NetworkInterface
type Liveness = original.Liveness
type Machine = original.Machine
type MachineCollection = original.MachineCollection
type MachineCollectionIterator = original.MachineCollectionIterator
type MachineCollectionPage = original.MachineCollectionPage
type MachineCountsByOperatingSystem = original.MachineCountsByOperatingSystem
type MachineGroup = original.MachineGroup
type MachineGroupCollection = original.MachineGroupCollection
type MachineGroupCollectionIterator = original.MachineGroupCollectionIterator
type MachineGroupCollectionPage = original.MachineGroupCollectionPage
type MachineGroupMapRequest = original.MachineGroupMapRequest
type MachineGroupProperties = original.MachineGroupProperties
type MachineProperties = original.MachineProperties
type MachineReference = original.MachineReference
type MachineReferenceWithHints = original.MachineReferenceWithHints
type MachineReferenceWithHintsProperties = original.MachineReferenceWithHintsProperties
type MachineResourcesConfiguration = original.MachineResourcesConfiguration
type MachinesSummary = original.MachinesSummary
type MachinesSummaryProperties = original.MachinesSummaryProperties
type Map = original.Map
type MapEdges = original.MapEdges
type MapNodes = original.MapNodes
type BasicMapRequest = original.BasicMapRequest
type MapRequest = original.MapRequest
type MapResponse = original.MapResponse
type NetworkConfiguration = original.NetworkConfiguration
type OperatingSystemConfiguration = original.OperatingSystemConfiguration
type Port = original.Port
type PortCollection = original.PortCollection
type PortCollectionIterator = original.PortCollectionIterator
type PortCollectionPage = original.PortCollectionPage
type PortProperties = original.PortProperties
type PortReference = original.PortReference
type PortReferenceProperties = original.PortReferenceProperties
type Process = original.Process
type ProcessCollection = original.ProcessCollection
type ProcessCollectionIterator = original.ProcessCollectionIterator
type ProcessCollectionPage = original.ProcessCollectionPage
type ProcessDetails = original.ProcessDetails
type ProcessProperties = original.ProcessProperties
type ProcessReference = original.ProcessReference
type ProcessReferenceProperties = original.ProcessReferenceProperties
type ProcessUser = original.ProcessUser
type BasicRelationship = original.BasicRelationship
type Relationship = original.Relationship
type RelationshipProperties = original.RelationshipProperties
type Resource = original.Resource
type BasicResourceReference = original.BasicResourceReference
type ResourceReference = original.ResourceReference
type SingleMachineDependencyMapRequest = original.SingleMachineDependencyMapRequest
type Summary = original.Summary
type SummaryProperties = original.SummaryProperties
type Timezone = original.Timezone
type VirtualMachineConfiguration = original.VirtualMachineConfiguration

func NewClientGroupsClient(subscriptionID string) ClientGroupsClient {
	return original.NewClientGroupsClient(subscriptionID)
}
func NewClientGroupsClientWithBaseURI(baseURI string, subscriptionID string) ClientGroupsClient {
	return original.NewClientGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMachinesClient(subscriptionID string) MachinesClient {
	return original.NewMachinesClient(subscriptionID)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewMachineGroupsClient(subscriptionID string) MachineGroupsClient {
	return original.NewMachineGroupsClient(subscriptionID)
}
func NewMachineGroupsClientWithBaseURI(baseURI string, subscriptionID string) MachineGroupsClient {
	return original.NewMachineGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMapsClient(subscriptionID string) MapsClient {
	return original.NewMapsClient(subscriptionID)
}
func NewMapsClientWithBaseURI(baseURI string, subscriptionID string) MapsClient {
	return original.NewMapsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPortsClient(subscriptionID string) PortsClient {
	return original.NewPortsClient(subscriptionID)
}
func NewPortsClientWithBaseURI(baseURI string, subscriptionID string) PortsClient {
	return original.NewPortsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProcessesClient(subscriptionID string) ProcessesClient {
	return original.NewProcessesClient(subscriptionID)
}
func NewProcessesClientWithBaseURI(baseURI string, subscriptionID string) ProcessesClient {
	return original.NewProcessesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSummariesClient(subscriptionID string) SummariesClient {
	return original.NewSummariesClient(subscriptionID)
}
func NewSummariesClientWithBaseURI(baseURI string, subscriptionID string) SummariesClient {
	return original.NewSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
