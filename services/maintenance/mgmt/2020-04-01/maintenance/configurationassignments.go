package maintenance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ConfigurationAssignmentsClient is the maintenance Client
type ConfigurationAssignmentsClient struct {
	BaseClient
}

// NewConfigurationAssignmentsClient creates an instance of the ConfigurationAssignmentsClient client.
func NewConfigurationAssignmentsClient(subscriptionID string) ConfigurationAssignmentsClient {
	return NewConfigurationAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationAssignmentsClientWithBaseURI creates an instance of the ConfigurationAssignmentsClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewConfigurationAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationAssignmentsClient {
	return ConfigurationAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate register configuration for resource.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
// configurationAssignmentName - configuration assignment name
// configurationAssignment - the configurationAssignment
func (client ConfigurationAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment) (result ConfigurationAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, configurationAssignment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConfigurationAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationAssignmentName": autorest.Encode("path", configurationAssignmentName),
		"providerName":                autorest.Encode("path", providerName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"resourceName":                autorest.Encode("path", resourceName),
		"resourceType":                autorest.Encode("path", resourceType),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}", pathParameters),
		autorest.WithJSON(configurationAssignment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result ConfigurationAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateParent register configuration for resource.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// resourceType - resource type
// resourceName - resource identifier
// configurationAssignmentName - configuration assignment name
// configurationAssignment - the configurationAssignment
func (client ConfigurationAssignmentsClient) CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment) (result ConfigurationAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsClient.CreateOrUpdateParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateParentPreparer(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, configurationAssignment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "CreateOrUpdateParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "CreateOrUpdateParent", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "CreateOrUpdateParent", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdateParentPreparer prepares the CreateOrUpdateParent request.
func (client ConfigurationAssignmentsClient) CreateOrUpdateParentPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationAssignmentName": autorest.Encode("path", configurationAssignmentName),
		"providerName":                autorest.Encode("path", providerName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"resourceName":                autorest.Encode("path", resourceName),
		"resourceParentName":          autorest.Encode("path", resourceParentName),
		"resourceParentType":          autorest.Encode("path", resourceParentType),
		"resourceType":                autorest.Encode("path", resourceType),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}", pathParameters),
		autorest.WithJSON(configurationAssignment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateParentSender sends the CreateOrUpdateParent request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsClient) CreateOrUpdateParentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateParentResponder handles the response to the CreateOrUpdateParent request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsClient) CreateOrUpdateParentResponder(resp *http.Response) (result ConfigurationAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete unregister configuration for resource.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
// configurationAssignmentName - unique configuration assignment name
func (client ConfigurationAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string) (result ConfigurationAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConfigurationAssignmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationAssignmentName": autorest.Encode("path", configurationAssignmentName),
		"providerName":                autorest.Encode("path", providerName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"resourceName":                autorest.Encode("path", resourceName),
		"resourceType":                autorest.Encode("path", resourceType),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsClient) DeleteResponder(resp *http.Response) (result ConfigurationAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteParent unregister configuration for resource.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// resourceType - resource type
// resourceName - resource identifier
// configurationAssignmentName - unique configuration assignment name
func (client ConfigurationAssignmentsClient) DeleteParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string) (result ConfigurationAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsClient.DeleteParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteParentPreparer(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "DeleteParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "DeleteParent", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "DeleteParent", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteParentPreparer prepares the DeleteParent request.
func (client ConfigurationAssignmentsClient) DeleteParentPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationAssignmentName": autorest.Encode("path", configurationAssignmentName),
		"providerName":                autorest.Encode("path", providerName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"resourceName":                autorest.Encode("path", resourceName),
		"resourceParentName":          autorest.Encode("path", resourceParentName),
		"resourceParentType":          autorest.Encode("path", resourceParentType),
		"resourceType":                autorest.Encode("path", resourceType),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteParentSender sends the DeleteParent request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsClient) DeleteParentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteParentResponder handles the response to the DeleteParent request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsClient) DeleteParentResponder(resp *http.Response) (result ConfigurationAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list configurationAssignments for resource.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
func (client ConfigurationAssignmentsClient) List(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result ListConfigurationAssignmentsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, providerName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ConfigurationAssignmentsClient) ListPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsClient) ListResponder(resp *http.Response) (result ListConfigurationAssignmentsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListParent list configurationAssignments for resource.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// resourceType - resource type
// resourceName - resource identifier
func (client ConfigurationAssignmentsClient) ListParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result ListConfigurationAssignmentsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationAssignmentsClient.ListParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListParentPreparer(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "ListParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "ListParent", resp, "Failure sending request")
		return
	}

	result, err = client.ListParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ConfigurationAssignmentsClient", "ListParent", resp, "Failure responding to request")
		return
	}

	return
}

// ListParentPreparer prepares the ListParent request.
func (client ConfigurationAssignmentsClient) ListParentPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":       autorest.Encode("path", providerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"resourceName":       autorest.Encode("path", resourceName),
		"resourceParentName": autorest.Encode("path", resourceParentName),
		"resourceParentType": autorest.Encode("path", resourceParentType),
		"resourceType":       autorest.Encode("path", resourceType),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListParentSender sends the ListParent request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationAssignmentsClient) ListParentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListParentResponder handles the response to the ListParent request. The method always
// closes the http.Response Body.
func (client ConfigurationAssignmentsClient) ListParentResponder(resp *http.Response) (result ListConfigurationAssignmentsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
