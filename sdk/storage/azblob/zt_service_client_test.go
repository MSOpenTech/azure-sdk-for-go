// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/stretchr/testify/assert"
	"strings"
	"time"
)

func (s *azblobTestSuite) TestGetAccountInfo() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	// Ensure the call succeeded. Don't test for specific account properties because we can't/don't want to set account properties.
	sAccInfo, err := svcClient.GetAccountInfo(context.Background())
	_assert.Nil(err)
	_assert.NotEqualValues(sAccInfo, ServiceGetAccountInfoResponse{})

	//Test on a container
	containerClient := svcClient.NewContainerClient(generateContainerName(testName))
	_, err = containerClient.Create(ctx, nil)
	defer containerClient.Delete(ctx, nil)
	_assert.Nil(err)

	cAccInfo, err := containerClient.GetAccountInfo(ctx)
	_assert.Nil(err)
	_assert.NotEqualValues(cAccInfo, ContainerGetAccountInfoResponse{})

	// test on a block blob URL. They all call the same thing on the base URL, so only one test is needed for that.
	blobClient := containerClient.NewBlockBlobClient(generateBlobName(testName))
	_, err = blobClient.Upload(ctx, azcore.NopCloser(strings.NewReader("blah")), nil)
	_assert.Nil(err)
	bAccInfo, err := blobClient.GetAccountInfo(ctx)
	_assert.Nil(err)
	_assert.NotEqualValues(bAccInfo, BlobGetAccountInfoResponse{})
}

func (s *azblobTestSuite) TestListContainersBasic() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	md := map[string]string{
		"foo": "foovalue",
		"bar": "barvalue",
	}

	containerName := generateContainerName(testName)
	containerClient := getContainerClient(containerName, svcClient)
	_, err = containerClient.Create(ctx, &CreateContainerOptions{Metadata: md})
	defer containerClient.Delete(ctx, nil)
	_assert.Nil(err)

	prefix := containerPrefix
	listOptions := ListContainersSegmentOptions{Prefix: &prefix, Include: ListContainersDetail{Metadata: true}}
	pager := svcClient.ListContainersSegment(&listOptions)

	count := 0
	for pager.NextPage(ctx) {
		resp := pager.PageResponse()

		for _, container := range resp.EnumerationResults.ContainerItems {
			_assert.NotNil(container.Name)

			if *container.Name == containerName {
				_assert.NotNil(container.Properties)
				_assert.NotNil(container.Properties.LastModified)
				_assert.NotNil(container.Properties.Etag)
				_assert.Equal(*container.Properties.LeaseStatus, LeaseStatusTypeUnlocked)
				_assert.Equal(*container.Properties.LeaseState, LeaseStateTypeAvailable)
				_assert.Nil(container.Properties.LeaseDuration)
				_assert.Nil(container.Properties.PublicAccess)
				_assert.NotNil(container.Metadata)

				unwrappedMeta := map[string]string{}
				for k, v := range container.Metadata {
					if v != nil {
						unwrappedMeta[k] = *v
					}
				}

				_assert.EqualValues(unwrappedMeta, md)
			}
		}
	}

	_assert.Nil(pager.Err())

	_assert.Nil(err)
	_assert.GreaterOrEqual(count, 0)
}

//func (s *azblobTestSuite) TestListContainersPaged() {
//	_assert := assert.New(s.T())
//	testName := s.T().Name()
//	_context := getTestContext(testName)
//	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
//	if err != nil {
//		s.Fail("Unable to fetch service client because " + err.Error())
//	}
//
//	const numContainers = 6
//	maxResults := int32(2)
//	const pagedContainersPrefix = "azcontainerpaged"
//
//	containers := make([]ContainerClient, numContainers)
//	expectedResults := make(map[string]bool)
//	for i := 0; i < numContainers; i++ {
//		containerName := pagedContainersPrefix + generateContainerName(testName) + string(i)
//		containerClient := createNewContainer(_assert, containerName, svcClient)
//		containers[i] = containerClient
//		expectedResults[containerName] = false
//	}
//
//	defer func() {
//		for i := range containers {
//			deleteContainer(_assert, containers[i])
//		}
//	}()
//
//	// list for a first time
//	prefix := containerPrefix + pagedContainersPrefix
//	listOptions := ListContainersSegmentOptions{MaxResults: &maxResults, Prefix: &prefix}
//	End := 0
//	results := make([]ContainerItem, 0)
//
//	pager := sa.ListContainersSegment(&listOptions)
//
//	for pager.NextPage(ctx) {
//		for _, container := range *pager.PageResponse().EnumerationResults.ContainerItems {
//			if container == nil {
//				continue
//			}
//
//			results = append(results, *container)
//			End += 1
//			_assert.(container.Name, chk.NotNil)
//		}
//	}
//
//	_assert.(pager.Err(), chk.IsNil)
//	_assert.(End, chk.Equals, numContainers)
//	_assert.(len(results), chk.Equals, numContainers)
//
//	// make sure each container we see is expected
//	for _, container := range results {
//		_, ok := expectedResults[*container.Name]
//		_assert.(ok, chk.Equals, true)
//
//		expectedResults[*container.Name] = true
//	}
//
//	// make sure every expected container was seen
//	for _, seen := range expectedResults {
//		_assert.(seen, chk.Equals, true)
//	}
//}

