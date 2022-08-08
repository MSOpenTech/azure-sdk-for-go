//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package monitor

const host = "https://api.loganalytics.io/v1"

type BatchQueryRequestMethod string

const (
	BatchQueryRequestMethodPOST BatchQueryRequestMethod = "POST"
)

// PossibleBatchQueryRequestMethodValues returns the possible values for the BatchQueryRequestMethod const type.
func PossibleBatchQueryRequestMethodValues() []BatchQueryRequestMethod {
	return []BatchQueryRequestMethod{
		BatchQueryRequestMethodPOST,
	}
}

type BatchQueryRequestPath string

const (
	BatchQueryRequestPathQuery BatchQueryRequestPath = "/query"
)

// PossibleBatchQueryRequestPathValues returns the possible values for the BatchQueryRequestPath const type.
func PossibleBatchQueryRequestPathValues() []BatchQueryRequestPath {
	return []BatchQueryRequestPath{
		BatchQueryRequestPathQuery,
	}
}

// LogsColumnType - The data type of this column.
type LogsColumnType string

const (
	LogsColumnTypeBool     LogsColumnType = "bool"
	LogsColumnTypeDatetime LogsColumnType = "datetime"
	LogsColumnTypeDecimal  LogsColumnType = "decimal"
	LogsColumnTypeDynamic  LogsColumnType = "dynamic"
	LogsColumnTypeGUID     LogsColumnType = "guid"
	LogsColumnTypeInt      LogsColumnType = "int"
	LogsColumnTypeLong     LogsColumnType = "long"
	LogsColumnTypeReal     LogsColumnType = "real"
	LogsColumnTypeString   LogsColumnType = "string"
	LogsColumnTypeTimespan LogsColumnType = "timespan"
)

// PossibleLogsColumnTypeValues returns the possible values for the LogsColumnType const type.
func PossibleLogsColumnTypeValues() []LogsColumnType {
	return []LogsColumnType{
		LogsColumnTypeBool,
		LogsColumnTypeDatetime,
		LogsColumnTypeDecimal,
		LogsColumnTypeDynamic,
		LogsColumnTypeGUID,
		LogsColumnTypeInt,
		LogsColumnTypeLong,
		LogsColumnTypeReal,
		LogsColumnTypeString,
		LogsColumnTypeTimespan,
	}
}

// MetadataColumnDataType - The data type of the column
type MetadataColumnDataType string

const (
	MetadataColumnDataTypeBool     MetadataColumnDataType = "bool"
	MetadataColumnDataTypeDatetime MetadataColumnDataType = "datetime"
	MetadataColumnDataTypeDecimal  MetadataColumnDataType = "decimal"
	MetadataColumnDataTypeDynamic  MetadataColumnDataType = "dynamic"
	MetadataColumnDataTypeGUID     MetadataColumnDataType = "guid"
	MetadataColumnDataTypeInt      MetadataColumnDataType = "int"
	MetadataColumnDataTypeLong     MetadataColumnDataType = "long"
	MetadataColumnDataTypeReal     MetadataColumnDataType = "real"
	MetadataColumnDataTypeString   MetadataColumnDataType = "string"
	MetadataColumnDataTypeTimespan MetadataColumnDataType = "timespan"
)

// PossibleMetadataColumnDataTypeValues returns the possible values for the MetadataColumnDataType const type.
func PossibleMetadataColumnDataTypeValues() []MetadataColumnDataType {
	return []MetadataColumnDataType{
		MetadataColumnDataTypeBool,
		MetadataColumnDataTypeDatetime,
		MetadataColumnDataTypeDecimal,
		MetadataColumnDataTypeDynamic,
		MetadataColumnDataTypeGUID,
		MetadataColumnDataTypeInt,
		MetadataColumnDataTypeLong,
		MetadataColumnDataTypeReal,
		MetadataColumnDataTypeString,
		MetadataColumnDataTypeTimespan,
	}
}
