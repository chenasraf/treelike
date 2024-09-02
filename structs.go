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
	rootPath      string
}

// default options factory
func DefaultOptions() *Options {
	return &Options{
		fromStdin:     false,
		fromFile:      "",
		extra:         strings.Builder{},
		charset:       "utf-8",
		trailingSlash: false,
		fullPath:      false,
		rootDot:       true,
		rootPath:      ".",
	}
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
