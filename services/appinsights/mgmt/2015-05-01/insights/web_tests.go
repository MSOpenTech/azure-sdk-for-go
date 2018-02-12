package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// WebTestsClient is the composite Swagger for Application Insights Management Client
type WebTestsClient struct {
	ManagementClient
}

// NewWebTestsClient creates an instance of the WebTestsClient client.
func NewWebTestsClient(p pipeline.Pipeline) WebTestsClient {
	return WebTestsClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates an Application Insights web test definition.
//
// resourceGroupName is the name of the resource group. webTestName is the name of the Application Insights webtest
// resource. webTestDefinition is properties that need to be specified to create or update an Application Insights web
// test definition.
func (client WebTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest) (*WebTest, error) {
	if err := validate([]validation{
		{targetValue: webTestDefinition,
			constraints: []constraint{{target: "webTestDefinition.WebTestProperties", name: null, rule: false,
				chain: []constraint{{target: "webTestDefinition.WebTestProperties.SyntheticMonitorID", name: null, rule: true, chain: nil},
					{target: "webTestDefinition.WebTestProperties.WebTestName", name: null, rule: true, chain: nil},
					{target: "webTestDefinition.WebTestProperties.Locations", name: null, rule: true, chain: nil},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, webTestName, webTestDefinition)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebTest), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WebTestsClient) createOrUpdatePreparer(resourceGroupName string, webTestName string, webTestDefinition WebTest) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/webtests/{webTestName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(webTestDefinition)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client WebTestsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WebTest{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes an Application Insights web test.
//
// resourceGroupName is the name of the resource group. webTestName is the name of the Application Insights webtest
// resource.
func (client WebTestsClient) Delete(ctx context.Context, resourceGroupName string, webTestName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, webTestName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client WebTestsClient) deletePreparer(resourceGroupName string, webTestName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/webtests/{webTestName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client WebTestsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusNoContent, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get get a specific Application Insights web test definition.
//
// resourceGroupName is the name of the resource group. webTestName is the name of the Application Insights webtest
// resource.
func (client WebTestsClient) Get(ctx context.Context, resourceGroupName string, webTestName string) (*WebTest, error) {
	req, err := client.getPreparer(resourceGroupName, webTestName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebTest), err
}

// getPreparer prepares the Get request.
func (client WebTestsClient) getPreparer(resourceGroupName string, webTestName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/webtests/{webTestName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client WebTestsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WebTest{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List get all Application Insights web test alerts definitioned within a subscription.
func (client WebTestsClient) List(ctx context.Context) (*WebTestListResult, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebTestListResult), err
}

// listPreparer prepares the List request.
func (client WebTestsClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/microsoft.insights/webtests"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client WebTestsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WebTestListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListByResourceGroup get all Application Insights web tests defined within a specified resource group.
//
// resourceGroupName is the name of the resource group.
func (client WebTestsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*WebTestListResult, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebTestListResult), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WebTestsClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/webtests"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client WebTestsClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WebTestListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// UpdateTags creates or updates an Application Insights web test definition.
//
// resourceGroupName is the name of the resource group. webTestName is the name of the Application Insights webtest
// resource. webTestTags is updated tag information to set into the web test instance.
func (client WebTestsClient) UpdateTags(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource) (*WebTest, error) {
	req, err := client.updateTagsPreparer(resourceGroupName, webTestName, webTestTags)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateTagsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*WebTest), err
}

// updateTagsPreparer prepares the UpdateTags request.
func (client WebTestsClient) updateTagsPreparer(resourceGroupName string, webTestName string, webTestTags TagsResource) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/webtests/{webTestName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(webTestTags)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateTagsResponder handles the response to the UpdateTags request.
func (client WebTestsClient) updateTagsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &WebTest{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
