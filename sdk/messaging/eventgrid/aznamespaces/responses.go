// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

// AcknowledgeEventsResponse contains the response from method ReceiverClient.AcknowledgeEvents.
type AcknowledgeEventsResponse struct {
	// The result of the Acknowledge operation.
	AcknowledgeEventsResult
}

// ReceiveEventsResponse contains the response from method ReceiverClient.ReceiveEvents.
type ReceiveEventsResponse struct {
	// Details of the Receive operation response.
	ReceiveEventsResult
}

// RejectEventsResponse contains the response from method ReceiverClient.RejectEvents.
type RejectEventsResponse struct {
	// The result of the Reject operation.
	RejectEventsResult
}

// ReleaseEventsResponse contains the response from method ReceiverClient.ReleaseEvents.
type ReleaseEventsResponse struct {
	// The result of the Release operation.
	ReleaseEventsResult
}

// RenewEventLocksResponse contains the response from method ReceiverClient.RenewEventLocks.
type RenewEventLocksResponse struct {
	// The result of the RenewLock operation.
	RenewEventLocksResult
}

// SendEventResponse contains the response from method SenderClient.SendEvent.
type SendEventResponse struct {
}

// SendEventsResponse contains the response from method SenderClient.SendEvents.
type SendEventsResponse struct {
}
