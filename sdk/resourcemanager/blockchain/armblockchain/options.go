//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// Name availability request payload.
	NameAvailabilityRequest *NameAvailabilityRequest
}

// LocationsClientListConsortiumsOptions contains the optional parameters for the LocationsClient.ListConsortiums method.
type LocationsClientListConsortiumsOptions struct {
	// placeholder for future optional parameters
}

// MemberOperationResultsClientGetOptions contains the optional parameters for the MemberOperationResultsClient.Get method.
type MemberOperationResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MembersClientBeginCreateOptions contains the optional parameters for the MembersClient.BeginCreate method.
type MembersClientBeginCreateOptions struct {
	// Payload to create a blockchain member.
	BlockchainMember *Member

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MembersClientBeginDeleteOptions contains the optional parameters for the MembersClient.BeginDelete method.
type MembersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MembersClientGetOptions contains the optional parameters for the MembersClient.Get method.
type MembersClientGetOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListAPIKeysOptions contains the optional parameters for the MembersClient.ListAPIKeys method.
type MembersClientListAPIKeysOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListAllOptions contains the optional parameters for the MembersClient.NewListAllPager method.
type MembersClientListAllOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListConsortiumMembersOptions contains the optional parameters for the MembersClient.NewListConsortiumMembersPager
// method.
type MembersClientListConsortiumMembersOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListOptions contains the optional parameters for the MembersClient.NewListPager method.
type MembersClientListOptions struct {
	// placeholder for future optional parameters
}

// MembersClientListRegenerateAPIKeysOptions contains the optional parameters for the MembersClient.ListRegenerateAPIKeys
// method.
type MembersClientListRegenerateAPIKeysOptions struct {
	// api key to be regenerate
	APIKey *APIKey
}

// MembersClientUpdateOptions contains the optional parameters for the MembersClient.Update method.
type MembersClientUpdateOptions struct {
	// Payload to update the blockchain member.
	BlockchainMember *MemberUpdate
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SKUsClientListOptions contains the optional parameters for the SKUsClient.List method.
type SKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientBeginCreateOptions contains the optional parameters for the TransactionNodesClient.BeginCreate method.
type TransactionNodesClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// Payload to create the transaction node.
	TransactionNode *TransactionNode
}

// TransactionNodesClientBeginDeleteOptions contains the optional parameters for the TransactionNodesClient.BeginDelete method.
type TransactionNodesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TransactionNodesClientGetOptions contains the optional parameters for the TransactionNodesClient.Get method.
type TransactionNodesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientListAPIKeysOptions contains the optional parameters for the TransactionNodesClient.ListAPIKeys method.
type TransactionNodesClientListAPIKeysOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientListOptions contains the optional parameters for the TransactionNodesClient.NewListPager method.
type TransactionNodesClientListOptions struct {
	// placeholder for future optional parameters
}

// TransactionNodesClientListRegenerateAPIKeysOptions contains the optional parameters for the TransactionNodesClient.ListRegenerateAPIKeys
// method.
type TransactionNodesClientListRegenerateAPIKeysOptions struct {
	// api key to be regenerated
	APIKey *APIKey
}

// TransactionNodesClientUpdateOptions contains the optional parameters for the TransactionNodesClient.Update method.
type TransactionNodesClientUpdateOptions struct {
	// Payload to create the transaction node.
	TransactionNode *TransactionNodeUpdate
}
