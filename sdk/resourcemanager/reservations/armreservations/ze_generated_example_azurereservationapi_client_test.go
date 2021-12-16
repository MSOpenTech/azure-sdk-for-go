//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/GetCatalog.json
func ExampleAzureReservationAPIClient_GetCatalog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewAzureReservationAPIClient(cred, nil)
	_, err = client.GetCatalog(ctx,
		"<subscription-id>",
		&armreservations.AzureReservationAPIGetCatalogOptions{ReservedResourceType: to.StringPtr("<reserved-resource-type>"),
			Location: to.StringPtr("<location>"),
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/GetAppliedReservations.json
func ExampleAzureReservationAPIClient_GetAppliedReservationList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewAzureReservationAPIClient(cred, nil)
	res, err := client.GetAppliedReservationList(ctx,
		"<subscription-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AppliedReservations.ID: %s\n", *res.ID)
}
