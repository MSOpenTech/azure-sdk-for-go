//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenaiassistants

// MessageContentClassification provides polymorphic access to related types.
// Call the interface's GetMessageContent() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MessageContent, *MessageImageFileContent, *MessageTextContent
type MessageContentClassification interface {
	// GetMessageContent returns the MessageContent content of the underlying type.
	GetMessageContent() *MessageContent
}

// MessageDeltaContentClassification provides polymorphic access to related types.
// Call the interface's GetMessageDeltaContent() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MessageDeltaContent, *MessageDeltaImageFileContent, *MessageDeltaTextContentObject
type MessageDeltaContentClassification interface {
	// GetMessageDeltaContent returns the MessageDeltaContent content of the underlying type.
	GetMessageDeltaContent() *MessageDeltaContent
}

// MessageDeltaTextAnnotationClassification provides polymorphic access to related types.
// Call the interface's GetMessageDeltaTextAnnotation() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MessageDeltaTextAnnotation, *MessageDeltaTextFileCitationAnnotationObject, *MessageDeltaTextFilePathAnnotationObject
type MessageDeltaTextAnnotationClassification interface {
	// GetMessageDeltaTextAnnotation returns the MessageDeltaTextAnnotation content of the underlying type.
	GetMessageDeltaTextAnnotation() *MessageDeltaTextAnnotation
}

// MessageTextAnnotationClassification provides polymorphic access to related types.
// Call the interface's GetMessageTextAnnotation() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MessageTextAnnotation, *MessageTextFileCitationAnnotation, *MessageTextFilePathAnnotation
type MessageTextAnnotationClassification interface {
	// GetMessageTextAnnotation returns the MessageTextAnnotation content of the underlying type.
	GetMessageTextAnnotation() *MessageTextAnnotation
}

// RequiredActionClassification provides polymorphic access to related types.
// Call the interface's GetRequiredAction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RequiredAction, *SubmitToolOutputsAction
type RequiredActionClassification interface {
	// GetRequiredAction returns the RequiredAction content of the underlying type.
	GetRequiredAction() *RequiredAction
}

// RequiredToolCallClassification provides polymorphic access to related types.
// Call the interface's GetRequiredToolCall() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RequiredFunctionToolCall, *RequiredToolCall
type RequiredToolCallClassification interface {
	// GetRequiredToolCall returns the RequiredToolCall content of the underlying type.
	GetRequiredToolCall() *RequiredToolCall
}

// RunStepCodeInterpreterToolCallOutputClassification provides polymorphic access to related types.
// Call the interface's GetRunStepCodeInterpreterToolCallOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RunStepCodeInterpreterImageOutput, *RunStepCodeInterpreterLogOutput, *RunStepCodeInterpreterToolCallOutput
type RunStepCodeInterpreterToolCallOutputClassification interface {
	// GetRunStepCodeInterpreterToolCallOutput returns the RunStepCodeInterpreterToolCallOutput content of the underlying type.
	GetRunStepCodeInterpreterToolCallOutput() *RunStepCodeInterpreterToolCallOutput
}

// RunStepDeltaCodeInterpreterOutputClassification provides polymorphic access to related types.
// Call the interface's GetRunStepDeltaCodeInterpreterOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RunStepDeltaCodeInterpreterImageOutput, *RunStepDeltaCodeInterpreterLogOutput, *RunStepDeltaCodeInterpreterOutput
type RunStepDeltaCodeInterpreterOutputClassification interface {
	// GetRunStepDeltaCodeInterpreterOutput returns the RunStepDeltaCodeInterpreterOutput content of the underlying type.
	GetRunStepDeltaCodeInterpreterOutput() *RunStepDeltaCodeInterpreterOutput
}

// RunStepDeltaDetailClassification provides polymorphic access to related types.
// Call the interface's GetRunStepDeltaDetail() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RunStepDeltaDetail, *RunStepDeltaMessageCreation, *RunStepDeltaToolCallObject
type RunStepDeltaDetailClassification interface {
	// GetRunStepDeltaDetail returns the RunStepDeltaDetail content of the underlying type.
	GetRunStepDeltaDetail() *RunStepDeltaDetail
}

// RunStepDeltaToolCallClassification provides polymorphic access to related types.
// Call the interface's GetRunStepDeltaToolCall() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RunStepDeltaCodeInterpreterToolCall, *RunStepDeltaFileSearchToolCall, *RunStepDeltaFunctionToolCall, *RunStepDeltaToolCall
type RunStepDeltaToolCallClassification interface {
	// GetRunStepDeltaToolCall returns the RunStepDeltaToolCall content of the underlying type.
	GetRunStepDeltaToolCall() *RunStepDeltaToolCall
}

// RunStepDetailsClassification provides polymorphic access to related types.
// Call the interface's GetRunStepDetails() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RunStepDetails, *RunStepMessageCreationDetails, *RunStepToolCallDetails
type RunStepDetailsClassification interface {
	// GetRunStepDetails returns the RunStepDetails content of the underlying type.
	GetRunStepDetails() *RunStepDetails
}

// RunStepToolCallClassification provides polymorphic access to related types.
// Call the interface's GetRunStepToolCall() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *RunStepCodeInterpreterToolCall, *RunStepFileSearchToolCall, *RunStepFunctionToolCall, *RunStepToolCall
type RunStepToolCallClassification interface {
	// GetRunStepToolCall returns the RunStepToolCall content of the underlying type.
	GetRunStepToolCall() *RunStepToolCall
}

// ToolDefinitionClassification provides polymorphic access to related types.
// Call the interface's GetToolDefinition() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CodeInterpreterToolDefinition, *FileSearchToolDefinition, *FunctionToolDefinition, *ToolDefinition
type ToolDefinitionClassification interface {
	// GetToolDefinition returns the ToolDefinition content of the underlying type.
	GetToolDefinition() *ToolDefinition
}

// VectorStoreChunkingStrategyRequestClassification provides polymorphic access to related types.
// Call the interface's GetVectorStoreChunkingStrategyRequest() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *VectorStoreAutoChunkingStrategyRequest, *VectorStoreChunkingStrategyRequest, *VectorStoreStaticChunkingStrategyRequest
type VectorStoreChunkingStrategyRequestClassification interface {
	// GetVectorStoreChunkingStrategyRequest returns the VectorStoreChunkingStrategyRequest content of the underlying type.
	GetVectorStoreChunkingStrategyRequest() *VectorStoreChunkingStrategyRequest
}

// VectorStoreChunkingStrategyResponseClassification provides polymorphic access to related types.
// Call the interface's GetVectorStoreChunkingStrategyResponse() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *VectorStoreAutoChunkingStrategyResponse, *VectorStoreChunkingStrategyResponse, *VectorStoreStaticChunkingStrategyResponse
type VectorStoreChunkingStrategyResponseClassification interface {
	// GetVectorStoreChunkingStrategyResponse returns the VectorStoreChunkingStrategyResponse content of the underlying type.
	GetVectorStoreChunkingStrategyResponse() *VectorStoreChunkingStrategyResponse
}
