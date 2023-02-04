//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookList.json
func ExampleWebhooksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myResourceGroup", "myRegistry", nil)
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
		// page.WebhookListResult = armcontainerregistry.WebhookListResult{
		// 	Value: []*armcontainerregistry.Webhook{
		// 		{
		// 			Name: to.Ptr("myWebhook"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armcontainerregistry.WebhookProperties{
		// 				Actions: []*armcontainerregistry.WebhookAction{
		// 					to.Ptr(armcontainerregistry.WebhookActionPush)},
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 					Scope: to.Ptr("myRepository"),
		// 					Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookGet.json
func ExampleWebhooksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myResourceGroup", "myRegistry", "myWebhook", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Webhook = armcontainerregistry.Webhook{
	// 	Name: to.Ptr("myWebhook"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.WebhookProperties{
	// 		Actions: []*armcontainerregistry.WebhookAction{
	// 			to.Ptr(armcontainerregistry.WebhookActionPush)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr("myRepository"),
	// 			Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookCreate.json
func ExampleWebhooksClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myResourceGroup", "myRegistry", "myWebhook", armcontainerregistry.WebhookCreateParameters{
		Location: to.Ptr("westus"),
		Properties: &armcontainerregistry.WebhookPropertiesCreateParameters{
			Actions: []*armcontainerregistry.WebhookAction{
				to.Ptr(armcontainerregistry.WebhookActionPush)},
			CustomHeaders: map[string]*string{
				"Authorization": to.Ptr("Basic 000000000000000000000000000000000000000000000000000"),
			},
			Scope:      to.Ptr("myRepository"),
			ServiceURI: to.Ptr("http://myservice.com"),
			Status:     to.Ptr(armcontainerregistry.WebhookStatusEnabled),
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.Webhook = armcontainerregistry.Webhook{
	// 	Name: to.Ptr("myWebhook"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.WebhookProperties{
	// 		Actions: []*armcontainerregistry.WebhookAction{
	// 			to.Ptr(armcontainerregistry.WebhookActionPush)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr("myRepository"),
	// 			Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookDelete.json
func ExampleWebhooksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "myResourceGroup", "myRegistry", "myWebhook", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookUpdate.json
func ExampleWebhooksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myWebhook", armcontainerregistry.WebhookUpdateParameters{
		Properties: &armcontainerregistry.WebhookPropertiesUpdateParameters{
			Actions: []*armcontainerregistry.WebhookAction{
				to.Ptr(armcontainerregistry.WebhookActionPush)},
			CustomHeaders: map[string]*string{
				"Authorization": to.Ptr("Basic 000000000000000000000000000000000000000000000000000"),
			},
			Scope:      to.Ptr("myRepository"),
			ServiceURI: to.Ptr("http://myservice.com"),
			Status:     to.Ptr(armcontainerregistry.WebhookStatusEnabled),
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.Webhook = armcontainerregistry.Webhook{
	// 	Name: to.Ptr("myWebhook"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.WebhookProperties{
	// 		Actions: []*armcontainerregistry.WebhookAction{
	// 			to.Ptr(armcontainerregistry.WebhookActionPush)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr("myRepository"),
	// 			Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookPing.json
func ExampleWebhooksClient_Ping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Ping(ctx, "myResourceGroup", "myRegistry", "myWebhook", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventInfo = armcontainerregistry.EventInfo{
	// 	ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookListEvents.json
func ExampleWebhooksClient_NewListEventsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListEventsPager("myResourceGroup", "myRegistry", "myWebhook", nil)
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
		// page.EventListResult = armcontainerregistry.EventListResult{
		// 	Value: []*armcontainerregistry.Event{
		// 		{
		// 			ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			EventRequestMessage: &armcontainerregistry.EventRequestMessage{
		// 				Method: to.Ptr("POST"),
		// 				Content: &armcontainerregistry.EventContent{
		// 					Action: to.Ptr("push"),
		// 					Actor: &armcontainerregistry.Actor{
		// 					},
		// 					ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Source: &armcontainerregistry.Source{
		// 						Addr: to.Ptr("xtal.local:5000"),
		// 						InstanceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 					Target: &armcontainerregistry.Target{
		// 						Digest: to.Ptr("sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf"),
		// 						Length: to.Ptr[int64](708),
		// 						MediaType: to.Ptr("application/vnd.docker.distribution.manifest.v2+json"),
		// 						Repository: to.Ptr("hello-world"),
		// 						Size: to.Ptr[int64](708),
		// 						Tag: to.Ptr("latest"),
		// 						URL: to.Ptr("http://192.168.100.227:5000/v2/hello-world/manifests/sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf"),
		// 					},
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
		// 					Request: &armcontainerregistry.Request{
		// 						Method: to.Ptr("GET"),
		// 						Addr: to.Ptr("192.168.64.11:42961"),
		// 						Host: to.Ptr("192.168.100.227:5000"),
		// 						ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						Useragent: to.Ptr("curl/7.38.0"),
		// 					},
		// 				},
		// 				Headers: map[string]*string{
		// 					"Authorization": to.Ptr("******"),
		// 					"Content-Length": to.Ptr("719"),
		// 					"Content-Type": to.Ptr("application/json"),
		// 				},
		// 				RequestURI: to.Ptr("http://myservice.com"),
		// 				Version: to.Ptr("1.1"),
		// 			},
		// 			EventResponseMessage: &armcontainerregistry.EventResponseMessage{
		// 				Headers: map[string]*string{
		// 					"Content-Length": to.Ptr("0"),
		// 				},
		// 				StatusCode: to.Ptr("200"),
		// 				Version: to.Ptr("1.1"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/WebhookGetCallbackConfig.json
func ExampleWebhooksClient_GetCallbackConfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetCallbackConfig(ctx, "myResourceGroup", "myRegistry", "myWebhook", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CallbackConfig = armcontainerregistry.CallbackConfig{
	// 	CustomHeaders: map[string]*string{
	// 		"Authorization": to.Ptr("Basic 000000000000000000000000000000000000000000000000000"),
	// 	},
	// 	ServiceURI: to.Ptr("http://myservice.com"),
	// }
}
