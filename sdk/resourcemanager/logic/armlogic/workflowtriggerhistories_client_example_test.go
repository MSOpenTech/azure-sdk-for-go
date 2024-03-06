//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_List.json
func ExampleWorkflowTriggerHistoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowTriggerHistoriesClient().NewListPager("testResourceGroup", "testWorkflowName", "testTriggerName", &armlogic.WorkflowTriggerHistoriesClientListOptions{Top: nil,
		Filter: nil,
	})
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
		// page.WorkflowTriggerHistoryListResult = armlogic.WorkflowTriggerHistoryListResult{
		// 	Value: []*armlogic.WorkflowTriggerHistory{
		// 		{
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
		// 			Name: to.Ptr("08586676746934337772206998657CU22"),
		// 			Type: to.Ptr("Microsoft.Logic/workflows/triggers/histories"),
		// 			Properties: &armlogic.WorkflowTriggerHistoryProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armlogic.Correlation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.298Z"); return t}()),
		// 				Fired: to.Ptr(true),
		// 				Run: &armlogic.ResourceReference{
		// 					Name: to.Ptr("08586676746934337772206998657CU22"),
		// 					Type: to.Ptr("Microsoft.Logic/workflows/runs"),
		// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
		// 				},
		// 				ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
		// 				Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_Get.json
func ExampleWorkflowTriggerHistoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggerHistoriesClient().Get(ctx, "testResourceGroup", "testWorkflowName", "testTriggerName", "08586676746934337772206998657CU22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerHistory = armlogic.WorkflowTriggerHistory{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
	// 	Name: to.Ptr("08586676746934337772206998657CU22"),
	// 	Type: to.Ptr("Microsoft.Logic/workflows/triggers/histories"),
	// 	Properties: &armlogic.WorkflowTriggerHistoryProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armlogic.Correlation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.298Z"); return t}()),
	// 		Fired: to.Ptr(true),
	// 		Run: &armlogic.ResourceReference{
	// 			Name: to.Ptr("08586676746934337772206998657CU22"),
	// 			Type: to.Ptr("Microsoft.Logic/workflows/runs"),
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
	// 		},
	// 		ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
	// 		Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_Resubmit.json
func ExampleWorkflowTriggerHistoriesClient_Resubmit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkflowTriggerHistoriesClient().Resubmit(ctx, "testResourceGroup", "testWorkflowName", "testTriggerName", "testHistoryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
