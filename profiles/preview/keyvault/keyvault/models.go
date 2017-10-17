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

package keyvault

import original "github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"

type ManagementClient = original.ManagementClient
type ActionType = original.ActionType

const (
	AutoRenew     ActionType = original.AutoRenew
	EmailContacts ActionType = original.EmailContacts
)

type DeletionRecoveryLevel = original.DeletionRecoveryLevel

const (
	Purgeable                        DeletionRecoveryLevel = original.Purgeable
	Recoverable                      DeletionRecoveryLevel = original.Recoverable
	RecoverableProtectedSubscription DeletionRecoveryLevel = original.RecoverableProtectedSubscription
	RecoverablePurgeable             DeletionRecoveryLevel = original.RecoverablePurgeable
)

type JSONWebKeyCurveName = original.JSONWebKeyCurveName

const (
	P256      JSONWebKeyCurveName = original.P256
	P384      JSONWebKeyCurveName = original.P384
	P521      JSONWebKeyCurveName = original.P521
	SECP256K1 JSONWebKeyCurveName = original.SECP256K1
)

type JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithm

const (
	RSA15      JSONWebKeyEncryptionAlgorithm = original.RSA15
	RSAOAEP    JSONWebKeyEncryptionAlgorithm = original.RSAOAEP
	RSAOAEP256 JSONWebKeyEncryptionAlgorithm = original.RSAOAEP256
)

type JSONWebKeyOperation = original.JSONWebKeyOperation

const (
	Decrypt   JSONWebKeyOperation = original.Decrypt
	Encrypt   JSONWebKeyOperation = original.Encrypt
	Sign      JSONWebKeyOperation = original.Sign
	UnwrapKey JSONWebKeyOperation = original.UnwrapKey
	Verify    JSONWebKeyOperation = original.Verify
	WrapKey   JSONWebKeyOperation = original.WrapKey
)

type JSONWebKeySignatureAlgorithm = original.JSONWebKeySignatureAlgorithm

const (
	ECDSA256 JSONWebKeySignatureAlgorithm = original.ECDSA256
	ES256    JSONWebKeySignatureAlgorithm = original.ES256
	ES384    JSONWebKeySignatureAlgorithm = original.ES384
	ES512    JSONWebKeySignatureAlgorithm = original.ES512
	PS256    JSONWebKeySignatureAlgorithm = original.PS256
	PS384    JSONWebKeySignatureAlgorithm = original.PS384
	PS512    JSONWebKeySignatureAlgorithm = original.PS512
	RS256    JSONWebKeySignatureAlgorithm = original.RS256
	RS384    JSONWebKeySignatureAlgorithm = original.RS384
	RS512    JSONWebKeySignatureAlgorithm = original.RS512
	RSNULL   JSONWebKeySignatureAlgorithm = original.RSNULL
)

type JSONWebKeyType = original.JSONWebKeyType

const (
	EC     JSONWebKeyType = original.EC
	ECHSM  JSONWebKeyType = original.ECHSM
	Oct    JSONWebKeyType = original.Oct
	RSA    JSONWebKeyType = original.RSA
	RSAHSM JSONWebKeyType = original.RSAHSM
)

type KeyUsageType = original.KeyUsageType

const (
	CRLSign          KeyUsageType = original.CRLSign
	DataEncipherment KeyUsageType = original.DataEncipherment
	DecipherOnly     KeyUsageType = original.DecipherOnly
	DigitalSignature KeyUsageType = original.DigitalSignature
	EncipherOnly     KeyUsageType = original.EncipherOnly
	KeyAgreement     KeyUsageType = original.KeyAgreement
	KeyCertSign      KeyUsageType = original.KeyCertSign
	KeyEncipherment  KeyUsageType = original.KeyEncipherment
	NonRepudiation   KeyUsageType = original.NonRepudiation
)

