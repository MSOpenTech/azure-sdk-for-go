//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnewrelic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Organizations_List_MaximumSet_Gen.json
func ExampleOrganizationsClient_NewListPager_organizationsListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnewrelic.NewOrganizationsClient("nqmcgifgaqlf", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("ruxvg@xqkmdhrnoo.hlmbpm", "egh", nil)
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
		// page.OrganizationsListResponse = armnewrelic.OrganizationsListResponse{
		// 	Value: []*armnewrelic.OrganizationResource{
		// 		{
		// 			ID: to.Ptr("ycdsgeiitvxcd"),
		// 			Properties: &armnewrelic.OrganizationProperties{
		// 				BillingSource: to.Ptr(armnewrelic.BillingSourceAZURE),
		// 				OrganizationID: to.Ptr("tjmcifofzirili"),
		// 				OrganizationName: to.Ptr("orgname"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Organizations_List_MinimumSet_Gen.json
func ExampleOrganizationsClient_NewListPager_organizationsListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnewrelic.NewOrganizationsClient("nqmcgifgaqlf", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("ruxvg@xqkmdhrnoo.hlmbpm", "egh", nil)
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
		// page.OrganizationsListResponse = armnewrelic.OrganizationsListResponse{
		// 	Value: []*armnewrelic.OrganizationResource{
		// 		{
		// 			ID: to.Ptr("ycdsgeiitvxcd"),
		// 	}},
		// }
	}
}
