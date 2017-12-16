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

package powerbidedicated

import original "github.com/Azure/azure-sdk-for-go/services/powerbidedicated/mgmt/2017-10-01/powerbidedicated"

type CapacitiesClient = original.CapacitiesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ProvisioningState = original.ProvisioningState

const (
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Paused       ProvisioningState = original.Paused
	Pausing      ProvisioningState = original.Pausing
	Preparing    ProvisioningState = original.Preparing
	Provisioning ProvisioningState = original.Provisioning
	Resuming     ProvisioningState = original.Resuming
	Scaling      ProvisioningState = original.Scaling
	Succeeded    ProvisioningState = original.Succeeded
	Suspended    ProvisioningState = original.Suspended
	Suspending   ProvisioningState = original.Suspending
	Updating     ProvisioningState = original.Updating
)

type SkuTier = original.SkuTier

const (
	PBIEAzure SkuTier = original.PBIEAzure
)

type State = original.State

const (
	StateDeleting     State = original.StateDeleting
	StateFailed       State = original.StateFailed
	StatePaused       State = original.StatePaused
	StatePausing      State = original.StatePausing
	StatePreparing    State = original.StatePreparing
	StateProvisioning State = original.StateProvisioning
	StateResuming     State = original.StateResuming
	StateScaling      State = original.StateScaling
	StateSucceeded    State = original.StateSucceeded
	StateSuspended    State = original.StateSuspended
	StateSuspending   State = original.StateSuspending
	StateUpdating     State = original.StateUpdating
)

type CapacitiesCreateFuture = original.CapacitiesCreateFuture
type CapacitiesDeleteFuture = original.CapacitiesDeleteFuture
type CapacitiesResumeFuture = original.CapacitiesResumeFuture
type CapacitiesSuspendFuture = original.CapacitiesSuspendFuture
type CapacitiesUpdateFuture = original.CapacitiesUpdateFuture
type DedicatedCapacities = original.DedicatedCapacities
type DedicatedCapacity = original.DedicatedCapacity
type DedicatedCapacityAdministrators = original.DedicatedCapacityAdministrators
type DedicatedCapacityMutableProperties = original.DedicatedCapacityMutableProperties
type DedicatedCapacityProperties = original.DedicatedCapacityProperties
type DedicatedCapacityUpdateParameters = original.DedicatedCapacityUpdateParameters
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type SkuDetailsForExistingResource = original.SkuDetailsForExistingResource
type SkuEnumerationForExistingResourceResult = original.SkuEnumerationForExistingResourceResult
type SkuEnumerationForNewResourceResult = original.SkuEnumerationForNewResourceResult
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewCapacitiesClient(subscriptionID string) CapacitiesClient {
	return original.NewCapacitiesClient(subscriptionID)
}
func NewCapacitiesClientWithBaseURI(baseURI string, subscriptionID string) CapacitiesClient {
	return original.NewCapacitiesClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
