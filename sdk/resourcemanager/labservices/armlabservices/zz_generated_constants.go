//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

const (
	moduleName    = "armlabservices"
	moduleVersion = "v0.2.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// ConnectionType - A connection type for access labs and VMs (Public, Private or None).
type ConnectionType string

const (
	ConnectionTypePublic  ConnectionType = "Public"
	ConnectionTypePrivate ConnectionType = "Private"
	ConnectionTypeNone    ConnectionType = "None"
)

// PossibleConnectionTypeValues returns the possible values for the ConnectionType const type.
func PossibleConnectionTypeValues() []ConnectionType {
	return []ConnectionType{
		ConnectionTypePublic,
		ConnectionTypePrivate,
		ConnectionTypeNone,
	}
}

// ToPtr returns a *ConnectionType pointing to the current value.
func (c ConnectionType) ToPtr() *ConnectionType {
	return &c
}

// CreateOption - Indicates what lab virtual machines are created from.
type CreateOption string

const (
	// CreateOptionImage - An image is used to create all lab user virtual machines. When this option is set, no template VM will
	// be created.
	CreateOptionImage CreateOption = "Image"
	// CreateOptionTemplateVM - A template VM will be used to create all lab user virtual machines.
	CreateOptionTemplateVM CreateOption = "TemplateVM"
)

// PossibleCreateOptionValues returns the possible values for the CreateOption const type.
func PossibleCreateOptionValues() []CreateOption {
	return []CreateOption{
		CreateOptionImage,
		CreateOptionTemplateVM,
	}
}

// ToPtr returns a *CreateOption pointing to the current value.
func (c CreateOption) ToPtr() *CreateOption {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// EnableState - Property enabled state.
type EnableState string

const (
	EnableStateEnabled  EnableState = "Enabled"
	EnableStateDisabled EnableState = "Disabled"
)

// PossibleEnableStateValues returns the possible values for the EnableState const type.
func PossibleEnableStateValues() []EnableState {
	return []EnableState{
		EnableStateEnabled,
		EnableStateDisabled,
	}
}

// ToPtr returns a *EnableState pointing to the current value.
func (c EnableState) ToPtr() *EnableState {
	return &c
}

// InvitationState - The lab user invitation state.
type InvitationState string

const (
	// InvitationStateNotSent - The invitation has not been sent.
	InvitationStateNotSent InvitationState = "NotSent"
	// InvitationStateSending - Currently sending the invitation.
	InvitationStateSending InvitationState = "Sending"
	// InvitationStateSent - The invitation has been successfully sent.
	InvitationStateSent InvitationState = "Sent"
	// InvitationStateFailed - There was an error while sending the invitation.
	InvitationStateFailed InvitationState = "Failed"
)

// PossibleInvitationStateValues returns the possible values for the InvitationState const type.
func PossibleInvitationStateValues() []InvitationState {
	return []InvitationState{
		InvitationStateNotSent,
		InvitationStateSending,
		InvitationStateSent,
		InvitationStateFailed,
	}
}

// ToPtr returns a *InvitationState pointing to the current value.
func (c InvitationState) ToPtr() *InvitationState {
	return &c
}

// LabServicesSKUTier - The tier of the SKU.
type LabServicesSKUTier string

const (
	LabServicesSKUTierPremium  LabServicesSKUTier = "Premium"
	LabServicesSKUTierStandard LabServicesSKUTier = "Standard"
)

// PossibleLabServicesSKUTierValues returns the possible values for the LabServicesSKUTier const type.
func PossibleLabServicesSKUTierValues() []LabServicesSKUTier {
	return []LabServicesSKUTier{
		LabServicesSKUTierPremium,
		LabServicesSKUTierStandard,
	}
}

// ToPtr returns a *LabServicesSKUTier pointing to the current value.
func (c LabServicesSKUTier) ToPtr() *LabServicesSKUTier {
	return &c
}

// LabState - The state of a virtual machine.
type LabState string

const (
	// LabStateDraft - The lab is currently in draft (has not been published).
	LabStateDraft LabState = "Draft"
	// LabStatePublishing - The lab is publishing.
	LabStatePublishing LabState = "Publishing"
	// LabStateScaling - The lab is scaling.
	LabStateScaling LabState = "Scaling"
	// LabStateSyncing - The lab is syncing users.
	LabStateSyncing LabState = "Syncing"
	// LabStatePublished - The lab has been published.
	LabStatePublished LabState = "Published"
)

// PossibleLabStateValues returns the possible values for the LabState const type.
func PossibleLabStateValues() []LabState {
	return []LabState{
		LabStateDraft,
		LabStatePublishing,
		LabStateScaling,
		LabStateSyncing,
		LabStatePublished,
	}
}

// ToPtr returns a *LabState pointing to the current value.
func (c LabState) ToPtr() *LabState {
	return &c
}

// OperationStatus - The operation status
type OperationStatus string

const (
	// OperationStatusNotStarted - The operation has been accepted but hasn't started.
	OperationStatusNotStarted OperationStatus = "NotStarted"
	// OperationStatusInProgress - The operation is running
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusSucceeded - The operation Succeeded
	OperationStatusSucceeded OperationStatus = "Succeeded"
	// OperationStatusFailed - The operation failed
	OperationStatusFailed OperationStatus = "Failed"
	// OperationStatusCanceled - Not supported yet
	OperationStatusCanceled OperationStatus = "Canceled"
)

// PossibleOperationStatusValues returns the possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{
		OperationStatusNotStarted,
		OperationStatusInProgress,
		OperationStatusSucceeded,
		OperationStatusFailed,
		OperationStatusCanceled,
	}
}

