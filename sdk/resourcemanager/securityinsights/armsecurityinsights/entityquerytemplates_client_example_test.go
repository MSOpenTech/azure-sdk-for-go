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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/entityQueryTemplates/GetEntityQueryTemplates.json
func ExampleEntityQueryTemplatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEntityQueryTemplatesClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.EntityQueryTemplatesClientListOptions{Kind: to.Ptr(armsecurityinsights.Enum15Activity)})
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
		// page.EntityQueryTemplateList = armsecurityinsights.EntityQueryTemplateList{
		// 	Value: []armsecurityinsights.EntityQueryTemplateClassification{
		// 		&armsecurityinsights.ActivityEntityQueryTemplate{
		// 			Name: to.Ptr("37ca3555-c135-4a73-a65e-9c1d00323f5d"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/entityQueryTemplates"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entityQueryTemplates/37ca3555-c135-4a73-a65e-9c1d00323f5d"),
		// 			Kind: to.Ptr(armsecurityinsights.EntityQueryTemplateKindActivity),
		// 			Properties: &armsecurityinsights.ActivityEntityQueryTemplateProperties{
		// 				Description: to.Ptr("Account deleted on host"),
		// 				Content: to.Ptr("On '{{Computer}}' the account '{{TargetAccount}}' was deleted by '{{AddedBy}}'"),
		// 				DataTypes: []*armsecurityinsights.DataTypeDefinitions{
		// 					{
		// 						DataType: to.Ptr("AuditLogs"),
		// 					},
		// 					{
		// 						DataType: to.Ptr("SecurityEvent"),
		// 				}},
		// 				EntitiesFilter: map[string][]*string{
		// 					"Host_OsFamily": []*string{
		// 						to.Ptr("Windows")},
		// 					},
		// 					InputEntityType: to.Ptr(armsecurityinsights.EntityTypeHost),
		// 					QueryDefinitions: &armsecurityinsights.ActivityEntityQueryTemplatePropertiesQueryDefinitions{
		// 						Query: to.Ptr("let GetAccountActions = (v_Host_Name:string, v_Host_NTDomain:string, v_Host_DnsDomain:string, v_Host_AzureID:string, v_Host_OMSAgentID:string){\nSecurityEvent\n| where EventID in (4725, 4726, 4767, 4720, 4722, 4723, 4724)\n// parsing for Host to handle variety of conventions coming from data\n| extend Host_HostName = case(\nComputer has '@', tostring(split(Computer, '@')[0]),\nComputer has '\\\\', tostring(split(Computer, '\\\\')[1]),\nComputer has '.', tostring(split(Computer, '.')[0]),\nComputer\n)\n| extend Host_NTDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', tostring(split(Computer, '.')[-2]), \nComputer\n)\n| extend Host_DnsDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', strcat_array(array_slice(split(Computer,'.'),-2,-1),'.'), \nComputer\n)\n| where (Host_HostName =~ v_Host_Name and Host_NTDomain =~ v_Host_NTDomain) \nor (Host_HostName =~ v_Host_Name and Host_DnsDomain =~ v_Host_DnsDomain) \nor v_Host_AzureID =~ _ResourceId \nor v_Host_OMSAgentID == SourceComputerId\n| project TimeGenerated, EventID, Activity, Computer, TargetAccount, TargetUserName, TargetDomainName, TargetSid, SubjectUserName, SubjectUserSid, _ResourceId, SourceComputerId\n| extend AddedBy = SubjectUserName\n// Future support for Activities\n| extend timestamp = TimeGenerated, HostCustomEntity = Computer, AccountCustomEntity = TargetAccount\n};\nGetAccountActions('{{Host_HostName}}', '{{Host_NTDomain}}', '{{Host_DnsDomain}}', '{{Host_AzureID}}', '{{Host_OMSAgentID}}')\n \n| where EventID == 4726 "),
		// 					},
		// 					RequiredInputFieldsSets: [][]*string{
		// 						[]*string{
		// 							to.Ptr("Host_HostName"),
		// 							to.Ptr("Host_NTDomain")},
		// 							[]*string{
		// 								to.Ptr("Host_HostName"),
		// 								to.Ptr("Host_DnsDomain")},
		// 								[]*string{
		// 									to.Ptr("Host_AzureID")},
		// 									[]*string{
		// 										to.Ptr("Host_OMSAgentID")}},
		// 										Title: to.Ptr("An account was deleted on this host"),
		// 									},
		// 								},
		// 								&armsecurityinsights.ActivityEntityQueryTemplate{
		// 									Name: to.Ptr("97a1d515-abf2-4231-9a35-985f9de0bb91"),
		// 									Type: to.Ptr("Microsoft.SecurityInsights/entityQueryTemplates"),
		// 									ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entityQueryTemplates/97a1d515-abf2-4231-9a35-985f9de0bb91"),
		// 									Kind: to.Ptr(armsecurityinsights.EntityQueryTemplateKindActivity),
		// 									Properties: &armsecurityinsights.ActivityEntityQueryTemplateProperties{
		// 										Description: to.Ptr("Account deleted on host"),
		// 										Content: to.Ptr("On '{{Computer}}' the account '{{TargetAccount}}' was deleted by '{{AddedBy}}'"),
		// 										DataTypes: []*armsecurityinsights.DataTypeDefinitions{
		// 											{
		// 												DataType: to.Ptr("AuditLogs"),
		// 											},
		// 											{
		// 												DataType: to.Ptr("SecurityEvent"),
		// 										}},
		// 										EntitiesFilter: map[string][]*string{
		// 											"Host_OsFamily": []*string{
		// 												to.Ptr("Windows")},
		// 											},
		// 											InputEntityType: to.Ptr(armsecurityinsights.EntityTypeHost),
		// 											QueryDefinitions: &armsecurityinsights.ActivityEntityQueryTemplatePropertiesQueryDefinitions{
		// 												Query: to.Ptr("let GetAccountActions = (v_Host_Name:string, v_Host_NTDomain:string, v_Host_DnsDomain:string, v_Host_AzureID:string, v_Host_OMSAgentID:string){\nSecurityEvent\n| where EventID in (4725, 4726, 4767, 4720, 4722, 4723, 4724)\n// parsing for Host to handle variety of conventions coming from data\n| extend Host_HostName = case(\nComputer has '@', tostring(split(Computer, '@')[0]),\nComputer has '\\\\', tostring(split(Computer, '\\\\')[1]),\nComputer has '.', tostring(split(Computer, '.')[0]),\nComputer\n)\n| extend Host_NTDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', tostring(split(Computer, '.')[-2]), \nComputer\n)\n| extend Host_DnsDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', strcat_array(array_slice(split(Computer,'.'),-2,-1),'.'), \nComputer\n)\n| where (Host_HostName =~ v_Host_Name and Host_NTDomain =~ v_Host_NTDomain) \nor (Host_HostName =~ v_Host_Name and Host_DnsDomain =~ v_Host_DnsDomain) \nor v_Host_AzureID =~ _ResourceId \nor v_Host_OMSAgentID == SourceComputerId\n| project TimeGenerated, EventID, Activity, Computer, TargetAccount, TargetUserName, TargetDomainName, TargetSid, SubjectUserName, SubjectUserSid, _ResourceId, SourceComputerId\n| extend AddedBy = SubjectUserName\n// Future support for Activities\n| extend timestamp = TimeGenerated, HostCustomEntity = Computer, AccountCustomEntity = TargetAccount\n};\nGetAccountActions('{{Host_HostName}}', '{{Host_NTDomain}}', '{{Host_DnsDomain}}', '{{Host_AzureID}}', '{{Host_OMSAgentID}}')\n \n| where EventID == 4726 "),
		// 											},
		// 											RequiredInputFieldsSets: [][]*string{
		// 												[]*string{
		// 													to.Ptr("Host_HostName"),
		// 													to.Ptr("Host_NTDomain")},
		// 													[]*string{
		// 														to.Ptr("Host_HostName"),
		// 														to.Ptr("Host_DnsDomain")},
		// 														[]*string{
		// 															to.Ptr("Host_AzureID")},
		// 															[]*string{
		// 																to.Ptr("Host_OMSAgentID")}},
		// 																Title: to.Ptr("An account was deleted on this host"),
		// 															},
		// 													}},
		// 												}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/entityQueryTemplates/GetActivityEntityQueryTemplateById.json
func ExampleEntityQueryTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntityQueryTemplatesClient().Get(ctx, "myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntityQueryTemplatesClientGetResponse{
	// 	                            EntityQueryTemplateClassification: &armsecurityinsights.ActivityEntityQueryTemplate{
	// 		Name: to.Ptr("07da3cc8-c8ad-4710-a44e-334cdcb7882b"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entityQueryTemplate"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entityQueryTemplates/07da3cc8-c8ad-4710-a44e-334cdcb7882b"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityQueryTemplateKindActivity),
	// 		Properties: &armsecurityinsights.ActivityEntityQueryTemplateProperties{
	// 			Description: to.Ptr("Account deleted on host"),
	// 			Content: to.Ptr("On '{{Computer}}' the account '{{TargetAccount}}' was deleted by '{{AddedBy}}'"),
	// 			DataTypes: []*armsecurityinsights.DataTypeDefinitions{
	// 				{
	// 					DataType: to.Ptr("AuditLogs"),
	// 				},
	// 				{
	// 					DataType: to.Ptr("SecurityEvent"),
	// 			}},
	// 			EntitiesFilter: map[string][]*string{
	// 				"Host_OsFamily": []*string{
	// 					to.Ptr("Windows")},
	// 				},
	// 				InputEntityType: to.Ptr(armsecurityinsights.EntityTypeHost),
	// 				QueryDefinitions: &armsecurityinsights.ActivityEntityQueryTemplatePropertiesQueryDefinitions{
	// 					Query: to.Ptr("let GetAccountActions = (v_Host_Name:string, v_Host_NTDomain:string, v_Host_DnsDomain:string, v_Host_AzureID:string, v_Host_OMSAgentID:string){\nSecurityEvent\n| where EventID in (4725, 4726, 4767, 4720, 4722, 4723, 4724)\n// parsing for Host to handle variety of conventions coming from data\n| extend Host_HostName = case(\nComputer has '@', tostring(split(Computer, '@')[0]),\nComputer has '\\\\', tostring(split(Computer, '\\\\')[1]),\nComputer has '.', tostring(split(Computer, '.')[0]),\nComputer\n)\n| extend Host_NTDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', tostring(split(Computer, '.')[-2]), \nComputer\n)\n| extend Host_DnsDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', strcat_array(array_slice(split(Computer,'.'),-2,-1),'.'), \nComputer\n)\n| where (Host_HostName =~ v_Host_Name and Host_NTDomain =~ v_Host_NTDomain) \nor (Host_HostName =~ v_Host_Name and Host_DnsDomain =~ v_Host_DnsDomain) \nor v_Host_AzureID =~ _ResourceId \nor v_Host_OMSAgentID == SourceComputerId\n| project TimeGenerated, EventID, Activity, Computer, TargetAccount, TargetUserName, TargetDomainName, TargetSid, SubjectUserName, SubjectUserSid, _ResourceId, SourceComputerId\n| extend AddedBy = SubjectUserName\n// Future support for Activities\n| extend timestamp = TimeGenerated, HostCustomEntity = Computer, AccountCustomEntity = TargetAccount\n};\nGetAccountActions('{{Host_HostName}}', '{{Host_NTDomain}}', '{{Host_DnsDomain}}', '{{Host_AzureID}}', '{{Host_OMSAgentID}}')\n \n| where EventID == 4726 "),
	// 				},
	// 				RequiredInputFieldsSets: [][]*string{
	// 					[]*string{
	// 						to.Ptr("Host_HostName"),
	// 						to.Ptr("Host_NTDomain")},
	// 						[]*string{
	// 							to.Ptr("Host_HostName"),
	// 							to.Ptr("Host_DnsDomain")},
	// 							[]*string{
	// 								to.Ptr("Host_AzureID")},
	// 								[]*string{
	// 									to.Ptr("Host_OMSAgentID")}},
	// 									Title: to.Ptr("An account was deleted on this host"),
	// 								},
	// 							},
	// 							                        }
}
