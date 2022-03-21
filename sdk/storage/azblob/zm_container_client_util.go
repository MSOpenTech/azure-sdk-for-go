// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

// ---------------------------------------------------------------------------------------------------------------------

// ContainerCreateOptions provides set of configurations for CreateContainer operation
type ContainerCreateOptions struct {
	// Specifies whether data in the container may be accessed publicly and the level of access
	Access *PublicAccessType

	// Optional. Specifies a user-defined name-value pair associated with the blob.
	Metadata map[string]string

	// Optional. Specifies the encryption scope settings to set on the container.
	CpkScope *ContainerCpkScopeInfo
}

func (o *ContainerCreateOptions) pointers() (*containerClientCreateOptions, *ContainerCpkScopeInfo) {
	if o == nil {
		return nil, nil
	}

	basicOptions := containerClientCreateOptions{
		Access:   o.Access,
		Metadata: o.Metadata,
	}

	return &basicOptions, o.CpkScope
}

// ContainerCreateResponse is wrapper around containerClientCreateResponse
type ContainerCreateResponse struct {
	containerClientCreateResponse
}

func toContainerCreateResponse(resp containerClientCreateResponse) ContainerCreateResponse {
	return ContainerCreateResponse{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

// ContainerDeleteOptions provides set of configurations for DeleteContainer operation
type ContainerDeleteOptions struct {
	LeaseAccessConditions    *LeaseAccessConditions
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *ContainerDeleteOptions) pointers() (*containerClientDeleteOptions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	return nil, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type ContainerDeleteResponse struct {
	containerClientDeleteResponse
}

func toContainerDeleteResponse(resp containerClientDeleteResponse) ContainerDeleteResponse {
	return ContainerDeleteResponse{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

// ContainerGetPropertiesOptions provides set of configurations for GetPropertiesContainer operation
type ContainerGetPropertiesOptions struct {
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *ContainerGetPropertiesOptions) pointers() (*containerClientGetPropertiesOptions, *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return nil, o.LeaseAccessConditions
}

type ContainerGetPropertiesResponse struct {
	containerClientGetPropertiesResponse
}

func toContainerGetPropertiesResponse(resp containerClientGetPropertiesResponse) ContainerGetPropertiesResponse {
	return ContainerGetPropertiesResponse{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

// ContainerSetMetadataOptions provides set of configurations for SetMetadataContainer operation
type ContainerSetMetadataOptions struct {
	Metadata                 map[string]string
	LeaseAccessConditions    *LeaseAccessConditions
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *ContainerSetMetadataOptions) pointers() (*containerClientSetMetadataOptions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}

	return &containerClientSetMetadataOptions{Metadata: o.Metadata}, o.LeaseAccessConditions, o.ModifiedAccessConditions
}

type ContainerSetMetadataResponse struct {
	containerClientSetMetadataResponse
}

func toContainerSetMetadataResponse(resp containerClientSetMetadataResponse) ContainerSetMetadataResponse {
	return ContainerSetMetadataResponse{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

// ContainerGetAccessPolicyOptions provides set of configurations for GetAccessPolicy operation
type ContainerGetAccessPolicyOptions struct {
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *ContainerGetAccessPolicyOptions) pointers() (*containerClientGetAccessPolicyOptions, *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return nil, o.LeaseAccessConditions
}

type ContainerGetAccessPolicyResponse struct {
	containerClientGetAccessPolicyResponse
}

func toContainerGetAccessPolicyResponse(resp containerClientGetAccessPolicyResponse) ContainerGetAccessPolicyResponse {
	return ContainerGetAccessPolicyResponse{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

// ContainerSetAccessPolicyOptions provides set of configurations for SetAccessPolicy operation
type ContainerSetAccessPolicyOptions struct {
	AccessConditions *ContainerAccessConditions
	// Specifies whether data in the container may be accessed publicly and the level of access
	Access *PublicAccessType
	// the acls for the container
	ContainerACL []*SignedIdentifier
}

func (o *ContainerSetAccessPolicyOptions) pointers() (*containerClientSetAccessPolicyOptions, *LeaseAccessConditions, *ModifiedAccessConditions) {
	if o == nil {
		return nil, nil, nil
	}
	mac, lac := o.AccessConditions.pointers()
	return &containerClientSetAccessPolicyOptions{
		Access:       o.Access,
		ContainerACL: o.ContainerACL,
	}, lac, mac
}

type ContainerSetAccessPolicyResponse struct {
	containerClientSetAccessPolicyResponse
}

func toContainerSetAccessPolicyResponse(resp containerClientSetAccessPolicyResponse) ContainerSetAccessPolicyResponse {
	return ContainerSetAccessPolicyResponse{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

type ContainerListBlobFlatSegmentOptions struct {
	// Include this parameter to specify one or more datasets to include in the response.
	Include []ListBlobsIncludeItem
	// A string value that identifies the portion of the list of containers to be returned with the next listing operation. The
	// operation returns the NextMarker value within the response body if the listing
	// operation did not return all containers remaining to be listed with the current page. The NextMarker value can be used
	// as the value for the marker parameter in a subsequent call to request the next
	// page of list items. The marker value is opaque to the client.
	Marker *string
	// Specifies the maximum number of containers to return. If the request does not specify maxresults, or specifies a value
	// greater than 5000, the server will return up to 5000 items. Note that if the
	// listing operation crosses a partition boundary, then the service will return a continuation token for retrieving the remainder
	// of the results. For this reason, it is possible that the service will
	// return fewer results than specified by maxresults, or than the default of 5000.
	Maxresults *int32
	// Filters the results to return only containers whose name begins with the specified prefix.
	Prefix *string
}

func (o *ContainerListBlobFlatSegmentOptions) pointers() *containerClientListBlobFlatSegmentOptions {
	return &containerClientListBlobFlatSegmentOptions{
		Include:    o.Include,
		Marker:     o.Marker,
		Maxresults: o.Maxresults,
		Prefix:     o.Prefix,
	}
}

type ContainerListBlobFlatSegmentPager struct {
	*containerClientListBlobFlatSegmentPager
}

func toContainerListBlobFlatSegmentPager(resp *containerClientListBlobFlatSegmentPager) *ContainerListBlobFlatSegmentPager {
	return &ContainerListBlobFlatSegmentPager{resp}
}

// ---------------------------------------------------------------------------------------------------------------------

type ContainerListBlobHierarchySegmentOptions struct {
	// Include this parameter to specify one or more datasets to include in the response.
	Include []ListBlobsIncludeItem
	// A string value that identifies the portion of the list of containers to be returned with the next listing operation. The
	// operation returns the NextMarker value within the response body if the listing
	// operation did not return all containers remaining to be listed with the current page. The NextMarker value can be used
	// as the value for the marker parameter in a subsequent call to request the next
	// page of list items. The marker value is opaque to the client.
	Marker *string
	// Specifies the maximum number of containers to return. If the request does not specify maxresults, or specifies a value
	// greater than 5000, the server will return up to 5000 items. Note that if the
	// listing operation crosses a partition boundary, then the service will return a continuation token for retrieving the remainder
	// of the results. For this reason, it is possible that the service will
	// return fewer results than specified by maxresults, or than the default of 5000.
	Maxresults *int32
	// Filters the results to return only containers whose name begins with the specified prefix.
	Prefix *string
}

func (o *ContainerListBlobHierarchySegmentOptions) pointers() *containerClientListBlobHierarchySegmentOptions {
	return &containerClientListBlobHierarchySegmentOptions{
		Include:    o.Include,
		Marker:     o.Marker,
		Maxresults: o.Maxresults,
		Prefix:     o.Prefix,
	}
}

type ContainerListBlobHierarchySegmentPager struct {
	containerClientListBlobHierarchySegmentPager
}

func toContainerListBlobHierarchySegmentPager(resp *containerClientListBlobHierarchySegmentPager) *ContainerListBlobHierarchySegmentPager {
	if resp == nil {
		return nil
	}
	return &ContainerListBlobHierarchySegmentPager{*resp}
}
