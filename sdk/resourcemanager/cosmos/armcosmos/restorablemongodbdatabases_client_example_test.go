//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBRestorableMongodbDatabaseList.json
func ExampleRestorableMongodbDatabasesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableMongodbDatabasesClient().NewListPager("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d", nil)
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
		// page.RestorableMongodbDatabasesListResult = armcosmos.RestorableMongodbDatabasesListResult{
		// 	Value: []*armcosmos.RestorableMongodbDatabaseGetResult{
		// 		{
		// 			Name: to.Ptr("59c21367-b98b-4a8e-abb7-b6f46600decc"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorableMongodbDatabases"),
		// 			ID: to.Ptr("/subscriptions/2296c272-5d55-40d9-bc05-4d56dc2d7588/providers/Microsoft.DocumentDb/locations/westus/restorableDatabaseAccounts/36f09704-6be3-4f33-aa05-17b73e504c75/restorableMongodbDatabases/59c21367-b98b-4a8e-abb7-b6f46600decc"),
		// 			Properties: &armcosmos.RestorableMongodbDatabaseProperties{
		// 				Resource: &armcosmos.RestorableMongodbDatabasePropertiesResource{
		// 					Rid: to.Ptr("DLB14gAAAA=="),
		// 					CanUndelete: to.Ptr("invalid"),
		// 					EventTimestamp: to.Ptr("2020-09-02T19:45:03Z"),
		// 					OperationType: to.Ptr(armcosmos.OperationTypeCreate),
		// 					OwnerID: to.Ptr("Database1"),
		// 					OwnerResourceID: to.Ptr("PD5DALigDgw="),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("8456cb17-cdb0-4c6a-8db8-d0ff3f886257"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorableMongodbDatabases"),
		// 			ID: to.Ptr("/subscriptions/2296c272-5d55-40d9-bc05-4d56dc2d7588/providers/Microsoft.DocumentDb/locations/westus/restorableDatabaseAccounts/d9b26648-2f53-4541-b3d8-3044f4f9810d/restorableMongodbDatabases/8456cb17-cdb0-4c6a-8db8-d0ff3f886257"),
		// 			Properties: &armcosmos.RestorableMongodbDatabaseProperties{
		// 				Resource: &armcosmos.RestorableMongodbDatabasePropertiesResource{
		// 					Rid: to.Ptr("ESXNLAAAAA=="),
		// 					CanUndelete: to.Ptr("notRestorable"),
		// 					CanUndeleteReason: to.Ptr("Database already exists. Only deleted resources can be restored within same account."),
		// 					EventTimestamp: to.Ptr("2020-09-02T19:53:42Z"),
		// 					OperationType: to.Ptr(armcosmos.OperationTypeDelete),
		// 					OwnerID: to.Ptr("Database1"),
		// 					OwnerResourceID: to.Ptr("PD5DALigDgw="),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
