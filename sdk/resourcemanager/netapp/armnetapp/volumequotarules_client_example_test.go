//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/VolumeQuotaRules_List.json
func ExampleVolumeQuotaRulesClient_NewListByVolumePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeQuotaRulesClient().NewListByVolumePager("myRG", "account-9957", "pool-5210", "volume-6387", nil)
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
		// page.VolumeQuotaRulesList = armnetapp.VolumeQuotaRulesList{
		// 	Value: []*armnetapp.VolumeQuotaRule{
		// 		{
		// 			Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
		// 			ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetapp.VolumeQuotaRulesProperties{
		// 				ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 				QuotaSizeInKiBs: to.Ptr[int64](100005),
		// 				QuotaTarget: to.Ptr("1821"),
		// 				QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/VolumeQuotaRules_Get.json
func ExampleVolumeQuotaRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeQuotaRulesClient().Get(ctx, "myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeQuotaRule = armnetapp.VolumeQuotaRule{
	// 	Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
	// 	ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.VolumeQuotaRulesProperties{
	// 		ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 		QuotaSizeInKiBs: to.Ptr[int64](100005),
	// 		QuotaTarget: to.Ptr("1821"),
	// 		QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/VolumeQuotaRules_Create.json
func ExampleVolumeQuotaRulesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeQuotaRulesClient().BeginCreate(ctx, "myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", armnetapp.VolumeQuotaRule{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.VolumeQuotaRulesProperties{
			QuotaSizeInKiBs: to.Ptr[int64](100005),
			QuotaTarget:     to.Ptr("1821"),
			QuotaType:       to.Ptr(armnetapp.TypeIndividualUserQuota),
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
	// res.VolumeQuotaRule = armnetapp.VolumeQuotaRule{
	// 	Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
	// 	ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.VolumeQuotaRulesProperties{
	// 		ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 		QuotaSizeInKiBs: to.Ptr[int64](100005),
	// 		QuotaTarget: to.Ptr("1821"),
	// 		QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/VolumeQuotaRules_Update.json
func ExampleVolumeQuotaRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeQuotaRulesClient().BeginUpdate(ctx, "myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", armnetapp.VolumeQuotaRulePatch{
		Properties: &armnetapp.VolumeQuotaRulesProperties{
			QuotaSizeInKiBs: to.Ptr[int64](100009),
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
	// res.VolumeQuotaRule = armnetapp.VolumeQuotaRule{
	// 	Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
	// 	ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.VolumeQuotaRulesProperties{
	// 		ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 		QuotaSizeInKiBs: to.Ptr[int64](100005),
	// 		QuotaTarget: to.Ptr("1821"),
	// 		QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43042075540b8d67cce7d3d9f70b9b9f5a359da/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/VolumeQuotaRules_Delete.json
func ExampleVolumeQuotaRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeQuotaRulesClient().BeginDelete(ctx, "myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
