//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscription_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscriptionLocation_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListByRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByRegionPager("<asc-location>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesResourceGroup_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesResourceGroupLocation_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListByResourceGroupAndRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupAndRegionPager("<resource-group-name>",
		"<asc-location>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<asc-location>",
		"<jit-network-access-policy-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/CreateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<asc-location>",
		"<jit-network-access-policy-name>",
		armsecurity.JitNetworkAccessPolicy{
			Kind:     to.Ptr("<kind>"),
			Location: to.Ptr("<location>"),
			Name:     to.Ptr("<name>"),
			Type:     to.Ptr("<type>"),
			ID:       to.Ptr("<id>"),
			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
				ProvisioningState: to.Ptr("<provisioning-state>"),
				Requests: []*armsecurity.JitNetworkAccessRequest{
					{
						Requestor:    to.Ptr("<requestor>"),
						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t }()),
						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
							{
								ID: to.Ptr("<id>"),
								Ports: []*armsecurity.JitNetworkAccessRequestPort{
									{
										AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
										EndTimeUTC:                 to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t }()),
										Number:                     to.Ptr[int32](3389),
										Status:                     to.Ptr(armsecurity.StatusInitiated),
										StatusReason:               to.Ptr(armsecurity.StatusReasonUserRequested),
									}},
							}},
					}},
				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
					{
						ID: to.Ptr("<id>"),
						Ports: []*armsecurity.JitNetworkAccessPortRule{
							{
								AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
								MaxRequestAccessDuration:   to.Ptr("<max-request-access-duration>"),
								Number:                     to.Ptr[int32](22),
								Protocol:                   to.Ptr(armsecurity.ProtocolAll),
							},
							{
								AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
								MaxRequestAccessDuration:   to.Ptr("<max-request-access-duration>"),
								Number:                     to.Ptr[int32](3389),
								Protocol:                   to.Ptr(armsecurity.ProtocolAll),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/DeleteJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<asc-location>",
		"<jit-network-access-policy-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/InitiateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Initiate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewJitNetworkAccessPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Initiate(ctx,
		"<resource-group-name>",
		"<asc-location>",
		"<jit-network-access-policy-name>",
		armsecurity.JitNetworkAccessPolicyInitiateRequest{
			Justification: to.Ptr("<justification>"),
			VirtualMachines: []*armsecurity.JitNetworkAccessPolicyInitiateVirtualMachine{
				{
					ID: to.Ptr("<id>"),
					Ports: []*armsecurity.JitNetworkAccessPolicyInitiatePort{
						{
							AllowedSourceAddressPrefix: to.Ptr("<allowed-source-address-prefix>"),
							Number:                     to.Ptr[int32](3389),
						}},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
