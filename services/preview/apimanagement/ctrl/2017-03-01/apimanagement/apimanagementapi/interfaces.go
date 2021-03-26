package apimanagementapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/apimanagement/ctrl/2017-03-01/apimanagement"
	"github.com/Azure/go-autorest/autorest"
)

// PolicyClientAPI contains the set of methods on the PolicyClient type.
type PolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, parameters apimanagement.PolicyContract) (result apimanagement.PolicyContract, err error)
	Delete(ctx context.Context, apimBaseURL string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string) (result apimanagement.PolicyContract, err error)
	List(ctx context.Context, apimBaseURL string, scope apimanagement.PolicyScopeContract) (result apimanagement.PolicyCollection, err error)
}

var _ PolicyClientAPI = (*apimanagement.PolicyClient)(nil)

// PolicySnippetsClientAPI contains the set of methods on the PolicySnippetsClient type.
type PolicySnippetsClientAPI interface {
	List(ctx context.Context, apimBaseURL string, scope apimanagement.PolicyScopeContract) (result apimanagement.PolicySnippetsCollection, err error)
}

var _ PolicySnippetsClientAPI = (*apimanagement.PolicySnippetsClient)(nil)

// RegionsClientAPI contains the set of methods on the RegionsClient type.
type RegionsClientAPI interface {
	List(ctx context.Context, apimBaseURL string) (result apimanagement.RegionListResult, err error)
}

var _ RegionsClientAPI = (*apimanagement.RegionsClient)(nil)

// APIClientAPI contains the set of methods on the APIClient type.
type APIClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, apiid string, parameters apimanagement.APICreateOrUpdateParameter, ifMatch string) (result apimanagement.APIContract, err error)
	Delete(ctx context.Context, apimBaseURL string, apiid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, apiid string) (result apimanagement.APIContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, apiid string, parameters apimanagement.APIUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ APIClientAPI = (*apimanagement.APIClient)(nil)

// APIOperationClientAPI contains the set of methods on the APIOperationClient type.
type APIOperationClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, apiid string, operationID string, parameters apimanagement.OperationContract) (result apimanagement.OperationContract, err error)
	Delete(ctx context.Context, apimBaseURL string, apiid string, operationID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, apiid string, operationID string) (result apimanagement.OperationContract, err error)
	ListByAPI(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.OperationCollectionPage, err error)
	ListByAPIComplete(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.OperationCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, apiid string, operationID string, parameters apimanagement.OperationUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ APIOperationClientAPI = (*apimanagement.APIOperationClient)(nil)

// APIOperationPolicyClientAPI contains the set of methods on the APIOperationPolicyClient type.
type APIOperationPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, apiid string, operationID string, parameters apimanagement.PolicyContract, ifMatch string) (result apimanagement.PolicyContract, err error)
	Delete(ctx context.Context, apimBaseURL string, apiid string, operationID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, apiid string, operationID string) (result apimanagement.PolicyContract, err error)
	ListByOperation(ctx context.Context, apimBaseURL string, apiid string, operationID string) (result apimanagement.PolicyCollection, err error)
}

var _ APIOperationPolicyClientAPI = (*apimanagement.APIOperationPolicyClient)(nil)

// APIProductClientAPI contains the set of methods on the APIProductClient type.
type APIProductClientAPI interface {
	ListByApis(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.ProductCollectionPage, err error)
	ListByApisComplete(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.ProductCollectionIterator, err error)
}

var _ APIProductClientAPI = (*apimanagement.APIProductClient)(nil)

// APIPolicyClientAPI contains the set of methods on the APIPolicyClient type.
type APIPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, apiid string, parameters apimanagement.PolicyContract, ifMatch string) (result apimanagement.PolicyContract, err error)
	Delete(ctx context.Context, apimBaseURL string, apiid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, apiid string) (result apimanagement.PolicyContract, err error)
	ListByAPI(ctx context.Context, apimBaseURL string, apiid string) (result apimanagement.PolicyCollection, err error)
}

var _ APIPolicyClientAPI = (*apimanagement.APIPolicyClient)(nil)

