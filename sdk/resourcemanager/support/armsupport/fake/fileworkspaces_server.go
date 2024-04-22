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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"net/http"
	"net/url"
	"regexp"
)

// FileWorkspacesServer is a fake server for instances of the armsupport.FileWorkspacesClient type.
type FileWorkspacesServer struct {
	// Create is the fake for method FileWorkspacesClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, fileWorkspaceName string, options *armsupport.FileWorkspacesClientCreateOptions) (resp azfake.Responder[armsupport.FileWorkspacesClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FileWorkspacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, fileWorkspaceName string, options *armsupport.FileWorkspacesClientGetOptions) (resp azfake.Responder[armsupport.FileWorkspacesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewFileWorkspacesServerTransport creates a new instance of FileWorkspacesServerTransport with the provided implementation.
// The returned FileWorkspacesServerTransport instance is connected to an instance of armsupport.FileWorkspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFileWorkspacesServerTransport(srv *FileWorkspacesServer) *FileWorkspacesServerTransport {
	return &FileWorkspacesServerTransport{srv: srv}
}

// FileWorkspacesServerTransport connects instances of armsupport.FileWorkspacesClient to instances of FileWorkspacesServer.
// Don't use this type directly, use NewFileWorkspacesServerTransport instead.
type FileWorkspacesServerTransport struct {
	srv *FileWorkspacesServer
}

// Do implements the policy.Transporter interface for FileWorkspacesServerTransport.
func (f *FileWorkspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FileWorkspacesClient.Create":
		resp, err = f.dispatchCreate(req)
	case "FileWorkspacesClient.Get":
		resp, err = f.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FileWorkspacesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if f.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Support/fileWorkspaces/(?P<fileWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	fileWorkspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileWorkspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Create(req.Context(), fileWorkspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileWorkspaceDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FileWorkspacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Support/fileWorkspaces/(?P<fileWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	fileWorkspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileWorkspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), fileWorkspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileWorkspaceDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
