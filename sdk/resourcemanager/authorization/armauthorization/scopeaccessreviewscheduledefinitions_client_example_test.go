//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewScheduleDefinitions.json
func ExampleScopeAccessReviewScheduleDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScopeAccessReviewScheduleDefinitionsClient().NewListPager("subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", &armauthorization.ScopeAccessReviewScheduleDefinitionsClientListOptions{Filter: nil})
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
		// page.AccessReviewScheduleDefinitionListResult = armauthorization.AccessReviewScheduleDefinitionListResult{
		// 	Value: []*armauthorization.AccessReviewScheduleDefinition{
		// 		{
		// 			Name: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
		// 			Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions"),
		// 			ID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
		// 			Properties: &armauthorization.AccessReviewScheduleDefinitionProperties{
		// 				CreatedBy: &armauthorization.AccessReviewActorIdentity{
		// 					PrincipalID: to.Ptr("a6c7aecb-cbfd-4763-87ef-e91b4bd509d9"),
		// 					PrincipalName: to.Ptr("Shubham Gupta"),
		// 					PrincipalType: to.Ptr(armauthorization.AccessReviewActorIdentityTypeUser),
		// 					UserPrincipalName: to.Ptr("shugup@microsoft.com"),
		// 				},
		// 				DescriptionForAdmins: to.Ptr("asdfasdf"),
		// 				DescriptionForReviewers: to.Ptr("asdfasdf"),
		// 				DisplayName: to.Ptr("Hello world"),
		// 				Reviewers: []*armauthorization.AccessReviewReviewer{
		// 					{
		// 						PrincipalID: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d "),
		// 						PrincipalType: to.Ptr(armauthorization.AccessReviewReviewerTypeUser),
		// 				}},
		// 				ReviewersType: to.Ptr(armauthorization.AccessReviewScheduleDefinitionReviewersTypeAssigned),
		// 				Scope: &armauthorization.AccessReviewScope{
		// 					ResourceID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
		// 					RoleDefinitionID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/roleDefinitions/b225c7ff-4338-4cdc-a790-6b34d987f7cd"),
		// 				},
		// 				Settings: &armauthorization.AccessReviewScheduleSettings{
		// 					AutoApplyDecisionsEnabled: to.Ptr(true),
		// 					DefaultDecision: to.Ptr(armauthorization.DefaultDecisionTypeApprove),
		// 					DefaultDecisionEnabled: to.Ptr(true),
		// 					InstanceDurationInDays: to.Ptr[int32](30),
		// 					JustificationRequiredOnApproval: to.Ptr(true),
		// 					MailNotificationsEnabled: to.Ptr(true),
		// 					RecommendationsEnabled: to.Ptr(true),
		// 					Recurrence: &armauthorization.AccessReviewRecurrenceSettings{
		// 						Range: &armauthorization.AccessReviewRecurrenceRange{
		// 							Type: to.Ptr(armauthorization.AccessReviewRecurrenceRangeTypeEndDate),
		// 							EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:17:30.513Z"); return t}()),
		// 							NumberOfOccurrences: to.Ptr[int32](1),
		// 							StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
		// 						},
		// 					},
		// 					ReminderNotificationsEnabled: to.Ptr(true),
		// 				},
		// 				Status: to.Ptr(armauthorization.AccessReviewScheduleDefinitionStatusInProgress),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewScheduleDefinition.json
