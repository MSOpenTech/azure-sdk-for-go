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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkspacesServer is a fake server for instances of the armoperationalinsights.WorkspacesClient type.
type WorkspacesServer struct {
	// BeginCreateOrUpdate is the fake for method WorkspacesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, parameters armoperationalinsights.Workspace, options *armoperationalinsights.WorkspacesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armoperationalinsights.WorkspacesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method WorkspacesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, options *armoperationalinsights.WorkspacesClientBeginDeleteOptions) (resp azfake.PollerResponder[armoperationalinsights.WorkspacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, options *armoperationalinsights.WorkspacesClientGetOptions) (resp azfake.Responder[armoperationalinsights.WorkspacesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkspacesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armoperationalinsights.WorkspacesClientListOptions) (resp azfake.PagerResponder[armoperationalinsights.WorkspacesClientListResponse])

	// NewListByResourceGroupPager is the fake for method WorkspacesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armoperationalinsights.WorkspacesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armoperationalinsights.WorkspacesClientListByResourceGroupResponse])

	// Update is the fake for method WorkspacesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, workspaceName string, parameters armoperationalinsights.WorkspacePatch, options *armoperationalinsights.WorkspacesClientUpdateOptions) (resp azfake.Responder[armoperationalinsights.WorkspacesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWorkspacesServerTransport creates a new instance of WorkspacesServerTransport with the provided implementation.
// The returned WorkspacesServerTransport instance is connected to an instance of armoperationalinsights.WorkspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspacesServerTransport(srv *WorkspacesServer) *WorkspacesServerTransport {
	return &WorkspacesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armoperationalinsights.WorkspacesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armoperationalinsights.WorkspacesClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armoperationalinsights.WorkspacesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armoperationalinsights.WorkspacesClientListByResourceGroupResponse]](),
	}
}

// WorkspacesServerTransport connects instances of armoperationalinsights.WorkspacesClient to instances of WorkspacesServer.
// Don't use this type directly, use NewWorkspacesServerTransport instead.
type WorkspacesServerTransport struct {
	srv                         *WorkspacesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armoperationalinsights.WorkspacesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armoperationalinsights.WorkspacesClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armoperationalinsights.WorkspacesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armoperationalinsights.WorkspacesClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for WorkspacesServerTransport.
func (w *WorkspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspacesClient.BeginCreateOrUpdate":
		resp, err = w.dispatchBeginCreateOrUpdate(req)
	case "WorkspacesClient.BeginDelete":
		resp, err = w.dispatchBeginDelete(req)
	case "WorkspacesClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspacesClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkspacesClient.NewListByResourceGroupPager":
		resp, err = w.dispatchNewListByResourceGroupPager(req)
	case "WorkspacesClient.Update":
		resp, err = w.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := w.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.Workspace](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		w.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		w.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		w.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := w.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
		if err != nil {
			return nil, err
		}
		forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armoperationalinsights.WorkspacesClientBeginDeleteOptions
		if forceParam != nil {
			options = &armoperationalinsights.WorkspacesClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := w.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		w.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		w.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		w.beginDelete.remove(req)
	}

	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Workspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := w.srv.NewListPager(nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
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

func (w *WorkspacesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := w.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		w.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		w.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (w *WorkspacesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.WorkspacePatch](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Update(req.Context(), resourceGroupNameParam, workspaceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Workspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
