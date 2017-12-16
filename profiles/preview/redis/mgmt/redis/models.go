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

package redis

import original "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2017-02-01/redis"

type FirewallRulesClient = original.FirewallRulesClient
type LinkedServerClient = original.LinkedServerClient
type DayOfWeek = original.DayOfWeek

const (
	Everyday  DayOfWeek = original.Everyday
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
	Weekend   DayOfWeek = original.Weekend
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type RebootType = original.RebootType

const (
	AllNodes      RebootType = original.AllNodes
	PrimaryNode   RebootType = original.PrimaryNode
	SecondaryNode RebootType = original.SecondaryNode
)

type ReplicationRole = original.ReplicationRole

const (
	ReplicationRolePrimary   ReplicationRole = original.ReplicationRolePrimary
	ReplicationRoleSecondary ReplicationRole = original.ReplicationRoleSecondary
)

type SkuFamily = original.SkuFamily

const (
	C SkuFamily = original.C
	P SkuFamily = original.P
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type AccessKeys = original.AccessKeys
type CreateParameters = original.CreateParameters
type CreateProperties = original.CreateProperties
type ExportRDBParameters = original.ExportRDBParameters
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleListResultIterator = original.FirewallRuleListResultIterator
type FirewallRuleListResultPage = original.FirewallRuleListResultPage
type FirewallRuleProperties = original.FirewallRuleProperties
type ForceRebootResponse = original.ForceRebootResponse
type ImportRDBParameters = original.ImportRDBParameters
type LinkedServer = original.LinkedServer
type LinkedServerCreateFuture = original.LinkedServerCreateFuture
type LinkedServerCreateParameters = original.LinkedServerCreateParameters
type LinkedServerCreateProperties = original.LinkedServerCreateProperties
type LinkedServerList = original.LinkedServerList
type LinkedServerProperties = original.LinkedServerProperties
type LinkedServerWithProperties = original.LinkedServerWithProperties
type LinkedServerWithPropertiesList = original.LinkedServerWithPropertiesList
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PatchSchedule = original.PatchSchedule
type Properties = original.Properties
type ProxyResource = original.ProxyResource
type RebootParameters = original.RebootParameters
type RedisCreateFuture = original.RedisCreateFuture
type RedisDeleteFuture = original.RedisDeleteFuture
type RedisExportDataFuture = original.RedisExportDataFuture
type RedisImportDataFuture = original.RedisImportDataFuture
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ResourceProperties = original.ResourceProperties
type ResourceType = original.ResourceType
type ScheduleEntries = original.ScheduleEntries
type ScheduleEntry = original.ScheduleEntry
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpdateParameters = original.UpdateParameters
type UpdateProperties = original.UpdateProperties
type OperationsClient = original.OperationsClient
type PatchSchedulesClient = original.PatchSchedulesClient
type Client = original.Client

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPatchSchedulesClient(subscriptionID string) PatchSchedulesClient {
	return original.NewPatchSchedulesClient(subscriptionID)
}
func NewPatchSchedulesClientWithBaseURI(baseURI string, subscriptionID string) PatchSchedulesClient {
	return original.NewPatchSchedulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedServerClient(subscriptionID string) LinkedServerClient {
	return original.NewLinkedServerClient(subscriptionID)
}
func NewLinkedServerClientWithBaseURI(baseURI string, subscriptionID string) LinkedServerClient {
	return original.NewLinkedServerClientWithBaseURI(baseURI, subscriptionID)
}
