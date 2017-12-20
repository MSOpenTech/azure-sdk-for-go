// +build go1.9

// Copyright 2017 Microsoft Corporation
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
// commit ID: 2014fbbf031942474ad27a5a66dffaed5347f3fb

package face

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/face"

type AttributeTypes = original.AttributeTypes

const (
	AttributeTypesAccessories AttributeTypes = original.AttributeTypesAccessories
	AttributeTypesAge         AttributeTypes = original.AttributeTypesAge
	AttributeTypesBlur        AttributeTypes = original.AttributeTypesBlur
	AttributeTypesEmotion     AttributeTypes = original.AttributeTypesEmotion
	AttributeTypesExposure    AttributeTypes = original.AttributeTypesExposure
	AttributeTypesFacialHair  AttributeTypes = original.AttributeTypesFacialHair
	AttributeTypesGender      AttributeTypes = original.AttributeTypesGender
	AttributeTypesGlasses     AttributeTypes = original.AttributeTypesGlasses
	AttributeTypesHair        AttributeTypes = original.AttributeTypesHair
	AttributeTypesHeadPose    AttributeTypes = original.AttributeTypesHeadPose
	AttributeTypesMakeup      AttributeTypes = original.AttributeTypesMakeup
	AttributeTypesNoise       AttributeTypes = original.AttributeTypesNoise
	AttributeTypesOcclusion   AttributeTypes = original.AttributeTypesOcclusion
	AttributeTypesSmile       AttributeTypes = original.AttributeTypesSmile
)

type AzureRegions = original.AzureRegions

const (
	Australiaeast  AzureRegions = original.Australiaeast
	Brazilsouth    AzureRegions = original.Brazilsouth
	Eastasia       AzureRegions = original.Eastasia
	Eastus         AzureRegions = original.Eastus
	Eastus2        AzureRegions = original.Eastus2
	Northeurope    AzureRegions = original.Northeurope
	Southcentralus AzureRegions = original.Southcentralus
	Southeastasia  AzureRegions = original.Southeastasia
	Westcentralus  AzureRegions = original.Westcentralus
	Westeurope     AzureRegions = original.Westeurope
	Westus         AzureRegions = original.Westus
	Westus2        AzureRegions = original.Westus2
)

type BlurLevels = original.BlurLevels

const (
	High   BlurLevels = original.High
	Low    BlurLevels = original.Low
	Medium BlurLevels = original.Medium
)

type ExposureLevels = original.ExposureLevels

const (
	GoodExposure  ExposureLevels = original.GoodExposure
	OverExposure  ExposureLevels = original.OverExposure
	UnderExposure ExposureLevels = original.UnderExposure
)

type Gender = original.Gender

const (
	Female     Gender = original.Female
	Genderless Gender = original.Genderless
	Male       Gender = original.Male
)

type GlassesTypes = original.GlassesTypes

const (
	NoGlasses       GlassesTypes = original.NoGlasses
	ReadingGlasses  GlassesTypes = original.ReadingGlasses
	Sunglasses      GlassesTypes = original.Sunglasses
	SwimmingGoggles GlassesTypes = original.SwimmingGoggles
)

type MatchingMode = original.MatchingMode

const (
	MatchFace   MatchingMode = original.MatchFace
	MatchPerson MatchingMode = original.MatchPerson
)

type NoiseLevels = original.NoiseLevels

const (
	NoiseLevelsHigh   NoiseLevels = original.NoiseLevelsHigh
	NoiseLevelsLow    NoiseLevels = original.NoiseLevelsLow
	NoiseLevelsMedium NoiseLevels = original.NoiseLevelsMedium
)

type TrainingStatusType = original.TrainingStatusType

const (
	Failed     TrainingStatusType = original.Failed
	Nonstarted TrainingStatusType = original.Nonstarted
	Running    TrainingStatusType = original.Running
	Succeeded  TrainingStatusType = original.Succeeded
)

type AccessoryItem = original.AccessoryItem
type APIError = original.APIError
type Attributes = original.Attributes
type BlurProperties = original.BlurProperties
type ColorProperty = original.ColorProperty
type CreateFaceListRequest = original.CreateFaceListRequest
type CreatePersonGroupRequest = original.CreatePersonGroupRequest
type CreatePersonRequest = original.CreatePersonRequest
type CreatePersonResult = original.CreatePersonResult
type DetectedFace = original.DetectedFace
type EmotionProperties = original.EmotionProperties
type Error = original.Error
type ExposureProperties = original.ExposureProperties
type FacialHairProperties = original.FacialHairProperties
type FindSimilarRequest = original.FindSimilarRequest
type GetFaceListResult = original.GetFaceListResult
type GroupRequest = original.GroupRequest
type GroupResponse = original.GroupResponse
type HairProperties = original.HairProperties
type HeadPoseProperties = original.HeadPoseProperties
type IdentifyRequest = original.IdentifyRequest
type IdentifyResultCandidate = original.IdentifyResultCandidate
type IdentifyResultItem = original.IdentifyResultItem
type ImageURL = original.ImageURL
type Landmarks = original.Landmarks
type ListDetectedFace = original.ListDetectedFace
type ListGetFaceListResult = original.ListGetFaceListResult
type ListIdentifyResultItem = original.ListIdentifyResultItem
type ListPersonGroupResult = original.ListPersonGroupResult
type ListPersonResult = original.ListPersonResult
type ListSimilarFaceResult = original.ListSimilarFaceResult
type MakeupProperties = original.MakeupProperties
type NoiseProperties = original.NoiseProperties
type OcclusionProperties = original.OcclusionProperties
type PersistedFaceResult = original.PersistedFaceResult
type PersonFaceResult = original.PersonFaceResult
type PersonGroupResult = original.PersonGroupResult
type PersonResult = original.PersonResult
type Position = original.Position
type Rectangle = original.Rectangle
type SimilarFaceResult = original.SimilarFaceResult
type TrainingStatus = original.TrainingStatus
type UpdatePersonFaceDataRequest = original.UpdatePersonFaceDataRequest
type VerifyPersonGroupRequest = original.VerifyPersonGroupRequest
type VerifyRequest = original.VerifyRequest
type VerifyResult = original.VerifyResult
type PersonClient = original.PersonClient
type PersonGroupClient = original.PersonGroupClient
type BaseClient = original.BaseClient
type Client = original.Client
type ListClient = original.ListClient

func NewClient(azureRegion AzureRegions) Client {
	return original.NewClient(azureRegion)
}
func NewListClient(azureRegion AzureRegions) ListClient {
	return original.NewListClient(azureRegion)
}
func NewPersonClient(azureRegion AzureRegions) PersonClient {
	return original.NewPersonClient(azureRegion)
}
func NewPersonGroupClient(azureRegion AzureRegions) PersonGroupClient {
	return original.NewPersonGroupClient(azureRegion)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func New(azureRegion AzureRegions) BaseClient {
	return original.New(azureRegion)
}
func NewWithoutDefaults(azureRegion AzureRegions) BaseClient {
	return original.NewWithoutDefaults(azureRegion)
}
