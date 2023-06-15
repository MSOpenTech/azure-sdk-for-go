//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"net/http"
)

// OperationsServer is a fake server for instances of the armsubscriptions.OperationsClient type.
type OperationsServer struct {
	// NewListPager is the fake for method OperationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsubscriptions.OperationsClientListOptions) (resp azfake.PagerResponder[armsubscriptions.OperationsClientListResponse])
}

// NewOperationsServerTransport creates a new instance of OperationsServerTransport with the provided implementation.
// The returned OperationsServerTransport instance is connected to an instance of armsubscriptions.OperationsClient by way of the
// undefined.Transporter field.
func NewOperationsServerTransport(srv *OperationsServer) *OperationsServerTransport {
	return &OperationsServerTransport{srv: srv}
}

// OperationsServerTransport connects instances of armsubscriptions.OperationsClient to instances of OperationsServer.
// Don't use this type directly, use NewOperationsServerTransport instead.
type OperationsServerTransport struct {
	srv          *OperationsServer
	newListPager *azfake.PagerResponder[armsubscriptions.OperationsClientListResponse]
}

// Do implements the policy.Transporter interface for OperationsServerTransport.
func (o *OperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationsClient.NewListPager":
		resp, err = o.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if o.newListPager == nil {
		resp := o.srv.NewListPager(nil)
		o.newListPager = &resp
		server.PagerResponderInjectNextLinks(o.newListPager, req, func(page *armsubscriptions.OperationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(o.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(o.newListPager) {
		o.newListPager = nil
	}
	return resp, nil
}
