package dynatraceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/dynatrace/mgmt/2021-09-01-preview/dynatrace"
)

// MonitorsClientAPI contains the set of methods on the MonitorsClient type.
type MonitorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, resource dynatrace.MonitorResource) (result dynatrace.MonitorsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.MonitorsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.MonitorResource, err error)
	GetAccountCredentials(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.AccountInfoSecure, err error)
	GetSSODetails(ctx context.Context, resourceGroupName string, monitorName string, request *dynatrace.SSODetailsRequest) (result dynatrace.SSODetailsResponse, err error)
	GetVMHostPayload(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.VMExtensionPayload, err error)
	ListAppServices(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.AppServiceListResponsePage, err error)
	ListAppServicesComplete(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.AppServiceListResponseIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result dynatrace.MonitorResourceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result dynatrace.MonitorResourceListResultIterator, err error)
	ListBySubscriptionID(ctx context.Context) (result dynatrace.MonitorResourceListResultPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context) (result dynatrace.MonitorResourceListResultIterator, err error)
	ListHosts(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.VMHostsListResponsePage, err error)
	ListHostsComplete(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.VMHostsListResponseIterator, err error)
	ListLinkableEnvironments(ctx context.Context, resourceGroupName string, monitorName string, request dynatrace.LinkableEnvironmentRequest) (result dynatrace.LinkableEnvironmentListResponsePage, err error)
	ListLinkableEnvironmentsComplete(ctx context.Context, resourceGroupName string, monitorName string, request dynatrace.LinkableEnvironmentRequest) (result dynatrace.LinkableEnvironmentListResponseIterator, err error)
	ListMonitoredResources(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.MonitoredResourceListResponsePage, err error)
	ListMonitoredResourcesComplete(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.MonitoredResourceListResponseIterator, err error)
	Update(ctx context.Context, resourceGroupName string, monitorName string, resource dynatrace.MonitorResourceUpdate) (result dynatrace.MonitorResource, err error)
}

var _ MonitorsClientAPI = (*dynatrace.MonitorsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result dynatrace.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result dynatrace.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*dynatrace.OperationsClient)(nil)

// TagRulesClientAPI contains the set of methods on the TagRulesClient type.
type TagRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, resource dynatrace.TagRule) (result dynatrace.TagRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string) (result dynatrace.TagRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string) (result dynatrace.TagRule, err error)
	List(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.TagRuleListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.TagRuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, resource dynatrace.TagRuleUpdate) (result dynatrace.TagRule, err error)
}

var _ TagRulesClientAPI = (*dynatrace.TagRulesClient)(nil)

// SingleSignOnClientAPI contains the set of methods on the SingleSignOnClient type.
type SingleSignOnClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, resource dynatrace.SingleSignOnResource) (result dynatrace.SingleSignOnCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, monitorName string, configurationName string) (result dynatrace.SingleSignOnResource, err error)
	List(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.SingleSignOnResourceListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result dynatrace.SingleSignOnResourceListResultIterator, err error)
}

var _ SingleSignOnClientAPI = (*dynatrace.SingleSignOnClient)(nil)