func ExampleScopeAccessReviewScheduleDefinitionsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopeAccessReviewScheduleDefinitionsClient().GetByID(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessReviewScheduleDefinition = armauthorization.AccessReviewScheduleDefinition{
	// 	Name: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 	Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions"),
	// 	ID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 	Properties: &armauthorization.AccessReviewScheduleDefinitionProperties{
	// 		CreatedBy: &armauthorization.AccessReviewActorIdentity{
	// 			PrincipalID: to.Ptr("a6c7aecb-cbfd-4763-87ef-e91b4bd509d9"),
	// 			PrincipalName: to.Ptr("Shubham Gupta"),
	// 			PrincipalType: to.Ptr(armauthorization.AccessReviewActorIdentityTypeUser),
	// 			UserPrincipalName: to.Ptr("shugup@microsoft.com"),
	// 		},
	// 		DescriptionForAdmins: to.Ptr("asdfasdf"),
	// 		DescriptionForReviewers: to.Ptr("asdfasdf"),
	// 		DisplayName: to.Ptr("Hello world"),
	// 		Reviewers: []*armauthorization.AccessReviewReviewer{
	// 			{
	// 				PrincipalID: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d "),
	// 				PrincipalType: to.Ptr(armauthorization.AccessReviewReviewerTypeUser),
	// 		}},
	// 		ReviewersType: to.Ptr(armauthorization.AccessReviewScheduleDefinitionReviewersTypeAssigned),
	// 		Scope: &armauthorization.AccessReviewScope{
	// 			ResourceID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 			RoleDefinitionID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/roleDefinitions/b225c7ff-4338-4cdc-a790-6b34d987f7cd"),
	// 		},
	// 		Settings: &armauthorization.AccessReviewScheduleSettings{
	// 			AutoApplyDecisionsEnabled: to.Ptr(true),
	// 			DefaultDecision: to.Ptr(armauthorization.DefaultDecisionTypeApprove),
	// 			DefaultDecisionEnabled: to.Ptr(true),
	// 			InstanceDurationInDays: to.Ptr[int32](30),
	// 			JustificationRequiredOnApproval: to.Ptr(true),
	// 			MailNotificationsEnabled: to.Ptr(true),
	// 			RecommendationsEnabled: to.Ptr(true),
	// 			Recurrence: &armauthorization.AccessReviewRecurrenceSettings{
	// 				Range: &armauthorization.AccessReviewRecurrenceRange{
	// 					Type: to.Ptr(armauthorization.AccessReviewRecurrenceRangeTypeEndDate),
	// 					EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:17:30.513Z"); return t}()),
	// 					NumberOfOccurrences: to.Ptr[int32](1),
	// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
	// 				},
	// 			},
	// 			ReminderNotificationsEnabled: to.Ptr(true),
	// 		},
	// 		Status: to.Ptr(armauthorization.AccessReviewScheduleDefinitionStatusInProgress),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/DeleteAccessReviewScheduleDefinition.json
func ExampleScopeAccessReviewScheduleDefinitionsClient_DeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScopeAccessReviewScheduleDefinitionsClient().DeleteByID(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/PutAccessReviewScheduleDefinition.json
func ExampleScopeAccessReviewScheduleDefinitionsClient_CreateOrUpdateByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopeAccessReviewScheduleDefinitionsClient().CreateOrUpdateByID(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", armauthorization.AccessReviewScheduleDefinitionProperties{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessReviewScheduleDefinition = armauthorization.AccessReviewScheduleDefinition{
	// 	Name: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 	Type: to.Ptr("Microsoft.Authorization/accessReviewScheduleDefinitions"),
	// 	ID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 	Properties: &armauthorization.AccessReviewScheduleDefinitionProperties{
	// 		CreatedBy: &armauthorization.AccessReviewActorIdentity{
	// 			PrincipalID: to.Ptr("a6c7aecb-cbfd-4763-87ef-e91b4bd509d9"),
	// 			PrincipalName: to.Ptr("Shubham Gupta"),
	// 			PrincipalType: to.Ptr(armauthorization.AccessReviewActorIdentityTypeUser),
	// 			UserPrincipalName: to.Ptr("shugup@microsoft.com"),
	// 		},
	// 		DescriptionForAdmins: to.Ptr("asdfasdf"),
	// 		DescriptionForReviewers: to.Ptr("asdfasdf"),
	// 		DisplayName: to.Ptr("Hello world"),
	// 		Reviewers: []*armauthorization.AccessReviewReviewer{
	// 			{
	// 				PrincipalID: to.Ptr("fa73e90b-5bf1-45fd-a182-35ce5fc0674d "),
	// 				PrincipalType: to.Ptr(armauthorization.AccessReviewReviewerTypeUser),
	// 		}},
	// 		ReviewersType: to.Ptr(armauthorization.AccessReviewScheduleDefinitionReviewersTypeAssigned),
	// 		Scope: &armauthorization.AccessReviewScope{
	// 			ResourceID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d"),
	// 			RoleDefinitionID: to.Ptr("/subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d/providers/Microsoft.Authorization/roleDefinitions/b225c7ff-4338-4cdc-a790-6b34d987f7cd"),
	// 		},
	// 		Settings: &armauthorization.AccessReviewScheduleSettings{
	// 			AutoApplyDecisionsEnabled: to.Ptr(true),
	// 			DefaultDecision: to.Ptr(armauthorization.DefaultDecisionTypeApprove),
	// 			DefaultDecisionEnabled: to.Ptr(true),
	// 			InstanceDurationInDays: to.Ptr[int32](30),
	// 			JustificationRequiredOnApproval: to.Ptr(true),
	// 			MailNotificationsEnabled: to.Ptr(true),
	// 			RecommendationsEnabled: to.Ptr(true),
	// 			Recurrence: &armauthorization.AccessReviewRecurrenceSettings{
	// 				Range: &armauthorization.AccessReviewRecurrenceRange{
	// 					Type: to.Ptr(armauthorization.AccessReviewRecurrenceRangeTypeEndDate),
	// 					EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:17:30.513Z"); return t}()),
	// 					NumberOfOccurrences: to.Ptr[int32](1),
	// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-03T21:02:30.667Z"); return t}()),
	// 				},
	// 			},
	// 			ReminderNotificationsEnabled: to.Ptr(true),
	// 		},
	// 		Status: to.Ptr(armauthorization.AccessReviewScheduleDefinitionStatusInProgress),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/StopAccessReviewScheduleDefinition.json
func ExampleScopeAccessReviewScheduleDefinitionsClient_Stop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScopeAccessReviewScheduleDefinitionsClient().Stop(ctx, "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d", "fa73e90b-5bf1-45fd-a182-35ce5fc0674d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
