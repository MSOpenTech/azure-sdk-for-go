//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package recording

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type sanitizerTests struct {
	suite.Suite
}

const authHeader string = "Authorization"
const customHeader1 string = "Fooheader"
const customHeader2 string = "Barheader"
const nonSanitizedHeader string = "notsanitized"

func TestRecordingSanitizer(t *testing.T) {
	suite.Run(t, new(sanitizerTests))
}

func (s *sanitizerTests) TestDefaultSanitizerSanitizesAuthHeader() {
	require := require.New(s.T())
	server, cleanup := mock.NewServer()
	server.SetResponse()
	defer cleanup()
	rt := NewMockRoundTripper(server)
	r, _ := recorder.NewAsMode(getTestFileName(s.T(), false), recorder.ModeRecording, rt)

	defaultSanitizer(r)

	req, _ := http.NewRequest(http.MethodPost, server.URL(), nil)
	req.Header.Add(authHeader, "superSecret")

	_, err := r.RoundTrip(req)
	require.NoError(err)
	err = r.Stop()
	require.NoError(err)

	require.Equal(SanitizedValue, req.Header.Get(authHeader))

	rec, err := cassette.Load(getTestFileName(s.T(), false))
	require.NoError(err)

	for _, i := range rec.Interactions {
		require.Equal(SanitizedValue, i.Request.Headers.Get(authHeader))
	}
}

func (s *sanitizerTests) TestAddSanitizedHeadersSanitizes() {
	require := require.New(s.T())
	server, cleanup := mock.NewServer()
	server.SetResponse()
	defer cleanup()
	rt := NewMockRoundTripper(server)
	r, _ := recorder.NewAsMode(getTestFileName(s.T(), false), recorder.ModeRecording, rt)

	target := defaultSanitizer(r)
	target.AddSanitizedHeaders(customHeader1, customHeader2)

	req, _ := http.NewRequest(http.MethodPost, server.URL(), nil)
	req.Header.Add(customHeader1, "superSecret")
	req.Header.Add(customHeader2, "verySecret")
	safeValue := "safeValue"
	req.Header.Add(nonSanitizedHeader, safeValue)

	_, err := r.RoundTrip(req)
	require.NoError(err)
	err = r.Stop()
	require.NoError(err)

	require.Equal(SanitizedValue, req.Header.Get(customHeader1))
	require.Equal(SanitizedValue, req.Header.Get(customHeader2))
	require.Equal(safeValue, req.Header.Get(nonSanitizedHeader))

	rec, err := cassette.Load(getTestFileName(s.T(), false))
	require.NoError(err)

	for _, i := range rec.Interactions {
		require.Equal(SanitizedValue, i.Request.Headers.Get(customHeader1))
		require.Equal(SanitizedValue, i.Request.Headers.Get(customHeader2))
		require.Equal(safeValue, i.Request.Headers.Get(nonSanitizedHeader))
	}
}

func (s *sanitizerTests) TestAddUrlSanitizerSanitizes() {
	require := require.New(s.T())
	secret := "secretvalue"
	secretBody := "some body content that contains a " + secret
	server, cleanup := mock.NewServer()
	server.SetResponse(mock.WithStatusCode(http.StatusCreated), mock.WithBody([]byte(secretBody)))
	defer cleanup()
	rt := NewMockRoundTripper(server)
	r, _ := recorder.NewAsMode(getTestFileName(s.T(), false), recorder.ModeRecording, rt)

	baseUrl := server.URL() + "/"

	target := defaultSanitizer(r)
	target.AddUrlSanitizer(func(url *string) {
		*url = strings.Replace(*url, secret, SanitizedValue, -1)
	})
	target.AddBodysanitizer(func(body *string) {
		*body = strings.Replace(*body, secret, SanitizedValue, -1)
	})

	req, _ := http.NewRequest(http.MethodPost, baseUrl+secret, closerFromString(secretBody))

	_, err := r.RoundTrip(req)
	require.NoError(err)
	err = r.Stop()
	require.NoError(err)

	rec, err := cassette.Load(getTestFileName(s.T(), false))
	require.NoError(err)

	for _, i := range rec.Interactions {
		require.NotContains(i.Response.Body, secret)
		require.NotContains(i.Request.URL, secret)
		require.NotContains(i.Request.Body, secret)
		require.Contains(i.Request.URL, SanitizedValue)
		require.Contains(i.Request.Body, SanitizedValue)
		require.Contains(i.Response.Body, SanitizedValue)
	}
}

func (s *sanitizerTests) TearDownSuite() {
	require := require.New(s.T())
	// cleanup test files
	err := os.RemoveAll("testfiles")
	require.NoError(err)
}

func getTestFileName(t *testing.T, addSuffix bool) string {
	name := "testfiles/" + t.Name()
	if addSuffix {
		name = name + ".yaml"
	}
	return name
}

type mockRoundTripper struct {
	server *mock.Server
}

func NewMockRoundTripper(server *mock.Server) *mockRoundTripper {
	return &mockRoundTripper{server: server}
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.server.Do(req)
}

