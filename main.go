package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: json-parser [file]")
		os.Exit(1)
	}

	file := strings.Join(os.Args[1:], "")
	fmt.Printf("Parsing file %s\n", file)
}
