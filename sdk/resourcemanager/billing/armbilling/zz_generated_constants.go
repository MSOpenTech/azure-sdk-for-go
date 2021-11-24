//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

const (
	module  = "armbilling"
	version = "v0.1.0"
)

// AcceptanceMode - The mode of acceptance for an agreement.
type AcceptanceMode string

const (
	AcceptanceModeClickToAccept AcceptanceMode = "ClickToAccept"
	AcceptanceModeESignEmbedded AcceptanceMode = "ESignEmbedded"
	AcceptanceModeESignOffline  AcceptanceMode = "ESignOffline"
)

// PossibleAcceptanceModeValues returns the possible values for the AcceptanceMode const type.
func PossibleAcceptanceModeValues() []AcceptanceMode {
	return []AcceptanceMode{
		AcceptanceModeClickToAccept,
		AcceptanceModeESignEmbedded,
		AcceptanceModeESignOffline,
	}
}

// ToPtr returns a *AcceptanceMode pointing to the current value.
func (c AcceptanceMode) ToPtr() *AcceptanceMode {
	return &c
}

// AccountStatus - The current status of the billing account.
type AccountStatus string

const (
	AccountStatusActive      AccountStatus = "Active"
	AccountStatusDeleted     AccountStatus = "Deleted"
	AccountStatusDisabled    AccountStatus = "Disabled"
	AccountStatusExpired     AccountStatus = "Expired"
	AccountStatusExtended    AccountStatus = "Extended"
	AccountStatusTerminated  AccountStatus = "Terminated"
	AccountStatusTransferred AccountStatus = "Transferred"
)

// PossibleAccountStatusValues returns the possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{
		AccountStatusActive,
		AccountStatusDeleted,
		AccountStatusDisabled,
		AccountStatusExpired,
		AccountStatusExtended,
		AccountStatusTerminated,
		AccountStatusTransferred,
	}
}

// ToPtr returns a *AccountStatus pointing to the current value.
func (c AccountStatus) ToPtr() *AccountStatus {
	return &c
}

// AccountType - The type of customer.
type AccountType string

const (
	AccountTypeEnterprise AccountType = "Enterprise"
	AccountTypeIndividual AccountType = "Individual"
	AccountTypePartner    AccountType = "Partner"
)

// PossibleAccountTypeValues returns the possible values for the AccountType const type.
func PossibleAccountTypeValues() []AccountType {
	return []AccountType{
		AccountTypeEnterprise,
		AccountTypeIndividual,
		AccountTypePartner,
	}
}

// ToPtr returns a *AccountType pointing to the current value.
func (c AccountType) ToPtr() *AccountType {
	return &c
}

// AddressValidationStatus - Status of the address validation.
type AddressValidationStatus string

const (
	AddressValidationStatusInvalid AddressValidationStatus = "Invalid"
	AddressValidationStatusValid   AddressValidationStatus = "Valid"
)

// PossibleAddressValidationStatusValues returns the possible values for the AddressValidationStatus const type.
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return []AddressValidationStatus{
		AddressValidationStatusInvalid,
		AddressValidationStatusValid,
	}
}

// ToPtr returns a *AddressValidationStatus pointing to the current value.
func (c AddressValidationStatus) ToPtr() *AddressValidationStatus {
	return &c
}

// AgreementType - The type of agreement.
type AgreementType string

const (
	AgreementTypeEnterpriseAgreement            AgreementType = "EnterpriseAgreement"
	AgreementTypeMicrosoftCustomerAgreement     AgreementType = "MicrosoftCustomerAgreement"
	AgreementTypeMicrosoftOnlineServicesProgram AgreementType = "MicrosoftOnlineServicesProgram"
	AgreementTypeMicrosoftPartnerAgreement      AgreementType = "MicrosoftPartnerAgreement"
)

// PossibleAgreementTypeValues returns the possible values for the AgreementType const type.
func PossibleAgreementTypeValues() []AgreementType {
	return []AgreementType{
		AgreementTypeEnterpriseAgreement,
		AgreementTypeMicrosoftCustomerAgreement,
		AgreementTypeMicrosoftOnlineServicesProgram,
		AgreementTypeMicrosoftPartnerAgreement,
	}
}

// ToPtr returns a *AgreementType pointing to the current value.
func (c AgreementType) ToPtr() *AgreementType {
	return &c
}

