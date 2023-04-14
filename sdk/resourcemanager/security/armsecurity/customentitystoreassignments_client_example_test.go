//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentGet_example.json
func ExampleCustomEntityStoreAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomEntityStoreAssignmentsClient().Get(ctx, "TestResourceGroup", "33e7cc6e-a139-4723-a0e5-76993aee0771", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomEntityStoreAssignment = armsecurity.CustomEntityStoreAssignment{
	// 	Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
	// 	ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
	// 		EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwtILC4uzy9KCcjPyUyu5OWqUShJzE5VMAQAlMJzABgAAAA="),
	// 		Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentCreate_example.json
func ExampleCustomEntityStoreAssignmentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomEntityStoreAssignmentsClient().Create(ctx, "TestResourceGroup", "33e7cc6e-a139-4723-a0e5-76993aee0771", armsecurity.CustomEntityStoreAssignmentRequest{
		Properties: &armsecurity.CustomEntityStoreAssignmentRequestProperties{
			Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomEntityStoreAssignment = armsecurity.CustomEntityStoreAssignment{
	// 	Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
	// 	ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
	// 		EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwtILC4uzy9KCcjPyUyu5OWqUShJzE5VMAQAlMJzABgAAAA="),
	// 		Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentDelete_example.json
func ExampleCustomEntityStoreAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCustomEntityStoreAssignmentsClient().Delete(ctx, "TestResourceGroup", "33e7cc6e-a139-4723-a0e5-76993aee0771", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentListByResourceGroup_example.json
func ExampleCustomEntityStoreAssignmentsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomEntityStoreAssignmentsClient().NewListByResourceGroupPager("TestResourceGroup", nil)
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
		// page.CustomEntityStoreAssignmentsListResult = armsecurity.CustomEntityStoreAssignmentsListResult{
		// 	Value: []*armsecurity.CustomEntityStoreAssignment{
		// 		{
		// 			Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
		// 			Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/33e7cc6e-a139-4723-a0e5-76993aee0771"),
		// 			Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
		// 				EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwtILC4uzy9KCcjPyUyu5OWqUShJzE5VMAQAlMJzABgAAAA="),
		// 				Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a400695c-4728-d5cc-8e19-4b5a76f209df"),
		// 			Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/a400695c-4728-d5cc-8e19-4b5a76f209df"),
		// 			Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
		// 				EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwvIz8lMrgzKz0nlqlFIrShJzUtR8Cz2SE3MKcmoVLBVUE9LzClOVQcA1IFnficAAAA="),
		// 				Principal: to.Ptr("aaduser=f6e2564c-f34a-9b61-416c-5e5e7e521118;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/24166cedd1055380f1b9d40df270bf51b287d7d9/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentListBySubscription_example.json
func ExampleCustomEntityStoreAssignmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomEntityStoreAssignmentsClient().NewListBySubscriptionPager(nil)
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
		// page.CustomEntityStoreAssignmentsListResult = armsecurity.CustomEntityStoreAssignmentsListResult{
		// 	Value: []*armsecurity.CustomEntityStoreAssignment{
		// 		{
		// 			Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
		// 			Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/33e7cc6e-a139-4723-a0e5-76993aee0771"),
		// 			Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
		// 				EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwtILC4uzy9KCcjPyUyu5OWqUShJzE5VMAQAlMJzABgAAAA="),
		// 				Principal: to.Ptr("aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a400695c-4728-d5cc-8e19-4b5a76f209df"),
		// 			Type: to.Ptr("Microsoft.Security/customEntityStoreAssignments"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customEntityStoreAssignments/a400695c-4728-d5cc-8e19-4b5a76f209df"),
		// 			Properties: &armsecurity.CustomEntityStoreAssignmentProperties{
		// 				EntityStoreDatabaseLink: to.Ptr("https://dataexplorer.azure.com/clusters/securitydatastore.centralus/databases/DiscoveryAwsKedamari?query=H4sIAAAAAAAAAwvIz8lMrgzKz0nlqlFIrShJzUtR8Cz2SE3MKcmoVLBVUE9LzClOVQcA1IFnficAAAA="),
		// 				Principal: to.Ptr("aaduser=f6e2564c-f34a-9b61-416c-5e5e7e521118;72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
