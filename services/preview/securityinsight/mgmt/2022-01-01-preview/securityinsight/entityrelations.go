package securityinsight

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EntityRelationsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type EntityRelationsClient struct {
	BaseClient
}

// NewEntityRelationsClient creates an instance of the EntityRelationsClient client.
func NewEntityRelationsClient(subscriptionID string) EntityRelationsClient {
	return NewEntityRelationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEntityRelationsClientWithBaseURI creates an instance of the EntityRelationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEntityRelationsClientWithBaseURI(baseURI string, subscriptionID string) EntityRelationsClient {
	return EntityRelationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetRelation gets an entity relation.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// entityID - entity ID
// relationName - relation Name
func (client EntityRelationsClient) GetRelation(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, relationName string) (result Relation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntityRelationsClient.GetRelation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.EntityRelationsClient", "GetRelation", err.Error())
	}

	req, err := client.GetRelationPreparer(ctx, resourceGroupName, workspaceName, entityID, relationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityRelationsClient", "GetRelation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRelationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.EntityRelationsClient", "GetRelation", resp, "Failure sending request")
		return
	}

	result, err = client.GetRelationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntityRelationsClient", "GetRelation", resp, "Failure responding to request")
		return
	}

	return
}

// GetRelationPreparer prepares the GetRelation request.
func (client EntityRelationsClient) GetRelationPreparer(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, relationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"entityId":          autorest.Encode("path", entityID),
		"relationName":      autorest.Encode("path", relationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/relations/{relationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRelationSender sends the GetRelation request. The method will close the
// http.Response Body if it receives an error.
func (client EntityRelationsClient) GetRelationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetRelationResponder handles the response to the GetRelation request. The method always
// closes the http.Response Body.
func (client EntityRelationsClient) GetRelationResponder(resp *http.Response) (result Relation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
