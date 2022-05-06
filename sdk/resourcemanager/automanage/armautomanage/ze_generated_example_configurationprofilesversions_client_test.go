//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/createOrUpdateConfigurationProfileVersion.json
func ExampleConfigurationProfilesVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<configuration-profile-name>",
		"<version-name>",
		"<resource-group-name>",
		armautomanage.ConfigurationProfile{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"Organization": to.Ptr("Administration"),
			},
			Properties: &armautomanage.ConfigurationProfileProperties{
				Configuration: map[string]interface{}{
					"Antimalware/Enable":                false,
					"AzureSecurityCenter/Enable":        true,
					"Backup/Enable":                     false,
					"BootDiagnostics/Enable":            true,
					"ChangeTrackingAndInventory/Enable": true,
					"GuestConfiguration/Enable":         true,
					"LogAnalytics/Enable":               true,
					"UpdateManagement/Enable":           true,
					"VMInsights/Enable":                 true,
				},
				Overrides: []interface{}{
					map[string]interface{}{
						"if": map[string]interface{}{
							"equals": "eastus",
							"field":  "$.location",
						},
						"priority": float64(100),
						"then": map[string]interface{}{
							"LogAnalytics/Enable":      true,
							"LogAnalytics/Reprovision": true,
							"LogAnalytics/Workspace":   "/subscriptions/abc/resourceGroups/xyz/providers/Microsoft.La/Workspaces/eastus",
						},
					},
					map[string]interface{}{
						"if": map[string]interface{}{
							"equals": "centralcanada",
							"field":  "$.location",
						},
						"priority": float64(200),
						"then": map[string]interface{}{
							"LogAnalytics/Enable":      true,
							"LogAnalytics/Reprovision": true,
							"LogAnalytics/Workspace":   "/subscriptions/abc/resourceGroups/xyz/providers/Microsoft.La/Workspaces/centralcanada",
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/getConfigurationProfileVersion.json
func ExampleConfigurationProfilesVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<configuration-profile-name>",
		"<version-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/deleteConfigurationProfileVersion.json
func ExampleConfigurationProfilesVersionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<configuration-profile-name>",
		"<version-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/updateConfigurationProfileVersion.json
func ExampleConfigurationProfilesVersionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"<configuration-profile-name>",
		"<version-name>",
		"<resource-group-name>",
		armautomanage.ConfigurationProfileUpdate{
			Tags: map[string]*string{
				"Organization": to.Ptr("Administration"),
			},
			Properties: &armautomanage.ConfigurationProfileProperties{
				Configuration: map[string]interface{}{
					"Antimalware/Enable":                false,
					"AzureSecurityCenter/Enable":        true,
					"Backup/Enable":                     false,
					"BootDiagnostics/Enable":            true,
					"ChangeTrackingAndInventory/Enable": true,
					"GuestConfiguration/Enable":         true,
					"LogAnalytics/Enable":               true,
					"UpdateManagement/Enable":           true,
					"VMInsights/Enable":                 true,
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automanage/resource-manager/Microsoft.Automanage/preview/2021-04-30-preview/examples/listConfigurationProfileVersions.json
func ExampleConfigurationProfilesVersionsClient_NewListChildResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomanage.NewConfigurationProfilesVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListChildResourcesPager("<configuration-profile-name>",
		"<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
