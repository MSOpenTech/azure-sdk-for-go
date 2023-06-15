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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"net/http"
	"net/url"
	"regexp"
)

// TagsServer is a fake server for instances of the armresources.TagsClient type.
type TagsServer struct {
	// CreateOrUpdate is the fake for method TagsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, tagName string, options *armresources.TagsClientCreateOrUpdateOptions) (resp azfake.Responder[armresources.TagsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtScope is the fake for method TagsClient.CreateOrUpdateAtScope
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateAtScope func(ctx context.Context, scope string, parameters armresources.TagsResource, options *armresources.TagsClientCreateOrUpdateAtScopeOptions) (resp azfake.Responder[armresources.TagsClientCreateOrUpdateAtScopeResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateValue is the fake for method TagsClient.CreateOrUpdateValue
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdateValue func(ctx context.Context, tagName string, tagValue string, options *armresources.TagsClientCreateOrUpdateValueOptions) (resp azfake.Responder[armresources.TagsClientCreateOrUpdateValueResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TagsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, tagName string, options *armresources.TagsClientDeleteOptions) (resp azfake.Responder[armresources.TagsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtScope is the fake for method TagsClient.DeleteAtScope
	// HTTP status codes to indicate success: http.StatusOK
	DeleteAtScope func(ctx context.Context, scope string, options *armresources.TagsClientDeleteAtScopeOptions) (resp azfake.Responder[armresources.TagsClientDeleteAtScopeResponse], errResp azfake.ErrorResponder)

	// DeleteValue is the fake for method TagsClient.DeleteValue
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteValue func(ctx context.Context, tagName string, tagValue string, options *armresources.TagsClientDeleteValueOptions) (resp azfake.Responder[armresources.TagsClientDeleteValueResponse], errResp azfake.ErrorResponder)

	// GetAtScope is the fake for method TagsClient.GetAtScope
	// HTTP status codes to indicate success: http.StatusOK
	GetAtScope func(ctx context.Context, scope string, options *armresources.TagsClientGetAtScopeOptions) (resp azfake.Responder[armresources.TagsClientGetAtScopeResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TagsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armresources.TagsClientListOptions) (resp azfake.PagerResponder[armresources.TagsClientListResponse])

	// UpdateAtScope is the fake for method TagsClient.UpdateAtScope
	// HTTP status codes to indicate success: http.StatusOK
	UpdateAtScope func(ctx context.Context, scope string, parameters armresources.TagsPatchResource, options *armresources.TagsClientUpdateAtScopeOptions) (resp azfake.Responder[armresources.TagsClientUpdateAtScopeResponse], errResp azfake.ErrorResponder)
}

// NewTagsServerTransport creates a new instance of TagsServerTransport with the provided implementation.
// The returned TagsServerTransport instance is connected to an instance of armresources.TagsClient by way of the
// undefined.Transporter field.
func NewTagsServerTransport(srv *TagsServer) *TagsServerTransport {
	return &TagsServerTransport{srv: srv}
}

// TagsServerTransport connects instances of armresources.TagsClient to instances of TagsServer.
// Don't use this type directly, use NewTagsServerTransport instead.
type TagsServerTransport struct {
	srv          *TagsServer
	newListPager *azfake.PagerResponder[armresources.TagsClientListResponse]
}

// Do implements the policy.Transporter interface for TagsServerTransport.
func (t *TagsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TagsClient.CreateOrUpdate":
		resp, err = t.dispatchCreateOrUpdate(req)
	case "TagsClient.CreateOrUpdateAtScope":
		resp, err = t.dispatchCreateOrUpdateAtScope(req)
	case "TagsClient.CreateOrUpdateValue":
		resp, err = t.dispatchCreateOrUpdateValue(req)
	case "TagsClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TagsClient.DeleteAtScope":
		resp, err = t.dispatchDeleteAtScope(req)
	case "TagsClient.DeleteValue":
		resp, err = t.dispatchDeleteValue(req)
	case "TagsClient.GetAtScope":
		resp, err = t.dispatchGetAtScope(req)
	case "TagsClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	case "TagsClient.UpdateAtScope":
		resp, err = t.dispatchUpdateAtScope(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TagsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagNames/(?P<tagName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	tagNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tagName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), tagNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TagsServerTransport) dispatchCreateOrUpdateAtScope(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdateAtScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtScope not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Resources/tags/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.TagsResource](req)
	if err != nil {
		return nil, err
	}
	scopeUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdateAtScope(req.Context(), scopeUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TagsServerTransport) dispatchCreateOrUpdateValue(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdateValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateValue not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagNames/(?P<tagName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagValues/(?P<tagValue>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	tagNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tagName")])
	if err != nil {
		return nil, err
	}
	tagValueUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tagValue")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdateValue(req.Context(), tagNameUnescaped, tagValueUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagValue, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TagsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagNames/(?P<tagName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	tagNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tagName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), tagNameUnescaped, nil)
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

func (t *TagsServerTransport) dispatchDeleteAtScope(req *http.Request) (*http.Response, error) {
	if t.srv.DeleteAtScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAtScope not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Resources/tags/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.DeleteAtScope(req.Context(), scopeUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TagsServerTransport) dispatchDeleteValue(req *http.Request) (*http.Response, error) {
	if t.srv.DeleteValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteValue not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagNames/(?P<tagName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagValues/(?P<tagValue>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	tagNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tagName")])
	if err != nil {
		return nil, err
	}
	tagValueUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tagValue")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.DeleteValue(req.Context(), tagNameUnescaped, tagValueUnescaped, nil)
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

func (t *TagsServerTransport) dispatchGetAtScope(req *http.Request) (*http.Response, error) {
	if t.srv.GetAtScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtScope not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Resources/tags/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.GetAtScope(req.Context(), scopeUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TagsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if t.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagNames`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := t.srv.NewListPager(nil)
		t.newListPager = &resp
		server.PagerResponderInjectNextLinks(t.newListPager, req, func(page *armresources.TagsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(t.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(t.newListPager) {
		t.newListPager = nil
	}
	return resp, nil
}

func (t *TagsServerTransport) dispatchUpdateAtScope(req *http.Request) (*http.Response, error) {
	if t.srv.UpdateAtScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateAtScope not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Resources/tags/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.TagsPatchResource](req)
	if err != nil {
		return nil, err
	}
	scopeUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.UpdateAtScope(req.Context(), scopeUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
