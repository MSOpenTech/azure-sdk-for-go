//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ConnectedEnvironmentsDaprComponents_List.json
func ExampleConnectedEnvironmentsDaprComponentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedEnvironmentsDaprComponentsClient().NewListPager("examplerg", "myenvironment", nil)
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
		// page.DaprComponentsCollection = armappcontainers.DaprComponentsCollection{
		// 	Value: []*armappcontainers.DaprComponent{
		// 		{
		// 			Name: to.Ptr("reddog"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments/daprcomponents"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/myenvironment/daprcomponents/reddog"),
		// 			Properties: &armappcontainers.DaprComponentProperties{
		// 				ComponentType: to.Ptr("state.azure.cosmosdb"),
		// 				IgnoreErrors: to.Ptr(false),
		// 				InitTimeout: to.Ptr("50s"),
		// 				Metadata: []*armappcontainers.DaprMetadata{
		// 					{
		// 						Name: to.Ptr("url"),
		// 						Value: to.Ptr("<COSMOS-URL>"),
		// 					},
		// 					{
		// 						Name: to.Ptr("database"),
		// 						Value: to.Ptr("itemsDB"),
		// 					},
		// 					{
		// 						Name: to.Ptr("collection"),
		// 						Value: to.Ptr("items"),
		// 					},
		// 					{
		// 						Name: to.Ptr("masterkey"),
		// 						SecretRef: to.Ptr("masterkey"),
		// 				}},
		// 				Scopes: []*string{
		// 					to.Ptr("container-app-1"),
		// 					to.Ptr("container-app-2")},
		// 					Secrets: []*armappcontainers.Secret{
		// 						{
		// 							Name: to.Ptr("masterkey"),
		// 					}},
		// 					ServiceComponentBind: []*armappcontainers.DaprComponentServiceBinding{
		// 						{
		// 							Name: to.Ptr("statestore"),
		// 							Metadata: &armappcontainers.DaprServiceBindMetadata{
		// 								Name: to.Ptr("daprcomponentBind"),
		// 								Value: to.Ptr("redis-bind"),
		// 							},
		// 							ServiceID: to.Ptr("/subscriptions/9f7371f1-b593-4c3c-84e2-9167806ad358/resourceGroups/ca-syn2-group/providers/Microsoft.App/containerapps/cappredis"),
		// 					}},
		// 					Version: to.Ptr("v1"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ConnectedEnvironmentsDaprComponents_Get.json
