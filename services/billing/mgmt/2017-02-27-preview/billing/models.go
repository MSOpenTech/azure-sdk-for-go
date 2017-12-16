package billing

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DownloadURL a secure URL that can be used to download a PDF invoice until the URL expires.
type DownloadURL struct {
	// ExpiryTime - The time in UTC at which this download URL will expire.
	ExpiryTime *date.Time `json:"expiryTime,omitempty"`
	// URL - The URL to the PDF file.
	URL *string `json:"url,omitempty"`
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error.
	Target *string `json:"target,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponse struct {
	Error *ErrorDetails `json:"error,omitempty"`
}

// Invoice an invoice resource can be used download a PDF version of an invoice.
type Invoice struct {
	autorest.Response `json:"-"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type               *string `json:"type,omitempty"`
	*InvoiceProperties `json:"properties,omitempty"`
}

// InvoiceProperties the properties of the invoice.
type InvoiceProperties struct {
	// InvoicePeriodStartDate - The start of the date range covered by the invoice.
	InvoicePeriodStartDate *date.Date `json:"invoicePeriodStartDate,omitempty"`
	// InvoicePeriodEndDate - The end of the date range covered by the invoice.
	InvoicePeriodEndDate *date.Date `json:"invoicePeriodEndDate,omitempty"`
	// DownloadURL - A secure link to download the PDF version of an invoice. The link will cease to work after its expiry time is reached.
	DownloadURL *DownloadURL `json:"downloadUrl,omitempty"`
}

// InvoicesListResult result of the request to list invoices. It contains a list of available invoices in reverse
// chronological order.
type InvoicesListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of invoices.
	Value *[]Invoice `json:"value,omitempty"`
	// NextLink - the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// InvoicesListResultIterator provides access to a complete listing of Invoice values.
type InvoicesListResultIterator struct {
	i    int
	page InvoicesListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *InvoicesListResultIterator) Next() error {
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
func (iter InvoicesListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter InvoicesListResultIterator) Response() InvoicesListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter InvoicesListResultIterator) Value() Invoice {
	if !iter.page.NotDone() {
		return Invoice{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (ilr InvoicesListResult) IsEmpty() bool {
	return ilr.Value == nil || len(*ilr.Value) == 0
}

// invoicesListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ilr InvoicesListResult) invoicesListResultPreparer() (*http.Request, error) {
	if ilr.NextLink == nil || len(to.String(ilr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ilr.NextLink)))
}

// InvoicesListResultPage contains a page of Invoice values.
type InvoicesListResultPage struct {
	fn  func(InvoicesListResult) (InvoicesListResult, error)
	ilr InvoicesListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *InvoicesListResultPage) Next() error {
	next, err := page.fn(page.ilr)
	if err != nil {
		return err
	}
	page.ilr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page InvoicesListResultPage) NotDone() bool {
	return !page.ilr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page InvoicesListResultPage) Response() InvoicesListResult {
	return page.ilr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page InvoicesListResultPage) Values() []Invoice {
	if page.ilr.IsEmpty() {
		return nil
	}
	return *page.ilr.Value
}

// Operation a Billing REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Billing
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list billing operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of billing operations supported by the Microsoft.Billing resource provider.
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

// Resource the Resource model definition.
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}
