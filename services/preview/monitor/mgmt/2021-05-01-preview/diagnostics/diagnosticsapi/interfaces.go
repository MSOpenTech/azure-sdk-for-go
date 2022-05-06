package diagnosticsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-05-01-preview/diagnostics"
	"github.com/Azure/go-autorest/autorest"
)

// AutoscaleSettingsClientAPI contains the set of methods on the AutoscaleSettingsClient type.
type AutoscaleSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, autoscaleSettingName string, parameters diagnostics.AutoscaleSettingResource) (result diagnostics.AutoscaleSettingResource, err error)
	Delete(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result diagnostics.AutoscaleSettingResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result diagnostics.AutoscaleSettingResourceCollectionPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result diagnostics.AutoscaleSettingResourceCollectionIterator, err error)
	ListBySubscription(ctx context.Context) (result diagnostics.AutoscaleSettingResourceCollectionPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result diagnostics.AutoscaleSettingResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, autoscaleSettingName string, autoscaleSettingResource diagnostics.AutoscaleSettingResourcePatch) (result diagnostics.AutoscaleSettingResource, err error)
}

var _ AutoscaleSettingsClientAPI = (*diagnostics.AutoscaleSettingsClient)(nil)

// PredictiveMetricClientAPI contains the set of methods on the PredictiveMetricClient type.
type PredictiveMetricClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, autoscaleSettingName string, timespan string, interval string, metricNamespace string, metricName string, aggregation string) (result diagnostics.PredictiveResponse, err error)
}

var _ PredictiveMetricClientAPI = (*diagnostics.PredictiveMetricClient)(nil)

// SettingsClientAPI contains the set of methods on the SettingsClient type.
type SettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceURI string, parameters diagnostics.SettingsResource, name string) (result diagnostics.SettingsResource, err error)
	Delete(ctx context.Context, resourceURI string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceURI string, name string) (result diagnostics.SettingsResource, err error)
	List(ctx context.Context, resourceURI string) (result diagnostics.SettingsResourceCollection, err error)
}

var _ SettingsClientAPI = (*diagnostics.SettingsClient)(nil)

// SettingsCategoryClientAPI contains the set of methods on the SettingsCategoryClient type.
type SettingsCategoryClientAPI interface {
	Get(ctx context.Context, resourceURI string, name string) (result diagnostics.SettingsCategoryResource, err error)
	List(ctx context.Context, resourceURI string) (result diagnostics.SettingsCategoryResourceCollection, err error)
}

var _ SettingsCategoryClientAPI = (*diagnostics.SettingsCategoryClient)(nil)

// ManagementGroupDiagnosticSettingsClientAPI contains the set of methods on the ManagementGroupDiagnosticSettingsClient type.
type ManagementGroupDiagnosticSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, managementGroupID string, parameters diagnostics.ManagementGroupDiagnosticSettingsResource, name string) (result diagnostics.ManagementGroupDiagnosticSettingsResource, err error)
	Delete(ctx context.Context, managementGroupID string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, managementGroupID string, name string) (result diagnostics.ManagementGroupDiagnosticSettingsResource, err error)
	List(ctx context.Context, managementGroupID string) (result diagnostics.ManagementGroupDiagnosticSettingsResourceCollection, err error)
}

var _ ManagementGroupDiagnosticSettingsClientAPI = (*diagnostics.ManagementGroupDiagnosticSettingsClient)(nil)

// SubscriptionDiagnosticSettingsClientAPI contains the set of methods on the SubscriptionDiagnosticSettingsClient type.
type SubscriptionDiagnosticSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters diagnostics.SubscriptionDiagnosticSettingsResource, name string) (result diagnostics.SubscriptionDiagnosticSettingsResource, err error)
	Delete(ctx context.Context, name string) (result autorest.Response, err error)
	Get(ctx context.Context, name string) (result diagnostics.SubscriptionDiagnosticSettingsResource, err error)
	List(ctx context.Context) (result diagnostics.SubscriptionDiagnosticSettingsResourceCollection, err error)
}

