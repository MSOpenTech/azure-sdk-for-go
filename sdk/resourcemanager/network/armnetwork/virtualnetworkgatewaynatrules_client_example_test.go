//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualNetworkGatewayNatRuleGet.json
func ExampleVirtualNetworkGatewayNatRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewayNatRulesClient().Get(ctx, "rg1", "gateway1", "natRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkGatewayNatRule = armnetwork.VirtualNetworkGatewayNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/natRules/natRule1"),
	// 	Name: to.Ptr("natRule1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 		Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 		ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("50.4.0.0/24"),
	// 				PortRange: to.Ptr("200-200"),
	// 		}},
	// 		InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("10.4.0.0/24"),
	// 				PortRange: to.Ptr("100-100"),
	// 		}},
	// 		IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default"),
	// 		Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualNetworkGatewayNatRulePut.json
func ExampleVirtualNetworkGatewayNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewayNatRulesClient().BeginCreateOrUpdate(ctx, "rg1", "gateway1", "natRule1", armnetwork.VirtualNetworkGatewayNatRule{
		Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
			Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
			ExternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("192.168.21.0/24"),
					PortRange:    to.Ptr("300-400"),
				}},
			InternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("10.4.0.0/24"),
					PortRange:    to.Ptr("200-300"),
				}},
			IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default"),
			Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
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
	// res.VirtualNetworkGatewayNatRule = armnetwork.VirtualNetworkGatewayNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/natRules/natRule1"),
	// 	Name: to.Ptr("natRule1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualNetworkGatewayNatRuleProperties{
	// 		Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 		ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("192.168.21.0/24"),
	// 				PortRange: to.Ptr("300-400"),
	// 		}},
	// 		InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("10.4.0.0/24"),
	// 				PortRange: to.Ptr("200-300"),
	// 		}},
	// 		IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default"),
	// 		Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualNetworkGatewayNatRuleDelete.json
func ExampleVirtualNetworkGatewayNatRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkGatewayNatRulesClient().BeginDelete(ctx, "rg1", "gateway1", "natRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualNetworkGatewayNatRuleList.json
func ExampleVirtualNetworkGatewayNatRulesClient_NewListByVirtualNetworkGatewayPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkGatewayNatRulesClient().NewListByVirtualNetworkGatewayPager("rg1", "gateway1", nil)
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
		// page.ListVirtualNetworkGatewayNatRulesResult = armnetwork.ListVirtualNetworkGatewayNatRulesResult{
		// }
	}
}
