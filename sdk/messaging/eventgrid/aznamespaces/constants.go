// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

// ReleaseDelay - Supported delays for release operation.
type ReleaseDelay string

const (
	// ReleaseDelayNoDelay - Release the event after 0 seconds.
	ReleaseDelayNoDelay ReleaseDelay = "0"
	// ReleaseDelayOneHour - Release the event after 3600 seconds.
	ReleaseDelayOneHour ReleaseDelay = "3600"
	// ReleaseDelayOneMinute - Release the event after 60 seconds.
	ReleaseDelayOneMinute ReleaseDelay = "60"
	// ReleaseDelayTenMinutes - Release the event after 600 seconds.
	ReleaseDelayTenMinutes ReleaseDelay = "600"
	// ReleaseDelayTenSeconds - Release the event after 10 seconds.
	ReleaseDelayTenSeconds ReleaseDelay = "10"
)

// PossibleReleaseDelayValues returns the possible values for the ReleaseDelay const type.
func PossibleReleaseDelayValues() []ReleaseDelay {
	return []ReleaseDelay{
		ReleaseDelayNoDelay,
		ReleaseDelayOneHour,
		ReleaseDelayOneMinute,
		ReleaseDelayTenMinutes,
		ReleaseDelayTenSeconds,
	}
}
