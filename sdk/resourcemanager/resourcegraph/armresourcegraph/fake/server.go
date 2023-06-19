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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	"net/http"
)

// Server is a fake server for instances of the armresourcegraph.Client type.
type Server struct {
	// Resources is the fake for method Client.Resources
	// HTTP status codes to indicate success: http.StatusOK
	Resources func(ctx context.Context, query armresourcegraph.QueryRequest, options *armresourcegraph.ClientResourcesOptions) (resp azfake.Responder[armresourcegraph.ClientResourcesResponse], errResp azfake.ErrorResponder)

	// ResourcesHistory is the fake for method Client.ResourcesHistory
	// HTTP status codes to indicate success: http.StatusOK
	ResourcesHistory func(ctx context.Context, request armresourcegraph.ResourcesHistoryRequest, options *armresourcegraph.ClientResourcesHistoryOptions) (resp azfake.Responder[armresourcegraph.ClientResourcesHistoryResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armresourcegraph.Client by way of the
// undefined.Transporter field.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{srv: srv}
}

// ServerTransport connects instances of armresourcegraph.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv *Server
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.Resources":
		resp, err = s.dispatchResources(req)
	case "Client.ResourcesHistory":
		resp, err = s.dispatchResourcesHistory(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchResources(req *http.Request) (*http.Response, error) {
	if s.srv.Resources == nil {
		return nil, &nonRetriableError{errors.New("fake for method Resources not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[armresourcegraph.QueryRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Resources(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueryResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchResourcesHistory(req *http.Request) (*http.Response, error) {
	if s.srv.ResourcesHistory == nil {
		return nil, &nonRetriableError{errors.New("fake for method ResourcesHistory not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[armresourcegraph.ResourcesHistoryRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ResourcesHistory(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Interface, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
