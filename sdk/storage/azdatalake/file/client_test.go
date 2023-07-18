//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package file_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/datalakeerror"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/file"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/testcommon"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

var proposedLeaseIDs = []*string{to.Ptr("c820a799-76d7-4ee2-6e15-546f19325c2c"), to.Ptr("326cc5e1-746e-4af8-4811-a50e6629a8ca")}

func Test(t *testing.T) {
	recordMode := recording.GetRecordMode()
	t.Logf("Running datalake Tests in %s mode\n", recordMode)
	if recordMode == recording.LiveMode {
		suite.Run(t, &RecordedTestSuite{})
		suite.Run(t, &UnrecordedTestSuite{})
	} else if recordMode == recording.PlaybackMode {
		suite.Run(t, &RecordedTestSuite{})
	} else if recordMode == recording.RecordingMode {
		suite.Run(t, &RecordedTestSuite{})
	}
}

func (s *RecordedTestSuite) BeforeTest(suite string, test string) {
	testcommon.BeforeTest(s.T(), suite, test)
}

func (s *RecordedTestSuite) AfterTest(suite string, test string) {
	testcommon.AfterTest(s.T(), suite, test)
}

func (s *UnrecordedTestSuite) BeforeTest(suite string, test string) {

}

func (s *UnrecordedTestSuite) AfterTest(suite string, test string) {

}

type RecordedTestSuite struct {
	suite.Suite
}

type UnrecordedTestSuite struct {
	suite.Suite
}

func validateFileDeleted(_require *require.Assertions, fileClient *file.Client) {
	_, err := fileClient.GetAccessControl(context.Background(), nil)
	_require.NotNil(err)

	testcommon.ValidateErrorCode(_require, err, datalakeerror.BlobNotFound)
}

func (s *RecordedTestSuite) TestCreateFileAndDelete() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileWithNilAccessConditions() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	createFileOpts := &file.CreateOptions{
		AccessConditions: nil,
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileIfModifiedSinceTrue() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, -10)

	createFileOpts := &file.CreateOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfModifiedSince: &currentTime,
			},
		},
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileIfModifiedSinceFalse() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, 10)

	createFileOpts := &file.CreateOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfModifiedSince: &currentTime,
			},
		},
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.NotNil(err)
	testcommon.ValidateErrorCode(_require, err, datalakeerror.ConditionNotMet)
}

func (s *RecordedTestSuite) TestCreateFileIfUnmodifiedSinceTrue() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, 10)

	createFileOpts := &file.CreateOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfUnmodifiedSince: &currentTime,
			},
		},
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileIfUnmodifiedSinceFalse() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, -10)

	createFileOpts := &file.CreateOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfUnmodifiedSince: &currentTime,
			},
		},
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.NotNil(err)

	testcommon.ValidateErrorCode(_require, err, datalakeerror.ConditionNotMet)
}

func (s *RecordedTestSuite) TestCreateFileIfETagMatch() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	etag := resp.ETag

	createFileOpts := &file.CreateOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfMatch: etag,
			},
		},
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileIfETagMatchFalse() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	etag := resp.ETag

	createFileOpts := &file.CreateOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfNoneMatch: etag,
			},
		},
	}

	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.NotNil(err)

	testcommon.ValidateErrorCode(_require, err, datalakeerror.ConditionNotMet)
}

func (s *RecordedTestSuite) TestCreateFileWithMetadataNotNil() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		Metadata: testcommon.BasicMetadata,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileWithEmptyMetadata() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		Metadata: nil,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileWithNilHTTPHeaders() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		HTTPHeaders: nil,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileWithHTTPHeaders() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		HTTPHeaders: &testcommon.BasicHeaders,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileWithExpiryAbsolute() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	expiryTimeAbsolute := time.Now().Add(8 * time.Second)
	expiry := file.CreationExpiryTypeAbsolute(expiryTimeAbsolute)
	createFileOpts := &file.CreateOptions{
		Expiry: expiry,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	defer testcommon.DeleteFile(context.Background(), _require, fClient)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)

	resp1, err := fClient.GetProperties(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp1.ExpiresOn)
	_require.Equal(expiryTimeAbsolute.UTC().Format(http.TimeFormat), (*resp1.ExpiresOn).UTC().Format(http.TimeFormat))
}

func (s *RecordedTestSuite) TestCreateFileWithExpiryRelativeToNow() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	expiry := file.CreationExpiryTypeRelativeToNow(8 * time.Second)
	createFileOpts := &file.CreateOptions{
		Expiry: expiry,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)

	resp1, err := fClient.GetProperties(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp1.ExpiresOn)

	time.Sleep(time.Second * 10)
	_, err = fClient.GetProperties(context.Background(), nil)
	testcommon.ValidateErrorCode(_require, err, datalakeerror.BlobNotFound)
}

