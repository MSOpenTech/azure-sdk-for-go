//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package file

import (
	"encoding/binary"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/generated"
)

const (
	_1MiB = 1024 * 1024

	// MaxUpdateRangeBytes indicates the maximum number of bytes that can be updated in a call to Client.UploadRange.
	MaxUpdateRangeBytes = 4 * 1024 * 1024 // 4MiB

	// MaxFileSize indicates the maximum size of the file allowed.
	MaxFileSize = 4 * 1024 * 1024 * 1024 * 1024 // 4 TiB
)

// CopyStatusType defines the states of the copy operation.
type CopyStatusType = generated.CopyStatusType

const (
	CopyStatusTypePending CopyStatusType = generated.CopyStatusTypePending
	CopyStatusTypeSuccess CopyStatusType = generated.CopyStatusTypeSuccess
	CopyStatusTypeAborted CopyStatusType = generated.CopyStatusTypeAborted
	CopyStatusTypeFailed  CopyStatusType = generated.CopyStatusTypeFailed
)

// PossibleCopyStatusTypeValues returns the possible values for the CopyStatusType const type.
func PossibleCopyStatusTypeValues() []CopyStatusType {
	return generated.PossibleCopyStatusTypeValues()
}

// PermissionCopyModeType determines the copy behavior of the security descriptor of the file.
//   - source: The security descriptor on the destination file is copied from the source file.
//   - override: The security descriptor on the destination file is determined via the x-ms-file-permission or x-ms-file-permission-key header.
type PermissionCopyModeType = generated.PermissionCopyModeType

const (
	PermissionCopyModeTypeSource   PermissionCopyModeType = generated.PermissionCopyModeTypeSource
	PermissionCopyModeTypeOverride PermissionCopyModeType = generated.PermissionCopyModeTypeOverride
)

// PossiblePermissionCopyModeTypeValues returns the possible values for the PermissionCopyModeType const type.
func PossiblePermissionCopyModeTypeValues() []PermissionCopyModeType {
	return generated.PossiblePermissionCopyModeTypeValues()
}

// RangeWriteType represents one of the following options.
//   - update: Writes the bytes specified by the request body into the specified range. The Range and Content-Length headers must match to perform the update.
//   - clear: Clears the specified range and releases the space used in storage for that range. To clear a range, set the Content-Length header to zero,
//     and set the Range header to a value that indicates the range to clear, up to maximum file size.
type RangeWriteType = generated.FileRangeWriteType

const (
	RangeWriteTypeUpdate RangeWriteType = generated.FileRangeWriteTypeUpdate
	RangeWriteTypeClear  RangeWriteType = generated.FileRangeWriteTypeClear
)

// PossibleRangeWriteTypeValues returns the possible values for the RangeWriteType const type.
func PossibleRangeWriteTypeValues() []RangeWriteType {
	return generated.PossibleFileRangeWriteTypeValues()
}

// TransferValidationType abstracts the various mechanisms used to verify a transfer.
type TransferValidationType = exported.TransferValidationType

// TransferValidationTypeMD5 is a TransferValidationType used to provide a precomputed MD5.
type TransferValidationTypeMD5 = exported.TransferValidationTypeMD5

// SourceContentValidationType abstracts the various mechanisms used to validate source content.
// This interface is not publicly implementable.
type SourceContentValidationType interface {
	Apply(generated.SourceContentSetter)
	notPubliclyImplementable()
}

// SourceContentValidationTypeCRC64 is a SourceContentValidationType used to provide a precomputed CRC64.
type SourceContentValidationTypeCRC64 uint64

// Apply implements the SourceContentValidationType interface for type SourceContentValidationTypeCRC64.
func (s SourceContentValidationTypeCRC64) Apply(src generated.SourceContentSetter) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(s))
	src.SetSourceContentCRC64(buf)
}

func (SourceContentValidationTypeCRC64) notPubliclyImplementable() {}
