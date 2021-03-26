package policyinsightsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2019-10-01-preview/policyinsights"
	"github.com/Azure/go-autorest/autorest/date"
)

// PolicyTrackedResourcesClientAPI contains the set of methods on the PolicyTrackedResourcesClient type.
type PolicyTrackedResourcesClientAPI interface {
	ListQueryResultsForManagementGroup(ctx context.Context, managementGroupName string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForManagementGroupComplete(ctx context.Context, managementGroupName string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsIterator, err error)
	ListQueryResultsForResource(ctx context.Context, resourceID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForResourceComplete(ctx context.Context, resourceID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsIterator, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, resourceGroupName string, subscriptionID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForResourceGroupComplete(ctx context.Context, resourceGroupName string, subscriptionID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsIterator, err error)
	ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForSubscriptionComplete(ctx context.Context, subscriptionID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsIterator, err error)
}

var _ PolicyTrackedResourcesClientAPI = (*policyinsights.PolicyTrackedResourcesClient)(nil)

// RemediationsClientAPI contains the set of methods on the RemediationsClient type.
type RemediationsClientAPI interface {
	CancelAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string) (result policyinsights.Remediation, err error)
	CancelAtResource(ctx context.Context, resourceID string, remediationName string) (result policyinsights.Remediation, err error)
	CancelAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string) (result policyinsights.Remediation, err error)
	CancelAtSubscription(ctx context.Context, subscriptionID string, remediationName string) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtResource(ctx context.Context, resourceID string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtSubscription(ctx context.Context, subscriptionID string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	DeleteAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string) (result policyinsights.Remediation, err error)
	DeleteAtResource(ctx context.Context, resourceID string, remediationName string) (result policyinsights.Remediation, err error)
	DeleteAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string) (result policyinsights.Remediation, err error)
	DeleteAtSubscription(ctx context.Context, subscriptionID string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtResource(ctx context.Context, resourceID string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtSubscription(ctx context.Context, subscriptionID string, remediationName string) (result policyinsights.Remediation, err error)
	ListDeploymentsAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtManagementGroupComplete(ctx context.Context, managementGroupID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultIterator, err error)
	ListDeploymentsAtResource(ctx context.Context, resourceID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtResourceComplete(ctx context.Context, resourceID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultIterator, err error)
	ListDeploymentsAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtResourceGroupComplete(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultIterator, err error)
	ListDeploymentsAtSubscription(ctx context.Context, subscriptionID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtSubscriptionComplete(ctx context.Context, subscriptionID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultIterator, err error)
	ListForManagementGroup(ctx context.Context, managementGroupID string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForManagementGroupComplete(ctx context.Context, managementGroupID string, top *int32, filter string) (result policyinsights.RemediationListResultIterator, err error)
	ListForResource(ctx context.Context, resourceID string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceID string, top *int32, filter string) (result policyinsights.RemediationListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, filter string) (result policyinsights.RemediationListResultIterator, err error)
	ListForSubscription(ctx context.Context, subscriptionID string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForSubscriptionComplete(ctx context.Context, subscriptionID string, top *int32, filter string) (result policyinsights.RemediationListResultIterator, err error)
}

var _ RemediationsClientAPI = (*policyinsights.RemediationsClient)(nil)

// PolicyEventsClientAPI contains the set of methods on the PolicyEventsClient type.
type PolicyEventsClientAPI interface {
	ListQueryResultsForManagementGroup(ctx context.Context, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForManagementGroupComplete(ctx context.Context, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForPolicyDefinition(ctx context.Context, subscriptionID string, policyDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForPolicyDefinitionComplete(ctx context.Context, subscriptionID string, policyDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForPolicySetDefinition(ctx context.Context, subscriptionID string, policySetDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForPolicySetDefinitionComplete(ctx context.Context, subscriptionID string, policySetDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForResource(ctx context.Context, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, expand string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForResourceComplete(ctx context.Context, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, expand string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForResourceGroupComplete(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForResourceGroupLevelPolicyAssignmentComplete(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForSubscriptionComplete(ctx context.Context, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
	ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx context.Context, subscriptionID string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsPage, err error)
	ListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx context.Context, subscriptionID string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyEventsQueryResultsIterator, err error)
}

var _ PolicyEventsClientAPI = (*policyinsights.PolicyEventsClient)(nil)

// PolicyStatesClientAPI contains the set of methods on the PolicyStatesClient type.
type PolicyStatesClientAPI interface {
	ListQueryResultsForManagementGroup(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForManagementGroupComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForPolicyDefinition(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policyDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForPolicyDefinitionComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policyDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForPolicySetDefinition(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policySetDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForPolicySetDefinitionComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policySetDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForResource(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, expand string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForResourceComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, expand string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForResourceGroupComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForResourceGroupLevelPolicyAssignmentComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForSubscription(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForSubscriptionComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsPage, err error)
	ListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, skipToken string) (result policyinsights.PolicyStatesQueryResultsIterator, err error)
	SummarizeForManagementGroup(ctx context.Context, managementGroupName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForPolicyDefinition(ctx context.Context, subscriptionID string, policyDefinitionName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForPolicySetDefinition(ctx context.Context, subscriptionID string, policySetDefinitionName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForResource(ctx context.Context, resourceID string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForResourceGroupLevelPolicyAssignment(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForSubscription(ctx context.Context, subscriptionID string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForSubscriptionLevelPolicyAssignment(ctx context.Context, subscriptionID string, policyAssignmentName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	TriggerResourceGroupEvaluation(ctx context.Context, subscriptionID string, resourceGroupName string) (result policyinsights.PolicyStatesTriggerResourceGroupEvaluationFuture, err error)
	TriggerSubscriptionEvaluation(ctx context.Context, subscriptionID string) (result policyinsights.PolicyStatesTriggerSubscriptionEvaluationFuture, err error)
}

var _ PolicyStatesClientAPI = (*policyinsights.PolicyStatesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result policyinsights.OperationsListResults, err error)
}

var _ OperationsClientAPI = (*policyinsights.OperationsClient)(nil)

// PolicyMetadataClientAPI contains the set of methods on the PolicyMetadataClient type.
type PolicyMetadataClientAPI interface {
	GetResource(ctx context.Context, resourceName string) (result policyinsights.PolicyMetadata, err error)
	List(ctx context.Context, top *int32) (result policyinsights.PolicyMetadataCollectionPage, err error)
	ListComplete(ctx context.Context, top *int32) (result policyinsights.PolicyMetadataCollectionIterator, err error)
}

var _ PolicyMetadataClientAPI = (*policyinsights.PolicyMetadataClient)(nil)
