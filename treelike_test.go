package main

import (
	"strings"
	"testing"
)

func TestParseDepth(t *testing.T) {
	tests := []struct {
		line       string
		indentSize int
		expected   int
	}{
		{"    node", 4, 1},
		{"        node", 4, 2},
		{"node", 4, 0},
		{"\tnode", 0, 1},
	}

	for _, test := range tests {
		result := parseDepth(test.line, test.indentSize)
		if result != test.expected {
			t.Errorf("parseDepth(%q, %d) = %d; want %d", test.line, test.indentSize, result, test.expected)
		}
	}
}

func TestParseInput(t *testing.T) {
	input := "root\n    child1\n    child2\n        grandchild1\n"
	root := parseInput(input)

	if root.name != "." {
		t.Errorf("Expected root name to be '.', got %s", root.name)
	}

	if len(root.children) != 1 {
		t.Fatalf("Expected root to have 1 child, got %d", len(root.children))
	}

	child1 := root.children[0]
	if child1.name != "root" {
		t.Errorf("Expected child1 name to be 'root', got %s", child1.name)
	}

	if len(child1.children) != 2 {
		t.Fatalf("Expected child1 to have 2 children, got %d", len(child1.children))
	}

	grandchild1 := child1.children[1].children[0]
	if grandchild1.name != "grandchild1" {
		t.Errorf("Expected grandchild1 name to be 'grandchild1', got %s", grandchild1.name)
	}
}

func TestDescribeTree(t *testing.T) {
	input := "root\n    child1\n    child2\n        grandchild1\n"
	root := parseInput(input)
	opts := Options{rootDot: true}
	result := describeTree(root, &opts)

	expected := ".\n└── root\n    ├── child1\n    └── child2\n        └── grandchild1"
	if result != expected {
		t.Errorf("describeTree() = %q; want %q", result, expected)
	}
}

func TestGetAsciiLine(t *testing.T) {
	root := &Node{name: ".", depth: 0, children: []*Node{}, parent: nil}
	child := &Node{name: "child", depth: 1, children: []*Node{}, parent: root}
	root.children = append(root.children, child)
	opts := Options{rootDot: true}

	result := getTreeLine(child, &opts)
	expected := "└── child"
	if result != expected {
		t.Errorf("getTreeLine() = %q; want %q", result, expected)
	}
}

func TestGetName(t *testing.T) {
	node := &Node{name: "node", depth: 0, children: []*Node{}, parent: nil}
	opts := Options{trailingSlash: true}

	result := getName(node, &opts)
	expected := "node"
	if result != expected {
		t.Errorf("getName() = %q; want %q", result, expected)
	}

	node.children = append(node.children, &Node{name: "child", depth: 1, children: []*Node{}, parent: node})
	result = getName(node, &opts)
	expected = "node/"
	if result != expected {
		t.Errorf("getName() = %q; want %q", result, expected)
	}
}

func TestIsLastChild(t *testing.T) {
	root := &Node{name: ".", depth: 0, children: []*Node{}, parent: nil}
	child1 := &Node{name: "child1", depth: 1, children: []*Node{}, parent: root}
	child2 := &Node{name: "child2", depth: 1, children: []*Node{}, parent: root}
	root.children = append(root.children, child1, child2)

	if isLastChild(child1) {
		t.Errorf("Expected child1 to not be the last child")
	}

	if !isLastChild(child2) {
		t.Errorf("Expected child2 to be the last child")
	}
}

func TestRemovePrefix(t *testing.T) {
	opts := Options{false, "", strings.Builder{}, "utf-8", false, false, true}
	CHILD, LAST_CHILD, DIRECTORY, EMPTY := getPrefixes(&opts)

	tests := []struct {
		str      string
		expected string
	}{
		{CHILD + "node", "node"},
		{LAST_CHILD + "node", "node"},
		{DIRECTORY + "node", "node"},
		{EMPTY + "node", "node"},
	}

	for _, test := range tests {
		result := removePrefix(test.str, &opts)
		if result != test.expected {
			t.Errorf("removePrefix(%q) = %q; want %q", test.str, result, test.expected)
		}
	}
}
func TestGetPrefixes(t *testing.T) {
	tests := []struct {
		charset           string
		expectedChild     string
		expectedLastChild string
		expectedDirectory string
		expectedEmpty     string
	}{
		{"utf-8", UTF8_CHILD, UTF8_LAST_CHILD, UTF8_DIRECTORY, UTF8_EMPTY},
		{"ascii", ASCII_CHILD, ASCII_LAST_CHILD, ASCII_DIRECTORY, ASCII_EMPTY},
	}

	for _, test := range tests {
		opts := &Options{charset: test.charset}
		child, lastChild, directory, empty := getPrefixes(opts)
		if child != test.expectedChild {
			t.Errorf("For charset %s, expected child prefix %s, but got %s", test.charset, test.expectedChild, child)
		}
		if lastChild != test.expectedLastChild {
			t.Errorf("For charset %s, expected last child prefix %s, but got %s", test.charset, test.expectedLastChild, lastChild)
		}
		if directory != test.expectedDirectory {
			t.Errorf("For charset %s, expected directory prefix %s, but got %s", test.charset, test.expectedDirectory, directory)
		}
		if empty != test.expectedEmpty {
			t.Errorf("For charset %s, expected empty prefix %s, but got %s", test.charset, test.expectedEmpty, empty)
		}
	}
}

func TestMultiRoot(t *testing.T) {
	input := "I\n am\n  a\n   superhero!\na\n what?\na\n superhero!\n"
	root := parseInput(input)
	opts := Options{rootDot: true}
	result := describeTree(root, &opts)
	expected := ".\n├── I\n│   └── am\n│       └── a\n│           └── superhero!\n├── a\n│   └── what?\n└── a\n    └── superhero!"
	if result != expected {
		t.Errorf("describeTree() = %q; want %q", result, expected)
	}
}

func TestWinNewLines(t *testing.T) {
	input := "root\r\n    child1\r\n    child2\r\n        grandchild1\r\n"
	root := parseInput(input)
	opts := Options{rootDot: true}
	result := describeTree(root, &opts)
	expected := ".\n└── root\n    ├── child1\n    └── child2\n        └── grandchild1"
	if result != expected {
		t.Errorf("describeTree() = %q; want %q", result, expected)
	}
}