// APISchemaClientAPI contains the set of methods on the APISchemaClient type.
type APISchemaClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, apiid string, schemaID string, parameters apimanagement.SchemaContract, ifMatch string) (result apimanagement.SchemaContract, err error)
	Delete(ctx context.Context, apimBaseURL string, apiid string, schemaID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, apiid string, schemaID string) (result apimanagement.SchemaContract, err error)
	ListByAPI(ctx context.Context, apimBaseURL string, apiid string) (result apimanagement.SchemaCollectionPage, err error)
	ListByAPIComplete(ctx context.Context, apimBaseURL string, apiid string) (result apimanagement.SchemaCollectionIterator, err error)
}

var _ APISchemaClientAPI = (*apimanagement.APISchemaClient)(nil)

// AuthorizationServerClientAPI contains the set of methods on the AuthorizationServerClient type.
type AuthorizationServerClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, authsid string, parameters apimanagement.AuthorizationServerContract) (result apimanagement.AuthorizationServerContract, err error)
	Delete(ctx context.Context, apimBaseURL string, authsid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, authsid string) (result apimanagement.AuthorizationServerContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.AuthorizationServerCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.AuthorizationServerCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, authsid string, parameters apimanagement.AuthorizationServerUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ AuthorizationServerClientAPI = (*apimanagement.AuthorizationServerClient)(nil)

// BackendClientAPI contains the set of methods on the BackendClient type.
type BackendClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, backendid string, parameters apimanagement.BackendContract) (result apimanagement.BackendContract, err error)
	Delete(ctx context.Context, apimBaseURL string, backendid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, backendid string) (result apimanagement.BackendContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.BackendCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.BackendCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, backendid string, parameters apimanagement.BackendUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ BackendClientAPI = (*apimanagement.BackendClient)(nil)

// CertificateClientAPI contains the set of methods on the CertificateClient type.
type CertificateClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, certificateID string, parameters apimanagement.CertificateCreateOrUpdateParameters, ifMatch string) (result apimanagement.CertificateContract, err error)
	Delete(ctx context.Context, apimBaseURL string, certificateID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, certificateID string) (result apimanagement.CertificateContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.CertificateCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.CertificateCollectionIterator, err error)
}

var _ CertificateClientAPI = (*apimanagement.CertificateClient)(nil)

// EmailTemplateClientAPI contains the set of methods on the EmailTemplateClient type.
type EmailTemplateClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, templateName apimanagement.TemplateName, parameters apimanagement.EmailTemplateUpdateParameters) (result apimanagement.EmailTemplateContract, err error)
	Delete(ctx context.Context, apimBaseURL string, templateName apimanagement.TemplateName, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, templateName apimanagement.TemplateName) (result apimanagement.EmailTemplateContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.EmailTemplateCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.EmailTemplateCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, templateName apimanagement.TemplateName, parameters apimanagement.EmailTemplateUpdateParameters) (result autorest.Response, err error)
}

var _ EmailTemplateClientAPI = (*apimanagement.EmailTemplateClient)(nil)

