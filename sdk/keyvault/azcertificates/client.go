//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcertificates

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azcertificates/internal/generated"
	shared "github.com/Azure/azure-sdk-for-go/sdk/keyvault/internal"
)

// Client is the struct for interacting with a Key Vault Certificates instance.
type Client struct {
	genClient *generated.KeyVaultClient
	vaultURL  string
}

// ClientOptions are the optional parameters for the NewClient function
type ClientOptions struct {
	azcore.ClientOptions
}

// converts ClientOptions to generated *generated.ConnectionOptions
func (c *ClientOptions) toConnectionOptions() *policy.ClientOptions {
	if c == nil {
		return &policy.ClientOptions{}
	}

	return &policy.ClientOptions{
		Logging:          c.Logging,
		Retry:            c.Retry,
		Telemetry:        c.Telemetry,
		Transport:        c.Transport,
		PerCallPolicies:  c.PerCallPolicies,
		PerRetryPolicies: c.PerRetryPolicies,
	}
}

// NewClient creates an instance of a Client for a Key Vault Certificate URL.
func NewClient(vaultURL string, credential azcore.TokenCredential, options *ClientOptions) (Client, error) {
	genOptions := options.toConnectionOptions()

	genOptions.PerRetryPolicies = append(
		genOptions.PerRetryPolicies,
		shared.NewKeyVaultChallengePolicy(credential),
	)

	conn := generated.NewConnection(genOptions)

	return Client{
		genClient: generated.NewKeyVaultClient(conn),
		vaultURL:  vaultURL,
	}, nil
}

