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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualHubBgpConnectionsServer is a fake server for instances of the armnetwork.VirtualHubBgpConnectionsClient type.
type VirtualHubBgpConnectionsServer struct {
	// NewListPager is the fake for method VirtualHubBgpConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualHubName string, options *armnetwork.VirtualHubBgpConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualHubBgpConnectionsClientListResponse])

	// BeginListAdvertisedRoutes is the fake for method VirtualHubBgpConnectionsClient.BeginListAdvertisedRoutes
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListAdvertisedRoutes func(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *armnetwork.VirtualHubBgpConnectionsClientBeginListAdvertisedRoutesOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse], errResp azfake.ErrorResponder)

	// BeginListLearnedRoutes is the fake for method VirtualHubBgpConnectionsClient.BeginListLearnedRoutes
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListLearnedRoutes func(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *armnetwork.VirtualHubBgpConnectionsClientBeginListLearnedRoutesOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListLearnedRoutesResponse], errResp azfake.ErrorResponder)
}

// NewVirtualHubBgpConnectionsServerTransport creates a new instance of VirtualHubBgpConnectionsServerTransport with the provided implementation.
// The returned VirtualHubBgpConnectionsServerTransport instance is connected to an instance of armnetwork.VirtualHubBgpConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualHubBgpConnectionsServerTransport(srv *VirtualHubBgpConnectionsServer) *VirtualHubBgpConnectionsServerTransport {
	return &VirtualHubBgpConnectionsServerTransport{
		srv:                       srv,
		newListPager:              newTracker[azfake.PagerResponder[armnetwork.VirtualHubBgpConnectionsClientListResponse]](),
		beginListAdvertisedRoutes: newTracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse]](),
		beginListLearnedRoutes:    newTracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListLearnedRoutesResponse]](),
	}
}

// VirtualHubBgpConnectionsServerTransport connects instances of armnetwork.VirtualHubBgpConnectionsClient to instances of VirtualHubBgpConnectionsServer.
// Don't use this type directly, use NewVirtualHubBgpConnectionsServerTransport instead.
type VirtualHubBgpConnectionsServerTransport struct {
	srv                       *VirtualHubBgpConnectionsServer
	newListPager              *tracker[azfake.PagerResponder[armnetwork.VirtualHubBgpConnectionsClientListResponse]]
	beginListAdvertisedRoutes *tracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse]]
	beginListLearnedRoutes    *tracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionsClientListLearnedRoutesResponse]]
}

// Do implements the policy.Transporter interface for VirtualHubBgpConnectionsServerTransport.
func (v *VirtualHubBgpConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualHubBgpConnectionsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualHubBgpConnectionsClient.BeginListAdvertisedRoutes":
		resp, err = v.dispatchBeginListAdvertisedRoutes(req)
	case "VirtualHubBgpConnectionsClient.BeginListLearnedRoutes":
		resp, err = v.dispatchBeginListLearnedRoutes(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bgpConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, virtualHubNameParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualHubBgpConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualHubBgpConnectionsServerTransport) dispatchBeginListAdvertisedRoutes(req *http.Request) (*http.Response, error) {
	if v.srv.BeginListAdvertisedRoutes == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListAdvertisedRoutes not implemented")}
	}
	beginListAdvertisedRoutes := v.beginListAdvertisedRoutes.get(req)
	if beginListAdvertisedRoutes == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bgpConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/advertisedRoutes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginListAdvertisedRoutes(req.Context(), resourceGroupNameParam, hubNameParam, connectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListAdvertisedRoutes = &respr
		v.beginListAdvertisedRoutes.add(req, beginListAdvertisedRoutes)
	}

	resp, err := server.PollerResponderNext(beginListAdvertisedRoutes, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginListAdvertisedRoutes.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListAdvertisedRoutes) {
		v.beginListAdvertisedRoutes.remove(req)
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionsServerTransport) dispatchBeginListLearnedRoutes(req *http.Request) (*http.Response, error) {
	if v.srv.BeginListLearnedRoutes == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListLearnedRoutes not implemented")}
	}
	beginListLearnedRoutes := v.beginListLearnedRoutes.get(req)
	if beginListLearnedRoutes == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bgpConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/learnedRoutes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginListLearnedRoutes(req.Context(), resourceGroupNameParam, hubNameParam, connectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListLearnedRoutes = &respr
		v.beginListLearnedRoutes.add(req, beginListLearnedRoutes)
	}

	resp, err := server.PollerResponderNext(beginListLearnedRoutes, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginListLearnedRoutes.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListLearnedRoutes) {
		v.beginListLearnedRoutes.remove(req)
	}

	return resp, nil
}
