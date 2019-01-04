// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package compute

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/compute/mgmt/2016-04-30-preview/compute"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessLevel = original.AccessLevel

const (
	None AccessLevel = original.None
	Read AccessLevel = original.Read
)

type CachingTypes = original.CachingTypes

const (
	CachingTypesNone      CachingTypes = original.CachingTypesNone
	CachingTypesReadOnly  CachingTypes = original.CachingTypesReadOnly
	CachingTypesReadWrite CachingTypes = original.CachingTypesReadWrite
)

type ComponentNames = original.ComponentNames

const (
	MicrosoftWindowsShellSetup ComponentNames = original.MicrosoftWindowsShellSetup
)

type DiskCreateOption = original.DiskCreateOption

const (
	Attach    DiskCreateOption = original.Attach
	Copy      DiskCreateOption = original.Copy
	Empty     DiskCreateOption = original.Empty
	FromImage DiskCreateOption = original.FromImage
	Import    DiskCreateOption = original.Import
	Restore   DiskCreateOption = original.Restore
)

type DiskCreateOptionTypes = original.DiskCreateOptionTypes

const (
	DiskCreateOptionTypesAttach    DiskCreateOptionTypes = original.DiskCreateOptionTypesAttach
	DiskCreateOptionTypesEmpty     DiskCreateOptionTypes = original.DiskCreateOptionTypesEmpty
	DiskCreateOptionTypesFromImage DiskCreateOptionTypes = original.DiskCreateOptionTypesFromImage
)

type InstanceViewTypes = original.InstanceViewTypes

const (
	InstanceView InstanceViewTypes = original.InstanceView
)

type OperatingSystemStateTypes = original.OperatingSystemStateTypes

const (
	Generalized OperatingSystemStateTypes = original.Generalized
	Specialized OperatingSystemStateTypes = original.Specialized
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	Linux   OperatingSystemTypes = original.Linux
	Windows OperatingSystemTypes = original.Windows
)

type PassNames = original.PassNames

const (
	OobeSystem PassNames = original.OobeSystem
)

type ProtocolTypes = original.ProtocolTypes

