//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

import "time"

// CloudError - An error message
type CloudError struct {
	// The error message body
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody - The body of an error message
type CloudErrorBody struct {
	// The error code
	Code *string `json:"code,omitempty"`

	// Extra error information
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A description of what caused the error
	Message *string `json:"message,omitempty"`

	// The target resource of the error message
	Target *string `json:"target,omitempty"`
}

// DNSForwardingRuleset - Describes a DNS forwarding ruleset.
type DNSForwardingRuleset struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the DNS forwarding ruleset.
	Properties *DNSForwardingRulesetProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; ETag of the DNS forwarding ruleset.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DNSForwardingRulesetListResult - The response to an enumeration operation on DNS forwarding rulesets.
type DNSForwardingRulesetListResult struct {
	// Enumeration of the DNS forwarding rulesets.
	Value []*DNSForwardingRuleset `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// DNSForwardingRulesetPatch - Describes a DNS forwarding ruleset PATCH operation.
type DNSForwardingRulesetPatch struct {
	// Tags for DNS Resolver.
	Tags map[string]*string `json:"tags,omitempty"`
}

// DNSForwardingRulesetProperties - Represents the properties of a DNS forwarding ruleset.
type DNSForwardingRulesetProperties struct {
	// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in
	// the ruleset to the target DNS servers.
	DNSResolverOutboundEndpoints []*SubResource `json:"dnsResolverOutboundEndpoints,omitempty"`

	// READ-ONLY; The current provisioning state of the DNS forwarding ruleset. This is a read-only property and any attempt to
	// set this value will be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The resourceGuid for the DNS forwarding ruleset.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// DNSForwardingRulesetsClientBeginCreateOrUpdateOptions contains the optional parameters for the DNSForwardingRulesetsClient.BeginCreateOrUpdate
// method.
type DNSForwardingRulesetsClientBeginCreateOrUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new resource to be created, but to prevent updating an existing resource. Other values will be ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DNSForwardingRulesetsClientBeginDeleteOptions contains the optional parameters for the DNSForwardingRulesetsClient.BeginDelete
// method.
type DNSForwardingRulesetsClientBeginDeleteOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DNSForwardingRulesetsClientBeginUpdateOptions contains the optional parameters for the DNSForwardingRulesetsClient.BeginUpdate
// method.
type DNSForwardingRulesetsClientBeginUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DNSForwardingRulesetsClientGetOptions contains the optional parameters for the DNSForwardingRulesetsClient.Get method.
type DNSForwardingRulesetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DNSForwardingRulesetsClientListByResourceGroupOptions contains the optional parameters for the DNSForwardingRulesetsClient.ListByResourceGroup
// method.
type DNSForwardingRulesetsClientListByResourceGroupOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// DNSForwardingRulesetsClientListByVirtualNetworkOptions contains the optional parameters for the DNSForwardingRulesetsClient.ListByVirtualNetwork
// method.
type DNSForwardingRulesetsClientListByVirtualNetworkOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// DNSForwardingRulesetsClientListOptions contains the optional parameters for the DNSForwardingRulesetsClient.List method.
type DNSForwardingRulesetsClientListOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// DNSResolver - Describes a DNS resolver.
type DNSResolver struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the DNS resolver.
	Properties *Properties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; ETag of the DNS resolver.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DNSResolversClientBeginCreateOrUpdateOptions contains the optional parameters for the DNSResolversClient.BeginCreateOrUpdate
// method.
type DNSResolversClientBeginCreateOrUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new resource to be created, but to prevent updating an existing resource. Other values will be ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DNSResolversClientBeginDeleteOptions contains the optional parameters for the DNSResolversClient.BeginDelete method.
type DNSResolversClientBeginDeleteOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DNSResolversClientBeginUpdateOptions contains the optional parameters for the DNSResolversClient.BeginUpdate method.
type DNSResolversClientBeginUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DNSResolversClientGetOptions contains the optional parameters for the DNSResolversClient.Get method.
type DNSResolversClientGetOptions struct {
	// placeholder for future optional parameters
}

// DNSResolversClientListByResourceGroupOptions contains the optional parameters for the DNSResolversClient.ListByResourceGroup
// method.
type DNSResolversClientListByResourceGroupOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// DNSResolversClientListByVirtualNetworkOptions contains the optional parameters for the DNSResolversClient.ListByVirtualNetwork
// method.
type DNSResolversClientListByVirtualNetworkOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// DNSResolversClientListOptions contains the optional parameters for the DNSResolversClient.List method.
type DNSResolversClientListOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// ForwardingRule - Describes a forwarding rule within a DNS forwarding ruleset.
type ForwardingRule struct {
	// Properties of the forwarding rule.
	Properties *ForwardingRuleProperties `json:"properties,omitempty"`

	// READ-ONLY; ETag of the forwarding rule.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ForwardingRuleListResult - The response to an enumeration operation on forwarding rules within a DNS forwarding ruleset.
type ForwardingRuleListResult struct {
	// Enumeration of the forwarding rules.
	Value []*ForwardingRule `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// ForwardingRulePatch - Describes a forwarding rule for PATCH operation.
type ForwardingRulePatch struct {
	// Updatable properties of the forwarding rule.
	Properties *ForwardingRulePatchProperties `json:"properties,omitempty"`
}

// ForwardingRulePatchProperties - Represents the updatable properties of a forwarding rule within a DNS forwarding ruleset.
type ForwardingRulePatchProperties struct {
	// The state of forwarding rule.
	ForwardingRuleState *ForwardingRuleState `json:"forwardingRuleState,omitempty"`

	// Metadata attached to the forwarding rule.
	Metadata map[string]*string `json:"metadata,omitempty"`

	// DNS servers to forward the DNS query to.
	TargetDNSServers []*TargetDNSServer `json:"targetDnsServers,omitempty"`
}

// ForwardingRuleProperties - Represents the properties of a forwarding rule within a DNS forwarding ruleset.
type ForwardingRuleProperties struct {
	// REQUIRED; The domain name for the forwarding rule.
	DomainName *string `json:"domainName,omitempty"`

	// REQUIRED; DNS servers to forward the DNS query to.
	TargetDNSServers []*TargetDNSServer `json:"targetDnsServers,omitempty"`

	// The state of forwarding rule.
	ForwardingRuleState *ForwardingRuleState `json:"forwardingRuleState,omitempty"`

	// Metadata attached to the forwarding rule.
	Metadata map[string]*string `json:"metadata,omitempty"`

	// READ-ONLY; The current provisioning state of the forwarding rule. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ForwardingRulesClientCreateOrUpdateOptions contains the optional parameters for the ForwardingRulesClient.CreateOrUpdate
// method.
type ForwardingRulesClientCreateOrUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new resource to be created, but to prevent updating an existing resource. Other values will be ignored.
	IfNoneMatch *string
}

// ForwardingRulesClientDeleteOptions contains the optional parameters for the ForwardingRulesClient.Delete method.
type ForwardingRulesClientDeleteOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
}

// ForwardingRulesClientGetOptions contains the optional parameters for the ForwardingRulesClient.Get method.
type ForwardingRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ForwardingRulesClientListOptions contains the optional parameters for the ForwardingRulesClient.List method.
type ForwardingRulesClientListOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// ForwardingRulesClientUpdateOptions contains the optional parameters for the ForwardingRulesClient.Update method.
type ForwardingRulesClientUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
}

