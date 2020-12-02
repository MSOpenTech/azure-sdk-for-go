package qnamaker

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

// EnvironmentType enumerates the values for environment type.
type EnvironmentType string

const (
	// Prod ...
	Prod EnvironmentType = "Prod"
	// Test ...
	Test EnvironmentType = "Test"
)

// PossibleEnvironmentTypeValues returns an array of possible values for the EnvironmentType const type.
func PossibleEnvironmentTypeValues() []EnvironmentType {
	return []EnvironmentType{Prod, Test}
}

// ErrorCodeType enumerates the values for error code type.
type ErrorCodeType string

const (
	// BadArgument ...
	BadArgument ErrorCodeType = "BadArgument"
	// EndpointKeysError ...
	EndpointKeysError ErrorCodeType = "EndpointKeysError"
	// ExtractionFailure ...
	ExtractionFailure ErrorCodeType = "ExtractionFailure"
	// Forbidden ...
	Forbidden ErrorCodeType = "Forbidden"
	// KbNotFound ...
	KbNotFound ErrorCodeType = "KbNotFound"
	// NotFound ...
	NotFound ErrorCodeType = "NotFound"
	// OperationNotFound ...
	OperationNotFound ErrorCodeType = "OperationNotFound"
	// QnaRuntimeError ...
	QnaRuntimeError ErrorCodeType = "QnaRuntimeError"
	// QuotaExceeded ...
	QuotaExceeded ErrorCodeType = "QuotaExceeded"
	// ServiceError ...
	ServiceError ErrorCodeType = "ServiceError"
	// SKULimitExceeded ...
	SKULimitExceeded ErrorCodeType = "SKULimitExceeded"
	// Unauthorized ...
	Unauthorized ErrorCodeType = "Unauthorized"
	// Unspecified ...
	Unspecified ErrorCodeType = "Unspecified"
	// ValidationFailure ...
	ValidationFailure ErrorCodeType = "ValidationFailure"
)

// PossibleErrorCodeTypeValues returns an array of possible values for the ErrorCodeType const type.
func PossibleErrorCodeTypeValues() []ErrorCodeType {
	return []ErrorCodeType{BadArgument, EndpointKeysError, ExtractionFailure, Forbidden, KbNotFound, NotFound, OperationNotFound, QnaRuntimeError, QuotaExceeded, ServiceError, SKULimitExceeded, Unauthorized, Unspecified, ValidationFailure}
}

// OperationStateType enumerates the values for operation state type.
type OperationStateType string

const (
	// Failed ...
	Failed OperationStateType = "Failed"
	// NotStarted ...
	NotStarted OperationStateType = "NotStarted"
	// Running ...
	Running OperationStateType = "Running"
	// Succeeded ...
	Succeeded OperationStateType = "Succeeded"
)

// PossibleOperationStateTypeValues returns an array of possible values for the OperationStateType const type.
func PossibleOperationStateTypeValues() []OperationStateType {
	return []OperationStateType{Failed, NotStarted, Running, Succeeded}
}

// StrictFiltersCompoundOperationType enumerates the values for strict filters compound operation type.
type StrictFiltersCompoundOperationType string

const (
	// AND ...
	AND StrictFiltersCompoundOperationType = "AND"
	// OR ...
	OR StrictFiltersCompoundOperationType = "OR"
)

// PossibleStrictFiltersCompoundOperationTypeValues returns an array of possible values for the StrictFiltersCompoundOperationType const type.
func PossibleStrictFiltersCompoundOperationTypeValues() []StrictFiltersCompoundOperationType {
	return []StrictFiltersCompoundOperationType{AND, OR}
}
