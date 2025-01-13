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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkSecurityPerimeterConfigurationsServer is a fake server for instances of the armcognitiveservices.NetworkSecurityPerimeterConfigurationsClient type.
type NetworkSecurityPerimeterConfigurationsServer struct {
	// Get is the fake for method NetworkSecurityPerimeterConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, nspConfigurationName string, options *armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientGetOptions) (resp azfake.Responder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NetworkSecurityPerimeterConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientListOptions) (resp azfake.PagerResponder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientListResponse])

	// BeginReconcile is the fake for method NetworkSecurityPerimeterConfigurationsClient.BeginReconcile
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReconcile func(ctx context.Context, resourceGroupName string, accountName string, nspConfigurationName string, options *armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (resp azfake.PollerResponder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientReconcileResponse], errResp azfake.ErrorResponder)
}

// NewNetworkSecurityPerimeterConfigurationsServerTransport creates a new instance of NetworkSecurityPerimeterConfigurationsServerTransport with the provided implementation.
// The returned NetworkSecurityPerimeterConfigurationsServerTransport instance is connected to an instance of armcognitiveservices.NetworkSecurityPerimeterConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkSecurityPerimeterConfigurationsServerTransport(srv *NetworkSecurityPerimeterConfigurationsServer) *NetworkSecurityPerimeterConfigurationsServerTransport {
	return &NetworkSecurityPerimeterConfigurationsServerTransport{
		srv:            srv,
		newListPager:   newTracker[azfake.PagerResponder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientListResponse]](),
		beginReconcile: newTracker[azfake.PollerResponder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientReconcileResponse]](),
	}
}

// NetworkSecurityPerimeterConfigurationsServerTransport connects instances of armcognitiveservices.NetworkSecurityPerimeterConfigurationsClient to instances of NetworkSecurityPerimeterConfigurationsServer.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsServerTransport instead.
type NetworkSecurityPerimeterConfigurationsServerTransport struct {
	srv            *NetworkSecurityPerimeterConfigurationsServer
	newListPager   *tracker[azfake.PagerResponder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientListResponse]]
	beginReconcile *tracker[azfake.PollerResponder[armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientReconcileResponse]]
}

// Do implements the policy.Transporter interface for NetworkSecurityPerimeterConfigurationsServerTransport.
func (n *NetworkSecurityPerimeterConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkSecurityPerimeterConfigurationsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkSecurityPerimeterConfigurationsClient.NewListPager":
		resp, err = n.dispatchNewListPager(req)
	case "NetworkSecurityPerimeterConfigurationsClient.BeginReconcile":
		resp, err = n.dispatchBeginReconcile(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<nspConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	nspConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("nspConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, nspConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSecurityPerimeterConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := n.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListPager(resourceGroupNameParam, accountNameParam, nil)
		newListPager = &resp
		n.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcognitiveservices.NetworkSecurityPerimeterConfigurationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		n.newListPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchBeginReconcile(req *http.Request) (*http.Response, error) {
	if n.srv.BeginReconcile == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReconcile not implemented")}
	}
	beginReconcile := n.beginReconcile.get(req)
	if beginReconcile == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<nspConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reconcile`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		nspConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("nspConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginReconcile(req.Context(), resourceGroupNameParam, accountNameParam, nspConfigurationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReconcile = &respr
		n.beginReconcile.add(req, beginReconcile)
	}

	resp, err := server.PollerResponderNext(beginReconcile, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginReconcile.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReconcile) {
		n.beginReconcile.remove(req)
	}

	return resp, nil
}
