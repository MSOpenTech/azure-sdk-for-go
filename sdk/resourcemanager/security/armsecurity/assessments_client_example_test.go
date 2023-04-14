//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/ListAssessments_example.json
func ExampleAssessmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessmentsClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", nil)
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
		// page.AssessmentList = armsecurity.AssessmentList{
		// 	Value: []*armsecurity.AssessmentResponse{
		// 		{
		// 			Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Type: to.Ptr("Microsoft.Security/assessments"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Properties: &armsecurity.AssessmentPropertiesResponse{
		// 				DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
		// 				ResourceDetails: &armsecurity.AzureResourceDetails{
		// 					Source: to.Ptr(armsecurity.SourceAzure),
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1"),
		// 				},
		// 				Status: &armsecurity.AssessmentStatusResponse{
		// 					Code: to.Ptr(armsecurity.AssessmentStatusCodeHealthy),
		// 					FirstEvaluationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
		// 					StatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Type: to.Ptr("Microsoft.Security/assessments"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2/providers/Microsoft.Security/assessments/21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Properties: &armsecurity.AssessmentPropertiesResponse{
		// 				AdditionalData: map[string]*string{
		// 					"linkedWorkspaceId": to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myLaWorkspace"),
		// 				},
		// 				DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
		// 				ResourceDetails: &armsecurity.AzureResourceDetails{
		// 					Source: to.Ptr(armsecurity.SourceAzure),
		// 					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2"),
		// 				},
		// 				Status: &armsecurity.AssessmentStatusResponse{
		// 					Description: to.Ptr("The effective policy for the assessment was evaluated to off - use Microsoft.Authorization/policyAssignments to turn this assessment on"),
		// 					Cause: to.Ptr("OffByPolicy"),
		// 					Code: to.Ptr(armsecurity.AssessmentStatusCodeNotApplicable),
		// 					FirstEvaluationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
		// 					StatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/GetAssessment_example.json
func ExampleAssessmentsClient_Get_getSecurityRecommendationTaskFromSecurityDataLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "21300918-b2e3-0346-785f-c77ff57d243b", &armsecurity.AssessmentsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentResponse = armsecurity.AssessmentResponse{
	// 	Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Type: to.Ptr("Microsoft.Security/assessments"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2/providers/Microsoft.Security/assessments/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Properties: &armsecurity.AssessmentPropertiesResponse{
	// 		AdditionalData: map[string]*string{
	// 			"linkedWorkspaceId": to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myLaWorkspace"),
	// 		},
	// 		DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 		ResourceDetails: &armsecurity.AzureResourceDetails{
	// 			Source: to.Ptr(armsecurity.SourceAzure),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2"),
	// 		},
	// 		Status: &armsecurity.AssessmentStatusResponse{
	// 			Description: to.Ptr("The effective policy for the assessment was evaluated to off - use Microsoft.Authorization/policyAssignments to turn this assessment on"),
	// 			Cause: to.Ptr("OffByPolicy"),
	// 			Code: to.Ptr(armsecurity.AssessmentStatusCodeNotApplicable),
	// 			FirstEvaluationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
	// 			StatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/GetAssessmentWithExpand_example.json
func ExampleAssessmentsClient_Get_getSecurityRecommendationTaskFromSecurityDataLocationWithExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "21300918-b2e3-0346-785f-c77ff57d243b", &armsecurity.AssessmentsClientGetOptions{Expand: to.Ptr(armsecurity.ExpandEnumLinks)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentResponse = armsecurity.AssessmentResponse{
	// 	Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Type: to.Ptr("Microsoft.Security/assessments"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2/providers/Microsoft.Security/assessments/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Properties: &armsecurity.AssessmentPropertiesResponse{
	// 		AdditionalData: map[string]*string{
	// 			"linkedWorkspaceId": to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myLaWorkspace"),
	// 		},
	// 		DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 		Links: &armsecurity.AssessmentLinks{
	// 			AzurePortalURI: to.Ptr("https://www.portal.azure.com/?fea#blade/Microsoft_Azure_Security/RecommendationsBlade/assessmentKey/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 		},
	// 		ResourceDetails: &armsecurity.AzureResourceDetails{
	// 			Source: to.Ptr(armsecurity.SourceAzure),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2"),
	// 		},
	// 		Status: &armsecurity.AssessmentStatusResponse{
	// 			Description: to.Ptr("The effective policy for the assessment was evaluated to off - use Microsoft.Authorization/policyAssignments to turn this assessment on"),
	// 			Cause: to.Ptr("OffByPolicy"),
	// 			Code: to.Ptr(armsecurity.AssessmentStatusCodeNotApplicable),
	// 			FirstEvaluationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
	// 			StatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-12T09:07:18.6759138Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/PutAssessment_example.json
func ExampleAssessmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().CreateOrUpdate(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "8bb8be0a-6010-4789-812f-e4d661c4ed0e", armsecurity.Assessment{
		Properties: &armsecurity.AssessmentProperties{
			ResourceDetails: &armsecurity.AzureResourceDetails{
				Source: to.Ptr(armsecurity.SourceAzure),
			},
			Status: &armsecurity.AssessmentStatus{
				Code: to.Ptr(armsecurity.AssessmentStatusCodeHealthy),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentResponse = armsecurity.AssessmentResponse{
	// 	Name: to.Ptr("8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
	// 	Type: to.Ptr("Microsoft.Security/assessments"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
	// 	Properties: &armsecurity.AssessmentPropertiesResponse{
	// 		DisplayName: to.Ptr("Install internal agent on VM"),
	// 		ResourceDetails: &armsecurity.AzureResourceDetails{
	// 			Source: to.Ptr(armsecurity.SourceAzure),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
	// 		},
	// 		Status: &armsecurity.AssessmentStatusResponse{
	// 			Code: to.Ptr(armsecurity.AssessmentStatusCodeHealthy),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/DeleteAssessment_example.json
func ExampleAssessmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssessmentsClient().Delete(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "8bb8be0a-6010-4789-812f-e4d661c4ed0e", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
