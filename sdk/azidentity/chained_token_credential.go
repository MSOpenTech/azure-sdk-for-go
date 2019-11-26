// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ChainedTokenCredential provides a TokenCredential implementation which chains multiple TokenCredential implementations to be tried in order
// until one of the GetToken methods returns a non-default AccessToken.
type ChainedTokenCredential struct {
	sources []azcore.TokenCredential
}

// NewChainedTokenCredential creates an instance of ChainedTokenCredential with the specified TokenCredential sources.
func NewChainedTokenCredential(sources ...azcore.TokenCredential) (*ChainedTokenCredential, error) {
	if len(sources) == 0 {
		return nil, &CredentialUnavailableError{CredentialType: "Chained Token Credential", Message: "Length of sources cannot be 0"}
	}

	for _, source := range sources {
		if source == nil {
			return nil, &CredentialUnavailableError{CredentialType: "Chained Token Credential", Message: "Sources cannot contain a nil TokenCredential"}
		}
	}

	return &ChainedTokenCredential{sources: sources}, nil
}

// GetToken sequentially calls TokenCredential.GetToken on all the specified sources, returning the first non default AccessToken.
func (c *ChainedTokenCredential) GetToken(ctx context.Context, opts *azcore.TokenRequestOptions) (*azcore.AccessToken, error) {
	var credErr *CredentialUnavailableError
	var token *azcore.AccessToken
	var err error
	var errList []*CredentialUnavailableError

	for i := 0; i < len(c.sources); i++ {
		token, err = c.sources[i].GetToken(ctx, opts)
		if errors.As(err, &credErr) {
			errList = append(errList, err.(*CredentialUnavailableError))
		} else {
			if token != nil && err == nil {
				return token, nil
			}
			return token, &ChainedCredentialError{Message: err.Error(), ErrorList: errList}
		}
	}

	if token == nil && len(errList) > 0 {
		err = &ChainedCredentialError{ErrorList: errList}
		return nil, err
	}

	return token, nil
}

// AuthenticationPolicy implements the azcore.Credential interface on ChainedTokenCredential.
func (c *ChainedTokenCredential) AuthenticationPolicy(options azcore.AuthenticationPolicyOptions) azcore.Policy {
	return newBearerTokenPolicy(c, options.Scopes)
}
