//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesPropertiesQuery.json
func ExampleResourceGraphClient_Resources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresourcegraph.NewResourceGraphClient(cred, nil)
	_, err = client.Resources(ctx,
		armresourcegraph.QueryRequest{
			Query: to.StringPtr("<query>"),
			Subscriptions: []*string{
				to.StringPtr("cfbbd179-59d2-4052-aa06-9270a38aa9d6")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesHistoryMgsGet.json
func ExampleResourceGraphClient_ResourcesHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresourcegraph.NewResourceGraphClient(cred, nil)
	_, err = client.ResourcesHistory(ctx,
		armresourcegraph.ResourcesHistoryRequest{
			ManagementGroups: []*string{
				to.StringPtr("e927f598-c1d4-4f72-8541-95d83a6a4ac8"),
				to.StringPtr("ProductionMG")},
			Options: &armresourcegraph.ResourcesHistoryRequestOptions{
				Interval: &armresourcegraph.DateTimeInterval{
					End:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T01:25:00.0000000Z"); return t }()),
					Start: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T01:00:00.0000000Z"); return t }()),
				},
			},
			Query: to.StringPtr("<query>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
