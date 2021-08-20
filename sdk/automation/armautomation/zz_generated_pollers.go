//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// DscCompilationJobCreatePoller provides polling facilities until the operation reaches a terminal state.
type DscCompilationJobCreatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DscCompilationJobCreateResponse will be returned.
	FinalResponse(ctx context.Context) (DscCompilationJobCreateResponse, error)
}

type dscCompilationJobCreatePoller struct {
	pt *armcore.LROPoller
}

func (p *dscCompilationJobCreatePoller) Done() bool {
	return p.pt.Done()
}

func (p *dscCompilationJobCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *dscCompilationJobCreatePoller) FinalResponse(ctx context.Context) (DscCompilationJobCreateResponse, error) {
	respType := DscCompilationJobCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DscCompilationJob)
	if err != nil {
		return DscCompilationJobCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *dscCompilationJobCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *dscCompilationJobCreatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (DscCompilationJobCreateResponse, error) {
	respType := DscCompilationJobCreateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.DscCompilationJob)
	if err != nil {
		return DscCompilationJobCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// DscNodeConfigurationCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DscNodeConfigurationCreateOrUpdatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DscNodeConfigurationCreateOrUpdateResponse will be returned.
	FinalResponse(ctx context.Context) (DscNodeConfigurationCreateOrUpdateResponse, error)
}

type dscNodeConfigurationCreateOrUpdatePoller struct {
	pt *armcore.LROPoller
}

func (p *dscNodeConfigurationCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

func (p *dscNodeConfigurationCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *dscNodeConfigurationCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DscNodeConfigurationCreateOrUpdateResponse, error) {
	respType := DscNodeConfigurationCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DscNodeConfiguration)
	if err != nil {
		return DscNodeConfigurationCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *dscNodeConfigurationCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *dscNodeConfigurationCreateOrUpdatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (DscNodeConfigurationCreateOrUpdateResponse, error) {
	respType := DscNodeConfigurationCreateOrUpdateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.DscNodeConfiguration)
	if err != nil {
		return DscNodeConfigurationCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// PrivateEndpointConnectionsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsCreateOrUpdatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final PrivateEndpointConnectionsCreateOrUpdateResponse will be returned.
	FinalResponse(ctx context.Context) (PrivateEndpointConnectionsCreateOrUpdateResponse, error)
}

type privateEndpointConnectionsCreateOrUpdatePoller struct {
	pt *armcore.LROPoller
}

func (p *privateEndpointConnectionsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

func (p *privateEndpointConnectionsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *privateEndpointConnectionsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *privateEndpointConnectionsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *privateEndpointConnectionsCreateOrUpdatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsCreateOrUpdateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// PrivateEndpointConnectionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsDeletePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final PrivateEndpointConnectionsDeleteResponse will be returned.
	FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error)
}

type privateEndpointConnectionsDeletePoller struct {
	pt *armcore.LROPoller
}

func (p *privateEndpointConnectionsDeletePoller) Done() bool {
	return p.pt.Done()
}

func (p *privateEndpointConnectionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *privateEndpointConnectionsDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *privateEndpointConnectionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *privateEndpointConnectionsDeletePoller) pollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RunbookDraftReplaceContentPoller provides polling facilities until the operation reaches a terminal state.
type RunbookDraftReplaceContentPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RunbookDraftReplaceContentResponse will be returned.
	FinalResponse(ctx context.Context) (RunbookDraftReplaceContentResponse, error)
}

type runbookDraftReplaceContentPoller struct {
	pt *armcore.LROPoller
}

func (p *runbookDraftReplaceContentPoller) Done() bool {
	return p.pt.Done()
}

func (p *runbookDraftReplaceContentPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *runbookDraftReplaceContentPoller) FinalResponse(ctx context.Context) (RunbookDraftReplaceContentResponse, error) {
	respType := RunbookDraftReplaceContentResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RunbookDraftReplaceContentResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *runbookDraftReplaceContentPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *runbookDraftReplaceContentPoller) pollUntilDone(ctx context.Context, freq time.Duration) (RunbookDraftReplaceContentResponse, error) {
	respType := RunbookDraftReplaceContentResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return RunbookDraftReplaceContentResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RunbookPublishPoller provides polling facilities until the operation reaches a terminal state.
type RunbookPublishPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RunbookPublishResponse will be returned.
	FinalResponse(ctx context.Context) (RunbookPublishResponse, error)
}

type runbookPublishPoller struct {
	pt *armcore.LROPoller
}

func (p *runbookPublishPoller) Done() bool {
	return p.pt.Done()
}

func (p *runbookPublishPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *runbookPublishPoller) FinalResponse(ctx context.Context) (RunbookPublishResponse, error) {
	respType := RunbookPublishResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RunbookPublishResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *runbookPublishPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *runbookPublishPoller) pollUntilDone(ctx context.Context, freq time.Duration) (RunbookPublishResponse, error) {
	respType := RunbookPublishResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return RunbookPublishResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
