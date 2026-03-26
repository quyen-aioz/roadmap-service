## Roadmap Service

Service for AIOZ Network Roadmap

## 🔧 Pre-commit Hook Setup (GolangCI-Lint)

This project uses a Git pre-commit hook to ensure code quality by running `golangci-lint` before every commit.

### 🚀 Setup Steps

#### 1. Create the hook file

```bash
mkdir -p .git/hooks
nano .git/hooks/pre-commit
```

#### 2. Add the following script

```bash
#!/bin/sh

echo "🔍 Running golangci-lint..."

GOLANGCI_LINT=$(go env GOPATH)/bin/golangci-lint

if [ ! -f "$GOLANGCI_LINT" ]; then
  echo "⚠️ golangci-lint not found. Installing..."
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

$GOLANGCI_LINT run

if [ $? -ne 0 ]; then
  echo "❌ Lint failed. Fix issues before committing."
  exit 1
fi

echo "✅ Lint passed"
```

#### 3. Make it executable

```bash
chmod +x .git/hooks/pre-commit
```

---

### ✅ How it works

- Runs `golangci-lint` before every commit
- Blocks commit if lint errors are found
- Automatically installs `golangci-lint` if missing

---

### 🧪 Test it

```bash
git add .
git commit -m "test"
```

If lint fails → commit is blocked
If lint passes → commit succeeds

---

### ⚠️ Notes

- Ensure `$GOPATH/bin` is in your `PATH`
- You can bypass the hook (not recommended):

```bash
git commit --no-verify
```
