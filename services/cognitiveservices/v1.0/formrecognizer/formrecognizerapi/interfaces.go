package formrecognizerapi

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
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/formrecognizer"
	"github.com/Azure/go-autorest/autorest"
	"github.com/satori/go.uuid"
	"io"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	AnalyzeWithModel(ctx context.Context, ID uuid.UUID, formStream io.ReadCloser, keys []string) (result formrecognizer.AnalyzeResult, err error)
	DeleteModel(ctx context.Context, ID uuid.UUID) (result autorest.Response, err error)
	GetExtractedKeys(ctx context.Context, ID uuid.UUID) (result formrecognizer.KeysResult, err error)
	GetModel(ctx context.Context, ID uuid.UUID) (result formrecognizer.ModelResult, err error)
	GetModels(ctx context.Context) (result formrecognizer.ModelsResult, err error)
	TrainModel(ctx context.Context, trainRequest formrecognizer.TrainRequest) (result formrecognizer.TrainResult, err error)
}

var _ BaseClientAPI = (*formrecognizer.BaseClient)(nil)
