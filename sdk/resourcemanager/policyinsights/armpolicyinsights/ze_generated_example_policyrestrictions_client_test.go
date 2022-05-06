//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtSubscriptionScope.json
func ExamplePolicyRestrictionsClient_CheckAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyRestrictionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckAtSubscriptionScope(ctx,
		armpolicyinsights.CheckRestrictionsRequest{
			PendingFields: []*armpolicyinsights.PendingField{
				{
					Field: to.Ptr("<field>"),
					Values: []*string{
						to.Ptr("myVMName")},
				},
				{
					Field: to.Ptr("<field>"),
					Values: []*string{
						to.Ptr("eastus"),
						to.Ptr("westus"),
						to.Ptr("westus2"),
						to.Ptr("westeurope")},
				},
				{
					Field: to.Ptr("<field>"),
				}},
			ResourceDetails: &armpolicyinsights.CheckRestrictionsResourceDetails{
				APIVersion: to.Ptr("<apiversion>"),
				ResourceContent: map[string]interface{}{
					"type": "Microsoft.Compute/virtualMachines",
					"properties": map[string]interface{}{
						"priority": "Spot",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtResourceGroupScope.json
func ExamplePolicyRestrictionsClient_CheckAtResourceGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyRestrictionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckAtResourceGroupScope(ctx,
		"<resource-group-name>",
		armpolicyinsights.CheckRestrictionsRequest{
			PendingFields: []*armpolicyinsights.PendingField{
				{
					Field: to.Ptr("<field>"),
					Values: []*string{
						to.Ptr("myVMName")},
				},
				{
					Field: to.Ptr("<field>"),
					Values: []*string{
						to.Ptr("eastus"),
						to.Ptr("westus"),
						to.Ptr("westus2"),
						to.Ptr("westeurope")},
				},
				{
					Field: to.Ptr("<field>"),
				}},
			ResourceDetails: &armpolicyinsights.CheckRestrictionsResourceDetails{
				APIVersion: to.Ptr("<apiversion>"),
				ResourceContent: map[string]interface{}{
					"type": "Microsoft.Compute/virtualMachines",
					"properties": map[string]interface{}{
						"priority": "Spot",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtManagementGroupScope.json
func ExamplePolicyRestrictionsClient_CheckAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyRestrictionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckAtManagementGroupScope(ctx,
		"<management-group-id>",
		armpolicyinsights.CheckManagementGroupRestrictionsRequest{
			PendingFields: []*armpolicyinsights.PendingField{
				{
					Field: to.Ptr("<field>"),
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
