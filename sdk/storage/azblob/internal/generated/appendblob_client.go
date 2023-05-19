//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

func (client *AppendBlobClient) Endpoint() string {
	return client.endpoint
}

func (client *AppendBlobClient) Pipeline() runtime.Pipeline {
	return client.internal.Pipeline()
}
