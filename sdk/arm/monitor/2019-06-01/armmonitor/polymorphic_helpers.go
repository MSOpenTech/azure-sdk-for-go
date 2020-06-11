// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import "encoding/json"

func unmarshalActionClassification(body []byte) (ActionClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b ActionClassification
	switch m["odata.type"] {
	case "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction":
		b = &AlertingAction{}
	case "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.LogToMetricAction":
		b = &LogToMetricAction{}
	default:
		b = &Action{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalActionClassificationArray(body []byte) (*[]ActionClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalActionClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalMetricAlertCriteriaClassification(body []byte) (MetricAlertCriteriaClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b MetricAlertCriteriaClassification
	switch m["odata.type"] {
	case OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria:
		b = &MetricAlertMultipleResourceMultipleMetricCriteria{}
	case OdatatypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria:
		b = &MetricAlertSingleResourceMultipleMetricCriteria{}
	case OdatatypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria:
		b = &WebtestLocationAvailabilityCriteria{}
	default:
		b = &MetricAlertCriteria{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalMetricAlertCriteriaClassificationArray(body []byte) (*[]MetricAlertCriteriaClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MetricAlertCriteriaClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMetricAlertCriteriaClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalMultiMetricCriteriaClassification(body []byte) (MultiMetricCriteriaClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b MultiMetricCriteriaClassification
	switch m["criterionType"] {
	case CriterionTypeDynamicThresholdCriterion:
		b = &DynamicMetricCriteria{}
	case CriterionTypeStaticThresholdCriterion:
		b = &MetricCriteria{}
	default:
		b = &MultiMetricCriteria{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalMultiMetricCriteriaClassificationArray(body []byte) (*[]MultiMetricCriteriaClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MultiMetricCriteriaClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMultiMetricCriteriaClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalRuleActionClassification(body []byte) (RuleActionClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b RuleActionClassification
	switch m["odata.type"] {
	case "Microsoft.Azure.Management.Insights.Models.RuleEmailAction":
		b = &RuleEmailAction{}
	case "Microsoft.Azure.Management.Insights.Models.RuleWebhookAction":
		b = &RuleWebhookAction{}
	default:
		b = &RuleAction{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalRuleActionClassificationArray(body []byte) (*[]RuleActionClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RuleActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRuleActionClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalRuleConditionClassification(body []byte) (RuleConditionClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b RuleConditionClassification
	switch m["odata.type"] {
	case "Microsoft.Azure.Management.Insights.Models.LocationThresholdRuleCondition":
		b = &LocationThresholdRuleCondition{}
	case "Microsoft.Azure.Management.Insights.Models.ManagementEventRuleCondition":
		b = &ManagementEventRuleCondition{}
	case "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition":
		b = &ThresholdRuleCondition{}
	default:
		b = &RuleCondition{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalRuleConditionClassificationArray(body []byte) (*[]RuleConditionClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RuleConditionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRuleConditionClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalRuleDataSourceClassification(body []byte) (RuleDataSourceClassification, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	var b RuleDataSourceClassification
	switch m["odata.type"] {
	case "Microsoft.Azure.Management.Insights.Models.RuleManagementEventDataSource":
		b = &RuleManagementEventDataSource{}
	case "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource":
		b = &RuleMetricDataSource{}
	default:
		b = &RuleDataSource{}
	}
	return b, json.Unmarshal(body, &b)
}

func unmarshalRuleDataSourceClassificationArray(body []byte) (*[]RuleDataSourceClassification, error) {
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(body, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RuleDataSourceClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRuleDataSourceClassification(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func strptr(s string) *string {
	return &s
}
