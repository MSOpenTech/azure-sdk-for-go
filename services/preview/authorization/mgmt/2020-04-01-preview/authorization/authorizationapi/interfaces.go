package authorizationapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2020-04-01-preview/authorization"
	"github.com/Azure/go-autorest/autorest"
)

// ClassicAdministratorsClientAPI contains the set of methods on the ClassicAdministratorsClient type.
type ClassicAdministratorsClientAPI interface {
	List(ctx context.Context) (result authorization.ClassicAdministratorListResultPage, err error)
	ListComplete(ctx context.Context) (result authorization.ClassicAdministratorListResultIterator, err error)
}

var _ ClassicAdministratorsClientAPI = (*authorization.ClassicAdministratorsClient)(nil)

// GlobalAdministratorClientAPI contains the set of methods on the GlobalAdministratorClient type.
type GlobalAdministratorClientAPI interface {
	ElevateAccess(ctx context.Context) (result autorest.Response, err error)
}

var _ GlobalAdministratorClientAPI = (*authorization.GlobalAdministratorClient)(nil)

// ProviderOperationsMetadataClientAPI contains the set of methods on the ProviderOperationsMetadataClient type.
type ProviderOperationsMetadataClientAPI interface {
	Get(ctx context.Context, resourceProviderNamespace string, expand string) (result authorization.ProviderOperationsMetadata, err error)
	List(ctx context.Context, expand string) (result authorization.ProviderOperationsMetadataListResultPage, err error)
	ListComplete(ctx context.Context, expand string) (result authorization.ProviderOperationsMetadataListResultIterator, err error)
}

var _ ProviderOperationsMetadataClientAPI = (*authorization.ProviderOperationsMetadataClient)(nil)

// PermissionsClientAPI contains the set of methods on the PermissionsClient type.
type PermissionsClientAPI interface {
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result authorization.PermissionGetResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result authorization.PermissionGetResultIterator, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string) (result authorization.PermissionGetResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, resourceGroupName string) (result authorization.PermissionGetResultIterator, err error)
}

var _ PermissionsClientAPI = (*authorization.PermissionsClient)(nil)

// RoleDefinitionsClientAPI contains the set of methods on the RoleDefinitionsClient type.
type RoleDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, roleDefinitionID string, roleDefinition authorization.RoleDefinition) (result authorization.RoleDefinition, err error)
	Delete(ctx context.Context, scope string, roleDefinitionID string) (result authorization.RoleDefinition, err error)
	Get(ctx context.Context, scope string, roleDefinitionID string) (result authorization.RoleDefinition, err error)
	GetByID(ctx context.Context, roleID string) (result authorization.RoleDefinition, err error)
	List(ctx context.Context, scope string, filter string) (result authorization.RoleDefinitionListResultPage, err error)
	ListComplete(ctx context.Context, scope string, filter string) (result authorization.RoleDefinitionListResultIterator, err error)
}

var _ RoleDefinitionsClientAPI = (*authorization.RoleDefinitionsClient)(nil)

// DenyAssignmentsClientAPI contains the set of methods on the DenyAssignmentsClient type.
type DenyAssignmentsClientAPI interface {
	Get(ctx context.Context, scope string, denyAssignmentID string) (result authorization.DenyAssignment, err error)
	GetByID(ctx context.Context, denyAssignmentID string) (result authorization.DenyAssignment, err error)
	List(ctx context.Context, filter string) (result authorization.DenyAssignmentListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result authorization.DenyAssignmentListResultIterator, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result authorization.DenyAssignmentListResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result authorization.DenyAssignmentListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result authorization.DenyAssignmentListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result authorization.DenyAssignmentListResultIterator, err error)
	ListForScope(ctx context.Context, scope string, filter string) (result authorization.DenyAssignmentListResultPage, err error)
	ListForScopeComplete(ctx context.Context, scope string, filter string) (result authorization.DenyAssignmentListResultIterator, err error)
}

var _ DenyAssignmentsClientAPI = (*authorization.DenyAssignmentsClient)(nil)

// RoleAssignmentMetricsClientAPI contains the set of methods on the RoleAssignmentMetricsClient type.
type RoleAssignmentMetricsClientAPI interface {
	GetMetricsForSubscription(ctx context.Context) (result authorization.RoleAssignmentMetricsResult, err error)
}

var _ RoleAssignmentMetricsClientAPI = (*authorization.RoleAssignmentMetricsClient)(nil)

// RoleAssignmentsClientAPI contains the set of methods on the RoleAssignmentsClient type.
type RoleAssignmentsClientAPI interface {
	Create(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (result authorization.RoleAssignment, err error)
	CreateByID(ctx context.Context, roleID string, parameters authorization.RoleAssignmentCreateParameters) (result authorization.RoleAssignment, err error)
	Delete(ctx context.Context, scope string, roleAssignmentName string, tenantID string) (result authorization.RoleAssignment, err error)
	DeleteByID(ctx context.Context, roleID string, tenantID string) (result authorization.RoleAssignment, err error)
	Get(ctx context.Context, scope string, roleAssignmentName string, tenantID string) (result authorization.RoleAssignment, err error)
	GetByID(ctx context.Context, roleID string, tenantID string) (result authorization.RoleAssignment, err error)
	List(ctx context.Context, filter string, tenantID string) (result authorization.RoleAssignmentListResultPage, err error)
	ListComplete(ctx context.Context, filter string, tenantID string) (result authorization.RoleAssignmentListResultIterator, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string, tenantID string) (result authorization.RoleAssignmentListResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string, tenantID string) (result authorization.RoleAssignmentListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string, tenantID string) (result authorization.RoleAssignmentListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, tenantID string) (result authorization.RoleAssignmentListResultIterator, err error)
	ListForScope(ctx context.Context, scope string, filter string, tenantID string) (result authorization.RoleAssignmentListResultPage, err error)
	ListForScopeComplete(ctx context.Context, scope string, filter string, tenantID string) (result authorization.RoleAssignmentListResultIterator, err error)
}

var _ RoleAssignmentsClientAPI = (*authorization.RoleAssignmentsClient)(nil)
