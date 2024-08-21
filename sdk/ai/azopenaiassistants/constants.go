//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenaiassistants

// APIResponseFormat - Must be one of text or json_object.
type APIResponseFormat string

const (
	// APIResponseFormatJSONObject - Using `json_object` format will limit the usage of ToolCall to only functions.
	APIResponseFormatJSONObject APIResponseFormat = "json_object"
	// APIResponseFormatText - `text` format should be used for requests involving any sort of ToolCall.
	APIResponseFormatText APIResponseFormat = "text"
)

// PossibleAPIResponseFormatValues returns the possible values for the APIResponseFormat const type.
func PossibleAPIResponseFormatValues() []APIResponseFormat {
	return []APIResponseFormat{
		APIResponseFormatJSONObject,
		APIResponseFormatText,
	}
}

// AssistantResponseFormatType - Represents the mode in which the model will handle the return format of a tool call.
type AssistantResponseFormatType string

const (
	// AssistantResponseFormatTypeAuto - Default value. Let the model handle the return format.
	AssistantResponseFormatTypeAuto AssistantResponseFormatType = "auto"
	// AssistantResponseFormatTypeJSONObject - Using `json_object` format will limit the usage of ToolCall to only functions.
	AssistantResponseFormatTypeJSONObject AssistantResponseFormatType = "json_object"
	// AssistantResponseFormatTypeText - `text` format should be used for requests involving any sort of ToolCall.
	AssistantResponseFormatTypeText AssistantResponseFormatType = "text"
)

// PossibleAssistantResponseFormatTypeValues returns the possible values for the AssistantResponseFormatType const type.
func PossibleAssistantResponseFormatTypeValues() []AssistantResponseFormatType {
	return []AssistantResponseFormatType{
		AssistantResponseFormatTypeAuto,
		AssistantResponseFormatTypeJSONObject,
		AssistantResponseFormatTypeText,
	}
}

// AssistantStreamEvent - Each event in a server-sent events stream has an event and data property:
// event: thread.created
// data: {"id": "thread_123", "object": "thread", ...}
// We emit events whenever a new object is created, transitions to a new state, or is being streamed in parts (deltas). For
// example, we emit thread.run.created when a new run is created,
// thread.run.completed when a run completes, and so on. When an Assistant chooses to create a message during a run, we emit
// a thread.message.created event, athread.message.in_progress event, many
// thread.message.delta events, and finally athread.message.completed event.
// We may add additional events over time, so we recommend handling unknown events gracefully in your code.
type AssistantStreamEvent string

