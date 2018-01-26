package consumption

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/shopspring/decimal"
	"net/http"
)

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// EqualTo ...
	EqualTo OperatorType = "EqualTo"
	// GreaterThan ...
	GreaterThan OperatorType = "GreaterThan"
	// GreaterThanOrEqualTo ...
	GreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// TimeGrainType enumerates the values for time grain type.
type TimeGrainType string

const (
	// Annually ...
	Annually TimeGrainType = "Annually"
	// Monthly ...
	Monthly TimeGrainType = "Monthly"
	// Quarterly ...
	Quarterly TimeGrainType = "Quarterly"
)

// Budget a budget resource.
type Budget struct {
	autorest.Response `json:"-"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// ETag - eTag of the resource. To handle concurrent update scenarion, this field will be used to determine whether the user is updating the latest version or not.
	ETag              *string `json:"eTag,omitempty"`
	*BudgetProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for Budget struct.
func (b *Budget) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties BudgetProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		b.BudgetProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		b.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		b.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		b.Type = &typeVar
	}

	v = m["eTag"]
	if v != nil {
		var eTag string
		err = json.Unmarshal(*m["eTag"], &eTag)
		if err != nil {
			return err
		}
		b.ETag = &eTag
	}

	return nil
}

// BudgetProperties the properties of the budget.
type BudgetProperties struct {
	// Category - The category of the budget, whether the budget tracks cost or something else.
	Category *string `json:"category,omitempty"`
	// Amount - The total amount of cost to track with the budget
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// TimeGrain - The time covered by a budget. Tracking of the amount will be reset based on the time grain. Possible values include: 'Monthly', 'Quarterly', 'Annually'
	TimeGrain TimeGrainType `json:"timeGrain,omitempty"`
	// TimePeriod - Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain preiod. There are no restrictions on the end date.
	TimePeriod *BudgetTimePeriod `json:"timePeriod,omitempty"`
	// CurrentSpend - The current amount of cost which is being tracked for a budget.
	CurrentSpend *CurrentSpend `json:"currentSpend,omitempty"`
	// Notifications - Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications *map[string]*Notification `json:"notifications,omitempty"`
}

// BudgetsListResult result of listing budgets. It contains a list of available budgets in the scope provided.
type BudgetsListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of budgets.
	Value *[]Budget `json:"value,omitempty"`
}

// BudgetTimePeriod the start and end date for a budget.
type BudgetTimePeriod struct {
	// StartDate - The start date for the budget.
	StartDate *date.Time `json:"startDate,omitempty"`
	// EndDate - The end date for the budget. If not provided, we default this to 10 years from the start date.
	EndDate *date.Time `json:"endDate,omitempty"`
}

// CurrentSpend the current amount of cost which is being tracked for a budget.
type CurrentSpend struct {
	// Amount - The total amount of cost which is being tracked by the budget.
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// Unit - The unit of measure for the budget amount.
	Unit *string `json:"unit,omitempty"`
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Notification the notification associated with a budget.
type Notification struct {
	// Enabled - The notification is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Operator - The comparison operator. Possible values include: 'EqualTo', 'GreaterThan', 'GreaterThanOrEqualTo'
	Operator OperatorType `json:"operator,omitempty"`
	// Threshold - Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *decimal.Decimal `json:"threshold,omitempty"`
	// ContactEmails - Email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails *[]string `json:"contactEmails,omitempty"`
	// ContactRoles - Contact roles to send the budget notification to when the threshold is exceeded.
	ContactRoles *[]string `json:"contactRoles,omitempty"`
}

// Operation a Consumption REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Consumption.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: UsageDetail, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of listing consumption operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of consumption operations supported by the Microsoft.Consumption resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
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

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
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

// ProxyResource the Resource model definition.
type ProxyResource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// ETag - eTag of the resource. To handle concurrent update scenarion, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `json:"eTag,omitempty"`
}
