//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworkRulesClient contains the methods for the VirtualNetworkRules group.
// Don't use this type directly, use NewVirtualNetworkRulesClient() instead.
type VirtualNetworkRulesClient struct {
	con            *connection
	subscriptionID string
}

// NewVirtualNetworkRulesClient creates a new instance of VirtualNetworkRulesClient with the specified values.
func NewVirtualNetworkRulesClient(con *connection, subscriptionID string) *VirtualNetworkRulesClient {
	return &VirtualNetworkRulesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates the specified virtual network rule. During update, the virtual network rule with the specified name will be replaced
// with this new virtual network rule.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, parameters CreateOrUpdateVirtualNetworkRuleParameters, options *VirtualNetworkRulesCreateOrUpdateOptions) (VirtualNetworkRulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, parameters, options)
	if err != nil {
		return VirtualNetworkRulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, parameters CreateOrUpdateVirtualNetworkRuleParameters, options *VirtualNetworkRulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualNetworkRulesClient) createOrUpdateHandleResponse(resp *http.Response) (VirtualNetworkRulesCreateOrUpdateResponse, error) {
	result := VirtualNetworkRulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes the specified virtual network rule from the specified Data Lake Store account.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesDeleteOptions) (VirtualNetworkRulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return VirtualNetworkRulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return VirtualNetworkRulesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworkRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the specified Data Lake Store virtual network rule.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) Get(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesGetOptions) (VirtualNetworkRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkRulesGetResponse, error) {
	result := VirtualNetworkRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworkRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByAccount - Lists the Data Lake Store virtual network rules within the specified Data Lake Store account.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) ListByAccount(resourceGroupName string, accountName string, options *VirtualNetworkRulesListByAccountOptions) *VirtualNetworkRulesListByAccountPager {
	return &VirtualNetworkRulesListByAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworkRulesListByAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkRuleListResult.NextLink)
		},
	}
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *VirtualNetworkRulesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *VirtualNetworkRulesListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *VirtualNetworkRulesClient) listByAccountHandleResponse(resp *http.Response) (VirtualNetworkRulesListByAccountResponse, error) {
	result := VirtualNetworkRulesListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRuleListResult); err != nil {
		return VirtualNetworkRulesListByAccountResponse{}, err
	}
	return result, nil
}

// listByAccountHandleError handles the ListByAccount error response.
func (client *VirtualNetworkRulesClient) listByAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Updates the specified virtual network rule.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) Update(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesUpdateOptions) (VirtualNetworkRulesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworkRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *VirtualNetworkRulesClient) updateHandleResponse(resp *http.Response) (VirtualNetworkRulesUpdateResponse, error) {
	result := VirtualNetworkRulesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *VirtualNetworkRulesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
