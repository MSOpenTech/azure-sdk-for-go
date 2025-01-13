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
	"net/http"
	"net/url"
	"regexp"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// ConnectedEnvironmentsStoragesServer is a fake server for instances of the armappcontainers.ConnectedEnvironmentsStoragesClient type.
type ConnectedEnvironmentsStoragesServer struct {
	// CreateOrUpdate is the fake for method ConnectedEnvironmentsStoragesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, storageEnvelope armappcontainers.ConnectedEnvironmentStorage, options *armappcontainers.ConnectedEnvironmentsStoragesClientCreateOrUpdateOptions) (resp azfake.Responder[armappcontainers.ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ConnectedEnvironmentsStoragesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, options *armappcontainers.ConnectedEnvironmentsStoragesClientDeleteOptions) (resp azfake.Responder[armappcontainers.ConnectedEnvironmentsStoragesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectedEnvironmentsStoragesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, storageName string, options *armappcontainers.ConnectedEnvironmentsStoragesClientGetOptions) (resp azfake.Responder[armappcontainers.ConnectedEnvironmentsStoragesClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ConnectedEnvironmentsStoragesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *armappcontainers.ConnectedEnvironmentsStoragesClientListOptions) (resp azfake.Responder[armappcontainers.ConnectedEnvironmentsStoragesClientListResponse], errResp azfake.ErrorResponder)
}

// NewConnectedEnvironmentsStoragesServerTransport creates a new instance of ConnectedEnvironmentsStoragesServerTransport with the provided implementation.
// The returned ConnectedEnvironmentsStoragesServerTransport instance is connected to an instance of armappcontainers.ConnectedEnvironmentsStoragesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectedEnvironmentsStoragesServerTransport(srv *ConnectedEnvironmentsStoragesServer) *ConnectedEnvironmentsStoragesServerTransport {
	return &ConnectedEnvironmentsStoragesServerTransport{srv: srv}
}

// ConnectedEnvironmentsStoragesServerTransport connects instances of armappcontainers.ConnectedEnvironmentsStoragesClient to instances of ConnectedEnvironmentsStoragesServer.
// Don't use this type directly, use NewConnectedEnvironmentsStoragesServerTransport instead.
type ConnectedEnvironmentsStoragesServerTransport struct {
	srv *ConnectedEnvironmentsStoragesServer
}

// Do implements the policy.Transporter interface for ConnectedEnvironmentsStoragesServerTransport.
func (c *ConnectedEnvironmentsStoragesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectedEnvironmentsStoragesClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ConnectedEnvironmentsStoragesClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ConnectedEnvironmentsStoragesClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectedEnvironmentsStoragesClient.List":
		resp, err = c.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectedEnvironmentsStoragesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/connectedEnvironments/(?P<connectedEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages/(?P<storageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.ConnectedEnvironmentStorage](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	connectedEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedEnvironmentName")])
	if err != nil {
		return nil, err
	}
	storageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, connectedEnvironmentNameParam, storageNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectedEnvironmentStorage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectedEnvironmentsStoragesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/connectedEnvironments/(?P<connectedEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages/(?P<storageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	connectedEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedEnvironmentName")])
	if err != nil {
		return nil, err
	}
	storageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, connectedEnvironmentNameParam, storageNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectedEnvironmentsStoragesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/connectedEnvironments/(?P<connectedEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages/(?P<storageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	connectedEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedEnvironmentName")])
	if err != nil {
		return nil, err
	}
	storageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, connectedEnvironmentNameParam, storageNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectedEnvironmentStorage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectedEnvironmentsStoragesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if c.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/connectedEnvironments/(?P<connectedEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	connectedEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectedEnvironmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.List(req.Context(), resourceGroupNameParam, connectedEnvironmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectedEnvironmentStoragesCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
