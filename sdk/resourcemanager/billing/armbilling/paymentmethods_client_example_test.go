//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)
import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"reflect"
	"time"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByBillingProfile.json
func ExamplePaymentMethodsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPaymentMethodsClient().NewListByBillingProfilePager("00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", "ABC1-A1CD-AB1-BP1", nil)
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
		// page.PaymentMethodLinksListResult = armbilling.PaymentMethodLinksListResult{
		// 	Value: []*armbilling.PaymentMethodLink{
		// 		{
		// 			Name: to.Ptr("ABCDABCDABC0"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/paymentMethodLinks"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31/billingProfiles/ABC1-A1CD-AB1-BP1/paymentMethodLinks/ABCDABCDABC0"),
		// 			Properties: &armbilling.PaymentMethodLinkProperties{
		// 				AccountHolderName: to.Ptr("abc"),
		// 				DisplayName: to.Ptr("Master Card"),
		// 				Expiration: to.Ptr("1/2035"),
		// 				Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 				LastFourDigits: to.Ptr("1270"),
		// 				Logos: []*armbilling.PaymentMethodLogo{
		// 					{
		// 						MimeType: to.Ptr("image/png"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
		// 					},
		// 					{
		// 						MimeType: to.Ptr("image/svg+xml"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
		// 				}},
		// 				PaymentMethodID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
		// 				PaymentMethodType: to.Ptr("mc"),
		// 				Status: to.Ptr(armbilling.PaymentMethodStatusActive),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByBillingProfile.json
func ExamplePaymentMethodsClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaymentMethodsClient().GetByBillingProfile(ctx, "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", "ABC1-A1CD-AB1-BP1", "ABCDABCDABC0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaymentMethodLink = armbilling.PaymentMethodLink{
	// 	Name: to.Ptr("ABCDABCDABC0"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/paymentMethodLinks"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31/billingProfiles/ABC1-A1CD-AB1-BP1/paymentMethodLinks/ABCDABCDABC0"),
	// 	Properties: &armbilling.PaymentMethodLinkProperties{
	// 		AccountHolderName: to.Ptr("abc"),
	// 		DisplayName: to.Ptr("Master Card"),
	// 		Expiration: to.Ptr("1/2035"),
	// 		Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
	// 		LastFourDigits: to.Ptr("1270"),
	// 		Logos: []*armbilling.PaymentMethodLogo{
	// 			{
	// 				MimeType: to.Ptr("image/png"),
	// 				URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
	// 			},
	// 			{
	// 				MimeType: to.Ptr("image/svg+xml"),
	// 				URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
	// 		}},
	// 		PaymentMethodID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
	// 		PaymentMethodType: to.Ptr("mc"),
	// 		Status: to.Ptr(armbilling.PaymentMethodStatusActive),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByBillingAccount.json
func ExamplePaymentMethodsClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPaymentMethodsClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", nil)
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
		// page.PaymentMethodsListResult = armbilling.PaymentMethodsListResult{
		// 	Value: []*armbilling.PaymentMethod{
		// 		{
		// 			Name: to.Ptr("21dd9edc-af71-4d62-80ce-37151d475326"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/paymentMethods"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31/paymentMethods/21dd9edc-af71-4d62-80ce-37151d475326"),
		// 			Properties: &armbilling.PaymentMethodProperties{
		// 				DisplayName: to.Ptr("Check/wire transfer"),
		// 				Family: to.Ptr(armbilling.PaymentMethodFamilyCheckWire),
		// 				PaymentMethodType: to.Ptr("check"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByBillingAccount.json
func ExamplePaymentMethodsClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaymentMethodsClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", "21dd9edc-af71-4d62-80ce-37151d475326", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaymentMethod = armbilling.PaymentMethod{
	// 	Name: to.Ptr("21dd9edc-af71-4d62-80ce-37151d475326"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/paymentMethods"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31/paymentMethods/21dd9edc-af71-4d62-80ce-37151d475326"),
	// 	Properties: &armbilling.PaymentMethodProperties{
	// 		DisplayName: to.Ptr("Check/wire transfer"),
	// 		Family: to.Ptr(armbilling.PaymentMethodFamilyCheckWire),
	// 		PaymentMethodType: to.Ptr("check"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByUser.json
func ExamplePaymentMethodsClient_NewListByUserPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPaymentMethodsClient().NewListByUserPager(nil)
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
		// page.PaymentMethodsListResult = armbilling.PaymentMethodsListResult{
		// 	Value: []*armbilling.PaymentMethod{
		// 		{
		// 			Name: to.Ptr("ABCDABCDABC0"),
		// 			Type: to.Ptr("Microsoft.Billing/paymentMethods"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
		// 			Properties: &armbilling.PaymentMethodProperties{
		// 				AccountHolderName: to.Ptr("abc"),
		// 				DisplayName: to.Ptr("Master Card"),
		// 				Expiration: to.Ptr("1/2035"),
		// 				Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 				LastFourDigits: to.Ptr("1270"),
		// 				Logos: []*armbilling.PaymentMethodLogo{
		// 					{
		// 						MimeType: to.Ptr("image/png"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_mc_rect.png"),
		// 					},
		// 					{
		// 						MimeType: to.Ptr("image/svg+xml"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_mc.svg"),
		// 				}},
		// 				PaymentMethodType: to.Ptr("mc"),
		// 				Status: to.Ptr(armbilling.PaymentMethodStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ABCDABCDABC1"),
		// 			Type: to.Ptr("Microsoft.Billing/paymentMethods"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC1"),
		// 			Properties: &armbilling.PaymentMethodProperties{
		// 				AccountHolderName: to.Ptr("abc"),
		// 				DisplayName: to.Ptr("Visa"),
		// 				Expiration: to.Ptr("1/2025"),
		// 				Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
		// 				LastFourDigits: to.Ptr("7373"),
		// 				Logos: []*armbilling.PaymentMethodLogo{
		// 					{
		// 						MimeType: to.Ptr("image/png"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
		// 					},
		// 					{
		// 						MimeType: to.Ptr("image/svg+xml"),
		// 						URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
		// 				}},
		// 				PaymentMethodType: to.Ptr("visa"),
		// 				Status: to.Ptr(armbilling.PaymentMethodStatusActive),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsDeleteByUser.json
func ExamplePaymentMethodsClient_DeleteByUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPaymentMethodsClient().DeleteByUser(ctx, "ABCDABCDABC0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByUser.json
func ExamplePaymentMethodsClient_GetByUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPaymentMethodsClient().GetByUser(ctx, "ABCDABCDABC0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaymentMethod = armbilling.PaymentMethod{
	// 	Name: to.Ptr("ABCDABCDABC0"),
	// 	Type: to.Ptr("Microsoft.Billing/paymentMethods"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/paymentMethods/ABCDABCDABC0"),
	// 	Properties: &armbilling.PaymentMethodProperties{
	// 		AccountHolderName: to.Ptr("abc"),
	// 		DisplayName: to.Ptr("Master Card"),
	// 		Expiration: to.Ptr("1/2035"),
	// 		Family: to.Ptr(armbilling.PaymentMethodFamilyCreditCard),
	// 		LastFourDigits: to.Ptr("1270"),
	// 		Logos: []*armbilling.PaymentMethodLogo{
	// 			{
	// 				MimeType: to.Ptr("image/png"),
	// 				URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa_rect.png"),
	// 			},
	// 			{
	// 				MimeType: to.Ptr("image/svg+xml"),
	// 				URL: to.Ptr("https://contoso.com/staticresourceservice/images/v4/logo_visa.svg"),
	// 		}},
	// 		PaymentMethodType: to.Ptr("mc"),
	// 		Status: to.Ptr(armbilling.PaymentMethodStatusActive),
	// 	},
	// }
}
