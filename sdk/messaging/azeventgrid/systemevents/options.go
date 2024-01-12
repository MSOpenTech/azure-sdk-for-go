//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package systemevents

// ClientDeleteMeClientPublishCloudEventsOptions contains the optional parameters for the ClientDeleteMeClient.PublishCloudEvents
// method.
type ClientDeleteMeClientPublishCloudEventsOptions struct {
	// Required only when publishing to partner namespaces with partner topic routing mode ChannelNameHeader.
	AegChannelName *string
}

// ClientDeleteMeClientPublishCustomEventEventsOptions contains the optional parameters for the ClientDeleteMeClient.PublishCustomEventEvents
// method.
type ClientDeleteMeClientPublishCustomEventEventsOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteMeClientPublishEventsOptions contains the optional parameters for the ClientDeleteMeClient.PublishEvents
// method.
type ClientDeleteMeClientPublishEventsOptions struct {
	// placeholder for future optional parameters
}
