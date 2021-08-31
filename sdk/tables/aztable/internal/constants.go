//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

const (
	module = "internal"
	version = "v0.1.0"
)

// GeoReplicationStatusType - The status of the secondary location.
type GeoReplicationStatusType string

const (
	GeoReplicationStatusTypeBootstrap GeoReplicationStatusType = "bootstrap"
	GeoReplicationStatusTypeLive GeoReplicationStatusType = "live"
	GeoReplicationStatusTypeUnavailable GeoReplicationStatusType = "unavailable"
)

// PossibleGeoReplicationStatusTypeValues returns the possible values for the GeoReplicationStatusType const type.
func PossibleGeoReplicationStatusTypeValues() []GeoReplicationStatusType {
	return []GeoReplicationStatusType{	
		GeoReplicationStatusTypeBootstrap,
		GeoReplicationStatusTypeLive,
		GeoReplicationStatusTypeUnavailable,
	}
}

// ToPtr returns a *GeoReplicationStatusType pointing to the current value.
func (c GeoReplicationStatusType) ToPtr() *GeoReplicationStatusType {
	return &c
}

type ODataMetadataFormat string

const (
	ODataMetadataFormatApplicationJSONODataFullmetadata ODataMetadataFormat = "application/json;odata=fullmetadata"
	ODataMetadataFormatApplicationJSONODataMinimalmetadata ODataMetadataFormat = "application/json;odata=minimalmetadata"
	ODataMetadataFormatApplicationJSONODataNometadata ODataMetadataFormat = "application/json;odata=nometadata"
)

// PossibleODataMetadataFormatValues returns the possible values for the ODataMetadataFormat const type.
func PossibleODataMetadataFormatValues() []ODataMetadataFormat {
	return []ODataMetadataFormat{	
		ODataMetadataFormatApplicationJSONODataFullmetadata,
		ODataMetadataFormatApplicationJSONODataMinimalmetadata,
		ODataMetadataFormatApplicationJSONODataNometadata,
	}
}

// ToPtr returns a *ODataMetadataFormat pointing to the current value.
func (c ODataMetadataFormat) ToPtr() *ODataMetadataFormat {
	return &c
}

type ResponseFormat string

const (
	ResponseFormatReturnContent ResponseFormat = "return-content"
	ResponseFormatReturnNoContent ResponseFormat = "return-no-content"
)

// PossibleResponseFormatValues returns the possible values for the ResponseFormat const type.
func PossibleResponseFormatValues() []ResponseFormat {
	return []ResponseFormat{	
		ResponseFormatReturnContent,
		ResponseFormatReturnNoContent,
	}
}

// ToPtr returns a *ResponseFormat pointing to the current value.
func (c ResponseFormat) ToPtr() *ResponseFormat {
	return &c
}

