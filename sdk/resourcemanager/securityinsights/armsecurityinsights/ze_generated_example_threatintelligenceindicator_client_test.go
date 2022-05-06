//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_CreateIndicator() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateIndicator(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsecurityinsights.ThreatIntelligenceIndicatorModel{
			Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				Description:        to.Ptr("<description>"),
				Confidence:         to.Ptr[int32](78),
				CreatedByRef:       to.Ptr("<created-by-ref>"),
				DisplayName:        to.Ptr("<display-name>"),
				ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{},
				GranularMarkings:   []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{},
				KillChainPhases:    []*armsecurityinsights.ThreatIntelligenceKillChainPhase{},
				Labels:             []*string{},
				Modified:           to.Ptr("<modified>"),
				Pattern:            to.Ptr("<pattern>"),
				PatternType:        to.Ptr("<pattern-type>"),
				Revoked:            to.Ptr(false),
				Source:             to.Ptr("<source>"),
				ThreatIntelligenceTags: []*string{
					to.Ptr("new schema")},
				ThreatTypes: []*string{
					to.Ptr("compromised")},
				ValidFrom:  to.Ptr("<valid-from>"),
				ValidUntil: to.Ptr("<valid-until>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/GetThreatIntelligenceById.json
func ExampleThreatIntelligenceIndicatorClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/UpdateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		armsecurityinsights.ThreatIntelligenceIndicatorModel{
			Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				Description:        to.Ptr("<description>"),
				Confidence:         to.Ptr[int32](78),
				CreatedByRef:       to.Ptr("<created-by-ref>"),
				DisplayName:        to.Ptr("<display-name>"),
				ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{},
				GranularMarkings:   []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{},
				KillChainPhases:    []*armsecurityinsights.ThreatIntelligenceKillChainPhase{},
				Labels:             []*string{},
				Modified:           to.Ptr("<modified>"),
				Pattern:            to.Ptr("<pattern>"),
				PatternType:        to.Ptr("<pattern-type>"),
				Revoked:            to.Ptr(false),
				Source:             to.Ptr("<source>"),
				ThreatIntelligenceTags: []*string{
					to.Ptr("new schema")},
				ThreatTypes: []*string{
					to.Ptr("compromised")},
				ValidFrom:  to.Ptr("<valid-from>"),
				ValidUntil: to.Ptr("<valid-until>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/DeleteThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/QueryThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_NewQueryIndicatorsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewQueryIndicatorsPager("<resource-group-name>",
		"<workspace-name>",
		armsecurityinsights.ThreatIntelligenceFilteringCriteria{
			MaxConfidence: to.Ptr[int32](80),
			MaxValidUntil: to.Ptr("<max-valid-until>"),
			MinConfidence: to.Ptr[int32](25),
			MinValidUntil: to.Ptr("<min-valid-until>"),
			PageSize:      to.Ptr[int32](100),
			SortBy: []*armsecurityinsights.ThreatIntelligenceSortingCriteria{
				{
					ItemKey:   to.Ptr("<item-key>"),
					SortOrder: to.Ptr(armsecurityinsights.ThreatIntelligenceSortingCriteriaEnumDescending),
				}},
			Sources: []*string{
				to.Ptr("Azure Sentinel")},
		},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/AppendTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_AppendTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.AppendTags(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		armsecurityinsights.ThreatIntelligenceAppendTags{
			ThreatIntelligenceTags: []*string{
				to.Ptr("tag1"),
				to.Ptr("tag2")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_ReplaceTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ReplaceTags(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		armsecurityinsights.ThreatIntelligenceIndicatorModel{
			Etag: to.Ptr("<etag>"),
			Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				ThreatIntelligenceTags: []*string{
					to.Ptr("patching tags")},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
