//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/GetAllMetadataOData.json
func ExampleMetadataClient_NewListPager_getAllMetadataWithODataFilterOrderbySkipTop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetadataClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.MetadataClientListOptions{Filter: nil,
		Orderby: nil,
		Top:     nil,
		Skip:    nil,
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
		// page.MetadataList = armsecurityinsights.MetadataList{
		// 	Value: []*armsecurityinsights.MetadataModel{
		// 		{
		// 			Name: to.Ptr("metadataName1"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName1"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
		// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName1"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("metadataName2"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName2"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("f5160682-0e10-4e23-8fcf-df3df49c5522"),
		// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName2"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/GetAllMetadata.json
func ExampleMetadataClient_NewListPager_getAllMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetadataClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.MetadataClientListOptions{Filter: nil,
		Orderby: nil,
		Top:     nil,
		Skip:    nil,
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
		// page.MetadataList = armsecurityinsights.MetadataList{
		// 	Value: []*armsecurityinsights.MetadataModel{
		// 		{
		// 			Name: to.Ptr("metadataName1"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName1"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
		// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("metadataName2"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName2"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("f5160682-0e10-4e23-8fcf-df3df49c5522"),
		// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName2"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("metadataName3"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.Insights/workbooks/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName3"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("f593501d-ec01-4057-8146-a1de35c461ef"),
		// 				Kind: to.Ptr(armsecurityinsights.KindWorkbook),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.Insights/workbooks/workbookName"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/GetMetadata.json
func ExampleMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetadataClient().Get(ctx, "myRg", "myWorkspace", "metadataName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetadataModel = armsecurityinsights.MetadataModel{
	// 	Name: to.Ptr("metadataName"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
	// 	ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName"),
	// 	Properties: &armsecurityinsights.MetadataProperties{
	// 		Author: &armsecurityinsights.MetadataAuthor{
	// 			Name: to.Ptr("User Name"),
	// 			Email: to.Ptr("email@microsoft.com"),
	// 		},
	// 		Categories: &armsecurityinsights.MetadataCategories{
	// 			Domains: []*string{
	// 				to.Ptr("Application"),
	// 				to.Ptr("Security – Insider Threat")},
	// 				Verticals: []*string{
	// 					to.Ptr("Healthcare")},
	// 				},
	// 				ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
	// 				ContentSchemaVersion: to.Ptr("2.0"),
	// 				CustomVersion: to.Ptr("1.0"),
	// 				Dependencies: &armsecurityinsights.MetadataDependencies{
	// 					Criteria: []*armsecurityinsights.MetadataDependencies{
	// 						{
	// 							Criteria: []*armsecurityinsights.MetadataDependencies{
	// 								{
	// 									ContentID: to.Ptr("045d06d0-ee72-4794-aba4-cf5646e4c756"),
	// 									Kind: to.Ptr(armsecurityinsights.KindDataConnector),
	// 								},
	// 								{
	// 									ContentID: to.Ptr("dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d"),
	// 									Kind: to.Ptr(armsecurityinsights.KindDataConnector),
	// 								},
	// 								{
	// 									ContentID: to.Ptr("de4dca9b-eb37-47d6-a56f-b8b06b261593"),
	// 									Kind: to.Ptr(armsecurityinsights.KindDataConnector),
	// 									Version: to.Ptr("2.0"),
	// 							}},
	// 							Operator: to.Ptr(armsecurityinsights.OperatorOR),
	// 						},
	// 						{
	// 							ContentID: to.Ptr("31ee11cc-9989-4de8-b176-5e0ef5c4dbab"),
	// 							Kind: to.Ptr(armsecurityinsights.KindPlaybook),
	// 							Version: to.Ptr("1.0"),
	// 						},
	// 						{
	// 							ContentID: to.Ptr("21ba424a-9438-4444-953a-7059539a7a1b"),
	// 							Kind: to.Ptr(armsecurityinsights.KindParser),
	// 					}},
	// 					Operator: to.Ptr(armsecurityinsights.OperatorAND),
	// 				},
	// 				FirstPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t}()),
	// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
	// 				LastPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t}()),
	// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
	// 				PreviewImages: []*string{
	// 					to.Ptr("firstImage.png"),
	// 					to.Ptr("secondImage.jpeg")},
	// 					PreviewImagesDark: []*string{
	// 						to.Ptr("firstImageDark.png"),
	// 						to.Ptr("secondImageDark.jpeg")},
	// 						Providers: []*string{
	// 							to.Ptr("Amazon"),
	// 							to.Ptr("Microsoft")},
	// 							Source: &armsecurityinsights.MetadataSource{
	// 								Name: to.Ptr("Contoso Solution 1.0"),
	// 								Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
	// 								SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
	// 							},
	// 							Support: &armsecurityinsights.MetadataSupport{
	// 								Name: to.Ptr("Microsoft"),
	// 								Email: to.Ptr("support@microsoft.com"),
	// 								Link: to.Ptr("https://support.microsoft.com/"),
	// 								Tier: to.Ptr(armsecurityinsights.SupportTierPartner),
	// 							},
	// 							ThreatAnalysisTactics: []*string{
	// 								to.Ptr("reconnaissance"),
	// 								to.Ptr("commandandcontrol")},
	// 								ThreatAnalysisTechniques: []*string{
	// 									to.Ptr("T1548"),
	// 									to.Ptr("T1548.001")},
	// 									Version: to.Ptr("1.0.0.0"),
	// 								},
	// 							}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/DeleteMetadata.json
func ExampleMetadataClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMetadataClient().Delete(ctx, "myRg", "myWorkspace", "metadataName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/PutMetadata.json
func ExampleMetadataClient_Create_createUpdateFullMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetadataClient().Create(ctx, "myRg", "myWorkspace", "metadataName", armsecurityinsights.MetadataModel{
		Properties: &armsecurityinsights.MetadataProperties{
			Author: &armsecurityinsights.MetadataAuthor{
				Name:  to.Ptr("User Name"),
				Email: to.Ptr("email@microsoft.com"),
			},
			Categories: &armsecurityinsights.MetadataCategories{
				Domains: []*string{
					to.Ptr("Application"),
					to.Ptr("Security – Insider Threat")},
				Verticals: []*string{
					to.Ptr("Healthcare")},
			},
			ContentID:            to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
			ContentSchemaVersion: to.Ptr("2.0"),
			CustomVersion:        to.Ptr("1.0"),
			Dependencies: &armsecurityinsights.MetadataDependencies{
				Criteria: []*armsecurityinsights.MetadataDependencies{
					{
						Criteria: []*armsecurityinsights.MetadataDependencies{
							{
								Name:      to.Ptr("Microsoft Defender for Endpoint"),
								ContentID: to.Ptr("045d06d0-ee72-4794-aba4-cf5646e4c756"),
								Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
							},
							{
								ContentID: to.Ptr("dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d"),
								Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
							},
							{
								ContentID: to.Ptr("de4dca9b-eb37-47d6-a56f-b8b06b261593"),
								Kind:      to.Ptr(armsecurityinsights.KindDataConnector),
								Version:   to.Ptr("2.0"),
							}},
						Operator: to.Ptr(armsecurityinsights.OperatorOR),
					},
					{
						ContentID: to.Ptr("31ee11cc-9989-4de8-b176-5e0ef5c4dbab"),
						Kind:      to.Ptr(armsecurityinsights.KindPlaybook),
						Version:   to.Ptr("1.0"),
					},
					{
						ContentID: to.Ptr("21ba424a-9438-4444-953a-7059539a7a1b"),
						Kind:      to.Ptr(armsecurityinsights.KindParser),
					}},
				Operator: to.Ptr(armsecurityinsights.OperatorAND),
			},
			FirstPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
			Kind:             to.Ptr(armsecurityinsights.KindAnalyticsRule),
			LastPublishDate:  to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t }()),
			ParentID:         to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
			PreviewImages: []*string{
				to.Ptr("firstImage.png"),
				to.Ptr("secondImage.jpeg")},
			PreviewImagesDark: []*string{
				to.Ptr("firstImageDark.png"),
				to.Ptr("secondImageDark.jpeg")},
			Providers: []*string{
				to.Ptr("Amazon"),
				to.Ptr("Microsoft")},
			Source: &armsecurityinsights.MetadataSource{
				Name:     to.Ptr("Contoso Solution 1.0"),
				Kind:     to.Ptr(armsecurityinsights.SourceKindSolution),
				SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
			},
			Support: &armsecurityinsights.MetadataSupport{
				Name:  to.Ptr("Microsoft"),
				Email: to.Ptr("support@microsoft.com"),
				Link:  to.Ptr("https://support.microsoft.com/"),
				Tier:  to.Ptr(armsecurityinsights.SupportTierPartner),
			},
			ThreatAnalysisTactics: []*string{
				to.Ptr("reconnaissance"),
				to.Ptr("commandandcontrol")},
			ThreatAnalysisTechniques: []*string{
				to.Ptr("T1548"),
				to.Ptr("T1548.001")},
			Version: to.Ptr("1.0.0.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetadataModel = armsecurityinsights.MetadataModel{
	// 	Name: to.Ptr("metadataName"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
	// 	ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName"),
	// 	Properties: &armsecurityinsights.MetadataProperties{
	// 		Author: &armsecurityinsights.MetadataAuthor{
	// 			Name: to.Ptr("User Name"),
	// 			Email: to.Ptr("email@microsoft.com"),
	// 		},
	// 		Categories: &armsecurityinsights.MetadataCategories{
	// 			Domains: []*string{
	// 				to.Ptr("Application"),
	// 				to.Ptr("Security – Insider Threat")},
	// 				Verticals: []*string{
	// 					to.Ptr("Healthcare")},
	// 				},
	// 				ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
	// 				ContentSchemaVersion: to.Ptr("2.0"),
	// 				CustomVersion: to.Ptr("1.0"),
	// 				Dependencies: &armsecurityinsights.MetadataDependencies{
	// 					Criteria: []*armsecurityinsights.MetadataDependencies{
	// 						{
	// 							Criteria: []*armsecurityinsights.MetadataDependencies{
	// 								{
	// 									ContentID: to.Ptr("045d06d0-ee72-4794-aba4-cf5646e4c756"),
	// 									Kind: to.Ptr(armsecurityinsights.KindDataConnector),
	// 								},
	// 								{
	// 									ContentID: to.Ptr("dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d"),
	// 									Kind: to.Ptr(armsecurityinsights.KindDataConnector),
	// 								},
	// 								{
	// 									ContentID: to.Ptr("de4dca9b-eb37-47d6-a56f-b8b06b261593"),
	// 									Kind: to.Ptr(armsecurityinsights.KindDataConnector),
	// 									Version: to.Ptr("2.0"),
	// 							}},
	// 							Operator: to.Ptr(armsecurityinsights.OperatorOR),
	// 						},
	// 						{
	// 							ContentID: to.Ptr("31ee11cc-9989-4de8-b176-5e0ef5c4dbab"),
	// 							Kind: to.Ptr(armsecurityinsights.KindPlaybook),
	// 							Version: to.Ptr("1.0"),
	// 						},
	// 						{
	// 							ContentID: to.Ptr("21ba424a-9438-4444-953a-7059539a7a1b"),
	// 							Kind: to.Ptr(armsecurityinsights.KindParser),
	// 					}},
	// 					Operator: to.Ptr(armsecurityinsights.OperatorAND),
	// 				},
	// 				FirstPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t}()),
	// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
	// 				LastPublishDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-05-18"); return t}()),
	// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
	// 				PreviewImages: []*string{
	// 					to.Ptr("firstImage.png"),
	// 					to.Ptr("secondImage.jpeg")},
	// 					PreviewImagesDark: []*string{
	// 						to.Ptr("firstImageDark.png"),
	// 						to.Ptr("secondImageDark.jpeg")},
	// 						Providers: []*string{
	// 							to.Ptr("Amazon"),
	// 							to.Ptr("Microsoft")},
	// 							Source: &armsecurityinsights.MetadataSource{
	// 								Name: to.Ptr("Contoso Solution 1.0"),
	// 								Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
	// 								SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
	// 							},
	// 							Support: &armsecurityinsights.MetadataSupport{
	// 								Name: to.Ptr("Microsoft"),
	// 								Email: to.Ptr("support@microsoft.com"),
	// 								Link: to.Ptr("https://support.microsoft.com/"),
	// 								Tier: to.Ptr(armsecurityinsights.SupportTierPartner),
	// 							},
	// 							ThreatAnalysisTactics: []*string{
	// 								to.Ptr("reconnaissance"),
	// 								to.Ptr("commandandcontrol")},
	// 								ThreatAnalysisTechniques: []*string{
	// 									to.Ptr("T1548"),
	// 									to.Ptr("T1548.001")},
	// 									Version: to.Ptr("1.0.0.0"),
	// 								},
	// 							}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/PutMetadataMinimal.json
func ExampleMetadataClient_Create_createUpdateMinimalMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetadataClient().Create(ctx, "myRg", "myWorkspace", "metadataName", armsecurityinsights.MetadataModel{
		Properties: &armsecurityinsights.MetadataProperties{
			ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
			Kind:      to.Ptr(armsecurityinsights.KindAnalyticsRule),
			ParentID:  to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetadataModel = armsecurityinsights.MetadataModel{
	// 	Name: to.Ptr("metadataName"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
	// 	ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName"),
	// 	Properties: &armsecurityinsights.MetadataProperties{
	// 		Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
	// 		ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/PatchMetadata.json
func ExampleMetadataClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetadataClient().Update(ctx, "myRg", "myWorkspace", "metadataName", armsecurityinsights.MetadataPatch{
		Properties: &armsecurityinsights.MetadataPropertiesPatch{
			Author: &armsecurityinsights.MetadataAuthor{
				Name:  to.Ptr("User Name"),
				Email: to.Ptr("email@microsoft.com"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetadataModel = armsecurityinsights.MetadataModel{
	// 	Name: to.Ptr("metadataName"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
	// 	ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName"),
	// 	Properties: &armsecurityinsights.MetadataProperties{
	// 		Author: &armsecurityinsights.MetadataAuthor{
	// 			Name: to.Ptr("User Name"),
	// 			Email: to.Ptr("email@microsoft.com"),
	// 		},
	// 		ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
	// 		Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
	// 		ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
	// 	},
	// }
}
