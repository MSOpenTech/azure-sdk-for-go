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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
	"net/http"
	"net/url"
	"regexp"
)

// DataTransferJobsServer is a fake server for instances of the armcosmos.DataTransferJobsClient type.
type DataTransferJobsServer struct {
	// Cancel is the fake for method DataTransferJobsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *armcosmos.DataTransferJobsClientCancelOptions) (resp azfake.Responder[armcosmos.DataTransferJobsClientCancelResponse], errResp azfake.ErrorResponder)

	// Complete is the fake for method DataTransferJobsClient.Complete
	// HTTP status codes to indicate success: http.StatusOK
	Complete func(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *armcosmos.DataTransferJobsClientCompleteOptions) (resp azfake.Responder[armcosmos.DataTransferJobsClientCompleteResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method DataTransferJobsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, accountName string, jobName string, jobCreateParameters armcosmos.CreateJobRequest, options *armcosmos.DataTransferJobsClientCreateOptions) (resp azfake.Responder[armcosmos.DataTransferJobsClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DataTransferJobsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *armcosmos.DataTransferJobsClientGetOptions) (resp azfake.Responder[armcosmos.DataTransferJobsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabaseAccountPager is the fake for method DataTransferJobsClient.NewListByDatabaseAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabaseAccountPager func(resourceGroupName string, accountName string, options *armcosmos.DataTransferJobsClientListByDatabaseAccountOptions) (resp azfake.PagerResponder[armcosmos.DataTransferJobsClientListByDatabaseAccountResponse])

	// Pause is the fake for method DataTransferJobsClient.Pause
	// HTTP status codes to indicate success: http.StatusOK
	Pause func(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *armcosmos.DataTransferJobsClientPauseOptions) (resp azfake.Responder[armcosmos.DataTransferJobsClientPauseResponse], errResp azfake.ErrorResponder)

	// Resume is the fake for method DataTransferJobsClient.Resume
	// HTTP status codes to indicate success: http.StatusOK
	Resume func(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *armcosmos.DataTransferJobsClientResumeOptions) (resp azfake.Responder[armcosmos.DataTransferJobsClientResumeResponse], errResp azfake.ErrorResponder)
}

// NewDataTransferJobsServerTransport creates a new instance of DataTransferJobsServerTransport with the provided implementation.
// The returned DataTransferJobsServerTransport instance is connected to an instance of armcosmos.DataTransferJobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDataTransferJobsServerTransport(srv *DataTransferJobsServer) *DataTransferJobsServerTransport {
	return &DataTransferJobsServerTransport{
		srv:                           srv,
		newListByDatabaseAccountPager: newTracker[azfake.PagerResponder[armcosmos.DataTransferJobsClientListByDatabaseAccountResponse]](),
	}
}

// DataTransferJobsServerTransport connects instances of armcosmos.DataTransferJobsClient to instances of DataTransferJobsServer.
// Don't use this type directly, use NewDataTransferJobsServerTransport instead.
type DataTransferJobsServerTransport struct {
	srv                           *DataTransferJobsServer
	newListByDatabaseAccountPager *tracker[azfake.PagerResponder[armcosmos.DataTransferJobsClientListByDatabaseAccountResponse]]
}

// Do implements the policy.Transporter interface for DataTransferJobsServerTransport.
func (d *DataTransferJobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DataTransferJobsClient.Cancel":
		resp, err = d.dispatchCancel(req)
	case "DataTransferJobsClient.Complete":
		resp, err = d.dispatchComplete(req)
	case "DataTransferJobsClient.Create":
		resp, err = d.dispatchCreate(req)
	case "DataTransferJobsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DataTransferJobsClient.NewListByDatabaseAccountPager":
		resp, err = d.dispatchNewListByDatabaseAccountPager(req)
	case "DataTransferJobsClient.Pause":
		resp, err = d.dispatchPause(req)
	case "DataTransferJobsClient.Resume":
		resp, err = d.dispatchResume(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if d.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Cancel(req.Context(), resourceGroupNameParam, accountNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataTransferJobGetResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchComplete(req *http.Request) (*http.Response, error) {
	if d.srv.Complete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Complete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/complete`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Complete(req.Context(), resourceGroupNameParam, accountNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataTransferJobGetResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if d.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcosmos.CreateJobRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, jobNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataTransferJobGetResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataTransferJobGetResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchNewListByDatabaseAccountPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByDatabaseAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabaseAccountPager not implemented")}
	}
	newListByDatabaseAccountPager := d.newListByDatabaseAccountPager.get(req)
	if newListByDatabaseAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByDatabaseAccountPager(resourceGroupNameParam, accountNameParam, nil)
		newListByDatabaseAccountPager = &resp
		d.newListByDatabaseAccountPager.add(req, newListByDatabaseAccountPager)
		server.PagerResponderInjectNextLinks(newListByDatabaseAccountPager, req, func(page *armcosmos.DataTransferJobsClientListByDatabaseAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabaseAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByDatabaseAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabaseAccountPager) {
		d.newListByDatabaseAccountPager.remove(req)
	}
	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchPause(req *http.Request) (*http.Response, error) {
	if d.srv.Pause == nil {
		return nil, &nonRetriableError{errors.New("fake for method Pause not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pause`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Pause(req.Context(), resourceGroupNameParam, accountNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataTransferJobGetResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataTransferJobsServerTransport) dispatchResume(req *http.Request) (*http.Response, error) {
	if d.srv.Resume == nil {
		return nil, &nonRetriableError{errors.New("fake for method Resume not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataTransferJobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resume`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Resume(req.Context(), resourceGroupNameParam, accountNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataTransferJobGetResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
