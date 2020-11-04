package attestation

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/stable/attestation/mgmt/2020-10-01/attestation"

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// CloudError an error response from Attestation.
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from Attestation.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for displaying in a user interface.
	Message *string `json:"message,omitempty"`
}

// JSONWebKey ...
type JSONWebKey struct {
	// Alg - The "alg" (algorithm) parameter identifies the algorithm intended for
	// use with the key.  The values used should either be registered in the
	// IANA "JSON Web Signature and Encryption Algorithms" registry
	// established by [JWA] or be a value that contains a Collision-
	// Resistant Name.
	Alg *string `json:"alg,omitempty"`
	// Crv - The "crv" (curve) parameter identifies the curve type
	Crv *string `json:"crv,omitempty"`
	// D - RSA private exponent or ECC private key
	D *string `json:"d,omitempty"`
	// Dp - RSA Private Key Parameter
	Dp *string `json:"dp,omitempty"`
	// Dq - RSA Private Key Parameter
	Dq *string `json:"dq,omitempty"`
	// E - RSA public exponent, in Base64
	E *string `json:"e,omitempty"`
	// K - Symmetric key
	K *string `json:"k,omitempty"`
	// Kid - The "kid" (key ID) parameter is used to match a specific key.  This
	// is used, for instance, to choose among a set of keys within a JWK Set
	// during key rollover.  The structure of the "kid" value is
	// unspecified.  When "kid" values are used within a JWK Set, different
	// keys within the JWK Set SHOULD use distinct "kid" values.  (One
	// example in which different keys might use the same "kid" value is if
	// they have different "kty" (key type) values but are considered to be
	// equivalent alternatives by the application using them.)  The "kid"
	// value is a case-sensitive string.
	Kid *string `json:"kid,omitempty"`
	// Kty - The "kty" (key type) parameter identifies the cryptographic algorithm
	// family used with the key, such as "RSA" or "EC". "kty" values should
	// either be registered in the IANA "JSON Web Key Types" registry
	// established by [JWA] or be a value that contains a Collision-
	// Resistant Name.  The "kty" value is a case-sensitive string.
	Kty *string `json:"kty,omitempty"`
	// N - RSA modulus, in Base64
	N *string `json:"n,omitempty"`
	// P - RSA secret prime
	P *string `json:"p,omitempty"`
	// Q - RSA secret prime, with p < q
	Q *string `json:"q,omitempty"`
	// Qi - RSA Private Key Parameter
	Qi *string `json:"qi,omitempty"`
	// Use - Use ("public key use") identifies the intended use of
	// the public key. The "use" parameter is employed to indicate whether
	// a public key is used for encrypting data or verifying the signature
	// on data. Values are commonly "sig" (signature) or "enc" (encryption).
	Use *string `json:"use,omitempty"`
	// X - X coordinate for the Elliptic Curve point
	X *string `json:"x,omitempty"`
	// X5c - The "x5c" (X.509 certificate chain) parameter contains a chain of one
	// or more PKIX certificates [RFC5280].  The certificate chain is
	// represented as a JSON array of certificate value strings.  Each
	// string in the array is a base64-encoded (Section 4 of [RFC4648] --
	// not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value.
	// The PKIX certificate containing the key value MUST be the first
	// certificate.
	X5c *[]string `json:"x5c,omitempty"`
	// Y - Y coordinate for the Elliptic Curve point
	Y *string `json:"y,omitempty"`
}

// JSONWebKeySet ...
type JSONWebKeySet struct {
	// Keys - The value of the "keys" parameter is an array of JWK values.  By
	// default, the order of the JWK values within the array does not imply
	// an order of preference among them, although applications of JWK Sets
	// can choose to assign a meaning to the order for their purposes, if
	// desired.
	Keys *[]JSONWebKey `json:"keys,omitempty"`
}