// AutoRenew - Indicates whether auto renewal is turned on or off for a product.
type AutoRenew string

const (
	AutoRenewOff AutoRenew = "Off"
	AutoRenewOn  AutoRenew = "On"
)

// PossibleAutoRenewValues returns the possible values for the AutoRenew const type.
func PossibleAutoRenewValues() []AutoRenew {
	return []AutoRenew{
		AutoRenewOff,
		AutoRenewOn,
	}
}

// ToPtr returns a *AutoRenew pointing to the current value.
func (c AutoRenew) ToPtr() *AutoRenew {
	return &c
}

// BillingFrequency - The frequency at which the product will be billed.
type BillingFrequency string

const (
	BillingFrequencyMonthly    BillingFrequency = "Monthly"
	BillingFrequencyOneTime    BillingFrequency = "OneTime"
	BillingFrequencyUsageBased BillingFrequency = "UsageBased"
)

// PossibleBillingFrequencyValues returns the possible values for the BillingFrequency const type.
func PossibleBillingFrequencyValues() []BillingFrequency {
	return []BillingFrequency{
		BillingFrequencyMonthly,
		BillingFrequencyOneTime,
		BillingFrequencyUsageBased,
	}
}

// ToPtr returns a *BillingFrequency pointing to the current value.
func (c BillingFrequency) ToPtr() *BillingFrequency {
	return &c
}

// BillingProfileSpendingLimit - The billing profile spending limit.
type BillingProfileSpendingLimit string

const (
	BillingProfileSpendingLimitOff BillingProfileSpendingLimit = "Off"
	BillingProfileSpendingLimitOn  BillingProfileSpendingLimit = "On"
)

// PossibleBillingProfileSpendingLimitValues returns the possible values for the BillingProfileSpendingLimit const type.
func PossibleBillingProfileSpendingLimitValues() []BillingProfileSpendingLimit {
	return []BillingProfileSpendingLimit{
		BillingProfileSpendingLimitOff,
		BillingProfileSpendingLimitOn,
	}
}

// ToPtr returns a *BillingProfileSpendingLimit pointing to the current value.
func (c BillingProfileSpendingLimit) ToPtr() *BillingProfileSpendingLimit {
	return &c
}

// BillingProfileStatus - The status of the billing profile.
type BillingProfileStatus string

const (
	BillingProfileStatusActive   BillingProfileStatus = "Active"
	BillingProfileStatusDisabled BillingProfileStatus = "Disabled"
	BillingProfileStatusWarned   BillingProfileStatus = "Warned"
)

// PossibleBillingProfileStatusValues returns the possible values for the BillingProfileStatus const type.
func PossibleBillingProfileStatusValues() []BillingProfileStatus {
	return []BillingProfileStatus{
		BillingProfileStatusActive,
		BillingProfileStatusDisabled,
		BillingProfileStatusWarned,
	}
}

// ToPtr returns a *BillingProfileStatus pointing to the current value.
func (c BillingProfileStatus) ToPtr() *BillingProfileStatus {
	return &c
}

// BillingProfileStatusReasonCode - Reason for the specified billing profile status.
type BillingProfileStatusReasonCode string

const (
	BillingProfileStatusReasonCodePastDue              BillingProfileStatusReasonCode = "PastDue"
	BillingProfileStatusReasonCodeSpendingLimitExpired BillingProfileStatusReasonCode = "SpendingLimitExpired"
	BillingProfileStatusReasonCodeSpendingLimitReached BillingProfileStatusReasonCode = "SpendingLimitReached"
)

// PossibleBillingProfileStatusReasonCodeValues returns the possible values for the BillingProfileStatusReasonCode const type.
func PossibleBillingProfileStatusReasonCodeValues() []BillingProfileStatusReasonCode {
	return []BillingProfileStatusReasonCode{
		BillingProfileStatusReasonCodePastDue,
		BillingProfileStatusReasonCodeSpendingLimitExpired,
		BillingProfileStatusReasonCodeSpendingLimitReached,
	}
}

// ToPtr returns a *BillingProfileStatusReasonCode pointing to the current value.
func (c BillingProfileStatusReasonCode) ToPtr() *BillingProfileStatusReasonCode {
	return &c
}

