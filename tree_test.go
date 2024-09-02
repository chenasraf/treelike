package main

import (
	"os"
	"strings"
	"testing"
)

func TestDescribeTree(t *testing.T) {
	input := "root\n    child1\n    child2\n        grandchild1\n"
	opts := DefaultOptions()
	opts.rootDot = true
	root := parseInput(input, opts)
	result := describeTree(root, opts)

	expected := ".\n└── root\n    ├── child1\n    └── child2\n        └── grandchild1"
	if result != expected {
		t.Errorf("describeTree()\n actual = %q\nwant   = %q", result, expected)
	}
}

func TestRemovePrefix(t *testing.T) {
	opts := DefaultOptions()
	CHILD, LAST_CHILD, DIRECTORY, EMPTY := getPrefixes(opts)

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
		result := removePrefix(test.str, opts)
		if result != test.expected {
			t.Errorf("removePrefix(%q)\n actual = %q\nwant   = %q", test.str, result, test.expected)
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
	opts := DefaultOptions()
	opts.rootDot = true
	root := parseInput(input, opts)
	result := describeTree(root, opts)
	expected := ".\n├── I\n│   └── am\n│       └── a\n│           └── superhero!\n├── a\n│   └── what?\n└── a\n    └── superhero!"
	if result != expected {
		t.Errorf("describeTree()\n actual = %q\nwant   = %q", result, expected)
	}
}

func TestWinNewLines(t *testing.T) {
	input := "root\r\n    child1\r\n    child2\r\n        grandchild1\r\n"
	opts := DefaultOptions()
	opts.rootDot = true
	root := parseInput(input, opts)
	result := describeTree(root, opts)
	expected := ".\n└── root\n    ├── child1\n    └── child2\n        └── grandchild1"
	if result != expected {
		t.Errorf("describeTree()\n actual = %q\nwant   = %q", result, expected)
	}
}

func TestSnapshots(t *testing.T) {
	// list all files in test_files/src
	// cwd := os.Getenv("PWD")
	dirPath := strings.Join([]string{"test_files", "src"}, string(os.PathSeparator))

	dir, err := os.Open(dirPath)

	if err != nil {
		t.Errorf("Error opening directory: %v", err)
	}

	entries, err := dir.Readdir(-1)
	if err != nil {
		t.Errorf("Error reading directory: %v", err)
	}

	defer dir.Close()

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		t.Logf("Testing %s", entry.Name())

		tests := make(map[string]*Options)

		tests[""] = DefaultOptions()
		tests["_ascii"] = DefaultOptions()
		tests["_ascii"].charset = "ascii"
		tests["_full_path"] = DefaultOptions()
		tests["_full_path"].fullPath = true
		tests["_root_path"] = DefaultOptions()
		tests["_root_path"].rootPath = "~"
		tests["_root_path"].fullPath = true
		tests["_no_root"] = DefaultOptions()
		tests["_no_root"].rootDot = false
		tests["_trailing_slash"] = DefaultOptions()
		tests["_trailing_slash"].trailingSlash = true

		contents, err := os.ReadFile(dirPath + string(os.PathSeparator) + entry.Name())

		if err != nil {
			t.Errorf("Error reading file: %v", err)
		}

		for suffix, opts := range tests {
			fileName := entry.Name()[:len(entry.Name())-4] + "_snapshot" + suffix + ".txt"
			filePath := strings.Join([]string{"test_files", "snapshots", fileName}, string(os.PathSeparator))

			t.Logf("Testing snapshot %s", fileName)

			want, err := os.ReadFile(filePath)
			if err != nil {
				t.Errorf("Error reading file: %v", err)
				break
			}
			actual := describeTree(parseInput(string(contents), opts), opts)

			if actual+"\n" != string(want) {
				t.Errorf("describeTree()\nactual = %q\nwant   = %q\n file %v", actual, want, filePath)
				break
			}
		}
	}

}
