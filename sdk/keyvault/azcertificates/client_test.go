//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcertificates

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

var ctx = context.TODO()

func TestNewClient(t *testing.T) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)
	client, err := NewClient("https://certvault.vault.azure.net", cred, nil)
	require.NoError(t, err)
	require.NotNil(t, client.genClient)
	require.Equal(t, "https://certvault.vault.azure.net", client.vaultURL)
}

func TestClient_BeginCreateCertificate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "beginCreate")
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	require.NoError(t, err)

	pollerResp, err := resp.PollUntilDone(ctx, delay())
	require.NoError(t, err)
	require.NotNil(t, pollerResp.ID)

	defer cleanUp(t, client, certName)

	// want to interface with x509 std library

	mid := base64.StdEncoding.EncodeToString(pollerResp.Csr)
	csr := fmt.Sprintf("-----BEGIN CERTIFICATE REQUEST-----\n%s\n-----END CERTIFICATE REQUEST-----", mid)

	// load certificate request
	csrblock, _ := pem.Decode([]byte(csr))
	require.NotNil(t, csrblock)
	req, err := x509.ParseCertificateRequest(csrblock.Bytes)
	require.NoError(t, err)
	require.NoError(t, req.CheckSignature())
}

func TestClient_BeginDeleteCertificate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "createCert")
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	require.NoError(t, err)

	pollerResp, err := resp.PollUntilDone(ctx, delay())
	require.NoError(t, err)
	require.NotNil(t, pollerResp.ID)

	delResp, err := client.BeginDeleteCertificate(ctx, certName, nil)
	require.NoError(t, err)

	delPollerResp, err := delResp.PollUntilDone(ctx, delay())
	require.NoError(t, err)
	require.Contains(t, *delPollerResp.ID, certName)

	_, err = client.GetCertificate(ctx, certName, nil)
	require.Error(t, err)

	deletedResp, err := client.GetDeletedCertificate(ctx, certName, nil)
	require.NoError(t, err)
	require.Contains(t, *deletedResp.ID, certName)

	_, err = client.PurgeDeletedCertificate(ctx, certName, nil)
	require.NoError(t, err)
}

func TestClient_GetCertificateOperation(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "cert")
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	require.NoError(t, err)

	_, err = resp.PollUntilDone(ctx, delay())
	require.NoError(t, err)

	resp2, err := client.GetCertificateOperation(ctx, certName, nil)
	require.NoError(t, err)
	require.NotNil(t, resp2.ID)

	cleanUp(t, client, certName)
}
func TestClient_CancelCertificateOperation(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "cert")
	require.NoError(t, err)

	_, err = client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	require.NoError(t, err)

	cancelResp, err := client.CancelCertificateOperation(ctx, certName, nil)
	require.NoError(t, err)
	require.Contains(t, *cancelResp.ID, certName)

	getResp, err := client.GetCertificateOperation(ctx, certName, nil)
	require.NoError(t, err)
	require.Equal(t, true, *getResp.CancellationRequested)

	_, err = client.DeleteCertificateOperation(ctx, certName, nil)
	require.NoError(t, err)

	// Get should fail now
	_, err = client.GetCertificateOperation(ctx, certName, nil)
	require.Error(t, err)
}

func TestClient_BackupCertificate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "cert")
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	require.NoError(t, err)
	_, err = resp.PollUntilDone(ctx, delay())
	require.NoError(t, err)

	backup, err := client.BackupCertificate(ctx, certName, nil)
	require.NoError(t, err)
	require.Greater(t, len(backup.Value), 0)

	cleanUp(t, client, certName)
}

func TestClient_ListCertificates(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	createdCount := 0
	for i := 0; i < 4; i++ {
		name, err := createRandomName(t, fmt.Sprintf("cert%d", i))
		require.NoError(t, err)
		createCert(t, client, name)
		defer cleanUp(t, client, name)
		createdCount++
	}

	time.Sleep(10 * delay())

	pager := client.ListCertificates(nil)
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		createdCount -= len(page.Certificates)
	}

	require.Equal(t, 0, createdCount)
}

