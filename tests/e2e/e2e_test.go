package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var binaryPath string

func TestMain(m *testing.M) {
	// 1. Build the binary once for all tests
	if runtime.GOOS == "windows" {
		binaryPath = filepath.Join(os.TempDir(), "tr2rl_e2e.exe")
	} else {
		binaryPath = filepath.Join(os.TempDir(), "tr2rl_e2e")
	}

	// Go build from the project root
	// rootDir, _ := os.Getwd()
	// projectRoot := filepath.Dir(filepath.Dir(rootDir)) // Unused

	// Wait, typical go test run is from root or we need to know where main.go is.
	// Let's assume we run `go test ./tests/e2e/...` from root.
	// Actually, `os.Getwd()` will be the package directory when running `go test`.

	// Harder to find main.go via relative path if we don't know where we started.
	// Robust way: assume we are in tests/e2e, so main is at ../../

	cmd := exec.Command("go", "build", "-o", binaryPath, "../../")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to build binary: %s\n%s\n", err, out)
		os.Exit(1)
	}

	// 2. Run tests
	code := m.Run()

	// 3. Cleanup
	os.Remove(binaryPath)
	os.Exit(code)
}

func runCLI(args ...string) (string, error) {
	cmd := exec.Command(binaryPath, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return strings.TrimSpace(out.String()), err
}

func TestFormat_ASCII(t *testing.T) {
	// Input: simple tree
	input := `root/
  child/
    file.txt`

	// Create temp input file
	tmpFile := filepath.Join(t.TempDir(), "ansi_test.txt")
	os.WriteFile(tmpFile, []byte(input), 0644)

	output, err := runCLI("format", tmpFile, "--style=ascii")
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	// Expected ASCII output (Windows compatible checks)
	if !strings.Contains(output, "`-- child") {
		t.Errorf("Expected ASCII style marker '`--', got:\n%s", output)
	}
}

func TestBuild_Verification(t *testing.T) {
	// Input: Valid tree
	input := `e2e_project/
├── src/
│   └── main.py
└── README.md`

	tmpInput := filepath.Join(t.TempDir(), "spec.tree")
	os.WriteFile(tmpInput, []byte(input), 0644)

	outputDir := t.TempDir()

	// Run Build (NOT dry run)
	out, err := runCLI("build", tmpInput, outputDir, "--dry-run=false", "--populate")
	if err != nil {
		t.Fatalf("Build failed: %v\nOutput: %s", err, out)
	}

	// Verify Filesystem
	expectedMain := filepath.Join(outputDir, "e2e_project", "src", "main.py")
	if _, err := os.Stat(expectedMain); os.IsNotExist(err) {
		t.Errorf("File verified missing: %s", expectedMain)
	}

	// Verify Populate Content
	content, _ := os.ReadFile(expectedMain)
	if !strings.Contains(string(content), "def main():") {
		t.Errorf("File content missing python boilerplate. Got: %s", string(content))
	}
}

func TestFormat_ProLevel(t *testing.T) {
	// "Nathan Friend" Style Input (Strict Indentation, mixed bullets)
	input := `Edit me to generate
  a
    nice
      tree
        diagram!
        :)
  Use indentation
    to indicate
      file
      and
      folder
      nesting.
    - You can even
      - use
        - markdown
        - bullets!`

	tmpFile := filepath.Join(t.TempDir(), "nathan.txt")
	os.WriteFile(tmpFile, []byte(input), 0644)

	output, err := runCLI("format", tmpFile, "--style=ascii")
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Verify deep nesting retention (a -> nice -> tree)
	if !strings.Contains(output, "`"+`-- tree/`) {
		t.Errorf("Deep nesting missing. Expected tree/ to be a leaf (or close). output:\n%s", output)
	}

	// Verify "Use indentation" survived the junk filter
	if !strings.Contains(output, "Use indentation/") {
		t.Errorf("Line with spaces 'Use indentation' was filtered out incorrectly!")
	}

	// Verify "a" is present
	// Need to be careful with matching "a/" vs "data/" etc.
	if !strings.Contains(output, "-- a/") {
		t.Errorf("Node 'a/' missing from output.")
	}
}

func TestFormat_MixedIndent(t *testing.T) {
	// "Mixed Relative Indentation" Input
	input := `c proj
    src
        hellow.c
    header
      hellow.h
    
    src1
       hellow1.c
    header1
      hellow1.h`

	tmpFile := filepath.Join(t.TempDir(), "mixed_indent.txt")
	os.WriteFile(tmpFile, []byte(input), 0644)

	output, err := runCLI("format", tmpFile, "--style=unicode")
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Verify siblings
	// All 4 folders should be at the same level (children of "c proj")
	// "c proj" is root.
	// So output should contain:
	// ├── header/
	// ├── header1/
	// ├── src/
	// └── src1/ (order depends on sort, but structure matters)

	if !strings.Contains(output, "├── header/") || !strings.Contains(output, "├── src/") {
		t.Errorf("Broken structure! header and src should be siblings. Output:\n%s", output)
	}
}
