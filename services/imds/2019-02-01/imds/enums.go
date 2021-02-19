package imds

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

// BypassCache enumerates the values for bypass cache.
type BypassCache string

const (
	// True ...
	True BypassCache = "true"
)

// PossibleBypassCacheValues returns an array of possible values for the BypassCache const type.
func PossibleBypassCacheValues() []BypassCache {
	return []BypassCache{True}
}

// Error enumerates the values for error.
type Error string

const (
	// AccessDenied ...
	AccessDenied Error = "access_denied"
	// BadRequest ...
	BadRequest Error = "bad_request"
	// Forbidden ...
	Forbidden Error = "forbidden"
	// InvalidRequest ...
	InvalidRequest Error = "invalid_request"
	// InvalidScope ...
	InvalidScope Error = "invalid_scope"
	// MethodNotAllowed ...
	MethodNotAllowed Error = "method_not_allowed"
	// NotFound ...
	NotFound Error = "not_found"
	// ServerError ...
	ServerError Error = "server_error"
	// ServiceUnavailable ...
	ServiceUnavailable Error = "service_unavailable"
	// TooManyRequests ...
	TooManyRequests Error = "too_many_requests"
	// UnauthorizedClient ...
	UnauthorizedClient Error = "unauthorized_client"
	// UnsupportedResponseType ...
	UnsupportedResponseType Error = "unsupported_response_type"
)

// PossibleErrorValues returns an array of possible values for the Error const type.
func PossibleErrorValues() []Error {
	return []Error{AccessDenied, BadRequest, Forbidden, InvalidRequest, InvalidScope, MethodNotAllowed, NotFound, ServerError, ServiceUnavailable, TooManyRequests, UnauthorizedClient, UnsupportedResponseType}
}
