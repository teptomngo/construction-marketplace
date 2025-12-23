Param(
  [Parameter(Mandatory = $true)]
  [ValidateSet("up","version","goto","force","down","drop")]
  [string]$Command,

  [string]$Arg
)

$ErrorActionPreference = "Stop"

# Load .env (local-only) if it exists
$envFile = Join-Path $PSScriptRoot "..\.env"
if (Test-Path $envFile) {
  Get-Content $envFile | ForEach-Object {
    $line = $_.Trim()
    if ($line -and -not $line.StartsWith("#") -and $line.Contains("=")) {
      $k, $v = $line.Split("=", 2)
      $k = $k.Trim()
      $v = $v.Trim()
      if ($k -and $v) { Set-Item -Path "Env:$k" -Value $v }
    }
  }
}

if (-not $env:DB_URL) {
  throw "DB_URL is not set. Add DB_URL to .env (copy from .env.example)."
}

# Ensure we join the same docker network as compose
$network = "conmesh_default"

$baseArgs = @(
  "run", "--rm",
  "--network", $network,
  "-v", "$((Get-Location).Path)\migrations:/migrations",
  "migrate/migrate",
  "-path", "/migrations",
  "-database", $env:DB_URL
)

if ($Arg) {
  docker @baseArgs $Command $Arg
} else {
  docker @baseArgs $Command
}
