//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_List_MaximumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_NewListPager_prefixListGlobalRulestackListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrefixListGlobalRulestackClient().NewListPager("praval", nil)
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
		// page.PrefixListGlobalRulestackResourceListResult = armpanngfw.PrefixListGlobalRulestackResourceListResult{
		// 	Value: []*armpanngfw.PrefixListGlobalRulestackResource{
		// 		{
		// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 			SystemData: &armpanngfw.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				CreatedBy: to.Ptr("praval"),
		// 				CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("praval"),
		// 				LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 			},
		// 			Properties: &armpanngfw.PrefixObject{
		// 				Description: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 				AuditComment: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 				PrefixList: []*string{
		// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 					ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_List_MinimumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_NewListPager_prefixListGlobalRulestackListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrefixListGlobalRulestackClient().NewListPager("praval", nil)
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
		// page.PrefixListGlobalRulestackResourceListResult = armpanngfw.PrefixListGlobalRulestackResourceListResult{
		// 	Value: []*armpanngfw.PrefixListGlobalRulestackResource{
		// 		{
		// 			ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/prefixlists/prefixlists1"),
		// 			Properties: &armpanngfw.PrefixObject{
		// 				PrefixList: []*string{
		// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_Get_MaximumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_Get_prefixListGlobalRulestackGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrefixListGlobalRulestackClient().Get(ctx, "praval", "armid1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrefixListGlobalRulestackResource = armpanngfw.PrefixListGlobalRulestackResource{
	// 	Name: to.Ptr("armid1"),
	// 	Type: to.Ptr("certificates"),
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalrulestacks/armid1/certificates/armid1"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Properties: &armpanngfw.PrefixObject{
	// 		Description: to.Ptr("string"),
	// 		AuditComment: to.Ptr("comment"),
	// 		Etag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 		PrefixList: []*string{
	// 			to.Ptr("1.0.0.0/24")},
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_Get_MinimumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_Get_prefixListGlobalRulestackGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrefixListGlobalRulestackClient().Get(ctx, "praval", "armid1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrefixListGlobalRulestackResource = armpanngfw.PrefixListGlobalRulestackResource{
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/prefixlists/armid1"),
	// 	Properties: &armpanngfw.PrefixObject{
	// 		PrefixList: []*string{
	// 			to.Ptr("1.0.0.0/24")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_BeginCreateOrUpdate_prefixListGlobalRulestackCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListGlobalRulestackClient().BeginCreateOrUpdate(ctx, "praval", "armid1", armpanngfw.PrefixListGlobalRulestackResource{
		Properties: &armpanngfw.PrefixObject{
			Description:  to.Ptr("string"),
			AuditComment: to.Ptr("comment"),
			Etag:         to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
			PrefixList: []*string{
				to.Ptr("1.0.0.0/24")},
			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
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
	// res.PrefixListGlobalRulestackResource = armpanngfw.PrefixListGlobalRulestackResource{
	// 	Name: to.Ptr("armid1"),
	// 	Type: to.Ptr("certificates"),
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalrulestacks/armid1/certificates/armid1"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Properties: &armpanngfw.PrefixObject{
	// 		Description: to.Ptr("string"),
	// 		AuditComment: to.Ptr("comment"),
	// 		Etag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 		PrefixList: []*string{
	// 			to.Ptr("1.0.0.0/24")},
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_BeginCreateOrUpdate_prefixListGlobalRulestackCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListGlobalRulestackClient().BeginCreateOrUpdate(ctx, "praval", "armid1", armpanngfw.PrefixListGlobalRulestackResource{
		Properties: &armpanngfw.PrefixObject{
			PrefixList: []*string{
				to.Ptr("1.0.0.0/24")},
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
	// res.PrefixListGlobalRulestackResource = armpanngfw.PrefixListGlobalRulestackResource{
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/praval/prefixlists/armid1"),
	// 	Properties: &armpanngfw.PrefixObject{
	// 		PrefixList: []*string{
	// 			to.Ptr("1.0.0.0/24")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_Delete_MaximumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_BeginDelete_prefixListGlobalRulestackDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListGlobalRulestackClient().BeginDelete(ctx, "praval", "armid1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListGlobalRulestack_Delete_MinimumSet_Gen.json
func ExamplePrefixListGlobalRulestackClient_BeginDelete_prefixListGlobalRulestackDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrefixListGlobalRulestackClient().BeginDelete(ctx, "praval", "armid1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
