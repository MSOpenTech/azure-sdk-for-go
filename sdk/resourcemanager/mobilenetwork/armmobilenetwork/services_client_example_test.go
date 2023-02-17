//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/ServiceDelete.json
func ExampleServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "testMobileNetwork", "TestService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/ServiceGet.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "testMobileNetwork", "TestService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testMobileNetwork", "TestService", armmobilenetwork.Service{
		Location: to.Ptr("eastus"),
		Properties: &armmobilenetwork.ServicePropertiesFormat{
			PccRules: []*armmobilenetwork.PccRuleConfiguration{
				{
					RuleName:       to.Ptr("default-rule"),
					RulePrecedence: to.Ptr[int32](255),
					RuleQosPolicy: &armmobilenetwork.PccRuleQosPolicy{
						FiveQi:                              to.Ptr[int32](9),
						AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
						MaximumBitRate: &armmobilenetwork.Ambr{
							Downlink: to.Ptr("1 Gbps"),
							Uplink:   to.Ptr("500 Mbps"),
						},
						PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
						PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
					},
					ServiceDataFlowTemplates: []*armmobilenetwork.ServiceDataFlowTemplate{
						{
							Direction: to.Ptr(armmobilenetwork.SdfDirectionUplink),
							Ports:     []*string{},
							RemoteIPList: []*string{
								to.Ptr("10.3.4.0/24")},
							TemplateName: to.Ptr("IP-to-server"),
							Protocol: []*string{
								to.Ptr("ip")},
						}},
					TrafficControl: to.Ptr(armmobilenetwork.TrafficControlPermissionEnabled),
				}},
			ServicePrecedence: to.Ptr[int32](255),
			ServiceQosPolicy: &armmobilenetwork.QosPolicy{
				FiveQi:                              to.Ptr[int32](9),
				AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
				MaximumBitRate: &armmobilenetwork.Ambr{
					Downlink: to.Ptr("1 Gbps"),
					Uplink:   to.Ptr("500 Mbps"),
				},
				PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
				PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
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
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/ServiceUpdateTags.json
func ExampleServicesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx, "rg1", "testMobileNetwork", "TestService", armmobilenetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/ServiceListByMobileNetwork.json
func ExampleServicesClient_NewListByMobileNetworkPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByMobileNetworkPager("testResourceGroupName", "testMobileNetwork", nil)
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
