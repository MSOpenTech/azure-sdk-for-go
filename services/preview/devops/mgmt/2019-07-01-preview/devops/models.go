package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/devops/mgmt/2019-07-01-preview/devops"

// Authorization authorization info used to access a resource (like code repository).
type Authorization struct {
	// AuthorizationType - Type of authorization.
	AuthorizationType *string `json:"authorizationType,omitempty"`
	// Parameters - Authorization parameters corresponding to the authorization type.
	Parameters map[string]*string `json:"parameters"`
}

// MarshalJSON is the custom marshaler for Authorization.
func (a Authorization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AuthorizationType != nil {
		objectMap["authorizationType"] = a.AuthorizationType
	}
	if a.Parameters != nil {
		objectMap["parameters"] = a.Parameters
	}
	return json.Marshal(objectMap)
}

// BootstrapConfiguration configuration used to bootstrap a Pipeline.
type BootstrapConfiguration struct {
	// Repository - Repository containing the source code for the pipeline.
	Repository *CodeRepository `json:"repository,omitempty"`
	// Template - Template used to bootstrap the pipeline.
	Template *PipelineTemplate `json:"template,omitempty"`
}

// CloudError an error response from the Pipelines Resource Provider.
type CloudError struct {
	// Error - Details of the error from the Pipelines Resource Provider.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from the Pipelines Resource Provider.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error or the method where the error occurred.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// CodeRepository repository containing the source code for a pipeline.
type CodeRepository struct {
	// RepositoryType - Type of code repository. Possible values include: 'GitHub', 'VstsGit'
	RepositoryType CodeRepositoryType `json:"repositoryType,omitempty"`
	// ID - Unique immutable identifier of the code repository.
	ID *string `json:"id,omitempty"`
	// DefaultBranch - Default branch used to configure Continuous Integration (CI) in the pipeline.
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	// Authorization - Authorization info to access the code repository.
	Authorization *Authorization `json:"authorization,omitempty"`
	// Properties - Repository-specific properties.
	Properties map[string]*string `json:"properties"`
}

// MarshalJSON is the custom marshaler for CodeRepository.
func (cr CodeRepository) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cr.RepositoryType != "" {
		objectMap["repositoryType"] = cr.RepositoryType
	}
	if cr.ID != nil {
		objectMap["id"] = cr.ID
	}
	if cr.DefaultBranch != nil {
		objectMap["defaultBranch"] = cr.DefaultBranch
	}
	if cr.Authorization != nil {
		objectMap["authorization"] = cr.Authorization
	}
	if cr.Properties != nil {
		objectMap["properties"] = cr.Properties
	}
	return json.Marshal(objectMap)
}

// InputDescriptor representation of a pipeline template input parameter.
type InputDescriptor struct {
	// ID - Identifier of the input parameter.
	ID *string `json:"id,omitempty"`
	// Description - Description of the input parameter.
	Description *string `json:"description,omitempty"`
	// Type - Data type of the value of the input parameter. Possible values include: 'InputDataTypeString', 'InputDataTypeSecureString', 'InputDataTypeInt', 'InputDataTypeBool', 'InputDataTypeAuthorization'
	Type InputDataType `json:"type,omitempty"`
	// PossibleValues - List of possible values for the input parameter.
	PossibleValues *[]InputValue `json:"possibleValues,omitempty"`
}

// InputValue representation of a pipeline template input parameter value.
type InputValue struct {
	// Value - Value of an input parameter.
	Value *string `json:"value,omitempty"`
	// DisplayValue - Description of the input parameter value.
	DisplayValue *string `json:"displayValue,omitempty"`
}

