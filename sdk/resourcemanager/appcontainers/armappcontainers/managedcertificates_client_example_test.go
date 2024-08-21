//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificate_Get.json
func ExampleManagedCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedCertificatesClient().Get(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedCertificate = armappcontainers.ManagedCertificate{
	// 	Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ManagedCertificateProperties{
	// 		DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
	// 		ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
	// 		SubjectName: to.Ptr("CN=my-subject-name.company.country.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificate_CreateOrUpdate.json
func ExampleManagedCertificatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedCertificatesClient().BeginCreateOrUpdate(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", &armappcontainers.ManagedCertificatesClientBeginCreateOrUpdateOptions{ManagedCertificateEnvelope: &armappcontainers.ManagedCertificate{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ManagedCertificateProperties{
			DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
			SubjectName:             to.Ptr("my-subject-name.company.country.net"),
		},
	},
	})
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
	// res.ManagedCertificate = armappcontainers.ManagedCertificate{
	// 	Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ManagedCertificateProperties{
	// 		DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
	// 		ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
	// 		SubjectName: to.Ptr("CN=my-subject-name.company.country.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificate_Delete.json
func ExampleManagedCertificatesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagedCertificatesClient().Delete(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificates_Patch.json
func ExampleManagedCertificatesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedCertificatesClient().Update(ctx, "examplerg", "testcontainerenv", "certificate-firendly-name", armappcontainers.ManagedCertificatePatch{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedCertificate = armappcontainers.ManagedCertificate{
	// 	Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ManagedCertificateProperties{
	// 		DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
	// 		ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
	// 		SubjectName: to.Ptr("CN=my-subject-name.company.country.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificates_ListByManagedEnvironment.json
func ExampleManagedCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedCertificatesClient().NewListPager("examplerg", "testcontainerenv", nil)
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
		// page.ManagedCertificateCollection = armappcontainers.ManagedCertificateCollection{
		// 	Value: []*armappcontainers.ManagedCertificate{
		// 		{
		// 			Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ManagedCertificateProperties{
		// 				DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationCNAME),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectName: to.Ptr("CN=my-subject-name.company.country.net"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name-root"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ManagedCertificateProperties{
		// 				DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationHTTP),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectName: to.Ptr("CN=company.country.net"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.App/ManagedEnvironments/managedCertificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/managedCertificates/certificate-firendly-name-txt"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.ManagedCertificateProperties{
		// 				DomainControlValidation: to.Ptr(armappcontainers.ManagedCertificateDomainControlValidationTXT),
		// 				ProvisioningState: to.Ptr(armappcontainers.CertificateProvisioningStateSucceeded),
		// 				SubjectName: to.Ptr("CN=txt.company.country.net"),
		// 			},
		// 	}},
		// }
	}
}
