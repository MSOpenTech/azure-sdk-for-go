package mixedrealityapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/mixedreality/mgmt/2021-01-01/mixedreality"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailabilityLocal(ctx context.Context, location string, checkNameAvailability mixedreality.CheckNameAvailabilityRequest) (result mixedreality.CheckNameAvailabilityResponse, err error)
}

var _ BaseClientAPI = (*mixedreality.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result mixedreality.OperationPagePage, err error)
	ListComplete(ctx context.Context) (result mixedreality.OperationPageIterator, err error)
}

var _ OperationsClientAPI = (*mixedreality.OperationsClient)(nil)

// SpatialAnchorsAccountsClientAPI contains the set of methods on the SpatialAnchorsAccountsClient type.
type SpatialAnchorsAccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, spatialAnchorsAccount mixedreality.SpatialAnchorsAccount) (result mixedreality.SpatialAnchorsAccount, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result mixedreality.SpatialAnchorsAccount, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result mixedreality.SpatialAnchorsAccountPagePage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result mixedreality.SpatialAnchorsAccountPageIterator, err error)
	ListBySubscription(ctx context.Context) (result mixedreality.SpatialAnchorsAccountPagePage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result mixedreality.SpatialAnchorsAccountPageIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result mixedreality.AccountKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, regenerate mixedreality.AccountKeyRegenerateRequest) (result mixedreality.AccountKeys, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, spatialAnchorsAccount mixedreality.SpatialAnchorsAccount) (result mixedreality.SpatialAnchorsAccount, err error)
}

var _ SpatialAnchorsAccountsClientAPI = (*mixedreality.SpatialAnchorsAccountsClient)(nil)

// RemoteRenderingAccountsClientAPI contains the set of methods on the RemoteRenderingAccountsClient type.
type RemoteRenderingAccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, remoteRenderingAccount mixedreality.RemoteRenderingAccount) (result mixedreality.RemoteRenderingAccount, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result mixedreality.RemoteRenderingAccount, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result mixedreality.RemoteRenderingAccountPagePage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result mixedreality.RemoteRenderingAccountPageIterator, err error)
	ListBySubscription(ctx context.Context) (result mixedreality.RemoteRenderingAccountPagePage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result mixedreality.RemoteRenderingAccountPageIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result mixedreality.AccountKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, regenerate mixedreality.AccountKeyRegenerateRequest) (result mixedreality.AccountKeys, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, remoteRenderingAccount mixedreality.RemoteRenderingAccount) (result mixedreality.RemoteRenderingAccount, err error)
}

var _ RemoteRenderingAccountsClientAPI = (*mixedreality.RemoteRenderingAccountsClient)(nil)
