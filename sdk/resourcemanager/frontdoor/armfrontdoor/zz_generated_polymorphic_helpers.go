//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import "encoding/json"

func unmarshalRouteConfigurationClassification(rawMsg json.RawMessage) (RouteConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RouteConfigurationClassification
	switch m["@odata.type"] {
	case "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration":
		b = &ForwardingConfiguration{}
	case "#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration":
		b = &RedirectConfiguration{}
	default:
		b = &RouteConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRouteConfigurationClassificationArray(rawMsg json.RawMessage) ([]RouteConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RouteConfigurationClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRouteConfigurationClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalRouteConfigurationClassificationMap(rawMsg json.RawMessage) (map[string]RouteConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]RouteConfigurationClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalRouteConfigurationClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
