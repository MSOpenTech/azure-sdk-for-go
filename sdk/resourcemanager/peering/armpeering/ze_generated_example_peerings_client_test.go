//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/GetPeering.json
func ExamplePeeringsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<peering-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Peering.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/CreateDirectPeering.json
func ExamplePeeringsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-name>",
		armpeering.Peering{
			Kind:     armpeering.KindDirect.ToPtr(),
			Location: to.StringPtr("<location>"),
			Properties: &armpeering.PeeringProperties{
				Direct: &armpeering.PeeringPropertiesDirect{
					Connections: []*armpeering.DirectConnection{
						{
							BandwidthInMbps: to.Int32Ptr(10000),
							BgpSession: &armpeering.BgpSession{
								MaxPrefixesAdvertisedV4: to.Int32Ptr(1000),
								MaxPrefixesAdvertisedV6: to.Int32Ptr(100),
								MD5AuthenticationKey:    to.StringPtr("<md5authentication-key>"),
								SessionPrefixV4:         to.StringPtr("<session-prefix-v4>"),
								SessionPrefixV6:         to.StringPtr("<session-prefix-v6>"),
							},
							ConnectionIdentifier:   to.StringPtr("<connection-identifier>"),
							PeeringDBFacilityID:    to.Int32Ptr(99999),
							SessionAddressProvider: armpeering.SessionAddressProviderPeer.ToPtr(),
							UseForPeeringService:   to.BoolPtr(false),
						},
						{
							BandwidthInMbps: to.Int32Ptr(10000),
							BgpSession: &armpeering.BgpSession{
								MaxPrefixesAdvertisedV4: to.Int32Ptr(1000),
								MaxPrefixesAdvertisedV6: to.Int32Ptr(100),
								MD5AuthenticationKey:    to.StringPtr("<md5authentication-key>"),
								SessionPrefixV4:         to.StringPtr("<session-prefix-v4>"),
								SessionPrefixV6:         to.StringPtr("<session-prefix-v6>"),
							},
							ConnectionIdentifier:   to.StringPtr("<connection-identifier>"),
							PeeringDBFacilityID:    to.Int32Ptr(99999),
							SessionAddressProvider: armpeering.SessionAddressProviderMicrosoft.ToPtr(),
							UseForPeeringService:   to.BoolPtr(true),
						}},
					DirectPeeringType: armpeering.DirectPeeringTypeEdge.ToPtr(),
					PeerAsn: &armpeering.SubResource{
						ID: to.StringPtr("<id>"),
					},
					UseForPeeringService: to.BoolPtr(false),
				},
				PeeringLocation: to.StringPtr("<peering-location>"),
			},
			SKU: &armpeering.PeeringSKU{
				Name: armpeering.NameBasicDirectFree.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Peering.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/DeletePeering.json
func ExamplePeeringsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<peering-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/UpdatePeeringTags.json
func ExamplePeeringsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<peering-name>",
		armpeering.ResourceTags{
			Tags: map[string]*string{
				"tag0": to.StringPtr("value0"),
				"tag1": to.StringPtr("value1"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Peering.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/ListPeeringsByResourceGroup.json
func ExamplePeeringsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Peering.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/ListPeeringsBySubscription.json
func ExamplePeeringsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Peering.ID: %s\n", *v.ID)
		}
	}
}
