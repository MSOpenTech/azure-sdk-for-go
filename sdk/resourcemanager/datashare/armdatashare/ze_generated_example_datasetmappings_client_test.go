//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Get.json
func ExampleDataSetMappingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetMappingsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-subscription-name>",
		"<data-set-mapping-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataSetMappingClassification.GetDataSetMapping().ID: %s\n", *res.GetDataSetMapping().ID)
}

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Create.json
func ExampleDataSetMappingsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetMappingsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-subscription-name>",
		"<data-set-mapping-name>",
		&armdatashare.BlobDataSetMapping{
			DataSetMapping: armdatashare.DataSetMapping{
				Kind: armdatashare.DataSetMappingKindBlob.ToPtr(),
			},
			Properties: &armdatashare.BlobMappingProperties{
				ContainerName:      to.StringPtr("<container-name>"),
				DataSetID:          to.StringPtr("<data-set-id>"),
				FilePath:           to.StringPtr("<file-path>"),
				ResourceGroup:      to.StringPtr("<resource-group>"),
				StorageAccountName: to.StringPtr("<storage-account-name>"),
				SubscriptionID:     to.StringPtr("<subscription-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataSetMappingClassification.GetDataSetMapping().ID: %s\n", *res.GetDataSetMapping().ID)
}

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Delete.json
func ExampleDataSetMappingsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetMappingsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-subscription-name>",
		"<data-set-mapping-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_ListByShareSubscription.json
func ExampleDataSetMappingsClient_ListByShareSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetMappingsClient("<subscription-id>", cred, nil)
	pager := client.ListByShareSubscription("<resource-group-name>",
		"<account-name>",
		"<share-subscription-name>",
		&armdatashare.DataSetMappingsListByShareSubscriptionOptions{SkipToken: nil,
			Filter:  nil,
			Orderby: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DataSetMappingClassification.GetDataSetMapping().ID: %s\n", *v.GetDataSetMapping().ID)
		}
	}
}
