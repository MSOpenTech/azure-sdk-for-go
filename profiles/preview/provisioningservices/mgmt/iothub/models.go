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

package iothub

import original "github.com/Azure/azure-sdk-for-go/services/provisioningservices/mgmt/2018-01-22/iothub"

type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type IotDpsResourceClient = original.IotDpsResourceClient

func NewIotDpsResourceClient(subscriptionID string) IotDpsResourceClient {
	return original.NewIotDpsResourceClient(subscriptionID)
}
func NewIotDpsResourceClientWithBaseURI(baseURI string, subscriptionID string) IotDpsResourceClient {
	return original.NewIotDpsResourceClientWithBaseURI(baseURI, subscriptionID)
}

type AccessRightsDescription = original.AccessRightsDescription

const (
	DeviceConnect           AccessRightsDescription = original.DeviceConnect
	EnrollmentRead          AccessRightsDescription = original.EnrollmentRead
	EnrollmentWrite         AccessRightsDescription = original.EnrollmentWrite
	RegistrationStatusRead  AccessRightsDescription = original.RegistrationStatusRead
	RegistrationStatusWrite AccessRightsDescription = original.RegistrationStatusWrite
	ServiceConfig           AccessRightsDescription = original.ServiceConfig
)

func PossibleAccessRightsDescriptionValues() []AccessRightsDescription {
	return original.PossibleAccessRightsDescriptionValues()
}

type AllocationPolicy = original.AllocationPolicy

const (
	GeoLatency AllocationPolicy = original.GeoLatency
	Hashed     AllocationPolicy = original.Hashed
	Static     AllocationPolicy = original.Static
)

func PossibleAllocationPolicyValues() []AllocationPolicy {
	return original.PossibleAllocationPolicyValues()
}

type CertificatePurpose = original.CertificatePurpose

const (
	ClientAuthentication CertificatePurpose = original.ClientAuthentication
	ServerAuthentication CertificatePurpose = original.ServerAuthentication
)

func PossibleCertificatePurposeValues() []CertificatePurpose {
	return original.PossibleCertificatePurposeValues()
}

type IotDpsSku = original.IotDpsSku

const (
	S1 IotDpsSku = original.S1
)

func PossibleIotDpsSkuValues() []IotDpsSku {
	return original.PossibleIotDpsSkuValues()
}

type NameUnavailabilityReason = original.NameUnavailabilityReason

const (
	AlreadyExists NameUnavailabilityReason = original.AlreadyExists
	Invalid       NameUnavailabilityReason = original.Invalid
)

func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return original.PossibleNameUnavailabilityReasonValues()
}

type State = original.State

const (
	Activating       State = original.Activating
	ActivationFailed State = original.ActivationFailed
	Active           State = original.Active
	Deleted          State = original.Deleted
	Deleting         State = original.Deleting
	DeletionFailed   State = original.DeletionFailed
	FailingOver      State = original.FailingOver
	FailoverFailed   State = original.FailoverFailed
	Resuming         State = original.Resuming
	Suspended        State = original.Suspended
	Suspending       State = original.Suspending
	Transitioning    State = original.Transitioning
)

func PossibleStateValues() []State {
	return original.PossibleStateValues()
}

type AsyncOperationResult = original.AsyncOperationResult
type CertificateBodyDescription = original.CertificateBodyDescription
type CertificateListDescription = original.CertificateListDescription
type CertificateProperties = original.CertificateProperties
type CertificateResponse = original.CertificateResponse
type DefinitionDescription = original.DefinitionDescription
type ErrorDetails = original.ErrorDetails
type ErrorMesssage = original.ErrorMesssage
type IotDpsPropertiesDescription = original.IotDpsPropertiesDescription
type IotDpsResourceCreateOrUpdateFuture = original.IotDpsResourceCreateOrUpdateFuture
type IotDpsResourceDeleteFuture = original.IotDpsResourceDeleteFuture
type IotDpsResourceUpdateFuture = original.IotDpsResourceUpdateFuture
type IotDpsSkuDefinition = original.IotDpsSkuDefinition
type IotDpsSkuDefinitionListResult = original.IotDpsSkuDefinitionListResult
type IotDpsSkuDefinitionListResultIterator = original.IotDpsSkuDefinitionListResultIterator
type IotDpsSkuDefinitionListResultPage = original.IotDpsSkuDefinitionListResultPage
type IotDpsSkuInfo = original.IotDpsSkuInfo
type NameAvailabilityInfo = original.NameAvailabilityInfo
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type ProvisioningServiceDescription = original.ProvisioningServiceDescription
type ProvisioningServiceDescriptionListResult = original.ProvisioningServiceDescriptionListResult
type ProvisioningServiceDescriptionListResultIterator = original.ProvisioningServiceDescriptionListResultIterator
type ProvisioningServiceDescriptionListResultPage = original.ProvisioningServiceDescriptionListResultPage
type Resource = original.Resource
type SharedAccessSignatureAuthorizationRuleAccessRightsDescription = original.SharedAccessSignatureAuthorizationRuleAccessRightsDescription
type SharedAccessSignatureAuthorizationRuleListResult = original.SharedAccessSignatureAuthorizationRuleListResult
type SharedAccessSignatureAuthorizationRuleListResultIterator = original.SharedAccessSignatureAuthorizationRuleListResultIterator
type SharedAccessSignatureAuthorizationRuleListResultPage = original.SharedAccessSignatureAuthorizationRuleListResultPage
type TagsResource = original.TagsResource
type VerificationCodeRequest = original.VerificationCodeRequest
type VerificationCodeResponse = original.VerificationCodeResponse
type VerificationCodeResponseProperties = original.VerificationCodeResponseProperties
type DpsCertificateClient = original.DpsCertificateClient

func NewDpsCertificateClient(subscriptionID string) DpsCertificateClient {
	return original.NewDpsCertificateClient(subscriptionID)
}
func NewDpsCertificateClientWithBaseURI(baseURI string, subscriptionID string) DpsCertificateClient {
	return original.NewDpsCertificateClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
