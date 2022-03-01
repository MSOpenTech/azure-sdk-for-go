//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azappconfig

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// AddConfigurationSettingOptions contains the optional parameters for the AddConfigurationSetting method.
type AddConfigurationSettingOptions struct {
}

// AddConfigurationSetting creates a configuration setting only if the setting does not already exist in the configuration store.
func (c *Client) AddConfigurationSetting(ctx context.Context, setting Setting, options *AddConfigurationSettingOptions) (AddConfigurationSettingResult, error) {
	etagAny := azcore.ETagAny
	resp, err := c.appConfigClient.PutKeyValue(ctx, *setting.Key, setting.toGeneratedPutOptions(nil, &etagAny))
	if err != nil {
		return AddConfigurationSettingResult{}, err
	}

	return (AddConfigurationSettingResult)(fromGeneratedPut(resp)), nil
}
