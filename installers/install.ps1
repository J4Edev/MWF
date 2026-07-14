# MWF v0.3 Installer
# Installs mwf.exe for the current Windows user

$ErrorActionPreference = "Stop"

$AppName = "MWF"
$InstallDir = "$env:LOCALAPPDATA\Programs\MWF"

Write-Host ""
Write-Host "================================"
Write-Host " MWF Windows Management Framework"
Write-Host " Installer v0.3"
Write-Host "================================"
Write-Host ""

# Locate binary
$SourceBinary = Join-Path $PSScriptRoot "..\dist\mwf.exe"

if (!(Test-Path $SourceBinary)) {
    Write-Host "ERROR: mwf.exe not found"
    Write-Host "Expected location:"
    Write-Host $SourceBinary
    exit 1
}

# Create install directory
Write-Host "[1/4] Creating install directory..."

if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}


# Copy executable
Write-Host "[2/4] Installing MWF..."

Copy-Item `
    -Path $SourceBinary `
    -Destination "$InstallDir\mwf.exe" `
    -Force


# Add to PATH
Write-Host "[3/4] Updating user PATH..."

$userPath = [Environment]::GetEnvironmentVariable(
    "Path",
    "User"
)

if ($userPath -notlike "*$InstallDir*") {

    if ([string]::IsNullOrEmpty($userPath)) {
        $newPath = $InstallDir
    }
    else {
        $newPath = "$userPath;$InstallDir"
    }

    [Environment]::SetEnvironmentVariable(
        "Path",
        $newPath,
        "User"
    )

    Write-Host "Added MWF to PATH"
}
else {
    Write-Host "MWF already exists in PATH"
}


# Verify installation
Write-Host "[4/4] Verifying installation..."

if (Test-Path "$InstallDir\mwf.exe") {

    Write-Host ""
    Write-Host "Installation complete!"
    Write-Host ""

    Write-Host "Installed:"
    Write-Host "$InstallDir\mwf.exe"

    Write-Host ""
    Write-Host "Restart PowerShell, then run:"
    Write-Host ""
    Write-Host "  mwf --help"
    Write-Host ""

}
else {

    Write-Host "Installation failed"
    exit 1
}