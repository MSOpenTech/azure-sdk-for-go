//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	testframework "github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal"
	"github.com/stretchr/testify/require"
	"io"
	"math/rand"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

// Index
// 0. AuthN/AuthZ
// 1. ServiceClient
// 2. ContainerClient
// 3. BlobClient
// 4. Data Generators
// 5. Utility functions

const (
	DefaultEndpointSuffix       = "core.windows.net/"
	DefaultBlobEndpointSuffix   = "blob.core.windows.net/"
	AccountNameEnvVar           = "AZURE_STORAGE_ACCOUNT_NAME"
	AccountKeyEnvVar            = "AZURE_STORAGE_ACCOUNT_KEY"
	DefaultEndpointSuffixEnvVar = "AZURE_STORAGE_ENDPOINT_SUFFIX"
)

const (
	containerPrefix             = "goc"
	blobPrefix                  = "gotestblob"
	blockBlobDefaultData        = "GoBlockBlobData"
	invalidHeaderErrorSubstring = "invalid header field" // error thrown by the http client
)

var (
	blobContentType        = "my_type"
	blobContentDisposition = "my_disposition"
	blobCacheControl       = "control"
	blobContentLanguage    = "my_language"
	blobContentEncoding    = "my_encoding"
)

var basicHeaders = BlobHTTPHeaders{
	BlobContentType:        &blobContentType,
	BlobContentDisposition: &blobContentDisposition,
	BlobCacheControl:       &blobCacheControl,
	BlobContentMD5:         nil,
	BlobContentLanguage:    &blobContentLanguage,
	BlobContentEncoding:    &blobContentEncoding,
}

var basicMetadata = map[string]string{"Foo": "bar"}

//nolint
var basicBlobTagsMap = map[string]string{
	"azure": "blob",
	"blob":  "sdk",
	"sdk":   "go",
}

//nolint
var specialCharBlobTagsMap = map[string]string{
	"+-./:=_ ":        "firsttag",
	"tag2":            "+-./:=_",
	"+-./:=_1":        "+-./:=_",
	"Microsoft Azure": "Azure Storage",
	"Storage+SDK":     "SDK/GO",
	"GO ":             ".Net",
}

// 1. AuthN/AuthZ --------------------------------------------------------------------------------------------------

type testAccountType string

const (
	testAccountDefault   testAccountType = ""
	testAccountSecondary testAccountType = "SECONDARY_"
	testAccountPremium   testAccountType = "PREMIUM_"
	//testAccountBlobStorage testAccountType = "BLOB_"
)

// getRequiredEnv gets an environment variable by name and returns an error if it is not found
func getRequiredEnv(name string) (string, error) {
	env, ok := os.LookupEnv(name)
	if ok {
		return env, nil
	} else {
		return "", errors.New("Required environment variable not set: " + name)
	}
}

func getAccountInfo(recording *testframework.Recording, accountType testAccountType) (string, string) {
	accountNameEnvVar := string(accountType) + AccountNameEnvVar
	accountKeyEnvVar := string(accountType) + AccountKeyEnvVar
	var err error
	accountName, accountKey := "", ""
	if recording == nil {
		accountName, err = getRequiredEnv(accountNameEnvVar)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		_ = err
		accountKey, err = getRequiredEnv(accountKeyEnvVar)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		_ = err
	} else {
		accountName, err = recording.GetEnvVar(accountNameEnvVar, testframework.NoSanitization)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		_ = err
		accountKey, err = recording.GetEnvVar(accountKeyEnvVar, testframework.Secret_Base64String)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		_ = err
	}
	return accountName, accountKey
}

func getGenericCredential(recording *testframework.Recording, accountType testAccountType) (*SharedKeyCredential, error) {
	accountName, accountKey := getAccountInfo(recording, accountType)
	if accountName == "" || accountKey == "" {
		return nil, errors.New(string(accountType) + AccountNameEnvVar + " and/or " + string(accountType) + AccountKeyEnvVar + " environment variables not specified.")
	}
	return NewSharedKeyCredential(accountName, accountKey)
}

func getServiceClient(recording *testframework.Recording, accountType testAccountType, options *ClientOptions) (*ServiceClient, error) {
	if recording != nil {
		if options == nil {
			options = &ClientOptions{
				Transport: recording,
				Retry:     policy.RetryOptions{MaxRetries: -1},
			}
		}
	}

	cred, err := getGenericCredential(recording, accountType)
	if err != nil {
		return nil, err
	}

	serviceURL, _ := url.Parse("https://" + cred.AccountName() + ".blob.core.windows.net/")
	serviceClient := NewServiceClientWithSharedKey(serviceURL.String(), cred, options)

	return serviceClient, err
}