// IPConfiguration - IP configuration.
type IPConfiguration struct {
	// Private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`

	// Private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIpAllocationMethod,omitempty"`

	// The reference to the subnet bound to the IP configuration.
	Subnet *SubResource `json:"subnet,omitempty"`
}

// InboundEndpoint - Describes an inbound endpoint for a DNS resolver.
type InboundEndpoint struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the inbound endpoint.
	Properties *InboundEndpointProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; ETag of the inbound endpoint.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// InboundEndpointListResult - The response to an enumeration operation on inbound endpoints for a DNS resolver.
type InboundEndpointListResult struct {
	// Enumeration of the inbound endpoints for a DNS resolver.
	Value []*InboundEndpoint `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// InboundEndpointPatch - Describes an inbound endpoint for a DNS resolver for PATCH operation.
type InboundEndpointPatch struct {
	// Tags for inbound endpoint.
	Tags map[string]*string `json:"tags,omitempty"`
}

// InboundEndpointProperties - Represents the properties of an inbound endpoint for a DNS resolver.
type InboundEndpointProperties struct {
	// IP configurations for the inbound endpoint.
	IPConfigurations []*IPConfiguration `json:"ipConfigurations,omitempty"`

	// READ-ONLY; The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set
	// this value will be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The resourceGuid property of the inbound endpoint resource.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// InboundEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the InboundEndpointsClient.BeginCreateOrUpdate
// method.
type InboundEndpointsClientBeginCreateOrUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new resource to be created, but to prevent updating an existing resource. Other values will be ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// InboundEndpointsClientBeginDeleteOptions contains the optional parameters for the InboundEndpointsClient.BeginDelete method.
type InboundEndpointsClientBeginDeleteOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// InboundEndpointsClientBeginUpdateOptions contains the optional parameters for the InboundEndpointsClient.BeginUpdate method.
type InboundEndpointsClientBeginUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// InboundEndpointsClientGetOptions contains the optional parameters for the InboundEndpointsClient.Get method.
type InboundEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// InboundEndpointsClientListOptions contains the optional parameters for the InboundEndpointsClient.List method.
type InboundEndpointsClientListOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// ListResult - The response to an enumeration operation on DNS resolvers.
type ListResult struct {
	// Enumeration of the DNS resolvers.
	Value []*DNSResolver `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// OutboundEndpoint - Describes an outbound endpoint for a DNS resolver.
type OutboundEndpoint struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of the outbound endpoint.
	Properties *OutboundEndpointProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; ETag of the outbound endpoint.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// OutboundEndpointListResult - The response to an enumeration operation on outbound endpoints for a DNS resolver.
type OutboundEndpointListResult struct {
	// Enumeration of the outbound endpoints for a DNS resolver.
	Value []*OutboundEndpoint `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// OutboundEndpointPatch - Describes an outbound endpoint for a DNS resolver for PATCH operation.
type OutboundEndpointPatch struct {
	// Tags for outbound endpoint.
	Tags map[string]*string `json:"tags,omitempty"`
}

// OutboundEndpointProperties - Represents the properties of an outbound endpoint for a DNS resolver.
type OutboundEndpointProperties struct {
	// The reference to the subnet used for the outbound endpoint.
	Subnet *SubResource `json:"subnet,omitempty"`

	// READ-ONLY; The current provisioning state of the outbound endpoint. This is a read-only property and any attempt to set
	// this value will be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The resourceGuid property of the outbound endpoint resource.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// OutboundEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the OutboundEndpointsClient.BeginCreateOrUpdate
// method.
type OutboundEndpointsClientBeginCreateOrUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new resource to be created, but to prevent updating an existing resource. Other values will be ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OutboundEndpointsClientBeginDeleteOptions contains the optional parameters for the OutboundEndpointsClient.BeginDelete
// method.
type OutboundEndpointsClientBeginDeleteOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OutboundEndpointsClientBeginUpdateOptions contains the optional parameters for the OutboundEndpointsClient.BeginUpdate
// method.
type OutboundEndpointsClientBeginUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OutboundEndpointsClientGetOptions contains the optional parameters for the OutboundEndpointsClient.Get method.
type OutboundEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// OutboundEndpointsClientListOptions contains the optional parameters for the OutboundEndpointsClient.List method.
type OutboundEndpointsClientListOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}

// Patch - Describes a DNS resolver for PATCH operation.
type Patch struct {
	// Tags for DNS Resolver.
	Tags map[string]*string `json:"tags,omitempty"`
}

// Properties - Represents the properties of a DNS resolver.
type Properties struct {
	// REQUIRED; The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *SubResource `json:"virtualNetwork,omitempty"`

	// READ-ONLY; The current status of the DNS resolver. This is a read-only property and any attempt to set this value will
	// be ignored.
	DNSResolverState *DNSResolverState `json:"dnsResolverState,omitempty" azure:"ro"`

	// READ-ONLY; The current provisioning state of the DNS resolver. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The resourceGuid property of the DNS resolver resource.
	ResourceGUID *string `json:"resourceGuid,omitempty" azure:"ro"`
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SubResource - Reference to another ARM resource.
type SubResource struct {
	// Resource ID.
	ID *string `json:"id,omitempty"`
}

// SubResourceListResult - The response to an enumeration operation on sub-resources.
type SubResourceListResult struct {
	// Enumeration of the sub-resources.
	Value []*SubResource `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TargetDNSServer - Describes a server to forward the DNS queries to.
type TargetDNSServer struct {
	// DNS server IP address.
	IPAddress *string `json:"ipAddress,omitempty"`

	// DNS server port.
	Port *int32 `json:"port,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// VirtualNetworkDNSForwardingRuleset - Reference to DNS forwarding ruleset and associated virtual network link.
type VirtualNetworkDNSForwardingRuleset struct {
	// DNS Forwarding Ruleset Resource ID.
	ID *string `json:"id,omitempty"`

	// Properties of the virtual network link sub-resource reference.
	Properties *VirtualNetworkLinkSubResourceProperties `json:"properties,omitempty"`
}

// VirtualNetworkDNSForwardingRulesetListResult - The response to an enumeration operation on Virtual Network DNS Forwarding
// Ruleset.
type VirtualNetworkDNSForwardingRulesetListResult struct {
	// Enumeration of the Virtual Network DNS Forwarding Ruleset.
	Value []*VirtualNetworkDNSForwardingRuleset `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// VirtualNetworkLink - Describes a virtual network link.
type VirtualNetworkLink struct {
	// Properties of the virtual network link.
	Properties *VirtualNetworkLinkProperties `json:"properties,omitempty"`

	// READ-ONLY; ETag of the virtual network link.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// VirtualNetworkLinkListResult - The response to an enumeration operation on virtual network links.
type VirtualNetworkLinkListResult struct {
	// Enumeration of the virtual network links.
	Value []*VirtualNetworkLink `json:"value,omitempty"`

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// VirtualNetworkLinkPatch - Describes a virtual network link for PATCH operation.
type VirtualNetworkLinkPatch struct {
	// Updatable properties of the virtual network link.
	Properties *VirtualNetworkLinkPatchProperties `json:"properties,omitempty"`
}

// VirtualNetworkLinkPatchProperties - Represents the updatable properties of the virtual network link.
type VirtualNetworkLinkPatchProperties struct {
	// Metadata attached to the virtual network link.
	Metadata map[string]*string `json:"metadata,omitempty"`
}

// VirtualNetworkLinkProperties - Represents the properties of a virtual network link.
type VirtualNetworkLinkProperties struct {
	// Metadata attached to the virtual network link.
	Metadata map[string]*string `json:"metadata,omitempty"`

	// The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *SubResource `json:"virtualNetwork,omitempty"`

	// READ-ONLY; The current provisioning state of the virtual network link. This is a read-only property and any attempt to
	// set this value will be ignored.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// VirtualNetworkLinkSubResourceProperties - The reference to the virtual network link that associates between the DNS forwarding
// ruleset and virtual network.
type VirtualNetworkLinkSubResourceProperties struct {
	// The reference to the virtual network link.
	VirtualNetworkLink *SubResource `json:"virtualNetworkLink,omitempty"`
}

// VirtualNetworkLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginCreateOrUpdate
// method.
type VirtualNetworkLinksClientBeginCreateOrUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Set to '*' to allow a new resource to be created, but to prevent updating an existing resource. Other values will be ignored.
	IfNoneMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualNetworkLinksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginDelete
// method.
type VirtualNetworkLinksClientBeginDeleteOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualNetworkLinksClientBeginUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginUpdate
// method.
type VirtualNetworkLinksClientBeginUpdateOptions struct {
	// ETag of the resource. Omit this value to always overwrite the current resource. Specify the last-seen ETag value to prevent
	// accidentally overwriting any concurrent changes.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualNetworkLinksClientGetOptions contains the optional parameters for the VirtualNetworkLinksClient.Get method.
type VirtualNetworkLinksClientGetOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkLinksClientListOptions contains the optional parameters for the VirtualNetworkLinksClient.List method.
type VirtualNetworkLinksClientListOptions struct {
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int32
}
