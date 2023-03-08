//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package rbac

// CreateOrUpdateRoleDefinitionOptions contains the optional parameters for the Client.CreateOrUpdateRoleDefinition
// method.
type CreateOrUpdateRoleDefinitionOptions struct {
	// placeholder for future optional parameters
}

// CreateRoleAssignmentOptions contains the optional parameters for the Client.CreateRoleAssignment method.
type CreateRoleAssignmentOptions struct {
	// placeholder for future optional parameters
}

// DeleteRoleAssignmentOptions contains the optional parameters for the Client.DeleteRoleAssignment method.
type DeleteRoleAssignmentOptions struct {
	// placeholder for future optional parameters
}

// DeleteRoleDefinitionOptions contains the optional parameters for the Client.DeleteRoleDefinition method.
type DeleteRoleDefinitionOptions struct {
	// placeholder for future optional parameters
}

// GetRoleAssignmentOptions contains the optional parameters for the Client.GetRoleAssignment method.
type GetRoleAssignmentOptions struct {
	// placeholder for future optional parameters
}

// GetRoleDefinitionOptions contains the optional parameters for the Client.GetRoleDefinition method.
type GetRoleDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ListRoleAssignmentsOptions contains the optional parameters for the Client.ListRoleAssignments method.
type ListRoleAssignmentsOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId
	// eq {id} to return all role assignments at, above or below the
	// scope for the specified principal.
	Filter *string
}

// ListRoleDefinitionsOptions contains the optional parameters for the Client.ListRoleDefinitions method.
type ListRoleDefinitionsOptions struct {
	// The filter to apply on the operation. Use atScopeAndBelow filter to search below the given scope as well.
	Filter *string
}

// Permission - Role definition permissions.
type Permission struct {
	// Action permissions that are granted.
	Actions []*string `json:"actions,omitempty"`

	// Data action permissions that are granted.
	DataActions []*DataAction `json:"dataActions,omitempty"`

	// Action permissions that are excluded but not denied. They may be granted by other role definitions assigned to a principal.
	NotActions []*string `json:"notActions,omitempty"`

	// Data action permissions that are excluded but not denied. They may be granted by other role definitions assigned to a principal.
	NotDataActions []*DataAction `json:"notDataActions,omitempty"`
}

// RoleAssignment - Role Assignments
type RoleAssignment struct {
	// Role assignment properties.
	Properties *RoleAssignmentPropertiesWithScope `json:"properties,omitempty"`

	// READ-ONLY; The role assignment ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The role assignment name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The role assignment type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RoleAssignmentCreateParameters - Role assignment create parameters.
type RoleAssignmentCreateParameters struct {
	// REQUIRED; Role assignment properties.
	Properties *RoleAssignmentProperties `json:"properties,omitempty"`
}

// RoleAssignmentListResult - Role assignment list operation result.
type RoleAssignmentListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Role assignment list.
	Value []*RoleAssignment `json:"value,omitempty"`
}

// RoleAssignmentProperties - Role assignment properties.
type RoleAssignmentProperties struct {
	// REQUIRED; The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user,
	// service principal, or security group.
	PrincipalID *string `json:"principalId,omitempty"`

	// REQUIRED; The role definition ID used in the role assignment.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
}

// RoleAssignmentPropertiesWithScope - Role assignment properties with scope.
type RoleAssignmentPropertiesWithScope struct {
	// The principal ID.
	PrincipalID *string `json:"principalId,omitempty"`

	// The role definition ID.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`

	// The role scope.
	Scope *RoleScope `json:"scope,omitempty"`
}

// RoleDefinition - Role definition.
type RoleDefinition struct {
	// Role definition properties.
	Properties *RoleDefinitionProperties `json:"properties,omitempty"`

	// READ-ONLY; The role definition ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The role definition name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The role definition type.
	Type *RoleDefinitionType `json:"type,omitempty" azure:"ro"`
}

// RoleDefinitionCreateParameters - Role definition create parameters.
type RoleDefinitionCreateParameters struct {
	// REQUIRED; Role definition properties.
	Properties *RoleDefinitionProperties `json:"properties,omitempty"`
}

// RoleDefinitionListResult - Role definition list operation result.
type RoleDefinitionListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Role definition list.
	Value []*RoleDefinition `json:"value,omitempty"`
}

// RoleDefinitionProperties - Role definition properties.
type RoleDefinitionProperties struct {
	// Role definition assignable scopes.
	AssignableScopes []*RoleScope `json:"assignableScopes,omitempty"`

	// The role definition description.
	Description *string `json:"description,omitempty"`

	// Role definition permissions.
	Permissions []*Permission `json:"permissions,omitempty"`

	// The role name.
	RoleName *string `json:"roleName,omitempty"`

	// The role type.
	RoleType *RoleType `json:"type,omitempty"`
}
