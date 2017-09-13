package relay

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.18.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen specifies the listen state for access rights.
	Listen AccessRights = "Listen"
	// Manage specifies the manage state for access rights.
	Manage AccessRights = "Manage"
	// Send specifies the send state for access rights.
	Send AccessRights = "Send"
)

// KeyType enumerates the values for key type.
type KeyType string

const (
	// PrimaryKey specifies the primary key state for key type.
	PrimaryKey KeyType = "PrimaryKey"
	// SecondaryKey specifies the secondary key state for key type.
	SecondaryKey KeyType = "SecondaryKey"
)

// ProvisioningStateEnum enumerates the values for provisioning state enum.
type ProvisioningStateEnum string

const (
	// Created specifies the created state for provisioning state enum.
	Created ProvisioningStateEnum = "Created"
	// Deleted specifies the deleted state for provisioning state enum.
	Deleted ProvisioningStateEnum = "Deleted"
	// Failed specifies the failed state for provisioning state enum.
	Failed ProvisioningStateEnum = "Failed"
	// Succeeded specifies the succeeded state for provisioning state enum.
	Succeeded ProvisioningStateEnum = "Succeeded"
	// Unknown specifies the unknown state for provisioning state enum.
	Unknown ProvisioningStateEnum = "Unknown"
	// Updating specifies the updating state for provisioning state enum.
	Updating ProvisioningStateEnum = "Updating"
)

// RelaytypeEnum enumerates the values for relaytype enum.
type RelaytypeEnum string

const (
	// HTTP specifies the http state for relaytype enum.
	HTTP RelaytypeEnum = "Http"
	// NetTCP specifies the net tcp state for relaytype enum.
	NetTCP RelaytypeEnum = "NetTcp"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Standard specifies the standard state for sku tier.
	Standard SkuTier = "Standard"
)

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// InvalidName specifies the invalid name state for unavailable reason.
	InvalidName UnavailableReason = "InvalidName"
	// NameInLockdown specifies the name in lockdown state for unavailable reason.
	NameInLockdown UnavailableReason = "NameInLockdown"
	// NameInUse specifies the name in use state for unavailable reason.
	NameInUse UnavailableReason = "NameInUse"
	// None specifies the none state for unavailable reason.
	None UnavailableReason = "None"
	// SubscriptionIsDisabled specifies the subscription is disabled state for unavailable reason.
	SubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// TooManyNamespaceInCurrentSubscription specifies the too many namespace in current subscription state for unavailable
	// reason.
	TooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// AccessKeys is namespace/Relay Connection String
type AccessKeys struct {
	autorest.Response         `json:"-"`
	PrimaryConnectionString   *string `json:"primaryConnectionString,omitempty"`
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
	PrimaryKey                *string `json:"primaryKey,omitempty"`
	SecondaryKey              *string `json:"secondaryKey,omitempty"`
	KeyName                   *string `json:"keyName,omitempty"`
}

// AuthorizationRule is description of a namespace authorization rule.
type AuthorizationRule struct {
	autorest.Response            `json:"-"`
	ID                           *string `json:"id,omitempty"`
	Name                         *string `json:"name,omitempty"`
	Type                         *string `json:"type,omitempty"`
	*AuthorizationRuleProperties `json:"properties,omitempty"`
}

// AuthorizationRuleProperties is authorization rule properties.
type AuthorizationRuleProperties struct {
	Rights *[]AccessRights `json:"rights,omitempty"`
}

// AuthorizationRuleListResult is the response from the list namespace operation.
type AuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]AuthorizationRule `json:"value,omitempty"`
	NextLink          *string              `json:"nextLink,omitempty"`
}

// AuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client AuthorizationRuleListResult) AuthorizationRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// CheckNameAvailability is description of the check name availability request properties.
type CheckNameAvailability struct {
	Name *string `json:"name,omitempty"`
}

// CheckNameAvailabilityResult is description of the check name availability request properties.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	Message           *string           `json:"message,omitempty"`
	NameAvailable     *bool             `json:"nameAvailable,omitempty"`
	Reason            UnavailableReason `json:"reason,omitempty"`
}

// ErrorResponse is error reponse indicates Relay service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// HybridConnection is description of hybrid connection resource.
type HybridConnection struct {
	autorest.Response           `json:"-"`
	ID                          *string `json:"id,omitempty"`
	Name                        *string `json:"name,omitempty"`
	Type                        *string `json:"type,omitempty"`
	*HybridConnectionProperties `json:"properties,omitempty"`
}

