package main

import (
	"os"
	"strings"
)

func helpText() strings.Builder {
	LE := getLE()
	var builder strings.Builder
	builder.WriteString("Usage: treelike [OPTIONS] [PATH]" + LE)
	builder.WriteString("Prints a tree-like representation of the input." + LE)
	builder.WriteString("" + LE)
	builder.WriteString("Options:" + LE)
	builder.WriteString("  -h, --help               Show this help message and exit" + LE)
	builder.WriteString("  -f, --file FILE          Read from FILE" + LE)
	builder.WriteString("   -, --stdin              Read from stdin" + LE)
	builder.WriteString("  -c, --charset CHARSET    Use CHARSET to display characters (utf-8, ascii)" + LE)
	builder.WriteString("  -s, --trailing-slash     Display trailing slash on directory" + LE)
	builder.WriteString("  -p, --full-path          Display full path" + LE)
	builder.WriteString("  -D, --no-root-dot        Do not display a root element" + LE)
	return builder
}

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

func getPrefixes(opts *Options) (string, string, string, string) {
	if opts.charset == "ascii" {
		return ASCII_CHILD, ASCII_LAST_CHILD, ASCII_DIRECTORY, ASCII_EMPTY
	}
	return UTF8_CHILD, UTF8_LAST_CHILD, UTF8_DIRECTORY, UTF8_EMPTY
}

func getLE() string {
	if os.IsPathSeparator('\\') {
		return LE_WIN
	}
	return LE_UNIX
}

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

func getName(node *Node, opts *Options) string {
	var chunks strings.Builder

	chunks.WriteString(node.name)

	if opts.trailingSlash && len(node.children) > 0 && node.name[len(node.name)-1] != '/' {
		chunks.WriteString("/")
	}

	str := chunks.String()

	if opts.fullPath && node.parent != nil {
		newOpts := Options{false, "", strings.Builder{}, opts.charset, true, opts.fullPath, opts.rootDot}
		str = getName(node.parent, &newOpts) + str
	}

	return str
}

func isLastChild(node *Node) bool {
	return node.parent != nil && node.parent.children[len(node.parent.children)-1] == node
}

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
