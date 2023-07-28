//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ServerConfigurationOptionList.json
func ExampleServerConfigurationOptionsClient_NewListByManagedInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerConfigurationOptionsClient().NewListByManagedInstancePager("testrg", "testinstance", nil)
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
		// page.ServerConfigurationOptionListResult = armsql.ServerConfigurationOptionListResult{
		// 	Value: []*armsql.ServerConfigurationOption{
		// 		{
		// 			Name: to.Ptr("allowPolybaseExport"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/serverConfigurationOptions"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/serverConfigurationOptions/allowPolybaseExport"),
		// 			Properties: &armsql.ServerConfigurationOptionProperties{
		// 				ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 				ServerConfigurationOptionValue: to.Ptr[int32](1),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ServerConfigurationOptionGet.json
func ExampleServerConfigurationOptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerConfigurationOptionsClient().Get(ctx, "testrg", "testinstance", armsql.ServerConfigurationOptionNameAllowPolybaseExport, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerConfigurationOption = armsql.ServerConfigurationOption{
	// 	Name: to.Ptr("allowPolybaseExport"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/serverConfigurationOptions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/serverConfigurationOptions/allowPolybaseExport"),
	// 	Properties: &armsql.ServerConfigurationOptionProperties{
	// 		ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 		ServerConfigurationOptionValue: to.Ptr[int32](1),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ServerConfigurationOptionUpdate.json
func ExampleServerConfigurationOptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerConfigurationOptionsClient().BeginCreateOrUpdate(ctx, "testrg", "testinstance", armsql.ServerConfigurationOptionNameAllowPolybaseExport, armsql.ServerConfigurationOption{
		Properties: &armsql.ServerConfigurationOptionProperties{
			ServerConfigurationOptionValue: to.Ptr[int32](1),
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
	// res.ServerConfigurationOption = armsql.ServerConfigurationOption{
	// 	Name: to.Ptr("allowPolybaseExport"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/serverConfigurationOptions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/serverConfigurationOptions/allowPolybaseExport"),
	// 	Properties: &armsql.ServerConfigurationOptionProperties{
	// 		ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 		ServerConfigurationOptionValue: to.Ptr[int32](1),
	// 	},
	// }
}