//nolint
func getConnectionString(recording *testframework.Recording, accountType testAccountType) string {
	accountName, accountKey := getAccountInfo(recording, accountType)
	connectionString := fmt.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=core.windows.net/",
		accountName, accountKey)
	return connectionString
}

//nolint
func getServiceClientFromConnectionString(recording *testframework.Recording, accountType testAccountType, options *ClientOptions) (*ServiceClient, error) {
	if recording != nil {
		if options == nil {
			options = &ClientOptions{
				Transport: recording,
				Retry:     policy.RetryOptions{MaxRetries: -1},
			}
		}
	}

	connectionString := getConnectionString(recording, accountType)
	primaryURL, cred, err := parseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}

	svcClient := NewServiceClientWithSharedKey(primaryURL, cred, options)
	return svcClient, err
}

// 2. ContainerClient --------------------------------------------------------------------------------------------------

func getContainerClient(containerName string, s *ServiceClient) *ContainerClient {
	return s.NewContainerClient(containerName)
}

func createNewContainer(_require *require.Assertions, containerName string, serviceClient *ServiceClient) *ContainerClient {
	containerClient := getContainerClient(containerName, serviceClient)

	cResp, err := containerClient.Create(ctx, nil)
	_require.Nil(err)
	_require.Equal(cResp.RawResponse.StatusCode, 201)
	return containerClient
}

func deleteContainer(_require *require.Assertions, containerClient *ContainerClient) {
	deleteContainerResp, err := containerClient.Delete(context.Background(), nil)
	_require.Nil(err)
	_require.Equal(deleteContainerResp.RawResponse.StatusCode, 202)
}

// 2. BlobClient -------------------------------------------------------------------------------------------------------

func createNewBlobs(_require *require.Assertions, blobNames []string, containerClient *ContainerClient) {
	for _, blobName := range blobNames {
		createNewBlockBlob(_require, blobName, containerClient)
	}
}

// 2a. BlockBlobClient -------------------------------------------------------------------------------------------------------

func getBlockBlobClient(blockBlobName string, containerClient *ContainerClient) *BlockBlobClient {
	return containerClient.NewBlockBlobClient(blockBlobName)
}

func createNewBlockBlob(_require *require.Assertions, blockBlobName string, containerClient *ContainerClient) *BlockBlobClient {
	bbClient := getBlockBlobClient(blockBlobName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)

	_require.Nil(err)
	_require.Equal(cResp.RawResponse.StatusCode, 201)

	return bbClient
}

func createNewBlockBlobWithCPK(_require *require.Assertions, blockBlobName string, containerClient *ContainerClient, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo) (bbClient *BlockBlobClient) {
	bbClient = getBlockBlobClient(blockBlobName, containerClient)

	uploadBlockBlobOptions := BlockBlobUploadOptions{
		CpkInfo:      cpkInfo,
		CpkScopeInfo: cpkScopeInfo,
	}
	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), &uploadBlockBlobOptions)
	_require.Nil(err)
	_require.Equal(cResp.RawResponse.StatusCode, 201)
	_require.Equal(*cResp.IsServerEncrypted, true)
	if cpkInfo != nil {
		_require.EqualValues(cResp.EncryptionKeySHA256, cpkInfo.EncryptionKeySHA256)
	}
	if cpkScopeInfo != nil {
		_require.EqualValues(cResp.EncryptionScope, cpkScopeInfo.EncryptionScope)
	}
	return
}

// 2b. AppendBlobClient -------------------------------------------------------------------------------------------------------

func getAppendBlobClient(appendBlobName string, containerClient *ContainerClient) *AppendBlobClient {
	return containerClient.NewAppendBlobClient(appendBlobName)
}

func createNewAppendBlob(_require *require.Assertions, appendBlobName string, containerClient *ContainerClient) *AppendBlobClient {
	abClient := getAppendBlobClient(appendBlobName, containerClient)

	appendBlobCreateResp, err := abClient.Create(ctx, nil)

	_require.Nil(err)
	_require.Equal(appendBlobCreateResp.RawResponse.StatusCode, 201)
	return abClient
}

// 2c. PageBlobClient -------------------------------------------------------------------------------------------------------

func getPageBlobClient(pageBlobName string, containerClient *ContainerClient) *PageBlobClient {
	return containerClient.NewPageBlobClient(pageBlobName)
}

func createNewPageBlob(_require *require.Assertions, pageBlobName string, containerClient *ContainerClient) *PageBlobClient {
	return createNewPageBlobWithSize(_require, pageBlobName, containerClient, internal.PageBlobPageBytes*10)
}

