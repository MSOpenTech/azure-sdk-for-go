//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// Contains common helpers for TESTS ONLY
package testcommon

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/service"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/share"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestAccountType string

const (
	TestAccountDefault    TestAccountType = ""
	TestAccountSecondary  TestAccountType = "SECONDARY_"
	TestAccountPremium    TestAccountType = "PREMIUM_"
	TestAccountSoftDelete TestAccountType = "SOFT_DELETE_"
)

const (
	DefaultEndpointSuffix       = "core.windows.net/"
	DefaultFileEndpointSuffix   = "file.core.windows.net/"
	AccountNameEnvVar           = "AZURE_STORAGE_ACCOUNT_NAME"
	AccountKeyEnvVar            = "AZURE_STORAGE_ACCOUNT_KEY"
	DefaultEndpointSuffixEnvVar = "AZURE_STORAGE_ENDPOINT_SUFFIX"
)

const (
	FakeStorageAccount = "fakestorage"
	FakeStorageURL     = "https://fakestorage.file.core.windows.net"
)

var (
	SampleSDDL = `O:S-1-5-32-548G:S-1-5-21-397955417-626881126-188441444-512D:(A;;RPWPCCDCLCSWRCWDWOGA;;;S-1-0-0)`
)

var BasicMetadata = map[string]*string{
	"foo": to.Ptr("foovalue"),
	"bar": to.Ptr("barvalue"),
}

func SetClientOptions(t *testing.T, opts *azcore.ClientOptions) {
	opts.Logging.AllowedHeaders = append(opts.Logging.AllowedHeaders, "X-Request-Mismatch", "X-Request-Mismatch-Error")

	transport, err := recording.NewRecordingHTTPClient(t, nil)
	require.NoError(t, err)
	opts.Transport = transport
}

func GetServiceClient(t *testing.T, accountType TestAccountType, options *service.ClientOptions) (*service.Client, error) {
	if options == nil {
		options = &service.ClientOptions{}
	}

	SetClientOptions(t, &options.ClientOptions)

	cred, err := GetGenericSharedKeyCredential(accountType)
	if err != nil {
		return nil, err
	}

	serviceClient, err := service.NewClientWithSharedKeyCredential("https://"+cred.AccountName()+".file.core.windows.net/", cred, options)

	return serviceClient, err
}

func GetServiceClientNoCredential(t *testing.T, sasUrl string, options *service.ClientOptions) (*service.Client, error) {
	if options == nil {
		options = &service.ClientOptions{}
	}

	SetClientOptions(t, &options.ClientOptions)

	serviceClient, err := service.NewClientWithNoCredential(sasUrl, options)

	return serviceClient, err
}

func GetGenericAccountInfo(accountType TestAccountType) (string, string) {
	if recording.GetRecordMode() == recording.PlaybackMode {
		return FakeStorageAccount, "ZmFrZQ=="
	}
	accountNameEnvVar := string(accountType) + AccountNameEnvVar
	accountKeyEnvVar := string(accountType) + AccountKeyEnvVar
	accountName, _ := GetRequiredEnv(accountNameEnvVar)
	accountKey, _ := GetRequiredEnv(accountKeyEnvVar)
	return accountName, accountKey
}

func GetGenericSharedKeyCredential(accountType TestAccountType) (*service.SharedKeyCredential, error) {
	accountName, accountKey := GetGenericAccountInfo(accountType)
	if accountName == "" || accountKey == "" {
		return nil, errors.New(string(accountType) + AccountNameEnvVar + " and/or " + string(accountType) + AccountKeyEnvVar + " environment variables not specified.")
	}
	return service.NewSharedKeyCredential(accountName, accountKey)
}

func GetGenericConnectionString(accountType TestAccountType) (*string, error) {
	accountName, accountKey := GetGenericAccountInfo(accountType)
	if accountName == "" || accountKey == "" {
		return nil, errors.New(string(accountType) + AccountNameEnvVar + " and/or " + string(accountType) + AccountKeyEnvVar + " environment variables not specified.")
	}
	connectionString := fmt.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=core.windows.net/",
		accountName, accountKey)
	return &connectionString, nil
}

func GetServiceClientFromConnectionString(t *testing.T, accountType TestAccountType, options *service.ClientOptions) (*service.Client, error) {
	if options == nil {
		options = &service.ClientOptions{}
	}
	SetClientOptions(t, &options.ClientOptions)

	transport, err := recording.NewRecordingHTTPClient(t, nil)
	require.NoError(t, err)
	options.Transport = transport

	cred, err := GetGenericConnectionString(accountType)
	if err != nil {
		return nil, err
	}
	svcClient, err := service.NewClientFromConnectionString(*cred, options)
	return svcClient, err
}

func GetShareClient(shareName string, s *service.Client) *share.Client {
	return s.NewShareClient(shareName)
}

func CreateNewShare(ctx context.Context, _require *require.Assertions, shareName string, svcClient *service.Client) *share.Client {
	shareClient := GetShareClient(shareName, svcClient)
	_, err := shareClient.Create(ctx, nil)
	_require.NoError(err)
	return shareClient
}

func DeleteShare(ctx context.Context, _require *require.Assertions, shareClient *share.Client) {
	_, err := shareClient.Delete(ctx, nil)
	_require.NoError(err)
}
