//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox

const (
	moduleName    = "armdatabox"
	moduleVersion = "v0.3.0"
)

type AccessProtocol string

const (
	// AccessProtocolSMB - Server Message Block protocol(SMB).
	AccessProtocolSMB AccessProtocol = "SMB"
	// AccessProtocolNFS - Network File System protocol(NFS).
	AccessProtocolNFS AccessProtocol = "NFS"
)

// PossibleAccessProtocolValues returns the possible values for the AccessProtocol const type.
func PossibleAccessProtocolValues() []AccessProtocol {
	return []AccessProtocol{
		AccessProtocolSMB,
		AccessProtocolNFS,
	}
}

// AddressType - Type of address.
type AddressType string

const (
	// AddressTypeNone - Address type not known.
	AddressTypeNone AddressType = "None"
	// AddressTypeResidential - Residential Address.
	AddressTypeResidential AddressType = "Residential"
	// AddressTypeCommercial - Commercial Address.
	AddressTypeCommercial AddressType = "Commercial"
)

// PossibleAddressTypeValues returns the possible values for the AddressType const type.
func PossibleAddressTypeValues() []AddressType {
	return []AddressType{
		AddressTypeNone,
		AddressTypeResidential,
		AddressTypeCommercial,
	}
}

// AddressValidationStatus - The address validation status.
type AddressValidationStatus string

const (
	// AddressValidationStatusValid - Address provided is valid.
	AddressValidationStatusValid AddressValidationStatus = "Valid"
	// AddressValidationStatusInvalid - Address provided is invalid or not supported.
	AddressValidationStatusInvalid AddressValidationStatus = "Invalid"
	// AddressValidationStatusAmbiguous - Address provided is ambiguous, please choose one of the alternate addresses returned.
	AddressValidationStatusAmbiguous AddressValidationStatus = "Ambiguous"
)

// PossibleAddressValidationStatusValues returns the possible values for the AddressValidationStatus const type.
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return []AddressValidationStatus{
		AddressValidationStatusValid,
		AddressValidationStatusInvalid,
		AddressValidationStatusAmbiguous,
	}
}

// ClassDiscriminator - Indicates the type of job details.
type ClassDiscriminator string

const (
	// ClassDiscriminatorDataBox - Data Box orders.
	ClassDiscriminatorDataBox ClassDiscriminator = "DataBox"
	// ClassDiscriminatorDataBoxDisk - Data Box Disk orders.
	ClassDiscriminatorDataBoxDisk ClassDiscriminator = "DataBoxDisk"
	// ClassDiscriminatorDataBoxHeavy - Data Box Heavy orders.
	ClassDiscriminatorDataBoxHeavy ClassDiscriminator = "DataBoxHeavy"
	// ClassDiscriminatorDataBoxCustomerDisk - Data Box Customer Disk orders.
	ClassDiscriminatorDataBoxCustomerDisk ClassDiscriminator = "DataBoxCustomerDisk"
)

// PossibleClassDiscriminatorValues returns the possible values for the ClassDiscriminator const type.
func PossibleClassDiscriminatorValues() []ClassDiscriminator {
	return []ClassDiscriminator{
		ClassDiscriminatorDataBox,
		ClassDiscriminatorDataBoxDisk,
		ClassDiscriminatorDataBoxHeavy,
		ClassDiscriminatorDataBoxCustomerDisk,
	}
}

// CopyStatus - The Status of the copy
type CopyStatus string

const (
	// CopyStatusCompleted - Data copy completed.
	CopyStatusCompleted CopyStatus = "Completed"
	// CopyStatusCompletedWithErrors - Data copy completed with errors.
	CopyStatusCompletedWithErrors CopyStatus = "CompletedWithErrors"
	// CopyStatusDeviceFormatted - Data copy failed. The Device was formatted by user.
	CopyStatusDeviceFormatted CopyStatus = "DeviceFormatted"
	// CopyStatusDeviceMetadataModified - Data copy failed. Device metadata was modified by user.
	CopyStatusDeviceMetadataModified CopyStatus = "DeviceMetadataModified"
	// CopyStatusDriveCorrupted - Copy failed due to corrupted drive.
	CopyStatusDriveCorrupted CopyStatus = "DriveCorrupted"
	// CopyStatusDriveNotDetected - Copy failed due to disk detection error.
	CopyStatusDriveNotDetected CopyStatus = "DriveNotDetected"
	// CopyStatusDriveNotReceived - No copy triggered as device was not received.
	CopyStatusDriveNotReceived CopyStatus = "DriveNotReceived"
	// CopyStatusFailed - Data copy failed. No data was copied.
	CopyStatusFailed CopyStatus = "Failed"
	// CopyStatusHardwareError - The Device has hit hardware issues.
	CopyStatusHardwareError CopyStatus = "HardwareError"
	// CopyStatusInProgress - Data copy is in progress.
	CopyStatusInProgress CopyStatus = "InProgress"
	// CopyStatusMetadataFilesModifiedOrRemoved - Copy failed due to modified or removed metadata files.
	CopyStatusMetadataFilesModifiedOrRemoved CopyStatus = "MetadataFilesModifiedOrRemoved"
	// CopyStatusNotReturned - No copy triggered as device was not returned.
	CopyStatusNotReturned CopyStatus = "NotReturned"
	// CopyStatusNotStarted - Data copy hasn't started yet.
	CopyStatusNotStarted CopyStatus = "NotStarted"
	// CopyStatusOtherServiceError - Copy failed due to service error.
	CopyStatusOtherServiceError CopyStatus = "OtherServiceError"
	// CopyStatusOtherUserError - Copy failed due to user error.
	CopyStatusOtherUserError CopyStatus = "OtherUserError"
	// CopyStatusStorageAccountNotAccessible - Data copy failed. Storage Account was not accessible during copy.
	CopyStatusStorageAccountNotAccessible CopyStatus = "StorageAccountNotAccessible"
	// CopyStatusUnsupportedData - Data copy failed. The Device data content is not supported.
	CopyStatusUnsupportedData CopyStatus = "UnsupportedData"
	// CopyStatusUnsupportedDrive - No copy triggered as device type is not supported.
	CopyStatusUnsupportedDrive CopyStatus = "UnsupportedDrive"
)