// ToPtr returns a *OperationStatus pointing to the current value.
func (c OperationStatus) ToPtr() *OperationStatus {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// OsState - The operating system state.
type OsState string

const (
	// OsStateGeneralized - Image does not contain any machine and user specific information.
	OsStateGeneralized OsState = "Generalized"
	// OsStateSpecialized - Image contains machine and user specific information.
	OsStateSpecialized OsState = "Specialized"
)

// PossibleOsStateValues returns the possible values for the OsState const type.
func PossibleOsStateValues() []OsState {
	return []OsState{
		OsStateGeneralized,
		OsStateSpecialized,
	}
}

// ToPtr returns a *OsState pointing to the current value.
func (c OsState) ToPtr() *OsState {
	return &c
}

// OsType - The operating system type.
type OsType string

const (
	OsTypeWindows OsType = "Windows"
	OsTypeLinux   OsType = "Linux"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeWindows,
		OsTypeLinux,
	}
}

// ToPtr returns a *OsType pointing to the current value.
func (c OsType) ToPtr() *OsType {
	return &c
}

// ProvisioningState - Resource provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCreating - Resource is in the process of being created.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateUpdating - New property values are being applied to the resource.
	ProvisioningStateUpdating ProvisioningState = "Updating"
	// ProvisioningStateDeleting - Resource is in the process of being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateSucceeded - Resource is in healthy state after creation or update operation.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateFailed - Previous operation on the resource has failed leaving resource in unhealthy state.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateLocked - The resource is locked and changes are currently blocked. This could be due to maintenance or
	// a scheduled operation. The state will go back to succeeded once the locking operation has finished.
	ProvisioningStateLocked ProvisioningState = "Locked"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateUpdating,
		ProvisioningStateDeleting,
		ProvisioningStateSucceeded,
		ProvisioningStateFailed,
		ProvisioningStateLocked,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// RecurrenceFrequency - Schedule recurrence frequencies.
type RecurrenceFrequency string

const (
	// RecurrenceFrequencyDaily - Schedule will run every days.
	RecurrenceFrequencyDaily RecurrenceFrequency = "Daily"
	// RecurrenceFrequencyWeekly - Schedule will run every week on days specified in weekDays.
	RecurrenceFrequencyWeekly RecurrenceFrequency = "Weekly"
)

// PossibleRecurrenceFrequencyValues returns the possible values for the RecurrenceFrequency const type.
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return []RecurrenceFrequency{
		RecurrenceFrequencyDaily,
		RecurrenceFrequencyWeekly,
	}
}

// ToPtr returns a *RecurrenceFrequency pointing to the current value.
func (c RecurrenceFrequency) ToPtr() *RecurrenceFrequency {
	return &c
}