// GroupClientAPI contains the set of methods on the GroupClient type.
type GroupClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, groupID string, parameters apimanagement.GroupCreateParameters) (result apimanagement.GroupContract, err error)
	Delete(ctx context.Context, apimBaseURL string, groupID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, groupID string) (result apimanagement.GroupContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, groupID string, parameters apimanagement.GroupUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ GroupClientAPI = (*apimanagement.GroupClient)(nil)

// GroupUserClientAPI contains the set of methods on the GroupUserClient type.
type GroupUserClientAPI interface {
	Create(ctx context.Context, apimBaseURL string, groupID string, UID string) (result apimanagement.UserContract, err error)
	Delete(ctx context.Context, apimBaseURL string, groupID string, UID string) (result autorest.Response, err error)
	List(ctx context.Context, apimBaseURL string, groupID string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, groupID string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionIterator, err error)
}

var _ GroupUserClientAPI = (*apimanagement.GroupUserClient)(nil)

// IdentityProviderClientAPI contains the set of methods on the IdentityProviderClient type.
type IdentityProviderClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, identityProviderName apimanagement.IdentityProviderType, parameters apimanagement.IdentityProviderContract) (result apimanagement.IdentityProviderContract, err error)
	Delete(ctx context.Context, apimBaseURL string, identityProviderName apimanagement.IdentityProviderType, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, identityProviderName apimanagement.IdentityProviderType) (result apimanagement.IdentityProviderContract, err error)
	List(ctx context.Context, apimBaseURL string) (result apimanagement.IdentityProviderList, err error)
	Update(ctx context.Context, apimBaseURL string, identityProviderName apimanagement.IdentityProviderType, parameters apimanagement.IdentityProviderUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ IdentityProviderClientAPI = (*apimanagement.IdentityProviderClient)(nil)

// LoggerClientAPI contains the set of methods on the LoggerClient type.
type LoggerClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, loggerid string, parameters apimanagement.LoggerContract) (result apimanagement.LoggerContract, err error)
	Delete(ctx context.Context, apimBaseURL string, loggerid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, loggerid string) (result apimanagement.LoggerContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.LoggerCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.LoggerCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, loggerid string, parameters apimanagement.LoggerUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ LoggerClientAPI = (*apimanagement.LoggerClient)(nil)

// OpenIDConnectProviderClientAPI contains the set of methods on the OpenIDConnectProviderClient type.
type OpenIDConnectProviderClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, opid string, parameters apimanagement.OpenidConnectProviderContract) (result apimanagement.OpenidConnectProviderContract, err error)
	Delete(ctx context.Context, apimBaseURL string, opid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, opid string) (result apimanagement.OpenidConnectProviderContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.OpenIDConnectProviderCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.OpenIDConnectProviderCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, opid string, parameters apimanagement.OpenidConnectProviderUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ OpenIDConnectProviderClientAPI = (*apimanagement.OpenIDConnectProviderClient)(nil)

// SignInSettingsClientAPI contains the set of methods on the SignInSettingsClient type.
type SignInSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, parameters apimanagement.PortalSigninSettings) (result apimanagement.PortalSigninSettings, err error)
	Get(ctx context.Context, apimBaseURL string) (result apimanagement.PortalSigninSettings, err error)
	Update(ctx context.Context, apimBaseURL string, parameters apimanagement.PortalSigninSettings, ifMatch string) (result autorest.Response, err error)
}

var _ SignInSettingsClientAPI = (*apimanagement.SignInSettingsClient)(nil)

// SignUpSettingsClientAPI contains the set of methods on the SignUpSettingsClient type.
type SignUpSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, parameters apimanagement.PortalSignupSettings) (result apimanagement.PortalSignupSettings, err error)
	Get(ctx context.Context, apimBaseURL string) (result apimanagement.PortalSignupSettings, err error)
	Update(ctx context.Context, apimBaseURL string, parameters apimanagement.PortalSignupSettings, ifMatch string) (result autorest.Response, err error)
}

var _ SignUpSettingsClientAPI = (*apimanagement.SignUpSettingsClient)(nil)

// DelegationSettingsClientAPI contains the set of methods on the DelegationSettingsClient type.
type DelegationSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, parameters apimanagement.PortalDelegationSettings) (result apimanagement.PortalDelegationSettings, err error)
	Get(ctx context.Context, apimBaseURL string) (result apimanagement.PortalDelegationSettings, err error)
	Update(ctx context.Context, apimBaseURL string, parameters apimanagement.PortalDelegationSettings, ifMatch string) (result autorest.Response, err error)
}

var _ DelegationSettingsClientAPI = (*apimanagement.DelegationSettingsClient)(nil)

// ProductClientAPI contains the set of methods on the ProductClient type.
type ProductClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, productID string, parameters apimanagement.ProductContract) (result apimanagement.ProductContract, err error)
	Delete(ctx context.Context, apimBaseURL string, productID string, ifMatch string, deleteSubscriptions *bool) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, productID string) (result apimanagement.ProductContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32, expandGroups *bool) (result apimanagement.ProductCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32, expandGroups *bool) (result apimanagement.ProductCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, productID string, parameters apimanagement.ProductUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ ProductClientAPI = (*apimanagement.ProductClient)(nil)

// ProductAPIClientAPI contains the set of methods on the ProductAPIClient type.
type ProductAPIClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, productID string, apiid string) (result apimanagement.APIContract, err error)
	Delete(ctx context.Context, apimBaseURL string, productID string, apiid string) (result autorest.Response, err error)
	ListByProduct(ctx context.Context, apimBaseURL string, productID string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionPage, err error)
	ListByProductComplete(ctx context.Context, apimBaseURL string, productID string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionIterator, err error)
}

var _ ProductAPIClientAPI = (*apimanagement.ProductAPIClient)(nil)

