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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// BackupOperationResultsServer is a fake server for instances of the armrecoveryservicesbackup.BackupOperationResultsClient type.
type BackupOperationResultsServer struct {
	// Get is the fake for method BackupOperationResultsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	Get func(ctx context.Context, vaultName string, resourceGroupName string, operationID string, options *armrecoveryservicesbackup.BackupOperationResultsClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.BackupOperationResultsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewBackupOperationResultsServerTransport creates a new instance of BackupOperationResultsServerTransport with the provided implementation.
// The returned BackupOperationResultsServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupOperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupOperationResultsServerTransport(srv *BackupOperationResultsServer) *BackupOperationResultsServerTransport {
	return &BackupOperationResultsServerTransport{srv: srv}
}

// BackupOperationResultsServerTransport connects instances of armrecoveryservicesbackup.BackupOperationResultsClient to instances of BackupOperationResultsServer.
// Don't use this type directly, use NewBackupOperationResultsServerTransport instead.
type BackupOperationResultsServerTransport struct {
	srv *BackupOperationResultsServer
}

// Do implements the policy.Transporter interface for BackupOperationResultsServerTransport.
func (b *BackupOperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BackupOperationResultsClient.Get":
		resp, err = b.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BackupOperationResultsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupOperationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