// Optional parameters for the Client.BeginCreateCertificate function
type BeginCreateCertificateOptions struct {
	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

func (b BeginCreateCertificateOptions) toGenerated() *generated.KeyVaultClientCreateCertificateOptions {
	return &generated.KeyVaultClientCreateCertificateOptions{}
}

// CreateCertificateResponse contains the response from method Client.BeginCreateCertificate.
type CreateCertificateResponse struct {
	CertificateOperation

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CreateCertificatePoller is the interface for the Client.BeginCreateCertificate operation.
type CreateCertificatePoller interface {
	// Done returns true if the LRO has reached a terminal state
	Done() bool

	// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
	// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
	// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
	Poll(context.Context) (*http.Response, error)

	// FinalResponse returns the final response after the operations has finished
	FinalResponse(context.Context) (CreateCertificateResponse, error)
}

// the poller returned by the Client.BeginCreateCertificate
type beginCreateCertificatePoller struct {
	certName       string
	certVersion    string
	vaultURL       string
	client         *generated.KeyVaultClient
	createResponse CreateCertificateResponse
	lastResponse   generated.KeyVaultClientGetCertificateResponse
	RawResponse    *http.Response
}

func (b *beginCreateCertificatePoller) Done() bool {
	return b.lastResponse.RawResponse.StatusCode == http.StatusOK
}

func (b *beginCreateCertificatePoller) Poll(ctx context.Context) (*http.Response, error) {
	resp, err := b.client.GetCertificate(ctx, b.vaultURL, b.certName, b.certVersion, nil)
	if err == nil {
		b.lastResponse = resp
		return resp.RawResponse, nil
	}

	if resp.RawResponse.StatusCode == http.StatusNotFound {
		// The certificate has not been fully created yet
		return b.createResponse.RawResponse, nil
	}

	// There was an error in this operation, return the original raw response and the error
	return b.createResponse.RawResponse, err
}

func (b *beginCreateCertificatePoller) FinalResponse(ctx context.Context) (CreateCertificateResponse, error) {
	return b.createResponse, nil
}

func (b *beginCreateCertificatePoller) pollUntilDone(ctx context.Context, t time.Duration) (CreateCertificateResponse, error) {
	for {
		resp, err := b.Poll(ctx)
		if err != nil {
			return CreateCertificateResponse{}, err
		}
		b.RawResponse = resp
		if b.Done() {
			break
		}
		time.Sleep(t)
	}
	return b.createResponse, nil
}

// CreateCertificatePollerResponse contains the response from the Client.BeginCreateCertificate method
type CreateCertificatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error occurs
	PollUntilDone func(context.Context, time.Duration) (CreateCertificateResponse, error)

	// Poller contains an initialized WidgetPoller
	Poller CreateCertificatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func (c *Client) BeginCreateCertificate(ctx context.Context, certName string, policy CertificatePolicy, options *BeginCreateCertificateOptions) (CreateCertificatePollerResponse, error) {
	if options == nil {
		options = &BeginCreateCertificateOptions{}
	}

	resp, err := c.genClient.CreateCertificate(
		ctx,
		c.vaultURL,
		certName,
		generated.CertificateCreateParameters{
			CertificatePolicy:     policy.toGeneratedCertificateCreateParameters(),
			Tags:                  options.Tags,
			CertificateAttributes: options.CertificateAttributes.toGenerated(),
		},
		options.toGenerated(),
	)

	if err != nil {
		return CreateCertificatePollerResponse{}, err
	}

	p := &beginCreateCertificatePoller{
		certName:    certName,
		certVersion: "",
		vaultURL:    c.vaultURL,
		client:      c.genClient,
		createResponse: CreateCertificateResponse{
			RawResponse: resp.RawResponse,
			CertificateOperation: CertificateOperation{
				CancellationRequested: resp.CancellationRequested,
				Csr:                   resp.Csr,
				Error:                 resp.Error,
				IssuerParameters:      resp.IssuerParameters,
				RequestID:             resp.RequestID,
				Status:                resp.Status,
				StatusDetails:         resp.StatusDetails,
				Target:                resp.Target,
				ID:                    resp.ID,
			},
		},
		lastResponse: generated.KeyVaultClientGetCertificateResponse{},
	}

	return CreateCertificatePollerResponse{
		Poller:        p,
		RawResponse:   resp.RawResponse,
		PollUntilDone: p.pollUntilDone,
	}, nil
}

// GetCertificateOptions contains the optional parameters for the Client.GetCertificate method.
type GetCertificateOptions struct {
	Version string
}

// GetCertificateResponse contains the result from method Client.GetCertificate.
type GetCertificateResponse struct {
	CertificateBundle

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GetCertificate - Gets information about a specific certificate. This operation requires the certificates/get permission.
// If the operation fails it returns the *KeyVaultError error type.
func (c *Client) GetCertificate(ctx context.Context, certName string, options *GetCertificateOptions) (GetCertificateResponse, error) {
	if options == nil {
		options = &GetCertificateOptions{}
	}

	resp, err := c.genClient.GetCertificate(ctx, c.vaultURL, certName, options.Version, nil)
	if err != nil {
		return GetCertificateResponse{}, err
	}

	return GetCertificateResponse{
		RawResponse: resp.RawResponse,
		CertificateBundle: CertificateBundle{
			Attributes:     certificateAttributesFromGenerated(resp.Attributes),
			Cer:            resp.Cer,
			ContentType:    resp.ContentType,
			Tags:           resp.Tags,
			ID:             resp.ID,
			Kid:            resp.Kid,
			Policy:         certificatePolicyFromGenerated(resp.Policy),
			Sid:            resp.Sid,
			X509Thumbprint: resp.X509Thumbprint,
		},
	}, nil
}

// GetCertificateOperationOptions contains the optional parameters for the Client.GetCertificateOperation method.
type GetCertificateOperationOptions struct{}

func (g *GetCertificateOperationOptions) toGenerated() *generated.KeyVaultClientGetCertificateOperationOptions {
	return &generated.KeyVaultClientGetCertificateOperationOptions{}
}

// GetCertificateOperationResponse contains the result from method Client.GetCertificateOperation.
type GetCertificateOperationResponse struct {
	CertificateOperation

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GetCertificateOperation - Gets the creation operation associated with a specified certificate. This operation requires the certificates/get permission.
// If the operation fails it returns the *KeyVaultError error type.
func (c *Client) GetCertificateOperation(ctx context.Context, certName string, options *GetCertificateOperationOptions) (GetCertificateOperationResponse, error) {
	resp, err := c.genClient.GetCertificateOperation(ctx, c.vaultURL, certName, options.toGenerated())
	if err != nil {
		return GetCertificateOperationResponse{}, err
	}

	return GetCertificateOperationResponse{
		RawResponse: resp.RawResponse,
		CertificateOperation: CertificateOperation{
			CancellationRequested: resp.CancellationRequested,
			Csr:                   resp.Csr,
			Error:                 resp.Error,
			IssuerParameters:      resp.IssuerParameters,
			RequestID:             resp.RequestID,
			Status:                resp.Status,
			StatusDetails:         resp.StatusDetails,
			Target:                resp.Target,
			ID:                    resp.ID,
		},
	}, nil
}

// BeginDeleteCertificateOptions contains the optional parameters for the Client.BeginDeleteCertificate method.
type BeginDeleteCertificateOptions struct{}

// convert public options to generated options struct
func (b *BeginDeleteCertificateOptions) toGenerated() *generated.KeyVaultClientDeleteCertificateOptions {
	return &generated.KeyVaultClientDeleteCertificateOptions{}
}

type DeleteCertificateResponse struct {
	DeletedCertificateBundle

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func deleteCertificateResponseFromGenerated(g *generated.KeyVaultClientDeleteCertificateResponse) DeleteCertificateResponse {
	if g == nil {
		return DeleteCertificateResponse{}
	}
	return DeleteCertificateResponse{
		RawResponse: g.RawResponse,
		DeletedCertificateBundle: DeletedCertificateBundle{
			RecoveryID:         g.RecoveryID,
			DeletedDate:        g.DeletedDate,
			ScheduledPurgeDate: g.ScheduledPurgeDate,
			CertificateBundle: CertificateBundle{
				Attributes:     certificateAttributesFromGenerated(g.Attributes),
				Cer:            g.Cer,
				ContentType:    g.ContentType,
				Tags:           g.Tags,
				ID:             g.ID,
				Kid:            g.Kid,
				Policy:         certificatePolicyFromGenerated(g.Policy),
				Sid:            g.Sid,
				X509Thumbprint: g.X509Thumbprint,
			},
		},
	}
}

// DeleteCertificatePoller is the interface for the Client.DeleteCertificate operation.
type DeleteCertificatePoller interface {
	// Done returns true if the LRO has reached a terminal state
	Done() bool

	// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
	// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
	// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
	Poll(context.Context) (*http.Response, error)

	// FinalResponse returns the final response after the operations has finished
	FinalResponse(context.Context) (DeleteCertificateResponse, error)
}

// The poller returned by the Client.BeginDeleteCertificate operation
type beginDeleteCertificatePoller struct {
	keyName        string // This is the key to Poll for in GetDeletedKey
	vaultURL       string
	client         *generated.KeyVaultClient
	deleteResponse generated.KeyVaultClientDeleteCertificateResponse
	lastResponse   generated.KeyVaultClientGetDeletedCertificateResponse
	RawResponse    *http.Response
}

// Done returns true if the LRO has reached a terminal state
func (s *beginDeleteCertificatePoller) Done() bool {
	return s.lastResponse.RawResponse != nil
}

// Poll fetches the latest state of the LRO. It returns an HTTP response or error.(
// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.)
func (s *beginDeleteCertificatePoller) Poll(ctx context.Context) (*http.Response, error) {
	resp, err := s.client.GetDeletedCertificate(ctx, s.vaultURL, s.keyName, nil)
	if err == nil {
		// Service recognizes DeletedKey, operation is done
		s.lastResponse = resp
		return resp.RawResponse, nil
	}

	var httpResponseErr azcore.HTTPResponse
	if errors.As(err, &httpResponseErr) {
		if httpResponseErr.RawResponse().StatusCode == http.StatusNotFound {
			// This is the expected result
			return s.deleteResponse.RawResponse, nil
		}
	}
	return s.deleteResponse.RawResponse, err
}

// FinalResponse returns the final response after the operations has finished
func (s *beginDeleteCertificatePoller) FinalResponse(ctx context.Context) (DeleteCertificateResponse, error) {
	return deleteCertificateResponseFromGenerated(&s.deleteResponse), nil
}

// pollUntilDone continually calls the Poll operation until the operation is completed. In between each
// Poll is a wait determined by the t parameter.
func (s *beginDeleteCertificatePoller) pollUntilDone(ctx context.Context, t time.Duration) (DeleteCertificateResponse, error) {
	for {
		resp, err := s.Poll(ctx)
		if err != nil {
			return DeleteCertificateResponse{}, err
		}
		s.RawResponse = resp
		if s.Done() {
			break
		}
		time.Sleep(t)
	}
	return deleteCertificateResponseFromGenerated(&s.deleteResponse), nil
}

// DeleteCertificatePollerResponse contains the response from the Client.BeginDeleteCertificate method
type DeleteCertificatePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error occurs
	PollUntilDone func(context.Context, time.Duration) (DeleteCertificateResponse, error)

	// Poller contains an initialized WidgetPoller
	Poller DeleteCertificatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BeginDeleteCertificate deletes a key from the keyvault. Delete cannot be applied to an individual version of a key. This operation
// requires the key/delete permission. This response contains a Poller struct that can be used to Poll for a response, or the
// response PollUntilDone function can be used to poll until completion.
func (c *Client) BeginDeleteCertificate(ctx context.Context, keyName string, options *BeginDeleteCertificateOptions) (DeleteCertificatePollerResponse, error) {
	if options == nil {
		options = &BeginDeleteCertificateOptions{}
	}
	resp, err := c.genClient.DeleteCertificate(ctx, c.vaultURL, keyName, options.toGenerated())
	if err != nil {
		return DeleteCertificatePollerResponse{}, err
	}

	getResp, err := c.genClient.GetDeletedCertificate(ctx, c.vaultURL, keyName, nil)
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		if httpErr.RawResponse().StatusCode != http.StatusNotFound {
			return DeleteCertificatePollerResponse{}, err
		}
	}

	s := &beginDeleteCertificatePoller{
		vaultURL:       c.vaultURL,
		keyName:        keyName,
		client:         c.genClient,
		deleteResponse: resp,
		lastResponse:   getResp,
	}

	return DeleteCertificatePollerResponse{
		Poller:        s,
		RawResponse:   resp.RawResponse,
		PollUntilDone: s.pollUntilDone,
	}, nil
}

// Optional parameters for the Client.PurgeDeletedCertificateOptions function
type PurgeDeletedCertificateOptions struct{}

func (p *PurgeDeletedCertificateOptions) toGenerated() *generated.KeyVaultClientPurgeDeletedCertificateOptions {
	return &generated.KeyVaultClientPurgeDeletedCertificateOptions{}
}

// PurgeDeletedCertificateResponse contains the response from method Client.PurgeDeletedCertificate.
type PurgeDeletedCertificateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func (c *Client) PurgeDeletedCertificate(ctx context.Context, certName string, options *PurgeDeletedCertificateOptions) (PurgeDeletedCertificateResponse, error) {
	resp, err := c.genClient.PurgeDeletedCertificate(ctx, c.vaultURL, certName, options.toGenerated())
	if err != nil {
		return PurgeDeletedCertificateResponse{}, err
	}

	return PurgeDeletedCertificateResponse{
		RawResponse: resp.RawResponse,
	}, nil
}

// Optional parameters for the Client.GetDeletedCertificateOptions function
type GetDeletedCertificateOptions struct{}

func (g *GetDeletedCertificateOptions) toGenerated() *generated.KeyVaultClientGetDeletedCertificateOptions {
	return &generated.KeyVaultClientGetDeletedCertificateOptions{}
}

type GetDeletedCertificateResponse struct {
	DeletedCertificateBundle

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func (c *Client) GetDeletedCertificate(ctx context.Context, certName string, options *GetDeletedCertificateOptions) (GetDeletedCertificateResponse, error) {
	resp, err := c.genClient.GetDeletedCertificate(ctx, c.vaultURL, certName, options.toGenerated())
	if err != nil {
		return GetDeletedCertificateResponse{}, err
	}

	return GetDeletedCertificateResponse{
		RawResponse: resp.RawResponse,
		DeletedCertificateBundle: DeletedCertificateBundle{
			RecoveryID:         resp.RecoveryID,
			DeletedDate:        resp.DeletedDate,
			ScheduledPurgeDate: resp.ScheduledPurgeDate,
			CertificateBundle: CertificateBundle{
				Attributes:     certificateAttributesFromGenerated(resp.Attributes),
				Cer:            resp.Cer,
				ContentType:    resp.ContentType,
				Tags:           resp.Tags,
				ID:             resp.ID,
				Kid:            resp.Kid,
				Policy:         certificatePolicyFromGenerated(resp.Policy),
				Sid:            resp.Sid,
				X509Thumbprint: resp.X509Thumbprint,
			},
		},
	}, nil
}

// Optional parameters for the Client.BackupCertificateOptions function
type BackupCertificateOptions struct{}

func (b *BackupCertificateOptions) toGenerated() *generated.KeyVaultClientBackupCertificateOptions {
	return &generated.KeyVaultClientBackupCertificateOptions{}
}

// BackupCertificateResponse contains the response from method Client.BackupCertificate.
type BackupCertificateResponse struct {
	// READ-ONLY; The backup blob containing the backed up certificate.
	Value []byte `json:"value,omitempty" azure:"ro"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func (c *Client) BackupCertificate(ctx context.Context, certName string, options *BackupCertificateOptions) (BackupCertificateResponse, error) {
	resp, err := c.genClient.BackupCertificate(ctx, c.vaultURL, certName, options.toGenerated())
	if err != nil {
		return BackupCertificateResponse{}, err
	}

	return BackupCertificateResponse{
		RawResponse: resp.RawResponse,
		Value:       resp.Value,
	}, nil
}

type ImportCertificateOptions struct {

	// The attributes of the certificate (optional).
	CertificateAttributes *CertificateAttributes `json:"attributes,omitempty"`

	// The management policy for the certificate.
	CertificatePolicy *CertificatePolicy `json:"policy,omitempty"`

	// If the private key in base64EncodedCertificate is encrypted, the password used for encryption.
	Password *string `json:"pwd,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

func (i *ImportCertificateOptions) toGenerated() *generated.KeyVaultClientImportCertificateOptions {
	return &generated.KeyVaultClientImportCertificateOptions{}
}

type ImportCertificateResponse struct {
	CertificateBundle

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

func (c *Client) ImportCertificate(ctx context.Context, certName string, base64EncodedCertificate string, options *ImportCertificateOptions) (ImportCertificateResponse, error) {
	if options == nil {
		options = &ImportCertificateOptions{}
	}
	resp, err := c.genClient.ImportCertificate(
		ctx,
		c.vaultURL,
		certName,
		generated.CertificateImportParameters{
			Base64EncodedCertificate: &base64EncodedCertificate,
			CertificateAttributes:    options.CertificateAttributes.toGenerated(),
			CertificatePolicy:        options.CertificatePolicy.toGeneratedCertificateCreateParameters(),
			Password:                 options.Password,
			Tags:                     options.Tags,
		},
		options.toGenerated(),
	)
	if err != nil {
		return ImportCertificateResponse{}, err
	}

	return ImportCertificateResponse{
		RawResponse: resp.RawResponse,
		CertificateBundle: CertificateBundle{
			Attributes:     certificateAttributesFromGenerated(resp.Attributes),
			Cer:            resp.Cer,
			ContentType:    resp.ContentType,
			Tags:           resp.Tags,
			ID:             resp.ID,
			Kid:            resp.Kid,
			Policy:         certificatePolicyFromGenerated(resp.Policy),
			Sid:            resp.Sid,
			X509Thumbprint: resp.X509Thumbprint,
		},
	}, nil
}
