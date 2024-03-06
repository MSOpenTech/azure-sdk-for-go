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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinition.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().CreateOrUpdate(ctx, "ResourceNaming", armpolicy.Definition{
		Properties: &armpolicy.DefinitionProperties{
			Description: to.Ptr("Force resource names to begin with given 'prefix' and/or end with given 'suffix'"),
			DisplayName: to.Ptr("Enforce resource naming convention"),
			Metadata: map[string]any{
				"category": "Naming",
			},
			Mode: to.Ptr("All"),
			Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
				"prefix": {
					Type: to.Ptr(armpolicy.ParameterTypeString),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("Resource name prefix"),
						DisplayName: to.Ptr("Prefix"),
					},
				},
				"suffix": {
					Type: to.Ptr(armpolicy.ParameterTypeString),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("Resource name suffix"),
						DisplayName: to.Ptr("Suffix"),
					},
				},
			},
			PolicyRule: map[string]any{
				"if": map[string]any{
					"not": map[string]any{
						"field": "name",
						"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
					},
				},
				"then": map[string]any{
					"effect": "deny",
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAdvancedParams.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinitionWithAdvancedParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().CreateOrUpdate(ctx, "EventHubDiagnosticLogs", armpolicy.Definition{
		Properties: &armpolicy.DefinitionProperties{
			Description: to.Ptr("Audit enabling of logs and retain them up to a year. This enables recreation of activity trails for investigation purposes when a security incident occurs or your network is compromised"),
			DisplayName: to.Ptr("Event Hubs should have diagnostic logging enabled"),
			Metadata: map[string]any{
				"category": "Event Hub",
			},
			Mode: to.Ptr("Indexed"),
			Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
				"requiredRetentionDays": {
					Type: to.Ptr(armpolicy.ParameterTypeInteger),
					AllowedValues: []any{
						float64(0),
						float64(30),
						float64(90),
						float64(180),
						float64(365)},
					DefaultValue: float64(365),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("The required diagnostic logs retention in days"),
						DisplayName: to.Ptr("Required retention (days)"),
					},
				},
			},
			PolicyRule: map[string]any{
				"if": map[string]any{
					"equals": "Microsoft.EventHub/namespaces",
					"field":  "type",
				},
				"then": map[string]any{
					"effect": "AuditIfNotExists",
					"details": map[string]any{
						"type": "Microsoft.Insights/diagnosticSettings",
						"existenceCondition": map[string]any{
							"allOf": []any{
								map[string]any{
									"equals": "true",
									"field":  "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled",
								},
								map[string]any{
									"equals": "[parameters('requiredRetentionDays')]",
									"field":  "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days",
								},
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyDefinition.json
func ExampleDefinitionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().Delete(ctx, "ResourceNaming", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinition.json
func ExampleDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionsClient().Get(ctx, "ResourceNaming", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Definition = armpolicy.Definition{
	// 	Name: to.Ptr("ResourceNaming"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 	Properties: &armpolicy.DefinitionProperties{
	// 		Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
	// 		DisplayName: to.Ptr("Naming Convention"),
	// 		Metadata: map[string]any{
	// 			"category": "Naming",
	// 		},
	// 		Mode: to.Ptr("All"),
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 			"prefix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("Resource name prefix"),
	// 					DisplayName: to.Ptr("Prefix"),
	// 				},
	// 			},
	// 			"suffix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("Resource name suffix"),
	// 					DisplayName: to.Ptr("Suffix"),
	// 				},
	// 			},
	// 		},
	// 		PolicyRule: map[string]any{
	// 			"if":map[string]any{
	// 				"not":map[string]any{
	// 					"field": "name",
	// 					"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
	// 				},
	// 			},
	// 			"then":map[string]any{
	// 				"effect": "deny",
	// 			},
	// 		},
	// 		PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getBuiltinPolicyDefinition.json
func ExampleDefinitionsClient_GetBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionsClient().GetBuiltIn(ctx, "7433c107-6db4-4ad1-b57a-a76dce0154a1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Definition = armpolicy.Definition{
	// 	Name: to.Ptr("7433c107-6db4-4ad1-b57a-a76dce0154a1"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
	// 	Properties: &armpolicy.DefinitionProperties{
	// 		Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
	// 		DisplayName: to.Ptr("Allowed storage account SKUs"),
	// 		Mode: to.Ptr("All"),
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 			"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeArray),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
	// 					DisplayName: to.Ptr("Allowed SKUs"),
	// 					StrongType: to.Ptr("StorageSKUs"),
	// 				},
	// 			},
	// 		},
	// 		PolicyRule: map[string]any{
	// 			"if":map[string]any{
	// 				"allOf":[]any{
	// 					map[string]any{
	// 						"equals": "Microsoft.Storage/storageAccounts",
	// 						"field": "type",
	// 					},
	// 					map[string]any{
	// 						"not":map[string]any{
	// 							"field": "Microsoft.Storage/storageAccounts/sku.name",
	// 							"in": "[parameters('listOfAllowedSKUs')]",
	// 						},
	// 					},
	// 				},
	// 			},
	// 			"then":map[string]any{
	// 				"effect": "Deny",
	// 			},
	// 		},
	// 		PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_CreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().CreateOrUpdateAtManagementGroup(ctx, "ResourceNaming", "MyManagementGroup", armpolicy.Definition{
		Properties: &armpolicy.DefinitionProperties{
			Description: to.Ptr("Force resource names to begin with given 'prefix' and/or end with given 'suffix'"),
			DisplayName: to.Ptr("Enforce resource naming convention"),
			Metadata: map[string]any{
				"category": "Naming",
			},
			Mode: to.Ptr("All"),
			Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
				"prefix": {
					Type: to.Ptr(armpolicy.ParameterTypeString),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("Resource name prefix"),
						DisplayName: to.Ptr("Prefix"),
					},
				},
				"suffix": {
					Type: to.Ptr(armpolicy.ParameterTypeString),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("Resource name suffix"),
						DisplayName: to.Ptr("Suffix"),
					},
				},
			},
			PolicyRule: map[string]any{
				"if": map[string]any{
					"not": map[string]any{
						"field": "name",
						"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
					},
				},
				"then": map[string]any{
					"effect": "deny",
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_DeleteAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().DeleteAtManagementGroup(ctx, "ResourceNaming", "MyManagementGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_GetAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionsClient().GetAtManagementGroup(ctx, "ResourceNaming", "MyManagementGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Definition = armpolicy.Definition{
	// 	Name: to.Ptr("ResourceNaming"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 	Properties: &armpolicy.DefinitionProperties{
	// 		Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
	// 		DisplayName: to.Ptr("Naming Convention"),
	// 		Metadata: map[string]any{
	// 			"category": "Naming",
	// 		},
	// 		Mode: to.Ptr("All"),
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 			"prefix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("Resource name prefix"),
	// 					DisplayName: to.Ptr("Prefix"),
	// 				},
	// 			},
	// 			"suffix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("Resource name suffix"),
	// 					DisplayName: to.Ptr("Suffix"),
	// 				},
	// 			},
	// 		},
	// 		PolicyRule: map[string]any{
	// 			"if":map[string]any{
	// 				"not":map[string]any{
	// 					"field": "name",
	// 					"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
	// 				},
	// 			},
	// 			"then":map[string]any{
	// 				"effect": "deny",
	// 			},
	// 		},
	// 		PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitions.json
func ExampleDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefinitionsClient().NewListPager(&armpolicy.DefinitionsClientListOptions{Filter: nil,
		Top: nil,
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
		// page.DefinitionListResult = armpolicy.DefinitionListResult{
		// 	Value: []*armpolicy.Definition{
		// 		{
		// 			Name: to.Ptr("7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
		// 				DisplayName: to.Ptr("Allowed storage account SKUs"),
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeArray),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
		// 							DisplayName: to.Ptr("Allowed SKUs"),
		// 							StrongType: to.Ptr("StorageSKUs"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"allOf":[]any{
		// 							map[string]any{
		// 								"equals": "Microsoft.Storage/storageAccounts",
		// 								"field": "type",
		// 							},
		// 							map[string]any{
		// 								"not":map[string]any{
		// 									"field": "Microsoft.Storage/storageAccounts/sku.name",
		// 									"in": "[parameters('listOfAllowedSKUs')]",
		// 								},
		// 							},
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "Deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ResourceNaming"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
		// 				DisplayName: to.Ptr("Naming Convention"),
		// 				Metadata: map[string]any{
		// 					"category": "Naming",
		// 				},
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"prefix": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("Resource name prefix"),
		// 							DisplayName: to.Ptr("Prefix"),
		// 						},
		// 					},
		// 					"suffix": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("Resource name suffix"),
		// 							DisplayName: to.Ptr("Suffix"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"not":map[string]any{
		// 							"field": "name",
		// 							"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AuditSoonToExpireCerts"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/AuditSoonToExpireCerts"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("Audit certificates that are stored in Azure Key Vault, that expire within 'X' number of days."),
		// 				DisplayName: to.Ptr("Audit KeyVault certificates that expire within specified number of days"),
		// 				Metadata: map[string]any{
		// 					"category": "KeyVault DataPlane",
		// 				},
		// 				Mode: to.Ptr("Microsoft.KeyVault.Data"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"daysToExpire": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeInteger),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("The number of days for a certificate to expire."),
		// 							DisplayName: to.Ptr("Days to expire"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"field": "Microsoft.KeyVault.Data/vaults/certificates/attributes/expiresOn",
		// 						"lessOrEquals": "[addDays(utcNow(), parameters('daysToExpire'))]",
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "audit",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listBuiltInPolicyDefinitions.json
func ExampleDefinitionsClient_NewListBuiltInPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefinitionsClient().NewListBuiltInPager(&armpolicy.DefinitionsClientListBuiltInOptions{Filter: nil,
		Top: nil,
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
		// page.DefinitionListResult = armpolicy.DefinitionListResult{
		// 	Value: []*armpolicy.Definition{
		// 		{
		// 			Name: to.Ptr("06a78e20-9358-41c9-923c-fb736d382a12"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/06a78e20-9358-41c9-923c-fb736d382a12"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("Audit DB level audit setting for SQL databases"),
		// 				DisplayName: to.Ptr("Audit SQL DB Level Audit Setting"),
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"setting": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						AllowedValues: []any{
		// 							"enabled",
		// 							"disabled"},
		// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 								DisplayName: to.Ptr("Audit Setting"),
		// 							},
		// 						},
		// 					},
		// 					PolicyRule: map[string]any{
		// 						"if":map[string]any{
		// 							"equals": "Microsoft.Sql/servers/databases",
		// 							"field": "type",
		// 						},
		// 						"then":map[string]any{
		// 							"effect": "AuditIfNotExists",
		// 							"details":map[string]any{
		// 								"name": "default",
		// 								"type": "Microsoft.Sql/servers/databases/auditingSettings",
		// 								"existenceCondition":map[string]any{
		// 									"allOf":[]any{
		// 										map[string]any{
		// 											"equals": "[parameters('setting')]",
		// 											"field": "Microsoft.Sql/auditingSettings.state",
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 				Properties: &armpolicy.DefinitionProperties{
		// 					Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
		// 					DisplayName: to.Ptr("Allowed storage account SKUs"),
		// 					Mode: to.Ptr("All"),
		// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 						"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
		// 							Type: to.Ptr(armpolicy.ParameterTypeArray),
		// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 								Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
		// 								DisplayName: to.Ptr("Allowed SKUs"),
		// 								StrongType: to.Ptr("StorageSKUs"),
		// 							},
		// 						},
		// 					},
		// 					PolicyRule: map[string]any{
		// 						"if":map[string]any{
		// 							"allOf":[]any{
		// 								map[string]any{
		// 									"equals": "Microsoft.Storage/storageAccounts",
		// 									"field": "type",
		// 								},
		// 								map[string]any{
		// 									"not":map[string]any{
		// 										"field": "Microsoft.Storage/storageAccounts/sku.name",
		// 										"in": "[parameters('listOfAllowedSKUs')]",
		// 									},
		// 								},
		// 							},
		// 						},
		// 						"then":map[string]any{
		// 							"effect": "Deny",
		// 						},
		// 					},
		// 					PolicyType: to.Ptr(armpolicy.PolicyTypeStatic),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("abeed54a-73c5-441d-8a8c-6b5e7a0c299e"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/abeed54a-73c5-441d-8a8c-6b5e7a0c299e"),
		// 				Properties: &armpolicy.DefinitionProperties{
		// 					Description: to.Ptr("Audit certificates that are stored in Azure Key Vault, that expire within 'X' number of days."),
		// 					DisplayName: to.Ptr("Audit KeyVault certificates that expire within specified number of days"),
		// 					Metadata: map[string]any{
		// 						"category": "KeyVault DataPlane",
		// 					},
		// 					Mode: to.Ptr("Microsoft.KeyVault.Data"),
		// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 						"daysToExpire": &armpolicy.ParameterDefinitionsValue{
		// 							Type: to.Ptr(armpolicy.ParameterTypeInteger),
		// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 								Description: to.Ptr("The number of days for a certificate to expire."),
		// 								DisplayName: to.Ptr("Days to expire"),
		// 							},
		// 						},
		// 					},
		// 					PolicyRule: map[string]any{
		// 						"if":map[string]any{
		// 							"field": "Microsoft.KeyVault.Data/vaults/certificates/attributes/expiresOn",
		// 							"lessOrEquals": "[addDays(utcNow(), parameters('daysToExpire'))]",
		// 						},
		// 						"then":map[string]any{
		// 							"effect": "audit",
		// 						},
		// 					},
		// 					PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitionsByManagementGroup.json
func ExampleDefinitionsClient_NewListByManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefinitionsClient().NewListByManagementGroupPager("MyManagementGroup", &armpolicy.DefinitionsClientListByManagementGroupOptions{Filter: nil,
		Top: nil,
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
		// page.DefinitionListResult = armpolicy.DefinitionListResult{
		// 	Value: []*armpolicy.Definition{
		// 		{
		// 			Name: to.Ptr("7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
		// 				DisplayName: to.Ptr("Allowed storage account SKUs"),
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeArray),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
		// 							DisplayName: to.Ptr("Allowed SKUs"),
		// 							StrongType: to.Ptr("StorageSKUs"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"allOf":[]any{
		// 							map[string]any{
		// 								"equals": "Microsoft.Storage/storageAccounts",
		// 								"field": "type",
		// 							},
		// 							map[string]any{
		// 								"not":map[string]any{
		// 									"field": "Microsoft.Storage/storageAccounts/sku.name",
		// 									"in": "[parameters('listOfAllowedSKUs')]",
		// 								},
		// 							},
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "Deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ResourceNaming"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
		// 				DisplayName: to.Ptr("Naming Convention"),
		// 				Metadata: map[string]any{
		// 					"category": "Naming",
		// 				},
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"prefix": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("Resource name prefix"),
		// 							DisplayName: to.Ptr("Prefix"),
		// 						},
		// 					},
		// 					"suffix": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("Resource name suffix"),
		// 							DisplayName: to.Ptr("Suffix"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"not":map[string]any{
		// 							"field": "name",
		// 							"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
		// 			},
		// 	}},
		// }
	}
}