const (
	// AssistantStreamEventDone - Event sent when the stream is done.
	AssistantStreamEventDone AssistantStreamEvent = "done"
	// AssistantStreamEventError - Event sent when an error occurs, such as an internal server error or a timeout.
	AssistantStreamEventError AssistantStreamEvent = "error"
	// AssistantStreamEventThreadCreated - Event sent when a new thread is created. The data of this event is of type AssistantThread
	AssistantStreamEventThreadCreated AssistantStreamEvent = "thread.created"
	// AssistantStreamEventThreadMessageCompleted - Event sent when a message is completed. The data of this event is of type
	// ThreadMessage
	AssistantStreamEventThreadMessageCompleted AssistantStreamEvent = "thread.message.completed"
	// AssistantStreamEventThreadMessageCreated - Event sent when a new message is created. The data of this event is of type
	// ThreadMessage
	AssistantStreamEventThreadMessageCreated AssistantStreamEvent = "thread.message.created"
	// AssistantStreamEventThreadMessageDelta - Event sent when a message is being streamed. The data of this event is of type
	// MessageDeltaChunk
	AssistantStreamEventThreadMessageDelta AssistantStreamEvent = "thread.message.delta"
	// AssistantStreamEventThreadMessageInProgress - Event sent when a message moves to `in_progress` status. The data of this
	// event is of type ThreadMessage
	AssistantStreamEventThreadMessageInProgress AssistantStreamEvent = "thread.message.in_progress"
	// AssistantStreamEventThreadMessageIncomplete - Event sent before a message is completed. The data of this event is of type
	// ThreadMessage
	AssistantStreamEventThreadMessageIncomplete AssistantStreamEvent = "thread.message.incomplete"
	// AssistantStreamEventThreadRunCancelled - Event sent when a run is cancelled. The data of this event is of type ThreadRun
	AssistantStreamEventThreadRunCancelled AssistantStreamEvent = "thread.run.cancelled"
	// AssistantStreamEventThreadRunCancelling - Event sent when a run moves to `cancelling` status. The data of this event is
	// of type ThreadRun
	AssistantStreamEventThreadRunCancelling AssistantStreamEvent = "thread.run.cancelling"
	// AssistantStreamEventThreadRunCompleted - Event sent when a run is completed. The data of this event is of type ThreadRun
	AssistantStreamEventThreadRunCompleted AssistantStreamEvent = "thread.run.completed"
	// AssistantStreamEventThreadRunCreated - Event sent when a new run is created. The data of this event is of type ThreadRun
	AssistantStreamEventThreadRunCreated AssistantStreamEvent = "thread.run.created"
	// AssistantStreamEventThreadRunExpired - Event sent when a run is expired. The data of this event is of type ThreadRun
	AssistantStreamEventThreadRunExpired AssistantStreamEvent = "thread.run.expired"
	// AssistantStreamEventThreadRunFailed - Event sent when a run fails. The data of this event is of type ThreadRun
	AssistantStreamEventThreadRunFailed AssistantStreamEvent = "thread.run.failed"
	// AssistantStreamEventThreadRunInProgress - Event sent when a run moves to `in_progress` status. The data of this event is
	// of type ThreadRun
	AssistantStreamEventThreadRunInProgress AssistantStreamEvent = "thread.run.in_progress"
	// AssistantStreamEventThreadRunQueued - Event sent when a run moves to `queued` status. The data of this event is of type
	// ThreadRun
	AssistantStreamEventThreadRunQueued AssistantStreamEvent = "thread.run.queued"
	// AssistantStreamEventThreadRunRequiresAction - Event sent when a run moves to `requires_action` status. The data of this
	// event is of type ThreadRun
	AssistantStreamEventThreadRunRequiresAction AssistantStreamEvent = "thread.run.requires_action"
	// AssistantStreamEventThreadRunStepCancelled - Event sent when a run step is cancelled. The data of this event is of type
	// RunStep
	AssistantStreamEventThreadRunStepCancelled AssistantStreamEvent = "thread.run.step.cancelled"
	// AssistantStreamEventThreadRunStepCompleted - Event sent when a run step is completed. The data of this event is of type
	// RunStep
	AssistantStreamEventThreadRunStepCompleted AssistantStreamEvent = "thread.run.step.completed"
	// AssistantStreamEventThreadRunStepCreated - Event sent when a new thread run step is created. The data of this event is
	// of type RunStep
	AssistantStreamEventThreadRunStepCreated AssistantStreamEvent = "thread.run.step.created"
	// AssistantStreamEventThreadRunStepDelta - Event sent when a run stepis being streamed. The data of this event is of type
	// RunStepDeltaChunk
	AssistantStreamEventThreadRunStepDelta AssistantStreamEvent = "thread.run.step.delta"
	// AssistantStreamEventThreadRunStepExpired - Event sent when a run step is expired. The data of this event is of type RunStep
	AssistantStreamEventThreadRunStepExpired AssistantStreamEvent = "thread.run.step.expired"
	// AssistantStreamEventThreadRunStepFailed - Event sent when a run step fails. The data of this event is of type RunStep
	AssistantStreamEventThreadRunStepFailed AssistantStreamEvent = "thread.run.step.failed"
	// AssistantStreamEventThreadRunStepInProgress - Event sent when a run step moves to `in_progress` status. The data of this
	// event is of type RunStep
	AssistantStreamEventThreadRunStepInProgress AssistantStreamEvent = "thread.run.step.in_progress"
)

// PossibleAssistantStreamEventValues returns the possible values for the AssistantStreamEvent const type.
func PossibleAssistantStreamEventValues() []AssistantStreamEvent {
	return []AssistantStreamEvent{
		AssistantStreamEventDone,
		AssistantStreamEventError,
		AssistantStreamEventThreadCreated,
		AssistantStreamEventThreadMessageCompleted,
		AssistantStreamEventThreadMessageCreated,
		AssistantStreamEventThreadMessageDelta,
		AssistantStreamEventThreadMessageInProgress,
		AssistantStreamEventThreadMessageIncomplete,
		AssistantStreamEventThreadRunCancelled,
		AssistantStreamEventThreadRunCancelling,
		AssistantStreamEventThreadRunCompleted,
		AssistantStreamEventThreadRunCreated,
		AssistantStreamEventThreadRunExpired,
		AssistantStreamEventThreadRunFailed,
		AssistantStreamEventThreadRunInProgress,
		AssistantStreamEventThreadRunQueued,
		AssistantStreamEventThreadRunRequiresAction,
		AssistantStreamEventThreadRunStepCancelled,
		AssistantStreamEventThreadRunStepCompleted,
		AssistantStreamEventThreadRunStepCreated,
		AssistantStreamEventThreadRunStepDelta,
		AssistantStreamEventThreadRunStepExpired,
		AssistantStreamEventThreadRunStepFailed,
		AssistantStreamEventThreadRunStepInProgress,
	}
}

// AssistantsAPIToolChoiceOptionMode - Specifies how the tool choice will be used
type AssistantsAPIToolChoiceOptionMode string

const (
	// AssistantsAPIToolChoiceOptionModeAuto - The model can pick between generating a message or calling a function.
	AssistantsAPIToolChoiceOptionModeAuto AssistantsAPIToolChoiceOptionMode = "auto"
	// AssistantsAPIToolChoiceOptionModeCodeInterpreter - Tool type `code_interpreter`
	AssistantsAPIToolChoiceOptionModeCodeInterpreter AssistantsAPIToolChoiceOptionMode = "code_interpreter"
	// AssistantsAPIToolChoiceOptionModeFileSearch - Tool type `file_search`
	AssistantsAPIToolChoiceOptionModeFileSearch AssistantsAPIToolChoiceOptionMode = "file_search"
	// AssistantsAPIToolChoiceOptionModeFunction - Tool type `function`
	AssistantsAPIToolChoiceOptionModeFunction AssistantsAPIToolChoiceOptionMode = "function"
	// AssistantsAPIToolChoiceOptionModeNone - The model will not call a function and instead generates a message.
	AssistantsAPIToolChoiceOptionModeNone AssistantsAPIToolChoiceOptionMode = "none"
)

