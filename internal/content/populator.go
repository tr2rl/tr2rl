package content

import (
	"fmt"
	"path/filepath"
	"strings"
)

// GetContent returns smart default content for a file based on its name/extension.
func GetContent(path string) string {
	base := filepath.Base(path)
	ext := strings.ToLower(filepath.Ext(path))

	// 1. Exact Filename Matches
	switch strings.ToLower(base) {
	case "makefile":
		return "all:\n\t@echo 'Hello World'\n"
	case "dockerfile":
		return "FROM alpine:latest\nCMD [\"echo\", \"Hello World\"]\n"
	case ".gitignore":
		return "# Ignore list\n.DS_Store\nnode_modules/\ndist/\nbin/\n"
	case "license", "license.txt", "license.md":
		return "MIT License\n\nCopyright (c) 2026\n"
	}

	// 2. Extension Matches (Specific Logic)
	switch ext {
	// Go
	case ".go":
		if base == "main.go" {
			return "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello World\")\n}\n"
		}
		return "package " + guessPackage(path) + "\n"

	// Web / JS / TS
	case ".html":
		return "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>" + base + "</title>\n</head>\n<body>\n    <h1>" + base + "</h1>\n</body>\n</html>\n"
	case ".css":
		return "/* " + base + " */\nbody {\n    font-family: sans-serif;\n    margin: 0;\n}\n"
	case ".js":
		return "// " + base + "\nconsole.log('Hello from " + base + "');\n"
	case ".jsx", ".tsx":
		compName := strings.TrimSuffix(strings.Title(base), ext)
		return "import React from 'react';\n\nexport const " + compName + " = () => {\n    return <div>" + compName + "</div>;\n};\n"
	case ".json":
		return "{}\n"
	case ".vue":
		return "<template>\n  <div>" + base + "</div>\n</template>\n\n<script>\nexport default {\n  name: '" + strings.TrimSuffix(base, ext) + "'\n}\n</script>\n\n<style scoped>\n</style>\n"

	// Python
	case ".py":
		if base == "main.py" || base == "app.py" {
			return "def main():\n    print(\"Hello from " + base + "\")\n\nif __name__ == \"__main__\":\n    main()\n"
		}
		return "# " + base + "\n"

	// Java / JVM
	case ".java":
		cls := strings.TrimSuffix(base, ext)
		return "public class " + cls + " {\n    public static void main(String[] args) {\n        System.out.println(\"Hello from " + cls + "\");\n    }\n}\n"
	case ".kt":
		return "fun main() {\n    println(\"Hello form " + base + "\")\n}\n"

	// C / C++
	case ".c":
		return "#include <stdio.h>\n\nint main() {\n    printf(\"Hello from " + base + "\\n\");\n    return 0;\n}\n"
	case ".cpp", ".cc":
		return "#include <iostream>\n\nint main() {\n    std::cout << \"Hello from " + base + "\" << std::endl;\n    return 0;\n}\n"
	case ".h", ".hpp":
		guard := strings.ToUpper(strings.ReplaceAll(base, ".", "_"))
		return "#ifndef " + guard + "\n#define " + guard + "\n\n// Content\n\n#endif // " + guard + "\n"

	// Rust
	case ".rs":
		return "fn main() {\n    println!(\"Hello from " + base + "\");\n}\n"

	// Shell
	case ".sh":
		return "#!/bin/bash\nset -euo pipefail\n\necho \"Running " + base + "\"\n"

	// Ruby
	case ".rb":
		return "# frozen_string_literal: true\n\nputs '" + base + "'\n"

	// PHP
	case ".php":
		return "<?php\n\necho \"Hello from " + base + "\";\n"

	// Config / Data
	case ".yaml", ".yml":
		return "# " + base + "\nversion: '1.0'\n"
	case ".toml":
		return "# " + base + "\n"
	case ".xml":
		return "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root>\n    <!-- " + base + " -->\n</root>\n"
	case ".md":
		title := strings.TrimSuffix(base, ext)
		return fmt.Sprintf("# %s\n\nFile: %s\n", strings.Title(title), base)
	case ".txt":
		return base + "\n"
	}

	// 3. Generic Fallback (Comment based on extension)
	// If we don't have a specific template, try to add a generic comment header
	if isSlashComment(ext) {
		return "// File: " + base + "\n"
	}
	if isHashComment(ext) {
		return "# File: " + base + "\n"
	}

	// Default: Empty
	return ""
}

func guessPackage(path string) string {
	// parent dir name
	dir := filepath.Base(filepath.Dir(path))
	if dir == "." || dir == "/" {
		return "main"
	}
	// Sanitize package name (lowercase, no symbols)
	clean := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		if r >= 'A' && r <= 'Z' {
			return r + 32 // toLower
		}
		return -1 // drop
	}, dir)

	if clean == "" {
		return "pkg"
	}
	return clean
}

func isSlashComment(ext string) bool {
	switch ext {
	case ".js", ".ts", ".jsx", ".tsx", ".java", ".c", ".cpp", ".cc", ".cs", ".go", ".rs", ".php", ".swift", ".kt", ".scala", ".dart", ".rust":
		return true
	}
	return false
}

func isHashComment(ext string) bool {
	switch ext {
	case ".py", ".rb", ".sh", ".pl", ".pm", ".yaml", ".yml", ".toml", ".conf", ".properties", ".dockerfile", ".makefile", ".r", ".el":
		return true
	}
	return false
}
