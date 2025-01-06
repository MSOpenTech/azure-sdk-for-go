//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListModelCapacities.json
func ExampleModelCapacitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewModelCapacitiesClient().NewListPager("OpenAI", "ada", "1", nil)
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
		// page.ModelCapacityListResult = armcognitiveservices.ModelCapacityListResult{
		// 	Value: []*armcognitiveservices.ModelCapacityListResultValueItem{
		// 		{
		// 			Name: to.Ptr("Standard"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/locations/models/skuCapacities"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionContext.SubscriptionId}/providers/Microsoft.CognitiveServices/locations/WestUS/models/OpenAI.ada.1/skuCapacities/Standard"),
		// 			Location: to.Ptr("WestUS"),
		// 			Properties: &armcognitiveservices.ModelSKUCapacityProperties{
		// 				AvailableCapacity: to.Ptr[float32](300),
		// 				AvailableFinetuneCapacity: to.Ptr[float32](20),
		// 				Model: &armcognitiveservices.DeploymentModel{
		// 					Name: to.Ptr("ada"),
		// 					Format: to.Ptr("OpenAI"),
		// 					Version: to.Ptr("1"),
		// 				},
		// 				SKUName: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
