package hardwaresecuritymodules

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
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/hardwaresecuritymodules/mgmt/2018-10-31-preview/hardwaresecuritymodules"

// APIEntityReference the API entity reference.
type APIEntityReference struct {
	// ID - The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
	ID *string `json:"id,omitempty"`
}

// DedicatedHsm resource information with extended details.
type DedicatedHsm struct {
	autorest.Response `json:"-"`
	// DedicatedHsmProperties - Properties of the dedicated HSM
	*DedicatedHsmProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The Azure Resource Manager resource ID for the dedicated HSM.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the dedicated HSM.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type of the dedicated HSM.
	Type *string `json:"type,omitempty"`
	// Location - The supported Azure location where the dedicated HSM should be created.
	Location *string `json:"location,omitempty"`
	// Sku - SKU details
	Sku *Sku `json:"sku,omitempty"`
	// Zones - The Dedicated Hsm zones.
	Zones *[]string `json:"zones,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for DedicatedHsm.
func (dh DedicatedHsm) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dh.DedicatedHsmProperties != nil {
		objectMap["properties"] = dh.DedicatedHsmProperties
	}
	if dh.Location != nil {
		objectMap["location"] = dh.Location
	}
	if dh.Sku != nil {
		objectMap["sku"] = dh.Sku
	}
	if dh.Zones != nil {
		objectMap["zones"] = dh.Zones
	}
	if dh.Tags != nil {
		objectMap["tags"] = dh.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DedicatedHsm struct.
func (dh *DedicatedHsm) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var dedicatedHsmProperties DedicatedHsmProperties
				err = json.Unmarshal(*v, &dedicatedHsmProperties)
				if err != nil {
					return err
				}
				dh.DedicatedHsmProperties = &dedicatedHsmProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				dh.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				dh.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				dh.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				dh.Location = &location
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				dh.Sku = &sku
			}
		case "zones":
			if v != nil {
				var zones []string
				err = json.Unmarshal(*v, &zones)
				if err != nil {
					return err
				}
				dh.Zones = &zones
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				dh.Tags = tags
			}
		}
	}

	return nil
}

// DedicatedHsmCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type DedicatedHsmCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DedicatedHsmCreateOrUpdateFuture) Result(client DedicatedHsmClient) (dh DedicatedHsm, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hardwaresecuritymodules.DedicatedHsmCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if dh.Response.Response, err = future.GetResult(sender); err == nil && dh.Response.Response.StatusCode != http.StatusNoContent {
		dh, err = client.CreateOrUpdateResponder(dh.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmCreateOrUpdateFuture", "Result", dh.Response.Response, "Failure responding to request")
		}
	}
	return
}

// DedicatedHsmDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type DedicatedHsmDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DedicatedHsmDeleteFuture) Result(client DedicatedHsmClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hardwaresecuritymodules.DedicatedHsmDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// DedicatedHsmError the error exception.
type DedicatedHsmError struct {
	// Error - READ-ONLY
	Error *Error `json:"error,omitempty"`
}

// DedicatedHsmListResult list of dedicated HSMs
type DedicatedHsmListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of dedicated HSMs.
	Value *[]DedicatedHsm `json:"value,omitempty"`
	// NextLink - The URL to get the next set of dedicated hsms.
	NextLink *string `json:"nextLink,omitempty"`
}

// DedicatedHsmListResultIterator provides access to a complete listing of DedicatedHsm values.
type DedicatedHsmListResultIterator struct {
	i    int
	page DedicatedHsmListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DedicatedHsmListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *DedicatedHsmListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DedicatedHsmListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DedicatedHsmListResultIterator) Response() DedicatedHsmListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DedicatedHsmListResultIterator) Value() DedicatedHsm {
	if !iter.page.NotDone() {
		return DedicatedHsm{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DedicatedHsmListResultIterator type.
func NewDedicatedHsmListResultIterator(page DedicatedHsmListResultPage) DedicatedHsmListResultIterator {
	return DedicatedHsmListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dhlr DedicatedHsmListResult) IsEmpty() bool {
	return dhlr.Value == nil || len(*dhlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dhlr DedicatedHsmListResult) hasNextLink() bool {
	return dhlr.NextLink != nil && len(*dhlr.NextLink) != 0
}

// dedicatedHsmListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dhlr DedicatedHsmListResult) dedicatedHsmListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !dhlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dhlr.NextLink)))
}

// DedicatedHsmListResultPage contains a page of DedicatedHsm values.
type DedicatedHsmListResultPage struct {
	fn   func(context.Context, DedicatedHsmListResult) (DedicatedHsmListResult, error)
	dhlr DedicatedHsmListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DedicatedHsmListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHsmListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dhlr)
		if err != nil {
			return err
		}
		page.dhlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DedicatedHsmListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DedicatedHsmListResultPage) NotDone() bool {
	return !page.dhlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DedicatedHsmListResultPage) Response() DedicatedHsmListResult {
	return page.dhlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DedicatedHsmListResultPage) Values() []DedicatedHsm {
	if page.dhlr.IsEmpty() {
		return nil
	}
	return *page.dhlr.Value
}

// Creates a new instance of the DedicatedHsmListResultPage type.
func NewDedicatedHsmListResultPage(cur DedicatedHsmListResult, getNextPage func(context.Context, DedicatedHsmListResult) (DedicatedHsmListResult, error)) DedicatedHsmListResultPage {
	return DedicatedHsmListResultPage{
		fn:   getNextPage,
		dhlr: cur,
	}
}

// DedicatedHsmOperation REST API operation
type DedicatedHsmOperation struct {
	// Name - The name of the Dedicated HSM Resource Provider Operation.
	Name *string `json:"name,omitempty"`
	// IsDataAction - READ-ONLY; Gets or sets a value indicating whether it is a data plane action
	IsDataAction *string                       `json:"isDataAction,omitempty"`
	Display      *DedicatedHsmOperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for DedicatedHsmOperation.
func (dho DedicatedHsmOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dho.Name != nil {
		objectMap["name"] = dho.Name
	}
	if dho.Display != nil {
		objectMap["display"] = dho.Display
	}
	return json.Marshal(objectMap)
}

// DedicatedHsmOperationDisplay ...
type DedicatedHsmOperationDisplay struct {
	// Provider - The Resource Provider of the operation
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - The object that represents the operation.
	Description *string `json:"description,omitempty"`
}

// DedicatedHsmOperationListResult result of the request to list Dedicated HSM Provider operations. It
// contains a list of operations.
type DedicatedHsmOperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Dedicated HSM Resource Provider operations.
	Value *[]DedicatedHsmOperation `json:"value,omitempty"`
}

// DedicatedHsmPatchParameters patchable properties of the dedicated HSM
type DedicatedHsmPatchParameters struct {
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for DedicatedHsmPatchParameters.
func (dhpp DedicatedHsmPatchParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dhpp.Tags != nil {
		objectMap["tags"] = dhpp.Tags
	}
	return json.Marshal(objectMap)
}

// DedicatedHsmProperties properties of the dedicated hsm
type DedicatedHsmProperties struct {
	// NetworkProfile - Specifies the network interfaces of the dedicated hsm.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	// StampID - This field will be used when RP does not support Availability zones.
	StampID *string `json:"stampId,omitempty"`
	// StatusMessage - READ-ONLY; Resource Status Message.
	StatusMessage *string `json:"statusMessage,omitempty"`
	// ProvisioningState - READ-ONLY; Provisioning state. Possible values include: 'Succeeded', 'Provisioning', 'Allocating', 'Connecting', 'Failed', 'CheckingQuota', 'Deleting'
	ProvisioningState JSONWebKeyType `json:"provisioningState,omitempty"`
}

// MarshalJSON is the custom marshaler for DedicatedHsmProperties.
func (dhp DedicatedHsmProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dhp.NetworkProfile != nil {
		objectMap["networkProfile"] = dhp.NetworkProfile
	}
	if dhp.StampID != nil {
		objectMap["stampId"] = dhp.StampID
	}
	return json.Marshal(objectMap)
}

// DedicatedHsmUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type DedicatedHsmUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DedicatedHsmUpdateFuture) Result(client DedicatedHsmClient) (dh DedicatedHsm, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hardwaresecuritymodules.DedicatedHsmUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if dh.Response.Response, err = future.GetResult(sender); err == nil && dh.Response.Response.StatusCode != http.StatusNoContent {
		dh, err = client.UpdateResponder(dh.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hardwaresecuritymodules.DedicatedHsmUpdateFuture", "Result", dh.Response.Response, "Failure responding to request")
		}
	}
	return
}

// Error the key vault server error.
type Error struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// InnerError - READ-ONLY
	InnerError *Error `json:"innererror,omitempty"`
}

// NetworkInterface the network interface definition.
type NetworkInterface struct {
	// ID - READ-ONLY; The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
	ID *string `json:"id,omitempty"`
	// PrivateIPAddress - Private Ip address of the interface
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`
}

