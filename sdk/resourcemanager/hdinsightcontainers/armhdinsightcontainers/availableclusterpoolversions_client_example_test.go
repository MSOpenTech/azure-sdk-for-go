//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListAvailableClusterPoolVersions.json
func ExampleAvailableClusterPoolVersionsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableClusterPoolVersionsClient().NewListByLocationPager("westus2", nil)
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
		// page.ClusterPoolVersionsListResult = armhdinsightcontainers.ClusterPoolVersionsListResult{
		// 	Value: []*armhdinsightcontainers.ClusterPoolVersion{
		// 		{
		// 			Name: to.Ptr("1.0"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HDInsight/locations/westus2/availableclusterpoolversions/1.0"),
		// 			Properties: &armhdinsightcontainers.ClusterPoolVersionProperties{
		// 				AksVersion: to.Ptr("1.24"),
		// 				ClusterPoolVersion: to.Ptr("1.0"),
		// 				IsPreview: to.Ptr(false),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1.1"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HDInsight/locations/westus2/availableclusterpoolversions/1.1"),
		// 			Properties: &armhdinsightcontainers.ClusterPoolVersionProperties{
		// 				AksVersion: to.Ptr("1.25"),
		// 				ClusterPoolVersion: to.Ptr("1.1"),
		// 				IsPreview: to.Ptr(false),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1.2"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HDInsight/locations/westus2/availableclusterpoolversions/1.2"),
		// 			Properties: &armhdinsightcontainers.ClusterPoolVersionProperties{
		// 				AksVersion: to.Ptr("1.26"),
		// 				ClusterPoolVersion: to.Ptr("1.2"),
		// 				IsPreview: to.Ptr(true),
		// 			},
		// 	}},
		// }
	}
}
