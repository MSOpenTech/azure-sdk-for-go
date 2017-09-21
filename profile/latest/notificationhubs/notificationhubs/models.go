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
// commit ID: b19c730e3a5c747d9055c95884b5c0310f7f2f16

package notificationhubs

import original "github.com/Azure/azure-sdk-for-go/service/notificationhubs/management/2014-09-01/notificationhubs"

type AccessRights = original.AccessRights

const (
	Listen	AccessRights	= original.Listen
	Manage	AccessRights	= original.Manage
	Send	AccessRights	= original.Send
)

type NamespaceType = original.NamespaceType

const (
	Messaging	NamespaceType	= original.Messaging
	NotificationHub	NamespaceType	= original.NotificationHub
)

type AdmCredential = original.AdmCredential
type AdmCredentialProperties = original.AdmCredentialProperties
type ApnsCredential = original.ApnsCredential
type ApnsCredentialProperties = original.ApnsCredentialProperties
type BaiduCredential = original.BaiduCredential
type BaiduCredentialProperties = original.BaiduCredentialProperties
type CheckAvailabilityParameters = original.CheckAvailabilityParameters
type CheckAvailabilityResource = original.CheckAvailabilityResource
type CreateOrUpdateParameters = original.CreateOrUpdateParameters
type GcmCredential = original.GcmCredential
type GcmCredentialProperties = original.GcmCredentialProperties
type ListResult = original.ListResult
type MpnsCredential = original.MpnsCredential
type MpnsCredentialProperties = original.MpnsCredentialProperties
type NamespaceCreateOrUpdateParameters = original.NamespaceCreateOrUpdateParameters
type NamespaceListResult = original.NamespaceListResult
type NamespaceProperties = original.NamespaceProperties
type NamespaceResource = original.NamespaceResource
type Properties = original.Properties
type Resource = original.Resource
type ResourceListKeys = original.ResourceListKeys
type ResourceType = original.ResourceType
type SharedAccessAuthorizationRuleCreateOrUpdateParameters = original.SharedAccessAuthorizationRuleCreateOrUpdateParameters
type SharedAccessAuthorizationRuleListResult = original.SharedAccessAuthorizationRuleListResult
type SharedAccessAuthorizationRuleProperties = original.SharedAccessAuthorizationRuleProperties
type SharedAccessAuthorizationRuleResource = original.SharedAccessAuthorizationRuleResource
type SubResource = original.SubResource
type WnsCredential = original.WnsCredential
type WnsCredentialProperties = original.WnsCredentialProperties
type NamespacesClient = original.NamespacesClient
type GroupClient = original.GroupClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
