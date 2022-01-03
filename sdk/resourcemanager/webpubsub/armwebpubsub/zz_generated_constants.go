//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

const (
	module  = "armwebpubsub"
	version = "v0.1.0"
)

// ACLAction - Default action when no other rule matches
type ACLAction string

const (
	ACLActionAllow ACLAction = "Allow"
	ACLActionDeny  ACLAction = "Deny"
)

// PossibleACLActionValues returns the possible values for the ACLAction const type.
func PossibleACLActionValues() []ACLAction {
	return []ACLAction{
		ACLActionAllow,
		ACLActionDeny,
	}
}

// ToPtr returns a *ACLAction pointing to the current value.
func (c ACLAction) ToPtr() *ACLAction {
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

// KeyType - The keyType to regenerate. Must be either 'primary' or 'secondary'(case-insensitive).
type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSalt      KeyType = "Salt"
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimary,
		KeyTypeSalt,
		KeyTypeSecondary,
	}
}

// ToPtr returns a *KeyType pointing to the current value.
func (c KeyType) ToPtr() *KeyType {
	return &c
}

// ManagedIdentityType - Represent the identity type: systemAssigned, userAssigned, None
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone           ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeUserAssigned   ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *ManagedIdentityType pointing to the current value.
func (c ManagedIdentityType) ToPtr() *ManagedIdentityType {
	return &c
}

// PrivateLinkServiceConnectionStatus - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusDisconnected,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateLinkServiceConnectionStatus pointing to the current value.
func (c PrivateLinkServiceConnectionStatus) ToPtr() *PrivateLinkServiceConnectionStatus {
	return &c
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// ScaleType - The scale type applicable to the sku.
type ScaleType string

const (
	ScaleTypeAutomatic ScaleType = "Automatic"
	ScaleTypeManual    ScaleType = "Manual"
	ScaleTypeNone      ScaleType = "None"
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

// SharedPrivateLinkResourceStatus - Status of the shared private link resource
type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = "Timeout"
)

// PossibleSharedPrivateLinkResourceStatusValues returns the possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{
		SharedPrivateLinkResourceStatusApproved,
		SharedPrivateLinkResourceStatusDisconnected,
		SharedPrivateLinkResourceStatusPending,
		SharedPrivateLinkResourceStatusRejected,
		SharedPrivateLinkResourceStatusTimeout,
	}
}

// ToPtr returns a *SharedPrivateLinkResourceStatus pointing to the current value.
func (c SharedPrivateLinkResourceStatus) ToPtr() *SharedPrivateLinkResourceStatus {
	return &c
}

// UpstreamAuthType - Gets or sets the type of auth. None or ManagedIdentity is supported now.
type UpstreamAuthType string

const (
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = "ManagedIdentity"
	UpstreamAuthTypeNone            UpstreamAuthType = "None"
)

// PossibleUpstreamAuthTypeValues returns the possible values for the UpstreamAuthType const type.
func PossibleUpstreamAuthTypeValues() []UpstreamAuthType {
	return []UpstreamAuthType{
		UpstreamAuthTypeManagedIdentity,
		UpstreamAuthTypeNone,
	}
}

// ToPtr returns a *UpstreamAuthType pointing to the current value.
func (c UpstreamAuthType) ToPtr() *UpstreamAuthType {
	return &c
}

// WebPubSubRequestType - Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
type WebPubSubRequestType string

const (
	WebPubSubRequestTypeClientConnection WebPubSubRequestType = "ClientConnection"
	WebPubSubRequestTypeRESTAPI          WebPubSubRequestType = "RESTAPI"
	WebPubSubRequestTypeServerConnection WebPubSubRequestType = "ServerConnection"
	WebPubSubRequestTypeTrace            WebPubSubRequestType = "Trace"
)

// PossibleWebPubSubRequestTypeValues returns the possible values for the WebPubSubRequestType const type.
func PossibleWebPubSubRequestTypeValues() []WebPubSubRequestType {
	return []WebPubSubRequestType{
		WebPubSubRequestTypeClientConnection,
		WebPubSubRequestTypeRESTAPI,
		WebPubSubRequestTypeServerConnection,
		WebPubSubRequestTypeTrace,
	}
}

// ToPtr returns a *WebPubSubRequestType pointing to the current value.
func (c WebPubSubRequestType) ToPtr() *WebPubSubRequestType {
	return &c
}

// WebPubSubSKUTier - Optional tier of this particular SKU. 'Standard' or 'Free'.
// Basic is deprecated, use Standard instead.
type WebPubSubSKUTier string

const (
	WebPubSubSKUTierBasic    WebPubSubSKUTier = "Basic"
	WebPubSubSKUTierFree     WebPubSubSKUTier = "Free"
	WebPubSubSKUTierPremium  WebPubSubSKUTier = "Premium"
	WebPubSubSKUTierStandard WebPubSubSKUTier = "Standard"
)

// PossibleWebPubSubSKUTierValues returns the possible values for the WebPubSubSKUTier const type.
func PossibleWebPubSubSKUTierValues() []WebPubSubSKUTier {
	return []WebPubSubSKUTier{
		WebPubSubSKUTierBasic,
		WebPubSubSKUTierFree,
		WebPubSubSKUTierPremium,
		WebPubSubSKUTierStandard,
	}
}

// ToPtr returns a *WebPubSubSKUTier pointing to the current value.
func (c WebPubSubSKUTier) ToPtr() *WebPubSubSKUTier {
	return &c
}
