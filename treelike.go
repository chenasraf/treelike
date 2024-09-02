package main

import (
	"fmt"
	"os"
)

func main() {
	opts := getOpts()

	input, err, code := parseRawInput(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(code)
	}
	node := parseInput(input.String(), opts)
	fmt.Println(describeTree(node, opts))
}
