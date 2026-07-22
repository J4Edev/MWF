# MWF Installer
# Downloads latest release from GitHub

$ErrorActionPreference = "Stop"

$Repo = "J4Edev/MWF"
$InstallDir = "$env:LOCALAPPDATA\Programs\MWF"

Write-Host ""
Write-Host "=============================="
Write-Host " MWF Installer"
Write-Host "=============================="
Write-Host ""

function Get-LatestRelease {

    $api = "https://api.github.com/repos/$Repo/releases/latest"

    Invoke-RestMethod `
        -Uri $api `
        -Headers @{
            "User-Agent"="MWF-Installer"
        }
}


Write-Host "[1/5] Checking latest release..."

$release = Get-LatestRelease

$version = $release.tag_name

Write-Host "Latest version:"
Write-Host $version


$exeAsset = $release.assets |
    Where-Object {
        $_.name -eq "mwf.exe"
    }

$checksumAsset = $release.assets |
    Where-Object {
        $_.name -eq "checksums.txt"
    }


if (!$exeAsset -or !$checksumAsset) {

    Write-Host "Release missing required files."
    exit 1
}


$temp = Join-Path $env:TEMP "mwf-install"

New-Item `
    -ItemType Directory `
    -Force `
    -Path $temp | Out-Null


$exePath = "$temp\mwf.exe"
$checksumPath = "$temp\checksums.txt"


Write-Host "[2/5] Downloading MWF..."

Invoke-WebRequest `
    $exeAsset.browser_download_url `
    -OutFile $exePath


Invoke-WebRequest `
    $checksumAsset.browser_download_url `
    -OutFile $checksumPath


Write-Host "[3/5] Verifying SHA256..."

$expected = (
    Get-Content $checksumPath |
    Select-String "mwf.exe"
).ToString().Split()[0]


$actual = (
    Get-FileHash `
        $exePath `
        -Algorithm SHA256
).Hash.ToLower()


if ($expected -ne $actual) {

    Write-Host "Checksum verification failed!"
    exit 1
}


Write-Host "Checksum verified."


Write-Host "[4/5] Installing..."

New-Item `
    -ItemType Directory `
    -Force `
    -Path $InstallDir | Out-Null


Copy-Item `
    $exePath `
    "$InstallDir\mwf.exe" `
    -Force


Write-Host "[5/5] Updating PATH..."

$userPath = [Environment]::GetEnvironmentVariable(
    "Path",
    "User"
)


if ($userPath -notlike "*$InstallDir*") {

    [Environment]::SetEnvironmentVariable(
        "Path",
        "$userPath;$InstallDir",
        "User"
    )
}


Write-Host ""
Write-Host "MWF installed successfully."
Write-Host ""
Write-Host "Restart PowerShell and run:"
Write-Host ""
Write-Host "mwf --help"