// RegistrationState - The user lab registration state.
type RegistrationState string

const (
	// RegistrationStateRegistered - User has not yet registered with the lab.
	RegistrationStateRegistered RegistrationState = "Registered"
	// RegistrationStateNotRegistered - User has registered with the lab.
	RegistrationStateNotRegistered RegistrationState = "NotRegistered"
)

// PossibleRegistrationStateValues returns the possible values for the RegistrationState const type.
func PossibleRegistrationStateValues() []RegistrationState {
	return []RegistrationState{
		RegistrationStateRegistered,
		RegistrationStateNotRegistered,
	}
}

// ToPtr returns a *RegistrationState pointing to the current value.
func (c RegistrationState) ToPtr() *RegistrationState {
	return &c
}

// RestrictionReasonCode - The reason for the restriction.
type RestrictionReasonCode string

const (
	RestrictionReasonCodeNotAvailableForSubscription RestrictionReasonCode = "NotAvailableForSubscription"
	RestrictionReasonCodeQuotaID                     RestrictionReasonCode = "QuotaId"
)

// PossibleRestrictionReasonCodeValues returns the possible values for the RestrictionReasonCode const type.
func PossibleRestrictionReasonCodeValues() []RestrictionReasonCode {
	return []RestrictionReasonCode{
		RestrictionReasonCodeNotAvailableForSubscription,
		RestrictionReasonCodeQuotaID,
	}
}

// ToPtr returns a *RestrictionReasonCode pointing to the current value.
func (c RestrictionReasonCode) ToPtr() *RestrictionReasonCode {
	return &c
}

// RestrictionType - The type of restriction.
type RestrictionType string

const (
	RestrictionTypeLocation RestrictionType = "Location"
)

// PossibleRestrictionTypeValues returns the possible values for the RestrictionType const type.
func PossibleRestrictionTypeValues() []RestrictionType {
	return []RestrictionType{
		RestrictionTypeLocation,
	}
}

// ToPtr returns a *RestrictionType pointing to the current value.
func (c RestrictionType) ToPtr() *RestrictionType {
	return &c
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// ScaleType - The localized name of the resource.
type ScaleType string

const (
	// ScaleTypeAutomatic - The user is permitted to scale this SKU in and out.
	ScaleTypeAutomatic ScaleType = "Automatic"
	// ScaleTypeManual - The user must manually scale this SKU in and out.
	ScaleTypeManual ScaleType = "Manual"
	// ScaleTypeNone - The capacity is not adjustable in any way.
	ScaleTypeNone ScaleType = "None"
)

// PossibleScaleTypeValues returns the possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{
		ScaleTypeAutomatic,
		ScaleTypeManual,
		ScaleTypeNone,
	}
}

// ToPtr returns a *ScaleType pointing to the current value.
func (c ScaleType) ToPtr() *ScaleType {
	return &c
}

// ShutdownOnIdleMode - Defines whether to shut down VM on idle and the criteria for idle detection.
type ShutdownOnIdleMode string

const (
	// ShutdownOnIdleModeNone - The VM won't be shut down when it is idle.
	ShutdownOnIdleModeNone ShutdownOnIdleMode = "None"
	// ShutdownOnIdleModeUserAbsence - The VM will be considered as idle when there is no keyboard or mouse input.
	ShutdownOnIdleModeUserAbsence ShutdownOnIdleMode = "UserAbsence"
	// ShutdownOnIdleModeLowUsage - The VM will be considered as idle when user is absent and the resource (CPU and disk) consumption
	// is low.
	ShutdownOnIdleModeLowUsage ShutdownOnIdleMode = "LowUsage"
)

// PossibleShutdownOnIdleModeValues returns the possible values for the ShutdownOnIdleMode const type.
func PossibleShutdownOnIdleModeValues() []ShutdownOnIdleMode {
	return []ShutdownOnIdleMode{
		ShutdownOnIdleModeNone,
		ShutdownOnIdleModeUserAbsence,
		ShutdownOnIdleModeLowUsage,
	}
}

// ToPtr returns a *ShutdownOnIdleMode pointing to the current value.
func (c ShutdownOnIdleMode) ToPtr() *ShutdownOnIdleMode {
	return &c
}

