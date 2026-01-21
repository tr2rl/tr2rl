package e2e

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStress_Chaos(t *testing.T) {
	// "Chaos Tree":
	// 1. Massive Indentation Jumps ("The Void")
	// 2. Subtle Misalignments (2 spaces vs 3 spaces)
	// 3. Special Characters in filenames
	// 4. Deep Nesting
	input := `root
  regular_child
                      void_child
   subtle_child_of_regular
  !@#$%^&()_+
  [brackets]
  {braces}
  unicode_🚀
`
	// Expected Structure:
	// root
	// ├── regular_child/
	// │   ├── subtle_child_of_regular
	// │   └── void_child
	// ├── !@#$%^&()_+
	// ├── [brackets]
	// ├── {braces}
	// └── unicode_🚀

	// Explanation of logic:
	// "regular_child" (2 spaces)
	// "void_child" (20+ spaces) -> Child of regular_child? Wait.
	// Parser stack: [root(0), regular(2)]. Next: void(22).
	// 22 > 2. Push void. Stack: [root, regular, void].
	// Next: subtle(3).
	// Pop void(22). Stack: [root, regular].
	// Top regular(2) < subtle(3). Stop. Push subtle.
	// Stack: [root, regular, subtle].
	// So "subtle" is a child of "regular", and "void" is a sibling of "subtle"?
	// No, "void" was popped. It's done.
	// "subtle" is added as child of "regular".
	// "void" was *also* a child of "regular".
	// So "regular" has children: "void", "subtle".
	// This seems correct visually.

	tmpFile := filepath.Join(t.TempDir(), "chaos.txt")
	os.WriteFile(tmpFile, []byte(input), 0644)

	// 1. Test Format
	output, err := runCLI("format", tmpFile, "--style=unicode")
	if err != nil {
		t.Fatalf("Format command failed: %v", err)
	}

	if !strings.Contains(output, "void_child") {
		t.Errorf("Massive indent jump failed. Output:\n%s", output)
	}
	if !strings.Contains(output, "unicode_🚀") {
		t.Errorf("Unicode filename failed.")
	}
	// Check subtle nesting
	// We expect "subtle" to be inside "regular"
	// Heuristic check: "regular" should appear before "subtle" and markers should align?
	// Hard to check exact tree structure with strings.Contains, but let's check basic presence.

	// 2. Test Build (Integration)
	buildDir := filepath.Join(t.TempDir(), "build_chaos")

	// Command: build <input> <output_dir> --dry-run=false
	_, err = runCLI("build", tmpFile, buildDir, "--dry-run=false")
	if err != nil {
		t.Fatalf("Build command failed: %v", err)
	}

	// Verify file creation
	// 1. void_child
	// Expected path: build_chaos/root/regular_child/void_child
	voidPath := filepath.Join(buildDir, "root", "regular_child", "void_child")
	if _, err := os.Stat(voidPath); os.IsNotExist(err) {
		t.Errorf("Build failed to create deep nested file: %s", voidPath)
	}

	// 2. Unicode file
	// Expected path: build_chaos/root/unicode_🚀
	uniPath := filepath.Join(buildDir, "root", "unicode_🚀")
	if _, err := os.Stat(uniPath); os.IsNotExist(err) {
		t.Errorf("Build failed to create unicode file: %s", uniPath)
	}

	// 3. Brackets
	// Expected path: build_chaos/root/[brackets]
	bracketPath := filepath.Join(buildDir, "root", "[brackets]")
	if _, err := os.Stat(bracketPath); os.IsNotExist(err) {
		t.Errorf("Build failed to create file with brackets: %s", bracketPath)
	}
}