// BillingRelationshipType - Identifies which services and purchases are paid by a billing profile.
type BillingRelationshipType string

const (
	BillingRelationshipTypeCSPPartner       BillingRelationshipType = "CSPPartner"
	BillingRelationshipTypeDirect           BillingRelationshipType = "Direct"
	BillingRelationshipTypeIndirectCustomer BillingRelationshipType = "IndirectCustomer"
	BillingRelationshipTypeIndirectPartner  BillingRelationshipType = "IndirectPartner"
)

// PossibleBillingRelationshipTypeValues returns the possible values for the BillingRelationshipType const type.
func PossibleBillingRelationshipTypeValues() []BillingRelationshipType {
	return []BillingRelationshipType{
		BillingRelationshipTypeCSPPartner,
		BillingRelationshipTypeDirect,
		BillingRelationshipTypeIndirectCustomer,
		BillingRelationshipTypeIndirectPartner,
	}
}

// ToPtr returns a *BillingRelationshipType pointing to the current value.
func (c BillingRelationshipType) ToPtr() *BillingRelationshipType {
	return &c
}

// BillingSubscriptionStatusType - The current billing status of the subscription.
type BillingSubscriptionStatusType string

const (
	BillingSubscriptionStatusTypeAbandoned BillingSubscriptionStatusType = "Abandoned"
	BillingSubscriptionStatusTypeActive    BillingSubscriptionStatusType = "Active"
	BillingSubscriptionStatusTypeDeleted   BillingSubscriptionStatusType = "Deleted"
	BillingSubscriptionStatusTypeInactive  BillingSubscriptionStatusType = "Inactive"
	BillingSubscriptionStatusTypeWarning   BillingSubscriptionStatusType = "Warning"
)

// PossibleBillingSubscriptionStatusTypeValues returns the possible values for the BillingSubscriptionStatusType const type.
func PossibleBillingSubscriptionStatusTypeValues() []BillingSubscriptionStatusType {
	return []BillingSubscriptionStatusType{
		BillingSubscriptionStatusTypeAbandoned,
		BillingSubscriptionStatusTypeActive,
		BillingSubscriptionStatusTypeDeleted,
		BillingSubscriptionStatusTypeInactive,
		BillingSubscriptionStatusTypeWarning,
	}
}

// ToPtr returns a *BillingSubscriptionStatusType pointing to the current value.
func (c BillingSubscriptionStatusType) ToPtr() *BillingSubscriptionStatusType {
	return &c
}

// Category - The category of the agreement signed by a customer.
type Category string

const (
	CategoryAffiliatePurchaseTerms     Category = "AffiliatePurchaseTerms"
	CategoryMicrosoftCustomerAgreement Category = "MicrosoftCustomerAgreement"
	CategoryOther                      Category = "Other"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryAffiliatePurchaseTerms,
		CategoryMicrosoftCustomerAgreement,
		CategoryOther,
	}
}

// ToPtr returns a *Category pointing to the current value.
func (c Category) ToPtr() *Category {
	return &c
}

// DocumentSource - The source of the document. ENF for Brazil and DRS for rest of the world.
type DocumentSource string

const (
	DocumentSourceDRS DocumentSource = "DRS"
	DocumentSourceENF DocumentSource = "ENF"
)

// PossibleDocumentSourceValues returns the possible values for the DocumentSource const type.
func PossibleDocumentSourceValues() []DocumentSource {
	return []DocumentSource{
		DocumentSourceDRS,
		DocumentSourceENF,
	}
}

// ToPtr returns a *DocumentSource pointing to the current value.
func (c DocumentSource) ToPtr() *DocumentSource {
	return &c
}

// DocumentType - The type of the document.
type DocumentType string

const (
	DocumentTypeCreditNote DocumentType = "CreditNote"
	DocumentTypeInvoice    DocumentType = "Invoice"
	DocumentTypeTaxReceipt DocumentType = "TaxReceipt"
	DocumentTypeVoidNote   DocumentType = "VoidNote"
)

// PossibleDocumentTypeValues returns the possible values for the DocumentType const type.
func PossibleDocumentTypeValues() []DocumentType {
	return []DocumentType{
		DocumentTypeCreditNote,
		DocumentTypeInvoice,
		DocumentTypeTaxReceipt,
		DocumentTypeVoidNote,
	}
}

