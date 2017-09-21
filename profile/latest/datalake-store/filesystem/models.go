// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: b19c730e3a5c747d9055c95884b5c0310f7f2f16

package filesystem

import original "github.com/Azure/azure-sdk-for-go/service/datalake-store/2016-11-01/filesystem"

const (
	DefaultAdlsFileSystemDNSSuffix = original.DefaultAdlsFileSystemDNSSuffix
)

type ManagementClient = original.ManagementClient
type GroupClient = original.GroupClient
type AppendModeType = original.AppendModeType

const (
	Autocreate AppendModeType = original.Autocreate
)

type ExpiryOptionType = original.ExpiryOptionType

const (
	Absolute		ExpiryOptionType	= original.Absolute
	NeverExpire		ExpiryOptionType	= original.NeverExpire
	RelativeToCreationDate	ExpiryOptionType	= original.RelativeToCreationDate
	RelativeToNow		ExpiryOptionType	= original.RelativeToNow
)

type FileType = original.FileType

const (
	DIRECTORY	FileType	= original.DIRECTORY
	FILE		FileType	= original.FILE
)

type SyncFlag = original.SyncFlag

const (
	CLOSE		SyncFlag	= original.CLOSE
	DATA		SyncFlag	= original.DATA
	METADATA	SyncFlag	= original.METADATA
)

type ACLStatus = original.ACLStatus
type ACLStatusResult = original.ACLStatusResult
type AdlsAccessControlException = original.AdlsAccessControlException
type AdlsBadOffsetException = original.AdlsBadOffsetException
type AdlsError = original.AdlsError
type AdlsFileAlreadyExistsException = original.AdlsFileAlreadyExistsException
type AdlsFileNotFoundException = original.AdlsFileNotFoundException
type AdlsIllegalArgumentException = original.AdlsIllegalArgumentException
type AdlsIOException = original.AdlsIOException
type AdlsRemoteException = original.AdlsRemoteException
type AdlsRuntimeException = original.AdlsRuntimeException
type AdlsSecurityException = original.AdlsSecurityException
type AdlsThrottledException = original.AdlsThrottledException
type AdlsUnsupportedOperationException = original.AdlsUnsupportedOperationException
type ContentSummary = original.ContentSummary
type ContentSummaryResult = original.ContentSummaryResult
type FileOperationResult = original.FileOperationResult
type FileStatuses = original.FileStatuses
type FileStatusesResult = original.FileStatusesResult
type FileStatusProperties = original.FileStatusProperties
type FileStatusResult = original.FileStatusResult
type ReadCloser = original.ReadCloser

func New() ManagementClient {
	return original.New()
}
func NewWithoutDefaults(adlsFileSystemDNSSuffix string) ManagementClient {
	return original.NewWithoutDefaults(adlsFileSystemDNSSuffix)
}
func NewGroupClient() GroupClient {
	return original.NewGroupClient()
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
