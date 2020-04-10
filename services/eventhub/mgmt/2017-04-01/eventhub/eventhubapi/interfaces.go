package eventhubapi

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
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result eventhub.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result eventhub.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*eventhub.OperationsClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters eventhub.CheckNameAvailabilityParameter) (result eventhub.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.EHNamespace) (result eventhub.NamespacesCreateOrUpdateFuture, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (result eventhub.AuthorizationRule, err error)
	CreateOrUpdateNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.NetworkRuleSet) (result eventhub.NetworkRuleSet, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NamespacesDeleteFuture, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.EHNamespace, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result eventhub.AuthorizationRule, err error)
	GetMessagingPlan(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.MessagingPlan, err error)
	GetNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NetworkRuleSet, err error)
	List(ctx context.Context) (result eventhub.EHNamespaceListResultPage, err error)
	ListComplete(ctx context.Context) (result eventhub.EHNamespaceListResultIterator, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.AuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.AuthorizationRuleListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result eventhub.EHNamespaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result eventhub.EHNamespaceListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result eventhub.AccessKeys, err error)
	ListNetworkRuleSets(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NetworkRuleSetListResultPage, err error)
	ListNetworkRuleSetsComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NetworkRuleSetListResultIterator, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters eventhub.RegenerateAccessKeyParameters) (result eventhub.AccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.EHNamespace) (result eventhub.EHNamespace, err error)
}

var _ NamespacesClientAPI = (*eventhub.NamespacesClient)(nil)

// DisasterRecoveryConfigsClientAPI contains the set of methods on the DisasterRecoveryConfigsClient type.
type DisasterRecoveryConfigsClientAPI interface {
	BreakPairing(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error)
	CheckNameAvailability(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.CheckNameAvailabilityParameter) (result eventhub.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters eventhub.ArmDisasterRecovery) (result eventhub.ArmDisasterRecovery, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error)
	FailOver(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result eventhub.ArmDisasterRecovery, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string) (result eventhub.AuthorizationRule, err error)
	List(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.ArmDisasterRecoveryListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.ArmDisasterRecoveryListResultIterator, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result eventhub.AuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result eventhub.AuthorizationRuleListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string) (result eventhub.AccessKeys, err error)
}

var _ DisasterRecoveryConfigsClientAPI = (*eventhub.DisasterRecoveryConfigsClient)(nil)

// EventHubsClientAPI contains the set of methods on the EventHubsClient type.
type EventHubsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters eventhub.Model) (result eventhub.Model, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.AuthorizationRule) (result eventhub.AuthorizationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.Model, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.AuthorizationRule, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.AuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.AuthorizationRuleListResultIterator, err error)
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string, skip *int32, top *int32) (result eventhub.ListResultPage, err error)
	ListByNamespaceComplete(ctx context.Context, resourceGroupName string, namespaceName string, skip *int32, top *int32) (result eventhub.ListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.AccessKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.RegenerateAccessKeyParameters) (result eventhub.AccessKeys, err error)
}

var _ EventHubsClientAPI = (*eventhub.EventHubsClient)(nil)

// ConsumerGroupsClientAPI contains the set of methods on the ConsumerGroupsClient type.
type ConsumerGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters eventhub.ConsumerGroup) (result eventhub.ConsumerGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result eventhub.ConsumerGroup, err error)
	ListByEventHub(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, skip *int32, top *int32) (result eventhub.ConsumerGroupListResultPage, err error)
	ListByEventHubComplete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, skip *int32, top *int32) (result eventhub.ConsumerGroupListResultIterator, err error)
}

var _ ConsumerGroupsClientAPI = (*eventhub.ConsumerGroupsClient)(nil)

// RegionsClientAPI contains the set of methods on the RegionsClient type.
type RegionsClientAPI interface {
	ListBySku(ctx context.Context, sku string) (result eventhub.MessagingRegionsListResultPage, err error)
	ListBySkuComplete(ctx context.Context, sku string) (result eventhub.MessagingRegionsListResultIterator, err error)
}

var _ RegionsClientAPI = (*eventhub.RegionsClient)(nil)
