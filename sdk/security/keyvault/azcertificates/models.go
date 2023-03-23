//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azcertificates

import "time"

// Action - The action that will be executed.
type Action struct {
	// The type of the action.
	ActionType *ActionType `json:"action_type,omitempty"`
}

// AdministratorDetails - Details of the organization administrator of the certificate issuer.
type AdministratorDetails struct {
	// Email address.
	EmailAddress *string `json:"email,omitempty"`

	// First name.
	FirstName *string `json:"first_name,omitempty"`

	// Last name.
	LastName *string `json:"last_name,omitempty"`

	// Phone number.
	Phone *string `json:"phone,omitempty"`
}

// BackupCertificateResult - The backup certificate result, containing the backup blob.
type BackupCertificateResult struct {
	// READ-ONLY; The backup blob containing the backed up certificate.
	Value []byte `json:"value,omitempty" azure:"ro"`
}

// CertificateAttributes - The certificate management attributes.
type CertificateAttributes struct {
	// Determines whether the object is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Expiry date in UTC.
	Expires *time.Time `json:"exp,omitempty"`

	// Not before date in UTC.
	NotBefore *time.Time `json:"nbf,omitempty"`

	// READ-ONLY; Creation time in UTC.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32 `json:"recoverableDays,omitempty" azure:"ro"`

	// READ-ONLY; Reflects the deletion recovery level currently in effect for certificates in the current vault. If it contains
	// 'Purgeable', the certificate can be permanently deleted by a privileged user; otherwise,
	// only the system can purge the certificate, at the end of the retention interval.
	RecoveryLevel *DeletionRecoveryLevel `json:"recoveryLevel,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

// CertificateBundle - A certificate bundle consists of a certificate (X509) plus its attributes.
type CertificateBundle struct {
	// The certificate attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// CER contents of x509 certificate.
	CER []byte `json:"cer,omitempty"`

	// The content type of the secret. eg. 'application/x-pem-file' or 'application/x-pkcs12',
	ContentType *string `json:"contentType,omitempty"`

	// Application specific metadata in the form of key-value pairs
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The certificate id.
	ID *ID `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The key id.
	KID *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; The management policy.
	Policy *CertificatePolicy `json:"policy,omitempty" azure:"ro"`

	// READ-ONLY; The secret id.
	SID *string `json:"sid,omitempty" azure:"ro"`

	// READ-ONLY; Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty" azure:"ro"`
}

// CertificateIssuerItem - The certificate issuer item containing certificate issuer metadata.
type CertificateIssuerItem struct {
	// Certificate Identifier.
	ID *string `json:"id,omitempty"`

	// The issuer provider.
	Provider *string `json:"provider,omitempty"`
}

// CertificateIssuerListResult - The certificate issuer list result.
type CertificateIssuerListResult struct {
	// READ-ONLY; The URL to get the next set of certificate issuers.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of certificate issuers in the key vault along with a link to the next page
	// of certificate issuers.
	Value []*CertificateIssuerItem `json:"value,omitempty" azure:"ro"`
}

// CertificateItem - The certificate item containing certificate metadata.
type CertificateItem struct {
	// The certificate management attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// Certificate identifier.
	ID *ID `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty"`
}

// CertificateListResult - The certificate list result.
type CertificateListResult struct {
	// READ-ONLY; The URL to get the next set of certificates.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of certificates in the key vault along with a link to the next page of
	// certificates.
	Value []*CertificateItem `json:"value,omitempty" azure:"ro"`
}

