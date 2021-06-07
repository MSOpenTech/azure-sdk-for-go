package training

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/gofrs/uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.1/customvision/training"

// Bool ...
type Bool struct {
	autorest.Response `json:"-"`
	Value             *bool `json:"value,omitempty"`
}

// BoundingBox bounding box that defines a region of an image.
type BoundingBox struct {
	// Left - Coordinate of the left boundary.
	Left *float64 `json:"left,omitempty"`
	// Top - Coordinate of the top boundary.
	Top *float64 `json:"top,omitempty"`
	// Width - Width.
	Width *float64 `json:"width,omitempty"`
	// Height - Height.
	Height *float64 `json:"height,omitempty"`
}

// CustomVisionError ...
type CustomVisionError struct {
	// Code - The error code. Possible values include: 'NoError', 'BadRequest', 'BadRequestExceededBatchSize', 'BadRequestNotSupported', 'BadRequestInvalidIds', 'BadRequestProjectName', 'BadRequestProjectNameNotUnique', 'BadRequestProjectDescription', 'BadRequestProjectUnknownDomain', 'BadRequestProjectUnknownClassification', 'BadRequestProjectUnsupportedDomainTypeChange', 'BadRequestProjectUnsupportedExportPlatform', 'BadRequestProjectImagePreprocessingSettings', 'BadRequestIterationName', 'BadRequestIterationNameNotUnique', 'BadRequestIterationDescription', 'BadRequestIterationIsNotTrained', 'BadRequestWorkspaceCannotBeModified', 'BadRequestWorkspaceNotDeletable', 'BadRequestTagName', 'BadRequestTagNameNotUnique', 'BadRequestTagDescription', 'BadRequestTagType', 'BadRequestMultipleNegativeTag', 'BadRequestImageTags', 'BadRequestImageRegions', 'BadRequestNegativeAndRegularTagOnSameImage', 'BadRequestRequiredParamIsNull', 'BadRequestIterationIsPublished', 'BadRequestInvalidPublishName', 'BadRequestInvalidPublishTarget', 'BadRequestUnpublishFailed', 'BadRequestIterationNotPublished', 'BadRequestSubscriptionAPI', 'BadRequestExceedProjectLimit', 'BadRequestExceedIterationPerProjectLimit', 'BadRequestExceedTagPerProjectLimit', 'BadRequestExceedTagPerImageLimit', 'BadRequestExceededQuota', 'BadRequestCannotMigrateProjectWithName', 'BadRequestNotLimitedTrial', 'BadRequestImageBatch', 'BadRequestImageStream', 'BadRequestImageURL', 'BadRequestImageFormat', 'BadRequestImageSizeBytes', 'BadRequestImageExceededCount', 'BadRequestTrainingNotNeeded', 'BadRequestTrainingNotNeededButTrainingPipelineUpdated', 'BadRequestTrainingValidationFailed', 'BadRequestClassificationTrainingValidationFailed', 'BadRequestMultiClassClassificationTrainingValidationFailed', 'BadRequestMultiLabelClassificationTrainingValidationFailed', 'BadRequestDetectionTrainingValidationFailed', 'BadRequestTrainingAlreadyInProgress', 'BadRequestDetectionTrainingNotAllowNegativeTag', 'BadRequestInvalidEmailAddress', 'BadRequestDomainNotSupportedForAdvancedTraining', 'BadRequestExportPlatformNotSupportedForAdvancedTraining', 'BadRequestReservedBudgetInHoursNotEnoughForAdvancedTraining', 'BadRequestExportValidationFailed', 'BadRequestExportAlreadyInProgress', 'BadRequestPredictionIdsMissing', 'BadRequestPredictionIdsExceededCount', 'BadRequestPredictionTagsExceededCount', 'BadRequestPredictionResultsExceededCount', 'BadRequestPredictionInvalidApplicationName', 'BadRequestPredictionInvalidQueryParameters', 'BadRequestInvalid', 'UnsupportedMediaType', 'Forbidden', 'ForbiddenUser', 'ForbiddenUserResource', 'ForbiddenUserSignupDisabled', 'ForbiddenUserSignupAllowanceExceeded', 'ForbiddenUserDoesNotExist', 'ForbiddenUserDisabled', 'ForbiddenUserInsufficientCapability', 'ForbiddenDRModeEnabled', 'ForbiddenInvalid', 'NotFound', 'NotFoundProject', 'NotFoundProjectDefaultIteration', 'NotFoundIteration', 'NotFoundIterationPerformance', 'NotFoundTag', 'NotFoundImage', 'NotFoundDomain', 'NotFoundApimSubscription', 'NotFoundInvalid', 'Conflict', 'ConflictInvalid', 'ErrorUnknown', 'ErrorProjectInvalidWorkspace', 'ErrorProjectInvalidPipelineConfiguration', 'ErrorProjectInvalidDomain', 'ErrorProjectTrainingRequestFailed', 'ErrorProjectExportRequestFailed', 'ErrorFeaturizationServiceUnavailable', 'ErrorFeaturizationQueueTimeout', 'ErrorFeaturizationInvalidFeaturizer', 'ErrorFeaturizationAugmentationUnavailable', 'ErrorFeaturizationUnrecognizedJob', 'ErrorFeaturizationAugmentationError', 'ErrorExporterInvalidPlatform', 'ErrorExporterInvalidFeaturizer', 'ErrorExporterInvalidClassifier', 'ErrorPredictionServiceUnavailable', 'ErrorPredictionModelNotFound', 'ErrorPredictionModelNotCached', 'ErrorPrediction', 'ErrorPredictionStorage', 'ErrorRegionProposal', 'ErrorInvalid'
	Code CustomVisionErrorCodes `json:"code,omitempty"`
	// Message - A message explaining the error reported by the service.
	Message *string `json:"message,omitempty"`
}

