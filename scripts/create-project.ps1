# Automated Project Creation Script
# Usage: .\create-project.ps1 -Name "project-name" -Description "Project description"

param(
    [Parameter(Mandatory=$true)]
    [string]$Name,
    
    [Parameter(Mandatory=$true)]
    [string]$Description
)

$TemplateDir = "C:\Users\Unkown\Desktop\Zone01\Spiros\combinatorics"
$User = "grsprs"

Write-Host "🚀 Creating project: $Name" -ForegroundColor Green

# 1. Create GitHub repo
gh repo create $Name --public --description $Description --clone
cd $Name

# 2. Initialize Go module
go mod init github.com/$User/$Name

# 3. Copy template files
Copy-Item "$TemplateDir\.github" -Recurse -Destination .
Copy-Item "$TemplateDir\.gitignore" -Destination .
Copy-Item "$TemplateDir\LICENSE" -Destination .
Copy-Item "$TemplateDir\CHANGELOG.md" -Destination .
Copy-Item "$TemplateDir\CONTRIBUTING.md" -Destination .
Copy-Item "$TemplateDir\SECURITY.md" -Destination .
Copy-Item "$TemplateDir\CODE_OF_CONDUCT.md" -Destination .
Copy-Item "$TemplateDir\AUTHORS.md" -Destination .
Copy-Item "$TemplateDir\codecov.yml" -Destination .

# 4. Create README
@"
# $Name

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CI Status](https://github.com/$User/$Name/workflows/CI/badge.svg)](https://github.com/$User/$Name/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/$User/$Name)](https://goreportcard.com/report/github.com/$User/$Name)
[![codecov](https://codecov.io/gh/$User/$Name/branch/main/graph/badge.svg)](https://codecov.io/gh/$User/$Name)
[![Go Reference](https://pkg.go.dev/badge/github.com/$User/$Name.svg)](https://pkg.go.dev/github.com/$User/$Name)

$Description

## Installation

``````bash
go get github.com/$User/$Name
``````

## License

MIT License - see [LICENSE](LICENSE) for details.

## Author

**Spiros Nikoloudakis** ([@grsprs](https://github.com/grsprs))

---

**Copyright © 2026 Spiros Nikoloudakis**
"@ | Out-File -FilePath "README.md" -Encoding UTF8

# 5. Initial commit
git add .
git commit -m "feat: initial project setup"
git push -u origin main

Write-Host "✅ Project created!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "1. Setup Codecov: https://codecov.io/" -ForegroundColor Cyan
Write-Host "2. Run: .\post-setup.ps1 -Name $Name -CodecovToken YOUR_TOKEN" -ForegroundColor Cyan
