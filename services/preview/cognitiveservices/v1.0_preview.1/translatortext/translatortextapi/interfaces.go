package translatortextapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0_preview.1/translatortext"
	"github.com/gofrs/uuid"
)

// TranslationClientAPI contains the set of methods on the TranslationClient type.
type TranslationClientAPI interface {
	CancelOperation(ctx context.Context, endpoint string, ID uuid.UUID) (result translatortext.SetObject, err error)
	GetDocumentFormats(ctx context.Context, endpoint string) (result translatortext.SetObject, err error)
	GetDocumentStatus(ctx context.Context, endpoint string, ID uuid.UUID, documentID uuid.UUID) (result translatortext.SetObject, err error)
	GetDocumentStorageSource(ctx context.Context, endpoint string) (result translatortext.SetObject, err error)
	GetGlossaryFormats(ctx context.Context, endpoint string) (result translatortext.SetObject, err error)
	GetOperationDocumentsStatus(ctx context.Context, endpoint string, ID uuid.UUID, top *int32, skip *int32) (result translatortext.SetObject, err error)
	GetOperations(ctx context.Context, endpoint string, top *int32, skip *int32) (result translatortext.SetObject, err error)
	GetOperationStatus(ctx context.Context, endpoint string, ID uuid.UUID) (result translatortext.SetObject, err error)
	SubmitBatchRequest(ctx context.Context, endpoint string, body *translatortext.BatchSubmissionRequest) (result translatortext.ErrorResponseV2, err error)
}

var _ TranslationClientAPI = (*translatortext.TranslationClient)(nil)