func createNewPageBlobWithSize(_require *require.Assertions, pageBlobName string,
	containerClient *ContainerClient, sizeInBytes int64) *PageBlobClient {
	pbClient := getPageBlobClient(pageBlobName, containerClient)

	pageBlobCreateResponse, err := pbClient.Create(ctx, sizeInBytes, nil)
	_require.Nil(err)
	_require.Equal(pageBlobCreateResponse.RawResponse.StatusCode, 201)
	return pbClient
}

func createNewPageBlobWithCPK(_require *require.Assertions, pageBlobName string, container *ContainerClient, sizeInBytes int64, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo) (pbClient *PageBlobClient) {
	pbClient = getPageBlobClient(pageBlobName, container)

	resp, err := pbClient.Create(ctx, sizeInBytes, &PageBlobCreateOptions{
		CpkInfo:      cpkInfo,
		CpkScopeInfo: cpkScopeInfo,
	})
	_require.Nil(err)
	_require.Equal(resp.RawResponse.StatusCode, 201)
	return
}

// 4. Data Generators --------------------------------------------------------------------------------------------------

// This function generates an entity name by concatenating the passed prefix,
// the name of the test requesting the entity name, and the minute, second, and nanoseconds of the call.
// This should make it easy to associate the entities with their test, uniquely identify
// them, and determine the order in which they were created.
// Note that this imposes a restriction on the length of test names
//nolint
func generateName(prefix string) string {
	// These next lines up through the for loop are obtaining and walking up the stack
	// trace to extract the test name, which is stored in name
	pc := make([]uintptr, 10)
	runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc)
	name := ""
	for f, next := frames.Next(); next; f, next = frames.Next() {
		name = f.Function
		if strings.Contains(name, "Suite") {
			break
		}
	}
	funcNameStart := strings.Index(name, "Test")
	name = name[funcNameStart+len("Test"):] // Just get the name of the test and not any of the garbage at the beginning
	name = strings.ToLower(name)            // Ensure it is a valid resource name
	currentTime := time.Now()
	name = fmt.Sprintf("%s%s%d%d%d", prefix, strings.ToLower(name), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond())
	return name
}

func generateEntityName(testName string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(testName), "/", ""), "test", "")
}
func generateContainerName(testName string) string {
	return containerPrefix + generateEntityName(testName)
}

func generateBlobName(testName string) string {
	return blobPrefix + generateEntityName(testName)
}

func getReaderToGeneratedBytes(n int) io.ReadSeekCloser {
	r, _ := generateData(n)
	return internal.NopCloser(r)
}

//nolint
func getRandomDataAndReader(n int) (*bytes.Reader, []byte) {
	data := make([]byte, n)
	rand.Read(data)
	return bytes.NewReader(data), data
}

const random64BString string = "2SDgZj6RkKYzJpu04sweQek4uWHO8ndPnYlZ0tnFS61hjnFZ5IkvIGGY44eKABov"

func generateData(sizeInBytes int) (io.ReadSeekCloser, []byte) {
	data := make([]byte, sizeInBytes)
	_len := len(random64BString)
	if sizeInBytes > _len {
		count := sizeInBytes / _len
		if sizeInBytes%_len != 0 {
			count = count + 1
		}
		copy(data[:], strings.Repeat(random64BString, count))
	} else {
		copy(data[:], random64BString)
	}
	return internal.NopCloser(bytes.NewReader(data)), data
}

// 5. Utility Functions ------------------------------------------------------------------------------------------------

//nolint
func getRelativeTimeGMT(amount time.Duration) time.Time {
	currentTime := time.Now().In(time.FixedZone("GMT", 0))
	currentTime = currentTime.Add(amount * time.Second)
	return currentTime
}

func getRelativeTimeFromAnchor(anchorTime *time.Time, amount time.Duration) time.Time {
	return anchorTime.Add(amount * time.Second)
}

func generateBlockIDsList(count int) []string {
	blockIDs := make([]string, count)
	for i := 0; i < count; i++ {
		blockIDs[i] = blockIDIntToBase64(i)
	}
	return blockIDs
}

// blockIDIntToBase64 functions convert an int block ID to a base-64 string and vice versa
func blockIDIntToBase64(blockID int) string {
	binaryBlockID := (&[4]byte{})[:]
	binary.LittleEndian.PutUint32(binaryBlockID, uint32(blockID))
	return base64.StdEncoding.EncodeToString(binaryBlockID)
}

func blobListToMap(list []string) map[string]bool {
	out := make(map[string]bool)

	for _, v := range list {
		out[v] = true
	}

	return out
}
