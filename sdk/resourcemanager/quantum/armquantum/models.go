//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

import "time"

// APIKey - Azure quantum workspace Api key details.
type APIKey struct {
	// The creation time of the api key.
	CreatedAt *time.Time

	// READ-ONLY; The Api key.
	Key *string
}

// APIKeys - List of api keys to be generated.
type APIKeys struct {
	// A list of api key names.
	Keys []*KeyType
}

// CheckNameAvailabilityParameters - Details of check name availability request body.
type CheckNameAvailabilityParameters struct {
	// Name for checking availability.
	Name *string

	// The resource type of Quantum Workspace.
	Type *string
}

// CheckNameAvailabilityResult - Result of check name availability.
type CheckNameAvailabilityResult struct {
	// Indicator of availability of the Quantum Workspace resource name.
	NameAvailable *bool

	// The reason of unavailability.
	Reason *string

	// READ-ONLY; The detailed info regarding the reason associated with the Namespace.
	Message *string
}

// ListKeysResult - Result of list Api keys and connection strings.
type ListKeysResult struct {
	// Indicator of enablement of the Quantum workspace Api keys.
	APIKeyEnabled *bool

	// The quantum workspace primary api key.
	PrimaryKey *APIKey

	// The quantum workspace secondary api key.
	SecondaryKey *APIKey

	// READ-ONLY; The connection string of the primary api key.
	PrimaryConnectionString *string

	// READ-ONLY; The connection string of the secondary api key.
	SecondaryConnectionString *string
}

// OfferingsListResult - The response of a list Providers operation.
type OfferingsListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of Providers.
	NextLink *string

	// Result of a list Providers operation.
	Value []*ProviderDescription
}

// Operation provided by provider
type Operation struct {
	// Properties of the operation
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Name of the operation
	Name *string
}

// OperationDisplay - Properties of the operation
type OperationDisplay struct {
	// Description of the operation
	Description *string

	// Operation name
	Operation *string

	// Provider name
	Provider *string

	// Resource name
	Resource *string
}

// OperationsList - Lists the operations available.
type OperationsList struct {
	// REQUIRED; Array of operations
	Value []*Operation

	// Url to follow for getting next page of operations.
	NextLink *string
}

// PricingDetail - Detailed pricing information for an sku.
type PricingDetail struct {
	// Unique id for this pricing information.
	ID *string

	// The unit cost of this sku.
	Value *string
}

// PricingDimension - Information about pricing dimension.
type PricingDimension struct {
	// Unique id of this pricing dimension.
	ID *string

	// The display name of this pricing dimension.
	Name *string
}

// Provider - Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
type Provider struct {
	// The provider's marketplace application display name.
	ApplicationName *string

	// A Uri identifying the specific instance of this provider.
	InstanceURI *string

	// Unique id of this provider.
	ProviderID *string

	// The sku associated with pricing information for this provider.
	ProviderSKU *string

	// Provisioning status field
	ProvisioningState *Status

	// Id to track resource usage for the provider.
	ResourceUsageID *string
}

// ProviderDescription - Information about an offering. A provider offering is an entity that offers Targets to run Azure
// Quantum Jobs.
type ProviderDescription struct {
	// Unique provider's id.
	ID *string

	// Provider properties.
	Properties *ProviderProperties

	// READ-ONLY; Provider's display name.
	Name *string
}

// ProviderProperties - Provider properties.
type ProviderProperties struct {
	// Azure Active Directory info.
	AAD *ProviderPropertiesAAD

	// Provider's Managed-Application info
	ManagedApplication *ProviderPropertiesManagedApplication

	// The list of pricing dimensions from the provider.
	PricingDimensions []*PricingDimension

	// The list of quota dimensions from the provider.
	QuotaDimensions []*QuotaDimension

	// The list of skus available from this provider.
	SKUs []*SKUDescription

	// The list of targets available from this provider.
	Targets []*TargetDescription

	// READ-ONLY; Company name.
	Company *string

	// READ-ONLY; Provider's default endpoint.
	DefaultEndpoint *string

	// READ-ONLY; A description about this provider.
	Description *string

	// READ-ONLY; Provider type.
	ProviderType *string
}

// ProviderPropertiesAAD - Azure Active Directory info.
type ProviderPropertiesAAD struct {
	// READ-ONLY; Provider's application id.
	ApplicationID *string

	// READ-ONLY; Provider's tenant id.
	TenantID *string
}

// ProviderPropertiesManagedApplication - Provider's Managed-Application info
type ProviderPropertiesManagedApplication struct {
	// READ-ONLY; Provider's offer id.
	OfferID *string

	// READ-ONLY; Provider's publisher id.
	PublisherID *string
}

// QuotaDimension - Information about a specific quota dimension.
type QuotaDimension struct {
	// A description about this quota dimension.
	Description *string

	// Unique id of this dimension.
	ID *string

	// The display name of this quota dimension.
	Name *string

	// The reset period of this quota dimension.
	Period *string

	// The max limit of this dimension.
	Quota *float32

	// The scope of this quota dimension.
	Scope *string

	// The standard unit of measurement used for this quota dimension.
	Unit *string

	// The standard unit of measurement used for this quota dimension in plural form.
	UnitPlural *string
}

// SKUDescription - Information about a specific sku.
type SKUDescription struct {
	// Flag to indicate whether the sku should be automatically added during workspace creation.
	AutoAdd *bool

	// Description about this sku.
	Description *string

	// Unique sku id.
	ID *string

	// Display name of this sku.
	Name *string

	// The list of pricing details for the sku.
	PricingDetails []*PricingDetail

	// The list of quota dimensions for this sku.
	QuotaDimensions []*QuotaDimension

	// Uri to subscribe to the restricted access sku.
	RestrictedAccessURI *string

	// The list of targets available for this sku.
	Targets []*string

	// Display name of this sku.
	Version *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TagsObject - Tags object for patch operations.
type TagsObject struct {
	// Resource tags.
	Tags map[string]*string
}

// TargetDescription - Information about a Target. A target is the component that can process a specific type of Job.
type TargetDescription struct {
	// List of content encodings accepted by this target.
	AcceptedContentEncodings []*string

	// List of data formats accepted by this target.
	AcceptedDataFormats []*string

	// A description about this target.
	Description *string

	// Unique target id.
	ID *string

	// Display name of this target.
	Name *string
}

// Workspace - The resource proxy definition object for quantum workspace.
type Workspace struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Managed Identity information.
	Identity *WorkspaceIdentity

	// Gets or sets the properties. Define quantum workspace's specific properties.
	Properties *WorkspaceResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WorkspaceIdentity - Managed Identity information.
type WorkspaceIdentity struct {
	// The identity type.
	Type *ResourceIdentityType

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// WorkspaceListResult - The response of a list Workspaces operation.
type WorkspaceListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of Workspaces.
	NextLink *string

	// Result of a list Workspaces operation.
	Value []*Workspace
}

// WorkspaceResourceProperties - Properties of a Workspace
type WorkspaceResourceProperties struct {
	// Indicator of enablement of the Quantum workspace Api keys.
	APIKeyEnabled *bool

	// List of Providers selected for this Workspace
	Providers []*Provider

	// ARM Resource Id of the storage account associated with this workspace.
	StorageAccount *string

	// READ-ONLY; The URI of the workspace endpoint.
	EndpointURI *string

	// READ-ONLY; Provisioning status field
	ProvisioningState *ProvisioningStatus

	// READ-ONLY; Whether the current workspace is ready to accept Jobs.
	Usable *UsableStatus
}
