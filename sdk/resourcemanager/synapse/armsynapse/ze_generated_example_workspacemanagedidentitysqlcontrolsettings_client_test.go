//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetManagedIdentitySqlControlSettings.json
func ExampleWorkspaceManagedIdentitySQLControlSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewWorkspaceManagedIdentitySQLControlSettingsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ManagedIdentitySQLControlSettingsModel.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateManagedIdentitySqlControlSettings.json
func ExampleWorkspaceManagedIdentitySQLControlSettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewWorkspaceManagedIdentitySQLControlSettingsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.ManagedIdentitySQLControlSettingsModel{
			Properties: &armsynapse.ManagedIdentitySQLControlSettingsModelProperties{
				GrantSQLControlToManagedIdentity: &armsynapse.ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity{
					DesiredState: armsynapse.ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredStateEnabled.ToPtr(),
				},
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
	log.Printf("ManagedIdentitySQLControlSettingsModel.ID: %s\n", *res.ID)
}