// HybridConnectionProperties is properties of the HybridConnection.
type HybridConnectionProperties struct {
	CreatedAt                   *date.Time `json:"createdAt,omitempty"`
	UpdatedAt                   *date.Time `json:"updatedAt,omitempty"`
	ListenerCount               *int32     `json:"listenerCount,omitempty"`
	RequiresClientAuthorization *bool      `json:"requiresClientAuthorization,omitempty"`
	UserMetadata                *string    `json:"userMetadata,omitempty"`
}

// HybridConnectionListResult is the response of the list hybrid connection operation.
type HybridConnectionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]HybridConnection `json:"value,omitempty"`
	NextLink          *string             `json:"nextLink,omitempty"`
}

// HybridConnectionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client HybridConnectionListResult) HybridConnectionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Namespace is description of a namespace resource.
type Namespace struct {
	autorest.Response    `json:"-"`
	ID                   *string             `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Location             *string             `json:"location,omitempty"`
	Tags                 *map[string]*string `json:"tags,omitempty"`
	Sku                  *Sku                `json:"sku,omitempty"`
	*NamespaceProperties `json:"properties,omitempty"`
}

// NamespaceListResult is the response from the list namespace operation.
type NamespaceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Namespace `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// NamespaceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client NamespaceListResult) NamespaceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// NamespaceProperties is properties of the namespace.
type NamespaceProperties struct {
	ProvisioningState  ProvisioningStateEnum `json:"provisioningState,omitempty"`
	CreatedAt          *date.Time            `json:"createdAt,omitempty"`
	UpdatedAt          *date.Time            `json:"updatedAt,omitempty"`
	ServiceBusEndpoint *string               `json:"serviceBusEndpoint,omitempty"`
	MetricID           *string               `json:"metricId,omitempty"`
}

// Operation is a Relay REST API operation.
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result of the request to list Relay operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RegenerateAccessKeyParameters is parameters supplied to the regenerate authorization rule operation, specifies which
// key neeeds to be reset.
type RegenerateAccessKeyParameters struct {
	KeyType KeyType `json:"keyType,omitempty"`
	Key     *string `json:"key,omitempty"`
}

// Resource is the resource definition.
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ResourceNamespacePatch is definition of resource.
type ResourceNamespacePatch struct {
	ID   *string             `json:"id,omitempty"`
	Name *string             `json:"name,omitempty"`
	Type *string             `json:"type,omitempty"`
	Tags *map[string]*string `json:"tags,omitempty"`
}

// Sku is SKU of the namespace.
type Sku struct {
	Name *string `json:"name,omitempty"`
	Tier SkuTier `json:"tier,omitempty"`
}

// TrackedResource is definition of resource.
type TrackedResource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// UpdateParameters is description of a namespace resource.
type UpdateParameters struct {
	ID                   *string             `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Tags                 *map[string]*string `json:"tags,omitempty"`
	Sku                  *Sku                `json:"sku,omitempty"`
	*NamespaceProperties `json:"properties,omitempty"`
}

// WcfRelay is description of the WCF relay resource.
type WcfRelay struct {
	autorest.Response   `json:"-"`
	ID                  *string `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	Type                *string `json:"type,omitempty"`
	*WcfRelayProperties `json:"properties,omitempty"`
}

// WcfRelayProperties is properties of the WCF relay.
type WcfRelayProperties struct {
	IsDynamic                   *bool         `json:"isDynamic,omitempty"`
	CreatedAt                   *date.Time    `json:"createdAt,omitempty"`
	UpdatedAt                   *date.Time    `json:"updatedAt,omitempty"`
	ListenerCount               *int32        `json:"listenerCount,omitempty"`
	RelayType                   RelaytypeEnum `json:"relayType,omitempty"`
	RequiresClientAuthorization *bool         `json:"requiresClientAuthorization,omitempty"`
	RequiresTransportSecurity   *bool         `json:"requiresTransportSecurity,omitempty"`
	UserMetadata                *string       `json:"userMetadata,omitempty"`
}

// WcfRelaysListResult is the response of the list WCF relay operation.
type WcfRelaysListResult struct {
	autorest.Response `json:"-"`
	Value             *[]WcfRelay `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// WcfRelaysListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client WcfRelaysListResult) WcfRelaysListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}
