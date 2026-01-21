// Package printer handles the visual rendering of the tree structure.
package printer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cytificlabs/tr2rl/internal/parser"
)

// PrintTree constructs and prints a hierarchical Unicode tree from a flat list of nodes.
// It handles directory grouping, sorting, and prefix generation (e.g., "├── ").
// Options controls the output format of the tree.
type Options struct {
	Style string // "unicode" (default) or "ascii"
}

// PrintTree constructs and prints a hierarchical Unicode tree from a flat list of nodes.
func PrintTree(nodes []parser.Node) {
	PrintTreeWithOptions(nodes, Options{Style: "unicode"})
}

// PrintTreeWithOptions prints the tree with specific formatting options.
func PrintTreeWithOptions(nodes []parser.Node, opts Options) {
	if len(nodes) == 0 {
		return
	}

	// 1. Build a Tree map structure for printing
	childrenMap := make(map[string][]parser.Node)
	roots := make([]parser.Node, 0)
	nodeMap := make(map[string]parser.Node)

	for _, n := range nodes {
		key := strings.TrimSuffix(n.Path, "/")
		nodeMap[key] = n
	}

	for _, n := range nodes {
		key := strings.TrimSuffix(n.Path, "/")
		lastSlash := strings.LastIndex(key, "/")
		if lastSlash == -1 {
			roots = append(roots, n)
		} else {
			parentPath := key[:lastSlash]
			childrenMap[parentPath] = append(childrenMap[parentPath], n)
		}
	}

	sortNodes(roots)

	for i, root := range roots {
		printNode(root, "", i == len(roots)-1, childrenMap, opts)
	}
}

func printNode(node parser.Node, prefix string, isLast bool, childrenMap map[string][]parser.Node, opts Options) {
	// Markers
	var marker, link, noLink string
	
	if opts.Style == "ascii" {
		// ASCII Style:
		// |-- child
		// |   `-- sub
		// `-- last
		if isLast {
			marker = "`-- "
		} else {
			marker = "|-- "
		}
		link = "|   "
		noLink = "    "
	} else {
		// Unicode Style (Default):
		// ├── child
		// │   └── sub
		// └── last
		if isLast {
			marker = "└── "
		} else {
			marker = "├── "
		}
		link = "│   "
		noLink = "    "
	}

	name := node.Path
	if idx := strings.LastIndex(strings.TrimSuffix(name, "/"), "/"); idx != -1 {
		name = strings.TrimSuffix(name, "/")
		name = name[idx+1:]
	}
	name = strings.TrimSuffix(name, "/")
	if node.Kind == parser.Dir {
		name += "/"
	}

	fmt.Printf("%s%s%s\n", prefix, marker, name)

	// Calculate prefix for children
	childPrefix := prefix
	if isLast {
		childPrefix += noLink
	} else {
		childPrefix += link
	}

	key := strings.TrimSuffix(node.Path, "/")
	children := childrenMap[key]
	sortNodes(children)

	for i, child := range children {
		printNode(child, childPrefix, i == len(children)-1, childrenMap, opts)
	}
}

func sortNodes(nodes []parser.Node) {
	sort.Slice(nodes, func(i, j int) bool {
		// Sort Dirs first, then Files. Both alphabetical.
		if nodes[i].Kind != nodes[j].Kind {
			return nodes[i].Kind == parser.Dir
		}
		return nodes[i].Path < nodes[j].Path
	})
}