func ExampleConnectedEnvironmentsDaprComponentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsDaprComponentsClient().Get(ctx, "examplerg", "myenvironment", "reddog", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DaprComponent = armappcontainers.DaprComponent{
	// 	Name: to.Ptr("reddog"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments/daprcomponents"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/jlaw-demo1/daprcomponents/reddog"),
	// 	Properties: &armappcontainers.DaprComponentProperties{
	// 		ComponentType: to.Ptr("state.azure.cosmosdb"),
	// 		IgnoreErrors: to.Ptr(false),
	// 		InitTimeout: to.Ptr("50s"),
	// 		Metadata: []*armappcontainers.DaprMetadata{
	// 			{
	// 				Name: to.Ptr("url"),
	// 				Value: to.Ptr("<COSMOS-URL>"),
	// 			},
	// 			{
	// 				Name: to.Ptr("database"),
	// 				Value: to.Ptr("itemsDB"),
	// 			},
	// 			{
	// 				Name: to.Ptr("collection"),
	// 				Value: to.Ptr("items"),
	// 			},
	// 			{
	// 				Name: to.Ptr("masterkey"),
	// 				SecretRef: to.Ptr("masterkey"),
	// 		}},
	// 		Scopes: []*string{
	// 			to.Ptr("container-app-1"),
	// 			to.Ptr("container-app-2")},
	// 			Secrets: []*armappcontainers.Secret{
	// 				{
	// 					Name: to.Ptr("masterkey"),
	// 			}},
	// 			ServiceComponentBind: []*armappcontainers.DaprComponentServiceBinding{
	// 				{
	// 					Name: to.Ptr("statestore"),
	// 					Metadata: &armappcontainers.DaprServiceBindMetadata{
	// 						Name: to.Ptr("daprcomponentBind"),
	// 						Value: to.Ptr("redis-bind"),
	// 					},
	// 					ServiceID: to.Ptr("/subscriptions/9f7371f1-b593-4c3c-84e2-9167806ad358/resourceGroups/ca-syn2-group/providers/Microsoft.App/containerapps/cappredis"),
	// 			}},
	// 			Version: to.Ptr("v1"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ConnectedEnvironmentsDaprComponents_CreateOrUpdate.json
func ExampleConnectedEnvironmentsDaprComponentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsDaprComponentsClient().CreateOrUpdate(ctx, "examplerg", "myenvironment", "reddog", armappcontainers.DaprComponent{
		Properties: &armappcontainers.DaprComponentProperties{
			ComponentType: to.Ptr("state.azure.cosmosdb"),
			IgnoreErrors:  to.Ptr(false),
			InitTimeout:   to.Ptr("50s"),
			Metadata: []*armappcontainers.DaprMetadata{
				{
					Name:  to.Ptr("url"),
					Value: to.Ptr("<COSMOS-URL>"),
				},
				{
					Name:  to.Ptr("database"),
					Value: to.Ptr("itemsDB"),
				},
				{
					Name:  to.Ptr("collection"),
					Value: to.Ptr("items"),
				},
				{
					Name:      to.Ptr("masterkey"),
					SecretRef: to.Ptr("masterkey"),
				}},
			Scopes: []*string{
				to.Ptr("container-app-1"),
				to.Ptr("container-app-2")},
			Secrets: []*armappcontainers.Secret{
				{
					Name:  to.Ptr("masterkey"),
					Value: to.Ptr("keyvalue"),
				}},
			ServiceComponentBind: []*armappcontainers.DaprComponentServiceBinding{
				{
					Name: to.Ptr("statestore"),
					Metadata: &armappcontainers.DaprServiceBindMetadata{
						Name:  to.Ptr("daprcomponentBind"),
						Value: to.Ptr("redis-bind"),
					},
					ServiceID: to.Ptr("/subscriptions/9f7371f1-b593-4c3c-84e2-9167806ad358/resourceGroups/ca-syn2-group/providers/Microsoft.App/containerapps/cappredis"),
				}},
			Version: to.Ptr("v1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DaprComponent = armappcontainers.DaprComponent{
	// 	Name: to.Ptr("reddog"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments/daprcomponents"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/jlaw-demo1/daprcomponents/reddog"),
	// 	Properties: &armappcontainers.DaprComponentProperties{
	// 		ComponentType: to.Ptr("state.azure.cosmosdb"),
	// 		IgnoreErrors: to.Ptr(false),
	// 		InitTimeout: to.Ptr("50s"),
	// 		Metadata: []*armappcontainers.DaprMetadata{
	// 			{
	// 				Name: to.Ptr("url"),
	// 				Value: to.Ptr("<COSMOS-URL>"),
	// 			},
	// 			{
	// 				Name: to.Ptr("database"),
	// 				Value: to.Ptr("itemsDB"),
	// 			},
	// 			{
	// 				Name: to.Ptr("collection"),
	// 				Value: to.Ptr("items"),
	// 			},
	// 			{
	// 				Name: to.Ptr("masterkey"),
	// 				SecretRef: to.Ptr("masterkey"),
	// 		}},
	// 		Scopes: []*string{
	// 			to.Ptr("container-app-1"),
	// 			to.Ptr("container-app-2")},
	// 			Secrets: []*armappcontainers.Secret{
	// 				{
	// 					Name: to.Ptr("masterkey"),
	// 			}},
	// 			ServiceComponentBind: []*armappcontainers.DaprComponentServiceBinding{
	// 				{
	// 					Name: to.Ptr("statestore"),
	// 					Metadata: &armappcontainers.DaprServiceBindMetadata{
	// 						Name: to.Ptr("daprcomponentBind"),
	// 						Value: to.Ptr("redis-bind"),
	// 					},
	// 					ServiceID: to.Ptr("/subscriptions/9f7371f1-b593-4c3c-84e2-9167806ad358/resourceGroups/ca-syn2-group/providers/Microsoft.App/containerapps/cappredis"),
	// 			}},
	// 			Version: to.Ptr("v1"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ConnectedEnvironmentsDaprComponents_Delete.json
func ExampleConnectedEnvironmentsDaprComponentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectedEnvironmentsDaprComponentsClient().Delete(ctx, "examplerg", "myenvironment", "reddog", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ConnectedEnvironmentsDaprComponents_ListSecrets.json
func ExampleConnectedEnvironmentsDaprComponentsClient_ListSecrets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsDaprComponentsClient().ListSecrets(ctx, "examplerg", "myenvironment", "reddog", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DaprSecretsCollection = armappcontainers.DaprSecretsCollection{
	// 	Value: []*armappcontainers.DaprSecret{
	// 		{
	// 			Name: to.Ptr("secret1"),
	// 			Value: to.Ptr("value1"),
	// 		},
	// 		{
	// 			Name: to.Ptr("secret2"),
	// 			Value: to.Ptr("value2"),
	// 	}},
	// }
}