type Action = original.Action
type AdministratorDetails = original.AdministratorDetails
type Attributes = original.Attributes
type BackupKeyResult = original.BackupKeyResult
type BackupSecretResult = original.BackupSecretResult
type CertificateAttributes = original.CertificateAttributes
type CertificateBundle = original.CertificateBundle
type CertificateCreateParameters = original.CertificateCreateParameters
type CertificateImportParameters = original.CertificateImportParameters
type CertificateIssuerItem = original.CertificateIssuerItem
type CertificateIssuerListResult = original.CertificateIssuerListResult
type CertificateIssuerSetParameters = original.CertificateIssuerSetParameters
type CertificateIssuerUpdateParameters = original.CertificateIssuerUpdateParameters
type CertificateItem = original.CertificateItem
type CertificateListResult = original.CertificateListResult
type CertificateMergeParameters = original.CertificateMergeParameters
type CertificateOperation = original.CertificateOperation
type CertificateOperationUpdateParameter = original.CertificateOperationUpdateParameter
type CertificatePolicy = original.CertificatePolicy
type CertificateUpdateParameters = original.CertificateUpdateParameters
type Contact = original.Contact
type Contacts = original.Contacts
type DeletedCertificateBundle = original.DeletedCertificateBundle
type DeletedCertificateItem = original.DeletedCertificateItem
type DeletedCertificateListResult = original.DeletedCertificateListResult
type DeletedKeyBundle = original.DeletedKeyBundle
type DeletedKeyItem = original.DeletedKeyItem
type DeletedKeyListResult = original.DeletedKeyListResult
type DeletedSecretBundle = original.DeletedSecretBundle
type DeletedSecretItem = original.DeletedSecretItem
type DeletedSecretListResult = original.DeletedSecretListResult
type Error = original.Error
type ErrorType = original.ErrorType
type IssuerAttributes = original.IssuerAttributes
type IssuerBundle = original.IssuerBundle
type IssuerCredentials = original.IssuerCredentials
type IssuerParameters = original.IssuerParameters
type JSONWebKey = original.JSONWebKey
type KeyAttributes = original.KeyAttributes
type KeyBundle = original.KeyBundle
type KeyCreateParameters = original.KeyCreateParameters
type KeyImportParameters = original.KeyImportParameters
type KeyItem = original.KeyItem
type KeyListResult = original.KeyListResult
type KeyOperationResult = original.KeyOperationResult
type KeyOperationsParameters = original.KeyOperationsParameters
type KeyProperties = original.KeyProperties
type KeyRestoreParameters = original.KeyRestoreParameters
type KeySignParameters = original.KeySignParameters
type KeyUpdateParameters = original.KeyUpdateParameters
type KeyVerifyParameters = original.KeyVerifyParameters
type KeyVerifyResult = original.KeyVerifyResult
type LifetimeAction = original.LifetimeAction
type OrganizationDetails = original.OrganizationDetails
type PendingCertificateSigningRequestResult = original.PendingCertificateSigningRequestResult
type SasDefinitionAttributes = original.SasDefinitionAttributes
type SasDefinitionBundle = original.SasDefinitionBundle
type SasDefinitionCreateParameters = original.SasDefinitionCreateParameters
type SasDefinitionItem = original.SasDefinitionItem
type SasDefinitionListResult = original.SasDefinitionListResult
type SasDefinitionUpdateParameters = original.SasDefinitionUpdateParameters
type SecretAttributes = original.SecretAttributes
type SecretBundle = original.SecretBundle
type SecretItem = original.SecretItem
type SecretListResult = original.SecretListResult
type SecretProperties = original.SecretProperties
type SecretRestoreParameters = original.SecretRestoreParameters
type SecretSetParameters = original.SecretSetParameters
type SecretUpdateParameters = original.SecretUpdateParameters
type StorageAccountAttributes = original.StorageAccountAttributes
type StorageAccountCreateParameters = original.StorageAccountCreateParameters
type StorageAccountItem = original.StorageAccountItem
type StorageAccountRegenerteKeyParameters = original.StorageAccountRegenerteKeyParameters
type StorageAccountUpdateParameters = original.StorageAccountUpdateParameters
type StorageBundle = original.StorageBundle
type StorageListResult = original.StorageListResult
type SubjectAlternativeNames = original.SubjectAlternativeNames
type Trigger = original.Trigger
type X509CertificateProperties = original.X509CertificateProperties

func New() ManagementClient {
	return original.New()
}
func NewWithoutDefaults() ManagementClient {
	return original.NewWithoutDefaults()
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
