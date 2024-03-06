//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Get.json
func ExampleDataSetMappingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Get(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", "DatasetMapping1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientGetResponse{
	// 	                            DataSetMappingClassification: &armdatashare.BlobDataSetMapping{
	// 		Name: to.Ptr("DatasetMapping1"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindBlob),
	// 		Properties: &armdatashare.BlobMappingProperties{
	// 			ContainerName: to.Ptr("C1"),
	// 			DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
	// 			FilePath: to.Ptr("file21"),
	// 			ResourceGroup: to.Ptr("SampleResourceGroup"),
	// 			StorageAccountName: to.Ptr("storage2"),
	// 			SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Create.json
func ExampleDataSetMappingsClient_Create_dataSetMappingsCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Create(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", "DatasetMapping1", &armdatashare.BlobDataSetMapping{
		Kind: to.Ptr(armdatashare.DataSetMappingKindBlob),
		Properties: &armdatashare.BlobMappingProperties{
			ContainerName:      to.Ptr("C1"),
			DataSetID:          to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
			FilePath:           to.Ptr("file21"),
			ResourceGroup:      to.Ptr("SampleResourceGroup"),
			StorageAccountName: to.Ptr("storage2"),
			SubscriptionID:     to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientCreateResponse{
	// 	                            DataSetMappingClassification: &armdatashare.BlobDataSetMapping{
	// 		Name: to.Ptr("DatasetMapping1"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindBlob),
	// 		Properties: &armdatashare.BlobMappingProperties{
	// 			ContainerName: to.Ptr("C1"),
	// 			DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
	// 			FilePath: to.Ptr("file21"),
	// 			ResourceGroup: to.Ptr("SampleResourceGroup"),
	// 			StorageAccountName: to.Ptr("storage2"),
	// 			SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SqlDB_Create.json
func ExampleDataSetMappingsClient_Create_dataSetMappingsSqlDbCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Create(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", "DatasetMapping1", &armdatashare.SQLDBTableDataSetMapping{
		Kind: to.Ptr(armdatashare.DataSetMappingKindSQLDBTable),
		Properties: &armdatashare.SQLDBTableDataSetMappingProperties{
			DataSetID:           to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
			DatabaseName:        to.Ptr("Database1"),
			SchemaName:          to.Ptr("dbo"),
			SQLServerResourceID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"),
			TableName:           to.Ptr("Table1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientCreateResponse{
	// 	                            DataSetMappingClassification: &armdatashare.SQLDBTableDataSetMapping{
	// 		Name: to.Ptr("DatasetMapping1"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindSQLDBTable),
	// 		Properties: &armdatashare.SQLDBTableDataSetMappingProperties{
	// 			DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
	// 			DatabaseName: to.Ptr("Database1"),
	// 			SchemaName: to.Ptr("dbo"),
	// 			SQLServerResourceID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"),
	// 			TableName: to.Ptr("Table1"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SqlDWDataSetToAdlsGen2File_Create.json
func ExampleDataSetMappingsClient_Create_dataSetMappingsSqlDwDataSetToAdlsGen2FileCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Create(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", "DatasetMapping1", &armdatashare.ADLSGen2FileDataSetMapping{
		Kind: to.Ptr(armdatashare.DataSetMappingKindAdlsGen2File),
		Properties: &armdatashare.ADLSGen2FileDataSetMappingProperties{
			DataSetID:          to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
			FilePath:           to.Ptr("file21"),
			FileSystem:         to.Ptr("fileSystem"),
			OutputType:         to.Ptr(armdatashare.OutputTypeCSV),
			ResourceGroup:      to.Ptr("SampleResourceGroup"),
			StorageAccountName: to.Ptr("storage2"),
			SubscriptionID:     to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientCreateResponse{
	// 	                            DataSetMappingClassification: &armdatashare.ADLSGen2FileDataSetMapping{
	// 		Name: to.Ptr("DatasetMapping1"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindAdlsGen2File),
	// 		Properties: &armdatashare.ADLSGen2FileDataSetMappingProperties{
	// 			DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
	// 			FilePath: to.Ptr("file21"),
	// 			FileSystem: to.Ptr("fileSystem"),
	// 			OutputType: to.Ptr(armdatashare.OutputTypeCSV),
	// 			ResourceGroup: to.Ptr("SampleResourceGroup"),
	// 			StorageAccountName: to.Ptr("storage2"),
	// 			SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SqlDW_Create.json
func ExampleDataSetMappingsClient_Create_dataSetMappingsSqlDwCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Create(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", "DatasetMapping1", &armdatashare.SQLDWTableDataSetMapping{
		Kind: to.Ptr(armdatashare.DataSetMappingKindSQLDWTable),
		Properties: &armdatashare.SQLDWTableDataSetMappingProperties{
			DataSetID:           to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
			DataWarehouseName:   to.Ptr("DataWarehouse1"),
			SchemaName:          to.Ptr("dbo"),
			SQLServerResourceID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"),
			TableName:           to.Ptr("Table1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientCreateResponse{
	// 	                            DataSetMappingClassification: &armdatashare.SQLDWTableDataSetMapping{
	// 		Name: to.Ptr("DatasetMapping1"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindSQLDWTable),
	// 		Properties: &armdatashare.SQLDWTableDataSetMappingProperties{
	// 			DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
	// 			DataWarehouseName: to.Ptr("DataWarehouse1"),
	// 			SchemaName: to.Ptr("dbo"),
	// 			SQLServerResourceID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"),
	// 			TableName: to.Ptr("Table1"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SynapseWorkspaceSqlPoolTable_Create.json
func ExampleDataSetMappingsClient_Create_dataSetMappingsSynapseWorkspaceSqlPoolTableCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Create(ctx, "SampleResourceGroup", "consumerAccount", "ShareSubscription1", "datasetMappingName1", &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMapping{
		Kind: to.Ptr(armdatashare.DataSetMappingKindSynapseWorkspaceSQLPoolTable),
		Properties: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMappingProperties{
			DataSetID:                              to.Ptr("3dc64e49-1fc3-4186-b3dc-d388c4d3076a"),
			SynapseWorkspaceSQLPoolTableResourceID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientCreateResponse{
	// 	                            DataSetMappingClassification: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMapping{
	// 		Name: to.Ptr("datasetMappingName"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/4e745bb7-c420-479b-b0d6-a0f92d48a227/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/consumerAccount/shareSubscriptions/ShareSubscription1/dataSetMappings/datasetMappingName1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindSynapseWorkspaceSQLPoolTable),
	// 		Properties: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMappingProperties{
	// 			DataSetID: to.Ptr("3dc64e49-1fc3-4186-b3dc-d388c4d3076a"),
	// 			DataSetMappingStatus: to.Ptr(armdatashare.DataSetMappingStatusOk),
	// 			ProvisioningState: to.Ptr(armdatashare.ProvisioningStateSucceeded),
	// 			SynapseWorkspaceSQLPoolTableResourceID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Delete.json
func ExampleDataSetMappingsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataSetMappingsClient().Delete(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", "DatasetMapping1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_ListByShareSubscription.json
func ExampleDataSetMappingsClient_NewListByShareSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataSetMappingsClient().NewListByShareSubscriptionPager("SampleResourceGroup", "Account1", "ShareSubscription1", &armdatashare.DataSetMappingsClientListByShareSubscriptionOptions{SkipToken: nil,
		Filter:  nil,
		Orderby: nil,
	})
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
		// page.DataSetMappingList = armdatashare.DataSetMappingList{
		// 	Value: []armdatashare.DataSetMappingClassification{
		// 		&armdatashare.BlobDataSetMapping{
		// 			Name: to.Ptr("DatasetMapping1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
		// 			Kind: to.Ptr(armdatashare.DataSetMappingKindBlob),
		// 			Properties: &armdatashare.BlobMappingProperties{
		// 				ContainerName: to.Ptr("C1"),
		// 				DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
		// 				FilePath: to.Ptr("file21"),
		// 				ResourceGroup: to.Ptr("SampleResourceGroup"),
		// 				StorageAccountName: to.Ptr("storage2"),
		// 				SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
		// 			},
		// 		},
		// 		&armdatashare.BlobDataSetMapping{
		// 			Name: to.Ptr("DatasetMapping1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/dataSetMappings/DatasetMapping1"),
		// 			Kind: to.Ptr(armdatashare.DataSetMappingKindBlob),
		// 			Properties: &armdatashare.BlobMappingProperties{
		// 				ContainerName: to.Ptr("C1"),
		// 				DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d227"),
		// 				FilePath: to.Ptr("file21"),
		// 				ResourceGroup: to.Ptr("SampleResourceGroup"),
		// 				StorageAccountName: to.Ptr("storage2"),
		// 				SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
		// 			},
		// 	}},
		// }
	}
}
