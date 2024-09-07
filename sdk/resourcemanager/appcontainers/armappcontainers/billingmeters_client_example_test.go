//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/BillingMeters_Get.json
func ExampleBillingMetersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBillingMetersClient().Get(ctx, "East US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BillingMeterCollection = armappcontainers.BillingMeterCollection{
	// 	Value: []*armappcontainers.BillingMeter{
	// 		{
	// 			Name: to.Ptr("GeneralPurposeDseriesCPU"),
	// 			Type: to.Ptr("Microsoft.App/billingMeters"),
	// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/GeneralPurposeDseriesCPU"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armappcontainers.BillingMeterProperties{
	// 				Category: to.Ptr("General purpose D-series"),
	// 				DisplayName: to.Ptr("General Purpose Cores per Second"),
	// 				MeterType: to.Ptr("CPU"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("GeneralPurposeDseriesMemory"),
	// 			Type: to.Ptr("Microsoft.App/billingMeters"),
	// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/GeneralPurposeDseriesMemory"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armappcontainers.BillingMeterProperties{
	// 				Category: to.Ptr("General purpose D-series"),
	// 				DisplayName: to.Ptr("General Purpose Memory GiB per Second"),
	// 				MeterType: to.Ptr("Memory"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("MemoryOptimizedEseriesCPU"),
	// 			Type: to.Ptr("Microsoft.App/billingMeters"),
	// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/MemoryOptimizedEseriesCPU"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armappcontainers.BillingMeterProperties{
	// 				Category: to.Ptr("Memory optimized E-series"),
	// 				DisplayName: to.Ptr("Memory Optimized Cores per Second"),
	// 				MeterType: to.Ptr("CPU"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("MemoryOptimizedEseriesMemory"),
	// 			Type: to.Ptr("Microsoft.App/billingMeters"),
	// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/MemoryOptimizedEseriesMemory"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armappcontainers.BillingMeterProperties{
	// 				Category: to.Ptr("Memory optimized E-series"),
	// 				DisplayName: to.Ptr("Memory Optimized Memory GiB per Second"),
	// 				MeterType: to.Ptr("Memory"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("ComputeOptimizedFseriesCPU"),
	// 			Type: to.Ptr("Microsoft.App/billingMeters"),
	// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/ComputeOptimizedFseriesCPU"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armappcontainers.BillingMeterProperties{
	// 				Category: to.Ptr("Compute optimized F-series"),
	// 				DisplayName: to.Ptr("Compute Optimized Cores per Second"),
	// 				MeterType: to.Ptr("CPU"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("GeneralComputeMemory"),
	// 			Type: to.Ptr("Microsoft.App/billingMeters"),
	// 			ID: to.Ptr("/subscriptions/55f240e3-3d66-44f6-8358-4e4f3d7a2e51/providers/Microsoft.App/GeneralComputeMemory"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armappcontainers.BillingMeterProperties{
	// 				Category: to.Ptr("Compute optimized F-series"),
	// 				DisplayName: to.Ptr("Compute Optimized Memory GiB per Second"),
	// 				MeterType: to.Ptr("Memory"),
	// 			},
	// 	}},
	// }
}
