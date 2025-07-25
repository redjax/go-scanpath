<#
    .SYNOPSIS
    My Go build script.

    .DESCRIPTION
    Accepts a common set of parameters to automate (re)building a Go app.

    .PARAMETER BinName
    The name of your executable, i.e. ./`$BinName.

    .PARAMETER BuildOS
    The OS to build for. Full ist available at https://github.com/golang/go/blob/master/src/internal/syslist/syslist.go

    .PARAMETER BuildArch
    The CPU architecture to build for. A full list does not seem to be available,
    but more info in this Gist: https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

    .PARAMETER BuildOutputDir
    The build artifact path, where build outputs will be saved.

    .PARAMETER BuildTarget
    The name of the file to build (the entrypoint for your app).

    .EXAMPLE
    .\build.ps1 -BinName "mycli" -BuildOS "windows" -BuildArch "amd64" -BuildOutputDir "dist/"
#>
Param(
    [Parameter(Mandatory = $false, HelpMessage = "The name of your executable, i.e. ./`$BinName.")]
    $BinName = $null, ## Change this value for each new project
    [Parameter(Mandatory = $false, HelpMessage = "The OS to build for. Full ist available at https://github.com/golang/go/blob/master/src/internal/syslist/syslist.go")]
    $BuildOS = "windows",
    [Parameter(Mandatory = $false, HelpMessage = "The CPU architecture to build for. A full list does not seem to be available, but more info in this Gist: https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63")]
    $BuildArch = "amd64",
    [Parameter(Mandatory = $false, HelpMessage = "The build artifact path, where build outputs will be saved.")]
    $BuildOutputDir = "bin/",
    [Parameter(Mandatory = $false, HelpMessage = "The name of the file to build (the entrypoint for your app).")]
    $BuildTarget = "./main.go"
)

Write-Debug "BinName: $BinName"
Write-Debug "BuildOS: $BuildOS"
Write-Debug "BuildArch: $BuildArch"
Write-Debug "BuildOutputDir: $BuildOutputDir"
Write-Debug "BuildTarget: $BuildTarget"

if ( $null -eq $BinName ) {
    Write-Warning "No bin name provided, pass the name of your executable using the -BinName flag"
    exit(1)
}

$env:GOOS = $BuildOS
$env:GOARCH = $BuildArch

$BuildOutput = Join-Path -Path $BuildOutputDir -ChildPath $BinName
Write-Debug "Build output: $BuildOutput"

Write-Host "Building $($BuildTarget), outputting to $($BuildOutput)" -ForegroundColor Cyan
Write-Information "-- [ Build start"
try {
    go build -o $BuildOutput $BuildTarget
    Write-Host "Build successful" -ForegroundColor Green
}
catch {
    Write-Error "Error building app. Details: $($_.Exception.Message)"
    exit(1)
}
finally {
    Write-Information "-- [ Build complete"
}

exit(0)
