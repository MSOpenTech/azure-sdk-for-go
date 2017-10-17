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
// commit ID: 714052a3359963ba46703fe033cf9dd4838bea11

package cognitiveservices

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2017-04-18/cognitiveservices"

type AccountsClient = original.AccountsClient
type CheckSkuAvailabilityClient = original.CheckSkuAvailabilityClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type KeyName = original.KeyName

const (
	Key1 KeyName = original.Key1
	Key2 KeyName = original.Key2
)

type Kind = original.Kind

const (
	Academic           Kind = original.Academic
	BingAutosuggest    Kind = original.BingAutosuggest
	BingSearch         Kind = original.BingSearch
	BingSpeech         Kind = original.BingSpeech
	BingSpellCheck     Kind = original.BingSpellCheck
	ComputerVision     Kind = original.ComputerVision
	ContentModerator   Kind = original.ContentModerator
	CustomSpeech       Kind = original.CustomSpeech
	Emotion            Kind = original.Emotion
	Face               Kind = original.Face
	LUIS               Kind = original.LUIS
	Recommendations    Kind = original.Recommendations
	SpeakerRecognition Kind = original.SpeakerRecognition
	Speech             Kind = original.Speech
	SpeechTranslation  Kind = original.SpeechTranslation
	TextAnalytics      Kind = original.TextAnalytics
	TextTranslation    Kind = original.TextTranslation
	WebLM              Kind = original.WebLM
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	Failed       ProvisioningState = original.Failed
	ResolvingDNS ProvisioningState = original.ResolvingDNS
	Succeeded    ProvisioningState = original.Succeeded
)

type SkuName = original.SkuName

const (
	F0 SkuName = original.F0
	P0 SkuName = original.P0
	P1 SkuName = original.P1
	P2 SkuName = original.P2
	S0 SkuName = original.S0
	S1 SkuName = original.S1
	S2 SkuName = original.S2
	S3 SkuName = original.S3
	S4 SkuName = original.S4
	S5 SkuName = original.S5
	S6 SkuName = original.S6
)

type SkuTier = original.SkuTier

const (
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type Account = original.Account
type AccountCreateParameters = original.AccountCreateParameters
type AccountEnumerateSkusResult = original.AccountEnumerateSkusResult
type AccountKeys = original.AccountKeys
type AccountListResult = original.AccountListResult
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type CheckSkuAvailabilityParameter = original.CheckSkuAvailabilityParameter
type CheckSkuAvailabilityResult = original.CheckSkuAvailabilityResult
type CheckSkuAvailabilityResultList = original.CheckSkuAvailabilityResultList
type Error = original.Error
type ErrorBody = original.ErrorBody
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type RegenerateKeyParameters = original.RegenerateKeyParameters
type ResourceAndSku = original.ResourceAndSku
type Sku = original.Sku
type OperationsClient = original.OperationsClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCheckSkuAvailabilityClient(subscriptionID string) CheckSkuAvailabilityClient {
	return original.NewCheckSkuAvailabilityClient(subscriptionID)
}
func NewCheckSkuAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckSkuAvailabilityClient {
	return original.NewCheckSkuAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}
