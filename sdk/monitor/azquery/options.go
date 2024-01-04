//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azquery

// LogsClientQueryBatchOptions contains the optional parameters for the LogsClient.QueryBatch method.
type LogsClientQueryBatchOptions struct {
	// placeholder for future optional parameters
}

// LogsClientQueryResourceOptions contains the optional parameters for the LogsClient.QueryResource method.
type LogsClientQueryResourceOptions struct {
	// Optional. The prefer header to set server timeout, query statistics and visualization information.
	Options *LogsQueryOptions
}

// LogsClientQueryWorkspaceOptions contains the optional parameters for the LogsClient.QueryWorkspace method.
type LogsClientQueryWorkspaceOptions struct {
	// Optional. The prefer header to set server timeout, query statistics and visualization information.
	Options *LogsQueryOptions
}

// MetricsBatchClientQueryBatchOptions contains the optional parameters for the MetricsBatchClient.QueryBatch method.
type MetricsBatchClientQueryBatchOptions struct {
	// The list of aggregation types to retrieve
	Aggregation []*AggregationType

	// The end time of the query. It is a string in the format 'yyyy-MM-ddTHH:mm:ss.fffZ'.
	EndTime *string

	// The filter is used to reduce the set of metric data returned.
	// Example:
	// Metric contains metadata A, B and C.
	// - Return all time series of C where A = a1 and B = b1 or b2
	// filter=A eq ‘a1’ and B eq ‘b1’ or B eq ‘b2’ and C eq ‘’
	// - Invalid variant:
	// filter=A eq ‘a1’ and B eq ‘b1’ and C eq ‘’ or B = ‘b2’
	// This is invalid because the logical or operator cannot separate two different metadata names.
	// - Return all time series where A = a1, B = b1 and C = c1:
	// filter=A eq ‘a1’ and B eq ‘b1’ and C eq ‘c1’
	// - Return all time series where A = a1
	// filter=A eq ‘a1’ and B eq ‘’ and C eq ‘’.
	Filter *string

	// The interval (i.e. timegrain) of the query in ISO 8601 duration format. Defaults to PT1M. Special case for 'FULL' value
	// that returns single datapoint for entire time span requested.Examples: PT15M,
	// PT1H, P1D, FULL
	Interval *string

	// The aggregation to use for sorting results and the direction of the sort. Only one order can be specified.Examples: sum
	// asc
	OrderBy *string

	// Dimension name(s) to rollup results by. For example if you only want to see metric values with a filter like 'City eq Seattle
	// or City eq Tacoma' but don't want to see separate values for each city,
	// you can specify 'RollUpBy=City' to see the results for Seattle and Tacoma rolled up into one timeseries.
	RollUpBy *string

	// The start time of the query. It is a string in the format 'yyyy-MM-ddTHH:mm:ss.fffZ'. If you have specified the endtime
	// parameter, then this parameter is required. If only starttime is specified, then
	// endtime defaults to the current time. If no time interval is specified, the default is 1 hour.
	StartTime *string

	// The maximum number of records to retrieve per resource ID in the request. Valid only if filter is specified. Defaults to
	// 10.
	Top *int32
}

// MetricsClientListDefinitionsOptions contains the optional parameters for the MetricsClient.NewListDefinitionsPager method.
type MetricsClientListDefinitionsOptions struct {
	// Metric namespace where the metrics you want reside.
	MetricNamespace *string
}

// MetricsClientListNamespacesOptions contains the optional parameters for the MetricsClient.NewListNamespacesPager method.
type MetricsClientListNamespacesOptions struct {
	// The ISO 8601 conform Date start time from which to query for metric namespaces.
	StartTime *string
}

// MetricsClientQueryResourceOptions contains the optional parameters for the MetricsClient.QueryResource method.
type MetricsClientQueryResourceOptions struct {
	// The list of aggregation types to retrieve
	Aggregation []*AggregationType

	// When set to true, if the timespan passed in is not supported by this metric, the API will return the result using the closest
	// supported timespan. When set to false, an error is returned for invalid
	// timespan parameters. Defaults to false.
	AutoAdjustTimegrain *bool

	// The $filter is used to reduce the set of metric data returned.
	// Example:
	// Metric contains metadata A, B and C.
	// - Return all time series of C where A = a1 and B = b1 or b2
	// $filter=A eq ‘a1’ and B eq ‘b1’ or B eq ‘b2’ and C eq ‘’
	// - Invalid variant:
	// $filter=A eq ‘a1’ and B eq ‘b1’ and C eq ‘’ or B = ‘b2’
	// This is invalid because the logical or operator cannot separate two different metadata names.
	// - Return all time series where A = a1, B = b1 and C = c1:
	// $filter=A eq ‘a1’ and B eq ‘b1’ and C eq ‘c1’
	// - Return all time series where A = a1
	// $filter=A eq ‘a1’ and B eq ‘’ and C eq ‘’.
	Filter *string

	// The interval (i.e. timegrain) of the query in ISO 8601 duration format. Defaults to PT1M. Special case for 'FULL' value
	// that returns single datapoint for entire time span requested.Examples: PT15M,
	// PT1H, P1D, FULL
	Interval *string

	// The names of the metrics (comma separated) to retrieve.
	MetricNames *string

	// Metric namespace where the metrics you want reside.
	MetricNamespace *string

	// The aggregation to use for sorting results and the direction of the sort. Only one order can be specified.Examples: sum
	// asc
	OrderBy *string

	// Reduces the set of data collected. The syntax allowed depends on the operation. See the operation's description for details.
	ResultType *ResultType

	// Dimension name(s) to rollup results by. For example if you only want to see metric values with a filter like 'City eq Seattle
	// or City eq Tacoma' but don't want to see separate values for each city,
	// you can specify 'RollUpBy=City' to see the results for Seattle and Tacoma rolled up into one timeseries.
	RollUpBy *string

	// The timespan of the query. It is a string with the following format 'startDateTimeISO/endDateTimeISO'.
	Timespan *TimeInterval

	// The maximum number of records to retrieve per resource ID in the request. Valid only if filter is specified. Defaults to
	// 10.
	Top *int32

	// When set to false, invalid filter parameter values will be ignored. When set to true, an error is returned for invalid
	// filter parameters. Defaults to true.
	ValidateDimensions *bool
}
