package postgresqlflexibleserversapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/postgresql/mgmt/2020-02-14-preview/postgresqlflexibleservers"
)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serverName string, parameters postgresqlflexibleservers.Server) (result postgresqlflexibleservers.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.Server, err error)
	List(ctx context.Context) (result postgresqlflexibleservers.ServerListResultPage, err error)
	ListComplete(ctx context.Context) (result postgresqlflexibleservers.ServerListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result postgresqlflexibleservers.ServerListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result postgresqlflexibleservers.ServerListResultIterator, err error)
	Restart(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ServersRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ServersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ServersStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, parameters postgresqlflexibleservers.ServerForUpdate) (result postgresqlflexibleservers.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*postgresqlflexibleservers.ServersClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters postgresqlflexibleservers.FirewallRule) (result postgresqlflexibleservers.FirewallRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result postgresqlflexibleservers.FirewallRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result postgresqlflexibleservers.FirewallRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.FirewallRuleListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.FirewallRuleListResultIterator, err error)
}

var _ FirewallRulesClientAPI = (*postgresqlflexibleservers.FirewallRulesClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string) (result postgresqlflexibleservers.Configuration, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ConfigurationListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ConfigurationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters postgresqlflexibleservers.Configuration) (result postgresqlflexibleservers.ConfigurationsUpdateFuture, err error)
}

var _ ConfigurationsClientAPI = (*postgresqlflexibleservers.ConfigurationsClient)(nil)

// ServerKeysClientAPI contains the set of methods on the ServerKeysClient type.
type ServerKeysClientAPI interface {
	CreateOrUpdate(ctx context.Context, serverName string, keyName string, parameters postgresqlflexibleservers.ServerKey, resourceGroupName string) (result postgresqlflexibleservers.ServerKeysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, serverName string, keyName string, resourceGroupName string) (result postgresqlflexibleservers.ServerKeysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, keyName string) (result postgresqlflexibleservers.ServerKey, err error)
	List(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ServerKeyListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serverName string) (result postgresqlflexibleservers.ServerKeyListResultIterator, err error)
}

var _ ServerKeysClientAPI = (*postgresqlflexibleservers.ServerKeysClient)(nil)

// CheckNameAvailabilityClientAPI contains the set of methods on the CheckNameAvailabilityClient type.
type CheckNameAvailabilityClientAPI interface {
	Execute(ctx context.Context, nameAvailabilityRequest postgresqlflexibleservers.NameAvailabilityRequest) (result postgresqlflexibleservers.NameAvailability, err error)
}

var _ CheckNameAvailabilityClientAPI = (*postgresqlflexibleservers.CheckNameAvailabilityClient)(nil)

// LocationBasedCapabilitiesClientAPI contains the set of methods on the LocationBasedCapabilitiesClient type.
type LocationBasedCapabilitiesClientAPI interface {
	Execute(ctx context.Context, locationName string) (result postgresqlflexibleservers.CapabilitiesListResultPage, err error)
	ExecuteComplete(ctx context.Context, locationName string) (result postgresqlflexibleservers.CapabilitiesListResultIterator, err error)
}

var _ LocationBasedCapabilitiesClientAPI = (*postgresqlflexibleservers.LocationBasedCapabilitiesClient)(nil)

// VirtualNetworkSubnetUsageClientAPI contains the set of methods on the VirtualNetworkSubnetUsageClient type.
type VirtualNetworkSubnetUsageClientAPI interface {
	Execute(ctx context.Context, locationName string, parameters postgresqlflexibleservers.VirtualNetworkSubnetUsageParameter) (result postgresqlflexibleservers.VirtualNetworkSubnetUsageResult, err error)
}

var _ VirtualNetworkSubnetUsageClientAPI = (*postgresqlflexibleservers.VirtualNetworkSubnetUsageClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result postgresqlflexibleservers.OperationListResult, err error)
}

var _ OperationsClientAPI = (*postgresqlflexibleservers.OperationsClient)(nil)