// ProductGroupClientAPI contains the set of methods on the ProductGroupClient type.
type ProductGroupClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, productID string, groupID string) (result apimanagement.GroupContract, err error)
	Delete(ctx context.Context, apimBaseURL string, productID string, groupID string) (result autorest.Response, err error)
	ListByProduct(ctx context.Context, apimBaseURL string, productID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListByProductComplete(ctx context.Context, apimBaseURL string, productID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
}

var _ ProductGroupClientAPI = (*apimanagement.ProductGroupClient)(nil)

// ProductSubscriptionsClientAPI contains the set of methods on the ProductSubscriptionsClient type.
type ProductSubscriptionsClientAPI interface {
	List(ctx context.Context, apimBaseURL string, productID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, productID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
}

var _ ProductSubscriptionsClientAPI = (*apimanagement.ProductSubscriptionsClient)(nil)

// ProductPolicyClientAPI contains the set of methods on the ProductPolicyClient type.
type ProductPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, productID string, parameters apimanagement.PolicyContract) (result apimanagement.PolicyContract, err error)
	Delete(ctx context.Context, apimBaseURL string, productID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, productID string) (result apimanagement.PolicyContract, err error)
	ListByProduct(ctx context.Context, apimBaseURL string, productID string) (result apimanagement.PolicyCollection, err error)
}

var _ ProductPolicyClientAPI = (*apimanagement.ProductPolicyClient)(nil)

