// Package parser provides the core logic for checking and converting text-based
// tree specifications into a structured list of nodes.
//
// It employs a "Magic Parsing" strategy that attempts to recover structure from
// messy inputs, such as mixed indentation, missing markers, and path lists.
package parser

import (
	"path"
	"strings"
)

// Parse turns a text tree (Windows ASCII tree, Unicode tree, indented lists, mixed)
// into a flat list of Nodes with normalized slash-separated paths.
func Parse(input string) Result {
	lines := ScanLines(input)

	// Pre-filter lines to remove comments and junk
	// This ensures lines[0] is the actual first tree item for root detection.
	validLines := make([]LineInfo, 0, len(lines))
	for _, l := range lines {
		if l.IsComment {
			continue
		}
		// Also skip empty CleanName here?
		if l.CleanName == "" {
			continue
		}

		// Windows Tree Headers
		if strings.HasPrefix(l.Raw, "Folder PATH listing") ||
			strings.HasPrefix(l.Raw, "Volume serial number") {
			continue
		}

		// Windows Drive Anchor "C:." -> Treat as "." (or skip if we want implicit root)
		// Actually, let's normalize it to "."
		if len(l.CleanName) == 3 && l.CleanName[1] == ':' && l.CleanName[2] == '.' {
			l.CleanName = "."
			l.IsPathLike = true // Force it to look like a path so it's not filtered later
			// We need to modify 'l' but 'l' is a copy.
			// But we append 'l' to validLines. So if we modify l, it works.
		}

		validLines = append(validLines, l)
	}
	lines = validLines

	if len(lines) == 0 {
		return Result{}
	}

	result := Result{
		Warnings: make([]string, 0),
	}

	// Phase 1: Heuristic Analysis
	pathLikeCount := 0
	markerCount := 0
	for _, l := range lines {
		if l.IsPathLike {
			pathLikeCount++
		}
		if l.Marker != "" {
			markerCount++
		}
	}

	// Decision: Is this a Path List or a Tree?
	isPathList := false
	if pathLikeCount > len(lines)/2 && markerCount == 0 {
		isPathList = true
		// result.Warnings = append(result.Warnings, "Detected Path List format")
	}

	if isPathList {
		result.Nodes = parsePathList(lines)
	} else {
		result.Nodes, result.Warnings = parseTree(lines, result.Warnings, markerCount == 0)
	}

	// Normalize Output
	norm := make([]string, 0, len(result.Nodes))
	for _, n := range result.Nodes {
		p := n.Path
		if n.Kind == Dir && !strings.HasSuffix(p, "/") {
			p += "/"
		}
		norm = append(norm, p)
	}
	result.Normalized = strings.Join(norm, "\n")

	return result
}

func parsePathList(lines []LineInfo) []Node {
	nodes := make([]Node, 0, len(lines))
	seen := make(map[string]bool)

	for _, l := range lines {
		clean := strings.TrimSpace(l.Raw) // Use raw for path lists, but trim
		// Remove ./ prefix if present
		clean = strings.TrimPrefix(clean, "./")
		clean = strings.ReplaceAll(clean, "\\", "/")

		if clean == "" {
			continue
		}

		kind := File
		if strings.HasSuffix(clean, "/") {
			kind = Dir
			clean = strings.TrimSuffix(clean, "/")
		}

		if !seen[clean] {
			nodes = append(nodes, Node{Path: clean, Kind: kind})
			seen[clean] = true
		}
	}
	return nodes
}

