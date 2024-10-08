package main

import (
	"os"
	"strings"
)

// describeTree generates a string representation of the tree structure starting from the given node.
// It recursively traverses the tree and collects lines representing each node and its children.
// The function ensures that blank lines are removed from the final output.
//
// Parameters:
//
//	node - The root node of the tree to describe.
//	opts - A pointer to an Options struct that specifies formatting options.
//
// Returns:
//
//	string - A string representation of the tree structure.
func describeTree(node *Node, opts *Options) string {
	lines := []string{getTreeLine(node, opts)}
	LE := getLE()

	for _, child := range node.children {
		next := describeTree(child, opts)
		for _, line := range strings.Split(next, LE) {
			if strings.TrimSpace(line) != "" {
				lines = append(lines, line)
			}
		}
	}

	var nonBlankLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonBlankLines = append(nonBlankLines, line)
		}
	}

	return strings.Join(nonBlankLines, LE)
}

// getPrefixes returns the appropriate tree drawing characters based on the specified charset in options.
// It supports both ASCII and UTF-8 character sets.
//
// Parameters:
//
//	opts - A pointer to an Options struct that specifies the charset.
//
// Returns:
//
//	string - The prefix for a child node.
//	string - The prefix for the last child node.
//	string - The prefix for a directory node.
//	string - The prefix for an empty node.
func getPrefixes(opts *Options) (string, string, string, string) {
	if opts.charset == "ascii" {
		return ASCII_CHILD, ASCII_LAST_CHILD, ASCII_DIRECTORY, ASCII_EMPTY
	}
	return UTF8_CHILD, UTF8_LAST_CHILD, UTF8_DIRECTORY, UTF8_EMPTY
}

// getLE returns the appropriate line ending based on the operating system.
// It returns "\r\n" for Windows and "\n" for Unix-based systems.
//
// Returns:
//
//	string - The line ending string.
func getLE() string {
	if os.IsPathSeparator('\\') {
		return LE_WIN
	}
	return LE_UNIX
}

// getTreeLine generates a string representing a single line of the tree structure for the given node.
// It constructs the line by adding appropriate tree drawing characters based on the node's position
// and the specified options. The function handles the root node, child nodes, and last child nodes
// differently to ensure the correct tree structure is represented.
//
// Parameters:
//
//	node - The node for which to generate the tree line.
//	opts - A pointer to an Options struct that specifies formatting options.
//
// Returns:
//
//	string - A string representing the tree line for the given node.
func getTreeLine(node *Node, opts *Options) string {
	if node.parent == nil {
		if opts.rootDot {
			return node.name
		} else {
			return ""
		}
	}

	CHILD, LAST_CHILD, DIRECTORY, EMPTY := getPrefixes(opts)

	var chunks strings.Builder

	if isLastChild(node) {
		chunks.WriteString(LAST_CHILD)
	} else {
		chunks.WriteString(CHILD)
	}

	chunks.WriteString(getName(node, opts))
	str := chunks.String()

	current := node.parent
	for current != nil && current.parent != nil {
		if isLastChild(current) {
			str = EMPTY + str
		} else {
			str = DIRECTORY + str
		}
		current = current.parent
	}

	if opts.rootDot {
		return str
	}
	return removePrefix(str, opts)
}

// getName generates the name of the node, optionally appending a trailing slash if the node has children
// and the trailingSlash option is enabled. If the fullPath option is enabled, it recursively constructs
// the full path of the node.
//
// Parameters:
//
//	node - The node for which to generate the name.
//	opts - A pointer to an Options struct that specifies formatting options.
//
// Returns:
//
//	string - The generated name of the node.
func getName(node *Node, opts *Options) string {
	var chunks strings.Builder

	chunks.WriteString(node.name)

	if opts.trailingSlash && len(node.children) > 0 && node.name[len(node.name)-1] != '/' {
		chunks.WriteString("/")
	}

	str := chunks.String()

	if opts.fullPath && node.parent != nil {
		newOpts := DefaultOptions()
		newOpts.charset = opts.charset
		newOpts.fullPath = opts.fullPath
		newOpts.rootPath = opts.rootPath
		newOpts.rootDot = opts.rootDot
		newOpts.trailingSlash = true
		str = getName(node.parent, newOpts) + str
	}

	return str
}

// isLastChild checks if the given node is the last child of its parent.
//
// Parameters:
//
//	node - The node to check.
//
// Returns:
//
//	bool - True if the node is the last child, false otherwise.
func isLastChild(node *Node) bool {
	return node.parent != nil && node.parent.children[len(node.parent.children)-1] == node
}

// removePrefix removes the tree drawing prefix from the given string based on the specified options.
// It ensures that the correct prefix is removed to maintain the tree structure.
//
// Parameters:
//
//	str - The string from which to remove the prefix.
//	opts - A pointer to an Options struct that specifies the charset.
//
// Returns:
//
//	string - The string with the prefix removed.
func removePrefix(str string, opts *Options) string {
	CHILD, LAST_CHILD, DIRECTORY, EMPTY := getPrefixes(opts)

	prefixes := []string{CHILD, LAST_CHILD, DIRECTORY, EMPTY}

	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return strings.Replace(str, prefix, "", 1)
		}
	}

	return str
}
