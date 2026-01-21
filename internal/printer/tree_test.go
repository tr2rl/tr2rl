package printer

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/tr2rl/tr2rl/internal/parser"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Normalize line endings for Windows
	return strings.ReplaceAll(buf.String(), "\r\n", "\n")
}

func TestPrintTree_Styles(t *testing.T) {
	// Sample nodes
	nodes := []parser.Node{
		{Path: "root/", Kind: parser.Dir},
		{Path: "root/child/", Kind: parser.Dir},
		{Path: "root/child/file.txt", Kind: parser.File},
		{Path: "root/file2.txt", Kind: parser.File},
	}

	tests := []struct {
		name     string
		style    string
		expected string
	}{
		{
			name:  "Unicode Style (Default)",
			style: "unicode",
			expected: `└── root/
    ├── child/
    │   └── file.txt
    └── file2.txt
`,
		},
		{
			name:  "ASCII Style",
			style: "ascii",
			expected: "`" + `-- root/
    |-- child/
    |   ` + "`" + `-- file.txt
    ` + "`" + `-- file2.txt
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				PrintTreeWithOptions(nodes, Options{Style: tt.style})
			})

			if output != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, output)
			}
		})
	}
}
