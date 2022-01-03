//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// MoveCollectionsBulkRemovePoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsBulkRemovePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsBulkRemovePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsBulkRemovePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsBulkRemoveResponse will be returned.
func (p *MoveCollectionsBulkRemovePoller) FinalResponse(ctx context.Context) (MoveCollectionsBulkRemoveResponse, error) {
	respType := MoveCollectionsBulkRemoveResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsBulkRemoveResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsBulkRemovePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveCollectionsCommitPoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsCommitPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsCommitPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsCommitPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsCommitResponse will be returned.
func (p *MoveCollectionsCommitPoller) FinalResponse(ctx context.Context) (MoveCollectionsCommitResponse, error) {
	respType := MoveCollectionsCommitResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsCommitResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsCommitPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveCollectionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsDeleteResponse will be returned.
func (p *MoveCollectionsDeletePoller) FinalResponse(ctx context.Context) (MoveCollectionsDeleteResponse, error) {
	respType := MoveCollectionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveCollectionsDiscardPoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsDiscardPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsDiscardPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsDiscardPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsDiscardResponse will be returned.
func (p *MoveCollectionsDiscardPoller) FinalResponse(ctx context.Context) (MoveCollectionsDiscardResponse, error) {
	respType := MoveCollectionsDiscardResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsDiscardResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsDiscardPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveCollectionsInitiateMovePoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsInitiateMovePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsInitiateMovePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsInitiateMovePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsInitiateMoveResponse will be returned.
func (p *MoveCollectionsInitiateMovePoller) FinalResponse(ctx context.Context) (MoveCollectionsInitiateMoveResponse, error) {
	respType := MoveCollectionsInitiateMoveResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsInitiateMoveResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsInitiateMovePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveCollectionsPreparePoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsPreparePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsPreparePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsPreparePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsPrepareResponse will be returned.
func (p *MoveCollectionsPreparePoller) FinalResponse(ctx context.Context) (MoveCollectionsPrepareResponse, error) {
	respType := MoveCollectionsPrepareResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsPrepareResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsPreparePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveCollectionsResolveDependenciesPoller provides polling facilities until the operation reaches a terminal state.
type MoveCollectionsResolveDependenciesPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveCollectionsResolveDependenciesPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveCollectionsResolveDependenciesPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveCollectionsResolveDependenciesResponse will be returned.
func (p *MoveCollectionsResolveDependenciesPoller) FinalResponse(ctx context.Context) (MoveCollectionsResolveDependenciesResponse, error) {
	respType := MoveCollectionsResolveDependenciesResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveCollectionsResolveDependenciesResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveCollectionsResolveDependenciesPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveResourcesCreatePoller provides polling facilities until the operation reaches a terminal state.
type MoveResourcesCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveResourcesCreatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveResourcesCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveResourcesCreateResponse will be returned.
func (p *MoveResourcesCreatePoller) FinalResponse(ctx context.Context) (MoveResourcesCreateResponse, error) {
	respType := MoveResourcesCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.MoveResource)
	if err != nil {
		return MoveResourcesCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveResourcesCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MoveResourcesDeletePoller provides polling facilities until the operation reaches a terminal state.
type MoveResourcesDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MoveResourcesDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MoveResourcesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MoveResourcesDeleteResponse will be returned.
func (p *MoveResourcesDeletePoller) FinalResponse(ctx context.Context) (MoveResourcesDeleteResponse, error) {
	respType := MoveResourcesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OperationStatus)
	if err != nil {
		return MoveResourcesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MoveResourcesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