// PossibleCopyStatusValues returns the possible values for the CopyStatus const type.
func PossibleCopyStatusValues() []CopyStatus {
	return []CopyStatus{
		CopyStatusCompleted,
		CopyStatusCompletedWithErrors,
		CopyStatusDeviceFormatted,
		CopyStatusDeviceMetadataModified,
		CopyStatusDriveCorrupted,
		CopyStatusDriveNotDetected,
		CopyStatusDriveNotReceived,
		CopyStatusFailed,
		CopyStatusHardwareError,
		CopyStatusInProgress,
		CopyStatusMetadataFilesModifiedOrRemoved,
		CopyStatusNotReturned,
		CopyStatusNotStarted,
		CopyStatusOtherServiceError,
		CopyStatusOtherUserError,
		CopyStatusStorageAccountNotAccessible,
		CopyStatusUnsupportedData,
		CopyStatusUnsupportedDrive,
	}
}

// CustomerResolutionCode - Resolution code provided by customer
type CustomerResolutionCode string

const (
	// CustomerResolutionCodeNone - No Resolution Yet
	CustomerResolutionCodeNone CustomerResolutionCode = "None"
	// CustomerResolutionCodeMoveToCleanUpDevice - Clean the device
	CustomerResolutionCodeMoveToCleanUpDevice CustomerResolutionCode = "MoveToCleanUpDevice"
	// CustomerResolutionCodeResume - Resume the job to same stage
	CustomerResolutionCodeResume CustomerResolutionCode = "Resume"
	// CustomerResolutionCodeRestart - Restart whole action.
	CustomerResolutionCodeRestart CustomerResolutionCode = "Restart"
	// CustomerResolutionCodeReachOutToOperation - Reach out to operation for further action.
	CustomerResolutionCodeReachOutToOperation CustomerResolutionCode = "ReachOutToOperation"
)

// PossibleCustomerResolutionCodeValues returns the possible values for the CustomerResolutionCode const type.
func PossibleCustomerResolutionCodeValues() []CustomerResolutionCode {
	return []CustomerResolutionCode{
		CustomerResolutionCodeNone,
		CustomerResolutionCodeMoveToCleanUpDevice,
		CustomerResolutionCodeResume,
		CustomerResolutionCodeRestart,
		CustomerResolutionCodeReachOutToOperation,
	}
}

// DataAccountType - Type of the account.
type DataAccountType string

const (
	// DataAccountTypeStorageAccount - Storage Accounts .
	DataAccountTypeStorageAccount DataAccountType = "StorageAccount"
	// DataAccountTypeManagedDisk - Azure Managed disk storage.
	DataAccountTypeManagedDisk DataAccountType = "ManagedDisk"
)

// PossibleDataAccountTypeValues returns the possible values for the DataAccountType const type.
func PossibleDataAccountTypeValues() []DataAccountType {
	return []DataAccountType{
		DataAccountTypeStorageAccount,
		DataAccountTypeManagedDisk,
	}
}

// DataCenterCode - DataCenter code.
type DataCenterCode string

