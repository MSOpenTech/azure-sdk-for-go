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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"regexp"
)

// SKUsServer is a fake server for instances of the armapimanagement.SKUsClient type.
type SKUsServer struct {
	// NewListPager is the fake for method SKUsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armapimanagement.SKUsClientListOptions) (resp azfake.PagerResponder[armapimanagement.SKUsClientListResponse])
}

// NewSKUsServerTransport creates a new instance of SKUsServerTransport with the provided implementation.
// The returned SKUsServerTransport instance is connected to an instance of armapimanagement.SKUsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSKUsServerTransport(srv *SKUsServer) *SKUsServerTransport {
	return &SKUsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armapimanagement.SKUsClientListResponse]](),
	}
}

// SKUsServerTransport connects instances of armapimanagement.SKUsClient to instances of SKUsServer.
// Don't use this type directly, use NewSKUsServerTransport instead.
type SKUsServerTransport struct {
	srv          *SKUsServer
	newListPager *tracker[azfake.PagerResponder[armapimanagement.SKUsClientListResponse]]
}

// Do implements the policy.Transporter interface for SKUsServerTransport.
func (s *SKUsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SKUsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SKUsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/skus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapimanagement.SKUsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}