func (s *azblobTestSuite) TestAccountListContainersEmptyPrefix() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerClient1 := createNewContainer(_assert, generateContainerName(testName)+"1", svcClient)
	defer deleteContainer(_assert, containerClient1)
	containerClient2 := createNewContainer(_assert, generateContainerName(testName)+"2", svcClient)
	defer deleteContainer(_assert, containerClient2)

	count := 0
	pager := svcClient.ListContainersSegment(nil)

	for pager.NextPage(ctx) {
		resp := pager.PageResponse()

		for _, container := range resp.EnumerationResults.ContainerItems {
			count++
			_assert.NotNil(container.Name)
		}
	}
	_assert.Nil(pager.Err())
	_assert.GreaterOrEqual(count, 2)
}

//// TODO re-enable after fixing error handling
////func (s *azblobTestSuite) TestAccountListContainersMaxResultsNegative() {
////	svcClient := getServiceClient()
////	containerClient, _ := createNewContainer(c, svcClient)
////	defer deleteContainer(_assert, containerClient)
////
////	illegalMaxResults := []int32{-2, 0}
////	for _, num := range illegalMaxResults {
////		options := ListContainersSegmentOptions{MaxResults: &num}
////
////		// getting the pager should still work
////		pager, err := svcClient.ListContainersSegment(context.Background(), 100, time.Hour, &options)
////		_assert.Nil(err)
////
////		// getting the next page should fail
////
////	}
////}
//
////func (s *azblobTestSuite) TestAccountListContainersMaxResultsExact() {
////	// If this test fails, ensure there are no extra containers prefixed with go in the account. These may be left over if a test is interrupted.
////	svcClient := getServiceClient()
////	containerClient1, containerName1 := createNewContainerWithSuffix(c, svcClient, "abc")
////	defer deleteContainer(containerClient1)
////	containerClient2, containerName2 := createNewContainerWithSuffix(c, svcClient, "abcde")
////	defer deleteContainer(containerClient2)
////
////	prefix := containerPrefix + "abc"
////	maxResults := int32(2)
////	options := ListContainersSegmentOptions{Prefix: &prefix, MaxResults: &maxResults}
////	pager, err := svcClient.ListContainersSegment(&options)
////	_assert.Nil(err)
////
////	// getting the next page should work
////	hasPage := pager.NextPage(context.Background())
////	_assert.(hasPage, chk.Equals, true)
////
////	page := pager.PageResponse()
////	_assert.Nil(err)
////	_assert.(*page.EnumerationResults.ContainerItems, chk.HasLen, 2)
////	_assert.(*(*page.EnumerationResults.ContainerItems)[0].Name, chk.DeepEquals, containerName1)
////	_assert.(*(*page.EnumerationResults.ContainerItems)[1].Name, chk.DeepEquals, containerName2)
////}

func (s *azblobTestSuite) TestAccountDeleteRetentionPolicy() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	days := to.Int32Ptr(5)
	enabled := to.BoolPtr(true)
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: enabled, Days: days}})
	_assert.Nil(err)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err := svcClient.GetProperties(ctx)
	_assert.Nil(err)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, *enabled)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, *days)

	disabled := false
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &disabled}})
	_assert.Nil(err)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err = svcClient.GetProperties(ctx)
	_assert.Nil(err)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, false)
	_assert.Nil(resp.StorageServiceProperties.DeleteRetentionPolicy.Days)
}

func (s *azblobTestSuite) TestAccountDeleteRetentionPolicyEmpty() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	days := to.Int32Ptr(5)
	enabled := to.BoolPtr(true)
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: enabled, Days: days}})
	_assert.Nil(err)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err := svcClient.GetProperties(ctx)
	_assert.Nil(err)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, *enabled)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, *days)

	// Empty retention policy causes an error, this is different from track 1.5
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{}})
	_assert.NotNil(err)
}

func (s *azblobTestSuite) TestAccountDeleteRetentionPolicyNil() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	days := to.Int32Ptr(5)
	enabled := to.BoolPtr(true)
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: enabled, Days: days}})
	_assert.Nil(err)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	resp, err := svcClient.GetProperties(ctx)
	_assert.Nil(err)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, *enabled)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, *days)

	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{})
	_assert.Nil(err)

	// From FE, 30 seconds is guaranteed to be enough.
	time.Sleep(time.Second * 30)

	// If an element of service properties is not passed, the service keeps the current settings.
	resp, err = svcClient.GetProperties(ctx)
	_assert.Nil(err)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Enabled, *enabled)
	_assert.EqualValues(*resp.StorageServiceProperties.DeleteRetentionPolicy.Days, *days)

	// Disable for other tests
	enabled = to.BoolPtr(false)
	svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: enabled}})
}

func (s *azblobTestSuite) TestAccountDeleteRetentionPolicyDaysTooSmall() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	days := int32(0) // Minimum days is 1. Validated on the client.
	enabled := true
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	_assert.NotNil(err)
}

func (s *azblobTestSuite) TestAccountDeleteRetentionPolicyDaysTooLarge() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	days := int32(366) // Max days is 365. Left to the service for validation.
	enabled := true
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled, Days: &days}})
	_assert.NotNil(err)

	validateStorageError(_assert, err, StorageErrorCodeInvalidXMLDocument)
}

func (s *azblobTestSuite) TestAccountDeleteRetentionPolicyDaysOmitted() {
	_assert := assert.New(s.T())
	testName := s.T().Name()
	_context := getTestContext(testName)
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	// Days is required if enabled is true.
	enabled := true
	_, err = svcClient.SetProperties(ctx, StorageServiceProperties{DeleteRetentionPolicy: &RetentionPolicy{Enabled: &enabled}})
	_assert.NotNil(err)

	validateStorageError(_assert, err, StorageErrorCodeInvalidXMLDocument)
}
