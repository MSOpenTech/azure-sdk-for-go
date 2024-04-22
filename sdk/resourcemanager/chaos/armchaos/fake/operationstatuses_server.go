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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
	"net/http"
	"net/url"
	"regexp"
)

// OperationStatusesServer is a fake server for instances of the armchaos.OperationStatusesClient type.
type OperationStatusesServer struct {
	// Get is the fake for method OperationStatusesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, asyncOperationID string, options *armchaos.OperationStatusesClientGetOptions) (resp azfake.Responder[armchaos.OperationStatusesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationStatusesServerTransport creates a new instance of OperationStatusesServerTransport with the provided implementation.
// The returned OperationStatusesServerTransport instance is connected to an instance of armchaos.OperationStatusesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationStatusesServerTransport(srv *OperationStatusesServer) *OperationStatusesServerTransport {
	return &OperationStatusesServerTransport{srv: srv}
}

// OperationStatusesServerTransport connects instances of armchaos.OperationStatusesClient to instances of OperationStatusesServer.
// Don't use this type directly, use NewOperationStatusesServerTransport instead.
type OperationStatusesServerTransport struct {
	srv *OperationStatusesServer
}

// Do implements the policy.Transporter interface for OperationStatusesServerTransport.
func (o *OperationStatusesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationStatusesClient.Get":
		resp, err = o.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationStatusesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationStatuses/(?P<asyncOperationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	asyncOperationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("asyncOperationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), locationParam, asyncOperationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