func parseTree(lines []LineInfo, warnings []string, indentedListMode bool) ([]Node, []string) {
	// Root Handling: Check if first line is a root
	// A line is ROOT if:
	// 1. Depth is 0 (or very low compared to next)
	// 2. It has no markers
	// 3. Next line is deeper OR has markers

	nodes := make([]Node, 0, len(lines))
	stack := make([]string, 0, 32)

	// Heuristic: If first line has markers, it's probably NOT a root (it's a child of CWD)
	// But if first line has NO markers, and second line DOES, first line is Root.

	rootIdx := -1
	if len(lines) > 0 {
		l0 := lines[0]
		if l0.Marker == "" {
			// Check if it looks like a root wrapper?
			if len(lines) > 1 {
				l1 := lines[1]
				// If next line has indentation OR markers, we are the root.
				// Even if current line has no slash, if it heads a tree, it's a dir.
				// Specially, if l1 has a marker like "|--", that's indentation 0 usually.
				// But physically, if l0 is above it, l0 is the parent.
				if l1.Indent >= l0.Indent || l1.Marker != "" {
					rootIdx = 0
				}
			} else if len(lines) == 1 && isDirLike(lines[0].CleanName) {
				// Single line directory
				rootIdx = 0
			}
		}
	}

	// Initialize state
	// We track:
	// 1. stack: names of current path components ["root", "src"]
	// 2. indentStack: indentation values for each component [0, 4]

	indentStack := make([]int, 0, 32)

	if rootIdx == 0 {
		// We have a declared root
		rootName := strings.TrimSuffix(lines[0].CleanName, "/")
		stack = append(stack, rootName)
		indentStack = append(indentStack, lines[0].Indent)
		nodes = append(nodes, Node{Path: rootName, Kind: Dir})
		// Start processing children from index 1
		lines = lines[1:]
	} else {
		// Implicit root (current directory)
		// We behave as if there is a root at Indent = -1
		// So first item (Indent >= 0) becomes a child of it.
		// No item pushed to stack yet.
	}

	for _, l := range lines {
		indent := l.Indent
		name := l.CleanName

		if name == "" {
			continue
		}

		// 2. Junk Filter
		// EXCEPTION: If it has a valid marker, TRUST IT.
		// EXCEPTION: If we are in Indented List Mode (no markers at all), TRUST IT.
		if !indentedListMode && l.Marker == "" && !l.IsPathLike && strings.Contains(name, " ") && !looksLikeFile(name) {
			continue
		}

		// RELATIVE INDENTATION LOGIC:
		// Pop the stack until we find a parent with strictly LESS indentation than current line.
		// Or until stack is empty (if implicit root).

		// If explicit root exists, we must NOT pop the root (index 0).
		// The root acts as Indent=-Infinity effectively, but practically it has an indent (e.g. 0).
		// Children must have Indent > RootIndent.

		minStackSize := 0
		if rootIdx == 0 {
			minStackSize = 1
		}

		for len(stack) > minStackSize {
			topIndent := indentStack[len(indentStack)-1]
			if topIndent >= indent {
				// Current line is same level or shallower -> Pop to find sibling/parent
				stack = stack[:len(stack)-1]
				indentStack = indentStack[:len(indentStack)-1]
			} else {
				// Top indent < Current indent -> Top is Parent. Stop popping.
				break
			}
		}

		// If rootIdx==0 and we popped everything down to root,
		// we verify current indent > root indent.
		// If not, it technically shouldn't be a child, but standard behavior is to just add it to root?
		// Or it's a sibling of root? (Impossible in single-root tree).
		// Let's assume everything else is child of root.

		// Append current
		stack = append(stack, name)
		indentStack = append(indentStack, indent)

		// Determine Kind
		kind := File
		if isDirLike(name) {
			kind = Dir
			name = strings.TrimSuffix(name, "/")
		} else {
			if !looksLikeFile(name) {
				// default to file
			}
		}

		fullPath := path.Join(stack...)
		nodes = append(nodes, Node{Path: fullPath, Kind: kind})
	}

	// Post-pass: Fix "File" that became a parent
	// If Node A is a parent of Node B, Node A must be a Directory.
	// Since `nodes` is ordered, we can check basic containment?
	// Or better: In the loop, when we push to stack, we are declaring it a parent.
	// So anything remaining in `stack` (except the very last item) IS acting as a directory.

	// Let's fix the loop to handle "Make Parent" logic.
	// Actually, `stack` represents the current active directory chain.
	// So every time we push to stack, the *previous* stack top (if any) was effectively a specific node.
	// But `stack` only stores strings.
	// We need to update the `Node` kind in `nodes` list.

	// Map-based correction
	pathToKind := make(map[string]int) // Index in nodes
	for i, n := range nodes {
		pathToKind[n.Path] = i
	}

	for _, n := range nodes {
		// If "A/B" exists, then "A" must be a Dir
		parent := path.Dir(n.Path)
		if parent != "." && parent != "/" {
			if idx, ok := pathToKind[parent]; ok {
				nodes[idx].Kind = Dir
			}
		}
	}

	return nodes, warnings
}
