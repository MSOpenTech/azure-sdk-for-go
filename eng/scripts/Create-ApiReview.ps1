Write-Host "$PSScriptRoot"
. (Join-Path $PSScriptRoot .. common scripts common.ps1)

<#
.DESCRIPTION
    Create .gosource APIVIew artifact for go
.PARAMETER ServiceDirectory
    Thee name of the ServiceDirectory
.PARAMETER ArtifactOutputDirectory
    Base output Directory path for the generated gosource artifacts
.PARAMETER ArtifactName
    Name of the group of artifacts to be created
#>
function Create-APIViewArtifact {
        Param(
        [Parameter(Mandatory=$True)]
        [string] $ServiceDirectory,
        [Parameter(Mandatory=$True)]
        [string] $ArtifactOutputDirectory,
        [Parameter(Mandatory=$True)]
        [string] $ArtifactName
    )

    $artifactList = @()

    $artifactsDirectoryPath = Join-Path -Path $ArtifactOutputDirectory $ArtifactName
    New-Item -ItemType Directory -Path $artifactsDirectoryPath -force

    foreach ($sdk in (Get-AllPackageInfoFromRepo $ServiceDirectory))
    {
        Write-Host "Creating API review artifact for $($sdk.Name)"
        New-Item -ItemType Directory -Path $artifactsDirectoryPath/$($sdk.Name) -force
        $fileName = Split-Path -Path $sdk.Name -Leaf
        Compress-Archive -Path $sdk.DirectoryPath -DestinationPath $artifactsDirectoryPath/$($sdk.Name)/$fileName -force
        Rename-Item $artifactsDirectoryPath/$($sdk.Name)/$fileName.zip -NewName "$fileName.gosource"

        $artifactList += [PSCustomObject]@{
            name = $sdk.Name
        }
    }
    return $artifactList
}

<#
.DESCRIPTION
    Create new automatic APIView from a CI run
.PARAMETER ServiceDirectory
    Thee name of the ServiceDirectory
.PARAMETER ArtifactOutputDirectory
    Base output Directory path for the generated gosource artifacts
.PARAMETER ArtifactName
    Name of the group of artifacts to be created
.PARAMETER ApiKey
    The APIview ApiKey
.PARAMETER SourceBranch
    SourceBranch
.PARAMETER DefaultBranch
    DefaultBranch
.PARAMETER ConfigFileDir
    Path to the ConfigFileDir as published in the pipeline
.PARAMETER RepoName
    The name of the repository
.PARAMETER BuildId
    The build Id of the pipeline run
.PARAMETER MarkPackageAsShipped
    Indicate weather to mark the package a s shipped
#>
function New-APIView-From-CI {
    Param(
        [Parameter(Mandatory=$True)]
        [string] $ServiceDirectory,
        [Parameter(Mandatory=$True)]
        [string] $ArtifactOutputDirectory,
        [string] $ArtifactName="APIViewArtifacts",
        [Parameter(Mandatory=$True)]
        [string] $ApiKey,
        [Parameter(Mandatory=$True)]
        [string] $SourceBranch,
        [Parameter(Mandatory=$True)]
        [string] $DefaultBranch,
        [Parameter(Mandatory=$True)]
        [string] $ConfigFileDir,
        [string] $RepoName,
        [string] $BuildId,
        [bool] $MarkPackageAsShipped = $false
    )
    $artifactList = Create-APIViewArtifact -ServiceDirectory $ServiceDirectory -ArtifactOutputDirectory $ArtifactOutputDirectory -ArtifactName $ArtifactName
    $createReviewScript = (Join-Path $PSScriptRoot .. common scripts Create-APIReview.ps1)

    Write-Host "Create Go APIView using generated artifacts"

    &($createReviewScript) -ArtifactList $artifactList -ArtifactPath $outPath -APIKey $ApiKey -SourceBranch $SourceBranch -DefaultBranch $DefaultBranch -ConfigFileDir $ConfigFileDir -RepoName $RepoName -BuildId $BuildId -MarkPackageAsShipped $MarkPackageAsShipped
}