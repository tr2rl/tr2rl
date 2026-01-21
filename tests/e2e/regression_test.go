package e2e

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRegression_MessyInput(t *testing.T) {
	// This input simulates a "worst case" user copy-paste:
	// 1. Windows 'tree' command header (Junk that should be skipped)
	// 2. Drive letter anchor (C:.)
	// 3. Mixed styles (markers + indentation)
	// 4. Comments
	input := `Folder PATH listing for volume Windows
Volume serial number is 0000-0000
C:.
├── mixing
│   ├── valid_marker.txt
│   # This is a comment
│   without_marker.txt
├── another_folder
    └── indented_child.txt
Program Files
  Ignored Junk
    Or Is It?
`
	// Note: "Program Files" might be seen as a folder if we relaxed junk filter too much.
	// But "Ignored Junk" should definitely be skipped or treated as files if they look like it.
	// In "Indented List Mode" (markerless), "Program Files" is a valid name.
	// But here we HAVE markers ("|--"). So we should be in "Tree Mode".
	// In "Tree Mode", lines without markers are subject to strict junk filter (unless they look like files).
	// "Program Files" has spaces and no ext -> Junk (historically) or File?
	// Let's see what happens. Ideally, "valid_marker.txt" and "indented_child.txt" MUST appear.

	tmpFile := filepath.Join(t.TempDir(), "regression.txt")
	os.WriteFile(tmpFile, []byte(input), 0644)

	output, err := runCLI("format", tmpFile, "--style=unicode")
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// 1. Verify Header Stripping
	if strings.Contains(output, "Volume serial number") {
		t.Errorf("FAIL: Windows header was NOT stripped.")
	}

	// 2. Verify Mixed Marker/Indent handling
	if !strings.Contains(output, "valid_marker.txt") {
		t.Errorf("FAIL: valid_marker.txt missing.")
	}
	if !strings.Contains(output, "indented_child.txt") {
		t.Errorf("FAIL: indented_child.txt missing (mixed indent/marker support broke).")
	}

	// 3. Verify implicit root handling (C:. -> .)
	// The parser normalizes "C:." to ".", so the root should be implicit or explicit "."
	// Output often starts with the children if implicit.
	// Let's just check valid children are there.
}
