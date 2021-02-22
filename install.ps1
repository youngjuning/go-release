#!/usr/bin/env pwsh
# Copyright 2020 youngjuning. All rights reserved. MIT license.
# TODO(everyone): Keep this script simple and easily auditable.
# ${l} latestReleaseURL
# ${s} shellName
# ${h} homeDirName
Write-Output "shellName：${s}"
Write-Output "latestReleaseURL：${l}"
Write-Output "homeDirName：${h}"

$ErrorActionPreference = 'Stop'

$BinDir = "$Home\${h}\bin"
$BinZip = "$BinDir\${s}.zip"
$BinExe = "$BinDir\${s}.exe"
$Target = 'x86_64-pc-windows-msvc'

# GitHub requires TLS 1.2
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

$BinUri = "${l}/download/${s}-${Target}.zip"

if (!(Test-Path $BinDir)) {
  New-Item $BinDir -ItemType Directory | Out-Null
}

Invoke-WebRequest $BinUri -OutFile $BinZip -UseBasicParsing

if (Get-Command Expand-Archive -ErrorAction SilentlyContinue) {
  Expand-Archive $BinZip -Destination $BinDir -Force
} else {
  if (Test-Path $BinExe) {
    Remove-Item $BinExe
  }
  Add-Type -AssemblyName System.IO.Compression.FileSystem
  [IO.Compression.ZipFile]::ExtractToDirectory($BinZip, $BinDir)
}

Remove-Item $BinZip

$User = [EnvironmentVariableTarget]::User
$Path = [Environment]::GetEnvironmentVariable('Path', $User)
if (!(";$Path;".ToLower() -like "*;$BinDir;*".ToLower())) {
  [Environment]::SetEnvironmentVariable('Path', "$Path;$BinDir", $User)
  $Env:Path += ";$BinDir"
}

Write-Output "${s} was installed successfully to $BinExe"
Write-Output "Run '${s}' to get started"
