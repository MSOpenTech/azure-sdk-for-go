// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armtrustedsigning_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
	"log"
)

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_CheckNameAvailability.json
func ExampleCodeSigningAccountsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCodeSigningAccountsClient().CheckNameAvailability(ctx, armtrustedsigning.CheckNameAvailability{
		Name: to.Ptr("sample-account"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armtrustedsigning.CodeSigningAccountsClientCheckNameAvailabilityResponse{
	// 	CheckNameAvailabilityResult: &armtrustedsigning.CheckNameAvailabilityResult{
	// 		NameAvailable: to.Ptr(true),
	// 	},
	// }
}

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_Create.json
func ExampleCodeSigningAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCodeSigningAccountsClient().BeginCreate(ctx, "MyResourceGroup", "MyAccount", armtrustedsigning.CodeSigningAccount{
		Location: to.Ptr("westus"),
		Properties: &armtrustedsigning.CodeSigningAccountProperties{
			SKU: &armtrustedsigning.AccountSKU{
				Name: to.Ptr(armtrustedsigning.SKUNameBasic),
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
	// res = armtrustedsigning.CodeSigningAccountsClientCreateResponse{
	// 	CodeSigningAccount: &armtrustedsigning.CodeSigningAccount{
	// 		Name: to.Ptr("MyAccount"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
	// 		ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armtrustedsigning.CodeSigningAccountProperties{
	// 			ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
	// 			SKU: &armtrustedsigning.AccountSKU{
	// 				Name: to.Ptr(armtrustedsigning.SKUNameBasic),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_Delete.json
func ExampleCodeSigningAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCodeSigningAccountsClient().BeginDelete(ctx, "MyResourceGroup", "MyAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_Get.json
func ExampleCodeSigningAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCodeSigningAccountsClient().Get(ctx, "MyResourceGroup", "MyAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armtrustedsigning.CodeSigningAccountsClientGetResponse{
	// 	CodeSigningAccount: &armtrustedsigning.CodeSigningAccount{
	// 		Name: to.Ptr("MyAccount"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armtrustedsigning.CodeSigningAccountProperties{
	// 			ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
	// 			SKU: &armtrustedsigning.AccountSKU{
	// 				Name: to.Ptr(armtrustedsigning.SKUNameBasic),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_ListByResourceGroup.json
func ExampleCodeSigningAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCodeSigningAccountsClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page = armtrustedsigning.CodeSigningAccountsClientListByResourceGroupResponse{
		// 	CodeSigningAccountListResult: armtrustedsigning.CodeSigningAccountListResult{
		// 		Value: []*armtrustedsigning.CodeSigningAccount{
		// 			{
		// 				Name: to.Ptr("alpha"),
		// 				Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Properties: &armtrustedsigning.CodeSigningAccountProperties{
		// 					ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
		// 					SKU: &armtrustedsigning.AccountSKU{
		// 						Name: to.Ptr(armtrustedsigning.SKUNameBasic),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_ListBySubscription.json
func ExampleCodeSigningAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCodeSigningAccountsClient().NewListBySubscriptionPager(nil)
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
		// page = armtrustedsigning.CodeSigningAccountsClientListBySubscriptionResponse{
		// 	CodeSigningAccountListResult: armtrustedsigning.CodeSigningAccountListResult{
		// 		Value: []*armtrustedsigning.CodeSigningAccount{
		// 			{
		// 				Name: to.Ptr("alpha"),
		// 				Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armtrustedsigning.CodeSigningAccountProperties{
		// 					ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
		// 					SKU: &armtrustedsigning.AccountSKU{
		// 						Name: to.Ptr(armtrustedsigning.SKUNameBasic),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning/TempTypeSpecFiles/CodeSigning.Management/examples/2024-02-05-preview/CodeSigningAccounts_Update.json
func ExampleCodeSigningAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCodeSigningAccountsClient().BeginUpdate(ctx, "MyResourceGroup", "MyAccount", armtrustedsigning.CodeSigningAccountPatch{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
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
	// res = armtrustedsigning.CodeSigningAccountsClientUpdateResponse{
	// 	CodeSigningAccount: &armtrustedsigning.CodeSigningAccount{
	// 		Name: to.Ptr("MyAccount"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armtrustedsigning.CodeSigningAccountProperties{
	// 			ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
	// 			SKU: &armtrustedsigning.AccountSKU{
	// 				Name: to.Ptr(armtrustedsigning.SKUNameBasic),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
