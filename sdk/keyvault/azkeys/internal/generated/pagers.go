//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// KeyVaultClientGetDeletedKeysPager provides operations for iterating over paged responses.
type KeyVaultClientGetDeletedKeysPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedKeysResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedKeysResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *KeyVaultClientGetDeletedKeysPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedKeyListResult.NextLink == nil || len(*p.current.DeletedKeyListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *KeyVaultClientGetDeletedKeysPager) NextPage(ctx context.Context) (KeyVaultClientGetDeletedKeysResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return KeyVaultClientGetDeletedKeysResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return KeyVaultClientGetDeletedKeysResponse{}, err
	}
	resp, err := p.client.Pl.Do(req)
	if err != nil {
		return KeyVaultClientGetDeletedKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return KeyVaultClientGetDeletedKeysResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDeletedKeysHandleResponse(resp)
	if err != nil {
		return KeyVaultClientGetDeletedKeysResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// KeyVaultClientGetKeyVersionsPager provides operations for iterating over paged responses.
type KeyVaultClientGetKeyVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetKeyVersionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetKeyVersionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *KeyVaultClientGetKeyVersionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *KeyVaultClientGetKeyVersionsPager) NextPage(ctx context.Context) (KeyVaultClientGetKeyVersionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return KeyVaultClientGetKeyVersionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return KeyVaultClientGetKeyVersionsResponse{}, err
	}
	resp, err := p.client.Pl.Do(req)
	if err != nil {
		return KeyVaultClientGetKeyVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return KeyVaultClientGetKeyVersionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getKeyVersionsHandleResponse(resp)
	if err != nil {
		return KeyVaultClientGetKeyVersionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// KeyVaultClientGetKeysPager provides operations for iterating over paged responses.
type KeyVaultClientGetKeysPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetKeysResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetKeysResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *KeyVaultClientGetKeysPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *KeyVaultClientGetKeysPager) NextPage(ctx context.Context) (KeyVaultClientGetKeysResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return KeyVaultClientGetKeysResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return KeyVaultClientGetKeysResponse{}, err
	}
	resp, err := p.client.Pl.Do(req)
	if err != nil {
		return KeyVaultClientGetKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return KeyVaultClientGetKeysResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.GetKeysHandleResponse(resp)
	if err != nil {
		return KeyVaultClientGetKeysResponse{}, err
	}
	p.current = result
	return p.current, nil
}
