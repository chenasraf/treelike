package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
