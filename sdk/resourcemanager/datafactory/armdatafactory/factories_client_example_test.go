//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_List.json
func ExampleFactoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFactoriesClient().NewListPager(nil)
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
		// page.FactoryListResponse = armdatafactory.FactoryListResponse{
		// 	Value: []*armdatafactory.Factory{
		// 		{
		// 			Name: to.Ptr("rpV2OrigDF-72c7d3d4-5e17-4ec6-91de-9ab433f15e79"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"0000aa0d-0000-0000-0000-5b0d58170000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/rg-yanzhang-dfv2/providers/Microsoft.DataFactory/factories/rpv2origdf-72c7d3d4-5e17-4ec6-91de-9ab433f15e79"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("399c3de2-6072-4326-bfa9-4d0c116f1a7b"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-29T13:39:35.615Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("df-dogfood-yanzhang-we"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"0000f301-0000-0000-0000-5b21b16c0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/rg-yanzhang-dfv2/providers/Microsoft.DataFactory/factories/df-dogfood-yanzhang-we"),
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("e8dd6df9-bad5-4dea-8fb8-0d13d1845d9e"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T00:06:04.666Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampleFactoryName-linked"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00008a02-0000-0000-0000-5b237f270000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName-linked"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("10743799-44d2-42fe-8c4d-5bc5c51c0684"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T08:56:07.182Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FactoryToUpgrade"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00003d04-0000-0000-0000-5b28962f0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/factorytoupgrade"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:35:35.713Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2018-06-01"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampleFactoryName"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00004004-0000-0000-0000-5b28979e0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"exampleTag": to.Ptr("exampleValue"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PurviewConfiguration: &armdatafactory.PurviewConfiguration{
		// 					PurviewResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Purview/accounts/examplePurview"),
		// 				},
		// 				RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
		// 					Type: to.Ptr("FactoryVSTSConfiguration"),
		// 					AccountName: to.Ptr("ADF"),
		// 					CollaborationBranch: to.Ptr("master"),
		// 					LastCommitID: to.Ptr(""),
		// 					RepositoryName: to.Ptr("repo"),
		// 					RootFolder: to.Ptr("/"),
		// 					ProjectName: to.Ptr("project"),
		// 					TenantID: to.Ptr(""),
		// 				},
		// 				Version: to.Ptr("2018-06-01"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rpV2OrigDF-72c7d3d4-5e17-4ec6-91de-9ab433f15e79"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"0000aa0d-0000-0000-0000-5b0d58170000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/rg-yanzhang-dfv2/providers/Microsoft.DataFactory/factories/rpv2origdf-72c7d3d4-5e17-4ec6-91de-9ab433f15e79"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("399c3de2-6072-4326-bfa9-4d0c116f1a7b"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-29T13:39:35.615Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("df-dogfood-yanzhang-we"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"0000f301-0000-0000-0000-5b21b16c0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/rg-yanzhang-dfv2/providers/Microsoft.DataFactory/factories/df-dogfood-yanzhang-we"),
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("e8dd6df9-bad5-4dea-8fb8-0d13d1845d9e"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T00:06:04.666Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampleFactoryName-linked"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00008a02-0000-0000-0000-5b237f270000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName-linked"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("10743799-44d2-42fe-8c4d-5bc5c51c0684"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T08:56:07.182Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FactoryToUpgrade"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00003d04-0000-0000-0000-5b28962f0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/factorytoupgrade"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:35:35.713Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2018-06-01"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampleFactoryName"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00004004-0000-0000-0000-5b28979e0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"exampleTag": to.Ptr("exampleValue"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PurviewConfiguration: &armdatafactory.PurviewConfiguration{
		// 					PurviewResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Purview/accounts/examplePurview"),
		// 				},
		// 				RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
		// 					Type: to.Ptr("FactoryVSTSConfiguration"),
		// 					AccountName: to.Ptr("ADF"),
		// 					CollaborationBranch: to.Ptr("master"),
		// 					LastCommitID: to.Ptr(""),
		// 					RepositoryName: to.Ptr("repo"),
		// 					RootFolder: to.Ptr("/"),
		// 					ProjectName: to.Ptr("project"),
		// 					TenantID: to.Ptr(""),
		// 				},
		// 				Version: to.Ptr("2018-06-01"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_ConfigureFactoryRepo.json
func ExampleFactoriesClient_ConfigureFactoryRepo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().ConfigureFactoryRepo(ctx, "East US", armdatafactory.FactoryRepoUpdate{
		FactoryResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
			Type:                to.Ptr("FactoryVSTSConfiguration"),
			AccountName:         to.Ptr("ADF"),
			CollaborationBranch: to.Ptr("master"),
			LastCommitID:        to.Ptr(""),
			RepositoryName:      to.Ptr("repo"),
			RootFolder:          to.Ptr("/"),
			ProjectName:         to.Ptr("project"),
			TenantID:            to.Ptr(""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Factory = armdatafactory.Factory{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories"),
	// 	ETag: to.Ptr("\"00004004-0000-0000-0000-5b28979e0000\""),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"exampleTag": to.Ptr("exampleValue"),
	// 	},
	// 	Properties: &armdatafactory.FactoryProperties{
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
	// 			Type: to.Ptr("FactoryVSTSConfiguration"),
	// 			AccountName: to.Ptr("ADF"),
	// 			CollaborationBranch: to.Ptr("master"),
	// 			LastCommitID: to.Ptr(""),
	// 			RepositoryName: to.Ptr("repo"),
	// 			RootFolder: to.Ptr("/"),
	// 			ProjectName: to.Ptr("project"),
	// 			TenantID: to.Ptr(""),
	// 		},
	// 		Version: to.Ptr("2018-06-01"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_ListByResourceGroup.json
func ExampleFactoriesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFactoriesClient().NewListByResourceGroupPager("exampleResourceGroup", nil)
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
		// page.FactoryListResponse = armdatafactory.FactoryListResponse{
		// 	Value: []*armdatafactory.Factory{
		// 		{
		// 			Name: to.Ptr("exampleFactoryName-linked"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00008a02-0000-0000-0000-5b237f270000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName-linked"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatafactory.FactoryIdentity{
		// 				Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("10743799-44d2-42fe-8c4d-5bc5c51c0684"),
		// 				TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T08:56:07.182Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2017-09-01-preview"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FactoryToUpgrade"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00003d04-0000-0000-0000-5b28962f0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/factorytoupgrade"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:35:35.713Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Version: to.Ptr("2018-06-01"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampleFactoryName"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 			ETag: to.Ptr("\"00004004-0000-0000-0000-5b28979e0000\""),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"exampleTag": to.Ptr("exampleValue"),
		// 			},
		// 			Properties: &armdatafactory.FactoryProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
		// 					Type: to.Ptr("FactoryVSTSConfiguration"),
		// 					AccountName: to.Ptr("ADF"),
		// 					CollaborationBranch: to.Ptr("master"),
		// 					LastCommitID: to.Ptr(""),
		// 					RepositoryName: to.Ptr("repo"),
		// 					RootFolder: to.Ptr("/"),
		// 					ProjectName: to.Ptr("project"),
		// 					TenantID: to.Ptr(""),
		// 				},
		// 				Version: to.Ptr("2018-06-01"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_CreateOrUpdate.json
func ExampleFactoriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.Factory{
		Location: to.Ptr("East US"),
	}, &armdatafactory.FactoriesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Factory = armdatafactory.Factory{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories"),
	// 	ETag: to.Ptr("\"00003e04-0000-0000-0000-5b28979e0000\""),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armdatafactory.FactoryProperties{
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Version: to.Ptr("2018-06-01"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_Update.json
func ExampleFactoriesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().Update(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.FactoryUpdateParameters{
		Tags: map[string]*string{
			"exampleTag": to.Ptr("exampleValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Factory = armdatafactory.Factory{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories"),
	// 	ETag: to.Ptr("\"00003f04-0000-0000-0000-5b28979e0000\""),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"exampleTag": to.Ptr("exampleValue"),
	// 	},
	// 	Properties: &armdatafactory.FactoryProperties{
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Version: to.Ptr("2018-06-01"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_Get.json
func ExampleFactoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", &armdatafactory.FactoriesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Factory = armdatafactory.Factory{
	// 	Name: to.Ptr("exampleFactoryName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories"),
	// 	ETag: to.Ptr("\"00004004-0000-0000-0000-5b28979e0000\""),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"exampleTag": to.Ptr("exampleValue"),
	// 	},
	// 	Properties: &armdatafactory.FactoryProperties{
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.004Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PurviewConfiguration: &armdatafactory.PurviewConfiguration{
	// 			PurviewResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Purview/accounts/examplePurview"),
	// 		},
	// 		RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
	// 			Type: to.Ptr("FactoryVSTSConfiguration"),
	// 			AccountName: to.Ptr("ADF"),
	// 			CollaborationBranch: to.Ptr("master"),
	// 			LastCommitID: to.Ptr(""),
	// 			RepositoryName: to.Ptr("repo"),
	// 			RootFolder: to.Ptr("/"),
	// 			ProjectName: to.Ptr("project"),
	// 			TenantID: to.Ptr(""),
	// 		},
	// 		Version: to.Ptr("2018-06-01"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_Delete.json
func ExampleFactoriesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFactoriesClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetGitHubAccessToken.json
func ExampleFactoriesClient_GetGitHubAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().GetGitHubAccessToken(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.GitHubAccessTokenRequest{
		GitHubAccessCode:         to.Ptr("some"),
		GitHubAccessTokenBaseURL: to.Ptr("some"),
		GitHubClientID:           to.Ptr("some"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubAccessTokenResponse = armdatafactory.GitHubAccessTokenResponse{
	// 	GitHubAccessToken: to.Ptr("myAccessTokenExample"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetDataPlaneAccess.json
func ExampleFactoriesClient_GetDataPlaneAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().GetDataPlaneAccess(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.UserAccessPolicy{
		AccessResourcePath: to.Ptr(""),
		ExpireTime:         to.Ptr("2018-11-10T09:46:20.2659347Z"),
		Permissions:        to.Ptr("r"),
		ProfileName:        to.Ptr("DefaultProfile"),
		StartTime:          to.Ptr("2018-11-10T02:46:20.2659347Z"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessPolicyResponse = armdatafactory.AccessPolicyResponse{
	// 	AccessToken: to.Ptr("**********"),
	// 	DataPlaneURL: to.Ptr("https://rpeastus.svc.datafactory.azure.com:4433"),
	// 	Policy: &armdatafactory.UserAccessPolicy{
	// 		AccessResourcePath: to.Ptr(""),
	// 		ExpireTime: to.Ptr("2018-11-10T09:46:20.2659347Z"),
	// 		Permissions: to.Ptr("r"),
	// 		ProfileName: to.Ptr("DefaultProfile"),
	// 		StartTime: to.Ptr("2018-11-10T02:46:20.2659347Z"),
	// 	},
	// }
}
