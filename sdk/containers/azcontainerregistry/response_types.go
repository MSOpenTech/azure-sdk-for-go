//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azcontainerregistry

import "io"

// authenticationClientExchangeAADAccessTokenForACRRefreshTokenResponse contains the response from method authenticationClient.ExchangeAADAccessTokenForACRRefreshToken.
type authenticationClientExchangeAADAccessTokenForACRRefreshTokenResponse struct {
	acrRefreshToken
}

// authenticationClientExchangeACRRefreshTokenForACRAccessTokenResponse contains the response from method authenticationClient.ExchangeACRRefreshTokenForACRAccessToken.
type authenticationClientExchangeACRRefreshTokenForACRAccessTokenResponse struct {
	acrAccessToken
}

// BlobClientCancelUploadResponse contains the response from method BlobClient.CancelUpload.
type BlobClientCancelUploadResponse struct {
	// placeholder for future response values
}

// BlobClientCheckBlobExistsResponse contains the response from method BlobClient.CheckBlobExists.
type BlobClientCheckBlobExistsResponse struct {
	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// BlobClientCheckChunkExistsResponse contains the response from method BlobClient.CheckChunkExists.
type BlobClientCheckChunkExistsResponse struct {
	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// ContentRange contains the information returned from the Content-Range header response.
	ContentRange *string
}

// BlobClientCompleteUploadResponse contains the response from method BlobClient.CompleteUpload.
type BlobClientCompleteUploadResponse struct {
	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string

	// Location contains the information returned from the Location header response.
	Location *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// BlobClientDeleteBlobResponse contains the response from method BlobClient.DeleteBlob.
type BlobClientDeleteBlobResponse struct {
	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// BlobClientGetBlobResponse contains the response from method BlobClient.GetBlob.
type BlobClientGetBlobResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser

	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// BlobClientGetChunkResponse contains the response from method BlobClient.GetChunk.
type BlobClientGetChunkResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser

	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// ContentRange contains the information returned from the Content-Range header response.
	ContentRange *string
}

// BlobClientGetUploadStatusResponse contains the response from method BlobClient.GetUploadStatus.
type BlobClientGetUploadStatusResponse struct {
	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// BlobClientMountBlobResponse contains the response from method BlobClient.MountBlob.
type BlobClientMountBlobResponse struct {
	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string

	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Location contains the information returned from the Location header response.
	Location *string
}

// BlobClientStartUploadResponse contains the response from method BlobClient.StartUpload.
type BlobClientStartUploadResponse struct {
	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Location contains the information returned from the Location header response.
	Location *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// BlobClientUploadChunkResponse contains the response from method BlobClient.UploadChunk.
type BlobClientUploadChunkResponse struct {
	// DockerUploadUUID contains the information returned from the Docker-Upload-UUID header response.
	DockerUploadUUID *string

	// Location contains the information returned from the Location header response.
	Location *string

	// Range contains the information returned from the Range header response.
	Range *string
}

// ClientDeleteManifestResponse contains the response from method Client.DeleteManifest.
type ClientDeleteManifestResponse struct {
	// placeholder for future response values
}

// ClientDeleteRepositoryResponse contains the response from method Client.DeleteRepository.
type ClientDeleteRepositoryResponse struct {
	// placeholder for future response values
}

// ClientDeleteTagResponse contains the response from method Client.DeleteTag.
type ClientDeleteTagResponse struct {
	// placeholder for future response values
}

// ClientGetManifestPropertiesResponse contains the response from method Client.GetManifestProperties.
type ClientGetManifestPropertiesResponse struct {
	ArtifactManifestProperties
}

// ClientGetManifestResponse contains the response from method Client.GetManifest.
type ClientGetManifestResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string
}

// ClientGetRepositoryPropertiesResponse contains the response from method Client.GetRepositoryProperties.
type ClientGetRepositoryPropertiesResponse struct {
	ContainerRepositoryProperties
}

// ClientGetTagPropertiesResponse contains the response from method Client.GetTagProperties.
type ClientGetTagPropertiesResponse struct {
	ArtifactTagProperties
}

// ClientListManifestsResponse contains the response from method Client.NewListManifestsPager.
type ClientListManifestsResponse struct {
	ACRManifests
	// Link contains the information returned from the Link header response.
	Link *string
}

// ClientListRepositoriesResponse contains the response from method Client.NewListRepositoriesPager.
type ClientListRepositoriesResponse struct {
	Repositories
	// Link contains the information returned from the Link header response.
	Link *string
}

// ClientListTagsResponse contains the response from method Client.NewListTagsPager.
type ClientListTagsResponse struct {
	TagList
	// Link contains the information returned from the Link header response.
	Link *string
}

// ClientUpdateManifestPropertiesResponse contains the response from method Client.UpdateManifestProperties.
type ClientUpdateManifestPropertiesResponse struct {
	ArtifactManifestProperties
}

// ClientUpdateRepositoryPropertiesResponse contains the response from method Client.UpdateRepositoryProperties.
type ClientUpdateRepositoryPropertiesResponse struct {
	ContainerRepositoryProperties
}

// ClientUpdateTagPropertiesResponse contains the response from method Client.UpdateTagProperties.
type ClientUpdateTagPropertiesResponse struct {
	ArtifactTagProperties
}

// ClientUploadManifestResponse contains the response from method Client.UploadManifest.
type ClientUploadManifestResponse struct {
	// ContentLength contains the information returned from the Content-Length header response.
	ContentLength *int64

	// DockerContentDigest contains the information returned from the Docker-Content-Digest header response.
	DockerContentDigest *string

	// Location contains the information returned from the Location header response.
	Location *string
}
