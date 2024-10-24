// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armiotoperations_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
	"log"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations/TempTypeSpecFiles/IoTOperations.Management/examples/2024-11-01/BrokerListener_CreateOrUpdate_Complex.json
func ExampleBrokerListenerClient_BeginCreateOrUpdate_brokerListenerCreateOrUpdateComplex() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerListenerClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.BrokerListenerResource{
		Properties: &armiotoperations.BrokerListenerProperties{
			ServiceType: to.Ptr(armiotoperations.ServiceTypeLoadBalancer),
			Ports: []*armiotoperations.ListenerPort{
				{
					Port:              to.Ptr[int32](8080),
					AuthenticationRef: to.Ptr("example-authentication"),
					Protocol:          to.Ptr(armiotoperations.BrokerProtocolTypeWebSockets),
				},
				{
					Port:              to.Ptr[int32](8443),
					AuthenticationRef: to.Ptr("example-authentication"),
					Protocol:          to.Ptr(armiotoperations.BrokerProtocolTypeWebSockets),
					TLS: &armiotoperations.TLSCertMethod{
						Mode: to.Ptr(armiotoperations.TLSCertMethodModeAutomatic),
						CertManagerCertificateSpec: &armiotoperations.CertManagerCertificateSpec{
							IssuerRef: &armiotoperations.CertManagerIssuerRef{
								Group: to.Ptr("jtmuladdkpasfpoyvewekmiy"),
								Name:  to.Ptr("example-issuer"),
								Kind:  to.Ptr(armiotoperations.CertManagerIssuerKindIssuer),
							},
						},
					},
				},
				{
					Port:              to.Ptr[int32](1883),
					AuthenticationRef: to.Ptr("example-authentication"),
				},
				{
					Port:              to.Ptr[int32](8883),
					AuthenticationRef: to.Ptr("example-authentication"),
					TLS: &armiotoperations.TLSCertMethod{
						Mode: to.Ptr(armiotoperations.TLSCertMethodModeManual),
						Manual: &armiotoperations.X509ManualCertificate{
							SecretRef: to.Ptr("example-secret"),
						},
					},
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
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
	// res = armiotoperations.BrokerListenerClientCreateOrUpdateResponse{
	// 	BrokerListenerResource: &armiotoperations.BrokerListenerResource{
	// 		Properties: &armiotoperations.BrokerListenerProperties{
	// 			ServiceName: to.Ptr("tpfiszlapdpxktx"),
	// 			ServiceType: to.Ptr(armiotoperations.ServiceTypeLoadBalancer),
	// 			Ports: []*armiotoperations.ListenerPort{
	// 				{
	// 					Port: to.Ptr[int32](8080),
	// 					AuthenticationRef: to.Ptr("example-authentication"),
	// 					Protocol: to.Ptr(armiotoperations.BrokerProtocolTypeWebSockets),
	// 				},
	// 				{
	// 					Port: to.Ptr[int32](8443),
	// 					AuthenticationRef: to.Ptr("example-authentication"),
	// 					Protocol: to.Ptr(armiotoperations.BrokerProtocolTypeWebSockets),
	// 					TLS: &armiotoperations.TLSCertMethod{
	// 						Mode: to.Ptr(armiotoperations.TLSCertMethodModeAutomatic),
	// 						CertManagerCertificateSpec: &armiotoperations.CertManagerCertificateSpec{
	// 							IssuerRef: &armiotoperations.CertManagerIssuerRef{
	// 								Group: to.Ptr("jtmuladdkpasfpoyvewekmiy"),
	// 								Name: to.Ptr("example-issuer"),
	// 								Kind: to.Ptr(armiotoperations.CertManagerIssuerKindIssuer),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Port: to.Ptr[int32](1883),
	// 					AuthenticationRef: to.Ptr("example-authentication"),
	// 				},
	// 				{
	// 					Port: to.Ptr[int32](8883),
	// 					AuthenticationRef: to.Ptr("example-authentication"),
	// 					TLS: &armiotoperations.TLSCertMethod{
	// 						Mode: to.Ptr(armiotoperations.TLSCertMethodModeManual),
	// 						Manual: &armiotoperations.X509ManualCertificate{
	// 							SecretRef: to.Ptr("example-secret"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/listeners/resource-name123"),
	// 		Name: to.Ptr("hoqjaachratt"),
	// 		Type: to.Ptr("hizbknwegcdaeh"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations/TempTypeSpecFiles/IoTOperations.Management/examples/2024-11-01/BrokerListener_CreateOrUpdate_MaximumSet_Gen.json
func ExampleBrokerListenerClient_BeginCreateOrUpdate_brokerListenerCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerListenerClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.BrokerListenerResource{
		Properties: &armiotoperations.BrokerListenerProperties{
			ServiceName: to.Ptr("tpfiszlapdpxktx"),
			Ports: []*armiotoperations.ListenerPort{
				{
					AuthenticationRef: to.Ptr("tjvdroaqqy"),
					AuthorizationRef:  to.Ptr("inxhvxnwswyrvt"),
					NodePort:          to.Ptr[int32](7281),
					Port:              to.Ptr[int32](1268),
					Protocol:          to.Ptr(armiotoperations.BrokerProtocolTypeMqtt),
					TLS: &armiotoperations.TLSCertMethod{
						Mode: to.Ptr(armiotoperations.TLSCertMethodModeAutomatic),
						CertManagerCertificateSpec: &armiotoperations.CertManagerCertificateSpec{
							Duration:    to.Ptr("qmpeffoksron"),
							SecretName:  to.Ptr("oagi"),
							RenewBefore: to.Ptr("hutno"),
							IssuerRef: &armiotoperations.CertManagerIssuerRef{
								Group: to.Ptr("jtmuladdkpasfpoyvewekmiy"),
								Kind:  to.Ptr(armiotoperations.CertManagerIssuerKindIssuer),
								Name:  to.Ptr("ocwoqpgucvjrsuudtjhb"),
							},
							PrivateKey: &armiotoperations.CertManagerPrivateKey{
								Algorithm:      to.Ptr(armiotoperations.PrivateKeyAlgorithmEc256),
								RotationPolicy: to.Ptr(armiotoperations.PrivateKeyRotationPolicyAlways),
							},
							San: &armiotoperations.SanForCert{
								DNS: []*string{
									to.Ptr("xhvmhrrhgfsapocjeebqtnzarlj"),
								},
								IP: []*string{
									to.Ptr("zbgugfzcgsmegevzktsnibyuyp"),
								},
							},
						},
						Manual: &armiotoperations.X509ManualCertificate{
							SecretRef: to.Ptr("secret-name"),
						},
					},
				},
			},
			ServiceType: to.Ptr(armiotoperations.ServiceTypeClusterIP),
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
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
	// res = armiotoperations.BrokerListenerClientCreateOrUpdateResponse{
	// 	BrokerListenerResource: &armiotoperations.BrokerListenerResource{
	// 		Properties: &armiotoperations.BrokerListenerProperties{
	// 			ServiceName: to.Ptr("tpfiszlapdpxktx"),
	// 			Ports: []*armiotoperations.ListenerPort{
	// 				{
	// 					AuthenticationRef: to.Ptr("tjvdroaqqy"),
	// 					AuthorizationRef: to.Ptr("inxhvxnwswyrvt"),
	// 					NodePort: to.Ptr[int32](7281),
	// 					Port: to.Ptr[int32](1268),
	// 					Protocol: to.Ptr(armiotoperations.BrokerProtocolTypeMqtt),
	// 					TLS: &armiotoperations.TLSCertMethod{
	// 						Mode: to.Ptr(armiotoperations.TLSCertMethodModeAutomatic),
	// 						CertManagerCertificateSpec: &armiotoperations.CertManagerCertificateSpec{
	// 							Duration: to.Ptr("qmpeffoksron"),
	// 							SecretName: to.Ptr("oagi"),
	// 							RenewBefore: to.Ptr("hutno"),
	// 							IssuerRef: &armiotoperations.CertManagerIssuerRef{
	// 								Group: to.Ptr("jtmuladdkpasfpoyvewekmiy"),
	// 								Kind: to.Ptr(armiotoperations.CertManagerIssuerKindIssuer),
	// 								Name: to.Ptr("ocwoqpgucvjrsuudtjhb"),
	// 							},
	// 							PrivateKey: &armiotoperations.CertManagerPrivateKey{
	// 								Algorithm: to.Ptr(armiotoperations.PrivateKeyAlgorithmEc256),
	// 								RotationPolicy: to.Ptr(armiotoperations.PrivateKeyRotationPolicyAlways),
	// 							},
	// 							San: &armiotoperations.SanForCert{
	// 								DNS: []*string{
	// 									to.Ptr("xhvmhrrhgfsapocjeebqtnzarlj"),
	// 								},
	// 								IP: []*string{
	// 									to.Ptr("zbgugfzcgsmegevzktsnibyuyp"),
	// 								},
	// 							},
	// 						},
	// 						Manual: &armiotoperations.X509ManualCertificate{
	// 							SecretRef: to.Ptr("secret-name"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ServiceType: to.Ptr(armiotoperations.ServiceTypeClusterIP),
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/listeners/resource-name123"),
	// 		Name: to.Ptr("hoqjaachratt"),
	// 		Type: to.Ptr("hizbknwegcdaeh"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations/TempTypeSpecFiles/IoTOperations.Management/examples/2024-11-01/BrokerListener_CreateOrUpdate_Simple.json
func ExampleBrokerListenerClient_BeginCreateOrUpdate_brokerListenerCreateOrUpdateSimple() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerListenerClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", armiotoperations.BrokerListenerResource{
		Properties: &armiotoperations.BrokerListenerProperties{
			Ports: []*armiotoperations.ListenerPort{
				{
					Port: to.Ptr[int32](1883),
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
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
	// res = armiotoperations.BrokerListenerClientCreateOrUpdateResponse{
	// 	BrokerListenerResource: &armiotoperations.BrokerListenerResource{
	// 		Properties: &armiotoperations.BrokerListenerProperties{
	// 			ServiceName: to.Ptr("tpfiszlapdpxktx"),
	// 			ServiceType: to.Ptr(armiotoperations.ServiceTypeLoadBalancer),
	// 			Ports: []*armiotoperations.ListenerPort{
	// 				{
	// 					Port: to.Ptr[int32](1883),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/listeners/resource-name123"),
	// 		Name: to.Ptr("hoqjaachratt"),
	// 		Type: to.Ptr("hizbknwegcdaeh"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations/TempTypeSpecFiles/IoTOperations.Management/examples/2024-11-01/BrokerListener_Delete_MaximumSet_Gen.json
func ExampleBrokerListenerClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBrokerListenerClient().BeginDelete(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations/TempTypeSpecFiles/IoTOperations.Management/examples/2024-11-01/BrokerListener_Get_MaximumSet_Gen.json
func ExampleBrokerListenerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBrokerListenerClient().Get(ctx, "rgiotoperations", "resource-name123", "resource-name123", "resource-name123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.BrokerListenerClientGetResponse{
	// 	BrokerListenerResource: &armiotoperations.BrokerListenerResource{
	// 		Properties: &armiotoperations.BrokerListenerProperties{
	// 			ServiceName: to.Ptr("tpfiszlapdpxktx"),
	// 			Ports: []*armiotoperations.ListenerPort{
	// 				{
	// 					AuthenticationRef: to.Ptr("tjvdroaqqy"),
	// 					AuthorizationRef: to.Ptr("inxhvxnwswyrvt"),
	// 					NodePort: to.Ptr[int32](7281),
	// 					Port: to.Ptr[int32](1268),
	// 					Protocol: to.Ptr(armiotoperations.BrokerProtocolTypeMqtt),
	// 					TLS: &armiotoperations.TLSCertMethod{
	// 						Mode: to.Ptr(armiotoperations.TLSCertMethodModeAutomatic),
	// 						CertManagerCertificateSpec: &armiotoperations.CertManagerCertificateSpec{
	// 							Duration: to.Ptr("qmpeffoksron"),
	// 							SecretName: to.Ptr("oagi"),
	// 							RenewBefore: to.Ptr("hutno"),
	// 							IssuerRef: &armiotoperations.CertManagerIssuerRef{
	// 								Group: to.Ptr("jtmuladdkpasfpoyvewekmiy"),
	// 								Kind: to.Ptr(armiotoperations.CertManagerIssuerKindIssuer),
	// 								Name: to.Ptr("ocwoqpgucvjrsuudtjhb"),
	// 							},
	// 							PrivateKey: &armiotoperations.CertManagerPrivateKey{
	// 								Algorithm: to.Ptr(armiotoperations.PrivateKeyAlgorithmEc256),
	// 								RotationPolicy: to.Ptr(armiotoperations.PrivateKeyRotationPolicyAlways),
	// 							},
	// 							San: &armiotoperations.SanForCert{
	// 								DNS: []*string{
	// 									to.Ptr("xhvmhrrhgfsapocjeebqtnzarlj"),
	// 								},
	// 								IP: []*string{
	// 									to.Ptr("zbgugfzcgsmegevzktsnibyuyp"),
	// 								},
	// 							},
	// 						},
	// 						Manual: &armiotoperations.X509ManualCertificate{
	// 							SecretRef: to.Ptr("secret-name"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ServiceType: to.Ptr(armiotoperations.ServiceTypeClusterIP),
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/listeners/resource-name123"),
	// 		Name: to.Ptr("hoqjaachratt"),
	// 		Type: to.Ptr("hizbknwegcdaeh"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations/TempTypeSpecFiles/IoTOperations.Management/examples/2024-11-01/BrokerListener_ListByResourceGroup_MaximumSet_Gen.json
func ExampleBrokerListenerClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBrokerListenerClient().NewListByResourceGroupPager("rgiotoperations", "resource-name123", "resource-name123", nil)
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
		// page = armiotoperations.BrokerListenerClientListByResourceGroupResponse{
		// 	BrokerListenerResourceListResult: armiotoperations.BrokerListenerResourceListResult{
		// 		Value: []*armiotoperations.BrokerListenerResource{
		// 			{
		// 				Properties: &armiotoperations.BrokerListenerProperties{
		// 					ServiceName: to.Ptr("tpfiszlapdpxktx"),
		// 					Ports: []*armiotoperations.ListenerPort{
		// 						{
		// 							AuthenticationRef: to.Ptr("tjvdroaqqy"),
		// 							AuthorizationRef: to.Ptr("inxhvxnwswyrvt"),
		// 							NodePort: to.Ptr[int32](7281),
		// 							Port: to.Ptr[int32](1268),
		// 							Protocol: to.Ptr(armiotoperations.BrokerProtocolTypeMqtt),
		// 							TLS: &armiotoperations.TLSCertMethod{
		// 								Mode: to.Ptr(armiotoperations.TLSCertMethodModeAutomatic),
		// 								CertManagerCertificateSpec: &armiotoperations.CertManagerCertificateSpec{
		// 									Duration: to.Ptr("qmpeffoksron"),
		// 									SecretName: to.Ptr("oagi"),
		// 									RenewBefore: to.Ptr("hutno"),
		// 									IssuerRef: &armiotoperations.CertManagerIssuerRef{
		// 										Group: to.Ptr("jtmuladdkpasfpoyvewekmiy"),
		// 										Kind: to.Ptr(armiotoperations.CertManagerIssuerKindIssuer),
		// 										Name: to.Ptr("ocwoqpgucvjrsuudtjhb"),
		// 									},
		// 									PrivateKey: &armiotoperations.CertManagerPrivateKey{
		// 										Algorithm: to.Ptr(armiotoperations.PrivateKeyAlgorithmEc256),
		// 										RotationPolicy: to.Ptr(armiotoperations.PrivateKeyRotationPolicyAlways),
		// 									},
		// 									San: &armiotoperations.SanForCert{
		// 										DNS: []*string{
		// 											to.Ptr("xhvmhrrhgfsapocjeebqtnzarlj"),
		// 										},
		// 										IP: []*string{
		// 											to.Ptr("zbgugfzcgsmegevzktsnibyuyp"),
		// 										},
		// 									},
		// 								},
		// 								Manual: &armiotoperations.X509ManualCertificate{
		// 									SecretRef: to.Ptr("secret-name"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ServiceType: to.Ptr(armiotoperations.ServiceTypeClusterIP),
		// 					ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
		// 				},
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/brokers/resource-name123/listeners/resource-name123"),
		// 				Name: to.Ptr("hoqjaachratt"),
		// 				Type: to.Ptr("hizbknwegcdaeh"),
		// 				SystemData: &armiotoperations.SystemData{
		// 					CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
		// 					CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("gnicpuszwd"),
		// 					LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
