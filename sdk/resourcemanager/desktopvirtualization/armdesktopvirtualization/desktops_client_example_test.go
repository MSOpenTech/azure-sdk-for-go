//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/Desktop_Get.json
func ExampleDesktopsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDesktopsClient().Get(ctx, "resourceGroup1", "applicationGroup1", "SessionDesktop", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Desktop = armdesktopvirtualization.Desktop{
	// 	Name: to.Ptr("applicationGroup1/SessionDesktop"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups/desktops"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1/applications/SessionDesktop"),
	// 	Properties: &armdesktopvirtualization.DesktopProperties{
	// 		Description: to.Ptr("des1"),
	// 		FriendlyName: to.Ptr("friendly"),
	// 		IconContent: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		IconHash: to.Ptr("bEQ1n2HysrGxCDvgt4bfOtkET8ydzh5SXGM0KGTBPVc"),
	// 		ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/Desktop_Update.json
func ExampleDesktopsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDesktopsClient().Update(ctx, "resourceGroup1", "applicationGroup1", "SessionDesktop", &armdesktopvirtualization.DesktopsClientUpdateOptions{Desktop: &armdesktopvirtualization.DesktopPatch{
		Properties: &armdesktopvirtualization.DesktopPatchProperties{
			Description:  to.Ptr("des1"),
			FriendlyName: to.Ptr("friendly"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Desktop = armdesktopvirtualization.Desktop{
	// 	Name: to.Ptr("applicationGroup1/SessionDesktop"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups/desktops"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1/desktops/SessionDesktop"),
	// 	Properties: &armdesktopvirtualization.DesktopProperties{
	// 		Description: to.Ptr("des1"),
	// 		FriendlyName: to.Ptr("friendly"),
	// 		IconContent: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		IconHash: to.Ptr("bEQ1n2HysrGxCDvgt4bfOtkET8ydzh5SXGM0KGTBPVc"),
	// 		ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/Desktop_List.json
func ExampleDesktopsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDesktopsClient().NewListPager("resourceGroup1", "applicationGroup1", &armdesktopvirtualization.DesktopsClientListOptions{PageSize: to.Ptr[int32](10),
		IsDescending: to.Ptr(true),
		InitialSkip:  to.Ptr[int32](0),
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
		// page.DesktopList = armdesktopvirtualization.DesktopList{
		// 	Value: []*armdesktopvirtualization.Desktop{
		// 		{
		// 			Name: to.Ptr("applicationGroup1/SessionDesktop"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups/desktops"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1/applications/SessionDesktop"),
		// 			Properties: &armdesktopvirtualization.DesktopProperties{
		// 				Description: to.Ptr("des1"),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				IconContent: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
		// 				IconHash: to.Ptr("bEQ1n2HysrGxCDvgt4bfOtkET8ydzh5SXGM0KGTBPVc"),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
