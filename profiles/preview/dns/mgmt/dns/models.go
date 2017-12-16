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

package dns

import original "github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2016-04-01/dns"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type HTTPStatusCode = original.HTTPStatusCode

const (
	Accepted                     HTTPStatusCode = original.Accepted
	Ambiguous                    HTTPStatusCode = original.Ambiguous
	BadGateway                   HTTPStatusCode = original.BadGateway
	BadRequest                   HTTPStatusCode = original.BadRequest
	Conflict                     HTTPStatusCode = original.Conflict
	Continue                     HTTPStatusCode = original.Continue
	Created                      HTTPStatusCode = original.Created
	ExpectationFailed            HTTPStatusCode = original.ExpectationFailed
	Forbidden                    HTTPStatusCode = original.Forbidden
	Found                        HTTPStatusCode = original.Found
	GatewayTimeout               HTTPStatusCode = original.GatewayTimeout
	Gone                         HTTPStatusCode = original.Gone
	HTTPVersionNotSupported      HTTPStatusCode = original.HTTPVersionNotSupported
	InternalServerError          HTTPStatusCode = original.InternalServerError
	LengthRequired               HTTPStatusCode = original.LengthRequired
	MethodNotAllowed             HTTPStatusCode = original.MethodNotAllowed
	Moved                        HTTPStatusCode = original.Moved
	MovedPermanently             HTTPStatusCode = original.MovedPermanently
	MultipleChoices              HTTPStatusCode = original.MultipleChoices
	NoContent                    HTTPStatusCode = original.NoContent
	NonAuthoritativeInformation  HTTPStatusCode = original.NonAuthoritativeInformation
	NotAcceptable                HTTPStatusCode = original.NotAcceptable
	NotFound                     HTTPStatusCode = original.NotFound
	NotImplemented               HTTPStatusCode = original.NotImplemented
	NotModified                  HTTPStatusCode = original.NotModified
	OK                           HTTPStatusCode = original.OK
	PartialContent               HTTPStatusCode = original.PartialContent
	PaymentRequired              HTTPStatusCode = original.PaymentRequired
	PreconditionFailed           HTTPStatusCode = original.PreconditionFailed
	ProxyAuthenticationRequired  HTTPStatusCode = original.ProxyAuthenticationRequired
	Redirect                     HTTPStatusCode = original.Redirect
	RedirectKeepVerb             HTTPStatusCode = original.RedirectKeepVerb
	RedirectMethod               HTTPStatusCode = original.RedirectMethod
	RequestedRangeNotSatisfiable HTTPStatusCode = original.RequestedRangeNotSatisfiable
	RequestEntityTooLarge        HTTPStatusCode = original.RequestEntityTooLarge
	RequestTimeout               HTTPStatusCode = original.RequestTimeout
	RequestURITooLong            HTTPStatusCode = original.RequestURITooLong
	ResetContent                 HTTPStatusCode = original.ResetContent
	SeeOther                     HTTPStatusCode = original.SeeOther
	ServiceUnavailable           HTTPStatusCode = original.ServiceUnavailable
	SwitchingProtocols           HTTPStatusCode = original.SwitchingProtocols
	TemporaryRedirect            HTTPStatusCode = original.TemporaryRedirect
	Unauthorized                 HTTPStatusCode = original.Unauthorized
	UnsupportedMediaType         HTTPStatusCode = original.UnsupportedMediaType
	Unused                       HTTPStatusCode = original.Unused
	UpgradeRequired              HTTPStatusCode = original.UpgradeRequired
	UseProxy                     HTTPStatusCode = original.UseProxy
)

type OperationStatus = original.OperationStatus

const (
	Failed     OperationStatus = original.Failed
	InProgress OperationStatus = original.InProgress
	Succeeded  OperationStatus = original.Succeeded
)

type RecordType = original.RecordType

const (
	A     RecordType = original.A
	AAAA  RecordType = original.AAAA
	CNAME RecordType = original.CNAME
	MX    RecordType = original.MX
	NS    RecordType = original.NS
	PTR   RecordType = original.PTR
	SOA   RecordType = original.SOA
	SRV   RecordType = original.SRV
	TXT   RecordType = original.TXT
)

type AaaaRecord = original.AaaaRecord
type ARecord = original.ARecord
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CnameRecord = original.CnameRecord
type MxRecord = original.MxRecord
type NsRecord = original.NsRecord
type PtrRecord = original.PtrRecord
type RecordSet = original.RecordSet
type RecordSetListResult = original.RecordSetListResult
type RecordSetListResultIterator = original.RecordSetListResultIterator
type RecordSetListResultPage = original.RecordSetListResultPage
type RecordSetProperties = original.RecordSetProperties
type RecordSetUpdateParameters = original.RecordSetUpdateParameters
type Resource = original.Resource
type SoaRecord = original.SoaRecord
type SrvRecord = original.SrvRecord
type SubResource = original.SubResource
type TxtRecord = original.TxtRecord
type Zone = original.Zone
type ZoneDeleteResult = original.ZoneDeleteResult
type ZoneListResult = original.ZoneListResult
type ZoneListResultIterator = original.ZoneListResultIterator
type ZoneListResultPage = original.ZoneListResultPage
type ZoneProperties = original.ZoneProperties
type ZonesDeleteFuture = original.ZonesDeleteFuture
type RecordSetsClient = original.RecordSetsClient
type ZonesClient = original.ZonesClient

func NewRecordSetsClient(subscriptionID string) RecordSetsClient {
	return original.NewRecordSetsClient(subscriptionID)
}
func NewRecordSetsClientWithBaseURI(baseURI string, subscriptionID string) RecordSetsClient {
	return original.NewRecordSetsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewZonesClient(subscriptionID string) ZonesClient {
	return original.NewZonesClient(subscriptionID)
}
func NewZonesClientWithBaseURI(baseURI string, subscriptionID string) ZonesClient {
	return original.NewZonesClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
