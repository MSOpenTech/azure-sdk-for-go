//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetPrivateStoreCollectionsList.json
func ExamplePrivateStoreCollectionClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	_, err = client.List(ctx,
		"<private-store-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/GetPrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	res, err := client.Get(ctx,
		"<private-store-id>",
		"<collection-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Collection.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/CreatePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<private-store-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionCreateOrUpdateOptions{Payload: &armmarketplace.Collection{
			Properties: &armmarketplace.CollectionProperties{
				AllSubscriptions: to.BoolPtr(false),
				Claim:            to.StringPtr("<claim>"),
				CollectionName:   to.StringPtr("<collection-name>"),
				SubscriptionsList: []*string{
					to.StringPtr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.StringPtr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Collection.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/DeletePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	_, err = client.Delete(ctx,
		"<private-store-id>",
		"<collection-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/PostPrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	_, err = client.Post(ctx,
		"<private-store-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionPostOptions{Payload: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/TransferOffers.json
func ExamplePrivateStoreCollectionClient_TransferOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	_, err = client.TransferOffers(ctx,
		"<private-store-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionTransferOffersOptions{Payload: &armmarketplace.TransferOffersProperties{
			Properties: &armmarketplace.TransferOffersDetails{
				OfferIDsList: []*string{
					to.StringPtr("marketplacetestthirdparty.md-test-third-party-2"),
					to.StringPtr("marketplacetestthirdparty.md-test-third-party-3")},
				Operation: to.StringPtr("<operation>"),
				TargetCollections: []*string{
					to.StringPtr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
					to.StringPtr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