// ToPtr returns a *DocumentType pointing to the current value.
func (c DocumentType) ToPtr() *DocumentType {
	return &c
}

// InvoiceDocumentType - The type of the document.
type InvoiceDocumentType string

const (
	InvoiceDocumentTypeCreditNote InvoiceDocumentType = "CreditNote"
	InvoiceDocumentTypeInvoice    InvoiceDocumentType = "Invoice"
)

// PossibleInvoiceDocumentTypeValues returns the possible values for the InvoiceDocumentType const type.
func PossibleInvoiceDocumentTypeValues() []InvoiceDocumentType {
	return []InvoiceDocumentType{
		InvoiceDocumentTypeCreditNote,
		InvoiceDocumentTypeInvoice,
	}
}

// ToPtr returns a *InvoiceDocumentType pointing to the current value.
func (c InvoiceDocumentType) ToPtr() *InvoiceDocumentType {
	return &c
}

// InvoiceSectionState - Identifies the state of an invoice section.
type InvoiceSectionState string

const (
	InvoiceSectionStateActive     InvoiceSectionState = "Active"
	InvoiceSectionStateRestricted InvoiceSectionState = "Restricted"
)

// PossibleInvoiceSectionStateValues returns the possible values for the InvoiceSectionState const type.
func PossibleInvoiceSectionStateValues() []InvoiceSectionState {
	return []InvoiceSectionState{
		InvoiceSectionStateActive,
		InvoiceSectionStateRestricted,
	}
}

// ToPtr returns a *InvoiceSectionState pointing to the current value.
func (c InvoiceSectionState) ToPtr() *InvoiceSectionState {
	return &c
}

// InvoiceStatus - The current status of the invoice.
type InvoiceStatus string

const (
	InvoiceStatusDue     InvoiceStatus = "Due"
	InvoiceStatusOverDue InvoiceStatus = "OverDue"
	InvoiceStatusPaid    InvoiceStatus = "Paid"
	InvoiceStatusVoid    InvoiceStatus = "Void"
)

// PossibleInvoiceStatusValues returns the possible values for the InvoiceStatus const type.
func PossibleInvoiceStatusValues() []InvoiceStatus {
	return []InvoiceStatus{
		InvoiceStatusDue,
		InvoiceStatusOverDue,
		InvoiceStatusPaid,
		InvoiceStatusVoid,
	}
}

// ToPtr returns a *InvoiceStatus pointing to the current value.
func (c InvoiceStatus) ToPtr() *InvoiceStatus {
	return &c
}

// InvoiceType - Invoice type.
type InvoiceType string

const (
	InvoiceTypeAzureMarketplace InvoiceType = "AzureMarketplace"
	InvoiceTypeAzureService     InvoiceType = "AzureService"
	InvoiceTypeAzureSupport     InvoiceType = "AzureSupport"
)

// PossibleInvoiceTypeValues returns the possible values for the InvoiceType const type.
func PossibleInvoiceTypeValues() []InvoiceType {
	return []InvoiceType{
		InvoiceTypeAzureMarketplace,
		InvoiceTypeAzureService,
		InvoiceTypeAzureSupport,
	}
}

// ToPtr returns a *InvoiceType pointing to the current value.
func (c InvoiceType) ToPtr() *InvoiceType {
	return &c
}

// MarketplacePurchasesPolicy - The policy that controls whether Azure marketplace purchases are allowed for a billing profile.
type MarketplacePurchasesPolicy string

const (
	MarketplacePurchasesPolicyAllAllowed      MarketplacePurchasesPolicy = "AllAllowed"
	MarketplacePurchasesPolicyNotAllowed      MarketplacePurchasesPolicy = "NotAllowed"
	MarketplacePurchasesPolicyOnlyFreeAllowed MarketplacePurchasesPolicy = "OnlyFreeAllowed"
)

// PossibleMarketplacePurchasesPolicyValues returns the possible values for the MarketplacePurchasesPolicy const type.
func PossibleMarketplacePurchasesPolicyValues() []MarketplacePurchasesPolicy {
	return []MarketplacePurchasesPolicy{
		MarketplacePurchasesPolicyAllAllowed,
		MarketplacePurchasesPolicyNotAllowed,
		MarketplacePurchasesPolicyOnlyFreeAllowed,
	}
}