func TestClient_ListCertificateVersions(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	name, err := createRandomName(t, "cert1")
	require.NoError(t, err)
	createCert(t, client, name)
	time.Sleep(10 * delay())
	defer cleanUp(t, client, name)

	pager := client.ListCertificateVersions(name, nil)
	count := 0
	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		count += len(resp.Certificates)
	}

	require.Equal(t, 1, count)

	// Add a second version
	createCert(t, client, name)
	time.Sleep(10 * delay())

	pager = client.ListCertificateVersions(name, nil)
	count = 0
	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		count += len(resp.Certificates)
	}

	require.Equal(t, 2, count)

	// Add a third version
	createCert(t, client, name)
	time.Sleep(10 * delay())

	pager = client.ListCertificateVersions(name, nil)
	count = 0
	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		count += len(resp.Certificates)
	}

	require.Equal(t, 3, count)
}

func TestClient_ImportCertificate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	importedName, err := createRandomName(t, "imported")
	require.NoError(t, err)

	importResp, err := client.ImportCertificate(ctx, importedName, certContentNotPasswordEncoded, nil)
	require.NoError(t, err)
	require.Contains(t, *importResp.ID, importedName)
	require.NotNil(t, importResp.Policy)

	cleanUp(t, client, importedName)
}

func TestClient_IssuerCRUD(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	issuerName, err := createRandomName(t, "issuer")
	require.NoError(t, err)

	resp, err := client.CreateIssuer(ctx, issuerName, "Test", &CreateIssuerOptions{
		Credentials: &IssuerCredentials{
			AccountID: to.Ptr("keyvaultuser"),
		},
		Enabled: to.Ptr(true),
		AdministratorContacts: []*AdministratorContact{
			{
				FirstName:    to.Ptr("John"),
				LastName:     to.Ptr("Doe"),
				EmailAddress: to.Ptr("admin@microsoft.com"),
				Phone:        to.Ptr("4255555555"),
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, *resp.Provider, "Test")
	require.Equal(t, *resp.Credentials.AccountID, "keyvaultuser")
	require.Contains(t, *resp.ID, fmt.Sprintf("/certificates/issuers/%s", issuerName))

	getResp, err := client.GetIssuer(ctx, issuerName, nil)
	require.NoError(t, err)
	require.Equal(t, *getResp.Provider, "Test")
	require.Equal(t, *getResp.Credentials.AccountID, "keyvaultuser")
	require.Contains(t, *getResp.ID, fmt.Sprintf("/certificates/issuers/%s", issuerName))

	issuerName2, err := createRandomName(t, "issuer2")
	require.NoError(t, err)

	_, err = client.CreateIssuer(ctx, issuerName2, "Test", &CreateIssuerOptions{
		Credentials: &IssuerCredentials{
			AccountID: to.Ptr("keyvaultuser2"),
		},
		Enabled: to.Ptr(true),
		AdministratorContacts: []*AdministratorContact{
			{
				FirstName:    to.Ptr("John"),
				LastName:     to.Ptr("Doe"),
				EmailAddress: to.Ptr("admin@microsoft.com"),
				Phone:        to.Ptr("4255555555"),
			},
		},
	})
	require.NoError(t, err)

	// List operation
	pager := client.ListPropertiesOfIssuers(nil)
	count := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		for _, issuer := range page.Issuers {
			require.Equal(t, "Test", *issuer.Provider)
			count += 1
		}
	}
	require.GreaterOrEqual(t, count, 2)

	// Update the certificate issuer
	updateResp, err := client.UpdateIssuer(ctx, issuerName2, &UpdateIssuerOptions{
		AdministratorContacts: []*AdministratorContact{
			{
				FirstName:    to.Ptr("Jane"),
				LastName:     to.Ptr("Doey"),
				EmailAddress: to.Ptr("admin2@microsoft.com"),
				Phone:        to.Ptr("4266666666"),
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(updateResp.AdministratorContacts))
	require.Equal(t, "Jane", *updateResp.AdministratorContacts[0].FirstName)
	require.Equal(t, "Doey", *updateResp.AdministratorContacts[0].LastName)
	require.Equal(t, "admin2@microsoft.com", *updateResp.AdministratorContacts[0].EmailAddress)
	require.Equal(t, "4266666666", *updateResp.AdministratorContacts[0].Phone)

	// Delete the first issuer
	_, err = client.DeleteIssuer(ctx, issuerName, nil)
	require.NoError(t, err)

	// Get on the first issuer fails
	_, err = client.GetIssuer(ctx, issuerName, nil)
	require.Error(t, err)
}

func TestClient_ContactsCRUD(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	contacts := Contacts{ContactList: []*Contact{
		{EmailAddress: to.Ptr("admin@microsoft.com"), Name: to.Ptr("John Doe"), Phone: to.Ptr("1111111111")},
		{EmailAddress: to.Ptr("admin@contoso.com"), Name: to.Ptr("Jane Doey"), Phone: to.Ptr("2222222222")},
	}}

	resp, err := client.SetContacts(ctx, contacts, nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(resp.ContactList))

	getResp, err := client.GetContacts(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(getResp.ContactList))
	require.Equal(t, "admin@microsoft.com", *getResp.ContactList[0].EmailAddress)
	require.Equal(t, "admin@contoso.com", *getResp.ContactList[1].EmailAddress)
	require.Equal(t, "John Doe", *getResp.ContactList[0].Name)
	require.Equal(t, "Jane Doey", *getResp.ContactList[1].Name)
	require.Equal(t, "1111111111", *getResp.ContactList[0].Phone)
	require.Equal(t, "2222222222", *getResp.ContactList[1].Phone)

	deleteResp, err := client.DeleteContacts(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(deleteResp.ContactList))

	// Get should fail
	_, err = client.GetContacts(ctx, nil)
	require.Error(t, err)
}

func TestPolicy(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "policyCertificate")
	require.NoError(t, err)

	policy := Policy{
		IssuerParameters: &IssuerParameters{
			CertificateTransparency: to.Ptr(false),
			Name:                    to.Ptr("Self"),
		},
		Exportable: to.Ptr(true),
		KeySize:    to.Ptr(int32(2048)),
		ReuseKey:   to.Ptr(true),
		KeyType:    to.Ptr(KeyTypeRSA),
		LifetimeActions: []*LifetimeAction{
			{Action: &Action{ActionType: to.Ptr(PolicyActionEmailContacts)}, Trigger: &Trigger{LifetimePercentage: to.Ptr(int32(98))}},
		},
		SecretProperties: &SecretProperties{
			ContentType: to.Ptr("application/x-pkcs12"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Ekus:             []*string{to.Ptr("1.3.6.1.5.5.7.3.1"), to.Ptr("1.3.6.1.5.5.7.3.2")},
			KeyUsage:         []*KeyUsage{to.Ptr(KeyUsageDecipherOnly)},
			Subject:          to.Ptr("CN=DefaultPolicy"),
			ValidityInMonths: to.Ptr(int32(12)),
			SubjectAlternativeNames: &SubjectAlternativeNames{
				DNSNames: []*string{to.Ptr("sdk.azure-int.net")},
			},
		},
	}

	_, err = client.BeginCreateCertificate(ctx, certName, policy, nil)
	require.NoError(t, err)

	receivedPolicy, err := client.GetCertificatePolicy(ctx, certName, nil)
	require.NoError(t, err)

	// Make sure policies are equal
	require.Equal(t, *policy.IssuerParameters.Name, *receivedPolicy.Policy.IssuerParameters.Name)
	require.Equal(t, *policy.Exportable, *receivedPolicy.Exportable)
	require.Equal(t, *policy.SecretProperties.ContentType, *receivedPolicy.SecretProperties.ContentType)

	// Update the policy
	policy.KeyType = to.Ptr(KeyTypeEC)
	policy.KeySize = to.Ptr(int32(256))
	policy.KeyCurveName = to.Ptr(KeyCurveNameP256)

	updateResp, err := client.UpdateCertificatePolicy(ctx, certName, policy, nil)
	require.NoError(t, err)

	require.Equal(t, *policy.IssuerParameters.Name, *updateResp.Policy.IssuerParameters.Name)
	require.Equal(t, *policy.Exportable, *updateResp.Exportable)
	require.Equal(t, *policy.SecretProperties.ContentType, *updateResp.SecretProperties.ContentType)
	require.Equal(t, *policy.KeyType, *updateResp.KeyType)
	require.Equal(t, *policy.KeySize, *updateResp.KeySize)
	require.Equal(t, *policy.KeyCurveName, *updateResp.KeyCurveName)

}

func TestCRUDOperations(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "cert")
	require.NoError(t, err)

	policy := Policy{
		IssuerParameters: &IssuerParameters{
			CertificateTransparency: to.Ptr(false),
			Name:                    to.Ptr("Self"),
		},
		Exportable: to.Ptr(true),
		KeySize:    to.Ptr(int32(2048)),
		ReuseKey:   to.Ptr(true),
		KeyType:    to.Ptr(KeyTypeRSA),
		LifetimeActions: []*LifetimeAction{
			{Action: &Action{ActionType: to.Ptr(PolicyActionEmailContacts)}, Trigger: &Trigger{LifetimePercentage: to.Ptr(int32(98))}},
		},
		SecretProperties: &SecretProperties{
			ContentType: to.Ptr("application/x-pkcs12"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Ekus:             []*string{to.Ptr("1.3.6.1.5.5.7.3.1"), to.Ptr("1.3.6.1.5.5.7.3.2")},
			KeyUsage:         []*KeyUsage{to.Ptr(KeyUsageDecipherOnly)},
			Subject:          to.Ptr("CN=DefaultPolicy"),
			ValidityInMonths: to.Ptr(int32(12)),
			SubjectAlternativeNames: &SubjectAlternativeNames{
				DNSNames: []*string{to.Ptr("sdk.azure-int.net")},
			},
		},
	}

	pollerResp, err := client.BeginCreateCertificate(ctx, certName, policy, nil)
	require.NoError(t, err)
	finalResp, err := pollerResp.PollUntilDone(ctx, time.Second)
	require.NoError(t, err)

	received, err := client.GetCertificate(ctx, certName, nil)
	require.NoError(t, err)
	require.NotNil(t, received.Policy)

	// Make sure certificates are the same
	require.Equal(t, *finalResp.ID, *received.ID)

	// // Make sure we can interface with x509 library
	// mid := base64.StdEncoding.EncodeToString(received.Cer)
	// cer := fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", mid)
	// block, _ := pem.Decode([]byte(cer))
	// require.NotNil(t, block)
	// parsedCert, err := x509.ParseCertificate(block.Bytes)
	// require.NoError(t, err)
	// require.NotNil(t, parsedCert)

	// Update the policy
	policy.KeyType = to.Ptr(KeyTypeEC)
	policy.KeySize = to.Ptr(int32(256))
	policy.KeyCurveName = to.Ptr(KeyCurveNameP256)

	updateResp, err := client.UpdateCertificatePolicy(ctx, certName, policy, nil)
	require.NoError(t, err)

	require.Equal(t, *policy.IssuerParameters.Name, *updateResp.Policy.IssuerParameters.Name)
	require.Equal(t, *policy.Exportable, *updateResp.Exportable)
	require.Equal(t, *policy.SecretProperties.ContentType, *updateResp.SecretProperties.ContentType)
	require.Equal(t, *policy.KeyType, *updateResp.KeyType)
	require.Equal(t, *policy.KeySize, *updateResp.KeySize)
	require.Equal(t, *policy.KeyCurveName, *updateResp.KeyCurveName)

	updatePropsResp, err := client.UpdateCertificateProperties(ctx, certName, &UpdateCertificatePropertiesOptions{
		Properties: &Properties{Tags: map[string]string{"tag1": "updated_values1"}},
	})
	require.NoError(t, err)
	require.Equal(t, "updated_values1", updatePropsResp.Properties.Tags["tag1"])
	require.Equal(t, *received.ID, *updatePropsResp.ID)
}

// https://stackoverflow.com/questions/42643048/signing-certificate-request-with-certificate-authority
// Much of this is thanks to this response, thanks @krostar
func TestMergeCertificate(t *testing.T) {
	recording.LiveOnly(t)
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "mergeCertificate")
	require.NoError(t, err)

	certPolicy := Policy{
		IssuerParameters: &IssuerParameters{
			Name:                    to.Ptr("Unknown"),
			CertificateTransparency: to.Ptr(false),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=MyCert"),
		},
	}

	// Load public key
	data, err := ioutil.ReadFile("testdata/ca.crt")
	require.NoError(t, err)
	block, _ := pem.Decode(data)
	require.NotNil(t, block)
	caCert, err := x509.ParseCertificate(block.Bytes)
	require.NoError(t, err)

	data, err = ioutil.ReadFile("testdata/ca.key")
	require.NoError(t, err)
	pkeyBlock, _ := pem.Decode(data)
	require.NotNil(t, pkeyBlock)
	require.Equal(t, pkeyBlock.Type, "RSA PRIVATE KEY")
	pkey, err := x509.ParsePKCS1PrivateKey(pkeyBlock.Bytes)
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, certPolicy, nil)
	require.NoError(t, err)
	_, err = resp.PollUntilDone(ctx, time.Second)
	require.NoError(t, err)
	defer cleanUp(t, client, certName)

	certOpResp, err := client.GetCertificateOperation(ctx, certName, nil)
	require.NoError(t, err)

	mid := base64.StdEncoding.EncodeToString(certOpResp.Csr)
	csr := fmt.Sprintf("-----BEGIN CERTIFICATE REQUEST-----\n%s\n-----END CERTIFICATE REQUEST-----", mid)

	// load certificate request
	csrblock, _ := pem.Decode([]byte(csr))
	require.NotNil(t, csrblock)
	req, err := x509.ParseCertificateRequest(csrblock.Bytes)
	require.NoError(t, err)
	require.NoError(t, req.CheckSignature())

	cert := x509.Certificate{
		SerialNumber:       big.NewInt(1),
		NotBefore:          time.Now(),
		NotAfter:           time.Now().Add(24 * time.Hour),
		Issuer:             caCert.Issuer,
		Subject:            req.Subject,
		PublicKey:          req.PublicKey,
		PublicKeyAlgorithm: req.PublicKeyAlgorithm,
		Signature:          req.Signature,
		SignatureAlgorithm: req.SignatureAlgorithm,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &cert, caCert, req.PublicKey, pkey)
	require.NoError(t, err)

	// Need to strip the BEGIN/END from the certificate
	certificateString := string(certBytes)
	certificateString = strings.Replace(certificateString, "-----Begin Certificate-----", "", 1)
	certificateString = strings.Replace(certificateString, "-----End Certificate-----", "", 1)

	mergeResp, err := client.MergeCertificate(ctx, certName, [][]byte{[]byte(certificateString)}, nil)
	require.NoError(t, err)
	require.NotNil(t, mergeResp.Policy)
}

func TestClient_BeginRecoverDeletedCertificate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "certRecover")
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
		},
	}, nil)
	require.NoError(t, err)
	defer cleanUp(t, client, certName)

	pollerResp, err := resp.PollUntilDone(ctx, delay())
	require.NoError(t, err)
	require.NotNil(t, pollerResp.ID)

	delResp, err := client.BeginDeleteCertificate(ctx, certName, nil)
	require.NoError(t, err)

	delPollerResp, err := delResp.PollUntilDone(ctx, delay())
	require.NoError(t, err)
	require.Contains(t, *delPollerResp.ID, certName)

	_, err = client.GetCertificate(ctx, certName, nil)
	require.Error(t, err)

	recover, err := client.BeginRecoverDeletedCertificate(ctx, certName, nil)
	require.NoError(t, err)

	recoveredResp, err := recover.PollUntilDone(ctx, time.Second)
	require.NoError(t, err)
	require.Contains(t, *recoveredResp.ID, certName)
}