var _ SubscriptionDiagnosticSettingsClientAPI = (*diagnostics.SubscriptionDiagnosticSettingsClient)(nil)

// PrivateLinkScopesClientAPI contains the set of methods on the PrivateLinkScopesClient type.
type PrivateLinkScopesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, azureMonitorPrivateLinkScopePayload diagnostics.AzureMonitorPrivateLinkScope) (result diagnostics.AzureMonitorPrivateLinkScope, err error)
	Delete(ctx context.Context, resourceGroupName string, scopeName string) (result diagnostics.PrivateLinkScopesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, scopeName string) (result diagnostics.AzureMonitorPrivateLinkScope, err error)
	List(ctx context.Context) (result diagnostics.AzureMonitorPrivateLinkScopeListResultPage, err error)
	ListComplete(ctx context.Context) (result diagnostics.AzureMonitorPrivateLinkScopeListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result diagnostics.AzureMonitorPrivateLinkScopeListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result diagnostics.AzureMonitorPrivateLinkScopeListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags diagnostics.TagsResource) (result diagnostics.AzureMonitorPrivateLinkScope, err error)
}

var _ PrivateLinkScopesClientAPI = (*diagnostics.PrivateLinkScopesClient)(nil)

// PrivateLinkScopeOperationStatusClientAPI contains the set of methods on the PrivateLinkScopeOperationStatusClient type.
type PrivateLinkScopeOperationStatusClientAPI interface {
	Get(ctx context.Context, asyncOperationID string, resourceGroupName string) (result diagnostics.OperationStatus, err error)
}

var _ PrivateLinkScopeOperationStatusClientAPI = (*diagnostics.PrivateLinkScopeOperationStatusClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, scopeName string, groupName string) (result diagnostics.PrivateLinkResource, err error)
	ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result diagnostics.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*diagnostics.PrivateLinkResourcesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, privateEndpointConnectionName string, parameters diagnostics.PrivateEndpointConnection) (result diagnostics.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, scopeName string, privateEndpointConnectionName string) (result diagnostics.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, scopeName string, privateEndpointConnectionName string) (result diagnostics.PrivateEndpointConnection, err error)
	ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result diagnostics.PrivateEndpointConnectionListResult, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*diagnostics.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkScopedResourcesClientAPI contains the set of methods on the PrivateLinkScopedResourcesClient type.
type PrivateLinkScopedResourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters diagnostics.ScopedResource) (result diagnostics.PrivateLinkScopedResourcesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, scopeName string, name string) (result diagnostics.PrivateLinkScopedResourcesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, scopeName string, name string) (result diagnostics.ScopedResource, err error)
	ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result diagnostics.ScopedResourceListResultPage, err error)
	ListByPrivateLinkScopeComplete(ctx context.Context, resourceGroupName string, scopeName string) (result diagnostics.ScopedResourceListResultIterator, err error)
}

var _ PrivateLinkScopedResourcesClientAPI = (*diagnostics.PrivateLinkScopedResourcesClient)(nil)

// ActionGroupsClientAPI contains the set of methods on the ActionGroupsClient type.
type ActionGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup diagnostics.ActionGroupResource) (result diagnostics.ActionGroupResource, err error)
	Delete(ctx context.Context, resourceGroupName string, actionGroupName string) (result autorest.Response, err error)
	EnableReceiver(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest diagnostics.EnableRequest) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, actionGroupName string) (result diagnostics.ActionGroupResource, err error)
	GetTestNotifications(ctx context.Context, notificationID string) (result diagnostics.TestNotificationDetailsResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result diagnostics.ActionGroupList, err error)
	ListBySubscriptionID(ctx context.Context) (result diagnostics.ActionGroupList, err error)
	PostTestNotifications(ctx context.Context, notificationRequest diagnostics.NotificationRequestBody) (result diagnostics.ActionGroupsPostTestNotificationsFuture, err error)
	Update(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch diagnostics.ActionGroupPatchBody) (result diagnostics.ActionGroupResource, err error)
}

var _ ActionGroupsClientAPI = (*diagnostics.ActionGroupsClient)(nil)
