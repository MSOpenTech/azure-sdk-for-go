//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Schedule/list.json
func ExampleSchedulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSchedulesClient().NewListPager("test-rg", "my-aml-workspace", &armmachinelearning.SchedulesClientListOptions{Skip: nil,
		ListViewType: nil,
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
		// page.ScheduleResourceArmPaginatedResult = armmachinelearning.ScheduleResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.Schedule{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armmachinelearning.ScheduleProperties{
		// 				Description: to.Ptr("string"),
		// 				Properties: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Action: &armmachinelearning.EndpointScheduleAction{
		// 					ActionType: to.Ptr(armmachinelearning.ScheduleActionTypeInvokeBatchEndpoint),
		// 					EndpointInvocationDefinition: map[string]any{
		// 						"00cd1396-a094-4d48-8d86-14c43a55a6af": nil,
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("string"),
		// 				IsEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armmachinelearning.ScheduleProvisioningStatusSucceeded),
		// 				Trigger: &armmachinelearning.CronTrigger{
		// 					EndTime: to.Ptr("string"),
		// 					StartTime: to.Ptr("string"),
		// 					TimeZone: to.Ptr("string"),
		// 					TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
		// 					Expression: to.Ptr("string"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Schedule/delete.json
func ExampleSchedulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSchedulesClient().BeginDelete(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Schedule/get.json
func ExampleSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSchedulesClient().Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armmachinelearning.Schedule{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
	// 	},
	// 	Properties: &armmachinelearning.ScheduleProperties{
	// 		Description: to.Ptr("string"),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Action: &armmachinelearning.EndpointScheduleAction{
	// 			ActionType: to.Ptr(armmachinelearning.ScheduleActionTypeInvokeBatchEndpoint),
	// 			EndpointInvocationDefinition: map[string]any{
	// 				"a108545b-def1-4c86-8e53-dbcb1de3a8bc": nil,
	// 			},
	// 		},
	// 		DisplayName: to.Ptr("string"),
	// 		IsEnabled: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armmachinelearning.ScheduleProvisioningStatusSucceeded),
	// 		Trigger: &armmachinelearning.CronTrigger{
	// 			EndTime: to.Ptr("string"),
	// 			StartTime: to.Ptr("string"),
	// 			TimeZone: to.Ptr("string"),
	// 			TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
	// 			Expression: to.Ptr("string"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Schedule/createOrUpdate.json
func ExampleSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSchedulesClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.Schedule{
		Properties: &armmachinelearning.ScheduleProperties{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			Action: &armmachinelearning.EndpointScheduleAction{
				ActionType: to.Ptr(armmachinelearning.ScheduleActionTypeInvokeBatchEndpoint),
				EndpointInvocationDefinition: map[string]any{
					"9965593e-526f-4b89-bb36-761138cf2794": nil,
				},
			},
			DisplayName: to.Ptr("string"),
			IsEnabled:   to.Ptr(false),
			Trigger: &armmachinelearning.CronTrigger{
				EndTime:     to.Ptr("string"),
				StartTime:   to.Ptr("string"),
				TimeZone:    to.Ptr("string"),
				TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
				Expression:  to.Ptr("string"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armmachinelearning.Schedule{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 	},
	// 	Properties: &armmachinelearning.ScheduleProperties{
	// 		Description: to.Ptr("string"),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Action: &armmachinelearning.EndpointScheduleAction{
	// 			ActionType: to.Ptr(armmachinelearning.ScheduleActionTypeInvokeBatchEndpoint),
	// 			EndpointInvocationDefinition: map[string]any{
	// 				"d77a9a9a-4bb5-4c0c-8a77-459be8b82b9f": nil,
	// 			},
	// 		},
	// 		DisplayName: to.Ptr("string"),
	// 		IsEnabled: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armmachinelearning.ScheduleProvisioningStatusSucceeded),
	// 		Trigger: &armmachinelearning.CronTrigger{
	// 			EndTime: to.Ptr("string"),
	// 			StartTime: to.Ptr("string"),
	// 			TimeZone: to.Ptr("string"),
	// 			TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
	// 			Expression: to.Ptr("string"),
	// 		},
	// 	},
	// }
}