const (
	DataCenterCodeAM2     DataCenterCode = "AM2"
	DataCenterCodeAMS06   DataCenterCode = "AMS06"
	DataCenterCodeAMS20   DataCenterCode = "AMS20"
	DataCenterCodeAUH20   DataCenterCode = "AUH20"
	DataCenterCodeAdHoc   DataCenterCode = "AdHoc"
	DataCenterCodeBJB     DataCenterCode = "BJB"
	DataCenterCodeBJS20   DataCenterCode = "BJS20"
	DataCenterCodeBL20    DataCenterCode = "BL20"
	DataCenterCodeBL7     DataCenterCode = "BL7"
	DataCenterCodeBN1     DataCenterCode = "BN1"
	DataCenterCodeBN7     DataCenterCode = "BN7"
	DataCenterCodeBOM01   DataCenterCode = "BOM01"
	DataCenterCodeBY1     DataCenterCode = "BY1"
	DataCenterCodeBY2     DataCenterCode = "BY2"
	DataCenterCodeBY21    DataCenterCode = "BY21"
	DataCenterCodeBY24    DataCenterCode = "BY24"
	DataCenterCodeCBR20   DataCenterCode = "CBR20"
	DataCenterCodeCH1     DataCenterCode = "CH1"
	DataCenterCodeCPQ02   DataCenterCode = "CPQ02"
	DataCenterCodeCPQ20   DataCenterCode = "CPQ20"
	DataCenterCodeCWL20   DataCenterCode = "CWL20"
	DataCenterCodeCYS04   DataCenterCode = "CYS04"
	DataCenterCodeDSM05   DataCenterCode = "DSM05"
	DataCenterCodeDUB07   DataCenterCode = "DUB07"
	DataCenterCodeFRA22   DataCenterCode = "FRA22"
	DataCenterCodeHKG20   DataCenterCode = "HKG20"
	DataCenterCodeInvalid DataCenterCode = "Invalid"
	DataCenterCodeJNB21   DataCenterCode = "JNB21"
	DataCenterCodeJNB22   DataCenterCode = "JNB22"
	DataCenterCodeLON24   DataCenterCode = "LON24"
	DataCenterCodeMAA01   DataCenterCode = "MAA01"
	DataCenterCodeMEL23   DataCenterCode = "MEL23"
	DataCenterCodeMNZ21   DataCenterCode = "MNZ21"
	DataCenterCodeMWH01   DataCenterCode = "MWH01"
	DataCenterCodeORK70   DataCenterCode = "ORK70"
	DataCenterCodeOSA02   DataCenterCode = "OSA02"
	DataCenterCodeOSA20   DataCenterCode = "OSA20"
	DataCenterCodeOSA22   DataCenterCode = "OSA22"
	DataCenterCodePAR22   DataCenterCode = "PAR22"
	DataCenterCodePNQ01   DataCenterCode = "PNQ01"
	DataCenterCodePUS20   DataCenterCode = "PUS20"
	DataCenterCodeSEL20   DataCenterCode = "SEL20"
	DataCenterCodeSEL21   DataCenterCode = "SEL21"
	DataCenterCodeSG2     DataCenterCode = "SG2"
	DataCenterCodeSHA03   DataCenterCode = "SHA03"
	DataCenterCodeSIN20   DataCenterCode = "SIN20"
	DataCenterCodeSN5     DataCenterCode = "SN5"
	DataCenterCodeSN6     DataCenterCode = "SN6"
	DataCenterCodeSN8     DataCenterCode = "SN8"
	DataCenterCodeSSE90   DataCenterCode = "SSE90"
	DataCenterCodeSVG20   DataCenterCode = "SVG20"
	DataCenterCodeSYD03   DataCenterCode = "SYD03"
	DataCenterCodeSYD23   DataCenterCode = "SYD23"
	DataCenterCodeTYO01   DataCenterCode = "TYO01"
	DataCenterCodeTYO22   DataCenterCode = "TYO22"
	DataCenterCodeYQB20   DataCenterCode = "YQB20"
	DataCenterCodeYTO20   DataCenterCode = "YTO20"
	DataCenterCodeYTO21   DataCenterCode = "YTO21"
	DataCenterCodeZRH20   DataCenterCode = "ZRH20"
)

