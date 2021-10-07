package postgresqlapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serverName string, parameters postgresql.ServerForCreate) (result postgresql.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.Server, err error)
	List(ctx context.Context) (result postgresql.ServerListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result postgresql.ServerListResult, err error)
	Restart(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServersRestartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, parameters postgresql.ServerUpdateParameters) (result postgresql.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*postgresql.ServersClient)(nil)

// ReplicasClientAPI contains the set of methods on the ReplicasClient type.
type ReplicasClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServerListResult, err error)
}

var _ ReplicasClientAPI = (*postgresql.ReplicasClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters postgresql.FirewallRule) (result postgresql.FirewallRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result postgresql.FirewallRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result postgresql.FirewallRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.FirewallRuleListResult, err error)
}

var _ FirewallRulesClientAPI = (*postgresql.FirewallRulesClient)(nil)

// VirtualNetworkRulesClientAPI contains the set of methods on the VirtualNetworkRulesClient type.
type VirtualNetworkRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters postgresql.VirtualNetworkRule) (result postgresql.VirtualNetworkRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result postgresql.VirtualNetworkRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result postgresql.VirtualNetworkRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.VirtualNetworkRuleListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.VirtualNetworkRuleListResultIterator, err error)
}

var _ VirtualNetworkRulesClientAPI = (*postgresql.VirtualNetworkRulesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters postgresql.Database) (result postgresql.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result postgresql.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result postgresql.Database, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.DatabaseListResult, err error)
}

var _ DatabasesClientAPI = (*postgresql.DatabasesClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters postgresql.Configuration) (result postgresql.ConfigurationsCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string) (result postgresql.Configuration, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ConfigurationListResult, err error)
}

var _ ConfigurationsClientAPI = (*postgresql.ConfigurationsClient)(nil)

// LogFilesClientAPI contains the set of methods on the LogFilesClient type.
type LogFilesClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.LogFileListResult, err error)
}

var _ LogFilesClientAPI = (*postgresql.LogFilesClient)(nil)

// ServerAdministratorsClientAPI contains the set of methods on the ServerAdministratorsClient type.
type ServerAdministratorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties postgresql.ServerAdministratorResource) (result postgresql.ServerAdministratorsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServerAdministratorsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServerAdministratorResource, err error)
	List(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServerAdministratorResourceListResult, err error)
}

var _ ServerAdministratorsClientAPI = (*postgresql.ServerAdministratorsClient)(nil)

// LocationBasedPerformanceTierClientAPI contains the set of methods on the LocationBasedPerformanceTierClient type.
type LocationBasedPerformanceTierClientAPI interface {
	List(ctx context.Context, locationName string) (result postgresql.PerformanceTierListResult, err error)
}

var _ LocationBasedPerformanceTierClientAPI = (*postgresql.LocationBasedPerformanceTierClient)(nil)

// CheckNameAvailabilityClientAPI contains the set of methods on the CheckNameAvailabilityClient type.
type CheckNameAvailabilityClientAPI interface {
	Execute(ctx context.Context, nameAvailabilityRequest postgresql.NameAvailabilityRequest) (result postgresql.NameAvailability, err error)
}

var _ CheckNameAvailabilityClientAPI = (*postgresql.CheckNameAvailabilityClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result postgresql.OperationListResult, err error)
}

var _ OperationsClientAPI = (*postgresql.OperationsClient)(nil)

// ServerSecurityAlertPoliciesClientAPI contains the set of methods on the ServerSecurityAlertPoliciesClient type.
type ServerSecurityAlertPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters postgresql.ServerSecurityAlertPolicy) (result postgresql.ServerSecurityAlertPoliciesCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result postgresql.ServerSecurityAlertPolicy, err error)
}

var _ ServerSecurityAlertPoliciesClientAPI = (*postgresql.ServerSecurityAlertPoliciesClient)(nil)
