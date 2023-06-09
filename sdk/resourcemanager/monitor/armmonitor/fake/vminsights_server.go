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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// VMInsightsServer is a fake server for instances of the armmonitor.VMInsightsClient type.
type VMInsightsServer struct {
	// GetOnboardingStatus is the fake for method VMInsightsClient.GetOnboardingStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetOnboardingStatus func(ctx context.Context, resourceURI string, options *armmonitor.VMInsightsClientGetOnboardingStatusOptions) (resp azfake.Responder[armmonitor.VMInsightsClientGetOnboardingStatusResponse], errResp azfake.ErrorResponder)
}

// NewVMInsightsServerTransport creates a new instance of VMInsightsServerTransport with the provided implementation.
// The returned VMInsightsServerTransport instance is connected to an instance of armmonitor.VMInsightsClient by way of the
// undefined.Transporter field.
func NewVMInsightsServerTransport(srv *VMInsightsServer) *VMInsightsServerTransport {
	return &VMInsightsServerTransport{srv: srv}
}

// VMInsightsServerTransport connects instances of armmonitor.VMInsightsClient to instances of VMInsightsServer.
// Don't use this type directly, use NewVMInsightsServerTransport instead.
type VMInsightsServerTransport struct {
	srv *VMInsightsServer
}

// Do implements the policy.Transporter interface for VMInsightsServerTransport.
func (v *VMInsightsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VMInsightsClient.GetOnboardingStatus":
		resp, err = v.dispatchGetOnboardingStatus(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VMInsightsServerTransport) dispatchGetOnboardingStatus(req *http.Request) (*http.Response, error) {
	if v.srv.GetOnboardingStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOnboardingStatus not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetOnboardingStatus(req.Context(), resourceURIUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VMInsightsOnboardingStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
