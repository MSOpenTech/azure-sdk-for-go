package managementgroups

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

// EntitiesClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type EntitiesClient struct {
	BaseClient
}

// NewEntitiesClient creates an instance of the EntitiesClient client.
func NewEntitiesClient() EntitiesClient {
	return NewEntitiesClientWithBaseURI(DefaultBaseURI)
}

// NewEntitiesClientWithBaseURI creates an instance of the EntitiesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEntitiesClientWithBaseURI(baseURI string) EntitiesClient {
	return EntitiesClient{NewWithBaseURI(baseURI)}
}

// List list all entities (Management Groups, Subscriptions, etc.) for the authenticated user.
// Parameters:
// skiptoken - page continuation token is only used if a previous operation returned a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a token
// parameter that specifies a starting point to use for subsequent calls.
// skip - number of entities to skip over when retrieving results. Passing this in will override $skipToken.
// top - number of elements to return when retrieving results. Passing this in will override $skipToken.
// selectParameter - this parameter specifies the fields to include in the response. Can include any
// combination of Name,DisplayName,Type,ParentDisplayNameChain,ParentChain, e.g.
// '$select=Name,DisplayName,Type,ParentDisplayNameChain,ParentNameChain'. When specified the $select parameter
// can override select in $skipToken.
// search - the $search parameter is used in conjunction with the $filter parameter to return three different
// outputs depending on the parameter passed in.
// With $search=AllowedParents the API will return the entity info of all groups that the requested entity will
// be able to reparent to as determined by the user's permissions.
// With $search=AllowedChildren the API will return the entity info of all entities that can be added as
// children of the requested entity.
// With $search=ParentAndFirstLevelChildren the API will return the parent and  first level of children that
// the user has either direct access to or indirect access via one of their descendants.
// With $search=ParentOnly the API will return only the group if the user has access to at least one of the
// descendants of the group.
// With $search=ChildrenOnly the API will return only the first level of children of the group entity info
// specified in $filter.  The user must have direct access to the children entities or one of it's descendants
// for it to show up in the results.
// filter - the filter parameter allows you to filter on the the name or display name fields. You can check for
// equality on the name field (e.g. name eq '{entityName}')  and you can check for substrings on either the
// name or display name fields(e.g. contains(name, '{substringToSearch}'), contains(displayName,
// '{substringToSearch')). Note that the '{entityName}' and '{substringToSearch}' fields are checked case
// insensitively.
// view - the view parameter allows clients to filter the type of data that is returned by the getEntities
// call.
// groupName - a filter which allows the get entities call to focus on a particular group (i.e. "$filter=name
// eq 'groupName'")
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client EntitiesClient) List(ctx context.Context, skiptoken string, skip *int32, top *int32, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (result EntityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntitiesClient.List")
		defer func() {
			sc := -1
			if result.elr.Response.Response != nil {
				sc = result.elr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, skiptoken, skip, top, selectParameter, search, filter, view, groupName, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.elr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", resp, "Failure sending request")
		return
	}

	result.elr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.elr.hasNextLink() && result.elr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EntitiesClient) ListPreparer(ctx context.Context, skiptoken string, skip *int32, top *int32, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (*http.Request, error) {
	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(string(search)) > 0 {
		queryParameters["$search"] = autorest.Encode("query", search)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(string(view)) > 0 {
		queryParameters["$view"] = autorest.Encode("query", view)
	}
	if len(groupName) > 0 {
		queryParameters["groupName"] = autorest.Encode("query", groupName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/getEntities"),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EntitiesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EntitiesClient) ListResponder(resp *http.Response) (result EntityListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EntitiesClient) listNextResults(ctx context.Context, lastResults EntityListResult) (result EntityListResult, err error) {
	req, err := lastResults.entityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EntitiesClient) ListComplete(ctx context.Context, skiptoken string, skip *int32, top *int32, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (result EntityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntitiesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, skiptoken, skip, top, selectParameter, search, filter, view, groupName, cacheControl)
	return
}
