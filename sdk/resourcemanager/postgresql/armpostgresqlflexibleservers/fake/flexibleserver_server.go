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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
	"net/http"
	"net/url"
	"regexp"
)

// FlexibleServerServer is a fake server for instances of the armpostgresqlflexibleservers.FlexibleServerClient type.
type FlexibleServerServer struct {
	// BeginStartLtrBackup is the fake for method FlexibleServerClient.BeginStartLtrBackup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartLtrBackup func(ctx context.Context, resourceGroupName string, serverName string, parameters armpostgresqlflexibleservers.LtrBackupRequest, options *armpostgresqlflexibleservers.FlexibleServerClientBeginStartLtrBackupOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.FlexibleServerClientStartLtrBackupResponse], errResp azfake.ErrorResponder)

	// TriggerLtrPreBackup is the fake for method FlexibleServerClient.TriggerLtrPreBackup
	// HTTP status codes to indicate success: http.StatusOK
	TriggerLtrPreBackup func(ctx context.Context, resourceGroupName string, serverName string, parameters armpostgresqlflexibleservers.LtrPreBackupRequest, options *armpostgresqlflexibleservers.FlexibleServerClientTriggerLtrPreBackupOptions) (resp azfake.Responder[armpostgresqlflexibleservers.FlexibleServerClientTriggerLtrPreBackupResponse], errResp azfake.ErrorResponder)
}

// NewFlexibleServerServerTransport creates a new instance of FlexibleServerServerTransport with the provided implementation.
// The returned FlexibleServerServerTransport instance is connected to an instance of armpostgresqlflexibleservers.FlexibleServerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFlexibleServerServerTransport(srv *FlexibleServerServer) *FlexibleServerServerTransport {
	return &FlexibleServerServerTransport{
		srv:                 srv,
		beginStartLtrBackup: newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.FlexibleServerClientStartLtrBackupResponse]](),
	}
}

// FlexibleServerServerTransport connects instances of armpostgresqlflexibleservers.FlexibleServerClient to instances of FlexibleServerServer.
// Don't use this type directly, use NewFlexibleServerServerTransport instead.
type FlexibleServerServerTransport struct {
	srv                 *FlexibleServerServer
	beginStartLtrBackup *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.FlexibleServerClientStartLtrBackupResponse]]
}

// Do implements the policy.Transporter interface for FlexibleServerServerTransport.
func (f *FlexibleServerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FlexibleServerClient.BeginStartLtrBackup":
		resp, err = f.dispatchBeginStartLtrBackup(req)
	case "FlexibleServerClient.TriggerLtrPreBackup":
		resp, err = f.dispatchTriggerLtrPreBackup(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FlexibleServerServerTransport) dispatchBeginStartLtrBackup(req *http.Request) (*http.Response, error) {
	if f.srv.BeginStartLtrBackup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartLtrBackup not implemented")}
	}
	beginStartLtrBackup := f.beginStartLtrBackup.get(req)
	if beginStartLtrBackup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startLtrBackup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.LtrBackupRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginStartLtrBackup(req.Context(), resourceGroupNameParam, serverNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStartLtrBackup = &respr
		f.beginStartLtrBackup.add(req, beginStartLtrBackup)
	}

	resp, err := server.PollerResponderNext(beginStartLtrBackup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginStartLtrBackup.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStartLtrBackup) {
		f.beginStartLtrBackup.remove(req)
	}

	return resp, nil
}

func (f *FlexibleServerServerTransport) dispatchTriggerLtrPreBackup(req *http.Request) (*http.Response, error) {
	if f.srv.TriggerLtrPreBackup == nil {
		return nil, &nonRetriableError{errors.New("fake for method TriggerLtrPreBackup not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ltrPreBackup`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.LtrPreBackupRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.TriggerLtrPreBackup(req.Context(), resourceGroupNameParam, serverNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LtrPreBackupResponse, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}