// PossibleAssistantsAPIToolChoiceOptionModeValues returns the possible values for the AssistantsAPIToolChoiceOptionMode const type.
func PossibleAssistantsAPIToolChoiceOptionModeValues() []AssistantsAPIToolChoiceOptionMode {
	return []AssistantsAPIToolChoiceOptionMode{
		AssistantsAPIToolChoiceOptionModeAuto,
		AssistantsAPIToolChoiceOptionModeCodeInterpreter,
		AssistantsAPIToolChoiceOptionModeFileSearch,
		AssistantsAPIToolChoiceOptionModeFunction,
		AssistantsAPIToolChoiceOptionModeNone,
	}
}

// DoneEvent - Terminal event indicating the successful end of a stream.
type DoneEvent string

const (
	// DoneEventDone - Event sent when the stream is done.
	DoneEventDone DoneEvent = "done"
)

// PossibleDoneEventValues returns the possible values for the DoneEvent const type.
func PossibleDoneEventValues() []DoneEvent {
	return []DoneEvent{
		DoneEventDone,
	}
}

// ErrorEvent - Terminal event indicating a server side error while streaming.
type ErrorEvent string

const (
	// ErrorEventError - Event sent when an error occurs, such as an internal server error or a timeout.
	ErrorEventError ErrorEvent = "error"
)

// PossibleErrorEventValues returns the possible values for the ErrorEvent const type.
func PossibleErrorEventValues() []ErrorEvent {
	return []ErrorEvent{
		ErrorEventError,
	}
}

// FilePurpose - The possible values denoting the intended usage of a file.
type FilePurpose string

const (
	// FilePurposeAssistants - Indicates a file is used as input to assistants.
	FilePurposeAssistants FilePurpose = "assistants"
	// FilePurposeAssistantsOutput - Indicates a file is used as output by assistants.
	FilePurposeAssistantsOutput FilePurpose = "assistants_output"
	// FilePurposeBatch - Indicates a file is used as input to .
	FilePurposeBatch FilePurpose = "batch"
	// FilePurposeBatchOutput - Indicates a file is used as output by a vector store batch operation.
	FilePurposeBatchOutput FilePurpose = "batch_output"
	// FilePurposeFineTune - Indicates a file is used for fine tuning input.
	FilePurposeFineTune FilePurpose = "fine-tune"
	// FilePurposeFineTuneResults - Indicates a file is used for fine tuning results.
	FilePurposeFineTuneResults FilePurpose = "fine-tune-results"
	// FilePurposeVision - Indicates a file is used as input to a vision operation.
	FilePurposeVision FilePurpose = "vision"
)

// PossibleFilePurposeValues returns the possible values for the FilePurpose const type.
func PossibleFilePurposeValues() []FilePurpose {
	return []FilePurpose{
		FilePurposeAssistants,
		FilePurposeAssistantsOutput,
		FilePurposeBatch,
		FilePurposeBatchOutput,
		FilePurposeFineTune,
		FilePurposeFineTuneResults,
		FilePurposeVision,
	}
}

// FileState - The state of the file.
type FileState string

const (
	// FileStateDeleted - The entity has been deleted but may still be referenced by other entities predating the deletion. It
	// can be categorized as a
	// terminal state.
	FileStateDeleted FileState = "deleted"
	// FileStateDeleting - The entity is in the process to be deleted. This state is not returned by Azure OpenAI and exposed
	// only for compatibility.
	// It can be categorized as an active state.
	FileStateDeleting FileState = "deleting"
	// FileStateError - The operation has completed processing with a failure and cannot be further consumed. It can be categorized
	// as a terminal state.
	FileStateError FileState = "error"
	// FileStatePending - The operation was created and is not queued to be processed in the future. It can be categorized as
	// an inactive state.
	FileStatePending FileState = "pending"
	// FileStateProcessed - The operation has successfully processed and is ready for consumption. It can be categorized as a
	// terminal state.
	FileStateProcessed FileState = "processed"
	// FileStateRunning - The operation has started to be processed. It can be categorized as an active state.
	FileStateRunning FileState = "running"
	// FileStateUploaded - The file has been uploaded but it's not yet processed. This state is not returned by Azure OpenAI and
	// exposed only for
	// compatibility. It can be categorized as an inactive state.
	FileStateUploaded FileState = "uploaded"
)

// PossibleFileStateValues returns the possible values for the FileState const type.
func PossibleFileStateValues() []FileState {
	return []FileState{
		FileStateDeleted,
		FileStateDeleting,
		FileStateError,
		FileStatePending,
		FileStateProcessed,
		FileStateRunning,
		FileStateUploaded,
	}
}

