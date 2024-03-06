//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple8000series

// AccessControlRecordsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccessControlRecordsClient.BeginCreateOrUpdate
// method.
type AccessControlRecordsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccessControlRecordsClientBeginDeleteOptions contains the optional parameters for the AccessControlRecordsClient.BeginDelete
// method.
type AccessControlRecordsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccessControlRecordsClientGetOptions contains the optional parameters for the AccessControlRecordsClient.Get method.
type AccessControlRecordsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccessControlRecordsClientListByManagerOptions contains the optional parameters for the AccessControlRecordsClient.NewListByManagerPager
// method.
type AccessControlRecordsClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientClearOptions contains the optional parameters for the AlertsClient.Clear method.
type AlertsClientClearOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientListByManagerOptions contains the optional parameters for the AlertsClient.NewListByManagerPager method.
type AlertsClientListByManagerOptions struct {
	// OData Filter options
	Filter *string
}

// AlertsClientSendTestEmailOptions contains the optional parameters for the AlertsClient.SendTestEmail method.
type AlertsClientSendTestEmailOptions struct {
	// placeholder for future optional parameters
}

// BackupPoliciesClientBeginBackupNowOptions contains the optional parameters for the BackupPoliciesClient.BeginBackupNow
// method.
type BackupPoliciesClientBeginBackupNowOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupPoliciesClient.BeginCreateOrUpdate
// method.
type BackupPoliciesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupPoliciesClientBeginDeleteOptions contains the optional parameters for the BackupPoliciesClient.BeginDelete method.
type BackupPoliciesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupPoliciesClientGetOptions contains the optional parameters for the BackupPoliciesClient.Get method.
type BackupPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// BackupPoliciesClientListByDeviceOptions contains the optional parameters for the BackupPoliciesClient.NewListByDevicePager
// method.
type BackupPoliciesClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// BackupSchedulesClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupSchedulesClient.BeginCreateOrUpdate
// method.
type BackupSchedulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupSchedulesClientBeginDeleteOptions contains the optional parameters for the BackupSchedulesClient.BeginDelete method.
type BackupSchedulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupSchedulesClientGetOptions contains the optional parameters for the BackupSchedulesClient.Get method.
type BackupSchedulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// BackupSchedulesClientListByBackupPolicyOptions contains the optional parameters for the BackupSchedulesClient.NewListByBackupPolicyPager
// method.
type BackupSchedulesClientListByBackupPolicyOptions struct {
	// placeholder for future optional parameters
}

// BackupsClientBeginCloneOptions contains the optional parameters for the BackupsClient.BeginClone method.
type BackupsClientBeginCloneOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
type BackupsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupsClientBeginRestoreOptions contains the optional parameters for the BackupsClient.BeginRestore method.
type BackupsClientBeginRestoreOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupsClientListByDeviceOptions contains the optional parameters for the BackupsClient.NewListByDevicePager method.
type BackupsClientListByDeviceOptions struct {
	// OData Filter options
	Filter *string
}

// BandwidthSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the BandwidthSettingsClient.BeginCreateOrUpdate
// method.
type BandwidthSettingsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BandwidthSettingsClientBeginDeleteOptions contains the optional parameters for the BandwidthSettingsClient.BeginDelete
// method.
type BandwidthSettingsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BandwidthSettingsClientGetOptions contains the optional parameters for the BandwidthSettingsClient.Get method.
type BandwidthSettingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BandwidthSettingsClientListByManagerOptions contains the optional parameters for the BandwidthSettingsClient.NewListByManagerPager
// method.
type BandwidthSettingsClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// CloudAppliancesClientBeginProvisionOptions contains the optional parameters for the CloudAppliancesClient.BeginProvision
// method.
type CloudAppliancesClientBeginProvisionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CloudAppliancesClientListSupportedConfigurationsOptions contains the optional parameters for the CloudAppliancesClient.NewListSupportedConfigurationsPager
// method.
type CloudAppliancesClientListSupportedConfigurationsOptions struct {
	// placeholder for future optional parameters
}

// DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginCreateOrUpdateAlertSettings
// method.
type DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginCreateOrUpdateTimeSettings
// method.
type DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions contains the optional parameters for the DeviceSettingsClient.BeginSyncRemotemanagementCertificate
// method.
type DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceSettingsClientBeginUpdateNetworkSettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginUpdateNetworkSettings
// method.
type DeviceSettingsClientBeginUpdateNetworkSettingsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceSettingsClientBeginUpdateSecuritySettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginUpdateSecuritySettings
// method.
type DeviceSettingsClientBeginUpdateSecuritySettingsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceSettingsClientGetAlertSettingsOptions contains the optional parameters for the DeviceSettingsClient.GetAlertSettings
// method.
type DeviceSettingsClientGetAlertSettingsOptions struct {
	// placeholder for future optional parameters
}

