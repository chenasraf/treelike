package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed version.txt
var VERSION embed.FS

// helpText generates and returns a strings.Builder containing the help text for the program.
// The help text includes usage instructions and descriptions of the available command-line options.
//
// Returns:
//
//	strings.Builder - A builder containing the formatted help text.
func helpText() strings.Builder {
	LE := getLE()
	var builder strings.Builder
	builder.WriteString("Usage: treelike [OPTIONS] [TREE-STRUCTURE]" + LE)
	builder.WriteString("Prints a tree-like representation of the input." + LE)
	builder.WriteString("" + LE)
	builder.WriteString("Options:" + LE)
	builder.WriteString("  -h, --help               Show this help message and exit" + LE)
	builder.WriteString("  -V, --version            Show the version number and exit" + LE)
	builder.WriteString("  -f, --file FILE          Read from FILE" + LE)
	builder.WriteString("   -, --stdin              Read from stdin" + LE)
	builder.WriteString("  -c, --charset CHARSET    Use CHARSET to display characters (utf-8, ascii)" + LE)
	builder.WriteString("  -s, --trailing-slash     Display trailing slash on directory" + LE)
	builder.WriteString("  -p, --full-path          Display full path" + LE)
	builder.WriteString("  -r, --root-path          Use PATH to change the name of the root node (default: .)" + LE)
	builder.WriteString("                           N/A if `--no-root-dot` is enabled" + LE)
	builder.WriteString("  -D, --no-root-dot        Do not display a root element" + LE)
	return builder

}

// getOpts parses command-line arguments and returns an Options struct populated with the parsed values.
// It supports various flags to customize the behavior of the program, such as reading from stdin,
// specifying a file, setting the charset, and more. If an invalid charset is provided, the function
// prints an error message and exits the program.
//
// Supported flags:
//
//	-h, --help            : Display help text and exit.
//	-, --stdin            : Read input from stdin.
//	-f, --file <filename> : Read input from the specified file.
//	-c, --charset <name>  : Set the charset (valid values are "utf-8" and "ascii").
//	-s, --trailing-slash  : Enable trailing slash in output.
//	-p, --full-path       : Enable full path in output.
//	-D, --no-root-dot     : Disable the root dot in output.
//
// Returns:
//
//	Options - A struct containing the parsed options.
func getOpts() *Options {
	opts := DefaultOptions()
	args := os.Args[1:]
	for len(args) > 0 {
		switch args[0] {
		case "-h", "--help":
			{
				help := helpText()
				fmt.Printf("%s", help.String())
				os.Exit(0)
			}
		case "-V", "--version":
			{
				version, err := VERSION.ReadFile("version.txt")
				if err != nil {
					fmt.Println("Error getting version:", err)
					os.Exit(0)
				}
				fmt.Println(string(version))
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
		case "-r", "--root-path":
			{
				opts.rootPath = args[1]
				args = args[2:]
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
