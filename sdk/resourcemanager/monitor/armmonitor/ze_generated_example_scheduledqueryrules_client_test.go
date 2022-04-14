//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.LogSearchRuleResource{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.LogSearchRule{
				Description: to.Ptr("<description>"),
				Action: &armmonitor.AlertingAction{
					ODataType: to.Ptr("<odata-type>"),
					AznsAction: &armmonitor.AzNsActionGroup{
						ActionGroup:          []*string{},
						CustomWebhookPayload: to.Ptr("<custom-webhook-payload>"),
						EmailSubject:         to.Ptr("<email-subject>"),
					},
					Severity: to.Ptr(armmonitor.AlertSeverityOne),
					Trigger: &armmonitor.TriggerCondition{
						MetricTrigger: &armmonitor.LogMetricTrigger{
							MetricColumn:      to.Ptr("<metric-column>"),
							MetricTriggerType: to.Ptr(armmonitor.MetricTriggerTypeConsecutive),
							Threshold:         to.Ptr[float64](5),
							ThresholdOperator: to.Ptr(armmonitor.ConditionalOperatorGreaterThan),
						},
						Threshold:         to.Ptr[float64](3),
						ThresholdOperator: to.Ptr(armmonitor.ConditionalOperatorGreaterThan),
					},
				},
				Enabled: to.Ptr(armmonitor.EnabledTrue),
				Schedule: &armmonitor.Schedule{
					FrequencyInMinutes:  to.Ptr[int32](15),
					TimeWindowInMinutes: to.Ptr[int32](15),
				},
				Source: &armmonitor.Source{
					DataSourceID: to.Ptr("<data-source-id>"),
					Query:        to.Ptr("<query>"),
					QueryType:    to.Ptr(armmonitor.QueryTypeResultCount),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/getScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<rule-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/patchScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.LogSearchRuleResourcePatch{
			Properties: &armmonitor.LogSearchRulePatch{
				Enabled: to.Ptr(armmonitor.EnabledTrue),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/deleteScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<rule-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/listScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListBySubscription(&armmonitor.ScheduledQueryRulesClientListBySubscriptionOptions{Filter: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/listScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armmonitor.ScheduledQueryRulesClientListByResourceGroupOptions{Filter: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
