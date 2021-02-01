package purview

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
// Code generated by Microsoft (R) AutoRest Code Generatoa.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/purview/mgmt/2020-12-01-preview/purview"

// Account describes the Purview account.
type Account struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Identity - The Purview account identity
	Identity *Identity `json:"identity,omitempty"`
	// Properties - The Purview account properties
	Properties *Properties `json:"properties,omitempty"`
	// DependsOn - The list of keys on which this entity depends on.
	DependsOn *[]string `json:"dependsOn,omitempty"`
	// Sku - The Purview account SKU.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})

	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Identity != nil {
		objectMap["identity"] = a.Identity
	}
	if a.Properties != nil {
		objectMap["properties"] = a.Properties
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	return json.Marshal(objectMap)
}

// Identity describes the Purview identity
type Identity struct {
	// Type - Identity type
	Type IdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})

	if i.Type != "" {
		objectMap["type"] = i.Type
	}

	return json.Marshal(objectMap)
}

// Properties describes the Purview properties
type Properties struct {
	NetworkAcls *NetworkAcls `json:"networkAcls,omitempty"`
}

// MarshalJSON is the custom marshaler for properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})

	if p.NetworkAcls != nil {
		objectMap["networkAcls"] = p.NetworkAcls
	}

	return json.Marshal(objectMap)
}

// NetworkAcls describes the Purview network ACLs
type NetworkAcls struct {
	DefaultAction NetworkACLAction `json:"defaultAction,omitempty"`
}

// MarshalJSON is the custom marshaler for network ACLs.
func (n NetworkAcls) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})

	if n.DefaultAction != "" {
		objectMap["defaultAction"] = n.DefaultAction
	}

	return json.Marshal(objectMap)
}

// Sku the SKU of the Purview account.
type Sku struct {
	// Name - Gets or sets the sku name. Required for account creation, optional for update.
	Name *string `json:"name,omitempty"`
	// Capacity - The number of capacity units. Possible values include: 4, 16.
	Capacity *int32 `json:"capacity,omitempty"`
}

// MarshalJSON is the custom marshaler for SKU.
func (s Sku) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})

	if s.Name != nil {
		objectMap["name"] = s.Name
	}

	if s.Capacity != nil {
		objectMap["capacity"] = s.Capacity
	}

	return json.Marshal(objectMap)
}

// AccountListResult the list of cognitive services accounts operation response.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of accounts.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; Gets the list of Cognitive Services accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
}

// AccountListResultPage contains a page of Account values.
type AccountListResultPage struct {
	fn  func(context.Context, AccountListResult) (AccountListResult, error)
	alr AccountListResult
}

// MarshalJSON is the custom marshaler for AccountListResult.
func (alr AccountListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if alr.NextLink != nil {
		objectMap["nextLink"] = alr.NextLink
	}
	return json.Marshal(objectMap)
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AccountListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (alr AccountListResult) hasNextLink() bool {
	return alr.NextLink != nil && len(*alr.NextLink) != 0
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.alr)
		if err != nil {
			return err
		}
		page.alr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// accountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AccountListResult) accountListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !alr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AccountListResultIterator provides access to a complete listing of Account values.
type AccountListResultIterator struct {
	i    int
	page AccountListResultPage
}

// Response returns the raw server response from the last page request.
func (page AccountListResultPage) Response() AccountListResult {
	return page.alr
}
