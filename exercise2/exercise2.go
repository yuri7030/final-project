package main

import (
    "fmt"
	"os"
	"strconv"
    "github.com/diegovanne/go23/exercise2/services"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run exercise2.go <elements>")
		fmt.Println("Example: go run exercise2.go 5 2 8 1 9 3")
		os.Exit(1)
	}

	input := os.Args[1:]
	fmt.Println("Input:", input)

	ints, floats, strings := inputparser.ParseInput(input)

	if len(ints) > 0 {
		sorting.SortInts(ints)
		fmt.Println("Sorted integer array:", ints)
	}

	if len(floats) > 0 {
		sorting.SortFloats(floats)
		fmt.Println("Sorted float64 array:", floats)
	}

	if len(strings) > 0 {
		sorting.SortStrings(strings)
		fmt.Println("Sorted string array:", strings)
	}
}