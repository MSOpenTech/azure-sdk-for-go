//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_ListByInformaticaOrganizationResource_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_NewListByInformaticaOrganizationResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerlessRuntimesClient().NewListByInformaticaOrganizationResourcePager("rgopenapi", "orgName", nil)
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
		// page.InformaticaServerlessRuntimeResourceListResult = arminformaticadatamgmt.InformaticaServerlessRuntimeResourceListResult{
		// 	Value: []*arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
		// 		{
		// 			Name: to.Ptr("byzccgftqjthb"),
		// 			Type: to.Ptr("due"),
		// 			ID: to.Ptr("vcdjzfgqjv"),
		// 			SystemData: &arminformaticadatamgmt.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
		// 				CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
		// 				CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
		// 				LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
		// 			},
		// 			Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
		// 				Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
		// 				AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
		// 					{
		// 						Key: to.Ptr("qcmc"),
		// 						Value: to.Ptr("unraxmnohdmvutt"),
		// 				}},
		// 				ApplicationType: to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
		// 				ComputeUnits: to.Ptr("bsctukmndvowse"),
		// 				ExecutionTimeout: to.Ptr("ruiougpypny"),
		// 				Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
		// 				ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
		// 				ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
		// 				ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
		// 					CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
		// 						{
		// 							ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
		// 								{
		// 									Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
		// 									Type: to.Ptr("lw"),
		// 									Customized: to.Ptr("j"),
		// 									DefaultValue: to.Ptr("zvgkqwmi"),
		// 									Platform: to.Ptr("dixfyeobngivyvf"),
		// 									Value: to.Ptr("mozgsetpwjmtyl"),
		// 							}},
		// 							EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
		// 							EngineVersion: to.Ptr("zlrlbg"),
		// 					}},
		// 					CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
		// 						{
		// 							ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
		// 								{
		// 									Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
		// 									Type: to.Ptr("lw"),
		// 									Customized: to.Ptr("j"),
		// 									DefaultValue: to.Ptr("zvgkqwmi"),
		// 									Platform: to.Ptr("dixfyeobngivyvf"),
		// 									Value: to.Ptr("mozgsetpwjmtyl"),
		// 							}},
		// 							EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
		// 							EngineVersion: to.Ptr("zlrlbg"),
		// 					}},
		// 				},
		// 				ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
		// 					NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
		// 						SubnetID: to.Ptr("s"),
		// 						VnetID: to.Ptr("uaqjvtubxccjs"),
		// 					},
		// 				},
		// 				ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
		// 					{
		// 						Name: to.Ptr("korveuycuwhs"),
		// 						Value: to.Ptr("uyiuegxnkgp"),
		// 				}},
		// 				ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
		// 					UserContextToken: to.Ptr("oludf"),
		// 				},
		// 				SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Get_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_Get_serverlessRuntimesGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().Get(ctx, "rgopenapi", "e3Y", "48-", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	Name: to.Ptr("byzccgftqjthb"),
	// 	Type: to.Ptr("due"),
	// 	ID: to.Ptr("vcdjzfgqjv"),
	// 	SystemData: &arminformaticadatamgmt.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
	// 		CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
	// 		LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 	},
	// 	Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
	// 		Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
	// 		AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
	// 			{
	// 				Key: to.Ptr("qcmc"),
	// 				Value: to.Ptr("unraxmnohdmvutt"),
	// 		}},
	// 		ApplicationType: to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
	// 		ComputeUnits: to.Ptr("bsctukmndvowse"),
	// 		ExecutionTimeout: to.Ptr("ruiougpypny"),
	// 		Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
	// 		ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
	// 		ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
	// 		ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
	// 			CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 			CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 		},
	// 		ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
	// 			NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
	// 				SubnetID: to.Ptr("s"),
	// 				VnetID: to.Ptr("uaqjvtubxccjs"),
	// 				VnetResourceGUID: to.Ptr("5328d299-1462-4be0-bef1-303a28e556a0"),
	// 			},
	// 		},
	// 		ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
	// 			{
	// 				Name: to.Ptr("korveuycuwhs"),
	// 				Value: to.Ptr("uyiuegxnkgp"),
	// 		}},
	// 		ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
	// 			UserContextToken: to.Ptr("oludf"),
	// 		},
	// 		SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Get_MinimumSet_Gen.json
