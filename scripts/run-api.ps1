# scripts/run-api.ps1
# Loads .env into the current PowerShell session, then runs the API.
# Does NOT modify .env or .env.example.

$envFile = Join-Path $PSScriptRoot "..\.env"

if (!(Test-Path $envFile)) {
  Write-Error ".env not found at: $envFile"
  exit 1
}

# Prevent stale session values from overriding .env
Remove-Item Env:DB_URL -ErrorAction SilentlyContinue

Get-Content $envFile |
  Where-Object { $_ -and -not $_.StartsWith("#") } |
  ForEach-Object {
    $k, $v = $_ -split "=", 2
    if ($k -and $v) { Set-Item -Path "Env:$k" -Value $v }
  }

go run .\cmd\api
