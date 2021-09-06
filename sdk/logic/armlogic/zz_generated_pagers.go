//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// IntegrationAccountAgreementsListPager provides operations for iterating over paged responses.
type IntegrationAccountAgreementsListPager struct {
	client    *IntegrationAccountAgreementsClient
	current   IntegrationAccountAgreementsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountAgreementsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountAgreementsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountAgreementsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountAgreementListResult.NextLink == nil || len(*p.current.IntegrationAccountAgreementListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountAgreementsListResponse page.
func (p *IntegrationAccountAgreementsListPager) PageResponse() IntegrationAccountAgreementsListResponse {
	return p.current
}

// IntegrationAccountCertificatesListPager provides operations for iterating over paged responses.
type IntegrationAccountCertificatesListPager struct {
	client    *IntegrationAccountCertificatesClient
	current   IntegrationAccountCertificatesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountCertificatesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountCertificatesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountCertificatesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountCertificateListResult.NextLink == nil || len(*p.current.IntegrationAccountCertificateListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountCertificatesListResponse page.
func (p *IntegrationAccountCertificatesListPager) PageResponse() IntegrationAccountCertificatesListResponse {
	return p.current
}

// IntegrationAccountMapsListPager provides operations for iterating over paged responses.
type IntegrationAccountMapsListPager struct {
	client    *IntegrationAccountMapsClient
	current   IntegrationAccountMapsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountMapsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountMapsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountMapsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountMapListResult.NextLink == nil || len(*p.current.IntegrationAccountMapListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountMapsListResponse page.
func (p *IntegrationAccountMapsListPager) PageResponse() IntegrationAccountMapsListResponse {
	return p.current
}

// IntegrationAccountPartnersListPager provides operations for iterating over paged responses.
type IntegrationAccountPartnersListPager struct {
	client    *IntegrationAccountPartnersClient
	current   IntegrationAccountPartnersListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountPartnersListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountPartnersListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountPartnersListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountPartnerListResult.NextLink == nil || len(*p.current.IntegrationAccountPartnerListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountPartnersListResponse page.
func (p *IntegrationAccountPartnersListPager) PageResponse() IntegrationAccountPartnersListResponse {
	return p.current
}

// IntegrationAccountSchemasListPager provides operations for iterating over paged responses.
type IntegrationAccountSchemasListPager struct {
	client    *IntegrationAccountSchemasClient
	current   IntegrationAccountSchemasListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountSchemasListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountSchemasListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountSchemasListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountSchemaListResult.NextLink == nil || len(*p.current.IntegrationAccountSchemaListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountSchemasListResponse page.
func (p *IntegrationAccountSchemasListPager) PageResponse() IntegrationAccountSchemasListResponse {
	return p.current
}

// IntegrationAccountSessionsListPager provides operations for iterating over paged responses.
type IntegrationAccountSessionsListPager struct {
	client    *IntegrationAccountSessionsClient
	current   IntegrationAccountSessionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountSessionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountSessionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountSessionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountSessionListResult.NextLink == nil || len(*p.current.IntegrationAccountSessionListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountSessionsListResponse page.
func (p *IntegrationAccountSessionsListPager) PageResponse() IntegrationAccountSessionsListResponse {
	return p.current
}

// IntegrationAccountsListByResourceGroupPager provides operations for iterating over paged responses.
type IntegrationAccountsListByResourceGroupPager struct {
	client    *IntegrationAccountsClient
	current   IntegrationAccountsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountListResult.NextLink == nil || len(*p.current.IntegrationAccountListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationAccountsListByResourceGroupResponse page.
func (p *IntegrationAccountsListByResourceGroupPager) PageResponse() IntegrationAccountsListByResourceGroupResponse {
	return p.current
}

// IntegrationAccountsListBySubscriptionPager provides operations for iterating over paged responses.
type IntegrationAccountsListBySubscriptionPager struct {
	client    *IntegrationAccountsClient
	current   IntegrationAccountsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationAccountsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationAccountsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationAccountsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationAccountListResult.NextLink == nil || len(*p.current.IntegrationAccountListResult.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IntegrationAccountsListBySubscriptionResponse page.
func (p *IntegrationAccountsListBySubscriptionPager) PageResponse() IntegrationAccountsListBySubscriptionResponse {
	return p.current
}

// IntegrationServiceEnvironmentManagedAPIOperationsListPager provides operations for iterating over paged responses.
type IntegrationServiceEnvironmentManagedAPIOperationsListPager struct {
	client    *IntegrationServiceEnvironmentManagedAPIOperationsClient
	current   IntegrationServiceEnvironmentManagedAPIOperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationServiceEnvironmentManagedAPIOperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationServiceEnvironmentManagedAPIOperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationServiceEnvironmentManagedAPIOperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.APIOperationListResult.NextLink == nil || len(*p.current.APIOperationListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationServiceEnvironmentManagedAPIOperationsListResponse page.
func (p *IntegrationServiceEnvironmentManagedAPIOperationsListPager) PageResponse() IntegrationServiceEnvironmentManagedAPIOperationsListResponse {
	return p.current
}

// IntegrationServiceEnvironmentManagedApisListPager provides operations for iterating over paged responses.
type IntegrationServiceEnvironmentManagedApisListPager struct {
	client    *IntegrationServiceEnvironmentManagedApisClient
	current   IntegrationServiceEnvironmentManagedApisListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationServiceEnvironmentManagedApisListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationServiceEnvironmentManagedApisListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationServiceEnvironmentManagedApisListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedAPIListResult.NextLink == nil || len(*p.current.ManagedAPIListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationServiceEnvironmentManagedApisListResponse page.
func (p *IntegrationServiceEnvironmentManagedApisListPager) PageResponse() IntegrationServiceEnvironmentManagedApisListResponse {
	return p.current
}

// IntegrationServiceEnvironmentSKUsListPager provides operations for iterating over paged responses.
type IntegrationServiceEnvironmentSKUsListPager struct {
	client    *IntegrationServiceEnvironmentSKUsClient
	current   IntegrationServiceEnvironmentSKUsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationServiceEnvironmentSKUsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationServiceEnvironmentSKUsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationServiceEnvironmentSKUsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationServiceEnvironmentSKUList.NextLink == nil || len(*p.current.IntegrationServiceEnvironmentSKUList.NextLink) == 0 {
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

// PageResponse returns the current IntegrationServiceEnvironmentSKUsListResponse page.
func (p *IntegrationServiceEnvironmentSKUsListPager) PageResponse() IntegrationServiceEnvironmentSKUsListResponse {
	return p.current
}

// IntegrationServiceEnvironmentsListByResourceGroupPager provides operations for iterating over paged responses.
type IntegrationServiceEnvironmentsListByResourceGroupPager struct {
	client    *IntegrationServiceEnvironmentsClient
	current   IntegrationServiceEnvironmentsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationServiceEnvironmentsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationServiceEnvironmentsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationServiceEnvironmentsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationServiceEnvironmentListResult.NextLink == nil || len(*p.current.IntegrationServiceEnvironmentListResult.NextLink) == 0 {
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

// PageResponse returns the current IntegrationServiceEnvironmentsListByResourceGroupResponse page.
func (p *IntegrationServiceEnvironmentsListByResourceGroupPager) PageResponse() IntegrationServiceEnvironmentsListByResourceGroupResponse {
	return p.current
}

// IntegrationServiceEnvironmentsListBySubscriptionPager provides operations for iterating over paged responses.
type IntegrationServiceEnvironmentsListBySubscriptionPager struct {
	client    *IntegrationServiceEnvironmentsClient
	current   IntegrationServiceEnvironmentsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IntegrationServiceEnvironmentsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IntegrationServiceEnvironmentsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IntegrationServiceEnvironmentsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IntegrationServiceEnvironmentListResult.NextLink == nil || len(*p.current.IntegrationServiceEnvironmentListResult.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IntegrationServiceEnvironmentsListBySubscriptionResponse page.
func (p *IntegrationServiceEnvironmentsListBySubscriptionPager) PageResponse() IntegrationServiceEnvironmentsListBySubscriptionResponse {
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

// WorkflowRunActionRepetitionsRequestHistoriesListPager provides operations for iterating over paged responses.
type WorkflowRunActionRepetitionsRequestHistoriesListPager struct {
	client    *WorkflowRunActionRepetitionsRequestHistoriesClient
	current   WorkflowRunActionRepetitionsRequestHistoriesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowRunActionRepetitionsRequestHistoriesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowRunActionRepetitionsRequestHistoriesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowRunActionRepetitionsRequestHistoriesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RequestHistoryListResult.NextLink == nil || len(*p.current.RequestHistoryListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowRunActionRepetitionsRequestHistoriesListResponse page.
func (p *WorkflowRunActionRepetitionsRequestHistoriesListPager) PageResponse() WorkflowRunActionRepetitionsRequestHistoriesListResponse {
	return p.current
}

// WorkflowRunActionRequestHistoriesListPager provides operations for iterating over paged responses.
type WorkflowRunActionRequestHistoriesListPager struct {
	client    *WorkflowRunActionRequestHistoriesClient
	current   WorkflowRunActionRequestHistoriesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowRunActionRequestHistoriesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowRunActionRequestHistoriesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowRunActionRequestHistoriesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RequestHistoryListResult.NextLink == nil || len(*p.current.RequestHistoryListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowRunActionRequestHistoriesListResponse page.
func (p *WorkflowRunActionRequestHistoriesListPager) PageResponse() WorkflowRunActionRequestHistoriesListResponse {
	return p.current
}

// WorkflowRunActionsListPager provides operations for iterating over paged responses.
type WorkflowRunActionsListPager struct {
	client    *WorkflowRunActionsClient
	current   WorkflowRunActionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowRunActionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowRunActionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowRunActionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowRunActionListResult.NextLink == nil || len(*p.current.WorkflowRunActionListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowRunActionsListResponse page.
func (p *WorkflowRunActionsListPager) PageResponse() WorkflowRunActionsListResponse {
	return p.current
}

// WorkflowRunsListPager provides operations for iterating over paged responses.
type WorkflowRunsListPager struct {
	client    *WorkflowRunsClient
	current   WorkflowRunsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowRunsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowRunsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowRunsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowRunListResult.NextLink == nil || len(*p.current.WorkflowRunListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowRunsListResponse page.
func (p *WorkflowRunsListPager) PageResponse() WorkflowRunsListResponse {
	return p.current
}

// WorkflowTriggerHistoriesListPager provides operations for iterating over paged responses.
type WorkflowTriggerHistoriesListPager struct {
	client    *WorkflowTriggerHistoriesClient
	current   WorkflowTriggerHistoriesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowTriggerHistoriesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowTriggerHistoriesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowTriggerHistoriesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowTriggerHistoryListResult.NextLink == nil || len(*p.current.WorkflowTriggerHistoryListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowTriggerHistoriesListResponse page.
func (p *WorkflowTriggerHistoriesListPager) PageResponse() WorkflowTriggerHistoriesListResponse {
	return p.current
}

// WorkflowTriggersListPager provides operations for iterating over paged responses.
type WorkflowTriggersListPager struct {
	client    *WorkflowTriggersClient
	current   WorkflowTriggersListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowTriggersListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowTriggersListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowTriggersListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowTriggerListResult.NextLink == nil || len(*p.current.WorkflowTriggerListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowTriggersListResponse page.
func (p *WorkflowTriggersListPager) PageResponse() WorkflowTriggersListResponse {
	return p.current
}

// WorkflowVersionsListPager provides operations for iterating over paged responses.
type WorkflowVersionsListPager struct {
	client    *WorkflowVersionsClient
	current   WorkflowVersionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowVersionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowVersionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowVersionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowVersionListResult.NextLink == nil || len(*p.current.WorkflowVersionListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowVersionsListResponse page.
func (p *WorkflowVersionsListPager) PageResponse() WorkflowVersionsListResponse {
	return p.current
}

// WorkflowsListByResourceGroupPager provides operations for iterating over paged responses.
type WorkflowsListByResourceGroupPager struct {
	client    *WorkflowsClient
	current   WorkflowsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowListResult.NextLink == nil || len(*p.current.WorkflowListResult.NextLink) == 0 {
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

// PageResponse returns the current WorkflowsListByResourceGroupResponse page.
func (p *WorkflowsListByResourceGroupPager) PageResponse() WorkflowsListByResourceGroupResponse {
	return p.current
}

// WorkflowsListBySubscriptionPager provides operations for iterating over paged responses.
type WorkflowsListBySubscriptionPager struct {
	client    *WorkflowsClient
	current   WorkflowsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WorkflowsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WorkflowsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WorkflowsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkflowListResult.NextLink == nil || len(*p.current.WorkflowListResult.NextLink) == 0 {
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
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WorkflowsListBySubscriptionResponse page.
func (p *WorkflowsListBySubscriptionPager) PageResponse() WorkflowsListBySubscriptionResponse {
	return p.current
}