const (
	HTTP  ProtocolTypes = original.HTTP
	HTTPS ProtocolTypes = original.HTTPS
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type SettingNames = original.SettingNames

const (
	AutoLogon          SettingNames = original.AutoLogon
	FirstLogonCommands SettingNames = original.FirstLogonCommands
)

type StatusLevelTypes = original.StatusLevelTypes

const (
	Error   StatusLevelTypes = original.Error
	Info    StatusLevelTypes = original.Info
	Warning StatusLevelTypes = original.Warning
)

type StorageAccountTypes = original.StorageAccountTypes

const (
	PremiumLRS  StorageAccountTypes = original.PremiumLRS
	StandardLRS StorageAccountTypes = original.StandardLRS
)

type UpgradeMode = original.UpgradeMode

const (
	Automatic UpgradeMode = original.Automatic
	Manual    UpgradeMode = original.Manual
)

type VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleType

const (
	VirtualMachineScaleSetSkuScaleTypeAutomatic VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleTypeAutomatic
	VirtualMachineScaleSetSkuScaleTypeNone      VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleTypeNone
)

type VirtualMachineSizeTypes = original.VirtualMachineSizeTypes

const (
	BasicA0        VirtualMachineSizeTypes = original.BasicA0
	BasicA1        VirtualMachineSizeTypes = original.BasicA1
	BasicA2        VirtualMachineSizeTypes = original.BasicA2
	BasicA3        VirtualMachineSizeTypes = original.BasicA3
	BasicA4        VirtualMachineSizeTypes = original.BasicA4
	StandardA0     VirtualMachineSizeTypes = original.StandardA0
	StandardA1     VirtualMachineSizeTypes = original.StandardA1
	StandardA10    VirtualMachineSizeTypes = original.StandardA10
	StandardA11    VirtualMachineSizeTypes = original.StandardA11
	StandardA2     VirtualMachineSizeTypes = original.StandardA2
	StandardA3     VirtualMachineSizeTypes = original.StandardA3
	StandardA4     VirtualMachineSizeTypes = original.StandardA4
	StandardA5     VirtualMachineSizeTypes = original.StandardA5
	StandardA6     VirtualMachineSizeTypes = original.StandardA6
	StandardA7     VirtualMachineSizeTypes = original.StandardA7
	StandardA8     VirtualMachineSizeTypes = original.StandardA8
	StandardA9     VirtualMachineSizeTypes = original.StandardA9
	StandardD1     VirtualMachineSizeTypes = original.StandardD1
	StandardD11    VirtualMachineSizeTypes = original.StandardD11
	StandardD11V2  VirtualMachineSizeTypes = original.StandardD11V2
	StandardD12    VirtualMachineSizeTypes = original.StandardD12
	StandardD12V2  VirtualMachineSizeTypes = original.StandardD12V2
	StandardD13    VirtualMachineSizeTypes = original.StandardD13
	StandardD13V2  VirtualMachineSizeTypes = original.StandardD13V2
	StandardD14    VirtualMachineSizeTypes = original.StandardD14
	StandardD14V2  VirtualMachineSizeTypes = original.StandardD14V2
	StandardD15V2  VirtualMachineSizeTypes = original.StandardD15V2
	StandardD1V2   VirtualMachineSizeTypes = original.StandardD1V2
	StandardD2     VirtualMachineSizeTypes = original.StandardD2
	StandardD2V2   VirtualMachineSizeTypes = original.StandardD2V2
	StandardD3     VirtualMachineSizeTypes = original.StandardD3
	StandardD3V2   VirtualMachineSizeTypes = original.StandardD3V2
	StandardD4     VirtualMachineSizeTypes = original.StandardD4
	StandardD4V2   VirtualMachineSizeTypes = original.StandardD4V2
	StandardD5V2   VirtualMachineSizeTypes = original.StandardD5V2
	StandardDS1    VirtualMachineSizeTypes = original.StandardDS1
	StandardDS11   VirtualMachineSizeTypes = original.StandardDS11
	StandardDS11V2 VirtualMachineSizeTypes = original.StandardDS11V2
	StandardDS12   VirtualMachineSizeTypes = original.StandardDS12
	StandardDS12V2 VirtualMachineSizeTypes = original.StandardDS12V2
	StandardDS13   VirtualMachineSizeTypes = original.StandardDS13
	StandardDS13V2 VirtualMachineSizeTypes = original.StandardDS13V2
	StandardDS14   VirtualMachineSizeTypes = original.StandardDS14
	StandardDS14V2 VirtualMachineSizeTypes = original.StandardDS14V2
	StandardDS15V2 VirtualMachineSizeTypes = original.StandardDS15V2
	StandardDS1V2  VirtualMachineSizeTypes = original.StandardDS1V2
	StandardDS2    VirtualMachineSizeTypes = original.StandardDS2
	StandardDS2V2  VirtualMachineSizeTypes = original.StandardDS2V2
	StandardDS3    VirtualMachineSizeTypes = original.StandardDS3
	StandardDS3V2  VirtualMachineSizeTypes = original.StandardDS3V2
	StandardDS4    VirtualMachineSizeTypes = original.StandardDS4
	StandardDS4V2  VirtualMachineSizeTypes = original.StandardDS4V2
	StandardDS5V2  VirtualMachineSizeTypes = original.StandardDS5V2
	StandardG1     VirtualMachineSizeTypes = original.StandardG1
	StandardG2     VirtualMachineSizeTypes = original.StandardG2
	StandardG3     VirtualMachineSizeTypes = original.StandardG3
	StandardG4     VirtualMachineSizeTypes = original.StandardG4
	StandardG5     VirtualMachineSizeTypes = original.StandardG5
	StandardGS1    VirtualMachineSizeTypes = original.StandardGS1
	StandardGS2    VirtualMachineSizeTypes = original.StandardGS2
	StandardGS3    VirtualMachineSizeTypes = original.StandardGS3
	StandardGS4    VirtualMachineSizeTypes = original.StandardGS4
	StandardGS5    VirtualMachineSizeTypes = original.StandardGS5
)

type APIEntityReference = original.APIEntityReference
type APIError = original.APIError
type APIErrorBase = original.APIErrorBase
type AccessURI = original.AccessURI
type AccessURIOutput = original.AccessURIOutput
type AccessURIRaw = original.AccessURIRaw
type AdditionalUnattendContent = original.AdditionalUnattendContent
type AvailabilitySet = original.AvailabilitySet
type AvailabilitySetListResult = original.AvailabilitySetListResult
type AvailabilitySetListResultIterator = original.AvailabilitySetListResultIterator
type AvailabilitySetListResultPage = original.AvailabilitySetListResultPage
type AvailabilitySetProperties = original.AvailabilitySetProperties
type AvailabilitySetsClient = original.AvailabilitySetsClient
type BaseClient = original.BaseClient
type BootDiagnostics = original.BootDiagnostics
type BootDiagnosticsInstanceView = original.BootDiagnosticsInstanceView
type CreationData = original.CreationData
type DataDisk = original.DataDisk
type DataDiskImage = original.DataDiskImage
type DiagnosticsProfile = original.DiagnosticsProfile
type Disk = original.Disk
type DiskEncryptionSettings = original.DiskEncryptionSettings
type DiskInstanceView = original.DiskInstanceView
type DiskList = original.DiskList
type DiskListIterator = original.DiskListIterator
type DiskListPage = original.DiskListPage
type DiskProperties = original.DiskProperties
type DiskUpdate = original.DiskUpdate
type DiskUpdateProperties = original.DiskUpdateProperties
type DisksClient = original.DisksClient
type DisksCreateOrUpdateFuture = original.DisksCreateOrUpdateFuture
type DisksDeleteFuture = original.DisksDeleteFuture
type DisksGrantAccessFuture = original.DisksGrantAccessFuture
type DisksRevokeAccessFuture = original.DisksRevokeAccessFuture
type DisksUpdateFuture = original.DisksUpdateFuture
type EncryptionSettings = original.EncryptionSettings
type GrantAccessData = original.GrantAccessData
type HardwareProfile = original.HardwareProfile
type Image = original.Image
type ImageDataDisk = original.ImageDataDisk
type ImageDiskReference = original.ImageDiskReference
type ImageListResult = original.ImageListResult
type ImageListResultIterator = original.ImageListResultIterator
type ImageListResultPage = original.ImageListResultPage
type ImageOSDisk = original.ImageOSDisk
type ImageProperties = original.ImageProperties
type ImageReference = original.ImageReference
type ImageStorageProfile = original.ImageStorageProfile
type ImagesClient = original.ImagesClient
type ImagesCreateOrUpdateFuture = original.ImagesCreateOrUpdateFuture
type ImagesDeleteFuture = original.ImagesDeleteFuture
type InnerError = original.InnerError
type InstanceViewStatus = original.InstanceViewStatus
type KeyVaultAndKeyReference = original.KeyVaultAndKeyReference
type KeyVaultAndSecretReference = original.KeyVaultAndSecretReference
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultSecretReference = original.KeyVaultSecretReference
type LinuxConfiguration = original.LinuxConfiguration
type ListUsagesResult = original.ListUsagesResult
type ListUsagesResultIterator = original.ListUsagesResultIterator
type ListUsagesResultPage = original.ListUsagesResultPage
type ListVirtualMachineExtensionImage = original.ListVirtualMachineExtensionImage
type ListVirtualMachineImageResource = original.ListVirtualMachineImageResource
type LongRunningOperationProperties = original.LongRunningOperationProperties
type ManagedDiskParameters = original.ManagedDiskParameters
type NetworkInterfaceReference = original.NetworkInterfaceReference
type NetworkInterfaceReferenceProperties = original.NetworkInterfaceReferenceProperties
type NetworkProfile = original.NetworkProfile
type OSDisk = original.OSDisk
type OSDiskImage = original.OSDiskImage
type OSProfile = original.OSProfile
type OperationStatusResponse = original.OperationStatusResponse
type Plan = original.Plan
type PurchasePlan = original.PurchasePlan
type Resource = original.Resource
type ResourceUpdate = original.ResourceUpdate
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type Sku = original.Sku
type Snapshot = original.Snapshot
type SnapshotList = original.SnapshotList
type SnapshotListIterator = original.SnapshotListIterator
type SnapshotListPage = original.SnapshotListPage
type SnapshotUpdate = original.SnapshotUpdate
type SnapshotsClient = original.SnapshotsClient
type SnapshotsCreateOrUpdateFuture = original.SnapshotsCreateOrUpdateFuture
type SnapshotsDeleteFuture = original.SnapshotsDeleteFuture
type SnapshotsGrantAccessFuture = original.SnapshotsGrantAccessFuture
type SnapshotsRevokeAccessFuture = original.SnapshotsRevokeAccessFuture
type SnapshotsUpdateFuture = original.SnapshotsUpdateFuture
type SourceVault = original.SourceVault
type StorageProfile = original.StorageProfile
type SubResource = original.SubResource
type SubResourceReadOnly = original.SubResourceReadOnly
type UpdateResource = original.UpdateResource
type UpgradePolicy = original.UpgradePolicy
type Usage = original.Usage
type UsageClient = original.UsageClient
type UsageName = original.UsageName
type VaultCertificate = original.VaultCertificate
type VaultSecretGroup = original.VaultSecretGroup
type VirtualHardDisk = original.VirtualHardDisk
type VirtualMachine = original.VirtualMachine
type VirtualMachineAgentInstanceView = original.VirtualMachineAgentInstanceView
type VirtualMachineCaptureParameters = original.VirtualMachineCaptureParameters
type VirtualMachineCaptureResult = original.VirtualMachineCaptureResult
type VirtualMachineCaptureResultProperties = original.VirtualMachineCaptureResultProperties
type VirtualMachineExtension = original.VirtualMachineExtension
type VirtualMachineExtensionHandlerInstanceView = original.VirtualMachineExtensionHandlerInstanceView
type VirtualMachineExtensionImage = original.VirtualMachineExtensionImage
type VirtualMachineExtensionImageProperties = original.VirtualMachineExtensionImageProperties
type VirtualMachineExtensionImagesClient = original.VirtualMachineExtensionImagesClient
type VirtualMachineExtensionInstanceView = original.VirtualMachineExtensionInstanceView
type VirtualMachineExtensionProperties = original.VirtualMachineExtensionProperties
type VirtualMachineExtensionUpdate = original.VirtualMachineExtensionUpdate
type VirtualMachineExtensionUpdateProperties = original.VirtualMachineExtensionUpdateProperties
type VirtualMachineExtensionsClient = original.VirtualMachineExtensionsClient
type VirtualMachineExtensionsCreateOrUpdateFuture = original.VirtualMachineExtensionsCreateOrUpdateFuture
type VirtualMachineExtensionsDeleteFuture = original.VirtualMachineExtensionsDeleteFuture
type VirtualMachineExtensionsListResult = original.VirtualMachineExtensionsListResult
type VirtualMachineExtensionsUpdateFuture = original.VirtualMachineExtensionsUpdateFuture
type VirtualMachineIdentity = original.VirtualMachineIdentity
type VirtualMachineImage = original.VirtualMachineImage
type VirtualMachineImageProperties = original.VirtualMachineImageProperties
type VirtualMachineImageResource = original.VirtualMachineImageResource
type VirtualMachineImagesClient = original.VirtualMachineImagesClient
type VirtualMachineInstanceView = original.VirtualMachineInstanceView
type VirtualMachineListResult = original.VirtualMachineListResult
type VirtualMachineListResultIterator = original.VirtualMachineListResultIterator
type VirtualMachineListResultPage = original.VirtualMachineListResultPage
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineScaleSet = original.VirtualMachineScaleSet
type VirtualMachineScaleSetDataDisk = original.VirtualMachineScaleSetDataDisk
type VirtualMachineScaleSetExtension = original.VirtualMachineScaleSetExtension
type VirtualMachineScaleSetExtensionProfile = original.VirtualMachineScaleSetExtensionProfile
type VirtualMachineScaleSetExtensionProperties = original.VirtualMachineScaleSetExtensionProperties
type VirtualMachineScaleSetIPConfiguration = original.VirtualMachineScaleSetIPConfiguration
type VirtualMachineScaleSetIPConfigurationProperties = original.VirtualMachineScaleSetIPConfigurationProperties
type VirtualMachineScaleSetIdentity = original.VirtualMachineScaleSetIdentity
type VirtualMachineScaleSetInstanceView = original.VirtualMachineScaleSetInstanceView
type VirtualMachineScaleSetInstanceViewStatusesSummary = original.VirtualMachineScaleSetInstanceViewStatusesSummary
type VirtualMachineScaleSetListResult = original.VirtualMachineScaleSetListResult
type VirtualMachineScaleSetListResultIterator = original.VirtualMachineScaleSetListResultIterator
type VirtualMachineScaleSetListResultPage = original.VirtualMachineScaleSetListResultPage
type VirtualMachineScaleSetListSkusResult = original.VirtualMachineScaleSetListSkusResult
type VirtualMachineScaleSetListSkusResultIterator = original.VirtualMachineScaleSetListSkusResultIterator
type VirtualMachineScaleSetListSkusResultPage = original.VirtualMachineScaleSetListSkusResultPage
type VirtualMachineScaleSetListWithLinkResult = original.VirtualMachineScaleSetListWithLinkResult
type VirtualMachineScaleSetListWithLinkResultIterator = original.VirtualMachineScaleSetListWithLinkResultIterator
type VirtualMachineScaleSetListWithLinkResultPage = original.VirtualMachineScaleSetListWithLinkResultPage
type VirtualMachineScaleSetManagedDiskParameters = original.VirtualMachineScaleSetManagedDiskParameters
type VirtualMachineScaleSetNetworkConfiguration = original.VirtualMachineScaleSetNetworkConfiguration
type VirtualMachineScaleSetNetworkConfigurationProperties = original.VirtualMachineScaleSetNetworkConfigurationProperties
type VirtualMachineScaleSetNetworkProfile = original.VirtualMachineScaleSetNetworkProfile
type VirtualMachineScaleSetOSDisk = original.VirtualMachineScaleSetOSDisk
type VirtualMachineScaleSetOSProfile = original.VirtualMachineScaleSetOSProfile
type VirtualMachineScaleSetProperties = original.VirtualMachineScaleSetProperties
type VirtualMachineScaleSetSku = original.VirtualMachineScaleSetSku
type VirtualMachineScaleSetSkuCapacity = original.VirtualMachineScaleSetSkuCapacity
type VirtualMachineScaleSetStorageProfile = original.VirtualMachineScaleSetStorageProfile
type VirtualMachineScaleSetVM = original.VirtualMachineScaleSetVM
type VirtualMachineScaleSetVMExtensionsSummary = original.VirtualMachineScaleSetVMExtensionsSummary
type VirtualMachineScaleSetVMInstanceIDs = original.VirtualMachineScaleSetVMInstanceIDs
type VirtualMachineScaleSetVMInstanceRequiredIDs = original.VirtualMachineScaleSetVMInstanceRequiredIDs
type VirtualMachineScaleSetVMInstanceView = original.VirtualMachineScaleSetVMInstanceView
type VirtualMachineScaleSetVMListResult = original.VirtualMachineScaleSetVMListResult
type VirtualMachineScaleSetVMListResultIterator = original.VirtualMachineScaleSetVMListResultIterator
type VirtualMachineScaleSetVMListResultPage = original.VirtualMachineScaleSetVMListResultPage
type VirtualMachineScaleSetVMProfile = original.VirtualMachineScaleSetVMProfile
type VirtualMachineScaleSetVMProperties = original.VirtualMachineScaleSetVMProperties
type VirtualMachineScaleSetVMsClient = original.VirtualMachineScaleSetVMsClient
type VirtualMachineScaleSetVMsDeallocateFuture = original.VirtualMachineScaleSetVMsDeallocateFuture
type VirtualMachineScaleSetVMsDeleteFuture = original.VirtualMachineScaleSetVMsDeleteFuture
type VirtualMachineScaleSetVMsPowerOffFuture = original.VirtualMachineScaleSetVMsPowerOffFuture
type VirtualMachineScaleSetVMsReimageAllFuture = original.VirtualMachineScaleSetVMsReimageAllFuture
type VirtualMachineScaleSetVMsReimageFuture = original.VirtualMachineScaleSetVMsReimageFuture
type VirtualMachineScaleSetVMsRestartFuture = original.VirtualMachineScaleSetVMsRestartFuture
type VirtualMachineScaleSetVMsStartFuture = original.VirtualMachineScaleSetVMsStartFuture
type VirtualMachineScaleSetsClient = original.VirtualMachineScaleSetsClient
type VirtualMachineScaleSetsCreateOrUpdateFuture = original.VirtualMachineScaleSetsCreateOrUpdateFuture
type VirtualMachineScaleSetsDeallocateFuture = original.VirtualMachineScaleSetsDeallocateFuture
type VirtualMachineScaleSetsDeleteFuture = original.VirtualMachineScaleSetsDeleteFuture
type VirtualMachineScaleSetsDeleteInstancesFuture = original.VirtualMachineScaleSetsDeleteInstancesFuture
type VirtualMachineScaleSetsPowerOffFuture = original.VirtualMachineScaleSetsPowerOffFuture
type VirtualMachineScaleSetsReimageAllFuture = original.VirtualMachineScaleSetsReimageAllFuture
type VirtualMachineScaleSetsReimageFuture = original.VirtualMachineScaleSetsReimageFuture
type VirtualMachineScaleSetsRestartFuture = original.VirtualMachineScaleSetsRestartFuture
type VirtualMachineScaleSetsStartFuture = original.VirtualMachineScaleSetsStartFuture
type VirtualMachineScaleSetsUpdateInstancesFuture = original.VirtualMachineScaleSetsUpdateInstancesFuture
type VirtualMachineSize = original.VirtualMachineSize
type VirtualMachineSizeListResult = original.VirtualMachineSizeListResult
type VirtualMachineSizesClient = original.VirtualMachineSizesClient
type VirtualMachineStatusCodeCount = original.VirtualMachineStatusCodeCount
type VirtualMachinesCaptureFuture = original.VirtualMachinesCaptureFuture
type VirtualMachinesClient = original.VirtualMachinesClient
type VirtualMachinesConvertToManagedDisksFuture = original.VirtualMachinesConvertToManagedDisksFuture
type VirtualMachinesCreateOrUpdateFuture = original.VirtualMachinesCreateOrUpdateFuture
type VirtualMachinesDeallocateFuture = original.VirtualMachinesDeallocateFuture
type VirtualMachinesDeleteFuture = original.VirtualMachinesDeleteFuture
type VirtualMachinesPowerOffFuture = original.VirtualMachinesPowerOffFuture
type VirtualMachinesRedeployFuture = original.VirtualMachinesRedeployFuture
type VirtualMachinesRestartFuture = original.VirtualMachinesRestartFuture
type VirtualMachinesStartFuture = original.VirtualMachinesStartFuture
type WinRMConfiguration = original.WinRMConfiguration
type WinRMListener = original.WinRMListener
type WindowsConfiguration = original.WindowsConfiguration

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailabilitySetListResultIterator(page AvailabilitySetListResultPage) AvailabilitySetListResultIterator {
	return original.NewAvailabilitySetListResultIterator(page)
}
func NewAvailabilitySetListResultPage(getNextPage func(context.Context, AvailabilitySetListResult) (AvailabilitySetListResult, error)) AvailabilitySetListResultPage {
	return original.NewAvailabilitySetListResultPage(getNextPage)
}
func NewAvailabilitySetsClient(subscriptionID string) AvailabilitySetsClient {
	return original.NewAvailabilitySetsClient(subscriptionID)
}
func NewAvailabilitySetsClientWithBaseURI(baseURI string, subscriptionID string) AvailabilitySetsClient {
	return original.NewAvailabilitySetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDiskListIterator(page DiskListPage) DiskListIterator {
	return original.NewDiskListIterator(page)
}
func NewDiskListPage(getNextPage func(context.Context, DiskList) (DiskList, error)) DiskListPage {
	return original.NewDiskListPage(getNextPage)
}
func NewDisksClient(subscriptionID string) DisksClient {
	return original.NewDisksClient(subscriptionID)
}
func NewDisksClientWithBaseURI(baseURI string, subscriptionID string) DisksClient {
	return original.NewDisksClientWithBaseURI(baseURI, subscriptionID)
}
func NewImageListResultIterator(page ImageListResultPage) ImageListResultIterator {
	return original.NewImageListResultIterator(page)
}
func NewImageListResultPage(getNextPage func(context.Context, ImageListResult) (ImageListResult, error)) ImageListResultPage {
	return original.NewImageListResultPage(getNextPage)
}
func NewImagesClient(subscriptionID string) ImagesClient {
	return original.NewImagesClient(subscriptionID)
}
func NewImagesClientWithBaseURI(baseURI string, subscriptionID string) ImagesClient {
	return original.NewImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewListUsagesResultIterator(page ListUsagesResultPage) ListUsagesResultIterator {
	return original.NewListUsagesResultIterator(page)
}
func NewListUsagesResultPage(getNextPage func(context.Context, ListUsagesResult) (ListUsagesResult, error)) ListUsagesResultPage {
	return original.NewListUsagesResultPage(getNextPage)
}
func NewSnapshotListIterator(page SnapshotListPage) SnapshotListIterator {
	return original.NewSnapshotListIterator(page)
}
func NewSnapshotListPage(getNextPage func(context.Context, SnapshotList) (SnapshotList, error)) SnapshotListPage {
	return original.NewSnapshotListPage(getNextPage)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageClient(subscriptionID string) UsageClient {
	return original.NewUsageClient(subscriptionID)
}
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return original.NewUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineExtensionImagesClient(subscriptionID string) VirtualMachineExtensionImagesClient {
	return original.NewVirtualMachineExtensionImagesClient(subscriptionID)
}
func NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionImagesClient {
	return original.NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineExtensionsClient(subscriptionID string) VirtualMachineExtensionsClient {
	return original.NewVirtualMachineExtensionsClient(subscriptionID)
}
func NewVirtualMachineExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionsClient {
	return original.NewVirtualMachineExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineImagesClient(subscriptionID string) VirtualMachineImagesClient {
	return original.NewVirtualMachineImagesClient(subscriptionID)
}
func NewVirtualMachineImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineImagesClient {
	return original.NewVirtualMachineImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineListResultIterator(page VirtualMachineListResultPage) VirtualMachineListResultIterator {
	return original.NewVirtualMachineListResultIterator(page)
}
func NewVirtualMachineListResultPage(getNextPage func(context.Context, VirtualMachineListResult) (VirtualMachineListResult, error)) VirtualMachineListResultPage {
	return original.NewVirtualMachineListResultPage(getNextPage)
}
func NewVirtualMachineScaleSetListResultIterator(page VirtualMachineScaleSetListResultPage) VirtualMachineScaleSetListResultIterator {
	return original.NewVirtualMachineScaleSetListResultIterator(page)
}
func NewVirtualMachineScaleSetListResultPage(getNextPage func(context.Context, VirtualMachineScaleSetListResult) (VirtualMachineScaleSetListResult, error)) VirtualMachineScaleSetListResultPage {
	return original.NewVirtualMachineScaleSetListResultPage(getNextPage)
}
func NewVirtualMachineScaleSetListSkusResultIterator(page VirtualMachineScaleSetListSkusResultPage) VirtualMachineScaleSetListSkusResultIterator {
	return original.NewVirtualMachineScaleSetListSkusResultIterator(page)
}
func NewVirtualMachineScaleSetListSkusResultPage(getNextPage func(context.Context, VirtualMachineScaleSetListSkusResult) (VirtualMachineScaleSetListSkusResult, error)) VirtualMachineScaleSetListSkusResultPage {
	return original.NewVirtualMachineScaleSetListSkusResultPage(getNextPage)
}
func NewVirtualMachineScaleSetListWithLinkResultIterator(page VirtualMachineScaleSetListWithLinkResultPage) VirtualMachineScaleSetListWithLinkResultIterator {
	return original.NewVirtualMachineScaleSetListWithLinkResultIterator(page)
}
func NewVirtualMachineScaleSetListWithLinkResultPage(getNextPage func(context.Context, VirtualMachineScaleSetListWithLinkResult) (VirtualMachineScaleSetListWithLinkResult, error)) VirtualMachineScaleSetListWithLinkResultPage {
	return original.NewVirtualMachineScaleSetListWithLinkResultPage(getNextPage)
}
func NewVirtualMachineScaleSetVMListResultIterator(page VirtualMachineScaleSetVMListResultPage) VirtualMachineScaleSetVMListResultIterator {
	return original.NewVirtualMachineScaleSetVMListResultIterator(page)
}
func NewVirtualMachineScaleSetVMListResultPage(getNextPage func(context.Context, VirtualMachineScaleSetVMListResult) (VirtualMachineScaleSetVMListResult, error)) VirtualMachineScaleSetVMListResultPage {
	return original.NewVirtualMachineScaleSetVMListResultPage(getNextPage)
}
func NewVirtualMachineScaleSetVMsClient(subscriptionID string) VirtualMachineScaleSetVMsClient {
	return original.NewVirtualMachineScaleSetVMsClient(subscriptionID)
}
func NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetVMsClient {
	return original.NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineScaleSetsClient(subscriptionID string) VirtualMachineScaleSetsClient {
	return original.NewVirtualMachineScaleSetsClient(subscriptionID)
}
func NewVirtualMachineScaleSetsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetsClient {
	return original.NewVirtualMachineScaleSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineSizesClient(subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClient(subscriptionID)
}
func NewVirtualMachineSizesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachinesClient(subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClient(subscriptionID)
}
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessLevelValues() []AccessLevel {
	return original.PossibleAccessLevelValues()
}
func PossibleCachingTypesValues() []CachingTypes {
	return original.PossibleCachingTypesValues()
}
func PossibleComponentNamesValues() []ComponentNames {
	return original.PossibleComponentNamesValues()
}
func PossibleDiskCreateOptionTypesValues() []DiskCreateOptionTypes {
	return original.PossibleDiskCreateOptionTypesValues()
}
func PossibleDiskCreateOptionValues() []DiskCreateOption {
	return original.PossibleDiskCreateOptionValues()
}
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return original.PossibleInstanceViewTypesValues()
}
func PossibleOperatingSystemStateTypesValues() []OperatingSystemStateTypes {
	return original.PossibleOperatingSystemStateTypesValues()
}
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return original.PossibleOperatingSystemTypesValues()
}
func PossiblePassNamesValues() []PassNames {
	return original.PossiblePassNamesValues()
}
func PossibleProtocolTypesValues() []ProtocolTypes {
	return original.PossibleProtocolTypesValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSettingNamesValues() []SettingNames {
	return original.PossibleSettingNamesValues()
}
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return original.PossibleStatusLevelTypesValues()
}
func PossibleStorageAccountTypesValues() []StorageAccountTypes {
	return original.PossibleStorageAccountTypesValues()
}
func PossibleUpgradeModeValues() []UpgradeMode {
	return original.PossibleUpgradeModeValues()
}
func PossibleVirtualMachineScaleSetSkuScaleTypeValues() []VirtualMachineScaleSetSkuScaleType {
	return original.PossibleVirtualMachineScaleSetSkuScaleTypeValues()
}
func PossibleVirtualMachineSizeTypesValues() []VirtualMachineSizeTypes {
	return original.PossibleVirtualMachineSizeTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
