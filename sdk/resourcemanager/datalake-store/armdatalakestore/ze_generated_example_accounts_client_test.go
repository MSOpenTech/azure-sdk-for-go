//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_List.json
func ExampleAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armdatalakestore.AccountsClientListOptions{Filter: to.Ptr("test_filter"),
		Top:     to.Ptr[int32](1),
		Skip:    to.Ptr[int32](1),
		Select:  to.Ptr("test_select"),
		Orderby: to.Ptr("test_orderby"),
		Count:   to.Ptr(false),
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_ListByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("contosorg",
		&armdatalakestore.AccountsClientListByResourceGroupOptions{Filter: to.Ptr("test_filter"),
			Top:     to.Ptr[int32](1),
			Skip:    to.Ptr[int32](1),
			Select:  to.Ptr("test_select"),
			Orderby: to.Ptr("test_orderby"),
			Count:   to.Ptr(false),
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contosorg",
		"contosoadla",
		armdatalakestore.CreateDataLakeStoreAccountParameters{
			Identity: &armdatalakestore.EncryptionIdentity{
				Type: to.Ptr("SystemAssigned"),
			},
			Location: to.Ptr("eastus2"),
			Properties: &armdatalakestore.CreateDataLakeStoreAccountProperties{
				DefaultGroup: to.Ptr("test_default_group"),
				EncryptionConfig: &armdatalakestore.EncryptionConfig{
					Type: to.Ptr(armdatalakestore.EncryptionConfigTypeUserManaged),
					KeyVaultMetaInfo: &armdatalakestore.KeyVaultMetaInfo{
						EncryptionKeyName:    to.Ptr("test_encryption_key_name"),
						EncryptionKeyVersion: to.Ptr("encryption_key_version"),
						KeyVaultResourceID:   to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
					},
				},
				EncryptionState:       to.Ptr(armdatalakestore.EncryptionStateEnabled),
				FirewallAllowAzureIPs: to.Ptr(armdatalakestore.FirewallAllowAzureIPsStateEnabled),
				FirewallRules: []*armdatalakestore.CreateFirewallRuleWithAccountParameters{
					{
						Name: to.Ptr("test_rule"),
						Properties: &armdatalakestore.CreateOrUpdateFirewallRuleProperties{
							EndIPAddress:   to.Ptr("2.2.2.2"),
							StartIPAddress: to.Ptr("1.1.1.1"),
						},
					}},
				FirewallState:          to.Ptr(armdatalakestore.FirewallStateEnabled),
				NewTier:                to.Ptr(armdatalakestore.TierTypeConsumption),
				TrustedIDProviderState: to.Ptr(armdatalakestore.TrustedIDProviderStateEnabled),
				TrustedIDProviders: []*armdatalakestore.CreateTrustedIDProviderWithAccountParameters{
					{
						Name: to.Ptr("test_trusted_id_provider_name"),
						Properties: &armdatalakestore.CreateOrUpdateTrustedIDProviderProperties{
							IDProvider: to.Ptr("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
						},
					}},
			},
			Tags: map[string]*string{
				"test_key": to.Ptr("test_value"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"contosorg",
		"contosoadla",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Update.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"contosorg",
		"contosoadla",
		armdatalakestore.UpdateDataLakeStoreAccountParameters{
			Properties: &armdatalakestore.UpdateDataLakeStoreAccountProperties{
				DefaultGroup: to.Ptr("test_default_group"),
				EncryptionConfig: &armdatalakestore.UpdateEncryptionConfig{
					KeyVaultMetaInfo: &armdatalakestore.UpdateKeyVaultMetaInfo{
						EncryptionKeyVersion: to.Ptr("encryption_key_version"),
					},
				},
				FirewallAllowAzureIPs:  to.Ptr(armdatalakestore.FirewallAllowAzureIPsStateEnabled),
				FirewallState:          to.Ptr(armdatalakestore.FirewallStateEnabled),
				NewTier:                to.Ptr(armdatalakestore.TierTypeConsumption),
				TrustedIDProviderState: to.Ptr(armdatalakestore.TrustedIDProviderStateEnabled),
			},
			Tags: map[string]*string{
				"test_key": to.Ptr("test_value"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Delete.json
func ExampleAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"contosorg",
		"contosoadla",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_EnableKeyVault.json
func ExampleAccountsClient_EnableKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.EnableKeyVault(ctx,
		"contosorg",
		"contosoadla",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_CheckNameAvailability.json
func ExampleAccountsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx,
		"EastUS2",
		armdatalakestore.CheckNameAvailabilityParameters{
			Name: to.Ptr("contosoadla"),
			Type: to.Ptr(armdatalakestore.CheckNameAvailabilityParametersTypeMicrosoftDataLakeStoreAccounts),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