// PossibleDataCenterCodeValues returns the possible values for the DataCenterCode const type.
func PossibleDataCenterCodeValues() []DataCenterCode {
	return []DataCenterCode{
		DataCenterCodeAM2,
		DataCenterCodeAMS06,
		DataCenterCodeAMS20,
		DataCenterCodeAUH20,
		DataCenterCodeAdHoc,
		DataCenterCodeBJB,
		DataCenterCodeBJS20,
		DataCenterCodeBL20,
		DataCenterCodeBL7,
		DataCenterCodeBN1,
		DataCenterCodeBN7,
		DataCenterCodeBOM01,
		DataCenterCodeBY1,
		DataCenterCodeBY2,
		DataCenterCodeBY21,
		DataCenterCodeBY24,
		DataCenterCodeCBR20,
		DataCenterCodeCH1,
		DataCenterCodeCPQ02,
		DataCenterCodeCPQ20,
		DataCenterCodeCWL20,
		DataCenterCodeCYS04,
		DataCenterCodeDSM05,
		DataCenterCodeDUB07,
		DataCenterCodeFRA22,
		DataCenterCodeHKG20,
		DataCenterCodeInvalid,
		DataCenterCodeJNB21,
		DataCenterCodeJNB22,
		DataCenterCodeLON24,
		DataCenterCodeMAA01,
		DataCenterCodeMEL23,
		DataCenterCodeMNZ21,
		DataCenterCodeMWH01,
		DataCenterCodeORK70,
		DataCenterCodeOSA02,
		DataCenterCodeOSA20,
		DataCenterCodeOSA22,
		DataCenterCodePAR22,
		DataCenterCodePNQ01,
		DataCenterCodePUS20,
		DataCenterCodeSEL20,
		DataCenterCodeSEL21,
		DataCenterCodeSG2,
		DataCenterCodeSHA03,
		DataCenterCodeSIN20,
		DataCenterCodeSN5,
		DataCenterCodeSN6,
		DataCenterCodeSN8,
		DataCenterCodeSSE90,
		DataCenterCodeSVG20,
		DataCenterCodeSYD03,
		DataCenterCodeSYD23,
		DataCenterCodeTYO01,
		DataCenterCodeTYO22,
		DataCenterCodeYQB20,
		DataCenterCodeYTO20,
		DataCenterCodeYTO21,
		DataCenterCodeZRH20,
	}
}

// DatacenterAddressType - Data center address type
type DatacenterAddressType string

const (
	// DatacenterAddressTypeDatacenterAddressLocation - Data center address location.
	DatacenterAddressTypeDatacenterAddressLocation DatacenterAddressType = "DatacenterAddressLocation"
	// DatacenterAddressTypeDatacenterAddressInstruction - Data center address instruction.
	DatacenterAddressTypeDatacenterAddressInstruction DatacenterAddressType = "DatacenterAddressInstruction"
)

// PossibleDatacenterAddressTypeValues returns the possible values for the DatacenterAddressType const type.
func PossibleDatacenterAddressTypeValues() []DatacenterAddressType {
	return []DatacenterAddressType{
		DatacenterAddressTypeDatacenterAddressLocation,
		DatacenterAddressTypeDatacenterAddressInstruction,
	}
}

// DoubleEncryption - Defines secondary layer of software-based encryption enablement.
type DoubleEncryption string

const (
	// DoubleEncryptionEnabled - Software-based encryption is enabled.
	DoubleEncryptionEnabled DoubleEncryption = "Enabled"
	// DoubleEncryptionDisabled - Software-based encryption is disabled.
	DoubleEncryptionDisabled DoubleEncryption = "Disabled"
)

// PossibleDoubleEncryptionValues returns the possible values for the DoubleEncryption const type.
func PossibleDoubleEncryptionValues() []DoubleEncryption {
	return []DoubleEncryption{
		DoubleEncryptionEnabled,
		DoubleEncryptionDisabled,
	}
}

// FilterFileType - Type of the filter file.
type FilterFileType string

const (
	// FilterFileTypeAzureBlob - Filter file is of the type AzureBlob.
	FilterFileTypeAzureBlob FilterFileType = "AzureBlob"
	// FilterFileTypeAzureFile - Filter file is of the type AzureFiles.
	FilterFileTypeAzureFile FilterFileType = "AzureFile"
)

// PossibleFilterFileTypeValues returns the possible values for the FilterFileType const type.
func PossibleFilterFileTypeValues() []FilterFileType {
	return []FilterFileType{
		FilterFileTypeAzureBlob,
		FilterFileTypeAzureFile,
	}
}

// JobDeliveryType - Delivery type of Job.
type JobDeliveryType string

const (
	// JobDeliveryTypeNonScheduled - Non Scheduled job.
	JobDeliveryTypeNonScheduled JobDeliveryType = "NonScheduled"
	// JobDeliveryTypeScheduled - Scheduled job.
	JobDeliveryTypeScheduled JobDeliveryType = "Scheduled"
)

// PossibleJobDeliveryTypeValues returns the possible values for the JobDeliveryType const type.
func PossibleJobDeliveryTypeValues() []JobDeliveryType {
	return []JobDeliveryType{
		JobDeliveryTypeNonScheduled,
		JobDeliveryTypeScheduled,
	}
}

// KekType - Type of encryption key used for key encryption.
type KekType string

const (
	// KekTypeMicrosoftManaged - Key encryption key is managed by Microsoft.
	KekTypeMicrosoftManaged KekType = "MicrosoftManaged"
	// KekTypeCustomerManaged - Key encryption key is managed by the Customer.
	KekTypeCustomerManaged KekType = "CustomerManaged"
)

// PossibleKekTypeValues returns the possible values for the KekType const type.
func PossibleKekTypeValues() []KekType {
	return []KekType{
		KekTypeMicrosoftManaged,
		KekTypeCustomerManaged,
	}
}