// ToPtr returns a *MarketplacePurchasesPolicy pointing to the current value.
func (c MarketplacePurchasesPolicy) ToPtr() *MarketplacePurchasesPolicy {
	return &c
}

// PaymentMethodFamily - The family of payment method.
type PaymentMethodFamily string

const (
	PaymentMethodFamilyCheckWire  PaymentMethodFamily = "CheckWire"
	PaymentMethodFamilyCreditCard PaymentMethodFamily = "CreditCard"
	PaymentMethodFamilyCredits    PaymentMethodFamily = "Credits"
	PaymentMethodFamilyNone       PaymentMethodFamily = "None"
)

// PossiblePaymentMethodFamilyValues returns the possible values for the PaymentMethodFamily const type.
func PossiblePaymentMethodFamilyValues() []PaymentMethodFamily {
	return []PaymentMethodFamily{
		PaymentMethodFamilyCheckWire,
		PaymentMethodFamilyCreditCard,
		PaymentMethodFamilyCredits,
		PaymentMethodFamilyNone,
	}
}

// ToPtr returns a *PaymentMethodFamily pointing to the current value.
func (c PaymentMethodFamily) ToPtr() *PaymentMethodFamily {
	return &c
}

// ProductStatusType - The current status of the product.
type ProductStatusType string

const (
	ProductStatusTypeActive    ProductStatusType = "Active"
	ProductStatusTypeAutoRenew ProductStatusType = "AutoRenew"
	ProductStatusTypeCancelled ProductStatusType = "Cancelled"
	ProductStatusTypeDisabled  ProductStatusType = "Disabled"
	ProductStatusTypeExpired   ProductStatusType = "Expired"
	ProductStatusTypeExpiring  ProductStatusType = "Expiring"
	ProductStatusTypeInactive  ProductStatusType = "Inactive"
	ProductStatusTypePastDue   ProductStatusType = "PastDue"
)

// PossibleProductStatusTypeValues returns the possible values for the ProductStatusType const type.
func PossibleProductStatusTypeValues() []ProductStatusType {
	return []ProductStatusType{
		ProductStatusTypeActive,
		ProductStatusTypeAutoRenew,
		ProductStatusTypeCancelled,
		ProductStatusTypeDisabled,
		ProductStatusTypeExpired,
		ProductStatusTypeExpiring,
		ProductStatusTypeInactive,
		ProductStatusTypePastDue,
	}
}

// ToPtr returns a *ProductStatusType pointing to the current value.
func (c ProductStatusType) ToPtr() *ProductStatusType {
	return &c
}

// ProductTransferValidationErrorCode - Error code of the transfer validation response.
type ProductTransferValidationErrorCode string

const (
	ProductTransferValidationErrorCodeCrossBillingAccountNotAllowed            ProductTransferValidationErrorCode = "CrossBillingAccountNotAllowed"
	ProductTransferValidationErrorCodeDestinationBillingProfilePastDue         ProductTransferValidationErrorCode = "DestinationBillingProfilePastDue"
	ProductTransferValidationErrorCodeInsufficientPermissionOnDestination      ProductTransferValidationErrorCode = "InsufficientPermissionOnDestination"
	ProductTransferValidationErrorCodeInsufficientPermissionOnSource           ProductTransferValidationErrorCode = "InsufficientPermissionOnSource"
	ProductTransferValidationErrorCodeInvalidSource                            ProductTransferValidationErrorCode = "InvalidSource"
	ProductTransferValidationErrorCodeNotAvailableForDestinationMarket         ProductTransferValidationErrorCode = "NotAvailableForDestinationMarket"
	ProductTransferValidationErrorCodeOneTimePurchaseProductTransferNotAllowed ProductTransferValidationErrorCode = "OneTimePurchaseProductTransferNotAllowed"
	ProductTransferValidationErrorCodeProductNotActive                         ProductTransferValidationErrorCode = "ProductNotActive"
	ProductTransferValidationErrorCodeProductTypeNotSupported                  ProductTransferValidationErrorCode = "ProductTypeNotSupported"
)

