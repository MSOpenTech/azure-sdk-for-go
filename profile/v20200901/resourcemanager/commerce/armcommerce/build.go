// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This file enables 'go generate' to regenerate this specific SDK
//go:generate pwsh ../../../../../eng/scripts/build.ps1 -skipBuild -cleanGenerated -format -generate ../profile/v20200901/resourcemanager/commerce/armcommerce
//go:generate pwsh -Command Remove-Item go.mod

package armcommerce
