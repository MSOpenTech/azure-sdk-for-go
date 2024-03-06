//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RoleGetAllInDevice.json
func ExampleRolesClient_NewListByDataBoxEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRolesClient().NewListByDataBoxEdgeDevicePager("testedgedevice", "GroupForEdgeAutomation", nil)
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
		// page.RoleList = armdataboxedge.RoleList{
		// 	Value: []armdataboxedge.RoleClassification{
		// 		&armdataboxedge.IoTRole{
		// 			Name: to.Ptr("IoTRole1"),
		// 			Type: to.Ptr("dataBoxEdgeDevices/roles"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/IoTRole1"),
		// 			Kind: to.Ptr(armdataboxedge.RoleTypesIOT),
		// 			Properties: &armdataboxedge.IoTRoleProperties{
		// 				HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
		// 				IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
		// 					Authentication: &armdataboxedge.Authentication{
		// 						SymmetricKey: &armdataboxedge.SymmetricKey{
		// 						},
		// 					},
		// 					DeviceID: to.Ptr("iotdevice"),
		// 					IoTHostHub: to.Ptr("iothub.azure-devices.net"),
		// 				},
		// 				IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
		// 					Authentication: &armdataboxedge.Authentication{
		// 						SymmetricKey: &armdataboxedge.SymmetricKey{
		// 						},
		// 					},
		// 					DeviceID: to.Ptr("iotEdge"),
		// 					IoTHostHub: to.Ptr("iothub.azure-devices.net"),
		// 				},
		// 				RoleStatus: to.Ptr(armdataboxedge.RoleStatusEnabled),
		// 				ShareMappings: []*armdataboxedge.MountPointMap{
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RoleGet.json
func ExampleRolesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRolesClient().Get(ctx, "testedgedevice", "IoTRole1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataboxedge.RolesClientGetResponse{
	// 	                            RoleClassification: &armdataboxedge.IoTRole{
	// 		Name: to.Ptr("IoTRole1"),
	// 		Type: to.Ptr("dataBoxEdgeDevices/roles"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/IoTRole1"),
	// 		Kind: to.Ptr(armdataboxedge.RoleTypesIOT),
	// 		Properties: &armdataboxedge.IoTRoleProperties{
	// 			HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
	// 			IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
	// 				Authentication: &armdataboxedge.Authentication{
	// 					SymmetricKey: &armdataboxedge.SymmetricKey{
	// 					},
	// 				},
	// 				DeviceID: to.Ptr("iotdevice"),
	// 				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
	// 			},
	// 			IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
	// 				Authentication: &armdataboxedge.Authentication{
	// 					SymmetricKey: &armdataboxedge.SymmetricKey{
	// 					},
	// 				},
	// 				DeviceID: to.Ptr("iotEdge"),
	// 				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
	// 			},
	// 			RoleStatus: to.Ptr(armdataboxedge.RoleStatusEnabled),
	// 			ShareMappings: []*armdataboxedge.MountPointMap{
	// 			},
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RolePut.json
func ExampleRolesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRolesClient().BeginCreateOrUpdate(ctx, "testedgedevice", "IoTRole1", "GroupForEdgeAutomation", &armdataboxedge.IoTRole{
		Kind: to.Ptr(armdataboxedge.RoleTypesIOT),
		Properties: &armdataboxedge.IoTRoleProperties{
			HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
			IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
				Authentication: &armdataboxedge.Authentication{
					SymmetricKey: &armdataboxedge.SymmetricKey{
						ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
							EncryptionAlgorithm:      to.Ptr(armdataboxedge.EncryptionAlgorithmAES256),
							EncryptionCertThumbprint: to.Ptr("348586569999244"),
							Value:                    to.Ptr("Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotDevice;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>"),
						},
					},
				},
				DeviceID:   to.Ptr("iotdevice"),
				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
			},
			IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
				Authentication: &armdataboxedge.Authentication{
					SymmetricKey: &armdataboxedge.SymmetricKey{
						ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
							EncryptionAlgorithm:      to.Ptr(armdataboxedge.EncryptionAlgorithmAES256),
							EncryptionCertThumbprint: to.Ptr("1245475856069999244"),
							Value:                    to.Ptr("Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotEdge;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>"),
						},
					},
				},
				DeviceID:   to.Ptr("iotEdge"),
				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
			},
			RoleStatus:    to.Ptr(armdataboxedge.RoleStatusEnabled),
			ShareMappings: []*armdataboxedge.MountPointMap{},
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
	// res = armdataboxedge.RolesClientCreateOrUpdateResponse{
	// 	                            RoleClassification: &armdataboxedge.IoTRole{
	// 		Name: to.Ptr("IoTRole1"),
	// 		Type: to.Ptr("dataBoxEdgeDevices/roles"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/IoTRole1"),
	// 		Kind: to.Ptr(armdataboxedge.RoleTypesIOT),
	// 		Properties: &armdataboxedge.IoTRoleProperties{
	// 			HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
	// 			IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
	// 				Authentication: &armdataboxedge.Authentication{
	// 					SymmetricKey: &armdataboxedge.SymmetricKey{
	// 					},
	// 				},
	// 				DeviceID: to.Ptr("iotdevice"),
	// 				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
	// 			},
	// 			IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
	// 				Authentication: &armdataboxedge.Authentication{
	// 					SymmetricKey: &armdataboxedge.SymmetricKey{
	// 					},
	// 				},
	// 				DeviceID: to.Ptr("iotEdge"),
	// 				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
	// 			},
	// 			RoleStatus: to.Ptr(armdataboxedge.RoleStatusEnabled),
	// 			ShareMappings: []*armdataboxedge.MountPointMap{
	// 			},
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RoleDelete.json
func ExampleRolesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRolesClient().BeginDelete(ctx, "testedgedevice", "IoTRole1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