// PossibleProductTransferValidationErrorCodeValues returns the possible values for the ProductTransferValidationErrorCode const type.
func PossibleProductTransferValidationErrorCodeValues() []ProductTransferValidationErrorCode {
	return []ProductTransferValidationErrorCode{
		ProductTransferValidationErrorCodeCrossBillingAccountNotAllowed,
		ProductTransferValidationErrorCodeDestinationBillingProfilePastDue,
		ProductTransferValidationErrorCodeInsufficientPermissionOnDestination,
		ProductTransferValidationErrorCodeInsufficientPermissionOnSource,
		ProductTransferValidationErrorCodeInvalidSource,
		ProductTransferValidationErrorCodeNotAvailableForDestinationMarket,
		ProductTransferValidationErrorCodeOneTimePurchaseProductTransferNotAllowed,
		ProductTransferValidationErrorCodeProductNotActive,
		ProductTransferValidationErrorCodeProductTypeNotSupported,
	}
}

// ToPtr returns a *ProductTransferValidationErrorCode pointing to the current value.
func (c ProductTransferValidationErrorCode) ToPtr() *ProductTransferValidationErrorCode {
	return &c
}

// ReservationPurchasesPolicy - The policy that controls whether Azure reservation purchases are allowed for a billing profile.
type ReservationPurchasesPolicy string

const (
	ReservationPurchasesPolicyAllowed    ReservationPurchasesPolicy = "Allowed"
	ReservationPurchasesPolicyNotAllowed ReservationPurchasesPolicy = "NotAllowed"
)

// PossibleReservationPurchasesPolicyValues returns the possible values for the ReservationPurchasesPolicy const type.
func PossibleReservationPurchasesPolicyValues() []ReservationPurchasesPolicy {
	return []ReservationPurchasesPolicy{
		ReservationPurchasesPolicyAllowed,
		ReservationPurchasesPolicyNotAllowed,
	}
}

// ToPtr returns a *ReservationPurchasesPolicy pointing to the current value.
func (c ReservationPurchasesPolicy) ToPtr() *ReservationPurchasesPolicy {
	return &c
}

// ReservationType - The type of transaction.
type ReservationType string

const (
	ReservationTypePurchase    ReservationType = "Purchase"
	ReservationTypeUsageCharge ReservationType = "Usage Charge"
)

// PossibleReservationTypeValues returns the possible values for the ReservationType const type.
func PossibleReservationTypeValues() []ReservationType {
	return []ReservationType{
		ReservationTypePurchase,
		ReservationTypeUsageCharge,
	}
}

// ToPtr returns a *ReservationType pointing to the current value.
func (c ReservationType) ToPtr() *ReservationType {
	return &c
}

// SpendingLimit - The billing profile spending limit.
type SpendingLimit string

const (
	SpendingLimitOff SpendingLimit = "Off"
	SpendingLimitOn  SpendingLimit = "On"
)

// PossibleSpendingLimitValues returns the possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{
		SpendingLimitOff,
		SpendingLimitOn,
	}
}

// ToPtr returns a *SpendingLimit pointing to the current value.
func (c SpendingLimit) ToPtr() *SpendingLimit {
	return &c
}

// SpendingLimitForBillingProfile - The billing profile spending limit.
type SpendingLimitForBillingProfile string

const (
	SpendingLimitForBillingProfileOff SpendingLimitForBillingProfile = "Off"
	SpendingLimitForBillingProfileOn  SpendingLimitForBillingProfile = "On"
)

// PossibleSpendingLimitForBillingProfileValues returns the possible values for the SpendingLimitForBillingProfile const type.
func PossibleSpendingLimitForBillingProfileValues() []SpendingLimitForBillingProfile {
	return []SpendingLimitForBillingProfile{
		SpendingLimitForBillingProfileOff,
		SpendingLimitForBillingProfileOn,
	}
}

// ToPtr returns a *SpendingLimitForBillingProfile pointing to the current value.
func (c SpendingLimitForBillingProfile) ToPtr() *SpendingLimitForBillingProfile {
	return &c
}

// StatusReasonCode - Reason for the specified billing profile status.
type StatusReasonCode string

const (
	StatusReasonCodePastDue              StatusReasonCode = "PastDue"
	StatusReasonCodeSpendingLimitExpired StatusReasonCode = "SpendingLimitExpired"
	StatusReasonCodeSpendingLimitReached StatusReasonCode = "SpendingLimitReached"
)

