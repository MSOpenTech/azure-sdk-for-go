// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package personalizer

import original "github.com/Azure/azure-sdk-for-go/services/preview/personalizer/v1.0/personalizer"

type ErrorCode = original.ErrorCode

const (
	BadRequest                  ErrorCode = original.BadRequest
	EvaluationNotFound          ErrorCode = original.EvaluationNotFound
	FrontEndNotFound            ErrorCode = original.FrontEndNotFound
	InternalServerError         ErrorCode = original.InternalServerError
	InvalidContainer            ErrorCode = original.InvalidContainer
	InvalidEvaluationContract   ErrorCode = original.InvalidEvaluationContract
	InvalidEventIDToActivate    ErrorCode = original.InvalidEventIDToActivate
	InvalidExportLogsRequest    ErrorCode = original.InvalidExportLogsRequest
	InvalidPolicyConfiguration  ErrorCode = original.InvalidPolicyConfiguration
	InvalidPolicyContract       ErrorCode = original.InvalidPolicyContract
	InvalidRankRequest          ErrorCode = original.InvalidRankRequest
	InvalidRewardRequest        ErrorCode = original.InvalidRewardRequest
	InvalidServiceConfiguration ErrorCode = original.InvalidServiceConfiguration
	LogsPropertiesNotFound      ErrorCode = original.LogsPropertiesNotFound
	MissingContainer            ErrorCode = original.MissingContainer
	ModelResetFailed            ErrorCode = original.ModelResetFailed
	RankNullResponse            ErrorCode = original.RankNullResponse
	ResourceNotFound            ErrorCode = original.ResourceNotFound
	UpdateConfigurationFailed   ErrorCode = original.UpdateConfigurationFailed
)

type EvaluationJobStatus = original.EvaluationJobStatus

const (
	Completed    EvaluationJobStatus = original.Completed
	Failed       EvaluationJobStatus = original.Failed
	NotSubmitted EvaluationJobStatus = original.NotSubmitted
	Pending      EvaluationJobStatus = original.Pending
)

type BaseClient = original.BaseClient
type ContainerStatus = original.ContainerStatus
type DateRange = original.DateRange
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Evaluation = original.Evaluation
type EvaluationContract = original.EvaluationContract
type EvaluationsClient = original.EvaluationsClient
type EventsClient = original.EventsClient
type InternalError = original.InternalError
type ListEvaluation = original.ListEvaluation
type LogClient = original.LogClient
type LogsProperties = original.LogsProperties
type LogsPropertiesDateRange = original.LogsPropertiesDateRange
type ModelClient = original.ModelClient
type ModelProperties = original.ModelProperties
type PolicyClient = original.PolicyClient
type PolicyContract = original.PolicyContract
type PolicyResult = original.PolicyResult
type PolicyResultSummary = original.PolicyResultSummary
type PolicyResultTotalSummary = original.PolicyResultTotalSummary
type RankRequest = original.RankRequest
type RankResponse = original.RankResponse
type RankableAction = original.RankableAction
type RankedAction = original.RankedAction
type ReadCloser = original.ReadCloser
type RewardRequest = original.RewardRequest
type ServiceConfiguration = original.ServiceConfiguration
type ServiceConfigurationClient = original.ServiceConfigurationClient

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewEvaluationsClient(endpoint string) EvaluationsClient {
	return original.NewEvaluationsClient(endpoint)
}
func NewEventsClient(endpoint string) EventsClient {
	return original.NewEventsClient(endpoint)
}
func NewLogClient(endpoint string) LogClient {
	return original.NewLogClient(endpoint)
}
func NewModelClient(endpoint string) ModelClient {
	return original.NewModelClient(endpoint)
}
func NewPolicyClient(endpoint string) PolicyClient {
	return original.NewPolicyClient(endpoint)
}
func NewServiceConfigurationClient(endpoint string) ServiceConfigurationClient {
	return original.NewServiceConfigurationClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleEvaluationJobStatusValues() []EvaluationJobStatus {
	return original.PossibleEvaluationJobStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