// Domain ...
type Domain struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - READ-ONLY
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Possible values include: 'Classification', 'ObjectDetection'
	Type DomainType `json:"type,omitempty"`
	// Exportable - READ-ONLY
	Exportable *bool `json:"exportable,omitempty"`
	// Enabled - READ-ONLY
	Enabled *bool `json:"enabled,omitempty"`
}

// MarshalJSON is the custom marshaler for Domain.
func (d Domain) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Export ...
type Export struct {
	autorest.Response `json:"-"`
	// Platform - READ-ONLY; Platform of the export. Possible values include: 'CoreML', 'TensorFlow', 'DockerFile', 'ONNX', 'VAIDK'
	Platform ExportPlatform `json:"platform,omitempty"`
	// Status - READ-ONLY; Status of the export. Possible values include: 'Exporting', 'Failed', 'Done'
	Status ExportStatus `json:"status,omitempty"`
	// DownloadURI - READ-ONLY; URI used to download the model.
	DownloadURI *string `json:"downloadUri,omitempty"`
	// Flavor - READ-ONLY; Flavor of the export. Possible values include: 'Linux', 'Windows', 'ONNX10', 'ONNX12', 'ARM', 'TensorFlowNormal', 'TensorFlowLite'
	Flavor ExportFlavor `json:"flavor,omitempty"`
	// NewerVersionAvailable - READ-ONLY; Indicates an updated version of the export package is available and should be re-exported for the latest changes.
	NewerVersionAvailable *bool `json:"newerVersionAvailable,omitempty"`
}

// MarshalJSON is the custom marshaler for Export.
func (e Export) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Image image model to be sent as JSON.
type Image struct {
	// ID - READ-ONLY; Id of the image.
	ID *uuid.UUID `json:"id,omitempty"`
	// Created - READ-ONLY; Date the image was created.
	Created *date.Time `json:"created,omitempty"`
	// Width - READ-ONLY; Width of the image.
	Width *int32 `json:"width,omitempty"`
	// Height - READ-ONLY; Height of the image.
	Height *int32 `json:"height,omitempty"`
	// ResizedImageURI - READ-ONLY; The URI to the (resized) image used for training.
	ResizedImageURI *string `json:"resizedImageUri,omitempty"`
	// ThumbnailURI - READ-ONLY; The URI to the thumbnail of the original image.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// OriginalImageURI - READ-ONLY; The URI to the original uploaded image.
	OriginalImageURI *string `json:"originalImageUri,omitempty"`
	// Tags - READ-ONLY; Tags associated with this image.
	Tags *[]ImageTag `json:"tags,omitempty"`
	// Regions - READ-ONLY; Regions associated with this image.
	Regions *[]ImageRegion `json:"regions,omitempty"`
}

