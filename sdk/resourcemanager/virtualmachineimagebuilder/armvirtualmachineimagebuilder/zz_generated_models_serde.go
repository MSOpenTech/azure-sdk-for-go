//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplate.
func (i ImageTemplate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", i.ID)
	populate(objectMap, "identity", i.Identity)
	populate(objectMap, "location", i.Location)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "tags", i.Tags)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateCustomizer.
func (i *ImageTemplateCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer { return i }

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateDistributor.
func (i *ImageTemplateDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor { return i }

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateDistributor.
func (i ImageTemplateDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "runOutputName", i.RunOutputName)
	objectMap["type"] = i.Type
	return json.Marshal(objectMap)
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateFileCustomizer.
func (i *ImageTemplateFileCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Type: i.Type,
		Name: i.Name,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateFileCustomizer.
func (i ImageTemplateFileCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "destination", i.Destination)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "sourceUri", i.SourceURI)
	objectMap["type"] = "File"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateFileCustomizer.
func (i *ImageTemplateFileCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "destination":
			err = unpopulate(val, &i.Destination)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &i.Name)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, &i.SHA256Checksum)
			delete(rawMsg, key)
		case "sourceUri":
			err = unpopulate(val, &i.SourceURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateIdentity.
func (i ImageTemplateIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateLastRunStatus.
func (i ImageTemplateLastRunStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", i.EndTime)
	populate(objectMap, "message", i.Message)
	populate(objectMap, "runState", i.RunState)
	populate(objectMap, "runSubState", i.RunSubState)
	populateTimeRFC3339(objectMap, "startTime", i.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateLastRunStatus.
func (i *ImageTemplateLastRunStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, &i.EndTime)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &i.Message)
			delete(rawMsg, key)
		case "runState":
			err = unpopulate(val, &i.RunState)
			delete(rawMsg, key)
		case "runSubState":
			err = unpopulate(val, &i.RunSubState)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &i.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateListResult.
func (i ImageTemplateListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateManagedImageDistributor.
func (i *ImageTemplateManagedImageDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor {
	return &ImageTemplateDistributor{
		Type:          i.Type,
		RunOutputName: i.RunOutputName,
		ArtifactTags:  i.ArtifactTags,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateManagedImageDistributor.
func (i ImageTemplateManagedImageDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "imageId", i.ImageID)
	populate(objectMap, "location", i.Location)
	populate(objectMap, "runOutputName", i.RunOutputName)
	objectMap["type"] = "ManagedImage"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateManagedImageDistributor.
func (i *ImageTemplateManagedImageDistributor) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactTags":
			err = unpopulate(val, &i.ArtifactTags)
			delete(rawMsg, key)
		case "imageId":
			err = unpopulate(val, &i.ImageID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &i.Location)
			delete(rawMsg, key)
		case "runOutputName":
			err = unpopulate(val, &i.RunOutputName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplateManagedImageSource.
func (i *ImageTemplateManagedImageSource) GetImageTemplateSource() *ImageTemplateSource {
	return &ImageTemplateSource{
		Type: i.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateManagedImageSource.
func (i ImageTemplateManagedImageSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "imageId", i.ImageID)
	objectMap["type"] = "ManagedImage"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateManagedImageSource.
func (i *ImageTemplateManagedImageSource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "imageId":
			err = unpopulate(val, &i.ImageID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplatePlatformImageSource.
func (i *ImageTemplatePlatformImageSource) GetImageTemplateSource() *ImageTemplateSource {
	return &ImageTemplateSource{
		Type: i.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplatePlatformImageSource.
func (i ImageTemplatePlatformImageSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "exactVersion", i.ExactVersion)
	populate(objectMap, "offer", i.Offer)
	populate(objectMap, "planInfo", i.PlanInfo)
	populate(objectMap, "publisher", i.Publisher)
	populate(objectMap, "sku", i.SKU)
	objectMap["type"] = "PlatformImage"
	populate(objectMap, "version", i.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplatePlatformImageSource.
func (i *ImageTemplatePlatformImageSource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "exactVersion":
			err = unpopulate(val, &i.ExactVersion)
			delete(rawMsg, key)
		case "offer":
			err = unpopulate(val, &i.Offer)
			delete(rawMsg, key)
		case "planInfo":
			err = unpopulate(val, &i.PlanInfo)
			delete(rawMsg, key)
		case "publisher":
			err = unpopulate(val, &i.Publisher)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, &i.SKU)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &i.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplatePowerShellCustomizer.
func (i *ImageTemplatePowerShellCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Type: i.Type,
		Name: i.Name,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplatePowerShellCustomizer.
func (i ImageTemplatePowerShellCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inline", i.Inline)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "runAsSystem", i.RunAsSystem)
	populate(objectMap, "runElevated", i.RunElevated)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "scriptUri", i.ScriptURI)
	objectMap["type"] = "PowerShell"
	populate(objectMap, "validExitCodes", i.ValidExitCodes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplatePowerShellCustomizer.
func (i *ImageTemplatePowerShellCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inline":
			err = unpopulate(val, &i.Inline)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &i.Name)
			delete(rawMsg, key)
		case "runAsSystem":
			err = unpopulate(val, &i.RunAsSystem)
			delete(rawMsg, key)
		case "runElevated":
			err = unpopulate(val, &i.RunElevated)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, &i.SHA256Checksum)
			delete(rawMsg, key)
		case "scriptUri":
			err = unpopulate(val, &i.ScriptURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		case "validExitCodes":
			err = unpopulate(val, &i.ValidExitCodes)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateProperties.
func (i ImageTemplateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "buildTimeoutInMinutes", i.BuildTimeoutInMinutes)
	populate(objectMap, "customize", i.Customize)
	populate(objectMap, "distribute", i.Distribute)
	populate(objectMap, "lastRunStatus", i.LastRunStatus)
	populate(objectMap, "provisioningError", i.ProvisioningError)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "source", i.Source)
	populate(objectMap, "vmProfile", i.VMProfile)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateProperties.
func (i *ImageTemplateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "buildTimeoutInMinutes":
			err = unpopulate(val, &i.BuildTimeoutInMinutes)
			delete(rawMsg, key)
		case "customize":
			i.Customize, err = unmarshalImageTemplateCustomizerClassificationArray(val)
			delete(rawMsg, key)
		case "distribute":
			i.Distribute, err = unmarshalImageTemplateDistributorClassificationArray(val)
			delete(rawMsg, key)
		case "lastRunStatus":
			err = unpopulate(val, &i.LastRunStatus)
			delete(rawMsg, key)
		case "provisioningError":
			err = unpopulate(val, &i.ProvisioningError)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &i.ProvisioningState)
			delete(rawMsg, key)
		case "source":
			i.Source, err = unmarshalImageTemplateSourceClassification(val)
			delete(rawMsg, key)
		case "vmProfile":
			err = unpopulate(val, &i.VMProfile)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateRestartCustomizer.
func (i *ImageTemplateRestartCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Type: i.Type,
		Name: i.Name,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateRestartCustomizer.
func (i ImageTemplateRestartCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", i.Name)
	populate(objectMap, "restartCheckCommand", i.RestartCheckCommand)
	populate(objectMap, "restartCommand", i.RestartCommand)
	populate(objectMap, "restartTimeout", i.RestartTimeout)
	objectMap["type"] = "WindowsRestart"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateRestartCustomizer.
func (i *ImageTemplateRestartCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, &i.Name)
			delete(rawMsg, key)
		case "restartCheckCommand":
			err = unpopulate(val, &i.RestartCheckCommand)
			delete(rawMsg, key)
		case "restartCommand":
			err = unpopulate(val, &i.RestartCommand)
			delete(rawMsg, key)
		case "restartTimeout":
			err = unpopulate(val, &i.RestartTimeout)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateSharedImageDistributor.
func (i *ImageTemplateSharedImageDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor {
	return &ImageTemplateDistributor{
		Type:          i.Type,
		RunOutputName: i.RunOutputName,
		ArtifactTags:  i.ArtifactTags,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateSharedImageDistributor.
func (i ImageTemplateSharedImageDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "excludeFromLatest", i.ExcludeFromLatest)
	populate(objectMap, "galleryImageId", i.GalleryImageID)
	populate(objectMap, "replicationRegions", i.ReplicationRegions)
	populate(objectMap, "runOutputName", i.RunOutputName)
	populate(objectMap, "storageAccountType", i.StorageAccountType)
	objectMap["type"] = "SharedImage"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateSharedImageDistributor.
func (i *ImageTemplateSharedImageDistributor) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactTags":
			err = unpopulate(val, &i.ArtifactTags)
			delete(rawMsg, key)
		case "excludeFromLatest":
			err = unpopulate(val, &i.ExcludeFromLatest)
			delete(rawMsg, key)
		case "galleryImageId":
			err = unpopulate(val, &i.GalleryImageID)
			delete(rawMsg, key)
		case "replicationRegions":
			err = unpopulate(val, &i.ReplicationRegions)
			delete(rawMsg, key)
		case "runOutputName":
			err = unpopulate(val, &i.RunOutputName)
			delete(rawMsg, key)
		case "storageAccountType":
			err = unpopulate(val, &i.StorageAccountType)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplateSharedImageVersionSource.
func (i *ImageTemplateSharedImageVersionSource) GetImageTemplateSource() *ImageTemplateSource {
	return &ImageTemplateSource{
		Type: i.Type,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateSharedImageVersionSource.
func (i ImageTemplateSharedImageVersionSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "imageVersionId", i.ImageVersionID)
	objectMap["type"] = "SharedImageVersion"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateSharedImageVersionSource.
func (i *ImageTemplateSharedImageVersionSource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "imageVersionId":
			err = unpopulate(val, &i.ImageVersionID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateShellCustomizer.
func (i *ImageTemplateShellCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Type: i.Type,
		Name: i.Name,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateShellCustomizer.
func (i ImageTemplateShellCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inline", i.Inline)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "scriptUri", i.ScriptURI)
	objectMap["type"] = "Shell"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateShellCustomizer.
func (i *ImageTemplateShellCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inline":
			err = unpopulate(val, &i.Inline)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &i.Name)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, &i.SHA256Checksum)
			delete(rawMsg, key)
		case "scriptUri":
			err = unpopulate(val, &i.ScriptURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplateSource.
func (i *ImageTemplateSource) GetImageTemplateSource() *ImageTemplateSource { return i }

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateUpdateParameters.
func (i ImageTemplateUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", i.Identity)
	populate(objectMap, "tags", i.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateVMProfile.
func (i ImageTemplateVMProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "osDiskSizeGB", i.OSDiskSizeGB)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	populate(objectMap, "vmSize", i.VMSize)
	populate(objectMap, "vnetConfig", i.VnetConfig)
	return json.Marshal(objectMap)
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateVhdDistributor.
func (i *ImageTemplateVhdDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor {
	return &ImageTemplateDistributor{
		Type:          i.Type,
		RunOutputName: i.RunOutputName,
		ArtifactTags:  i.ArtifactTags,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateVhdDistributor.
func (i ImageTemplateVhdDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "runOutputName", i.RunOutputName)
	objectMap["type"] = "VHD"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateVhdDistributor.
func (i *ImageTemplateVhdDistributor) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactTags":
			err = unpopulate(val, &i.ArtifactTags)
			delete(rawMsg, key)
		case "runOutputName":
			err = unpopulate(val, &i.RunOutputName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateWindowsUpdateCustomizer.
func (i *ImageTemplateWindowsUpdateCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Type: i.Type,
		Name: i.Name,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateWindowsUpdateCustomizer.
func (i ImageTemplateWindowsUpdateCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "filters", i.Filters)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "searchCriteria", i.SearchCriteria)
	objectMap["type"] = "WindowsUpdate"
	populate(objectMap, "updateLimit", i.UpdateLimit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateWindowsUpdateCustomizer.
func (i *ImageTemplateWindowsUpdateCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "filters":
			err = unpopulate(val, &i.Filters)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &i.Name)
			delete(rawMsg, key)
		case "searchCriteria":
			err = unpopulate(val, &i.SearchCriteria)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &i.Type)
			delete(rawMsg, key)
		case "updateLimit":
			err = unpopulate(val, &i.UpdateLimit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RunOutputCollection.
func (r RunOutputCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
