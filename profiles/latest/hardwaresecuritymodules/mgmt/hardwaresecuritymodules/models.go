//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package hardwaresecuritymodules

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/hardwaresecuritymodules/mgmt/2021-11-30/hardwaresecuritymodules"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type IdentityType = original.IdentityType

const (
	IdentityTypeApplication     IdentityType = original.IdentityTypeApplication
	IdentityTypeKey             IdentityType = original.IdentityTypeKey
	IdentityTypeManagedIdentity IdentityType = original.IdentityTypeManagedIdentity
	IdentityTypeUser            IdentityType = original.IdentityTypeUser
)

type JSONWebKeyType = original.JSONWebKeyType

const (
	JSONWebKeyTypeAllocating    JSONWebKeyType = original.JSONWebKeyTypeAllocating
	JSONWebKeyTypeCheckingQuota JSONWebKeyType = original.JSONWebKeyTypeCheckingQuota
	JSONWebKeyTypeConnecting    JSONWebKeyType = original.JSONWebKeyTypeConnecting
	JSONWebKeyTypeDeleting      JSONWebKeyType = original.JSONWebKeyTypeDeleting
	JSONWebKeyTypeFailed        JSONWebKeyType = original.JSONWebKeyTypeFailed
	JSONWebKeyTypeProvisioning  JSONWebKeyType = original.JSONWebKeyTypeProvisioning
	JSONWebKeyTypeSucceeded     JSONWebKeyType = original.JSONWebKeyTypeSucceeded
)

type SkuName = original.SkuName

const (
	SkuNamePayShield10KLMK1CPS250    SkuName = original.SkuNamePayShield10KLMK1CPS250
	SkuNamePayShield10KLMK1CPS2500   SkuName = original.SkuNamePayShield10KLMK1CPS2500
	SkuNamePayShield10KLMK1CPS60     SkuName = original.SkuNamePayShield10KLMK1CPS60
	SkuNamePayShield10KLMK2CPS250    SkuName = original.SkuNamePayShield10KLMK2CPS250
	SkuNamePayShield10KLMK2CPS2500   SkuName = original.SkuNamePayShield10KLMK2CPS2500
	SkuNamePayShield10KLMK2CPS60     SkuName = original.SkuNamePayShield10KLMK2CPS60
	SkuNameSafeNetLunaNetworkHSMA790 SkuName = original.SkuNameSafeNetLunaNetworkHSMA790
)

type APIEntityReference = original.APIEntityReference
type BaseClient = original.BaseClient
type DedicatedHsm = original.DedicatedHsm
type DedicatedHsmClient = original.DedicatedHsmClient
type DedicatedHsmCreateOrUpdateFuture = original.DedicatedHsmCreateOrUpdateFuture
type DedicatedHsmDeleteFuture = original.DedicatedHsmDeleteFuture
type DedicatedHsmError = original.DedicatedHsmError
type DedicatedHsmListResult = original.DedicatedHsmListResult
type DedicatedHsmListResultIterator = original.DedicatedHsmListResultIterator
type DedicatedHsmListResultPage = original.DedicatedHsmListResultPage
type DedicatedHsmOperation = original.DedicatedHsmOperation
type DedicatedHsmOperationDisplay = original.DedicatedHsmOperationDisplay
type DedicatedHsmOperationListResult = original.DedicatedHsmOperationListResult
type DedicatedHsmPatchParameters = original.DedicatedHsmPatchParameters
type DedicatedHsmProperties = original.DedicatedHsmProperties
type DedicatedHsmUpdateFuture = original.DedicatedHsmUpdateFuture
type EndpointDependency = original.EndpointDependency
type EndpointDetail = original.EndpointDetail
type Error = original.Error
type NetworkInterface = original.NetworkInterface
type NetworkProfile = original.NetworkProfile
type OperationsClient = original.OperationsClient
type OutboundEnvironmentEndpoint = original.OutboundEnvironmentEndpoint
type OutboundEnvironmentEndpointCollection = original.OutboundEnvironmentEndpointCollection
type OutboundEnvironmentEndpointCollectionIterator = original.OutboundEnvironmentEndpointCollectionIterator
type OutboundEnvironmentEndpointCollectionPage = original.OutboundEnvironmentEndpointCollectionPage
type Resource = original.Resource
type ResourceListResult = original.ResourceListResult
type Sku = original.Sku
type SystemData = original.SystemData

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDedicatedHsmClient(subscriptionID string) DedicatedHsmClient {
	return original.NewDedicatedHsmClient(subscriptionID)
}
func NewDedicatedHsmClientWithBaseURI(baseURI string, subscriptionID string) DedicatedHsmClient {
	return original.NewDedicatedHsmClientWithBaseURI(baseURI, subscriptionID)
}
func NewDedicatedHsmListResultIterator(page DedicatedHsmListResultPage) DedicatedHsmListResultIterator {
	return original.NewDedicatedHsmListResultIterator(page)
}
func NewDedicatedHsmListResultPage(cur DedicatedHsmListResult, getNextPage func(context.Context, DedicatedHsmListResult) (DedicatedHsmListResult, error)) DedicatedHsmListResultPage {
	return original.NewDedicatedHsmListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutboundEnvironmentEndpointCollectionIterator(page OutboundEnvironmentEndpointCollectionPage) OutboundEnvironmentEndpointCollectionIterator {
	return original.NewOutboundEnvironmentEndpointCollectionIterator(page)
}
func NewOutboundEnvironmentEndpointCollectionPage(cur OutboundEnvironmentEndpointCollection, getNextPage func(context.Context, OutboundEnvironmentEndpointCollection) (OutboundEnvironmentEndpointCollection, error)) OutboundEnvironmentEndpointCollectionPage {
	return original.NewOutboundEnvironmentEndpointCollectionPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return original.PossibleJSONWebKeyTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