// IncompleteRunDetails - The reason why the run is incomplete. This will point to which specific token limit was reached
// over the course of the run.
type IncompleteRunDetails string

const (
	// IncompleteRunDetailsMaxCompletionTokens - Maximum completion tokens exceeded
	IncompleteRunDetailsMaxCompletionTokens IncompleteRunDetails = "max_completion_tokens"
	// IncompleteRunDetailsMaxPromptTokens - Maximum prompt tokens exceeded
	IncompleteRunDetailsMaxPromptTokens IncompleteRunDetails = "max_prompt_tokens"
)

// PossibleIncompleteRunDetailsValues returns the possible values for the IncompleteRunDetails const type.
func PossibleIncompleteRunDetailsValues() []IncompleteRunDetails {
	return []IncompleteRunDetails{
		IncompleteRunDetailsMaxCompletionTokens,
		IncompleteRunDetailsMaxPromptTokens,
	}
}

// ListSortOrder - The available sorting options when requesting a list of response objects.
type ListSortOrder string

const (
	// ListSortOrderAscending - Specifies an ascending sort order.
	ListSortOrderAscending ListSortOrder = "asc"
	// ListSortOrderDescending - Specifies a descending sort order.
	ListSortOrderDescending ListSortOrder = "desc"
)

// PossibleListSortOrderValues returns the possible values for the ListSortOrder const type.
func PossibleListSortOrderValues() []ListSortOrder {
	return []ListSortOrder{
		ListSortOrderAscending,
		ListSortOrderDescending,
	}
}

// MessageIncompleteDetailsReason - A set of reasons describing why a message is marked as incomplete.
type MessageIncompleteDetailsReason string

const (
	// MessageIncompleteDetailsReasonContentFilter - The run generating the message was terminated due to content filter flagging.
	MessageIncompleteDetailsReasonContentFilter MessageIncompleteDetailsReason = "content_filter"
	// MessageIncompleteDetailsReasonMaxTokens - The run generating the message exhausted available tokens before completion.
	MessageIncompleteDetailsReasonMaxTokens MessageIncompleteDetailsReason = "max_tokens"
	// MessageIncompleteDetailsReasonRunCancelled - The run generating the message was cancelled before completion.
	MessageIncompleteDetailsReasonRunCancelled MessageIncompleteDetailsReason = "run_cancelled"
	// MessageIncompleteDetailsReasonRunExpired - The run generating the message expired.
	MessageIncompleteDetailsReasonRunExpired MessageIncompleteDetailsReason = "run_expired"
	// MessageIncompleteDetailsReasonRunFailed - The run generating the message failed.
	MessageIncompleteDetailsReasonRunFailed MessageIncompleteDetailsReason = "run_failed"
)

// PossibleMessageIncompleteDetailsReasonValues returns the possible values for the MessageIncompleteDetailsReason const type.
func PossibleMessageIncompleteDetailsReasonValues() []MessageIncompleteDetailsReason {
	return []MessageIncompleteDetailsReason{
		MessageIncompleteDetailsReasonContentFilter,
		MessageIncompleteDetailsReasonMaxTokens,
		MessageIncompleteDetailsReasonRunCancelled,
		MessageIncompleteDetailsReasonRunExpired,
		MessageIncompleteDetailsReasonRunFailed,
	}
}

// MessageRole - The possible values for roles attributed to messages in a thread.
type MessageRole string

const (
	// MessageRoleAssistant - The role representing the assistant.
	MessageRoleAssistant MessageRole = "assistant"
	// MessageRoleUser - The role representing the end-user.
	MessageRoleUser MessageRole = "user"
)

// PossibleMessageRoleValues returns the possible values for the MessageRole const type.
func PossibleMessageRoleValues() []MessageRole {
	return []MessageRole{
		MessageRoleAssistant,
		MessageRoleUser,
	}
}

// MessageStatus - The possible execution status values for a thread message.
type MessageStatus string

const (
	// MessageStatusCompleted - This message was successfully completed by a run.
	MessageStatusCompleted MessageStatus = "completed"
	// MessageStatusInProgress - A run is currently creating this message.
	MessageStatusInProgress MessageStatus = "in_progress"
	// MessageStatusIncomplete - This message is incomplete. See incomplete_details for more information.
	MessageStatusIncomplete MessageStatus = "incomplete"
)

// PossibleMessageStatusValues returns the possible values for the MessageStatus const type.
func PossibleMessageStatusValues() []MessageStatus {
	return []MessageStatus{
		MessageStatusCompleted,
		MessageStatusInProgress,
		MessageStatusIncomplete,
	}
}

// MessageStreamEvent - Message operation related streaming events
type MessageStreamEvent string

