package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// parseDepth calculates the depth of a line based on its leading whitespace characters.
// The depth is determined by counting the number of leading spaces and tabs.
// If an indent size is provided, the depth is divided by the indent size to normalize it.
//
// Parameters:
//
//	line - The input string line whose depth is to be calculated.
//	indentSize - The size of the indentation to be used for normalization.
//
// Returns:
//
//	int - The calculated depth of the line.
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

// parseInput parses a string input representing a tree structure and returns the root node of the tree.
// The input string should use indentation to represent the depth of each node in the tree.
// The function handles different line endings and adjusts the indentation size based on the input.
//
// Parameters:
//
//	input - A string representing the tree structure with nodes and indentation.
//
// Returns:
//
//	*Node - The root node of the parsed tree structure.
func parseInput(input string) *Node {
	input = strings.Replace(input, "\r", "", -1)
	root := &Node{".", 0, []*Node{}, nil}
	current := root
	indentSize := 0
	LE := LE_UNIX

	for _, line := range strings.Split(input, LE) {
		if line == "" {
			continue
		}
		depth := parseDepth(line, indentSize)
		if depth > 0 && indentSize == 0 {
			indentSize = depth
			depth /= indentSize
		}
		if depth < 0 {
			depth = 0
		}
		name := line[depth*indentSize:]
		if depth <= current.depth && current.parent != nil {
			for current != nil && current.depth >= depth {
				current = current.parent
			}
		}
		if current == nil {
			current = root
		}
		current.children = append(current.children, &Node{name, depth, []*Node{}, current})
		current = current.children[len(current.children)-1]
	}
	return root
}

// parseRawInput reads input based on the provided options and returns it as a strings.Builder.
// It can read from stdin, a file, or an extra string provided in the options.
// If an error occurs during reading, it returns the error with a description and an error code.
//
// Parameters:
//
//	opts - A pointer to an Options struct that specifies the input source.
//
// Returns:
//
//	strings.Builder - The input read from the specified source.
//	error - An error object if an error occurred, otherwise nil.
//	int - An error code: 0 for success, 1 for reading errors, 2 for missing input source.
func parseRawInput(opts *Options) (strings.Builder, error, int) {
	var input strings.Builder

	if opts.fromStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input.WriteString(scanner.Text() + "\n")
		}
		if err := scanner.Err(); err != nil {
			return input, fmt.Errorf("error reading from stdin: %w", err), 1
		}
	} else if opts.fromFile != "" {
		file, err := os.Open(opts.fromFile)
		if err != nil {
			return input, fmt.Errorf("error opening file %s: %w", opts.fromFile, err), 1
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input.WriteString(scanner.Text() + "\n")
		}
		if err := scanner.Err(); err != nil {
			return input, fmt.Errorf("error reading from file: %w", err), 1
		}
	} else if opts.extra.Len() > 0 {
		input.WriteString(opts.extra.String())
	} else {
		help := helpText()
		return input, fmt.Errorf(help.String()), 2
	}

	return input, nil, 0
}
