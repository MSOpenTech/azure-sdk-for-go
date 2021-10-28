//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// BackupsGetResponse contains the response from method Backups.Get.
type BackupsGetResponse struct {
	BackupsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BackupsGetResult contains the result from method Backups.Get.
type BackupsGetResult struct {
	ServerBackup
}

// BackupsListByServerResponse contains the response from method Backups.ListByServer.
type BackupsListByServerResponse struct {
	BackupsListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BackupsListByServerResult contains the result from method Backups.ListByServer.
type BackupsListByServerResult struct {
	ServerBackupListResult
}

// CheckNameAvailabilityExecuteResponse contains the response from method CheckNameAvailability.Execute.
type CheckNameAvailabilityExecuteResponse struct {
	CheckNameAvailabilityExecuteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CheckNameAvailabilityExecuteResult contains the result from method CheckNameAvailability.Execute.
type CheckNameAvailabilityExecuteResult struct {
	NameAvailability
}

// CheckVirtualNetworkSubnetUsageExecuteResponse contains the response from method CheckVirtualNetworkSubnetUsage.Execute.
type CheckVirtualNetworkSubnetUsageExecuteResponse struct {
	CheckVirtualNetworkSubnetUsageExecuteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CheckVirtualNetworkSubnetUsageExecuteResult contains the result from method CheckVirtualNetworkSubnetUsage.Execute.
type CheckVirtualNetworkSubnetUsageExecuteResult struct {
	VirtualNetworkSubnetUsageResult
}

// ConfigurationsBatchUpdatePollerResponse contains the response from method Configurations.BatchUpdate.
type ConfigurationsBatchUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ConfigurationsBatchUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ConfigurationsBatchUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ConfigurationsBatchUpdateResponse, error) {
	respType := ConfigurationsBatchUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ConfigurationListResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ConfigurationsBatchUpdatePollerResponse from the provided client and resume token.
func (l *ConfigurationsBatchUpdatePollerResponse) Resume(ctx context.Context, client *ConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ConfigurationsClient.BatchUpdate", token, client.pl, client.batchUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &ConfigurationsBatchUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ConfigurationsBatchUpdateResponse contains the response from method Configurations.BatchUpdate.
type ConfigurationsBatchUpdateResponse struct {
	ConfigurationsBatchUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsBatchUpdateResult contains the result from method Configurations.BatchUpdate.
type ConfigurationsBatchUpdateResult struct {
	ConfigurationListResult
}

// ConfigurationsGetResponse contains the response from method Configurations.Get.
type ConfigurationsGetResponse struct {
	ConfigurationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsGetResult contains the result from method Configurations.Get.
type ConfigurationsGetResult struct {
	Configuration
}

// ConfigurationsListByServerResponse contains the response from method Configurations.ListByServer.
type ConfigurationsListByServerResponse struct {
	ConfigurationsListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsListByServerResult contains the result from method Configurations.ListByServer.
type ConfigurationsListByServerResult struct {
	ConfigurationListResult
}

// ConfigurationsUpdatePollerResponse contains the response from method Configurations.Update.
type ConfigurationsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ConfigurationsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ConfigurationsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ConfigurationsUpdateResponse, error) {
	respType := ConfigurationsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Configuration)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ConfigurationsUpdatePollerResponse from the provided client and resume token.
func (l *ConfigurationsUpdatePollerResponse) Resume(ctx context.Context, client *ConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ConfigurationsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ConfigurationsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ConfigurationsUpdateResponse contains the response from method Configurations.Update.
type ConfigurationsUpdateResponse struct {
	ConfigurationsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsUpdateResult contains the result from method Configurations.Update.
type ConfigurationsUpdateResult struct {
	Configuration
}

// DatabasesCreateOrUpdatePollerResponse contains the response from method Databases.CreateOrUpdate.
type DatabasesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DatabasesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DatabasesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DatabasesCreateOrUpdateResponse, error) {
	respType := DatabasesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Database)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DatabasesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *DatabasesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *DatabasesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DatabasesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &DatabasesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DatabasesCreateOrUpdateResponse contains the response from method Databases.CreateOrUpdate.
type DatabasesCreateOrUpdateResponse struct {
	DatabasesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesCreateOrUpdateResult contains the result from method Databases.CreateOrUpdate.
type DatabasesCreateOrUpdateResult struct {
	Database
}

// DatabasesDeletePollerResponse contains the response from method Databases.Delete.
type DatabasesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DatabasesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DatabasesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DatabasesDeleteResponse, error) {
	respType := DatabasesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DatabasesDeletePollerResponse from the provided client and resume token.
func (l *DatabasesDeletePollerResponse) Resume(ctx context.Context, client *DatabasesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DatabasesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &DatabasesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DatabasesDeleteResponse contains the response from method Databases.Delete.
type DatabasesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesGetResponse contains the response from method Databases.Get.
type DatabasesGetResponse struct {
	DatabasesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesGetResult contains the result from method Databases.Get.
type DatabasesGetResult struct {
	Database
}

// DatabasesListByServerResponse contains the response from method Databases.ListByServer.
type DatabasesListByServerResponse struct {
	DatabasesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DatabasesListByServerResult contains the result from method Databases.ListByServer.
type DatabasesListByServerResult struct {
	DatabaseListResult
}

// FirewallRulesCreateOrUpdatePollerResponse contains the response from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FirewallRulesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FirewallRulesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FirewallRulesCreateOrUpdateResponse, error) {
	respType := FirewallRulesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.FirewallRule)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a FirewallRulesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *FirewallRulesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *FirewallRulesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FirewallRulesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &FirewallRulesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// FirewallRulesCreateOrUpdateResponse contains the response from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdateResponse struct {
	FirewallRulesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesCreateOrUpdateResult contains the result from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdateResult struct {
	FirewallRule
}

// FirewallRulesDeletePollerResponse contains the response from method FirewallRules.Delete.
type FirewallRulesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *FirewallRulesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l FirewallRulesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (FirewallRulesDeleteResponse, error) {
	respType := FirewallRulesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a FirewallRulesDeletePollerResponse from the provided client and resume token.
func (l *FirewallRulesDeletePollerResponse) Resume(ctx context.Context, client *FirewallRulesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("FirewallRulesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &FirewallRulesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// FirewallRulesDeleteResponse contains the response from method FirewallRules.Delete.
type FirewallRulesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesGetResponse contains the response from method FirewallRules.Get.
type FirewallRulesGetResponse struct {
	FirewallRulesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesGetResult contains the result from method FirewallRules.Get.
type FirewallRulesGetResult struct {
	FirewallRule
}

// FirewallRulesListByServerResponse contains the response from method FirewallRules.ListByServer.
type FirewallRulesListByServerResponse struct {
	FirewallRulesListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesListByServerResult contains the result from method FirewallRules.ListByServer.
type FirewallRulesListByServerResult struct {
	FirewallRuleListResult
}

// GetPrivateDNSZoneSuffixExecuteResponse contains the response from method GetPrivateDNSZoneSuffix.Execute.
type GetPrivateDNSZoneSuffixExecuteResponse struct {
	GetPrivateDNSZoneSuffixExecuteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GetPrivateDNSZoneSuffixExecuteResult contains the result from method GetPrivateDNSZoneSuffix.Execute.
type GetPrivateDNSZoneSuffixExecuteResult struct {
	GetPrivateDNSZoneSuffixResponse
}

// LocationBasedCapabilitiesListResponse contains the response from method LocationBasedCapabilities.List.
type LocationBasedCapabilitiesListResponse struct {
	LocationBasedCapabilitiesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LocationBasedCapabilitiesListResult contains the result from method LocationBasedCapabilities.List.
type LocationBasedCapabilitiesListResult struct {
	CapabilitiesListResult
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// ReplicasListByServerResponse contains the response from method Replicas.ListByServer.
type ReplicasListByServerResponse struct {
	ReplicasListByServerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReplicasListByServerResult contains the result from method Replicas.ListByServer.
type ReplicasListByServerResult struct {
	ServerListResult
}

// ServersCreatePollerResponse contains the response from method Servers.Create.
type ServersCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersCreateResponse, error) {
	respType := ServersCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Server)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersCreatePollerResponse from the provided client and resume token.
func (l *ServersCreatePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &ServersCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersCreateResponse contains the response from method Servers.Create.
type ServersCreateResponse struct {
	ServersCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersCreateResult contains the result from method Servers.Create.
type ServersCreateResult struct {
	Server
}

// ServersDeletePollerResponse contains the response from method Servers.Delete.
type ServersDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersDeleteResponse, error) {
	respType := ServersDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersDeletePollerResponse from the provided client and resume token.
func (l *ServersDeletePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ServersDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersDeleteResponse contains the response from method Servers.Delete.
type ServersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersFailoverPollerResponse contains the response from method Servers.Failover.
type ServersFailoverPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersFailoverPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersFailoverPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersFailoverResponse, error) {
	respType := ServersFailoverResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersFailoverPollerResponse from the provided client and resume token.
func (l *ServersFailoverPollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Failover", token, client.pl, client.failoverHandleError)
	if err != nil {
		return err
	}
	poller := &ServersFailoverPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersFailoverResponse contains the response from method Servers.Failover.
type ServersFailoverResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersGetResponse contains the response from method Servers.Get.
type ServersGetResponse struct {
	ServersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersGetResult contains the result from method Servers.Get.
type ServersGetResult struct {
	Server
}

// ServersListByResourceGroupResponse contains the response from method Servers.ListByResourceGroup.
type ServersListByResourceGroupResponse struct {
	ServersListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListByResourceGroupResult contains the result from method Servers.ListByResourceGroup.
type ServersListByResourceGroupResult struct {
	ServerListResult
}

// ServersListResponse contains the response from method Servers.List.
type ServersListResponse struct {
	ServersListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersListResult contains the result from method Servers.List.
type ServersListResult struct {
	ServerListResult
}

// ServersRestartPollerResponse contains the response from method Servers.Restart.
type ServersRestartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersRestartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersRestartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersRestartResponse, error) {
	respType := ServersRestartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersRestartPollerResponse from the provided client and resume token.
func (l *ServersRestartPollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Restart", token, client.pl, client.restartHandleError)
	if err != nil {
		return err
	}
	poller := &ServersRestartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersRestartResponse contains the response from method Servers.Restart.
type ServersRestartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersStartPollerResponse contains the response from method Servers.Start.
type ServersStartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersStartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersStartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersStartResponse, error) {
	respType := ServersStartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersStartPollerResponse from the provided client and resume token.
func (l *ServersStartPollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Start", token, client.pl, client.startHandleError)
	if err != nil {
		return err
	}
	poller := &ServersStartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersStartResponse contains the response from method Servers.Start.
type ServersStartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersStopPollerResponse contains the response from method Servers.Stop.
type ServersStopPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersStopPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersStopPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersStopResponse, error) {
	respType := ServersStopResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersStopPollerResponse from the provided client and resume token.
func (l *ServersStopPollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Stop", token, client.pl, client.stopHandleError)
	if err != nil {
		return err
	}
	poller := &ServersStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersStopResponse contains the response from method Servers.Stop.
type ServersStopResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdatePollerResponse contains the response from method Servers.Update.
type ServersUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ServersUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ServersUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ServersUpdateResponse, error) {
	respType := ServersUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Server)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ServersUpdatePollerResponse from the provided client and resume token.
func (l *ServersUpdatePollerResponse) Resume(ctx context.Context, client *ServersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ServersClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ServersUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ServersUpdateResponse contains the response from method Servers.Update.
type ServersUpdateResponse struct {
	ServersUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ServersUpdateResult contains the result from method Servers.Update.
type ServersUpdateResult struct {
	Server
}
