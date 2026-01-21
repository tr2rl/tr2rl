# 🤝 Contributing to tr2rl

First off, thanks for taking the time to contribute! 🚀

`tr2rl` is designed to be the "missing bridge" between AI and filesystems. We value **robustness** (it should never crash on bad input) and **safety** (it should never destroy user data).

## 🛠️ Development Setup

1.  **Prerequisites**: Go 1.21+
2.  **Clone**: `git clone https://github.com/tr2rl/tr2rl`
3.  **Run**: `go run main.go --help`

## 📂 Project Structure

*   `cmd/`: CLI commands (built with Cobra).
*   `internal/parser/`: The **"Magic Parsing"** logic. (Go here if you want to improve how we handle messy text).
*   `internal/fs/`: Filesystem operations. (Crucial: Maintains `DryRun` safety checks).
*   `examples/`: Real-world tree files for diverse stacks.
*   `testdata/`: "Torture test" files for integration testing.

## 🧪 Testing

We take testing seriously because we touch the user's filesystem.

**Run the suite:**
```bash
go test ./... -v
```

**Add a new test case:**
1.  Add a tricky tree file to `testdata/`.
2.  Run `go run main.go spec testdata/your_file.tree --json` to see how it parses.

## 📝 Pull Request Guidelines

1.  **Fork** the repo on GitHub.
2.  **Clone** your fork locally.
3.  **Branch**: Create a feature branch (`git checkout -b feature/amazing-feature`).
4.  **Commit**: Commit changes (atomic commits preferred).
5.  **Test**: Ensure `go test ./...` passes.
6.  **Push**: Push to your branch and submit a PR!

### Style Guide
*   **Safety First**: Any feature that writes to disk MUST respect the `DryRun` flag.
*   **Zero Dependencies**: We try to keep the binary static and small. Avoid adding external packages unless absolutely necessary.
