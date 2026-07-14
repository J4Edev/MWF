# MWF v0.3 Uninstaller

$ErrorActionPreference = "Stop"

$InstallDir = "$env:LOCALAPPDATA\Programs\MWF"

Write-Host "Removing MWF..."

if (Test-Path $InstallDir) {

    Remove-Item `
        -Path $InstallDir `
        -Recurse `
        -Force

    Write-Host "Removed files"
}
else {
    Write-Host "MWF installation not found"
}


# Remove from PATH

$userPath = [Environment]::GetEnvironmentVariable(
    "Path",
    "User"
)

if ($userPath -like "*$InstallDir*") {

    $newPath = (
        $userPath.Split(";") |
        Where-Object { $_ -ne $InstallDir }
    ) -join ";"

    [Environment]::SetEnvironmentVariable(
        "Path",
        $newPath,
        "User"
    )

    Write-Host "Removed PATH entry"
}


Write-Host ""
Write-Host "MWF has been uninstalled."