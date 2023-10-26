//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azwebpubsub

// ContentType - Content type for upload
type ContentType string

const (
	// ContentTypeApplicationJSON - Content Type 'application/json'
	ContentTypeApplicationJSON ContentType = "application/json"
	// ContentTypeApplicationOctetStream - Content Type 'application/octet-stream'
	ContentTypeApplicationOctetStream ContentType = "application/octet-stream"
	// ContentTypeTextPlain - Content Type 'text/plain'
	ContentTypeTextPlain ContentType = "text/plain"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{	
		ContentTypeApplicationJSON,
		ContentTypeApplicationOctetStream,
		ContentTypeTextPlain,
	}
}

type WebPubSubPermission string

const (
	WebPubSubPermissionJoinLeaveGroup WebPubSubPermission = "joinLeaveGroup"
	WebPubSubPermissionSendToGroup WebPubSubPermission = "sendToGroup"
)

// PossibleWebPubSubPermissionValues returns the possible values for the WebPubSubPermission const type.
func PossibleWebPubSubPermissionValues() []WebPubSubPermission {
	return []WebPubSubPermission{	
		WebPubSubPermissionJoinLeaveGroup,
		WebPubSubPermissionSendToGroup,
	}
}

