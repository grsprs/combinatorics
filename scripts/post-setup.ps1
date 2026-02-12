# Post-Setup Automation Script
# Usage: .\post-setup.ps1 -Name "project-name" -CodecovToken "token"

param(
    [Parameter(Mandatory=$true)]
    [string]$Name,
    
    [Parameter(Mandatory=$true)]
    [string]$CodecovToken
)

Write-Host "🔧 Running post-setup for: $Name" -ForegroundColor Green

# 1. Add Codecov token
gh secret set CODECOV_TOKEN --body $CodecovToken --repo grsprs/$Name

# 2. Enable branch protection
gh api repos/grsprs/$Name/branches/main/protection `
  --method PUT `
  --field required_status_checks='{"strict":true,"contexts":["Test","Lint","Security"]}' `
  --field enforce_admins=false `
  --field required_pull_request_reviews='{"required_approving_review_count":1}' `
  --field restrictions=null

# 3. Add topics
gh repo edit grsprs/$Name --add-topic go --add-topic golang

# 4. Enable features
gh repo edit grsprs/$Name --enable-issues --enable-discussions

Write-Host "✅ Post-setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "Repository ready at: https://github.com/grsprs/$Name" -ForegroundColor Cyan