// Operation properties of an Operation.
type Operation struct {
	// Name - READ-ONLY; Name of the operation.
	Name *string `json:"name,omitempty"`
	// IsDataAction - Indicates whether the operation applies to data-plane.
	IsDataAction *string `json:"isDataAction,omitempty"`
	// OperationDisplayValue - Display information of the operation.
	*OperationDisplayValue `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.IsDataAction != nil {
		objectMap["isDataAction"] = o.IsDataAction
	}
	if o.OperationDisplayValue != nil {
		objectMap["display"] = o.OperationDisplayValue
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Operation struct.
func (o *Operation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				o.Name = &name
			}
		case "isDataAction":
			if v != nil {
				var isDataAction string
				err = json.Unmarshal(*v, &isDataAction)
				if err != nil {
					return err
				}
				o.IsDataAction = &isDataAction
			}
		case "display":
			if v != nil {
				var operationDisplayValue OperationDisplayValue
				err = json.Unmarshal(*v, &operationDisplayValue)
				if err != nil {
					return err
				}
				o.OperationDisplayValue = &operationDisplayValue
			}
		}
	}

	return nil
}

// OperationDisplayValue display information of an operation.
type OperationDisplayValue struct {
	// Operation - READ-ONLY; Friendly name of the operation.
	Operation *string `json:"operation,omitempty"`
	// Resource - READ-ONLY; Friendly name of the resource type the operation applies to.
	Resource *string `json:"resource,omitempty"`
	// Description - READ-ONLY; Friendly description of the operation.
	Description *string `json:"description,omitempty"`
	// Provider - READ-ONLY; Friendly name of the resource provider.
	Provider *string `json:"provider,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDisplayValue.
func (odv OperationDisplayValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// OperationListResult result of a request to list all operations supported by Microsoft.DevOps resource
// provider.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of operations supported by Microsoft.DevOps resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The URL to get the next set of operations, if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationListResult.
func (olr OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if olr.NextLink != nil {
		objectMap["nextLink"] = olr.NextLink
	}
	return json.Marshal(objectMap)
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// OrganizationReference reference to an Azure DevOps Organization.
type OrganizationReference struct {
	// ID - READ-ONLY; Unique immutable identifier for the Azure DevOps Organization.
	ID *string `json:"id,omitempty"`
	// Name - Name of the Azure DevOps Organization.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for OrganizationReference.
func (or OrganizationReference) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if or.Name != nil {
		objectMap["name"] = or.Name
	}
	return json.Marshal(objectMap)
}

// Pipeline azure DevOps Pipeline used to configure Continuous Integration (CI) & Continuous Delivery (CD)
// for Azure resources.
type Pipeline struct {
	autorest.Response `json:"-"`
	// PipelineProperties - Custom properties of the Pipeline.
	*PipelineProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Resource Type
	Type *string `json:"type,omitempty"`
	// Tags - Resource Tags
	Tags map[string]*string `json:"tags"`
	// Location - Resource Location
	Location *string `json:"location,omitempty"`
	// Name - READ-ONLY; Resource Name
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for Pipeline.
func (p Pipeline) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.PipelineProperties != nil {
		objectMap["properties"] = p.PipelineProperties
	}
	if p.Tags != nil {
		objectMap["tags"] = p.Tags
	}
	if p.Location != nil {
		objectMap["location"] = p.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Pipeline struct.
func (p *Pipeline) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var pipelineProperties PipelineProperties
				err = json.Unmarshal(*v, &pipelineProperties)
				if err != nil {
					return err
				}
				p.PipelineProperties = &pipelineProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				p.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				p.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				p.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				p.Location = &location
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				p.Name = &name
			}
		}
	}

	return nil
}

// PipelineListResult result of a request to list all Azure Pipelines under a given scope.
type PipelineListResult struct {
	autorest.Response `json:"-"`
	// Value - List of pipelines.
	Value *[]Pipeline `json:"value,omitempty"`
	// NextLink - URL to get the next set of Pipelines, if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// PipelineListResultIterator provides access to a complete listing of Pipeline values.
type PipelineListResultIterator struct {
	i    int
	page PipelineListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PipelineListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *PipelineListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PipelineListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PipelineListResultIterator) Response() PipelineListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PipelineListResultIterator) Value() Pipeline {
	if !iter.page.NotDone() {
		return Pipeline{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the PipelineListResultIterator type.
func NewPipelineListResultIterator(page PipelineListResultPage) PipelineListResultIterator {
	return PipelineListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (plr PipelineListResult) IsEmpty() bool {
	return plr.Value == nil || len(*plr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (plr PipelineListResult) hasNextLink() bool {
	return plr.NextLink != nil && len(*plr.NextLink) != 0
}

// pipelineListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (plr PipelineListResult) pipelineListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !plr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(plr.NextLink)))
}

// PipelineListResultPage contains a page of Pipeline values.
type PipelineListResultPage struct {
	fn  func(context.Context, PipelineListResult) (PipelineListResult, error)
	plr PipelineListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PipelineListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.plr)
		if err != nil {
			return err
		}
		page.plr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PipelineListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PipelineListResultPage) NotDone() bool {
	return !page.plr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PipelineListResultPage) Response() PipelineListResult {
	return page.plr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PipelineListResultPage) Values() []Pipeline {
	if page.plr.IsEmpty() {
		return nil
	}
	return *page.plr.Value
}

// Creates a new instance of the PipelineListResultPage type.
func NewPipelineListResultPage(cur PipelineListResult, getNextPage func(context.Context, PipelineListResult) (PipelineListResult, error)) PipelineListResultPage {
	return PipelineListResultPage{
		fn:  getNextPage,
		plr: cur,
	}
}

// PipelineProperties custom properties of a Pipeline.
type PipelineProperties struct {
	// PipelineID - READ-ONLY; Unique identifier of the Azure Pipeline within the Azure DevOps Project.
	PipelineID *int32 `json:"pipelineId,omitempty"`
	// Organization - Reference to the Azure DevOps Organization containing the Pipeline.
	Organization *OrganizationReference `json:"organization,omitempty"`
	// Project - Reference to the Azure DevOps Project containing the Pipeline.
	Project *ProjectReference `json:"project,omitempty"`
	// BootstrapConfiguration - Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration *BootstrapConfiguration `json:"bootstrapConfiguration,omitempty"`
}

// MarshalJSON is the custom marshaler for PipelineProperties.
func (pp PipelineProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pp.Organization != nil {
		objectMap["organization"] = pp.Organization
	}
	if pp.Project != nil {
		objectMap["project"] = pp.Project
	}
	if pp.BootstrapConfiguration != nil {
		objectMap["bootstrapConfiguration"] = pp.BootstrapConfiguration
	}
	return json.Marshal(objectMap)
}

// PipelinesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type PipelinesCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(PipelinesClient) (Pipeline, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *PipelinesCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for PipelinesCreateOrUpdateFuture.Result.
func (future *PipelinesCreateOrUpdateFuture) result(client PipelinesClient) (p Pipeline, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelinesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		p.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("devops.PipelinesCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if p.Response.Response, err = future.GetResult(sender); err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
		p, err = client.CreateOrUpdateResponder(p.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devops.PipelinesCreateOrUpdateFuture", "Result", p.Response.Response, "Failure responding to request")
		}
	}
	return
}

// PipelineTemplate template used to bootstrap the pipeline.
type PipelineTemplate struct {
	// ID - Unique identifier of the pipeline template.
	ID *string `json:"id,omitempty"`
	// Parameters - Dictionary of input parameters used in the pipeline template.
	Parameters map[string]*string `json:"parameters"`
}

// MarshalJSON is the custom marshaler for PipelineTemplate.
func (pt PipelineTemplate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pt.ID != nil {
		objectMap["id"] = pt.ID
	}
	if pt.Parameters != nil {
		objectMap["parameters"] = pt.Parameters
	}
	return json.Marshal(objectMap)
}

// PipelineTemplateDefinition definition of a pipeline template.
type PipelineTemplateDefinition struct {
	// ID - Unique identifier of the pipeline template.
	ID *string `json:"id,omitempty"`
	// Description - Description of the pipeline enabled by the template.
	Description *string `json:"description,omitempty"`
	// Inputs - List of input parameters required by the template to create a pipeline.
	Inputs *[]InputDescriptor `json:"inputs,omitempty"`
}

// PipelineTemplateDefinitionListResult result of a request to list all pipeline template definitions.
type PipelineTemplateDefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - List of pipeline template definitions.
	Value *[]PipelineTemplateDefinition `json:"value,omitempty"`
	// NextLink - The URL to get the next set of pipeline template definitions, if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// PipelineTemplateDefinitionListResultIterator provides access to a complete listing of
// PipelineTemplateDefinition values.
type PipelineTemplateDefinitionListResultIterator struct {
	i    int
	page PipelineTemplateDefinitionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PipelineTemplateDefinitionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineTemplateDefinitionListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *PipelineTemplateDefinitionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PipelineTemplateDefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PipelineTemplateDefinitionListResultIterator) Response() PipelineTemplateDefinitionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PipelineTemplateDefinitionListResultIterator) Value() PipelineTemplateDefinition {
	if !iter.page.NotDone() {
		return PipelineTemplateDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the PipelineTemplateDefinitionListResultIterator type.
func NewPipelineTemplateDefinitionListResultIterator(page PipelineTemplateDefinitionListResultPage) PipelineTemplateDefinitionListResultIterator {
	return PipelineTemplateDefinitionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ptdlr PipelineTemplateDefinitionListResult) IsEmpty() bool {
	return ptdlr.Value == nil || len(*ptdlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ptdlr PipelineTemplateDefinitionListResult) hasNextLink() bool {
	return ptdlr.NextLink != nil && len(*ptdlr.NextLink) != 0
}

// pipelineTemplateDefinitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ptdlr PipelineTemplateDefinitionListResult) pipelineTemplateDefinitionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !ptdlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ptdlr.NextLink)))
}

// PipelineTemplateDefinitionListResultPage contains a page of PipelineTemplateDefinition values.
type PipelineTemplateDefinitionListResultPage struct {
	fn    func(context.Context, PipelineTemplateDefinitionListResult) (PipelineTemplateDefinitionListResult, error)
	ptdlr PipelineTemplateDefinitionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PipelineTemplateDefinitionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineTemplateDefinitionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ptdlr)
		if err != nil {
			return err
		}
		page.ptdlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PipelineTemplateDefinitionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PipelineTemplateDefinitionListResultPage) NotDone() bool {
	return !page.ptdlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PipelineTemplateDefinitionListResultPage) Response() PipelineTemplateDefinitionListResult {
	return page.ptdlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PipelineTemplateDefinitionListResultPage) Values() []PipelineTemplateDefinition {
	if page.ptdlr.IsEmpty() {
		return nil
	}
	return *page.ptdlr.Value
}

// Creates a new instance of the PipelineTemplateDefinitionListResultPage type.
func NewPipelineTemplateDefinitionListResultPage(cur PipelineTemplateDefinitionListResult, getNextPage func(context.Context, PipelineTemplateDefinitionListResult) (PipelineTemplateDefinitionListResult, error)) PipelineTemplateDefinitionListResultPage {
	return PipelineTemplateDefinitionListResultPage{
		fn:    getNextPage,
		ptdlr: cur,
	}
}

// PipelineUpdateParameters request payload used to update an existing Azure Pipeline.
type PipelineUpdateParameters struct {
	// Tags - Dictionary of key-value pairs to be set as tags on the Azure Pipeline. This will overwrite any existing tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PipelineUpdateParameters.
func (pup PipelineUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pup.Tags != nil {
		objectMap["tags"] = pup.Tags
	}
	return json.Marshal(objectMap)
}

// ProjectReference reference to an Azure DevOps Project.
type ProjectReference struct {
	// ID - READ-ONLY; Unique immutable identifier of the Azure DevOps Project.
	ID *string `json:"id,omitempty"`
	// Name - Name of the Azure DevOps Project.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for ProjectReference.
func (pr ProjectReference) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.Name != nil {
		objectMap["name"] = pr.Name
	}
	return json.Marshal(objectMap)
}

// Resource an Azure Resource Manager (ARM) resource.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Resource Type
	Type *string `json:"type,omitempty"`
	// Tags - Resource Tags
	Tags map[string]*string `json:"tags"`
	// Location - Resource Location
	Location *string `json:"location,omitempty"`
	// Name - READ-ONLY; Resource Name
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	return json.Marshal(objectMap)
}
