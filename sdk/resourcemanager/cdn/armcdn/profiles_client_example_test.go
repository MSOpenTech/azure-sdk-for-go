//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_List.json
func ExampleProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProfileListResult = armcdn.ProfileListResult{
		// 	Value: []*armcdn.Profile{
		// 		{
		// 			Name: to.Ptr("profile1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG1/providers/Microsoft.Cdn/profiles/profile1"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("frontdoor"),
		// 			Properties: &armcdn.ProfileProperties{
		// 				FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
		// 				LogScrubbing: &armcdn.ProfileLogScrubbing{
		// 					ScrubbingRules: []*armcdn.ProfileScrubbingRules{
		// 					},
		// 					State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
		// 				},
		// 				OriginResponseTimeoutSeconds: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("profile2"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG1/providers/Microsoft.Cdn/profiles/profile2"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("frontdoor"),
		// 			Properties: &armcdn.ProfileProperties{
		// 				FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
		// 				LogScrubbing: &armcdn.ProfileLogScrubbing{
		// 					ScrubbingRules: []*armcdn.ProfileScrubbingRules{
		// 					},
		// 					State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
		// 				},
		// 				OriginResponseTimeoutSeconds: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_ListByResourceGroup.json
func ExampleProfilesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListByResourceGroupPager("RG", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProfileListResult = armcdn.ProfileListResult{
		// 	Value: []*armcdn.Profile{
		// 		{
		// 			Name: to.Ptr("profile1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("frontdoor"),
		// 			Properties: &armcdn.ProfileProperties{
		// 				FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
		// 				LogScrubbing: &armcdn.ProfileLogScrubbing{
		// 					ScrubbingRules: []*armcdn.ProfileScrubbingRules{
		// 					},
		// 					State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
		// 				},
		// 				OriginResponseTimeoutSeconds: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("profile2"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile2"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr("frontdoor"),
		// 			Properties: &armcdn.ProfileProperties{
		// 				FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
		// 				LogScrubbing: &armcdn.ProfileLogScrubbing{
		// 					ScrubbingRules: []*armcdn.ProfileScrubbingRules{
		// 					},
		// 					State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
		// 				},
		// 				OriginResponseTimeoutSeconds: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
		// 				ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
		// 			},
		// 			SKU: &armcdn.SKU{
		// 				Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_Get.json
func ExampleProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().Get(ctx, "RG", "profile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armcdn.Profile{
	// 	Name: to.Ptr("profile1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Kind: to.Ptr("frontdoor"),
	// 	Properties: &armcdn.ProfileProperties{
	// 		FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
	// 		LogScrubbing: &armcdn.ProfileLogScrubbing{
	// 			ScrubbingRules: []*armcdn.ProfileScrubbingRules{
	// 			},
	// 			State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
	// 		},
	// 		OriginResponseTimeoutSeconds: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
	// 	},
	// 	SKU: &armcdn.SKU{
	// 		Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_Create.json
func ExampleProfilesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginCreate(ctx, "RG", "profile1", armcdn.Profile{
		Location: to.Ptr("global"),
		SKU: &armcdn.SKU{
			Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armcdn.Profile{
	// 	Name: to.Ptr("profile1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Kind: to.Ptr("frontdoor"),
	// 	Properties: &armcdn.ProfileProperties{
	// 		FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
	// 		OriginResponseTimeoutSeconds: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
	// 	},
	// 	SKU: &armcdn.SKU{
	// 		Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_Update.json
func ExampleProfilesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginUpdate(ctx, "RG", "profile1", armcdn.ProfileUpdateParameters{
		Tags: map[string]*string{
			"additionalProperties": to.Ptr("Tag1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armcdn.Profile{
	// 	Name: to.Ptr("profile1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 		"additionalProperties": to.Ptr("Tag1"),
	// 	},
	// 	Kind: to.Ptr("frontdoor"),
	// 	Properties: &armcdn.ProfileProperties{
	// 		FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
	// 		LogScrubbing: &armcdn.ProfileLogScrubbing{
	// 			ScrubbingRules: []*armcdn.ProfileScrubbingRules{
	// 			},
	// 			State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
	// 		},
	// 		OriginResponseTimeoutSeconds: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
	// 		ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
	// 	},
	// 	SKU: &armcdn.SKU{
	// 		Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_Delete.json
func ExampleProfilesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginDelete(ctx, "RG", "profile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_CanMigrate.json
func ExampleProfilesClient_BeginCanMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginCanMigrate(ctx, "RG", armcdn.CanMigrateParameters{
		ClassicResourceReference: &armcdn.ResourceReference{
			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoors/frontdoorname"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CanMigrateResult = armcdn.CanMigrateResult{
	// 	Type: to.Ptr("Microsoft.Cdn/canmigrate"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/yaoshitest/providers/Microsoft.Cdn/operationresults/operationid/profileresults/DummyProfile/canmigrateresults/DummyProfile"),
	// 	Properties: &armcdn.CanMigrateProperties{
	// 		CanMigrate: to.Ptr(true),
	// 		DefaultSKU: to.Ptr(armcdn.CanMigrateDefaultSKUStandardAzureFrontDoor),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_Migrate.json
func ExampleProfilesClient_BeginMigrate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginMigrate(ctx, "RG", armcdn.MigrationParameters{
		ClassicResourceReference: &armcdn.ResourceReference{
			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoors/frontdoorname"),
		},
		ProfileName: to.Ptr("profile1"),
		SKU: &armcdn.SKU{
			Name: to.Ptr(armcdn.SKUNameStandardAzureFrontDoor),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrateResult = armcdn.MigrateResult{
	// 	Type: to.Ptr("Microsoft.Cdn/migrate"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/operationresults/operationid/profileresults/profile1/migrateresults/profile1"),
	// 	Properties: &armcdn.MigrateResultProperties{
	// 		MigratedProfileResourceID: &armcdn.ResourceReference{
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_MigrationCommit.json
func ExampleProfilesClient_BeginMigrationCommit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProfilesClient().BeginMigrationCommit(ctx, "RG", "profile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_GenerateSsoUri.json
func ExampleProfilesClient_GenerateSsoURI() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().GenerateSsoURI(ctx, "RG", "profile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SsoURI = armcdn.SsoURI{
	// 	SsoURIValue: to.Ptr("https://someuri.com"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_ListSupportedOptimizationTypes.json
func ExampleProfilesClient_ListSupportedOptimizationTypes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().ListSupportedOptimizationTypes(ctx, "RG", "profile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SupportedOptimizationTypesListResult = armcdn.SupportedOptimizationTypesListResult{
	// 	SupportedOptimizationTypes: []*armcdn.OptimizationType{
	// 		to.Ptr(armcdn.OptimizationTypeGeneralWebDelivery),
	// 		to.Ptr(armcdn.OptimizationTypeDynamicSiteAcceleration)},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Profiles_ListResourceUsage.json
func ExampleProfilesClient_NewListResourceUsagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListResourceUsagePager("RG", "profile1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ResourceUsageListResult = armcdn.ResourceUsageListResult{
		// 	Value: []*armcdn.ResourceUsage{
		// 		{
		// 			CurrentValue: to.Ptr[int32](0),
		// 			Limit: to.Ptr[int32](25),
		// 			ResourceType: to.Ptr("endpoint"),
		// 			Unit: to.Ptr(armcdn.ResourceUsageUnitCount),
		// 	}},
		// }
	}
}
