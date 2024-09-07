//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/DeleteScriptAction.json
func ExampleScriptActionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScriptActionsClient().Delete(ctx, "rg1", "cluster1", "scriptName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetLinuxHadoopScriptAction.json
func ExampleScriptActionsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScriptActionsClient().NewListByClusterPager("rg1", "cluster1", nil)
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
		// page.ScriptActionsList = armhdinsight.ScriptActionsList{
		// 	Value: []*armhdinsight.RuntimeScriptActionDetail{
		// 		{
		// 			Name: to.Ptr("app-Install"),
		// 			ApplicationName: to.Ptr("app"),
		// 			Roles: []*string{
		// 				to.Ptr("edgenode")},
		// 				URI: to.Ptr("https://app.com/azure/app_install.sh"),
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetScriptActionById.json
func ExampleScriptActionsClient_GetExecutionDetail() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptActionsClient().GetExecutionDetail(ctx, "rg1", "cluster1", "391145124054712", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuntimeScriptActionDetail = armhdinsight.RuntimeScriptActionDetail{
	// 	Name: to.Ptr("Test"),
	// 	ApplicationName: to.Ptr("app1"),
	// 	Roles: []*string{
	// 		to.Ptr("headnode"),
	// 		to.Ptr("workernode")},
	// 		URI: to.Ptr("http://testurl.com/install.ssh"),
	// 		DebugInformation: to.Ptr(""),
	// 		EndTime: to.Ptr("2017-03-22T21:34:39.293"),
	// 		ExecutionSummary: []*armhdinsight.ScriptActionExecutionSummary{
	// 		},
	// 		Operation: to.Ptr("PostClusterCreateScriptActionRequest"),
	// 		ScriptExecutionID: to.Ptr[int64](391145124054712),
	// 		StartTime: to.Ptr("2017-03-22T21:34:39.293"),
	// 		Status: to.Ptr("ValidationFailed"),
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetScriptExecutionAsyncOperationStatus.json
func ExampleScriptActionsClient_GetExecutionAsyncOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptActionsClient().GetExecutionAsyncOperationStatus(ctx, "rg1", "cluster1", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AsyncOperationResult = armhdinsight.AsyncOperationResult{
	// 	Status: to.Ptr(armhdinsight.AsyncOperationStateInProgress),
	// }
}