func ExampleServerlessRuntimesClient_Get_serverlessRuntimesGetMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().Get(ctx, "rgopenapi", "YC", "___", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	ID: to.Ptr("cadokiejnrth"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CreateOrUpdate_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_BeginCreateOrUpdate_serverlessRuntimesCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerlessRuntimesClient().BeginCreateOrUpdate(ctx, "rgopenapi", "__C", "0j-__", arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
		Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
			Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
			AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
				{
					Key:   to.Ptr("qcmc"),
					Value: to.Ptr("unraxmnohdmvutt"),
				}},
			ApplicationType:           to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
			ComputeUnits:              to.Ptr("bsctukmndvowse"),
			ExecutionTimeout:          to.Ptr("ruiougpypny"),
			Platform:                  to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
			ProvisioningState:         to.Ptr(arminformaticadatamgmt.ProvisioningStateAccepted),
			ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
			ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
				CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
					{
						ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
							{
								Name:         to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
								Type:         to.Ptr("lw"),
								Customized:   to.Ptr("j"),
								DefaultValue: to.Ptr("zvgkqwmi"),
								Platform:     to.Ptr("dixfyeobngivyvf"),
								Value:        to.Ptr("mozgsetpwjmtyl"),
							}},
						EngineName:    to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
						EngineVersion: to.Ptr("zlrlbg"),
					}},
				CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
					{
						ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
							{
								Name:         to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
								Type:         to.Ptr("lw"),
								Customized:   to.Ptr("j"),
								DefaultValue: to.Ptr("zvgkqwmi"),
								Platform:     to.Ptr("dixfyeobngivyvf"),
								Value:        to.Ptr("mozgsetpwjmtyl"),
							}},
						EngineName:    to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
						EngineVersion: to.Ptr("zlrlbg"),
					}},
			},
			ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
				NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
					SubnetID:         to.Ptr("s"),
					VnetID:           to.Ptr("uaqjvtubxccjs"),
					VnetResourceGUID: to.Ptr("5328d299-1462-4be0-bef1-303a28e556a0"),
				},
			},
			ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
				{
					Name:  to.Ptr("korveuycuwhs"),
					Value: to.Ptr("uyiuegxnkgp"),
				}},
			ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
				UserContextToken: to.Ptr("oludf"),
			},
			SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
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
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	Name: to.Ptr("byzccgftqjthb"),
	// 	Type: to.Ptr("due"),
	// 	ID: to.Ptr("vcdjzfgqjv"),
	// 	SystemData: &arminformaticadatamgmt.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
	// 		CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
	// 		LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 	},
	// 	Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
	// 		Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
	// 		AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
	// 			{
	// 				Key: to.Ptr("qcmc"),
	// 				Value: to.Ptr("unraxmnohdmvutt"),
	// 		}},
	// 		ApplicationType: to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
	// 		ComputeUnits: to.Ptr("bsctukmndvowse"),
	// 		ExecutionTimeout: to.Ptr("ruiougpypny"),
	// 		Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
	// 		ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
	// 		ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
	// 		ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
	// 			CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 			CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 		},
	// 		ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
	// 			NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
	// 				SubnetID: to.Ptr("s"),
	// 				VnetID: to.Ptr("uaqjvtubxccjs"),
	// 				VnetResourceGUID: to.Ptr("5328d299-1462-4be0-bef1-303a28e556a0"),
	// 			},
	// 		},
	// 		ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
	// 			{
	// 				Name: to.Ptr("korveuycuwhs"),
	// 				Value: to.Ptr("uyiuegxnkgp"),
	// 		}},
	// 		ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
	// 			UserContextToken: to.Ptr("oludf"),
	// 		},
	// 		SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CreateOrUpdate_MinimumSet_Gen.json
