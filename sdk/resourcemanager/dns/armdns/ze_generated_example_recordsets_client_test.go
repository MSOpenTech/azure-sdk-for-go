//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/PatchARecordset.json
func ExampleRecordSetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<zone-name>",
		"<relative-record-set-name>",
		armdns.RecordTypeA,
		armdns.RecordSet{
			Properties: &armdns.RecordSetProperties{
				Metadata: map[string]*string{
					"key2": to.StringPtr("value2"),
				},
			},
		},
		&armdns.RecordSetsUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RecordSet.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateARecordset.json
func ExampleRecordSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<zone-name>",
		"<relative-record-set-name>",
		armdns.RecordTypeA,
		armdns.RecordSet{
			Properties: &armdns.RecordSetProperties{
				ARecords: []*armdns.ARecord{
					{
						IPv4Address: to.StringPtr("<ipv4address>"),
					}},
				TTL: to.Int64Ptr(3600),
				Metadata: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
			},
		},
		&armdns.RecordSetsCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RecordSet.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeleteARecordset.json
func ExampleRecordSetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<zone-name>",
		"<relative-record-set-name>",
		armdns.RecordTypeA,
		&armdns.RecordSetsDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetARecordset.json
func ExampleRecordSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<zone-name>",
		"<relative-record-set-name>",
		armdns.RecordTypeA,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RecordSet.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListARecordset.json
func ExampleRecordSetsClient_ListByType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	pager := client.ListByType("<resource-group-name>",
		"<zone-name>",
		armdns.RecordTypeA,
		&armdns.RecordSetsListByTypeOptions{Top: nil,
			Recordsetnamesuffix: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RecordSet.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListRecordSetsByZone.json
func ExampleRecordSetsClient_ListByDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	pager := client.ListByDNSZone("<resource-group-name>",
		"<zone-name>",
		&armdns.RecordSetsListByDNSZoneOptions{Top: nil,
			Recordsetnamesuffix: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RecordSet.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListRecordSetsByZone.json
func ExampleRecordSetsClient_ListAllByDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdns.NewRecordSetsClient("<subscription-id>", cred, nil)
	pager := client.ListAllByDNSZone("<resource-group-name>",
		"<zone-name>",
		&armdns.RecordSetsListAllByDNSZoneOptions{Top: nil,
			RecordSetNameSuffix: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RecordSet.ID: %s\n", *v.ID)
		}
	}
}