// CertificateOperation - A certificate operation is returned in case of asynchronous requests.
type CertificateOperation struct {
	// The certificate signing request (CSR) that is being used in the certificate operation.
	CSR []byte `json:"csr,omitempty"`

	// Indicates if cancellation was requested on the certificate operation.
	CancellationRequested *bool `json:"cancellation_requested,omitempty"`

	// Error encountered, if any, during the certificate operation.
	Error *Error `json:"error,omitempty"`

	// Parameters for the issuer of the X509 component of a certificate.
	IssuerParameters *IssuerParameters `json:"issuer,omitempty"`

	// Identifier for the certificate operation.
	RequestID *string `json:"request_id,omitempty"`

	// Status of the certificate operation.
	Status *string `json:"status,omitempty"`

	// The status details of the certificate operation.
	StatusDetails *string `json:"status_details,omitempty"`

	// Location which contains the result of the certificate operation.
	Target *string `json:"target,omitempty"`

	// READ-ONLY; The certificate id.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// CertificatePolicy - Management policy for a certificate.
type CertificatePolicy struct {
	// The certificate attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// Parameters for the issuer of the X509 component of a certificate.
	IssuerParameters *IssuerParameters `json:"issuer,omitempty"`

	// Properties of the key backing a certificate.
	KeyProperties *KeyProperties `json:"key_props,omitempty"`

	// Actions that will be performed by Key Vault over the lifetime of a certificate.
	LifetimeActions []*LifetimeAction `json:"lifetime_actions,omitempty"`

	// Properties of the secret backing a certificate.
	SecretProperties *SecretProperties `json:"secret_props,omitempty"`

	// Properties of the X509 component of a certificate.
	X509CertificateProperties *X509CertificateProperties `json:"x509_props,omitempty"`

	// READ-ONLY; The certificate id.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// BackupCertificateOptions contains the optional parameters for the Client.BackupCertificate method.
type BackupCertificateOptions struct {
	// placeholder for future optional parameters
}

// CreateCertificateOptions contains the optional parameters for the Client.CreateCertificate method.
type CreateCertificateOptions struct {
	// placeholder for future optional parameters
}

// DeleteCertificateContactsOptions contains the optional parameters for the Client.DeleteCertificateContacts method.
type DeleteCertificateContactsOptions struct {
	// placeholder for future optional parameters
}

// DeleteCertificateIssuerOptions contains the optional parameters for the Client.DeleteCertificateIssuer method.
type DeleteCertificateIssuerOptions struct {
	// placeholder for future optional parameters
}

// DeleteCertificateOperationOptions contains the optional parameters for the Client.DeleteCertificateOperation method.
type DeleteCertificateOperationOptions struct {
	// placeholder for future optional parameters
}

// DeleteCertificateOptions contains the optional parameters for the Client.DeleteCertificate method.
type DeleteCertificateOptions struct {
	// placeholder for future optional parameters
}

// GetCertificateContactsOptions contains the optional parameters for the Client.GetCertificateContacts method.
type GetCertificateContactsOptions struct {
	// placeholder for future optional parameters
}

// GetCertificateIssuerOptions contains the optional parameters for the Client.GetCertificateIssuer method.
type GetCertificateIssuerOptions struct {
	// placeholder for future optional parameters
}

// GetCertificateOperationOptions contains the optional parameters for the Client.GetCertificateOperation method.
type GetCertificateOperationOptions struct {
	// placeholder for future optional parameters
}

// GetCertificateOptions contains the optional parameters for the Client.GetCertificate method.
type GetCertificateOptions struct {
	// placeholder for future optional parameters
}

// GetCertificatePolicyOptions contains the optional parameters for the Client.GetCertificatePolicy method.
type GetCertificatePolicyOptions struct {
	// placeholder for future optional parameters
}

// GetDeletedCertificateOptions contains the optional parameters for the Client.GetDeletedCertificate method.
type GetDeletedCertificateOptions struct {
	// placeholder for future optional parameters
}

// ImportCertificateOptions contains the optional parameters for the Client.ImportCertificate method.
type ImportCertificateOptions struct {
	// placeholder for future optional parameters
}

// ListCertificateIssuersOptions contains the optional parameters for the Client.NewListCertificateIssuersPager method.
type ListCertificateIssuersOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// ListCertificateVersionsOptions contains the optional parameters for the Client.NewListCertificateVersionsPager method.
type ListCertificateVersionsOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// ListCertificatesOptions contains the optional parameters for the Client.NewListCertificatesPager method.
type ListCertificatesOptions struct {
	// Specifies whether to include certificates which are not completely provisioned.
	IncludePending *bool
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// ListDeletedCertificatesOptions contains the optional parameters for the Client.NewListDeletedCertificatesPager method.
type ListDeletedCertificatesOptions struct {
	// Specifies whether to include certificates which are not completely provisioned.
	IncludePending *bool
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// MergeCertificateOptions contains the optional parameters for the Client.MergeCertificate method.
type MergeCertificateOptions struct {
	// placeholder for future optional parameters
}

// PurgeDeletedCertificateOptions contains the optional parameters for the Client.PurgeDeletedCertificate method.
type PurgeDeletedCertificateOptions struct {
	// placeholder for future optional parameters
}

// RecoverDeletedCertificateOptions contains the optional parameters for the Client.RecoverDeletedCertificate method.
type RecoverDeletedCertificateOptions struct {
	// placeholder for future optional parameters
}

// RestoreCertificateOptions contains the optional parameters for the Client.RestoreCertificate method.
type RestoreCertificateOptions struct {
	// placeholder for future optional parameters
}

// SetCertificateContactsOptions contains the optional parameters for the Client.SetCertificateContacts method.
type SetCertificateContactsOptions struct {
	// placeholder for future optional parameters
}

// SetCertificateIssuerOptions contains the optional parameters for the Client.SetCertificateIssuer method.
type SetCertificateIssuerOptions struct {
	// placeholder for future optional parameters
}

// UpdateCertificateIssuerOptions contains the optional parameters for the Client.UpdateCertificateIssuer method.
type UpdateCertificateIssuerOptions struct {
	// placeholder for future optional parameters
}

// UpdateCertificateOperationOptions contains the optional parameters for the Client.UpdateCertificateOperation method.
type UpdateCertificateOperationOptions struct {
	// placeholder for future optional parameters
}

// UpdateCertificateOptions contains the optional parameters for the Client.UpdateCertificate method.
type UpdateCertificateOptions struct {
	// placeholder for future optional parameters
}

// UpdateCertificatePolicyOptions contains the optional parameters for the Client.UpdateCertificatePolicy method.
type UpdateCertificatePolicyOptions struct {
	// placeholder for future optional parameters
}

// Contact - The contact information for the vault certificates.
type Contact struct {
	// Email address.
	EmailAddress *string `json:"email,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	// Phone number.
	Phone *string `json:"phone,omitempty"`
}

// Contacts - The contacts for the vault certificates.
type Contacts struct {
	// The contact list for the vault certificates.
	ContactList []*Contact `json:"contacts,omitempty"`

	// READ-ONLY; Identifier for the contacts collection.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// CreateCertificateParameters - The certificate create parameters.
type CreateCertificateParameters struct {
	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy `json:"policy,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// DeletedCertificateBundle - A Deleted Certificate consisting of its previous id, attributes and its tags, as well as information
// on when it will be purged.
type DeletedCertificateBundle struct {
	// The certificate attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// CER contents of x509 certificate.
	CER []byte `json:"cer,omitempty"`

	// The content type of the secret. eg. 'application/x-pem-file' or 'application/x-pkcs12',
	ContentType *string `json:"contentType,omitempty"`

	// The url of the recovery object, used to identify and recover the deleted certificate.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// Application specific metadata in the form of key-value pairs
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The time when the certificate was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The certificate id.
	ID *ID `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The key id.
	KID *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; The management policy.
	Policy *CertificatePolicy `json:"policy,omitempty" azure:"ro"`

	// READ-ONLY; The secret id.
	SID *string `json:"sid,omitempty" azure:"ro"`

	// READ-ONLY; The time when the certificate is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`

	// READ-ONLY; Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty" azure:"ro"`
}

// DeletedCertificateItem - The deleted certificate item containing metadata about the deleted certificate.
type DeletedCertificateItem struct {
	// The certificate management attributes.
	Attributes *CertificateAttributes `json:"attributes,omitempty"`

	// Certificate identifier.
	ID *ID `json:"id,omitempty"`

	// The url of the recovery object, used to identify and recover the deleted certificate.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// Thumbprint of the certificate.
	X509Thumbprint []byte `json:"x5t,omitempty"`

	// READ-ONLY; The time when the certificate was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the certificate is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// DeletedCertificateListResult - A list of certificates that have been deleted in this vault.
type DeletedCertificateListResult struct {
	// READ-ONLY; The URL to get the next set of deleted certificates.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of deleted certificates in the vault along with a link to the next page
	// of deleted certificates
	Value []*DeletedCertificateItem `json:"value,omitempty" azure:"ro"`
}

// Error - The key vault server error.
type Error struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The key vault server error.
	InnerError *Error `json:"innererror,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ImportCertificateParameters - The certificate import parameters.
type ImportCertificateParameters struct {
	// REQUIRED; Base64 encoded representation of the certificate object to import. This certificate needs to contain the private
	// key.
	Base64EncodedCertificate *string `json:"value,omitempty"`

	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy `json:"policy,omitempty"`

	// If the private key in base64EncodedCertificate is encrypted, the password used for encryption.
	Password *string `json:"pwd,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// IssuerAttributes - The attributes of an issuer managed by the Key Vault service.
type IssuerAttributes struct {
	// Determines whether the issuer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// READ-ONLY; Creation time in UTC.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

// IssuerBundle - The issuer for Key Vault certificate.
type IssuerBundle struct {
	// Attributes of the issuer object.
	Attributes *IssuerAttributes `json:"attributes,omitempty"`

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials `json:"credentials,omitempty"`

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`

	// The issuer provider.
	Provider *string `json:"provider,omitempty"`

	// READ-ONLY; Identifier for the issuer object.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// IssuerCredentials - The credentials to be used for the certificate issuer.
type IssuerCredentials struct {
	// The user name/account name/account id.
	AccountID *string `json:"account_id,omitempty"`

	// The password/secret/account key.
	Password *string `json:"pwd,omitempty"`
}

// IssuerParameters - Parameters for the issuer of the X509 component of a certificate.
type IssuerParameters struct {
	// Indicates if the certificates generated under this policy should be published to certificate transparency logs.
	CertificateTransparency *bool `json:"cert_transparency,omitempty"`

	// Certificate type as supported by the provider (optional); for example 'OV-SSL', 'EV-SSL'
	CertificateType *string `json:"cty,omitempty"`

	// Name of the referenced issuer object or reserved names; for example, 'Self' or 'Unknown'.
	Name *string `json:"name,omitempty"`
}

// KeyProperties - Properties of the key pair backing a certificate.
type KeyProperties struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Curve *JSONWebKeyCurveName `json:"crv,omitempty"`

	// Indicates if the private key can be exported.
	Exportable *bool `json:"exportable,omitempty"`

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32 `json:"key_size,omitempty"`

	// The type of key pair to be used for the certificate.
	KeyType *JSONWebKeyType `json:"kty,omitempty"`

	// Indicates if the same key pair will be used on certificate renewal.
	ReuseKey *bool `json:"reuse_key,omitempty"`
}

// LifetimeAction - Action and its trigger that will be performed by Key Vault over the lifetime of a certificate.
type LifetimeAction struct {
	// The action that will be executed.
	Action *Action `json:"action,omitempty"`

	// The condition that will execute the action.
	Trigger *Trigger `json:"trigger,omitempty"`
}

// MergeCertificateParameters - The certificate merge parameters
type MergeCertificateParameters struct {
	// REQUIRED; The certificate or the certificate chain to merge.
	X509Certificates [][]byte `json:"x5c,omitempty"`

	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// OrganizationDetails - Details of the organization of the certificate issuer.
type OrganizationDetails struct {
	// Details of the organization administrator.
	AdminDetails []*AdministratorDetails `json:"admin_details,omitempty"`

	// Id of the organization.
	ID *string `json:"id,omitempty"`
}

// RestoreCertificateParameters - The certificate restore parameters.
type RestoreCertificateParameters struct {
	// REQUIRED; The backup blob associated with a certificate bundle.
	CertificateBundleBackup []byte `json:"value,omitempty"`
}

// SecretProperties - Properties of the key backing a certificate.
type SecretProperties struct {
	// The media type (MIME type).
	ContentType *string `json:"contentType,omitempty"`
}

// SetCertificateIssuerParameters - The certificate issuer set parameters.
type SetCertificateIssuerParameters struct {
	// REQUIRED; The issuer provider.
	Provider *string `json:"provider,omitempty"`

	// Attributes of the issuer object.
	Attributes *IssuerAttributes `json:"attributes,omitempty"`

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials `json:"credentials,omitempty"`

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`
}

// SubjectAlternativeNames - The subject alternate names of a X509 object.
type SubjectAlternativeNames struct {
	// Domain names.
	DNSNames []*string `json:"dns_names,omitempty"`

	// Email addresses.
	Emails []*string `json:"emails,omitempty"`

	// User principal names.
	UPNs []*string `json:"upns,omitempty"`
}

// Trigger - A condition to be satisfied for an action to be executed.
type Trigger struct {
	// Days before expiry to attempt renewal. Value should be between 1 and validityinmonths multiplied by 27. If validityinmonths
	// is 36, then value should be between 1 and 972 (36 * 27).
	DaysBeforeExpiry *int32 `json:"days_before_expiry,omitempty"`

	// Percentage of lifetime at which to trigger. Value should be between 1 and 99.
	LifetimePercentage *int32 `json:"lifetime_percentage,omitempty"`
}

// UpdateCertificateIssuerParameters - The certificate issuer update parameters.
type UpdateCertificateIssuerParameters struct {
	// Attributes of the issuer object.
	Attributes *IssuerAttributes `json:"attributes,omitempty"`

	// The credentials to be used for the issuer.
	Credentials *IssuerCredentials `json:"credentials,omitempty"`

	// Details of the organization as provided to the issuer.
	OrganizationDetails *OrganizationDetails `json:"org_details,omitempty"`

	// The issuer provider.
	Provider *string `json:"provider,omitempty"`
}

// UpdateCertificateOperationParameter - The certificate operation update parameters.
type UpdateCertificateOperationParameter struct {
	// REQUIRED; Indicates if cancellation was requested on the certificate operation.
	CancellationRequested *bool `json:"cancellation_requested,omitempty"`
}

// UpdateCertificateParameters - The certificate update parameters.
type UpdateCertificateParameters struct {
	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy `json:"policy,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// X509CertificateProperties - Properties of the X509 component of a certificate.
type X509CertificateProperties struct {
	// The enhanced key usage.
	EKUs []*string `json:"ekus,omitempty"`

	// List of key usages.
	KeyUsage []*KeyUsageType `json:"key_usage,omitempty"`

	// The subject name. Should be a valid X509 distinguished Name.
	Subject *string `json:"subject,omitempty"`

	// The subject alternative names.
	SubjectAlternativeNames *SubjectAlternativeNames `json:"sans,omitempty"`

	// The duration that the certificate is valid in months.
	ValidityInMonths *int32 `json:"validity_months,omitempty"`
}
