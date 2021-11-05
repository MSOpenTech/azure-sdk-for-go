//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

const (
	module  = "armmanagementgroups"
	version = "v0.2.1"
)

type Enum0 string

const (
	Enum0Ancestors Enum0 = "ancestors"
	Enum0Children  Enum0 = "children"
	Enum0Path      Enum0 = "path"
)

// PossibleEnum0Values returns the possible values for the Enum0 const type.
func PossibleEnum0Values() []Enum0 {
	return []Enum0{
		Enum0Ancestors,
		Enum0Children,
		Enum0Path,
	}
}

// ToPtr returns a *Enum0 pointing to the current value.
func (c Enum0) ToPtr() *Enum0 {
	return &c
}

type Enum2 string

const (
	Enum2AllowedChildren             Enum2 = "AllowedChildren"
	Enum2AllowedParents              Enum2 = "AllowedParents"
	Enum2ChildrenOnly                Enum2 = "ChildrenOnly"
	Enum2ParentAndFirstLevelChildren Enum2 = "ParentAndFirstLevelChildren"
	Enum2ParentOnly                  Enum2 = "ParentOnly"
)

// PossibleEnum2Values returns the possible values for the Enum2 const type.
func PossibleEnum2Values() []Enum2 {
	return []Enum2{
		Enum2AllowedChildren,
		Enum2AllowedParents,
		Enum2ChildrenOnly,
		Enum2ParentAndFirstLevelChildren,
		Enum2ParentOnly,
	}
}

// ToPtr returns a *Enum2 pointing to the current value.
func (c Enum2) ToPtr() *Enum2 {
	return &c
}

type Enum3 string

const (
	Enum3Audit             Enum3 = "Audit"
	Enum3FullHierarchy     Enum3 = "FullHierarchy"
	Enum3GroupsOnly        Enum3 = "GroupsOnly"
	Enum3SubscriptionsOnly Enum3 = "SubscriptionsOnly"
)

// PossibleEnum3Values returns the possible values for the Enum3 const type.
func PossibleEnum3Values() []Enum3 {
	return []Enum3{
		Enum3Audit,
		Enum3FullHierarchy,
		Enum3GroupsOnly,
		Enum3SubscriptionsOnly,
	}
}

// ToPtr returns a *Enum3 pointing to the current value.
func (c Enum3) ToPtr() *Enum3 {
	return &c
}

// ManagementGroupChildType - The type of child resource.
type ManagementGroupChildType string

const (
	ManagementGroupChildTypeMicrosoftManagementManagementGroups ManagementGroupChildType = "Microsoft.Management/managementGroups"
	ManagementGroupChildTypeSubscriptions                       ManagementGroupChildType = "/subscriptions"
)

// PossibleManagementGroupChildTypeValues returns the possible values for the ManagementGroupChildType const type.
func PossibleManagementGroupChildTypeValues() []ManagementGroupChildType {
	return []ManagementGroupChildType{
		ManagementGroupChildTypeMicrosoftManagementManagementGroups,
		ManagementGroupChildTypeSubscriptions,
	}
}

// ToPtr returns a *ManagementGroupChildType pointing to the current value.
func (c ManagementGroupChildType) ToPtr() *ManagementGroupChildType {
	return &c
}

// Permissions - The users specific permissions to this item.
type Permissions string

const (
	PermissionsDelete   Permissions = "delete"
	PermissionsEdit     Permissions = "edit"
	PermissionsNoaccess Permissions = "noaccess"
	PermissionsView     Permissions = "view"
)

// PossiblePermissionsValues returns the possible values for the Permissions const type.
func PossiblePermissionsValues() []Permissions {
	return []Permissions{
		PermissionsDelete,
		PermissionsEdit,
		PermissionsNoaccess,
		PermissionsView,
	}
}

// ToPtr returns a *Permissions pointing to the current value.
func (c Permissions) ToPtr() *Permissions {
	return &c
}

// Reason - Required if nameAvailable == false. Invalid indicates the name provided does not match the resource provider's naming requirements (incorrect
// length, unsupported characters, etc.) AlreadyExists
// indicates that the name is already in use and is therefore unavailable.
type Reason string

const (
	ReasonInvalid       Reason = "Invalid"
	ReasonAlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonInvalid,
		ReasonAlreadyExists,
	}
}

// ToPtr returns a *Reason pointing to the current value.
func (c Reason) ToPtr() *Reason {
	return &c
}

// Status - The status of the Tenant Backfill
type Status string

const (
	StatusNotStarted               Status = "NotStarted"
	StatusNotStartedButGroupsExist Status = "NotStartedButGroupsExist"
	StatusStarted                  Status = "Started"
	StatusFailed                   Status = "Failed"
	StatusCancelled                Status = "Cancelled"
	StatusCompleted                Status = "Completed"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusNotStarted,
		StatusNotStartedButGroupsExist,
		StatusStarted,
		StatusFailed,
		StatusCancelled,
		StatusCompleted,
	}
}

// ToPtr returns a *Status pointing to the current value.
func (c Status) ToPtr() *Status {
	return &c
}
