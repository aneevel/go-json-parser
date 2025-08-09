package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: json-parser [file]")
		os.Exit(1)
	}
	fmt.Println("JSON Parser")
}
