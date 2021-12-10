//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvisualstudio

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AccountsCheckNameAvailabilityResponse contains the response from method Accounts.CheckNameAvailability.
type AccountsCheckNameAvailabilityResponse struct {
	AccountsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCheckNameAvailabilityResult contains the result from method Accounts.CheckNameAvailability.
type AccountsCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// AccountsCreateOrUpdateResponse contains the response from method Accounts.CreateOrUpdate.
type AccountsCreateOrUpdateResponse struct {
	AccountsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCreateOrUpdateResult contains the result from method Accounts.CreateOrUpdate.
type AccountsCreateOrUpdateResult struct {
	AccountResource
}

// AccountsDeleteResponse contains the response from method Accounts.Delete.
type AccountsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsGetResponse contains the response from method Accounts.Get.
type AccountsGetResponse struct {
	AccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsGetResult contains the result from method Accounts.Get.
type AccountsGetResult struct {
	AccountResource
}

// AccountsListByResourceGroupResponse contains the response from method Accounts.ListByResourceGroup.
type AccountsListByResourceGroupResponse struct {
	AccountsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsListByResourceGroupResult contains the result from method Accounts.ListByResourceGroup.
type AccountsListByResourceGroupResult struct {
	AccountResourceListResult
}

// AccountsUpdateResponse contains the response from method Accounts.Update.
type AccountsUpdateResponse struct {
	AccountsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsUpdateResult contains the result from method Accounts.Update.
type AccountsUpdateResult struct {
	AccountResource
}

// ExtensionsCreateResponse contains the response from method Extensions.Create.
type ExtensionsCreateResponse struct {
	ExtensionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsCreateResult contains the result from method Extensions.Create.
type ExtensionsCreateResult struct {
	ExtensionResource
}

// ExtensionsDeleteResponse contains the response from method Extensions.Delete.
type ExtensionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsGetResponse contains the response from method Extensions.Get.
type ExtensionsGetResponse struct {
	ExtensionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsGetResult contains the result from method Extensions.Get.
type ExtensionsGetResult struct {
	ExtensionResource
}

// ExtensionsListByAccountResponse contains the response from method Extensions.ListByAccount.
type ExtensionsListByAccountResponse struct {
	ExtensionsListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsListByAccountResult contains the result from method Extensions.ListByAccount.
type ExtensionsListByAccountResult struct {
	ExtensionResourceListResult
}

// ExtensionsUpdateResponse contains the response from method Extensions.Update.
type ExtensionsUpdateResponse struct {
	ExtensionsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExtensionsUpdateResult contains the result from method Extensions.Update.
type ExtensionsUpdateResult struct {
	ExtensionResource
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// ProjectsCreatePollerResponse contains the response from method Projects.Create.
type ProjectsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ProjectsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ProjectsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ProjectsCreateResponse, error) {
	respType := ProjectsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ProjectResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ProjectsCreatePollerResponse from the provided client and resume token.
func (l *ProjectsCreatePollerResponse) Resume(ctx context.Context, client *ProjectsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ProjectsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &ProjectsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ProjectsCreateResponse contains the response from method Projects.Create.
type ProjectsCreateResponse struct {
	ProjectsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProjectsCreateResult contains the result from method Projects.Create.
type ProjectsCreateResult struct {
	ProjectResource
}

// ProjectsGetJobStatusResponse contains the response from method Projects.GetJobStatus.
type ProjectsGetJobStatusResponse struct {
	ProjectsGetJobStatusResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProjectsGetJobStatusResult contains the result from method Projects.GetJobStatus.
type ProjectsGetJobStatusResult struct {
	ProjectResource
}

// ProjectsGetResponse contains the response from method Projects.Get.
type ProjectsGetResponse struct {
	ProjectsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProjectsGetResult contains the result from method Projects.Get.
type ProjectsGetResult struct {
	ProjectResource
}

// ProjectsListByResourceGroupResponse contains the response from method Projects.ListByResourceGroup.
type ProjectsListByResourceGroupResponse struct {
	ProjectsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProjectsListByResourceGroupResult contains the result from method Projects.ListByResourceGroup.
type ProjectsListByResourceGroupResult struct {
	ProjectResourceListResult
}

// ProjectsUpdateResponse contains the response from method Projects.Update.
type ProjectsUpdateResponse struct {
	ProjectsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ProjectsUpdateResult contains the result from method Projects.Update.
type ProjectsUpdateResult struct {
	ProjectResource
}
