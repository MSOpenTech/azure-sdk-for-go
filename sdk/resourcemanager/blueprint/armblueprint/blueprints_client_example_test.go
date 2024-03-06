//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_Create.json
func ExampleBlueprintsClient_CreateOrUpdate_managementGroupBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBlueprintsClient().CreateOrUpdate(ctx, "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", "simpleBlueprint", armblueprint.Blueprint{
		Properties: &armblueprint.Properties{
			Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
			Parameters: map[string]*armblueprint.ParameterDefinition{
				"costCenter": {
					Type: to.Ptr(armblueprint.TemplateParameterTypeString),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
					},
				},
				"owners": {
					Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
					},
				},
				"storageAccountType": {
					Type: to.Ptr(armblueprint.TemplateParameterTypeString),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("storage account type."),
					},
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
				"storageRG": {
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
						DisplayName: to.Ptr("storage resource group"),
					},
				},
			},
			TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/ResourceGroupWithTags.json
func ExampleBlueprintsClient_CreateOrUpdate_resourceGroupWithTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBlueprintsClient().CreateOrUpdate(ctx, "providers/Microsoft.Management/managementGroups/{ManagementGroupId}", "simpleBlueprint", armblueprint.Blueprint{
		Properties: &armblueprint.Properties{
			Description: to.Ptr("An example blueprint containing an RG with two tags."),
			ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
				"myRGName": {
					Name:     to.Ptr("myRGName"),
					Location: to.Ptr("westus"),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("My Resource Group"),
					},
					Tags: map[string]*string{
						"costcenter":  to.Ptr("123456"),
						"nameOnlyTag": to.Ptr(""),
					},
				},
			},
			TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/Blueprint_Create.json
func ExampleBlueprintsClient_CreateOrUpdate_subscriptionBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBlueprintsClient().CreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "simpleBlueprint", armblueprint.Blueprint{
		Properties: &armblueprint.Properties{
			Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
			Parameters: map[string]*armblueprint.ParameterDefinition{
				"costCenter": {
					Type: to.Ptr(armblueprint.TemplateParameterTypeString),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
					},
				},
				"owners": {
					Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
					},
				},
				"storageAccountType": {
					Type: to.Ptr(armblueprint.TemplateParameterTypeString),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("storage account type."),
					},
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
				"storageRG": {
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
						DisplayName: to.Ptr("storage resource group"),
					},
				},
			},
			TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_Get.json
