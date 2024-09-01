package main

import "strings"

type Options struct {
	fromStdin     bool
	fromFile      string
	extra         strings.Builder
	charset       string
	trailingSlash bool
	fullPath      bool
	rootDot       bool
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
