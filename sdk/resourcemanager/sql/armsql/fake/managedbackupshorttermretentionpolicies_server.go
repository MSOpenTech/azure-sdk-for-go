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

// ManagedBackupShortTermRetentionPoliciesServer is a fake server for instances of the armsql.ManagedBackupShortTermRetentionPoliciesClient type.
type ManagedBackupShortTermRetentionPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedBackupShortTermRetentionPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName armsql.ManagedShortTermRetentionPolicyName, parameters armsql.ManagedBackupShortTermRetentionPolicy, options *armsql.ManagedBackupShortTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedBackupShortTermRetentionPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName armsql.ManagedShortTermRetentionPolicyName, options *armsql.ManagedBackupShortTermRetentionPoliciesClientGetOptions) (resp azfake.Responder[armsql.ManagedBackupShortTermRetentionPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method ManagedBackupShortTermRetentionPoliciesClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, managedInstanceName string, databaseName string, options *armsql.ManagedBackupShortTermRetentionPoliciesClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse])

	// BeginUpdate is the fake for method ManagedBackupShortTermRetentionPoliciesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName armsql.ManagedShortTermRetentionPolicyName, parameters armsql.ManagedBackupShortTermRetentionPolicy, options *armsql.ManagedBackupShortTermRetentionPoliciesClientBeginUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewManagedBackupShortTermRetentionPoliciesServerTransport creates a new instance of ManagedBackupShortTermRetentionPoliciesServerTransport with the provided implementation.
// The returned ManagedBackupShortTermRetentionPoliciesServerTransport instance is connected to an instance of armsql.ManagedBackupShortTermRetentionPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedBackupShortTermRetentionPoliciesServerTransport(srv *ManagedBackupShortTermRetentionPoliciesServer) *ManagedBackupShortTermRetentionPoliciesServerTransport {
	return &ManagedBackupShortTermRetentionPoliciesServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientCreateOrUpdateResponse]](),
		newListByDatabasePager: newTracker[azfake.PagerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse]](),
		beginUpdate:            newTracker[azfake.PollerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientUpdateResponse]](),
	}
}

// ManagedBackupShortTermRetentionPoliciesServerTransport connects instances of armsql.ManagedBackupShortTermRetentionPoliciesClient to instances of ManagedBackupShortTermRetentionPoliciesServer.
// Don't use this type directly, use NewManagedBackupShortTermRetentionPoliciesServerTransport instead.
type ManagedBackupShortTermRetentionPoliciesServerTransport struct {
	srv                    *ManagedBackupShortTermRetentionPoliciesServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientCreateOrUpdateResponse]]
	newListByDatabasePager *tracker[azfake.PagerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse]]
	beginUpdate            *tracker[azfake.PollerResponder[armsql.ManagedBackupShortTermRetentionPoliciesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ManagedBackupShortTermRetentionPoliciesServerTransport.
func (m *ManagedBackupShortTermRetentionPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedBackupShortTermRetentionPoliciesClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "ManagedBackupShortTermRetentionPoliciesClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedBackupShortTermRetentionPoliciesClient.NewListByDatabasePager":
		resp, err = m.dispatchNewListByDatabasePager(req)
	case "ManagedBackupShortTermRetentionPoliciesClient.BeginUpdate":
		resp, err = m.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedBackupShortTermRetentionPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupShortTermRetentionPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedBackupShortTermRetentionPolicy](req)
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := parseWithCast(matches[regex.SubexpIndex("policyName")], func(v string) (armsql.ManagedShortTermRetentionPolicyName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.ManagedShortTermRetentionPolicyName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, policyNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedBackupShortTermRetentionPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupShortTermRetentionPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := parseWithCast(matches[regex.SubexpIndex("policyName")], func(v string) (armsql.ManagedShortTermRetentionPolicyName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.ManagedShortTermRetentionPolicyName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, policyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedBackupShortTermRetentionPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedBackupShortTermRetentionPoliciesServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := m.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupShortTermRetentionPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByDatabasePager(resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		m.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		m.newListByDatabasePager.remove(req)
	}
	return resp, nil
}

func (m *ManagedBackupShortTermRetentionPoliciesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupShortTermRetentionPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedBackupShortTermRetentionPolicy](req)
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := parseWithCast(matches[regex.SubexpIndex("policyName")], func(v string) (armsql.ManagedShortTermRetentionPolicyName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.ManagedShortTermRetentionPolicyName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, policyNameParam, body, nil)
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
