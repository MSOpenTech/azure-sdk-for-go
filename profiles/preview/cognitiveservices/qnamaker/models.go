// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package qnamaker

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v4.0/qnamaker"

type AlterationsClient = original.AlterationsClient
type BaseClient = original.BaseClient
type EndpointClient = original.EndpointClient
type KnowledgebaseClient = original.KnowledgebaseClient
type EnvironmentType = original.EnvironmentType

const (
	Prod EnvironmentType = original.Prod
	Test EnvironmentType = original.Test
)

type ErrorCodeType = original.ErrorCodeType

const (
	BadArgument       ErrorCodeType = original.BadArgument
	EndpointKeysError ErrorCodeType = original.EndpointKeysError
	ExtractionFailure ErrorCodeType = original.ExtractionFailure
	Forbidden         ErrorCodeType = original.Forbidden
	KbNotFound        ErrorCodeType = original.KbNotFound
	NotFound          ErrorCodeType = original.NotFound
	OperationNotFound ErrorCodeType = original.OperationNotFound
	QnaRuntimeError   ErrorCodeType = original.QnaRuntimeError
	QuotaExceeded     ErrorCodeType = original.QuotaExceeded
	ServiceError      ErrorCodeType = original.ServiceError
	SKULimitExceeded  ErrorCodeType = original.SKULimitExceeded
	Unauthorized      ErrorCodeType = original.Unauthorized
	Unspecified       ErrorCodeType = original.Unspecified
	ValidationFailure ErrorCodeType = original.ValidationFailure
)

type KnowledgebaseEnvironmentType = original.KnowledgebaseEnvironmentType

const (
	KnowledgebaseEnvironmentTypeProd KnowledgebaseEnvironmentType = original.KnowledgebaseEnvironmentTypeProd
	KnowledgebaseEnvironmentTypeTest KnowledgebaseEnvironmentType = original.KnowledgebaseEnvironmentTypeTest
)

type OperationStateType = original.OperationStateType

const (
	Failed     OperationStateType = original.Failed
	NotStarted OperationStateType = original.NotStarted
	Running    OperationStateType = original.Running
	Succeeded  OperationStateType = original.Succeeded
)

type AlterationsDTO = original.AlterationsDTO
type CreateKbDTO = original.CreateKbDTO
type CreateKbInputDTO = original.CreateKbInputDTO
type DeleteKbContentsDTO = original.DeleteKbContentsDTO
type EndpointKeysDTO = original.EndpointKeysDTO
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type FileDTO = original.FileDTO
type InnerErrorModel = original.InnerErrorModel
type KnowledgebaseDTO = original.KnowledgebaseDTO
type KnowledgebasesDTO = original.KnowledgebasesDTO
type MetadataDTO = original.MetadataDTO
type Operation = original.Operation
type QnADocumentsDTO = original.QnADocumentsDTO
type QnADTO = original.QnADTO
type ReplaceKbDTO = original.ReplaceKbDTO
type UpdateKbContentsDTO = original.UpdateKbContentsDTO
type UpdateKbOperationDTO = original.UpdateKbOperationDTO
type UpdateKbOperationDTOAdd = original.UpdateKbOperationDTOAdd
type UpdateKbOperationDTODelete = original.UpdateKbOperationDTODelete
type UpdateKbOperationDTOUpdate = original.UpdateKbOperationDTOUpdate
type UpdateMetadataDTO = original.UpdateMetadataDTO
type UpdateQnaDTO = original.UpdateQnaDTO
type UpdateQnaDTOMetadata = original.UpdateQnaDTOMetadata
type UpdateQnaDTOQuestions = original.UpdateQnaDTOQuestions
type UpdateQuestionsDTO = original.UpdateQuestionsDTO
type WordAlterationsDTO = original.WordAlterationsDTO
type OperationsClient = original.OperationsClient

func NewAlterationsClient(endpoint string) AlterationsClient {
	return original.NewAlterationsClient(endpoint)
}
func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func NewEndpointClient(endpoint string) EndpointClient {
	return original.NewEndpointClient(endpoint)
}
func NewKnowledgebaseClient(endpoint string) KnowledgebaseClient {
	return original.NewKnowledgebaseClient(endpoint)
}
func PossibleEnvironmentTypeValues() []EnvironmentType {
	return original.PossibleEnvironmentTypeValues()
}
func PossibleErrorCodeTypeValues() []ErrorCodeType {
	return original.PossibleErrorCodeTypeValues()
}
func PossibleKnowledgebaseEnvironmentTypeValues() []KnowledgebaseEnvironmentType {
	return original.PossibleKnowledgebaseEnvironmentTypeValues()
}
func PossibleOperationStateTypeValues() []OperationStateType {
	return original.PossibleOperationStateTypeValues()
}
func NewOperationsClient(endpoint string) OperationsClient {
	return original.NewOperationsClient(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
