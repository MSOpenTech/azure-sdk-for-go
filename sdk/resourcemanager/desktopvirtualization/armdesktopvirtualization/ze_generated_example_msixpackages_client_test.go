//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_Get.json
func ExampleMSIXPackagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_Create.json
func ExampleMSIXPackagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		armdesktopvirtualization.MSIXPackage{
			Properties: &armdesktopvirtualization.MSIXPackageProperties{
				DisplayName:           to.Ptr("<display-name>"),
				ImagePath:             to.Ptr("<image-path>"),
				IsActive:              to.Ptr(false),
				IsRegularRegistration: to.Ptr(false),
				LastUpdated:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t }()),
				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
					{
						Description:    to.Ptr("<description>"),
						AppID:          to.Ptr("<app-id>"),
						AppUserModelID: to.Ptr("<app-user-model-id>"),
						FriendlyName:   to.Ptr("<friendly-name>"),
						IconImageName:  to.Ptr("<icon-image-name>"),
						RawIcon:        []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
						RawPNG:         []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
					}},
				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
					{
						DependencyName: to.Ptr("<dependency-name>"),
						MinVersion:     to.Ptr("<min-version>"),
						Publisher:      to.Ptr("<publisher>"),
					}},
				PackageFamilyName:   to.Ptr("<package-family-name>"),
				PackageName:         to.Ptr("<package-name>"),
				PackageRelativePath: to.Ptr("<package-relative-path>"),
				Version:             to.Ptr("<version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_Delete.json
func ExampleMSIXPackagesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_Update.json
func ExampleMSIXPackagesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		&armdesktopvirtualization.MSIXPackagesClientUpdateOptions{MsixPackage: &armdesktopvirtualization.MSIXPackagePatch{
			Properties: &armdesktopvirtualization.MSIXPackagePatchProperties{
				DisplayName:           to.Ptr("<display-name>"),
				IsActive:              to.Ptr(true),
				IsRegularRegistration: to.Ptr(false),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_List.json
func ExampleMSIXPackagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager("<resource-group-name>",
		"<host-pool-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
