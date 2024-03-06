//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationPolicies_List.json
func ExampleReplicationPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationPoliciesClient().NewListPager("vault1", "resourceGroupPS1", nil)
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
		// page.PolicyCollection = armrecoveryservicessiterecovery.PolicyCollection{
		// 	Value: []*armrecoveryservicessiterecovery.Policy{
		// 		{
		// 			Name: to.Ptr("protectionprofile1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
		// 			Properties: &armrecoveryservicessiterecovery.PolicyProperties{
		// 				FriendlyName: to.Ptr("protectionprofile1"),
		// 				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyDetails{
		// 					InstanceType: to.Ptr("HyperVReplicaAzure"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationPolicies_Get.json
func ExampleReplicationPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationPoliciesClient().Get(ctx, "vault1", "resourceGroupPS1", "protectionprofile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Policy = armrecoveryservicessiterecovery.Policy{
	// 	Name: to.Ptr("protectionprofile1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
	// 	Properties: &armrecoveryservicessiterecovery.PolicyProperties{
	// 		FriendlyName: to.Ptr("protectionprofile1"),
	// 		ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyDetails{
	// 			InstanceType: to.Ptr("HyperVReplicaAzure"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationPolicies_Create.json
func ExampleReplicationPoliciesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationPoliciesClient().BeginCreate(ctx, "vault1", "resourceGroupPS1", "protectionprofile1", armrecoveryservicessiterecovery.CreatePolicyInput{
		Properties: &armrecoveryservicessiterecovery.CreatePolicyInputProperties{
			ProviderSpecificInput: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyInput{
				InstanceType: to.Ptr("HyperVReplicaAzure"),
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
	// res.Policy = armrecoveryservicessiterecovery.Policy{
	// 	Name: to.Ptr("protectionprofile1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
	// 	Properties: &armrecoveryservicessiterecovery.PolicyProperties{
	// 		FriendlyName: to.Ptr("protectionprofile1"),
	// 		ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyDetails{
	// 			InstanceType: to.Ptr("HyperVReplicaAzure"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationPolicies_Delete.json
func ExampleReplicationPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationPoliciesClient().BeginDelete(ctx, "vault1", "resourceGroupPS1", "protectionprofile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationPolicies_Update.json
func ExampleReplicationPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationPoliciesClient().BeginUpdate(ctx, "vault1", "resourceGroupPS1", "protectionprofile1", armrecoveryservicessiterecovery.UpdatePolicyInput{
		Properties: &armrecoveryservicessiterecovery.UpdatePolicyInputProperties{
			ReplicationProviderSettings: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyInput{
				InstanceType: to.Ptr("HyperVReplicaAzure"),
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
	// res.Policy = armrecoveryservicessiterecovery.Policy{
	// 	Name: to.Ptr("protectionprofile1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationPolicies"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
	// 	Properties: &armrecoveryservicessiterecovery.PolicyProperties{
	// 		FriendlyName: to.Ptr("protectionprofile1"),
	// 		ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyDetails{
	// 			InstanceType: to.Ptr("HyperVReplicaAzure"),
	// 		},
	// 	},
	// }
}
