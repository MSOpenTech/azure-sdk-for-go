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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkflowTriggerHistoriesServer is a fake server for instances of the armappservice.WorkflowTriggerHistoriesClient type.
type WorkflowTriggerHistoriesServer struct {
	// Get is the fake for method WorkflowTriggerHistoriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *armappservice.WorkflowTriggerHistoriesClientGetOptions) (resp azfake.Responder[armappservice.WorkflowTriggerHistoriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowTriggerHistoriesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, name string, workflowName string, triggerName string, options *armappservice.WorkflowTriggerHistoriesClientListOptions) (resp azfake.PagerResponder[armappservice.WorkflowTriggerHistoriesClientListResponse])

	// BeginResubmit is the fake for method WorkflowTriggerHistoriesClient.BeginResubmit
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginResubmit func(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *armappservice.WorkflowTriggerHistoriesClientBeginResubmitOptions) (resp azfake.PollerResponder[armappservice.WorkflowTriggerHistoriesClientResubmitResponse], errResp azfake.ErrorResponder)
}

// NewWorkflowTriggerHistoriesServerTransport creates a new instance of WorkflowTriggerHistoriesServerTransport with the provided implementation.
// The returned WorkflowTriggerHistoriesServerTransport instance is connected to an instance of armappservice.WorkflowTriggerHistoriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowTriggerHistoriesServerTransport(srv *WorkflowTriggerHistoriesServer) *WorkflowTriggerHistoriesServerTransport {
	return &WorkflowTriggerHistoriesServerTransport{
		srv:           srv,
		newListPager:  newTracker[azfake.PagerResponder[armappservice.WorkflowTriggerHistoriesClientListResponse]](),
		beginResubmit: newTracker[azfake.PollerResponder[armappservice.WorkflowTriggerHistoriesClientResubmitResponse]](),
	}
}

// WorkflowTriggerHistoriesServerTransport connects instances of armappservice.WorkflowTriggerHistoriesClient to instances of WorkflowTriggerHistoriesServer.
// Don't use this type directly, use NewWorkflowTriggerHistoriesServerTransport instead.
type WorkflowTriggerHistoriesServerTransport struct {
	srv           *WorkflowTriggerHistoriesServer
	newListPager  *tracker[azfake.PagerResponder[armappservice.WorkflowTriggerHistoriesClientListResponse]]
	beginResubmit *tracker[azfake.PollerResponder[armappservice.WorkflowTriggerHistoriesClientResubmitResponse]]
}

// Do implements the policy.Transporter interface for WorkflowTriggerHistoriesServerTransport.
func (w *WorkflowTriggerHistoriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkflowTriggerHistoriesClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkflowTriggerHistoriesClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkflowTriggerHistoriesClient.BeginResubmit":
		resp, err = w.dispatchBeginResubmit(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkflowTriggerHistoriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/histories/(?P<historyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	historyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, triggerNameParam, historyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowTriggerHistory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowTriggerHistoriesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/histories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
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
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armappservice.WorkflowTriggerHistoriesClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armappservice.WorkflowTriggerHistoriesClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, nameParam, workflowNameParam, triggerNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.WorkflowTriggerHistoriesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkflowTriggerHistoriesServerTransport) dispatchBeginResubmit(req *http.Request) (*http.Response, error) {
	if w.srv.BeginResubmit == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginResubmit not implemented")}
	}
	beginResubmit := w.beginResubmit.get(req)
	if beginResubmit == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/histories/(?P<historyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resubmit`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
		if err != nil {
			return nil, err
		}
		historyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginResubmit(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, triggerNameParam, historyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginResubmit = &respr
		w.beginResubmit.add(req, beginResubmit)
	}

	resp, err := server.PollerResponderNext(beginResubmit, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		w.beginResubmit.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginResubmit) {
		w.beginResubmit.remove(req)
	}

	return resp, nil
}
