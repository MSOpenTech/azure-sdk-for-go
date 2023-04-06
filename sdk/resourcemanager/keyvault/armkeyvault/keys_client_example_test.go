//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/createKey.json
func ExampleKeysClient_CreateIfNotExist() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKeysClient().CreateIfNotExist(ctx, "sample-group", "sample-vault-name", "sample-key-name", armkeyvault.KeyCreateParameters{
		Properties: &armkeyvault.KeyProperties{
			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Key = armkeyvault.Key{
	// 	Name: to.Ptr("sample-key-name"),
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults/keys"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkeyvault.KeyProperties{
	// 		Attributes: &armkeyvault.KeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/getKey.json
func ExampleKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKeysClient().Get(ctx, "sample-group", "sample-vault-name", "sample-key-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Key = armkeyvault.Key{
	// 	Name: to.Ptr("sample-key-name"),
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults/keys"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkeyvault.KeyProperties{
	// 		Attributes: &armkeyvault.KeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/listKeys.json
func ExampleKeysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKeysClient().NewListPager("sample-group", "sample-vault-name", nil)
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
		// page.KeyListResult = armkeyvault.KeyListResult{
		// 	Value: []*armkeyvault.Key{
		// 		{
		// 			Name: to.Ptr("sample-key-name-1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name-1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.KeyProperties{
		// 				Attributes: &armkeyvault.KeyAttributes{
		// 					Created: to.Ptr[int64](1596493796),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1596493796),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name-1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-key-name-2"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name-2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.KeyProperties{
		// 				Attributes: &armkeyvault.KeyAttributes{
		// 					Created: to.Ptr[int64](1596493797),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1596493797),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name-2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/getKeyVersion.json
func ExampleKeysClient_GetVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKeysClient().GetVersion(ctx, "sample-group", "sample-vault-name", "sample-key-name", "fd618d9519b74f9aae94ade66b876acc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Key = armkeyvault.Key{
	// 	Name: to.Ptr("fd618d9519b74f9aae94ade66b876acc"),
	// 	Type: to.Ptr("Microsoft.KeyVault/vaults/keys/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name/versions/fd618d9519b74f9aae94ade66b876acc"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkeyvault.KeyProperties{
	// 		Attributes: &armkeyvault.KeyAttributes{
	// 			Created: to.Ptr[int64](1598533051),
	// 			Enabled: to.Ptr(true),
	// 			RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
	// 			Updated: to.Ptr[int64](1598533051),
	// 		},
	// 		KeyOps: []*armkeyvault.JSONWebKeyOperation{
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationEncrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationDecrypt),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationSign),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationVerify),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationWrapKey),
	// 			to.Ptr(armkeyvault.JSONWebKeyOperationUnwrapKey)},
	// 			KeySize: to.Ptr[int32](2048),
	// 			KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
	// 			KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/fd618d9519b74f9aae94ade66b876acc"),
	// 			Kty: to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/listKeyVersions.json
func ExampleKeysClient_NewListVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKeysClient().NewListVersionsPager("sample-group", "sample-vault-name", "sample-key-name", nil)
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
		// page.KeyListResult = armkeyvault.KeyListResult{
		// 	Value: []*armkeyvault.Key{
		// 		{
		// 			Name: to.Ptr("c2296aa24acf4daf86942bff5aca73dd"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/keys/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name/versions/c2296aa24acf4daf86942bff5aca73dd"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.KeyProperties{
		// 				Attributes: &armkeyvault.KeyAttributes{
		// 					Created: to.Ptr[int64](1598641074),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1598641074),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
		// 				KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/c2296aa24acf4daf86942bff5aca73dd"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("d5a04667b6f44b0ca62825f5eae93da6"),
		// 			Type: to.Ptr("Microsoft.KeyVault/vaults/keys/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault-name/keys/sample-key-name/versions/d5a04667b6f44b0ca62825f5eae93da6"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkeyvault.KeyProperties{
		// 				Attributes: &armkeyvault.KeyAttributes{
		// 					Created: to.Ptr[int64](1598641295),
		// 					Enabled: to.Ptr(true),
		// 					RecoveryLevel: to.Ptr(armkeyvault.DeletionRecoveryLevelPurgeable),
		// 					Updated: to.Ptr[int64](1598641295),
		// 				},
		// 				KeyURI: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name"),
		// 				KeyURIWithVersion: to.Ptr("https://sample-vault-name.vault.azure.net:443/keys/sample-key-name/d5a04667b6f44b0ca62825f5eae93da6"),
		// 			},
		// 	}},
		// }
	}
}