// LogCollectionLevel - Level of the logs to be collected.
type LogCollectionLevel string

const (
	// LogCollectionLevelError - Only Errors will be collected in the logs.
	LogCollectionLevelError LogCollectionLevel = "Error"
	// LogCollectionLevelVerbose - Verbose logging (includes Errors, CRC, size information and others).
	LogCollectionLevelVerbose LogCollectionLevel = "Verbose"
)

// PossibleLogCollectionLevelValues returns the possible values for the LogCollectionLevel const type.
func PossibleLogCollectionLevelValues() []LogCollectionLevel {
	return []LogCollectionLevel{
		LogCollectionLevelError,
		LogCollectionLevelVerbose,
	}
}

// NotificationStageName - Name of the stage.
type NotificationStageName string

const (
	// NotificationStageNameAtAzureDC - Notification at device received at Azure datacenter stage.
	NotificationStageNameAtAzureDC NotificationStageName = "AtAzureDC"
	// NotificationStageNameCreated - Notification at job created stage.
	NotificationStageNameCreated NotificationStageName = "Created"
	// NotificationStageNameDataCopy - Notification at data copy started stage.
	NotificationStageNameDataCopy NotificationStageName = "DataCopy"
	// NotificationStageNameDelivered - Notification at device delivered stage.
	NotificationStageNameDelivered NotificationStageName = "Delivered"
	// NotificationStageNameDevicePrepared - Notification at device prepared stage.
	NotificationStageNameDevicePrepared NotificationStageName = "DevicePrepared"
	// NotificationStageNameDispatched - Notification at device dispatched stage.
	NotificationStageNameDispatched NotificationStageName = "Dispatched"
	// NotificationStageNamePickedUp - Notification at device picked up from user stage.
	NotificationStageNamePickedUp NotificationStageName = "PickedUp"
	// NotificationStageNameShippedToCustomer - Notification at shipped devices to customer stage.
	NotificationStageNameShippedToCustomer NotificationStageName = "ShippedToCustomer"
)

// PossibleNotificationStageNameValues returns the possible values for the NotificationStageName const type.
func PossibleNotificationStageNameValues() []NotificationStageName {
	return []NotificationStageName{
		NotificationStageNameAtAzureDC,
		NotificationStageNameCreated,
		NotificationStageNameDataCopy,
		NotificationStageNameDelivered,
		NotificationStageNameDevicePrepared,
		NotificationStageNameDispatched,
		NotificationStageNamePickedUp,
		NotificationStageNameShippedToCustomer,
	}
}

// OverallValidationStatus - Overall validation status.
type OverallValidationStatus string

const (
	// OverallValidationStatusAllValidToProceed - Every input request is valid.
	OverallValidationStatusAllValidToProceed OverallValidationStatus = "AllValidToProceed"
	// OverallValidationStatusInputsRevisitRequired - Some input requests are not valid.
	OverallValidationStatusInputsRevisitRequired OverallValidationStatus = "InputsRevisitRequired"
	// OverallValidationStatusCertainInputValidationsSkipped - Certain input validations skipped.
	OverallValidationStatusCertainInputValidationsSkipped OverallValidationStatus = "CertainInputValidationsSkipped"
)

// PossibleOverallValidationStatusValues returns the possible values for the OverallValidationStatus const type.
func PossibleOverallValidationStatusValues() []OverallValidationStatus {
	return []OverallValidationStatus{
		OverallValidationStatusAllValidToProceed,
		OverallValidationStatusInputsRevisitRequired,
		OverallValidationStatusCertainInputValidationsSkipped,
	}
}

// SKUDisabledReason - Reason why the Sku is disabled.
type SKUDisabledReason string

const (
	// SKUDisabledReasonNone - SKU is not disabled.
	SKUDisabledReasonNone SKUDisabledReason = "None"
	// SKUDisabledReasonCountry - SKU is not available in the requested country.
	SKUDisabledReasonCountry SKUDisabledReason = "Country"
	// SKUDisabledReasonRegion - SKU is not available to push data to the requested Azure region.
	SKUDisabledReasonRegion SKUDisabledReason = "Region"
	// SKUDisabledReasonFeature - Required features are not enabled for the SKU.
	SKUDisabledReasonFeature SKUDisabledReason = "Feature"
	// SKUDisabledReasonOfferType - Subscription does not have required offer types for the SKU.
	SKUDisabledReasonOfferType SKUDisabledReason = "OfferType"
	// SKUDisabledReasonNoSubscriptionInfo - Subscription has not registered to Microsoft.DataBox and Service does not have the
	// subscription notification.
	SKUDisabledReasonNoSubscriptionInfo SKUDisabledReason = "NoSubscriptionInfo"
)

