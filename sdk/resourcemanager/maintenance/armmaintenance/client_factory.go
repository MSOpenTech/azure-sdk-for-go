//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewApplyUpdateForResourceGroupClient creates a new instance of ApplyUpdateForResourceGroupClient.
func (c *ClientFactory) NewApplyUpdateForResourceGroupClient() *ApplyUpdateForResourceGroupClient {
	subClient, _ := NewApplyUpdateForResourceGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewApplyUpdatesClient creates a new instance of ApplyUpdatesClient.
func (c *ClientFactory) NewApplyUpdatesClient() *ApplyUpdatesClient {
	subClient, _ := NewApplyUpdatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationAssignmentsClient creates a new instance of ConfigurationAssignmentsClient.
func (c *ClientFactory) NewConfigurationAssignmentsClient() *ConfigurationAssignmentsClient {
	subClient, _ := NewConfigurationAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationAssignmentsForResourceGroupClient creates a new instance of ConfigurationAssignmentsForResourceGroupClient.
func (c *ClientFactory) NewConfigurationAssignmentsForResourceGroupClient() *ConfigurationAssignmentsForResourceGroupClient {
	subClient, _ := NewConfigurationAssignmentsForResourceGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationAssignmentsForSubscriptionsClient creates a new instance of ConfigurationAssignmentsForSubscriptionsClient.
func (c *ClientFactory) NewConfigurationAssignmentsForSubscriptionsClient() *ConfigurationAssignmentsForSubscriptionsClient {
	subClient, _ := NewConfigurationAssignmentsForSubscriptionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationAssignmentsWithinSubscriptionClient creates a new instance of ConfigurationAssignmentsWithinSubscriptionClient.
func (c *ClientFactory) NewConfigurationAssignmentsWithinSubscriptionClient() *ConfigurationAssignmentsWithinSubscriptionClient {
	subClient, _ := NewConfigurationAssignmentsWithinSubscriptionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient.
func (c *ClientFactory) NewConfigurationsClient() *ConfigurationsClient {
	subClient, _ := NewConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationsForResourceGroupClient creates a new instance of ConfigurationsForResourceGroupClient.
func (c *ClientFactory) NewConfigurationsForResourceGroupClient() *ConfigurationsForResourceGroupClient {
	subClient, _ := NewConfigurationsForResourceGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPublicMaintenanceConfigurationsClient creates a new instance of PublicMaintenanceConfigurationsClient.
func (c *ClientFactory) NewPublicMaintenanceConfigurationsClient() *PublicMaintenanceConfigurationsClient {
	subClient, _ := NewPublicMaintenanceConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewUpdatesClient creates a new instance of UpdatesClient.
func (c *ClientFactory) NewUpdatesClient() *UpdatesClient {
	subClient, _ := NewUpdatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
