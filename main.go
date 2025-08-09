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
	fmt.Printf("Parsing file %s\n", file)

	data, err := os.ReadFile(file)
	check(err)

	// TODO: Remove when no longer testing
	_, err = os.Stdout.Write(data)
	check(err)

}
