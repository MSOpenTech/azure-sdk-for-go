//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azcertificatemanager

type CertificateFileFormat string

const (
	CertificateFileFormatApplicationPkixCert           CertificateFileFormat = "application/pkix-cert"
	CertificateFileFormatApplicationXPemFile           CertificateFileFormat = "application/x-pem-file"
	CertificateFileFormatApplicationXPkcs7Certificates CertificateFileFormat = "application/x-pkcs7-certificates"
)

// PossibleCertificateFileFormatValues returns the possible values for the CertificateFileFormat const type.
func PossibleCertificateFileFormatValues() []CertificateFileFormat {
	return []CertificateFileFormat{
		CertificateFileFormatApplicationPkixCert,
		CertificateFileFormatApplicationXPemFile,
		CertificateFileFormatApplicationXPkcs7Certificates,
	}
}

type CertificateFormat string

const (
	CertificateFormatPem   CertificateFormat = "pem"
	CertificateFormatPkcs7 CertificateFormat = "pkcs7"
)

// PossibleCertificateFormatValues returns the possible values for the CertificateFormat const type.
func PossibleCertificateFormatValues() []CertificateFormat {
	return []CertificateFormat{
		CertificateFormatPem,
		CertificateFormatPkcs7,
	}
}

type ChainFormat string

const (
	ChainFormatP7B ChainFormat = "p7b"
	ChainFormatPem ChainFormat = "pem"
)

// PossibleChainFormatValues returns the possible values for the ChainFormat const type.
func PossibleChainFormatValues() []ChainFormat {
	return []ChainFormat{
		ChainFormatP7B,
		ChainFormatPem,
	}
}

type ExtendedKeyUsage string

const (
	ExtendedKeyUsageClientAuth      ExtendedKeyUsage = "ClientAuth"
	ExtendedKeyUsageCodeSigning     ExtendedKeyUsage = "CodeSigning"
	ExtendedKeyUsageEmailProtection ExtendedKeyUsage = "EmailProtection"
	ExtendedKeyUsageIPSecEndSystem  ExtendedKeyUsage = "IpsecEndSystem"
	ExtendedKeyUsageIPSecTunnel     ExtendedKeyUsage = "IpsecTunnel"
	ExtendedKeyUsageIPSecUser       ExtendedKeyUsage = "IpsecUser"
	ExtendedKeyUsageMACAddress      ExtendedKeyUsage = "MACAddress"
	ExtendedKeyUsageOCSPSigning     ExtendedKeyUsage = "OCSPSigning"
	ExtendedKeyUsageServerAuth      ExtendedKeyUsage = "ServerAuth"
	ExtendedKeyUsageSmartcardLogon  ExtendedKeyUsage = "SmartcardLogon"
	ExtendedKeyUsageTimeStamping    ExtendedKeyUsage = "TimeStamping"
)

// PossibleExtendedKeyUsageValues returns the possible values for the ExtendedKeyUsage const type.
func PossibleExtendedKeyUsageValues() []ExtendedKeyUsage {
	return []ExtendedKeyUsage{
		ExtendedKeyUsageClientAuth,
		ExtendedKeyUsageCodeSigning,
		ExtendedKeyUsageEmailProtection,
		ExtendedKeyUsageIPSecEndSystem,
		ExtendedKeyUsageIPSecTunnel,
		ExtendedKeyUsageIPSecUser,
		ExtendedKeyUsageMACAddress,
		ExtendedKeyUsageOCSPSigning,
		ExtendedKeyUsageServerAuth,
		ExtendedKeyUsageSmartcardLogon,
		ExtendedKeyUsageTimeStamping,
	}
}

type Include string

const (
	IncludeLogs Include = "logs"
)

// PossibleIncludeValues returns the possible values for the Include const type.
func PossibleIncludeValues() []Include {
	return []Include{
		IncludeLogs,
	}
}

type KeyUsage string

const (
	KeyUsageCrlSign          KeyUsage = "CrlSign"
	KeyUsageDataEncipherment KeyUsage = "DataEncipherment"
	KeyUsageDecipherOnly     KeyUsage = "DecipherOnly"
	KeyUsageDigitalSignature KeyUsage = "DigitalSignature"
	KeyUsageEncipherOnly     KeyUsage = "EncipherOnly"
	KeyUsageKeyAgreement     KeyUsage = "KeyAgreement"
	KeyUsageKeyCertSign      KeyUsage = "KeyCertSign"
	KeyUsageKeyEncipherment  KeyUsage = "KeyEncipherment"
	KeyUsageNonRepudiation   KeyUsage = "NonRepudiation"
)

// PossibleKeyUsageValues returns the possible values for the KeyUsage const type.
func PossibleKeyUsageValues() []KeyUsage {
	return []KeyUsage{
		KeyUsageCrlSign,
		KeyUsageDataEncipherment,
		KeyUsageDecipherOnly,
		KeyUsageDigitalSignature,
		KeyUsageEncipherOnly,
		KeyUsageKeyAgreement,
		KeyUsageKeyCertSign,
		KeyUsageKeyEncipherment,
		KeyUsageNonRepudiation,
	}
}

type RevocationReason string

const (
	RevocationReasonAACompromise         RevocationReason = "aACompromise"
	RevocationReasonAffiliationChanged   RevocationReason = "affiliationChanged"
	RevocationReasonCACompromise         RevocationReason = "cACompromise"
	RevocationReasonCertificateHold      RevocationReason = "certificateHold"
	RevocationReasonCessationOfOperation RevocationReason = "cessationOfOperation"
	RevocationReasonKeyCompromise        RevocationReason = "keyCompromise"
	RevocationReasonPrivilegeWithdrawn   RevocationReason = "privilegeWithdrawn"
	RevocationReasonRemoveFromCRL        RevocationReason = "removeFromCRL"
	RevocationReasonSuperseded           RevocationReason = "superseded"
	RevocationReasonUnspecified          RevocationReason = "unspecified"
)

// PossibleRevocationReasonValues returns the possible values for the RevocationReason const type.
func PossibleRevocationReasonValues() []RevocationReason {
	return []RevocationReason{
		RevocationReasonAACompromise,
		RevocationReasonAffiliationChanged,
		RevocationReasonCACompromise,
		RevocationReasonCertificateHold,
		RevocationReasonCessationOfOperation,
		RevocationReasonKeyCompromise,
		RevocationReasonPrivilegeWithdrawn,
		RevocationReasonRemoveFromCRL,
		RevocationReasonSuperseded,
		RevocationReasonUnspecified,
	}
}
