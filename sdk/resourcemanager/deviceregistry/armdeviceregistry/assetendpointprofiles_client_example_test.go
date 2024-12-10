//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd37988c515821dcd48192ec1db5cb39c39e3c97/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/List_AssetEndpointProfiles_Subscription.json
func ExampleAssetEndpointProfilesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssetEndpointProfilesClient().NewListBySubscriptionPager(nil)
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
		// page.AssetEndpointProfileListResult = armdeviceregistry.AssetEndpointProfileListResult{
		// 	Value: []*armdeviceregistry.AssetEndpointProfile{
		// 		{
		// 			Name: to.Ptr("my-assetendpointprofile"),
		// 			Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
		// 			SystemData: &armdeviceregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 				"site": to.Ptr("building-1"),
		// 			},
		// 			ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 				ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
		// 				UserAuthentication: &armdeviceregistry.UserAuthentication{
		// 					Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeAnonymous),
		// 				},
		// 				UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-assetendpointprofile1"),
		// 			Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile1"),
		// 			SystemData: &armdeviceregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 				"site": to.Ptr("building-2"),
		// 			},
		// 			ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 				ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				TargetAddress: to.Ptr("https://www.example.com/myTargetAddress1"),
		// 				UserAuthentication: &armdeviceregistry.UserAuthentication{
		// 					Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeCertificate),
		// 					X509Credentials: &armdeviceregistry.X509Credentials{
		// 						CertificateReference: to.Ptr("certificatRef"),
		// 					},
		// 				},
		// 				UUID: to.Ptr("7824a74f-21e1-4458-ae06-604d3a241d2c"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-assetendpointprofile2"),
		// 			Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile2"),
		// 			SystemData: &armdeviceregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 				"site": to.Ptr("building-2"),
		// 			},
		// 			ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 				ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				TargetAddress: to.Ptr("https://www.example.com/myTargetAddress2"),
		// 				TransportAuthentication: &armdeviceregistry.TransportAuthentication{
		// 					OwnCertificates: []*armdeviceregistry.OwnCertificate{
		// 						{
		// 							CertPasswordReference: to.Ptr("passwordRef"),
		// 							CertSecretReference: to.Ptr("secretRef"),
		// 							CertThumbprint: to.Ptr("myThumbprint"),
		// 					}},
		// 				},
		// 				UserAuthentication: &armdeviceregistry.UserAuthentication{
		// 					Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeUsernamePassword),
		// 					UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
		// 						PasswordReference: to.Ptr("passwordRef"),
		// 						UsernameReference: to.Ptr("usernameRef"),
		// 					},
		// 				},
		// 				UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd37988c515821dcd48192ec1db5cb39c39e3c97/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/List_AssetEndpointProfiles_ResourceGroup.json
func ExampleAssetEndpointProfilesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssetEndpointProfilesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.AssetEndpointProfileListResult = armdeviceregistry.AssetEndpointProfileListResult{
		// 	Value: []*armdeviceregistry.AssetEndpointProfile{
		// 		{
		// 			Name: to.Ptr("my-assetendpointprofile"),
		// 			Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
		// 			SystemData: &armdeviceregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 				"site": to.Ptr("building-1"),
		// 			},
		// 			ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 				ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
		// 				UserAuthentication: &armdeviceregistry.UserAuthentication{
		// 					Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeAnonymous),
		// 				},
		// 				UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-assetendpointprofile1"),
		// 			Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile1"),
		// 			SystemData: &armdeviceregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 				"site": to.Ptr("building-2"),
		// 			},
		// 			ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 				ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				TargetAddress: to.Ptr("https://www.example.com/myTargetAddress1"),
		// 				UserAuthentication: &armdeviceregistry.UserAuthentication{
		// 					Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeCertificate),
		// 					X509Credentials: &armdeviceregistry.X509Credentials{
		// 						CertificateReference: to.Ptr("certificatRef"),
		// 					},
		// 				},
		// 				UUID: to.Ptr("7824a74f-21e1-4458-ae06-604d3a241d2c"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-assetendpointprofile2"),
		// 			Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile2"),
		// 			SystemData: &armdeviceregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
		// 				CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 				LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West Europe"),
		// 			Tags: map[string]*string{
		// 				"site": to.Ptr("building-2"),
		// 			},
		// 			ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 				ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				TargetAddress: to.Ptr("https://www.example.com/myTargetAddress2"),
		// 				TransportAuthentication: &armdeviceregistry.TransportAuthentication{
		// 					OwnCertificates: []*armdeviceregistry.OwnCertificate{
		// 						{
		// 							CertPasswordReference: to.Ptr("passwordRef"),
		// 							CertSecretReference: to.Ptr("secretRef"),
		// 							CertThumbprint: to.Ptr("myThumbprint"),
		// 					}},
		// 				},
		// 				UserAuthentication: &armdeviceregistry.UserAuthentication{
		// 					Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeUsernamePassword),
		// 					UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
		// 						PasswordReference: to.Ptr("passwordRef"),
		// 						UsernameReference: to.Ptr("usernameRef"),
		// 					},
		// 				},
		// 				UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd37988c515821dcd48192ec1db5cb39c39e3c97/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Get_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetEndpointProfilesClient().Get(ctx, "myResourceGroup", "my-assetendpointprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssetEndpointProfile = armdeviceregistry.AssetEndpointProfile{
	// 	Name: to.Ptr("my-assetendpointprofile"),
	// 	Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 	SystemData: &armdeviceregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
	// 		CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West Europe"),
	// 	Tags: map[string]*string{
	// 		"site": to.Ptr("building-1"),
	// 	},
	// 	ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 		ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 		TransportAuthentication: &armdeviceregistry.TransportAuthentication{
	// 			OwnCertificates: []*armdeviceregistry.OwnCertificate{
	// 				{
	// 					CertPasswordReference: to.Ptr("passwordRef"),
	// 					CertSecretReference: to.Ptr("secretRef"),
	// 					CertThumbprint: to.Ptr("myThumbprint"),
	// 			}},
	// 		},
	// 		UserAuthentication: &armdeviceregistry.UserAuthentication{
	// 			Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeUsernamePassword),
	// 			UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
	// 				PasswordReference: to.Ptr("passwordRef"),
	// 				UsernameReference: to.Ptr("usernameRef"),
	// 			},
	// 		},
	// 		UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd37988c515821dcd48192ec1db5cb39c39e3c97/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Create_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginCreateOrReplace(ctx, "myResourceGroup", "my-assetendpointprofile", armdeviceregistry.AssetEndpointProfile{
		Location: to.Ptr("West Europe"),
		Tags: map[string]*string{
			"site": to.Ptr("building-1"),
		},
		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
			UserAuthentication: &armdeviceregistry.UserAuthentication{
				Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeAnonymous),
			},
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
	// res.AssetEndpointProfile = armdeviceregistry.AssetEndpointProfile{
	// 	Name: to.Ptr("my-assetendpointprofile"),
	// 	Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 	SystemData: &armdeviceregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
	// 		CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West Europe"),
	// 	Tags: map[string]*string{
	// 		"site": to.Ptr("building-1"),
	// 	},
	// 	ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 		ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 		UserAuthentication: &armdeviceregistry.UserAuthentication{
	// 			Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeAnonymous),
	// 		},
	// 		UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd37988c515821dcd48192ec1db5cb39c39e3c97/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Update_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginUpdate(ctx, "myResourceGroup", "my-assetendpointprofile", armdeviceregistry.AssetEndpointProfileUpdate{
		Properties: &armdeviceregistry.AssetEndpointProfileUpdateProperties{
			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
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
	// res.AssetEndpointProfile = armdeviceregistry.AssetEndpointProfile{
	// 	Name: to.Ptr("my-assetendpointprofile"),
	// 	Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 	SystemData: &armdeviceregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.251Z"); return t}()),
	// 		CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.092Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 		LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West Europe"),
	// 	Tags: map[string]*string{
	// 		"site": to.Ptr("building-1"),
	// 	},
	// 	ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 		ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 		UserAuthentication: &armdeviceregistry.UserAuthentication{
	// 			Mode: to.Ptr(armdeviceregistry.UserAuthenticationModeAnonymous),
	// 		},
	// 		UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd37988c515821dcd48192ec1db5cb39c39e3c97/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Delete_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginDelete(ctx, "myResourceGroup", "my-assetendpointprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