// UsageUnit - The unit details.
type UsageUnit string

const (
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
	}
}

// ToPtr returns a *UsageUnit pointing to the current value.
func (c UsageUnit) ToPtr() *UsageUnit {
	return &c
}

// VirtualMachineState - The state of a virtual machine.
type VirtualMachineState string

const (
	// VirtualMachineStateStopped - The VM is currently stopped.
	VirtualMachineStateStopped VirtualMachineState = "Stopped"
	// VirtualMachineStateStarting - The VM is starting.
	VirtualMachineStateStarting VirtualMachineState = "Starting"
	// VirtualMachineStateRunning - The VM is running.
	VirtualMachineStateRunning VirtualMachineState = "Running"
	// VirtualMachineStateStopping - The VM is stopping.
	VirtualMachineStateStopping VirtualMachineState = "Stopping"
	// VirtualMachineStateResettingPassword - The VM password is being reset.
	VirtualMachineStateResettingPassword VirtualMachineState = "ResettingPassword"
	// VirtualMachineStateReimaging - The VM is being reimaged.
	VirtualMachineStateReimaging VirtualMachineState = "Reimaging"
	// VirtualMachineStateRedeploying - The VM is being redeployed.
	VirtualMachineStateRedeploying VirtualMachineState = "Redeploying"
)

// PossibleVirtualMachineStateValues returns the possible values for the VirtualMachineState const type.
func PossibleVirtualMachineStateValues() []VirtualMachineState {
	return []VirtualMachineState{
		VirtualMachineStateStopped,
		VirtualMachineStateStarting,
		VirtualMachineStateRunning,
		VirtualMachineStateStopping,
		VirtualMachineStateResettingPassword,
		VirtualMachineStateReimaging,
		VirtualMachineStateRedeploying,
	}
}

// ToPtr returns a *VirtualMachineState pointing to the current value.
func (c VirtualMachineState) ToPtr() *VirtualMachineState {
	return &c
}

// VirtualMachineType - The type of the lab virtual machine.
type VirtualMachineType string

const (
	// VirtualMachineTypeUser - A user VM
	VirtualMachineTypeUser VirtualMachineType = "User"
	// VirtualMachineTypeTemplate - A template VM
	VirtualMachineTypeTemplate VirtualMachineType = "Template"
)

// PossibleVirtualMachineTypeValues returns the possible values for the VirtualMachineType const type.
func PossibleVirtualMachineTypeValues() []VirtualMachineType {
	return []VirtualMachineType{
		VirtualMachineTypeUser,
		VirtualMachineTypeTemplate,
	}
}

// ToPtr returns a *VirtualMachineType pointing to the current value.
func (c VirtualMachineType) ToPtr() *VirtualMachineType {
	return &c
}

// WeekDay - Days of the week.
type WeekDay string

const (
	// WeekDaySunday - Schedule will run on Sunday
	WeekDaySunday WeekDay = "Sunday"
	// WeekDayMonday - Schedule will run on Monday
	WeekDayMonday WeekDay = "Monday"
	// WeekDayTuesday - Schedule will run on Tuesday
	WeekDayTuesday WeekDay = "Tuesday"
	// WeekDayWednesday - Schedule will run on Wednesday
	WeekDayWednesday WeekDay = "Wednesday"
	// WeekDayThursday - Schedule will run on Thursday
	WeekDayThursday WeekDay = "Thursday"
	// WeekDayFriday - Schedule will run on Friday
	WeekDayFriday WeekDay = "Friday"
	// WeekDaySaturday - Schedule will run on Saturday
	WeekDaySaturday WeekDay = "Saturday"
)

// PossibleWeekDayValues returns the possible values for the WeekDay const type.
func PossibleWeekDayValues() []WeekDay {
	return []WeekDay{
		WeekDaySunday,
		WeekDayMonday,
		WeekDayTuesday,
		WeekDayWednesday,
		WeekDayThursday,
		WeekDayFriday,
		WeekDaySaturday,
	}
}

// ToPtr returns a *WeekDay pointing to the current value.
func (c WeekDay) ToPtr() *WeekDay {
	return &c
}
