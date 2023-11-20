//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/deletePolicyExemption.json
func ExampleExemptionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewExemptionsClient().Delete(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/createOrUpdatePolicyExemption.json
func ExampleExemptionsClient_CreateOrUpdate_createOrUpdateAPolicyExemption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().CreateOrUpdate(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", armpolicy.Exemption{
		Properties: &armpolicy.ExemptionProperties{
			Description:       to.Ptr("Exempt demo cluster from limit sku"),
			DisplayName:       to.Ptr("Exempt demo cluster"),
			ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
			Metadata: map[string]any{
				"reason": "Temporary exemption for a expensive VM demo",
			},
			PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
			PolicyDefinitionReferenceIDs: []*string{
				to.Ptr("Limit_Skus")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Exemption = armpolicy.Exemption{
	// 	Name: to.Ptr("DemoExpensiveVM"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 	Properties: &armpolicy.ExemptionProperties{
	// 		Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 		DisplayName: to.Ptr("Exempt demo cluster"),
	// 		ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 		Metadata: map[string]any{
	// 			"reason": "Temporary exemption for a expensive VM demo",
	// 		},
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		PolicyDefinitionReferenceIDs: []*string{
	// 			to.Ptr("Limit_Skus")},
	// 		},
	// 		SystemData: &armpolicy.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/createOrUpdatePolicyExemptionWithResourceSelectors.json
func ExampleExemptionsClient_CreateOrUpdate_createOrUpdateAPolicyExemptionWithResourceSelectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().CreateOrUpdate(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", armpolicy.Exemption{
		Properties: &armpolicy.ExemptionProperties{
			Description:               to.Ptr("Exempt demo cluster from limit sku"),
			AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
			DisplayName:               to.Ptr("Exempt demo cluster"),
			ExemptionCategory:         to.Ptr(armpolicy.ExemptionCategoryWaiver),
			Metadata: map[string]any{
				"reason": "Temporary exemption for a expensive VM demo",
			},
			PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
			PolicyDefinitionReferenceIDs: []*string{
				to.Ptr("Limit_Skus")},
			ResourceSelectors: []*armpolicy.ResourceSelector{
				{
					Name: to.Ptr("SDPRegions"),
					Selectors: []*armpolicy.Selector{
						{
							In: []*string{
								to.Ptr("eastus2euap"),
								to.Ptr("centraluseuap")},
							Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Exemption = armpolicy.Exemption{
	// 	Name: to.Ptr("DemoExpensiveVM"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 	Properties: &armpolicy.ExemptionProperties{
	// 		Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 		AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
	// 		DisplayName: to.Ptr("Exempt demo cluster"),
	// 		ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 		Metadata: map[string]any{
	// 			"reason": "Temporary exemption for a expensive VM demo",
	// 		},
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		PolicyDefinitionReferenceIDs: []*string{
	// 			to.Ptr("Limit_Skus")},
	// 			ResourceSelectors: []*armpolicy.ResourceSelector{
	// 				{
	// 					Name: to.Ptr("SDPRegions"),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							In: []*string{
	// 								to.Ptr("eastus2euap"),
	// 								to.Ptr("centraluseuap")},
	// 								Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
	// 						}},
	// 				}},
	// 			},
	// 			SystemData: &armpolicy.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
	// 				CreatedBy: to.Ptr("string"),
	// 				CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("string"),
	// 				LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/getPolicyExemption.json
func ExampleExemptionsClient_Get_retrieveAPolicyExemption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().Get(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Exemption = armpolicy.Exemption{
	// 	Name: to.Ptr("DemoExpensiveVM"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 	Properties: &armpolicy.ExemptionProperties{
	// 		Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 		DisplayName: to.Ptr("Exempt demo cluster"),
	// 		ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 		Metadata: map[string]any{
	// 			"reason": "Temporary exemption for a expensive VM demo",
	// 		},
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		PolicyDefinitionReferenceIDs: []*string{
	// 			to.Ptr("Limit_Skus")},
	// 		},
	// 		SystemData: &armpolicy.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/getPolicyExemptionWithResourceSelectors.json
func ExampleExemptionsClient_Get_retrieveAPolicyExemptionWithResourceSelectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().Get(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Exemption = armpolicy.Exemption{
	// 	Name: to.Ptr("DemoExpensiveVM"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 	Properties: &armpolicy.ExemptionProperties{
	// 		Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 		AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
	// 		DisplayName: to.Ptr("Exempt demo cluster"),
	// 		ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 		Metadata: map[string]any{
	// 			"reason": "Temporary exemption for a expensive VM demo",
	// 		},
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		PolicyDefinitionReferenceIDs: []*string{
	// 			to.Ptr("Limit_Skus")},
	// 			ResourceSelectors: []*armpolicy.ResourceSelector{
	// 				{
	// 					Name: to.Ptr("SDPRegions"),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							In: []*string{
	// 								to.Ptr("eastus2euap"),
	// 								to.Ptr("centraluseuap")},
	// 								Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
	// 						}},
	// 				}},
	// 			},
	// 			SystemData: &armpolicy.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
	// 				CreatedBy: to.Ptr("string"),
	// 				CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("string"),
	// 				LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/updatePolicyExemptionWithResourceSelectors.json
func ExampleExemptionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().Update(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", armpolicy.ExemptionUpdate{
		Properties: &armpolicy.ExemptionUpdateProperties{
			AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
			ResourceSelectors: []*armpolicy.ResourceSelector{
				{
					Name: to.Ptr("SDPRegions"),
					Selectors: []*armpolicy.Selector{
						{
							In: []*string{
								to.Ptr("eastus2euap"),
								to.Ptr("centraluseuap")},
							Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Exemption = armpolicy.Exemption{
	// 	Name: to.Ptr("DemoExpensiveVM"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 	Properties: &armpolicy.ExemptionProperties{
	// 		Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 		AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
	// 		DisplayName: to.Ptr("Exempt demo cluster"),
	// 		ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 		Metadata: map[string]any{
	// 			"reason": "Temporary exemption for a expensive VM demo",
	// 		},
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		PolicyDefinitionReferenceIDs: []*string{
	// 			to.Ptr("Limit_Skus")},
	// 			ResourceSelectors: []*armpolicy.ResourceSelector{
	// 				{
	// 					Name: to.Ptr("SDPRegions"),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							In: []*string{
	// 								to.Ptr("eastus2euap"),
	// 								to.Ptr("centraluseuap")},
	// 								Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
	// 						}},
	// 				}},
	// 			},
	// 			SystemData: &armpolicy.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
	// 				CreatedBy: to.Ptr("string"),
	// 				CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("string"),
	// 				LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/listPolicyExemptionsForSubscription.json
func ExampleExemptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExemptionsClient().NewListPager(&armpolicy.ExemptionsClientListOptions{Filter: to.Ptr("atScope()")})
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
		// page.ExemptionListResult = armpolicy.ExemptionListResult{
		// 	Value: []*armpolicy.Exemption{
		// 		{
		// 			Name: to.Ptr("TestVMSub"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyExemptions/TestVMSub"),
		// 			Properties: &armpolicy.ExemptionProperties{
		// 				Description: to.Ptr("Exempt demo cluster from limit sku"),
		// 				DisplayName: to.Ptr("Exempt demo cluster"),
		// 				ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
		// 				Metadata: map[string]any{
		// 					"reason": "Temporary exemption for a expensive VM demo",
		// 				},
		// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 				PolicyDefinitionReferenceIDs: []*string{
		// 					to.Ptr("Limit_Skus")},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestVNetSub"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyExemptions/TestVNetSub"),
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					Description: to.Ptr("Exempt jump box open ports from limit ports policy"),
		// 					DisplayName: to.Ptr("Exempt jump box open ports"),
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryMitigated),
		// 					Metadata: map[string]any{
		// 						"reason": "Need to open RDP port to corp net",
		// 					},
		// 					PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyAssignments/LimitPorts"),
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/listPolicyExemptionsForResourceGroup.json
func ExampleExemptionsClient_NewListForResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExemptionsClient().NewListForResourceGroupPager("TestResourceGroup", &armpolicy.ExemptionsClientListForResourceGroupOptions{Filter: to.Ptr("atScope()")})
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
		// page.ExemptionListResult = armpolicy.ExemptionListResult{
		// 	Value: []*armpolicy.Exemption{
		// 		{
		// 			Name: to.Ptr("TestVMSub"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyExemptions/TestVMSub"),
		// 			Properties: &armpolicy.ExemptionProperties{
		// 				Description: to.Ptr("Exempt demo cluster from limit sku"),
		// 				DisplayName: to.Ptr("Exempt demo cluster"),
		// 				ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
		// 				Metadata: map[string]any{
		// 					"reason": "Temporary exemption for a expensive VM demo",
		// 				},
		// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 				PolicyDefinitionReferenceIDs: []*string{
		// 					to.Ptr("Limit_Skus")},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestVNetRG"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Authorization/policyExemptions/TestVNetRG"),
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					Description: to.Ptr("Exempt jump box open ports from limit ports policy"),
		// 					DisplayName: to.Ptr("Exempt jump box open ports"),
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryMitigated),
		// 					Metadata: map[string]any{
		// 						"reason": "Need to open RDP port to corp net",
		// 					},
		// 					PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyAssignments/LimitPorts"),
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/listPolicyExemptionsForResource.json
func ExampleExemptionsClient_NewListForResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExemptionsClient().NewListForResourcePager("TestResourceGroup", "Microsoft.Compute", "virtualMachines/MyTestVm", "domainNames", "MyTestComputer.cloudapp.net", &armpolicy.ExemptionsClientListForResourceOptions{Filter: nil})
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
		// page.ExemptionListResult = armpolicy.ExemptionListResult{
		// 	Value: []*armpolicy.Exemption{
		// 		{
		// 			Name: to.Ptr("DemoExpensiveVMGroup"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVMGroup"),
		// 			Properties: &armpolicy.ExemptionProperties{
		// 				Description: to.Ptr("Exempt demo cluster from limit sku"),
		// 				DisplayName: to.Ptr("Exempt demo cluster"),
		// 				ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
		// 				Metadata: map[string]any{
		// 					"reason": "Temporary exemption for a expensive VM demo",
		// 				},
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 				PolicyDefinitionReferenceIDs: []*string{
		// 					to.Ptr("Limit_Skus")},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("jumpBoxExemption"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Compute/virtualMachines/MyTestVm/providers/Microsoft.Authorization/policyExemptions/jumpBoxExemption"),
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					Description: to.Ptr("Exempt jump box open ports from limit ports policy"),
		// 					DisplayName: to.Ptr("Exempt jump box open ports"),
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryMitigated),
		// 					Metadata: map[string]any{
		// 						"reason": "Need to open RDP port to corp net",
		// 					},
		// 					PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/LimitPorts"),
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/listPolicyExemptionsForManagementGroup.json
func ExampleExemptionsClient_NewListForManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExemptionsClient().NewListForManagementGroupPager("DevOrg", &armpolicy.ExemptionsClientListForManagementGroupOptions{Filter: to.Ptr("atScope()")})
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
		// page.ExemptionListResult = armpolicy.ExemptionListResult{
		// 	Value: []*armpolicy.Exemption{
		// 		{
		// 			Name: to.Ptr("ResearchBudgetExemption"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyExemptions/ResearchBudgetExemption"),
		// 			Properties: &armpolicy.ExemptionProperties{
		// 				Description: to.Ptr("Exempt demo cluster from limit sku"),
		// 				DisplayName: to.Ptr("Exempt demo cluster"),
		// 				ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
		// 				Metadata: map[string]any{
		// 					"reason": "Temporary exemption for a expensive VM demo",
		// 				},
		// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/HardwareDivision/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 				PolicyDefinitionReferenceIDs: []*string{
		// 					to.Ptr("Limit_Skus")},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("VNetIsMonitored"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyExemptions/VNetIsMonitored"),
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					Description: to.Ptr("Exempt jump box open ports from limit ports policy"),
		// 					DisplayName: to.Ptr("Exempt jump box open ports"),
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryMitigated),
		// 					Metadata: map[string]any{
		// 						"reason": "Need to open RDP port to corp net",
		// 					},
		// 					PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/HardwareDivision/providers/Microsoft.Authorization/policyAssignments/LimitPorts"),
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}
