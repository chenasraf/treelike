package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func helpText() {
	fmt.Println("Usage: treelike [OPTIONS] [PATH]")
}

func main() {
	args := os.Args[1:]
	fromStdin := false
	fromFile := ""
	_ = fromFile

	var input strings.Builder

	for len(args) > 0 {
		switch args[0] {
		case "-h", "--help":
			{
				helpText()
				os.Exit(0)
			}
		case "-", "--stdin":
			{
				fromStdin = true
				args = args[1:]
			}
		case "-f", "--file":
			{
				fromFile = args[1]
				args = args[1:]
			}
		default:
			{
				input.WriteString(args[0] + "\n")
				args = args[1:]
			}
		}
	}

	if fromStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input.WriteString(scanner.Text() + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(1)
		}
	} else if fromFile != "" {
		file, err := os.Open(fromFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", fromFile, err)
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
	} else {
		helpText()
		os.Exit(1)
	}
	fmt.Printf("input:\n%s", input.String())
}
