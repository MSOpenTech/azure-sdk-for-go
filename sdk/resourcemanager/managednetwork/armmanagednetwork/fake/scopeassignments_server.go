//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
	"net/http"
	"net/url"
	"regexp"
)

// ScopeAssignmentsServer is a fake server for instances of the armmanagednetwork.ScopeAssignmentsClient type.
type ScopeAssignmentsServer struct {
	// CreateOrUpdate is the fake for method ScopeAssignmentsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, scope string, scopeAssignmentName string, parameters armmanagednetwork.ScopeAssignment, options *armmanagednetwork.ScopeAssignmentsClientCreateOrUpdateOptions) (resp azfake.Responder[armmanagednetwork.ScopeAssignmentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ScopeAssignmentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, scope string, scopeAssignmentName string, options *armmanagednetwork.ScopeAssignmentsClientDeleteOptions) (resp azfake.Responder[armmanagednetwork.ScopeAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ScopeAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, scopeAssignmentName string, options *armmanagednetwork.ScopeAssignmentsClientGetOptions) (resp azfake.Responder[armmanagednetwork.ScopeAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ScopeAssignmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, options *armmanagednetwork.ScopeAssignmentsClientListOptions) (resp azfake.PagerResponder[armmanagednetwork.ScopeAssignmentsClientListResponse])
}

// NewScopeAssignmentsServerTransport creates a new instance of ScopeAssignmentsServerTransport with the provided implementation.
// The returned ScopeAssignmentsServerTransport instance is connected to an instance of armmanagednetwork.ScopeAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAssignmentsServerTransport(srv *ScopeAssignmentsServer) *ScopeAssignmentsServerTransport {
	return &ScopeAssignmentsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmanagednetwork.ScopeAssignmentsClientListResponse]](),
	}
}

// ScopeAssignmentsServerTransport connects instances of armmanagednetwork.ScopeAssignmentsClient to instances of ScopeAssignmentsServer.
// Don't use this type directly, use NewScopeAssignmentsServerTransport instead.
type ScopeAssignmentsServerTransport struct {
	srv          *ScopeAssignmentsServer
	newListPager *tracker[azfake.PagerResponder[armmanagednetwork.ScopeAssignmentsClientListResponse]]
}

// Do implements the policy.Transporter interface for ScopeAssignmentsServerTransport.
func (s *ScopeAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ScopeAssignmentsClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "ScopeAssignmentsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "ScopeAssignmentsClient.Get":
		resp, err = s.dispatchGet(req)
	case "ScopeAssignmentsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ScopeAssignmentsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetwork/scopeAssignments/(?P<scopeAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmanagednetwork.ScopeAssignment](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scopeAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("scopeAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), scopeParam, scopeAssignmentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScopeAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAssignmentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetwork/scopeAssignments/(?P<scopeAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scopeAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("scopeAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), scopeParam, scopeAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetwork/scopeAssignments/(?P<scopeAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scopeAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("scopeAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), scopeParam, scopeAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScopeAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAssignmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetwork/scopeAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(scopeParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmanagednetwork.ScopeAssignmentsClientListResponse, createLink func() string) {
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