// MarshalJSON is the custom marshaler for Image.
func (i Image) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageCreateResult ...
type ImageCreateResult struct {
	// SourceURL - READ-ONLY; Source URL of the image.
	SourceURL *string `json:"sourceUrl,omitempty"`
	// Status - READ-ONLY; Status of the image creation. Possible values include: 'ImageCreateStatusOK', 'ImageCreateStatusOKDuplicate', 'ImageCreateStatusErrorSource', 'ImageCreateStatusErrorImageFormat', 'ImageCreateStatusErrorImageSize', 'ImageCreateStatusErrorStorage', 'ImageCreateStatusErrorLimitExceed', 'ImageCreateStatusErrorTagLimitExceed', 'ImageCreateStatusErrorRegionLimitExceed', 'ImageCreateStatusErrorUnknown', 'ImageCreateStatusErrorNegativeAndRegularTagOnSameImage'
	Status ImageCreateStatus `json:"status,omitempty"`
	// Image - READ-ONLY; The image.
	Image *Image `json:"image,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageCreateResult.
func (icr ImageCreateResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageCreateSummary ...
type ImageCreateSummary struct {
	autorest.Response `json:"-"`
	// IsBatchSuccessful - READ-ONLY; True if all of the images in the batch were created successfully, otherwise false.
	IsBatchSuccessful *bool `json:"isBatchSuccessful,omitempty"`
	// Images - READ-ONLY; List of the image creation results.
	Images *[]ImageCreateResult `json:"images,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageCreateSummary.
func (ics ImageCreateSummary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageFileCreateBatch ...
type ImageFileCreateBatch struct {
	Images *[]ImageFileCreateEntry `json:"images,omitempty"`
	TagIds *[]uuid.UUID            `json:"tagIds,omitempty"`
}

// ImageFileCreateEntry ...
type ImageFileCreateEntry struct {
	Name     *string      `json:"name,omitempty"`
	Contents *[]byte      `json:"contents,omitempty"`
	TagIds   *[]uuid.UUID `json:"tagIds,omitempty"`
	Regions  *[]Region    `json:"regions,omitempty"`
}

// ImageIDCreateBatch ...
type ImageIDCreateBatch struct {
	Images *[]ImageIDCreateEntry `json:"images,omitempty"`
	TagIds *[]uuid.UUID          `json:"tagIds,omitempty"`
}

// ImageIDCreateEntry ...
type ImageIDCreateEntry struct {
	// ID - Id of the image.
	ID      *uuid.UUID   `json:"id,omitempty"`
	TagIds  *[]uuid.UUID `json:"tagIds,omitempty"`
	Regions *[]Region    `json:"regions,omitempty"`
}

// ImagePerformance image performance model.
type ImagePerformance struct {
	// Predictions - READ-ONLY
	Predictions *[]Prediction `json:"predictions,omitempty"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// Width - READ-ONLY
	Width *int32 `json:"width,omitempty"`
	// Height - READ-ONLY
	Height *int32 `json:"height,omitempty"`
	// ImageURI - READ-ONLY
	ImageURI *string `json:"imageUri,omitempty"`
	// ThumbnailURI - READ-ONLY
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// Tags - READ-ONLY
	Tags *[]ImageTag `json:"tags,omitempty"`
	// Regions - READ-ONLY
	Regions *[]ImageRegion `json:"regions,omitempty"`
}

// MarshalJSON is the custom marshaler for ImagePerformance.
func (IP ImagePerformance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImagePrediction result of an image prediction request.
type ImagePrediction struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Prediction Id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY; Project Id.
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY; Iteration Id.
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY; Date this prediction was created.
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY; List of predictions.
	Predictions *[]Prediction `json:"predictions,omitempty"`
}

// MarshalJSON is the custom marshaler for ImagePrediction.
func (IP ImagePrediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageProcessingSettings represents image preprocessing settings used by image augmentation.
type ImageProcessingSettings struct {
	// AugmentationMethods - Gets or sets enabled image transforms. The key corresponds to the transform name. If value is set to true, then correspondent transform is enabled. Otherwise this transform will not be used.
	// Augmentation will be uniformly distributed among enabled transforms.
	AugmentationMethods map[string]*bool `json:"augmentationMethods"`
}

// MarshalJSON is the custom marshaler for ImageProcessingSettings.
func (ips ImageProcessingSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ips.AugmentationMethods != nil {
		objectMap["augmentationMethods"] = ips.AugmentationMethods
	}
	return json.Marshal(objectMap)
}

// ImageRegion ...
type ImageRegion struct {
	// RegionID - READ-ONLY
	RegionID *uuid.UUID `json:"regionId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// TagID - Id of the tag associated with this region.
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// Left - Coordinate of the left boundary.
	Left *float64 `json:"left,omitempty"`
	// Top - Coordinate of the top boundary.
	Top *float64 `json:"top,omitempty"`
	// Width - Width.
	Width *float64 `json:"width,omitempty"`
	// Height - Height.
	Height *float64 `json:"height,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageRegion.
func (ir ImageRegion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ir.TagID != nil {
		objectMap["tagId"] = ir.TagID
	}
	if ir.Left != nil {
		objectMap["left"] = ir.Left
	}
	if ir.Top != nil {
		objectMap["top"] = ir.Top
	}
	if ir.Width != nil {
		objectMap["width"] = ir.Width
	}
	if ir.Height != nil {
		objectMap["height"] = ir.Height
	}
	return json.Marshal(objectMap)
}

// ImageRegionCreateBatch batch of image region information to create.
type ImageRegionCreateBatch struct {
	Regions *[]ImageRegionCreateEntry `json:"regions,omitempty"`
}

// ImageRegionCreateEntry entry associating a region to an image.
type ImageRegionCreateEntry struct {
	// ImageID - Id of the image.
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// TagID - Id of the tag associated with this region.
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// Left - Coordinate of the left boundary.
	Left *float64 `json:"left,omitempty"`
	// Top - Coordinate of the top boundary.
	Top *float64 `json:"top,omitempty"`
	// Width - Width.
	Width *float64 `json:"width,omitempty"`
	// Height - Height.
	Height *float64 `json:"height,omitempty"`
}

// ImageRegionCreateResult ...
type ImageRegionCreateResult struct {
	// ImageID - READ-ONLY
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// RegionID - READ-ONLY
	RegionID *uuid.UUID `json:"regionId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// TagID - Id of the tag associated with this region.
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// Left - Coordinate of the left boundary.
	Left *float64 `json:"left,omitempty"`
	// Top - Coordinate of the top boundary.
	Top *float64 `json:"top,omitempty"`
	// Width - Width.
	Width *float64 `json:"width,omitempty"`
	// Height - Height.
	Height *float64 `json:"height,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageRegionCreateResult.
func (ircr ImageRegionCreateResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ircr.TagID != nil {
		objectMap["tagId"] = ircr.TagID
	}
	if ircr.Left != nil {
		objectMap["left"] = ircr.Left
	}
	if ircr.Top != nil {
		objectMap["top"] = ircr.Top
	}
	if ircr.Width != nil {
		objectMap["width"] = ircr.Width
	}
	if ircr.Height != nil {
		objectMap["height"] = ircr.Height
	}
	return json.Marshal(objectMap)
}

// ImageRegionCreateSummary ...
type ImageRegionCreateSummary struct {
	autorest.Response `json:"-"`
	Created           *[]ImageRegionCreateResult `json:"created,omitempty"`
	Duplicated        *[]ImageRegionCreateEntry  `json:"duplicated,omitempty"`
	Exceeded          *[]ImageRegionCreateEntry  `json:"exceeded,omitempty"`
}

// ImageRegionProposal ...
type ImageRegionProposal struct {
	autorest.Response `json:"-"`
	// ProjectID - READ-ONLY
	ProjectID *uuid.UUID `json:"projectId,omitempty"`
	// ImageID - READ-ONLY
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// Proposals - READ-ONLY
	Proposals *[]RegionProposal `json:"proposals,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageRegionProposal.
func (irp ImageRegionProposal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageTag ...
type ImageTag struct {
	// TagID - READ-ONLY
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageTag.
func (it ImageTag) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageTagCreateBatch batch of image tags.
type ImageTagCreateBatch struct {
	// Tags - Image Tag entries to include in this batch.
	Tags *[]ImageTagCreateEntry `json:"tags,omitempty"`
}

// ImageTagCreateEntry entry associating a tag to an image.
type ImageTagCreateEntry struct {
	// ImageID - Id of the image.
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// TagID - Id of the tag.
	TagID *uuid.UUID `json:"tagId,omitempty"`
}

// ImageTagCreateSummary ...
type ImageTagCreateSummary struct {
	autorest.Response `json:"-"`
	Created           *[]ImageTagCreateEntry `json:"created,omitempty"`
	Duplicated        *[]ImageTagCreateEntry `json:"duplicated,omitempty"`
	Exceeded          *[]ImageTagCreateEntry `json:"exceeded,omitempty"`
}

// ImageURL image url.
type ImageURL struct {
	// URL - Url of the image.
	URL *string `json:"url,omitempty"`
}

// ImageURLCreateBatch ...
type ImageURLCreateBatch struct {
	Images *[]ImageURLCreateEntry `json:"images,omitempty"`
	TagIds *[]uuid.UUID           `json:"tagIds,omitempty"`
}

// ImageURLCreateEntry ...
type ImageURLCreateEntry struct {
	// URL - Url of the image.
	URL     *string      `json:"url,omitempty"`
	TagIds  *[]uuid.UUID `json:"tagIds,omitempty"`
	Regions *[]Region    `json:"regions,omitempty"`
}

// Int32 ...
type Int32 struct {
	autorest.Response `json:"-"`
	Value             *int32 `json:"value,omitempty"`
}

// Iteration iteration model to be sent over JSON.
type Iteration struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the id of the iteration.
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - Gets or sets the name of the iteration.
	Name *string `json:"name,omitempty"`
	// Status - READ-ONLY; Gets the current iteration status.
	Status *string `json:"status,omitempty"`
	// Created - READ-ONLY; Gets the time this iteration was completed.
	Created *date.Time `json:"created,omitempty"`
	// LastModified - READ-ONLY; Gets the time this iteration was last modified.
	LastModified *date.Time `json:"lastModified,omitempty"`
	// TrainedAt - READ-ONLY; Gets the time this iteration was last modified.
	TrainedAt *date.Time `json:"trainedAt,omitempty"`
	// ProjectID - READ-ONLY; Gets the project id of the iteration.
	ProjectID *uuid.UUID `json:"projectId,omitempty"`
	// Exportable - READ-ONLY; Whether the iteration can be exported to another format for download.
	Exportable *bool `json:"exportable,omitempty"`
	// ExportableTo - READ-ONLY; A set of platforms this iteration can export to.
	ExportableTo *[]string `json:"exportableTo,omitempty"`
	// DomainID - READ-ONLY; Get or sets a guid of the domain the iteration has been trained on.
	DomainID *uuid.UUID `json:"domainId,omitempty"`
	// ClassificationType - READ-ONLY; Gets the classification type of the project. Possible values include: 'Multiclass', 'Multilabel'
	ClassificationType Classifier `json:"classificationType,omitempty"`
	// TrainingType - READ-ONLY; Gets the training type of the iteration. Possible values include: 'TypeRegular', 'TypeAdvanced'
	TrainingType Type `json:"trainingType,omitempty"`
	// ReservedBudgetInHours - READ-ONLY; Gets the reserved advanced training budget for the iteration.
	ReservedBudgetInHours *int32 `json:"reservedBudgetInHours,omitempty"`
	// TrainingTimeInMinutes - READ-ONLY; Gets the training time for the iteration.
	TrainingTimeInMinutes *int32 `json:"trainingTimeInMinutes,omitempty"`
	// PublishName - READ-ONLY; Name of the published model.
	PublishName *string `json:"publishName,omitempty"`
	// OriginalPublishResourceID - READ-ONLY; Resource Provider Id this iteration was originally published to.
	OriginalPublishResourceID *string `json:"originalPublishResourceId,omitempty"`
}

// MarshalJSON is the custom marshaler for Iteration.
func (i Iteration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.Name != nil {
		objectMap["name"] = i.Name
	}
	return json.Marshal(objectMap)
}

// IterationPerformance represents the detailed performance data for a trained iteration.
type IterationPerformance struct {
	autorest.Response `json:"-"`
	// PerTagPerformance - READ-ONLY; Gets the per-tag performance details for this iteration.
	PerTagPerformance *[]TagPerformance `json:"perTagPerformance,omitempty"`
	// Precision - READ-ONLY; Gets the precision.
	Precision *float64 `json:"precision,omitempty"`
	// PrecisionStdDeviation - READ-ONLY; Gets the standard deviation for the precision.
	PrecisionStdDeviation *float64 `json:"precisionStdDeviation,omitempty"`
	// Recall - READ-ONLY; Gets the recall.
	Recall *float64 `json:"recall,omitempty"`
	// RecallStdDeviation - READ-ONLY; Gets the standard deviation for the recall.
	RecallStdDeviation *float64 `json:"recallStdDeviation,omitempty"`
	// AveragePrecision - READ-ONLY; Gets the average precision when applicable.
	AveragePrecision *float64 `json:"averagePrecision,omitempty"`
}

// MarshalJSON is the custom marshaler for IterationPerformance.
func (IP IterationPerformance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ListDomain ...
type ListDomain struct {
	autorest.Response `json:"-"`
	Value             *[]Domain `json:"value,omitempty"`
}

// ListExport ...
type ListExport struct {
	autorest.Response `json:"-"`
	Value             *[]Export `json:"value,omitempty"`
}

// ListImage ...
type ListImage struct {
	autorest.Response `json:"-"`
	Value             *[]Image `json:"value,omitempty"`
}

// ListImagePerformance ...
type ListImagePerformance struct {
	autorest.Response `json:"-"`
	Value             *[]ImagePerformance `json:"value,omitempty"`
}

// ListIteration ...
type ListIteration struct {
	autorest.Response `json:"-"`
	Value             *[]Iteration `json:"value,omitempty"`
}

// ListProject ...
type ListProject struct {
	autorest.Response `json:"-"`
	Value             *[]Project `json:"value,omitempty"`
}

// ListSuggestedTagAndRegion ...
type ListSuggestedTagAndRegion struct {
	autorest.Response `json:"-"`
	Value             *[]SuggestedTagAndRegion `json:"value,omitempty"`
}

// ListTag ...
type ListTag struct {
	autorest.Response `json:"-"`
	Value             *[]Tag `json:"value,omitempty"`
}

// Prediction prediction result.
type Prediction struct {
	// Probability - READ-ONLY; Probability of the tag.
	Probability *float64 `json:"probability,omitempty"`
	// TagID - READ-ONLY; Id of the predicted tag.
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// TagName - READ-ONLY; Name of the predicted tag.
	TagName *string `json:"tagName,omitempty"`
	// BoundingBox - READ-ONLY; Bounding box of the prediction.
	BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
}

// MarshalJSON is the custom marshaler for Prediction.
func (p Prediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// PredictionQueryResult query result of the prediction images that were sent to your prediction endpoint.
type PredictionQueryResult struct {
	autorest.Response `json:"-"`
	// Token - Prediction Query Token.
	Token *PredictionQueryToken `json:"token,omitempty"`
	// Results - READ-ONLY; Result of an prediction request.
	Results *[]StoredImagePrediction `json:"results,omitempty"`
}

// MarshalJSON is the custom marshaler for PredictionQueryResult.
func (pqr PredictionQueryResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pqr.Token != nil {
		objectMap["token"] = pqr.Token
	}
	return json.Marshal(objectMap)
}

// PredictionQueryTag ...
type PredictionQueryTag struct {
	ID           *uuid.UUID `json:"id,omitempty"`
	MinThreshold *float64   `json:"minThreshold,omitempty"`
	MaxThreshold *float64   `json:"maxThreshold,omitempty"`
}

// PredictionQueryToken ...
type PredictionQueryToken struct {
	Session      *string `json:"session,omitempty"`
	Continuation *string `json:"continuation,omitempty"`
	MaxCount     *int32  `json:"maxCount,omitempty"`
	// OrderBy - Possible values include: 'Newest', 'Oldest', 'Suggested'
	OrderBy     OrderBy               `json:"orderBy,omitempty"`
	Tags        *[]PredictionQueryTag `json:"tags,omitempty"`
	IterationID *uuid.UUID            `json:"iterationId,omitempty"`
	StartTime   *date.Time            `json:"startTime,omitempty"`
	EndTime     *date.Time            `json:"endTime,omitempty"`
	Application *string               `json:"application,omitempty"`
}

// Project represents a project.
type Project struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the project id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - Gets or sets the name of the project.
	Name *string `json:"name,omitempty"`
	// Description - Gets or sets the description of the project.
	Description *string `json:"description,omitempty"`
	// Settings - Gets or sets the project settings.
	Settings *ProjectSettings `json:"settings,omitempty"`
	// Created - READ-ONLY; Gets the date this project was created.
	Created *date.Time `json:"created,omitempty"`
	// LastModified - READ-ONLY; Gets the date this project was last modified.
	LastModified *date.Time `json:"lastModified,omitempty"`
	// ThumbnailURI - READ-ONLY; Gets the thumbnail url representing the image.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// DrModeEnabled - READ-ONLY; Gets if the Disaster Recovery (DR) mode is on, indicating the project is temporarily read-only.
	DrModeEnabled *bool `json:"drModeEnabled,omitempty"`
}

// MarshalJSON is the custom marshaler for Project.
func (p Project) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Description != nil {
		objectMap["description"] = p.Description
	}
	if p.Settings != nil {
		objectMap["settings"] = p.Settings
	}
	return json.Marshal(objectMap)
}

// ProjectSettings represents settings associated with a project.
type ProjectSettings struct {
	// DomainID - Gets or sets the id of the Domain to use with this project.
	DomainID *uuid.UUID `json:"domainId,omitempty"`
	// ClassificationType - Gets or sets the classification type of the project. Possible values include: 'Multiclass', 'Multilabel'
	ClassificationType Classifier `json:"classificationType,omitempty"`
	// TargetExportPlatforms - A list of ExportPlatform that the trained model should be able to support.
	TargetExportPlatforms *[]string `json:"targetExportPlatforms,omitempty"`
	// UseNegativeSet - READ-ONLY; Indicates if negative set is being used.
	UseNegativeSet *bool `json:"useNegativeSet,omitempty"`
	// DetectionParameters - READ-ONLY; Detection parameters in use, if any.
	DetectionParameters *string `json:"detectionParameters,omitempty"`
	// ImageProcessingSettings - Gets or sets image preprocessing settings.
	ImageProcessingSettings *ImageProcessingSettings `json:"imageProcessingSettings,omitempty"`
}

// MarshalJSON is the custom marshaler for ProjectSettings.
func (ps ProjectSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ps.DomainID != nil {
		objectMap["domainId"] = ps.DomainID
	}
	if ps.ClassificationType != "" {
		objectMap["classificationType"] = ps.ClassificationType
	}
	if ps.TargetExportPlatforms != nil {
		objectMap["targetExportPlatforms"] = ps.TargetExportPlatforms
	}
	if ps.ImageProcessingSettings != nil {
		objectMap["imageProcessingSettings"] = ps.ImageProcessingSettings
	}
	return json.Marshal(objectMap)
}

// Region ...
type Region struct {
	// TagID - Id of the tag associated with this region.
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// Left - Coordinate of the left boundary.
	Left *float64 `json:"left,omitempty"`
	// Top - Coordinate of the top boundary.
	Top *float64 `json:"top,omitempty"`
	// Width - Width.
	Width *float64 `json:"width,omitempty"`
	// Height - Height.
	Height *float64 `json:"height,omitempty"`
}

// RegionProposal ...
type RegionProposal struct {
	// Confidence - READ-ONLY
	Confidence *float64 `json:"confidence,omitempty"`
	// BoundingBox - READ-ONLY
	BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
}

// MarshalJSON is the custom marshaler for RegionProposal.
func (rp RegionProposal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SetInt32 ...
type SetInt32 struct {
	autorest.Response `json:"-"`
	Value             map[string]*int32 `json:"value"`
}

// MarshalJSON is the custom marshaler for SetInt32.
func (si3 SetInt32) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if si3.Value != nil {
		objectMap["value"] = si3.Value
	}
	return json.Marshal(objectMap)
}

// StoredImagePrediction result of an image prediction request.
type StoredImagePrediction struct {
	// ResizedImageURI - READ-ONLY; The URI to the (resized) prediction image.
	ResizedImageURI *string `json:"resizedImageUri,omitempty"`
	// ThumbnailURI - READ-ONLY; The URI to the thumbnail of the original prediction image.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// OriginalImageURI - READ-ONLY; The URI to the original prediction image.
	OriginalImageURI *string `json:"originalImageUri,omitempty"`
	// Domain - READ-ONLY; Domain used for the prediction.
	Domain *uuid.UUID `json:"domain,omitempty"`
	// ID - READ-ONLY; Prediction Id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY; Project Id.
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY; Iteration Id.
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY; Date this prediction was created.
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY; List of predictions.
	Predictions *[]Prediction `json:"predictions,omitempty"`
}

// MarshalJSON is the custom marshaler for StoredImagePrediction.
func (sip StoredImagePrediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// StoredSuggestedTagAndRegion result of a suggested tags and regions request of the untagged image.
type StoredSuggestedTagAndRegion struct {
	// Width - READ-ONLY; Width of the resized image.
	Width *int32 `json:"width,omitempty"`
	// Height - READ-ONLY; Height of the resized image.
	Height *int32 `json:"height,omitempty"`
	// ResizedImageURI - READ-ONLY; The URI to the (resized) prediction image.
	ResizedImageURI *string `json:"resizedImageUri,omitempty"`
	// ThumbnailURI - READ-ONLY; The URI to the thumbnail of the original prediction image.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// OriginalImageURI - READ-ONLY; The URI to the original prediction image.
	OriginalImageURI *string `json:"originalImageUri,omitempty"`
	// Domain - READ-ONLY; Domain used for the prediction.
	Domain *uuid.UUID `json:"domain,omitempty"`
	// ID - READ-ONLY; Prediction Id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY; Project Id.
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY; Iteration Id.
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY; Date this prediction was created.
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY; List of predictions.
	Predictions *[]Prediction `json:"predictions,omitempty"`
	// PredictionUncertainty - READ-ONLY; Uncertainty (entropy) of suggested tags or regions per image.
	PredictionUncertainty *float64 `json:"predictionUncertainty,omitempty"`
}

// MarshalJSON is the custom marshaler for StoredSuggestedTagAndRegion.
func (sstar StoredSuggestedTagAndRegion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SuggestedTagAndRegion result of a suggested tags and regions request.
type SuggestedTagAndRegion struct {
	// ID - READ-ONLY; Prediction Id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY; Project Id.
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY; Iteration Id.
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY; Date this prediction was created.
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY; List of predictions.
	Predictions *[]Prediction `json:"predictions,omitempty"`
	// PredictionUncertainty - READ-ONLY; Uncertainty (entropy) of suggested tags or regions per image.
	PredictionUncertainty *float64 `json:"predictionUncertainty,omitempty"`
}

// MarshalJSON is the custom marshaler for SuggestedTagAndRegion.
func (star SuggestedTagAndRegion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// SuggestedTagAndRegionQuery the array of result images and token containing session and continuation Ids
// for the next query.
type SuggestedTagAndRegionQuery struct {
	autorest.Response `json:"-"`
	// Token - Contains properties we need to fetch suggested tags for.
	Token *SuggestedTagAndRegionQueryToken `json:"token,omitempty"`
	// Results - READ-ONLY; Result of a suggested tags and regions request of the untagged image.
	Results *[]StoredSuggestedTagAndRegion `json:"results,omitempty"`
}

// MarshalJSON is the custom marshaler for SuggestedTagAndRegionQuery.
func (starq SuggestedTagAndRegionQuery) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if starq.Token != nil {
		objectMap["token"] = starq.Token
	}
	return json.Marshal(objectMap)
}

// SuggestedTagAndRegionQueryToken contains properties we need to fetch suggested tags for. For the first
// call, Session and continuation set to null.
// Then on subsequent calls, uses the session/continuation from the previous SuggestedTagAndRegionQuery
// result to fetch additional results.
type SuggestedTagAndRegionQueryToken struct {
	// TagIds - Existing TagIds in project to filter suggested tags on.
	TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
	// Threshold - Confidence threshold to filter suggested tags on.
	Threshold *float64 `json:"threshold,omitempty"`
	// Session - SessionId for database query. Initially set to null but later used to paginate.
	Session *string `json:"session,omitempty"`
	// Continuation - Continuation Id for database pagination. Initially null but later used to paginate.
	Continuation *string `json:"continuation,omitempty"`
	// MaxCount - Maximum number of results you want to be returned in the response.
	MaxCount *int32 `json:"maxCount,omitempty"`
	// SortBy - OrderBy. Ordering mechanism for your results. Possible values include: 'UncertaintyAscending', 'UncertaintyDescending'
	SortBy SortBy `json:"sortBy,omitempty"`
}

// Tag represents a Tag.
type Tag struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the Tag ID.
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - Gets or sets the name of the tag.
	Name *string `json:"name,omitempty"`
	// Description - Gets or sets the description of the tag.
	Description *string `json:"description,omitempty"`
	// Type - Gets or sets the type of the tag. Possible values include: 'Regular', 'Negative'
	Type TagType `json:"type,omitempty"`
	// ImageCount - READ-ONLY; Gets the number of images with this tag.
	ImageCount *int32 `json:"imageCount,omitempty"`
}

// MarshalJSON is the custom marshaler for Tag.
func (t Tag) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if t.Name != nil {
		objectMap["name"] = t.Name
	}
	if t.Description != nil {
		objectMap["description"] = t.Description
	}
	if t.Type != "" {
		objectMap["type"] = t.Type
	}
	return json.Marshal(objectMap)
}

// TagFilter model that query for counting of images whose suggested tags match given tags and their
// probability are greater than or equal to the given threshold.
type TagFilter struct {
	// TagIds - Existing TagIds in project to get suggested tags count for.
	TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
	// Threshold - Confidence threshold to filter suggested tags on.
	Threshold *float64 `json:"threshold,omitempty"`
}

// TagPerformance represents performance data for a particular tag in a trained iteration.
type TagPerformance struct {
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - READ-ONLY
	Name *string `json:"name,omitempty"`
	// Precision - READ-ONLY; Gets the precision.
	Precision *float64 `json:"precision,omitempty"`
	// PrecisionStdDeviation - READ-ONLY; Gets the standard deviation for the precision.
	PrecisionStdDeviation *float64 `json:"precisionStdDeviation,omitempty"`
	// Recall - READ-ONLY; Gets the recall.
	Recall *float64 `json:"recall,omitempty"`
	// RecallStdDeviation - READ-ONLY; Gets the standard deviation for the recall.
	RecallStdDeviation *float64 `json:"recallStdDeviation,omitempty"`
	// AveragePrecision - READ-ONLY; Gets the average precision when applicable.
	AveragePrecision *float64 `json:"averagePrecision,omitempty"`
}

// MarshalJSON is the custom marshaler for TagPerformance.
func (tp TagPerformance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
