package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed version.txt
var VERSION embed.FS

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
func getOpts() Options {
	opts := Options{false, "", strings.Builder{}, "utf-8", false, false, true}
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
		default:
			{
				opts.extra.WriteString(args[0] + "\n")
				args = args[1:]
			}
		}
	}
	return opts
}