const (
	// MessageStreamEventThreadMessageCompleted - Event sent when a message is completed. The data of this event is of type ThreadMessage
	MessageStreamEventThreadMessageCompleted MessageStreamEvent = "thread.message.completed"
	// MessageStreamEventThreadMessageCreated - Event sent when a new message is created. The data of this event is of type ThreadMessage
	MessageStreamEventThreadMessageCreated MessageStreamEvent = "thread.message.created"
	// MessageStreamEventThreadMessageDelta - Event sent when a message is being streamed. The data of this event is of type MessageDeltaChunk
	MessageStreamEventThreadMessageDelta MessageStreamEvent = "thread.message.delta"
	// MessageStreamEventThreadMessageInProgress - Event sent when a message moves to `in_progress` status. The data of this event
	// is of type ThreadMessage
	MessageStreamEventThreadMessageInProgress MessageStreamEvent = "thread.message.in_progress"
	// MessageStreamEventThreadMessageIncomplete - Event sent before a message is completed. The data of this event is of type
	// ThreadMessage
	MessageStreamEventThreadMessageIncomplete MessageStreamEvent = "thread.message.incomplete"
)

// PossibleMessageStreamEventValues returns the possible values for the MessageStreamEvent const type.
func PossibleMessageStreamEventValues() []MessageStreamEvent {
	return []MessageStreamEvent{
		MessageStreamEventThreadMessageCompleted,
		MessageStreamEventThreadMessageCreated,
		MessageStreamEventThreadMessageDelta,
		MessageStreamEventThreadMessageInProgress,
		MessageStreamEventThreadMessageIncomplete,
	}
}

// RunStatus - Possible values for the status of an assistant thread run.
type RunStatus string

const (
	// RunStatusCancelled - Represents a run that has been cancelled.
	RunStatusCancelled RunStatus = "cancelled"
	// RunStatusCancelling - Represents a run that is in the process of cancellation.
	RunStatusCancelling RunStatus = "cancelling"
	// RunStatusCompleted - Represents a run that successfully completed.
	RunStatusCompleted RunStatus = "completed"
	// RunStatusExpired - Represents a run that expired before it could otherwise finish.
	RunStatusExpired RunStatus = "expired"
	// RunStatusFailed - Represents a run that failed.
	RunStatusFailed RunStatus = "failed"
	// RunStatusInProgress - Represents a run that is in progress.
	RunStatusInProgress RunStatus = "in_progress"
	// RunStatusQueued - Represents a run that is queued to start.
	RunStatusQueued RunStatus = "queued"
	// RunStatusRequiresAction - Represents a run that needs another operation, such as tool output submission, to continue.
	RunStatusRequiresAction RunStatus = "requires_action"
)

// PossibleRunStatusValues returns the possible values for the RunStatus const type.
func PossibleRunStatusValues() []RunStatus {
	return []RunStatus{
		RunStatusCancelled,
		RunStatusCancelling,
		RunStatusCompleted,
		RunStatusExpired,
		RunStatusFailed,
		RunStatusInProgress,
		RunStatusQueued,
		RunStatusRequiresAction,
	}
}

// RunStepErrorCode - Possible error code values attributable to a failed run step.
type RunStepErrorCode string

const (
	// RunStepErrorCodeRateLimitExceeded - Represents an error indicating configured rate limits were exceeded.
	RunStepErrorCodeRateLimitExceeded RunStepErrorCode = "rate_limit_exceeded"
	// RunStepErrorCodeServerError - Represents a server error.
	RunStepErrorCodeServerError RunStepErrorCode = "server_error"
)

// PossibleRunStepErrorCodeValues returns the possible values for the RunStepErrorCode const type.
func PossibleRunStepErrorCodeValues() []RunStepErrorCode {
	return []RunStepErrorCode{
		RunStepErrorCodeRateLimitExceeded,
		RunStepErrorCodeServerError,
	}
}

// RunStepStatus - Possible values for the status of a run step.
type RunStepStatus string

const (
	// RunStepStatusCancelled - Represents a run step that was cancelled.
	RunStepStatusCancelled RunStepStatus = "cancelled"
	// RunStepStatusCompleted - Represents a run step that successfully completed.
	RunStepStatusCompleted RunStepStatus = "completed"
	// RunStepStatusExpired - Represents a run step that expired before otherwise finishing.
	RunStepStatusExpired RunStepStatus = "expired"
	// RunStepStatusFailed - Represents a run step that failed.
	RunStepStatusFailed RunStepStatus = "failed"
	// RunStepStatusInProgress - Represents a run step still in progress.
	RunStepStatusInProgress RunStepStatus = "in_progress"
)

// PossibleRunStepStatusValues returns the possible values for the RunStepStatus const type.
func PossibleRunStepStatusValues() []RunStepStatus {
	return []RunStepStatus{
		RunStepStatusCancelled,
		RunStepStatusCompleted,
		RunStepStatusExpired,
		RunStepStatusFailed,
		RunStepStatusInProgress,
	}
}

// RunStepStreamEvent - Run step operation related streaming events
type RunStepStreamEvent string

