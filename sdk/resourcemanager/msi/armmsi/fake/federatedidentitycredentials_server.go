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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// FederatedIdentityCredentialsServer is a fake server for instances of the armmsi.FederatedIdentityCredentialsClient type.
type FederatedIdentityCredentialsServer struct {
	// CreateOrUpdate is the fake for method FederatedIdentityCredentialsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, parameters armmsi.FederatedIdentityCredential, options *armmsi.FederatedIdentityCredentialsClientCreateOrUpdateOptions) (resp azfake.Responder[armmsi.FederatedIdentityCredentialsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method FederatedIdentityCredentialsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, options *armmsi.FederatedIdentityCredentialsClientDeleteOptions) (resp azfake.Responder[armmsi.FederatedIdentityCredentialsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FederatedIdentityCredentialsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, options *armmsi.FederatedIdentityCredentialsClientGetOptions) (resp azfake.Responder[armmsi.FederatedIdentityCredentialsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FederatedIdentityCredentialsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armmsi.FederatedIdentityCredentialsClientListOptions) (resp azfake.PagerResponder[armmsi.FederatedIdentityCredentialsClientListResponse])
}

// NewFederatedIdentityCredentialsServerTransport creates a new instance of FederatedIdentityCredentialsServerTransport with the provided implementation.
// The returned FederatedIdentityCredentialsServerTransport instance is connected to an instance of armmsi.FederatedIdentityCredentialsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFederatedIdentityCredentialsServerTransport(srv *FederatedIdentityCredentialsServer) *FederatedIdentityCredentialsServerTransport {
	return &FederatedIdentityCredentialsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmsi.FederatedIdentityCredentialsClientListResponse]](),
	}
}

// FederatedIdentityCredentialsServerTransport connects instances of armmsi.FederatedIdentityCredentialsClient to instances of FederatedIdentityCredentialsServer.
// Don't use this type directly, use NewFederatedIdentityCredentialsServerTransport instead.
type FederatedIdentityCredentialsServerTransport struct {
	srv          *FederatedIdentityCredentialsServer
	newListPager *tracker[azfake.PagerResponder[armmsi.FederatedIdentityCredentialsClientListResponse]]
}

// Do implements the policy.Transporter interface for FederatedIdentityCredentialsServerTransport.
func (f *FederatedIdentityCredentialsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FederatedIdentityCredentialsClient.CreateOrUpdate":
		resp, err = f.dispatchCreateOrUpdate(req)
	case "FederatedIdentityCredentialsClient.Delete":
		resp, err = f.dispatchDelete(req)
	case "FederatedIdentityCredentialsClient.Get":
		resp, err = f.dispatchGet(req)
	case "FederatedIdentityCredentialsClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FederatedIdentityCredentialsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedIdentity/userAssignedIdentities/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/federatedIdentityCredentials/(?P<federatedIdentityCredentialResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmsi.FederatedIdentityCredential](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	federatedIdentityCredentialResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("federatedIdentityCredentialResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, federatedIdentityCredentialResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FederatedIdentityCredential, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FederatedIdentityCredentialsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if f.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedIdentity/userAssignedIdentities/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/federatedIdentityCredentials/(?P<federatedIdentityCredentialResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	federatedIdentityCredentialResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("federatedIdentityCredentialResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Delete(req.Context(), resourceGroupNameParam, resourceNameParam, federatedIdentityCredentialResourceNameParam, nil)
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

func (f *FederatedIdentityCredentialsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedIdentity/userAssignedIdentities/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/federatedIdentityCredentials/(?P<federatedIdentityCredentialResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	federatedIdentityCredentialResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("federatedIdentityCredentialResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, federatedIdentityCredentialResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FederatedIdentityCredential, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FederatedIdentityCredentialsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedIdentity/userAssignedIdentities/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/federatedIdentityCredentials`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam := getOptional(skiptokenUnescaped)
		var options *armmsi.FederatedIdentityCredentialsClientListOptions
		if topParam != nil || skiptokenParam != nil {
			options = &armmsi.FederatedIdentityCredentialsClientListOptions{
				Top:       topParam,
				Skiptoken: skiptokenParam,
			}
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, resourceNameParam, options)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmsi.FederatedIdentityCredentialsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}
