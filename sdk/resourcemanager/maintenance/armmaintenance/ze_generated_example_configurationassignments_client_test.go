//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_GetParent.json
func ExampleConfigurationAssignmentsClient_GetParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.GetParent(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-parent-type>",
		"<resource-parent-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_CreateOrUpdateParent.json
func ExampleConfigurationAssignmentsClient_CreateOrUpdateParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateParent(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-parent-type>",
		"<resource-parent-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		armmaintenance.ConfigurationAssignment{
			Properties: &armmaintenance.ConfigurationAssignmentProperties{
				MaintenanceConfigurationID: to.StringPtr("<maintenance-configuration-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_DeleteParent.json
func ExampleConfigurationAssignmentsClient_DeleteParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.DeleteParent(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-parent-type>",
		"<resource-parent-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_Get.json
func ExampleConfigurationAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_CreateOrUpdate.json
func ExampleConfigurationAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		armmaintenance.ConfigurationAssignment{
			Properties: &armmaintenance.ConfigurationAssignmentProperties{
				MaintenanceConfigurationID: to.StringPtr("<maintenance-configuration-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_Delete.json
func ExampleConfigurationAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.Delete(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-type>",
		"<resource-name>",
		"<configuration-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_ListParent.json
func ExampleConfigurationAssignmentsClient_ListParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.ListParent(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-parent-type>",
		"<resource-parent-name>",
		"<resource-type>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/ConfigurationAssignments_List.json
func ExampleConfigurationAssignmentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		"<provider-name>",
		"<resource-type>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