const (
	// RunStepStreamEventThreadRunStepCancelled - Event sent when a run step is cancelled. The data of this event is of type RunStep
	RunStepStreamEventThreadRunStepCancelled RunStepStreamEvent = "thread.run.step.cancelled"
	// RunStepStreamEventThreadRunStepCompleted - Event sent when a run step is completed. The data of this event is of type RunStep
	RunStepStreamEventThreadRunStepCompleted RunStepStreamEvent = "thread.run.step.completed"
	// RunStepStreamEventThreadRunStepCreated - Event sent when a new thread run step is created. The data of this event is of
	// type RunStep
	RunStepStreamEventThreadRunStepCreated RunStepStreamEvent = "thread.run.step.created"
	// RunStepStreamEventThreadRunStepDelta - Event sent when a run stepis being streamed. The data of this event is of type RunStepDeltaChunk
	RunStepStreamEventThreadRunStepDelta RunStepStreamEvent = "thread.run.step.delta"
	// RunStepStreamEventThreadRunStepExpired - Event sent when a run step is expired. The data of this event is of type RunStep
	RunStepStreamEventThreadRunStepExpired RunStepStreamEvent = "thread.run.step.expired"
	// RunStepStreamEventThreadRunStepFailed - Event sent when a run step fails. The data of this event is of type RunStep
	RunStepStreamEventThreadRunStepFailed RunStepStreamEvent = "thread.run.step.failed"
	// RunStepStreamEventThreadRunStepInProgress - Event sent when a run step moves to `in_progress` status. The data of this
	// event is of type RunStep
	RunStepStreamEventThreadRunStepInProgress RunStepStreamEvent = "thread.run.step.in_progress"
)

// PossibleRunStepStreamEventValues returns the possible values for the RunStepStreamEvent const type.
func PossibleRunStepStreamEventValues() []RunStepStreamEvent {
	return []RunStepStreamEvent{
		RunStepStreamEventThreadRunStepCancelled,
		RunStepStreamEventThreadRunStepCompleted,
		RunStepStreamEventThreadRunStepCreated,
		RunStepStreamEventThreadRunStepDelta,
		RunStepStreamEventThreadRunStepExpired,
		RunStepStreamEventThreadRunStepFailed,
		RunStepStreamEventThreadRunStepInProgress,
	}
}

// RunStepType - The possible types of run steps.
type RunStepType string

const (
	// RunStepTypeMessageCreation - Represents a run step to create a message.
	RunStepTypeMessageCreation RunStepType = "message_creation"
	// RunStepTypeToolCalls - Represents a run step that calls tools.
	RunStepTypeToolCalls RunStepType = "tool_calls"
)

// PossibleRunStepTypeValues returns the possible values for the RunStepType const type.
func PossibleRunStepTypeValues() []RunStepType {
	return []RunStepType{
		RunStepTypeMessageCreation,
		RunStepTypeToolCalls,
	}
}

// RunStreamEvent - Run operation related streaming events
type RunStreamEvent string

const (
	// RunStreamEventThreadRunCancelled - Event sent when a run is cancelled. The data of this event is of type ThreadRun
	RunStreamEventThreadRunCancelled RunStreamEvent = "thread.run.cancelled"
	// RunStreamEventThreadRunCancelling - Event sent when a run moves to `cancelling` status. The data of this event is of type
	// ThreadRun
	RunStreamEventThreadRunCancelling RunStreamEvent = "thread.run.cancelling"
	// RunStreamEventThreadRunCompleted - Event sent when a run is completed. The data of this event is of type ThreadRun
	RunStreamEventThreadRunCompleted RunStreamEvent = "thread.run.completed"
	// RunStreamEventThreadRunCreated - Event sent when a new run is created. The data of this event is of type ThreadRun
	RunStreamEventThreadRunCreated RunStreamEvent = "thread.run.created"
	// RunStreamEventThreadRunExpired - Event sent when a run is expired. The data of this event is of type ThreadRun
	RunStreamEventThreadRunExpired RunStreamEvent = "thread.run.expired"
	// RunStreamEventThreadRunFailed - Event sent when a run fails. The data of this event is of type ThreadRun
	RunStreamEventThreadRunFailed RunStreamEvent = "thread.run.failed"
	// RunStreamEventThreadRunInProgress - Event sent when a run moves to `in_progress` status. The data of this event is of type
	// ThreadRun
	RunStreamEventThreadRunInProgress RunStreamEvent = "thread.run.in_progress"
	// RunStreamEventThreadRunQueued - Event sent when a run moves to `queued` status. The data of this event is of type ThreadRun
	RunStreamEventThreadRunQueued RunStreamEvent = "thread.run.queued"
	// RunStreamEventThreadRunRequiresAction - Event sent when a run moves to `requires_action` status. The data of this event
	// is of type ThreadRun
	RunStreamEventThreadRunRequiresAction RunStreamEvent = "thread.run.requires_action"
)

// PossibleRunStreamEventValues returns the possible values for the RunStreamEvent const type.
func PossibleRunStreamEventValues() []RunStreamEvent {
	return []RunStreamEvent{
		RunStreamEventThreadRunCancelled,
		RunStreamEventThreadRunCancelling,
		RunStreamEventThreadRunCompleted,
		RunStreamEventThreadRunCreated,
		RunStreamEventThreadRunExpired,
		RunStreamEventThreadRunFailed,
		RunStreamEventThreadRunInProgress,
		RunStreamEventThreadRunQueued,
		RunStreamEventThreadRunRequiresAction,
	}
}

