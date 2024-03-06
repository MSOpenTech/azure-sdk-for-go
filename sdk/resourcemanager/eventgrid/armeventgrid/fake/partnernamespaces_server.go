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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PartnerNamespacesServer is a fake server for instances of the armeventgrid.PartnerNamespacesClient type.
type PartnerNamespacesServer struct {
	// BeginCreateOrUpdate is the fake for method PartnerNamespacesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceInfo armeventgrid.PartnerNamespace, options *armeventgrid.PartnerNamespacesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armeventgrid.PartnerNamespacesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PartnerNamespacesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *armeventgrid.PartnerNamespacesClientBeginDeleteOptions) (resp azfake.PollerResponder[armeventgrid.PartnerNamespacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PartnerNamespacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *armeventgrid.PartnerNamespacesClientGetOptions) (resp azfake.Responder[armeventgrid.PartnerNamespacesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method PartnerNamespacesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armeventgrid.PartnerNamespacesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armeventgrid.PartnerNamespacesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method PartnerNamespacesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armeventgrid.PartnerNamespacesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armeventgrid.PartnerNamespacesClientListBySubscriptionResponse])

	// ListSharedAccessKeys is the fake for method PartnerNamespacesClient.ListSharedAccessKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListSharedAccessKeys func(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *armeventgrid.PartnerNamespacesClientListSharedAccessKeysOptions) (resp azfake.Responder[armeventgrid.PartnerNamespacesClientListSharedAccessKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKey is the fake for method PartnerNamespacesClient.RegenerateKey
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKey func(ctx context.Context, resourceGroupName string, partnerNamespaceName string, regenerateKeyRequest armeventgrid.PartnerNamespaceRegenerateKeyRequest, options *armeventgrid.PartnerNamespacesClientRegenerateKeyOptions) (resp azfake.Responder[armeventgrid.PartnerNamespacesClientRegenerateKeyResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method PartnerNamespacesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceUpdateParameters armeventgrid.PartnerNamespaceUpdateParameters, options *armeventgrid.PartnerNamespacesClientBeginUpdateOptions) (resp azfake.PollerResponder[armeventgrid.PartnerNamespacesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPartnerNamespacesServerTransport creates a new instance of PartnerNamespacesServerTransport with the provided implementation.
// The returned PartnerNamespacesServerTransport instance is connected to an instance of armeventgrid.PartnerNamespacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPartnerNamespacesServerTransport(srv *PartnerNamespacesServer) *PartnerNamespacesServerTransport {
	return &PartnerNamespacesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armeventgrid.PartnerNamespacesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armeventgrid.PartnerNamespacesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armeventgrid.PartnerNamespacesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armeventgrid.PartnerNamespacesClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armeventgrid.PartnerNamespacesClientUpdateResponse]](),
	}
}

// PartnerNamespacesServerTransport connects instances of armeventgrid.PartnerNamespacesClient to instances of PartnerNamespacesServer.
// Don't use this type directly, use NewPartnerNamespacesServerTransport instead.
type PartnerNamespacesServerTransport struct {
	srv                         *PartnerNamespacesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armeventgrid.PartnerNamespacesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armeventgrid.PartnerNamespacesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armeventgrid.PartnerNamespacesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armeventgrid.PartnerNamespacesClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armeventgrid.PartnerNamespacesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for PartnerNamespacesServerTransport.
func (p *PartnerNamespacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PartnerNamespacesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PartnerNamespacesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PartnerNamespacesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PartnerNamespacesClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "PartnerNamespacesClient.NewListBySubscriptionPager":
		resp, err = p.dispatchNewListBySubscriptionPager(req)
	case "PartnerNamespacesClient.ListSharedAccessKeys":
		resp, err = p.dispatchListSharedAccessKeys(req)
	case "PartnerNamespacesClient.RegenerateKey":
		resp, err = p.dispatchRegenerateKey(req)
	case "PartnerNamespacesClient.BeginUpdate":
		resp, err = p.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces/(?P<partnerNamespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.PartnerNamespace](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		partnerNamespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerNamespaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, partnerNamespaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces/(?P<partnerNamespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		partnerNamespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerNamespaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, partnerNamespaceNameParam, nil)
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

func (p *PartnerNamespacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces/(?P<partnerNamespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	partnerNamespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerNamespaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, partnerNamespaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PartnerNamespace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := p.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		var options *armeventgrid.PartnerNamespacesClientListByResourceGroupOptions
		if filterParam != nil || topParam != nil {
			options = &armeventgrid.PartnerNamespacesClientListByResourceGroupOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		p.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armeventgrid.PartnerNamespacesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		p.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := p.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		var options *armeventgrid.PartnerNamespacesClientListBySubscriptionOptions
		if filterParam != nil || topParam != nil {
			options = &armeventgrid.PartnerNamespacesClientListBySubscriptionOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		p.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armeventgrid.PartnerNamespacesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		p.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchListSharedAccessKeys(req *http.Request) (*http.Response, error) {
	if p.srv.ListSharedAccessKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSharedAccessKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces/(?P<partnerNamespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	partnerNamespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerNamespaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ListSharedAccessKeys(req.Context(), resourceGroupNameParam, partnerNamespaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PartnerNamespaceSharedAccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchRegenerateKey(req *http.Request) (*http.Response, error) {
	if p.srv.RegenerateKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces/(?P<partnerNamespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventgrid.PartnerNamespaceRegenerateKeyRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	partnerNamespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerNamespaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.RegenerateKey(req.Context(), resourceGroupNameParam, partnerNamespaceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PartnerNamespaceSharedAccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PartnerNamespacesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventGrid/partnerNamespaces/(?P<partnerNamespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armeventgrid.PartnerNamespaceUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		partnerNamespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("partnerNamespaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, partnerNamespaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}