// DeviceSettingsClientGetNetworkSettingsOptions contains the optional parameters for the DeviceSettingsClient.GetNetworkSettings
// method.
type DeviceSettingsClientGetNetworkSettingsOptions struct {
	// placeholder for future optional parameters
}

// DeviceSettingsClientGetSecuritySettingsOptions contains the optional parameters for the DeviceSettingsClient.GetSecuritySettings
// method.
type DeviceSettingsClientGetSecuritySettingsOptions struct {
	// placeholder for future optional parameters
}

// DeviceSettingsClientGetTimeSettingsOptions contains the optional parameters for the DeviceSettingsClient.GetTimeSettings
// method.
type DeviceSettingsClientGetTimeSettingsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientAuthorizeForServiceEncryptionKeyRolloverOptions contains the optional parameters for the DevicesClient.AuthorizeForServiceEncryptionKeyRollover
// method.
type DevicesClientAuthorizeForServiceEncryptionKeyRolloverOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientBeginConfigureOptions contains the optional parameters for the DevicesClient.BeginConfigure method.
type DevicesClientBeginConfigureOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginDeactivateOptions contains the optional parameters for the DevicesClient.BeginDeactivate method.
type DevicesClientBeginDeactivateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginDeleteOptions contains the optional parameters for the DevicesClient.BeginDelete method.
type DevicesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginFailoverOptions contains the optional parameters for the DevicesClient.BeginFailover method.
type DevicesClientBeginFailoverOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginInstallUpdatesOptions contains the optional parameters for the DevicesClient.BeginInstallUpdates method.
type DevicesClientBeginInstallUpdatesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginScanForUpdatesOptions contains the optional parameters for the DevicesClient.BeginScanForUpdates method.
type DevicesClientBeginScanForUpdatesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
type DevicesClientGetOptions struct {
	// Specify $expand=details to populate additional fields related to the device or $expand=rolloverdetails to populate additional
	// fields related to the service data encryption key rollover on device
	Expand *string
}

// DevicesClientGetUpdateSummaryOptions contains the optional parameters for the DevicesClient.GetUpdateSummary method.
type DevicesClientGetUpdateSummaryOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListByManagerOptions contains the optional parameters for the DevicesClient.NewListByManagerPager method.
type DevicesClientListByManagerOptions struct {
	// Specify $expand=details to populate additional fields related to the device or $expand=rolloverdetails to populate additional
	// fields related to the service data encryption key rollover on device
	Expand *string
}

// DevicesClientListFailoverSetsOptions contains the optional parameters for the DevicesClient.NewListFailoverSetsPager method.
type DevicesClientListFailoverSetsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListFailoverTargetsOptions contains the optional parameters for the DevicesClient.NewListFailoverTargetsPager
// method.
type DevicesClientListFailoverTargetsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListMetricDefinitionOptions contains the optional parameters for the DevicesClient.NewListMetricDefinitionPager
// method.
type DevicesClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListMetricsOptions contains the optional parameters for the DevicesClient.NewListMetricsPager method.
type DevicesClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientUpdateOptions contains the optional parameters for the DevicesClient.Update method.
type DevicesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// HardwareComponentGroupsClientBeginChangeControllerPowerStateOptions contains the optional parameters for the HardwareComponentGroupsClient.BeginChangeControllerPowerState
// method.
type HardwareComponentGroupsClientBeginChangeControllerPowerStateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// HardwareComponentGroupsClientListByDeviceOptions contains the optional parameters for the HardwareComponentGroupsClient.NewListByDevicePager
// method.
type HardwareComponentGroupsClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// JobsClientBeginCancelOptions contains the optional parameters for the JobsClient.BeginCancel method.
type JobsClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
type JobsClientGetOptions struct {
	// placeholder for future optional parameters
}

// JobsClientListByDeviceOptions contains the optional parameters for the JobsClient.NewListByDevicePager method.
type JobsClientListByDeviceOptions struct {
	// OData Filter options
	Filter *string
}

// JobsClientListByManagerOptions contains the optional parameters for the JobsClient.NewListByManagerPager method.
type JobsClientListByManagerOptions struct {
	// OData Filter options
	Filter *string
}

