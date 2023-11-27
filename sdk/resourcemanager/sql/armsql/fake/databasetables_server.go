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
)

// DatabaseTablesServer is a fake server for instances of the armsql.DatabaseTablesClient type.
type DatabaseTablesServer struct {
	// Get is the fake for method DatabaseTablesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *armsql.DatabaseTablesClientGetOptions) (resp azfake.Responder[armsql.DatabaseTablesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySchemaPager is the fake for method DatabaseTablesClient.NewListBySchemaPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySchemaPager func(resourceGroupName string, serverName string, databaseName string, schemaName string, options *armsql.DatabaseTablesClientListBySchemaOptions) (resp azfake.PagerResponder[armsql.DatabaseTablesClientListBySchemaResponse])
}

// NewDatabaseTablesServerTransport creates a new instance of DatabaseTablesServerTransport with the provided implementation.
// The returned DatabaseTablesServerTransport instance is connected to an instance of armsql.DatabaseTablesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatabaseTablesServerTransport(srv *DatabaseTablesServer) *DatabaseTablesServerTransport {
	return &DatabaseTablesServerTransport{
		srv:                  srv,
		newListBySchemaPager: newTracker[azfake.PagerResponder[armsql.DatabaseTablesClientListBySchemaResponse]](),
	}
}

// DatabaseTablesServerTransport connects instances of armsql.DatabaseTablesClient to instances of DatabaseTablesServer.
// Don't use this type directly, use NewDatabaseTablesServerTransport instead.
type DatabaseTablesServerTransport struct {
	srv                  *DatabaseTablesServer
	newListBySchemaPager *tracker[azfake.PagerResponder[armsql.DatabaseTablesClientListBySchemaResponse]]
}

// Do implements the policy.Transporter interface for DatabaseTablesServerTransport.
func (d *DatabaseTablesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DatabaseTablesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DatabaseTablesClient.NewListBySchemaPager":
		resp, err = d.dispatchNewListBySchemaPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DatabaseTablesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
	if err != nil {
		return nil, err
	}
	tableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, tableNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatabaseTable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatabaseTablesServerTransport) dispatchNewListBySchemaPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySchemaPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySchemaPager not implemented")}
	}
	newListBySchemaPager := d.newListBySchemaPager.get(req)
	if newListBySchemaPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		schemaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.DatabaseTablesClientListBySchemaOptions
		if filterParam != nil {
			options = &armsql.DatabaseTablesClientListBySchemaOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListBySchemaPager(resourceGroupNameParam, serverNameParam, databaseNameParam, schemaNameParam, options)
		newListBySchemaPager = &resp
		d.newListBySchemaPager.add(req, newListBySchemaPager)
		server.PagerResponderInjectNextLinks(newListBySchemaPager, req, func(page *armsql.DatabaseTablesClientListBySchemaResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySchemaPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBySchemaPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySchemaPager) {
		d.newListBySchemaPager.remove(req)
	}
	return resp, nil
}
