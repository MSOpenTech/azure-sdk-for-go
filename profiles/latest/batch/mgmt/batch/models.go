// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package batch

import original "github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2017-09-01/batch"

type LocationClient = original.LocationClient
type AccountKeyType = original.AccountKeyType

const (
	Primary   AccountKeyType = original.Primary
	Secondary AccountKeyType = original.Secondary
)

type AllocationState = original.AllocationState

const (
	Resizing AllocationState = original.Resizing
	Steady   AllocationState = original.Steady
	Stopping AllocationState = original.Stopping
)

type AutoUserScope = original.AutoUserScope

const (
	AutoUserScopePool AutoUserScope = original.AutoUserScopePool
	AutoUserScopeTask AutoUserScope = original.AutoUserScopeTask
)

type CachingType = original.CachingType

const (
	None      CachingType = original.None
	ReadOnly  CachingType = original.ReadOnly
	ReadWrite CachingType = original.ReadWrite
)

type CertificateFormat = original.CertificateFormat

const (
	Cer CertificateFormat = original.Cer
	Pfx CertificateFormat = original.Pfx
)

type CertificateProvisioningState = original.CertificateProvisioningState

const (
	Deleting  CertificateProvisioningState = original.Deleting
	Failed    CertificateProvisioningState = original.Failed
	Succeeded CertificateProvisioningState = original.Succeeded
)

type CertificateStoreLocation = original.CertificateStoreLocation

const (
	CurrentUser  CertificateStoreLocation = original.CurrentUser
	LocalMachine CertificateStoreLocation = original.LocalMachine
)

type CertificateVisibility = original.CertificateVisibility

const (
	CertificateVisibilityRemoteUser CertificateVisibility = original.CertificateVisibilityRemoteUser
	CertificateVisibilityStartTask  CertificateVisibility = original.CertificateVisibilityStartTask
	CertificateVisibilityTask       CertificateVisibility = original.CertificateVisibilityTask
)

type ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOption

const (
	Requeue        ComputeNodeDeallocationOption = original.Requeue
	RetainedData   ComputeNodeDeallocationOption = original.RetainedData
	TaskCompletion ComputeNodeDeallocationOption = original.TaskCompletion
	Terminate      ComputeNodeDeallocationOption = original.Terminate
)

type ComputeNodeFillType = original.ComputeNodeFillType

const (
	Pack   ComputeNodeFillType = original.Pack
	Spread ComputeNodeFillType = original.Spread
)

type ElevationLevel = original.ElevationLevel

const (
	Admin    ElevationLevel = original.Admin
	NonAdmin ElevationLevel = original.NonAdmin
)

type InboundEndpointProtocol = original.InboundEndpointProtocol

const (
	TCP InboundEndpointProtocol = original.TCP
	UDP InboundEndpointProtocol = original.UDP
)

type InterNodeCommunicationState = original.InterNodeCommunicationState

const (
	Disabled InterNodeCommunicationState = original.Disabled
	Enabled  InterNodeCommunicationState = original.Enabled
)

type NameAvailabilityReason = original.NameAvailabilityReason

const (
	AlreadyExists NameAvailabilityReason = original.AlreadyExists
	Invalid       NameAvailabilityReason = original.Invalid
)

type NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccess

const (
	Allow NetworkSecurityGroupRuleAccess = original.Allow
	Deny  NetworkSecurityGroupRuleAccess = original.Deny
)

type PackageState = original.PackageState

const (
	Active   PackageState = original.Active
	Pending  PackageState = original.Pending
	Unmapped PackageState = original.Unmapped
)

type PoolAllocationMode = original.PoolAllocationMode

const (
	BatchService     PoolAllocationMode = original.BatchService
	UserSubscription PoolAllocationMode = original.UserSubscription
)

type PoolProvisioningState = original.PoolProvisioningState

