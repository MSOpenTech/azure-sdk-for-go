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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedClusterSnapshotsServer is a fake server for instances of the armcontainerservice.ManagedClusterSnapshotsClient type.
type ManagedClusterSnapshotsServer struct {
	// CreateOrUpdate is the fake for method ManagedClusterSnapshotsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, parameters armcontainerservice.ManagedClusterSnapshot, options *armcontainerservice.ManagedClusterSnapshotsClientCreateOrUpdateOptions) (resp azfake.Responder[armcontainerservice.ManagedClusterSnapshotsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagedClusterSnapshotsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClusterSnapshotsClientDeleteOptions) (resp azfake.Responder[armcontainerservice.ManagedClusterSnapshotsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedClusterSnapshotsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClusterSnapshotsClientGetOptions) (resp azfake.Responder[armcontainerservice.ManagedClusterSnapshotsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagedClusterSnapshotsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcontainerservice.ManagedClusterSnapshotsClientListOptions) (resp azfake.PagerResponder[armcontainerservice.ManagedClusterSnapshotsClientListResponse])

	// NewListByResourceGroupPager is the fake for method ManagedClusterSnapshotsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcontainerservice.ManagedClusterSnapshotsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcontainerservice.ManagedClusterSnapshotsClientListByResourceGroupResponse])

	// UpdateTags is the fake for method ManagedClusterSnapshotsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, resourceName string, parameters armcontainerservice.TagsObject, options *armcontainerservice.ManagedClusterSnapshotsClientUpdateTagsOptions) (resp azfake.Responder[armcontainerservice.ManagedClusterSnapshotsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewManagedClusterSnapshotsServerTransport creates a new instance of ManagedClusterSnapshotsServerTransport with the provided implementation.
// The returned ManagedClusterSnapshotsServerTransport instance is connected to an instance of armcontainerservice.ManagedClusterSnapshotsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedClusterSnapshotsServerTransport(srv *ManagedClusterSnapshotsServer) *ManagedClusterSnapshotsServerTransport {
	return &ManagedClusterSnapshotsServerTransport{
		srv:                         srv,
		newListPager:                newTracker[azfake.PagerResponder[armcontainerservice.ManagedClusterSnapshotsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armcontainerservice.ManagedClusterSnapshotsClientListByResourceGroupResponse]](),
	}
}

// ManagedClusterSnapshotsServerTransport connects instances of armcontainerservice.ManagedClusterSnapshotsClient to instances of ManagedClusterSnapshotsServer.
// Don't use this type directly, use NewManagedClusterSnapshotsServerTransport instead.
type ManagedClusterSnapshotsServerTransport struct {
	srv                         *ManagedClusterSnapshotsServer
	newListPager                *tracker[azfake.PagerResponder[armcontainerservice.ManagedClusterSnapshotsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armcontainerservice.ManagedClusterSnapshotsClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for ManagedClusterSnapshotsServerTransport.
func (m *ManagedClusterSnapshotsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedClusterSnapshotsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagedClusterSnapshotsClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ManagedClusterSnapshotsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedClusterSnapshotsClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	case "ManagedClusterSnapshotsClient.NewListByResourceGroupPager":
		resp, err = m.dispatchNewListByResourceGroupPager(req)
	case "ManagedClusterSnapshotsClient.UpdateTags":
		resp, err = m.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedClusterSnapshotsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclustersnapshots/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcontainerservice.ManagedClusterSnapshot](req)
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
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedClusterSnapshot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedClusterSnapshotsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclustersnapshots/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
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

func (m *ManagedClusterSnapshotsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclustersnapshots/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedClusterSnapshot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedClusterSnapshotsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclustersnapshots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListPager(nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerservice.ManagedClusterSnapshotsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedClusterSnapshotsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclustersnapshots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armcontainerservice.ManagedClusterSnapshotsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		m.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedClusterSnapshotsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if m.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedclustersnapshots/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcontainerservice.TagsObject](req)
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
	respr, errRespr := m.srv.UpdateTags(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedClusterSnapshot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
