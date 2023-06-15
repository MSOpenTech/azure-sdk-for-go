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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"regexp"
)

// AzureFirewallFqdnTagsServer is a fake server for instances of the armnetwork.AzureFirewallFqdnTagsClient type.
type AzureFirewallFqdnTagsServer struct {
	// NewListAllPager is the fake for method AzureFirewallFqdnTagsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.AzureFirewallFqdnTagsClientListAllOptions) (resp azfake.PagerResponder[armnetwork.AzureFirewallFqdnTagsClientListAllResponse])
}

// NewAzureFirewallFqdnTagsServerTransport creates a new instance of AzureFirewallFqdnTagsServerTransport with the provided implementation.
// The returned AzureFirewallFqdnTagsServerTransport instance is connected to an instance of armnetwork.AzureFirewallFqdnTagsClient by way of the
// undefined.Transporter field.
func NewAzureFirewallFqdnTagsServerTransport(srv *AzureFirewallFqdnTagsServer) *AzureFirewallFqdnTagsServerTransport {
	return &AzureFirewallFqdnTagsServerTransport{srv: srv}
}

// AzureFirewallFqdnTagsServerTransport connects instances of armnetwork.AzureFirewallFqdnTagsClient to instances of AzureFirewallFqdnTagsServer.
// Don't use this type directly, use NewAzureFirewallFqdnTagsServerTransport instead.
type AzureFirewallFqdnTagsServerTransport struct {
	srv             *AzureFirewallFqdnTagsServer
	newListAllPager *azfake.PagerResponder[armnetwork.AzureFirewallFqdnTagsClientListAllResponse]
}

// Do implements the policy.Transporter interface for AzureFirewallFqdnTagsServerTransport.
func (a *AzureFirewallFqdnTagsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureFirewallFqdnTagsClient.NewListAllPager":
		resp, err = a.dispatchNewListAllPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureFirewallFqdnTagsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if a.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewallFqdnTags`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAllPager(nil)
		a.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(a.newListAllPager, req, func(page *armnetwork.AzureFirewallFqdnTagsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListAllPager) {
		a.newListAllPager = nil
	}
	return resp, nil
}
