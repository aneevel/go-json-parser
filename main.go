package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: json-parser [file]")
		os.Exit(1)
	}

	file := strings.Join(os.Args[1:], "")

	// TODO: Remove when no longer testing
	fmt.Printf("Parsing file: %s\n", file)

	data, err := os.ReadFile(file)
	check(err)

	parser := NewParser(data)

	status := parser.Parse()

	if status == 0 {
		fmt.Println("Valid JSON")
	} else {
		fmt.Println("Invalid JSON")
	}

	fmt.Printf("Parsed output: %v\n", status)
}
