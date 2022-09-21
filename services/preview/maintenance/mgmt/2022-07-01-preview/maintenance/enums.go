package maintenance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// ImpactType enumerates the values for impact type.
type ImpactType string

const (
	// Freeze Pending updates can freeze network or disk io operation on resource.
	Freeze ImpactType = "Freeze"
	// None Pending updates has no impact on resource.
	None ImpactType = "None"
	// Redeploy Pending updates can redeploy resource.
	Redeploy ImpactType = "Redeploy"
	// Restart Pending updates can cause resource to restart.
	Restart ImpactType = "Restart"
)

// PossibleImpactTypeValues returns an array of possible values for the ImpactType const type.
func PossibleImpactTypeValues() []ImpactType {
	return []ImpactType{Freeze, None, Redeploy, Restart}
}

// RebootOptions enumerates the values for reboot options.
type RebootOptions string

const (
	// Always ...
	Always RebootOptions = "Always"
	// IfRequired ...
	IfRequired RebootOptions = "IfRequired"
	// Never ...
	Never RebootOptions = "Never"
)

// PossibleRebootOptionsValues returns an array of possible values for the RebootOptions const type.
func PossibleRebootOptionsValues() []RebootOptions {
	return []RebootOptions{Always, IfRequired, Never}
}

// Scope enumerates the values for scope.
type Scope string

const (
	// ScopeExtension This maintenance scope controls extension installation on VM/VMSS
	ScopeExtension Scope = "Extension"
	// ScopeHost This maintenance scope controls installation of azure platform updates i.e. services on
	// physical nodes hosting customer VMs.
	ScopeHost Scope = "Host"
	// ScopeInGuestPatch This maintenance scope controls installation of windows and linux packages on VM/VMSS
	ScopeInGuestPatch Scope = "InGuestPatch"
	// ScopeOSImage This maintenance scope controls os image installation on VM/VMSS
	ScopeOSImage Scope = "OSImage"
	// ScopeResource This maintenance scope controls the default update maintenance of the Azure Resource
	ScopeResource Scope = "Resource"
	// ScopeSQLDB This maintenance scope controls installation of SQL server platform updates.
	ScopeSQLDB Scope = "SQLDB"
	// ScopeSQLManagedInstance This maintenance scope controls installation of SQL managed instance platform
	// update.
	ScopeSQLManagedInstance Scope = "SQLManagedInstance"
)

// PossibleScopeValues returns an array of possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{ScopeExtension, ScopeHost, ScopeInGuestPatch, ScopeOSImage, ScopeResource, ScopeSQLDB, ScopeSQLManagedInstance}
}

// TaskScope enumerates the values for task scope.
type TaskScope string

const (
	// TaskScopeGlobal ...
	TaskScopeGlobal TaskScope = "Global"
	// TaskScopeResource ...
	TaskScopeResource TaskScope = "Resource"
)

// PossibleTaskScopeValues returns an array of possible values for the TaskScope const type.
func PossibleTaskScopeValues() []TaskScope {
	return []TaskScope{TaskScopeGlobal, TaskScopeResource}
}

// UpdateStatus enumerates the values for update status.
type UpdateStatus string

const (
	// Completed All updates are successfully applied.
	Completed UpdateStatus = "Completed"
	// InProgress Updates installation are in progress.
	InProgress UpdateStatus = "InProgress"
	// Pending There are pending updates to be installed.
	Pending UpdateStatus = "Pending"
	// RetryLater Updates installation failed and should be retried later.
	RetryLater UpdateStatus = "RetryLater"
	// RetryNow Updates installation failed but are ready to retry again.
	RetryNow UpdateStatus = "RetryNow"
)

// PossibleUpdateStatusValues returns an array of possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{Completed, InProgress, Pending, RetryLater, RetryNow}
}

// Visibility enumerates the values for visibility.
type Visibility string

const (
	// Custom Only visible to users with permissions.
	Custom Visibility = "Custom"
	// Public Visible to all users.
	Public Visibility = "Public"
)

// PossibleVisibilityValues returns an array of possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{Custom, Public}
}
