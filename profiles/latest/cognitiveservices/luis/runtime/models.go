//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package runtime

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/luis/runtime"

type BaseClient = original.BaseClient
type DynamicList = original.DynamicList
type Error = original.Error
type ErrorBody = original.ErrorBody
type ExternalEntity = original.ExternalEntity
type Intent = original.Intent
type Prediction = original.Prediction
type PredictionClient = original.PredictionClient
type PredictionRequest = original.PredictionRequest
type PredictionRequestOptions = original.PredictionRequestOptions
type PredictionResponse = original.PredictionResponse
type RequestList = original.RequestList
type Sentiment = original.Sentiment

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewPredictionClient(endpoint string) PredictionClient {
	return original.NewPredictionClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
