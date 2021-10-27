//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azkeys

import (
	"context"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

var pathToPackage = "sdk/keyvault/azkeys/testdata"

var fakeURL = "https://fakekvurl.vault.azure.net/"

func TestMain(m *testing.M) {
	// Initialize
	if recording.GetRecordMode() == "record" {
		vaultUrl := os.Getenv("AZURE_KEYVAULT_URL")
		err := recording.AddURISanitizer("https://fakekvurl.vault.azure.net/", vaultUrl, nil)
		if err != nil {
			panic(err)
		}
	}

	// Run tests
	exitVal := m.Run()

	// 3. Reset
	// TODO: Add after sanitizer PR
	err := recording.ResetSanitizers(nil)
	if err != nil {
		panic(err)
	}

	// 4. Error out if applicable
	os.Exit(exitVal)
}
func createRandomName(t *testing.T, prefix string) (string, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(t.Name()))
	return prefix + fmt.Sprint(h.Sum32()), err
}

type recordingPolicy struct {
	options recording.RecordingOptions
	t       *testing.T
}

func (r recordingPolicy) Host() string {
	if r.options.UseHTTPS {
		return "localhost:5001"
	}
	return "localhost:5000"
}

func (r recordingPolicy) Scheme() string {
	if r.options.UseHTTPS {
		return "https"
	}
	return "http"
}

func NewRecordingPolicy(t *testing.T, o *recording.RecordingOptions) policy.Policy {
	if o == nil {
		o = &recording.RecordingOptions{UseHTTPS: true}
	}
	p := &recordingPolicy{options: *o, t: t}
	return p
}

func (p *recordingPolicy) Do(req *policy.Request) (resp *http.Response, err error) {
	if recording.GetRecordMode() != "live" {
		originalURLHost := req.Raw().URL.Host
		req.Raw().URL.Scheme = p.Scheme()
		req.Raw().URL.Host = p.Host()
		req.Raw().Host = p.Host()

		req.Raw().Header.Set(recording.UpstreamURIHeader, fmt.Sprintf("%v://%v", p.Scheme(), originalURLHost))
		req.Raw().Header.Set(recording.ModeHeader, recording.GetRecordMode())
		req.Raw().Header.Set(recording.IDHeader, recording.GetRecordingId(p.t))
	}
	return req.Next()
}

func lookupEnvVar(s string) string {
	ret, ok := os.LookupEnv(s)
	if !ok {
		panic(fmt.Sprintf("Could not find env var: '%s'", s))
	}
	return ret
}

func createClient(t *testing.T, testType string) (*Client, error) {
	vaultUrl := recording.GetEnvVariable("AZURE_KEYVAULT_URL", fakeURL)
	var credOptions *azidentity.ClientSecretCredentialOptions
	if testType == HSMTEST {
		vaultUrl = recording.GetEnvVariable("AZURE_MANAGEDHSM_URL", fakeURL)
	}
	if recording.GetRecordMode() == "playback" {
		vaultUrl = fakeURL
	}

	p := NewRecordingPolicy(t, &recording.RecordingOptions{UseHTTPS: true})
	client, err := recording.GetHTTPClient(t)
	require.NoError(t, err)

	options := &ClientOptions{
		ClientOptions: azcore.ClientOptions{
			PerCallPolicies: []policy.Policy{p},
			Transport:       client,
		},
	}

	var cred azcore.TokenCredential
	if recording.GetRecordMode() != "playback" {
		tenantId := lookupEnvVar("KEYVAULT_TENANT_ID")
		clientId := lookupEnvVar("KEYVAULT_CLIENT_ID")
		clientSecret := lookupEnvVar("KEYVAULT_CLIENT_SECRET")
		cred, err = azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, credOptions)
		require.NoError(t, err)
	} else {
		cred = NewFakeCredential("fake", "fake")
	}

	return NewClient(vaultUrl, cred, options)
}

func delay() time.Duration {
	if recording.GetRecordMode() == "playback" {
		return 1 * time.Microsecond
	}
	return 250 * time.Millisecond
}

func cleanUpKey(t *testing.T, client *Client, key string) {
	resp, err := client.BeginDeleteKey(context.Background(), key, nil)
	if err != nil {
		return
	}

	_, err = resp.PollUntilDone(context.Background(), delay())
	require.NoError(t, err)

	_, err = client.PurgeDeletedKey(context.Background(), key, nil)
	require.NoError(t, err)
}

type FakeCredential struct {
	accountName string
	accountKey  string
}

func NewFakeCredential(accountName, accountKey string) *FakeCredential {
	return &FakeCredential{
		accountName: accountName,
		accountKey:  accountKey,
	}
}

func (f *FakeCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (*azcore.AccessToken, error) {
	return &azcore.AccessToken{
		Token:     "faketoken",
		ExpiresOn: time.Date(2040, time.January, 1, 1, 1, 1, 1, time.UTC),
	}, nil
}

func toBytes(s string, t *testing.T) []byte {
	if len(s)%2 == 1 {
		s = fmt.Sprintf("0%s", s)
	}
	ret, err := hex.DecodeString(s)
	require.NoError(t, err)
	return ret
}