// ManagersClientCreateExtendedInfoOptions contains the optional parameters for the ManagersClient.CreateExtendedInfo method.
type ManagersClientCreateExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientCreateOrUpdateOptions contains the optional parameters for the ManagersClient.CreateOrUpdate method.
type ManagersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientDeleteExtendedInfoOptions contains the optional parameters for the ManagersClient.DeleteExtendedInfo method.
type ManagersClientDeleteExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientDeleteOptions contains the optional parameters for the ManagersClient.Delete method.
type ManagersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetActivationKeyOptions contains the optional parameters for the ManagersClient.GetActivationKey method.
type ManagersClientGetActivationKeyOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetDevicePublicEncryptionKeyOptions contains the optional parameters for the ManagersClient.GetDevicePublicEncryptionKey
// method.
type ManagersClientGetDevicePublicEncryptionKeyOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetEncryptionSettingsOptions contains the optional parameters for the ManagersClient.GetEncryptionSettings
// method.
type ManagersClientGetEncryptionSettingsOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetExtendedInfoOptions contains the optional parameters for the ManagersClient.GetExtendedInfo method.
type ManagersClientGetExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetOptions contains the optional parameters for the ManagersClient.Get method.
type ManagersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetPublicEncryptionKeyOptions contains the optional parameters for the ManagersClient.GetPublicEncryptionKey
// method.
type ManagersClientGetPublicEncryptionKeyOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListByResourceGroupOptions contains the optional parameters for the ManagersClient.NewListByResourceGroupPager
// method.
type ManagersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListFeatureSupportStatusOptions contains the optional parameters for the ManagersClient.NewListFeatureSupportStatusPager
// method.
type ManagersClientListFeatureSupportStatusOptions struct {
	// OData Filter options
	Filter *string
}

// ManagersClientListMetricDefinitionOptions contains the optional parameters for the ManagersClient.NewListMetricDefinitionPager
// method.
type ManagersClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListMetricsOptions contains the optional parameters for the ManagersClient.NewListMetricsPager method.
type ManagersClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListOptions contains the optional parameters for the ManagersClient.NewListPager method.
type ManagersClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientRegenerateActivationKeyOptions contains the optional parameters for the ManagersClient.RegenerateActivationKey
// method.
type ManagersClientRegenerateActivationKeyOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientUpdateExtendedInfoOptions contains the optional parameters for the ManagersClient.UpdateExtendedInfo method.
type ManagersClientUpdateExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientUpdateOptions contains the optional parameters for the ManagersClient.Update method.
type ManagersClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// StorageAccountCredentialsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginCreateOrUpdate
// method.
type StorageAccountCredentialsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageAccountCredentialsClientBeginDeleteOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginDelete
// method.
type StorageAccountCredentialsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageAccountCredentialsClientGetOptions contains the optional parameters for the StorageAccountCredentialsClient.Get
// method.
type StorageAccountCredentialsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StorageAccountCredentialsClientListByManagerOptions contains the optional parameters for the StorageAccountCredentialsClient.NewListByManagerPager
// method.
type StorageAccountCredentialsClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// VolumeContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the VolumeContainersClient.BeginCreateOrUpdate
// method.
type VolumeContainersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VolumeContainersClientBeginDeleteOptions contains the optional parameters for the VolumeContainersClient.BeginDelete method.
type VolumeContainersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VolumeContainersClientGetOptions contains the optional parameters for the VolumeContainersClient.Get method.
type VolumeContainersClientGetOptions struct {
	// placeholder for future optional parameters
}

// VolumeContainersClientListByDeviceOptions contains the optional parameters for the VolumeContainersClient.NewListByDevicePager
// method.
type VolumeContainersClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// VolumeContainersClientListMetricDefinitionOptions contains the optional parameters for the VolumeContainersClient.NewListMetricDefinitionPager
// method.
type VolumeContainersClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// VolumeContainersClientListMetricsOptions contains the optional parameters for the VolumeContainersClient.NewListMetricsPager
// method.
type VolumeContainersClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientBeginCreateOrUpdateOptions contains the optional parameters for the VolumesClient.BeginCreateOrUpdate method.
type VolumesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VolumesClientBeginDeleteOptions contains the optional parameters for the VolumesClient.BeginDelete method.
type VolumesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VolumesClientGetOptions contains the optional parameters for the VolumesClient.Get method.
type VolumesClientGetOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListByDeviceOptions contains the optional parameters for the VolumesClient.NewListByDevicePager method.
type VolumesClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListByVolumeContainerOptions contains the optional parameters for the VolumesClient.NewListByVolumeContainerPager
// method.
type VolumesClientListByVolumeContainerOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListMetricDefinitionOptions contains the optional parameters for the VolumesClient.NewListMetricDefinitionPager
// method.
type VolumesClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListMetricsOptions contains the optional parameters for the VolumesClient.NewListMetricsPager method.
type VolumesClientListMetricsOptions struct {
	// placeholder for future optional parameters
}