const (
	PoolProvisioningStateDeleting  PoolProvisioningState = original.PoolProvisioningStateDeleting
	PoolProvisioningStateSucceeded PoolProvisioningState = original.PoolProvisioningStateSucceeded
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCancelled ProvisioningState = original.ProvisioningStateCancelled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateInvalid   ProvisioningState = original.ProvisioningStateInvalid
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type StorageAccountType = original.StorageAccountType

const (
	PremiumLRS  StorageAccountType = original.PremiumLRS
	StandardLRS StorageAccountType = original.StandardLRS
)

type Account = original.Account
type AccountCreateFuture = original.AccountCreateFuture
type AccountCreateParameters = original.AccountCreateParameters
type AccountCreateProperties = original.AccountCreateProperties
type AccountDeleteFuture = original.AccountDeleteFuture
type AccountKeys = original.AccountKeys
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountProperties = original.AccountProperties
type AccountRegenerateKeyParameters = original.AccountRegenerateKeyParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountUpdateProperties = original.AccountUpdateProperties
type ActivateApplicationPackageParameters = original.ActivateApplicationPackageParameters
type Application = original.Application
type ApplicationCreateParameters = original.ApplicationCreateParameters
type ApplicationPackage = original.ApplicationPackage
type ApplicationPackageReference = original.ApplicationPackageReference
type ApplicationUpdateParameters = original.ApplicationUpdateParameters
type AutoScaleRun = original.AutoScaleRun
type AutoScaleRunError = original.AutoScaleRunError
type AutoScaleSettings = original.AutoScaleSettings
type AutoStorageBaseProperties = original.AutoStorageBaseProperties
type AutoStorageProperties = original.AutoStorageProperties
type AutoUserSpecification = original.AutoUserSpecification
type Certificate = original.Certificate
type CertificateBaseProperties = original.CertificateBaseProperties
type CertificateCreateFuture = original.CertificateCreateFuture
type CertificateCreateOrUpdateParameters = original.CertificateCreateOrUpdateParameters
type CertificateCreateOrUpdateProperties = original.CertificateCreateOrUpdateProperties
type CertificateDeleteFuture = original.CertificateDeleteFuture
type CertificateProperties = original.CertificateProperties
type CertificateReference = original.CertificateReference
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CloudServiceConfiguration = original.CloudServiceConfiguration
type DataDisk = original.DataDisk
type DeleteCertificateError = original.DeleteCertificateError
type DeploymentConfiguration = original.DeploymentConfiguration
type EnvironmentSetting = original.EnvironmentSetting
type FixedScaleSettings = original.FixedScaleSettings
type ImageReference = original.ImageReference
type InboundNatPool = original.InboundNatPool
type KeyVaultReference = original.KeyVaultReference
type LinuxUserConfiguration = original.LinuxUserConfiguration
type ListApplicationsResult = original.ListApplicationsResult
type ListApplicationsResultIterator = original.ListApplicationsResultIterator
type ListApplicationsResultPage = original.ListApplicationsResultPage
type ListCertificatesResult = original.ListCertificatesResult
type ListCertificatesResultIterator = original.ListCertificatesResultIterator
type ListCertificatesResultPage = original.ListCertificatesResultPage
type ListPoolsResult = original.ListPoolsResult
type ListPoolsResultIterator = original.ListPoolsResultIterator
type ListPoolsResultPage = original.ListPoolsResultPage
type LocationQuota = original.LocationQuota
type MetadataItem = original.MetadataItem
type NetworkConfiguration = original.NetworkConfiguration
type NetworkSecurityGroupRule = original.NetworkSecurityGroupRule
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OSDisk = original.OSDisk
type Pool = original.Pool
type PoolCreateFuture = original.PoolCreateFuture
type PoolDeleteFuture = original.PoolDeleteFuture
type PoolEndpointConfiguration = original.PoolEndpointConfiguration
type PoolProperties = original.PoolProperties
type ProxyResource = original.ProxyResource
type ResizeError = original.ResizeError
type ResizeOperationStatus = original.ResizeOperationStatus
type Resource = original.Resource
type ResourceFile = original.ResourceFile
type ScaleSettings = original.ScaleSettings
type StartTask = original.StartTask
type TaskSchedulingPolicy = original.TaskSchedulingPolicy
type UserAccount = original.UserAccount
type UserIdentity = original.UserIdentity
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type WindowsConfiguration = original.WindowsConfiguration
type PoolClient = original.PoolClient
type CertificateClient = original.CertificateClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type OperationsClient = original.OperationsClient
type ApplicationClient = original.ApplicationClient
type AccountClient = original.AccountClient
type ApplicationPackageClient = original.ApplicationPackageClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationClient(subscriptionID string) ApplicationClient {
	return original.NewApplicationClient(subscriptionID)
}
func NewApplicationClientWithBaseURI(baseURI string, subscriptionID string) ApplicationClient {
	return original.NewApplicationClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccountClient(subscriptionID string) AccountClient {
	return original.NewAccountClient(subscriptionID)
}
func NewAccountClientWithBaseURI(baseURI string, subscriptionID string) AccountClient {
	return original.NewAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationPackageClient(subscriptionID string) ApplicationPackageClient {
	return original.NewApplicationPackageClient(subscriptionID)
}
func NewApplicationPackageClientWithBaseURI(baseURI string, subscriptionID string) ApplicationPackageClient {
	return original.NewApplicationPackageClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func NewCertificateClient(subscriptionID string) CertificateClient {
	return original.NewCertificateClient(subscriptionID)
}
func NewCertificateClientWithBaseURI(baseURI string, subscriptionID string) CertificateClient {
	return original.NewCertificateClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewLocationClient(subscriptionID string) LocationClient {
	return original.NewLocationClient(subscriptionID)
}
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string) LocationClient {
	return original.NewLocationClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return original.PossibleAccountKeyTypeValues()
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleAutoUserScopeValues() []AutoUserScope {
	return original.PossibleAutoUserScopeValues()
}
func PossibleCachingTypeValues() []CachingType {
	return original.PossibleCachingTypeValues()
}
func PossibleCertificateFormatValues() []CertificateFormat {
	return original.PossibleCertificateFormatValues()
}
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return original.PossibleCertificateProvisioningStateValues()
}
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return original.PossibleCertificateStoreLocationValues()
}
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return original.PossibleCertificateVisibilityValues()
}
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return original.PossibleComputeNodeDeallocationOptionValues()
}
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return original.PossibleComputeNodeFillTypeValues()
}
func PossibleElevationLevelValues() []ElevationLevel {
	return original.PossibleElevationLevelValues()
}
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return original.PossibleInboundEndpointProtocolValues()
}
func PossibleInterNodeCommunicationStateValues() []InterNodeCommunicationState {
	return original.PossibleInterNodeCommunicationStateValues()
}
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return original.PossibleNameAvailabilityReasonValues()
}
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return original.PossibleNetworkSecurityGroupRuleAccessValues()
}
func PossiblePackageStateValues() []PackageState {
	return original.PossiblePackageStateValues()
}
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return original.PossiblePoolAllocationModeValues()
}
func PossiblePoolProvisioningStateValues() []PoolProvisioningState {
	return original.PossiblePoolProvisioningStateValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return original.PossibleStorageAccountTypeValues()
}
func NewPoolClient(subscriptionID string) PoolClient {
	return original.NewPoolClient(subscriptionID)
}
func NewPoolClientWithBaseURI(baseURI string, subscriptionID string) PoolClient {
	return original.NewPoolClientWithBaseURI(baseURI, subscriptionID)
}