func (s *RecordedTestSuite) TestCreateFileWithNeverExpire() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		Expiry: file.CreationExpiryTypeNever{},
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)

	resp1, err := fClient.GetProperties(context.Background(), nil)
	_require.Nil(err)
	_require.Nil(resp1.ExpiresOn)
}

func (s *RecordedTestSuite) TestCreateFileWithLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		ProposedLeaseID: proposedLeaseIDs[0],
		LeaseDuration:   to.Ptr(int64(15)),
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)

	// should fail since leased
	_, err = fClient.Create(context.Background(), createFileOpts)
	_require.NotNil(err)

	time.Sleep(time.Second * 15)
	resp, err = fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestCreateFileWithPermissions() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	perms := "0777"
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		Permissions: &perms,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)

	//TODO: GetProperties() when you figured out how to add permissions into response
}

func (s *RecordedTestSuite) TestCreateFileWithOwnerGroupACLUmask() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	owner := "4cf4e284-f6a8-4540-b53e-c3469af032dc"
	group := owner
	acl := "user::rwx,group::r-x,other::rwx"
	umask := "0000"
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	createFileOpts := &file.CreateOptions{
		Owner: &owner,
		Group: &group,
		ACL:   &acl,
		Umask: &umask,
	}

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), createFileOpts)
	_require.Nil(err)
	_require.NotNil(resp)

	//TODO: GetProperties() when you figured out how to add o,g, ACL into response
}

func (s *RecordedTestSuite) TestDeleteFileWithNilAccessConditions() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	_, err = fClient.Create(context.Background(), nil)
	_require.Nil(err)

	deleteOpts := &file.DeleteOptions{
		AccessConditions: nil,
	}

	resp, err := fClient.Delete(context.Background(), deleteOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestDeleteFileIfModifiedSinceTrue() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, -10)

	deleteOpts := &file.DeleteOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfModifiedSince: &currentTime,
			},
		},
	}

	resp1, err := fClient.Delete(context.Background(), deleteOpts)
	_require.Nil(err)
	_require.NotNil(resp1)
}

func (s *RecordedTestSuite) TestDeleteFileIfModifiedSinceFalse() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, 10)

	deleteOpts := &file.DeleteOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfModifiedSince: &currentTime,
			},
		},
	}

	_, err = fClient.Delete(context.Background(), deleteOpts)
	_require.NotNil(err)
	testcommon.ValidateErrorCode(_require, err, datalakeerror.ConditionNotMet)
}

func (s *RecordedTestSuite) TestDeleteFileIfUnmodifiedSinceTrue() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, 10)

	deleteOpts := &file.DeleteOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfUnmodifiedSince: &currentTime,
			},
		},
	}

	_, err = fClient.Delete(context.Background(), deleteOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestDeleteFileIfUnmodifiedSinceFalse() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	currentTime := testcommon.GetRelativeTimeFromAnchor(resp.Date, -10)

	deleteOpts := &file.DeleteOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfUnmodifiedSince: &currentTime,
			},
		},
	}

	_, err = fClient.Delete(context.Background(), deleteOpts)
	_require.NotNil(err)

	testcommon.ValidateErrorCode(_require, err, datalakeerror.ConditionNotMet)
}

func (s *RecordedTestSuite) TestDeleteFileIfETagMatch() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	etag := resp.ETag

	deleteOpts := &file.DeleteOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfMatch: etag,
			},
		},
	}

	_, err = fClient.Delete(context.Background(), deleteOpts)
	_require.Nil(err)
	_require.NotNil(resp)
}

func (s *RecordedTestSuite) TestDeleteFileIfETagMatchFalse() {
	_require := require.New(s.T())
	testName := s.T().Name()

	filesystemName := testcommon.GenerateFilesystemName(testName)
	fsClient, err := testcommon.GetFilesystemClient(filesystemName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)
	defer testcommon.DeleteFilesystem(context.Background(), _require, fsClient)

	_, err = fsClient.Create(context.Background(), nil)
	_require.Nil(err)

	fileName := testcommon.GenerateFileName(testName)
	fClient, err := testcommon.GetFileClient(filesystemName, fileName, s.T(), testcommon.TestAccountDatalake, nil)
	_require.NoError(err)

	resp, err := fClient.Create(context.Background(), nil)
	_require.Nil(err)
	_require.NotNil(resp)

	etag := resp.ETag

	deleteOpts := &file.DeleteOptions{
		AccessConditions: &file.AccessConditions{
			ModifiedAccessConditions: &file.ModifiedAccessConditions{
				IfNoneMatch: etag,
			},
		},
	}

	_, err = fClient.Delete(context.Background(), deleteOpts)
	_require.NotNil(err)

	testcommon.ValidateErrorCode(_require, err, datalakeerror.ConditionNotMet)
}
