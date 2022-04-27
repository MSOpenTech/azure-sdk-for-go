//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AuthorizedGroundstation.
func (a AuthorizedGroundstation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateDateType(objectMap, "expirationDate", a.ExpirationDate)
	populate(objectMap, "groundStation", a.GroundStation)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AuthorizedGroundstation.
func (a *AuthorizedGroundstation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationDate":
			err = unpopulateDateType(val, &a.ExpirationDate)
			delete(rawMsg, key)
		case "groundStation":
			err = unpopulate(val, &a.GroundStation)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AvailableContactsListResult.
func (a AvailableContactsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AvailableContactsProperties.
func (a AvailableContactsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endAzimuthDegrees", a.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", a.EndElevationDegrees)
	populate(objectMap, "maximumElevationDegrees", a.MaximumElevationDegrees)
	populateTimeRFC3339(objectMap, "rxEndTime", a.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", a.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", a.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", a.StartElevationDegrees)
	populateTimeRFC3339(objectMap, "txEndTime", a.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", a.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AvailableContactsProperties.
func (a *AvailableContactsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endAzimuthDegrees":
			err = unpopulate(val, &a.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &a.EndElevationDegrees)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &a.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &a.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &a.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &a.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &a.StartElevationDegrees)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &a.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &a.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AvailableGroundStationListResult.
func (a AvailableGroundStationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactInstanceProperties.
func (c ContactInstanceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endAzimuthDegrees", c.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", c.EndElevationDegrees)
	populate(objectMap, "maximumElevationDegrees", c.MaximumElevationDegrees)
	populateTimeRFC3339(objectMap, "rxEndTime", c.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", c.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", c.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", c.StartElevationDegrees)
	populateTimeRFC3339(objectMap, "txEndTime", c.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", c.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactInstanceProperties.
func (c *ContactInstanceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endAzimuthDegrees":
			err = unpopulate(val, &c.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &c.EndElevationDegrees)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &c.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &c.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &c.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &c.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &c.StartElevationDegrees)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &c.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &c.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContactListResult.
func (c ContactListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactParameters.
func (c ContactParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactProfile", c.ContactProfile)
	populateTimeRFC3339(objectMap, "endTime", c.EndTime)
	populate(objectMap, "groundStationName", c.GroundStationName)
	populateTimeRFC3339(objectMap, "startTime", c.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactParameters.
func (c *ContactParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contactProfile":
			err = unpopulate(val, &c.ContactProfile)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &c.EndTime)
			delete(rawMsg, key)
		case "groundStationName":
			err = unpopulate(val, &c.GroundStationName)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &c.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfile.
func (c ContactProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileLink.
func (c ContactProfileLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "channels", c.Channels)
	populate(objectMap, "direction", c.Direction)
	populate(objectMap, "eirpdBW", c.EirpdBW)
	populate(objectMap, "gainOverTemperature", c.GainOverTemperature)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "polarization", c.Polarization)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileListResult.
func (c ContactProfileListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfileProperties.
func (c ContactProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoTrackingConfiguration", c.AutoTrackingConfiguration)
	populate(objectMap, "eventHubUri", c.EventHubURI)
	populate(objectMap, "links", c.Links)
	populate(objectMap, "minimumElevationDegrees", c.MinimumElevationDegrees)
	populate(objectMap, "minimumViableContactDuration", c.MinimumViableContactDuration)
	populate(objectMap, "networkConfiguration", c.NetworkConfiguration)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfilesProperties.
func (c ContactProfilesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoTrackingConfiguration", c.AutoTrackingConfiguration)
	populate(objectMap, "eventHubUri", c.EventHubURI)
	populate(objectMap, "links", c.Links)
	populate(objectMap, "minimumElevationDegrees", c.MinimumElevationDegrees)
	populate(objectMap, "minimumViableContactDuration", c.MinimumViableContactDuration)
	populate(objectMap, "networkConfiguration", c.NetworkConfiguration)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContactsProperties.
func (c ContactsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "antennaConfiguration", c.AntennaConfiguration)
	populate(objectMap, "contactProfile", c.ContactProfile)
	populate(objectMap, "endAzimuthDegrees", c.EndAzimuthDegrees)
	populate(objectMap, "endElevationDegrees", c.EndElevationDegrees)
	populate(objectMap, "errorMessage", c.ErrorMessage)
	populate(objectMap, "groundStationName", c.GroundStationName)
	populate(objectMap, "maximumElevationDegrees", c.MaximumElevationDegrees)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populateTimeRFC3339(objectMap, "reservationEndTime", c.ReservationEndTime)
	populateTimeRFC3339(objectMap, "reservationStartTime", c.ReservationStartTime)
	populateTimeRFC3339(objectMap, "rxEndTime", c.RxEndTime)
	populateTimeRFC3339(objectMap, "rxStartTime", c.RxStartTime)
	populate(objectMap, "startAzimuthDegrees", c.StartAzimuthDegrees)
	populate(objectMap, "startElevationDegrees", c.StartElevationDegrees)
	populate(objectMap, "status", c.Status)
	populateTimeRFC3339(objectMap, "txEndTime", c.TxEndTime)
	populateTimeRFC3339(objectMap, "txStartTime", c.TxStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContactsProperties.
func (c *ContactsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "antennaConfiguration":
			err = unpopulate(val, &c.AntennaConfiguration)
			delete(rawMsg, key)
		case "contactProfile":
			err = unpopulate(val, &c.ContactProfile)
			delete(rawMsg, key)
		case "endAzimuthDegrees":
			err = unpopulate(val, &c.EndAzimuthDegrees)
			delete(rawMsg, key)
		case "endElevationDegrees":
			err = unpopulate(val, &c.EndElevationDegrees)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &c.ErrorMessage)
			delete(rawMsg, key)
		case "groundStationName":
			err = unpopulate(val, &c.GroundStationName)
			delete(rawMsg, key)
		case "maximumElevationDegrees":
			err = unpopulate(val, &c.MaximumElevationDegrees)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "reservationEndTime":
			err = unpopulateTimeRFC3339(val, &c.ReservationEndTime)
			delete(rawMsg, key)
		case "reservationStartTime":
			err = unpopulateTimeRFC3339(val, &c.ReservationStartTime)
			delete(rawMsg, key)
		case "rxEndTime":
			err = unpopulateTimeRFC3339(val, &c.RxEndTime)
			delete(rawMsg, key)
		case "rxStartTime":
			err = unpopulateTimeRFC3339(val, &c.RxStartTime)
			delete(rawMsg, key)
		case "startAzimuthDegrees":
			err = unpopulate(val, &c.StartAzimuthDegrees)
			delete(rawMsg, key)
		case "startElevationDegrees":
			err = unpopulate(val, &c.StartElevationDegrees)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &c.Status)
			delete(rawMsg, key)
		case "txEndTime":
			err = unpopulateTimeRFC3339(val, &c.TxEndTime)
			delete(rawMsg, key)
		case "txStartTime":
			err = unpopulateTimeRFC3339(val, &c.TxStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContactsPropertiesAntennaConfiguration.
func (c ContactsPropertiesAntennaConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "destinationIp", c.DestinationIP)
	populate(objectMap, "sourceIps", c.SourceIPs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationResult.
func (o OperationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", o.EndTime)
	populate(objectMap, "error", o.Error)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "percentComplete", o.PercentComplete)
	populate(objectMap, "properties", &o.Properties)
	populateTimeRFC3339(objectMap, "startTime", o.StartTime)
	populate(objectMap, "status", o.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResult.
func (o *OperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, &o.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, &o.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &o.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &o.Name)
			delete(rawMsg, key)
		case "percentComplete":
			err = unpopulate(val, &o.PercentComplete)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &o.Properties)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &o.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &o.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceIDListResult.
func (r ResourceIDListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Spacecraft.
func (s Spacecraft) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftLink.
func (s SpacecraftLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizations", s.Authorizations)
	populate(objectMap, "bandwidthMHz", s.BandwidthMHz)
	populate(objectMap, "centerFrequencyMHz", s.CenterFrequencyMHz)
	populate(objectMap, "direction", s.Direction)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "polarization", s.Polarization)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftListResult.
func (s SpacecraftListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SpacecraftsProperties.
func (s SpacecraftsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "links", s.Links)
	populate(objectMap, "noradId", s.NoradID)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "titleLine", s.TitleLine)
	populate(objectMap, "tleLine1", s.TleLine1)
	populate(objectMap, "tleLine2", s.TleLine2)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
