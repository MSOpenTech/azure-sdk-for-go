// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package commerce

import original "github.com/Azure/azure-sdk-for-go/services/azsadmin/mgmt/2015-06-01-preview/commerce"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type Resource = original.Resource
type UsageAggregate = original.UsageAggregate
type UsageAggregateModel = original.UsageAggregateModel
type UsageAggregatePage = original.UsageAggregatePage
type SubscriberUsageAggregatesClient = original.SubscriberUsageAggregatesClient

func NewSubscriberUsageAggregatesClient(subscriptionID string) SubscriberUsageAggregatesClient {
	return original.NewSubscriberUsageAggregatesClient(subscriptionID)
}
func NewSubscriberUsageAggregatesClientWithBaseURI(baseURI string, subscriptionID string) SubscriberUsageAggregatesClient {
	return original.NewSubscriberUsageAggregatesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
