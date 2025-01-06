//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CreateAccount.json
func ExampleAccountsClient_BeginCreate_createAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "myResourceGroup", "testCreate1", armcognitiveservices.Account{
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr("Emotion"),
		Location: to.Ptr("West US"),
		Properties: &armcognitiveservices.AccountProperties{
			Encryption: &armcognitiveservices.Encryption{
				KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
				KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
					KeyName:     to.Ptr("KeyName"),
					KeyVaultURI: to.Ptr("https://pltfrmscrts-use-pc-dev.vault.azure.net/"),
					KeyVersion:  to.Ptr("891CF236-D241-4738-9462-D506AF493DFA"),
				},
			},
			UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
				{
					ResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
				}},
		},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
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
	// res.Account = armcognitiveservices.Account{
	// 	Name: to.Ptr("testCreate1"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/testCreate1"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T08%3A00%3A05.445595Z'\""),
	// 	Identity: &armcognitiveservices.Identity{
	// 		Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("b5cf119e-a5c2-42c7-802f-592e0efb169f"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Kind: to.Ptr("Emotion"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.AccountProperties{
	// 		Encryption: &armcognitiveservices.Encryption{
	// 			KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
	// 			KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
	// 				KeyName: to.Ptr("FakeKeyName"),
	// 				KeyVaultURI: to.Ptr("https://pltfrmscrts-use-pc-dev.vault.azure.net/"),
	// 				KeyVersion: to.Ptr("891CF236-D241-4738-9462-D506AF493DFA"),
	// 			},
	// 		},
	// 		Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 		UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
	// 		}},
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CreateAccountMin.json
func ExampleAccountsClient_BeginCreate_createAccountMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "myResourceGroup", "testCreate1", armcognitiveservices.Account{
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
		},
		Kind:       to.Ptr("CognitiveServices"),
		Location:   to.Ptr("West US"),
		Properties: &armcognitiveservices.AccountProperties{},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
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
	// res.Account = armcognitiveservices.Account{
	// 	Name: to.Ptr("testCreate1"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/testCreate1"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T08%3A00%3A05.445595Z'\""),
	// 	Identity: &armcognitiveservices.Identity{
	// 		Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("b5cf119e-a5c2-42c7-802f-592e0efb169f"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Kind: to.Ptr("Emotion"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.AccountProperties{
	// 		Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/UpdateAccount.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginUpdate(ctx, "bvttest", "bingSearch", armcognitiveservices.Account{
		Location: to.Ptr("global"),
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S2"),
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
	// res.Account = armcognitiveservices.Account{
	// 	Name: to.Ptr("bingSearch"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/bingSearch"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T07%3A46%3A21.5618831Z'\""),
	// 	Kind: to.Ptr("Bing.Search"),
	// 	Location: to.Ptr("global"),
	// 	Properties: &armcognitiveservices.AccountProperties{
	// 		Endpoint: to.Ptr("https://api.cognitive.microsoft.com/bing/v5.0"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/DeleteAccount.json
func ExampleAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginDelete(ctx, "myResourceGroup", "PropTest01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/GetAccount.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "myResourceGroup", "myAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armcognitiveservices.Account{
	// 	Name: to.Ptr("myAccount"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T04%3A42%3A19.7067387Z'\""),
	// 	Kind: to.Ptr("Emotion"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcognitiveservices.AccountProperties{
	// 		Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("F0"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"ExpiredDate": to.Ptr("2017/09/01"),
	// 		"Owner": to.Ptr("felixwa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListAccountsByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.AccountListResult = armcognitiveservices.AccountListResult{
		// 	Value: []*armcognitiveservices.Account{
		// 		{
		// 			Name: to.Ptr("myAccount"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-10T04%3A42%3A19.7067387Z'\""),
		// 			Kind: to.Ptr("Emotion"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("F0"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"ExpiredDate": to.Ptr("2017/09/01"),
		// 				"Owner": to.Ptr("felixwa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestPropertyWU2"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/TestPropertyWU2"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-07T04%3A32%3A38.9187216Z'\""),
		// 			Kind: to.Ptr("Face"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/face/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListAccountsBySubscription.json
func ExampleAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(nil)
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
		// page.AccountListResult = armcognitiveservices.AccountListResult{
		// 	Value: []*armcognitiveservices.Account{
		// 		{
		// 			Name: to.Ptr("bingSearch"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/bingSearch"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-27T11%3A19%3A08.762494Z'\""),
		// 			Kind: to.Ptr("Bing.Search"),
		// 			Location: to.Ptr("global"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://api.cognitive.microsoft.com/bing/v5.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CrisProd"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/CrisProd"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-31T08%3A57%3A07.4499566Z'\""),
		// 			Kind: to.Ptr("CRIS"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/sts/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"can't delete it successfully": to.Ptr("v-yunjin"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rayrptest0308"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/rayrptest0308"),
		// 			Etag: to.Ptr("W/\"datetime'2017-03-27T11%3A15%3A23.5232645Z'\""),
		// 			Kind: to.Ptr("Face"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/face/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("raytest02"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/raytest02"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-04T02%3A07%3A07.3957572Z'\""),
		// 			Kind: to.Ptr("Emotion"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.AccountProperties{
		// 				Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListKeys.json
func ExampleAccountsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListKeys(ctx, "myResourceGroup", "myAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeys = armcognitiveservices.APIKeys{
	// 	Key1: to.Ptr("KEY1"),
	// 	Key2: to.Ptr("KEY2"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/RegenerateKey.json
func ExampleAccountsClient_RegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().RegenerateKey(ctx, "myResourceGroup", "myAccount", armcognitiveservices.RegenerateKeyParameters{
		KeyName: to.Ptr(armcognitiveservices.KeyNameKey2),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeys = armcognitiveservices.APIKeys{
	// 	Key1: to.Ptr("KEY1"),
	// 	Key2: to.Ptr("KEY2"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListSkus.json
func ExampleAccountsClient_ListSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListSKUs(ctx, "myResourceGroup", "myAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountSKUListResult = armcognitiveservices.AccountSKUListResult{
	// 	Value: []*armcognitiveservices.AccountSKU{
	// 		{
	// 			ResourceType: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 			SKU: &armcognitiveservices.SKU{
	// 				Name: to.Ptr("F0"),
	// 				Tier: to.Ptr(armcognitiveservices.SKUTierFree),
	// 			},
	// 		},
	// 		{
	// 			ResourceType: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 			SKU: &armcognitiveservices.SKU{
	// 				Name: to.Ptr("S0"),
	// 				Tier: to.Ptr(armcognitiveservices.SKUTierStandard),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/GetUsages.json
func ExampleAccountsClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListUsages(ctx, "myResourceGroup", "TestUsage02", &armcognitiveservices.AccountsClientListUsagesOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UsageListResult = armcognitiveservices.UsageListResult{
	// 	Value: []*armcognitiveservices.Usage{
	// 		{
	// 			Name: &armcognitiveservices.MetricName{
	// 				LocalizedValue: to.Ptr("Face.Transactions"),
	// 				Value: to.Ptr("Face.Transactions"),
	// 			},
	// 			CurrentValue: to.Ptr[float64](3),
	// 			Limit: to.Ptr[float64](30000),
	// 			NextResetTime: to.Ptr("2018-03-28T09:33:51Z"),
	// 			QuotaPeriod: to.Ptr("30.00:00:00"),
	// 			Status: to.Ptr(armcognitiveservices.QuotaUsageStatusIncluded),
	// 			Unit: to.Ptr(armcognitiveservices.UnitTypeCount),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListAccountModels.json
func ExampleAccountsClient_NewListModelsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListModelsPager("resourceGroupName", "accountName", nil)
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
		// page.AccountModelListResult = armcognitiveservices.AccountModelListResult{
		// 	Value: []*armcognitiveservices.AccountModel{
		// 		{
		// 			Name: to.Ptr("ada.1"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("1"),
		// 			BaseModel: &armcognitiveservices.DeploymentModel{
		// 				Name: to.Ptr("ada"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Version: to.Ptr("1"),
		// 			},
		// 			Capabilities: map[string]*string{
		// 				"fineTune": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				FineTune: to.Ptr("2024-01-01T00:00:00Z"),
		// 				Inference: to.Ptr("2024-01-01T00:00:00Z"),
		// 			},
		// 			MaxCapacity: to.Ptr[int32](10),
		// 		},
		// 		{
		// 			Name: to.Ptr("davinci"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("1"),
		// 			Capabilities: map[string]*string{
		// 				"fineTune": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				FineTune: to.Ptr("2024-01-01T00:00:00Z"),
		// 				Inference: to.Ptr("2024-01-01T00:00:00Z"),
		// 			},
		// 			MaxCapacity: to.Ptr[int32](10),
		// 	}},
		// }
	}
}
