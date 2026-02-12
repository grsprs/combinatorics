# Project Automation Guide

## Quick Start - New Project in 2 Minutes

### Prerequisites (One-time setup):
```powershell
# Install GitHub CLI
winget install GitHub.cli

# Install VS Code extensions
code --install-extension golang.go
code --install-extension ryanluker.vscode-coverage-gutters
code --install-extension eamodio.gitlens
code --install-extension redhat.vscode-yaml
code --install-extension yzhang.markdown-all-in-one
```

### Create New Project:
```powershell
# 1. Create project (30 seconds)
cd C:\Users\Unkown\Desktop\Zone01\Spiros\combinatorics\scripts
.\create-project.ps1 -Name "my-project" -Description "My awesome Go library"

# 2. Setup Codecov (1 minute)
# - Go to https://codecov.io/
# - Add repository
# - Copy token

# 3. Run post-setup (30 seconds)
.\post-setup.ps1 -Name "my-project" -CodecovToken "YOUR_TOKEN"

# 4. Open in VS Code
cd ..\my-project
code .
```

## What Gets Automated:

### ✅ Repository Setup
- GitHub repo creation
- Branch protection rules
- Issue/PR templates
- CODEOWNERS file

### ✅ CI/CD Pipeline
- GitHub Actions workflow
- Test automation
- Linting (golangci-lint)
- Security scanning (gosec, govulncheck)
- Coverage reporting (Codecov)

### ✅ Documentation
- README with badges
- LICENSE (MIT)
- CHANGELOG
- CONTRIBUTING
- SECURITY
- CODE_OF_CONDUCT
- AUTHORS

### ✅ Development Environment
- VS Code settings
- VS Code tasks
- Extension recommendations
- Makefile

### ✅ Quality Tools
- Codecov integration
- Go Report Card ready
- pkg.go.dev ready

## VS Code Features

### Run Tests (Ctrl+Shift+P → "Tasks: Run Test Task"):
- Automatically runs tests on save
- Shows coverage in editor (green/red lines)
- Coverage report in browser

### Build CLI (Ctrl+Shift+P → "Tasks: Run Build Task"):
- Builds executable
- Output in `bin/` directory

### Keyboard Shortcuts:
- `Ctrl+Shift+P` → "Coverage Gutters: Display Coverage"
- `Ctrl+Shift+T` → Run tests
- `Ctrl+Shift+B` → Build

## Manual Steps (Still Required):

### After Project Creation:
1. **Write code** in your packages
2. **Write tests** (aim for 100% coverage)
3. **Update README** with usage examples
4. **First release**:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```
5. **Submit to Go Report Card**: https://goreportcard.com/

### Optional:
- Create GitHub Release with notes
- Announce on social media
- Submit to awesome-go lists

## Time Savings

| Task | Manual | Automated | Saved |
|------|--------|-----------|-------|
| Repo setup | 30 min | 30 sec | 29.5 min |
| CI/CD config | 45 min | 0 sec | 45 min |
| Documentation | 60 min | 0 sec | 60 min |
| GitHub settings | 15 min | 30 sec | 14.5 min |
| VS Code setup | 10 min | 0 sec | 10 min |
| **Total** | **2.5 hours** | **1 minute** | **2.5 hours** |

## Troubleshooting

### GitHub CLI not authenticated:
```bash
gh auth login
```

### Codecov token not working:
- Verify token is correct
- Check GitHub secrets: https://github.com/grsprs/PROJECT/settings/secrets/actions

### Branch protection fails:
- Ensure you have admin access
- Wait for first CI run to complete

## Template Repository (Alternative)

Create once, reuse forever:
```bash
# Create template from combinatorics
gh repo create go-project-template --template grsprs/combinatorics --public

# Use template for new projects
gh repo create my-project --template grsprs/go-project-template --public
```

## Summary

**One command to rule them all:**
```powershell
.\create-project.ps1 -Name "project" -Description "desc" && .\post-setup.ps1 -Name "project" -CodecovToken "token"
```

**Result:** Production-ready Go project in 2 minutes! 🚀
