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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"net/http"
	"net/url"
	"regexp"
)

// FrontendEndpointsServer is a fake server for instances of the armfrontdoor.FrontendEndpointsClient type.
type FrontendEndpointsServer struct {
	// BeginDisableHTTPS is the fake for method FrontendEndpointsClient.BeginDisableHTTPS
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDisableHTTPS func(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, options *armfrontdoor.FrontendEndpointsClientBeginDisableHTTPSOptions) (resp azfake.PollerResponder[armfrontdoor.FrontendEndpointsClientDisableHTTPSResponse], errResp azfake.ErrorResponder)

	// BeginEnableHTTPS is the fake for method FrontendEndpointsClient.BeginEnableHTTPS
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginEnableHTTPS func(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, customHTTPSConfiguration armfrontdoor.CustomHTTPSConfiguration, options *armfrontdoor.FrontendEndpointsClientBeginEnableHTTPSOptions) (resp azfake.PollerResponder[armfrontdoor.FrontendEndpointsClientEnableHTTPSResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FrontendEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, options *armfrontdoor.FrontendEndpointsClientGetOptions) (resp azfake.Responder[armfrontdoor.FrontendEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFrontDoorPager is the fake for method FrontendEndpointsClient.NewListByFrontDoorPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFrontDoorPager func(resourceGroupName string, frontDoorName string, options *armfrontdoor.FrontendEndpointsClientListByFrontDoorOptions) (resp azfake.PagerResponder[armfrontdoor.FrontendEndpointsClientListByFrontDoorResponse])
}

// NewFrontendEndpointsServerTransport creates a new instance of FrontendEndpointsServerTransport with the provided implementation.
// The returned FrontendEndpointsServerTransport instance is connected to an instance of armfrontdoor.FrontendEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFrontendEndpointsServerTransport(srv *FrontendEndpointsServer) *FrontendEndpointsServerTransport {
	return &FrontendEndpointsServerTransport{
		srv:                     srv,
		beginDisableHTTPS:       newTracker[azfake.PollerResponder[armfrontdoor.FrontendEndpointsClientDisableHTTPSResponse]](),
		beginEnableHTTPS:        newTracker[azfake.PollerResponder[armfrontdoor.FrontendEndpointsClientEnableHTTPSResponse]](),
		newListByFrontDoorPager: newTracker[azfake.PagerResponder[armfrontdoor.FrontendEndpointsClientListByFrontDoorResponse]](),
	}
}

// FrontendEndpointsServerTransport connects instances of armfrontdoor.FrontendEndpointsClient to instances of FrontendEndpointsServer.
// Don't use this type directly, use NewFrontendEndpointsServerTransport instead.
type FrontendEndpointsServerTransport struct {
	srv                     *FrontendEndpointsServer
	beginDisableHTTPS       *tracker[azfake.PollerResponder[armfrontdoor.FrontendEndpointsClientDisableHTTPSResponse]]
	beginEnableHTTPS        *tracker[azfake.PollerResponder[armfrontdoor.FrontendEndpointsClientEnableHTTPSResponse]]
	newListByFrontDoorPager *tracker[azfake.PagerResponder[armfrontdoor.FrontendEndpointsClientListByFrontDoorResponse]]
}

// Do implements the policy.Transporter interface for FrontendEndpointsServerTransport.
func (f *FrontendEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FrontendEndpointsClient.BeginDisableHTTPS":
		resp, err = f.dispatchBeginDisableHTTPS(req)
	case "FrontendEndpointsClient.BeginEnableHTTPS":
		resp, err = f.dispatchBeginEnableHTTPS(req)
	case "FrontendEndpointsClient.Get":
		resp, err = f.dispatchGet(req)
	case "FrontendEndpointsClient.NewListByFrontDoorPager":
		resp, err = f.dispatchNewListByFrontDoorPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FrontendEndpointsServerTransport) dispatchBeginDisableHTTPS(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDisableHTTPS == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDisableHTTPS not implemented")}
	}
	beginDisableHTTPS := f.beginDisableHTTPS.get(req)
	if beginDisableHTTPS == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/frontDoors/(?P<frontDoorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontendEndpoints/(?P<frontendEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disableHttps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		frontDoorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontDoorName")])
		if err != nil {
			return nil, err
		}
		frontendEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDisableHTTPS(req.Context(), resourceGroupNameParam, frontDoorNameParam, frontendEndpointNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDisableHTTPS = &respr
		f.beginDisableHTTPS.add(req, beginDisableHTTPS)
	}

	resp, err := server.PollerResponderNext(beginDisableHTTPS, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginDisableHTTPS.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDisableHTTPS) {
		f.beginDisableHTTPS.remove(req)
	}

	return resp, nil
}

func (f *FrontendEndpointsServerTransport) dispatchBeginEnableHTTPS(req *http.Request) (*http.Response, error) {
	if f.srv.BeginEnableHTTPS == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginEnableHTTPS not implemented")}
	}
	beginEnableHTTPS := f.beginEnableHTTPS.get(req)
	if beginEnableHTTPS == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/frontDoors/(?P<frontDoorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontendEndpoints/(?P<frontendEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/enableHttps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfrontdoor.CustomHTTPSConfiguration](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		frontDoorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontDoorName")])
		if err != nil {
			return nil, err
		}
		frontendEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginEnableHTTPS(req.Context(), resourceGroupNameParam, frontDoorNameParam, frontendEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginEnableHTTPS = &respr
		f.beginEnableHTTPS.add(req, beginEnableHTTPS)
	}

	resp, err := server.PollerResponderNext(beginEnableHTTPS, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginEnableHTTPS.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginEnableHTTPS) {
		f.beginEnableHTTPS.remove(req)
	}

	return resp, nil
}

func (f *FrontendEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/frontDoors/(?P<frontDoorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontendEndpoints/(?P<frontendEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	frontDoorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontDoorName")])
	if err != nil {
		return nil, err
	}
	frontendEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontendEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, frontDoorNameParam, frontendEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FrontendEndpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FrontendEndpointsServerTransport) dispatchNewListByFrontDoorPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByFrontDoorPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFrontDoorPager not implemented")}
	}
	newListByFrontDoorPager := f.newListByFrontDoorPager.get(req)
	if newListByFrontDoorPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/frontDoors/(?P<frontDoorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/frontendEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		frontDoorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("frontDoorName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByFrontDoorPager(resourceGroupNameParam, frontDoorNameParam, nil)
		newListByFrontDoorPager = &resp
		f.newListByFrontDoorPager.add(req, newListByFrontDoorPager)
		server.PagerResponderInjectNextLinks(newListByFrontDoorPager, req, func(page *armfrontdoor.FrontendEndpointsClientListByFrontDoorResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFrontDoorPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByFrontDoorPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFrontDoorPager) {
		f.newListByFrontDoorPager.remove(req)
	}
	return resp, nil
}