// OperationList list of supported operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// SystemData - READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// Value - List of supported operations.
	Value *[]OperationsDefinition `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationList.
func (ol OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ol.Value != nil {
		objectMap["value"] = ol.Value
	}
	return json.Marshal(objectMap)
}

// OperationsDefinition definition object with the name and properties of an operation.
type OperationsDefinition struct {
	// Name - Name of the operation.
	Name *string `json:"name,omitempty"`
	// Display - Display object with properties of the operation.
	Display *OperationsDisplayDefinition `json:"display,omitempty"`
}

// OperationsDisplayDefinition display object with properties of the operation.
type OperationsDisplayDefinition struct {
	// Provider - Resource provider of the operation.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource for the operation.
	Resource *string `json:"resource,omitempty"`
	// Operation - Short description of the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// Provider attestation service response message.
type Provider struct {
	autorest.Response `json:"-"`
	// SystemData - READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// StatusResult - Describes Attestation service status.
	*StatusResult `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Provider.
func (p Provider) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.StatusResult != nil {
		objectMap["properties"] = p.StatusResult
	}
	if p.Tags != nil {
		objectMap["tags"] = p.Tags
	}
	if p.Location != nil {
		objectMap["location"] = p.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Provider struct.
func (p *Provider) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				p.SystemData = &systemData
			}
		case "properties":
			if v != nil {
				var statusResult StatusResult
				err = json.Unmarshal(*v, &statusResult)
				if err != nil {
					return err
				}
				p.StatusResult = &statusResult
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				p.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				p.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				p.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				p.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				p.Type = &typeVar
			}
		}
	}

	return nil
}

// ProviderListResult attestation Providers List.
type ProviderListResult struct {
	autorest.Response `json:"-"`
	// SystemData - READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData `json:"systemData,omitempty"`
	// Value - Attestation Provider array.
	Value *[]Provider `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for ProviderListResult.
func (plr ProviderListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if plr.Value != nil {
		objectMap["value"] = plr.Value
	}
	return json.Marshal(objectMap)
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// ServiceCreationParams parameters for creating an attestation service instance
type ServiceCreationParams struct {
	// Location - The supported Azure location where the attestation service instance should be created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags that will be assigned to the attestation service instance.
	Tags map[string]*string `json:"tags"`
	// Properties - Properties of the attestation service instance
	Properties *ServiceCreationSpecificParams `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceCreationParams.
func (scp ServiceCreationParams) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if scp.Location != nil {
		objectMap["location"] = scp.Location
	}
	if scp.Tags != nil {
		objectMap["tags"] = scp.Tags
	}
	if scp.Properties != nil {
		objectMap["properties"] = scp.Properties
	}
	return json.Marshal(objectMap)
}

// ServiceCreationSpecificParams client supplied parameters used to create a new attestation service
// instance.
type ServiceCreationSpecificParams struct {
	// PolicySigningCertificates - JSON Web Key Set defining a set of X.509 Certificates that will represent the parent certificate for the signing certificate used for policy operations
	PolicySigningCertificates *JSONWebKeySet `json:"policySigningCertificates,omitempty"`
}

// ServicePatchParams parameters for patching an attestation service instance
type ServicePatchParams struct {
	// Tags - The tags that will be assigned to the attestation service instance.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ServicePatchParams.
func (spp ServicePatchParams) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if spp.Tags != nil {
		objectMap["tags"] = spp.Tags
	}
	return json.Marshal(objectMap)
}

// StatusResult status of attestation service.
type StatusResult struct {
	// TrustModel - Trust model for the attestation service instance.
	TrustModel *string `json:"trustModel,omitempty"`
	// Status - Status of attestation service. Possible values include: 'Ready', 'NotReady', 'Error'
	Status ServiceStatus `json:"status,omitempty"`
	// AttestURI - Gets the uri of attestation service
	AttestURI *string `json:"attestUri,omitempty"`
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The type of identity that last modified the resource.
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
