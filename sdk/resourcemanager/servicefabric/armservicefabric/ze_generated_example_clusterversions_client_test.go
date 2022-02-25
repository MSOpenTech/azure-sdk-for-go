//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGet_example.json
func ExampleClusterVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClusterVersionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<location>",
		"<cluster-version>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClusterVersionsClientGetResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGetByEnvironment_example.json
func ExampleClusterVersionsClient_GetByEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClusterVersionsClient("<subscription-id>", cred, nil)
	res, err := client.GetByEnvironment(ctx,
		"<location>",
		armservicefabric.ClusterVersionsEnvironment("Windows"),
		"<cluster-version>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClusterVersionsClientGetByEnvironmentResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsList_example.json
func ExampleClusterVersionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClusterVersionsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<location>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClusterVersionsClientListResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsListByEnvironment.json
func ExampleClusterVersionsClient_ListByEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClusterVersionsClient("<subscription-id>", cred, nil)
	res, err := client.ListByEnvironment(ctx,
		"<location>",
		armservicefabric.ClusterVersionsEnvironment("Windows"),
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClusterVersionsClientListByEnvironmentResult)
}
