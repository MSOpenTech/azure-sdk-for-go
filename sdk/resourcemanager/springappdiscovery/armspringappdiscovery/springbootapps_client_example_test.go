//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armspringappdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Get_MaximumSet_Gen.json
func ExampleSpringbootappsClient_Get_springbootappsGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSpringbootappsClient().Get(ctx, "rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd", "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SpringbootappsModel = armspringappdiscovery.SpringbootappsModel{
	// 	Name: to.Ptr("enyeyrgonjdauhscqy"),
	// 	Type: to.Ptr("hvtpwijptwksiyxdrmnpsv"),
	// 	ID: to.Ptr("ftbefgkqxhwszzienudx"),
	// 	SystemData: &armspringappdiscovery.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("kmiwjeuplqvqwk"),
	// 		CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("pmmzrztbtuuj"),
	// 		LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
	// 	},
	// 	Properties: &armspringappdiscovery.SpringbootappsProperties{
	// 		AppName: to.Ptr("wrauwfegjfccym"),
	// 		AppPort: to.Ptr[int32](12),
	// 		AppType: to.Ptr("axzunlh"),
	// 		ApplicationConfigurations: []*armspringappdiscovery.SpringbootappsPropertiesApplicationConfigurationsItem{
	// 			{
	// 				Key: to.Ptr("wrbnwhqxjextxgdfbonuynvs"),
	// 				Value: to.Ptr("jnmaf"),
	// 		}},
	// 		ArtifactName: to.Ptr("wrauwfegjfccym"),
	// 		BindingPorts: []*int32{
	// 			to.Ptr[int32](11)},
	// 			BuildJdkVersion: to.Ptr("ipzruwqqulkpvhzymqegntz"),
	// 			Certificates: []*string{
	// 				to.Ptr("xpiqqob")},
	// 				Checksum: to.Ptr("gpzumvbzfnhhmuehveanctiamr"),
	// 				Dependencies: []*string{
	// 					to.Ptr("zrtted")},
	// 					Environments: []*string{
	// 						to.Ptr("afhprevtcx")},
	// 						InstanceCount: to.Ptr[int32](5),
	// 						Instances: []*armspringappdiscovery.SpringbootappsPropertiesInstancesItem{
	// 							{
	// 								InstanceCount: to.Ptr[int32](5),
	// 								JvmMemoryInMB: to.Ptr[int32](128),
	// 								MachineArmID: to.Ptr("lsstlommxuskyhnwyxh"),
	// 						}},
	// 						JarFileLocation: to.Ptr("wfptqclncaqycyfbfih"),
	// 						JvmMemoryInMB: to.Ptr[int32](1),
	// 						JvmOptions: []*string{
	// 							to.Ptr("nytejjoytevmvlgnfwb")},
	// 							LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 							LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 							MachineArmIDs: []*string{
	// 								to.Ptr("lsstlommxuskyhnwyxh")},
	// 								Miscs: []*armspringappdiscovery.SpringbootappsPropertiesMiscsItem{
	// 									{
	// 										Key: to.Ptr("fobsfetkynfmkziei"),
	// 										Value: to.Ptr("k"),
	// 								}},
	// 								ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
	// 								RuntimeJdkVersion: to.Ptr("eblzujbsulpeilykqyjso"),
	// 								Servers: []*string{
	// 									to.Ptr("gvfhsohasdx")},
	// 									SiteName: to.Ptr("nzzyrevhsz"),
	// 									SpringBootVersion: to.Ptr("euggigfiii"),
	// 									StaticContentLocations: []*string{
	// 										to.Ptr("wvvajfkbtmjftir")},
	// 									},
	// 								}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Get_MinimumSet_Gen.json
func ExampleSpringbootappsClient_Get_springbootappsGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSpringbootappsClient().Get(ctx, "rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd", "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SpringbootappsModel = armspringappdiscovery.SpringbootappsModel{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.OffAzureSpringBoot/springbootsites/testsite1/springbootapps/app1"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Update_MaximumSet_Gen.json
func ExampleSpringbootappsClient_BeginUpdate_springbootappsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSpringbootappsClient().BeginUpdate(ctx, "rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd", "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf", armspringappdiscovery.SpringbootappsPatch{}, nil)
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
	// res.SpringbootappsModel = armspringappdiscovery.SpringbootappsModel{
	// 	Name: to.Ptr("enyeyrgonjdauhscqy"),
	// 	Type: to.Ptr("hvtpwijptwksiyxdrmnpsv"),
	// 	ID: to.Ptr("ftbefgkqxhwszzienudx"),
	// 	SystemData: &armspringappdiscovery.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("kmiwjeuplqvqwk"),
	// 		CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("pmmzrztbtuuj"),
	// 		LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
	// 	},
	// 	Properties: &armspringappdiscovery.SpringbootappsProperties{
	// 		AppName: to.Ptr("wrauwfegjfccym"),
	// 		AppPort: to.Ptr[int32](12),
	// 		AppType: to.Ptr("axzunlh"),
	// 		ApplicationConfigurations: []*armspringappdiscovery.SpringbootappsPropertiesApplicationConfigurationsItem{
	// 			{
	// 				Key: to.Ptr("wrbnwhqxjextxgdfbonuynvs"),
	// 				Value: to.Ptr("jnmaf"),
	// 		}},
	// 		ArtifactName: to.Ptr("wrauwfegjfccym"),
	// 		BindingPorts: []*int32{
	// 			to.Ptr[int32](11)},
	// 			BuildJdkVersion: to.Ptr("ipzruwqqulkpvhzymqegntz"),
	// 			Certificates: []*string{
	// 				to.Ptr("xpiqqob")},
	// 				Checksum: to.Ptr("gpzumvbzfnhhmuehveanctiamr"),
	// 				Dependencies: []*string{
	// 					to.Ptr("zrtted")},
	// 					Environments: []*string{
	// 						to.Ptr("afhprevtcx")},
	// 						InstanceCount: to.Ptr[int32](5),
	// 						Instances: []*armspringappdiscovery.SpringbootappsPropertiesInstancesItem{
	// 							{
	// 								InstanceCount: to.Ptr[int32](5),
	// 								JvmMemoryInMB: to.Ptr[int32](128),
	// 								MachineArmID: to.Ptr("lsstlommxuskyhnwyxh"),
	// 						}},
	// 						JarFileLocation: to.Ptr("wfptqclncaqycyfbfih"),
	// 						JvmMemoryInMB: to.Ptr[int32](1),
	// 						JvmOptions: []*string{
	// 							to.Ptr("nytejjoytevmvlgnfwb")},
	// 							MachineArmIDs: []*string{
	// 								to.Ptr("lsstlommxuskyhnwyxh")},
	// 								Miscs: []*armspringappdiscovery.SpringbootappsPropertiesMiscsItem{
	// 									{
	// 										Key: to.Ptr("fobsfetkynfmkziei"),
	// 										Value: to.Ptr("k"),
	// 								}},
	// 								ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
	// 								RuntimeJdkVersion: to.Ptr("eblzujbsulpeilykqyjso"),
	// 								Servers: []*string{
	// 									to.Ptr("gvfhsohasdx")},
	// 									SiteName: to.Ptr("nzzyrevhsz"),
	// 									SpringBootVersion: to.Ptr("euggigfiii"),
	// 									StaticContentLocations: []*string{
	// 										to.Ptr("wvvajfkbtmjftir")},
	// 									},
	// 								}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Update_MinimumSet_Gen.json
func ExampleSpringbootappsClient_BeginUpdate_springbootappsUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSpringbootappsClient().BeginUpdate(ctx, "rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd", "ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf", armspringappdiscovery.SpringbootappsPatch{}, nil)
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
	// res.SpringbootappsModel = armspringappdiscovery.SpringbootappsModel{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_ListByResourceGroup_MaximumSet_Gen.json
func ExampleSpringbootappsClient_NewListByResourceGroupPager_springbootappsListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSpringbootappsClient().NewListByResourceGroupPager("rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd", nil)
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
		// page.SpringbootappsListResult = armspringappdiscovery.SpringbootappsListResult{
		// 	Value: []*armspringappdiscovery.SpringbootappsModel{
		// 		{
		// 			Name: to.Ptr("enyeyrgonjdauhscqy"),
		// 			Type: to.Ptr("hvtpwijptwksiyxdrmnpsv"),
		// 			ID: to.Ptr("ftbefgkqxhwszzienudx"),
		// 			SystemData: &armspringappdiscovery.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 				CreatedBy: to.Ptr("kmiwjeuplqvqwk"),
		// 				CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("pmmzrztbtuuj"),
		// 				LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 			},
		// 			Properties: &armspringappdiscovery.SpringbootappsProperties{
		// 				AppName: to.Ptr("wrauwfegjfccym"),
		// 				AppPort: to.Ptr[int32](12),
		// 				AppType: to.Ptr("axzunlh"),
		// 				ApplicationConfigurations: []*armspringappdiscovery.SpringbootappsPropertiesApplicationConfigurationsItem{
		// 					{
		// 						Key: to.Ptr("wrbnwhqxjextxgdfbonuynvs"),
		// 						Value: to.Ptr("jnmaf"),
		// 				}},
		// 				ArtifactName: to.Ptr("wrauwfegjfccym"),
		// 				BindingPorts: []*int32{
		// 					to.Ptr[int32](11)},
		// 					BuildJdkVersion: to.Ptr("ipzruwqqulkpvhzymqegntz"),
		// 					Certificates: []*string{
		// 						to.Ptr("xpiqqob")},
		// 						Checksum: to.Ptr("gpzumvbzfnhhmuehveanctiamr"),
		// 						Dependencies: []*string{
		// 							to.Ptr("zrtted")},
		// 							Environments: []*string{
		// 								to.Ptr("afhprevtcx")},
		// 								InstanceCount: to.Ptr[int32](5),
		// 								Instances: []*armspringappdiscovery.SpringbootappsPropertiesInstancesItem{
		// 									{
		// 										InstanceCount: to.Ptr[int32](5),
		// 										JvmMemoryInMB: to.Ptr[int32](128),
		// 										MachineArmID: to.Ptr("lsstlommxuskyhnwyxh"),
		// 								}},
		// 								JarFileLocation: to.Ptr("wfptqclncaqycyfbfih"),
		// 								JvmMemoryInMB: to.Ptr[int32](1),
		// 								JvmOptions: []*string{
		// 									to.Ptr("nytejjoytevmvlgnfwb")},
		// 									LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 									LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 									MachineArmIDs: []*string{
		// 										to.Ptr("lsstlommxuskyhnwyxh")},
		// 										Miscs: []*armspringappdiscovery.SpringbootappsPropertiesMiscsItem{
		// 											{
		// 												Key: to.Ptr("fobsfetkynfmkziei"),
		// 												Value: to.Ptr("k"),
		// 										}},
		// 										ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
		// 										RuntimeJdkVersion: to.Ptr("eblzujbsulpeilykqyjso"),
		// 										Servers: []*string{
		// 											to.Ptr("gvfhsohasdx")},
		// 											SiteName: to.Ptr("nzzyrevhsz"),
		// 											SpringBootVersion: to.Ptr("euggigfiii"),
		// 											StaticContentLocations: []*string{
		// 												to.Ptr("wvvajfkbtmjftir")},
		// 											},
		// 									}},
		// 								}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_ListByResourceGroup_MinimumSet_Gen.json
func ExampleSpringbootappsClient_NewListByResourceGroupPager_springbootappsListByResourceGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSpringbootappsClient().NewListByResourceGroupPager("rgspringbootapps", "pdfosfhtemfsaglvwjdyqlyeipucrd", nil)
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
		// page.SpringbootappsListResult = armspringappdiscovery.SpringbootappsListResult{
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_ListBySubscription_MaximumSet_Gen.json
func ExampleSpringbootappsClient_NewListBySubscriptionPager_springbootappsListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSpringbootappsClient().NewListBySubscriptionPager("pdfosfhtemfsaglvwjdyqlyeipucrd", nil)
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
		// page.SpringbootappsListResult = armspringappdiscovery.SpringbootappsListResult{
		// 	Value: []*armspringappdiscovery.SpringbootappsModel{
		// 		{
		// 			Name: to.Ptr("enyeyrgonjdauhscqy"),
		// 			Type: to.Ptr("hvtpwijptwksiyxdrmnpsv"),
		// 			ID: to.Ptr("ftbefgkqxhwszzienudx"),
		// 			SystemData: &armspringappdiscovery.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 				CreatedBy: to.Ptr("kmiwjeuplqvqwk"),
		// 				CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("pmmzrztbtuuj"),
		// 				LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 			},
		// 			Properties: &armspringappdiscovery.SpringbootappsProperties{
		// 				AppName: to.Ptr("wrauwfegjfccym"),
		// 				AppPort: to.Ptr[int32](12),
		// 				AppType: to.Ptr("axzunlh"),
		// 				ApplicationConfigurations: []*armspringappdiscovery.SpringbootappsPropertiesApplicationConfigurationsItem{
		// 					{
		// 						Key: to.Ptr("wrbnwhqxjextxgdfbonuynvs"),
		// 						Value: to.Ptr("jnmaf"),
		// 				}},
		// 				ArtifactName: to.Ptr("wrauwfegjfccym"),
		// 				BindingPorts: []*int32{
		// 					to.Ptr[int32](11)},
		// 					BuildJdkVersion: to.Ptr("ipzruwqqulkpvhzymqegntz"),
		// 					Certificates: []*string{
		// 						to.Ptr("xpiqqob")},
		// 						Checksum: to.Ptr("gpzumvbzfnhhmuehveanctiamr"),
		// 						Dependencies: []*string{
		// 							to.Ptr("zrtted")},
		// 							Environments: []*string{
		// 								to.Ptr("afhprevtcx")},
		// 								InstanceCount: to.Ptr[int32](5),
		// 								Instances: []*armspringappdiscovery.SpringbootappsPropertiesInstancesItem{
		// 									{
		// 										InstanceCount: to.Ptr[int32](5),
		// 										JvmMemoryInMB: to.Ptr[int32](128),
		// 										MachineArmID: to.Ptr("lsstlommxuskyhnwyxh"),
		// 								}},
		// 								JarFileLocation: to.Ptr("wfptqclncaqycyfbfih"),
		// 								JvmMemoryInMB: to.Ptr[int32](1),
		// 								JvmOptions: []*string{
		// 									to.Ptr("nytejjoytevmvlgnfwb")},
		// 									LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 									LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:28:24.094Z"); return t}()),
		// 									MachineArmIDs: []*string{
		// 										to.Ptr("lsstlommxuskyhnwyxh")},
		// 										Miscs: []*armspringappdiscovery.SpringbootappsPropertiesMiscsItem{
		// 											{
		// 												Key: to.Ptr("fobsfetkynfmkziei"),
		// 												Value: to.Ptr("k"),
		// 										}},
		// 										ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
		// 										RuntimeJdkVersion: to.Ptr("eblzujbsulpeilykqyjso"),
		// 										Servers: []*string{
		// 											to.Ptr("gvfhsohasdx")},
		// 											SiteName: to.Ptr("nzzyrevhsz"),
		// 											SpringBootVersion: to.Ptr("euggigfiii"),
		// 											StaticContentLocations: []*string{
		// 												to.Ptr("wvvajfkbtmjftir")},
		// 											},
		// 									}},
		// 								}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_ListBySubscription_MinimumSet_Gen.json
func ExampleSpringbootappsClient_NewListBySubscriptionPager_springbootappsListBySubscriptionMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSpringbootappsClient().NewListBySubscriptionPager("pdfosfhtemfsaglvwjdyqlyeipucrd", nil)
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
		// page.SpringbootappsListResult = armspringappdiscovery.SpringbootappsListResult{
		// }
	}
}