// PossibleSKUDisabledReasonValues returns the possible values for the SKUDisabledReason const type.
func PossibleSKUDisabledReasonValues() []SKUDisabledReason {
	return []SKUDisabledReason{
		SKUDisabledReasonNone,
		SKUDisabledReasonCountry,
		SKUDisabledReasonRegion,
		SKUDisabledReasonFeature,
		SKUDisabledReasonOfferType,
		SKUDisabledReasonNoSubscriptionInfo,
	}
}

type SKUName string

const (
	// SKUNameDataBox - Data Box.
	SKUNameDataBox SKUName = "DataBox"
	// SKUNameDataBoxDisk - Data Box Disk.
	SKUNameDataBoxDisk SKUName = "DataBoxDisk"
	// SKUNameDataBoxHeavy - Data Box Heavy.
	SKUNameDataBoxHeavy SKUName = "DataBoxHeavy"
	// SKUNameDataBoxCustomerDisk - Data Box Customer Disk
	SKUNameDataBoxCustomerDisk SKUName = "DataBoxCustomerDisk"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameDataBox,
		SKUNameDataBoxDisk,
		SKUNameDataBoxHeavy,
		SKUNameDataBoxCustomerDisk,
	}
}

// ShareDestinationFormatType - Type of the share.
type ShareDestinationFormatType string

const (
	// ShareDestinationFormatTypeUnknownType - Unknown format.
	ShareDestinationFormatTypeUnknownType ShareDestinationFormatType = "UnknownType"
	// ShareDestinationFormatTypeHCS - Storsimple data format.
	ShareDestinationFormatTypeHCS ShareDestinationFormatType = "HCS"
	// ShareDestinationFormatTypeBlockBlob - Azure storage block blob format.
	ShareDestinationFormatTypeBlockBlob ShareDestinationFormatType = "BlockBlob"
	// ShareDestinationFormatTypePageBlob - Azure storage page blob format.
	ShareDestinationFormatTypePageBlob ShareDestinationFormatType = "PageBlob"
	// ShareDestinationFormatTypeAzureFile - Azure storage file format.
	ShareDestinationFormatTypeAzureFile ShareDestinationFormatType = "AzureFile"
	// ShareDestinationFormatTypeManagedDisk - Azure Compute Disk.
	ShareDestinationFormatTypeManagedDisk ShareDestinationFormatType = "ManagedDisk"
)

// PossibleShareDestinationFormatTypeValues returns the possible values for the ShareDestinationFormatType const type.
func PossibleShareDestinationFormatTypeValues() []ShareDestinationFormatType {
	return []ShareDestinationFormatType{
		ShareDestinationFormatTypeUnknownType,
		ShareDestinationFormatTypeHCS,
		ShareDestinationFormatTypeBlockBlob,
		ShareDestinationFormatTypePageBlob,
		ShareDestinationFormatTypeAzureFile,
		ShareDestinationFormatTypeManagedDisk,
	}
}

// StageName - Name of the stage which is in progress.
type StageName string

const (
	// StageNameAborted - Order has been aborted.
	StageNameAborted StageName = "Aborted"
	// StageNameAtAzureDC - Device has been received at Azure datacenter from the user.
	StageNameAtAzureDC StageName = "AtAzureDC"
	// StageNameAwaitingShipmentDetails - Awaiting shipment details of device from customer.
	StageNameAwaitingShipmentDetails StageName = "AwaitingShipmentDetails"
	// StageNameCancelled - Order has been cancelled.
	StageNameCancelled StageName = "Cancelled"
	// StageNameCompleted - Order has completed.
	StageNameCompleted StageName = "Completed"
	// StageNameCompletedWithErrors - Order has completed with errors.
	StageNameCompletedWithErrors StageName = "CompletedWithErrors"
	// StageNameCompletedWithWarnings - Order has completed with warnings.
	StageNameCompletedWithWarnings StageName = "CompletedWithWarnings"
	// StageNameCreated - Job created by the customer.
	StageNameCreated StageName = "Created"
	// StageNameDataCopy - Data copy from the device at Azure datacenter.
	StageNameDataCopy StageName = "DataCopy"
	// StageNameDelivered - Device has been delivered to the user of the order.
	StageNameDelivered StageName = "Delivered"
	// StageNameDeviceOrdered - An order has been created.
	StageNameDeviceOrdered StageName = "DeviceOrdered"
	// StageNameDevicePrepared - A device has been prepared for the order.
	StageNameDevicePrepared StageName = "DevicePrepared"
	// StageNameDispatched - Device has been dispatched to the user of the order.
	StageNameDispatched StageName = "Dispatched"
	// StageNameFailedIssueDetectedAtAzureDC - Order has failed due to issue detected at Azure datacenter.
	StageNameFailedIssueDetectedAtAzureDC StageName = "Failed_IssueDetectedAtAzureDC"
	// StageNameFailedIssueReportedAtCustomer - Order has failed due to issue reported by user.
	StageNameFailedIssueReportedAtCustomer StageName = "Failed_IssueReportedAtCustomer"
	// StageNamePickedUp - Device has been picked up from user and in transit to Azure datacenter.
	StageNamePickedUp StageName = "PickedUp"
	// StageNamePreparingToShipFromAzureDC - Preparing the device to ship to customer.
	StageNamePreparingToShipFromAzureDC StageName = "PreparingToShipFromAzureDC"
	// StageNameReadyToDispatchFromAzureDC - Device is ready to be handed to customer from Azure DC.
	StageNameReadyToDispatchFromAzureDC StageName = "ReadyToDispatchFromAzureDC"
	// StageNameReadyToReceiveAtAzureDC - Device can be dropped off at Azure DC.
	StageNameReadyToReceiveAtAzureDC StageName = "ReadyToReceiveAtAzureDC"
	// StageNameShippedToAzureDC - User shipped the device to AzureDC.
	StageNameShippedToAzureDC StageName = "ShippedToAzureDC"
	// StageNameShippedToCustomer - Shipped the device to customer.
	StageNameShippedToCustomer StageName = "ShippedToCustomer"
)

