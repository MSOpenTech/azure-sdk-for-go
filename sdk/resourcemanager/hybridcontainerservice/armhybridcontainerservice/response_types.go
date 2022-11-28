//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcontainerservice

// AgentPoolClientCreateOrUpdateResponse contains the response from method AgentPoolClient.CreateOrUpdate.
type AgentPoolClientCreateOrUpdateResponse struct {
	AgentPool
}

// AgentPoolClientDeleteResponse contains the response from method AgentPoolClient.Delete.
type AgentPoolClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolClientGetResponse contains the response from method AgentPoolClient.Get.
type AgentPoolClientGetResponse struct {
	AgentPool
}

// AgentPoolClientListByProvisionedClusterResponse contains the response from method AgentPoolClient.ListByProvisionedCluster.
type AgentPoolClientListByProvisionedClusterResponse struct {
	AgentPoolListResult
}

// AgentPoolClientUpdateResponse contains the response from method AgentPoolClient.Update.
type AgentPoolClientUpdateResponse struct {
	AgentPool
}

// ClientListOrchestratorsResponse contains the response from method Client.ListOrchestrators.
type ClientListOrchestratorsResponse struct {
	OrchestratorVersionProfileListResult
}

// ClientListVMSKUsResponse contains the response from method Client.ListVMSKUs.
type ClientListVMSKUsResponse struct {
	VMSKUListResult
}

// HybridIdentityMetadataClientDeleteResponse contains the response from method HybridIdentityMetadataClient.Delete.
type HybridIdentityMetadataClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridIdentityMetadataClientGetResponse contains the response from method HybridIdentityMetadataClient.Get.
type HybridIdentityMetadataClientGetResponse struct {
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientListByClusterResponse contains the response from method HybridIdentityMetadataClient.ListByCluster.
type HybridIdentityMetadataClientListByClusterResponse struct {
	HybridIdentityMetadataList
}

// HybridIdentityMetadataClientPutResponse contains the response from method HybridIdentityMetadataClient.Put.
type HybridIdentityMetadataClientPutResponse struct {
	HybridIdentityMetadata
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	ResourceProviderOperationList
}

// ProvisionedClustersClientCreateOrUpdateResponse contains the response from method ProvisionedClustersClient.CreateOrUpdate.
type ProvisionedClustersClientCreateOrUpdateResponse struct {
	ProvisionedClustersResponse
}

// ProvisionedClustersClientDeleteResponse contains the response from method ProvisionedClustersClient.Delete.
type ProvisionedClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ProvisionedClustersClientGetResponse contains the response from method ProvisionedClustersClient.Get.
type ProvisionedClustersClientGetResponse struct {
	ProvisionedClustersResponse
}

// ProvisionedClustersClientListByResourceGroupResponse contains the response from method ProvisionedClustersClient.ListByResourceGroup.
type ProvisionedClustersClientListByResourceGroupResponse struct {
	ProvisionedClustersResponseListResult
}

// ProvisionedClustersClientListBySubscriptionResponse contains the response from method ProvisionedClustersClient.ListBySubscription.
type ProvisionedClustersClientListBySubscriptionResponse struct {
	ProvisionedClustersResponseListResult
}

// ProvisionedClustersClientUpdateResponse contains the response from method ProvisionedClustersClient.Update.
type ProvisionedClustersClientUpdateResponse struct {
	ProvisionedClustersResponse
}

// StorageSpacesClientCreateOrUpdateResponse contains the response from method StorageSpacesClient.CreateOrUpdate.
type StorageSpacesClientCreateOrUpdateResponse struct {
	StorageSpaces
}

// StorageSpacesClientDeleteResponse contains the response from method StorageSpacesClient.Delete.
type StorageSpacesClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageSpacesClientListByResourceGroupResponse contains the response from method StorageSpacesClient.ListByResourceGroup.
type StorageSpacesClientListByResourceGroupResponse struct {
	StorageSpacesListResult
}

// StorageSpacesClientListBySubscriptionResponse contains the response from method StorageSpacesClient.ListBySubscription.
type StorageSpacesClientListBySubscriptionResponse struct {
	StorageSpacesListResult
}

// StorageSpacesClientRetrieveResponse contains the response from method StorageSpacesClient.Retrieve.
type StorageSpacesClientRetrieveResponse struct {
	StorageSpaces
}

// StorageSpacesClientUpdateResponse contains the response from method StorageSpacesClient.Update.
type StorageSpacesClientUpdateResponse struct {
	StorageSpaces
}

// VirtualNetworksClientCreateOrUpdateResponse contains the response from method VirtualNetworksClient.CreateOrUpdate.
type VirtualNetworksClientCreateOrUpdateResponse struct {
	VirtualNetworks
}

// VirtualNetworksClientDeleteResponse contains the response from method VirtualNetworksClient.Delete.
type VirtualNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworksClientListByResourceGroupResponse contains the response from method VirtualNetworksClient.ListByResourceGroup.
type VirtualNetworksClientListByResourceGroupResponse struct {
	VirtualNetworksListResult
}

// VirtualNetworksClientListBySubscriptionResponse contains the response from method VirtualNetworksClient.ListBySubscription.
type VirtualNetworksClientListBySubscriptionResponse struct {
	VirtualNetworksListResult
}

// VirtualNetworksClientRetrieveResponse contains the response from method VirtualNetworksClient.Retrieve.
type VirtualNetworksClientRetrieveResponse struct {
	VirtualNetworks
}

// VirtualNetworksClientUpdateResponse contains the response from method VirtualNetworksClient.Update.
type VirtualNetworksClientUpdateResponse struct {
	VirtualNetworks
}
