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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NotificationServer is a fake server for instances of the armapimanagement.NotificationClient type.
type NotificationServer struct {
	// CreateOrUpdate is the fake for method NotificationClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, notificationName armapimanagement.NotificationName, options *armapimanagement.NotificationClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.NotificationClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NotificationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, notificationName armapimanagement.NotificationName, options *armapimanagement.NotificationClientGetOptions) (resp azfake.Responder[armapimanagement.NotificationClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method NotificationClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, options *armapimanagement.NotificationClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.NotificationClientListByServiceResponse])
}

// NewNotificationServerTransport creates a new instance of NotificationServerTransport with the provided implementation.
// The returned NotificationServerTransport instance is connected to an instance of armapimanagement.NotificationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNotificationServerTransport(srv *NotificationServer) *NotificationServerTransport {
	return &NotificationServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.NotificationClientListByServiceResponse]](),
	}
}

// NotificationServerTransport connects instances of armapimanagement.NotificationClient to instances of NotificationServer.
// Don't use this type directly, use NewNotificationServerTransport instead.
type NotificationServerTransport struct {
	srv                   *NotificationServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.NotificationClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for NotificationServerTransport.
func (n *NotificationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NotificationClient.CreateOrUpdate":
		resp, err = n.dispatchCreateOrUpdate(req)
	case "NotificationClient.Get":
		resp, err = n.dispatchGet(req)
	case "NotificationClient.NewListByServicePager":
		resp, err = n.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NotificationServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notifications/(?P<notificationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	notificationNameParam, err := parseWithCast(matches[regex.SubexpIndex("notificationName")], func(v string) (armapimanagement.NotificationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.NotificationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.NotificationClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.NotificationClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := n.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, notificationNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notifications/(?P<notificationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	notificationNameParam, err := parseWithCast(matches[regex.SubexpIndex("notificationName")], func(v string) (armapimanagement.NotificationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.NotificationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, notificationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := n.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notifications`
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
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.NotificationClientListByServiceOptions
		if topParam != nil || skipParam != nil {
			options = &armapimanagement.NotificationClientListByServiceOptions{
				Top:  topParam,
				Skip: skipParam,
			}
		}
		resp := n.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, options)
		newListByServicePager = &resp
		n.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.NotificationClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		n.newListByServicePager.remove(req)
	}
	return resp, nil
}