// PossibleStageNameValues returns the possible values for the StageName const type.
func PossibleStageNameValues() []StageName {
	return []StageName{
		StageNameAborted,
		StageNameAtAzureDC,
		StageNameAwaitingShipmentDetails,
		StageNameCancelled,
		StageNameCompleted,
		StageNameCompletedWithErrors,
		StageNameCompletedWithWarnings,
		StageNameCreated,
		StageNameDataCopy,
		StageNameDelivered,
		StageNameDeviceOrdered,
		StageNameDevicePrepared,
		StageNameDispatched,
		StageNameFailedIssueDetectedAtAzureDC,
		StageNameFailedIssueReportedAtCustomer,
		StageNamePickedUp,
		StageNamePreparingToShipFromAzureDC,
		StageNameReadyToDispatchFromAzureDC,
		StageNameReadyToReceiveAtAzureDC,
		StageNameShippedToAzureDC,
		StageNameShippedToCustomer,
	}
}

// StageStatus - Status of the job stage.
type StageStatus string

const (
	// StageStatusNone - No status available yet.
	StageStatusNone StageStatus = "None"
	// StageStatusInProgress - Stage is in progress.
	StageStatusInProgress StageStatus = "InProgress"
	// StageStatusSucceeded - Stage has succeeded.
	StageStatusSucceeded StageStatus = "Succeeded"
	// StageStatusFailed - Stage has failed.
	StageStatusFailed StageStatus = "Failed"
	// StageStatusCancelled - Stage has been cancelled.
	StageStatusCancelled StageStatus = "Cancelled"
	// StageStatusCancelling - Stage is cancelling.
	StageStatusCancelling StageStatus = "Cancelling"
	// StageStatusSucceededWithErrors - Stage has succeeded with errors.
	StageStatusSucceededWithErrors StageStatus = "SucceededWithErrors"
	// StageStatusWaitingForCustomerAction - Stage is stuck until customer takes some action.
	StageStatusWaitingForCustomerAction StageStatus = "WaitingForCustomerAction"
	// StageStatusSucceededWithWarnings - Stage has succeeded with warnings.
	StageStatusSucceededWithWarnings StageStatus = "SucceededWithWarnings"
	// StageStatusWaitingForCustomerActionForKek - Stage is waiting for customer action for kek action items.
	StageStatusWaitingForCustomerActionForKek StageStatus = "WaitingForCustomerActionForKek"
	// StageStatusWaitingForCustomerActionForCleanUp - Stage is waiting for customer action for clean up.
	StageStatusWaitingForCustomerActionForCleanUp StageStatus = "WaitingForCustomerActionForCleanUp"
	// StageStatusCustomerActionPerformedForCleanUp - Stage has performed customer action for clean up.
	StageStatusCustomerActionPerformedForCleanUp StageStatus = "CustomerActionPerformedForCleanUp"
	// StageStatusCustomerActionPerformed - Stage has performed customer action.
	StageStatusCustomerActionPerformed StageStatus = "CustomerActionPerformed"
)

// PossibleStageStatusValues returns the possible values for the StageStatus const type.
func PossibleStageStatusValues() []StageStatus {
	return []StageStatus{
		StageStatusNone,
		StageStatusInProgress,
		StageStatusSucceeded,
		StageStatusFailed,
		StageStatusCancelled,
		StageStatusCancelling,
		StageStatusSucceededWithErrors,
		StageStatusWaitingForCustomerAction,
		StageStatusSucceededWithWarnings,
		StageStatusWaitingForCustomerActionForKek,
		StageStatusWaitingForCustomerActionForCleanUp,
		StageStatusCustomerActionPerformedForCleanUp,
		StageStatusCustomerActionPerformed,
	}
}

// TransferConfigurationType - Type of the configuration for transfer.
type TransferConfigurationType string

