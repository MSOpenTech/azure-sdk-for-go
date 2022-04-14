//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_Get.json
func ExampleCostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewCostsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		&armdevtestlabs.CostsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.json
func ExampleCostsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewCostsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.LabCost{
			Properties: &armdevtestlabs.LabCostProperties{
				CurrencyCode:  to.Ptr("<currency-code>"),
				EndDateTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T23:59:59Z"); return t }()),
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00Z"); return t }()),
				TargetCost: &armdevtestlabs.TargetCostProperties{
					CostThresholds: []*armdevtestlabs.CostThresholdProperties{
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](25),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("<threshold-id>"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusEnabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](50),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusEnabled),
							ThresholdID:                  to.Ptr("<threshold-id>"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](75),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("<threshold-id>"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](100),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("<threshold-id>"),
						},
						{
							DisplayOnChart: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							PercentageThreshold: &armdevtestlabs.PercentageCostThresholdProperties{
								ThresholdValue: to.Ptr[float64](125),
							},
							SendNotificationWhenExceeded: to.Ptr(armdevtestlabs.CostThresholdStatusDisabled),
							ThresholdID:                  to.Ptr("<threshold-id>"),
						}},
					CycleEndDateTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T00:00:00.000Z"); return t }()),
					CycleStartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00.000Z"); return t }()),
					CycleType:          to.Ptr(armdevtestlabs.ReportingCycleTypeCalendarMonth),
					Status:             to.Ptr(armdevtestlabs.TargetCostStatusEnabled),
					Target:             to.Ptr[int32](100),
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
