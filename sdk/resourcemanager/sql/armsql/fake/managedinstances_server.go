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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ManagedInstancesServer is a fake server for instances of the armsql.ManagedInstancesClient type.
type ManagedInstancesServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedInstancesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters armsql.ManagedInstance, options *armsql.ManagedInstancesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedInstancesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ManagedInstancesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, managedInstanceName string, options *armsql.ManagedInstancesClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.ManagedInstancesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFailover is the fake for method ManagedInstancesClient.BeginFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailover func(ctx context.Context, resourceGroupName string, managedInstanceName string, options *armsql.ManagedInstancesClientBeginFailoverOptions) (resp azfake.PollerResponder[armsql.ManagedInstancesClientFailoverResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, options *armsql.ManagedInstancesClientGetOptions) (resp azfake.Responder[armsql.ManagedInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagedInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsql.ManagedInstancesClientListOptions) (resp azfake.PagerResponder[armsql.ManagedInstancesClientListResponse])

	// NewListByInstancePoolPager is the fake for method ManagedInstancesClient.NewListByInstancePoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInstancePoolPager func(resourceGroupName string, instancePoolName string, options *armsql.ManagedInstancesClientListByInstancePoolOptions) (resp azfake.PagerResponder[armsql.ManagedInstancesClientListByInstancePoolResponse])

	// NewListByManagedInstancePager is the fake for method ManagedInstancesClient.NewListByManagedInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagedInstancePager func(resourceGroupName string, managedInstanceName string, options *armsql.ManagedInstancesClientListByManagedInstanceOptions) (resp azfake.PagerResponder[armsql.ManagedInstancesClientListByManagedInstanceResponse])

	// NewListByResourceGroupPager is the fake for method ManagedInstancesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsql.ManagedInstancesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsql.ManagedInstancesClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method ManagedInstancesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters armsql.ManagedInstanceUpdate, options *armsql.ManagedInstancesClientBeginUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedInstancesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewManagedInstancesServerTransport creates a new instance of ManagedInstancesServerTransport with the provided implementation.
// The returned ManagedInstancesServerTransport instance is connected to an instance of armsql.ManagedInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedInstancesServerTransport(srv *ManagedInstancesServer) *ManagedInstancesServerTransport {
	return &ManagedInstancesServerTransport{
		srv:                           srv,
		beginCreateOrUpdate:           newTracker[azfake.PollerResponder[armsql.ManagedInstancesClientCreateOrUpdateResponse]](),
		beginDelete:                   newTracker[azfake.PollerResponder[armsql.ManagedInstancesClientDeleteResponse]](),
		beginFailover:                 newTracker[azfake.PollerResponder[armsql.ManagedInstancesClientFailoverResponse]](),
		newListPager:                  newTracker[azfake.PagerResponder[armsql.ManagedInstancesClientListResponse]](),
		newListByInstancePoolPager:    newTracker[azfake.PagerResponder[armsql.ManagedInstancesClientListByInstancePoolResponse]](),
		newListByManagedInstancePager: newTracker[azfake.PagerResponder[armsql.ManagedInstancesClientListByManagedInstanceResponse]](),
		newListByResourceGroupPager:   newTracker[azfake.PagerResponder[armsql.ManagedInstancesClientListByResourceGroupResponse]](),
		beginUpdate:                   newTracker[azfake.PollerResponder[armsql.ManagedInstancesClientUpdateResponse]](),
	}
}

// ManagedInstancesServerTransport connects instances of armsql.ManagedInstancesClient to instances of ManagedInstancesServer.
// Don't use this type directly, use NewManagedInstancesServerTransport instead.
type ManagedInstancesServerTransport struct {
	srv                           *ManagedInstancesServer
	beginCreateOrUpdate           *tracker[azfake.PollerResponder[armsql.ManagedInstancesClientCreateOrUpdateResponse]]
	beginDelete                   *tracker[azfake.PollerResponder[armsql.ManagedInstancesClientDeleteResponse]]
	beginFailover                 *tracker[azfake.PollerResponder[armsql.ManagedInstancesClientFailoverResponse]]
	newListPager                  *tracker[azfake.PagerResponder[armsql.ManagedInstancesClientListResponse]]
	newListByInstancePoolPager    *tracker[azfake.PagerResponder[armsql.ManagedInstancesClientListByInstancePoolResponse]]
	newListByManagedInstancePager *tracker[azfake.PagerResponder[armsql.ManagedInstancesClientListByManagedInstanceResponse]]
	newListByResourceGroupPager   *tracker[azfake.PagerResponder[armsql.ManagedInstancesClientListByResourceGroupResponse]]
	beginUpdate                   *tracker[azfake.PollerResponder[armsql.ManagedInstancesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ManagedInstancesServerTransport.
func (m *ManagedInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedInstancesClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "ManagedInstancesClient.BeginDelete":
		resp, err = m.dispatchBeginDelete(req)
	case "ManagedInstancesClient.BeginFailover":
		resp, err = m.dispatchBeginFailover(req)
	case "ManagedInstancesClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedInstancesClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	case "ManagedInstancesClient.NewListByInstancePoolPager":
		resp, err = m.dispatchNewListByInstancePoolPager(req)
	case "ManagedInstancesClient.NewListByManagedInstancePager":
		resp, err = m.dispatchNewListByManagedInstancePager(req)
	case "ManagedInstancesClient.NewListByResourceGroupPager":
		resp, err = m.dispatchNewListByResourceGroupPager(req)
	case "ManagedInstancesClient.BeginUpdate":
		resp, err = m.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedInstance](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, managedInstanceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchBeginFailover(req *http.Request) (*http.Response, error) {
	if m.srv.BeginFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFailover not implemented")}
	}
	beginFailover := m.beginFailover.get(req)
	if beginFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failover`
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
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		replicaTypeUnescaped, err := url.QueryUnescape(qp.Get("replicaType"))
		if err != nil {
			return nil, err
		}
		replicaTypeParam := getOptional(armsql.ReplicaType(replicaTypeUnescaped))
		var options *armsql.ManagedInstancesClientBeginFailoverOptions
		if replicaTypeParam != nil {
			options = &armsql.ManagedInstancesClientBeginFailoverOptions{
				ReplicaType: replicaTypeParam,
			}
		}
		respr, errRespr := m.srv.BeginFailover(req.Context(), resourceGroupNameParam, managedInstanceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFailover = &respr
		m.beginFailover.add(req, beginFailover)
	}

	resp, err := server.PollerResponderNext(beginFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFailover) {
		m.beginFailover.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armsql.ManagedInstancesClientGetOptions
	if expandParam != nil {
		options = &armsql.ManagedInstancesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armsql.ManagedInstancesClientListOptions
		if expandParam != nil {
			options = &armsql.ManagedInstancesClientListOptions{
				Expand: expandParam,
			}
		}
		resp := m.srv.NewListPager(options)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsql.ManagedInstancesClientListResponse, createLink func() string) {
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

func (m *ManagedInstancesServerTransport) dispatchNewListByInstancePoolPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByInstancePoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInstancePoolPager not implemented")}
	}
	newListByInstancePoolPager := m.newListByInstancePoolPager.get(req)
	if newListByInstancePoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/instancePools/(?P<instancePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedInstances`
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
		instancePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("instancePoolName")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armsql.ManagedInstancesClientListByInstancePoolOptions
		if expandParam != nil {
			options = &armsql.ManagedInstancesClientListByInstancePoolOptions{
				Expand: expandParam,
			}
		}
		resp := m.srv.NewListByInstancePoolPager(resourceGroupNameParam, instancePoolNameParam, options)
		newListByInstancePoolPager = &resp
		m.newListByInstancePoolPager.add(req, newListByInstancePoolPager)
		server.PagerResponderInjectNextLinks(newListByInstancePoolPager, req, func(page *armsql.ManagedInstancesClientListByInstancePoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInstancePoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByInstancePoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInstancePoolPager) {
		m.newListByInstancePoolPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchNewListByManagedInstancePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByManagedInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagedInstancePager not implemented")}
	}
	newListByManagedInstancePager := m.newListByManagedInstancePager.get(req)
	if newListByManagedInstancePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topqueries`
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
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		numberOfQueriesUnescaped, err := url.QueryUnescape(qp.Get("numberOfQueries"))
		if err != nil {
			return nil, err
		}
		numberOfQueriesParam, err := parseOptional(numberOfQueriesUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		databasesUnescaped, err := url.QueryUnescape(qp.Get("databases"))
		if err != nil {
			return nil, err
		}
		databasesParam := getOptional(databasesUnescaped)
		startTimeUnescaped, err := url.QueryUnescape(qp.Get("startTime"))
		if err != nil {
			return nil, err
		}
		startTimeParam := getOptional(startTimeUnescaped)
		endTimeUnescaped, err := url.QueryUnescape(qp.Get("endTime"))
		if err != nil {
			return nil, err
		}
		endTimeParam := getOptional(endTimeUnescaped)
		intervalUnescaped, err := url.QueryUnescape(qp.Get("interval"))
		if err != nil {
			return nil, err
		}
		intervalParam := getOptional(armsql.QueryTimeGrainType(intervalUnescaped))
		aggregationFunctionUnescaped, err := url.QueryUnescape(qp.Get("aggregationFunction"))
		if err != nil {
			return nil, err
		}
		aggregationFunctionParam := getOptional(armsql.AggregationFunctionType(aggregationFunctionUnescaped))
		observationMetricUnescaped, err := url.QueryUnescape(qp.Get("observationMetric"))
		if err != nil {
			return nil, err
		}
		observationMetricParam := getOptional(armsql.MetricType(observationMetricUnescaped))
		var options *armsql.ManagedInstancesClientListByManagedInstanceOptions
		if numberOfQueriesParam != nil || databasesParam != nil || startTimeParam != nil || endTimeParam != nil || intervalParam != nil || aggregationFunctionParam != nil || observationMetricParam != nil {
			options = &armsql.ManagedInstancesClientListByManagedInstanceOptions{
				NumberOfQueries:     numberOfQueriesParam,
				Databases:           databasesParam,
				StartTime:           startTimeParam,
				EndTime:             endTimeParam,
				Interval:            intervalParam,
				AggregationFunction: aggregationFunctionParam,
				ObservationMetric:   observationMetricParam,
			}
		}
		resp := m.srv.NewListByManagedInstancePager(resourceGroupNameParam, managedInstanceNameParam, options)
		newListByManagedInstancePager = &resp
		m.newListByManagedInstancePager.add(req, newListByManagedInstancePager)
		server.PagerResponderInjectNextLinks(newListByManagedInstancePager, req, func(page *armsql.ManagedInstancesClientListByManagedInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByManagedInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByManagedInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByManagedInstancePager) {
		m.newListByManagedInstancePager.remove(req)
	}
	return resp, nil
}

func (m *ManagedInstancesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances`
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
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armsql.ManagedInstancesClientListByResourceGroupOptions
		if expandParam != nil {
			options = &armsql.ManagedInstancesClientListByResourceGroupOptions{
				Expand: expandParam,
			}
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsql.ManagedInstancesClientListByResourceGroupResponse, createLink func() string) {
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

func (m *ManagedInstancesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedInstanceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		m.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		m.beginUpdate.remove(req)
	}

	return resp, nil
}
