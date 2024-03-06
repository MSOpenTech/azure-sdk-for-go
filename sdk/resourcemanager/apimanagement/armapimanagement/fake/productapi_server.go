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

// ProductAPIServer is a fake server for instances of the armapimanagement.ProductAPIClient type.
type ProductAPIServer struct {
	// CheckEntityExists is the fake for method ProductAPIClient.CheckEntityExists
	// HTTP status codes to indicate success: http.StatusNoContent
	CheckEntityExists func(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *armapimanagement.ProductAPIClientCheckEntityExistsOptions) (resp azfake.Responder[armapimanagement.ProductAPIClientCheckEntityExistsResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method ProductAPIClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *armapimanagement.ProductAPIClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.ProductAPIClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProductAPIClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *armapimanagement.ProductAPIClientDeleteOptions) (resp azfake.Responder[armapimanagement.ProductAPIClientDeleteResponse], errResp azfake.ErrorResponder)

	// NewListByProductPager is the fake for method ProductAPIClient.NewListByProductPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProductPager func(resourceGroupName string, serviceName string, productID string, options *armapimanagement.ProductAPIClientListByProductOptions) (resp azfake.PagerResponder[armapimanagement.ProductAPIClientListByProductResponse])
}

// NewProductAPIServerTransport creates a new instance of ProductAPIServerTransport with the provided implementation.
// The returned ProductAPIServerTransport instance is connected to an instance of armapimanagement.ProductAPIClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProductAPIServerTransport(srv *ProductAPIServer) *ProductAPIServerTransport {
	return &ProductAPIServerTransport{
		srv:                   srv,
		newListByProductPager: newTracker[azfake.PagerResponder[armapimanagement.ProductAPIClientListByProductResponse]](),
	}
}

// ProductAPIServerTransport connects instances of armapimanagement.ProductAPIClient to instances of ProductAPIServer.
// Don't use this type directly, use NewProductAPIServerTransport instead.
type ProductAPIServerTransport struct {
	srv                   *ProductAPIServer
	newListByProductPager *tracker[azfake.PagerResponder[armapimanagement.ProductAPIClientListByProductResponse]]
}

// Do implements the policy.Transporter interface for ProductAPIServerTransport.
func (p *ProductAPIServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProductAPIClient.CheckEntityExists":
		resp, err = p.dispatchCheckEntityExists(req)
	case "ProductAPIClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "ProductAPIClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "ProductAPIClient.NewListByProductPager":
		resp, err = p.dispatchNewListByProductPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProductAPIServerTransport) dispatchCheckEntityExists(req *http.Request) (*http.Response, error) {
	if p.srv.CheckEntityExists == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckEntityExists not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CheckEntityExists(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, apiIDParam, nil)
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

func (p *ProductAPIServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, apiIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).APIContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProductAPIServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, apiIDParam, nil)
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

func (p *ProductAPIServerTransport) dispatchNewListByProductPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByProductPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProductPager not implemented")}
	}
	newListByProductPager := p.newListByProductPager.get(req)
	if newListByProductPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
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
		var options *armapimanagement.ProductAPIClientListByProductOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.ProductAPIClientListByProductOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := p.srv.NewListByProductPager(resourceGroupNameParam, serviceNameParam, productIDParam, options)
		newListByProductPager = &resp
		p.newListByProductPager.add(req, newListByProductPager)
		server.PagerResponderInjectNextLinks(newListByProductPager, req, func(page *armapimanagement.ProductAPIClientListByProductResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProductPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByProductPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProductPager) {
		p.newListByProductPager.remove(req)
	}
	return resp, nil
}
