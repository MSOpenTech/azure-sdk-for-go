//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ListAppServicePlans.json
func ExampleAppServicePlansClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	pager := client.List(&armappservice.AppServicePlansListOptions{Detailed: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("AppServicePlan.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ListAppServicePlansByResourceGroup.json
func ExampleAppServicePlansClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("AppServicePlan.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetAppServicePlan.json
func ExampleAppServicePlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AppServicePlan.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/CreateOrUpdateAppServicePlan.json
func ExampleAppServicePlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.AppServicePlan{
			Resource: armappservice.Resource{
				Kind:     to.StringPtr("<kind>"),
				Location: to.StringPtr("<location>"),
			},
			Properties: &armappservice.AppServicePlanProperties{},
			SKU: &armappservice.SKUDescription{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(1),
				Family:   to.StringPtr("<family>"),
				Size:     to.StringPtr("<size>"),
				Tier:     to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AppServicePlan.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/DeleteAppServicePlan.json
func ExampleAppServicePlansClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/PatchAppServicePlan.json
func ExampleAppServicePlansClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewAppServicePlansClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.AppServicePlanPatchResource{
			ProxyOnlyResource: armappservice.ProxyOnlyResource{
				Kind: to.StringPtr("<kind>"),
			},
			Properties: &armappservice.AppServicePlanPatchResourceProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AppServicePlan.ID: %s\n", *res.ID)
}
