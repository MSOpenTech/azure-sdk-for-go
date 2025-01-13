// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
	"log"
)

// Generated from example definition: 2024-11-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armiotoperations.OperationsClientListResponse{
		// 	OperationListResult: armiotoperations.OperationListResult{
		// 		Value: []*armiotoperations.Operation{
		// 			{
		// 				Name: to.Ptr("xzxqfusky"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armiotoperations.OperationDisplay{
		// 					Provider: to.Ptr("lrveskajtuwf"),
		// 					Resource: to.Ptr("d"),
		// 					Operation: to.Ptr("icuckgobartrrgmirax"),
		// 					Description: to.Ptr("dsbfnxzvnoqdm"),
		// 				},
		// 				Origin: to.Ptr(armiotoperations.OriginUser),
		// 				ActionType: to.Ptr(armiotoperations.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
