// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armtrustedsigning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CertificateProfilesClient contains the methods for the CertificateProfiles group.
// Don't use this type directly, use NewCertificateProfilesClient() instead.
type CertificateProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCertificateProfilesClient creates a new instance of CertificateProfilesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificateProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateProfilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CertificateProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-05-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Trusted Signing account name.
//   - profileName - Certificate profile name.
//   - resource - Parameters to create the certificate profile
//   - options - CertificateProfilesClientBeginCreateOptions contains the optional parameters for the CertificateProfilesClient.BeginCreate
//     method.
func (client *CertificateProfilesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, profileName string, resource CertificateProfile, options *CertificateProfilesClientBeginCreateOptions) (*runtime.Poller[CertificateProfilesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, profileName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateProfilesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateProfilesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-05-preview
func (client *CertificateProfilesClient) create(ctx context.Context, resourceGroupName string, accountName string, profileName string, resource CertificateProfile, options *CertificateProfilesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateProfilesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, profileName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *CertificateProfilesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, profileName string, resource CertificateProfile, _ *CertificateProfilesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CodeSigning/codeSigningAccounts/{accountName}/certificateProfiles/{profileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-05-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Trusted Signing account name.
//   - profileName - Certificate profile name.
//   - options - CertificateProfilesClientBeginDeleteOptions contains the optional parameters for the CertificateProfilesClient.BeginDelete
//     method.
func (client *CertificateProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, profileName string, options *CertificateProfilesClientBeginDeleteOptions) (*runtime.Poller[CertificateProfilesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, profileName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-05-preview
func (client *CertificateProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, profileName string, options *CertificateProfilesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateProfilesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, profileName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificateProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, profileName string, _ *CertificateProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CodeSigning/codeSigningAccounts/{accountName}/certificateProfiles/{profileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get details of a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-05-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Trusted Signing account name.
//   - profileName - Certificate profile name.
//   - options - CertificateProfilesClientGetOptions contains the optional parameters for the CertificateProfilesClient.Get method.
func (client *CertificateProfilesClient) Get(ctx context.Context, resourceGroupName string, accountName string, profileName string, options *CertificateProfilesClientGetOptions) (CertificateProfilesClientGetResponse, error) {
	var err error
	const operationName = "CertificateProfilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, profileName, options)
	if err != nil {
		return CertificateProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CertificateProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, profileName string, _ *CertificateProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CodeSigning/codeSigningAccounts/{accountName}/certificateProfiles/{profileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificateProfilesClient) getHandleResponse(resp *http.Response) (CertificateProfilesClientGetResponse, error) {
	result := CertificateProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateProfile); err != nil {
		return CertificateProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByCodeSigningAccountPager - List certificate profiles under a trusted signing account.
//
// Generated from API version 2024-02-05-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Trusted Signing account name.
//   - options - CertificateProfilesClientListByCodeSigningAccountOptions contains the optional parameters for the CertificateProfilesClient.NewListByCodeSigningAccountPager
//     method.
func (client *CertificateProfilesClient) NewListByCodeSigningAccountPager(resourceGroupName string, accountName string, options *CertificateProfilesClientListByCodeSigningAccountOptions) *runtime.Pager[CertificateProfilesClientListByCodeSigningAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateProfilesClientListByCodeSigningAccountResponse]{
		More: func(page CertificateProfilesClientListByCodeSigningAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateProfilesClientListByCodeSigningAccountResponse) (CertificateProfilesClientListByCodeSigningAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificateProfilesClient.NewListByCodeSigningAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCodeSigningAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return CertificateProfilesClientListByCodeSigningAccountResponse{}, err
			}
			return client.listByCodeSigningAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCodeSigningAccountCreateRequest creates the ListByCodeSigningAccount request.
func (client *CertificateProfilesClient) listByCodeSigningAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *CertificateProfilesClientListByCodeSigningAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CodeSigning/codeSigningAccounts/{accountName}/certificateProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCodeSigningAccountHandleResponse handles the ListByCodeSigningAccount response.
func (client *CertificateProfilesClient) listByCodeSigningAccountHandleResponse(resp *http.Response) (CertificateProfilesClientListByCodeSigningAccountResponse, error) {
	result := CertificateProfilesClientListByCodeSigningAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateProfileListResult); err != nil {
		return CertificateProfilesClientListByCodeSigningAccountResponse{}, err
	}
	return result, nil
}

// RevokeCertificate - Revoke a certificate under a certificate profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-05-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Trusted Signing account name.
//   - profileName - Certificate profile name.
//   - body - Parameters to revoke the certificate profile
//   - options - CertificateProfilesClientRevokeCertificateOptions contains the optional parameters for the CertificateProfilesClient.RevokeCertificate
//     method.
func (client *CertificateProfilesClient) RevokeCertificate(ctx context.Context, resourceGroupName string, accountName string, profileName string, body RevokeCertificate, options *CertificateProfilesClientRevokeCertificateOptions) (CertificateProfilesClientRevokeCertificateResponse, error) {
	var err error
	const operationName = "CertificateProfilesClient.RevokeCertificate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revokeCertificateCreateRequest(ctx, resourceGroupName, accountName, profileName, body, options)
	if err != nil {
		return CertificateProfilesClientRevokeCertificateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateProfilesClientRevokeCertificateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CertificateProfilesClientRevokeCertificateResponse{}, err
	}
	return CertificateProfilesClientRevokeCertificateResponse{}, nil
}

// revokeCertificateCreateRequest creates the RevokeCertificate request.
func (client *CertificateProfilesClient) revokeCertificateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, profileName string, body RevokeCertificate, _ *CertificateProfilesClientRevokeCertificateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CodeSigning/codeSigningAccounts/{accountName}/certificateProfiles/{profileName}/revokeCertificate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
