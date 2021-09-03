//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClustersListByResourceGroupPager provides operations for iterating over paged responses.
type ClustersListByResourceGroupPager struct {
	client    *ClustersClient
	current   ClustersListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterListResult.NextLink == nil || len(*p.current.ClusterListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersListByResourceGroupResponse page.
func (p *ClustersListByResourceGroupPager) PageResponse() ClustersListByResourceGroupResponse {
	return p.current
}

// ConsumerGroupsListByEventHubPager provides operations for iterating over paged responses.
type ConsumerGroupsListByEventHubPager struct {
	client    *ConsumerGroupsClient
	current   ConsumerGroupsListByEventHubResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConsumerGroupsListByEventHubResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConsumerGroupsListByEventHubPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConsumerGroupsListByEventHubPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ConsumerGroupListResult.NextLink == nil || len(*p.current.ConsumerGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByEventHubHandleError(resp)
		return false
	}
	result, err := p.client.listByEventHubHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConsumerGroupsListByEventHubResponse page.
func (p *ConsumerGroupsListByEventHubPager) PageResponse() ConsumerGroupsListByEventHubResponse {
	return p.current
}

// DisasterRecoveryConfigsListAuthorizationRulesPager provides operations for iterating over paged responses.
type DisasterRecoveryConfigsListAuthorizationRulesPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DisasterRecoveryConfigsListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DisasterRecoveryConfigsListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DisasterRecoveryConfigsListAuthorizationRulesResponse page.
func (p *DisasterRecoveryConfigsListAuthorizationRulesPager) PageResponse() DisasterRecoveryConfigsListAuthorizationRulesResponse {
	return p.current
}

// DisasterRecoveryConfigsListPager provides operations for iterating over paged responses.
type DisasterRecoveryConfigsListPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DisasterRecoveryConfigsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DisasterRecoveryConfigsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArmDisasterRecoveryListResult.NextLink == nil || len(*p.current.ArmDisasterRecoveryListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DisasterRecoveryConfigsListResponse page.
func (p *DisasterRecoveryConfigsListPager) PageResponse() DisasterRecoveryConfigsListResponse {
	return p.current
}

// EventHubsListAuthorizationRulesPager provides operations for iterating over paged responses.
type EventHubsListAuthorizationRulesPager struct {
	client    *EventHubsClient
	current   EventHubsListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EventHubsListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EventHubsListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EventHubsListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current EventHubsListAuthorizationRulesResponse page.
func (p *EventHubsListAuthorizationRulesPager) PageResponse() EventHubsListAuthorizationRulesResponse {
	return p.current
}

// EventHubsListByNamespacePager provides operations for iterating over paged responses.
type EventHubsListByNamespacePager struct {
	client    *EventHubsClient
	current   EventHubsListByNamespaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EventHubsListByNamespaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EventHubsListByNamespacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EventHubsListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EventHubListResult.NextLink == nil || len(*p.current.EventHubListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByNamespaceHandleError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current EventHubsListByNamespaceResponse page.
func (p *EventHubsListByNamespacePager) PageResponse() EventHubsListByNamespaceResponse {
	return p.current
}

// NamespacesListAuthorizationRulesPager provides operations for iterating over paged responses.
type NamespacesListAuthorizationRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesListAuthorizationRulesResponse page.
func (p *NamespacesListAuthorizationRulesPager) PageResponse() NamespacesListAuthorizationRulesResponse {
	return p.current
}

// NamespacesListByResourceGroupPager provides operations for iterating over paged responses.
type NamespacesListByResourceGroupPager struct {
	client    *NamespacesClient
	current   NamespacesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EHNamespaceListResult.NextLink == nil || len(*p.current.EHNamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesListByResourceGroupResponse page.
func (p *NamespacesListByResourceGroupPager) PageResponse() NamespacesListByResourceGroupResponse {
	return p.current
}

// NamespacesListIPFilterRulesPager provides operations for iterating over paged responses.
type NamespacesListIPFilterRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListIPFilterRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesListIPFilterRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesListIPFilterRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesListIPFilterRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IPFilterRuleListResult.NextLink == nil || len(*p.current.IPFilterRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listIPFilterRulesHandleError(resp)
		return false
	}
	result, err := p.client.listIPFilterRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesListIPFilterRulesResponse page.
func (p *NamespacesListIPFilterRulesPager) PageResponse() NamespacesListIPFilterRulesResponse {
	return p.current
}

// NamespacesListPager provides operations for iterating over paged responses.
type NamespacesListPager struct {
	client    *NamespacesClient
	current   NamespacesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EHNamespaceListResult.NextLink == nil || len(*p.current.EHNamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesListResponse page.
func (p *NamespacesListPager) PageResponse() NamespacesListResponse {
	return p.current
}

// NamespacesListVirtualNetworkRulesPager provides operations for iterating over paged responses.
type NamespacesListVirtualNetworkRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListVirtualNetworkRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesListVirtualNetworkRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesListVirtualNetworkRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesListVirtualNetworkRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkRuleListResult.NextLink == nil || len(*p.current.VirtualNetworkRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listVirtualNetworkRulesHandleError(resp)
		return false
	}
	result, err := p.client.listVirtualNetworkRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesListVirtualNetworkRulesResponse page.
func (p *NamespacesListVirtualNetworkRulesPager) PageResponse() NamespacesListVirtualNetworkRulesResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// PrivateEndpointConnectionsListPager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsListPager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsListResponse page.
func (p *PrivateEndpointConnectionsListPager) PageResponse() PrivateEndpointConnectionsListResponse {
	return p.current
}

// RegionsListBySKUPager provides operations for iterating over paged responses.
type RegionsListBySKUPager struct {
	client    *RegionsClient
	current   RegionsListBySKUResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RegionsListBySKUResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RegionsListBySKUPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RegionsListBySKUPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MessagingRegionsListResult.NextLink == nil || len(*p.current.MessagingRegionsListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySKUHandleError(resp)
		return false
	}
	result, err := p.client.listBySKUHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RegionsListBySKUResponse page.
func (p *RegionsListBySKUPager) PageResponse() RegionsListBySKUResponse {
	return p.current
}
