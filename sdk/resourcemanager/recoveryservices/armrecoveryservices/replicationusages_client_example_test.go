//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/ListReplicationUsages.json
func ExampleReplicationUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationUsagesClient().NewListPager("avrai7517RG1", "avrai7517Vault1", nil)
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
		// page.ReplicationUsageList = armrecoveryservices.ReplicationUsageList{
		// 	Value: []*armrecoveryservices.ReplicationUsage{
		// 		{
		// 			JobsSummary: &armrecoveryservices.JobsSummary{
		// 				FailedJobs: to.Ptr[int32](0),
		// 				InProgressJobs: to.Ptr[int32](0),
		// 				SuspendedJobs: to.Ptr[int32](0),
		// 			},
		// 			MonitoringSummary: &armrecoveryservices.MonitoringSummary{
		// 				DeprecatedProviderCount: to.Ptr[int32](0),
		// 				EventsCount: to.Ptr[int32](0),
		// 				SupportedProviderCount: to.Ptr[int32](0),
		// 				UnHealthyProviderCount: to.Ptr[int32](0),
		// 				UnHealthyVMCount: to.Ptr[int32](0),
		// 				UnsupportedProviderCount: to.Ptr[int32](0),
		// 			},
		// 			ProtectedItemCount: to.Ptr[int32](2),
		// 			RecoveryPlanCount: to.Ptr[int32](1),
		// 			RegisteredServersCount: to.Ptr[int32](2),
		// 	}},
		// }
	}
}