const (
	// TransferConfigurationTypeTransferAll - Transfer all the data.
	TransferConfigurationTypeTransferAll TransferConfigurationType = "TransferAll"
	// TransferConfigurationTypeTransferUsingFilter - Transfer using filter.
	TransferConfigurationTypeTransferUsingFilter TransferConfigurationType = "TransferUsingFilter"
)

// PossibleTransferConfigurationTypeValues returns the possible values for the TransferConfigurationType const type.
func PossibleTransferConfigurationTypeValues() []TransferConfigurationType {
	return []TransferConfigurationType{
		TransferConfigurationTypeTransferAll,
		TransferConfigurationTypeTransferUsingFilter,
	}
}

// TransferType - Type of the transfer.
type TransferType string

const (
	// TransferTypeImportToAzure - Import data to azure.
	TransferTypeImportToAzure TransferType = "ImportToAzure"
	// TransferTypeExportFromAzure - Export data from azure.
	TransferTypeExportFromAzure TransferType = "ExportFromAzure"
)

// PossibleTransferTypeValues returns the possible values for the TransferType const type.
func PossibleTransferTypeValues() []TransferType {
	return []TransferType{
		TransferTypeImportToAzure,
		TransferTypeExportFromAzure,
	}
}

// TransportShipmentTypes - Transport Shipment Type supported for given region.
type TransportShipmentTypes string

const (
	// TransportShipmentTypesCustomerManaged - Shipment Logistics is handled by the customer.
	TransportShipmentTypesCustomerManaged TransportShipmentTypes = "CustomerManaged"
	// TransportShipmentTypesMicrosoftManaged - Shipment Logistics is handled by Microsoft.
	TransportShipmentTypesMicrosoftManaged TransportShipmentTypes = "MicrosoftManaged"
)

// PossibleTransportShipmentTypesValues returns the possible values for the TransportShipmentTypes const type.
func PossibleTransportShipmentTypesValues() []TransportShipmentTypes {
	return []TransportShipmentTypes{
		TransportShipmentTypesCustomerManaged,
		TransportShipmentTypesMicrosoftManaged,
	}
}

// ValidationInputDiscriminator - Identifies the type of validation request.
type ValidationInputDiscriminator string

const (
	// ValidationInputDiscriminatorValidateAddress - Identify request and response of address validation.
	ValidationInputDiscriminatorValidateAddress ValidationInputDiscriminator = "ValidateAddress"
	// ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob - Identify request and response for validation of
	// subscription permission to create job.
	ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob ValidationInputDiscriminator = "ValidateSubscriptionIsAllowedToCreateJob"
	// ValidationInputDiscriminatorValidatePreferences - Identify request and response of preference validation.
	ValidationInputDiscriminatorValidatePreferences ValidationInputDiscriminator = "ValidatePreferences"
	// ValidationInputDiscriminatorValidateCreateOrderLimit - Identify request and response of create order limit for subscription
	// validation.
	ValidationInputDiscriminatorValidateCreateOrderLimit ValidationInputDiscriminator = "ValidateCreateOrderLimit"
	// ValidationInputDiscriminatorValidateSKUAvailability - Identify request and response of active job limit for sku availability.
	ValidationInputDiscriminatorValidateSKUAvailability ValidationInputDiscriminator = "ValidateSkuAvailability"
	// ValidationInputDiscriminatorValidateDataTransferDetails - Identify request and response of data transfer details validation.
	ValidationInputDiscriminatorValidateDataTransferDetails ValidationInputDiscriminator = "ValidateDataTransferDetails"
)

// PossibleValidationInputDiscriminatorValues returns the possible values for the ValidationInputDiscriminator const type.
func PossibleValidationInputDiscriminatorValues() []ValidationInputDiscriminator {
	return []ValidationInputDiscriminator{
		ValidationInputDiscriminatorValidateAddress,
		ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob,
		ValidationInputDiscriminatorValidatePreferences,
		ValidationInputDiscriminatorValidateCreateOrderLimit,
		ValidationInputDiscriminatorValidateSKUAvailability,
		ValidationInputDiscriminatorValidateDataTransferDetails,
	}
}

// ValidationStatus - Create order limit validation status.
type ValidationStatus string

const (
	// ValidationStatusValid - Validation is successful
	ValidationStatusValid ValidationStatus = "Valid"
	// ValidationStatusInvalid - Validation is not successful
	ValidationStatusInvalid ValidationStatus = "Invalid"
	// ValidationStatusSkipped - Validation is skipped
	ValidationStatusSkipped ValidationStatus = "Skipped"
)

// PossibleValidationStatusValues returns the possible values for the ValidationStatus const type.
func PossibleValidationStatusValues() []ValidationStatus {
	return []ValidationStatus{
		ValidationStatusValid,
		ValidationStatusInvalid,
		ValidationStatusSkipped,
	}
}
