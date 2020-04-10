// +build go1.9

// Copyright 2020 Microsoft Corporation
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

package maps

import original "github.com/Azure/azure-sdk-for-go/services/maps/mgmt/2018-05-01/maps"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type Account = original.Account
type AccountCreateParameters = original.AccountCreateParameters
type AccountKeys = original.AccountKeys
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type Accounts = original.Accounts
type AccountsClient = original.AccountsClient
type AccountsMoveRequest = original.AccountsMoveRequest
type BaseClient = original.BaseClient
type Error = original.Error
type ErrorDetailsItem = original.ErrorDetailsItem
type KeySpecification = original.KeySpecification
type Operations = original.Operations
type OperationsValueItem = original.OperationsValueItem
type OperationsValueItemDisplay = original.OperationsValueItemDisplay
type Resource = original.Resource
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
