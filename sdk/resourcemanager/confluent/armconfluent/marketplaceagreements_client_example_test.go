//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/MarketplaceAgreements_List.json
func ExampleMarketplaceAgreementsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMarketplaceAgreementsClient().NewListPager(nil)
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
		// page.AgreementResourceListResponse = armconfluent.AgreementResourceListResponse{
		// 	Value: []*armconfluent.AgreementResource{
		// 		{
		// 			Name: to.Ptr("planid1"),
		// 			Type: to.Ptr("Microsoft.Confluent/agreements"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Confluent/agreements/default"),
		// 			Properties: &armconfluent.AgreementProperties{
		// 				Accepted: to.Ptr(true),
		// 				LicenseTextLink: to.Ptr("test.licenseLink1"),
		// 				Plan: to.Ptr("planid1"),
		// 				PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink1"),
		// 				Product: to.Ptr("offid1"),
		// 				Publisher: to.Ptr("pubid1"),
		// 				RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.121Z"); return t}()),
		// 				Signature: to.Ptr("ASDFSDAFWEFASDGWERLWER"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("planid2"),
		// 			Type: to.Ptr("Microsoft.MarketplaceOrdering/offertypes"),
		// 			ID: to.Ptr("id2"),
		// 			Properties: &armconfluent.AgreementProperties{
		// 				Accepted: to.Ptr(true),
		// 				LicenseTextLink: to.Ptr("test.licenseLin2k"),
		// 				Plan: to.Ptr("planid2"),
		// 				PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink2"),
		// 				Product: to.Ptr("offid2"),
		// 				Publisher: to.Ptr("pubid2"),
		// 				RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-14T11:33:07.121Z"); return t}()),
		// 				Signature: to.Ptr("ASDFSDAFWEFASDGWERLWER"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/MarketplaceAgreements_Create.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMarketplaceAgreementsClient().Create(ctx, &armconfluent.MarketplaceAgreementsClientCreateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgreementResource = armconfluent.AgreementResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Confluent/agreements"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Confluent/agreements/default"),
	// 	Properties: &armconfluent.AgreementProperties{
	// 		Accepted: to.Ptr(true),
	// 		LicenseTextLink: to.Ptr("test.licenseLink1"),
	// 		Plan: to.Ptr("planid1"),
	// 		PrivacyPolicyLink: to.Ptr("test.privacyPolicyLink1"),
	// 		Product: to.Ptr("offid1"),
	// 		Publisher: to.Ptr("pubid1"),
	// 		RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-05T17:33:07.121Z"); return t}()),
	// 		Signature: to.Ptr("YKWOQOKH2BCKZ46O7SCKHANWEENRFRU5WB4LXDFUYWCBWTS4AG4SGQXCOZYIR5ZJCZTXRMZKYZMO2BJSL5YKPLAR4LBFRUNS6CRYE7A"),
	// 	},
	// 	SystemData: &armconfluent.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 	},
	// }
}
