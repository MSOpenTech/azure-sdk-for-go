// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package config

import (
	"encoding/json"
)

type Config struct {
	Track2Requests   Track2ReleaseRequests   `json:"track2Requests,omitempty"`
	TypeSpecRequests TypeSpecReleaseRequests `json:"typespecRequests,omitempty"`
	RefreshInfo      RefreshInfo             `json:"refresh,omitempty"`
	AdditionalFlags  []string                `json:"additionalOptions,omitempty"`
}

func (c Config) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c Config) AdditionalOptions() ([]Option, error) {
	return parseAdditionalOptions(c.AdditionalFlags)
}