func ExampleBlueprintsClient_Get_managementGroupBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlueprintsClient().Get(ctx, "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", "simpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Blueprint = armblueprint.Blueprint{
	// 	Name: to.Ptr("simpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/blueprints"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 	Properties: &armblueprint.Properties{
	// 		Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
	// 		Parameters: map[string]*armblueprint.ParameterDefinition{
	// 			"costCenter": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
	// 				},
	// 			},
	// 			"owners": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("storage account type."),
	// 				},
	// 			},
	// 		},
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
	// 			"storageRG": &armblueprint.ResourceGroupDefinition{
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
	// 					DisplayName: to.Ptr("storage resource group"),
	// 				},
	// 			},
	// 		},
	// 		TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/Blueprint_Get.json
func ExampleBlueprintsClient_Get_subscriptionBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlueprintsClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "simpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Blueprint = armblueprint.Blueprint{
	// 	Name: to.Ptr("simpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/blueprints"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 	Properties: &armblueprint.Properties{
	// 		Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
	// 		Parameters: map[string]*armblueprint.ParameterDefinition{
	// 			"costCenter": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
	// 				},
	// 			},
	// 			"owners": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("storage account type."),
	// 				},
	// 			},
	// 		},
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
	// 			"storageRG": &armblueprint.ResourceGroupDefinition{
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
	// 					DisplayName: to.Ptr("storage resource group"),
	// 				},
	// 			},
	// 		},
	// 		TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_Delete.json
func ExampleBlueprintsClient_Delete_managementGroupBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlueprintsClient().Delete(ctx, "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", "simpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Blueprint = armblueprint.Blueprint{
	// 	Name: to.Ptr("simpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/blueprints"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 	Properties: &armblueprint.Properties{
	// 		Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
	// 		Parameters: map[string]*armblueprint.ParameterDefinition{
	// 			"costCenter": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
	// 				},
	// 			},
	// 			"owners": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("storage account type."),
	// 				},
	// 			},
	// 		},
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
	// 			"storageRG": &armblueprint.ResourceGroupDefinition{
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
	// 					DisplayName: to.Ptr("storage resource group"),
	// 				},
	// 			},
	// 		},
	// 		TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/Blueprint_Delete.json
func ExampleBlueprintsClient_Delete_subscriptionBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlueprintsClient().Delete(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "simpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Blueprint = armblueprint.Blueprint{
	// 	Name: to.Ptr("simpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/blueprints"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 	Properties: &armblueprint.Properties{
	// 		Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
	// 		Parameters: map[string]*armblueprint.ParameterDefinition{
	// 			"costCenter": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
	// 				},
	// 			},
	// 			"owners": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("storage account type."),
	// 				},
	// 			},
	// 		},
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
	// 			"storageRG": &armblueprint.ResourceGroupDefinition{
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
	// 					DisplayName: to.Ptr("storage resource group"),
	// 				},
	// 			},
	// 		},
	// 		TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_List.json
func ExampleBlueprintsClient_NewListPager_managementGroupBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBlueprintsClient().NewListPager("providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", nil)
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
		// page.List = armblueprint.List{
		// 	Value: []*armblueprint.Blueprint{
		// 		{
		// 			Name: to.Ptr("simpleBlueprint"),
		// 			Type: to.Ptr("Microsoft.Blueprint/blueprints"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
		// 			Properties: &armblueprint.Properties{
		// 				Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
		// 				Parameters: map[string]*armblueprint.ParameterDefinition{
		// 					"costCenter": &armblueprint.ParameterDefinition{
		// 						Type: to.Ptr(armblueprint.TemplateParameterTypeString),
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
		// 						},
		// 					},
		// 					"owners": &armblueprint.ParameterDefinition{
		// 						Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
		// 						},
		// 					},
		// 					"storageAccountType": &armblueprint.ParameterDefinition{
		// 						Type: to.Ptr(armblueprint.TemplateParameterTypeString),
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							DisplayName: to.Ptr("storage account type."),
		// 						},
		// 					},
		// 				},
		// 				ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
		// 					"storageRG": &armblueprint.ResourceGroupDefinition{
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
		// 							DisplayName: to.Ptr("storage resource group"),
		// 						},
		// 					},
		// 				},
		// 				TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/Blueprint_List.json
func ExampleBlueprintsClient_NewListPager_subscriptionBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBlueprintsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", nil)
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
		// page.List = armblueprint.List{
		// 	Value: []*armblueprint.Blueprint{
		// 		{
		// 			Name: to.Ptr("simpleBlueprint"),
		// 			Type: to.Ptr("Microsoft.Blueprint/blueprints"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
		// 			Properties: &armblueprint.Properties{
		// 				Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
		// 				Parameters: map[string]*armblueprint.ParameterDefinition{
		// 					"costCenter": &armblueprint.ParameterDefinition{
		// 						Type: to.Ptr(armblueprint.TemplateParameterTypeString),
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
		// 						},
		// 					},
		// 					"owners": &armblueprint.ParameterDefinition{
		// 						Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
		// 						},
		// 					},
		// 					"storageAccountType": &armblueprint.ParameterDefinition{
		// 						Type: to.Ptr(armblueprint.TemplateParameterTypeString),
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							DisplayName: to.Ptr("storage account type."),
		// 						},
		// 					},
		// 				},
		// 				ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
		// 					"storageRG": &armblueprint.ResourceGroupDefinition{
		// 						Metadata: &armblueprint.ParameterDefinitionMetadata{
		// 							Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
		// 							DisplayName: to.Ptr("storage resource group"),
		// 						},
		// 					},
		// 				},
		// 				TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
		// 			},
		// 	}},
		// }
	}
}
