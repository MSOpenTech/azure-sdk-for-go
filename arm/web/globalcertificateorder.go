package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// GlobalCertificateOrderClient is the use these APIs to manage Azure Websites
// resources through the Azure Resource Manager. All task operations conform to
// the HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see
// https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx.
type GlobalCertificateOrderClient struct {
	ManagementClient
}

// NewGlobalCertificateOrderClient creates an instance of the
// GlobalCertificateOrderClient client.
func NewGlobalCertificateOrderClient(subscriptionID string) GlobalCertificateOrderClient {
	return NewGlobalCertificateOrderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGlobalCertificateOrderClientWithBaseURI creates an instance of the
// GlobalCertificateOrderClient client.
func NewGlobalCertificateOrderClientWithBaseURI(baseURI string, subscriptionID string) GlobalCertificateOrderClient {
	return GlobalCertificateOrderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAllCertificateOrders sends the get all certificate orders request.
func (client GlobalCertificateOrderClient) GetAllCertificateOrders() (result CertificateOrderCollection, err error) {
	req, err := client.GetAllCertificateOrdersPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", nil, "Failure preparing request")
	}

	resp, err := client.GetAllCertificateOrdersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", resp, "Failure sending request")
	}

	result, err = client.GetAllCertificateOrdersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", resp, "Failure responding to request")
	}

	return
}

// GetAllCertificateOrdersPreparer prepares the GetAllCertificateOrders request.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAllCertificateOrdersSender sends the GetAllCertificateOrders request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetAllCertificateOrdersResponder handles the response to the GetAllCertificateOrders request. The method always
// closes the http.Response Body.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersResponder(resp *http.Response) (result CertificateOrderCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAllCertificateOrdersNextResults retrieves the next set of results, if any.
func (client GlobalCertificateOrderClient) GetAllCertificateOrdersNextResults(lastResults CertificateOrderCollection) (result CertificateOrderCollection, err error) {
	req, err := lastResults.CertificateOrderCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetAllCertificateOrdersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", resp, "Failure sending next results request")
	}

	result, err = client.GetAllCertificateOrdersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "GetAllCertificateOrders", resp, "Failure responding to next results request")
	}

	return
}

// ValidateCertificatePurchaseInformation sends the validate certificate
// purchase information request.
//
// certificateOrder is certificate order
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformation(certificateOrder CertificateOrder) (result SetObject, err error) {
	req, err := client.ValidateCertificatePurchaseInformationPreparer(certificateOrder)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "ValidateCertificatePurchaseInformation", nil, "Failure preparing request")
	}

	resp, err := client.ValidateCertificatePurchaseInformationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "ValidateCertificatePurchaseInformation", resp, "Failure sending request")
	}

	result, err = client.ValidateCertificatePurchaseInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalCertificateOrderClient", "ValidateCertificatePurchaseInformation", resp, "Failure responding to request")
	}

	return
}

// ValidateCertificatePurchaseInformationPreparer prepares the ValidateCertificatePurchaseInformation request.
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformationPreparer(certificateOrder CertificateOrder) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/validateCertificateRegistrationInformation", pathParameters),
		autorest.WithJSON(certificateOrder),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ValidateCertificatePurchaseInformationSender sends the ValidateCertificatePurchaseInformation request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ValidateCertificatePurchaseInformationResponder handles the response to the ValidateCertificatePurchaseInformation request. The method always
// closes the http.Response Body.
func (client GlobalCertificateOrderClient) ValidateCertificatePurchaseInformationResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
