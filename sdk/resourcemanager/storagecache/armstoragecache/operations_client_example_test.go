//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.APIOperationListResult = armstoragecache.APIOperationListResult{
		// 	Value: []*armstoragecache.APIOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.StorageCache/caches/write"),
		// 			Display: &armstoragecache.APIOperationDisplay{
		// 				Operation: to.Ptr("Create or Update Cache"),
		// 				Provider: to.Ptr("Azure Storage Cache"),
		// 				Resource: to.Ptr("Caches"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorageCache/caches/delete"),
		// 			Display: &armstoragecache.APIOperationDisplay{
		// 				Operation: to.Ptr("Delete Cache"),
		// 				Provider: to.Ptr("Azure Storage Cache"),
		// 				Resource: to.Ptr("Caches"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorageCache/caches/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstoragecache.APIOperationDisplay{
		// 				Description: to.Ptr("Reads Cache Metric Definitions."),
		// 				Operation: to.Ptr("Get Cache Metric Definitions"),
		// 				Provider: to.Ptr("Microsoft Azure HPC Cache"),
		// 				Resource: to.Ptr("StorageCache Metric Definitions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 			Properties: &armstoragecache.APIOperationProperties{
		// 				ServiceSpecification: &armstoragecache.APIOperationPropertiesServiceSpecification{
		// 					MetricSpecifications: []*armstoragecache.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("ClientIOPS"),
		// 							AggregationType: to.Ptr("Average"),
		// 							DisplayDescription: to.Ptr("The rate of client file operations processed by the Cache."),
		// 							DisplayName: to.Ptr("Total Client IOPS"),
		// 							MetricClass: to.Ptr("Transactions"),
		// 							SupportedAggregationTypes: []*armstoragecache.MetricAggregationType{
		// 								to.Ptr(armstoragecache.MetricAggregationTypeMinimum),
		// 								to.Ptr(armstoragecache.MetricAggregationTypeMaximum),
		// 								to.Ptr(armstoragecache.MetricAggregationTypeAverage)},
		// 								Unit: to.Ptr("Count"),
		// 							},
		// 							{
		// 								Name: to.Ptr("ClientLatency"),
		// 								AggregationType: to.Ptr("Average"),
		// 								DisplayDescription: to.Ptr("Average latency of client file operations to the Cache."),
		// 								DisplayName: to.Ptr("Average Client Latency"),
		// 								MetricClass: to.Ptr("Latency"),
		// 								SupportedAggregationTypes: []*armstoragecache.MetricAggregationType{
		// 									to.Ptr(armstoragecache.MetricAggregationTypeMinimum),
		// 									to.Ptr(armstoragecache.MetricAggregationTypeMaximum),
		// 									to.Ptr(armstoragecache.MetricAggregationTypeAverage)},
		// 									Unit: to.Ptr("Milliseconds"),
		// 							}},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
