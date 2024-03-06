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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"net/http"
	"net/url"
	"regexp"
)

// PipelineRunsServer is a fake server for instances of the armcontainerregistry.PipelineRunsClient type.
type PipelineRunsServer struct {
	// BeginCreate is the fake for method PipelineRunsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters armcontainerregistry.PipelineRun, options *armcontainerregistry.PipelineRunsClientBeginCreateOptions) (resp azfake.PollerResponder[armcontainerregistry.PipelineRunsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PipelineRunsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *armcontainerregistry.PipelineRunsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerregistry.PipelineRunsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PipelineRunsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *armcontainerregistry.PipelineRunsClientGetOptions) (resp azfake.Responder[armcontainerregistry.PipelineRunsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PipelineRunsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, options *armcontainerregistry.PipelineRunsClientListOptions) (resp azfake.PagerResponder[armcontainerregistry.PipelineRunsClientListResponse])
}

// NewPipelineRunsServerTransport creates a new instance of PipelineRunsServerTransport with the provided implementation.
// The returned PipelineRunsServerTransport instance is connected to an instance of armcontainerregistry.PipelineRunsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPipelineRunsServerTransport(srv *PipelineRunsServer) *PipelineRunsServerTransport {
	return &PipelineRunsServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armcontainerregistry.PipelineRunsClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armcontainerregistry.PipelineRunsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armcontainerregistry.PipelineRunsClientListResponse]](),
	}
}

// PipelineRunsServerTransport connects instances of armcontainerregistry.PipelineRunsClient to instances of PipelineRunsServer.
// Don't use this type directly, use NewPipelineRunsServerTransport instead.
type PipelineRunsServerTransport struct {
	srv          *PipelineRunsServer
	beginCreate  *tracker[azfake.PollerResponder[armcontainerregistry.PipelineRunsClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armcontainerregistry.PipelineRunsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armcontainerregistry.PipelineRunsClientListResponse]]
}

// Do implements the policy.Transporter interface for PipelineRunsServerTransport.
func (p *PipelineRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PipelineRunsClient.BeginCreate":
		resp, err = p.dispatchBeginCreate(req)
	case "PipelineRunsClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PipelineRunsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PipelineRunsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PipelineRunsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := p.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelineRuns/(?P<pipelineRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.PipelineRun](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		pipelineRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineRunName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreate(req.Context(), resourceGroupNameParam, registryNameParam, pipelineRunNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		p.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		p.beginCreate.remove(req)
	}

	return resp, nil
}

func (p *PipelineRunsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelineRuns/(?P<pipelineRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		pipelineRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineRunName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, registryNameParam, pipelineRunNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PipelineRunsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelineRuns/(?P<pipelineRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	pipelineRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineRunName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, registryNameParam, pipelineRunNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PipelineRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PipelineRunsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelineRuns`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, registryNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerregistry.PipelineRunsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}
