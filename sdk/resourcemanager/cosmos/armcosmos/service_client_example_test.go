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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBServicesList.json
func ExampleServiceClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", "ddb1", nil)
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
		// page.ServiceResourceListResult = armcosmos.ServiceResourceListResult{
		// 	Value: []*armcosmos.ServiceResource{
		// 		{
		// 			Name: to.Ptr("sqlDedicatedGateway"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/sqlDedicatedGateway"),
		// 			Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
		// 				InstanceCount: to.Ptr[int32](1),
		// 				InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
		// 				ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
		// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
		// 				Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
		// 					{
		// 						Name: to.Ptr("sqlDedicatedGateway-westus2"),
		// 						Location: to.Ptr("West US 2"),
		// 						Status: to.Ptr(armcosmos.ServiceStatusRunning),
		// 						SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
		// 				}},
		// 				SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBDataTransferServiceCreate.json
func ExampleServiceClient_BeginCreate_dataTransferServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rg1", "ddb1", "DataTransfer", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.ServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeDataTransfer),
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
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("DataTransfer"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/DataTransfer"),
	// 	Properties: &armcosmos.DataTransferServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeDataTransfer),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.DataTransferRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("DataTransfer-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBGraphAPIComputeServiceCreate.json
func ExampleServiceClient_BeginCreate_graphApiComputeServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rg1", "ddb1", "GraphAPICompute", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.ServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeGraphAPICompute),
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
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("GraphAPICompute"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/GraphAPICompute"),
	// 	Properties: &armcosmos.GraphAPIComputeServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeGraphAPICompute),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute.gremlin.cosmos.windows-int.net/"),
	// 		Locations: []*armcosmos.GraphAPIComputeRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("GraphAPICompute-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute-westus.gremlin.cosmos.windows-int.net/"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBMaterializedViewsBuilderServiceCreate.json
func ExampleServiceClient_BeginCreate_materializedViewsBuilderServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rg1", "ddb1", "MaterializedViewsBuilder", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.ServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeMaterializedViewsBuilder),
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
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("MaterializedViewsBuilder"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/MaterializedViewsBuilder"),
	// 	Properties: &armcosmos.MaterializedViewsBuilderServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeMaterializedViewsBuilder),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.MaterializedViewsBuilderRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("MaterializedViewsBuilder-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBSqlDedicatedGatewayServiceCreate.json
func ExampleServiceClient_BeginCreate_sqlDedicatedGatewayServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rg1", "ddb1", "SqlDedicatedGateway", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.ServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
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
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("SqlDedicatedGateway"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/SqlDedicatedGateway"),
	// 	Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("SqlDedicatedGateway-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
	// 		}},
	// 		SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBDataTransferServiceGet.json
func ExampleServiceClient_Get_dataTransferServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "ddb1", "DataTransfer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("DataTransfer"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/DataTransfer"),
	// 	Properties: &armcosmos.DataTransferServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeDataTransfer),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.DataTransferRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("DataTransfer-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBGraphAPIComputeServiceGet.json
func ExampleServiceClient_Get_graphApiComputeServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "ddb1", "GraphAPICompute", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("GraphAPICompute"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/GraphAPICompute"),
	// 	Properties: &armcosmos.GraphAPIComputeServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeGraphAPICompute),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute.gremlin.cosmos.windows-int.net/"),
	// 		Locations: []*armcosmos.GraphAPIComputeRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("GraphAPICompute-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute-westus.gremlin.cosmos.windows-int.net/"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBMaterializedViewsBuilderServiceGet.json
func ExampleServiceClient_Get_materializedViewsBuilderServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "ddb1", "MaterializedViewsBuilder", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("MaterializedViewsBuilder"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/MaterializedViewsBuilder"),
	// 	Properties: &armcosmos.MaterializedViewsBuilderServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeMaterializedViewsBuilder),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.MaterializedViewsBuilderRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("MaterializedViewsBuilder-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBSqlDedicatedGatewayServiceGet.json
func ExampleServiceClient_Get_sqlDedicatedGatewayServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "ddb1", "SqlDedicatedGateway", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("SqlDedicatedGateway"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/SqlDedicatedGateway"),
	// 	Properties: &armcosmos.SQLDedicatedGatewayServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		Locations: []*armcosmos.SQLDedicatedGatewayRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("SqlDedicatedGateway-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway-westus.sqlx.cosmos.windows-int.net/"),
	// 		}},
	// 		SQLDedicatedGatewayEndpoint: to.Ptr("https://sqlDedicatedGateway.sqlx.cosmos.windows-int.net/"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBDataTransferServiceDelete.json
func ExampleServiceClient_BeginDelete_dataTransferServiceDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "ddb1", "DataTransfer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBGraphAPIComputeServiceDelete.json
func ExampleServiceClient_BeginDelete_graphApiComputeServiceDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "ddb1", "GraphAPICompute", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBMaterializedViewsBuilderServiceDelete.json
func ExampleServiceClient_BeginDelete_materializedViewsBuilderServiceDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "ddb1", "MaterializedViewsBuilder", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBSqlDedicatedGatewayServiceDelete.json
func ExampleServiceClient_BeginDelete_sqlDedicatedGatewayServiceDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "ddb1", "SqlDedicatedGateway", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