// PropertyClientAPI contains the set of methods on the PropertyClient type.
type PropertyClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, propID string, parameters apimanagement.PropertyContract) (result apimanagement.PropertyContract, err error)
	Delete(ctx context.Context, apimBaseURL string, propID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, propID string) (result apimanagement.PropertyContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.PropertyCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.PropertyCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, propID string, parameters apimanagement.PropertyUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ PropertyClientAPI = (*apimanagement.PropertyClient)(nil)

// QuotaByCounterKeysClientAPI contains the set of methods on the QuotaByCounterKeysClient type.
type QuotaByCounterKeysClientAPI interface {
	List(ctx context.Context, apimBaseURL string, quotaCounterKey string) (result apimanagement.QuotaCounterCollection, err error)
	Update(ctx context.Context, apimBaseURL string, quotaCounterKey string, parameters apimanagement.QuotaCounterValueContractProperties) (result autorest.Response, err error)
}

var _ QuotaByCounterKeysClientAPI = (*apimanagement.QuotaByCounterKeysClient)(nil)

// QuotaByPeriodKeysClientAPI contains the set of methods on the QuotaByPeriodKeysClient type.
type QuotaByPeriodKeysClientAPI interface {
	Get(ctx context.Context, apimBaseURL string, quotaCounterKey string, quotaPeriodKey string) (result apimanagement.QuotaCounterContract, err error)
	Update(ctx context.Context, apimBaseURL string, quotaCounterKey string, quotaPeriodKey string, parameters apimanagement.QuotaCounterValueContractProperties) (result autorest.Response, err error)
}

var _ QuotaByPeriodKeysClientAPI = (*apimanagement.QuotaByPeriodKeysClient)(nil)

// ReportsClientAPI contains the set of methods on the ReportsClient type.
type ReportsClientAPI interface {
	ListByAPI(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListByAPIComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
	ListByGeo(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListByGeoComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
	ListByOperation(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListByOperationComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
	ListByProduct(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListByProductComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
	ListByRequest(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.RequestReportCollection, err error)
	ListBySubscription(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListBySubscriptionComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
	ListByTime(ctx context.Context, apimBaseURL string, interval string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListByTimeComplete(ctx context.Context, apimBaseURL string, interval string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
	ListByUser(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionPage, err error)
	ListByUserComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.ReportCollectionIterator, err error)
}

var _ ReportsClientAPI = (*apimanagement.ReportsClient)(nil)

// SubscriptionClientAPI contains the set of methods on the SubscriptionClient type.
type SubscriptionClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, sid string, parameters apimanagement.SubscriptionCreateParameters, notify string) (result apimanagement.SubscriptionContract, err error)
	Delete(ctx context.Context, apimBaseURL string, sid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, apimBaseURL string, sid string) (result apimanagement.SubscriptionContract, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
	RegeneratePrimaryKey(ctx context.Context, apimBaseURL string, sid string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, apimBaseURL string, sid string) (result autorest.Response, err error)
	Update(ctx context.Context, apimBaseURL string, sid string, parameters apimanagement.SubscriptionUpdateParameters, ifMatch string, notify string) (result autorest.Response, err error)
}

var _ SubscriptionClientAPI = (*apimanagement.SubscriptionClient)(nil)

// TenantAccessClientAPI contains the set of methods on the TenantAccessClient type.
type TenantAccessClientAPI interface {
	Get(ctx context.Context, apimBaseURL string) (result apimanagement.AccessInformationContract, err error)
	RegeneratePrimaryKey(ctx context.Context, apimBaseURL string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, apimBaseURL string) (result autorest.Response, err error)
	Update(ctx context.Context, apimBaseURL string, parameters apimanagement.AccessInformationUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ TenantAccessClientAPI = (*apimanagement.TenantAccessClient)(nil)

// TenantAccessGitClientAPI contains the set of methods on the TenantAccessGitClient type.
type TenantAccessGitClientAPI interface {
	Get(ctx context.Context, apimBaseURL string) (result apimanagement.AccessInformationContract, err error)
	RegeneratePrimaryKey(ctx context.Context, apimBaseURL string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, apimBaseURL string) (result autorest.Response, err error)
}

var _ TenantAccessGitClientAPI = (*apimanagement.TenantAccessGitClient)(nil)

// TenantConfigurationClientAPI contains the set of methods on the TenantConfigurationClient type.
type TenantConfigurationClientAPI interface {
	Deploy(ctx context.Context, apimBaseURL string, parameters apimanagement.DeployConfigurationParameters) (result apimanagement.TenantConfigurationDeployFuture, err error)
	GetSyncState(ctx context.Context, apimBaseURL string) (result apimanagement.TenantConfigurationSyncStateContract, err error)
	Save(ctx context.Context, apimBaseURL string, parameters apimanagement.SaveConfigurationParameter) (result apimanagement.TenantConfigurationSaveFuture, err error)
	Validate(ctx context.Context, apimBaseURL string, parameters apimanagement.DeployConfigurationParameters) (result apimanagement.TenantConfigurationValidateFuture, err error)
}

var _ TenantConfigurationClientAPI = (*apimanagement.TenantConfigurationClient)(nil)

// UserClientAPI contains the set of methods on the UserClient type.
type UserClientAPI interface {
	CreateOrUpdate(ctx context.Context, apimBaseURL string, UID string, parameters apimanagement.UserCreateParameters) (result apimanagement.UserContract, err error)
	Delete(ctx context.Context, apimBaseURL string, UID string, ifMatch string, deleteSubscriptions string, notify string) (result autorest.Response, err error)
	GenerateSsoURL(ctx context.Context, apimBaseURL string, UID string) (result apimanagement.GenerateSsoURLResult, err error)
	Get(ctx context.Context, apimBaseURL string, UID string) (result apimanagement.UserContract, err error)
	GetSharedAccessToken(ctx context.Context, apimBaseURL string, UID string, parameters apimanagement.UserTokenParameters) (result apimanagement.UserTokenResult, err error)
	List(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionIterator, err error)
	Update(ctx context.Context, apimBaseURL string, UID string, parameters apimanagement.UserUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ UserClientAPI = (*apimanagement.UserClient)(nil)

// UserGroupClientAPI contains the set of methods on the UserGroupClient type.
type UserGroupClientAPI interface {
	List(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
}

var _ UserGroupClientAPI = (*apimanagement.UserGroupClient)(nil)

// UserSubscriptionClientAPI contains the set of methods on the UserSubscriptionClient type.
type UserSubscriptionClientAPI interface {
	List(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListComplete(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
}

var _ UserSubscriptionClientAPI = (*apimanagement.UserSubscriptionClient)(nil)

// UserIdentitiesClientAPI contains the set of methods on the UserIdentitiesClient type.
type UserIdentitiesClientAPI interface {
	List(ctx context.Context, apimBaseURL string, UID string) (result apimanagement.UserIdentityCollection, err error)
}

var _ UserIdentitiesClientAPI = (*apimanagement.UserIdentitiesClient)(nil)

// APIExportClientAPI contains the set of methods on the APIExportClient type.
type APIExportClientAPI interface {
	Get(ctx context.Context, apimBaseURL string, apiid string) (result apimanagement.APIExportResult, err error)
}

var _ APIExportClientAPI = (*apimanagement.APIExportClient)(nil)
