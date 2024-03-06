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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
	"net/http"
	"net/url"
	"regexp"
)

// IscsiDisksServer is a fake server for instances of the armstorsimple1200series.IscsiDisksClient type.
type IscsiDisksServer struct {
	// BeginCreateOrUpdate is the fake for method IscsiDisksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, iscsiDisk armstorsimple1200series.ISCSIDisk, options *armstorsimple1200series.IscsiDisksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armstorsimple1200series.IscsiDisksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method IscsiDisksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiDisksClientBeginDeleteOptions) (resp azfake.PollerResponder[armstorsimple1200series.IscsiDisksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IscsiDisksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiDisksClientGetOptions) (resp azfake.Responder[armstorsimple1200series.IscsiDisksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDevicePager is the fake for method IscsiDisksClient.NewListByDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDevicePager func(deviceName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiDisksClientListByDeviceOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListByDeviceResponse])

	// NewListByIscsiServerPager is the fake for method IscsiDisksClient.NewListByIscsiServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByIscsiServerPager func(deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiDisksClientListByIscsiServerOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListByIscsiServerResponse])

	// NewListMetricDefinitionPager is the fake for method IscsiDisksClient.NewListMetricDefinitionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricDefinitionPager func(deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiDisksClientListMetricDefinitionOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListMetricDefinitionResponse])

	// NewListMetricsPager is the fake for method IscsiDisksClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiDisksClientListMetricsOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListMetricsResponse])
}

// NewIscsiDisksServerTransport creates a new instance of IscsiDisksServerTransport with the provided implementation.
// The returned IscsiDisksServerTransport instance is connected to an instance of armstorsimple1200series.IscsiDisksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIscsiDisksServerTransport(srv *IscsiDisksServer) *IscsiDisksServerTransport {
	return &IscsiDisksServerTransport{
		srv:                          srv,
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armstorsimple1200series.IscsiDisksClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armstorsimple1200series.IscsiDisksClientDeleteResponse]](),
		newListByDevicePager:         newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListByDeviceResponse]](),
		newListByIscsiServerPager:    newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListByIscsiServerResponse]](),
		newListMetricDefinitionPager: newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListMetricDefinitionResponse]](),
		newListMetricsPager:          newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListMetricsResponse]](),
	}
}

// IscsiDisksServerTransport connects instances of armstorsimple1200series.IscsiDisksClient to instances of IscsiDisksServer.
// Don't use this type directly, use NewIscsiDisksServerTransport instead.
type IscsiDisksServerTransport struct {
	srv                          *IscsiDisksServer
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armstorsimple1200series.IscsiDisksClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armstorsimple1200series.IscsiDisksClientDeleteResponse]]
	newListByDevicePager         *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListByDeviceResponse]]
	newListByIscsiServerPager    *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListByIscsiServerResponse]]
	newListMetricDefinitionPager *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListMetricDefinitionResponse]]
	newListMetricsPager          *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiDisksClientListMetricsResponse]]
}

// Do implements the policy.Transporter interface for IscsiDisksServerTransport.
func (i *IscsiDisksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IscsiDisksClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "IscsiDisksClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "IscsiDisksClient.Get":
		resp, err = i.dispatchGet(req)
	case "IscsiDisksClient.NewListByDevicePager":
		resp, err = i.dispatchNewListByDevicePager(req)
	case "IscsiDisksClient.NewListByIscsiServerPager":
		resp, err = i.dispatchNewListByIscsiServerPager(req)
	case "IscsiDisksClient.NewListMetricDefinitionPager":
		resp, err = i.dispatchNewListMetricDefinitionPager(req)
	case "IscsiDisksClient.NewListMetricsPager":
		resp, err = i.dispatchNewListMetricsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple1200series.ISCSIDisk](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		diskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, iscsiServerNameParam, diskNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		diskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), deviceNameParam, iscsiServerNameParam, diskNameParam, resourceGroupNameParam, managerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
	if err != nil {
		return nil, err
	}
	diskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), deviceNameParam, iscsiServerNameParam, diskNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ISCSIDisk, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchNewListByDevicePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDevicePager not implemented")}
	}
	newListByDevicePager := i.newListByDevicePager.get(req)
	if newListByDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByDevicePager(deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByDevicePager = &resp
		i.newListByDevicePager.add(req, newListByDevicePager)
	}
	resp, err := server.PagerResponderNext(newListByDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDevicePager) {
		i.newListByDevicePager.remove(req)
	}
	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchNewListByIscsiServerPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByIscsiServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByIscsiServerPager not implemented")}
	}
	newListByIscsiServerPager := i.newListByIscsiServerPager.get(req)
	if newListByIscsiServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByIscsiServerPager(deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByIscsiServerPager = &resp
		i.newListByIscsiServerPager.add(req, newListByIscsiServerPager)
	}
	resp, err := server.PagerResponderNext(newListByIscsiServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByIscsiServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByIscsiServerPager) {
		i.newListByIscsiServerPager.remove(req)
	}
	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchNewListMetricDefinitionPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListMetricDefinitionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricDefinitionPager not implemented")}
	}
	newListMetricDefinitionPager := i.newListMetricDefinitionPager.get(req)
	if newListMetricDefinitionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metricsDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		diskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListMetricDefinitionPager(deviceNameParam, iscsiServerNameParam, diskNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListMetricDefinitionPager = &resp
		i.newListMetricDefinitionPager.add(req, newListMetricDefinitionPager)
	}
	resp, err := server.PagerResponderNext(newListMetricDefinitionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListMetricDefinitionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricDefinitionPager) {
		i.newListMetricDefinitionPager.remove(req)
	}
	return resp, nil
}

func (i *IscsiDisksServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := i.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		diskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armstorsimple1200series.IscsiDisksClientListMetricsOptions
		if filterParam != nil {
			options = &armstorsimple1200series.IscsiDisksClientListMetricsOptions{
				Filter: filterParam,
			}
		}
		resp := i.srv.NewListMetricsPager(deviceNameParam, iscsiServerNameParam, diskNameParam, resourceGroupNameParam, managerNameParam, options)
		newListMetricsPager = &resp
		i.newListMetricsPager.add(req, newListMetricsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListMetricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricsPager) {
		i.newListMetricsPager.remove(req)
	}
	return resp, nil
}