// PossibleStatusReasonCodeValues returns the possible values for the StatusReasonCode const type.
func PossibleStatusReasonCodeValues() []StatusReasonCode {
	return []StatusReasonCode{
		StatusReasonCodePastDue,
		StatusReasonCodeSpendingLimitExpired,
		StatusReasonCodeSpendingLimitReached,
	}
}

// ToPtr returns a *StatusReasonCode pointing to the current value.
func (c StatusReasonCode) ToPtr() *StatusReasonCode {
	return &c
}

// StatusReasonCodeForBillingProfile - Reason for the specified billing profile status.
type StatusReasonCodeForBillingProfile string

const (
	StatusReasonCodeForBillingProfilePastDue              StatusReasonCodeForBillingProfile = "PastDue"
	StatusReasonCodeForBillingProfileSpendingLimitExpired StatusReasonCodeForBillingProfile = "SpendingLimitExpired"
	StatusReasonCodeForBillingProfileSpendingLimitReached StatusReasonCodeForBillingProfile = "SpendingLimitReached"
)

// PossibleStatusReasonCodeForBillingProfileValues returns the possible values for the StatusReasonCodeForBillingProfile const type.
func PossibleStatusReasonCodeForBillingProfileValues() []StatusReasonCodeForBillingProfile {
	return []StatusReasonCodeForBillingProfile{
		StatusReasonCodeForBillingProfilePastDue,
		StatusReasonCodeForBillingProfileSpendingLimitExpired,
		StatusReasonCodeForBillingProfileSpendingLimitReached,
	}
}

// ToPtr returns a *StatusReasonCodeForBillingProfile pointing to the current value.
func (c StatusReasonCodeForBillingProfile) ToPtr() *StatusReasonCodeForBillingProfile {
	return &c
}

// SubscriptionTransferValidationErrorCode - Error code of the transfer validation response.
type SubscriptionTransferValidationErrorCode string

const (
	SubscriptionTransferValidationErrorCodeBillingAccountInactive              SubscriptionTransferValidationErrorCode = "BillingAccountInactive"
	SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed       SubscriptionTransferValidationErrorCode = "CrossBillingAccountNotAllowed"
	SubscriptionTransferValidationErrorCodeDestinationBillingProfileInactive   SubscriptionTransferValidationErrorCode = "DestinationBillingProfileInactive"
	SubscriptionTransferValidationErrorCodeDestinationBillingProfileNotFound   SubscriptionTransferValidationErrorCode = "DestinationBillingProfileNotFound"
	SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue    SubscriptionTransferValidationErrorCode = "DestinationBillingProfilePastDue"
	SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionInactive   SubscriptionTransferValidationErrorCode = "DestinationInvoiceSectionInactive"
	SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionNotFound   SubscriptionTransferValidationErrorCode = "DestinationInvoiceSectionNotFound"
	SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination SubscriptionTransferValidationErrorCode = "InsufficientPermissionOnDestination"
	SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource      SubscriptionTransferValidationErrorCode = "InsufficientPermissionOnSource"
	SubscriptionTransferValidationErrorCodeInvalidDestination                  SubscriptionTransferValidationErrorCode = "InvalidDestination"
	SubscriptionTransferValidationErrorCodeInvalidSource                       SubscriptionTransferValidationErrorCode = "InvalidSource"
	SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination  SubscriptionTransferValidationErrorCode = "MarketplaceNotEnabledOnDestination"
	SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket    SubscriptionTransferValidationErrorCode = "NotAvailableForDestinationMarket"
	SubscriptionTransferValidationErrorCodeProductInactive                     SubscriptionTransferValidationErrorCode = "ProductInactive"
	SubscriptionTransferValidationErrorCodeProductNotFound                     SubscriptionTransferValidationErrorCode = "ProductNotFound"
	SubscriptionTransferValidationErrorCodeProductTypeNotSupported             SubscriptionTransferValidationErrorCode = "ProductTypeNotSupported"
	SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue         SubscriptionTransferValidationErrorCode = "SourceBillingProfilePastDue"
	SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive        SubscriptionTransferValidationErrorCode = "SourceInvoiceSectionInactive"
	SubscriptionTransferValidationErrorCodeSubscriptionNotActive               SubscriptionTransferValidationErrorCode = "SubscriptionNotActive"
	SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported        SubscriptionTransferValidationErrorCode = "SubscriptionTypeNotSupported"
)