func TestClient_RestoreCertificateBackup(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	certName, err := createRandomName(t, "certRestore")
	require.NoError(t, err)

	resp, err := client.BeginCreateCertificate(ctx, certName, Policy{
		IssuerParameters: &IssuerParameters{
			Name: to.Ptr("Self"),
		},
		X509CertificateProperties: &X509CertificateProperties{
			Subject: to.Ptr("CN=DefaultPolicy"),
			SubjectAlternativeNames: &SubjectAlternativeNames{
				Upns: []*string{to.Ptr("john.doe@domain.com")},
			},
		},
	}, nil)
	require.NoError(t, err)
	defer cleanUp(t, client, certName)

	pollerResp, err := resp.PollUntilDone(ctx, delay())
	require.NoError(t, err)
	require.NotNil(t, pollerResp.ID)

	// Create a backup
	certificateBackupResp, err := client.BackupCertificate(ctx, certName, nil)
	require.NoError(t, err)

	// Delete the certificate
	deletePoller, err := client.BeginDeleteCertificate(ctx, certName, nil)
	require.NoError(t, err)

	_, err = deletePoller.PollUntilDone(ctx, delay())
	require.NoError(t, err)

	// Purge the cert
	_, err = client.PurgeDeletedCertificate(ctx, certName, nil)
	require.NoError(t, err)

	// Restore the cert
	// Poll until no exception
	count := 0
	for {
		resp, err := client.RestoreCertificateBackup(ctx, certificateBackupResp.Value, nil)
		if err == nil {
			require.NotNil(t, resp.Policy)
			break
		}
		var respErr *azcore.ResponseError
		if errors.As(err, &respErr) {
			if respErr.RawResponse.StatusCode != 409 {
				require.NoError(t, err)
			}
		} else {
			require.NoError(t, err)
		}
		count += 1
		if count > 25 {
			require.NoError(t, err)
		}
		recording.Sleep(5 * time.Second)
	}
}

func TestClient_ListDeletedCertificates(t *testing.T) {
	stop := startTest(t)
	defer stop()

	client, err := createClient(t)
	require.NoError(t, err)

	var names []string
	createdCount := 0
	for i := 0; i < 4; i++ {
		name, err := createRandomName(t, fmt.Sprintf("delCert%d", i))
		require.NoError(t, err)
		names = append(names, name)
		createCert(t, client, name)
		createdCount++
	}

	for _, name := range names {
		poller, err := client.BeginDeleteCertificate(ctx, name, nil)
		require.NoError(t, err)
		_, err = poller.PollUntilDone(ctx, delay())
		require.NoError(t, err)
	}

	time.Sleep(10 * delay())

	pager := client.ListDeletedCertificates(nil)
	deletedCount := 0
	for pager.More() {
		page, err := pager.NextPage(ctx)
		require.NoError(t, err)
		for i := range page.Certificates {
			purgeCert(t, client, names[i])
			deletedCount += 1
		}
	}
	require.Equal(t, 4, createdCount)
}