// MarshalJSON is the custom marshaler for NetworkInterface.
func (ni NetworkInterface) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ni.PrivateIPAddress != nil {
		objectMap["privateIpAddress"] = ni.PrivateIPAddress
	}
	return json.Marshal(objectMap)
}

// NetworkProfile ...
type NetworkProfile struct {
	// Subnet - Specifies the identifier of the subnet.
	Subnet *APIEntityReference `json:"subnet,omitempty"`
	// NetworkInterfaces - Specifies the list of resource Ids for the network interfaces associated with the dedicated HSM.
	NetworkInterfaces *[]NetworkInterface `json:"networkInterfaces,omitempty"`
}

// Resource dedicated HSM resource
type Resource struct {
	// ID - READ-ONLY; The Azure Resource Manager resource ID for the dedicated HSM.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the dedicated HSM.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type of the dedicated HSM.
	Type *string `json:"type,omitempty"`
	// Location - The supported Azure location where the dedicated HSM should be created.
	Location *string `json:"location,omitempty"`
	// Sku - SKU details
	Sku *Sku `json:"sku,omitempty"`
	// Zones - The Dedicated Hsm zones.
	Zones *[]string `json:"zones,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Sku != nil {
		objectMap["sku"] = r.Sku
	}
	if r.Zones != nil {
		objectMap["zones"] = r.Zones
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceListResult list of dedicated HSM resources.
type ResourceListResult struct {
	// Value - The list of dedicated HSM resources.
	Value *[]Resource `json:"value,omitempty"`
	// NextLink - The URL to get the next set of dedicated HSM resources.
	NextLink *string `json:"nextLink,omitempty"`
}

// Sku ...
type Sku struct {
	// Name - SKU of the dedicated HSM. Possible values include: 'SafeNetLunaNetworkHSMA790'
	Name Name `json:"name,omitempty"`
}