// ThreadStreamEvent - Thread operation related streaming events
type ThreadStreamEvent string

const (
	// ThreadStreamEventThreadCreated - Event sent when a new thread is created. The data of this event is of type AssistantThread
	ThreadStreamEventThreadCreated ThreadStreamEvent = "thread.created"
)

// PossibleThreadStreamEventValues returns the possible values for the ThreadStreamEvent const type.
func PossibleThreadStreamEventValues() []ThreadStreamEvent {
	return []ThreadStreamEvent{
		ThreadStreamEventThreadCreated,
	}
}

// TruncationStrategy - The truncation strategy to use for the thread. The default is auto. If set to last_messages, the thread
// will be truncated to the lastMessages count most recent messages in the thread. When set to auto
// , messages in the middle of the thread will be dropped to fit the context length of the model, max_prompt_tokens.
type TruncationStrategy string

const (
	// TruncationStrategyAuto - Default value. Messages in the middle of the thread will be dropped to fit the context length
	// of the model.
	TruncationStrategyAuto TruncationStrategy = "auto"
	// TruncationStrategyLastMessages - The thread will truncate to the `lastMessages` count of recent messages.
	TruncationStrategyLastMessages TruncationStrategy = "last_messages"
)

// PossibleTruncationStrategyValues returns the possible values for the TruncationStrategy const type.
func PossibleTruncationStrategyValues() []TruncationStrategy {
	return []TruncationStrategy{
		TruncationStrategyAuto,
		TruncationStrategyLastMessages,
	}
}

// VectorStoreChunkingStrategyRequestType - Type of chunking strategy
type VectorStoreChunkingStrategyRequestType string

const (
	VectorStoreChunkingStrategyRequestTypeAuto   VectorStoreChunkingStrategyRequestType = "auto"
	VectorStoreChunkingStrategyRequestTypeStatic VectorStoreChunkingStrategyRequestType = "static"
)

// PossibleVectorStoreChunkingStrategyRequestTypeValues returns the possible values for the VectorStoreChunkingStrategyRequestType const type.
func PossibleVectorStoreChunkingStrategyRequestTypeValues() []VectorStoreChunkingStrategyRequestType {
	return []VectorStoreChunkingStrategyRequestType{
		VectorStoreChunkingStrategyRequestTypeAuto,
		VectorStoreChunkingStrategyRequestTypeStatic,
	}
}

// VectorStoreChunkingStrategyResponseType - Type of chunking strategy
type VectorStoreChunkingStrategyResponseType string

const (
	VectorStoreChunkingStrategyResponseTypeOther  VectorStoreChunkingStrategyResponseType = "other"
	VectorStoreChunkingStrategyResponseTypeStatic VectorStoreChunkingStrategyResponseType = "static"
)

// PossibleVectorStoreChunkingStrategyResponseTypeValues returns the possible values for the VectorStoreChunkingStrategyResponseType const type.
func PossibleVectorStoreChunkingStrategyResponseTypeValues() []VectorStoreChunkingStrategyResponseType {
	return []VectorStoreChunkingStrategyResponseType{
		VectorStoreChunkingStrategyResponseTypeOther,
		VectorStoreChunkingStrategyResponseTypeStatic,
	}
}

// VectorStoreExpirationPolicyAnchor - Describes the relationship between the days and the expiration of this vector store
type VectorStoreExpirationPolicyAnchor string

const (
	// VectorStoreExpirationPolicyAnchorLastActiveAt - The expiration policy is based on the last time the vector store was active.
	VectorStoreExpirationPolicyAnchorLastActiveAt VectorStoreExpirationPolicyAnchor = "last_active_at"
)

// PossibleVectorStoreExpirationPolicyAnchorValues returns the possible values for the VectorStoreExpirationPolicyAnchor const type.
func PossibleVectorStoreExpirationPolicyAnchorValues() []VectorStoreExpirationPolicyAnchor {
	return []VectorStoreExpirationPolicyAnchor{
		VectorStoreExpirationPolicyAnchorLastActiveAt,
	}
}

// VectorStoreFileBatchStatus - The status of the vector store file batch.
type VectorStoreFileBatchStatus string

const (
	// VectorStoreFileBatchStatusCancelled - The vector store file batch was cancelled.
	VectorStoreFileBatchStatusCancelled VectorStoreFileBatchStatus = "cancelled"
	// VectorStoreFileBatchStatusCompleted - the vector store file batch is ready for use.
	VectorStoreFileBatchStatusCompleted VectorStoreFileBatchStatus = "completed"
	// VectorStoreFileBatchStatusFailed - The vector store file batch failed to process.
	VectorStoreFileBatchStatusFailed VectorStoreFileBatchStatus = "failed"
	// VectorStoreFileBatchStatusInProgress - The vector store is still processing this file batch.
	VectorStoreFileBatchStatusInProgress VectorStoreFileBatchStatus = "in_progress"
)

