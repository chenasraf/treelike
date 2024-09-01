package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func helpText() {
	fmt.Println("Usage: treelike [OPTIONS] [PATH]")
	fmt.Println("Prints a tree-like representation of the input.")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h, --help               Show this help message and exit")
	fmt.Println("  -f, --file FILE          Read from FILE")
	fmt.Println("   -, --stdin              Read from stdin")
	fmt.Println("  -c, --charset CHARSET    Use CHARSET to display characters (utf-8, ascii)")
	fmt.Println("  -s, --trailing-slash     Display trailing slash on directory")
	fmt.Println("  -p, --full-path          Display full path")
	fmt.Println("  -D, --no-root-dot        Do not display a root element")
}

type Options struct {
	fromStdin     bool
	fromFile      string
	extra         strings.Builder
	charset       string
	trailingSlash bool
	fullPath      bool
	rootDot       bool
}

func getOpts() Options {
	opts := Options{false, "", strings.Builder{}, "utf-8", false, false, true}
	args := os.Args[1:]
	for len(args) > 0 {
		switch args[0] {
		case "-h", "--help":
			{
				helpText()
				os.Exit(0)
			}
		case "-", "--stdin":
			{
				opts.fromStdin = true
				args = args[1:]
			}
		case "-f", "--file":
			{
				opts.fromFile = args[1]
				args = args[2:]
			}
		case "-c", "--charset":
			{
				opts.charset = args[1]
				if opts.charset != "utf-8" && opts.charset != "ascii" {
					fmt.Fprintf(os.Stderr, "Invalid charset: %s\n", opts.charset)
					os.Exit(1)
				}
				args = args[2:]
			}
		case "-s", "--trailing-slash":
			{
				opts.trailingSlash = true
				args = args[1:]
			}
		case "-p", "--full-path":
			{
				opts.fullPath = true
				args = args[1:]
			}
		case "-D", "--no-root-dot":
			{
				opts.rootDot = false
				args = args[1:]
			}
		default:
			{
				opts.extra.WriteString(args[0] + "\n")
				args = args[1:]
			}
		}
	}
	return opts
}

type Node struct {
	// name of node
	name string
	// depth of node
	depth int
	// children of node
	children []*Node
	// parent of node
	parent *Node
}

func parseDepth(line string, indentSize int) int {
	depth := 0
	for j := 0; j < len(line); j++ {
		if line[j] != ' ' && line[j] != '\t' {
			break
		}
		depth++
	}
	if indentSize > 0 {
		depth /= indentSize
	}
	return depth
}

func parseInput(input string) *Node {
	root := &Node{".", 0, []*Node{}, nil}
	current := root
	indentSize := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		depth := parseDepth(line, indentSize)
		if depth > 0 && indentSize == 0 {
			indentSize = depth
			depth /= indentSize
		}
		name := line[depth*indentSize:]
		if depth <= current.depth && current.parent != nil {
			for current.depth >= depth {
				current = current.parent
			}
		}
		current.children = append(current.children, &Node{name, depth, []*Node{}, current})
		current = current.children[len(current.children)-1]
	}
	return root
}

func describeTree(node *Node, opts *Options) string {
	lines := []string{getTreeLine(node, opts)}

	for _, child := range node.children {
		next := describeTree(child, opts)
		for _, line := range strings.Split(next, "\n") {
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

	return strings.Join(nonBlankLines, "\n")
}

const (
	UTF8_CHILD      string = "├── "
	UTF8_LAST_CHILD string = "└── "
	UTF8_DIRECTORY  string = "│   "
	UTF8_EMPTY      string = "    "
)

const (
	ASCII_CHILD      string = "|-- "
	ASCII_LAST_CHILD string = "`-- "
	ASCII_DIRECTORY  string = "|   "
	ASCII_EMPTY      string = "    "
)

func getPrefixes(opts *Options) (string, string, string, string) {
	if opts.charset == "ascii" {
		return ASCII_CHILD, ASCII_LAST_CHILD, ASCII_DIRECTORY, ASCII_EMPTY
	}
	return UTF8_CHILD, UTF8_LAST_CHILD, UTF8_DIRECTORY, UTF8_EMPTY
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

func main() {
	var input strings.Builder

	opts := getOpts()

	if opts.fromStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input.WriteString(scanner.Text() + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(1)
		}
	} else if opts.fromFile != "" {
		file, err := os.Open(opts.fromFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", opts.fromFile, err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input.WriteString(scanner.Text() + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from file:", err)
			os.Exit(1)
		}
	} else if opts.extra.Len() > 0 {
		input.WriteString(opts.extra.String())
	} else {
		helpText()
		os.Exit(1)
	}
	node := parseInput(input.String())
	fmt.Println(describeTree(node, &opts))
}