func TestUriSanitizer(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	err = AddUriSanitizer("replacement", "bing", nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://www.bing.com/")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)

	require.Equal(t, data.Entries[0].RequestUri, "https://www.replacement.com/")
}

func TestHeaderRegexSanitizer(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	err = AddUriSanitizer("replacement", "bing", nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://www.bing.com/")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))
	req.Header.Set("testproxy-header", "fakevalue")
	req.Header.Set("FakeStorageLocation", "https://fakeaccount.blob.core.windows.net")
	req.Header.Set("ComplexRegex", "https://fakeaccount.table.core.windows.net")

	err = AddHeaderRegexSanitizer("testproxy-header", "Sanitized", "", "", nil)
	require.NoError(t, err)

	err = AddHeaderRegexSanitizer("FakeStorageLocation", "Sanitized", "https\\:\\/\\/(?<account>[a-z]+)\\.blob\\.core\\.windows\\.net", "", nil)
	require.NoError(t, err)

	err = AddHeaderRegexSanitizer("ComplexRegex", "Sanitized", "https\\:\\/\\/(?<account>[a-z]+)\\.(?:table|blob|queue)\\.core\\.windows\\.net", "account", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)

	require.Equal(t, data.Entries[0].RequestHeaders["testproxy-header"], "Sanitized")
	require.Equal(t, data.Entries[0].RequestHeaders["fakestoragelocation"], "Sanitized")
	require.Equal(t, data.Entries[0].RequestHeaders["complexregex"], "https://fakeaccount.table.core.windows.net")
}

func TestBodyKeySanitizer(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	err = AddUriSanitizer("replacement", "bing", nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://jsonplaceholder.typicode.com/posts/1")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))

	bodyValue := map[string]string{
		"key1": "value1",
	}
	marshalled, err := json.Marshal(bodyValue)
	require.NoError(t, err)

	req.Body = ioutil.NopCloser(bytes.NewReader(marshalled))

	err = AddBodyKeySanitizer("$.key1", "Sanitized", "", "", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)
}

func TestBodyRegexSanitizer(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	err = AddUriSanitizer("replacement", "bing", nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://jsonplaceholder.typicode.com/posts/1")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))

	bodyValue := map[string]string{
		"key1": "value1",
	}
	marshalled, err := json.Marshal(bodyValue)
	require.NoError(t, err)

	req.Body = ioutil.NopCloser(bytes.NewReader(marshalled))

	err = AddBodyKeySanitizer("$.key1", "Sanitized", "", "", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)
}

func reset(t *testing.T) {
	err := ResetSanitizers(nil)
	require.NoError(t, err)
}

func TestRemoveHeaderSanitizer(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://jsonplaceholder.typicode.com/posts/1")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))
	req.Header.Set("FakeStorageLocation", "https://fakeaccount.blob.core.windows.net")
	req.Header.Set("ComplexRegexRemove", "https://fakeaccount.table.core.windows.net")

	err = AddRemoveHeaderSanitizer([]string{"ComplexRegexRemove", "FakeStorageLocation"}, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)

	require.Equal(t, data.Entries[0].RequestHeaders["fakestoragelocation"], "")
	require.Equal(t, data.Entries[0].RequestHeaders["complexregexremove"], "")
}

func TestContinuationSanitizer(t *testing.T) {

}

func TestGeneralRegexSanitizer(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://jsonplaceholder.typicode.com/posts")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))

	err = AddGeneralRegexSanitizer("Sanitized", "sunt (<replace> aut) facere", "replace", nil)
	require.NoError(t, err)

	_, err = client.Do(req)
	require.NoError(t, err)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)

	// Header should be there because all sanitizers were removed
	require.Equal(t, data.Entries[0].RequestHeaders["fakestoragelocation"], "https://fakeaccount.blob.core.windows.net")

}

func TestOAuthResponseSanitizer(t *testing.T) {

}

func TestReplaceRequestSubscriptionId(t *testing.T) {

}

func TestResetSanitizers(t *testing.T) {
	os.Setenv("AZURE_RECORD_MODE", "record")
	defer os.Unsetenv("AZURE_RECORD_MODE")
	defer reset(t)

	err := ResetSanitizers(nil)
	require.NoError(t, err)

	err = StartRecording(t, packagePath, nil)
	require.NoError(t, err)

	client, err := GetHTTPClient(t)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "https://localhost:5001", nil)
	require.NoError(t, err)

	req.Header.Set(UpstreamUriHeader, "https://jsonplaceholder.typicode.com/posts/1")
	req.Header.Set(ModeHeader, GetRecordMode())
	req.Header.Set(IdHeader, GetRecordingId(t))
	req.Header.Set("FakeStorageLocation", "https://fakeaccount.blob.core.windows.net")

	// Add a sanitizer
	err = AddRemoveHeaderSanitizer([]string{"FakeStorageLocation"}, nil)
	require.NoError(t, err)

	// Remove all sanitizers
	err = ResetSanitizers(nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotNil(t, GetRecordingId(t))

	err = StopRecording(t, nil)
	require.NoError(t, err)

	// Make sure the file is there
	jsonFile, err := os.Open(fmt.Sprintf("./recordings/%s.json", t.Name()))
	require.NoError(t, err)
	defer jsonFile.Close()

	var data RecordingFileStruct
	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	err = json.Unmarshal(byteValue, &data)
	require.NoError(t, err)

	// Header should be there because all sanitizers were removed
	require.Equal(t, data.Entries[0].RequestHeaders["fakestoragelocation"], "https://fakeaccount.blob.core.windows.net")
}
