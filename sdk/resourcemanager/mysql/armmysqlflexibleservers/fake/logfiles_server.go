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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
	"net/http"
	"net/url"
	"regexp"
)

// LogFilesServer is a fake server for instances of the armmysqlflexibleservers.LogFilesClient type.
type LogFilesServer struct {
	// NewListByServerPager is the fake for method LogFilesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armmysqlflexibleservers.LogFilesClientListByServerOptions) (resp azfake.PagerResponder[armmysqlflexibleservers.LogFilesClientListByServerResponse])
}

// NewLogFilesServerTransport creates a new instance of LogFilesServerTransport with the provided implementation.
// The returned LogFilesServerTransport instance is connected to an instance of armmysqlflexibleservers.LogFilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLogFilesServerTransport(srv *LogFilesServer) *LogFilesServerTransport {
	return &LogFilesServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armmysqlflexibleservers.LogFilesClientListByServerResponse]](),
	}
}

// LogFilesServerTransport connects instances of armmysqlflexibleservers.LogFilesClient to instances of LogFilesServer.
// Don't use this type directly, use NewLogFilesServerTransport instead.
type LogFilesServerTransport struct {
	srv                  *LogFilesServer
	newListByServerPager *tracker[azfake.PagerResponder[armmysqlflexibleservers.LogFilesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for LogFilesServerTransport.
func (l *LogFilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LogFilesClient.NewListByServerPager":
		resp, err = l.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LogFilesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := l.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/logFiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		l.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armmysqlflexibleservers.LogFilesClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		l.newListByServerPager.remove(req)
	}
	return resp, nil
}
