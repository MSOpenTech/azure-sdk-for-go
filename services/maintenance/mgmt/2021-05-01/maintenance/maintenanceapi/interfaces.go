package maintenanceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/maintenance/mgmt/2021-05-01/maintenance"
)

// PublicMaintenanceConfigurationsClientAPI contains the set of methods on the PublicMaintenanceConfigurationsClient type.
type PublicMaintenanceConfigurationsClientAPI interface {
	Get(ctx context.Context, resourceName string) (result maintenance.Configuration, err error)
	List(ctx context.Context) (result maintenance.ListMaintenanceConfigurationsResult, err error)
}

var _ PublicMaintenanceConfigurationsClientAPI = (*maintenance.PublicMaintenanceConfigurationsClient)(nil)

// ApplyUpdatesClientAPI contains the set of methods on the ApplyUpdatesClient type.
type ApplyUpdatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result maintenance.ApplyUpdate, err error)
	CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result maintenance.ApplyUpdate, err error)
	Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result maintenance.ApplyUpdate, err error)
	GetParent(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result maintenance.ApplyUpdate, err error)
	List(ctx context.Context) (result maintenance.ListApplyUpdate, err error)
}

var _ ApplyUpdatesClientAPI = (*maintenance.ApplyUpdatesClient)(nil)

// ConfigurationAssignmentsClientAPI contains the set of methods on the ConfigurationAssignmentsClient type.
type ConfigurationAssignmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment maintenance.ConfigurationAssignment) (result maintenance.ConfigurationAssignment, err error)
	CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment maintenance.ConfigurationAssignment) (result maintenance.ConfigurationAssignment, err error)
	Delete(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string) (result maintenance.ConfigurationAssignment, err error)
	DeleteParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string) (result maintenance.ConfigurationAssignment, err error)
	List(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result maintenance.ListConfigurationAssignmentsResult, err error)
	ListParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result maintenance.ListConfigurationAssignmentsResult, err error)
}

var _ ConfigurationAssignmentsClientAPI = (*maintenance.ConfigurationAssignmentsClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, configuration maintenance.Configuration) (result maintenance.Configuration, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result maintenance.Configuration, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result maintenance.Configuration, err error)
	List(ctx context.Context) (result maintenance.ListMaintenanceConfigurationsResult, err error)
	UpdateMethod(ctx context.Context, resourceGroupName string, resourceName string, configuration maintenance.Configuration) (result maintenance.Configuration, err error)
}

var _ ConfigurationsClientAPI = (*maintenance.ConfigurationsClient)(nil)

// ConfigurationsForResourceGroupClientAPI contains the set of methods on the ConfigurationsForResourceGroupClient type.
type ConfigurationsForResourceGroupClientAPI interface {
	List(ctx context.Context, resourceGroupName string) (result maintenance.ListMaintenanceConfigurationsResult, err error)
}

var _ ConfigurationsForResourceGroupClientAPI = (*maintenance.ConfigurationsForResourceGroupClient)(nil)

// ApplyUpdateForResourceGroupClientAPI contains the set of methods on the ApplyUpdateForResourceGroupClient type.
type ApplyUpdateForResourceGroupClientAPI interface {
	List(ctx context.Context, resourceGroupName string) (result maintenance.ListApplyUpdate, err error)
}

var _ ApplyUpdateForResourceGroupClientAPI = (*maintenance.ApplyUpdateForResourceGroupClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result maintenance.OperationsListResult, err error)
}

var _ OperationsClientAPI = (*maintenance.OperationsClient)(nil)

// UpdatesClientAPI contains the set of methods on the UpdatesClient type.
type UpdatesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result maintenance.ListUpdatesResult, err error)
	ListParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result maintenance.ListUpdatesResult, err error)
}

var _ UpdatesClientAPI = (*maintenance.UpdatesClient)(nil)