func ExampleServerlessRuntimesClient_BeginCreateOrUpdate_serverlessRuntimesCreateOrUpdateMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerlessRuntimesClient().BeginCreateOrUpdate(ctx, "rgopenapi", "-4Z__7", "J", arminformaticadatamgmt.InformaticaServerlessRuntimeResource{}, nil)
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
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	ID: to.Ptr("cadokiejnrth"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Update_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_Update_serverlessRuntimesUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().Update(ctx, "rgopenapi", "W5", "t_", arminformaticadatamgmt.InformaticaServerlessRuntimeResourceUpdate{
		Properties: &arminformaticadatamgmt.ServerlessRuntimePropertiesCustomUpdate{
			Description: to.Ptr("ocprslpljoikxyduackzqnkuhyzrh"),
			AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
				{
					Key:   to.Ptr("qcmc"),
					Value: to.Ptr("unraxmnohdmvutt"),
				}},
			ApplicationType:           to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
			ComputeUnits:              to.Ptr("uncwbpu"),
			ExecutionTimeout:          to.Ptr("tjyfytuywriabt"),
			Platform:                  to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
			ServerlessAccountLocation: to.Ptr("goaugkyfanqfnvcmntreibqrswfpis"),
			ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigPropertiesUpdate{
				CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
					{
						ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
							{
								Name:         to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
								Type:         to.Ptr("lw"),
								Customized:   to.Ptr("j"),
								DefaultValue: to.Ptr("zvgkqwmi"),
								Platform:     to.Ptr("dixfyeobngivyvf"),
								Value:        to.Ptr("mozgsetpwjmtyl"),
							}},
						EngineName:    to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
						EngineVersion: to.Ptr("zlrlbg"),
					}},
				CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
					{
						ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
							{
								Name:         to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
								Type:         to.Ptr("lw"),
								Customized:   to.Ptr("j"),
								DefaultValue: to.Ptr("zvgkqwmi"),
								Platform:     to.Ptr("dixfyeobngivyvf"),
								Value:        to.Ptr("mozgsetpwjmtyl"),
							}},
						EngineName:    to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
						EngineVersion: to.Ptr("zlrlbg"),
					}},
			},
			ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfileUpdate{
				NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfigurationUpdate{
					SubnetID:         to.Ptr("dctcuhgttxhcarwcrgdmsfwksyrzj"),
					VnetID:           to.Ptr("tnsqwwoxydeqqffumdnxlkkb"),
					VnetResourceGUID: to.Ptr("5328d299-1462-4be0-bef1-303a28e556a0"),
				},
			},
			ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
				{
					Name:  to.Ptr("korveuycuwhs"),
					Value: to.Ptr("uyiuegxnkgp"),
				}},
			ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextPropertiesUpdate{
				UserContextToken: to.Ptr("ctgebtvjhwh"),
			},
			SupplementaryFileLocation: to.Ptr("csxaqzpxu"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	Name: to.Ptr("byzccgftqjthb"),
	// 	Type: to.Ptr("due"),
	// 	ID: to.Ptr("vcdjzfgqjv"),
	// 	SystemData: &arminformaticadatamgmt.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
	// 		CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
	// 		LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 	},
	// 	Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
	// 		Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
	// 		AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
	// 			{
	// 				Key: to.Ptr("qcmc"),
	// 				Value: to.Ptr("unraxmnohdmvutt"),
	// 		}},
	// 		ApplicationType: to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
	// 		ComputeUnits: to.Ptr("bsctukmndvowse"),
	// 		ExecutionTimeout: to.Ptr("ruiougpypny"),
	// 		Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
	// 		ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
	// 		ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
	// 		ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
	// 			CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 			CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 		},
	// 		ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
	// 			NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
	// 				SubnetID: to.Ptr("dctcuhgttxhcarwcrgdmsfwksyrzj"),
	// 				VnetID: to.Ptr("tnsqwwoxydeqqffumdnxlkkb"),
	// 				VnetResourceGUID: to.Ptr("5328d299-1462-4be0-bef1-303a28e556a0"),
	// 			},
	// 		},
	// 		ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
	// 			{
	// 				Name: to.Ptr("korveuycuwhs"),
	// 				Value: to.Ptr("uyiuegxnkgp"),
	// 		}},
	// 		ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
	// 			UserContextToken: to.Ptr("ctgebtvjhwh"),
	// 		},
	// 		SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Update_MinimumSet_Gen.json
func ExampleServerlessRuntimesClient_Update_serverlessRuntimesUpdateMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().Update(ctx, "rgopenapi", "_f--", "8Zr__", arminformaticadatamgmt.InformaticaServerlessRuntimeResourceUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	ID: to.Ptr("cadokiejnrth"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Delete_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerlessRuntimesClient().BeginDelete(ctx, "rgopenapi", "orgName", "serverlessRuntimeName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CheckDependencies_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_CheckDependencies_serverlessRuntimesCheckDependencies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().CheckDependencies(ctx, "rgopenapi", "3P", "M", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckDependenciesResponse = arminformaticadatamgmt.CheckDependenciesResponse{
	// 	Count: to.Ptr[int32](28),
	// 	ID: to.Ptr("uwjsqpxr"),
	// 	References: []*arminformaticadatamgmt.ServerlessRuntimeDependency{
	// 		{
	// 			Path: to.Ptr("yxbpmcmfhhtht"),
	// 			Description: to.Ptr("vlkyqkevlrpge"),
	// 			AppContextID: to.Ptr("t"),
	// 			DocumentType: to.Ptr("jpcz"),
	// 			ID: to.Ptr("uzp"),
	// 			LastUpdatedTime: to.Ptr("yyf"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CheckDependencies_MinimumSet_Gen.json
func ExampleServerlessRuntimesClient_CheckDependencies_serverlessRuntimesCheckDependenciesMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().CheckDependencies(ctx, "rgopenapi", "_-", "_2_", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckDependenciesResponse = arminformaticadatamgmt.CheckDependenciesResponse{
	// 	Count: to.Ptr[int32](28),
	// 	ID: to.Ptr("uwjsqpxr"),
	// 	References: []*arminformaticadatamgmt.ServerlessRuntimeDependency{
	// 		{
	// 			Path: to.Ptr("yxbpmcmfhhtht"),
	// 			Description: to.Ptr("vlkyqkevlrpge"),
	// 			AppContextID: to.Ptr("t"),
	// 			DocumentType: to.Ptr("jpcz"),
	// 			ID: to.Ptr("uzp"),
	// 			LastUpdatedTime: to.Ptr("yyf"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_ServerlessResourceById_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_ServerlessResourceByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerlessRuntimesClient().ServerlessResourceByID(ctx, "rgopenapi", "_RD_R", "serverlessRuntimeName159", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaServerlessRuntimeResource = arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
	// 	Name: to.Ptr("byzccgftqjthb"),
	// 	Type: to.Ptr("due"),
	// 	ID: to.Ptr("vcdjzfgqjv"),
	// 	SystemData: &arminformaticadatamgmt.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
	// 		CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
	// 		LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 	},
	// 	Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
	// 		Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
	// 		AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
	// 			{
	// 				Key: to.Ptr("qcmc"),
	// 				Value: to.Ptr("unraxmnohdmvutt"),
	// 		}},
	// 		ApplicationType: to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
	// 		ComputeUnits: to.Ptr("bsctukmndvowse"),
	// 		ExecutionTimeout: to.Ptr("ruiougpypny"),
	// 		Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
	// 		ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
	// 		ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
	// 		ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
	// 			CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 			CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 		},
	// 		ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
	// 			NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
	// 				SubnetID: to.Ptr("s"),
	// 				VnetID: to.Ptr("uaqjvtubxccjs"),
	// 				VnetResourceGUID: to.Ptr("5328d299-1462-4be0-bef1-303a28e556a0"),
	// 			},
	// 		},
	// 		ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
	// 			{
	// 				Name: to.Ptr("korveuycuwhs"),
	// 				Value: to.Ptr("uyiuegxnkgp"),
	// 		}},
	// 		ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
	// 			UserContextToken: to.Ptr("oludf"),
	// 		},
	// 		SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_StartFailedServerlessRuntime_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_StartFailedServerlessRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServerlessRuntimesClient().StartFailedServerlessRuntime(ctx, "rgopenapi", "9M4", "-25-G_", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
