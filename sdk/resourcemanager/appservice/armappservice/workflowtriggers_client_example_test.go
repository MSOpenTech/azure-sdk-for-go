//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggers_List.json
func ExampleWorkflowTriggersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowTriggersClient().NewListPager("test-resource-group", "test-name", "test-workflow", &armappservice.WorkflowTriggersClientListOptions{Top: nil,
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
		// page.WorkflowTriggerListResult = armappservice.WorkflowTriggerListResult{
		// 	Value: []*armappservice.WorkflowTrigger{
		// 		{
		// 			ID: to.Ptr("/workflows/test-workflow/triggers/manual"),
		// 			Name: to.Ptr("manual"),
		// 			Type: to.Ptr("/workflows/triggers"),
		// 			Properties: &armappservice.WorkflowTriggerProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:47:49.528Z"); return t}()),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T17:32:30.249Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armappservice.WorkflowTriggerProvisioningStateSucceeded),
		// 				State: to.Ptr(armappservice.WorkflowStateEnabled),
		// 				Workflow: &armappservice.ResourceReference{
		// 					Name: to.Ptr("08586676800160476478"),
		// 					Type: to.Ptr("/workflows/versions"),
		// 					ID: to.Ptr("/workflows/test-workflow/versions/08586676800160476478"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggers_Get.json
func ExampleWorkflowTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggersClient().Get(ctx, "test-resource-group", "test-name", "test-workflow", "manual", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTrigger = armappservice.WorkflowTrigger{
	// 	ID: to.Ptr("/workflows/test-workflow/triggers/manual"),
	// 	Name: to.Ptr("manual"),
	// 	Type: to.Ptr("/workflows/triggers"),
	// 	Properties: &armappservice.WorkflowTriggerProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:47:49.528Z"); return t}()),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T17:32:30.249Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armappservice.WorkflowTriggerProvisioningStateSucceeded),
	// 		State: to.Ptr(armappservice.WorkflowStateEnabled),
	// 		Workflow: &armappservice.ResourceReference{
	// 			Name: to.Ptr("08586676800160476478"),
	// 			Type: to.Ptr("/workflows/versions"),
	// 			ID: to.Ptr("/workflows/test-workflow/versions/08586676800160476478"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggers_ListCallbackUrl.json
func ExampleWorkflowTriggersClient_ListCallbackURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggersClient().ListCallbackURL(ctx, "test-resource-group", "test-name", "test-workflow", "manual", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerCallbackURL = armappservice.WorkflowTriggerCallbackURL{
	// 	Method: to.Ptr("POST"),
	// 	BasePath: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/fb9c8d79b15f41ce9b12861862f43546/versions/08587100027316071865/triggers/manualTrigger/paths/invoke"),
	// 	Queries: &armappservice.WorkflowTriggerListCallbackURLQueries{
	// 		APIVersion: to.Ptr("2018-07-01-preview"),
	// 		Sig: to.Ptr("IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk"),
	// 		Sp: to.Ptr("/versions/08587100027316071865/triggers/manualTrigger/run"),
	// 		Sv: to.Ptr("1.0"),
	// 	},
	// 	Value: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/fb9c8d79b15f41ce9b12861862f43546/versions/08587100027316071865/triggers/manualTrigger/paths/invoke?api-version=2015-08-01-preview&sp=%2Fversions%2F08587100027316071865%2Ftriggers%2FmanualTrigger%2Frun&sv=1.0&sig=IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggers_Run.json
func ExampleWorkflowTriggersClient_BeginRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkflowTriggersClient().BeginRun(ctx, "test-resource-group", "test-name", "test-workflow", "recurrence", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowTriggers_GetSchemaJson.json
func ExampleWorkflowTriggersClient_GetSchemaJSON() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggersClient().GetSchemaJSON(ctx, "testResourceGroup", "test-name", "testWorkflow", "testTrigger", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JSONSchema = armappservice.JSONSchema{
	// 	Content: to.Ptr("JsonContent"),
	// 	Title: to.Ptr("JsonTitle"),
	// }
}