// PossibleSubscriptionTransferValidationErrorCodeValues returns the possible values for the SubscriptionTransferValidationErrorCode const type.
func PossibleSubscriptionTransferValidationErrorCodeValues() []SubscriptionTransferValidationErrorCode {
	return []SubscriptionTransferValidationErrorCode{
		SubscriptionTransferValidationErrorCodeBillingAccountInactive,
		SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed,
		SubscriptionTransferValidationErrorCodeDestinationBillingProfileInactive,
		SubscriptionTransferValidationErrorCodeDestinationBillingProfileNotFound,
		SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue,
		SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionInactive,
		SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionNotFound,
		SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination,
		SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource,
		SubscriptionTransferValidationErrorCodeInvalidDestination,
		SubscriptionTransferValidationErrorCodeInvalidSource,
		SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination,
		SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket,
		SubscriptionTransferValidationErrorCodeProductInactive,
		SubscriptionTransferValidationErrorCodeProductNotFound,
		SubscriptionTransferValidationErrorCodeProductTypeNotSupported,
		SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue,
		SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive,
		SubscriptionTransferValidationErrorCodeSubscriptionNotActive,
		SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported,
	}
}

// ToPtr returns a *SubscriptionTransferValidationErrorCode pointing to the current value.
func (c SubscriptionTransferValidationErrorCode) ToPtr() *SubscriptionTransferValidationErrorCode {
	return &c
}

// TargetCloud - Possible cloud environments.
type TargetCloud string

const (
	TargetCloudUSGov TargetCloud = "USGov"
	TargetCloudUSNat TargetCloud = "USNat"
	TargetCloudUSSec TargetCloud = "USSec"
)

// PossibleTargetCloudValues returns the possible values for the TargetCloud const type.
func PossibleTargetCloudValues() []TargetCloud {
	return []TargetCloud{
		TargetCloudUSGov,
		TargetCloudUSNat,
		TargetCloudUSSec,
	}
}

// ToPtr returns a *TargetCloud pointing to the current value.
func (c TargetCloud) ToPtr() *TargetCloud {
	return &c
}

// TransactionTypeKind - The kind of transaction. Options are all or reservation.
type TransactionTypeKind string

const (
	TransactionTypeKindAll         TransactionTypeKind = "all"
	TransactionTypeKindReservation TransactionTypeKind = "reservation"
)

// PossibleTransactionTypeKindValues returns the possible values for the TransactionTypeKind const type.
func PossibleTransactionTypeKindValues() []TransactionTypeKind {
	return []TransactionTypeKind{
		TransactionTypeKindAll,
		TransactionTypeKindReservation,
	}
}

// ToPtr returns a *TransactionTypeKind pointing to the current value.
func (c TransactionTypeKind) ToPtr() *TransactionTypeKind {
	return &c
}

// ViewCharges - The policy that controls whether the users in customer's organization can view charges at pay-as-you-go prices.
type ViewCharges string

const (
	ViewChargesAllowed    ViewCharges = "Allowed"
	ViewChargesNotAllowed ViewCharges = "NotAllowed"
)

// PossibleViewChargesValues returns the possible values for the ViewCharges const type.
func PossibleViewChargesValues() []ViewCharges {
	return []ViewCharges{
		ViewChargesAllowed,
		ViewChargesNotAllowed,
	}
}

// ToPtr returns a *ViewCharges pointing to the current value.
func (c ViewCharges) ToPtr() *ViewCharges {
	return &c
}

// ViewChargesPolicy - The policy that controls whether users with Azure RBAC access to a subscription can view its charges.
type ViewChargesPolicy string

const (
	ViewChargesPolicyAllowed    ViewChargesPolicy = "Allowed"
	ViewChargesPolicyNotAllowed ViewChargesPolicy = "NotAllowed"
)

// PossibleViewChargesPolicyValues returns the possible values for the ViewChargesPolicy const type.
func PossibleViewChargesPolicyValues() []ViewChargesPolicy {
	return []ViewChargesPolicy{
		ViewChargesPolicyAllowed,
		ViewChargesPolicyNotAllowed,
	}
}

// ToPtr returns a *ViewChargesPolicy pointing to the current value.
func (c ViewChargesPolicy) ToPtr() *ViewChargesPolicy {
	return &c
}
