// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// AppServiceCertificateOrderPoller provides polling facilities until the operation reaches a terminal state.
type AppServiceCertificateOrderPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final AppServiceCertificateOrderResponse will be returned.
	FinalResponse(ctx context.Context) (AppServiceCertificateOrderResponse, error)
}

type appServiceCertificateOrderPoller struct {
	pt *armcore.LROPoller
}

func (p *appServiceCertificateOrderPoller) Done() bool {
	return p.pt.Done()
}

func (p *appServiceCertificateOrderPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *appServiceCertificateOrderPoller) FinalResponse(ctx context.Context) (AppServiceCertificateOrderResponse, error) {
	respType := AppServiceCertificateOrderResponse{AppServiceCertificateOrder: &AppServiceCertificateOrder{}}
	resp, err := p.pt.FinalResponse(ctx, respType.AppServiceCertificateOrder)
	if err != nil {
		return AppServiceCertificateOrderResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *appServiceCertificateOrderPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *appServiceCertificateOrderPoller) pollUntilDone(ctx context.Context, freq time.Duration) (AppServiceCertificateOrderResponse, error) {
	respType := AppServiceCertificateOrderResponse{AppServiceCertificateOrder: &AppServiceCertificateOrder{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.AppServiceCertificateOrder)
	if err != nil {
		return AppServiceCertificateOrderResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// AppServiceCertificateResourcePoller provides polling facilities until the operation reaches a terminal state.
type AppServiceCertificateResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final AppServiceCertificateResourceResponse will be returned.
	FinalResponse(ctx context.Context) (AppServiceCertificateResourceResponse, error)
}

type appServiceCertificateResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *appServiceCertificateResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *appServiceCertificateResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *appServiceCertificateResourcePoller) FinalResponse(ctx context.Context) (AppServiceCertificateResourceResponse, error) {
	respType := AppServiceCertificateResourceResponse{AppServiceCertificateResource: &AppServiceCertificateResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.AppServiceCertificateResource)
	if err != nil {
		return AppServiceCertificateResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *appServiceCertificateResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *appServiceCertificateResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (AppServiceCertificateResourceResponse, error) {
	respType := AppServiceCertificateResourceResponse{AppServiceCertificateResource: &AppServiceCertificateResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.AppServiceCertificateResource)
	if err != nil {
		return AppServiceCertificateResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// AppServiceEnvironmentResourcePoller provides polling facilities until the operation reaches a terminal state.
type AppServiceEnvironmentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final AppServiceEnvironmentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (AppServiceEnvironmentResourceResponse, error)
}

type appServiceEnvironmentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *appServiceEnvironmentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *appServiceEnvironmentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *appServiceEnvironmentResourcePoller) FinalResponse(ctx context.Context) (AppServiceEnvironmentResourceResponse, error) {
	respType := AppServiceEnvironmentResourceResponse{AppServiceEnvironmentResource: &AppServiceEnvironmentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.AppServiceEnvironmentResource)
	if err != nil {
		return AppServiceEnvironmentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *appServiceEnvironmentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *appServiceEnvironmentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (AppServiceEnvironmentResourceResponse, error) {
	respType := AppServiceEnvironmentResourceResponse{AppServiceEnvironmentResource: &AppServiceEnvironmentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.AppServiceEnvironmentResource)
	if err != nil {
		return AppServiceEnvironmentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// AppServicePlanPoller provides polling facilities until the operation reaches a terminal state.
type AppServicePlanPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final AppServicePlanResponse will be returned.
	FinalResponse(ctx context.Context) (AppServicePlanResponse, error)
}

type appServicePlanPoller struct {
	pt *armcore.LROPoller
}

func (p *appServicePlanPoller) Done() bool {
	return p.pt.Done()
}

func (p *appServicePlanPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *appServicePlanPoller) FinalResponse(ctx context.Context) (AppServicePlanResponse, error) {
	respType := AppServicePlanResponse{AppServicePlan: &AppServicePlan{}}
	resp, err := p.pt.FinalResponse(ctx, respType.AppServicePlan)
	if err != nil {
		return AppServicePlanResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *appServicePlanPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *appServicePlanPoller) pollUntilDone(ctx context.Context, freq time.Duration) (AppServicePlanResponse, error) {
	respType := AppServicePlanResponse{AppServicePlan: &AppServicePlan{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.AppServicePlan)
	if err != nil {
		return AppServicePlanResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// DomainPoller provides polling facilities until the operation reaches a terminal state.
type DomainPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DomainResponse will be returned.
	FinalResponse(ctx context.Context) (DomainResponse, error)
}

type domainPoller struct {
	pt *armcore.LROPoller
}

func (p *domainPoller) Done() bool {
	return p.pt.Done()
}

func (p *domainPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *domainPoller) FinalResponse(ctx context.Context) (DomainResponse, error) {
	respType := DomainResponse{Domain: &Domain{}}
	resp, err := p.pt.FinalResponse(ctx, respType.Domain)
	if err != nil {
		return DomainResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *domainPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *domainPoller) pollUntilDone(ctx context.Context, freq time.Duration) (DomainResponse, error) {
	respType := DomainResponse{Domain: &Domain{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.Domain)
	if err != nil {
		return DomainResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// FunctionEnvelopePoller provides polling facilities until the operation reaches a terminal state.
type FunctionEnvelopePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final FunctionEnvelopeResponse will be returned.
	FinalResponse(ctx context.Context) (FunctionEnvelopeResponse, error)
}

type functionEnvelopePoller struct {
	pt *armcore.LROPoller
}

func (p *functionEnvelopePoller) Done() bool {
	return p.pt.Done()
}

func (p *functionEnvelopePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *functionEnvelopePoller) FinalResponse(ctx context.Context) (FunctionEnvelopeResponse, error) {
	respType := FunctionEnvelopeResponse{FunctionEnvelope: &FunctionEnvelope{}}
	resp, err := p.pt.FinalResponse(ctx, respType.FunctionEnvelope)
	if err != nil {
		return FunctionEnvelopeResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *functionEnvelopePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *functionEnvelopePoller) pollUntilDone(ctx context.Context, freq time.Duration) (FunctionEnvelopeResponse, error) {
	respType := FunctionEnvelopeResponse{FunctionEnvelope: &FunctionEnvelope{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.FunctionEnvelope)
	if err != nil {
		return FunctionEnvelopeResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// HTTPPoller provides polling facilities until the operation reaches a terminal state.
type HTTPPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final *http.Response will be returned.
	FinalResponse(ctx context.Context) (*http.Response, error)
}

type httpPoller struct {
	pt *armcore.LROPoller
}

func (p *httpPoller) Done() bool {
	return p.pt.Done()
}

func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *httpPoller) FinalResponse(ctx context.Context) (*http.Response, error) {
	return p.pt.FinalResponse(ctx, nil)
}

func (p *httpPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpPoller) pollUntilDone(ctx context.Context, freq time.Duration) (*http.Response, error) {
	return p.pt.PollUntilDone(ctx, freq, nil)
}

// KubeEnvironmentPoller provides polling facilities until the operation reaches a terminal state.
type KubeEnvironmentPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final KubeEnvironmentResponse will be returned.
	FinalResponse(ctx context.Context) (KubeEnvironmentResponse, error)
}

type kubeEnvironmentPoller struct {
	pt *armcore.LROPoller
}

func (p *kubeEnvironmentPoller) Done() bool {
	return p.pt.Done()
}

func (p *kubeEnvironmentPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *kubeEnvironmentPoller) FinalResponse(ctx context.Context) (KubeEnvironmentResponse, error) {
	respType := KubeEnvironmentResponse{KubeEnvironment: &KubeEnvironment{}}
	resp, err := p.pt.FinalResponse(ctx, respType.KubeEnvironment)
	if err != nil {
		return KubeEnvironmentResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *kubeEnvironmentPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *kubeEnvironmentPoller) pollUntilDone(ctx context.Context, freq time.Duration) (KubeEnvironmentResponse, error) {
	respType := KubeEnvironmentResponse{KubeEnvironment: &KubeEnvironment{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.KubeEnvironment)
	if err != nil {
		return KubeEnvironmentResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// MSDeployStatusPoller provides polling facilities until the operation reaches a terminal state.
type MSDeployStatusPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final MSDeployStatusResponse will be returned.
	FinalResponse(ctx context.Context) (MSDeployStatusResponse, error)
}

type msDeployStatusPoller struct {
	pt *armcore.LROPoller
}

func (p *msDeployStatusPoller) Done() bool {
	return p.pt.Done()
}

func (p *msDeployStatusPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *msDeployStatusPoller) FinalResponse(ctx context.Context) (MSDeployStatusResponse, error) {
	respType := MSDeployStatusResponse{MSDeployStatus: &MSDeployStatus{}}
	resp, err := p.pt.FinalResponse(ctx, respType.MSDeployStatus)
	if err != nil {
		return MSDeployStatusResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *msDeployStatusPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *msDeployStatusPoller) pollUntilDone(ctx context.Context, freq time.Duration) (MSDeployStatusResponse, error) {
	respType := MSDeployStatusResponse{MSDeployStatus: &MSDeployStatus{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.MSDeployStatus)
	if err != nil {
		return MSDeployStatusResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// NetworkTraceArrayPoller provides polling facilities until the operation reaches a terminal state.
type NetworkTraceArrayPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final NetworkTraceArrayResponse will be returned.
	FinalResponse(ctx context.Context) (NetworkTraceArrayResponse, error)
}

type networkTraceArrayPoller struct {
	pt *armcore.LROPoller
}

func (p *networkTraceArrayPoller) Done() bool {
	return p.pt.Done()
}

func (p *networkTraceArrayPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *networkTraceArrayPoller) FinalResponse(ctx context.Context) (NetworkTraceArrayResponse, error) {
	respType := NetworkTraceArrayResponse{NetworkTraceArray: []*NetworkTrace{}}
	resp, err := p.pt.FinalResponse(ctx, &respType.NetworkTraceArray)
	if err != nil {
		return NetworkTraceArrayResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *networkTraceArrayPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *networkTraceArrayPoller) pollUntilDone(ctx context.Context, freq time.Duration) (NetworkTraceArrayResponse, error) {
	respType := NetworkTraceArrayResponse{NetworkTraceArray: []*NetworkTrace{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.NetworkTraceArray)
	if err != nil {
		return NetworkTraceArrayResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ObjectPoller provides polling facilities until the operation reaches a terminal state.
type ObjectPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final ObjectResponse will be returned.
	FinalResponse(ctx context.Context) (ObjectResponse, error)
}

type objectPoller struct {
	pt *armcore.LROPoller
}

func (p *objectPoller) Done() bool {
	return p.pt.Done()
}

func (p *objectPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *objectPoller) FinalResponse(ctx context.Context) (ObjectResponse, error) {
	respType := ObjectResponse{}
	resp, err := p.pt.FinalResponse(ctx, respType.Object)
	if err != nil {
		return ObjectResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *objectPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *objectPoller) pollUntilDone(ctx context.Context, freq time.Duration) (ObjectResponse, error) {
	respType := ObjectResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.Object)
	if err != nil {
		return ObjectResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// OperationPoller provides polling facilities until the operation reaches a terminal state.
type OperationPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final OperationResponse will be returned.
	FinalResponse(ctx context.Context) (OperationResponse, error)
}

type operationPoller struct {
	pt *armcore.LROPoller
}

func (p *operationPoller) Done() bool {
	return p.pt.Done()
}

func (p *operationPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *operationPoller) FinalResponse(ctx context.Context) (OperationResponse, error) {
	respType := OperationResponse{Operation: &Operation{}}
	resp, err := p.pt.FinalResponse(ctx, respType.Operation)
	if err != nil {
		return OperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *operationPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *operationPoller) pollUntilDone(ctx context.Context, freq time.Duration) (OperationResponse, error) {
	respType := OperationResponse{Operation: &Operation{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.Operation)
	if err != nil {
		return OperationResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RemotePrivateEndpointConnectionARMResourcePoller provides polling facilities until the operation reaches a terminal state.
type RemotePrivateEndpointConnectionARMResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RemotePrivateEndpointConnectionARMResourceResponse will be returned.
	FinalResponse(ctx context.Context) (RemotePrivateEndpointConnectionARMResourceResponse, error)
}

type remotePrivateEndpointConnectionARMResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *remotePrivateEndpointConnectionARMResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *remotePrivateEndpointConnectionARMResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *remotePrivateEndpointConnectionARMResourcePoller) FinalResponse(ctx context.Context) (RemotePrivateEndpointConnectionARMResourceResponse, error) {
	respType := RemotePrivateEndpointConnectionARMResourceResponse{RemotePrivateEndpointConnectionARMResource: &RemotePrivateEndpointConnectionARMResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.RemotePrivateEndpointConnectionARMResource)
	if err != nil {
		return RemotePrivateEndpointConnectionARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *remotePrivateEndpointConnectionARMResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *remotePrivateEndpointConnectionARMResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (RemotePrivateEndpointConnectionARMResourceResponse, error) {
	respType := RemotePrivateEndpointConnectionARMResourceResponse{RemotePrivateEndpointConnectionARMResource: &RemotePrivateEndpointConnectionARMResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.RemotePrivateEndpointConnectionARMResource)
	if err != nil {
		return RemotePrivateEndpointConnectionARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// SiteExtensionInfoPoller provides polling facilities until the operation reaches a terminal state.
type SiteExtensionInfoPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final SiteExtensionInfoResponse will be returned.
	FinalResponse(ctx context.Context) (SiteExtensionInfoResponse, error)
}

type siteExtensionInfoPoller struct {
	pt *armcore.LROPoller
}

func (p *siteExtensionInfoPoller) Done() bool {
	return p.pt.Done()
}

func (p *siteExtensionInfoPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *siteExtensionInfoPoller) FinalResponse(ctx context.Context) (SiteExtensionInfoResponse, error) {
	respType := SiteExtensionInfoResponse{SiteExtensionInfo: &SiteExtensionInfo{}}
	resp, err := p.pt.FinalResponse(ctx, respType.SiteExtensionInfo)
	if err != nil {
		return SiteExtensionInfoResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *siteExtensionInfoPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *siteExtensionInfoPoller) pollUntilDone(ctx context.Context, freq time.Duration) (SiteExtensionInfoResponse, error) {
	respType := SiteExtensionInfoResponse{SiteExtensionInfo: &SiteExtensionInfo{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.SiteExtensionInfo)
	if err != nil {
		return SiteExtensionInfoResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// SitePoller provides polling facilities until the operation reaches a terminal state.
type SitePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final SiteResponse will be returned.
	FinalResponse(ctx context.Context) (SiteResponse, error)
}

type sitePoller struct {
	pt *armcore.LROPoller
}

func (p *sitePoller) Done() bool {
	return p.pt.Done()
}

func (p *sitePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *sitePoller) FinalResponse(ctx context.Context) (SiteResponse, error) {
	respType := SiteResponse{Site: &Site{}}
	resp, err := p.pt.FinalResponse(ctx, respType.Site)
	if err != nil {
		return SiteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *sitePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *sitePoller) pollUntilDone(ctx context.Context, freq time.Duration) (SiteResponse, error) {
	respType := SiteResponse{Site: &Site{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.Site)
	if err != nil {
		return SiteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// SiteSourceControlPoller provides polling facilities until the operation reaches a terminal state.
type SiteSourceControlPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final SiteSourceControlResponse will be returned.
	FinalResponse(ctx context.Context) (SiteSourceControlResponse, error)
}

type siteSourceControlPoller struct {
	pt *armcore.LROPoller
}

func (p *siteSourceControlPoller) Done() bool {
	return p.pt.Done()
}

func (p *siteSourceControlPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *siteSourceControlPoller) FinalResponse(ctx context.Context) (SiteSourceControlResponse, error) {
	respType := SiteSourceControlResponse{SiteSourceControl: &SiteSourceControl{}}
	resp, err := p.pt.FinalResponse(ctx, respType.SiteSourceControl)
	if err != nil {
		return SiteSourceControlResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *siteSourceControlPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *siteSourceControlPoller) pollUntilDone(ctx context.Context, freq time.Duration) (SiteSourceControlResponse, error) {
	respType := SiteSourceControlResponse{SiteSourceControl: &SiteSourceControl{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.SiteSourceControl)
	if err != nil {
		return SiteSourceControlResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// StaticSiteARMResourcePoller provides polling facilities until the operation reaches a terminal state.
type StaticSiteARMResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final StaticSiteARMResourceResponse will be returned.
	FinalResponse(ctx context.Context) (StaticSiteARMResourceResponse, error)
}

type staticSiteARMResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *staticSiteARMResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *staticSiteARMResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *staticSiteARMResourcePoller) FinalResponse(ctx context.Context) (StaticSiteARMResourceResponse, error) {
	respType := StaticSiteARMResourceResponse{StaticSiteARMResource: &StaticSiteARMResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.StaticSiteARMResource)
	if err != nil {
		return StaticSiteARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *staticSiteARMResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *staticSiteARMResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (StaticSiteARMResourceResponse, error) {
	respType := StaticSiteARMResourceResponse{StaticSiteARMResource: &StaticSiteARMResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.StaticSiteARMResource)
	if err != nil {
		return StaticSiteARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// StaticSiteCustomDomainOverviewARMResourcePoller provides polling facilities until the operation reaches a terminal state.
type StaticSiteCustomDomainOverviewARMResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final StaticSiteCustomDomainOverviewARMResourceResponse will be returned.
	FinalResponse(ctx context.Context) (StaticSiteCustomDomainOverviewARMResourceResponse, error)
}

type staticSiteCustomDomainOverviewARMResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *staticSiteCustomDomainOverviewARMResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *staticSiteCustomDomainOverviewARMResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *staticSiteCustomDomainOverviewARMResourcePoller) FinalResponse(ctx context.Context) (StaticSiteCustomDomainOverviewARMResourceResponse, error) {
	respType := StaticSiteCustomDomainOverviewARMResourceResponse{StaticSiteCustomDomainOverviewARMResource: &StaticSiteCustomDomainOverviewARMResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.StaticSiteCustomDomainOverviewARMResource)
	if err != nil {
		return StaticSiteCustomDomainOverviewARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *staticSiteCustomDomainOverviewARMResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *staticSiteCustomDomainOverviewARMResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (StaticSiteCustomDomainOverviewARMResourceResponse, error) {
	respType := StaticSiteCustomDomainOverviewARMResourceResponse{StaticSiteCustomDomainOverviewARMResource: &StaticSiteCustomDomainOverviewARMResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.StaticSiteCustomDomainOverviewARMResource)
	if err != nil {
		return StaticSiteCustomDomainOverviewARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// StaticSiteUserProvidedFunctionAppARMResourcePoller provides polling facilities until the operation reaches a terminal state.
type StaticSiteUserProvidedFunctionAppARMResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final StaticSiteUserProvidedFunctionAppARMResourceResponse will be returned.
	FinalResponse(ctx context.Context) (StaticSiteUserProvidedFunctionAppARMResourceResponse, error)
}

type staticSiteUserProvidedFunctionAppARMResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *staticSiteUserProvidedFunctionAppARMResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *staticSiteUserProvidedFunctionAppARMResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *staticSiteUserProvidedFunctionAppARMResourcePoller) FinalResponse(ctx context.Context) (StaticSiteUserProvidedFunctionAppARMResourceResponse, error) {
	respType := StaticSiteUserProvidedFunctionAppARMResourceResponse{StaticSiteUserProvidedFunctionAppARMResource: &StaticSiteUserProvidedFunctionAppARMResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.StaticSiteUserProvidedFunctionAppARMResource)
	if err != nil {
		return StaticSiteUserProvidedFunctionAppARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *staticSiteUserProvidedFunctionAppARMResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *staticSiteUserProvidedFunctionAppARMResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (StaticSiteUserProvidedFunctionAppARMResourceResponse, error) {
	respType := StaticSiteUserProvidedFunctionAppARMResourceResponse{StaticSiteUserProvidedFunctionAppARMResource: &StaticSiteUserProvidedFunctionAppARMResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.StaticSiteUserProvidedFunctionAppARMResource)
	if err != nil {
		return StaticSiteUserProvidedFunctionAppARMResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// StorageMigrationResponsePoller provides polling facilities until the operation reaches a terminal state.
type StorageMigrationResponsePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final StorageMigrationResponseResponse will be returned.
	FinalResponse(ctx context.Context) (StorageMigrationResponseResponse, error)
}

type storageMigrationResponsePoller struct {
	pt *armcore.LROPoller
}

func (p *storageMigrationResponsePoller) Done() bool {
	return p.pt.Done()
}

func (p *storageMigrationResponsePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *storageMigrationResponsePoller) FinalResponse(ctx context.Context) (StorageMigrationResponseResponse, error) {
	respType := StorageMigrationResponseResponse{StorageMigrationResponse: &StorageMigrationResponse{}}
	resp, err := p.pt.FinalResponse(ctx, respType.StorageMigrationResponse)
	if err != nil {
		return StorageMigrationResponseResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *storageMigrationResponsePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *storageMigrationResponsePoller) pollUntilDone(ctx context.Context, freq time.Duration) (StorageMigrationResponseResponse, error) {
	respType := StorageMigrationResponseResponse{StorageMigrationResponse: &StorageMigrationResponse{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.StorageMigrationResponse)
	if err != nil {
		return StorageMigrationResponseResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// UserPoller provides polling facilities until the operation reaches a terminal state.
type UserPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final UserResponse will be returned.
	FinalResponse(ctx context.Context) (UserResponse, error)
}

type userPoller struct {
	pt *armcore.LROPoller
}

func (p *userPoller) Done() bool {
	return p.pt.Done()
}

func (p *userPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *userPoller) FinalResponse(ctx context.Context) (UserResponse, error) {
	respType := UserResponse{User: &User{}}
	resp, err := p.pt.FinalResponse(ctx, respType.User)
	if err != nil {
		return UserResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *userPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *userPoller) pollUntilDone(ctx context.Context, freq time.Duration) (UserResponse, error) {
	respType := UserResponse{User: &User{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.User)
	if err != nil {
		return UserResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// WebAppCollectionPagerPoller provides polling facilities until the operation reaches a terminal state.
type WebAppCollectionPagerPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final WebAppCollectionPager will be returned.
	FinalResponse(ctx context.Context) (WebAppCollectionPager, error)
}

type webAppCollectionPagerPoller struct {
	pt          *armcore.LROPoller
	errHandler  webAppCollectionHandleError
	respHandler webAppCollectionHandleResponse
	statusCodes []int
}

func (p *webAppCollectionPagerPoller) Done() bool {
	return p.pt.Done()
}

func (p *webAppCollectionPagerPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *webAppCollectionPagerPoller) FinalResponse(ctx context.Context) (WebAppCollectionPager, error) {
	respType := &webAppCollectionPager{}
	resp, err := p.pt.FinalResponse(ctx, respType)
	if err != nil {
		return nil, err
	}
	return p.handleResponse(&azcore.Response{Response: resp})
}

func (p *webAppCollectionPagerPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *webAppCollectionPagerPoller) pollUntilDone(ctx context.Context, freq time.Duration) (WebAppCollectionPager, error) {
	respType := &webAppCollectionPager{}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType)
	if err != nil {
		return nil, err
	}
	return p.handleResponse(&azcore.Response{Response: resp})
}

func (p *webAppCollectionPagerPoller) handleResponse(resp *azcore.Response) (WebAppCollectionPager, error) {
	return &webAppCollectionPager{
		pipeline:  p.pt.Pipeline,
		resp:      resp,
		errorer:   p.errHandler,
		responder: p.respHandler,
		advancer: func(ctx context.Context, resp WebAppCollectionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.WebAppCollection.NextLink)
		},
		statusCodes: p.statusCodes,
	}, nil
}

// WorkerPoolResourcePoller provides polling facilities until the operation reaches a terminal state.
type WorkerPoolResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final WorkerPoolResourceResponse will be returned.
	FinalResponse(ctx context.Context) (WorkerPoolResourceResponse, error)
}

type workerPoolResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *workerPoolResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *workerPoolResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *workerPoolResourcePoller) FinalResponse(ctx context.Context) (WorkerPoolResourceResponse, error) {
	respType := WorkerPoolResourceResponse{WorkerPoolResource: &WorkerPoolResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.WorkerPoolResource)
	if err != nil {
		return WorkerPoolResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *workerPoolResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *workerPoolResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (WorkerPoolResourceResponse, error) {
	respType := WorkerPoolResourceResponse{WorkerPoolResource: &WorkerPoolResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.WorkerPoolResource)
	if err != nil {
		return WorkerPoolResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
