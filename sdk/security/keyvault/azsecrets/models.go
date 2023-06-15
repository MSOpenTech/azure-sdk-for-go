//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azsecrets

import "time"

// BackupSecretResult - The backup secret result, containing the backup blob.
type BackupSecretResult struct {
	// READ-ONLY; The backup blob containing the backed up secret.
	Value []byte `json:"value,omitempty" azure:"ro"`
}

// BackupSecretOptions contains the optional parameters for the Client.BackupSecret method.
type BackupSecretOptions struct {
	// placeholder for future optional parameters
}

// DeleteSecretOptions contains the optional parameters for the Client.DeleteSecret method.
type DeleteSecretOptions struct {
	// placeholder for future optional parameters
}

// GetDeletedSecretOptions contains the optional parameters for the Client.GetDeletedSecret method.
type GetDeletedSecretOptions struct {
	// placeholder for future optional parameters
}

// GetSecretOptions contains the optional parameters for the Client.GetSecret method.
type GetSecretOptions struct {
	// placeholder for future optional parameters
}

// ListDeletedSecretPropertiesOptions contains the optional parameters for the Client.NewListDeletedSecretPropertiesPager
// method.
type ListDeletedSecretPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ListSecretPropertiesOptions contains the optional parameters for the Client.NewListSecretPropertiesPager method.
type ListSecretPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ListSecretPropertiesVersionsOptions contains the optional parameters for the Client.NewListSecretPropertiesVersionsPager
// method.
type ListSecretPropertiesVersionsOptions struct {
	// placeholder for future optional parameters
}

// PurgeDeletedSecretOptions contains the optional parameters for the Client.PurgeDeletedSecret method.
type PurgeDeletedSecretOptions struct {
	// placeholder for future optional parameters
}

// RecoverDeletedSecretOptions contains the optional parameters for the Client.RecoverDeletedSecret method.
type RecoverDeletedSecretOptions struct {
	// placeholder for future optional parameters
}

// RestoreSecretOptions contains the optional parameters for the Client.RestoreSecret method.
type RestoreSecretOptions struct {
	// placeholder for future optional parameters
}

// SetSecretOptions contains the optional parameters for the Client.SetSecret method.
type SetSecretOptions struct {
	// placeholder for future optional parameters
}

// UpdateSecretPropertiesOptions contains the optional parameters for the Client.UpdateSecretProperties method.
type UpdateSecretPropertiesOptions struct {
	// placeholder for future optional parameters
}

// DeletedSecret - A Deleted Secret consisting of its previous id, attributes and its tags, as well as information on when
// it will be purged.
type DeletedSecret struct {
	// The secret management attributes.
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// The content type of the secret.
	ContentType *string `json:"contentType,omitempty"`

	// The secret id.
	ID *ID `json:"id,omitempty"`

	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// The secret value.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV
	// certificate.
	KID *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a secret backing a certificate, then managed
	// will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// DeletedSecretProperties - The deleted secret item containing metadata about the deleted secret.
type DeletedSecretProperties struct {
	// The secret management attributes.
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// Secret identifier.
	ID *ID `json:"id,omitempty"`

	// The url of the recovery object, used to identify and recover the deleted secret.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The time when the secret was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a key backing a certificate, then managed
	// will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`

	// READ-ONLY; The time when the secret is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// DeletedSecretPropertiesListResult - The deleted secret list result
type DeletedSecretPropertiesListResult struct {
	// READ-ONLY; The URL to get the next set of deleted secrets.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of the deleted secrets in the vault along with a link to the next page
	// of deleted secrets
	Value []*DeletedSecretProperties `json:"value,omitempty" azure:"ro"`
}

// RestoreSecretParameters - The secret restore parameters.
type RestoreSecretParameters struct {
	// REQUIRED; The backup blob associated with a secret bundle.
	SecretBackup []byte `json:"value,omitempty"`
}

// Secret - A secret consisting of a value, id and its attributes.
type Secret struct {
	// The secret management attributes.
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// The content type of the secret.
	ContentType *string `json:"contentType,omitempty"`

	// The secret id.
	ID *ID `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// The secret value.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV
	// certificate.
	KID *string `json:"kid,omitempty" azure:"ro"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a secret backing a certificate, then managed
	// will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

// SecretAttributes - The secret management attributes.
type SecretAttributes struct {
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

	// READ-ONLY; Reflects the deletion recovery level currently in effect for secrets in the current vault. If it contains 'Purgeable',
	// the secret can be permanently deleted by a privileged user; otherwise, only the
	// system can purge the secret, at the end of the retention interval.
	RecoveryLevel *string `json:"recoveryLevel,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

// SecretProperties - The secret item containing secret metadata.
type SecretProperties struct {
	// The secret management attributes.
	Attributes *SecretAttributes `json:"attributes,omitempty"`

	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// Secret identifier.
	ID *ID `json:"id,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a key backing a certificate, then managed
	// will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

// SecretPropertiesListResult - The secret list result.
type SecretPropertiesListResult struct {
	// READ-ONLY; The URL to get the next set of secrets.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of secrets in the key vault along with a link to the next page of secrets.
	Value []*SecretProperties `json:"value,omitempty" azure:"ro"`
}

// SetSecretParameters - The secret set parameters.
type SetSecretParameters struct {
	// REQUIRED; The value of the secret.
	Value *string `json:"value,omitempty"`

	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// The secret management attributes.
	SecretAttributes *SecretAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// UpdateSecretPropertiesParameters - The secret update parameters.
type UpdateSecretPropertiesParameters struct {
	// Type of the secret value such as a password.
	ContentType *string `json:"contentType,omitempty"`

	// The secret management attributes.
	SecretAttributes *SecretAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}
