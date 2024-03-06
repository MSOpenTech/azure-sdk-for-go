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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
)

// OperationsForMonitorServer is a fake server for instances of the armmonitor.OperationsForMonitorClient type.
type OperationsForMonitorServer struct {
	// NewListPager is the fake for method OperationsForMonitorClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armmonitor.OperationsForMonitorClientListOptions) (resp azfake.PagerResponder[armmonitor.OperationsForMonitorClientListResponse])
}

// NewOperationsForMonitorServerTransport creates a new instance of OperationsForMonitorServerTransport with the provided implementation.
// The returned OperationsForMonitorServerTransport instance is connected to an instance of armmonitor.OperationsForMonitorClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationsForMonitorServerTransport(srv *OperationsForMonitorServer) *OperationsForMonitorServerTransport {
	return &OperationsForMonitorServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmonitor.OperationsForMonitorClientListResponse]](),
	}
}

// OperationsForMonitorServerTransport connects instances of armmonitor.OperationsForMonitorClient to instances of OperationsForMonitorServer.
// Don't use this type directly, use NewOperationsForMonitorServerTransport instead.
type OperationsForMonitorServerTransport struct {
	srv          *OperationsForMonitorServer
	newListPager *tracker[azfake.PagerResponder[armmonitor.OperationsForMonitorClientListResponse]]
}

// Do implements the policy.Transporter interface for OperationsForMonitorServerTransport.
func (o *OperationsForMonitorServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationsForMonitorClient.NewListPager":
		resp, err = o.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationsForMonitorServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := o.newListPager.get(req)
	if newListPager == nil {
		resp := o.srv.NewListPager(nil)
		newListPager = &resp
		o.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmonitor.OperationsForMonitorClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		o.newListPager.remove(req)
	}
	return resp, nil
}
