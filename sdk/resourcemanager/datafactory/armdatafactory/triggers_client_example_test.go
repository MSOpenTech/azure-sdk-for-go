//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_ListByFactory.json
func ExampleTriggersClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTriggersClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.TriggerListResponse = armdatafactory.TriggerListResponse{
		// 	Value: []*armdatafactory.TriggerResource{
		// 		{
		// 			Name: to.Ptr("exampleTrigger"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
		// 			Etag: to.Ptr("0a008ed4-0000-0000-0000-5b245c740000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
		// 			Properties: &armdatafactory.ScheduleTrigger{
		// 				Type: to.Ptr("ScheduleTrigger"),
		// 				Description: to.Ptr("Example description"),
		// 				RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStarted),
		// 				Pipelines: []*armdatafactory.TriggerPipelineReference{
		// 					{
		// 						Parameters: map[string]any{
		// 							"OutputBlobNameList": []any{
		// 								"exampleoutput.csv",
		// 							},
		// 						},
		// 						PipelineReference: &armdatafactory.PipelineReference{
		// 							Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
		// 							ReferenceName: to.Ptr("examplePipeline"),
		// 						},
		// 				}},
		// 				TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
		// 					Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t}()),
		// 						Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
		// 						Interval: to.Ptr[int32](4),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t}()),
		// 						TimeZone: to.Ptr("UTC"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_QueryByFactory.json
func ExampleTriggersClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().QueryByFactory(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.TriggerFilterParameters{
		ParentTriggerName: to.Ptr("exampleTrigger"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerQueryResponse = armdatafactory.TriggerQueryResponse{
	// 	Value: []*armdatafactory.TriggerResource{
	// 		{
	// 			Name: to.Ptr("exampleRerunTrigger"),
	// 			Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
	// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleRerunTrigger"),
	// 			Properties: &armdatafactory.RerunTumblingWindowTrigger{
	// 				Type: to.Ptr("RerunTumblingWindowTrigger"),
	// 				Description: to.Ptr("Example description"),
	// 				TypeProperties: &armdatafactory.RerunTumblingWindowTriggerTypeProperties{
	// 					ParentTrigger: "exampleTrigger",
	// 					RequestedEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t}()),
	// 					RequestedStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t}()),
	// 					RerunConcurrency: to.Ptr[int32](4),
	// 				},
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
func ExampleTriggersClient_CreateOrUpdate_triggersCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", armdatafactory.TriggerResource{
		Properties: &armdatafactory.ScheduleTrigger{
			Type: to.Ptr("ScheduleTrigger"),
			Pipelines: []*armdatafactory.TriggerPipelineReference{
				{
					Parameters: map[string]any{
						"OutputBlobNameList": []any{
							"exampleoutput.csv",
						},
					},
					PipelineReference: &armdatafactory.PipelineReference{
						Type:          to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
						ReferenceName: to.Ptr("examplePipeline"),
					},
				}},
			TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
				Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
					EndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:13.844Z"); return t }()),
					Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
					Interval:  to.Ptr[int32](4),
					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:13.844Z"); return t }()),
					TimeZone:  to.Ptr("UTC"),
				},
			},
		},
	}, &armdatafactory.TriggersClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerResource = armdatafactory.TriggerResource{
	// 	Name: to.Ptr("exampleTrigger"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
	// 	Etag: to.Ptr("0a008ad4-0000-0000-0000-5b245c6e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
	// 	Properties: &armdatafactory.ScheduleTrigger{
	// 		Type: to.Ptr("ScheduleTrigger"),
	// 		RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStopped),
	// 		Pipelines: []*armdatafactory.TriggerPipelineReference{
	// 			{
	// 				Parameters: map[string]any{
	// 					"OutputBlobNameList": []any{
	// 						"exampleoutput.csv",
	// 					},
	// 				},
	// 				PipelineReference: &armdatafactory.PipelineReference{
	// 					Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
	// 					ReferenceName: to.Ptr("examplePipeline"),
	// 				},
	// 		}},
	// 		TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
	// 			Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:13.844Z"); return t}()),
	// 				Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
	// 				Interval: to.Ptr[int32](4),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:13.844Z"); return t}()),
	// 				TimeZone: to.Ptr("UTC"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Update.json
func ExampleTriggersClient_CreateOrUpdate_triggersUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", armdatafactory.TriggerResource{
		Properties: &armdatafactory.ScheduleTrigger{
			Type:        to.Ptr("ScheduleTrigger"),
			Description: to.Ptr("Example description"),
			Pipelines: []*armdatafactory.TriggerPipelineReference{
				{
					Parameters: map[string]any{
						"OutputBlobNameList": []any{
							"exampleoutput.csv",
						},
					},
					PipelineReference: &armdatafactory.PipelineReference{
						Type:          to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
						ReferenceName: to.Ptr("examplePipeline"),
					},
				}},
			TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
				Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
					EndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t }()),
					Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
					Interval:  to.Ptr[int32](4),
					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t }()),
					TimeZone:  to.Ptr("UTC"),
				},
			},
		},
	}, &armdatafactory.TriggersClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerResource = armdatafactory.TriggerResource{
	// 	Name: to.Ptr("exampleTrigger"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
	// 	Etag: to.Ptr("0a008dd4-0000-0000-0000-5b245c6f0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
	// 	Properties: &armdatafactory.ScheduleTrigger{
	// 		Type: to.Ptr("ScheduleTrigger"),
	// 		Description: to.Ptr("Example description"),
	// 		RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStopped),
	// 		Pipelines: []*armdatafactory.TriggerPipelineReference{
	// 			{
	// 				Parameters: map[string]any{
	// 					"OutputBlobNameList": []any{
	// 						"exampleoutput.csv",
	// 					},
	// 				},
	// 				PipelineReference: &armdatafactory.PipelineReference{
	// 					Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
	// 					ReferenceName: to.Ptr("examplePipeline"),
	// 				},
	// 		}},
	// 		TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
	// 			Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:55:14.905Z"); return t}()),
	// 				Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
	// 				Interval: to.Ptr[int32](4),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:39:14.905Z"); return t}()),
	// 				TimeZone: to.Ptr("UTC"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Get.json
func ExampleTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", &armdatafactory.TriggersClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerResource = armdatafactory.TriggerResource{
	// 	Name: to.Ptr("exampleTrigger"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/triggers"),
	// 	Etag: to.Ptr("1500544f-0000-0200-0000-5cbe09100000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/triggers/exampleTrigger"),
	// 	Properties: &armdatafactory.ScheduleTrigger{
	// 		Type: to.Ptr("ScheduleTrigger"),
	// 		RuntimeState: to.Ptr(armdatafactory.TriggerRuntimeStateStopped),
	// 		Pipelines: []*armdatafactory.TriggerPipelineReference{
	// 			{
	// 				Parameters: map[string]any{
	// 					"OutputBlobNameList": []any{
	// 						"exampleoutput.csv",
	// 					},
	// 				},
	// 				PipelineReference: &armdatafactory.PipelineReference{
	// 					Type: to.Ptr(armdatafactory.PipelineReferenceTypePipelineReference),
	// 					ReferenceName: to.Ptr("examplePipeline"),
	// 				},
	// 		}},
	// 		TypeProperties: &armdatafactory.ScheduleTriggerTypeProperties{
	// 			Recurrence: &armdatafactory.ScheduleTriggerRecurrence{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-22T18:48:52.528Z"); return t}()),
	// 				Frequency: to.Ptr(armdatafactory.RecurrenceFrequencyMinute),
	// 				Interval: to.Ptr[int32](4),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-22T18:32:52.527Z"); return t}()),
	// 				TimeZone: to.Ptr("UTC"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Delete.json
func ExampleTriggersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTriggersClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_SubscribeToEvents.json
func ExampleTriggersClient_BeginSubscribeToEvents() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginSubscribeToEvents(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", nil)
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
	// res.TriggerSubscriptionOperationStatus = armdatafactory.TriggerSubscriptionOperationStatus{
	// 	Status: to.Ptr(armdatafactory.EventSubscriptionStatusEnabled),
	// 	TriggerName: to.Ptr("exampleTrigger"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_GetEventSubscriptionStatus.json
func ExampleTriggersClient_GetEventSubscriptionStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().GetEventSubscriptionStatus(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerSubscriptionOperationStatus = armdatafactory.TriggerSubscriptionOperationStatus{
	// 	Status: to.Ptr(armdatafactory.EventSubscriptionStatusEnabled),
	// 	TriggerName: to.Ptr("exampleTrigger"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_UnsubscribeFromEvents.json
func ExampleTriggersClient_BeginUnsubscribeFromEvents() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginUnsubscribeFromEvents(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", nil)
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
	// res.TriggerSubscriptionOperationStatus = armdatafactory.TriggerSubscriptionOperationStatus{
	// 	Status: to.Ptr(armdatafactory.EventSubscriptionStatusDisabled),
	// 	TriggerName: to.Ptr("exampleTrigger"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Start.json
func ExampleTriggersClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginStart(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Stop.json
func ExampleTriggersClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginStop(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleTrigger", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
