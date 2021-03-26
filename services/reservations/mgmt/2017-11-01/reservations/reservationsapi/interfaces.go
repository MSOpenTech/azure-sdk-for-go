package reservationsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/reservations/mgmt/2017-11-01/reservations"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetAppliedReservationList(ctx context.Context, subscriptionID string) (result reservations.AppliedReservations, err error)
	GetCatalog(ctx context.Context, subscriptionID string) (result reservations.ListCatalog, err error)
}

var _ BaseClientAPI = (*reservations.BaseClient)(nil)

// OrderClientAPI contains the set of methods on the OrderClient type.
type OrderClientAPI interface {
	Get(ctx context.Context, reservationOrderID string) (result reservations.OrderResponse, err error)
	List(ctx context.Context) (result reservations.OrderListPage, err error)
	ListComplete(ctx context.Context) (result reservations.OrderListIterator, err error)
}

var _ OrderClientAPI = (*reservations.OrderClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Get(ctx context.Context, reservationID string, reservationOrderID string) (result reservations.Response, err error)
	List(ctx context.Context, reservationOrderID string) (result reservations.ListPage, err error)
	ListComplete(ctx context.Context, reservationOrderID string) (result reservations.ListIterator, err error)
	ListRevisions(ctx context.Context, reservationID string, reservationOrderID string) (result reservations.ListPage, err error)
	ListRevisionsComplete(ctx context.Context, reservationID string, reservationOrderID string) (result reservations.ListIterator, err error)
	Merge(ctx context.Context, reservationOrderID string, body reservations.MergeRequest) (result reservations.ReservationMergeFuture, err error)
	Split(ctx context.Context, reservationOrderID string, body reservations.SplitRequest) (result reservations.SplitFuture, err error)
	Update(ctx context.Context, reservationOrderID string, reservationID string, parameters reservations.Patch) (result reservations.ReservationUpdateFuture, err error)
}

var _ ClientAPI = (*reservations.Client)(nil)

// OperationClientAPI contains the set of methods on the OperationClient type.
type OperationClientAPI interface {
	List(ctx context.Context) (result reservations.OperationListPage, err error)
	ListComplete(ctx context.Context) (result reservations.OperationListIterator, err error)
}

var _ OperationClientAPI = (*reservations.OperationClient)(nil)
