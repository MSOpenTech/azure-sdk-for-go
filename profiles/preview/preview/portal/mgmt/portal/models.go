// +build go1.9

// Copyright 2020 Microsoft Corporation
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

package portal

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/portal/mgmt/2019-01-01-preview/portal"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Configuration = original.Configuration
type ConfigurationList = original.ConfigurationList
type ConfigurationProperties = original.ConfigurationProperties
type Dashboard = original.Dashboard
type DashboardLens = original.DashboardLens
type DashboardListResult = original.DashboardListResult
type DashboardListResultIterator = original.DashboardListResultIterator
type DashboardListResultPage = original.DashboardListResultPage
type DashboardParts = original.DashboardParts
type DashboardPartsPosition = original.DashboardPartsPosition
type DashboardProperties = original.DashboardProperties
type DashboardsClient = original.DashboardsClient
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type OperationsClient = original.OperationsClient
type PatchableDashboard = original.PatchableDashboard
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceProviderOperation = original.ResourceProviderOperation
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type TenantConfigurationsClient = original.TenantConfigurationsClient
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDashboardListResultIterator(page DashboardListResultPage) DashboardListResultIterator {
	return original.NewDashboardListResultIterator(page)
}
func NewDashboardListResultPage(cur DashboardListResult, getNextPage func(context.Context, DashboardListResult) (DashboardListResult, error)) DashboardListResultPage {
	return original.NewDashboardListResultPage(cur, getNextPage)
}
func NewDashboardsClient(subscriptionID string) DashboardsClient {
	return original.NewDashboardsClient(subscriptionID)
}
func NewDashboardsClientWithBaseURI(baseURI string, subscriptionID string) DashboardsClient {
	return original.NewDashboardsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return original.NewResourceProviderOperationListIterator(page)
}
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return original.NewResourceProviderOperationListPage(cur, getNextPage)
}
func NewTenantConfigurationsClient(subscriptionID string) TenantConfigurationsClient {
	return original.NewTenantConfigurationsClient(subscriptionID)
}
func NewTenantConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) TenantConfigurationsClient {
	return original.NewTenantConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
