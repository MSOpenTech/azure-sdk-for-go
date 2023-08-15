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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"regexp"
)

// QueueServer is a fake server for instances of the armstorage.QueueClient type.
type QueueServer struct {
	// Create is the fake for method QueueClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue armstorage.Queue, options *armstorage.QueueClientCreateOptions) (resp azfake.Responder[armstorage.QueueClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method QueueClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *armstorage.QueueClientDeleteOptions) (resp azfake.Responder[armstorage.QueueClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method QueueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, queueName string, options *armstorage.QueueClientGetOptions) (resp azfake.Responder[armstorage.QueueClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method QueueClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armstorage.QueueClientListOptions) (resp azfake.PagerResponder[armstorage.QueueClientListResponse])

	// Update is the fake for method QueueClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue armstorage.Queue, options *armstorage.QueueClientUpdateOptions) (resp azfake.Responder[armstorage.QueueClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewQueueServerTransport creates a new instance of QueueServerTransport with the provided implementation.
// The returned QueueServerTransport instance is connected to an instance of armstorage.QueueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQueueServerTransport(srv *QueueServer) *QueueServerTransport {
	return &QueueServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armstorage.QueueClientListResponse]](),
	}
}

// QueueServerTransport connects instances of armstorage.QueueClient to instances of QueueServer.
// Don't use this type directly, use NewQueueServerTransport instead.
type QueueServerTransport struct {
	srv          *QueueServer
	newListPager *tracker[azfake.PagerResponder[armstorage.QueueClientListResponse]]
}

// Do implements the policy.Transporter interface for QueueServerTransport.
func (q *QueueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "QueueClient.Create":
		resp, err = q.dispatchCreate(req)
	case "QueueClient.Delete":
		resp, err = q.dispatchDelete(req)
	case "QueueClient.Get":
		resp, err = q.dispatchGet(req)
	case "QueueClient.NewListPager":
		resp, err = q.dispatchNewListPager(req)
	case "QueueClient.Update":
		resp, err = q.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *QueueServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if q.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queueServices/default/queues/(?P<queueName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.Queue](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	queueNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("queueName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Create(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, queueNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Queue, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueueServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if q.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queueServices/default/queues/(?P<queueName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	queueNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("queueName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Delete(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, queueNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if q.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queueServices/default/queues/(?P<queueName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	queueNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("queueName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Get(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, queueNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Queue, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueueServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if q.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := q.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queueServices/default/queues`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam := getOptional(maxpagesizeUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armstorage.QueueClientListOptions
		if maxpagesizeParam != nil || filterParam != nil {
			options = &armstorage.QueueClientListOptions{
				Maxpagesize: maxpagesizeParam,
				Filter:      filterParam,
			}
		}
		resp := q.srv.NewListPager(resourceGroupNameUnescaped, accountNameUnescaped, options)
		newListPager = &resp
		q.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstorage.QueueClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		q.newListPager.remove(req)
	}
	return resp, nil
}

func (q *QueueServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if q.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queueServices/default/queues/(?P<queueName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.Queue](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	queueNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("queueName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Update(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, queueNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Queue, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
