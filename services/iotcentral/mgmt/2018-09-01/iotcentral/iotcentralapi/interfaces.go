package iotcentralapi

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
	"github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2018-09-01/iotcentral"
)

// AppsClientAPI contains the set of methods on the AppsClient type.
type AppsClientAPI interface {
	CheckNameAvailability(ctx context.Context, operationInputs iotcentral.OperationInputs) (result iotcentral.AppAvailabilityInfo, err error)
	CheckSubdomainAvailability(ctx context.Context, operationInputs iotcentral.OperationInputs) (result iotcentral.AppAvailabilityInfo, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, app iotcentral.App) (result iotcentral.AppsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result iotcentral.AppsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result iotcentral.App, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result iotcentral.AppListResultPage, err error)
	ListBySubscription(ctx context.Context) (result iotcentral.AppListResultPage, err error)
	Templates(ctx context.Context) (result iotcentral.AppTemplates, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, appPatch iotcentral.AppPatch) (result iotcentral.AppsUpdateFuture, err error)
}

var _ AppsClientAPI = (*iotcentral.AppsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result iotcentral.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*iotcentral.OperationsClient)(nil)