// PossibleVectorStoreFileBatchStatusValues returns the possible values for the VectorStoreFileBatchStatus const type.
func PossibleVectorStoreFileBatchStatusValues() []VectorStoreFileBatchStatus {
	return []VectorStoreFileBatchStatus{
		VectorStoreFileBatchStatusCancelled,
		VectorStoreFileBatchStatusCompleted,
		VectorStoreFileBatchStatusFailed,
		VectorStoreFileBatchStatusInProgress,
	}
}

// VectorStoreFileErrorCode - Error code variants for vector store file processing
type VectorStoreFileErrorCode string

const (
	// VectorStoreFileErrorCodeFileNotFound - The file was not found.
	VectorStoreFileErrorCodeFileNotFound VectorStoreFileErrorCode = "file_not_found"
	// VectorStoreFileErrorCodeInternalError - An internal error occurred.
	VectorStoreFileErrorCodeInternalError VectorStoreFileErrorCode = "internal_error"
	// VectorStoreFileErrorCodeParsingError - The file could not be parsed.
	VectorStoreFileErrorCodeParsingError VectorStoreFileErrorCode = "parsing_error"
	// VectorStoreFileErrorCodeUnhandledMimeType - The file has an unhandled mime type.
	VectorStoreFileErrorCodeUnhandledMimeType VectorStoreFileErrorCode = "unhandled_mime_type"
)

// PossibleVectorStoreFileErrorCodeValues returns the possible values for the VectorStoreFileErrorCode const type.
func PossibleVectorStoreFileErrorCodeValues() []VectorStoreFileErrorCode {
	return []VectorStoreFileErrorCode{
		VectorStoreFileErrorCodeFileNotFound,
		VectorStoreFileErrorCodeInternalError,
		VectorStoreFileErrorCodeParsingError,
		VectorStoreFileErrorCodeUnhandledMimeType,
	}
}

// VectorStoreFileStatus - Vector store file status
type VectorStoreFileStatus string

const (
	// VectorStoreFileStatusCancelled - The file was cancelled.
	VectorStoreFileStatusCancelled VectorStoreFileStatus = "cancelled"
	// VectorStoreFileStatusCompleted - The file has been successfully processed.
	VectorStoreFileStatusCompleted VectorStoreFileStatus = "completed"
	// VectorStoreFileStatusFailed - The file has failed to process.
	VectorStoreFileStatusFailed VectorStoreFileStatus = "failed"
	// VectorStoreFileStatusInProgress - The file is currently being processed.
	VectorStoreFileStatusInProgress VectorStoreFileStatus = "in_progress"
)

// PossibleVectorStoreFileStatusValues returns the possible values for the VectorStoreFileStatus const type.
func PossibleVectorStoreFileStatusValues() []VectorStoreFileStatus {
	return []VectorStoreFileStatus{
		VectorStoreFileStatusCancelled,
		VectorStoreFileStatusCompleted,
		VectorStoreFileStatusFailed,
		VectorStoreFileStatusInProgress,
	}
}

// VectorStoreFileStatusFilter - Query parameter filter for vector store file retrieval endpoint
type VectorStoreFileStatusFilter string

const (
	// VectorStoreFileStatusFilterCancelled - Retrieve only files that were cancelled
	VectorStoreFileStatusFilterCancelled VectorStoreFileStatusFilter = "cancelled"
	// VectorStoreFileStatusFilterCompleted - Retrieve only files that have been successfully processed
	VectorStoreFileStatusFilterCompleted VectorStoreFileStatusFilter = "completed"
	// VectorStoreFileStatusFilterFailed - Retrieve only files that have failed to process
	VectorStoreFileStatusFilterFailed VectorStoreFileStatusFilter = "failed"
	// VectorStoreFileStatusFilterInProgress - Retrieve only files that are currently being processed
	VectorStoreFileStatusFilterInProgress VectorStoreFileStatusFilter = "in_progress"
)

// PossibleVectorStoreFileStatusFilterValues returns the possible values for the VectorStoreFileStatusFilter const type.
func PossibleVectorStoreFileStatusFilterValues() []VectorStoreFileStatusFilter {
	return []VectorStoreFileStatusFilter{
		VectorStoreFileStatusFilterCancelled,
		VectorStoreFileStatusFilterCompleted,
		VectorStoreFileStatusFilterFailed,
		VectorStoreFileStatusFilterInProgress,
	}
}

// VectorStoreStatus - Vector store possible status
type VectorStoreStatus string

const (
	// VectorStoreStatusCompleted - completed status indicates that this vector store is ready for use.
	VectorStoreStatusCompleted VectorStoreStatus = "completed"
	// VectorStoreStatusExpired - expired status indicates that this vector store has expired and is no longer available for use.
	VectorStoreStatusExpired VectorStoreStatus = "expired"
	// VectorStoreStatusInProgress - in_progress status indicates that this vector store is still processing files.
	VectorStoreStatusInProgress VectorStoreStatus = "in_progress"
)

// PossibleVectorStoreStatusValues returns the possible values for the VectorStoreStatus const type.
func PossibleVectorStoreStatusValues() []VectorStoreStatus {
	return []VectorStoreStatus{
		VectorStoreStatusCompleted,
		VectorStoreStatusExpired,
		VectorStoreStatusInProgress,
	}
}
