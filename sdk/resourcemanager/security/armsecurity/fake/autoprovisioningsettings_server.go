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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// AutoProvisioningSettingsServer is a fake server for instances of the armsecurity.AutoProvisioningSettingsClient type.
type AutoProvisioningSettingsServer struct {
	// Create is the fake for method AutoProvisioningSettingsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, settingName string, setting armsecurity.AutoProvisioningSetting, options *armsecurity.AutoProvisioningSettingsClientCreateOptions) (resp azfake.Responder[armsecurity.AutoProvisioningSettingsClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AutoProvisioningSettingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, settingName string, options *armsecurity.AutoProvisioningSettingsClientGetOptions) (resp azfake.Responder[armsecurity.AutoProvisioningSettingsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AutoProvisioningSettingsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsecurity.AutoProvisioningSettingsClientListOptions) (resp azfake.PagerResponder[armsecurity.AutoProvisioningSettingsClientListResponse])
}

// NewAutoProvisioningSettingsServerTransport creates a new instance of AutoProvisioningSettingsServerTransport with the provided implementation.
// The returned AutoProvisioningSettingsServerTransport instance is connected to an instance of armsecurity.AutoProvisioningSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAutoProvisioningSettingsServerTransport(srv *AutoProvisioningSettingsServer) *AutoProvisioningSettingsServerTransport {
	return &AutoProvisioningSettingsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurity.AutoProvisioningSettingsClientListResponse]](),
	}
}

// AutoProvisioningSettingsServerTransport connects instances of armsecurity.AutoProvisioningSettingsClient to instances of AutoProvisioningSettingsServer.
// Don't use this type directly, use NewAutoProvisioningSettingsServerTransport instead.
type AutoProvisioningSettingsServerTransport struct {
	srv          *AutoProvisioningSettingsServer
	newListPager *tracker[azfake.PagerResponder[armsecurity.AutoProvisioningSettingsClientListResponse]]
}

// Do implements the policy.Transporter interface for AutoProvisioningSettingsServerTransport.
func (a *AutoProvisioningSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AutoProvisioningSettingsClient.Create":
		resp, err = a.dispatchCreate(req)
	case "AutoProvisioningSettingsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AutoProvisioningSettingsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AutoProvisioningSettingsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if a.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/autoProvisioningSettings/(?P<settingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurity.AutoProvisioningSetting](req)
	if err != nil {
		return nil, err
	}
	settingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("settingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Create(req.Context(), settingNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AutoProvisioningSetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutoProvisioningSettingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/autoProvisioningSettings/(?P<settingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	settingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("settingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), settingNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AutoProvisioningSetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutoProvisioningSettingsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/autoProvisioningSettings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.AutoProvisioningSettingsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
