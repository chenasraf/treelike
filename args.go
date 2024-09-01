package main

import (
	"fmt"
	"os"
	"strings"
)

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
